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
	foo(3, 3)
}

func foo(int, int) (string, error) {
	var godebugInput1 int
	var godebugInput2 int
	var godebugResult1 string
	var godebugResult2 error
	ctx, ok := godebug.EnterFunc(func() {
		godebugResult1, godebugResult2 = foo(godebugInput1, godebugInput2)
	})
	if !ok {
		return godebugResult1, godebugResult2
	}
	defer godebug.ExitFunc()
	godebug.Line(ctx, unnamed_input_in_goScope)
	return "hello", nil
}
