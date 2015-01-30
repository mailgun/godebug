package main

import "github.com/mailgun/godebug/lib"

var init_in_goScope = godebug.EnteringNewScope()

func init() {
	a = 5
}

var a = 3

type Foo int

func (f *Foo) init() {
	ctx, ok := godebug.EnterFunc(func() {
		f.init()
	})
	if !ok {
		return
	}
	defer godebug.ExitFunc()
	godebugScope := init_in_goScope.EnteringNewChildScope()
	godebugScope.Declare("f", &f)
	godebug.Line(ctx, godebugScope)
	*f = 1337
}

func main() {
}
