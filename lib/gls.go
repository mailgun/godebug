package godebug

import (
	"bytes"
	"runtime"
	"sync"

	"github.com/jtolds/gls"
)

/*
	This file is inspired by https://github.com/jtolds/gls.
	My goal is to implement a subset of that API with lower overhead.

	I use the goroutine number from the stack trace as an identifier
	and assume that SetValues is not called recursively.
*/

type contextManager interface {
	GetValue(key interface{}) (value interface{}, ok bool)
	SetValues(newValues gls.Values, contextCall func())
}

func getPreferredContextManager() contextManager {
	return gls.NewContextManager()
}

type stackTraceManager struct {
	sync.Mutex
	m map[string]gls.Values
}

func newStackTraceManager() *stackTraceManager {
	return &stackTraceManager{
		m: make(map[string]gls.Values),
	}
}

func (mgr *stackTraceManager) GetValue(key interface{}) (value interface{}, ok bool) {
	id := goid()
	mgr.Lock()
	m, ok := mgr.m[id]
	mgr.Unlock()
	if ok {
		value, ok = m[key]
		return
	}
	return nil, false
}

func (mgr *stackTraceManager) SetValues(newValues gls.Values, contextCall func()) {
	id := goid()
	mgr.Lock()
	mgr.m[id] = newValues
	mgr.Unlock()

	defer func() {
		mgr.Lock()
		delete(mgr.m, id)
		mgr.Unlock()
	}()

	contextCall()
}

func goid() string {
	buf := make([]byte, 15)
	buf = buf[:runtime.Stack(buf, false)]
	return string(bytes.Split(buf, []byte(" "))[1])
}
