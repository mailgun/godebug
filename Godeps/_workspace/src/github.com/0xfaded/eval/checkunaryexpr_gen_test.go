package eval

import (
	"testing"
)

// Test + Int
func TestCheckUnaryExprAddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `+ 4`, env, NewConstInt64(+ 4), ConstInt)
}

// Test + Rune
func TestCheckUnaryExprAddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `+ '@'`, env, NewConstRune(+ '@'), ConstRune)
}

// Test + Float
func TestCheckUnaryExprAddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `+ 2.0`, env, NewConstFloat64(+ 2.0), ConstFloat)
}

// Test + Complex
func TestCheckUnaryExprAddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `+ 8.0i`, env, NewConstComplex128(+ 8.0i), ConstComplex)
}

// Test + Bool
func TestCheckUnaryExprAddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `+ true`, env,
		`invalid operation: + untyped bool`,
	)

}

// Test + String
func TestCheckUnaryExprAddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `+ "abc"`, env,
		`invalid operation: + untyped string`,
	)

}

// Test + Nil
func TestCheckUnaryExprAddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `+ nil`, env,
		`invalid operation: + nil`,
	)

}

// Test - Int
func TestCheckUnaryExprSubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `- 4`, env, NewConstInt64(- 4), ConstInt)
}

// Test - Rune
func TestCheckUnaryExprSubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `- '@'`, env, NewConstRune(- '@'), ConstRune)
}

// Test - Float
func TestCheckUnaryExprSubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `- 2.0`, env, NewConstFloat64(- 2.0), ConstFloat)
}

// Test - Complex
func TestCheckUnaryExprSubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `- 8.0i`, env, NewConstComplex128(- 8.0i), ConstComplex)
}

// Test - Bool
func TestCheckUnaryExprSubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `- true`, env,
		`invalid operation: - untyped bool`,
	)

}

// Test - String
func TestCheckUnaryExprSubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `- "abc"`, env,
		`invalid operation: - untyped string`,
	)

}

// Test - Nil
func TestCheckUnaryExprSubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `- nil`, env,
		`invalid operation: - nil`,
	)

}

// Test ^ Int
func TestCheckUnaryExprXorInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `^ 4`, env, NewConstInt64(^ 4), ConstInt)
}

// Test ^ Rune
func TestCheckUnaryExprXorRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `^ '@'`, env, NewConstRune(^ '@'), ConstRune)
}

// Test ^ Float
func TestCheckUnaryExprXorFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `^ 2.0`, env,
		`illegal constant expression ^ untyped number`,
	)

}

// Test ^ Complex
func TestCheckUnaryExprXorComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `^ 8.0i`, env,
		`illegal constant expression ^ untyped number`,
	)

}

// Test ^ Bool
func TestCheckUnaryExprXorBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `^ true`, env,
		`invalid operation: ^ untyped bool`,
	)

}

// Test ^ String
func TestCheckUnaryExprXorString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `^ "abc"`, env,
		`invalid operation: ^ untyped string`,
	)

}

// Test ^ Nil
func TestCheckUnaryExprXorNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `^ nil`, env,
		`invalid operation: ^ nil`,
	)

}

// Test ! Int
func TestCheckUnaryExprNotInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `! 4`, env,
		`invalid operation: ! untyped number`,
	)

}

// Test ! Rune
func TestCheckUnaryExprNotRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `! '@'`, env,
		`invalid operation: ! untyped number`,
	)

}

// Test ! Float
func TestCheckUnaryExprNotFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `! 2.0`, env,
		`invalid operation: ! untyped number`,
	)

}

// Test ! Complex
func TestCheckUnaryExprNotComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `! 8.0i`, env,
		`invalid operation: ! untyped number`,
	)

}

// Test ! Bool
func TestCheckUnaryExprNotBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `! true`, env, (! true), ConstBool)
}

// Test ! String
func TestCheckUnaryExprNotString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `! "abc"`, env,
		`invalid operation: ! untyped string`,
	)

}

// Test ! Nil
func TestCheckUnaryExprNotNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `! nil`, env,
		`invalid operation: ! nil`,
	)

}
