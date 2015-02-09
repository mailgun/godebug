package main

import "github.com/mailgun/godebug/lib"

var regression_in_goScope = godebug.EnteringNewScope()

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	{
		scope := regression_in_goScope.EnteringNewChildScope()

		for _, s := range []string{"foo"} {
			godebug.SLine(ctx, scope, "for _, s := range []string{\"foo\"} {")
			scope.Declare("s", &s)
			godebug.Line(ctx, regression_in_goScope)
			_ = s
		}
		godebug.SLine(ctx, scope, "for _, s := range []string{\"foo\"} {")
	}
	godebug.Line(ctx, regression_in_goScope)

	c := make(chan bool)
	scope := regression_in_goScope.EnteringNewChildScope()
	scope.Declare("c", &c)
	godebug.Line(ctx, scope)
	go func() {
		fn := func(ctx *godebug.Context) {
			godebug.Line(ctx, scope)
			c <- true
		}
		if ctx, ok := godebug.EnterFuncLit(fn); ok {
			fn(ctx)
		}
		godebug.ExitFunc()
	}()
	godebug.Line(ctx, scope)
	<-c
}
