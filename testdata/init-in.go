package main

// Don't generate debug calls for init.
func init() {
	a = 5
}

var a = 3

type Foo int

// Do generate debug calls for methods named init.
func (f *Foo) init() {
	*f = 1337
}

func main() {
}
