package main

import "github.com/mailgun/godebug/lib"

var struct_in_go_scope = godebug.EnteringNewFile(main_pkg_scope, struct_in_go_contents)

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.Line(ctx, struct_in_go_scope, 4)
	type myType struct {
		A int
		B string
		C bool
		d int
	}
	godebug.Line(ctx, struct_in_go_scope, 10)
	var v myType
	scope := struct_in_go_scope.EnteringNewChildScope()
	scope.Declare("v", &v)
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, scope, 11)
	godebug.Line(ctx, scope, 12)

	_ = v
}

var struct_in_go_contents = `package main

func main() {
	type myType struct {
		A int
		B string
		C bool
		d int
	}
	var v myType
	_ = "breakpoint"
	_ = v
}
`


var main_pkg_scope = &godebug.Scope{}

func init() {
	main_pkg_scope.Vars = map[string]interface{}{
	}
	main_pkg_scope.Consts = map[string]interface{}{
	}
	main_pkg_scope.Funcs = map[string]interface{}{
		"main": main,
	}
}