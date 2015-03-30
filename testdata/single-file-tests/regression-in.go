package main

func main() {
	// Nested scope in the first declaration in a function.
	foo := func(i int) int {
		return i
	}(3)
	_ = foo

	// String literal in range statement.
	// Blank identifier in range statement.
	for _, s := range []string{"foo"} {
		_ = s
	}

	// go statement with function literal.
	c := make(chan bool)
	go func() {
		c <- true
	}()
	<-c

	// String literal in defer statement.
	defer println("Hello")

	// String literal in else-if statement.
	if false {
	} else if s := "hello"; s == "hello" {
		println(s)
	}

	// Comma-ok in else-if
	m := map[string]int{"test": 5}
	if false {
	} else if _, ok := m["test"]; ok {
		println("test")
	}
}
