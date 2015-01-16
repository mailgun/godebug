package main

import (
	"fmt"

	"github.com/bradfitz/iter"
	"github.com/jeremyschlatter/godebug"
)

func main() {
	godebug.Line()
	x := mul(1, 2)
	godebug.RecordVars(&x, "x")
	defer godebug.OutOfScope("x")
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
	godebug.RecordVars(&n, "n", &m, "m")
	defer godebug.OutOfScope("n", "m")
	godebug.Line()
	var r int
	godebug.RecordVars(&r, "r")
	defer godebug.OutOfScope("r")
	godebug.Line()
	for range iter.N(m) {
		godebug.Line()
		r += m
		godebug.SLine("for range iter.N(m) {")
	}
	godebug.Line()
	return r
}
