package main

func main() {
	type myType struct {
		A int
		B string
		C bool
		d int
	}
	var v myType
	_ = "breakpoint"
	_ = v
}
