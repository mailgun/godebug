package eval

import (
	"reflect"
	"go/ast"
)

func checkCompositeLit(lit *ast.CompositeLit, env Env) (*CompositeLit, []error) {
	return checkCompositeLitR(lit, nil, env)
}

// Recursively check composite literals, where a child composite lit's type depends the
// parent's type For example, the expression [][]int{{1,2},{3,4}} contains two
// slice lits, {1,2} and {3,4}, but their types are inferenced from the parent [][]int{}.
func checkCompositeLitR(lit *ast.CompositeLit, t reflect.Type, env Env) (*CompositeLit, []error) {
	alit := &CompositeLit{CompositeLit: lit}

	// We won't generate any errors here if the given type does not match lit.Type.
	// The caller will need to detect the type incompatibility.
	if lit.Type != nil {
		var errs []error
		lit.Type, t, _, errs = checkType(lit.Type, env)
		if errs != nil {
			return alit, errs
		}
	} else if t == nil {
		return alit, []error{ErrMissingCompositeLitType{alit}}
	}

	alit.knownType = knownType{t}
	if alit.CompositeLit.Elts != nil {
		alit.Elts = make([]Expr, len(alit.CompositeLit.Elts))
	}

	switch t.Kind() {
	case reflect.Map:
		return checkCompositeLitMap(alit, t, env)
	case reflect.Array, reflect.Slice:
		return checkCompositeLitArrayOrSlice(alit, t, env)
	case reflect.Struct:
		return checkCompositeLitStruct(alit, t, env)
	default:
		panic("eval: unimplemented composite lit " + t.Kind().String())
	}
}

func checkCompositeLitMap(lit *CompositeLit, t reflect.Type, env Env) (*CompositeLit, []error) {
	var errs []error
	kT := t.Key()

	// Don't check for duplicate interface{} keys. This is a gc bug
	// http://code.google.com/p/go/issues/detail?id=7214
	var seen map[interface{}] bool
	if kT.Kind() != reflect.Interface {
		seen = make(map[interface{}] bool, len(lit.Elts))
	}
	eltT := t.Elem()

	for i := range lit.Elts {
		if kv, ok := lit.CompositeLit.Elts[i].(*ast.KeyValueExpr); !ok {
			elt, moreErrs := CheckExpr(lit.CompositeLit.Elts[i], env)
			lit.Elts[i] = elt
			if moreErrs != nil {
				errs = append(errs, moreErrs...)
			}
			errs = append(errs, ErrMissingMapKey{elt})
		} else {
			k, ok, moreErrs := checkExprAssignableTo(kv.Key, kT, env)
			if !ok {
				if len(k.KnownType()) != 0 {
					kF := fakeCheckExpr(kv.Key, env)
					kF.setKnownType(knownType(k.KnownType()))
					errs = append(errs, ErrBadMapKey{kF, kT})
				}
			} else {
				errs = append(errs, moreErrs...)
			}
			if seen != nil && k.IsConst() {
				var constKey interface{}
				if k.KnownType()[0] == ConstNil {
					constKey = nil
				} else if cT, ok := k.KnownType()[0].(ConstType); ok {
					c, _ := promoteConstToTyped(cT, constValue(k.Const()),
						cT.DefaultPromotion(), k)
					constKey = reflect.Value(c).Interface()
				} else {
					constKey = k.Const().Interface()
				}
				if seen[constKey] {
					errs = append(errs, ErrDuplicateMapKey{k})
				}
				seen[constKey] = true
			}
			v, moreErrs := checkMapValue(kv.Value, eltT, env)
			if moreErrs != nil {
				errs = append(errs, moreErrs...)
			}
			lit.Elts[i] = &KeyValueExpr{KeyValueExpr: kv, Key: k, Value: v}
		}
	}
	return lit, errs
}

func checkCompositeLitArrayOrSlice(lit *CompositeLit, t reflect.Type, env Env) (*CompositeLit, []error) {
	var errs, moreErrs []error
	eltT := t.Elem()
	maxIndex, curIndex := -1, 0
	outOfBounds := false
	length := -1
	if t.Kind() == reflect.Array {
		length = t.Len()
	}
	used := make(map[int] bool, len(lit.Elts))
	// Check all keys are valid and calculate array or slice length.
	// Elements with key are placed at the keyed position.
	// Elements without are placed in the next position.
	// For example, []int{1, 2:1, 1} -> [1, 0, 1, 1]
	for i := range lit.Elts {
		var value ast.Expr
		var key Expr
		var avalue *Expr
		skipIndexChecks := false
		kv, ok := lit.CompositeLit.Elts[i].(*ast.KeyValueExpr)
		if !ok {
			avalue, value = &lit.Elts[i], lit.CompositeLit.Elts[i]
		} else {
			akv := &KeyValueExpr{KeyValueExpr: kv}
			lit.Elts[i] = akv
			avalue, value = &akv.Value, kv.Value
			// Check the array key
			var index int
			key, index, ok, moreErrs = checkArrayIndex(kv.Key, env);
			if !ok || moreErrs != nil {
				// NOTE[crc] Haven't checked the gc implementation, but
				// from experimentation it seems that only undefined
				// idents are reported. This filter should perhaps be part
				// of checkArrayIndex
				for _, err := range moreErrs {
					if _, ok := err.(ErrUndefined); ok {
						errs = append(errs, err)
					}
				}
				errs = append(errs, ErrBadArrayKey{key})
				// Don't include this element in index calculations
				curIndex -= 1
				skipIndexChecks = true
			} else {
				lit.indices = append(lit.indices, struct{pos, index int}{i, index})
				curIndex = index
			}
		}
		// finally check the value
		v, moreErrs := checkArrayValue(value, eltT, env)
		*avalue = v

		// These errors slide in before the check errors, but we need v
		if !skipIndexChecks {
			if maxIndex < curIndex {
				maxIndex = curIndex
			}
			if !outOfBounds && length != -1 && curIndex >= length {
				outOfBounds = true
				errs = append(errs, ErrArrayKeyOutOfBounds{v, t, curIndex})
			}
			// has this index been used already
			if used[curIndex] {
				errs = append(errs, ErrDuplicateArrayKey{key, curIndex})
			}
			used[curIndex] = true
		}

		// Add the check errors
		if moreErrs != nil {
			errs = append(errs, moreErrs...)
		}

		curIndex += 1
	}
	lit.indices = append(lit.indices, struct{pos, index int}{-1, -1})
	if length == -1 {
		lit.length = maxIndex + 1
	} else {
		lit.length = length
	}
	return lit, errs
}

func checkCompositeLitStruct(lit *CompositeLit, t reflect.Type, env Env) (*CompositeLit, []error) {
	var errs, moreErrs []error

	// X{} is treated as if it has zero KeyValue'd elements, i.e. unspecified
	// elements are set to zero. This is always valid
	if len(lit.Elts) == 0 {
		return lit, nil
	}

	// gc first checks if there are ANY keys present, and then decides how
	// to process the initialisers.
	keysPresent := false
	for _, elt := range lit.CompositeLit.Elts {
		_, ok := elt.(*ast.KeyValueExpr)
		keysPresent = keysPresent || ok
	}

	if keysPresent {
		seen := make(map[string] bool, len(lit.Elts))
		mixed := false
		for i := 0; i < len(lit.Elts); i += 1 {
			kv, ok := lit.CompositeLit.Elts[i].(*ast.KeyValueExpr)
			if !ok {
				elt := fakeCheckExpr(lit.CompositeLit.Elts[i], env)
				lit.Elts[i] = elt
				if !mixed {
					// This error only gets reported once
					mixed = true
					errs = append(errs, ErrMixedStructValues{elt})
				}
				continue
			}

			key := fakeCheckExpr(kv.Key, env)
			var value Expr
			// Check the key is a struct member
			if ident, ok := key.(*Ident); !ok {
				value = fakeCheckExpr(kv.Value, env)
				errs = append(errs, ErrInvalidStructField{key})
			} else {
				name := ident.Name
				if field, ok := t.FieldByName(name); !ok {
					value = fakeCheckExpr(kv.Value, env)
					errs = append(errs, ErrUnknownStructField{key, t, name})
				} else {
					if seen[name] {
						errs = append(errs, ErrDuplicateStructField{ident, name})
					}
					seen[name] = true
					lit.fields = append(lit.fields, field.Index[0])
					value, moreErrs = checkStructField(kv.Value, field, env)
					if moreErrs != nil {
						errs = append(errs, moreErrs...)
					}
				}
			}
			lit.Elts[i] = &KeyValueExpr{KeyValueExpr: kv, Key: key, Value: value}
		}
	} else {
		numFields := t.NumField()
		var i int
		for i = 0; i < numFields && i < len(lit.Elts); i += 1 {
			field := t.Field(i)
			lit.Elts[i], moreErrs = checkStructField(lit.CompositeLit.Elts[i], field, env)
			if moreErrs != nil {
				errs = append(errs, moreErrs...)
			}
			lit.fields = append(lit.fields, i)
		}
		if numFields != len(lit.Elts) {
			errs = append(errs, ErrWrongNumberOfStructValues{lit})
		}
		// Remaining fields are type checked reguardless of use
		for ; i < len(lit.Elts); i += 1 {
			lit.Elts[i], moreErrs = CheckExpr(lit.CompositeLit.Elts[i], env)
			if moreErrs != nil {
				errs = append(errs, moreErrs...)
			}
		}
	}
	return lit, errs
}

func checkMapValue(expr ast.Expr, eltT reflect.Type, env Env) (Expr, []error) {
	switch eltT.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct:
		if lit, ok := expr.(*ast.CompositeLit); ok {
			return checkCompositeLitR(lit, eltT, env)
		}
	}

	aexpr, ok, errs := checkExprAssignableTo(expr, eltT, env)
	if !ok {
		// NOTE[crc] this hack removes conversion errors from consts other
		// than strings and nil to match the output of gc.
		if ccerr, ok := errs[0].(ErrBadConstConversion); ok {
			if ccerr.from == ConstNil {
				// No ErrBadMapValue for nil
				return aexpr, errs
			} else if ccerr.from != ConstString {
				// gc implementation only displays string conversion errors
				errs = nil
			}
		}
		errs = append(errs, ErrBadMapValue{aexpr, eltT})
	}
	return aexpr, errs
}

func checkArrayValue(expr ast.Expr, eltT reflect.Type, env Env) (Expr, []error) {
	switch eltT.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct:
		if lit, ok := expr.(*ast.CompositeLit); ok {
			return checkCompositeLitR(lit, eltT, env)
		}
	}

	aexpr, ok, errs := checkExprAssignableTo(expr, eltT, env)
	if !ok {
		// NOTE[crc] this hack removes conversion errors from consts other
		// than strings and nil to match the output of gc.
		if ccerr, ok := errs[0].(ErrBadConstConversion); ok {
			if ccerr.from == ConstNil {
				// No ErrBadArrayValue for nil
				return aexpr, errs
			} else if ccerr.from != ConstString {
				// gc implementation only displays string conversion errors
				errs = nil
			}
		}
		errs = append(errs, ErrBadArrayValue{aexpr, eltT})
	}
	return aexpr, errs
}

func checkStructField(expr ast.Expr, field reflect.StructField, env Env) (Expr, []error) {
	aexpr, ok, errs := checkExprAssignableTo(expr, field.Type, env)
	if !ok {
		errs = append([]error{}, ErrBadStructValue{aexpr, field.Type})
	}
	return aexpr, errs
}
