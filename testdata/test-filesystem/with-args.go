package main

import (
	"flag"

	"github.com/mailgun/godebug/lib"
)

func main() {
	var foo string
	flag.StringVar(&foo, "foo", "foo's default value", "a string flag")
	godebug.SetTrace()
	flag.Parse()
	_ = foo
}
