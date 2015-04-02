package testpkg

// This file was written for the test cases in testdata/test.txt

import "github.com/mailgun/godebug/lib"

func Func1() {
	godebug.SetTrace()
	_ = "inside Func1"
}

func Func2() {
	_ = "inside Func2"
}
