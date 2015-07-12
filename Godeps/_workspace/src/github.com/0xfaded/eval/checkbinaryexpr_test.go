package eval

import (
	"testing"
)

// This file contains three groups of tests.
// Key: ck. check, ov. overflow, tr. truncation, bo. bad operation
// 1. Untyped op Untyped. These tests are divided accordingly
//
//	    | integer |  rune   |floating | complex |  bool   | string  |   nil   |
//          +---------+---------+---------|---------+---------+---------+---------+
// integer  |         |         |         |         |         |         |         |
// rune     |         |         |         |         |         |         |         |
// floating |         |         |         |         |         |         |         |
// complex  |         |         |         |         |         |         |         |
// bool     |         |         |         |         |         |         |         |
// string   |         |         |         |         |         |         |         |
// nil      |         |         |         |         |         |         |         |

// integer op X tests
func TestBasicCheckConstBinaryIntegerInteger(t *testing.T) {
	env := MakeSimpleEnv()

	// Valid
	expectConst(t, "5 / 2", env, NewConstInt64(5 / 2), ConstInt)
	expectConst(t, "5 % 2", env, NewConstInt64(5 % 2), ConstInt)
	expectConst(t, "5 & 2", env, NewConstInt64(5 & 2), ConstInt)
	expectConst(t, "5 == 2", env, 5 == 2, ConstBool)
	expectConst(t, "5 <= 2", env, 5 <= 2, ConstBool)
}

func TestBasicCheckConstBinaryIntegerRune(t *testing.T) {
	env := MakeSimpleEnv()

	// Valid
	expectConst(t, "5 - 'a'", env, NewConstInt64(5 - 'a'), ConstRune)
	expectConst(t, "5 % 'a'", env, NewConstInt64(5 % 'a'), ConstRune)
	expectConst(t, "5 | 'a'", env, NewConstInt64(5 | 'a'), ConstRune)
	expectConst(t, "5 != 'a'", env, 5 != 'a', ConstBool)
	expectConst(t, "5 >= 'a'", env, 5 >= 'a', ConstBool)
}

func TestBasicCheckConstBinaryIntegerFloating(t *testing.T) {
	env := MakeSimpleEnv()

	// Valid
	expectConst(t, "5 / 2.0", env, NewConstFloat64(5 / 2.0), ConstFloat)
	expectConst(t, "5 != 1.5", env, 5 != 1.5, ConstBool)
	expectConst(t, "5 < 1.5", env, 5 < 1.5, ConstBool)

	// Invalid
	expectCheckError(t, "5 % 1.0", env, "illegal constant expression: floating-point % operation")
	expectCheckError(t, "5 | 1.0", env, "illegal constant expression: untyped number | untyped number")
}

func TestBasicCheckConstBinaryIntegerComplex(t *testing.T) {
	env := MakeSimpleEnv()

	// Vaild
	expectConst(t, "5 / 1.25i", env, NewConstComplex128(5 / 1.25i), ConstComplex)
	expectConst(t, "5 != 1.5i", env, 5 != 1.5i, ConstBool)

	// Invalid
	expectCheckError(t, "5 > 1.5i", env, "illegal constant expression: untyped number > untyped number")
	expectCheckError(t, "5 % 2.0i", env, "illegal constant expression: untyped number % untyped number")
	expectCheckError(t, "5 & 1.5i", env, "illegal constant expression: untyped number & untyped number")
}

func TestBasicCheckConstBinaryIntegerBool(t *testing.T) {
	env := MakeSimpleEnv()

	// Invalid
	expectCheckError(t, "5 + true", env,
		"cannot convert true to type int",
		"invalid operation: 5 + true (mismatched types int and bool)",
	)
	expectCheckError(t, "5 % true", env,
		"cannot convert true to type int",
		"invalid operation: 5 % true (mismatched types int and bool)",
	)
	expectCheckError(t, "5 & true", env,
		"cannot convert true to type int",
		"invalid operation: 5 & true (mismatched types int and bool)",
	)
	expectCheckError(t, "5 == true", env,
		"cannot convert true to type int",
		"invalid operation: 5 == true (mismatched types int and bool)",
	)
	expectCheckError(t, "5 < true", env,
		"cannot convert true to type int",
		"invalid operation: 5 < true (mismatched types int and bool)",
	)
}

func TestBasicCheckConstBinaryIntegerString(t *testing.T) {
	env := MakeSimpleEnv()

	// Invalid
	expectCheckError(t, `5 + "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 5 + "abc" (mismatched types int and string)`,
	)
	expectCheckError(t, `5 % "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 5 % "abc" (mismatched types int and string)`,
	)
	expectCheckError(t, `5 & "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 5 & "abc" (mismatched types int and string)`,
	)
	expectCheckError(t, `5 == "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 5 == "abc" (mismatched types int and string)`,
	)
	expectCheckError(t, `5 < "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 5 < "abc" (mismatched types int and string)`,
	)
}

func TestBasicCheckConstBinaryIntegerNil(t *testing.T) {
	env := MakeSimpleEnv()

	// Invalid
	expectCheckError(t, "5 + nil", env,
		"cannot convert nil to type int",
		"invalid operation: 5 + nil (mismatched types int and <T>)",
	)
	expectCheckError(t, "5 % nil", env,
		"cannot convert nil to type int",
		"invalid operation: 5 % nil (mismatched types int and <T>)",
	)
	expectCheckError(t, "5 & nil", env,
		"cannot convert nil to type int",
		"invalid operation: 5 & nil (mismatched types int and <T>)",
	)
	expectCheckError(t, "5 == nil", env,
		"cannot convert nil to type int",
		"invalid operation: 5 == nil (mismatched types int and <T>)",
	)
	expectCheckError(t, "5 < nil", env,
		"cannot convert nil to type int",
		"invalid operation: 5 < nil (mismatched types int and <T>)",
	)
}

// rune op X tests
func TestBasicCheckConstBinaryRuneInteger(t *testing.T) {
	env := MakeSimpleEnv()

	// Valid
	expectConst(t, "'a' / 2", env, NewConstRune('a' / 2), ConstRune)
	expectConst(t, "'a' % 2", env, NewConstRune('a' % 2), ConstRune)
	expectConst(t, "'a' & 2", env, NewConstRune('a' & 2), ConstRune)
	expectConst(t, "'a' != 2", env, 'a' != 2, ConstBool)
	expectConst(t, "'a' >= 2", env, 'a' >= 2, ConstBool)
}

func TestBasicCheckConstBinaryRuneRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, "'a' / '\\x04'", env, NewConstRune('a' / '\x04'), ConstRune)
	expectConst(t, "'a' % '\\x04'", env, NewConstRune('a' % '\x04'), ConstRune)
	expectConst(t, "'a' ^ '\\x04'", env, NewConstRune('a' ^ '\x04'), ConstRune)
	expectConst(t, "'a' != '\\x04'", env, 'a' != '\x04', ConstBool)
	expectConst(t, "'a' > '\\x04'", env, 'a' > '\x04', ConstBool)
}

func TestBasicCheckConstBinaryRuneFloating(t *testing.T) {
	env := MakeSimpleEnv()

	// Valid
	expectConst(t, "'d' * 1.25", env, NewConstFloat64('d' * 1.25), ConstFloat)
	expectConst(t, "'d' < 1.25", env, 'd' < 1.25, ConstBool)
	expectConst(t, "'d' != 100.0", env, 'd' != 100.0, ConstBool)

	// Invalid
	expectCheckError(t, "'d' % 1.0", env, "illegal constant expression: floating-point % operation")
	expectCheckError(t, "'d' &^ 1.0", env, "illegal constant expression: untyped number &^ untyped number")
}

func TestBasicCheckConstBinaryRuneComplex(t *testing.T) {
	env := MakeSimpleEnv()

	// Valid
	expectConst(t, "'d' / 4i", env, NewConstComplex128('d' / 4i), ConstComplex)
	expectConst(t, "'a' == 1i", env, 'a' == 1i, ConstBool)

	// Invalid
	expectCheckError(t, "'d' > 1.5i", env, "illegal constant expression: untyped number > untyped number")
	expectCheckError(t, "'d' % 1.0i", env, "illegal constant expression: untyped number % untyped number")
	expectCheckError(t, "'d' ^ 1.0i", env, "illegal constant expression: untyped number ^ untyped number")
}

func TestBasicCheckConstBinaryRuneBool(t *testing.T) {
	env := MakeSimpleEnv()

	// Invalid
	expectCheckError(t, "'a' + true", env,
		"cannot convert true to type rune",
		"invalid operation: 'a' + true (mismatched types rune and bool)",
	)
	expectCheckError(t, "'a' % true", env,
		"cannot convert true to type rune",
		"invalid operation: 'a' % true (mismatched types rune and bool)",
	)
	expectCheckError(t, "'a' & true", env,
		"cannot convert true to type rune",
		"invalid operation: 'a' & true (mismatched types rune and bool)",
	)
	expectCheckError(t, "'a' == true", env,
		"cannot convert true to type rune",
		"invalid operation: 'a' == true (mismatched types rune and bool)",
	)
	expectCheckError(t, "'a' < true", env,
		"cannot convert true to type rune",
		"invalid operation: 'a' < true (mismatched types rune and bool)",
	)
}

func TestBasicCheckConstBinaryRuneString(t *testing.T) {
	env := MakeSimpleEnv()

	// Invalid
	expectCheckError(t, `'a' + "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: 'a' + "abc" (mismatched types rune and string)`,
	)
	expectCheckError(t, `'a' % "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: 'a' % "abc" (mismatched types rune and string)`,
	)
	expectCheckError(t, `'a' & "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: 'a' & "abc" (mismatched types rune and string)`,
	)
	expectCheckError(t, `'a' == "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: 'a' == "abc" (mismatched types rune and string)`,
	)
	expectCheckError(t, `'a' < "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: 'a' < "abc" (mismatched types rune and string)`,
	)
}

func TestBasicCheckConstBinaryRuneNil(t *testing.T) {
	env := MakeSimpleEnv()

	// Invalid
	expectCheckError(t, "'a' + nil", env,
		"cannot convert nil to type rune",
		"invalid operation: 'a' + nil (mismatched types rune and <T>)",
	)
	expectCheckError(t, "'a' % nil", env,
		"cannot convert nil to type rune",
		"invalid operation: 'a' % nil (mismatched types rune and <T>)",
	)
	expectCheckError(t, "'a' & nil", env,
		"cannot convert nil to type rune",
		"invalid operation: 'a' & nil (mismatched types rune and <T>)",
	)
	expectCheckError(t, "'a' == nil", env,
		"cannot convert nil to type rune",
		"invalid operation: 'a' == nil (mismatched types rune and <T>)",
	)
	expectCheckError(t, "'a' < nil", env,
		"cannot convert nil to type rune",
		"invalid operation: 'a' < nil (mismatched types rune and <T>)",
	)
}

// floating op X tests
func TestBasicCheckConstBinaryFloatingInteger(t *testing.T) {
	env := MakeSimpleEnv()

	// Valid
	expectConst(t, "1.5 + 2", env, NewConstFloat64(1.5 + 2), ConstFloat)
	expectConst(t, "1.5 < 100", env, 1.5 < 100, ConstBool)
	expectConst(t, "1.5 == 5", env, 1.5 == 5, ConstBool)

	// Invalid
	expectCheckError(t, "1.5 % 2", env, "illegal constant expression: floating-point % operation")
	expectCheckError(t, "1.5 &^ 3", env, "illegal constant expression: untyped number &^ untyped number")
}

func TestBasicCheckConstBinaryFloatingRune(t *testing.T) {
	env := MakeSimpleEnv()

	// Valid
	expectConst(t, "1.5 - 'a'", env, NewConstFloat64(1.5 - 'a'), ConstFloat)
	expectConst(t, "1.5 < 'a'", env, 1.5 < 'a', ConstBool)
	expectConst(t, "1.5 == 'a'", env, 1.5 == 'a', ConstBool)

	// Invalid
	expectCheckError(t, "1.5 % 'a'", env, "illegal constant expression: floating-point % operation")
	expectCheckError(t, "1.5 | 'a'", env, "illegal constant expression: untyped number | untyped number")
}

func TestBasicCheckConstBinaryFloatingFloating(t *testing.T) {
	env := MakeSimpleEnv()

	// Valid
	expectConst(t, "2.5 * 1.25", env, NewConstFloat64(2.5 * 1.25), ConstFloat)
	expectConst(t, "1.5 < 1.25", env, 1.5 < 1.25, ConstBool)
	expectConst(t, "1.5 == 1.25", env, 1.5 == 1.25, ConstBool)

	// Invalid
	expectCheckError(t, "1.5 % 1.5", env, "illegal constant expression: floating-point % operation")
	expectCheckError(t, "1.5 | 1.5", env, "illegal constant expression: untyped number | untyped number")
}

func TestBasicCheckConstBinaryFloatingComplex(t *testing.T) {
	env := MakeSimpleEnv()

	// Valid
	expectConst(t, "2.5 / 4i", env, NewConstComplex128(2.5 / 4i), ConstComplex)
	expectConst(t, "1.5 == 1.25i", env, 1.5 == 1.25i, ConstBool)

	// Invalid
	expectCheckError(t, "1.5 < 1.25i", env, "illegal constant expression: untyped number < untyped number")
	expectCheckError(t, "1.5 % 1.5i", env, "illegal constant expression: untyped number % untyped number")
	expectCheckError(t, "1.5 | 1.5i", env, "illegal constant expression: untyped number | untyped number")
}

func TestBasicCheckConstBinaryFloatBool(t *testing.T) {
	env := MakeSimpleEnv()

	// Invalid
	expectCheckError(t, "1.0 + true", env,
		"cannot convert true to type float64",
		"invalid operation: 1 + true (mismatched types float64 and bool)",
	)
	expectCheckError(t, "1.0 % true", env,
		"cannot convert true to type float64",
		"invalid operation: 1 % true (mismatched types float64 and bool)",
	)
	expectCheckError(t, "1.0 & true", env,
		"cannot convert true to type float64",
		"invalid operation: 1 & true (mismatched types float64 and bool)",
	)
	expectCheckError(t, "1.0 == true", env,
		"cannot convert true to type float64",
		"invalid operation: 1 == true (mismatched types float64 and bool)",
	)
	expectCheckError(t, "1.0 < true", env,
		"cannot convert true to type float64",
		"invalid operation: 1 < true (mismatched types float64 and bool)",
	)
}

func TestBasicCheckConstBinaryFloatString(t *testing.T) {
	env := MakeSimpleEnv()

	// Invalid
	expectCheckError(t, `1.0 + "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 1 + "abc" (mismatched types float64 and string)`,
	)
	expectCheckError(t, `1.0 % "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 1 % "abc" (mismatched types float64 and string)`,
	)
	expectCheckError(t, `1.0 & "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 1 & "abc" (mismatched types float64 and string)`,
	)
	expectCheckError(t, `1.0 == "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 1 == "abc" (mismatched types float64 and string)`,
	)
	expectCheckError(t, `1.0 < "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 1 < "abc" (mismatched types float64 and string)`,
	)
}

func TestBasicCheckConstBinaryFloatNil(t *testing.T) {
	env := MakeSimpleEnv()

	// Invalid
	expectCheckError(t, "1.0 + nil", env,
		"cannot convert nil to type float64",
		"invalid operation: 1 + nil (mismatched types float64 and <T>)",
	)
	expectCheckError(t, "1.0 % nil", env,
		"cannot convert nil to type float64",
		"invalid operation: 1 % nil (mismatched types float64 and <T>)",
	)
	expectCheckError(t, "1.0 & nil", env,
		"cannot convert nil to type float64",
		"invalid operation: 1 & nil (mismatched types float64 and <T>)",
	)
	expectCheckError(t, "1.0 == nil", env,
		"cannot convert nil to type float64",
		"invalid operation: 1 == nil (mismatched types float64 and <T>)",
	)
	expectCheckError(t, "1.0 < nil", env,
		"cannot convert nil to type float64",
		"invalid operation: 1 < nil (mismatched types float64 and <T>)",
	)
}

// complex op X tests
func TestBasicCheckConstBinaryComplexInteger(t *testing.T) {
	env := MakeSimpleEnv()

	// Invalid
	expectConst(t, "2.5i + 2", env, NewConstComplex128(2.5i + 2), ConstComplex)
	expectConst(t, "2.5i == 2", env, 2.5i == 2, ConstBool)

	// Invalid
	expectCheckError(t, "2.5i < 2", env, "illegal constant expression: untyped number < untyped number")
	expectCheckError(t, "2.5i % 2", env, "illegal constant expression: untyped number % untyped number")
	expectCheckError(t, "2.5i | 2", env, "illegal constant expression: untyped number | untyped number")
}

func TestBasicCheckConstBinaryComplexRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, "2.5i - 'a'", env, NewConstComplex128(2.5i - 'a'), ConstComplex)
	expectConst(t, "2.5i == 'a'", env, 2.5i == 'a', ConstBool)

	// Invalid
	expectCheckError(t, "2.5i < 'a'", env, "illegal constant expression: untyped number < untyped number")
	expectCheckError(t, "2.5i % 'a'", env, "illegal constant expression: untyped number % untyped number")
	expectCheckError(t, "2.5i | 'a'", env, "illegal constant expression: untyped number | untyped number")
}

func TestBasicCheckConstBinaryComplexComplexing(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, "2.5i * 1.25", env, NewConstComplex128(2.5i * 1.25), ConstComplex)
	expectConst(t, "2.5i == 2.0", env, 2.5i == 2.0, ConstBool)

	// Invalid
	expectCheckError(t, "2.5i < 2.0", env, "illegal constant expression: untyped number < untyped number")
	expectCheckError(t, "2.5i % 2.0", env, "illegal constant expression: untyped number % untyped number")
	expectCheckError(t, "2.5i | 2.0", env, "illegal constant expression: untyped number | untyped number")
}

func TestBasicCheckConstBinaryComplexComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, "3i / 4i", env, NewConstComplex128(3i / 4i), ConstComplex)
	expectConst(t, "2.5i == 2i", env, 2.5i == 2i, ConstBool)

	// Invalid
	expectCheckError(t, "2.5i < 2i", env, "illegal constant expression: untyped number < untyped number")
	expectCheckError(t, "2.5i % 2i", env, "illegal constant expression: untyped number % untyped number")
	expectCheckError(t, "2.5i | 2i", env, "illegal constant expression: untyped number | untyped number")
}

func TestBasicCheckConstBinaryComplexBool(t *testing.T) {
	env := MakeSimpleEnv()

	// Invalid
	expectCheckError(t, "1.0i + true", env,
		"cannot convert true to type complex128",
		"invalid operation: 1i + true (mismatched types complex128 and bool)",
	)
	expectCheckError(t, "1.0i % true", env,
		"cannot convert true to type complex128",
		"invalid operation: 1i % true (mismatched types complex128 and bool)",
	)
	expectCheckError(t, "1.0i & true", env,
		"cannot convert true to type complex128",
		"invalid operation: 1i & true (mismatched types complex128 and bool)",
	)
	expectCheckError(t, "1.0i == true", env,
		"cannot convert true to type complex128",
		"invalid operation: 1i == true (mismatched types complex128 and bool)",
	)
	expectCheckError(t, "1.0i < true", env,
		"cannot convert true to type complex128",
		"invalid operation: 1i < true (mismatched types complex128 and bool)",
	)
}

func TestBasicCheckConstBinaryComplexString(t *testing.T) {
	env := MakeSimpleEnv()

	// Invalid
	expectCheckError(t, `1.0i + "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 1i + "abc" (mismatched types complex128 and string)`,
	)
	expectCheckError(t, `1.0i % "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 1i % "abc" (mismatched types complex128 and string)`,
	)
	expectCheckError(t, `1.0i & "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 1i & "abc" (mismatched types complex128 and string)`,
	)
	expectCheckError(t, `1.0i == "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 1i == "abc" (mismatched types complex128 and string)`,
	)
	expectCheckError(t, `1.0i < "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 1i < "abc" (mismatched types complex128 and string)`,
	)
}

func TestBasicCheckConstBinaryComplexNil(t *testing.T) {
	env := MakeSimpleEnv()

	// Invalid
	expectCheckError(t, "1.0i + nil", env,
		"cannot convert nil to type complex128",
		"invalid operation: 1i + nil (mismatched types complex128 and <T>)",
	)
	expectCheckError(t, "1.0i % nil", env,
		"cannot convert nil to type complex128",
		"invalid operation: 1i % nil (mismatched types complex128 and <T>)",
	)
	expectCheckError(t, "1.0i & nil", env,
		"cannot convert nil to type complex128",
		"invalid operation: 1i & nil (mismatched types complex128 and <T>)",
	)
	expectCheckError(t, "1.0i == nil", env,
		"cannot convert nil to type complex128",
		"invalid operation: 1i == nil (mismatched types complex128 and <T>)",
	)
	expectCheckError(t, "1.0i < nil", env,
		"cannot convert nil to type complex128",
		"invalid operation: 1i < nil (mismatched types complex128 and <T>)",
	)
}

// bool op X tests
func TestBasicCheckConstBinaryBoolBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, "true == true", env, true, ConstBool)
	expectConst(t, "true != true", env, false, ConstBool)
	expectConst(t, "true == false", env, false, ConstBool)
	expectConst(t, "true != false", env, true, ConstBool)
}
