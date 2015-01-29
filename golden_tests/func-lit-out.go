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
	godebugScope := func_lit_in_goScope.EnteringNewChildScope()
	godebugScope.Declare("hi", &hi, "there", &there)
	godebug.Line(ctx, godebugScope)
	fmt.Println(hi, there)
	godebug.Line(ctx, godebugScope)
	bar()
}

var foo = func(a, _ int) (b, _ string) {
	var godebugResult2 string
	var ctx *godebug.Context
	fn := func() {
		b, godebugResult2 = func() (b, _ string) {
			godebugScope := func_lit_in_goScope.EnteringNewChildScope()
			godebugScope.Declare("a", &a, "b", &b)
			godebug.Line(ctx, godebugScope)
			return "Hello", "World"
		}()
	}
	var ok bool
	ctx, ok = godebug.EnterFunc(fn)
	if ok {
		fn()
	}
	godebug.ExitFunc()
	return b, godebugResult2
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
