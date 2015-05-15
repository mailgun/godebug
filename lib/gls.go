// +build !js

package godebug

// Pure-Go implementation of goroutine local storage.

import "github.com/mailgun/godebug/Godeps/_workspace/src/github.com/jtolds/gls"

var Go = gls.Go

func getContextManager() contextManager {
	return wrapper{gls.NewContextManager()}
}

type wrapper struct {
	*gls.ContextManager
}

func (w wrapper) SetValues(contextCall func(), keyVal ...interface{}) {
	if len(keyVal)%2 != 0 {
		panic("bad argument to SetValues")
	}
	vals := make(gls.Values)
	for i := 0; i < len(keyVal); i += 2 {
		vals[keyVal[i]] = keyVal[i+1]
	}

	w.ContextManager.SetValues(vals, contextCall)
}
