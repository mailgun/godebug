package main

import (
	godebug "fmt"
	_godebug "github.com/mailgun/godebug/lib"
)

var _name_conflicts_in_goScope = _godebug.EnteringNewScope()

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
	defer _godebug.ExitFunc()
	_godebug.Line(_ctx, _name_conflicts_in_goScope)
	var fn, ok, _ok, ctx, result1, input1, receiver, name_conflicts_in_goScope, scope int
	__scope := _name_conflicts_in_goScope.EnteringNewChildScope()
	__scope.Declare("fn", &fn, "ok", &ok, "_ok", &_ok, "ctx", &ctx, "result1", &result1, "input1", &input1, "receiver", &receiver, "name_conflicts_in_goScope", &name_conflicts_in_goScope, "scope", &scope)
	_godebug.Line(_ctx, __scope)
	godebug.Println(fn, ok, _ok, ctx, result1, input1, receiver, name_conflicts_in_goScope, scope, _scope)
	_godebug.Line(_ctx, __scope)
	return 3
}

var f = func() {
	fn := func(_ctx *_godebug.Context) {
		_godebug.Line(_ctx, _name_conflicts_in_goScope)
		var fn, ok, _ok, ctx, result1, input1, receiver, name_conflicts_in_goScope, scope int
		__scope := _name_conflicts_in_goScope.EnteringNewChildScope()
		__scope.Declare("fn", &fn, "ok", &ok, "_ok", &_ok, "ctx", &ctx, "result1", &result1, "input1", &input1, "receiver", &receiver, "name_conflicts_in_goScope", &name_conflicts_in_goScope, "scope", &scope)
		_godebug.Line(_ctx, __scope)
		godebug.Println(fn, ok, _ok, ctx, result1, input1, receiver, name_conflicts_in_goScope, scope, _scope)
	}
	if _ctx, __ok := _godebug.EnterFuncLit(fn); __ok {
		fn(_ctx)
	}
	_godebug.ExitFunc()
}

var _scope = 7

func main() {
	_ctx, __ok := _godebug.EnterFunc(main)
	if !__ok {
		return
	}
	_godebug.Line(_ctx, _name_conflicts_in_goScope)
	f()
	_godebug.Line(_ctx, _name_conflicts_in_goScope)
	foo := "hello"
	__scope := _name_conflicts_in_goScope.EnteringNewChildScope()
	__scope.Declare("foo", &foo)
	_godebug.Line(_ctx, __scope)
	{
		_godebug.Line(_ctx, __scope)
		scope := 3
		__scope := __scope.EnteringNewChildScope()
		__scope.Declare("scope", &scope)
		_godebug.Line(_ctx, __scope)
		{
			_godebug.Line(_ctx, __scope)
			godebug.Println(2)
		}
		_godebug.Line(_ctx, __scope)
		godebug.Println(scope)
	}
	_godebug.Line(_ctx, __scope)
	godebug.Println(foo)
}
