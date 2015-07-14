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

	// Constant declaration.
	_ = "breakpoint"
	const n = 10
	_ = n

	name1(5)
	name2()
	T{}.name3()
}

func _switch() int {
	// Terminating switch statement in function with return value.
	switch {
	case false:
		return 4
	default:
		return 5
	}
}

func _select() int {
	// Terminating select statement in function with return value.
	select {
	case <-make(chan bool):
		return 4
	default:
		return 5
	}
}

// Function shares a name with an input parameter.
func name1(name1 int) {
	if true {
		_ = name1
	}
}

// Function shares a name with an output parameter.
func name2() (name2 string) {
	if true {
		name2 = "foo"
	}
	return name2
}

type T struct{}

// Function shares a name with its receiver
func (name3 T) name3() {
	if true {
		_ = name3
	}
}

var nestedSwitch = func() {
	var foo interface{} = 5
	// Type switch nested inside expression switch
	switch {
	default:
		switch foo.(type) {
		case int:
		}
	}
}

func init() {
	doFallthrough()
}

// Fallthrough should work.
func doFallthrough() {
	fellthrough := false
	switch {
	case true:
		fallthrough
	case false:
		fellthrough = true
	}
	if !fellthrough {
		panic("fallthrough statement did not work")
	}
}

func a() int {
	return 0
}

// Don't repeat switch initialization, use correct scope inside switch.
func switchInit() {
	_ = "breakpoint"
	switch a := a(); {
	default:
		_ = a
	}
	_ = "the variable a should be out of scope"
}

func constants() {
	const tooSmallForInt32 = (-1 << 31) - 1
	const tooBigForInt64 = 1 << 63
}

func unexportedField() {
	var f struct{ bar int }
	f.bar = 5
	_ = "breakpoint"
}

func init() {
	switchInit()
	unexportedField()
}
