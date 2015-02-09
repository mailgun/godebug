package main

import "github.com/mailgun/godebug/lib"

var init_in_goScope = godebug.EnteringNewScope()

func init() {
	a = 5
}

var a = 3

type Foo int

func (f *Foo) init() {
	ctx, ok := godebug.EnterFunc(f.init)
	if !ok {
		return
	}
	defer godebug.ExitFunc(ctx)
	scope := init_in_goScope.EnteringNewChildScope()
	scope.Declare("f", &f)
	godebug.Line(ctx, scope)
	*f = 1337
}

func main() {
}
