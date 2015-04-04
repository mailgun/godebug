package main
import "github.com/mailgun/godebug/lib"
var regression_in_go_scope = godebug.EnteringNewScope(regression_in_go_contents)
func main() {
	ctx, _ok := godebug.EnterFunc(main)
	if !_ok {
		return
	}
	godebug.Line(ctx, regression_in_go_scope, 7)
	foo := func(i int) int {
		var result1 int
		fn := func(ctx *godebug.Context) {
			result1 = func() int {
				scope := regression_in_go_scope.EnteringNewChildScope()
				scope.Declare("i", &i)
				godebug.Line(ctx, scope, 8)
				return i
			}()
		}
		if ctx, _ok := godebug.EnterFuncLit(fn); _ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
		return result1
	}(3)
	scope := regression_in_go_scope.EnteringNewChildScope()
	scope.Declare("foo", &foo)
	godebug.Line(ctx, scope, 10)
	_ = foo
	{
		scope := scope.EnteringNewChildScope()
		for _, s := range []string{"foo"} {
			godebug.Line(ctx, scope, 14)
			scope.Declare("s", &s)
			godebug.Line(ctx, scope, 15)
			_ = s
		}
		godebug.Line(ctx, scope, 14)
	}
	godebug.Line(ctx, scope, 19)
	c := make(chan bool)
	scope.Declare("c", &c)
	godebug.Line(ctx, scope, 20)
	go func() {
		fn := func(ctx *godebug.Context) {
			godebug.Line(ctx, scope, 21)
			c <- true
		}
		if ctx, _ok := godebug.EnterFuncLit(fn); _ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
	}()
	godebug.Line(ctx, scope, 23)
	<-c
	godebug.Line(ctx, scope, 26)
	defer println("Hello")
	defer godebug.Defer(ctx, scope, 26)
	godebug.Line(ctx, scope, 29)
	if false {
	} else {
		godebug.ElseIfSimpleStmt(ctx, scope, 30)
		s := "hello"
		godebug.ElseIfExpr(ctx, scope, 30)
		if s == "hello" {
			godebug.Line(ctx, scope, 31)
			println(s)
		}
	}
	godebug.Line(ctx, scope, 35)
	m := map[string]int{"test": 5}
	scope.Declare("m", &m)
	godebug.Line(ctx, scope, 36)
	if false {
	} else {
		godebug.ElseIfSimpleStmt(ctx, scope, 37)
		_, ok := m["test"]
		godebug.ElseIfExpr(ctx, scope, 37)
		if ok {
			godebug.Line(ctx, scope, 38)
			println("test")
		}
	}
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, scope, 43)
	const n = 10
	scope.Constant("n", n)
	godebug.Line(ctx, scope, 44)
	_ = n
}
func _switch() int {
	var result1 int
	ctx, _ok := godebug.EnterFunc(func() {
		result1 = _switch()
	})
	if !_ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, regression_in_go_scope, 49)
	switch {
	case godebug.Case(ctx, regression_in_go_scope, 50):
		panic("impossible")
	case false:
		godebug.Line(ctx, regression_in_go_scope, 51)
		return 4
	default:
		godebug.Line(ctx, regression_in_go_scope, 52)
		godebug.Line(ctx, regression_in_go_scope, 53)
		return 5
	}
}
func _select() int {
	var result1 int
	ctx, _ok := godebug.EnterFunc(func() {
		result1 = _select()
	})
	if !_ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	godebug.Select(ctx, regression_in_go_scope, 59)
	select {
	case <-godebug.Comm(ctx, regression_in_go_scope, 60):
		panic("impossible")
	case <-make(chan bool):
		godebug.Line(ctx, regression_in_go_scope, 60)
		godebug.Line(ctx, regression_in_go_scope, 61)
		return 4
	default:
		godebug.Line(ctx, regression_in_go_scope, 62)
		godebug.Line(ctx, regression_in_go_scope, 63)
		return 5
	case <-godebug.EndSelect(ctx, regression_in_go_scope):
		panic("impossible")
	}
}

var regression_in_go_contents = `package main

import "github.com/mailgun/godebug/lib"

func main() {
	// Nested scope in the first declaration in a function.
	foo := func(i int) int {
		return i
	}(3)
	_ = foo

	// String literal in range statement.
	// Blank identifier in range statement.
	for _, s := range []string{"foo"} {
		_ = s
	}

	// go statement with function literal.
	c := make(chan bool)
	go func() {
		c <- true
	}()
	<-c

	// String literal in defer statement.
	defer println("Hello")

	// String literal in else-if statement.
	if false {
	} else if s := "hello"; s == "hello" {
		println(s)
	}

	// Comma-ok in else-if
	m := map[string]int{"test": 5}
	if false {
	} else if _, ok := m["test"]; ok {
		println("test")
	}

	// Constant declaration.
	godebug.SetTrace()
	const n = 10
	_ = n
}

func _switch() int {
	// Terminating switch statement in function with return value.
	switch {
	case false:
		return 4
	default:
		return 5
	}
}

func _select() int {
	// Terminating select statement in function with return value.
	select {
	case <-make(chan bool):
		return 4
	default:
		return 5
	}
}
`
