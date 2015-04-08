package main
import (
	"fmt"
	"github.com/mailgun/godebug/lib"
)
var switch_in_go_scope = godebug.EnteringNewScope(switch_in_go_contents)
func foo() interface{} {
	var result1 interface{}
	ctx, ok := godebug.EnterFunc(func() {
		result1 = foo()
	})
	if !ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, switch_in_go_scope, 10)
	return "hi"
}
func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, switch_in_go_scope, 16)
	switch {
	case godebug.Case(ctx, switch_in_go_scope, 17):
		fallthrough
	case false:
		godebug.Line(ctx, switch_in_go_scope, 18)
		fmt.Println("false")
	case godebug.Case(ctx, switch_in_go_scope, 19):
		fallthrough
	case true:
		godebug.Line(ctx, switch_in_go_scope, 20)
		fmt.Println("true")
	}
	godebug.Line(ctx, switch_in_go_scope, 23)
	i := 3
	scope := switch_in_go_scope.EnteringNewChildScope()
	scope.Declare("i", &i)
	godebug.Line(ctx, scope, 25)
	switch i {
	case godebug.Case(ctx, scope, 26):
		fallthrough
	case foo():
	default:
		godebug.Line(ctx, scope, 27)
	case godebug.Case(ctx, scope, 28):
		fallthrough
	case 5, 4, 1:
	case godebug.Case(ctx, scope, 29):
		fallthrough
	case 2:
	}
	godebug.Line(ctx, scope, 32)
	var ifc interface{} = i
	scope.Declare("ifc", &ifc)
	godebug.Line(ctx, scope, 34)
	switch ifc.(type) {
	case string:
		godebug.Line(ctx, scope, 35)
	case bool:
		godebug.Line(ctx, scope, 36)
	}
	{
		godebug.Line(ctx, scope, 39)
		b := 2
		scope := scope.EnteringNewChildScope()
		scope.Declare("b", &b)
		switch b == 6 {
		case godebug.Case(ctx, scope, 40):
			fallthrough
		case true:
		case godebug.Case(ctx, scope, 41):
			fallthrough
		case false:
		}
	}
	godebug.Line(ctx, scope, 44)
	switch b := ifc; i := ifc.(type) {
	case string:
		godebug.Line(ctx, scope, 45)
	case int:
		godebug.Line(ctx, scope, 46)
	default:
		godebug.Line(ctx, scope, 47)
		godebug.Line(ctx, scope, 48)
		_, _ = i, b
	}
}

var switch_in_go_contents = `package main

import (
	"fmt"

	"github.com/mailgun/godebug/lib"
)

func foo() interface{} {
	return "hi"
}

func main() {
	godebug.SetTrace()

	switch {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	}

	i := 3

	switch i {
	case foo():
	default:
	case 5, 4, 1:
	case 2:
	}

	var ifc interface{} = i

	switch ifc.(type) {
	case string:
	case bool:
	}

	switch b := 2; b == 6 {
	case true:
	case false:
	}

	switch b := ifc; i := ifc.(type) {
	case string:
	case int:
	default:
		_, _ = i, b
	}
}
`
