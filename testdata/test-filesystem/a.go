package main

import (
	"foo"

	"github.com/mailgun/godebug/lib"
)

func main() {
	godebug.SetTrace()
	foo.HelloWorld()
	foo.HelloWorld()
}
