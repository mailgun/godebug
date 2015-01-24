package main

import (
	"fmt"

	"github.com/mailgun/godebug/lib"
)

var example_in_goScope = godebug.EnteringNewScope()

func main() {
	godebug.Line()
	x := mul(1, 2)
	godebugScope := example_in_goScope.EnteringNewChildScope()
	defer godebugScope.End()
	godebugScope.Declare("x", &x)
	godebug.SetTrace()
	godebug.Line()
	x = mul(x, x)
	godebug.Line()
	if x == 4 {
		godebug.Line()
		fmt.Println("It works! x == 4.")
	} else if n := func() int {
		godebug.ElseIfSimpleStmt("} else if n := 2; n == 3 {")
		return 2
	}(); func() bool {
		godebug.ElseIfExpr("} else if n := 2; n == 3 {")
		return n == 3
	}() {
		godebug.Line()
		fmt.Println("Math is broken. Ah!")
	} else {
		godebug.SLine("} else {")
		godebug.Line()
		fmt.Println("What's going on? x ==", x)
	}
}

func add(n, m int) int {
	var godebugResult1 int
	if !godebug.EnterFunc(func() {
		godebugResult1 = add(n, m)
	}) {
		return godebugResult1
	}

	defer godebug.ExitFunc()
	godebugScope := example_in_goScope.EnteringNewChildScope()
	defer godebugScope.End()
	godebugScope.Declare("n", &n, "m", &m)
	godebug.Line()
	if n == 0 {
		godebug.Line()
		return m
	}
	godebug.Line()
	if m == 0 {
		godebug.Line()
		return n
	}
	godebug.Line()
	return n + m
}

func mul(n, m int) int {
	var godebugResult1 int
	if !godebug.EnterFunc(func() {
		godebugResult1 = mul(n, m)
	}) {
		return godebugResult1
	}

	defer godebug.ExitFunc()
	godebugScope := example_in_goScope.EnteringNewChildScope()
	defer godebugScope.End()
	godebugScope.Declare("n", &n, "m", &m)
	godebug.Line()
	var x int
	godebugScope.Declare("x", &x)
	godebug.Line()
	for i := 0; i < m; i++ {
		godebug.Line()
		x = add(x, m)
		godebug.SLine("for i := 0; i < m; i++ {")
	}
	godebug.Line()
	return x
}
