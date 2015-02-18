package main
import "github.com/mailgun/godebug/lib"
var regression_in_go_scope = godebug.EnteringNewScope(regression_in_go_contents)
func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.Line(ctx, regression_in_go_scope, 5)
	foo := func(i int) int {
		var result1 int
		fn := func(ctx *godebug.Context) {
			result1 = func() int {
				scope := regression_in_go_scope.EnteringNewChildScope()
				scope.Declare("i", &i)
				godebug.Line(ctx, scope, 6)
				return i
			}()
		}
		if ctx, ok := godebug.EnterFuncLit(fn); ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
		return result1
	}(3)
	scope := regression_in_go_scope.EnteringNewChildScope()
	scope.Declare("foo", &foo)
	godebug.Line(ctx, scope, 8)
	_ = foo
	{
		scope := scope.EnteringNewChildScope()
		for _, s := range []string{"foo"} {
			godebug.Line(ctx, scope, 12)
			scope.Declare("s", &s)
			godebug.Line(ctx, scope, 13)
			_ = s
		}
		godebug.Line(ctx, scope, 12)
	}
	godebug.Line(ctx, scope, 17)
	c := make(chan bool)
	scope.Declare("c", &c)
	godebug.Line(ctx, scope, 18)
	go func() {
		fn := func(ctx *godebug.Context) {
			godebug.Line(ctx, scope, 19)
			c <- true
		}
		if ctx, ok := godebug.EnterFuncLit(fn); ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
	}()
	godebug.Line(ctx, scope, 21)
	<-c
	godebug.Line(ctx, scope, 24)
	defer println("Hello")
	defer godebug.Defer(ctx, scope, 24)
	godebug.Line(ctx, scope, 27)
	if false {
	} else if s := func() string {
		godebug.ElseIfSimpleStmt(ctx, scope, 28)
		return "hello"
	}(); func() bool {
		godebug.ElseIfExpr(ctx, scope, 28)
		return s == "hello"
	}() {
		godebug.Line(ctx, scope, 29)
		println(s)
	}
}

var regression_in_go_contents = `package main

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
}
`
