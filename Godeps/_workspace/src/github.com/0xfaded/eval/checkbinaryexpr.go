package eval

import (
	"reflect"

	"go/ast"
	"go/token"
)

// from go.h
const (
	mpscale uint64 = 29
	mpprec = 16
)

func checkBinaryExpr(binary *ast.BinaryExpr, env Env) (*BinaryExpr, []error) {
	aexpr := &BinaryExpr{BinaryExpr: binary}
	x, y, ok, errs := checkBinaryOperands(binary.X, binary.Y, env)
	aexpr.X, aexpr.Y = x, y
	if !ok {
		return aexpr, errs
	}
	xt, yt := x.KnownType()[0], y.KnownType()[0]

	xc, xuntyped := xt.(ConstType)
	yc, yuntyped := yt.(ConstType)
	op := aexpr.Op()
	if op == token.SHL || op == token.SHR {
		var count uint64
		if yuntyped {
			if yt == ConstNil {
				return aexpr, append(errs, ErrBadConstConversion{y, yt, uintType})
			} else {
				c, moreErrs := promoteConstToTyped(yc, constValue(y.Const()), uintType, y)
				yy := reflect.Value(c)
				if yy.IsValid() || yc == ConstString {
					if moreErrs != nil {
						errs = append(errs, moreErrs...)
					}
				}
				if !yy.IsValid() {
					return aexpr, append(errs, ErrInvalidBinaryOperation{aexpr})
				}
				count = yy.Uint()
			}
		} else {
			if !isUnsignedInt(yt) {
				return aexpr, append(errs, ErrInvalidBinaryOperation{aexpr})
			}
			if y.IsConst() {
				count = y.Const().Uint()
			}

		}
		if xuntyped {
			if xn, ok := x.Const().Interface().(*ConstNumber); ok && !xn.Value.IsInteger() {
				var to ConstType
				if xc.IsReal() {
					to = ConstInt
				} else {
					to = ConstFloat
				}
				return aexpr, append(errs, ErrTruncatedConstant{x, to, xn})
			}
			aexpr.knownType = knownType{ConstShiftedInt}
		} else {
			aexpr.knownType = x.KnownType()
		}
		if !isShiftable(xt) {
			return aexpr, append(errs, ErrInvalidBinaryOperation{aexpr})
		}
		if y.IsConst() && x.IsConst() {
			aexpr.knownType = knownType{ConstInt}
			if count > mpscale * mpprec {
				return aexpr, append(errs, ErrStupidShift{y, count})
			}
			xx, _ := convertTypedToConstNumber(x.Const())
			var r *ConstNumber
			if op == token.SHL {
				r = new(ConstNumber).Lsh(xx, uint(count))
			} else {
				r = new(ConstNumber).Rsh(xx, uint(count))
			}
			if !xuntyped {
				c, moreErrs := promoteConstToTyped(ConstShiftedInt, constValueOf(r), xt, aexpr)
				errs = append(errs, moreErrs...)
				aexpr.constValue = c
			} else {
				aexpr.constValue = constValueOf(r)
			}
		}
	} else if x.IsConst() && y.IsConst() {
		if xuntyped && yuntyped {
			yv := y.Const()
			xv := x.Const()
			if promoted, moreErrs := promoteConsts(xc, yc, x, y, xv, yv); moreErrs != nil {
				errs = append(errs, moreErrs...)
				errs = append(errs, ErrInvalidBinaryOperation{aexpr})
			} else {
				if isBooleanOp(op) {
					aexpr.knownType = []reflect.Type{ConstBool}
				} else {
					aexpr.knownType = knownType{promoted}
				}
				aexpr.constValue, moreErrs = evalConstUntypedBinaryExpr(aexpr, promoted)
				if moreErrs != nil {
					errs = append(errs, moreErrs...)
				}
			}
		} else if yuntyped {
			z, moreErrs := evalConstTypedUntypedBinaryExpr(aexpr, x, y, true)
			if moreErrs != nil {
				errs = append(errs, moreErrs...)
			} else {
				aexpr.knownType = knownType{reflect.Value(z).Type()}
				aexpr.constValue = z
			}
		} else if xuntyped {
			z, moreErrs := evalConstTypedUntypedBinaryExpr(aexpr, y, x, false)
			if moreErrs != nil {
				errs = append(errs, moreErrs...)
			} else {
				aexpr.knownType = knownType{reflect.Value(z).Type()}
				aexpr.constValue = z
			}
		} else {
			if z, moreErrs := evalConstTypedBinaryExpr(aexpr, x, y); moreErrs != nil {
				errs = append(errs, moreErrs...)
			} else {
				aexpr.knownType = knownType{reflect.Value(z).Type()}
				aexpr.constValue = z
			}
		}
	} else {
		if yuntyped {
			// the old switcheroo
			xt, yt = yt, xt
			xc, yc = yc, xc
			x, y = y, x
			xuntyped = true
		}
		yk := yt.Kind()
		errExpr := aexpr

		// special cases for const nil
		// note that only (slice|map|func|interface|ptr) == nil are legal
		// other expressions containing ConstNil do not produces a mismatched
		// types error (ErrInvalidBinaryOperation)
		if xt == ConstNil {
			if (op == token.EQL || op == token.NEQ) &&
				(yk == reflect.Slice || yk == reflect.Map || yk == reflect.Func ||
				yk == reflect.Interface || yk == reflect.Ptr) {
				aexpr.knownType = knownType{boolType}
			} else if yk == reflect.String || yk == reflect.Slice || yk == reflect.Interface ||
				yk == reflect.Ptr || yk == reflect.Map {
				// Except strings, they do produce mismatched types
				// instead of bad conversions
				err := ErrInvalidBinaryOperation{aexpr}
				errs = append(errs, err)
			} else {
				err := ErrBadConstConversion{x, xt, yt}
				errs = append(errs, err)
			}
			// http://code.google.com/p/go/issues/detail?id=7206
			if yk == reflect.Array || yk == reflect.Uintptr {
				errs = append(errs, ErrUntypedNil{x})
			}
			return aexpr, errs
		}

		xk := xt.Kind()
		var operandT reflect.Type
		// Identical types are always valid, except non comparable structs
		// and types with can only be compared to nil
                if unhackType(xt) == unhackType(yt) {
			if !comparableToNilOnly(xt) && (xk != reflect.Struct || isStructComparable(xt)) {
				operandT = xt
			}
                } else if xuntyped && attemptBinaryOpConversion(yt) {
	                c, moreErrs := promoteConstToTyped(xc, constValue(x.Const()), yt, x)
			errs = append(errs, moreErrs...)
			v := reflect.Value(c)
			if v.IsValid() {
				operandT = yt
				// Check for divide by zero. Note we only need to do this
				// if y was the untyped constant, and also note that
				// it may have been truncated
				if yuntyped && (op == token.QUO || op == token.REM) &&
					isOpDefinedOn(op, operandT) &&
					v.Interface() == reflect.Zero(yt).Interface() {

					errs = append(errs, ErrDivideByZero{errExpr})
				}
			}
		// An interface is comprable if its paired operand implements it.
		// To match gc error output, if the operator produces a boolean and
		// one operand is a type that satisfies but is not the the other
		// operand's type, wrap that node in a type cast. This will only be
		// used by errors.
		} else if yk == reflect.Interface && xt.Implements(yt) {
			operandT = yt
			if isBooleanOp(op) {
				errExpr = new(BinaryExpr)
				*errExpr = *aexpr
				errExpr.X = wrapConcreteTypeWithInterface(x, yt)
			}
		} else if xk == reflect.Interface && yt.Implements(xt) {
			operandT = xt
			if isBooleanOp(op) {
				errExpr = new(BinaryExpr)
				*errExpr = *aexpr
				errExpr.Y = wrapConcreteTypeWithInterface(y, xt)
			}
                }

		if operandT != nil {
			if !isOpDefinedOn(op, operandT) {
				errs = append(errs, ErrInvalidBinaryOperation{errExpr})
			} else if isBooleanOp(op) {
				aexpr.knownType = knownType{boolType}
			} else {
				aexpr.knownType = knownType{operandT}
			}
                } else {
                        errs = append(errs, ErrInvalidBinaryOperation{errExpr})
		}
        }
	return aexpr, errs
}

// Evaluates a const binary Expr. May return a sensical constValue
// even if ErrTruncatedConst errors are present
func evalConstUntypedBinaryExpr(binary *BinaryExpr, promotedType ConstType) (constValue, []error) {
	x := binary.X.Const()
	y := binary.Y.Const()
	switch promotedType.(type) {
	case ConstIntType, ConstRuneType, ConstFloatType, ConstComplexType:
		xx := x.Interface().(*ConstNumber)
		yy := y.Interface().(*ConstNumber)
		return evalConstBinaryNumericExpr(binary, xx, yy)
	case ConstStringType:
		xx := x.String()
		yy := y.String()
		return evalConstBinaryStringExpr(binary, xx, yy)
	case ConstBoolType:
		xx := x.Bool()
		yy := y.Bool()
		return evalConstBinaryBoolExpr(binary, xx, yy)
	default:
		// It is possible that both x and y are ConstNil, however no operator is defined, not even ==
		return constValue{}, []error{ErrInvalidBinaryOperation{binary}}
	}

}

func evalConstBinaryNumericExpr(binary *BinaryExpr, x, y *ConstNumber) (constValue, []error) {
	var errs []error

	op := binary.Op()
	switch op {
	case token.ADD:
		return constValueOf(new(ConstNumber).Add(x, y)), nil
	case token.SUB:
		return constValueOf(new(ConstNumber).Sub(x, y)), nil
	case token.MUL:
		return constValueOf(new(ConstNumber).Mul(x, y)), nil
	case token.QUO:
		if y.Value.IsZero() {
			return constValue{}, []error{ErrDivideByZero{binary}}
		}
		return constValueOf(new(ConstNumber).Quo(x, y)), nil
	case token.REM:
		if y.Value.IsZero() {
			return constValue{}, []error{ErrDivideByZero{binary}}
		} else if !(x.Type.IsIntegral() && y.Type.IsIntegral()) {
			return constValue{}, []error{ErrInvalidBinaryOperation{binary}}
		} else {
			return constValueOf(new(ConstNumber).Rem(x, y)), nil
		}
	case token.AND, token.OR, token.XOR, token.AND_NOT:
		if !(x.Type.IsIntegral() && y.Type.IsIntegral()) {
			return constValue{}, []error{ErrInvalidBinaryOperation{binary}}
		}

		switch op {
		case token.AND:
			return constValueOf(new(ConstNumber).And(x, y)), nil
		case token.OR:
			return constValueOf(new(ConstNumber).Or(x, y)), nil
		case token.XOR:
			return constValueOf(new(ConstNumber).Xor(x, y)), nil
		case token.AND_NOT:
			return constValueOf(new(ConstNumber).AndNot(x, y)), nil
		default:
			panic("go-interactive: impossible")
		}

	case token.EQL:
		return constValueOf(x.Value.Equals(&y.Value)), nil
	case token.NEQ:
		return constValueOf(!x.Value.Equals(&y.Value)), nil

	case token.LEQ, token.GEQ, token.LSS, token.GTR:
		var b bool
		if !(x.Type.IsReal() && y.Type.IsReal()) {
			return constValue{}, []error{ErrInvalidBinaryOperation{binary}}
		}
		cmp := x.Value.Re.Cmp(&y.Value.Re)
		switch op {
		case token.NEQ:
			b = cmp != 0
		case token.LEQ:
			b = cmp <= 0
		case token.GEQ:
			b = cmp >= 0
		case token.LSS:
			b = cmp < 0
		case token.GTR:
			b = cmp > 0
		}
		return constValueOf(b), errs
	default:
		return constValue{}, []error{ErrInvalidBinaryOperation{binary}}
	}
}

func evalConstBinaryStringExpr(binary *BinaryExpr, x, y string) (constValue, []error) {
	switch binary.Op() {
	case token.ADD:
		return constValueOf(x + y), nil
	case token.EQL:
		return constValueOf(x == y), nil
	case token.NEQ:
		return constValueOf(x != y), nil
	case token.LEQ:
		return constValueOf(x <= y), nil
	case token.GEQ:
		return constValueOf(x >= y), nil
	case token.LSS:
		return constValueOf(x < y), nil
	case token.GTR:
		return constValueOf(x > y), nil
	default:
		return constValue{}, []error{ErrInvalidBinaryOperation{binary}}
	}
}

func evalConstBinaryBoolExpr(binary *BinaryExpr, x, y bool) (constValue, []error) {
	switch binary.Op() {
	case token.EQL:
		return constValueOf(x == y), nil
	case token.NEQ:
		return constValueOf(x != y), nil
	case token. LAND:
		return constValueOf(x && y), nil
	case token.LOR:
		return constValueOf(x || y), nil
	default:
		return constValue{}, []error{ErrInvalidBinaryOperation{binary}}
	}
}

// Evaluate x op y
func evalConstTypedUntypedBinaryExpr(binary *BinaryExpr, typedExpr, untypedExpr Expr, reversed bool) (
	constValue, []error) {

	xt := untypedExpr.KnownType()[0].(ConstType)
	yt := typedExpr.KnownType()[0]
	op := binary.Op()

	// x must be convertible to target type
	xUntyped := untypedExpr.Const()
	x, xConvErrs := promoteConstToTyped(xt, constValue(xUntyped), yt, untypedExpr)

        // If the untyped operand is nil, gc simply says it could not convert the nil type
        if xt == ConstNil {
                // ... unless, its a string. In which case, report mismatched types
                if yt.Kind() == reflect.String {
		        return constValue{}, []error{ErrInvalidBinaryOperation{binary}}
                } else {
		        return constValue{}, xConvErrs
                }
	} else if !isOpDefinedOn(op, yt) {
		return constValue{}, append(xConvErrs, ErrInvalidBinaryOperation{binary})
	}

        // Check for an impossible conversion. This occurs when the types
        // are incompatible, such as string(1.5). For other errors, such as
        // integer overflows, the type check should continue as if the conversion
        // succeeded
        if !reflect.Value(x).IsValid() {
		errs := append(xConvErrs, ErrInvalidBinaryOperation{binary})
		return constValue{}, errs
        }

	switch xt.(type) {
	case ConstIntType, ConstRuneType, ConstFloatType, ConstComplexType:
		xx, xok := convertTypedToConstNumber(reflect.Value(x))
		yy, yok := convertTypedToConstNumber(typedExpr.Const())

		// If a child node errored, then it is possible that, typedExpr.Const() is
		// actually a *ConstNumber to avoid loss of precision in error messages.
		if !yok {
			yy, yok = reflect.Value(typedExpr.Const()).Interface().(*ConstNumber)
		}

		if !xok || !yok {
			// This is a non numeric expression. Return the errors encountered so far
			return constValue{}, append(xConvErrs, ErrInvalidBinaryOperation{binary})
		}

		if reversed {
			xx, yy = yy, xx
		}

		z, errs := evalConstBinaryNumericExpr(binary, xx, yy)
		if errs != nil {
			return constValueOf(z), append(xConvErrs, errs...)
		}
		errs = append(xConvErrs, errs...)

		var zt ConstType
		var rt reflect.Type
		if isBooleanOp(op) {
			zt = ConstBool
			rt = reflect.TypeOf(false)
		} else {
			zt = reflect.Value(z).Interface().(*ConstNumber).Type
			rt = yt
		}

		r, moreErrs := promoteConstToTyped(zt, z, rt, binary)
		return constValue(r), append(errs, moreErrs...)

	case ConstStringType:
		if yt.Kind() == reflect.String {
			xstring := reflect.Value(x).String()
			ystring := typedExpr.Const().String()

			if reversed {
				xstring, ystring = ystring, xstring
			}

			z, errs := evalConstBinaryStringExpr(binary, xstring, ystring)
                        if errs != nil {
                                return constValue{}, errs
                        }

                        var zt ConstType
                        var rt reflect.Type
                        if isBooleanOp(op) {
                                zt = ConstBool
                                rt = reflect.TypeOf(false)
                        } else {
                                zt = ConstString
                                rt = reflect.TypeOf("")
                        }

			r, errs := promoteConstToTyped(zt, z, rt, binary)
			return constValue(r), errs
		}

	case ConstBoolType:
		if yt.Kind() == reflect.Bool {
			xbool := reflect.Value(x).Bool()
			ybool := typedExpr.Const().Bool()

			if reversed {
				xbool, ybool = ybool, xbool
			}

			z, errs := evalConstBinaryBoolExpr(binary, xbool, ybool)
			return constValue(z), errs
		}
	}
	return constValue{}, append(xConvErrs, ErrInvalidBinaryOperation{binary})
}

func evalConstTypedBinaryExpr(binary *BinaryExpr, xexpr, yexpr Expr) (constValue, []error) {

	// These are known not to be ConstTypes
	xt := xexpr.KnownType()[0]
	yt := yexpr.KnownType()[0]
	op := binary.Op()

        // Check that the types are compatible, handling the special alias type for runes
        // For the sake of error messages, for expressions involving int32 and rune, the
        // resulting type is that of the left operand
	var zt reflect.Type
        if xt == yt {
		zt = xt
        } else if xt == RuneType && yt == RuneType.Type || xt == ByteType && yt == ByteType.Type {
                zt = RuneType
        } else if yt == RuneType && xt == RuneType.Type || yt == ByteType && yt == ByteType.Type {
                zt = xt
        } else {
		return constValue{}, []error{ErrInvalidBinaryOperation{binary}}
	}

	x, xok := convertTypedToConstNumber(xexpr.Const())
	y, yok := convertTypedToConstNumber(yexpr.Const())

	if xok && yok {
		z, errs := evalConstBinaryNumericExpr(binary, x, y)
                if isBooleanOp(op) {
                        return constValue(z), errs
                }
                if errs != nil {
                        if _, ok := errs[0].(ErrInvalidBinaryOperation); ok {
                                // This happens if the operator is not defined on x and y
		                return constValue(z), errs
                        }
                }
		from := reflect.Value(z).Interface().(*ConstNumber).Type
		r, moreErrs := promoteConstToTyped(from, z, zt, binary)
		return constValue(r), append(errs, moreErrs...)
	} else if !xok && !yok {
		switch zt.Kind() {
		case reflect.String:
			xstring := xexpr.Const().String()
			ystring := yexpr.Const().String()
			z, errs := evalConstBinaryStringExpr(binary, xstring, ystring)
                        if isBooleanOp(op) {
                                return constValue(z), errs
                        }
			r, moreErrs := promoteConstToTyped(ConstString, z, zt, binary)
			return constValue(r), append(errs, moreErrs...)

		case reflect.Bool:
			xbool := xexpr.Const().Bool()
			ybool := yexpr.Const().Bool()
			z, errs := evalConstBinaryBoolExpr(binary, xbool, ybool)
			return constValue(z), errs
		}
	}
	panic("go-interactive: impossible")
}

func checkBinaryOperands(xexpr, yexpr ast.Expr, env Env) (Expr, Expr, bool, []error) {
	var xok, yok bool
	var err error

	x, errs := CheckExpr(xexpr, env)
	if errs == nil || x.IsConst() {
		if _, err = expectSingleType(x); err != nil {
			errs = append(errs, err)
		} else {
			xok = true
		}
	}

	y, moreErrs := CheckExpr(yexpr, env)
	if moreErrs == nil || y.IsConst() {
		if _, err = expectSingleType(y); err != nil {
			errs = append(moreErrs, err)
		} else {
			yok = true
		}
	}
	errs = append(errs, moreErrs...)
	return x, y, xok && yok, errs
}


func wrapConcreteTypeWithInterface(operand Expr, interfaceT reflect.Type) Expr {
	// Rig the token positions to such that typeConv.(Len|Pos) match operand
	typeConv := &CallExpr{CallExpr: new(ast.CallExpr)}
	typeConv.Fun = &Ident{Ident: &ast.Ident{Name: "", NamePos: operand.Pos()}}
	typeConv.Lparen = operand.Pos()
	typeConv.Rparen = operand.End() - 1
	typeConv.knownType = knownType{interfaceT}
	typeConv.Args = []Expr{operand}
	typeConv.isTypeConversion = true
	return typeConv;
}
