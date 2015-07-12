package eval

import (
	"reflect"
)

type envSource int
const (
	envUnknown envSource = iota
	envVar
	envFunc
	envConst
)

// A Environment used for evaluation
type Env interface {
	// Return a pointer value to the variable ident if defined in the top scope, or reflect.Value{} otherwise
	Var(ident string) reflect.Value

	// Return the const value ident if defined in the top scope, or reflect.Value{} otherwise
	Const(ident string) reflect.Value

	// Return the func value ident if defined in the top scope, or reflect.Value{} otherwise
	Func(ident string) reflect.Value

	// Return the type ident if defined in the top scope, or reflect.Value{} otherwise
	Type(ident string) reflect.Type

	// Return the environment containing vars, consts, funcs and types of pkg, or nil if not defined.
	// Unlike other lookup methods, packages exist only in the root scope.
	Pkg(pkg string) Env

	// Create a new block scope. Only the behaviour of the returned Env should change
	PushScope() Env

	// Pop the top block scope. Only the behaviour of the returned Env should change
	PopScope() Env

	// Add var ident to the top scope. The value is always a pointer value, and this same value should be
	// returned by Var(ident). It is up to the implementation how to handle duplicate identifiers.
	AddVar(ident string, v reflect.Value)

	// Add const ident to the top scope. It is up to the implementation how to handle duplicate identifiers.
	AddConst(ident string, c reflect.Value)

	// Add func ident to the top scope. It is up to the implementation how to handle duplicate identifiers.
	AddFunc(ident string, f reflect.Value)

	// Add type ident to the top scope. It is up to the implementation how to handle duplicate identifiers.
	AddType(ident string, t reflect.Type)

	// Add pkg to the root scope. It is up to the implementation how to handle duplicate identifiers.
	AddPkg(pkg string, p Env)
}

func MakeSimpleEnv() *SimpleEnv {
	return &SimpleEnv {
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{},
		Consts: map[string]reflect.Value{},
		Types: map[string]reflect.Type{},
		Pkgs: map[string]Env{},
	}
}

type SimpleEnv struct {
	// path relative to GOROOT or GOPATH. e.g. github.com/0xfaded/eval
	Path string
	Parent *SimpleEnv
	Vars map[string]reflect.Value
	Funcs map[string]reflect.Value
	Consts map[string]reflect.Value
	Types map[string]reflect.Type
	Pkgs map[string]Env
}

func (env *SimpleEnv) Var(ident string) reflect.Value {
	return env.Vars[ident]
}

func (env *SimpleEnv) Func(ident string) reflect.Value {
	return env.Funcs[ident]
}

func (env *SimpleEnv) Const(ident string) reflect.Value {
	return env.Consts[ident]
}

func (env *SimpleEnv) Type(ident string) reflect.Type {
	return env.Types[ident]
}

func (env *SimpleEnv) Pkg(pkg string) Env {
	for env.Parent != nil {
		env = env.Parent
	}
	return env.Pkgs[pkg]
}

func (env *SimpleEnv) PushScope() Env {
	top := MakeSimpleEnv()
	top.Parent = env
	return top
}

func (env *SimpleEnv) PopScope() Env {
	if env.Parent == nil {
		return nil
	} else {
		return env.Parent
	}
}

func (env *SimpleEnv) AddVar(ident string, v reflect.Value) {
	env.Vars[ident] = v
}

func (env *SimpleEnv) AddFunc(ident string, f reflect.Value) {
	env.Funcs[ident] = f
}

func (env *SimpleEnv) AddConst(ident string, c reflect.Value) {
	env.Consts[ident] = c
}

func (env *SimpleEnv) AddType(ident string, t reflect.Type) {
	env.Types[ident] = t
}

func (env *SimpleEnv) AddPkg(pkg string, p Env) {
	for env.Parent != nil {
		env = env.Parent
	}
	env.Pkgs[pkg] = p
}
