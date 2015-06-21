package main

type Foo int

func (f Foo) Double() Foo {
	return f * 2
}

func (Foo) Seven() Foo {
	return Foo(7)
}

func (_ Foo) Bar() int {
	return 0
}

func main() {
	Foo(3).Double()
}
