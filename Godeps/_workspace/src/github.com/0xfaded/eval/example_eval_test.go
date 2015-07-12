package eval_test

import (
	"fmt"
	"github.com/mailgun/godebug/Godeps/_workspace/src/github.com/0xfaded/eval" //makeEnv creates an environment to use in eval
	"go/parser"
	"reflect"
)

const constant1 = "A constant"

func makeEnv() *eval.SimpleEnv {
	env := eval.MakeSimpleEnv()
	env.Consts["constant1"] = reflect.ValueOf(constant1)
	var1 := 1
	env.Vars["var1"] = reflect.ValueOf(&var1)
	env.Funcs["add"] = reflect.ValueOf(func(a, b int) int { return a + b })
	fmtPkg := eval.MakeSimpleEnv()
	fmtPkg.Funcs["Sprintf"] = reflect.ValueOf(fmt.Sprintf)
	env.Pkgs["fmt"] = fmtPkg
	return env

}

func ExampleSimple() {
	result, panik, compileErrs := eval.Eval(`append([]int{1,2}[:1], map[int]int{1:2}[1])[1]`)
	_ = panik
	_ = compileErrs
	fmt.Printf("%+v\n", result[0].Interface())
}

func ExampleWithEnvironment() {
	env := makeEnv() // Create evaluation environment
	result, panik, compileErrs := eval.EvalEnv(`fmt.Sprintf("%s %d", constant1, add(var1, 1) + 1)`, env)
	_ = panik
	_ = compileErrs
	fmt.Printf("%+v\n", result[0].Interface())
}

// ExpectResult check the evaluation of a string with an expected result.
// More importantly though, does the steps to evaluate a string:
//   1. Parse expression using parser.ParseExpr (go/parser)
//   2. Type check expression using evalCheckExpr (0xfaded/eval)
//   3. run eval.EvalExpr (0xfaded/eval)
func ExampleFullApi() {
	expr := `fmt.Sprintf("%s %d", constant1, add(var1, 1) + 1)`
	env := makeEnv() // Create evaluation environment
	if e, err := parser.ParseExpr(expr); err != nil {
		fmt.Printf("Failed to parse expression '%s' (%v)\n", expr, err)
	} else if cexpr, errs := eval.CheckExpr(e, env); len(errs) != 0 {
		fmt.Printf("Error checking expression '%s' (%v)\n", expr, errs)
	} else if results, err := eval.EvalExpr(cexpr, env); err != nil {
		fmt.Printf("Panic evaluating expression '%s' (%v)\n", expr, err)
	} else {
		fmt.Printf("Expression '%s' yielded '%+v'\n", expr, results[0].Interface())
	}
}
