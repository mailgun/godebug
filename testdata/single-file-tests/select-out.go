package main

import (
	"fmt"
	"github.com/mailgun/godebug/lib"
)

var select_in_go_scope = godebug.EnteringNewFile(main_pkg_scope, select_in_go_contents)

func foo() chan int {
	var result1 chan int
	ctx, _ok := godebug.EnterFunc(func() {
		result1 = foo()
	})
	if !_ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, select_in_go_scope, 6)
	return make(chan int)
}

func bar() int {
	var result1 int
	ctx, _ok := godebug.EnterFunc(func() {
		result1 = bar()
	})
	if !_ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, select_in_go_scope, 10)
	return 0
}

func main() {
	ctx, _ok := godebug.EnterFunc(main)
	if !_ok {
		return
	}
	godebug.Line(ctx, select_in_go_scope, 14)
	c := make([]chan int, 10)
	scope := select_in_go_scope.EnteringNewChildScope()
	scope.Declare("c", &c)
	{
		scope := scope.EnteringNewChildScope()
		for i := range c {
			godebug.Line(ctx, scope, 15)
			scope.Declare("i", &i)
			godebug.Line(ctx, scope, 16)
			c[i] = make(chan int, 1)
		}
		godebug.Line(ctx, scope, 15)
	}
	godebug.Line(ctx, scope, 19)

	var r1 int
	scope.Declare("r1", &r1)
	godebug.Line(ctx, scope, 20)
	var ok bool
	scope.Declare("ok", &ok)
	godebug.Line(ctx, scope, 22)

	_, _ = r1, ok
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, scope, 24)
	godebug.Line(ctx, scope, 29)

	go func() {
		fn := func(ctx *godebug.Context) {
			godebug.Select(ctx, scope, 30)
			select {
			case <-godebug.EndSelect(ctx, scope):
				panic("impossible")
			}
		}
		if ctx, _ok := godebug.EnterFuncLit(fn); _ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
	}()
	godebug.Select(ctx, scope, 33)

	select {
	default:
		godebug.Line(ctx, scope, 34)
	case <-godebug.EndSelect(ctx, scope):
		panic("impossible")
	}
	godebug.Line(ctx, scope, 37)

	c[0] <- 0
	godebug.Select(ctx, scope, 39)

	select {
	case <-godebug.Comm(ctx, scope, 40):
		panic("impossible")
	case <-c[0]:
		godebug.Line(ctx, scope, 40)
	case <-godebug.EndSelect(ctx, scope):
		panic("impossible")
	}
	godebug.Line(ctx, scope, 46)

	c[0] <- 0
	godebug.Select(ctx, scope, 47)
	select {
	case <-godebug.Comm(ctx, scope, 48):
		panic("impossible")
	case <-c[0]:
		godebug.Line(ctx, scope, 48)
		godebug.Line(ctx, scope, 49)
		hi := "hello"
		scope := scope.EnteringNewChildScope()
		scope.Declare("hi", &hi)
		godebug.Line(ctx, scope, 50)
		fmt.Println(hi)
	default:
		godebug.Line(ctx, scope, 51)
	case <-godebug.Comm(ctx, scope, 52):
		panic("impossible")
	case <-c[1]:
		godebug.Line(ctx, scope, 52)
	case <-godebug.EndSelect(ctx, scope):
		panic("impossible")
	}
	godebug.Line(ctx, scope, 55)

	c[0] <- 0
	{
		godebug.Line(ctx, scope, 57)
		hi := "hi"
		scope := scope.EnteringNewChildScope()
		scope.Declare("hi", &hi)
		godebug.Select(ctx, scope, 58)
		select {
		case <-godebug.Comm(ctx, scope, 59):
			panic("impossible")
		case <-c[0]:
			godebug.Line(ctx, scope, 59)
			godebug.Line(ctx, scope, 60)
			hi := "hello"
			scope := scope.EnteringNewChildScope()
			scope.Declare("hi", &hi)
			godebug.Line(ctx, scope, 61)
			fmt.Println(hi)
		default:
			godebug.Line(ctx, scope, 62)
		case <-godebug.Comm(ctx, scope, 63):
			panic("impossible")
		case <-c[1]:
			godebug.Line(ctx, scope, 63)
		case <-godebug.EndSelect(ctx, scope):
			panic("impossible")
		}
		godebug.Line(ctx, scope, 65)
		_ = hi
	}
	godebug.Select(ctx, scope, 69)

	select {
	case <-godebug.Comm(ctx, scope, 70):
		panic("impossible")
	case <-c[0]:
		godebug.Line(ctx, scope, 70)
	default:
		godebug.Line(ctx, scope, 71)
		godebug.Line(ctx, scope, 72)
		hi := "hello"
		scope := scope.EnteringNewChildScope()
		scope.Declare("hi", &hi)
		godebug.Line(ctx, scope, 73)
		fmt.Println(hi)
	case <-godebug.Comm(ctx, scope, 74):
		panic("impossible")
	case <-c[1]:
		godebug.Line(ctx, scope, 74)
	case <-godebug.EndSelect(ctx, scope):
		panic("impossible")
	}
	godebug.Line(ctx, scope, 80)

	c[9] <- 1
	godebug.Select(ctx, scope, 82)

	select {
	case <-godebug.Comm(ctx, scope, 84):
		panic("impossible")

	case <-c[0]:
		godebug.Line(ctx, scope, 84)
	case <-godebug.Comm(ctx, scope, 85):
		panic("impossible")
	case _ = <-c[1]:
		godebug.Line(ctx, scope, 85)
	case <-godebug.Comm(ctx, scope, 86):
		panic("impossible")
	case r1 = <-c[2]:
		godebug.Line(ctx, scope, 86)
	case <-godebug.Comm(ctx, scope, 87):
		panic("impossible")
	case r2 := <-c[3]:
		godebug.Line(ctx, scope, 87)
		scope := scope.EnteringNewChildScope()
		scope.Declare("r2", &r2)
		godebug.Line(ctx, scope, 88)
		_ = r2
	case <-godebug.Comm(ctx, scope, 90):
		panic("impossible")

	case _, _ = <-c[4]:
		godebug.Line(ctx, scope, 90)
	case <-godebug.Comm(ctx, scope, 91):
		panic("impossible")
	case r1, _ = <-c[5]:
		godebug.Line(ctx, scope, 91)
	case <-godebug.Comm(ctx, scope, 92):
		panic("impossible")
	case _, ok = <-c[6]:
		godebug.Line(ctx, scope, 92)
	case <-godebug.Comm(ctx, scope, 93):
		panic("impossible")
	case _, ok1 := <-c[7]:
		godebug.Line(ctx, scope, 93)
		scope := scope.EnteringNewChildScope()
		scope.Declare("ok1", &ok1)
		godebug.Line(ctx, scope, 94)
		_ = ok1
	case <-godebug.Comm(ctx, scope, 95):
		panic("impossible")
	case r1, ok = <-c[8]:
		godebug.Line(ctx, scope, 95)
	case <-godebug.Comm(ctx, scope, 96):
		panic("impossible")
	case r2, ok := <-c[9]:
		godebug.Line(ctx, scope, 96)
		scope := scope.EnteringNewChildScope()
		scope.Declare("r2", &r2, "ok", &ok)
		godebug.Line(ctx, scope, 97)
		_, _ = r2, ok
	case <-godebug.Comm(ctx, scope, 99):
		panic("impossible")

	case <-foo():
		godebug.Line(ctx, scope, 99)
	case <-godebug.Comm(ctx, scope, 100):
		panic("impossible")
	case _ = <-foo():
		godebug.Line(ctx, scope, 100)
	case <-godebug.Comm(ctx, scope, 101):
		panic("impossible")
	case r1 = <-foo():
		godebug.Line(ctx, scope, 101)
	case <-godebug.Comm(ctx, scope, 102):
		panic("impossible")
	case r2 := <-foo():
		godebug.Line(ctx, scope, 102)
		scope := scope.EnteringNewChildScope()
		scope.Declare("r2", &r2)
		godebug.Line(ctx, scope, 103)
		_ = r2
	case <-godebug.Comm(ctx, scope, 105):
		panic("impossible")

	case _, _ = <-foo():
		godebug.Line(ctx, scope, 105)
	case <-godebug.Comm(ctx, scope, 106):
		panic("impossible")
	case r1, _ = <-foo():
		godebug.Line(ctx, scope, 106)
	case <-godebug.Comm(ctx, scope, 107):
		panic("impossible")
	case _, ok = <-foo():
		godebug.Line(ctx, scope, 107)
	case <-godebug.Comm(ctx, scope, 108):
		panic("impossible")
	case _, ok1 := <-foo():
		godebug.Line(ctx, scope, 108)
		scope := scope.EnteringNewChildScope()
		scope.Declare("ok1", &ok1)
		godebug.Line(ctx, scope, 109)
		_ = ok1
	case <-godebug.Comm(ctx, scope, 110):
		panic("impossible")
	case r1, ok = <-foo():
		godebug.Line(ctx, scope, 110)
	case <-godebug.Comm(ctx, scope, 111):
		panic("impossible")
	case r2, ok := <-foo():
		godebug.Line(ctx, scope, 111)
		scope := scope.EnteringNewChildScope()
		scope.Declare("r2", &r2, "ok", &ok)
		godebug.Line(ctx, scope, 112)
		_, _ = r2, ok
	case <-godebug.EndSelect(ctx, scope):
		panic("impossible")

	}
	godebug.Line(ctx, scope, 119)

	c[0], c[1] = make(chan int), make(chan int)
	godebug.Line(ctx, scope, 121)

	go func() {
		fn := func(ctx *godebug.Context) {
			godebug.Line(ctx, scope, 122)
			<-c[1]
		}
		if ctx, _ok := godebug.EnterFuncLit(fn); _ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
	}()
	godebug.Select(ctx, scope, 125)

	select {
	case <-godebug.Comm(ctx, scope, 127):
		panic("impossible")

	case c[0] <- 0:
		godebug.Line(ctx, scope, 127)
	case <-godebug.Comm(ctx, scope, 128):
		panic("impossible")
	case c[1] <- bar():
		godebug.Line(ctx, scope, 128)
		godebug.Line(ctx, scope, 129)
		fmt.Println("sent")
	case <-godebug.Comm(ctx, scope, 131):
		panic("impossible")

	case foo() <- 0:
		godebug.Line(ctx, scope, 131)
	case <-godebug.Comm(ctx, scope, 132):
		panic("impossible")
	case foo() <- bar():
		godebug.Line(ctx, scope, 132)
	case <-godebug.EndSelect(ctx, scope):
		panic("impossible")

	}
}

var select_in_go_contents = `package main

import "fmt"

func foo() chan int {
	return make(chan int)
}

func bar() int {
	return 0
}

func main() {
	c := make([]chan int, 10)
	for i := range c {
		c[i] = make(chan int, 1) // buffered
	}

	var r1 int
	var ok bool

	_, _ = r1, ok

	_ = "breakpoint"

	// -------------------
	// Check simple cases.

	go func() {
		select {}
	}()

	select {
	default:
	}

	c[0] <- 0

	select {
	case <-c[0]:
	}

	// -------------------
	// Check case bodies.

	c[0] <- 0
	select {
	case <-c[0]:
		hi := "hello"
		fmt.Println(hi)
	default:
	case <-c[1]:
	}

	c[0] <- 0
	{
		hi := "hi"
		select {
		case <-c[0]:
			hi := "hello"
			fmt.Println(hi)
		default:
		case <-c[1]:
		}
		_ = hi
	}

	// No send. default case will proceed.
	select {
	case <-c[0]:
	default:
		hi := "hello"
		fmt.Println(hi)
	case <-c[1]:
	}

	// -------------------
	// Check different ways of assigning from receives.

	c[9] <- 1

	select {

	case <-c[0]:
	case _ = <-c[1]:
	case r1 = <-c[2]:
	case r2 := <-c[3]:
		_ = r2

	case _, _ = <-c[4]:
	case r1, _ = <-c[5]:
	case _, ok = <-c[6]:
	case _, ok1 := <-c[7]:
		_ = ok1
	case r1, ok = <-c[8]:
	case r2, ok := <-c[9]: // This is the case that will proceed.
		_, _ = r2, ok

	case <-foo():
	case _ = <-foo():
	case r1 = <-foo():
	case r2 := <-foo():
		_ = r2

	case _, _ = <-foo():
	case r1, _ = <-foo():
	case _, ok = <-foo():
	case _, ok1 := <-foo():
		_ = ok1
	case r1, ok = <-foo():
	case r2, ok := <-foo():
		_, _ = r2, ok

	}

	// -------------------
	// Check sends.

	c[0], c[1] = make(chan int), make(chan int) // unbuffered

	go func() {
		<-c[1]
	}()

	select {

	case c[0] <- 0:
	case c[1] <- bar():
		fmt.Println("sent")

	case foo() <- 0:
	case foo() <- bar():

	}
}
`


var main_pkg_scope = &godebug.Scope{}

func init() {
	main_pkg_scope.Vars = map[string]interface{}{
	}
	main_pkg_scope.Consts = map[string]interface{}{
	}
	main_pkg_scope.Funcs = map[string]interface{}{
		"foo": foo,
		"bar": bar,
		"main": main,
	}
}