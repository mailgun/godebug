package main

func Varargs(i ...int) int {
	return 6
}

func main() {
	Varargs(1, 2, 3, 4)
}
