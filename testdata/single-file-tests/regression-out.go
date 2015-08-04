package main

import "github.com/mailgun/godebug/lib"

var regression_in_go_scope = godebug.EnteringNewFile(main_pkg_scope, regression_in_go_contents)

func main() {
	ctx, _ok := godebug.EnterFunc(main)
	if !_ok {
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
		if ctx, _ok := godebug.EnterFuncLit(fn); _ok {
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
		if ctx, _ok := godebug.EnterFuncLit(fn); _ok {
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
	} else {
		godebug.ElseIfSimpleStmt(ctx, scope, 28)
		s := "hello"
		godebug.ElseIfExpr(ctx, scope, 28)
		if s == "hello" {
			godebug.Line(ctx, scope, 29)
			println(s)
		}
	}
	godebug.Line(ctx, scope, 33)

	m := map[string]int{"test": 5}
	scope.Declare("m", &m)
	godebug.Line(ctx, scope, 34)
	if false {
	} else {
		godebug.ElseIfSimpleStmt(ctx, scope, 35)
		_, ok := m["test"]
		godebug.ElseIfExpr(ctx, scope, 35)
		if ok {
			godebug.Line(ctx, scope, 36)
			println("test")
		}
	}
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, scope, 40)
	godebug.Line(ctx, scope, 41)

	const n = 10
	scope.Constant("n", n)
	godebug.Line(ctx, scope, 42)
	_ = n
	godebug.Line(ctx, scope, 44)

	name1(5)
	godebug.Line(ctx, scope, 45)
	name2()
	godebug.Line(ctx, scope, 46)
	T{}.name3()
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
	godebug.Line(ctx, regression_in_go_scope, 51)

	switch {
	case godebug.Case(ctx, regression_in_go_scope, 52):
		fallthrough
	case false:
		godebug.Line(ctx, regression_in_go_scope, 53)
		return 4
	default:
		godebug.Line(ctx, regression_in_go_scope, 54)
		godebug.Line(ctx, regression_in_go_scope, 55)
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
	godebug.Select(ctx, regression_in_go_scope, 61)

	select {
	case <-godebug.Comm(ctx, regression_in_go_scope, 62):
		panic("impossible")
	case <-make(chan bool):
		godebug.Line(ctx, regression_in_go_scope, 62)
		godebug.Line(ctx, regression_in_go_scope, 63)
		return 4
	default:
		godebug.Line(ctx, regression_in_go_scope, 64)
		godebug.Line(ctx, regression_in_go_scope, 65)
		return 5
	case <-godebug.EndSelect(ctx, regression_in_go_scope):
		panic("impossible")
	}
}

func name1(_name1 int) {
	ctx, _ok := godebug.EnterFunc(func() {
		name1(_name1)
	})
	if !_ok {
		return
	}
	defer godebug.ExitFunc(ctx)
	scope := regression_in_go_scope.EnteringNewChildScope()
	scope.Declare("name1", &_name1)
	godebug.Line(ctx, scope, 71)
	if true {
		godebug.Line(ctx, scope, 72)
		_ = _name1
	}
}

func name2() (_name2 string) {
	ctx, _ok := godebug.EnterFunc(func() {
		_name2 = name2()
	})
	if !_ok {
		return _name2
	}
	defer godebug.ExitFunc(ctx)
	scope := regression_in_go_scope.EnteringNewChildScope()
	scope.Declare("name2", &_name2)
	godebug.Line(ctx, scope, 78)
	if true {
		godebug.Line(ctx, scope, 79)
		_name2 = "foo"
	}
	godebug.Line(ctx, scope, 81)
	return _name2
}

type T struct{}

func (_name3 T) name3() {
	ctx, _ok := godebug.EnterFunc(_name3.name3)
	if !_ok {
		return
	}
	defer godebug.ExitFunc(ctx)
	scope := regression_in_go_scope.EnteringNewChildScope()
	scope.Declare("name3", &_name3)
	godebug.Line(ctx, scope, 88)
	if true {
		godebug.Line(ctx, scope, 89)
		_ = _name3
	}
}

var nestedSwitch = func() {
	fn := func(ctx *godebug.Context) {
		godebug.Line(ctx, regression_in_go_scope, 94)
		var foo interface {
		} = 5
		scope := regression_in_go_scope.EnteringNewChildScope()
		scope.Declare("foo", &foo)
		godebug.Line(ctx, scope, 96)
		switch {
		default:
			godebug.Line(ctx, scope, 97)
			godebug.Line(ctx, scope, 98)
			switch foo.(type) {
			case int:
				godebug.Line(ctx, scope, 99)
			}
		}
	}
	if ctx, _ok := godebug.EnterFuncLit(fn); _ok {
		defer godebug.ExitFunc(ctx)
		fn(ctx)
	}
}

func init() {
	doFallthrough()
}

func doFallthrough() {
	ctx, _ok := godebug.EnterFunc(doFallthrough)
	if !_ok {
		return
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, regression_in_go_scope, 110)
	fellthrough := false
	scope := regression_in_go_scope.EnteringNewChildScope()
	scope.Declare("fellthrough", &fellthrough)
	godebug.Line(ctx, scope, 111)
	switch {
	case godebug.Case(ctx, scope, 112):
		fallthrough
	case true:
		godebug.Line(ctx, scope, 113)
		fallthrough
	case godebug.Case(ctx, scope, 114):
		fallthrough
	case false:
		godebug.Line(ctx, scope, 115)
		fellthrough = true
	}
	godebug.Line(ctx, scope, 117)
	if !fellthrough {
		godebug.Line(ctx, scope, 118)
		panic("fallthrough statement did not work")
	}
}

func a() int {
	var result1 int
	ctx, _ok := godebug.EnterFunc(func() {
		result1 = a()
	})
	if !_ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, regression_in_go_scope, 123)
	return 0
}

func switchInit() {
	ctx, _ok := godebug.EnterFunc(switchInit)
	if !_ok {
		return
	}
	defer godebug.ExitFunc(ctx)
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, regression_in_go_scope, 128)
	{
		godebug.Line(ctx, regression_in_go_scope, 129)
		a := a()
		scope := regression_in_go_scope.EnteringNewChildScope()
		scope.Declare("a", &a)
		switch {
		default:
			godebug.Line(ctx, scope, 130)
			godebug.Line(ctx, scope, 131)
			_ = a
		}
	}
	godebug.Line(ctx, regression_in_go_scope, 133)
	_ = "the variable a should be out of scope"
}

func constants() {
	ctx, _ok := godebug.EnterFunc(constants)
	if !_ok {
		return
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, regression_in_go_scope, 137)
	const tooSmallForInt32 = (-1 << 31) - 1
	scope := regression_in_go_scope.EnteringNewChildScope()
	scope.Constant("tooSmallForInt32", int64(tooSmallForInt32))
	godebug.Line(ctx, scope, 138)
	const tooBigForInt64 = 1 << 63
	scope.Constant("tooBigForInt64", uint64(tooBigForInt64))
}

func unexportedField() {
	ctx, _ok := godebug.EnterFunc(unexportedField)
	if !_ok {
		return
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, regression_in_go_scope, 142)
	var f struct{ bar int }
	scope := regression_in_go_scope.EnteringNewChildScope()
	scope.Declare("f", &f)
	godebug.Line(ctx, scope, 143)
	f.bar = 5
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, scope, 144)

}

func init() {
	switchInit()
	unexportedField()
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

	// Comma-ok in else-if
	m := map[string]int{"test": 5}
	if false {
	} else if _, ok := m["test"]; ok {
		println("test")
	}

	// Constant declaration.
	_ = "breakpoint"
	const n = 10
	_ = n

	name1(5)
	name2()
	T{}.name3()
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

// Function shares a name with an input parameter.
func name1(name1 int) {
	if true {
		_ = name1
	}
}

// Function shares a name with an output parameter.
func name2() (name2 string) {
	if true {
		name2 = "foo"
	}
	return name2
}

type T struct{}

// Function shares a name with its receiver
func (name3 T) name3() {
	if true {
		_ = name3
	}
}

var nestedSwitch = func() {
	var foo interface{} = 5
	// Type switch nested inside expression switch
	switch {
	default:
		switch foo.(type) {
		case int:
		}
	}
}

func init() {
	doFallthrough()
}

// Fallthrough should work.
func doFallthrough() {
	fellthrough := false
	switch {
	case true:
		fallthrough
	case false:
		fellthrough = true
	}
	if !fellthrough {
		panic("fallthrough statement did not work")
	}
}

func a() int {
	return 0
}

// Don't repeat switch initialization, use correct scope inside switch.
func switchInit() {
	_ = "breakpoint"
	switch a := a(); {
	default:
		_ = a
	}
	_ = "the variable a should be out of scope"
}

func constants() {
	const tooSmallForInt32 = (-1 << 31) - 1
	const tooBigForInt64 = 1 << 63
}

func unexportedField() {
	var f struct{ bar int }
	f.bar = 5
	_ = "breakpoint"
}

func init() {
	switchInit()
	unexportedField()
}
`


var main_pkg_scope = &godebug.Scope{}

func init() {
	main_pkg_scope.Vars = map[string]interface{}{
		"nestedSwitch": &nestedSwitch,
	}
	main_pkg_scope.Consts = map[string]interface{}{
	}
	main_pkg_scope.Funcs = map[string]interface{}{
		"main": main,
		"_switch": _switch,
		"_select": _select,
		"name1": name1,
		"name2": name2,
		"doFallthrough": doFallthrough,
		"a": a,
		"switchInit": switchInit,
		"constants": constants,
		"unexportedField": unexportedField,
	}
}