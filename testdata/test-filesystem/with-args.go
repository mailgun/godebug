package main

import "flag"

func main() {
	var foo string
	flag.StringVar(&foo, "foo", "foo's default value", "a string flag")
	_ = "breakpoint"
	flag.Parse()
	_ = foo
}
