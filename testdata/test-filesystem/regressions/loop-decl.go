package main

func main() {
	a()
	b()
}

func a() {
	for i := 0; i < 1; i++ {
		_ = "breakpoint"
		_ = i
	}
}

func b() {
	for _, s := range []string{"hello"} {
		_ = "breakpoint"
		_ = s
	}
}
