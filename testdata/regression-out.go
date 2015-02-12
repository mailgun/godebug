package main

import "github.com/mailgun/godebug/lib"

var regression_in_goScope = godebug.EnteringNewScope()

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.Line(ctx, regression_in_goScope)

	foo := func(i int) int {
		var result1 int
		fn := func(ctx *godebug.Context) {
			result1 = func() int {
				scope := regression_in_goScope.EnteringNewChildScope()
				scope.Declare("i", &i)
				godebug.Line(ctx, scope)
				return i
			}()
		}
		if ctx, ok := godebug.EnterFuncLit(fn); ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
		return result1
	}(3)
	scope := regression_in_goScope.EnteringNewChildScope()
	scope.Declare("foo", &foo)
	godebug.Line(ctx, scope)

	_ = foo
	{
		scope := scope.EnteringNewChildScope()

		for _, s := range []string{"foo"} {
			godebug.SLine(ctx, scope, "for _, s := range []string{\"foo\"} {")
			scope.Declare("s", &s)
			godebug.Line(ctx, scope)
			_ = s
		}
		godebug.SLine(ctx, scope, "for _, s := range []string{\"foo\"} {")
	}
	godebug.Line(ctx, scope)

	c := make(chan bool)
	scope.Declare("c", &c)
	godebug.Line(ctx, scope)
	go func() {
		fn := func(ctx *godebug.Context) {
			godebug.Line(ctx, scope)
			c <- true
		}
		if ctx, ok := godebug.EnterFuncLit(fn); ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
	}()
	godebug.Line(ctx, scope)
	<-c
	godebug.Line(ctx, scope)

	defer println("Hello")
	defer godebug.SLine(ctx, scope, "<Running deferred function>: defer println(\"Hello\")")
	godebug.Line(ctx, scope)

	if false {
	} else if s := func() string {
		godebug.ElseIfSimpleStmt(ctx, scope, "} else if s := \"hello\"; s == \"hello\" {")
		return "hello"
	}(); func() bool {
		godebug.ElseIfExpr(ctx, scope, "} else if s := \"hello\"; s == \"hello\" {")
		return s == "hello"
	}() {
		godebug.Line(ctx, scope)
		println(s)
	}
}
