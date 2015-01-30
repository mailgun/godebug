package main

import "github.com/mailgun/godebug/lib"

var variadic_in_goScope = godebug.EnteringNewScope()

func Varargs(i ...int) int {
	var godebugResult1 int
	ctx, ok := godebug.EnterFunc(func() {
		godebugResult1 = Varargs(i...)
	})
	if !ok {
		return godebugResult1
	}
	defer godebug.ExitFunc()
	godebugScope := variadic_in_goScope.EnteringNewChildScope()
	godebugScope.Declare("i", &i)
	godebug.Line(ctx, godebugScope)
	return 6
}

func main() {
	ctx, ok := godebug.EnterFunc(func() {
		main()
	})
	if !ok {
		return
	}
	godebug.Line(ctx, variadic_in_goScope)
	Varargs(1, 2, 3, 4)
}
