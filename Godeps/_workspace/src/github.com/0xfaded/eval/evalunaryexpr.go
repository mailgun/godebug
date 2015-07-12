package eval

import (
	"reflect"

	"go/token"
)

func evalUnaryExpr(unary *UnaryExpr, env Env) ([]reflect.Value, error) {
	if unary.IsConst() {
		return []reflect.Value{unary.Const()}, nil
	}

	xx, err := EvalExpr(unary.X, env)
	if err != nil {
		return []reflect.Value{}, err
	}
	x := xx[0]

	// handle & and <- first
	if unary.Op == token.AND {
		return []reflect.Value{x.Addr()}, nil
	} else if unary.Op == token.ARROW {
		// TODO[crc] use x.Recv in contexts where blocking is acceptable
		v, ok := x.TryRecv()
		if !v.IsValid() {
			v = reflect.New(x.Type().Elem()).Elem()
		}
		// TODO[crc] also return ok once assignability context is implemented
		return []reflect.Value{v, reflect.ValueOf(ok)}, nil
	}

	var r reflect.Value
	switch x.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		r, err = evalUnaryIntExpr(x, unary.Op)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		r, err = evalUnaryUintExpr(x, unary.Op)
	case reflect.Float32, reflect.Float64:
		r, err = evalUnaryFloatExpr(x, unary.Op)
	case reflect.Complex64, reflect.Complex128:
		r, err = evalUnaryComplexExpr(x, unary.Op)
	case reflect.Bool:
		r, err = evalUnaryBoolExpr(x, unary.Op)
	default:
		panic("eval: impossible unary op " + unary.Op.String())
	}
	return []reflect.Value{r}, err
}

func evalUnaryIntExpr(x reflect.Value, op token.Token) (reflect.Value, error) {
	var r int64
	var err error

	xx := x.Int()
	switch op {
	case token.ADD: r = +xx
	case token.SUB: r = -xx
	case token.XOR: r = ^xx
	default:
		panic("eval: impossible unary op " + op.String())
	}
	return reflect.ValueOf(r).Convert(x.Type()), err
}

func evalUnaryUintExpr(x reflect.Value, op token.Token) (reflect.Value, error) {
	var err error
	var r uint64

	xx := x.Uint()
	switch op {
	case token.ADD: r = +xx
	case token.SUB: r = -xx
	case token.XOR: r = ^xx
	default:
		panic("eval: impossible unary op " + op.String())
	}
	return reflect.ValueOf(r).Convert(x.Type()), err
}

func evalUnaryFloatExpr(x reflect.Value, op token.Token) (reflect.Value, error) {
	var err error
	var r float64

	xx := x.Float()
	switch op {
	case token.ADD: r = + xx
	case token.SUB: r = - xx
	default:
		panic("eval: impossible unary op " + op.String())
	}
	return reflect.ValueOf(r).Convert(x.Type()), err
}

func evalUnaryComplexExpr(x reflect.Value, op token.Token) (reflect.Value, error) {
	var err error
	var r complex128

	xx := x.Complex()
	switch op {
	case token.ADD: r = +xx
	case token.SUB: r = -xx
	default:
		panic("eval: impossible unary op " + op.String())
	}
	return reflect.ValueOf(r).Convert(x.Type()), err
}

func evalUnaryBoolExpr(x reflect.Value, op token.Token) (reflect.Value, error) {
	var err error
	var r bool

	xx := x.Bool()
	switch op {
	case token.NOT: r = !xx
	default:
		panic("eval: impossible unary op " + op.String())
	}
	return reflect.ValueOf(r).Convert(x.Type()), err
}
