package eval

import (
	"reflect"
)

func evalTypeAssertExpr(assert *TypeAssertExpr, env Env) ([]reflect.Value, error) {
	x := assert.X
	if vs, err := EvalExpr(x, env); err != nil {
		return []reflect.Value{}, err
	} else {
		v := vs[0]
		xT := x.KnownType()[0]
		aT := assert.KnownType()[0]
		if v.IsNil() {
			vs := []reflect.Value{hackedNew(aT).Elem(), reflect.ValueOf(false)}
			return vs, PanicInterfaceConversion{aT: aT}
		}
		dynamic := v.Elem()
		dT := dynamic.Type()
		if aT.Kind() == reflect.Interface {
			if !dT.Implements(aT) {
				vs := []reflect.Value{hackedNew(aT).Elem(), reflect.ValueOf(false)}
				return vs, PanicInterfaceConversion{xT, aT, nil}
			}
		} else {
			if dT != aT {
				vs := []reflect.Value{hackedNew(aT).Elem(), reflect.ValueOf(false)}
				return vs, PanicInterfaceConversion{xT, aT, dT}
			}
		}
		r := reflect.New(aT).Elem()
		r.Set(dynamic)
		return []reflect.Value{r, reflect.ValueOf(true)}, nil
	}
}
