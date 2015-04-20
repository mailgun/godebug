package testpkg

// This file was written for the test cases in testdata/test.txt

import "testing"

func TestA(t *testing.T) {
	_ = "before Func1 -- debugger should not be paused yet"
	Func1()
	_ = "after Func1"
}

func TestB(t *testing.T) {
	_ = "in TestB"
	Func2()
	_ = "still in TestB"
}

func TestC(*testing.T) {
	_ = "breakpoint"
	_ = "in TestC"
}
