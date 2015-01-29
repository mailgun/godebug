package main

import "github.com/mailgun/godebug/lib"

var method_in_goScope = godebug.EnteringNewScope()

type Foo int

func (f Foo) Double() Foo {
	var godebugResult1 Foo
	ctx, ok := godebug.EnterFunc(func() {
		godebugResult1 = f.Double()
	})
	if !ok {
		return godebugResult1
	}
	defer godebug.ExitFunc()
	godebugScope := method_in_goScope.EnteringNewChildScope()
	godebugScope.Declare("f", &f)
	godebug.Line(ctx, godebugScope)
	return f * 2
}

func (Foo) Seven() Foo {
	var godebugResult1 Foo
	var godebugReceiver Foo
	ctx, ok := godebug.EnterFunc(func() {
		godebugResult1 = godebugReceiver.Seven()
	})
	if !ok {
		return godebugResult1
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
