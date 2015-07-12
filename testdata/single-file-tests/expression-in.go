package main

type Foo struct {
	A int
	B []string
}

func plusTwo(x int) (int, string) {
	return x + 2, "done"
}

const c = 42

var v = []int{1, 2, 3}

func main() {
	f := Foo{
		A: 12,
		B: []string{"hello", "world"},
	}
	_ = f
	_ = "breakpoint"
}
