package main

import godebug "fmt"

type Foo int

func (Foo) DoStuff(int) int {
	var fn, ok, _ok, ctx, result1, input1, receiver, name_conflicts_in_goScope, scope int
	godebug.Println(fn, ok, _ok, ctx, result1, input1, receiver, name_conflicts_in_goScope, scope, _scope)
	return 3
}

var f = func() {
	var fn, ok, _ok, ctx, result1, input1, receiver, name_conflicts_in_goScope, scope int
	godebug.Println(fn, ok, _ok, ctx, result1, input1, receiver, name_conflicts_in_goScope, scope, _scope)
}

var _scope = 7

func main() {
	f()
	foo := "hello"
	{
		scope := 3
		{
			godebug.Println(2)
		}
		godebug.Println(scope)
	}
	godebug.Println(foo)
}
