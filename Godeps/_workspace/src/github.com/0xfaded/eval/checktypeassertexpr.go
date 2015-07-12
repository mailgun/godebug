package eval

import (
	"reflect"

	"go/ast"
)

func checkTypeAssertExpr(assert *ast.TypeAssertExpr, env Env) (*TypeAssertExpr, []error) {
	aexpr := &TypeAssertExpr{TypeAssertExpr: assert}
	x, errs := CheckExpr(assert.X, env)
	aexpr.X = x

	if errs != nil && !x.IsConst() {
		return aexpr, errs
	} else if xT, err := expectSingleType(x); err != nil {
		errs = append(errs, err)
	} else if xT == ConstNil {
		errs = append(errs, ErrUntypedNil{x})
	} else if xT.Kind() != reflect.Interface {
		errs = append(errs, ErrInvalidTypeAssert{aexpr})
	} else {
		typ, t, _, moreErrs := checkType(assert.Type, env)
		aexpr.Type = typ
		errs = append(errs, moreErrs...)
		if t != nil {
			aexpr.knownType = knownType{t}
			if t.Kind() != reflect.Interface && !unhackType(t).Implements(xT) {
				errs = append(errs, ErrImpossibleTypeAssert{aexpr})
			}
		}
	}
	return aexpr, errs
}
