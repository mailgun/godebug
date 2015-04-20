package main

import "fmt"

func foo() chan int {
	return make(chan int)
}

func bar() int {
	return 0
}

func main() {
	c := make([]chan int, 10)
	for i := range c {
		c[i] = make(chan int, 1) // buffered
	}

	var r1 int
	var ok bool

	_, _ = r1, ok

	_ = "breakpoint"

	// -------------------
	// Check simple cases.

	go func() {
		select {}
	}()

	select {
	default:
	}

	c[0] <- 0

	select {
	case <-c[0]:
	}

	// -------------------
	// Check case bodies.

	c[0] <- 0
	select {
	case <-c[0]:
		hi := "hello"
		fmt.Println(hi)
	default:
	case <-c[1]:
	}

	c[0] <- 0
	{
		hi := "hi"
		select {
		case <-c[0]:
			hi := "hello"
			fmt.Println(hi)
		default:
		case <-c[1]:
		}
		_ = hi
	}

	// No send. default case will proceed.
	select {
	case <-c[0]:
	default:
		hi := "hello"
		fmt.Println(hi)
	case <-c[1]:
	}

	// -------------------
	// Check different ways of assigning from receives.

	c[9] <- 1

	select {

	case <-c[0]:
	case _ = <-c[1]:
	case r1 = <-c[2]:
	case r2 := <-c[3]:
		_ = r2

	case _, _ = <-c[4]:
	case r1, _ = <-c[5]:
	case _, ok = <-c[6]:
	case _, ok1 := <-c[7]:
		_ = ok1
	case r1, ok = <-c[8]:
	case r2, ok := <-c[9]: // This is the case that will proceed.
		_, _ = r2, ok

	case <-foo():
	case _ = <-foo():
	case r1 = <-foo():
	case r2 := <-foo():
		_ = r2

	case _, _ = <-foo():
	case r1, _ = <-foo():
	case _, ok = <-foo():
	case _, ok1 := <-foo():
		_ = ok1
	case r1, ok = <-foo():
	case r2, ok := <-foo():
		_, _ = r2, ok

	}

	// -------------------
	// Check sends.

	c[0], c[1] = make(chan int), make(chan int) // unbuffered

	go func() {
		<-c[1]
	}()

	select {

	case c[0] <- 0:
	case c[1] <- bar():
		fmt.Println("sent")

	case foo() <- 0:
	case foo() <- bar():

	}
}
