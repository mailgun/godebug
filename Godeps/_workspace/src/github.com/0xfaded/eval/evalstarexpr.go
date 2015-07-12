package eval

import (
	"reflect"
)

func evalStarExpr(starExpr *StarExpr, env Env) (reflect.Value, error) {
	if vs, err := EvalExpr(starExpr.X, env); err != nil {
		return reflect.Value{}, err
	} else {
		v := vs[0]
		if v.IsNil() {
			return reflect.Value{}, PanicInvalidDereference{}
		}
		return v.Elem(), nil
	}
}
