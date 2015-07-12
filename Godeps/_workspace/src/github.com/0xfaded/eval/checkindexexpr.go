package eval

import (
	"reflect"
	"go/ast"
)

func checkIndexExpr(index *ast.IndexExpr, env Env) (*IndexExpr, []error) {
	aexpr := &IndexExpr{IndexExpr: index}
	x, errs := CheckExpr(index.X, env)
	aexpr.X = x
	if errs != nil && !x.IsConst() {
		return aexpr, errs
	}

	t, err := expectSingleType(x)
	if err != nil {
		return aexpr, append(errs, err)
	}

	// index of array pointer is short hand for dereference and then index
	if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Array {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.Map:
		aexpr.knownType = knownType{t.Elem()}
		i, ok, moreErrs := checkExprAssignableTo(index.Index, t.Key(), env)
		aexpr.Index = i
		if moreErrs != nil {
			errs = append(errs, moreErrs...)
		}
		if !ok {
			errs = append(errs, ErrBadMapIndex{i, t.Key()})
		}
		return aexpr, errs
	case reflect.String:
		aexpr.knownType = knownType{u8}
		i, moreErrs := checkIndexVectorExpr(x, index.Index, env)
		if moreErrs != nil {
			errs = append(errs, moreErrs...)
		}
		aexpr.Index = i
		return aexpr, errs
	case reflect.Array, reflect.Slice:
		aexpr.knownType = knownType{t.Elem()}
		i, moreErrs := checkIndexVectorExpr(x, index.Index, env)
		if moreErrs != nil {
			errs = append(errs, moreErrs...)
		}
		aexpr.Index = i
		return aexpr, errs
	default:
		aexpr.Index = fakeCheckExpr(index.Index, env)
		return aexpr, append(errs, ErrInvalidIndexOperation{aexpr})
	}
}

func checkIndexVectorExpr(x Expr, index ast.Expr, env Env) (Expr, []error) {
	t := x.KnownType()[0]
	i, iint, ok, errs := checkInteger(index, env)
	if errs != nil && !i.IsConst() {
		// Type check of index failed
	} else if !ok {
		// Type check of index passed but this node is not an integer
		printableIndex := fakeCheckExpr(index, env)
		printableIndex.setKnownType(i.KnownType())
		errs = append(errs, ErrNonIntegerIndex{printableIndex})
	} else if i.IsConst() {
		// If we know the index at compile time, we must assert it is in bounds.
		if iint < 0 {
			errs = append(errs, ErrIndexOutOfBounds{i, x, iint})
		} else if t.Kind() == reflect.Array {
			if iint >= t.Len() {
				errs = append(errs, ErrIndexOutOfBounds{i, x, iint})
			}
		} else if t.Kind() == reflect.String && x.IsConst() {
			str := x.Const()
			if iint >= str.Len() {
				errs = append(errs, ErrIndexOutOfBounds{i, x, iint})
			}
		}
	}
	return i, errs
}
