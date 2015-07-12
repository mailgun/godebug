eval - A library for providing an eval function in Go
============================================================================
[![Build Status](https://travis-ci.org/0xfaded/eval.png)](https://travis-ci.org/0xfaded/eval)

This project adds an *Eval()* function to go.

Right now only, Go expressions are handled.

Using
-----

The simplest invocation is

```
	results, panik, compileErrs := Eval("1 + 2")
	// results = []reflect.Value{reflect.ValueOf(3)}
	// panik = error(nil)
	// compileErrs = []error(nil)
```
Results are reflect.Values, and in reality Eval is nothing more than a fancy wrapper around the reflect package. Runtime panics should be detected during evaluation and returned as panik, if an actual runtime panic
occurs please file a bug. Any parse or type errors are returned as compileErrs; it is nil otherwise.

EvalEnv evaluates an expression inside an environment containing variables, constants, functions, types and packages.
```
	type T int
	x := 1
	sum := func(xs ...int) (result int) {
		for _, x := range xs {
			result += x
		}
		return result
	}
	env := eval.MakeSimpleEnv()
	// Note the &x
	env.Vars["x"] = reflect.ValueOf(&x)
	env.Funcs["sum"] = reflect.ValueOf(f)
	env.Types["T"] = reflect.TypeOf(T(0))
	pkg := MakeSimpleEnv()
	pkg.Consts["C"] = reflect.ValueOf(2)
	env.Pkgs["pkg"] = pkg
	results, panik, compileErrs := EvalEnv("T(sum(1, x, pkg.C))", env)
```

Extended API
------------

The extented API allows step by step execution of the evaluator. The dance is three part
 - Parse
 - Check
 - Eval

```
	env := eval.MakeSimpleEnv()
	if expr, err := parser.ParseExpr(expr); err != nil {
		fmt.Printf("parse error: %s\n", err)
	} else if cexpr, errs := eval.CheckExpr(expr, env); len(errs) != 0 {
		for _, cerr := range errs {
			fmt.Printf("%v\n", cerr)
		}
	} else if vals, _, err := eval.EvalExpr(cexpr, env); err != nil {
		fmt.Printf("eval error: %s\n", err)
	} else {
	  // do something with pointer to reflect.Value array vals, e.g.:
	  fmt.Println(vals[0].Interface())
	}
```


The program [repl.go](https://github.com/0xfaded/eval/tree/master/demo/repl.go) is a full Go program showing this.

Limitations
-----------

Eval is currently limited to the functionality of the reflect package.

Most noteably, the following are not implemented as they cannot be created:
 - struct literals
 - array literals
 - function literals

Struct and Array composite named types can still be constructed. E.g.
```
	env.Types["A"] = reflect.TypeOf([2]int{})
	EvalEnv("A{1, 2}", env)
```
In theory this could also work for named Function literals, but this has not been implemented.


See Also
--------

* [What's left to do?](https://github.com/0xfaded/eval/wiki/What's-left-to-do%3F)
* [go-fish](https://github.com/rocky/go-fish): an interactive read, eval, print loop which uses this to handle the *eval()* step. In that project, see the program *make_eval* for how to create a complete environment given an initial import.
* [gack](https://github.com/0xfaded/gack): another experimental REPL which implements import by continuously recompiling the executable.
* [gub debugger](https://github.com/rocky/ssa-interp): a debugger that uses this to handle the *eval* debugger command
