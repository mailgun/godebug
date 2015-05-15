/*
Package atexit implements a simple atexit library.

Example:
	package main

	import (
		"fmt"

		"bitbucket.org/JeremySchlatter/go-atexit"
	)

	func main() {
		atexit.TrapSignals()
		defer atexit.CallExitFuncs()

		atexit.Run(func() {
			fmt.Println("exited")
		})
	}
*/
package atexit

import (
	"os"
	"os/signal"
	"syscall"
)

var exitFuncs []func()

// TrapSignals registers a handler for SIGINT and SIGTERM
// that calls CallExitFuncs() followed by os.Exit(1).
func TrapSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		CallExitFuncs()
		os.Exit(1)
	}()
}

// Run registers the function f to be called at the program's exit.
func Run(f func()) {
	exitFuncs = append(exitFuncs, f)
}

// CallExitFuncs calls all functions previously registered through Run, in reverse order.
func CallExitFuncs() {
	for i := len(exitFuncs) - 1; i >= 0; i-- {
		exitFuncs[i]()
	}
}
