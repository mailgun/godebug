package main

import (
	"fmt"
	"github.com/mailgun/godebug/lib"
)

var example_in_go_scope = godebug.EnteringNewFile(main_pkg_scope, example_in_go_contents)

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.Line(ctx, example_in_go_scope, 6)
	x := mul(1, 2)
	scope := example_in_go_scope.EnteringNewChildScope()
	scope.Declare("x", &x)
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, scope, 7)
	godebug.Line(ctx, scope, 8)

	x = mul(x, x)
	godebug.Line(ctx, scope, 9)
	if x == 4 {
		godebug.Line(ctx, scope, 10)
		fmt.Println("It works! x == 4.")
	} else {
		godebug.ElseIfSimpleStmt(ctx, scope, 11)
		n := 2
		godebug.ElseIfExpr(ctx, scope, 11)
		if n == 3 {
			godebug.Line(ctx, scope, 12)
			fmt.Println("Math is broken. Ah!")
		} else {
			godebug.Line(ctx, scope, 13)
			godebug.Line(ctx, scope, 14)
			fmt.Println("What's going on? x ==", x)
		}
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
	godebug.Line(ctx, scope, 19)
	if n == 0 {
		godebug.Line(ctx, scope, 20)
		return m
	}
	godebug.Line(ctx, scope, 22)
	if m == 0 {
		godebug.Line(ctx, scope, 23)
		return n
	}
	godebug.Line(ctx, scope, 25)
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
	godebug.Line(ctx, scope, 29)
	var x int
	scope.Declare("x", &x)
	{
		scope := scope.EnteringNewChildScope()
		for i := 0; i < m; i++ {
			godebug.Line(ctx, scope, 30)
			scope.Declare("i", &i)
			godebug.Line(ctx, scope, 31)
			x = add(x, m)
		}
		godebug.Line(ctx, scope, 30)
	}
	godebug.Line(ctx, scope, 33)
	return x
}

var example_in_go_contents = `package main

import "fmt"

func main() {
	x := mul(1, 2)
	_ = "breakpoint"
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


var main_pkg_scope = &godebug.Scope{}

func init() {
	main_pkg_scope.Vars = map[string]interface{}{
	}
	main_pkg_scope.Consts = map[string]interface{}{
	}
	main_pkg_scope.Funcs = map[string]interface{}{
		"main": main,
		"add": add,
		"mul": mul,
	}
}