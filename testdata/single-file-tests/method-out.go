package main

import "github.com/mailgun/godebug/lib"

var method_in_go_scope = godebug.EnteringNewFile(main_pkg_scope, method_in_go_contents)

type Foo int

func (f Foo) Double() Foo {
	var result1 Foo
	ctx, ok := godebug.EnterFunc(func() {
		result1 = f.Double()
	})
	if !ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	scope := method_in_go_scope.EnteringNewChildScope()
	scope.Declare("f", &f)
	godebug.Line(ctx, scope, 6)
	return f * 2
}

func (Foo) Seven() Foo {
	var result1 Foo
	var receiver Foo
	ctx, ok := godebug.EnterFunc(func() {
		result1 = receiver.Seven()
	})
	if !ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, method_in_go_scope, 10)
	return Foo(7)
}

func (_ Foo) Bar() int {
	var result1 int
	var receiver Foo
	ctx, ok := godebug.EnterFunc(func() {
		result1 = receiver.Bar()
	})
	if !ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, method_in_go_scope, 14)
	return 0
}

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.Line(ctx, method_in_go_scope, 18)
	Foo(3).Double()
}

var method_in_go_contents = `package main

type Foo int

func (f Foo) Double() Foo {
	return f * 2
}

func (Foo) Seven() Foo {
	return Foo(7)
}

func (_ Foo) Bar() int {
	return 0
}

func main() {
	Foo(3).Double()
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