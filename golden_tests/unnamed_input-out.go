package main

import "github.com/mailgun/godebug/lib"

var unnamed_input_in_goScope = godebug.EnteringNewScope()

func main() {
	godebug.Line()
	foo(3)
}

func foo(int) string {
	var godebugInput1 int
	var godebugResult1 string
	if !godebug.EnterFunc(func() {
		godebugResult1 = foo(godebugInput1)
	}) {
		return godebugResult1
	}
	defer godebug.ExitFunc()
	godebug.Line()
	return "hello"
}
