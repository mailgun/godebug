package main

import "github.com/mailgun/godebug/lib"

var unnamed_input_in_goScope = godebug.EnteringNewScope()

func main() {
	ctx, ok := godebug.EnterFunc(func() {
		main()
	})
	if !ok {
		return
	}
	godebug.Line(ctx, unnamed_input_in_goScope)
	foo(3)
}

func foo(int) string {
	var godebugInput1 int
	var godebugResult1 string
	ctx, ok := godebug.EnterFunc(func() {
		godebugResult1 = foo(godebugInput1)
	})
	if !ok {
		return godebugResult1
	}
	defer godebug.ExitFunc()
	godebug.Line(ctx, unnamed_input_in_goScope)
	return "hello"
}
