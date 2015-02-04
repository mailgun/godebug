package main

import "github.com/mailgun/godebug/lib"

var variadic_in_goScope = godebug.EnteringNewScope()

func Varargs(i ...int) int {
	var result1 int
	ctx, ok := godebug.EnterFunc(func() {
		result1 = Varargs(i...)
	})
	if !ok {
		return result1
	}
	defer godebug.ExitFunc()
	scope := variadic_in_goScope.EnteringNewChildScope()
	scope.Declare("i", &i)
	godebug.Line(ctx, scope)
	return 6
}

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.Line(ctx, variadic_in_goScope)
	Varargs(1, 2, 3, 4)
}
