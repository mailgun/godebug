package main

import (
	"fmt"

	"github.com/mailgun/godebug/lib"
)

var example_in_goScope = godebug.EnteringNewScope()

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.Line(ctx, example_in_goScope)
	x := mul(1, 2)
	scope := example_in_goScope.EnteringNewChildScope()
	scope.Declare("x", &x)
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, scope)
	x = mul(x, x)
	godebug.Line(ctx, scope)
	if x == 4 {
		godebug.Line(ctx, scope)
		fmt.Println("It works! x == 4.")
	} else if n := func() int {
		godebug.ElseIfSimpleStmt(ctx, scope, "} else if n := 2; n == 3 {")
		return 2
	}(); func() bool {
		godebug.ElseIfExpr(ctx, scope, "} else if n := 2; n == 3 {")
		return n == 3
	}() {
		godebug.Line(ctx, scope)
		fmt.Println("Math is broken. Ah!")
	} else {
		godebug.SLine(ctx, scope, "} else {")
		godebug.Line(ctx, scope)
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
	defer godebug.ExitFunc()
	scope := example_in_goScope.EnteringNewChildScope()
	scope.Declare("n", &n, "m", &m)
	godebug.Line(ctx, scope)
	if n == 0 {
		godebug.Line(ctx, scope)
		return m
	}
	godebug.Line(ctx, scope)
	if m == 0 {
		godebug.Line(ctx, scope)
		return n
	}
	godebug.Line(ctx, scope)
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
	defer godebug.ExitFunc()
	scope := example_in_goScope.EnteringNewChildScope()
	scope.Declare("n", &n, "m", &m)
	godebug.Line(ctx, scope)
	var x int
	scope.Declare("x", &x)
	godebug.Line(ctx, scope)
	for i := 0; i < m; i++ {
		godebug.Line(ctx, scope)
		x = add(x, m)
		godebug.SLine(ctx, scope, "for i := 0; i < m; i++ {")
	}
	godebug.Line(ctx, scope)
	return x
}
