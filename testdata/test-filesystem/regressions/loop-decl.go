package main

import "github.com/mailgun/godebug/lib"

func main() {
	a()
	b()
}

func a() {
	for i := 0; i < 1; i++ {
		godebug.SetTrace()
		_ = i
	}
}

func b() {
	for _, s := range []string{"hello"} {
		godebug.SetTrace()
		_ = s
	}
}
