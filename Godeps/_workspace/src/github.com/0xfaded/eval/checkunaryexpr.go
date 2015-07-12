package eval

import (
	"reflect"

	"go/ast"
	"go/token"
)

func checkUnaryExpr(unary *ast.UnaryExpr, env Env) (*UnaryExpr, []error) {
	aexpr := &UnaryExpr{UnaryExpr: unary}

	x, errs := CheckExpr(unary.X, env)
	if errs == nil || x.IsConst() {
		if t, err := expectSingleType(x); err != nil {
			errs = append(errs, err)
		} else if unary.Op == token.AND { // address of
			if !isAddressableOrCompositeLit(x) {
				printableX := fakeCheckExpr(unary.X, env)
				printableX.setKnownType(knownType{t})
				errs = append(errs, ErrInvalidAddressOf{printableX})
			}
			t := x.KnownType()[0]
			if ct, ok := t.(ConstType); ok {
				if ct == ConstNil {
					errs = append(errs, ErrUntypedNil{x})
				} else {
					ptrT := reflect.PtrTo(unhackType(ct.DefaultPromotion()))
					aexpr.knownType = knownType{ptrT}
				}
			} else {
				ptrT := reflect.PtrTo(unhackType(t))
				aexpr.knownType = knownType{ptrT}
			}
			aexpr.X = x
		} else if unary.Op == token.ARROW { // <-
			if (t.Kind() != reflect.Chan) || (t.ChanDir() | reflect.RecvDir == 0) {
				errs = append(errs, ErrInvalidRecvFrom{x})
			}
		} else {
			aexpr.X = x
			// All numeric and bool unary expressions do not change type
			aexpr.knownType = knownType(x.KnownType())
			if x.IsConst() {
				if ct, ok := t.(ConstType); ok {
					c, moreErrs := evalConstUnaryExpr(aexpr, x, ct)
					if moreErrs != nil {
						errs = append(errs, moreErrs...)
					}
					aexpr.constValue = c

				} else {
					c, moreErrs := evalConstTypedUnaryExpr(aexpr, x)
					if moreErrs != nil {
						errs = append(errs, moreErrs...)
					}
					aexpr.constValue = c
				}
			} else {
				if !isUnaryOpDefinedOn(unary.Op, t) {
					errs = append(errs, ErrInvalidUnaryOperation{aexpr})
				}
			}
		}
        }
	aexpr.X = x
	return aexpr, errs
}

func evalConstUnaryExpr(unary *UnaryExpr, x Expr, xT ConstType) (constValue, []error) {
	switch xT.(type) {
	case ConstIntType, ConstRuneType, ConstFloatType, ConstComplexType:
		xx := x.Const().Interface().(*ConstNumber)
		return evalConstUnaryNumericExpr(unary, xx)
	case ConstBoolType:
		xx := x.Const().Bool()
		return evalConstUnaryBoolExpr(unary, xx)
	default:
		return constValue{}, []error{ErrInvalidUnaryOperation{unary}}
	}
}

func evalConstTypedUnaryExpr(unary *UnaryExpr, x Expr) (constValue, []error) {
	t := x.KnownType()[0]
	if xx, ok := convertTypedToConstNumber(x.Const()); ok {
		zz, errs := evalConstUnaryNumericExpr(unary, xx)
		if !reflect.Value(zz).IsValid() {
			return constValue{}, errs
		}
		rr, moreErrs := promoteConstToTyped(xx.Type, zz, t, x)
		return rr, append(errs, moreErrs...)
	} else if t.Kind() == reflect.Bool {
		return evalConstUnaryBoolExpr(unary, x.Const().Bool())
	}
	return constValue{}, []error{ErrInvalidUnaryOperation{unary}}
}

func evalConstUnaryNumericExpr(unary *UnaryExpr, x *ConstNumber) (constValue, []error) {
	switch unary.Op {
	case token.ADD:
		return constValueOf(x), nil
	case token.SUB:
		zero := &ConstNumber{Type: x.Type}
		return constValueOf(zero.Sub(zero, x)), nil
	case token.XOR:
		if x.Type.IsIntegral() {
			minusOne := NewConstInt64(-1)
			return constValueOf(minusOne.Xor(minusOne, x)), nil
		}
	}
	return constValue{}, []error{ErrInvalidUnaryOperation{unary}}
}

func evalConstUnaryBoolExpr(unary *UnaryExpr, x bool) (constValue, []error) {
	switch unary.Op {
	case token.NOT:
		return constValueOf(!x), nil
	default:
		return constValue{}, []error{ErrInvalidUnaryOperation{unary}}
	}
}
