// +build js

// i/o for javascript
//
// When running on node.js, we have access to stdin and stdout as normal,
// so nothing needs to change. When we run in a browser, we do i/o through
// [jq-console](https://replit.github.io/jq-console/).

package godebug

import "github.com/gopherjs/gopherjs/js"

// OnReady lets us know that we are running in a browser gives us access to an
// instance of jqconsole.
func OnReady(jqconsole *js.Object) {
	js.Global.Set("goPrintToConsole", js.InternalObject(func(b []byte) {
		jqconsole.Call("Write", string(b), "jqconsole-output")
	}))

	promptUser = func() (response string, ok bool) {
		s := make(chan string)
		jqconsole.Call("Prompt", true, js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
			s <- args[0].String()
			return nil
		}))
		return <-s, true
	}
}
