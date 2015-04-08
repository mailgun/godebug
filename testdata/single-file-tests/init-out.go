package main

import "github.com/mailgun/godebug/lib"

var init_in_go_scope = godebug.EnteringNewScope(init_in_go_contents)

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
	scope := init_in_go_scope.EnteringNewChildScope()
	scope.Declare("f", &f)
	godebug.Line(ctx, scope, 14)
	*f = 1337
}

func main() {
}

var init_in_go_contents = `package main

// Don't generate debug calls for init.
func init() {
	a = 5
}

var a = 3

type Foo int

// Do generate debug calls for methods named init.
func (f *Foo) init() {
	*f = 1337
}

func main() {
}
`
