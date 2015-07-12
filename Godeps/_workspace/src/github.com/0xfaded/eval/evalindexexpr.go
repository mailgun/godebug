package eval

import (
	"reflect"
)

func evalIndexExpr(index *IndexExpr, env Env) ([]reflect.Value, error) {
	xs, err := EvalExpr(index.X, env)
	if err != nil {
		return []reflect.Value{}, err
	}
	x := xs[0]

	t := index.X.KnownType()[0]
	switch t.Kind() {
	case reflect.Map:
		k, err := evalTypedExpr(index.Index, knownType{t.Key()}, env)
		if err != nil {
			return []reflect.Value{}, err
		}
		v := x.MapIndex(k[0])
		ok := v.IsValid()
		if !ok {
			v = reflect.New(t.Key()).Elem()
		}
		return []reflect.Value{v, reflect.ValueOf(ok)}, nil
	case reflect.Ptr:
		// Short hand for array pointers
		x = x.Elem()
		fallthrough
	default:
		i, err := evalInteger(index.Index, env)
		if err != nil {
			return []reflect.Value{}, err
		}
		if !(0 <= i && i < x.Len()) {
			return []reflect.Value{}, PanicIndexOutOfBounds{}
		}
		return []reflect.Value{x.Index(i)}, nil
	}
}
