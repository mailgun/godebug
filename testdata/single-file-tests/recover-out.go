package main
import (
	"log"
	"github.com/mailgun/godebug/lib"
)
var recover_in_go_scope = godebug.EnteringNewScope(recover_in_go_contents)
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
			godebug.Line(ctx, recover_in_go_scope, 10)
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
			godebug.Line(ctx, recover_in_go_scope, 14)
			if r := <-(<-_godebug_recover_chan_); r == nil {
				scope := recover_in_go_scope.EnteringNewChildScope()
				scope.Declare("r", &r)
				godebug.Line(ctx, scope, 15)
				log.Fatal("r2: Expected panic, but it didn't happen.")
			}
			godebug.Line(ctx, recover_in_go_scope, 17)
			if r := <-(<-_godebug_recover_chan_); r != nil {
				scope := recover_in_go_scope.EnteringNewChildScope()
				scope.Declare("r", &r)
				godebug.Line(ctx, scope, 18)
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
			godebug.Line(ctx, recover_in_go_scope, 23)
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
			godebug.Line(ctx, recover_in_go_scope, 27)
			if r := <-(<-_godebug_recover_chan_); r == nil {
				scope := recover_in_go_scope.EnteringNewChildScope()
				scope.Declare("r", &r)
				godebug.Line(ctx, scope, 28)
				log.Fatal("r4: Expected panic, but it didn't happen.")
			}
			godebug.Line(ctx, recover_in_go_scope, 30)
			if r := <-(<-_godebug_recover_chan_); r != nil {
				scope := recover_in_go_scope.EnteringNewChildScope()
				scope.Declare("r", &r)
				godebug.Line(ctx, scope, 31)
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
		quit := make(chan struct {
		})
		_godebug_recover_chan_ := make(chan chan interface {
		})
		rr := make(chan interface {
		})
		godebug.Go(func() {
			defer close(quit)
			fn := func(ctx *godebug.Context) {
				godebug.Line(ctx, scope, 43)
				recoverer()
				godebug.Line(ctx, scope, 44)
				if r := <-(<-_godebug_recover_chan_); r == nil {
					scope := scope.EnteringNewChildScope()
					scope.Declare("r", &r)
					godebug.Line(ctx, scope, 45)
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
}
`
