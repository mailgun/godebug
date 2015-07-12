package godebug

// This file implements expression evaluation for the "print" command.
// Almost all of the work is done by the library github.com/0xfaded/eval.

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/mailgun/godebug/Godeps/_workspace/src/github.com/0xfaded/eval"
)

// Scope represents a lexical scope for variable bindings.
type Scope struct {
	Vars, Consts, Funcs map[string]interface{}
	parent              *Scope
	fileText            []string
}

// EnteringNewFile returns a new Scope and internally sets
// the current scope to be the returned scope.
func EnteringNewFile(parent *Scope, fileText string) *Scope {
	return &Scope{
		Vars:     make(map[string]interface{}),
		Consts:   make(map[string]interface{}),
		Funcs:    make(map[string]interface{}),
		parent:   parent,
		fileText: parseLines(fileText),
	}
}

func parseLines(text string) []string {
	lines := strings.Split(text, "\n")

	// Trailing newline is not a separate line
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines
}

// EnteringNewChildScope returns a new Scope that is the
// child of s and internally sets the current scope to be
// the returned scope.
func (s *Scope) EnteringNewChildScope() *Scope {
	return &Scope{
		Vars:     make(map[string]interface{}),
		Consts:   make(map[string]interface{}),
		Funcs:    make(map[string]interface{}),
		parent:   s,
		fileText: s.fileText,
	}
}

func (s *Scope) getIdent(name string) (i interface{}, ok bool) {
	// TODO: This can race with other goroutines setting the value you are printing.
	for scope := s; scope != nil; scope = scope.parent {
		if i, ok = scope.Vars[name]; ok {
			return dereference(i), true
		}
		if i, ok = scope.Consts[name]; ok {
			return i, true
		}
		if i, ok = scope.Funcs[name]; ok {
			return i, true
		}
	}
	return nil, false
}

// Declare creates new variable bindings in s from a list of name, value pairs.
// The values should be pointers to the values in the program rather than copies
// of them so that s can track changes to them.
func (s *Scope) Declare(namevalue ...interface{}) {
	s.addIdents(s.Vars, "Declare", namevalue...)
}

// Constant is like Declare, but for constants. The values must be passed directly.
func (s *Scope) Constant(namevalue ...interface{}) {
	s.addIdents(s.Consts, "Constant", namevalue...)
}

func (s *Scope) addIdents(to map[string]interface{}, funcName string, namevalue ...interface{}) {
	var i int
	for i = 0; i+1 < len(namevalue); i += 2 {
		name, ok := namevalue[i].(string)
		if !ok {
			panic(fmt.Sprintf("programming error: got odd-numbered argument to %s that was not a string", funcName))
		}
		to[name] = namevalue[i+1]
	}
	if i != len(namevalue) {
		panic(fmt.Sprintf("programming error: called %s with odd number of arguments", funcName))
	}
}

// ----------- Implementation of the github.com/0xfaded/eval.Env interface ------------------ //

func (s *Scope) Var(ident string) reflect.Value {
	return reflect.ValueOf(s.Vars[ident])
}

func (s *Scope) Func(ident string) reflect.Value {
	return reflect.ValueOf(s.Funcs[ident])
}

func (s *Scope) Const(ident string) reflect.Value {
	return reflect.ValueOf(s.Consts[ident])
}

func (s *Scope) Type(ident string) reflect.Type {
	return nil
}

func (s *Scope) Pkg(pkg string) eval.Env {
	return nil
}

func (s *Scope) PushScope() eval.Env {
	return s.EnteringNewChildScope()
}

func (s *Scope) PopScope() eval.Env {
	if s.parent == nil {
		return nil
	} else {
		return s.parent
	}
}

func (s *Scope) AddVar(ident string, v reflect.Value) {
	panic("not implemented")
}

func (s *Scope) AddFunc(ident string, f reflect.Value) {
	panic("not implemented")
}

func (s *Scope) AddConst(ident string, c reflect.Value) {
	panic("not implemented")
}

func (s *Scope) AddType(ident string, t reflect.Type) {
	panic("not implemented")
}

func (s *Scope) AddPkg(pkg string, p eval.Env) {
	panic("not implemented")
}
