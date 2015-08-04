package main

import (
	"fmt"
	"github.com/mailgun/godebug/lib"
)

var func_lit_in_go_scope = godebug.EnteringNewFile(main_pkg_scope, func_lit_in_go_contents)

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.Line(ctx, func_lit_in_go_scope, 6)
	hi, there := foo(7, 12)
	scope := func_lit_in_go_scope.EnteringNewChildScope()
	scope.Declare("hi", &hi, "there", &there)
	godebug.Line(ctx, scope, 7)
	fmt.Println(hi, there)
	godebug.Line(ctx, scope, 8)
	bar()
}

var foo = func(a, _ int) (b, _ string) {
	var result2 string
	fn := func(ctx *godebug.Context) {
		b, result2 = func() (b, _ string) {
			scope := func_lit_in_go_scope.EnteringNewChildScope()
			scope.Declare("a", &a, "b", &b)
			godebug.Line(ctx, scope, 12)
			return "Hello", "World"
		}()
	}
	if ctx, ok := godebug.EnterFuncLit(fn); ok {
		defer godebug.ExitFunc(ctx)
		fn(ctx)
	}
	return b, result2
}

var bar = func() {
	fn := func(ctx *godebug.Context) {
		godebug.Line(ctx, func_lit_in_go_scope, 16)
		fmt.Println("No inputs or outputs")
	}
	if ctx, ok := godebug.EnterFuncLit(fn); ok {
		defer godebug.ExitFunc(ctx)
		fn(ctx)
	}
}

var func_lit_in_go_contents = `package main

import "fmt"

func main() {
	hi, there := foo(7, 12)
	fmt.Println(hi, there)
	bar()
}

var foo = func(a, _ int) (b, _ string) {
	return "Hello", "World"
}

var bar = func() {
	fmt.Println("No inputs or outputs")
}
`


var main_pkg_scope = &godebug.Scope{}

func init() {
	main_pkg_scope.Vars = map[string]interface{}{
		"foo": &foo,
		"bar": &bar,
	}
	main_pkg_scope.Consts = map[string]interface{}{
	}
	main_pkg_scope.Funcs = map[string]interface{}{
		"main": main,
	}
}