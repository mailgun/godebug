package main

import (
	"log"

	"github.com/mailgun/godebug/lib"
)

var recover_in_goScope = godebug.EnteringNewScope()

func r1() {
	quit := make(chan struct {
	})
	_godebug_recover_chan_ := make(chan chan interface {
	})
	rr := make(chan interface {
	})
	godebug.Go(func() {
		defer close(quit)
		fn := func(ctx *godebug.Context) {
			godebug.Line(ctx, recover_in_goScope)
			<-(<-_godebug_recover_chan_)
		}
		if ctx, ok := godebug.EnterFuncLit(fn); ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
	})
	for {
		select {
		case <-quit:
			return
		case _godebug_recover_chan_ <- rr:
			rr <- recover()
		}
	}
}

func r2() {
	quit := make(chan struct {
	})
	_godebug_recover_chan_ := make(chan chan interface {
	})
	rr := make(chan interface {
	})
	godebug.Go(func() {
		defer close(quit)
		fn := func(ctx *godebug.Context) {
			godebug.Line(ctx, recover_in_goScope)
			if r := <-(<-_godebug_recover_chan_); r == nil {
				scope := recover_in_goScope.EnteringNewChildScope()
				scope.Declare("r", &r)
				godebug.Line(ctx, scope)
				log.Fatal("r2: Expected panic, but it didn't happen.")
			}
			godebug.Line(ctx, recover_in_goScope)
			if r := <-(<-_godebug_recover_chan_); r != nil {
				scope := recover_in_goScope.EnteringNewChildScope()
				scope.Declare("r", &r)
				godebug.Line(ctx, scope)
				log.Fatal("r2: Second recover should return nil.")
			}
		}
		if ctx, ok := godebug.EnterFuncLit(fn); ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
	})
	for {
		select {
		case <-quit:
			return
		case _godebug_recover_chan_ <- rr:
			rr <- recover()
		}
	}
}

var r3 = func() {
	quit := make(chan struct {
	})
	_godebug_recover_chan_ := make(chan chan interface {
	})
	rr := make(chan interface {
	})
	godebug.Go(func() {
		defer close(quit)
		fn := func(ctx *godebug.Context) {
			godebug.Line(ctx, recover_in_goScope)
			<-(<-_godebug_recover_chan_)
		}
		if ctx, ok := godebug.EnterFuncLit(fn); ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
	})
	for {
		select {
		case <-quit:
			return
		case _godebug_recover_chan_ <- rr:
			rr <- recover()
		}
	}
}

var r4 = func() {
	quit := make(chan struct {
	})
	_godebug_recover_chan_ := make(chan chan interface {
	})
	rr := make(chan interface {
	})
	godebug.Go(func() {
		defer close(quit)
		fn := func(ctx *godebug.Context) {
			godebug.Line(ctx, recover_in_goScope)
			if r := <-(<-_godebug_recover_chan_); r == nil {
				scope := recover_in_goScope.EnteringNewChildScope()
				scope.Declare("r", &r)
				godebug.Line(ctx, scope)
				log.Fatal("r4: Expected panic, but it didn't happen.")
			}
			godebug.Line(ctx, recover_in_goScope)
			if r := <-(<-_godebug_recover_chan_); r != nil {
				scope := recover_in_goScope.EnteringNewChildScope()
				scope.Declare("r", &r)
				godebug.Line(ctx, scope)
				log.Fatal("r4: Second recover should return nil.")
			}
		}
		if ctx, ok := godebug.EnterFuncLit(fn); ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
	})
	for {
		select {
		case <-quit:
			return
		case _godebug_recover_chan_ <- rr:
			rr <- recover()
		}
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
	scope := recover_in_goScope.EnteringNewChildScope()
	scope.Declare("recoverer", &recoverer)
	godebug.Line(ctx, scope)
	defer recoverer()
	defer godebug.SLine(ctx, scope, "<Running deferred function>: defer recoverer()")
	godebug.Line(ctx, scope)
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
	scope := recover_in_goScope.EnteringNewChildScope()
	scope.Declare("recoverer", &recoverer)
	godebug.Line(ctx, scope)
	defer func() {
		quit := make(chan struct {
		})
		_godebug_recover_chan_ := make(chan chan interface {
		})
		rr := make(chan interface {
		})
		godebug.Go(func() {
			defer close(quit)
			fn := func(ctx *godebug.Context) {
				godebug.Line(ctx, scope)
				recoverer()
				godebug.Line(ctx, scope)
				if r := <-(<-_godebug_recover_chan_); r == nil {
					scope := scope.EnteringNewChildScope()
					scope.Declare("r", &r)
					godebug.Line(ctx, scope)
					log.Fatal("doNestedRecover: Expected to still be panicking, but we aren't.")
				}
			}
			if ctx, ok := godebug.EnterFuncLit(fn); ok {
				defer godebug.ExitFunc(ctx)
				fn(ctx)
			}
		})
		for {
			select {
			case <-quit:
				return
			case _godebug_recover_chan_ <- rr:
				rr <- recover()
			}
		}
	}()
	defer godebug.SLine(ctx, scope, "<Running deferred function>: defer func() {")
	godebug.Line(ctx, scope)
	panic("doNestedRecover: panic")
}

func main() {
	ctx, ok := godebug.EnterFunc(main)
	if !ok {
		return
	}
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, recover_in_goScope)
	doPanic(r1)
	godebug.Line(ctx, recover_in_goScope)
	doPanic(r2)
	godebug.Line(ctx, recover_in_goScope)
	doPanic(r3)
	godebug.Line(ctx, recover_in_goScope)
	doPanic(r4)
	godebug.Line(ctx, recover_in_goScope)
	doNestedRecover(r1)
	godebug.Line(ctx, recover_in_goScope)
	doNestedRecover(r3)
}
