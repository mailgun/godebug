package main

import (
	"fmt"

	"github.com/mailgun/godebug/lib"
)

var example_in_goScope = godebug.EnteringNewScope()

func main() {
	ctx, ok := godebug.EnterFunc(func() {
		main()
	})
	if !ok {
		return
	}
	godebug.Line(ctx, example_in_goScope)
	x := mul(1, 2)
	godebugScope := example_in_goScope.EnteringNewChildScope()
	godebugScope.Declare("x", &x)
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, godebugScope)
	x = mul(x, x)
	godebug.Line(ctx, godebugScope)
	if x == 4 {
		godebug.Line(ctx, godebugScope)
		fmt.Println("It works! x == 4.")
	} else if n := func() int {
		godebug.ElseIfSimpleStmt(ctx, godebugScope, "} else if n := 2; n == 3 {")
		return 2
	}(); func() bool {
		godebug.ElseIfExpr(ctx, godebugScope, "} else if n := 2; n == 3 {")
		return n == 3
	}() {
		godebug.Line(ctx, godebugScope)
		fmt.Println("Math is broken. Ah!")
	} else {
		godebug.SLine(ctx, godebugScope, "} else {")
		godebug.Line(ctx, godebugScope)
		fmt.Println("What's going on? x ==", x)
	}
}

func add(n, m int) int {
	var godebugResult1 int
	ctx, ok := godebug.EnterFunc(func() {
		godebugResult1 = add(n, m)
	})
	if !ok {
		return godebugResult1
	}
	defer godebug.ExitFunc()
	godebugScope := example_in_goScope.EnteringNewChildScope()
	godebugScope.Declare("n", &n, "m", &m)
	godebug.Line(ctx, godebugScope)
	if n == 0 {
		godebug.Line(ctx, godebugScope)
		return m
	}
	godebug.Line(ctx, godebugScope)
	if m == 0 {
		godebug.Line(ctx, godebugScope)
		return n
	}
	godebug.Line(ctx, godebugScope)
	return n + m
}

func mul(n, m int) int {
	var godebugResult1 int
	ctx, ok := godebug.EnterFunc(func() {
		godebugResult1 = mul(n, m)
	})
	if !ok {
		return godebugResult1
	}
	defer godebug.ExitFunc()
	godebugScope := example_in_goScope.EnteringNewChildScope()
	godebugScope.Declare("n", &n, "m", &m)
	godebug.Line(ctx, godebugScope)
	var x int
	godebugScope.Declare("x", &x)
	godebug.Line(ctx, godebugScope)
	for i := 0; i < m; i++ {
		godebug.Line(ctx, godebugScope)
		x = add(x, m)
		godebug.SLine(ctx, godebugScope, "for i := 0; i < m; i++ {")
	}
	godebug.Line(ctx, godebugScope)
	return x
}
