// Example repl is a simple REPL (read-eval-print loop) for GO using
// http://github.com/0xfaded/eval to the heavy lifting to implement
// the eval() part.
//
// The intent here is to show how more to use the library, rather than
// be a full-featured REPL.
//
// A more complete REPL including command history, tab completion and
// readline editing is available as a separate package:
// http://github.com/rocky/go-fish
//
// (rocky) My intent here is also to have something that I can debug in
// the ssa-debugger tortoise/gub.sh. Right now that can't handle the
// unsafe package, pointers, and calls to C code. So that let's out
// go-gnureadline and lineedit.
package main

import (
	"bufio"
	"fmt"
	"go/parser"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/mailgun/godebug/Godeps/_workspace/src/github.com/0xfaded/eval"
)

// Simple replacement for GNU readline
func readline(prompt string, in *bufio.Reader) (string, error) {
	fmt.Printf(prompt)
	line, err := in.ReadString('\n')
	if err == nil {
		line = strings.TrimRight(line, "\r\n")
	}
	return line, err
}

func intro_text() {
	fmt.Printf(`=== A simple Go eval REPL ===

Results of expression are stored in variable slice "results".
The environment is stored in global variable "env".

Enter expressions to be evaluated at the "go>" prompt.

To see all results, type: "results".

To quit, enter: "quit" or Ctrl-D (EOF).
`)

}

// REPL is the a read, eval, and print loop.
func REPL(env *eval.SimpleEnv) {

	var err error

	// A place to store result values of expressions entered
	// interactively
	results := make([]interface{}, 0, 10)
	env.Vars["results"] = reflect.ValueOf(&results)

	exprs := 0
	in := bufio.NewReader(os.Stdin)
	line, err := readline("go> ", in)
	for line != "quit" {
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if expr, err := parser.ParseExpr(line); err != nil {
			if pair := eval.FormatErrorPos(line, err.Error()); len(pair) == 2 {
				fmt.Println(pair[0])
				fmt.Println(pair[1])
			}
			fmt.Printf("parse error: %s\n", err)
		} else if cexpr, errs := eval.CheckExpr(expr, env); len(errs) != 0 {
			for _, cerr := range errs {
				fmt.Printf("check error: %v\n", cerr)
			}
		} else if vals, err := eval.EvalExpr(cexpr, env); err != nil {
			fmt.Printf("panic: %s\n", err)
		} else if len(vals) == 0 {
			fmt.Printf("Kind=Slice\nvoid\n")
		} else if len(vals) == 1 {
			value := (vals)[0]
			if value.IsValid() {
				kind := value.Kind().String()
				typ := value.Type().String()
				if typ != kind {
					fmt.Printf("Kind = %v\n", kind)
					fmt.Printf("Type = %v\n", typ)
				} else {
					fmt.Printf("Kind = Type = %v\n", kind)
				}
				fmt.Printf("results[%d] = %s\n", exprs, eval.Inspect(value))
				exprs += 1
				results = append(results, (vals)[0].Interface())
			} else {
				fmt.Printf("%s\n", value)
			}
		} else {
			fmt.Printf("Kind = Multi-Value\n")
			size := len(vals)
			for i, v := range vals {
				fmt.Printf("%s", eval.Inspect(v))
				if i < size-1 {
					fmt.Printf(", ")
				}
			}
			fmt.Printf("\n")
			exprs += 1
			results = append(results, vals)
		}

		line, err = readline("go> ", in)
	}
}

type XI interface {
	x()
}
type YI interface {
	y()
}
type ZI interface {
	x()
}

type X int
type Y int
type Z int

func (X) x() {}
func (Y) y() {}
func (Z) x() {}

// Create an eval.Env environment to use in evaluation.
// This is a bit ugly here, because we are rolling everything by hand, but
// we want some sort of environment to show off in demo'ing.
// The artifical environment we create here consists of
//   fmt:
//      fns: fmt.Println, fmt.Printf
//   os:
//      types: MyInt
//      vars: Stdout, Args
//   main:
//      type Alice
//      var  alice, aliceptr
//
// (REPL also adds var results to main)
//
// See make_env in github.com/rocky/go-fish for an automated way to
// create more complete environment from a starting import.
func makeBogusEnv() *eval.SimpleEnv {

	// A copule of things from the fmt package.
	var fmt_funcs map[string]reflect.Value = make(map[string]reflect.Value)
	fmt_funcs["Println"] = reflect.ValueOf(fmt.Println)
	fmt_funcs["Printf"] = reflect.ValueOf(fmt.Printf)

	// A simple type for demo
	type MyInt int

	// A stripped down package environment.  See
	// http://github.com/rocky/go-fish and repl_imports.go for a more
	// complete environment.
	pkgs := map[string]eval.Env{
		"fmt": &eval.SimpleEnv{
			Vars:   make(map[string]reflect.Value),
			Consts: make(map[string]reflect.Value),
			Funcs:  fmt_funcs,
			Types:  make(map[string]reflect.Type),
			Pkgs:   nil,
		}, "os": &eval.SimpleEnv{
			Vars: map[string]reflect.Value{
				"Stdout": reflect.ValueOf(&os.Stdout),
				"Args":   reflect.ValueOf(&os.Args)},
			Consts: make(map[string]reflect.Value),
			Funcs:  make(map[string]reflect.Value),
			Types: map[string]reflect.Type{
				"MyInt": reflect.TypeOf(*new(MyInt))},
			Pkgs: nil,
		},
	}

	mainEnv := eval.MakeSimpleEnv()
	mainEnv.Pkgs = pkgs

	// Some "alice" things for testing
	type Alice struct {
		Bob    int
		Secret string
	}

	type R rune

	alice := Alice{1, "shhh"}
	alicePtr := &alice
	foo := 10
	ints := []int{1, 2, 3, 4}
	add := func(a, b int) int {
		return a + b
	}
	sum := func(as ...int) int {
		r := 0
		for _, a := range as {
			r += a
		}
		return r
	}

	mainEnv.Vars["alice"] = reflect.ValueOf(&alice)
	mainEnv.Vars["alicePtr"] = reflect.ValueOf(&alicePtr)
	mainEnv.Vars["foo"] = reflect.ValueOf(&foo)
	mainEnv.Vars["ints"] = reflect.ValueOf(&ints)
	mainEnv.Consts["bar"] = reflect.ValueOf(eval.NewConstInt64(5))
	mainEnv.Funcs["add"] = reflect.ValueOf(add)
	mainEnv.Funcs["sum"] = reflect.ValueOf(sum)
	mainEnv.Types["Alice"] = reflect.TypeOf(Alice{})
	mainEnv.Types["R"] = reflect.TypeOf(R(0))

	var xi *XI = new(XI)
	var yi *YI = new(YI)
	var zi *ZI = new(ZI)
	*xi = XI(X(0))
	*yi = YI(Y(0))
	*zi = ZI(Z(0))
	mainEnv.Types["XI"] = reflect.TypeOf(xi).Elem()
	mainEnv.Types["YI"] = reflect.TypeOf(yi).Elem()
	mainEnv.Types["ZI"] = reflect.TypeOf(zi).Elem()
	mainEnv.Types["X"] = reflect.TypeOf(X(0))
	mainEnv.Types["Y"] = reflect.TypeOf(Y(0))
	mainEnv.Types["Z"] = reflect.TypeOf(Z(0))

	return mainEnv
}

func main() {
	env := makeBogusEnv()
	intro_text()
	REPL(env)
}
