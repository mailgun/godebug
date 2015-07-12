package eval

import (
	"testing"
	"reflect"
)

// Test a
func TestCheckStarExprA(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)

	expectCheckError(t, `*a`, env,
		`invalid indirect of a (type int)`,
	)

}

// Test b
func TestCheckStarExprB(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)

	expectType(t, `*b`, env, reflect.TypeOf(*b))
}

// Test &a
func TestCheckStarExprAtA(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)

	expectType(t, `*&a`, env, reflect.TypeOf(*&a))
}

// Test &b
func TestCheckStarExprAtB(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)

	expectType(t, `*&b`, env, reflect.TypeOf(*&b))
}

// Test int(1)
func TestCheckStarExprInt(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)

	expectCheckError(t, `*int(1)`, env,
		`invalid indirect of int(1) (type int)`,
	)

}

// Test 1.4
func TestCheckStarExprNumber(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)

	expectCheckError(t, `*1.4`, env,
		`invalid indirect of 1.4 (type untyped number)`,
	)

}

// Test 'a'
func TestCheckStarExprRune(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)

	expectCheckError(t, `*'a'`, env,
		`invalid indirect of 'a' (type untyped number)`,
	)

}

// Test true
func TestCheckStarExprBool(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)

	expectCheckError(t, `*true`, env,
		`invalid indirect of true (type untyped bool)`,
	)

}

// Test "a"
func TestCheckStarExprString(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)

	expectCheckError(t, `*"a"`, env,
		`invalid indirect of "a" (type untyped string)`,
	)

}

// Test nil
func TestCheckStarExprNil(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)

	expectCheckError(t, `*nil`, env,
		`invalid indirect of nil`,
	)

}

// Test *b
func TestCheckStarExprStarB(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)

	expectCheckError(t, `**b`, env,
		`invalid indirect of *b (type int)`,
	)

}
