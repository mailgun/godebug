package eval

import (
	"reflect"
)

func evalIdent(ident *Ident, env Env) (reflect.Value, error) {
	if ident.IsConst() {
		return ident.Const(), nil
	}

	name := ident.Name
	switch ident.source {
	case envVar:
		for searchEnv := env; searchEnv != nil; searchEnv = searchEnv.PopScope() {
			if v := searchEnv.Var(name); v.IsValid() {
				return v.Elem(), nil
			}
		}
	case envFunc:
		for searchEnv := env; searchEnv != nil; searchEnv = searchEnv.PopScope() {
			if v := searchEnv.Func(name); v.IsValid() {
				return v, nil
			}
		}
	}
        panic(dytc("missing identifier '"+name+"'"))
}

