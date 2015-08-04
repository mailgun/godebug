package main

import "github.com/mailgun/godebug/lib"

var variadic_in_go_scope = godebug.EnteringNewFile(main_pkg_scope, variadic_in_go_contents)

func Varargs(i ...int) int {
	var result1 int
	ctx, ok := godebug.EnterFunc(func() {
		result1 = Varargs(i...)
	})
	if !ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	scope := variadic_in_go_scope.EnteringNewChildScope()
	scope.Declare("i", &i)
	godebug.Line(ctx, scope, 4)
	return 6
}

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.Line(ctx, variadic_in_go_scope, 8)
	Varargs(1, 2, 3, 4)
}

var variadic_in_go_contents = `package main

func Varargs(i ...int) int {
	return 6
}

func main() {
	Varargs(1, 2, 3, 4)
}
`


var main_pkg_scope = &godebug.Scope{}

func init() {
	main_pkg_scope.Vars = map[string]interface{}{
	}
	main_pkg_scope.Consts = map[string]interface{}{
	}
	main_pkg_scope.Funcs = map[string]interface{}{
		"Varargs": Varargs,
		"main": main,
	}
}