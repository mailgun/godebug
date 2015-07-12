package eval

import (
	"reflect"
)

func evalSelectorExpr(selector *SelectorExpr, env Env) (reflect.Value, error) {

	if selector.pkgName != "" {
		vs, err := evalIdent(selector.Sel, env.Pkg(selector.pkgName))
		return vs, err
	}

	vs, err := EvalExpr(selector.X, env)
	if err != nil {
		return reflect.Value{}, err
	}
	v := vs[0]
	t := v.Type()
	if selector.field != nil {
		if t.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		return v.FieldByIndex(selector.field), nil
	}

	if selector.isPtrReceiver {
		v = v.Addr()
	}
	return v.Method(selector.method), nil
}

