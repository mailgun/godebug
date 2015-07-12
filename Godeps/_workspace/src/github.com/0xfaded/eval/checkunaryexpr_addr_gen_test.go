package eval

import (
	"testing"
	"reflect"
)

// Test a
func TestCheckAddrExprA(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expectType(t, `&a`, env, reflect.TypeOf(&a))
}

// Test int(1)
func TestCheckAddrExprInt(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expectCheckError(t, `&int(1)`, env,
		`cannot take the address of int(1)`,
	)

}

// Test 1.4
func TestCheckAddrExprNumber(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expectCheckError(t, `&1.4`, env,
		`cannot take the address of 1.4`,
	)

}

// Test 'a'
func TestCheckAddrExprRune(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expectCheckError(t, `&'a'`, env,
		`cannot take the address of 'a'`,
	)

}

// Test true
func TestCheckAddrExprBool(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expectCheckError(t, `&true`, env,
		`cannot take the address of true`,
	)

}

// Test "a"
func TestCheckAddrExprString(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expectCheckError(t, `&"a"`, env,
		`cannot take the address of "a"`,
	)

}

// Test nil
func TestCheckAddrExprNil(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expectCheckError(t, `&nil`, env,
		`cannot take the address of nil`,
		`use of untyped nil`,
	)

}

// Test  &a
func TestCheckAddrExprAtA(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expectCheckError(t, `& &a`, env,
		`cannot take the address of &a`,
	)

}

// Test  *a
func TestCheckAddrExprStarB(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expectCheckError(t, `& *a`, env,
		`invalid indirect of a (type int)`,
	)

}

// Test []int{1}
func TestCheckAddrExprSlice(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expectType(t, `&[]int{1}`, env, reflect.TypeOf(&[]int{1}))
}

// Test []int{1}[0]
func TestCheckAddrExprSliceElt(t *testing.T) {

	a := 1
	b := &a
	_ = b
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expectType(t, `&[]int{1}[0]`, env, reflect.TypeOf(&[]int{1}[0]))
}
