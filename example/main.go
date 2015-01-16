package main

import (
	"fmt"

	"github.com/bradfitz/iter"
)

func main() {
	Line()
	x := mul(1, 2)
	RecordVars(&x, "x")
	defer OutOfScope("x")
	SetTrace()
	Line()
	x = mul(x, x)
	Line()
	if x == 4 {
		Line()
		fmt.Println("It works! x == 4.")
	} else if n := func() int {
		ElseIfSimpleStmt("} else if n := 2; n == 3 {")
		return 2
	}(); func() bool {
		ElseIfExpr("} else if n := 2; n == 3 {")
		return n == 3
	}() {
		Line()
		fmt.Println("Math is broken. Run.")
	} else {
		SLine("} else {")
		Line()
		fmt.Println("What's going on? x ==", x)
	}
}

func add(n, m int) int {
	EnterFunc()
	defer ExitFunc()
	Line()
	if n == 0 {
		Line()
		return m
	}
	Line()
	if m == 0 {
		Line()
		return n
	}
	Line()
	return n + m
}

func mul(n, m int) int {
	EnterFunc()
	defer ExitFunc()
	RecordVars(&n, "n", &m, "m")
	defer OutOfScope("n", "m")
	Line()
	var r int
	RecordVars(&r, "r")
	defer OutOfScope("r")
	Line()
	for range iter.N(m) {
		Line()
		r += m
		SLine("for range iter.N(m) {")
	}
	Line()
	return r
}
