package main

func main() {
	foo(3, 3)
}

func foo(int, int) (string, error) {
	return "hello", nil
}
