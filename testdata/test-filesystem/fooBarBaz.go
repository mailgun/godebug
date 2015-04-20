package main

import (
	"bar"
	"baz"
	"foo"
)

func main() {
	_ = "breakpoint"
	foo.Foo()
	bar.Bar()
	baz.Baz()
}
