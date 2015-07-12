package eval

import (
	"reflect"
)

func evalCallExpr(call *CallExpr, env Env) ([]reflect.Value, error) {
	if call.IsConst() {
		return []reflect.Value{call.Const()}, nil
	} else if call.isBuiltin {
		return evalCallBuiltinExpr(call, env)
	} else if call.isTypeConversion {
		return evalCallTypeExpr(call, env)
	} else {
		return evalCallFunExpr(call, env)
	}
}

func evalCallTypeExpr(call *CallExpr, env Env) ([]reflect.Value, error) {
	// Arg0 can only be const if it is ConstNil, otherwise the entire expression
	// would be const and evalCallExpr will have already returned.
	arg := call.Args[0]
	t := arg.KnownType()[0]
	if t == ConstNil {
		// This has already been typechecked to be a nil-able type
		return []reflect.Value{hackedNew(call.KnownType()[0]).Elem()}, nil
	} else if v, err := EvalExpr(arg, env); err != nil {
		return nil, nil
	} else {
		cast := v[0].Convert(unhackType(call.KnownType()[0]))
		return []reflect.Value{cast}, nil
	}
}

func evalCallFunExpr(call *CallExpr, env Env) ([]reflect.Value, error) {
	v, err := EvalExpr(call.Fun, env)
	if err != nil {
		return nil, err
	}

	fun := v[0]
	ft := fun.Type()
	numIn := ft.NumIn()

	// Evaluate arguments
	args := make([]reflect.Value, len(call.Args))
	if call.arg0MultiValued {
		if argp, err := EvalExpr(call.Args[0], env); err != nil {
			return nil, err
		} else {
			args = argp
		}
	} else if len(args) != 0 {
		var i int
		for i = 0; i < numIn - 1; i += 1 {
			arg := call.Args[i]
			argType := knownType{ft.In(i)}
			if argV, err := evalTypedExpr(arg, argType, env); err != nil {
				return nil, err
			} else {
				args[i] = argV[0]
			}
		}
		argNT := ft.In(i)
		if ft.IsVariadic() && !call.argNEllipsis {
			argNT = argNT.Elem()
		}
		argNKnownType := knownType{argNT}
		for ; i < len(call.Args); i += 1 {
			arg := call.Args[i]
			if argV, err := evalTypedExpr(arg, argNKnownType, env); err != nil {
				return nil, err
			} else {
				args[i] = argV[0]
			}
		}
	}

	var out []reflect.Value
	if call.argNEllipsis {
		out = fun.CallSlice(args)
	} else {
		out = fun.Call(args)
	}
	return out, nil
}
