package main

import "foo"

func main() {
	_ = "breakpoint"
	foo.HelloWorld()
	foo.HelloWorld()
}
