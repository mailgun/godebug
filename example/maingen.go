package main

import (
	"fmt"

	"github.com/bradfitz/iter"
	"github.com/jeremyschlatter/godebug"
)

var main_goScope = godebug.EnteringNewScope()

func main() {
	godebugScope := main_goScope.EnteringNewChildScope()
	godebug.Line()
	x := mul(1, 2)
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
		fmt.Println("Math is broken. Run.")
	} else {
		godebug.SLine("} else {")
		godebug.Line()
		fmt.Println("What's going on? x ==", x)
	}
}

func add(n, m int) int {
	godebug.EnterFunc()
	defer godebug.ExitFunc()
	godebugScope := main_goScope.EnteringNewChildScope()
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
	godebug.EnterFunc()
	defer godebug.ExitFunc()
	godebugScope := main_goScope.EnteringNewChildScope()
	defer godebugScope.End()
	godebugScope.Declare("n", &n, "m", &m)
	godebug.Line()
	var x int
	godebugScope.Declare("x", &x)
	godebug.Line()
	for range iter.N(m) {
		godebug.Line()
		x = add(x, m)
		godebug.SLine("for range iter.N(m) {")
	}
	godebug.Line()
	return x
}
