package eval

import (
	"reflect"
)

func evalFuncLit(lit *FuncLit, env Env) (reflect.Value, error) {
	ft := lit.Type.KnownType()[0]
	env = env.PushScope()
	v := reflect.MakeFunc(ft, func(in []reflect.Value) (out []reflect.Value) {
		i := 0
		for _, field := range lit.Type.Params.List {
			if field.Names != nil {
				for _, name := range field.Names {
					if name.Name != "_" {
						v := reflect.New(ft.In(i))
						v.Elem().Set(in[i])
						env.AddVar(name.Name, v)
					}
					i += 1
				}
			}
		}
		i = 0
		if lit.Type.Results != nil {
			for _, field := range lit.Type.Results.List {
				for _, name := range field.Names {
					if name.Name != "_" {
						env.AddVar(name.Name, reflect.New(ft.Out(i)))
					}
					i += 1
				}
			}
		}
		if last, err := InterpStmt(lit.Body, env); err != nil {
			panic(err)
		} else if ft.NumOut() == 0 {
			// void func which may terminate on a non-ReturnStmt
		} else if ret := last.Last.(*ReturnStmt); len(ret.Results) == ft.NumOut() {
			for i, result := range ret.Results {
				if r, err := evalTypedExpr(result, knownType{ft.Out(i)}, last.Env); err != nil {
					panic(err)
				} else {
					out = append(out, r[0])
				}
			}
		} else if len(ret.Results) == 1 {
			// return multi()
			if rs, err := EvalExpr(ret.Results[0], env); err != nil {
				panic(err)
			} else {
				out = rs
			}
		} else {
			for _, field := range lit.Type.Results.List {
				for _, name := range field.Names {
					if name.Name != "_" {
						out = append(out, env.Var(name.Name).Elem())
					} else {
						out = append(out, reflect.Zero(field.KnownType()[0]))
					}
				}
			}
		}
		return out
	})
	return v, nil
}
