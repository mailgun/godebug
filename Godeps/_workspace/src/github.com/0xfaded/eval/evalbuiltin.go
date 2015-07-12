package eval

import (
	"reflect"
)

func evalCallBuiltinExpr(call *CallExpr, env Env) ([]reflect.Value, error) {
	ident := call.Fun.(*Ident)
	switch ident.Name {
	case "complex":
		return evalBuiltinComplexExpr(call, env)
	case "real":
		return evalBuiltinRealExpr(call, env)
	case "imag":
		return evalBuiltinImagExpr(call, env)
	case "new":
		return evalBuiltinNewExpr(call, env)
	case "make":
		return evalBuiltinMakeExpr(call, env)
	case "len":
		return evalBuiltinLenExpr(call, env)
	case "cap":
		return evalBuiltinCapExpr(call, env)
	case "append":
		return evalBuiltinAppendExpr(call, env)
	case "copy":
		return evalBuiltinCopyExpr(call, env)
	case "delete":
		return evalBuiltinDeleteExpr(call, env)
	case "panic":
		return evalBuiltinPanicExpr(call, env)
	default:
		panic("eval: unimplemented builtin " + ident.Name)
	}
}

func evalBuiltinComplexExpr(call *CallExpr, env Env) ([]reflect.Value, error) {
	var err error

	resT := call.KnownType()[0]
	argT := knownType{comprisingFloatType(resT)}

	var re, im []reflect.Value
	if re, err = evalTypedExpr(call.Args[0], argT, env); err != nil {
		return nil, err
	} else if im, err = evalTypedExpr(call.Args[1], argT, env); err != nil {
		return nil, err
	}
	cplx := builtinComplex(re[0], im[0])
	return []reflect.Value{cplx}, nil
}

func evalBuiltinRealExpr(call *CallExpr, env Env) ([]reflect.Value, error) {
	var err error

	resT := call.KnownType()[0]
	argT := knownType{comprisingFloatType(resT)}

	var cplx []reflect.Value
	if cplx, err = evalTypedExpr(call.Args[0], argT, env); err != nil {
		return nil, err
	}
	re := builtinReal(cplx[0])
	return []reflect.Value{re}, nil
}

func evalBuiltinImagExpr(call *CallExpr, env Env) ([]reflect.Value, error) {
	var err error

	resT := call.KnownType()[0]
	argT := knownType{comprisingFloatType(resT)}

	var cplx []reflect.Value
	if cplx, err = evalTypedExpr(call.Args[0], argT, env); err != nil {
		return nil, err
	}
	im := builtinImag(cplx[0])
	return []reflect.Value{im}, nil
}

func evalBuiltinNewExpr(call *CallExpr, env Env) ([]reflect.Value, error) {
	resT := call.KnownType()[0]
	ptr := builtinNew(resT.Elem())
	return []reflect.Value{ptr}, nil
}

func evalBuiltinMakeExpr(call *CallExpr, env Env) ([]reflect.Value, error) {
	resT := call.KnownType()[0]
	length, capacity := 0, 0
	var err error
	if len(call.Args) > 1 {
		if length, err = evalInteger(call.Args[1], env); err != nil {
			return nil, err
		}
	}
	if len(call.Args) > 2 {
		if capacity, err = evalInteger(call.Args[2], env); err != nil {
			return nil, err
		}
	}
	var res reflect.Value
	switch resT.Kind() {
	case reflect.Slice:
		res = reflect.MakeSlice(resT, length, capacity)
	case reflect.Map:
		res = reflect.MakeMap(resT)
	case reflect.Chan:
		res = reflect.MakeChan(resT, length)
	default:
		panic(dytc("make(bad type)"))
	}
	return []reflect.Value{res}, nil
}

func evalBuiltinLenExpr(call *CallExpr, env Env) ([]reflect.Value, error) {
	if arg0, err := EvalExpr(call.Args[0], env); err != nil {
		return nil, err
	} else {
		length := builtinLen(arg0[0])
		return []reflect.Value{length}, nil
	}
}

func evalBuiltinCapExpr(call *CallExpr, env Env) ([]reflect.Value, error) {
	if arg0, err := EvalExpr(call.Args[0], env); err != nil {
		return nil, err
	} else {
		capacity := builtinCap(arg0[0])
		return []reflect.Value{capacity}, nil
	}
}

func evalBuiltinAppendExpr(call *CallExpr, env Env) ([]reflect.Value, error) {
	sliceT := call.KnownType()
	head, err := evalTypedExpr(call.Args[0], sliceT, env)
	if err != nil {
		return nil, err
	}
	var tail reflect.Value
	if call.argNEllipsis {
		xs, err := EvalExpr(call.Args[1], env)
		if err != nil {
			return nil, err
		}
		tail = xs[0]
	} else {
		numXs := len(call.Args) - 1
		tail = reflect.MakeSlice(sliceT[0], numXs, numXs)
		xT := knownType{sliceT[0].Elem()}
		for i := 1; i < len(call.Args); i += 1 {
			if x, err := evalTypedExpr(call.Args[i], xT, env); err != nil {
				return nil, err
			} else {
				tail.Index(i-1).Set(x[0])
			}
		}
	}

	res := builtinAppend(head[0], tail)
	return []reflect.Value{res}, nil
}

func evalBuiltinCopyExpr(call *CallExpr, env Env) ([]reflect.Value, error) {
	if x, err := EvalExpr(call.Args[0], env); err != nil {
		return nil, err
	} else if y, err := EvalExpr(call.Args[1], env); err != nil {
		return nil, err
	} else {
		n := builtinCopy(x[0], y[0])
		return []reflect.Value{n}, nil
	}
}

func evalBuiltinDeleteExpr(call *CallExpr, env Env) ([]reflect.Value, error) {
	m := call.Args[0]
	mT := m.KnownType()[0]
	if x, err := EvalExpr(m, env); err != nil {
		return nil, err
	} else if y, err := evalTypedExpr(call.Args[1], knownType{mT.Key()}, env); err != nil {
		return nil, err
	} else {
		builtinDelete(x[0], y[0])
		return []reflect.Value{}, nil
	}
}

func evalBuiltinPanicExpr(call *CallExpr, env Env) ([]reflect.Value, error) {
	if arg0, err := evalTypedExpr(call.Args[0], knownType{emptyInterface}, env); err != nil {
		return nil, err
	} else {
		err := builtinPanic(arg0[0])
		return []reflect.Value{}, err
	}
}

