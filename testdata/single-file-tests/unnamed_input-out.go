package main

import "github.com/mailgun/godebug/lib"

var unnamed_input_in_go_scope = godebug.EnteringNewScope(unnamed_input_in_go_contents)

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.Line(ctx, unnamed_input_in_go_scope, 4)
	foo(3, 3)
}

func foo(int, int) (string, error) {
	var input1 int
	var input2 int
	var result1 string
	var result2 error
	ctx, ok := godebug.EnterFunc(func() {
		result1, result2 = foo(input1, input2)
	})
	if !ok {
		return result1, result2
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, unnamed_input_in_go_scope, 8)
	return "hello", nil
}

var unnamed_input_in_go_contents = `package main

func main() {
	foo(3, 3)
}

func foo(int, int) (string, error) {
	return "hello", nil
}
`
