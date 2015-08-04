package main

import (
	godebug "fmt"
	_godebug "github.com/mailgun/godebug/lib"
)

var name_conflicts_in_go_scope = _godebug.EnteringNewFile(main_pkg_scope, name_conflicts_in_go_contents)

type Foo int

func (Foo) DoStuff(int) int {
	var _input1 int
	var _result1 int
	var _receiver Foo
	_ctx, __ok := _godebug.EnterFunc(func() {
		_result1 = _receiver.DoStuff(_input1)
	})
	if !__ok {
		return _result1
	}
	defer _godebug.ExitFunc(_ctx)
	_godebug.Line(_ctx, name_conflicts_in_go_scope, 8)
	var fn, ok, _ok, ctx, result1, input1, receiver, name_conflicts_in_goScope, scope int
	__scope := name_conflicts_in_go_scope.EnteringNewChildScope()
	__scope.Declare("fn", &fn, "ok", &ok, "_ok", &_ok, "ctx", &ctx, "result1", &result1, "input1", &input1, "receiver", &receiver, "name_conflicts_in_goScope", &name_conflicts_in_goScope, "scope", &scope)
	_godebug.Line(_ctx, __scope, 9)
	godebug.Println(fn, ok, _ok, ctx, result1, input1, receiver, name_conflicts_in_goScope, scope, _scope)
	_godebug.Line(_ctx, __scope, 10)
	return 3
}

var f = func() {
	fn := func(_ctx *_godebug.Context) {
		_godebug.Line(_ctx, name_conflicts_in_go_scope, 14)
		var fn, ok, _ok, ctx, result1, input1, receiver, name_conflicts_in_goScope, scope int
		__scope := name_conflicts_in_go_scope.EnteringNewChildScope()
		__scope.Declare("fn", &fn, "ok", &ok, "_ok", &_ok, "ctx", &ctx, "result1", &result1, "input1", &input1, "receiver", &receiver, "name_conflicts_in_goScope", &name_conflicts_in_goScope, "scope", &scope)
		_godebug.Line(_ctx, __scope, 15)
		godebug.Println(fn, ok, _ok, ctx, result1, input1, receiver, name_conflicts_in_goScope, scope, _scope)
	}
	if _ctx, __ok := _godebug.EnterFuncLit(fn); __ok {
		defer _godebug.ExitFunc(_ctx)
		fn(_ctx)
	}
}

var _scope = 7

func main() {
	_ctx, __ok := _godebug.EnterFunc(main)
	if !__ok {
		return
	}
	_godebug.Line(_ctx, name_conflicts_in_go_scope, 21)
	f()
	_godebug.Line(_ctx, name_conflicts_in_go_scope, 22)
	foo := "hello"
	__scope := name_conflicts_in_go_scope.EnteringNewChildScope()
	__scope.Declare("foo", &foo)
	{
		_godebug.Line(_ctx, __scope, 24)
		scope := 3
		__scope := __scope.EnteringNewChildScope()
		__scope.Declare("scope", &scope)
		{
			_godebug.Line(_ctx, __scope, 26)
			godebug.Println(2)
		}
		_godebug.Line(_ctx, __scope, 28)
		godebug.Println(scope)
	}
	_godebug.Line(_ctx, __scope, 30)
	godebug.Println(foo)
}

var name_conflicts_in_go_contents = `package main

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
`


var main_pkg_scope = &_godebug.Scope{}

func init() {
	main_pkg_scope.Vars = map[string]interface{}{
		"f": &f,
		"_scope": &_scope,
	}
	main_pkg_scope.Consts = map[string]interface{}{
	}
	main_pkg_scope.Funcs = map[string]interface{}{
		"main": main,
	}
}