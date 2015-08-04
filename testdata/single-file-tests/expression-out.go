package main

import "github.com/mailgun/godebug/lib"

var expression_in_go_scope = godebug.EnteringNewFile(main_pkg_scope, expression_in_go_contents)

type Foo struct {
	A int
	B []string
}

func plusTwo(x int) (int, string) {
	var result1 int
	var result2 string
	ctx, ok := godebug.EnterFunc(func() {
		result1, result2 = plusTwo(x)
	})
	if !ok {
		return result1, result2
	}
	defer godebug.ExitFunc(ctx)
	scope := expression_in_go_scope.EnteringNewChildScope()
	scope.Declare("x", &x)
	godebug.Line(ctx, scope, 9)
	return x + 2, "done"
}

const c = 42

var v = []int{1, 2, 3}

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.Line(ctx, expression_in_go_scope, 17)
	f := Foo{
		A: 12,
		B: []string{"hello", "world"},
	}
	scope := expression_in_go_scope.EnteringNewChildScope()
	scope.Declare("f", &f)
	godebug.Line(ctx, scope, 21)

	_ = f
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, scope, 22)

}

var expression_in_go_contents = `package main

type Foo struct {
	A int
	B []string
}

func plusTwo(x int) (int, string) {
	return x + 2, "done"
}

const c = 42

var v = []int{1, 2, 3}

func main() {
	f := Foo{
		A: 12,
		B: []string{"hello", "world"},
	}
	_ = f
	_ = "breakpoint"
}
`


var main_pkg_scope = &godebug.Scope{}

func init() {
	main_pkg_scope.Vars = map[string]interface{}{
		"v": &v,
	}
	main_pkg_scope.Consts = map[string]interface{}{
		"c": c,
	}
	main_pkg_scope.Funcs = map[string]interface{}{
		"plusTwo": plusTwo,
		"main": main,
	}
}