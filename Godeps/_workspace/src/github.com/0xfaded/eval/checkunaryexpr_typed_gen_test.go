package eval

import (
	"testing"
	"reflect"
)

// Test + Int32
func TestCheckUnaryTypedExprAddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `+ int32(4)`, env, + int32(4), reflect.TypeOf(+ int32(4)))
}

// Test + Float64
func TestCheckUnaryTypedExprAddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `+ float64(2)`, env, + float64(2), reflect.TypeOf(+ float64(2)))
}

// Test + Complex128
func TestCheckUnaryTypedExprAddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `+ complex128(8i)`, env, + complex128(8i), reflect.TypeOf(+ complex128(8i)))
}

// Test + Bool
func TestCheckUnaryTypedExprAddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `+ bool(true)`, env,
		`invalid operation: + bool`,
	)

}

// Test + String
func TestCheckUnaryTypedExprAddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `+ string("abc")`, env,
		`invalid operation: + string`,
	)

}

// Test - Int32
func TestCheckUnaryTypedExprSubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `- int32(4)`, env, - int32(4), reflect.TypeOf(- int32(4)))
}

// Test - Float64
func TestCheckUnaryTypedExprSubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `- float64(2)`, env, - float64(2), reflect.TypeOf(- float64(2)))
}

// Test - Complex128
func TestCheckUnaryTypedExprSubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `- complex128(8i)`, env, - complex128(8i), reflect.TypeOf(- complex128(8i)))
}

// Test - Bool
func TestCheckUnaryTypedExprSubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `- bool(true)`, env,
		`invalid operation: - bool`,
	)

}

// Test - String
func TestCheckUnaryTypedExprSubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `- string("abc")`, env,
		`invalid operation: - string`,
	)

}

// Test ^ Int32
func TestCheckUnaryTypedExprXorInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `^ int32(4)`, env, ^ int32(4), reflect.TypeOf(^ int32(4)))
}

// Test ^ Float64
func TestCheckUnaryTypedExprXorFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `^ float64(2)`, env,
		`invalid operation: ^ float64`,
	)

}

// Test ^ Complex128
func TestCheckUnaryTypedExprXorComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `^ complex128(8i)`, env,
		`invalid operation: ^ complex128`,
	)

}

// Test ^ Bool
func TestCheckUnaryTypedExprXorBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `^ bool(true)`, env,
		`invalid operation: ^ bool`,
	)

}

// Test ^ String
func TestCheckUnaryTypedExprXorString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `^ string("abc")`, env,
		`invalid operation: ^ string`,
	)

}

// Test ! Int32
func TestCheckUnaryTypedExprNotInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `! int32(4)`, env,
		`invalid operation: ! int32`,
	)

}

// Test ! Float64
func TestCheckUnaryTypedExprNotFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `! float64(2)`, env,
		`invalid operation: ! float64`,
	)

}

// Test ! Complex128
func TestCheckUnaryTypedExprNotComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `! complex128(8i)`, env,
		`invalid operation: ! complex128`,
	)

}

// Test ! Bool
func TestCheckUnaryTypedExprNotBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `! bool(true)`, env, ! bool(true), reflect.TypeOf(! bool(true)))
}

// Test ! String
func TestCheckUnaryTypedExprNotString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `! string("abc")`, env,
		`invalid operation: ! string`,
	)

}
