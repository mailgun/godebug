package eval

import (
	"reflect"
	"go/ast"
)

func checkIdent(ident *ast.Ident, env Env) (_ *Ident, errs []error) {
	aexpr := &Ident{Ident: ident}
	switch aexpr.Name {
	case "nil":
		aexpr.constValue = constValueOf(UntypedNil{})
		aexpr.knownType = []reflect.Type{ConstNil}

	case "true":
		aexpr.constValue = constValueOf(true)
		aexpr.knownType = []reflect.Type{ConstBool}

	case "false":
		aexpr.constValue = constValueOf(false)
		aexpr.knownType = []reflect.Type{ConstBool}
        default:
		for searchEnv := env; searchEnv != nil; searchEnv = searchEnv.PopScope() {
			if v := searchEnv.Var(aexpr.Name); v.IsValid() {
				aexpr.knownType = knownType{v.Elem().Type()}
				aexpr.source = envVar
				return aexpr, errs
			} else if v := searchEnv.Func(aexpr.Name); v.IsValid() {
				aexpr.knownType = knownType{v.Type()}
				aexpr.source = envFunc
				return aexpr, errs
			} else if v := searchEnv.Const(aexpr.Name); v.IsValid() {
				if n, ok := v.Interface().(*ConstNumber); ok {
					aexpr.knownType = knownType{n.Type}
				} else {
					aexpr.knownType = knownType{v.Type()}
				}
				aexpr.constValue = constValue(v)
				aexpr.source = envConst
				return aexpr, errs
			}
		}
		return aexpr, append(errs, ErrUndefined{aexpr})
        }
	return aexpr, errs
}
