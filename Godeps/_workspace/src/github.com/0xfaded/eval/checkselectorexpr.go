package eval

import (
	"reflect"
	"go/ast"
)

func checkSelectorExpr(selector *ast.SelectorExpr, env Env) (*SelectorExpr, []error) {
	aexpr := &SelectorExpr{SelectorExpr: selector}

	// First check if this is a package identifier
	if ident, ok := selector.X.(*ast.Ident); ok {
		if pkg := env.Pkg(ident.Name); pkg != nil {
			// Lookup this ident in the context of the package.
			sel, errs := checkIdent(aexpr.SelectorExpr.Sel, pkg)
			if len(errs) == 1 {
				if undefined, ok := (errs[0]).(ErrUndefined); ok {
					undefined.Expr = aexpr
					errs[0] = undefined
				}
			}
			// This selector node is really a single identifier.
			// Convey the type information to the parent.
			aexpr.constValue = sel.constValue
			aexpr.knownType = sel.knownType
			aexpr.pkgName = ident.Name
			aexpr.X = &Ident{Ident: ident}
			aexpr.Sel = sel
			return aexpr, errs
		}
	}

	x, errs := CheckExpr(selector.X, env)
	aexpr.X = x
	aexpr.Sel = &Ident{Ident: selector.Sel}
	if errs != nil && !x.IsConst() {
		return aexpr, errs
	}

	t, err := expectSingleType(x)
	if err != nil {
		return aexpr, append(errs, err)
	} else if t == ConstNil {
		return aexpr, append(errs, ErrUntypedNil{x})
	}

	name := aexpr.Sel.Name
	// Check structs for field selectors
	switch t.Kind() {
	case reflect.Struct:
		if field, ok := t.FieldByName(name); ok {
			aexpr.field = field.Index
			aexpr.knownType = knownType{field.Type}
			return aexpr, errs
		}
	case reflect.Ptr:
		// auto-indirect of *struct types
		// filter out const types and non structs
		if _, ok := t.(ConstType); ok {
			break
		} else if t.Elem().Kind() != reflect.Struct {
			break
		} else if field, ok := t.Elem().FieldByName(name); ok {
			aexpr.field = field.Index
			aexpr.knownType = knownType{field.Type}
			return aexpr, errs
		}
	}

	// Type.Method() on a non interface type returns a receiver type
	// accepting X.KnownType as it's first argument. Value.Method() Binds the
	// method to the value, and hence does not have the first argument.
	if t.Kind() == reflect.Interface {
		if method, ok := t.MethodByName(name); ok {
			aexpr.knownType = knownType{method.Type}
			aexpr.method = method.Index
			return aexpr, errs
		}
	} else {
		for i := 0; i < 2; i += 1 {
			if method, ok := t.MethodByName(name); ok {
				zero := reflect.Zero(t)
				bound := zero.MethodByName(name)
				aexpr.knownType = knownType{bound.Type()}
				aexpr.method = method.Index
				aexpr.isPtrReceiver = i != 0
				return aexpr, errs
			}
			// Check for ptr receivers
			t = reflect.PtrTo(t)
		}
	}

	return aexpr, append(errs, ErrUndefinedFieldOrMethod{aexpr})
}
