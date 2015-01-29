package main

import "fmt"

func main() {
	hi, there := foo(7, 12)
	fmt.Println(hi, there)
	bar()
}

var foo = func(a, _ int) (b, _ string) {
	return "Hello", "World"
}

var bar = func() {
	fmt.Println("No inputs or outputs")
}
