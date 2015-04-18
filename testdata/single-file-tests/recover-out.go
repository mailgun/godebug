package main

import (
	"log"

	"github.com/mailgun/godebug/lib"
)

var recover_in_go_scope = godebug.EnteringNewScope(recover_in_go_contents)

func r1() {
	_r := make(chan chan interface {
	})
	recovers, panicChan := godebug.EnterFuncWithRecovers(_r, func(ctx *godebug.Context) {
		godebug.Line(ctx, recover_in_go_scope, 10)
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
		godebug.Line(ctx, recover_in_go_scope, 14)
		if r := <-(<-_r); r == nil {
			scope := recover_in_go_scope.EnteringNewChildScope()
			scope.Declare("r", &r)
			godebug.Line(ctx, scope, 15)
			log.Fatal("r2: Expected panic, but it didn't happen.")
		}
		godebug.Line(ctx, recover_in_go_scope, 17)
		if r := <-(<-_r); r != nil {
			scope := recover_in_go_scope.EnteringNewChildScope()
			scope.Declare("r", &r)
			godebug.Line(ctx, scope, 18)
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
		godebug.Line(ctx, recover_in_go_scope, 23)
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
		godebug.Line(ctx, recover_in_go_scope, 27)
		if r := <-(<-_r); r == nil {
			scope := recover_in_go_scope.EnteringNewChildScope()
			scope.Declare("r", &r)
			godebug.Line(ctx, scope, 28)
			log.Fatal("r4: Expected panic, but it didn't happen.")
		}
		godebug.Line(ctx, recover_in_go_scope, 30)
		if r := <-(<-_r); r != nil {
			scope := recover_in_go_scope.EnteringNewChildScope()
			scope.Declare("r", &r)
			godebug.Line(ctx, scope, 31)
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
	godebug.Line(ctx, scope, 36)
	defer recoverer()
	defer godebug.Defer(ctx, scope, 36)
	godebug.Line(ctx, scope, 37)
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
	godebug.Line(ctx, scope, 41)
	defer func() {
		_r := make(chan chan interface {
		})
		recovers, panicChan := godebug.EnterFuncWithRecovers(_r, func(ctx *godebug.Context) {
			godebug.Line(ctx, scope, 43)
			recoverer()
			godebug.Line(ctx, scope, 44)
			if r := <-(<-_r); r == nil {
				scope := scope.EnteringNewChildScope()
				scope.Declare("r", &r)
				godebug.Line(ctx, scope, 45)
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
	defer godebug.Defer(ctx, scope, 41)
	godebug.Line(ctx, scope, 48)
	panic("doNestedRecover: panic")
}

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, recover_in_go_scope, 53)
	doPanic(r1)
	godebug.Line(ctx, recover_in_go_scope, 54)
	doPanic(r2)
	godebug.Line(ctx, recover_in_go_scope, 55)
	doPanic(r3)
	godebug.Line(ctx, recover_in_go_scope, 56)
	doPanic(r4)
	godebug.Line(ctx, recover_in_go_scope, 57)
	doNestedRecover(r1)
	godebug.Line(ctx, recover_in_go_scope, 58)
	doNestedRecover(r3)
	godebug.Line(ctx, recover_in_go_scope, 60)

	recovererWithParams(2, "foo")
	godebug.Line(ctx, recover_in_go_scope, 62)

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
			godebug.Line(ctx, scope, 66)
			<-(<-_r)
			godebug.Line(ctx, scope, 67)
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
	godebug.Line(ctx, recover_in_go_scope, 71)
	defer func() {
		_r := make(chan chan interface {
		})
		recovers, panicChan := godebug.EnterFuncWithRecovers(_r, func(ctx *godebug.Context) {
			godebug.Line(ctx, recover_in_go_scope, 72)
			<-(<-_r)
		})
		for rr := range recovers {
			rr <- recover()
		}
		if v, ok := <-panicChan; ok {
			panic(v)
		}
	}()
	defer godebug.Defer(ctx, recover_in_go_scope, 71)
	godebug.Line(ctx, recover_in_go_scope, 74)
	recoverThenPanic()
}

func recoverThenPanic() {
	_r := make(chan chan interface {
	})
	recovers, panicChan := godebug.EnterFuncWithRecovers(_r, func(ctx *godebug.Context) {
		godebug.Line(ctx, recover_in_go_scope, 78)
		<-(<-_r)
		godebug.Line(ctx, recover_in_go_scope, 79)
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

import (
	"log"

	"github.com/mailgun/godebug/lib"
)

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
	godebug.SetTrace()
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
