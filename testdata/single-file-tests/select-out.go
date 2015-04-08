package main

import (
	"fmt"

	"github.com/mailgun/godebug/lib"
)

var select_in_go_scope = godebug.EnteringNewScope(select_in_go_contents)

func foo() chan int {
	var result1 chan int
	ctx, _ok := godebug.EnterFunc(func() {
		result1 = foo()
	})
	if !_ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, select_in_go_scope, 10)
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
	godebug.Line(ctx, select_in_go_scope, 14)
	return 0
}

func main() {
	ctx, _ok := godebug.EnterFunc(main)
	if !_ok {
		return
	}
	godebug.Line(ctx, select_in_go_scope, 18)
	c := make([]chan int, 10)
	scope := select_in_go_scope.EnteringNewChildScope()
	scope.Declare("c", &c)
	{
		scope := scope.EnteringNewChildScope()
		for i := range c {
			godebug.Line(ctx, scope, 19)
			scope.Declare("i", &i)
			godebug.Line(ctx, scope, 20)
			c[i] = make(chan int, 1)
		}
		godebug.Line(ctx, scope, 19)
	}
	godebug.Line(ctx, scope, 23)

	var r1 int
	scope.Declare("r1", &r1)
	godebug.Line(ctx, scope, 24)
	var ok bool
	scope.Declare("ok", &ok)
	godebug.Line(ctx, scope, 26)

	_, _ = r1, ok

	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, scope, 33)

	go func() {
		fn := func(ctx *godebug.Context) {
			godebug.Select(ctx, scope, 34)
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
	godebug.Select(ctx, scope, 37)

	select {
	default:
		godebug.Line(ctx, scope, 38)
	case <-godebug.EndSelect(ctx, scope):
		panic("impossible")
	}
	godebug.Line(ctx, scope, 41)

	c[0] <- 0
	godebug.Select(ctx, scope, 43)

	select {
	case <-godebug.Comm(ctx, scope, 44):
		panic("impossible")
	case <-c[0]:
		godebug.Line(ctx, scope, 44)
	case <-godebug.EndSelect(ctx, scope):
		panic("impossible")
	}
	godebug.Line(ctx, scope, 50)

	c[0] <- 0
	godebug.Select(ctx, scope, 51)
	select {
	case <-godebug.Comm(ctx, scope, 52):
		panic("impossible")
	case <-c[0]:
		godebug.Line(ctx, scope, 52)
		godebug.Line(ctx, scope, 53)
		hi := "hello"
		scope := scope.EnteringNewChildScope()
		scope.Declare("hi", &hi)
		godebug.Line(ctx, scope, 54)
		fmt.Println(hi)
	default:
		godebug.Line(ctx, scope, 55)
	case <-godebug.Comm(ctx, scope, 56):
		panic("impossible")
	case <-c[1]:
		godebug.Line(ctx, scope, 56)
	case <-godebug.EndSelect(ctx, scope):
		panic("impossible")
	}
	godebug.Line(ctx, scope, 59)

	c[0] <- 0
	{
		godebug.Line(ctx, scope, 61)
		hi := "hi"
		scope := scope.EnteringNewChildScope()
		scope.Declare("hi", &hi)
		godebug.Select(ctx, scope, 62)
		select {
		case <-godebug.Comm(ctx, scope, 63):
			panic("impossible")
		case <-c[0]:
			godebug.Line(ctx, scope, 63)
			godebug.Line(ctx, scope, 64)
			hi := "hello"
			scope := scope.EnteringNewChildScope()
			scope.Declare("hi", &hi)
			godebug.Line(ctx, scope, 65)
			fmt.Println(hi)
		default:
			godebug.Line(ctx, scope, 66)
		case <-godebug.Comm(ctx, scope, 67):
			panic("impossible")
		case <-c[1]:
			godebug.Line(ctx, scope, 67)
		case <-godebug.EndSelect(ctx, scope):
			panic("impossible")
		}
		godebug.Line(ctx, scope, 69)
		_ = hi
	}
	godebug.Select(ctx, scope, 73)

	select {
	case <-godebug.Comm(ctx, scope, 74):
		panic("impossible")
	case <-c[0]:
		godebug.Line(ctx, scope, 74)
	default:
		godebug.Line(ctx, scope, 75)
		godebug.Line(ctx, scope, 76)
		hi := "hello"
		scope := scope.EnteringNewChildScope()
		scope.Declare("hi", &hi)
		godebug.Line(ctx, scope, 77)
		fmt.Println(hi)
	case <-godebug.Comm(ctx, scope, 78):
		panic("impossible")
	case <-c[1]:
		godebug.Line(ctx, scope, 78)
	case <-godebug.EndSelect(ctx, scope):
		panic("impossible")
	}
	godebug.Line(ctx, scope, 84)

	c[9] <- 1
	godebug.Select(ctx, scope, 86)

	select {
	case <-godebug.Comm(ctx, scope, 88):
		panic("impossible")

	case <-c[0]:
		godebug.Line(ctx, scope, 88)
	case <-godebug.Comm(ctx, scope, 89):
		panic("impossible")
	case _ = <-c[1]:
		godebug.Line(ctx, scope, 89)
	case <-godebug.Comm(ctx, scope, 90):
		panic("impossible")
	case r1 = <-c[2]:
		godebug.Line(ctx, scope, 90)
	case <-godebug.Comm(ctx, scope, 91):
		panic("impossible")
	case r2 := <-c[3]:
		godebug.Line(ctx, scope, 91)
		scope := scope.EnteringNewChildScope()
		scope.Declare("r2", &r2)
		godebug.Line(ctx, scope, 92)
		_ = r2
	case <-godebug.Comm(ctx, scope, 94):
		panic("impossible")

	case _, _ = <-c[4]:
		godebug.Line(ctx, scope, 94)
	case <-godebug.Comm(ctx, scope, 95):
		panic("impossible")
	case r1, _ = <-c[5]:
		godebug.Line(ctx, scope, 95)
	case <-godebug.Comm(ctx, scope, 96):
		panic("impossible")
	case _, ok = <-c[6]:
		godebug.Line(ctx, scope, 96)
	case <-godebug.Comm(ctx, scope, 97):
		panic("impossible")
	case _, ok1 := <-c[7]:
		godebug.Line(ctx, scope, 97)
		scope := scope.EnteringNewChildScope()
		scope.Declare("ok1", &ok1)
		godebug.Line(ctx, scope, 98)
		_ = ok1
	case <-godebug.Comm(ctx, scope, 99):
		panic("impossible")
	case r1, ok = <-c[8]:
		godebug.Line(ctx, scope, 99)
	case <-godebug.Comm(ctx, scope, 100):
		panic("impossible")
	case r2, ok := <-c[9]:
		godebug.Line(ctx, scope, 100)
		scope := scope.EnteringNewChildScope()
		scope.Declare("r2", &r2, "ok", &ok)
		godebug.Line(ctx, scope, 101)
		_, _ = r2, ok
	case <-godebug.Comm(ctx, scope, 103):
		panic("impossible")

	case <-foo():
		godebug.Line(ctx, scope, 103)
	case <-godebug.Comm(ctx, scope, 104):
		panic("impossible")
	case _ = <-foo():
		godebug.Line(ctx, scope, 104)
	case <-godebug.Comm(ctx, scope, 105):
		panic("impossible")
	case r1 = <-foo():
		godebug.Line(ctx, scope, 105)
	case <-godebug.Comm(ctx, scope, 106):
		panic("impossible")
	case r2 := <-foo():
		godebug.Line(ctx, scope, 106)
		scope := scope.EnteringNewChildScope()
		scope.Declare("r2", &r2)
		godebug.Line(ctx, scope, 107)
		_ = r2
	case <-godebug.Comm(ctx, scope, 109):
		panic("impossible")

	case _, _ = <-foo():
		godebug.Line(ctx, scope, 109)
	case <-godebug.Comm(ctx, scope, 110):
		panic("impossible")
	case r1, _ = <-foo():
		godebug.Line(ctx, scope, 110)
	case <-godebug.Comm(ctx, scope, 111):
		panic("impossible")
	case _, ok = <-foo():
		godebug.Line(ctx, scope, 111)
	case <-godebug.Comm(ctx, scope, 112):
		panic("impossible")
	case _, ok1 := <-foo():
		godebug.Line(ctx, scope, 112)
		scope := scope.EnteringNewChildScope()
		scope.Declare("ok1", &ok1)
		godebug.Line(ctx, scope, 113)
		_ = ok1
	case <-godebug.Comm(ctx, scope, 114):
		panic("impossible")
	case r1, ok = <-foo():
		godebug.Line(ctx, scope, 114)
	case <-godebug.Comm(ctx, scope, 115):
		panic("impossible")
	case r2, ok := <-foo():
		godebug.Line(ctx, scope, 115)
		scope := scope.EnteringNewChildScope()
		scope.Declare("r2", &r2, "ok", &ok)
		godebug.Line(ctx, scope, 116)
		_, _ = r2, ok
	case <-godebug.EndSelect(ctx, scope):
		panic("impossible")

	}
	godebug.Line(ctx, scope, 123)

	c[0], c[1] = make(chan int), make(chan int)
	godebug.Line(ctx, scope, 125)

	go func() {
		fn := func(ctx *godebug.Context) {
			godebug.Line(ctx, scope, 126)
			<-c[1]
		}
		if ctx, _ok := godebug.EnterFuncLit(fn); _ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
	}()
	godebug.Select(ctx, scope, 129)

	select {
	case <-godebug.Comm(ctx, scope, 131):
		panic("impossible")

	case c[0] <- 0:
		godebug.Line(ctx, scope, 131)
	case <-godebug.Comm(ctx, scope, 132):
		panic("impossible")
	case c[1] <- bar():
		godebug.Line(ctx, scope, 132)
		godebug.Line(ctx, scope, 133)
		fmt.Println("sent")
	case <-godebug.Comm(ctx, scope, 135):
		panic("impossible")

	case foo() <- 0:
		godebug.Line(ctx, scope, 135)
	case <-godebug.Comm(ctx, scope, 136):
		panic("impossible")
	case foo() <- bar():
		godebug.Line(ctx, scope, 136)
	case <-godebug.EndSelect(ctx, scope):
		panic("impossible")

	}
}

var select_in_go_contents = `package main

import (
	"fmt"

	"github.com/mailgun/godebug/lib"
)

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

	godebug.SetTrace()

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
