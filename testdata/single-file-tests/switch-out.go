package main

import (
	"fmt"
	"github.com/mailgun/godebug/lib"
)

var switch_in_go_scope = godebug.EnteringNewFile(main_pkg_scope, switch_in_go_contents)

func foo() interface{} {
	var result1 interface{}
	ctx, ok := godebug.EnterFunc(func() {
		result1 = foo()
	})
	if !ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, switch_in_go_scope, 6)
	return "hi"
}

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, switch_in_go_scope, 10)
	godebug.Line(ctx, switch_in_go_scope, 12)

	switch {
	case godebug.Case(ctx, switch_in_go_scope, 13):
		fallthrough
	case false:
		godebug.Line(ctx, switch_in_go_scope, 14)
		fmt.Println("false")
	case godebug.Case(ctx, switch_in_go_scope, 15):
		fallthrough
	case true:
		godebug.Line(ctx, switch_in_go_scope, 16)
		fmt.Println("true")
	}
	godebug.Line(ctx, switch_in_go_scope, 19)

	i := 3
	scope := switch_in_go_scope.EnteringNewChildScope()
	scope.Declare("i", &i)
	godebug.Line(ctx, scope, 21)

	switch i {
	case godebug.Case(ctx, scope, 22):
		fallthrough
	case foo():
	default:
		godebug.Line(ctx, scope, 23)
	case godebug.Case(ctx, scope, 24):
		fallthrough
	case 5, 4, 1:
	case godebug.Case(ctx, scope, 25):
		fallthrough
	case 2:
	}
	godebug.Line(ctx, scope, 28)

	var ifc interface{} = i
	scope.Declare("ifc", &ifc)
	godebug.Line(ctx, scope, 30)

	switch ifc.(type) {
	case string:
		godebug.Line(ctx, scope, 31)
	case bool:
		godebug.Line(ctx, scope, 32)
	}
	{
		godebug.Line(ctx, scope, 35)
		b := 2
		scope := scope.EnteringNewChildScope()
		scope.Declare("b", &b)
		switch b == 6 {
		case godebug.Case(ctx, scope, 36):
			fallthrough
		case true:
		case godebug.Case(ctx, scope, 37):
			fallthrough
		case false:
		}
	}
	godebug.Line(ctx, scope, 40)

	switch b := ifc; i := ifc.(type) {
	case string:
		godebug.Line(ctx, scope, 41)
	case int:
		godebug.Line(ctx, scope, 42)
	default:
		godebug.Line(ctx, scope, 43)
		godebug.Line(ctx, scope, 44)
		_, _ = i, b
	}
}

var switch_in_go_contents = `package main

import "fmt"

func foo() interface{} {
	return "hi"
}

func main() {
	_ = "breakpoint"

	switch {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	}

	i := 3

	switch i {
	case foo():
	default:
	case 5, 4, 1:
	case 2:
	}

	var ifc interface{} = i

	switch ifc.(type) {
	case string:
	case bool:
	}

	switch b := 2; b == 6 {
	case true:
	case false:
	}

	switch b := ifc; i := ifc.(type) {
	case string:
	case int:
	default:
		_, _ = i, b
	}
}
`


var main_pkg_scope = &godebug.Scope{}

func init() {
	main_pkg_scope.Vars = map[string]interface{}{
	}
	main_pkg_scope.Consts = map[string]interface{}{
	}
	main_pkg_scope.Funcs = map[string]interface{}{
		"foo": foo,
		"main": main,
	}
}