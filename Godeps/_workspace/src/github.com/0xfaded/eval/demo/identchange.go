// Demos replacing the default identifier lookup value mechanism with
// our own custom version.

package main

import (
	"fmt"
	"github.com/mailgun/godebug/Godeps/_workspace/src/github.com/0xfaded/eval"
	"go/parser"
	"reflect"
	"strings"
)

func expectResult(expr string, env eval.Env, expected interface{}) {
	if e, err := parser.ParseExpr(expr); err != nil {
		fmt.Printf("Failed to parse expression '%s' (%v)\n", expr, err)
		return
	} else if cexpr, errs := eval.CheckExpr(e, env); len(errs) != 0 {
		fmt.Printf("Error checking expression '%s' (%v)\n", expr, errs)
	} else if results, err := eval.EvalExpr(cexpr, env); err != nil {
		fmt.Printf("Error evaluating expression '%s' (%v)\n", expr, err)
		return
	} else {
		fmt.Printf("Expression '%s' yielded '%+v', expected '%+v'\n",
			expr, results[0].Interface(), expected)
	}
}

type CustomEnv struct {
	// Encapsulate the default functionality
	eval.Env
}

// At a minimum, You will probably want to define PushScope and PopScope
func (env *CustomEnv) PushScope() eval.Env {
	// The newly created underlying
	top := env.Env.PushScope()

	// Wrap the env and return it
	return &CustomEnv{top}
}

func (env *CustomEnv) PopScope() eval.Env {
	if top := env.Env.PopScope(); top == nil {
		return nil
	} else {
		return &CustomEnv{top}
	}
}

// Our custom Var will add one to any and all ints
func (env *CustomEnv) Var(ident string) reflect.Value {
	// Note that variables are always pointer values.
	if v := env.Env.Var(ident); v.IsValid() && v.Type().Elem().Kind() == reflect.Int {
		i := v.Elem().Int()
		plusOne := reflect.New(v.Type().Elem()).Elem()
		plusOne.Set(reflect.ValueOf(int(i + 1)))
		return plusOne.Addr()
	} else {
		return v
	}
}

// Our custom Const will lowercase all strings
func (env *CustomEnv) Const(ident string) reflect.Value {
	if v := env.Env.Const(ident); v.IsValid() && v.Type().Kind() == reflect.String {
		s := v.String()
		return (reflect.ValueOf(strings.ToLower(s)))
	} else {
		return v
	}
}

func main() {
	simpleEnv := eval.MakeSimpleEnv()
	v := 4
	c := "CONSTANT"
	simpleEnv.Vars["v"] = reflect.ValueOf(&v)
	simpleEnv.Consts["c"] = reflect.ValueOf(c)
	env := &CustomEnv{simpleEnv}
	expectResult("v + 1", env, "6")
	expectResult("c + \" value\"", env, "constant value")

}
