package eval

import (
	"testing"
)

// TODO the make() type check in gc 1.2 is quite broken. Delete these tests
// and enable the generated make() tests in 1.3

func TestCheckBuiltinMakeMissingArgs(t *testing.T) {
	env := MakeSimpleEnv()
	expectCheckError(t, "make()", env, "missing argument to make")
}

func TestCheckBuiltinMakeUntypedIntFirstArg(t *testing.T) {
	env := MakeSimpleEnv()
	expectCheckError(t, "make(1)", env, "1 is not a type")
}

func TestCheckBuiltinMakeTypedIntFirstArg(t *testing.T) {
	env := MakeSimpleEnv()
	expectCheckError(t, "make(int(1))", env, "int(1) is not a type")
}

func TestCheckBuiltinMakeTypedNilFirstArg(t *testing.T) {
	env := MakeSimpleEnv()
	expectCheckError(t, "make(nil)", env, "nil is not a type")
}

func TestCheckBuiltinMakeSecondArgNotInt(t *testing.T) {
	env := MakeSimpleEnv()
	expectCheckError(t, "make([]int, true)", env, "make: non-integer len argument true")
}

func TestCheckBuiltinMakeThirdArgNotInt(t *testing.T) {
	env := MakeSimpleEnv()
	expectCheckError(t, "make([]int, 1, bool(true))", env, "make: non-integer cap argument bool(true)")
}

func TestCheckBuiltinMakeChanTooManyArgs(t *testing.T) {
	env := MakeSimpleEnv()
	expectCheckError(t, "make(chan<- int, 1, 1)", env, "too many arguments to make: make(chan<- int, 1, 1)")
}

func TestCheckBuiltinMakeSliceTooFew(t *testing.T) {
	env := MakeSimpleEnv()
	expectCheckError(t, "make([]int)", env, "too few arguments to make: make([]int)")
}

func TestCheckBuiltinMakeSliceTooManyArgs(t *testing.T) {
	env := MakeSimpleEnv()
	expectCheckError(t, "make([]int, 1, 1, 1)", env, "too many arguments to make: make([]int, 1, 1, 1)")
}

func TestCheckBuiltinMakeSliceLenGtrThanCap(t *testing.T) {
	env := MakeSimpleEnv()
	expectCheckError(t, "make([]int, 3, 2)", env, "len larger than cap in make([]int, 3, 2)")
}

func TestCheckBuiltinMakeBadType(t *testing.T) {
	env := MakeSimpleEnv()
	expectCheckError(t, "make(int)", env, "cannot make type int")
}

