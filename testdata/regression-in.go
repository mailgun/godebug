package main

func main() {
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
}
