package main

import (
	"fmt"
	"github.com/mailgun/godebug/lib"
)

var func_lit_in_goScope = godebug.EnteringNewScope()

func main() {
	ctx, ok := godebug.EnterFunc(func() {
		main()
	})
	if !ok {
		return
	}
	godebug.Line(ctx, func_lit_in_goScope)
	hi, there := foo(7, 12)
	scope := func_lit_in_goScope.EnteringNewChildScope()
	scope.Declare("hi", &hi, "there", &there)
	godebug.Line(ctx, scope)
	fmt.Println(hi, there)
	godebug.Line(ctx, scope)
	bar()
}

var foo = func(a, _ int) (b, _ string) {
	var result2 string
	var ctx *godebug.Context
	fn := func() {
		b, result2 = func() (b, _ string) {
			scope := func_lit_in_goScope.EnteringNewChildScope()
			scope.Declare("a", &a, "b", &b)
			godebug.Line(ctx, scope)
			return "Hello", "World"
		}()
	}
	var ok bool
	ctx, ok = godebug.EnterFunc(fn)
	if ok {
		fn()
	}
	godebug.ExitFunc()
	return b, result2
}

var bar = func() {
	var ctx *godebug.Context
	fn := func() {
		godebug.Line(ctx, func_lit_in_goScope)
		fmt.Println("No inputs or outputs")
	}
	var ok bool
	ctx, ok = godebug.EnterFunc(fn)
	if ok {
		fn()
	}
	godebug.ExitFunc()
}
