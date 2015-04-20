package main

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
