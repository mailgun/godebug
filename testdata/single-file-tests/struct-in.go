package main

import "github.com/mailgun/godebug/lib"

func main() {
	type myType struct {
		A int
		B string
		C bool
		d int
	}
	var v myType
	godebug.SetTrace()
	_ = v
}
