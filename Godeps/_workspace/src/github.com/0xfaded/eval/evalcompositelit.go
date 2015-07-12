package eval

import (
	"errors"
	"fmt"
	"reflect"
)

func evalCompositeLit(lit *CompositeLit, env Env) (reflect.Value, error) {
	t := lit.KnownType()[0]

	switch t.Kind() {
	case reflect.Map:
		return evalCompositeLitMap(t, lit, env)
	case reflect.Array, reflect.Slice:
		return evalCompositeLitArrayOrSlice(t, lit, env)
	case reflect.Struct:
		return evalCompositeLitStruct(t, lit, env)
	default:
		return reflect.Value{}, errors.New(fmt.Sprintf("eval: unimplemented type for composite literal %s", t.Name()))
	}
}

func evalCompositeLitMap(t reflect.Type, lit *CompositeLit, env Env) (reflect.Value, error) {

	m := reflect.New(t).Elem()
	m.Set(reflect.MakeMap(t))

	kT := knownType{t.Key()}
	vT := knownType{t.Elem()}
	for _, elt := range lit.Elts {
		kv := elt.(*KeyValueExpr)
		k, err := evalTypedExpr(kv.Key, kT, env)
		if err != nil {
			return reflect.Value{}, err
		}
		if kT[0].Kind() == reflect.Interface {
			dynamicT := k[0].Elem().Type()
			if !isStaticTypeComparable(dynamicT) {
				return reflect.Value{}, PanicUnhashableType{dynamicT}
			}
		}
		v, err := evalTypedExpr(kv.Value, vT, env)
		if err != nil {
			return reflect.Value{}, err
		}
		m.SetMapIndex(k[0], v[0])
	}
	return m, nil
}

func evalCompositeLitArrayOrSlice(t reflect.Type, lit *CompositeLit, env Env) (reflect.Value, error) {

	v := reflect.New(t).Elem()
	if t.Kind() == reflect.Slice {
		v.Set(reflect.MakeSlice(t, lit.length, lit.length))
	}

	eT := knownType{t.Elem()}
	for src, dst, i := 0, 0, 0; src < len(lit.Elts); src, dst = src + 1, dst + 1 {
		var elt Expr
		if lit.indices[i].pos == src {
			elt = lit.Elts[src].(*KeyValueExpr).Value
			dst = lit.indices[i].index
			i += 1
		} else {
			elt = lit.Elts[src]
		}
		if elem, err := evalTypedExpr(elt, eT, env); err != nil {
			return reflect.Value{}, err
		} else {
			v.Index(dst).Set(elem[0])
		}
	}
	return v, nil
}

func evalCompositeLitStruct(t reflect.Type, lit *CompositeLit, env Env) (reflect.Value, error) {
	v := reflect.New(t).Elem()
	for i, f := range lit.fields {
		var elt Expr
		if kv, ok := lit.Elts[i].(*KeyValueExpr); ok {
			elt = kv.Value
		} else {
			elt = lit.Elts[i]
		}
		field := v.Field(f)
		if elem, err := evalTypedExpr(elt, knownType{field.Type()}, env); err != nil {
			return reflect.Value{}, err
		} else {
			field.Set(elem[0])
		}
	}
	return v, nil
}
