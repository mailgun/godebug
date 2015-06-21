package main

import (
	"fmt"
	"github.com/mailgun/godebug/lib"
)

var multiple_var_declaration_in_go_scope = godebug.EnteringNewScope(multiple_var_declaration_in_go_contents)

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.Line(ctx, multiple_var_declaration_in_go_scope, 8)
	func() {
		fn := func(ctx *godebug.Context) {
			godebug.Line(ctx, multiple_var_declaration_in_go_scope, 9)
			var (
				x   int
				y   int
				err error
			)
			scope := multiple_var_declaration_in_go_scope.EnteringNewChildScope()
			scope.Declare("x", &x, "y", &y, "err", &err)
			godebug.Line(ctx, scope, 14)
			x = 2
			godebug.Line(ctx, scope, 15)
			y = 3
			godebug.Line(ctx, scope, 16)
			err = nil
			godebug.Line(ctx, scope, 17)
			if err != nil {
				godebug.Line(ctx, scope, 18)
				fmt.Printf("%d\n", x+y)
			}
		}
		if ctx, ok := godebug.EnterFuncLit(fn); ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
	}()
}

var multiple_var_declaration_in_go_contents = `package main

import (
	"fmt"
)

func main() {
	func() {
		var (
			x   int
			y   int
			err error
		)
		x = 2
		y = 3
		err = nil
		if err != nil {
			fmt.Printf("%d\n", x+y)
		}
	}()
}
`
