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
		for s := range []string{"foo"} {
			godebug.SLine(ctx, scope, "for s := range []string{\"foo\"} {")
			scope.Declare("s", &s)
			godebug.Line(ctx, regression_in_goScope)
			_ = s
		}
		godebug.SLine(ctx, scope, "for s := range []string{\"foo\"} {")
	}
}
