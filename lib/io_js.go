// +build js

// i/o for javascript
//
// godebug can run on node.js with no i/o changes.
//
// If we're embedded in a web page or web worker, whoever embeds us needs to define up a couple functions:
//
//     godebugOutput
//         Called when godebug wants to write something to stdout.
//         Has one parameter:
//             - String: The data to display.
//
//     godebugPrompt
//         Called by godebug whenever it is ready for user input. Takes no parameters. Should display a prompt.
//         Does not need to block until input is ready. When input is ready, call the global function "godebugInput",
//         which is documented below.
//
// If godebug starts running in an environment that has "godebugPrompt" defined, it will create the following function
// in the global scope:
//
//     godebugInput
//         Should be called when the user responds to a prompt.
//         Takes one parameter:
//             - String: Text input from the user.

package godebug

import "github.com/gopherjs/gopherjs/js"

func init() {
	prompt := js.Global.Get("godebugPrompt")
	if !prompt.Bool() {
		return
	}

	// Expose an input function to javascript for responses to prompts.
	input := make(chan string)
	js.Global.Set("godebugInput", js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
		input <- args[0].String()
		return nil
	}))

	// Hook up gopherJS's output function.
	js.Global.Set("goPrintToConsole", js.InternalObject(func(b []byte) {
		js.Global.Call("godebugOutput", string(b))
	}))

	// Override our internal prompt function.
	promptUser = func() (response string, ok bool) {
		prompt.Invoke()
		response = <-input
		return response, true
	}
}
