package eval

import (
	"reflect"

	"go/ast"
)

func checkStarExpr(star *ast.StarExpr, env Env) (*StarExpr, []error) {
	aexpr := &StarExpr{StarExpr: star}
	x, errs := CheckExpr(star.X, env)

	if errs != nil && !x.IsConst() {
		return aexpr, errs
	} else if t, err := expectSingleType(x); err != nil {
		errs = append(errs, err)
	} else if t == ConstNil {
		errs = append(errs, ErrInvalidIndirect{x})
	} else if t.Kind() != reflect.Ptr {
		printableX := fakeCheckExpr(star.X, env)
		printableX.setKnownType(x.KnownType())
		errs = append(errs, ErrInvalidIndirect{printableX})
	} else {
		aexpr.knownType = knownType{t.Elem()}
	}
	aexpr.X = x
	return aexpr, errs
}
