package main

import "fmt"

func foo() interface{} {
	return "hi"
}

func main() {
	_ = "breakpoint"

	switch {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	}

	i := 3

	switch i {
	case foo():
	default:
	case 5, 4, 1:
	case 2:
	}

	var ifc interface{} = i

	switch ifc.(type) {
	case string:
	case bool:
	}

	switch b := 2; b == 6 {
	case true:
	case false:
	}

	switch b := ifc; i := ifc.(type) {
	case string:
	case int:
	default:
		_, _ = i, b
	}
}
