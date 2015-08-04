package main

import (
	"log"
	"github.com/mailgun/godebug/lib"
)

var recover_in_go_scope = godebug.EnteringNewFile(main_pkg_scope, recover_in_go_contents)

func r1() {
	_r := make(chan chan interface {
	})
	recovers, panicChan := godebug.EnterFuncWithRecovers(_r, func(ctx *godebug.Context) {
		godebug.Line(ctx, recover_in_go_scope, 6)
		<-(<-_r)
	})
	for rr := range recovers {
		rr <- recover()
	}
	if v, ok := <-panicChan; ok {
		panic(v)
	}
}

func r2() {
	_r := make(chan chan interface {
	})
	recovers, panicChan := godebug.EnterFuncWithRecovers(_r, func(ctx *godebug.Context) {
		godebug.Line(ctx, recover_in_go_scope, 10)
		if r := <-(<-_r); r == nil {
			scope := recover_in_go_scope.EnteringNewChildScope()
			scope.Declare("r", &r)
			godebug.Line(ctx, scope, 11)
			log.Fatal("r2: Expected panic, but it didn't happen.")
		}
		godebug.Line(ctx, recover_in_go_scope, 13)
		if r := <-(<-_r); r != nil {
			scope := recover_in_go_scope.EnteringNewChildScope()
			scope.Declare("r", &r)
			godebug.Line(ctx, scope, 14)
			log.Fatal("r2: Second recover should return nil.")
		}
	})
	for rr := range recovers {
		rr <- recover()
	}
	if v, ok := <-panicChan; ok {
		panic(v)
	}
}

var r3 = func() {
	_r := make(chan chan interface {
	})
	recovers, panicChan := godebug.EnterFuncWithRecovers(_r, func(ctx *godebug.Context) {
		godebug.Line(ctx, recover_in_go_scope, 19)
		<-(<-_r)
	})
	for rr := range recovers {
		rr <- recover()
	}
	if v, ok := <-panicChan; ok {
		panic(v)
	}
}

var r4 = func() {
	_r := make(chan chan interface {
	})
	recovers, panicChan := godebug.EnterFuncWithRecovers(_r, func(ctx *godebug.Context) {
		godebug.Line(ctx, recover_in_go_scope, 23)
		if r := <-(<-_r); r == nil {
			scope := recover_in_go_scope.EnteringNewChildScope()
			scope.Declare("r", &r)
			godebug.Line(ctx, scope, 24)
			log.Fatal("r4: Expected panic, but it didn't happen.")
		}
		godebug.Line(ctx, recover_in_go_scope, 26)
		if r := <-(<-_r); r != nil {
			scope := recover_in_go_scope.EnteringNewChildScope()
			scope.Declare("r", &r)
			godebug.Line(ctx, scope, 27)
			log.Fatal("r4: Second recover should return nil.")
		}
	})
	for rr := range recovers {
		rr <- recover()
	}
	if v, ok := <-panicChan; ok {
		panic(v)
	}
}

func doPanic(recoverer func()) {
	ctx, ok := godebug.EnterFunc(func() {
		doPanic(recoverer)
	})
	if !ok {
		return
	}
	defer godebug.ExitFunc(ctx)
	scope := recover_in_go_scope.EnteringNewChildScope()
	scope.Declare("recoverer", &recoverer)
	godebug.Line(ctx, scope, 32)
	defer recoverer()
	defer godebug.Defer(ctx, scope, 32)
	godebug.Line(ctx, scope, 33)
	panic("doPanic: panic")
}

func doNestedRecover(recoverer func()) {
	ctx, ok := godebug.EnterFunc(func() {
		doNestedRecover(recoverer)
	})
	if !ok {
		return
	}
	defer godebug.ExitFunc(ctx)
	scope := recover_in_go_scope.EnteringNewChildScope()
	scope.Declare("recoverer", &recoverer)
	godebug.Line(ctx, scope, 37)
	defer func() {
		_r := make(chan chan interface {
		})
		recovers, panicChan := godebug.EnterFuncWithRecovers(_r, func(ctx *godebug.Context) {
			godebug.Line(ctx, scope, 39)
			recoverer()
			godebug.Line(ctx, scope, 40)
			if r := <-(<-_r); r == nil {
				scope := scope.EnteringNewChildScope()
				scope.Declare("r", &r)
				godebug.Line(ctx, scope, 41)
				log.Fatal("doNestedRecover: Expected to still be panicking, but we aren't.")
			}
		})
		for rr := range recovers {
			rr <- recover()
		}
		if v, ok := <-panicChan; ok {
			panic(v)
		}
	}()
	defer godebug.Defer(ctx, scope, 37)
	godebug.Line(ctx, scope, 44)
	panic("doNestedRecover: panic")
}

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, recover_in_go_scope, 48)
	godebug.Line(ctx, recover_in_go_scope, 49)

	doPanic(r1)
	godebug.Line(ctx, recover_in_go_scope, 50)
	doPanic(r2)
	godebug.Line(ctx, recover_in_go_scope, 51)
	doPanic(r3)
	godebug.Line(ctx, recover_in_go_scope, 52)
	doPanic(r4)
	godebug.Line(ctx, recover_in_go_scope, 53)
	doNestedRecover(r1)
	godebug.Line(ctx, recover_in_go_scope, 54)
	doNestedRecover(r3)
	godebug.Line(ctx, recover_in_go_scope, 56)

	recovererWithParams(2, "foo")
	godebug.Line(ctx, recover_in_go_scope, 58)

	doNestedPanic()
}

func recovererWithParams(i int, s string) bool {
	var result1 bool
	_r := make(chan chan interface {
	})
	recovers, panicChan := godebug.EnterFuncWithRecovers(_r, func(ctx *godebug.Context) {
		result1 = func() bool {
			scope := recover_in_go_scope.EnteringNewChildScope()
			scope.Declare("i", &i, "s", &s)
			godebug.Line(ctx, scope, 62)
			<-(<-_r)
			godebug.Line(ctx, scope, 63)
			return true
		}()
	})
	for rr := range recovers {
		rr <- recover()
	}
	if v, ok := <-panicChan; ok {
		panic(v)
	}
	return result1
}

func doNestedPanic() {
	ctx, ok := godebug.EnterFunc(doNestedPanic)
	if !ok {
		return
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, recover_in_go_scope, 67)
	defer func() {
		_r := make(chan chan interface {
		})
		recovers, panicChan := godebug.EnterFuncWithRecovers(_r, func(ctx *godebug.Context) {
			godebug.Line(ctx, recover_in_go_scope, 68)
			<-(<-_r)
		})
		for rr := range recovers {
			rr <- recover()
		}
		if v, ok := <-panicChan; ok {
			panic(v)
		}
	}()
	defer godebug.Defer(ctx, recover_in_go_scope, 67)
	godebug.Line(ctx, recover_in_go_scope, 70)
	recoverThenPanic()
}

func recoverThenPanic() {
	_r := make(chan chan interface {
	})
	recovers, panicChan := godebug.EnterFuncWithRecovers(_r, func(ctx *godebug.Context) {
		godebug.Line(ctx, recover_in_go_scope, 74)
		<-(<-_r)
		godebug.Line(ctx, recover_in_go_scope, 75)
		panic("panic")
	})
	for rr := range recovers {
		rr <- recover()
	}
	if v, ok := <-panicChan; ok {
		panic(v)
	}
}

var recover_in_go_contents = `package main

import "log"

func r1() {
	recover()
}

func r2() {
	if r := recover(); r == nil {
		log.Fatal("r2: Expected panic, but it didn't happen.")
	}
	if r := recover(); r != nil {
		log.Fatal("r2: Second recover should return nil.")
	}
}

var r3 = func() {
	recover()
}

var r4 = func() {
	if r := recover(); r == nil {
		log.Fatal("r4: Expected panic, but it didn't happen.")
	}
	if r := recover(); r != nil {
		log.Fatal("r4: Second recover should return nil.")
	}
}

func doPanic(recoverer func()) {
	defer recoverer()
	panic("doPanic: panic")
}

func doNestedRecover(recoverer func()) {
	defer func() {
		// The call to recover inside recoverer should not work.
		recoverer()
		if r := recover(); r == nil {
			log.Fatal("doNestedRecover: Expected to still be panicking, but we aren't.")
		}
	}()
	panic("doNestedRecover: panic")
}

func main() {
	_ = "breakpoint"
	doPanic(r1)
	doPanic(r2)
	doPanic(r3)
	doPanic(r4)
	doNestedRecover(r1)
	doNestedRecover(r3)

	recovererWithParams(2, "foo")

	doNestedPanic()
}

func recovererWithParams(i int, s string) bool {
	recover()
	return true
}

func doNestedPanic() {
	defer func() {
		recover()
	}()
	recoverThenPanic()
}

func recoverThenPanic() {
	recover()
	panic("panic")
}
`


var main_pkg_scope = &godebug.Scope{}

func init() {
	main_pkg_scope.Vars = map[string]interface{}{
		"r3": &r3,
		"r4": &r4,
	}
	main_pkg_scope.Consts = map[string]interface{}{
	}
	main_pkg_scope.Funcs = map[string]interface{}{
		"r1": r1,
		"r2": r2,
		"doPanic": doPanic,
		"doNestedRecover": doNestedRecover,
		"main": main,
		"recovererWithParams": recovererWithParams,
		"doNestedPanic": doNestedPanic,
		"recoverThenPanic": recoverThenPanic,
	}
}