// +build js

// goroutine local storage for gopherjs

package godebug

import "github.com/gopherjs/gopherjs/js"

var registry []jsContext

func getContextManager() contextManager {
	cm := make(jsContext)
	registry = append(registry, cm)
	return cm
}

// Go preserves goroutine-local-storage across new goroutine invocations.
var Go = func(cb func()) {
	for _, mgr := range registry {
		values, ok := mgr.getValues()
		if ok {
			mgrCopy := mgr
			cbCopy := cb
			cb = func() { mgrCopy.setValues(values, cbCopy) }
		}
	}

	go cb()
}

var glsIDs idPool

type jsContext map[int]map[interface{}]interface{}

func (c jsContext) GetValue(key interface{}) (value interface{}, ok bool) {
	m, ok := c.getValues()
	if !ok {
		return nil, false
	}
	value, ok = m[key]
	return
}

// Unlike the implementation in github.com/jtolds/gls, this implementation
// assumes that SetValues is not called multiple times in the same stack for
// the same context manager.
func (c jsContext) SetValues(contextCall func(), keyVal ...interface{}) {
	if len(keyVal)%2 != 0 {
		panic("bad argument to SetValues")
	}

	values := make(map[interface{}]interface{})
	for i := 0; i < len(keyVal); i += 2 {
		values[keyVal[i]] = keyVal[i+1]
	}
	c.setValues(values, contextCall)
}

func (c jsContext) setValues(values map[interface{}]interface{}, contextCall func()) {
	id := glsIDs.Acquire()
	defer glsIDs.Release(id)
	c[int(id)] = values
	defer delete(c, int(id))

	curGoroutine().Set("godebugKey", int(id))
	defer curGoroutine().Delete("godebugKey")

	contextCall()
}

func (c jsContext) getValues() (values map[interface{}]interface{}, ok bool) {
	id := curGoroutine().Get("godebugKey")
	if id == js.Undefined {
		return nil, false
	}

	values, ok = c[id.Int()]
	return
}

func curGoroutine() *js.Object {
	return js.Global.Get("$curGoroutine")
}
