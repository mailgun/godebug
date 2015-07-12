package eval

import (
	"reflect"

	"go/ast"
	"go/token"
)

func checkCallBuiltinExpr(call *CallExpr, env Env) (*CallExpr, []error, bool) {
	var errs []error
	ident, ok := call.CallExpr.Fun.(*ast.Ident)
	if !ok {
		return call, nil, false
	}
	switch ident.Name {
	case "complex":
		call, errs = checkBuiltinComplex(call, env)
	case "real":
		call, errs = checkBuiltinRealImag(call, env, true)
	case "imag":
		call, errs = checkBuiltinRealImag(call, env, false)
	case "new":
		call, errs = checkBuiltinNew(call, env)
	case "make":
		call, errs = checkBuiltinMake(call, env)
	case "len":
		call, errs = checkBuiltinLenCap(call, env, true)
	case "cap":
		call, errs = checkBuiltinLenCap(call, env, false)
	case "append":
		call, errs = checkBuiltinAppend(call, env)
	case "copy":
		call, errs = checkBuiltinCopyExpr(call, env)
	case "delete":
		call, errs = checkBuiltinDeleteExpr(call, env)
	case "panic":
		call, errs = checkBuiltinPanicExpr(call, env)
	default:
		return call, nil, false
	}
	call.Fun = &Ident{Ident: ident}
	call.isBuiltin = true
	return call, errs, true
}

func checkBuiltinComplex(call *CallExpr, env Env) (*CallExpr, []error) {
	var errs []error
	if call.argNEllipsis = call.Ellipsis != token.NoPos; call.argNEllipsis {
		errs = append(errs, ErrBuiltinInvalidEllipsis{call})
	}
	if len(call.CallExpr.Args) != 2 {
		fakeCheckRemainingArgs(call, 0, env)
		return call, append(errs, ErrBuiltinWrongNumberOfArgs{call})
	}
	x, y, ok, moreErrs := checkBinaryOperands(call.CallExpr.Args[0], call.CallExpr.Args[1], env)
	call.Args = []Expr{x, y}
	if moreErrs != nil {
		errs = append(errs, moreErrs...)
	}
	if !ok {
		return call, errs
	}
	xt, yt := x.KnownType()[0], y.KnownType()[0]
	xct, xctok := xt.(ConstType)
	yct, yctok := yt.(ConstType)
	if xctok && yctok {
		if xct.IsNumeric() && yct.IsNumeric() {
			call.knownType = knownType{c128}
			xc, xerrs := promoteConstToTyped(xct, constValue(x.Const()), f64, x)
			if xerrs != nil {
				errs = append(errs, xerrs...)
			}
			yc, yerrs := promoteConstToTyped(yct, constValue(y.Const()), f64, y)
			if yerrs != nil {
				errs = append(errs, yerrs...)
			}
			if reflect.Value(xc).IsValid() && reflect.Value(yc).IsValid() {
				xf := float64(reflect.Value(xc).Float())
				yf := float64(reflect.Value(yc).Float())
				call.constValue = constValueOf(complex(xf, yf))
				return call, errs
			}
		}
	} else if xctok {
		if attemptBinaryOpConversion(yt) {
			xc, xerrs := promoteConstToTyped(xct, constValue(x.Const()), yt, x)
			if xerrs != nil {
				errs = append(errs, xerrs...)
				if xt == ConstNil {
					// No MismatchedTypes error for nils
					return call, errs
				}
			}
			xv := reflect.Value(xc)
			if xv.IsValid() {
				if yt.Kind() == reflect.Float32 {
					call.knownType = knownType{c64}
					if y.IsConst() {
						xf := float32(xv.Float())
						yf := float32(y.Const().Float())
						call.constValue = constValueOf(complex(xf, yf))
					}
					return call, errs
				} else if yt.Kind() == reflect.Float64 {
					call.knownType = knownType{c128}
					if y.IsConst() {
						xf := float64(xv.Float())
						yf := float64(y.Const().Float())
						call.constValue = constValueOf(complex(xf, yf))
					}
					return call, errs
				}
			}
		} else {
			if xt == ConstNil && isNillable(yt) {
				errs = append(errs, ErrBuiltinWrongArgType{y, call})
				return call, errs
			}
		}
	} else if yctok {
		if attemptBinaryOpConversion(xt) {
			yc, yerrs := promoteConstToTyped(yct, constValue(y.Const()), xt, y)
			if yerrs != nil {
				errs = append(errs, yerrs...)
				if yt == ConstNil {
					// No MismatchedTypes error for nils
					return call, errs
				}
			} else if yt == ConstNil {
				errs = append(errs, ErrBuiltinWrongArgType{x, call})
				return call, errs
			}
			yv := reflect.Value(yc)
			if yv.IsValid() {
				if xt.Kind() == reflect.Float32 {
					call.knownType = knownType{c64}
					if x.IsConst() {
						xf := float32(x.Const().Float())
						yf := float32(yv.Float())
						call.constValue = constValueOf(complex(xf, yf))
					}
					return call, errs
				} else if xt.Kind() == reflect.Float64 {
					call.knownType = knownType{c128}
					if x.IsConst() {
						xf := float64(x.Const().Float())
						yf := float64(yv.Float())
						call.constValue = constValueOf(complex(xf, yf))
					}
					return call, errs
				}
			}
		} else {
			if yt == ConstNil && isNillable(xt) {
				errs = append(errs, ErrBuiltinWrongArgType{x, call})
				return call, errs
			}
		}
	} else if xt == yt {
		if xt.Kind() == reflect.Float32 {
			call.knownType = knownType{c64}
			if x.IsConst() && y.IsConst() {
				xf := float32(x.Const().Float())
				yf := float32(y.Const().Float())
				call.constValue = constValueOf(complex(xf, yf))
			}
			return call, errs
		} else if xt.Kind() == reflect.Float64 {
			call.knownType = knownType{c128}
			if x.IsConst() && y.IsConst() {
				xf := float64(x.Const().Float())
				yf := float64(y.Const().Float())
				call.constValue = constValueOf(complex(xf, yf))
			}
			return call, errs
		}
	}
	if unhackType(xt) == unhackType(yt) {
		errs = append(errs, ErrBuiltinWrongArgType{x, call})
	} else {
		errs = append(errs, ErrBuiltinMismatchedArgs{call, xt, yt})
	}
	return call, errs
}

func checkBuiltinRealImag(call *CallExpr, env Env, isReal bool) (*CallExpr, []error) {
	var errs []error
	if call.argNEllipsis = call.Ellipsis != token.NoPos; call.argNEllipsis {
		errs = append(errs, ErrBuiltinInvalidEllipsis{call})
	}
	if len(call.CallExpr.Args) != 1 {
		fakeCheckRemainingArgs(call, 0, env)
		return call, append(errs, ErrBuiltinWrongNumberOfArgs{call})
	}
	x, moreErrs := CheckExpr(call.CallExpr.Args[0], env)
	if moreErrs != nil {
		errs = append(errs, moreErrs...)
	}
	call.Args = []Expr{x}
	if moreErrs != nil && !x.IsConst() {
		return call, errs
	}
	xt, err := expectSingleType(x)
	if err != nil {
		return call, append(errs, err)
	}
	if xt == ConstNil {
		return call, append(errs, ErrUntypedNil{x})
	}

	if ct, ok := xt.(ConstType); ok {
		if ct == ConstComplex {
			xc, moreErrs := promoteConstToTyped(ct, constValue(x.Const()), c128, x)
			if moreErrs != nil {
				errs = append(errs, moreErrs...)
			}
			xv := reflect.Value(xc)
			if xv.IsValid() {
				call.knownType = knownType{f64}
				c := complex128(reflect.Value(xc).Complex())
				if isReal {
					call.constValue = constValueOf(real(c))
				} else {
					call.constValue = constValueOf(imag(c))
				}
				return call, errs
			}
		}
	} else if xt.Kind() == reflect.Complex128 {
		call.knownType = knownType{f64}
		if x.IsConst() {
			c := complex128(x.Const().Complex())
			if isReal {
				call.constValue = constValueOf(real(c))
			} else {
				call.constValue = constValueOf(imag(c))
			}
		}
		return call, errs
	} else if xt.Kind() == reflect.Complex64 {
		call.knownType = knownType{f32}
		if x.IsConst() {
			c := complex64(x.Const().Complex())
			if isReal {
				call.constValue = constValueOf(real(c))
			} else {
				call.constValue = constValueOf(imag(c))
			}
		}
		return call, errs
	}
	errs = append(errs, ErrBuiltinWrongArgType{x, call})
	return call, errs
}

func checkBuiltinNew(call *CallExpr, env Env) (*CallExpr, []error) {
	var errs []error
	if call.argNEllipsis = call.Ellipsis != token.NoPos; call.argNEllipsis {
		errs = append(errs, ErrBuiltinInvalidEllipsis{call})
	}
	if len(call.CallExpr.Args) == 0 {
		return call, append(errs, ErrBuiltinWrongNumberOfArgs{call})
	}
	x, of, isType, moreErrs := checkType(call.CallExpr.Args[0], env)
	if !isType {
		x, moreErrs = CheckExpr(call.CallExpr.Args[0], env)
		if moreErrs != nil {
			errs = append(errs, moreErrs...)
		}
		call.Args = []Expr{x}
		fakeCheckRemainingArgs(call, 1, env)
		if moreErrs == nil {
			errs = append(errs, ErrBuiltinNonTypeArg{x})
		}
		return call, errs
	} else if len(call.CallExpr.Args) != 1 {
		fakeCheckRemainingArgs(call, 0, env)
		return call, append(errs, ErrBuiltinWrongNumberOfArgs{call})
	} else if moreErrs != nil {
		return call, append(errs, moreErrs...)
	} else {
		call.Args = []Expr{x}
		call.knownType = knownType{reflect.PtrTo(of)}
		return call, nil
	}
}

func checkBuiltinMake(call *CallExpr, env Env) (*CallExpr, []error) {
	if len(call.CallExpr.Args) == 0 {
		return call, []error{ErrBuiltinWrongNumberOfArgs{call}}
	}
	x, of, isType, errs := checkType(call.CallExpr.Args[0], env)
	if !isType {
		fakeCheckRemainingArgs(call, 0, env)
		return call, []error{ErrBuiltinNonTypeArg{call.Args[0]}}
	}
	call.knownType = knownType{of}
	call.Args = make([]Expr, len(call.CallExpr.Args))
	call.Args[0] = x
	if errs != nil {
		fakeCheckRemainingArgs(call, 1, env)
		return call, errs
	}
	var narg int
	skipOrdering := false
	switch of.Kind() {
	case reflect.Slice:
		if len(call.CallExpr.Args) == 1 {
			errs = append(errs, ErrBuiltinWrongNumberOfArgs{call})
		}
		narg = 3
	case reflect.Map, reflect.Chan:
		skipOrdering = true
		narg = 2
	default:
		return call, append(errs, ErrMakeBadType{x, of})
	}
	var args [3]int
	for i := 1; i < narg && i < len(call.CallExpr.Args); i += 1 {
		arg, iint, ok, moreErrs := checkInteger(call.CallExpr.Args[i], env)
		call.Args[i] = arg
		args[i] = iint
		if !ok {
			skipOrdering = true
			errs = append(errs, ErrMakeNonIntegerArg{arg, i})
		} else if moreErrs != nil {
			// Type check passed but is non integral
			errs = append(errs, moreErrs...)
		}
	}
	if len(call.CallExpr.Args) > narg {
		fakeCheckRemainingArgs(call, narg, env)
		errs = append(errs, ErrBuiltinWrongNumberOfArgs{call})
	} else if !skipOrdering{
		if args[1] > args[2] {
			errs = append(errs, ErrMakeLenGtrThanCap{call, args[1], args[2]})
		}
	}
	return call, errs
}

type callRecvWalker bool
func (found *callRecvWalker) visit(expr Expr) bool {
	if *found {
		return false
	}
	if call, ok := expr.(*CallExpr); ok && !call.isTypeConversion {
		*found = true
		return false
	}
	if unary, ok := expr.(*UnaryExpr); ok && unary.Op == token.ARROW {
		*found = true
		return false
	}
	return true
}

func checkBuiltinLenCap(call *CallExpr, env Env, isLen bool) (*CallExpr, []error) {
	call.knownType = knownType{intType}
	var errs []error
	if call.argNEllipsis = call.Ellipsis != token.NoPos; call.argNEllipsis {
		errs = append(errs, ErrBuiltinInvalidEllipsis{call})
	}
	if len(call.CallExpr.Args) != 1 {
		fakeCheckRemainingArgs(call, 0, env)
		return call, append(errs, ErrBuiltinWrongNumberOfArgs{call})
	}

	x, moreErrs := CheckExpr(call.CallExpr.Args[0], env)
	if moreErrs != nil {
		errs = append(errs, moreErrs...)
	}
	call.Args = []Expr{x}
	if errs != nil && !x.IsConst() {
		return call, errs
	}
	xt, err := expectSingleType(x)
	if err != nil {
		return call, append(errs, err)
	}
	if xt == ConstNil {
		return call, append(errs, ErrUntypedNil{x})
	}
	switch xt.Kind() {
	case reflect.Chan, reflect.Slice: // do nothing
	case reflect.Map:
		if !isLen {
			errs = append(errs, ErrBuiltinWrongArgType{x, call})
		}
	case reflect.Ptr:
		xt := xt.Elem()
		if xt.Kind() != reflect.Array {
			break
		}
		fallthrough
	case reflect.Array:
		w := new(callRecvWalker)
		walk(x, w)
		if !*w {
			call.constValue = constValueOf(xt.Len())
		}
	case reflect.String:
		if !isLen {
			errs = append(errs, ErrBuiltinWrongArgType{x, call})
		} else if x.IsConst() {
			call.constValue = constValueOf(x.Const().Len())
		}
	default:
		errs = append(errs, ErrBuiltinWrongArgType{x, call})
	}
	return call, errs
}

func checkBuiltinAppend(call *CallExpr, env Env) (*CallExpr, []error) {
	if len(call.CallExpr.Args) < 1 {
		fakeCheckRemainingArgs(call, 0, env)
		return call, []error{ErrBuiltinWrongNumberOfArgs{call}}
	}
	call.Args = make([]Expr, len(call.CallExpr.Args))
	slice, errs := CheckExpr(call.CallExpr.Args[0], env)
	call.Args[0] = slice
	var sliceT reflect.Type
	var isSlice bool
	if errs == nil || slice.IsConst() {
		var err error
		sliceT, err = expectSingleType(slice)
		if err != nil {
			fakeCheckRemainingArgs(call, 1, env)
			return call, append(errs, err)
		}
		if sliceT != ConstNil {
			isSlice = sliceT.Kind() == reflect.Slice
			call.knownType = knownType{sliceT}
		}
	}
	if call.Ellipsis != token.NoPos {
		call.argNEllipsis = true
		if len(call.CallExpr.Args) == 1 {
			return call, append(errs, ErrAppendFirstArgNotVariadic{slice})
		} else if len(call.CallExpr.Args) != 2 {
			fakeCheckRemainingArgs(call, 1, env)
			return call, append(errs, ErrBuiltinWrongNumberOfArgs{call})
		} else {
			arg1, moreErrs := CheckExpr(call.CallExpr.Args[1], env)
			call.Args[1] = arg1
			if moreErrs != nil && !slice.IsConst() {
				return call, append(errs, moreErrs...)
			}
			arg1T, err := expectSingleType(arg1)
			if err != nil {
				errs = append(errs, err)
			} else if isSlice {
				if arg1T != sliceT && !(sliceT == byteSlice && arg1T.Kind() == reflect.String) {
					errs = append(errs, ErrBuiltinWrongArgType{arg1, call})
				}
			} else if sliceT != nil {
				errs = append(errs, ErrAppendFirstArgNotSlice{slice})
			}
		}
	} else {
		skipTypeCheck := make([]bool, len(call.Args))
		for i := 1; i < len(call.CallExpr.Args); i += 1 {
			argI, moreErrs := CheckExpr(call.CallExpr.Args[i], env)
			call.Args[i] = argI
			if moreErrs != nil {
				errs = append(errs, moreErrs...)
			}
			if moreErrs == nil || argI.IsConst() {
				if _, err := expectSingleType(argI); err != nil {
					skipTypeCheck[i] = true
					errs = append(errs, err)
				}
			} else {
				skipTypeCheck[i] = true
			}
		}
		if isSlice {
			eltT := sliceT.Elem()
			for i := 1; i < len(call.CallExpr.Args); i += 1 {
				if skipTypeCheck[i] {
					continue
				}
				argI := call.Args[i]
				ok := false
				if argI.IsConst() {
					var ct ConstType
					if ct, ok = argI.KnownType()[0].(ConstType); ok {
						x, moreErrs := promoteConstToTyped(ct, constValue(argI.Const()), eltT, argI)
						if !reflect.Value(x).IsValid() {
							errs = append(errs, ErrBuiltinWrongArgType{argI, call})
						} else if moreErrs != nil {
							errs = append(errs, moreErrs...)
						}
					}
				}
				if !ok && unhackType(argI.KnownType()[0]) != unhackType(eltT) {
					errs = append(errs, ErrBuiltinWrongArgType{argI, call})
				}
			}
		} else if sliceT != nil {
			errs = append(errs, ErrAppendFirstArgNotSlice{slice})
		}
	}
	return call, errs
}

func checkBuiltinCopyExpr(call *CallExpr, env Env) (*CallExpr, []error) {
	call.knownType = knownType{intType}
	var errs []error
	if call.argNEllipsis = call.Ellipsis != token.NoPos; call.argNEllipsis {
		errs = append(errs, ErrBuiltinInvalidEllipsis{call})
	}
	if len(call.CallExpr.Args) != 2 {
		fakeCheckRemainingArgs(call, 0, env)
		return call, append(errs, ErrBuiltinWrongNumberOfArgs{call})
	}

	var err error
	var xt, yt reflect.Type
	x, xErrs := CheckExpr(call.CallExpr.Args[0], env)
	if xErrs != nil {
		errs = append(errs, xErrs...)
	}
	if xErrs == nil || x.IsConst() {
		xt, err = expectSingleType(x)
		if err != nil {
			errs = append(errs, err)
		}
	}
	y, yErrs := CheckExpr(call.CallExpr.Args[1], env)
	if yErrs != nil {
		errs = append(errs, yErrs...)
	}
	call.Args = []Expr{x, y}
	if yErrs == nil || y.IsConst() {
		yt, err = expectSingleType(y)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if xt != nil && yt != nil {
		var xk, yk reflect.Kind
		if xt == ConstNil {
			errs = append(errs, ErrUntypedNil{x})
		} else {
			xk = xt.Kind()
		}
		if yt == ConstNil {
			errs = append(errs, ErrUntypedNil{y})
		} else {
			yk = yt.Kind()
		}
		if xk != reflect.Slice || yk != reflect.Slice && yk != reflect.String {
			errs = append(errs, ErrCopyArgsMustBeSlices{call, xt, yt})
		} else if yt.Kind() == reflect.String {
			if xt != byteSlice {
				errs = append(errs, ErrCopyArgsHaveDifferentEltTypes{call, xt, yt})
			}
		} else if unhackType(xt.Elem()) != unhackType(yt.Elem()) {
			errs = append(errs, ErrCopyArgsHaveDifferentEltTypes{call, xt, yt})
		}
	}
	return call, errs
}

func checkBuiltinDeleteExpr(call *CallExpr, env Env) (*CallExpr, []error) {
	call.knownType = knownType{intType}
	var errs []error
	if call.argNEllipsis = call.Ellipsis != token.NoPos; call.argNEllipsis {
		errs = append(errs, ErrBuiltinInvalidEllipsis{call})
	}
	if len(call.CallExpr.Args) != 2 {
		fakeCheckRemainingArgs(call, 0, env)
		return call, append(errs, ErrBuiltinWrongNumberOfArgs{call})
	}
	var mapT, keyT reflect.Type
	m, moreErrs := CheckExpr(call.CallExpr.Args[0], env)
	if moreErrs != nil {
		errs = append(errs, moreErrs...)
	}
	if moreErrs == nil || m.IsConst() {
		var err error
		mapT, err = expectSingleType(m)
		if err != nil {
			errs = append(errs, moreErrs...)
		}
	}

	key, moreErrs := CheckExpr(call.CallExpr.Args[1], env)
	if moreErrs != nil {
		errs = append(errs, moreErrs...)
	}
	call.Args = []Expr{m, key}
	if moreErrs == nil || key.IsConst() {
		var err error
		keyT, err = expectSingleType(key)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if mapT != nil {
		if mapT == ConstNil || mapT.Kind() != reflect.Map {
			errs = append(errs, ErrDeleteFirstArgNotMap{m})
		} else if keyT != nil {
			ok, convErrs := exprAssignableTo(key, mapT.Elem())
			if !ok {
				errs = append(errs, ErrBuiltinWrongArgType{key, call})
			} else if convErrs != nil {
				errs = append(errs, convErrs...)
			}
		}
	}
	return call, errs
}

func checkBuiltinPanicExpr(call *CallExpr, env Env) (*CallExpr, []error) {
	var errs []error
	if call.argNEllipsis = call.Ellipsis != token.NoPos; call.argNEllipsis {
		errs = append(errs, ErrBuiltinInvalidEllipsis{call})
	}
	if len(call.CallExpr.Args) != 1 {
		fakeCheckRemainingArgs(call, 0, env)
		return call, append(errs, ErrBuiltinWrongNumberOfArgs{call})
	}
	x, moreErrs := CheckExpr(call.CallExpr.Args[0], env)
	if moreErrs != nil {
		errs = append(errs, moreErrs...)
	}
	call.Args = []Expr{x}
	if moreErrs != nil && !x.IsConst() {
		return call, errs
	}
	_, err := expectSingleType(x)
	if err != nil {
		return call, append(errs, err)
	}
	return call, errs
}

func fakeCheckRemainingArgs(call *CallExpr, from int, env Env) {
	call.Args = append(call.Args[:from], make([]Expr, len(call.CallExpr.Args)-from)...)
	for i := from; i < len(call.Args); i += 1 {
		call.Args[i] = fakeCheckExpr(call.CallExpr.Args[i], env)
	}
}
