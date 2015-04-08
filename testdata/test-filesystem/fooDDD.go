package main

// fooDDD = fooD[ot]D[ot]D[ot]

import (
	"foo"
	"foo/subfoo"

	"github.com/mailgun/godebug/lib"
)

func main() {
	godebug.SetTrace()
	foo.Foo()
	subfoo.SubFoo()
}
