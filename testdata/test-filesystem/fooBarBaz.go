package main

import (
	"bar"
	"baz"
	"foo"

	"github.com/mailgun/godebug/lib"
)

func main() {
	godebug.SetTrace()
	foo.Foo()
	bar.Bar()
	baz.Baz()
}
