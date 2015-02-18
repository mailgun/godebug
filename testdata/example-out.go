package main
import (
	"fmt"
	"github.com/mailgun/godebug/lib"
)
var example_in_go_scope = godebug.EnteringNewScope(example_in_go_contents)
func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.Line(ctx, example_in_go_scope, 10)
	x := mul(1, 2)
	scope := example_in_go_scope.EnteringNewChildScope()
	scope.Declare("x", &x)
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, scope, 12)
	x = mul(x, x)
	godebug.Line(ctx, scope, 13)
	if x == 4 {
		godebug.Line(ctx, scope, 14)
		fmt.Println("It works! x == 4.")
	} else if n := func() int {
		godebug.ElseIfSimpleStmt(ctx, scope, 15)
		return 2
	}(); func() bool {
		godebug.ElseIfExpr(ctx, scope, 15)
		return n == 3
	}() {
		godebug.Line(ctx, scope, 16)
		fmt.Println("Math is broken. Ah!")
	} else {
		godebug.Line(ctx, scope, 17)
		godebug.Line(ctx, scope, 18)
		fmt.Println("What's going on? x ==", x)
	}
}
func add(n, m int) int {
	var result1 int
	ctx, ok := godebug.EnterFunc(func() {
		result1 = add(n, m)
	})
	if !ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	scope := example_in_go_scope.EnteringNewChildScope()
	scope.Declare("n", &n, "m", &m)
	godebug.Line(ctx, scope, 23)
	if n == 0 {
		godebug.Line(ctx, scope, 24)
		return m
	}
	godebug.Line(ctx, scope, 26)
	if m == 0 {
		godebug.Line(ctx, scope, 27)
		return n
	}
	godebug.Line(ctx, scope, 29)
	return n + m
}
func mul(n, m int) int {
	var result1 int
	ctx, ok := godebug.EnterFunc(func() {
		result1 = mul(n, m)
	})
	if !ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	scope := example_in_go_scope.EnteringNewChildScope()
	scope.Declare("n", &n, "m", &m)
	godebug.Line(ctx, scope, 33)
	var x int
	scope.Declare("x", &x)
	{
		scope := scope.EnteringNewChildScope()
		for i := 0; i < m; i++ {
			godebug.Line(ctx, scope, 34)
			scope.Declare("i", &i)
			godebug.Line(ctx, scope, 35)
			x = add(x, m)
		}
		godebug.Line(ctx, scope, 34)
	}
	godebug.Line(ctx, scope, 37)
	return x
}

var example_in_go_contents = `package main

import (
	"fmt"

	"github.com/mailgun/godebug/lib"
)

func main() {
	x := mul(1, 2)
	godebug.SetTrace()
	x = mul(x, x)
	if x == 4 {
		fmt.Println("It works! x == 4.")
	} else if n := 2; n == 3 {
		fmt.Println("Math is broken. Ah!")
	} else {
		fmt.Println("What's going on? x ==", x)
	}
}

func add(n, m int) int {
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}
	return n + m
}

func mul(n, m int) int {
	var x int
	for i := 0; i < m; i++ {
		x = add(x, m)
	}
	return x
}
`
