package main

import "github.com/mailgun/godebug/lib"

var method_in_goScope = godebug.EnteringNewScope()

type Foo int

func (f Foo) Double() Foo {
	var result1 Foo
	ctx, ok := godebug.EnterFunc(func() {
		result1 = f.Double()
	})
	if !ok {
		return result1
	}
	defer godebug.ExitFunc()
	scope := method_in_goScope.EnteringNewChildScope()
	scope.Declare("f", &f)
	godebug.Line(ctx, scope)
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
	defer godebug.ExitFunc()
	godebug.Line(ctx, method_in_goScope)
	return Foo(7)
}

func main() {
	ctx, ok := godebug.EnterFunc(func() {
		main()
	})
	if !ok {
		return
	}
	godebug.Line(ctx, method_in_goScope)
	Foo(3).Double()
}
