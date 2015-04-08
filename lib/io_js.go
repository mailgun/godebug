// +build never

// i/o for javascript

package godebug

import (
	"github.com/gopherjs/gopherjs/js"
)

// Variables used by debug.go
var (
	input   = &jsInput{}
	outputW jsOutput
)

// Variables used just in this file
var (
	document  = js.Global.Get("document")
	outputDiv = js.Global.Get("document").Call("getElementById", "#output")
)

type jsInput struct {
	s string
}

func (s *jsInput) Scan() bool {

	// scan stuff
	return true
}

func (s *jsInput) Text() string {
	return s.s
}

type jsOutput struct{}

func (jsOutput) Write(p []byte) (n int, err error) {
	js.Global.Get("console").Call("log", string(p))
	/*
		span := document.Call("createElement", "span")
		span.Set("className", "stdout")
		span.Set("innerHTML", string(p))
		outputDiv.Call("appendChild", span)
		return len(p), nil
	*/
}
