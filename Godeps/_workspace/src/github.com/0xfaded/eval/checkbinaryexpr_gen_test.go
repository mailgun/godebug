package eval

import (
	"testing"
)

// Test Int + Int
func TestCheckBinaryExprIntAddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 + 4`, env, NewConstInt64(4 + 4), ConstInt)
}

// Test Int + Rune
func TestCheckBinaryExprIntAddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 + '@'`, env, NewConstRune(4 + '@'), ConstRune)
}

// Test Int + Float
func TestCheckBinaryExprIntAddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 + 2.0`, env, NewConstFloat64(4 + 2.0), ConstFloat)
}

// Test Int + Complex
func TestCheckBinaryExprIntAddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 + 8.0i`, env, NewConstComplex128(4 + 8.0i), ConstComplex)
}

// Test Int + Bool
func TestCheckBinaryExprIntAddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 + true`, env,
		`cannot convert true to type int`,
		`invalid operation: 4 + true (mismatched types int and bool)`,
	)

}

// Test Int + String
func TestCheckBinaryExprIntAddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 + "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 4 + "abc" (mismatched types int and string)`,
	)

}

// Test Int + Nil
func TestCheckBinaryExprIntAddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 + nil`, env,
		`cannot convert nil to type int`,
		`invalid operation: 4 + nil (mismatched types int and <T>)`,
	)

}

// Test Int - Int
func TestCheckBinaryExprIntSubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 - 4`, env, NewConstInt64(4 - 4), ConstInt)
}

// Test Int - Rune
func TestCheckBinaryExprIntSubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 - '@'`, env, NewConstRune(4 - '@'), ConstRune)
}

// Test Int - Float
func TestCheckBinaryExprIntSubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 - 2.0`, env, NewConstFloat64(4 - 2.0), ConstFloat)
}

// Test Int - Complex
func TestCheckBinaryExprIntSubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 - 8.0i`, env, NewConstComplex128(4 - 8.0i), ConstComplex)
}

// Test Int - Bool
func TestCheckBinaryExprIntSubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 - true`, env,
		`cannot convert true to type int`,
		`invalid operation: 4 - true (mismatched types int and bool)`,
	)

}

// Test Int - String
func TestCheckBinaryExprIntSubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 - "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 4 - "abc" (mismatched types int and string)`,
	)

}

// Test Int - Nil
func TestCheckBinaryExprIntSubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 - nil`, env,
		`cannot convert nil to type int`,
		`invalid operation: 4 - nil (mismatched types int and <T>)`,
	)

}

// Test Int * Int
func TestCheckBinaryExprIntMulInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 * 4`, env, NewConstInt64(4 * 4), ConstInt)
}

// Test Int * Rune
func TestCheckBinaryExprIntMulRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 * '@'`, env, NewConstRune(4 * '@'), ConstRune)
}

// Test Int * Float
func TestCheckBinaryExprIntMulFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 * 2.0`, env, NewConstFloat64(4 * 2.0), ConstFloat)
}

// Test Int * Complex
func TestCheckBinaryExprIntMulComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 * 8.0i`, env, NewConstComplex128(4 * 8.0i), ConstComplex)
}

// Test Int * Bool
func TestCheckBinaryExprIntMulBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 * true`, env,
		`cannot convert true to type int`,
		`invalid operation: 4 * true (mismatched types int and bool)`,
	)

}

// Test Int * String
func TestCheckBinaryExprIntMulString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 * "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 4 * "abc" (mismatched types int and string)`,
	)

}

// Test Int * Nil
func TestCheckBinaryExprIntMulNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 * nil`, env,
		`cannot convert nil to type int`,
		`invalid operation: 4 * nil (mismatched types int and <T>)`,
	)

}

// Test Int / Int
func TestCheckBinaryExprIntQuoInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 / 4`, env, NewConstInt64(4 / 4), ConstInt)
}

// Test Int / Rune
func TestCheckBinaryExprIntQuoRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 / '@'`, env, NewConstRune(4 / '@'), ConstRune)
}

// Test Int / Float
func TestCheckBinaryExprIntQuoFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 / 2.0`, env, NewConstFloat64(4 / 2.0), ConstFloat)
}

// Test Int / Complex
func TestCheckBinaryExprIntQuoComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 / 8.0i`, env, NewConstComplex128(4 / 8.0i), ConstComplex)
}

// Test Int / Bool
func TestCheckBinaryExprIntQuoBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 / true`, env,
		`cannot convert true to type int`,
		`invalid operation: 4 / true (mismatched types int and bool)`,
	)

}

// Test Int / String
func TestCheckBinaryExprIntQuoString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 / "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 4 / "abc" (mismatched types int and string)`,
	)

}

// Test Int / Nil
func TestCheckBinaryExprIntQuoNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 / nil`, env,
		`cannot convert nil to type int`,
		`invalid operation: 4 / nil (mismatched types int and <T>)`,
	)

}

// Test Int % Int
func TestCheckBinaryExprIntRemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 % 4`, env, NewConstInt64(4 % 4), ConstInt)
}

// Test Int % Rune
func TestCheckBinaryExprIntRemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 % '@'`, env, NewConstRune(4 % '@'), ConstRune)
}

// Test Int % Float
func TestCheckBinaryExprIntRemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 % 2.0`, env,
		`illegal constant expression: floating-point % operation`,
	)

}

// Test Int % Complex
func TestCheckBinaryExprIntRemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 % 8.0i`, env,
		`illegal constant expression: untyped number % untyped number`,
	)

}

// Test Int % Bool
func TestCheckBinaryExprIntRemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 % true`, env,
		`cannot convert true to type int`,
		`invalid operation: 4 % true (mismatched types int and bool)`,
	)

}

// Test Int % String
func TestCheckBinaryExprIntRemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 % "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 4 % "abc" (mismatched types int and string)`,
	)

}

// Test Int % Nil
func TestCheckBinaryExprIntRemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 % nil`, env,
		`cannot convert nil to type int`,
		`invalid operation: 4 % nil (mismatched types int and <T>)`,
	)

}

// Test Int & Int
func TestCheckBinaryExprIntAndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 & 4`, env, NewConstInt64(4 & 4), ConstInt)
}

// Test Int & Rune
func TestCheckBinaryExprIntAndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 & '@'`, env, NewConstRune(4 & '@'), ConstRune)
}

// Test Int & Float
func TestCheckBinaryExprIntAndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 & 2.0`, env,
		`illegal constant expression: untyped number & untyped number`,
	)

}

// Test Int & Complex
func TestCheckBinaryExprIntAndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 & 8.0i`, env,
		`illegal constant expression: untyped number & untyped number`,
	)

}

// Test Int & Bool
func TestCheckBinaryExprIntAndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 & true`, env,
		`cannot convert true to type int`,
		`invalid operation: 4 & true (mismatched types int and bool)`,
	)

}

// Test Int & String
func TestCheckBinaryExprIntAndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 & "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 4 & "abc" (mismatched types int and string)`,
	)

}

// Test Int & Nil
func TestCheckBinaryExprIntAndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 & nil`, env,
		`cannot convert nil to type int`,
		`invalid operation: 4 & nil (mismatched types int and <T>)`,
	)

}

// Test Int | Int
func TestCheckBinaryExprIntOrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 | 4`, env, NewConstInt64(4 | 4), ConstInt)
}

// Test Int | Rune
func TestCheckBinaryExprIntOrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 | '@'`, env, NewConstRune(4 | '@'), ConstRune)
}

// Test Int | Float
func TestCheckBinaryExprIntOrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 | 2.0`, env,
		`illegal constant expression: untyped number | untyped number`,
	)

}

// Test Int | Complex
func TestCheckBinaryExprIntOrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 | 8.0i`, env,
		`illegal constant expression: untyped number | untyped number`,
	)

}

// Test Int | Bool
func TestCheckBinaryExprIntOrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 | true`, env,
		`cannot convert true to type int`,
		`invalid operation: 4 | true (mismatched types int and bool)`,
	)

}

// Test Int | String
func TestCheckBinaryExprIntOrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 | "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 4 | "abc" (mismatched types int and string)`,
	)

}

// Test Int | Nil
func TestCheckBinaryExprIntOrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 | nil`, env,
		`cannot convert nil to type int`,
		`invalid operation: 4 | nil (mismatched types int and <T>)`,
	)

}

// Test Int ^ Int
func TestCheckBinaryExprIntXorInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 ^ 4`, env, NewConstInt64(4 ^ 4), ConstInt)
}

// Test Int ^ Rune
func TestCheckBinaryExprIntXorRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 ^ '@'`, env, NewConstRune(4 ^ '@'), ConstRune)
}

// Test Int ^ Float
func TestCheckBinaryExprIntXorFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 ^ 2.0`, env,
		`illegal constant expression: untyped number ^ untyped number`,
	)

}

// Test Int ^ Complex
func TestCheckBinaryExprIntXorComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 ^ 8.0i`, env,
		`illegal constant expression: untyped number ^ untyped number`,
	)

}

// Test Int ^ Bool
func TestCheckBinaryExprIntXorBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 ^ true`, env,
		`cannot convert true to type int`,
		`invalid operation: 4 ^ true (mismatched types int and bool)`,
	)

}

// Test Int ^ String
func TestCheckBinaryExprIntXorString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 ^ "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 4 ^ "abc" (mismatched types int and string)`,
	)

}

// Test Int ^ Nil
func TestCheckBinaryExprIntXorNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 ^ nil`, env,
		`cannot convert nil to type int`,
		`invalid operation: 4 ^ nil (mismatched types int and <T>)`,
	)

}

// Test Int &^ Int
func TestCheckBinaryExprIntAndNotInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 &^ 4`, env, NewConstInt64(4 &^ 4), ConstInt)
}

// Test Int &^ Rune
func TestCheckBinaryExprIntAndNotRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 &^ '@'`, env, NewConstRune(4 &^ '@'), ConstRune)
}

// Test Int &^ Float
func TestCheckBinaryExprIntAndNotFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 &^ 2.0`, env,
		`illegal constant expression: untyped number &^ untyped number`,
	)

}

// Test Int &^ Complex
func TestCheckBinaryExprIntAndNotComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 &^ 8.0i`, env,
		`illegal constant expression: untyped number &^ untyped number`,
	)

}

// Test Int &^ Bool
func TestCheckBinaryExprIntAndNotBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 &^ true`, env,
		`cannot convert true to type int`,
		`invalid operation: 4 &^ true (mismatched types int and bool)`,
	)

}

// Test Int &^ String
func TestCheckBinaryExprIntAndNotString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 &^ "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 4 &^ "abc" (mismatched types int and string)`,
	)

}

// Test Int &^ Nil
func TestCheckBinaryExprIntAndNotNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 &^ nil`, env,
		`cannot convert nil to type int`,
		`invalid operation: 4 &^ nil (mismatched types int and <T>)`,
	)

}

// Test Int == Int
func TestCheckBinaryExprIntEqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 == 4`, env, (4 == 4), ConstBool)
}

// Test Int == Rune
func TestCheckBinaryExprIntEqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 == '@'`, env, (4 == '@'), ConstBool)
}

// Test Int == Float
func TestCheckBinaryExprIntEqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 == 2.0`, env, (4 == 2.0), ConstBool)
}

// Test Int == Complex
func TestCheckBinaryExprIntEqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 == 8.0i`, env, (4 == 8.0i), ConstBool)
}

// Test Int == Bool
func TestCheckBinaryExprIntEqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 == true`, env,
		`cannot convert true to type int`,
		`invalid operation: 4 == true (mismatched types int and bool)`,
	)

}

// Test Int == String
func TestCheckBinaryExprIntEqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 == "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 4 == "abc" (mismatched types int and string)`,
	)

}

// Test Int == Nil
func TestCheckBinaryExprIntEqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 == nil`, env,
		`cannot convert nil to type int`,
		`invalid operation: 4 == nil (mismatched types int and <T>)`,
	)

}

// Test Int != Int
func TestCheckBinaryExprIntNeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 != 4`, env, (4 != 4), ConstBool)
}

// Test Int != Rune
func TestCheckBinaryExprIntNeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 != '@'`, env, (4 != '@'), ConstBool)
}

// Test Int != Float
func TestCheckBinaryExprIntNeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 != 2.0`, env, (4 != 2.0), ConstBool)
}

// Test Int != Complex
func TestCheckBinaryExprIntNeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 != 8.0i`, env, (4 != 8.0i), ConstBool)
}

// Test Int != Bool
func TestCheckBinaryExprIntNeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 != true`, env,
		`cannot convert true to type int`,
		`invalid operation: 4 != true (mismatched types int and bool)`,
	)

}

// Test Int != String
func TestCheckBinaryExprIntNeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 != "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 4 != "abc" (mismatched types int and string)`,
	)

}

// Test Int != Nil
func TestCheckBinaryExprIntNeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 != nil`, env,
		`cannot convert nil to type int`,
		`invalid operation: 4 != nil (mismatched types int and <T>)`,
	)

}

// Test Int <= Int
func TestCheckBinaryExprIntLeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 <= 4`, env, (4 <= 4), ConstBool)
}

// Test Int <= Rune
func TestCheckBinaryExprIntLeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 <= '@'`, env, (4 <= '@'), ConstBool)
}

// Test Int <= Float
func TestCheckBinaryExprIntLeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 <= 2.0`, env, (4 <= 2.0), ConstBool)
}

// Test Int <= Complex
func TestCheckBinaryExprIntLeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 <= 8.0i`, env,
		`illegal constant expression: untyped number <= untyped number`,
	)

}

// Test Int <= Bool
func TestCheckBinaryExprIntLeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 <= true`, env,
		`cannot convert true to type int`,
		`invalid operation: 4 <= true (mismatched types int and bool)`,
	)

}

// Test Int <= String
func TestCheckBinaryExprIntLeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 <= "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 4 <= "abc" (mismatched types int and string)`,
	)

}

// Test Int <= Nil
func TestCheckBinaryExprIntLeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 <= nil`, env,
		`cannot convert nil to type int`,
		`invalid operation: 4 <= nil (mismatched types int and <T>)`,
	)

}

// Test Int >= Int
func TestCheckBinaryExprIntGeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 >= 4`, env, (4 >= 4), ConstBool)
}

// Test Int >= Rune
func TestCheckBinaryExprIntGeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 >= '@'`, env, (4 >= '@'), ConstBool)
}

// Test Int >= Float
func TestCheckBinaryExprIntGeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 >= 2.0`, env, (4 >= 2.0), ConstBool)
}

// Test Int >= Complex
func TestCheckBinaryExprIntGeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 >= 8.0i`, env,
		`illegal constant expression: untyped number >= untyped number`,
	)

}

// Test Int >= Bool
func TestCheckBinaryExprIntGeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 >= true`, env,
		`cannot convert true to type int`,
		`invalid operation: 4 >= true (mismatched types int and bool)`,
	)

}

// Test Int >= String
func TestCheckBinaryExprIntGeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 >= "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 4 >= "abc" (mismatched types int and string)`,
	)

}

// Test Int >= Nil
func TestCheckBinaryExprIntGeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 >= nil`, env,
		`cannot convert nil to type int`,
		`invalid operation: 4 >= nil (mismatched types int and <T>)`,
	)

}

// Test Int < Int
func TestCheckBinaryExprIntLssInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 < 4`, env, (4 < 4), ConstBool)
}

// Test Int < Rune
func TestCheckBinaryExprIntLssRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 < '@'`, env, (4 < '@'), ConstBool)
}

// Test Int < Float
func TestCheckBinaryExprIntLssFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 < 2.0`, env, (4 < 2.0), ConstBool)
}

// Test Int < Complex
func TestCheckBinaryExprIntLssComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 < 8.0i`, env,
		`illegal constant expression: untyped number < untyped number`,
	)

}

// Test Int < Bool
func TestCheckBinaryExprIntLssBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 < true`, env,
		`cannot convert true to type int`,
		`invalid operation: 4 < true (mismatched types int and bool)`,
	)

}

// Test Int < String
func TestCheckBinaryExprIntLssString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 < "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 4 < "abc" (mismatched types int and string)`,
	)

}

// Test Int < Nil
func TestCheckBinaryExprIntLssNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 < nil`, env,
		`cannot convert nil to type int`,
		`invalid operation: 4 < nil (mismatched types int and <T>)`,
	)

}

// Test Int > Int
func TestCheckBinaryExprIntGtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 > 4`, env, (4 > 4), ConstBool)
}

// Test Int > Rune
func TestCheckBinaryExprIntGtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 > '@'`, env, (4 > '@'), ConstBool)
}

// Test Int > Float
func TestCheckBinaryExprIntGtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 > 2.0`, env, (4 > 2.0), ConstBool)
}

// Test Int > Complex
func TestCheckBinaryExprIntGtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 > 8.0i`, env,
		`illegal constant expression: untyped number > untyped number`,
	)

}

// Test Int > Bool
func TestCheckBinaryExprIntGtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 > true`, env,
		`cannot convert true to type int`,
		`invalid operation: 4 > true (mismatched types int and bool)`,
	)

}

// Test Int > String
func TestCheckBinaryExprIntGtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 > "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: 4 > "abc" (mismatched types int and string)`,
	)

}

// Test Int > Nil
func TestCheckBinaryExprIntGtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 > nil`, env,
		`cannot convert nil to type int`,
		`invalid operation: 4 > nil (mismatched types int and <T>)`,
	)

}

// Test Int >> Int
func TestCheckBinaryExprIntRhlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 >> 4`, env, NewConstInt64(4 >> 4), ConstInt)
}

// Test Int >> Rune
func TestCheckBinaryExprIntRhlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 >> '@'`, env, NewConstInt64(4 >> '@'), ConstInt)
}

// Test Int >> Float
func TestCheckBinaryExprIntRhlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `4 >> 2.0`, env, NewConstInt64(4 >> 2.0), ConstInt)
}

// Test Int >> Complex
func TestCheckBinaryExprIntRhlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 >> 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int >> Bool
func TestCheckBinaryExprIntRhlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 >> true`, env,
		`invalid operation: 4 >> true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Int >> String
func TestCheckBinaryExprIntRhlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 >> "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: 4 >> "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Int >> Nil
func TestCheckBinaryExprIntRhlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `4 >> nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Rune + Int
func TestCheckBinaryExprRuneAddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' + 4`, env, NewConstRune('@' + 4), ConstRune)
}

// Test Rune + Rune
func TestCheckBinaryExprRuneAddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' + '@'`, env, NewConstRune('@' + '@'), ConstRune)
}

// Test Rune + Float
func TestCheckBinaryExprRuneAddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' + 2.0`, env, NewConstFloat64('@' + 2.0), ConstFloat)
}

// Test Rune + Complex
func TestCheckBinaryExprRuneAddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' + 8.0i`, env, NewConstComplex128('@' + 8.0i), ConstComplex)
}

// Test Rune + Bool
func TestCheckBinaryExprRuneAddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' + true`, env,
		`cannot convert true to type rune`,
		`invalid operation: '@' + true (mismatched types rune and bool)`,
	)

}

// Test Rune + String
func TestCheckBinaryExprRuneAddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' + "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: '@' + "abc" (mismatched types rune and string)`,
	)

}

// Test Rune + Nil
func TestCheckBinaryExprRuneAddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' + nil`, env,
		`cannot convert nil to type rune`,
		`invalid operation: '@' + nil (mismatched types rune and <T>)`,
	)

}

// Test Rune - Int
func TestCheckBinaryExprRuneSubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' - 4`, env, NewConstRune('@' - 4), ConstRune)
}

// Test Rune - Rune
func TestCheckBinaryExprRuneSubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' - '@'`, env, NewConstRune('@' - '@'), ConstRune)
}

// Test Rune - Float
func TestCheckBinaryExprRuneSubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' - 2.0`, env, NewConstFloat64('@' - 2.0), ConstFloat)
}

// Test Rune - Complex
func TestCheckBinaryExprRuneSubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' - 8.0i`, env, NewConstComplex128('@' - 8.0i), ConstComplex)
}

// Test Rune - Bool
func TestCheckBinaryExprRuneSubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' - true`, env,
		`cannot convert true to type rune`,
		`invalid operation: '@' - true (mismatched types rune and bool)`,
	)

}

// Test Rune - String
func TestCheckBinaryExprRuneSubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' - "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: '@' - "abc" (mismatched types rune and string)`,
	)

}

// Test Rune - Nil
func TestCheckBinaryExprRuneSubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' - nil`, env,
		`cannot convert nil to type rune`,
		`invalid operation: '@' - nil (mismatched types rune and <T>)`,
	)

}

// Test Rune * Int
func TestCheckBinaryExprRuneMulInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' * 4`, env, NewConstRune('@' * 4), ConstRune)
}

// Test Rune * Rune
func TestCheckBinaryExprRuneMulRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' * '@'`, env, NewConstRune('@' * '@'), ConstRune)
}

// Test Rune * Float
func TestCheckBinaryExprRuneMulFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' * 2.0`, env, NewConstFloat64('@' * 2.0), ConstFloat)
}

// Test Rune * Complex
func TestCheckBinaryExprRuneMulComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' * 8.0i`, env, NewConstComplex128('@' * 8.0i), ConstComplex)
}

// Test Rune * Bool
func TestCheckBinaryExprRuneMulBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' * true`, env,
		`cannot convert true to type rune`,
		`invalid operation: '@' * true (mismatched types rune and bool)`,
	)

}

// Test Rune * String
func TestCheckBinaryExprRuneMulString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' * "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: '@' * "abc" (mismatched types rune and string)`,
	)

}

// Test Rune * Nil
func TestCheckBinaryExprRuneMulNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' * nil`, env,
		`cannot convert nil to type rune`,
		`invalid operation: '@' * nil (mismatched types rune and <T>)`,
	)

}

// Test Rune / Int
func TestCheckBinaryExprRuneQuoInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' / 4`, env, NewConstRune('@' / 4), ConstRune)
}

// Test Rune / Rune
func TestCheckBinaryExprRuneQuoRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' / '@'`, env, NewConstRune('@' / '@'), ConstRune)
}

// Test Rune / Float
func TestCheckBinaryExprRuneQuoFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' / 2.0`, env, NewConstFloat64('@' / 2.0), ConstFloat)
}

// Test Rune / Complex
func TestCheckBinaryExprRuneQuoComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' / 8.0i`, env, NewConstComplex128('@' / 8.0i), ConstComplex)
}

// Test Rune / Bool
func TestCheckBinaryExprRuneQuoBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' / true`, env,
		`cannot convert true to type rune`,
		`invalid operation: '@' / true (mismatched types rune and bool)`,
	)

}

// Test Rune / String
func TestCheckBinaryExprRuneQuoString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' / "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: '@' / "abc" (mismatched types rune and string)`,
	)

}

// Test Rune / Nil
func TestCheckBinaryExprRuneQuoNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' / nil`, env,
		`cannot convert nil to type rune`,
		`invalid operation: '@' / nil (mismatched types rune and <T>)`,
	)

}

// Test Rune % Int
func TestCheckBinaryExprRuneRemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' % 4`, env, NewConstRune('@' % 4), ConstRune)
}

// Test Rune % Rune
func TestCheckBinaryExprRuneRemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' % '@'`, env, NewConstRune('@' % '@'), ConstRune)
}

// Test Rune % Float
func TestCheckBinaryExprRuneRemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' % 2.0`, env,
		`illegal constant expression: floating-point % operation`,
	)

}

// Test Rune % Complex
func TestCheckBinaryExprRuneRemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' % 8.0i`, env,
		`illegal constant expression: untyped number % untyped number`,
	)

}

// Test Rune % Bool
func TestCheckBinaryExprRuneRemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' % true`, env,
		`cannot convert true to type rune`,
		`invalid operation: '@' % true (mismatched types rune and bool)`,
	)

}

// Test Rune % String
func TestCheckBinaryExprRuneRemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' % "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: '@' % "abc" (mismatched types rune and string)`,
	)

}

// Test Rune % Nil
func TestCheckBinaryExprRuneRemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' % nil`, env,
		`cannot convert nil to type rune`,
		`invalid operation: '@' % nil (mismatched types rune and <T>)`,
	)

}

// Test Rune & Int
func TestCheckBinaryExprRuneAndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' & 4`, env, NewConstRune('@' & 4), ConstRune)
}

// Test Rune & Rune
func TestCheckBinaryExprRuneAndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' & '@'`, env, NewConstRune('@' & '@'), ConstRune)
}

// Test Rune & Float
func TestCheckBinaryExprRuneAndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' & 2.0`, env,
		`illegal constant expression: untyped number & untyped number`,
	)

}

// Test Rune & Complex
func TestCheckBinaryExprRuneAndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' & 8.0i`, env,
		`illegal constant expression: untyped number & untyped number`,
	)

}

// Test Rune & Bool
func TestCheckBinaryExprRuneAndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' & true`, env,
		`cannot convert true to type rune`,
		`invalid operation: '@' & true (mismatched types rune and bool)`,
	)

}

// Test Rune & String
func TestCheckBinaryExprRuneAndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' & "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: '@' & "abc" (mismatched types rune and string)`,
	)

}

// Test Rune & Nil
func TestCheckBinaryExprRuneAndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' & nil`, env,
		`cannot convert nil to type rune`,
		`invalid operation: '@' & nil (mismatched types rune and <T>)`,
	)

}

// Test Rune | Int
func TestCheckBinaryExprRuneOrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' | 4`, env, NewConstRune('@' | 4), ConstRune)
}

// Test Rune | Rune
func TestCheckBinaryExprRuneOrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' | '@'`, env, NewConstRune('@' | '@'), ConstRune)
}

// Test Rune | Float
func TestCheckBinaryExprRuneOrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' | 2.0`, env,
		`illegal constant expression: untyped number | untyped number`,
	)

}

// Test Rune | Complex
func TestCheckBinaryExprRuneOrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' | 8.0i`, env,
		`illegal constant expression: untyped number | untyped number`,
	)

}

// Test Rune | Bool
func TestCheckBinaryExprRuneOrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' | true`, env,
		`cannot convert true to type rune`,
		`invalid operation: '@' | true (mismatched types rune and bool)`,
	)

}

// Test Rune | String
func TestCheckBinaryExprRuneOrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' | "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: '@' | "abc" (mismatched types rune and string)`,
	)

}

// Test Rune | Nil
func TestCheckBinaryExprRuneOrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' | nil`, env,
		`cannot convert nil to type rune`,
		`invalid operation: '@' | nil (mismatched types rune and <T>)`,
	)

}

// Test Rune ^ Int
func TestCheckBinaryExprRuneXorInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' ^ 4`, env, NewConstRune('@' ^ 4), ConstRune)
}

// Test Rune ^ Rune
func TestCheckBinaryExprRuneXorRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' ^ '@'`, env, NewConstRune('@' ^ '@'), ConstRune)
}

// Test Rune ^ Float
func TestCheckBinaryExprRuneXorFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' ^ 2.0`, env,
		`illegal constant expression: untyped number ^ untyped number`,
	)

}

// Test Rune ^ Complex
func TestCheckBinaryExprRuneXorComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' ^ 8.0i`, env,
		`illegal constant expression: untyped number ^ untyped number`,
	)

}

// Test Rune ^ Bool
func TestCheckBinaryExprRuneXorBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' ^ true`, env,
		`cannot convert true to type rune`,
		`invalid operation: '@' ^ true (mismatched types rune and bool)`,
	)

}

// Test Rune ^ String
func TestCheckBinaryExprRuneXorString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' ^ "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: '@' ^ "abc" (mismatched types rune and string)`,
	)

}

// Test Rune ^ Nil
func TestCheckBinaryExprRuneXorNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' ^ nil`, env,
		`cannot convert nil to type rune`,
		`invalid operation: '@' ^ nil (mismatched types rune and <T>)`,
	)

}

// Test Rune &^ Int
func TestCheckBinaryExprRuneAndNotInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' &^ 4`, env, NewConstRune('@' &^ 4), ConstRune)
}

// Test Rune &^ Rune
func TestCheckBinaryExprRuneAndNotRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' &^ '@'`, env, NewConstRune('@' &^ '@'), ConstRune)
}

// Test Rune &^ Float
func TestCheckBinaryExprRuneAndNotFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' &^ 2.0`, env,
		`illegal constant expression: untyped number &^ untyped number`,
	)

}

// Test Rune &^ Complex
func TestCheckBinaryExprRuneAndNotComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' &^ 8.0i`, env,
		`illegal constant expression: untyped number &^ untyped number`,
	)

}

// Test Rune &^ Bool
func TestCheckBinaryExprRuneAndNotBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' &^ true`, env,
		`cannot convert true to type rune`,
		`invalid operation: '@' &^ true (mismatched types rune and bool)`,
	)

}

// Test Rune &^ String
func TestCheckBinaryExprRuneAndNotString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' &^ "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: '@' &^ "abc" (mismatched types rune and string)`,
	)

}

// Test Rune &^ Nil
func TestCheckBinaryExprRuneAndNotNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' &^ nil`, env,
		`cannot convert nil to type rune`,
		`invalid operation: '@' &^ nil (mismatched types rune and <T>)`,
	)

}

// Test Rune == Int
func TestCheckBinaryExprRuneEqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' == 4`, env, ('@' == 4), ConstBool)
}

// Test Rune == Rune
func TestCheckBinaryExprRuneEqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' == '@'`, env, ('@' == '@'), ConstBool)
}

// Test Rune == Float
func TestCheckBinaryExprRuneEqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' == 2.0`, env, ('@' == 2.0), ConstBool)
}

// Test Rune == Complex
func TestCheckBinaryExprRuneEqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' == 8.0i`, env, ('@' == 8.0i), ConstBool)
}

// Test Rune == Bool
func TestCheckBinaryExprRuneEqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' == true`, env,
		`cannot convert true to type rune`,
		`invalid operation: '@' == true (mismatched types rune and bool)`,
	)

}

// Test Rune == String
func TestCheckBinaryExprRuneEqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' == "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: '@' == "abc" (mismatched types rune and string)`,
	)

}

// Test Rune == Nil
func TestCheckBinaryExprRuneEqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' == nil`, env,
		`cannot convert nil to type rune`,
		`invalid operation: '@' == nil (mismatched types rune and <T>)`,
	)

}

// Test Rune != Int
func TestCheckBinaryExprRuneNeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' != 4`, env, ('@' != 4), ConstBool)
}

// Test Rune != Rune
func TestCheckBinaryExprRuneNeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' != '@'`, env, ('@' != '@'), ConstBool)
}

// Test Rune != Float
func TestCheckBinaryExprRuneNeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' != 2.0`, env, ('@' != 2.0), ConstBool)
}

// Test Rune != Complex
func TestCheckBinaryExprRuneNeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' != 8.0i`, env, ('@' != 8.0i), ConstBool)
}

// Test Rune != Bool
func TestCheckBinaryExprRuneNeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' != true`, env,
		`cannot convert true to type rune`,
		`invalid operation: '@' != true (mismatched types rune and bool)`,
	)

}

// Test Rune != String
func TestCheckBinaryExprRuneNeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' != "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: '@' != "abc" (mismatched types rune and string)`,
	)

}

// Test Rune != Nil
func TestCheckBinaryExprRuneNeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' != nil`, env,
		`cannot convert nil to type rune`,
		`invalid operation: '@' != nil (mismatched types rune and <T>)`,
	)

}

// Test Rune <= Int
func TestCheckBinaryExprRuneLeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' <= 4`, env, ('@' <= 4), ConstBool)
}

// Test Rune <= Rune
func TestCheckBinaryExprRuneLeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' <= '@'`, env, ('@' <= '@'), ConstBool)
}

// Test Rune <= Float
func TestCheckBinaryExprRuneLeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' <= 2.0`, env, ('@' <= 2.0), ConstBool)
}

// Test Rune <= Complex
func TestCheckBinaryExprRuneLeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' <= 8.0i`, env,
		`illegal constant expression: untyped number <= untyped number`,
	)

}

// Test Rune <= Bool
func TestCheckBinaryExprRuneLeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' <= true`, env,
		`cannot convert true to type rune`,
		`invalid operation: '@' <= true (mismatched types rune and bool)`,
	)

}

// Test Rune <= String
func TestCheckBinaryExprRuneLeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' <= "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: '@' <= "abc" (mismatched types rune and string)`,
	)

}

// Test Rune <= Nil
func TestCheckBinaryExprRuneLeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' <= nil`, env,
		`cannot convert nil to type rune`,
		`invalid operation: '@' <= nil (mismatched types rune and <T>)`,
	)

}

// Test Rune >= Int
func TestCheckBinaryExprRuneGeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' >= 4`, env, ('@' >= 4), ConstBool)
}

// Test Rune >= Rune
func TestCheckBinaryExprRuneGeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' >= '@'`, env, ('@' >= '@'), ConstBool)
}

// Test Rune >= Float
func TestCheckBinaryExprRuneGeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' >= 2.0`, env, ('@' >= 2.0), ConstBool)
}

// Test Rune >= Complex
func TestCheckBinaryExprRuneGeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' >= 8.0i`, env,
		`illegal constant expression: untyped number >= untyped number`,
	)

}

// Test Rune >= Bool
func TestCheckBinaryExprRuneGeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' >= true`, env,
		`cannot convert true to type rune`,
		`invalid operation: '@' >= true (mismatched types rune and bool)`,
	)

}

// Test Rune >= String
func TestCheckBinaryExprRuneGeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' >= "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: '@' >= "abc" (mismatched types rune and string)`,
	)

}

// Test Rune >= Nil
func TestCheckBinaryExprRuneGeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' >= nil`, env,
		`cannot convert nil to type rune`,
		`invalid operation: '@' >= nil (mismatched types rune and <T>)`,
	)

}

// Test Rune < Int
func TestCheckBinaryExprRuneLssInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' < 4`, env, ('@' < 4), ConstBool)
}

// Test Rune < Rune
func TestCheckBinaryExprRuneLssRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' < '@'`, env, ('@' < '@'), ConstBool)
}

// Test Rune < Float
func TestCheckBinaryExprRuneLssFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' < 2.0`, env, ('@' < 2.0), ConstBool)
}

// Test Rune < Complex
func TestCheckBinaryExprRuneLssComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' < 8.0i`, env,
		`illegal constant expression: untyped number < untyped number`,
	)

}

// Test Rune < Bool
func TestCheckBinaryExprRuneLssBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' < true`, env,
		`cannot convert true to type rune`,
		`invalid operation: '@' < true (mismatched types rune and bool)`,
	)

}

// Test Rune < String
func TestCheckBinaryExprRuneLssString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' < "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: '@' < "abc" (mismatched types rune and string)`,
	)

}

// Test Rune < Nil
func TestCheckBinaryExprRuneLssNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' < nil`, env,
		`cannot convert nil to type rune`,
		`invalid operation: '@' < nil (mismatched types rune and <T>)`,
	)

}

// Test Rune > Int
func TestCheckBinaryExprRuneGtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' > 4`, env, ('@' > 4), ConstBool)
}

// Test Rune > Rune
func TestCheckBinaryExprRuneGtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' > '@'`, env, ('@' > '@'), ConstBool)
}

// Test Rune > Float
func TestCheckBinaryExprRuneGtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' > 2.0`, env, ('@' > 2.0), ConstBool)
}

// Test Rune > Complex
func TestCheckBinaryExprRuneGtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' > 8.0i`, env,
		`illegal constant expression: untyped number > untyped number`,
	)

}

// Test Rune > Bool
func TestCheckBinaryExprRuneGtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' > true`, env,
		`cannot convert true to type rune`,
		`invalid operation: '@' > true (mismatched types rune and bool)`,
	)

}

// Test Rune > String
func TestCheckBinaryExprRuneGtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' > "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: '@' > "abc" (mismatched types rune and string)`,
	)

}

// Test Rune > Nil
func TestCheckBinaryExprRuneGtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' > nil`, env,
		`cannot convert nil to type rune`,
		`invalid operation: '@' > nil (mismatched types rune and <T>)`,
	)

}

// Test Rune >> Int
func TestCheckBinaryExprRuneRhlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' >> 4`, env, NewConstInt64('@' >> 4), ConstInt)
}

// Test Rune >> Rune
func TestCheckBinaryExprRuneRhlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' >> '@'`, env, NewConstInt64('@' >> '@'), ConstInt)
}

// Test Rune >> Float
func TestCheckBinaryExprRuneRhlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `'@' >> 2.0`, env, NewConstInt64('@' >> 2.0), ConstInt)
}

// Test Rune >> Complex
func TestCheckBinaryExprRuneRhlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' >> 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Rune >> Bool
func TestCheckBinaryExprRuneRhlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' >> true`, env,
		`invalid operation: '@' >> true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Rune >> String
func TestCheckBinaryExprRuneRhlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' >> "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: '@' >> "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Rune >> Nil
func TestCheckBinaryExprRuneRhlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `'@' >> nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Float + Int
func TestCheckBinaryExprFloatAddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 + 4`, env, NewConstFloat64(2.0 + 4), ConstFloat)
}

// Test Float + Rune
func TestCheckBinaryExprFloatAddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 + '@'`, env, NewConstFloat64(2.0 + '@'), ConstFloat)
}

// Test Float + Float
func TestCheckBinaryExprFloatAddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 + 2.0`, env, NewConstFloat64(2.0 + 2.0), ConstFloat)
}

// Test Float + Complex
func TestCheckBinaryExprFloatAddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 + 8.0i`, env, NewConstComplex128(2.0 + 8.0i), ConstComplex)
}

// Test Float + Bool
func TestCheckBinaryExprFloatAddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 + true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 2 + true (mismatched types float64 and bool)`,
	)

}

// Test Float + String
func TestCheckBinaryExprFloatAddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 + "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 2 + "abc" (mismatched types float64 and string)`,
	)

}

// Test Float + Nil
func TestCheckBinaryExprFloatAddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 + nil`, env,
		`cannot convert nil to type float64`,
		`invalid operation: 2 + nil (mismatched types float64 and <T>)`,
	)

}

// Test Float - Int
func TestCheckBinaryExprFloatSubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 - 4`, env, NewConstFloat64(2.0 - 4), ConstFloat)
}

// Test Float - Rune
func TestCheckBinaryExprFloatSubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 - '@'`, env, NewConstFloat64(2.0 - '@'), ConstFloat)
}

// Test Float - Float
func TestCheckBinaryExprFloatSubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 - 2.0`, env, NewConstFloat64(2.0 - 2.0), ConstFloat)
}

// Test Float - Complex
func TestCheckBinaryExprFloatSubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 - 8.0i`, env, NewConstComplex128(2.0 - 8.0i), ConstComplex)
}

// Test Float - Bool
func TestCheckBinaryExprFloatSubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 - true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 2 - true (mismatched types float64 and bool)`,
	)

}

// Test Float - String
func TestCheckBinaryExprFloatSubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 - "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 2 - "abc" (mismatched types float64 and string)`,
	)

}

// Test Float - Nil
func TestCheckBinaryExprFloatSubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 - nil`, env,
		`cannot convert nil to type float64`,
		`invalid operation: 2 - nil (mismatched types float64 and <T>)`,
	)

}

// Test Float * Int
func TestCheckBinaryExprFloatMulInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 * 4`, env, NewConstFloat64(2.0 * 4), ConstFloat)
}

// Test Float * Rune
func TestCheckBinaryExprFloatMulRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 * '@'`, env, NewConstFloat64(2.0 * '@'), ConstFloat)
}

// Test Float * Float
func TestCheckBinaryExprFloatMulFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 * 2.0`, env, NewConstFloat64(2.0 * 2.0), ConstFloat)
}

// Test Float * Complex
func TestCheckBinaryExprFloatMulComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 * 8.0i`, env, NewConstComplex128(2.0 * 8.0i), ConstComplex)
}

// Test Float * Bool
func TestCheckBinaryExprFloatMulBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 * true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 2 * true (mismatched types float64 and bool)`,
	)

}

// Test Float * String
func TestCheckBinaryExprFloatMulString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 * "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 2 * "abc" (mismatched types float64 and string)`,
	)

}

// Test Float * Nil
func TestCheckBinaryExprFloatMulNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 * nil`, env,
		`cannot convert nil to type float64`,
		`invalid operation: 2 * nil (mismatched types float64 and <T>)`,
	)

}

// Test Float / Int
func TestCheckBinaryExprFloatQuoInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 / 4`, env, NewConstFloat64(2.0 / 4), ConstFloat)
}

// Test Float / Rune
func TestCheckBinaryExprFloatQuoRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 / '@'`, env, NewConstFloat64(2.0 / '@'), ConstFloat)
}

// Test Float / Float
func TestCheckBinaryExprFloatQuoFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 / 2.0`, env, NewConstFloat64(2.0 / 2.0), ConstFloat)
}

// Test Float / Complex
func TestCheckBinaryExprFloatQuoComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 / 8.0i`, env, NewConstComplex128(2.0 / 8.0i), ConstComplex)
}

// Test Float / Bool
func TestCheckBinaryExprFloatQuoBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 / true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 2 / true (mismatched types float64 and bool)`,
	)

}

// Test Float / String
func TestCheckBinaryExprFloatQuoString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 / "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 2 / "abc" (mismatched types float64 and string)`,
	)

}

// Test Float / Nil
func TestCheckBinaryExprFloatQuoNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 / nil`, env,
		`cannot convert nil to type float64`,
		`invalid operation: 2 / nil (mismatched types float64 and <T>)`,
	)

}

// Test Float % Int
func TestCheckBinaryExprFloatRemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 % 4`, env,
		`illegal constant expression: floating-point % operation`,
	)

}

// Test Float % Rune
func TestCheckBinaryExprFloatRemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 % '@'`, env,
		`illegal constant expression: floating-point % operation`,
	)

}

// Test Float % Float
func TestCheckBinaryExprFloatRemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 % 2.0`, env,
		`illegal constant expression: floating-point % operation`,
	)

}

// Test Float % Complex
func TestCheckBinaryExprFloatRemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 % 8.0i`, env,
		`illegal constant expression: untyped number % untyped number`,
	)

}

// Test Float % Bool
func TestCheckBinaryExprFloatRemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 % true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 2 % true (mismatched types float64 and bool)`,
	)

}

// Test Float % String
func TestCheckBinaryExprFloatRemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 % "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 2 % "abc" (mismatched types float64 and string)`,
	)

}

// Test Float % Nil
func TestCheckBinaryExprFloatRemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 % nil`, env,
		`cannot convert nil to type float64`,
		`invalid operation: 2 % nil (mismatched types float64 and <T>)`,
	)

}

// Test Float & Int
func TestCheckBinaryExprFloatAndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 & 4`, env,
		`illegal constant expression: untyped number & untyped number`,
	)

}

// Test Float & Rune
func TestCheckBinaryExprFloatAndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 & '@'`, env,
		`illegal constant expression: untyped number & untyped number`,
	)

}

// Test Float & Float
func TestCheckBinaryExprFloatAndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 & 2.0`, env,
		`illegal constant expression: untyped number & untyped number`,
	)

}

// Test Float & Complex
func TestCheckBinaryExprFloatAndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 & 8.0i`, env,
		`illegal constant expression: untyped number & untyped number`,
	)

}

// Test Float & Bool
func TestCheckBinaryExprFloatAndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 & true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 2 & true (mismatched types float64 and bool)`,
	)

}

// Test Float & String
func TestCheckBinaryExprFloatAndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 & "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 2 & "abc" (mismatched types float64 and string)`,
	)

}

// Test Float & Nil
func TestCheckBinaryExprFloatAndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 & nil`, env,
		`cannot convert nil to type float64`,
		`invalid operation: 2 & nil (mismatched types float64 and <T>)`,
	)

}

// Test Float | Int
func TestCheckBinaryExprFloatOrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 | 4`, env,
		`illegal constant expression: untyped number | untyped number`,
	)

}

// Test Float | Rune
func TestCheckBinaryExprFloatOrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 | '@'`, env,
		`illegal constant expression: untyped number | untyped number`,
	)

}

// Test Float | Float
func TestCheckBinaryExprFloatOrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 | 2.0`, env,
		`illegal constant expression: untyped number | untyped number`,
	)

}

// Test Float | Complex
func TestCheckBinaryExprFloatOrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 | 8.0i`, env,
		`illegal constant expression: untyped number | untyped number`,
	)

}

// Test Float | Bool
func TestCheckBinaryExprFloatOrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 | true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 2 | true (mismatched types float64 and bool)`,
	)

}

// Test Float | String
func TestCheckBinaryExprFloatOrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 | "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 2 | "abc" (mismatched types float64 and string)`,
	)

}

// Test Float | Nil
func TestCheckBinaryExprFloatOrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 | nil`, env,
		`cannot convert nil to type float64`,
		`invalid operation: 2 | nil (mismatched types float64 and <T>)`,
	)

}

// Test Float ^ Int
func TestCheckBinaryExprFloatXorInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 ^ 4`, env,
		`illegal constant expression: untyped number ^ untyped number`,
	)

}

// Test Float ^ Rune
func TestCheckBinaryExprFloatXorRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 ^ '@'`, env,
		`illegal constant expression: untyped number ^ untyped number`,
	)

}

// Test Float ^ Float
func TestCheckBinaryExprFloatXorFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 ^ 2.0`, env,
		`illegal constant expression: untyped number ^ untyped number`,
	)

}

// Test Float ^ Complex
func TestCheckBinaryExprFloatXorComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 ^ 8.0i`, env,
		`illegal constant expression: untyped number ^ untyped number`,
	)

}

// Test Float ^ Bool
func TestCheckBinaryExprFloatXorBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 ^ true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 2 ^ true (mismatched types float64 and bool)`,
	)

}

// Test Float ^ String
func TestCheckBinaryExprFloatXorString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 ^ "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 2 ^ "abc" (mismatched types float64 and string)`,
	)

}

// Test Float ^ Nil
func TestCheckBinaryExprFloatXorNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 ^ nil`, env,
		`cannot convert nil to type float64`,
		`invalid operation: 2 ^ nil (mismatched types float64 and <T>)`,
	)

}

// Test Float &^ Int
func TestCheckBinaryExprFloatAndNotInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 &^ 4`, env,
		`illegal constant expression: untyped number &^ untyped number`,
	)

}

// Test Float &^ Rune
func TestCheckBinaryExprFloatAndNotRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 &^ '@'`, env,
		`illegal constant expression: untyped number &^ untyped number`,
	)

}

// Test Float &^ Float
func TestCheckBinaryExprFloatAndNotFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 &^ 2.0`, env,
		`illegal constant expression: untyped number &^ untyped number`,
	)

}

// Test Float &^ Complex
func TestCheckBinaryExprFloatAndNotComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 &^ 8.0i`, env,
		`illegal constant expression: untyped number &^ untyped number`,
	)

}

// Test Float &^ Bool
func TestCheckBinaryExprFloatAndNotBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 &^ true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 2 &^ true (mismatched types float64 and bool)`,
	)

}

// Test Float &^ String
func TestCheckBinaryExprFloatAndNotString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 &^ "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 2 &^ "abc" (mismatched types float64 and string)`,
	)

}

// Test Float &^ Nil
func TestCheckBinaryExprFloatAndNotNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 &^ nil`, env,
		`cannot convert nil to type float64`,
		`invalid operation: 2 &^ nil (mismatched types float64 and <T>)`,
	)

}

// Test Float == Int
func TestCheckBinaryExprFloatEqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 == 4`, env, (2.0 == 4), ConstBool)
}

// Test Float == Rune
func TestCheckBinaryExprFloatEqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 == '@'`, env, (2.0 == '@'), ConstBool)
}

// Test Float == Float
func TestCheckBinaryExprFloatEqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 == 2.0`, env, (2.0 == 2.0), ConstBool)
}

// Test Float == Complex
func TestCheckBinaryExprFloatEqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 == 8.0i`, env, (2.0 == 8.0i), ConstBool)
}

// Test Float == Bool
func TestCheckBinaryExprFloatEqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 == true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 2 == true (mismatched types float64 and bool)`,
	)

}

// Test Float == String
func TestCheckBinaryExprFloatEqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 == "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 2 == "abc" (mismatched types float64 and string)`,
	)

}

// Test Float == Nil
func TestCheckBinaryExprFloatEqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 == nil`, env,
		`cannot convert nil to type float64`,
		`invalid operation: 2 == nil (mismatched types float64 and <T>)`,
	)

}

// Test Float != Int
func TestCheckBinaryExprFloatNeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 != 4`, env, (2.0 != 4), ConstBool)
}

// Test Float != Rune
func TestCheckBinaryExprFloatNeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 != '@'`, env, (2.0 != '@'), ConstBool)
}

// Test Float != Float
func TestCheckBinaryExprFloatNeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 != 2.0`, env, (2.0 != 2.0), ConstBool)
}

// Test Float != Complex
func TestCheckBinaryExprFloatNeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 != 8.0i`, env, (2.0 != 8.0i), ConstBool)
}

// Test Float != Bool
func TestCheckBinaryExprFloatNeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 != true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 2 != true (mismatched types float64 and bool)`,
	)

}

// Test Float != String
func TestCheckBinaryExprFloatNeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 != "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 2 != "abc" (mismatched types float64 and string)`,
	)

}

// Test Float != Nil
func TestCheckBinaryExprFloatNeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 != nil`, env,
		`cannot convert nil to type float64`,
		`invalid operation: 2 != nil (mismatched types float64 and <T>)`,
	)

}

// Test Float <= Int
func TestCheckBinaryExprFloatLeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 <= 4`, env, (2.0 <= 4), ConstBool)
}

// Test Float <= Rune
func TestCheckBinaryExprFloatLeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 <= '@'`, env, (2.0 <= '@'), ConstBool)
}

// Test Float <= Float
func TestCheckBinaryExprFloatLeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 <= 2.0`, env, (2.0 <= 2.0), ConstBool)
}

// Test Float <= Complex
func TestCheckBinaryExprFloatLeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 <= 8.0i`, env,
		`illegal constant expression: untyped number <= untyped number`,
	)

}

// Test Float <= Bool
func TestCheckBinaryExprFloatLeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 <= true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 2 <= true (mismatched types float64 and bool)`,
	)

}

// Test Float <= String
func TestCheckBinaryExprFloatLeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 <= "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 2 <= "abc" (mismatched types float64 and string)`,
	)

}

// Test Float <= Nil
func TestCheckBinaryExprFloatLeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 <= nil`, env,
		`cannot convert nil to type float64`,
		`invalid operation: 2 <= nil (mismatched types float64 and <T>)`,
	)

}

// Test Float >= Int
func TestCheckBinaryExprFloatGeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 >= 4`, env, (2.0 >= 4), ConstBool)
}

// Test Float >= Rune
func TestCheckBinaryExprFloatGeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 >= '@'`, env, (2.0 >= '@'), ConstBool)
}

// Test Float >= Float
func TestCheckBinaryExprFloatGeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 >= 2.0`, env, (2.0 >= 2.0), ConstBool)
}

// Test Float >= Complex
func TestCheckBinaryExprFloatGeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 >= 8.0i`, env,
		`illegal constant expression: untyped number >= untyped number`,
	)

}

// Test Float >= Bool
func TestCheckBinaryExprFloatGeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 >= true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 2 >= true (mismatched types float64 and bool)`,
	)

}

// Test Float >= String
func TestCheckBinaryExprFloatGeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 >= "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 2 >= "abc" (mismatched types float64 and string)`,
	)

}

// Test Float >= Nil
func TestCheckBinaryExprFloatGeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 >= nil`, env,
		`cannot convert nil to type float64`,
		`invalid operation: 2 >= nil (mismatched types float64 and <T>)`,
	)

}

// Test Float < Int
func TestCheckBinaryExprFloatLssInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 < 4`, env, (2.0 < 4), ConstBool)
}

// Test Float < Rune
func TestCheckBinaryExprFloatLssRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 < '@'`, env, (2.0 < '@'), ConstBool)
}

// Test Float < Float
func TestCheckBinaryExprFloatLssFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 < 2.0`, env, (2.0 < 2.0), ConstBool)
}

// Test Float < Complex
func TestCheckBinaryExprFloatLssComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 < 8.0i`, env,
		`illegal constant expression: untyped number < untyped number`,
	)

}

// Test Float < Bool
func TestCheckBinaryExprFloatLssBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 < true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 2 < true (mismatched types float64 and bool)`,
	)

}

// Test Float < String
func TestCheckBinaryExprFloatLssString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 < "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 2 < "abc" (mismatched types float64 and string)`,
	)

}

// Test Float < Nil
func TestCheckBinaryExprFloatLssNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 < nil`, env,
		`cannot convert nil to type float64`,
		`invalid operation: 2 < nil (mismatched types float64 and <T>)`,
	)

}

// Test Float > Int
func TestCheckBinaryExprFloatGtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 > 4`, env, (2.0 > 4), ConstBool)
}

// Test Float > Rune
func TestCheckBinaryExprFloatGtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 > '@'`, env, (2.0 > '@'), ConstBool)
}

// Test Float > Float
func TestCheckBinaryExprFloatGtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 > 2.0`, env, (2.0 > 2.0), ConstBool)
}

// Test Float > Complex
func TestCheckBinaryExprFloatGtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 > 8.0i`, env,
		`illegal constant expression: untyped number > untyped number`,
	)

}

// Test Float > Bool
func TestCheckBinaryExprFloatGtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 > true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 2 > true (mismatched types float64 and bool)`,
	)

}

// Test Float > String
func TestCheckBinaryExprFloatGtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 > "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 2 > "abc" (mismatched types float64 and string)`,
	)

}

// Test Float > Nil
func TestCheckBinaryExprFloatGtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 > nil`, env,
		`cannot convert nil to type float64`,
		`invalid operation: 2 > nil (mismatched types float64 and <T>)`,
	)

}

// Test Float >> Int
func TestCheckBinaryExprFloatRhlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 >> 4`, env, NewConstInt64(2.0 >> 4), ConstInt)
}

// Test Float >> Rune
func TestCheckBinaryExprFloatRhlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 >> '@'`, env, NewConstInt64(2.0 >> '@'), ConstInt)
}

// Test Float >> Float
func TestCheckBinaryExprFloatRhlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `2.0 >> 2.0`, env, NewConstInt64(2.0 >> 2.0), ConstInt)
}

// Test Float >> Complex
func TestCheckBinaryExprFloatRhlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 >> 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Float >> Bool
func TestCheckBinaryExprFloatRhlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 >> true`, env,
		`invalid operation: 2 >> true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Float >> String
func TestCheckBinaryExprFloatRhlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 >> "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: 2 >> "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Float >> Nil
func TestCheckBinaryExprFloatRhlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `2.0 >> nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Complex + Int
func TestCheckBinaryExprComplexAddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i + 4`, env, NewConstComplex128(8.0i + 4), ConstComplex)
}

// Test Complex + Rune
func TestCheckBinaryExprComplexAddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i + '@'`, env, NewConstComplex128(8.0i + '@'), ConstComplex)
}

// Test Complex + Float
func TestCheckBinaryExprComplexAddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i + 2.0`, env, NewConstComplex128(8.0i + 2.0), ConstComplex)
}

// Test Complex + Complex
func TestCheckBinaryExprComplexAddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i + 8.0i`, env, NewConstComplex128(8.0i + 8.0i), ConstComplex)
}

// Test Complex + Bool
func TestCheckBinaryExprComplexAddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i + true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: 8i + true (mismatched types complex128 and bool)`,
	)

}

// Test Complex + String
func TestCheckBinaryExprComplexAddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i + "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 8i + "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex + Nil
func TestCheckBinaryExprComplexAddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i + nil`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: 8i + nil (mismatched types complex128 and <T>)`,
	)

}

// Test Complex - Int
func TestCheckBinaryExprComplexSubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i - 4`, env, NewConstComplex128(8.0i - 4), ConstComplex)
}

// Test Complex - Rune
func TestCheckBinaryExprComplexSubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i - '@'`, env, NewConstComplex128(8.0i - '@'), ConstComplex)
}

// Test Complex - Float
func TestCheckBinaryExprComplexSubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i - 2.0`, env, NewConstComplex128(8.0i - 2.0), ConstComplex)
}

// Test Complex - Complex
func TestCheckBinaryExprComplexSubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i - 8.0i`, env, NewConstComplex128(8.0i - 8.0i), ConstComplex)
}

// Test Complex - Bool
func TestCheckBinaryExprComplexSubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i - true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: 8i - true (mismatched types complex128 and bool)`,
	)

}

// Test Complex - String
func TestCheckBinaryExprComplexSubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i - "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 8i - "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex - Nil
func TestCheckBinaryExprComplexSubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i - nil`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: 8i - nil (mismatched types complex128 and <T>)`,
	)

}

// Test Complex * Int
func TestCheckBinaryExprComplexMulInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i * 4`, env, NewConstComplex128(8.0i * 4), ConstComplex)
}

// Test Complex * Rune
func TestCheckBinaryExprComplexMulRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i * '@'`, env, NewConstComplex128(8.0i * '@'), ConstComplex)
}

// Test Complex * Float
func TestCheckBinaryExprComplexMulFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i * 2.0`, env, NewConstComplex128(8.0i * 2.0), ConstComplex)
}

// Test Complex * Complex
func TestCheckBinaryExprComplexMulComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i * 8.0i`, env, NewConstComplex128(8.0i * 8.0i), ConstComplex)
}

// Test Complex * Bool
func TestCheckBinaryExprComplexMulBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i * true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: 8i * true (mismatched types complex128 and bool)`,
	)

}

// Test Complex * String
func TestCheckBinaryExprComplexMulString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i * "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 8i * "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex * Nil
func TestCheckBinaryExprComplexMulNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i * nil`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: 8i * nil (mismatched types complex128 and <T>)`,
	)

}

// Test Complex / Int
func TestCheckBinaryExprComplexQuoInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i / 4`, env, NewConstComplex128(8.0i / 4), ConstComplex)
}

// Test Complex / Rune
func TestCheckBinaryExprComplexQuoRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i / '@'`, env, NewConstComplex128(8.0i / '@'), ConstComplex)
}

// Test Complex / Float
func TestCheckBinaryExprComplexQuoFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i / 2.0`, env, NewConstComplex128(8.0i / 2.0), ConstComplex)
}

// Test Complex / Complex
func TestCheckBinaryExprComplexQuoComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i / 8.0i`, env, NewConstComplex128(8.0i / 8.0i), ConstComplex)
}

// Test Complex / Bool
func TestCheckBinaryExprComplexQuoBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i / true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: 8i / true (mismatched types complex128 and bool)`,
	)

}

// Test Complex / String
func TestCheckBinaryExprComplexQuoString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i / "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 8i / "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex / Nil
func TestCheckBinaryExprComplexQuoNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i / nil`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: 8i / nil (mismatched types complex128 and <T>)`,
	)

}

// Test Complex % Int
func TestCheckBinaryExprComplexRemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i % 4`, env,
		`illegal constant expression: untyped number % untyped number`,
	)

}

// Test Complex % Rune
func TestCheckBinaryExprComplexRemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i % '@'`, env,
		`illegal constant expression: untyped number % untyped number`,
	)

}

// Test Complex % Float
func TestCheckBinaryExprComplexRemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i % 2.0`, env,
		`illegal constant expression: untyped number % untyped number`,
	)

}

// Test Complex % Complex
func TestCheckBinaryExprComplexRemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i % 8.0i`, env,
		`illegal constant expression: untyped number % untyped number`,
	)

}

// Test Complex % Bool
func TestCheckBinaryExprComplexRemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i % true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: 8i % true (mismatched types complex128 and bool)`,
	)

}

// Test Complex % String
func TestCheckBinaryExprComplexRemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i % "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 8i % "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex % Nil
func TestCheckBinaryExprComplexRemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i % nil`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: 8i % nil (mismatched types complex128 and <T>)`,
	)

}

// Test Complex & Int
func TestCheckBinaryExprComplexAndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i & 4`, env,
		`illegal constant expression: untyped number & untyped number`,
	)

}

// Test Complex & Rune
func TestCheckBinaryExprComplexAndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i & '@'`, env,
		`illegal constant expression: untyped number & untyped number`,
	)

}

// Test Complex & Float
func TestCheckBinaryExprComplexAndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i & 2.0`, env,
		`illegal constant expression: untyped number & untyped number`,
	)

}

// Test Complex & Complex
func TestCheckBinaryExprComplexAndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i & 8.0i`, env,
		`illegal constant expression: untyped number & untyped number`,
	)

}

// Test Complex & Bool
func TestCheckBinaryExprComplexAndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i & true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: 8i & true (mismatched types complex128 and bool)`,
	)

}

// Test Complex & String
func TestCheckBinaryExprComplexAndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i & "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 8i & "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex & Nil
func TestCheckBinaryExprComplexAndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i & nil`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: 8i & nil (mismatched types complex128 and <T>)`,
	)

}

// Test Complex | Int
func TestCheckBinaryExprComplexOrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i | 4`, env,
		`illegal constant expression: untyped number | untyped number`,
	)

}

// Test Complex | Rune
func TestCheckBinaryExprComplexOrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i | '@'`, env,
		`illegal constant expression: untyped number | untyped number`,
	)

}

// Test Complex | Float
func TestCheckBinaryExprComplexOrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i | 2.0`, env,
		`illegal constant expression: untyped number | untyped number`,
	)

}

// Test Complex | Complex
func TestCheckBinaryExprComplexOrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i | 8.0i`, env,
		`illegal constant expression: untyped number | untyped number`,
	)

}

// Test Complex | Bool
func TestCheckBinaryExprComplexOrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i | true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: 8i | true (mismatched types complex128 and bool)`,
	)

}

// Test Complex | String
func TestCheckBinaryExprComplexOrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i | "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 8i | "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex | Nil
func TestCheckBinaryExprComplexOrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i | nil`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: 8i | nil (mismatched types complex128 and <T>)`,
	)

}

// Test Complex ^ Int
func TestCheckBinaryExprComplexXorInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i ^ 4`, env,
		`illegal constant expression: untyped number ^ untyped number`,
	)

}

// Test Complex ^ Rune
func TestCheckBinaryExprComplexXorRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i ^ '@'`, env,
		`illegal constant expression: untyped number ^ untyped number`,
	)

}

// Test Complex ^ Float
func TestCheckBinaryExprComplexXorFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i ^ 2.0`, env,
		`illegal constant expression: untyped number ^ untyped number`,
	)

}

// Test Complex ^ Complex
func TestCheckBinaryExprComplexXorComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i ^ 8.0i`, env,
		`illegal constant expression: untyped number ^ untyped number`,
	)

}

// Test Complex ^ Bool
func TestCheckBinaryExprComplexXorBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i ^ true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: 8i ^ true (mismatched types complex128 and bool)`,
	)

}

// Test Complex ^ String
func TestCheckBinaryExprComplexXorString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i ^ "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 8i ^ "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex ^ Nil
func TestCheckBinaryExprComplexXorNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i ^ nil`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: 8i ^ nil (mismatched types complex128 and <T>)`,
	)

}

// Test Complex &^ Int
func TestCheckBinaryExprComplexAndNotInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i &^ 4`, env,
		`illegal constant expression: untyped number &^ untyped number`,
	)

}

// Test Complex &^ Rune
func TestCheckBinaryExprComplexAndNotRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i &^ '@'`, env,
		`illegal constant expression: untyped number &^ untyped number`,
	)

}

// Test Complex &^ Float
func TestCheckBinaryExprComplexAndNotFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i &^ 2.0`, env,
		`illegal constant expression: untyped number &^ untyped number`,
	)

}

// Test Complex &^ Complex
func TestCheckBinaryExprComplexAndNotComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i &^ 8.0i`, env,
		`illegal constant expression: untyped number &^ untyped number`,
	)

}

// Test Complex &^ Bool
func TestCheckBinaryExprComplexAndNotBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i &^ true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: 8i &^ true (mismatched types complex128 and bool)`,
	)

}

// Test Complex &^ String
func TestCheckBinaryExprComplexAndNotString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i &^ "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 8i &^ "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex &^ Nil
func TestCheckBinaryExprComplexAndNotNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i &^ nil`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: 8i &^ nil (mismatched types complex128 and <T>)`,
	)

}

// Test Complex == Int
func TestCheckBinaryExprComplexEqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i == 4`, env, (8.0i == 4), ConstBool)
}

// Test Complex == Rune
func TestCheckBinaryExprComplexEqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i == '@'`, env, (8.0i == '@'), ConstBool)
}

// Test Complex == Float
func TestCheckBinaryExprComplexEqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i == 2.0`, env, (8.0i == 2.0), ConstBool)
}

// Test Complex == Complex
func TestCheckBinaryExprComplexEqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i == 8.0i`, env, (8.0i == 8.0i), ConstBool)
}

// Test Complex == Bool
func TestCheckBinaryExprComplexEqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i == true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: 8i == true (mismatched types complex128 and bool)`,
	)

}

// Test Complex == String
func TestCheckBinaryExprComplexEqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i == "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 8i == "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex == Nil
func TestCheckBinaryExprComplexEqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i == nil`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: 8i == nil (mismatched types complex128 and <T>)`,
	)

}

// Test Complex != Int
func TestCheckBinaryExprComplexNeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i != 4`, env, (8.0i != 4), ConstBool)
}

// Test Complex != Rune
func TestCheckBinaryExprComplexNeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i != '@'`, env, (8.0i != '@'), ConstBool)
}

// Test Complex != Float
func TestCheckBinaryExprComplexNeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i != 2.0`, env, (8.0i != 2.0), ConstBool)
}

// Test Complex != Complex
func TestCheckBinaryExprComplexNeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `8.0i != 8.0i`, env, (8.0i != 8.0i), ConstBool)
}

// Test Complex != Bool
func TestCheckBinaryExprComplexNeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i != true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: 8i != true (mismatched types complex128 and bool)`,
	)

}

// Test Complex != String
func TestCheckBinaryExprComplexNeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i != "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 8i != "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex != Nil
func TestCheckBinaryExprComplexNeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i != nil`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: 8i != nil (mismatched types complex128 and <T>)`,
	)

}

// Test Complex <= Int
func TestCheckBinaryExprComplexLeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i <= 4`, env,
		`illegal constant expression: untyped number <= untyped number`,
	)

}

// Test Complex <= Rune
func TestCheckBinaryExprComplexLeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i <= '@'`, env,
		`illegal constant expression: untyped number <= untyped number`,
	)

}

// Test Complex <= Float
func TestCheckBinaryExprComplexLeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i <= 2.0`, env,
		`illegal constant expression: untyped number <= untyped number`,
	)

}

// Test Complex <= Complex
func TestCheckBinaryExprComplexLeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i <= 8.0i`, env,
		`illegal constant expression: untyped number <= untyped number`,
	)

}

// Test Complex <= Bool
func TestCheckBinaryExprComplexLeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i <= true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: 8i <= true (mismatched types complex128 and bool)`,
	)

}

// Test Complex <= String
func TestCheckBinaryExprComplexLeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i <= "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 8i <= "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex <= Nil
func TestCheckBinaryExprComplexLeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i <= nil`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: 8i <= nil (mismatched types complex128 and <T>)`,
	)

}

// Test Complex >= Int
func TestCheckBinaryExprComplexGeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i >= 4`, env,
		`illegal constant expression: untyped number >= untyped number`,
	)

}

// Test Complex >= Rune
func TestCheckBinaryExprComplexGeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i >= '@'`, env,
		`illegal constant expression: untyped number >= untyped number`,
	)

}

// Test Complex >= Float
func TestCheckBinaryExprComplexGeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i >= 2.0`, env,
		`illegal constant expression: untyped number >= untyped number`,
	)

}

// Test Complex >= Complex
func TestCheckBinaryExprComplexGeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i >= 8.0i`, env,
		`illegal constant expression: untyped number >= untyped number`,
	)

}

// Test Complex >= Bool
func TestCheckBinaryExprComplexGeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i >= true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: 8i >= true (mismatched types complex128 and bool)`,
	)

}

// Test Complex >= String
func TestCheckBinaryExprComplexGeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i >= "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 8i >= "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex >= Nil
func TestCheckBinaryExprComplexGeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i >= nil`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: 8i >= nil (mismatched types complex128 and <T>)`,
	)

}

// Test Complex < Int
func TestCheckBinaryExprComplexLssInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i < 4`, env,
		`illegal constant expression: untyped number < untyped number`,
	)

}

// Test Complex < Rune
func TestCheckBinaryExprComplexLssRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i < '@'`, env,
		`illegal constant expression: untyped number < untyped number`,
	)

}

// Test Complex < Float
func TestCheckBinaryExprComplexLssFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i < 2.0`, env,
		`illegal constant expression: untyped number < untyped number`,
	)

}

// Test Complex < Complex
func TestCheckBinaryExprComplexLssComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i < 8.0i`, env,
		`illegal constant expression: untyped number < untyped number`,
	)

}

// Test Complex < Bool
func TestCheckBinaryExprComplexLssBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i < true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: 8i < true (mismatched types complex128 and bool)`,
	)

}

// Test Complex < String
func TestCheckBinaryExprComplexLssString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i < "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 8i < "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex < Nil
func TestCheckBinaryExprComplexLssNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i < nil`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: 8i < nil (mismatched types complex128 and <T>)`,
	)

}

// Test Complex > Int
func TestCheckBinaryExprComplexGtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i > 4`, env,
		`illegal constant expression: untyped number > untyped number`,
	)

}

// Test Complex > Rune
func TestCheckBinaryExprComplexGtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i > '@'`, env,
		`illegal constant expression: untyped number > untyped number`,
	)

}

// Test Complex > Float
func TestCheckBinaryExprComplexGtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i > 2.0`, env,
		`illegal constant expression: untyped number > untyped number`,
	)

}

// Test Complex > Complex
func TestCheckBinaryExprComplexGtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i > 8.0i`, env,
		`illegal constant expression: untyped number > untyped number`,
	)

}

// Test Complex > Bool
func TestCheckBinaryExprComplexGtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i > true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: 8i > true (mismatched types complex128 and bool)`,
	)

}

// Test Complex > String
func TestCheckBinaryExprComplexGtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i > "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: 8i > "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex > Nil
func TestCheckBinaryExprComplexGtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i > nil`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: 8i > nil (mismatched types complex128 and <T>)`,
	)

}

// Test Complex >> Int
func TestCheckBinaryExprComplexRhlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i >> 4`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Complex >> Rune
func TestCheckBinaryExprComplexRhlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i >> '@'`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Complex >> Float
func TestCheckBinaryExprComplexRhlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i >> 2.0`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Complex >> Complex
func TestCheckBinaryExprComplexRhlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i >> 8.0i`, env,
		`constant 0+8i truncated to real`,
		`constant 0+8i truncated to real`,
	)

}

// Test Complex >> Bool
func TestCheckBinaryExprComplexRhlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i >> true`, env,
		`invalid operation: 8i >> true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Complex >> String
func TestCheckBinaryExprComplexRhlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i >> "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: 8i >> "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Complex >> Nil
func TestCheckBinaryExprComplexRhlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `8.0i >> nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Bool + Int
func TestCheckBinaryExprBoolAddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true + 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true + 4 (mismatched types bool and int)`,
	)

}

// Test Bool + Rune
func TestCheckBinaryExprBoolAddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true + '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true + '@' (mismatched types bool and rune)`,
	)

}

// Test Bool + Float
func TestCheckBinaryExprBoolAddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true + 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true + 2 (mismatched types bool and float64)`,
	)

}

// Test Bool + Complex
func TestCheckBinaryExprBoolAddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true + 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true + 8i (mismatched types bool and complex128)`,
	)

}

// Test Bool + Bool
func TestCheckBinaryExprBoolAddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true + true`, env,
		`invalid operation: true + true (operator + not defined on bool)`,
	)

}

// Test Bool + String
func TestCheckBinaryExprBoolAddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true + "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true + "abc" (mismatched types bool and string)`,
	)

}

// Test Bool + Nil
func TestCheckBinaryExprBoolAddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true + nil`, env,
		`cannot convert nil to type bool`,
		`invalid operation: true + nil (mismatched types bool and <T>)`,
	)

}

// Test Bool - Int
func TestCheckBinaryExprBoolSubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true - 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true - 4 (mismatched types bool and int)`,
	)

}

// Test Bool - Rune
func TestCheckBinaryExprBoolSubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true - '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true - '@' (mismatched types bool and rune)`,
	)

}

// Test Bool - Float
func TestCheckBinaryExprBoolSubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true - 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true - 2 (mismatched types bool and float64)`,
	)

}

// Test Bool - Complex
func TestCheckBinaryExprBoolSubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true - 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true - 8i (mismatched types bool and complex128)`,
	)

}

// Test Bool - Bool
func TestCheckBinaryExprBoolSubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true - true`, env,
		`invalid operation: true - true (operator - not defined on bool)`,
	)

}

// Test Bool - String
func TestCheckBinaryExprBoolSubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true - "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true - "abc" (mismatched types bool and string)`,
	)

}

// Test Bool - Nil
func TestCheckBinaryExprBoolSubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true - nil`, env,
		`cannot convert nil to type bool`,
		`invalid operation: true - nil (mismatched types bool and <T>)`,
	)

}

// Test Bool * Int
func TestCheckBinaryExprBoolMulInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true * 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true * 4 (mismatched types bool and int)`,
	)

}

// Test Bool * Rune
func TestCheckBinaryExprBoolMulRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true * '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true * '@' (mismatched types bool and rune)`,
	)

}

// Test Bool * Float
func TestCheckBinaryExprBoolMulFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true * 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true * 2 (mismatched types bool and float64)`,
	)

}

// Test Bool * Complex
func TestCheckBinaryExprBoolMulComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true * 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true * 8i (mismatched types bool and complex128)`,
	)

}

// Test Bool * Bool
func TestCheckBinaryExprBoolMulBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true * true`, env,
		`invalid operation: true * true (operator * not defined on bool)`,
	)

}

// Test Bool * String
func TestCheckBinaryExprBoolMulString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true * "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true * "abc" (mismatched types bool and string)`,
	)

}

// Test Bool * Nil
func TestCheckBinaryExprBoolMulNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true * nil`, env,
		`cannot convert nil to type bool`,
		`invalid operation: true * nil (mismatched types bool and <T>)`,
	)

}

// Test Bool / Int
func TestCheckBinaryExprBoolQuoInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true / 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true / 4 (mismatched types bool and int)`,
	)

}

// Test Bool / Rune
func TestCheckBinaryExprBoolQuoRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true / '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true / '@' (mismatched types bool and rune)`,
	)

}

// Test Bool / Float
func TestCheckBinaryExprBoolQuoFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true / 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true / 2 (mismatched types bool and float64)`,
	)

}

// Test Bool / Complex
func TestCheckBinaryExprBoolQuoComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true / 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true / 8i (mismatched types bool and complex128)`,
	)

}

// Test Bool / Bool
func TestCheckBinaryExprBoolQuoBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true / true`, env,
		`invalid operation: true / true (operator / not defined on bool)`,
	)

}

// Test Bool / String
func TestCheckBinaryExprBoolQuoString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true / "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true / "abc" (mismatched types bool and string)`,
	)

}

// Test Bool / Nil
func TestCheckBinaryExprBoolQuoNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true / nil`, env,
		`cannot convert nil to type bool`,
		`invalid operation: true / nil (mismatched types bool and <T>)`,
	)

}

// Test Bool % Int
func TestCheckBinaryExprBoolRemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true % 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true % 4 (mismatched types bool and int)`,
	)

}

// Test Bool % Rune
func TestCheckBinaryExprBoolRemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true % '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true % '@' (mismatched types bool and rune)`,
	)

}

// Test Bool % Float
func TestCheckBinaryExprBoolRemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true % 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true % 2 (mismatched types bool and float64)`,
	)

}

// Test Bool % Complex
func TestCheckBinaryExprBoolRemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true % 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true % 8i (mismatched types bool and complex128)`,
	)

}

// Test Bool % Bool
func TestCheckBinaryExprBoolRemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true % true`, env,
		`invalid operation: true % true (operator % not defined on bool)`,
	)

}

// Test Bool % String
func TestCheckBinaryExprBoolRemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true % "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true % "abc" (mismatched types bool and string)`,
	)

}

// Test Bool % Nil
func TestCheckBinaryExprBoolRemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true % nil`, env,
		`cannot convert nil to type bool`,
		`invalid operation: true % nil (mismatched types bool and <T>)`,
	)

}

// Test Bool & Int
func TestCheckBinaryExprBoolAndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true & 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true & 4 (mismatched types bool and int)`,
	)

}

// Test Bool & Rune
func TestCheckBinaryExprBoolAndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true & '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true & '@' (mismatched types bool and rune)`,
	)

}

// Test Bool & Float
func TestCheckBinaryExprBoolAndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true & 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true & 2 (mismatched types bool and float64)`,
	)

}

// Test Bool & Complex
func TestCheckBinaryExprBoolAndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true & 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true & 8i (mismatched types bool and complex128)`,
	)

}

// Test Bool & Bool
func TestCheckBinaryExprBoolAndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true & true`, env,
		`invalid operation: true & true (operator & not defined on bool)`,
	)

}

// Test Bool & String
func TestCheckBinaryExprBoolAndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true & "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true & "abc" (mismatched types bool and string)`,
	)

}

// Test Bool & Nil
func TestCheckBinaryExprBoolAndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true & nil`, env,
		`cannot convert nil to type bool`,
		`invalid operation: true & nil (mismatched types bool and <T>)`,
	)

}

// Test Bool | Int
func TestCheckBinaryExprBoolOrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true | 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true | 4 (mismatched types bool and int)`,
	)

}

// Test Bool | Rune
func TestCheckBinaryExprBoolOrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true | '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true | '@' (mismatched types bool and rune)`,
	)

}

// Test Bool | Float
func TestCheckBinaryExprBoolOrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true | 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true | 2 (mismatched types bool and float64)`,
	)

}

// Test Bool | Complex
func TestCheckBinaryExprBoolOrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true | 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true | 8i (mismatched types bool and complex128)`,
	)

}

// Test Bool | Bool
func TestCheckBinaryExprBoolOrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true | true`, env,
		`invalid operation: true | true (operator | not defined on bool)`,
	)

}

// Test Bool | String
func TestCheckBinaryExprBoolOrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true | "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true | "abc" (mismatched types bool and string)`,
	)

}

// Test Bool | Nil
func TestCheckBinaryExprBoolOrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true | nil`, env,
		`cannot convert nil to type bool`,
		`invalid operation: true | nil (mismatched types bool and <T>)`,
	)

}

// Test Bool ^ Int
func TestCheckBinaryExprBoolXorInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true ^ 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true ^ 4 (mismatched types bool and int)`,
	)

}

// Test Bool ^ Rune
func TestCheckBinaryExprBoolXorRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true ^ '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true ^ '@' (mismatched types bool and rune)`,
	)

}

// Test Bool ^ Float
func TestCheckBinaryExprBoolXorFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true ^ 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true ^ 2 (mismatched types bool and float64)`,
	)

}

// Test Bool ^ Complex
func TestCheckBinaryExprBoolXorComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true ^ 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true ^ 8i (mismatched types bool and complex128)`,
	)

}

// Test Bool ^ Bool
func TestCheckBinaryExprBoolXorBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true ^ true`, env,
		`invalid operation: true ^ true (operator ^ not defined on bool)`,
	)

}

// Test Bool ^ String
func TestCheckBinaryExprBoolXorString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true ^ "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true ^ "abc" (mismatched types bool and string)`,
	)

}

// Test Bool ^ Nil
func TestCheckBinaryExprBoolXorNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true ^ nil`, env,
		`cannot convert nil to type bool`,
		`invalid operation: true ^ nil (mismatched types bool and <T>)`,
	)

}

// Test Bool &^ Int
func TestCheckBinaryExprBoolAndNotInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true &^ 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true &^ 4 (mismatched types bool and int)`,
	)

}

// Test Bool &^ Rune
func TestCheckBinaryExprBoolAndNotRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true &^ '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true &^ '@' (mismatched types bool and rune)`,
	)

}

// Test Bool &^ Float
func TestCheckBinaryExprBoolAndNotFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true &^ 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true &^ 2 (mismatched types bool and float64)`,
	)

}

// Test Bool &^ Complex
func TestCheckBinaryExprBoolAndNotComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true &^ 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true &^ 8i (mismatched types bool and complex128)`,
	)

}

// Test Bool &^ Bool
func TestCheckBinaryExprBoolAndNotBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true &^ true`, env,
		`invalid operation: true &^ true (operator &^ not defined on bool)`,
	)

}

// Test Bool &^ String
func TestCheckBinaryExprBoolAndNotString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true &^ "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true &^ "abc" (mismatched types bool and string)`,
	)

}

// Test Bool &^ Nil
func TestCheckBinaryExprBoolAndNotNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true &^ nil`, env,
		`cannot convert nil to type bool`,
		`invalid operation: true &^ nil (mismatched types bool and <T>)`,
	)

}

// Test Bool == Int
func TestCheckBinaryExprBoolEqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true == 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true == 4 (mismatched types bool and int)`,
	)

}

// Test Bool == Rune
func TestCheckBinaryExprBoolEqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true == '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true == '@' (mismatched types bool and rune)`,
	)

}

// Test Bool == Float
func TestCheckBinaryExprBoolEqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true == 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true == 2 (mismatched types bool and float64)`,
	)

}

// Test Bool == Complex
func TestCheckBinaryExprBoolEqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true == 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true == 8i (mismatched types bool and complex128)`,
	)

}

// Test Bool == Bool
func TestCheckBinaryExprBoolEqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `true == true`, env, (true == true), ConstBool)
}

// Test Bool == String
func TestCheckBinaryExprBoolEqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true == "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true == "abc" (mismatched types bool and string)`,
	)

}

// Test Bool == Nil
func TestCheckBinaryExprBoolEqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true == nil`, env,
		`cannot convert nil to type bool`,
		`invalid operation: true == nil (mismatched types bool and <T>)`,
	)

}

// Test Bool != Int
func TestCheckBinaryExprBoolNeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true != 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true != 4 (mismatched types bool and int)`,
	)

}

// Test Bool != Rune
func TestCheckBinaryExprBoolNeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true != '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true != '@' (mismatched types bool and rune)`,
	)

}

// Test Bool != Float
func TestCheckBinaryExprBoolNeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true != 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true != 2 (mismatched types bool and float64)`,
	)

}

// Test Bool != Complex
func TestCheckBinaryExprBoolNeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true != 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true != 8i (mismatched types bool and complex128)`,
	)

}

// Test Bool != Bool
func TestCheckBinaryExprBoolNeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `true != true`, env, (true != true), ConstBool)
}

// Test Bool != String
func TestCheckBinaryExprBoolNeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true != "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true != "abc" (mismatched types bool and string)`,
	)

}

// Test Bool != Nil
func TestCheckBinaryExprBoolNeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true != nil`, env,
		`cannot convert nil to type bool`,
		`invalid operation: true != nil (mismatched types bool and <T>)`,
	)

}

// Test Bool <= Int
func TestCheckBinaryExprBoolLeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true <= 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true <= 4 (mismatched types bool and int)`,
	)

}

// Test Bool <= Rune
func TestCheckBinaryExprBoolLeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true <= '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true <= '@' (mismatched types bool and rune)`,
	)

}

// Test Bool <= Float
func TestCheckBinaryExprBoolLeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true <= 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true <= 2 (mismatched types bool and float64)`,
	)

}

// Test Bool <= Complex
func TestCheckBinaryExprBoolLeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true <= 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true <= 8i (mismatched types bool and complex128)`,
	)

}

// Test Bool <= Bool
func TestCheckBinaryExprBoolLeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true <= true`, env,
		`invalid operation: true <= true (operator <= not defined on bool)`,
	)

}

// Test Bool <= String
func TestCheckBinaryExprBoolLeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true <= "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true <= "abc" (mismatched types bool and string)`,
	)

}

// Test Bool <= Nil
func TestCheckBinaryExprBoolLeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true <= nil`, env,
		`cannot convert nil to type bool`,
		`invalid operation: true <= nil (mismatched types bool and <T>)`,
	)

}

// Test Bool >= Int
func TestCheckBinaryExprBoolGeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true >= 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true >= 4 (mismatched types bool and int)`,
	)

}

// Test Bool >= Rune
func TestCheckBinaryExprBoolGeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true >= '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true >= '@' (mismatched types bool and rune)`,
	)

}

// Test Bool >= Float
func TestCheckBinaryExprBoolGeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true >= 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true >= 2 (mismatched types bool and float64)`,
	)

}

// Test Bool >= Complex
func TestCheckBinaryExprBoolGeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true >= 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true >= 8i (mismatched types bool and complex128)`,
	)

}

// Test Bool >= Bool
func TestCheckBinaryExprBoolGeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true >= true`, env,
		`invalid operation: true >= true (operator >= not defined on bool)`,
	)

}

// Test Bool >= String
func TestCheckBinaryExprBoolGeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true >= "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true >= "abc" (mismatched types bool and string)`,
	)

}

// Test Bool >= Nil
func TestCheckBinaryExprBoolGeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true >= nil`, env,
		`cannot convert nil to type bool`,
		`invalid operation: true >= nil (mismatched types bool and <T>)`,
	)

}

// Test Bool < Int
func TestCheckBinaryExprBoolLssInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true < 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true < 4 (mismatched types bool and int)`,
	)

}

// Test Bool < Rune
func TestCheckBinaryExprBoolLssRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true < '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true < '@' (mismatched types bool and rune)`,
	)

}

// Test Bool < Float
func TestCheckBinaryExprBoolLssFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true < 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true < 2 (mismatched types bool and float64)`,
	)

}

// Test Bool < Complex
func TestCheckBinaryExprBoolLssComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true < 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true < 8i (mismatched types bool and complex128)`,
	)

}

// Test Bool < Bool
func TestCheckBinaryExprBoolLssBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true < true`, env,
		`invalid operation: true < true (operator < not defined on bool)`,
	)

}

// Test Bool < String
func TestCheckBinaryExprBoolLssString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true < "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true < "abc" (mismatched types bool and string)`,
	)

}

// Test Bool < Nil
func TestCheckBinaryExprBoolLssNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true < nil`, env,
		`cannot convert nil to type bool`,
		`invalid operation: true < nil (mismatched types bool and <T>)`,
	)

}

// Test Bool > Int
func TestCheckBinaryExprBoolGtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true > 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true > 4 (mismatched types bool and int)`,
	)

}

// Test Bool > Rune
func TestCheckBinaryExprBoolGtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true > '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true > '@' (mismatched types bool and rune)`,
	)

}

// Test Bool > Float
func TestCheckBinaryExprBoolGtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true > 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true > 2 (mismatched types bool and float64)`,
	)

}

// Test Bool > Complex
func TestCheckBinaryExprBoolGtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true > 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true > 8i (mismatched types bool and complex128)`,
	)

}

// Test Bool > Bool
func TestCheckBinaryExprBoolGtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true > true`, env,
		`invalid operation: true > true (operator > not defined on bool)`,
	)

}

// Test Bool > String
func TestCheckBinaryExprBoolGtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true > "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true > "abc" (mismatched types bool and string)`,
	)

}

// Test Bool > Nil
func TestCheckBinaryExprBoolGtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true > nil`, env,
		`cannot convert nil to type bool`,
		`invalid operation: true > nil (mismatched types bool and <T>)`,
	)

}

// Test Bool >> Int
func TestCheckBinaryExprBoolRhlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true >> 4`, env,
		`invalid operation: true >> 4 (shift of type untyped bool)`,
	)

}

// Test Bool >> Rune
func TestCheckBinaryExprBoolRhlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true >> '@'`, env,
		`invalid operation: true >> 64 (shift of type untyped bool)`,
	)

}

// Test Bool >> Float
func TestCheckBinaryExprBoolRhlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true >> 2.0`, env,
		`invalid operation: true >> 2 (shift of type untyped bool)`,
	)

}

// Test Bool >> Complex
func TestCheckBinaryExprBoolRhlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true >> 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: true >> 0 (shift of type untyped bool)`,
	)

}

// Test Bool >> Bool
func TestCheckBinaryExprBoolRhlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true >> true`, env,
		`invalid operation: true >> true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Bool >> String
func TestCheckBinaryExprBoolRhlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true >> "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: true >> "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Bool >> Nil
func TestCheckBinaryExprBoolRhlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `true >> nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test String + Int
func TestCheckBinaryExprStringAddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" + 4`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: "abc" + 4 (mismatched types string and int)`,
	)

}

// Test String + Rune
func TestCheckBinaryExprStringAddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" + '@'`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: "abc" + '@' (mismatched types string and rune)`,
	)

}

// Test String + Float
func TestCheckBinaryExprStringAddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" + 2.0`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: "abc" + 2 (mismatched types string and float64)`,
	)

}

// Test String + Complex
func TestCheckBinaryExprStringAddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" + 8.0i`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: "abc" + 8i (mismatched types string and complex128)`,
	)

}

// Test String + Bool
func TestCheckBinaryExprStringAddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" + true`, env,
		`cannot convert "abc" to type int`,
		`cannot convert true to type int`,
		`invalid operation: "abc" + true (mismatched types string and bool)`,
	)

}

// Test String + String
func TestCheckBinaryExprStringAddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `"abc" + "abc"`, env, ("abc" + "abc"), ConstString)
}

// Test String + Nil
func TestCheckBinaryExprStringAddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" + nil`, env,
		`cannot convert "abc" to type int`,
		`cannot convert nil to type int`,
		`invalid operation: "abc" + nil (mismatched types string and <T>)`,
	)

}

// Test String - Int
func TestCheckBinaryExprStringSubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" - 4`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: "abc" - 4 (mismatched types string and int)`,
	)

}

// Test String - Rune
func TestCheckBinaryExprStringSubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" - '@'`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: "abc" - '@' (mismatched types string and rune)`,
	)

}

// Test String - Float
func TestCheckBinaryExprStringSubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" - 2.0`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: "abc" - 2 (mismatched types string and float64)`,
	)

}

// Test String - Complex
func TestCheckBinaryExprStringSubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" - 8.0i`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: "abc" - 8i (mismatched types string and complex128)`,
	)

}

// Test String - Bool
func TestCheckBinaryExprStringSubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" - true`, env,
		`cannot convert "abc" to type int`,
		`cannot convert true to type int`,
		`invalid operation: "abc" - true (mismatched types string and bool)`,
	)

}

// Test String - String
func TestCheckBinaryExprStringSubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" - "abc"`, env,
		`invalid operation: "abc" - "abc" (operator - not defined on string)`,
	)

}

// Test String - Nil
func TestCheckBinaryExprStringSubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" - nil`, env,
		`cannot convert "abc" to type int`,
		`cannot convert nil to type int`,
		`invalid operation: "abc" - nil (mismatched types string and <T>)`,
	)

}

// Test String * Int
func TestCheckBinaryExprStringMulInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" * 4`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: "abc" * 4 (mismatched types string and int)`,
	)

}

// Test String * Rune
func TestCheckBinaryExprStringMulRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" * '@'`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: "abc" * '@' (mismatched types string and rune)`,
	)

}

// Test String * Float
func TestCheckBinaryExprStringMulFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" * 2.0`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: "abc" * 2 (mismatched types string and float64)`,
	)

}

// Test String * Complex
func TestCheckBinaryExprStringMulComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" * 8.0i`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: "abc" * 8i (mismatched types string and complex128)`,
	)

}

// Test String * Bool
func TestCheckBinaryExprStringMulBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" * true`, env,
		`cannot convert "abc" to type int`,
		`cannot convert true to type int`,
		`invalid operation: "abc" * true (mismatched types string and bool)`,
	)

}

// Test String * String
func TestCheckBinaryExprStringMulString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" * "abc"`, env,
		`invalid operation: "abc" * "abc" (operator * not defined on string)`,
	)

}

// Test String * Nil
func TestCheckBinaryExprStringMulNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" * nil`, env,
		`cannot convert "abc" to type int`,
		`cannot convert nil to type int`,
		`invalid operation: "abc" * nil (mismatched types string and <T>)`,
	)

}

// Test String / Int
func TestCheckBinaryExprStringQuoInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" / 4`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: "abc" / 4 (mismatched types string and int)`,
	)

}

// Test String / Rune
func TestCheckBinaryExprStringQuoRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" / '@'`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: "abc" / '@' (mismatched types string and rune)`,
	)

}

// Test String / Float
func TestCheckBinaryExprStringQuoFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" / 2.0`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: "abc" / 2 (mismatched types string and float64)`,
	)

}

// Test String / Complex
func TestCheckBinaryExprStringQuoComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" / 8.0i`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: "abc" / 8i (mismatched types string and complex128)`,
	)

}

// Test String / Bool
func TestCheckBinaryExprStringQuoBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" / true`, env,
		`cannot convert "abc" to type int`,
		`cannot convert true to type int`,
		`invalid operation: "abc" / true (mismatched types string and bool)`,
	)

}

// Test String / String
func TestCheckBinaryExprStringQuoString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" / "abc"`, env,
		`invalid operation: "abc" / "abc" (operator / not defined on string)`,
	)

}

// Test String / Nil
func TestCheckBinaryExprStringQuoNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" / nil`, env,
		`cannot convert "abc" to type int`,
		`cannot convert nil to type int`,
		`invalid operation: "abc" / nil (mismatched types string and <T>)`,
	)

}

// Test String % Int
func TestCheckBinaryExprStringRemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" % 4`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: "abc" % 4 (mismatched types string and int)`,
	)

}

// Test String % Rune
func TestCheckBinaryExprStringRemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" % '@'`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: "abc" % '@' (mismatched types string and rune)`,
	)

}

// Test String % Float
func TestCheckBinaryExprStringRemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" % 2.0`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: "abc" % 2 (mismatched types string and float64)`,
	)

}

// Test String % Complex
func TestCheckBinaryExprStringRemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" % 8.0i`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: "abc" % 8i (mismatched types string and complex128)`,
	)

}

// Test String % Bool
func TestCheckBinaryExprStringRemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" % true`, env,
		`cannot convert "abc" to type int`,
		`cannot convert true to type int`,
		`invalid operation: "abc" % true (mismatched types string and bool)`,
	)

}

// Test String % String
func TestCheckBinaryExprStringRemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" % "abc"`, env,
		`invalid operation: "abc" % "abc" (operator % not defined on string)`,
	)

}

// Test String % Nil
func TestCheckBinaryExprStringRemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" % nil`, env,
		`cannot convert "abc" to type int`,
		`cannot convert nil to type int`,
		`invalid operation: "abc" % nil (mismatched types string and <T>)`,
	)

}

// Test String & Int
func TestCheckBinaryExprStringAndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" & 4`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: "abc" & 4 (mismatched types string and int)`,
	)

}

// Test String & Rune
func TestCheckBinaryExprStringAndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" & '@'`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: "abc" & '@' (mismatched types string and rune)`,
	)

}

// Test String & Float
func TestCheckBinaryExprStringAndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" & 2.0`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: "abc" & 2 (mismatched types string and float64)`,
	)

}

// Test String & Complex
func TestCheckBinaryExprStringAndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" & 8.0i`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: "abc" & 8i (mismatched types string and complex128)`,
	)

}

// Test String & Bool
func TestCheckBinaryExprStringAndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" & true`, env,
		`cannot convert "abc" to type int`,
		`cannot convert true to type int`,
		`invalid operation: "abc" & true (mismatched types string and bool)`,
	)

}

// Test String & String
func TestCheckBinaryExprStringAndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" & "abc"`, env,
		`invalid operation: "abc" & "abc" (operator & not defined on string)`,
	)

}

// Test String & Nil
func TestCheckBinaryExprStringAndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" & nil`, env,
		`cannot convert "abc" to type int`,
		`cannot convert nil to type int`,
		`invalid operation: "abc" & nil (mismatched types string and <T>)`,
	)

}

// Test String | Int
func TestCheckBinaryExprStringOrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" | 4`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: "abc" | 4 (mismatched types string and int)`,
	)

}

// Test String | Rune
func TestCheckBinaryExprStringOrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" | '@'`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: "abc" | '@' (mismatched types string and rune)`,
	)

}

// Test String | Float
func TestCheckBinaryExprStringOrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" | 2.0`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: "abc" | 2 (mismatched types string and float64)`,
	)

}

// Test String | Complex
func TestCheckBinaryExprStringOrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" | 8.0i`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: "abc" | 8i (mismatched types string and complex128)`,
	)

}

// Test String | Bool
func TestCheckBinaryExprStringOrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" | true`, env,
		`cannot convert "abc" to type int`,
		`cannot convert true to type int`,
		`invalid operation: "abc" | true (mismatched types string and bool)`,
	)

}

// Test String | String
func TestCheckBinaryExprStringOrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" | "abc"`, env,
		`invalid operation: "abc" | "abc" (operator | not defined on string)`,
	)

}

// Test String | Nil
func TestCheckBinaryExprStringOrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" | nil`, env,
		`cannot convert "abc" to type int`,
		`cannot convert nil to type int`,
		`invalid operation: "abc" | nil (mismatched types string and <T>)`,
	)

}

// Test String ^ Int
func TestCheckBinaryExprStringXorInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" ^ 4`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: "abc" ^ 4 (mismatched types string and int)`,
	)

}

// Test String ^ Rune
func TestCheckBinaryExprStringXorRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" ^ '@'`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: "abc" ^ '@' (mismatched types string and rune)`,
	)

}

// Test String ^ Float
func TestCheckBinaryExprStringXorFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" ^ 2.0`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: "abc" ^ 2 (mismatched types string and float64)`,
	)

}

// Test String ^ Complex
func TestCheckBinaryExprStringXorComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" ^ 8.0i`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: "abc" ^ 8i (mismatched types string and complex128)`,
	)

}

// Test String ^ Bool
func TestCheckBinaryExprStringXorBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" ^ true`, env,
		`cannot convert "abc" to type int`,
		`cannot convert true to type int`,
		`invalid operation: "abc" ^ true (mismatched types string and bool)`,
	)

}

// Test String ^ String
func TestCheckBinaryExprStringXorString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" ^ "abc"`, env,
		`invalid operation: "abc" ^ "abc" (operator ^ not defined on string)`,
	)

}

// Test String ^ Nil
func TestCheckBinaryExprStringXorNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" ^ nil`, env,
		`cannot convert "abc" to type int`,
		`cannot convert nil to type int`,
		`invalid operation: "abc" ^ nil (mismatched types string and <T>)`,
	)

}

// Test String &^ Int
func TestCheckBinaryExprStringAndNotInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" &^ 4`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: "abc" &^ 4 (mismatched types string and int)`,
	)

}

// Test String &^ Rune
func TestCheckBinaryExprStringAndNotRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" &^ '@'`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: "abc" &^ '@' (mismatched types string and rune)`,
	)

}

// Test String &^ Float
func TestCheckBinaryExprStringAndNotFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" &^ 2.0`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: "abc" &^ 2 (mismatched types string and float64)`,
	)

}

// Test String &^ Complex
func TestCheckBinaryExprStringAndNotComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" &^ 8.0i`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: "abc" &^ 8i (mismatched types string and complex128)`,
	)

}

// Test String &^ Bool
func TestCheckBinaryExprStringAndNotBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" &^ true`, env,
		`cannot convert "abc" to type int`,
		`cannot convert true to type int`,
		`invalid operation: "abc" &^ true (mismatched types string and bool)`,
	)

}

// Test String &^ String
func TestCheckBinaryExprStringAndNotString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" &^ "abc"`, env,
		`invalid operation: "abc" &^ "abc" (operator &^ not defined on string)`,
	)

}

// Test String &^ Nil
func TestCheckBinaryExprStringAndNotNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" &^ nil`, env,
		`cannot convert "abc" to type int`,
		`cannot convert nil to type int`,
		`invalid operation: "abc" &^ nil (mismatched types string and <T>)`,
	)

}

// Test String == Int
func TestCheckBinaryExprStringEqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" == 4`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: "abc" == 4 (mismatched types string and int)`,
	)

}

// Test String == Rune
func TestCheckBinaryExprStringEqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" == '@'`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: "abc" == '@' (mismatched types string and rune)`,
	)

}

// Test String == Float
func TestCheckBinaryExprStringEqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" == 2.0`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: "abc" == 2 (mismatched types string and float64)`,
	)

}

// Test String == Complex
func TestCheckBinaryExprStringEqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" == 8.0i`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: "abc" == 8i (mismatched types string and complex128)`,
	)

}

// Test String == Bool
func TestCheckBinaryExprStringEqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" == true`, env,
		`cannot convert "abc" to type int`,
		`cannot convert true to type int`,
		`invalid operation: "abc" == true (mismatched types string and bool)`,
	)

}

// Test String == String
func TestCheckBinaryExprStringEqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `"abc" == "abc"`, env, ("abc" == "abc"), ConstBool)
}

// Test String == Nil
func TestCheckBinaryExprStringEqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" == nil`, env,
		`cannot convert "abc" to type int`,
		`cannot convert nil to type int`,
		`invalid operation: "abc" == nil (mismatched types string and <T>)`,
	)

}

// Test String != Int
func TestCheckBinaryExprStringNeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" != 4`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: "abc" != 4 (mismatched types string and int)`,
	)

}

// Test String != Rune
func TestCheckBinaryExprStringNeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" != '@'`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: "abc" != '@' (mismatched types string and rune)`,
	)

}

// Test String != Float
func TestCheckBinaryExprStringNeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" != 2.0`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: "abc" != 2 (mismatched types string and float64)`,
	)

}

// Test String != Complex
func TestCheckBinaryExprStringNeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" != 8.0i`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: "abc" != 8i (mismatched types string and complex128)`,
	)

}

// Test String != Bool
func TestCheckBinaryExprStringNeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" != true`, env,
		`cannot convert "abc" to type int`,
		`cannot convert true to type int`,
		`invalid operation: "abc" != true (mismatched types string and bool)`,
	)

}

// Test String != String
func TestCheckBinaryExprStringNeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `"abc" != "abc"`, env, ("abc" != "abc"), ConstBool)
}

// Test String != Nil
func TestCheckBinaryExprStringNeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" != nil`, env,
		`cannot convert "abc" to type int`,
		`cannot convert nil to type int`,
		`invalid operation: "abc" != nil (mismatched types string and <T>)`,
	)

}

// Test String <= Int
func TestCheckBinaryExprStringLeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" <= 4`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: "abc" <= 4 (mismatched types string and int)`,
	)

}

// Test String <= Rune
func TestCheckBinaryExprStringLeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" <= '@'`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: "abc" <= '@' (mismatched types string and rune)`,
	)

}

// Test String <= Float
func TestCheckBinaryExprStringLeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" <= 2.0`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: "abc" <= 2 (mismatched types string and float64)`,
	)

}

// Test String <= Complex
func TestCheckBinaryExprStringLeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" <= 8.0i`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: "abc" <= 8i (mismatched types string and complex128)`,
	)

}

// Test String <= Bool
func TestCheckBinaryExprStringLeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" <= true`, env,
		`cannot convert "abc" to type int`,
		`cannot convert true to type int`,
		`invalid operation: "abc" <= true (mismatched types string and bool)`,
	)

}

// Test String <= String
func TestCheckBinaryExprStringLeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `"abc" <= "abc"`, env, ("abc" <= "abc"), ConstBool)
}

// Test String <= Nil
func TestCheckBinaryExprStringLeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" <= nil`, env,
		`cannot convert "abc" to type int`,
		`cannot convert nil to type int`,
		`invalid operation: "abc" <= nil (mismatched types string and <T>)`,
	)

}

// Test String >= Int
func TestCheckBinaryExprStringGeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" >= 4`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: "abc" >= 4 (mismatched types string and int)`,
	)

}

// Test String >= Rune
func TestCheckBinaryExprStringGeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" >= '@'`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: "abc" >= '@' (mismatched types string and rune)`,
	)

}

// Test String >= Float
func TestCheckBinaryExprStringGeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" >= 2.0`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: "abc" >= 2 (mismatched types string and float64)`,
	)

}

// Test String >= Complex
func TestCheckBinaryExprStringGeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" >= 8.0i`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: "abc" >= 8i (mismatched types string and complex128)`,
	)

}

// Test String >= Bool
func TestCheckBinaryExprStringGeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" >= true`, env,
		`cannot convert "abc" to type int`,
		`cannot convert true to type int`,
		`invalid operation: "abc" >= true (mismatched types string and bool)`,
	)

}

// Test String >= String
func TestCheckBinaryExprStringGeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `"abc" >= "abc"`, env, ("abc" >= "abc"), ConstBool)
}

// Test String >= Nil
func TestCheckBinaryExprStringGeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" >= nil`, env,
		`cannot convert "abc" to type int`,
		`cannot convert nil to type int`,
		`invalid operation: "abc" >= nil (mismatched types string and <T>)`,
	)

}

// Test String < Int
func TestCheckBinaryExprStringLssInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" < 4`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: "abc" < 4 (mismatched types string and int)`,
	)

}

// Test String < Rune
func TestCheckBinaryExprStringLssRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" < '@'`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: "abc" < '@' (mismatched types string and rune)`,
	)

}

// Test String < Float
func TestCheckBinaryExprStringLssFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" < 2.0`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: "abc" < 2 (mismatched types string and float64)`,
	)

}

// Test String < Complex
func TestCheckBinaryExprStringLssComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" < 8.0i`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: "abc" < 8i (mismatched types string and complex128)`,
	)

}

// Test String < Bool
func TestCheckBinaryExprStringLssBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" < true`, env,
		`cannot convert "abc" to type int`,
		`cannot convert true to type int`,
		`invalid operation: "abc" < true (mismatched types string and bool)`,
	)

}

// Test String < String
func TestCheckBinaryExprStringLssString(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `"abc" < "abc"`, env, ("abc" < "abc"), ConstBool)
}

// Test String < Nil
func TestCheckBinaryExprStringLssNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" < nil`, env,
		`cannot convert "abc" to type int`,
		`cannot convert nil to type int`,
		`invalid operation: "abc" < nil (mismatched types string and <T>)`,
	)

}

// Test String > Int
func TestCheckBinaryExprStringGtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" > 4`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: "abc" > 4 (mismatched types string and int)`,
	)

}

// Test String > Rune
func TestCheckBinaryExprStringGtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" > '@'`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: "abc" > '@' (mismatched types string and rune)`,
	)

}

// Test String > Float
func TestCheckBinaryExprStringGtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" > 2.0`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: "abc" > 2 (mismatched types string and float64)`,
	)

}

// Test String > Complex
func TestCheckBinaryExprStringGtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" > 8.0i`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: "abc" > 8i (mismatched types string and complex128)`,
	)

}

// Test String > Bool
func TestCheckBinaryExprStringGtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" > true`, env,
		`cannot convert "abc" to type int`,
		`cannot convert true to type int`,
		`invalid operation: "abc" > true (mismatched types string and bool)`,
	)

}

// Test String > String
func TestCheckBinaryExprStringGtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `"abc" > "abc"`, env, ("abc" > "abc"), ConstBool)
}

// Test String > Nil
func TestCheckBinaryExprStringGtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" > nil`, env,
		`cannot convert "abc" to type int`,
		`cannot convert nil to type int`,
		`invalid operation: "abc" > nil (mismatched types string and <T>)`,
	)

}

// Test String >> Int
func TestCheckBinaryExprStringRhlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" >> 4`, env,
		`invalid operation: "abc" >> 4 (shift of type untyped string)`,
	)

}

// Test String >> Rune
func TestCheckBinaryExprStringRhlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" >> '@'`, env,
		`invalid operation: "abc" >> 64 (shift of type untyped string)`,
	)

}

// Test String >> Float
func TestCheckBinaryExprStringRhlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" >> 2.0`, env,
		`invalid operation: "abc" >> 2 (shift of type untyped string)`,
	)

}

// Test String >> Complex
func TestCheckBinaryExprStringRhlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" >> 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: "abc" >> 0 (shift of type untyped string)`,
	)

}

// Test String >> Bool
func TestCheckBinaryExprStringRhlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" >> true`, env,
		`invalid operation: "abc" >> true (shift count type bool, must be unsigned integer)`,
	)

}

// Test String >> String
func TestCheckBinaryExprStringRhlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" >> "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: "abc" >> "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test String >> Nil
func TestCheckBinaryExprStringRhlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `"abc" >> nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Nil + Int
func TestCheckBinaryExprNilAddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil + 4`, env,
		`cannot convert nil to type int`,
		`invalid operation: nil + 4 (mismatched types <T> and int)`,
	)

}

// Test Nil + Rune
func TestCheckBinaryExprNilAddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil + '@'`, env,
		`cannot convert nil to type rune`,
		`invalid operation: nil + '@' (mismatched types <T> and rune)`,
	)

}

// Test Nil + Float
func TestCheckBinaryExprNilAddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil + 2.0`, env,
		`cannot convert nil to type float64`,
		`invalid operation: nil + 2 (mismatched types <T> and float64)`,
	)

}

// Test Nil + Complex
func TestCheckBinaryExprNilAddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil + 8.0i`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: nil + 8i (mismatched types <T> and complex128)`,
	)

}

// Test Nil + Bool
func TestCheckBinaryExprNilAddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil + true`, env,
		`cannot convert nil to type int`,
		`cannot convert true to type int`,
		`invalid operation: nil + true (mismatched types <T> and bool)`,
	)

}

// Test Nil + String
func TestCheckBinaryExprNilAddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil + "abc"`, env,
		`cannot convert nil to type int`,
		`cannot convert "abc" to type int`,
		`invalid operation: nil + "abc" (mismatched types <T> and string)`,
	)

}

// Test Nil + Nil
func TestCheckBinaryExprNilAddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil + nil`, env,
		`invalid operation: nil + nil (operator + not defined on nil)`,
	)

}

// Test Nil - Int
func TestCheckBinaryExprNilSubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil - 4`, env,
		`cannot convert nil to type int`,
		`invalid operation: nil - 4 (mismatched types <T> and int)`,
	)

}

// Test Nil - Rune
func TestCheckBinaryExprNilSubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil - '@'`, env,
		`cannot convert nil to type rune`,
		`invalid operation: nil - '@' (mismatched types <T> and rune)`,
	)

}

// Test Nil - Float
func TestCheckBinaryExprNilSubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil - 2.0`, env,
		`cannot convert nil to type float64`,
		`invalid operation: nil - 2 (mismatched types <T> and float64)`,
	)

}

// Test Nil - Complex
func TestCheckBinaryExprNilSubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil - 8.0i`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: nil - 8i (mismatched types <T> and complex128)`,
	)

}

// Test Nil - Bool
func TestCheckBinaryExprNilSubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil - true`, env,
		`cannot convert nil to type int`,
		`cannot convert true to type int`,
		`invalid operation: nil - true (mismatched types <T> and bool)`,
	)

}

// Test Nil - String
func TestCheckBinaryExprNilSubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil - "abc"`, env,
		`cannot convert nil to type int`,
		`cannot convert "abc" to type int`,
		`invalid operation: nil - "abc" (mismatched types <T> and string)`,
	)

}

// Test Nil - Nil
func TestCheckBinaryExprNilSubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil - nil`, env,
		`invalid operation: nil - nil (operator - not defined on nil)`,
	)

}

// Test Nil * Int
func TestCheckBinaryExprNilMulInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil * 4`, env,
		`cannot convert nil to type int`,
		`invalid operation: nil * 4 (mismatched types <T> and int)`,
	)

}

// Test Nil * Rune
func TestCheckBinaryExprNilMulRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil * '@'`, env,
		`cannot convert nil to type rune`,
		`invalid operation: nil * '@' (mismatched types <T> and rune)`,
	)

}

// Test Nil * Float
func TestCheckBinaryExprNilMulFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil * 2.0`, env,
		`cannot convert nil to type float64`,
		`invalid operation: nil * 2 (mismatched types <T> and float64)`,
	)

}

// Test Nil * Complex
func TestCheckBinaryExprNilMulComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil * 8.0i`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: nil * 8i (mismatched types <T> and complex128)`,
	)

}

// Test Nil * Bool
func TestCheckBinaryExprNilMulBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil * true`, env,
		`cannot convert nil to type int`,
		`cannot convert true to type int`,
		`invalid operation: nil * true (mismatched types <T> and bool)`,
	)

}

// Test Nil * String
func TestCheckBinaryExprNilMulString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil * "abc"`, env,
		`cannot convert nil to type int`,
		`cannot convert "abc" to type int`,
		`invalid operation: nil * "abc" (mismatched types <T> and string)`,
	)

}

// Test Nil * Nil
func TestCheckBinaryExprNilMulNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil * nil`, env,
		`invalid operation: nil * nil (operator * not defined on nil)`,
	)

}

// Test Nil / Int
func TestCheckBinaryExprNilQuoInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil / 4`, env,
		`cannot convert nil to type int`,
		`invalid operation: nil / 4 (mismatched types <T> and int)`,
	)

}

// Test Nil / Rune
func TestCheckBinaryExprNilQuoRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil / '@'`, env,
		`cannot convert nil to type rune`,
		`invalid operation: nil / '@' (mismatched types <T> and rune)`,
	)

}

// Test Nil / Float
func TestCheckBinaryExprNilQuoFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil / 2.0`, env,
		`cannot convert nil to type float64`,
		`invalid operation: nil / 2 (mismatched types <T> and float64)`,
	)

}

// Test Nil / Complex
func TestCheckBinaryExprNilQuoComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil / 8.0i`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: nil / 8i (mismatched types <T> and complex128)`,
	)

}

// Test Nil / Bool
func TestCheckBinaryExprNilQuoBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil / true`, env,
		`cannot convert nil to type int`,
		`cannot convert true to type int`,
		`invalid operation: nil / true (mismatched types <T> and bool)`,
	)

}

// Test Nil / String
func TestCheckBinaryExprNilQuoString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil / "abc"`, env,
		`cannot convert nil to type int`,
		`cannot convert "abc" to type int`,
		`invalid operation: nil / "abc" (mismatched types <T> and string)`,
	)

}

// Test Nil / Nil
func TestCheckBinaryExprNilQuoNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil / nil`, env,
		`invalid operation: nil / nil (operator / not defined on nil)`,
	)

}

// Test Nil % Int
func TestCheckBinaryExprNilRemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil % 4`, env,
		`cannot convert nil to type int`,
		`invalid operation: nil % 4 (mismatched types <T> and int)`,
	)

}

// Test Nil % Rune
func TestCheckBinaryExprNilRemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil % '@'`, env,
		`cannot convert nil to type rune`,
		`invalid operation: nil % '@' (mismatched types <T> and rune)`,
	)

}

// Test Nil % Float
func TestCheckBinaryExprNilRemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil % 2.0`, env,
		`cannot convert nil to type float64`,
		`invalid operation: nil % 2 (mismatched types <T> and float64)`,
	)

}

// Test Nil % Complex
func TestCheckBinaryExprNilRemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil % 8.0i`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: nil % 8i (mismatched types <T> and complex128)`,
	)

}

// Test Nil % Bool
func TestCheckBinaryExprNilRemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil % true`, env,
		`cannot convert nil to type int`,
		`cannot convert true to type int`,
		`invalid operation: nil % true (mismatched types <T> and bool)`,
	)

}

// Test Nil % String
func TestCheckBinaryExprNilRemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil % "abc"`, env,
		`cannot convert nil to type int`,
		`cannot convert "abc" to type int`,
		`invalid operation: nil % "abc" (mismatched types <T> and string)`,
	)

}

// Test Nil % Nil
func TestCheckBinaryExprNilRemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil % nil`, env,
		`invalid operation: nil % nil (operator % not defined on nil)`,
	)

}

// Test Nil & Int
func TestCheckBinaryExprNilAndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil & 4`, env,
		`cannot convert nil to type int`,
		`invalid operation: nil & 4 (mismatched types <T> and int)`,
	)

}

// Test Nil & Rune
func TestCheckBinaryExprNilAndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil & '@'`, env,
		`cannot convert nil to type rune`,
		`invalid operation: nil & '@' (mismatched types <T> and rune)`,
	)

}

// Test Nil & Float
func TestCheckBinaryExprNilAndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil & 2.0`, env,
		`cannot convert nil to type float64`,
		`invalid operation: nil & 2 (mismatched types <T> and float64)`,
	)

}

// Test Nil & Complex
func TestCheckBinaryExprNilAndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil & 8.0i`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: nil & 8i (mismatched types <T> and complex128)`,
	)

}

// Test Nil & Bool
func TestCheckBinaryExprNilAndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil & true`, env,
		`cannot convert nil to type int`,
		`cannot convert true to type int`,
		`invalid operation: nil & true (mismatched types <T> and bool)`,
	)

}

// Test Nil & String
func TestCheckBinaryExprNilAndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil & "abc"`, env,
		`cannot convert nil to type int`,
		`cannot convert "abc" to type int`,
		`invalid operation: nil & "abc" (mismatched types <T> and string)`,
	)

}

// Test Nil & Nil
func TestCheckBinaryExprNilAndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil & nil`, env,
		`invalid operation: nil & nil (operator & not defined on nil)`,
	)

}

// Test Nil | Int
func TestCheckBinaryExprNilOrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil | 4`, env,
		`cannot convert nil to type int`,
		`invalid operation: nil | 4 (mismatched types <T> and int)`,
	)

}

// Test Nil | Rune
func TestCheckBinaryExprNilOrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil | '@'`, env,
		`cannot convert nil to type rune`,
		`invalid operation: nil | '@' (mismatched types <T> and rune)`,
	)

}

// Test Nil | Float
func TestCheckBinaryExprNilOrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil | 2.0`, env,
		`cannot convert nil to type float64`,
		`invalid operation: nil | 2 (mismatched types <T> and float64)`,
	)

}

// Test Nil | Complex
func TestCheckBinaryExprNilOrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil | 8.0i`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: nil | 8i (mismatched types <T> and complex128)`,
	)

}

// Test Nil | Bool
func TestCheckBinaryExprNilOrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil | true`, env,
		`cannot convert nil to type int`,
		`cannot convert true to type int`,
		`invalid operation: nil | true (mismatched types <T> and bool)`,
	)

}

// Test Nil | String
func TestCheckBinaryExprNilOrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil | "abc"`, env,
		`cannot convert nil to type int`,
		`cannot convert "abc" to type int`,
		`invalid operation: nil | "abc" (mismatched types <T> and string)`,
	)

}

// Test Nil | Nil
func TestCheckBinaryExprNilOrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil | nil`, env,
		`invalid operation: nil | nil (operator | not defined on nil)`,
	)

}

// Test Nil ^ Int
func TestCheckBinaryExprNilXorInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil ^ 4`, env,
		`cannot convert nil to type int`,
		`invalid operation: nil ^ 4 (mismatched types <T> and int)`,
	)

}

// Test Nil ^ Rune
func TestCheckBinaryExprNilXorRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil ^ '@'`, env,
		`cannot convert nil to type rune`,
		`invalid operation: nil ^ '@' (mismatched types <T> and rune)`,
	)

}

// Test Nil ^ Float
func TestCheckBinaryExprNilXorFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil ^ 2.0`, env,
		`cannot convert nil to type float64`,
		`invalid operation: nil ^ 2 (mismatched types <T> and float64)`,
	)

}

// Test Nil ^ Complex
func TestCheckBinaryExprNilXorComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil ^ 8.0i`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: nil ^ 8i (mismatched types <T> and complex128)`,
	)

}

// Test Nil ^ Bool
func TestCheckBinaryExprNilXorBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil ^ true`, env,
		`cannot convert nil to type int`,
		`cannot convert true to type int`,
		`invalid operation: nil ^ true (mismatched types <T> and bool)`,
	)

}

// Test Nil ^ String
func TestCheckBinaryExprNilXorString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil ^ "abc"`, env,
		`cannot convert nil to type int`,
		`cannot convert "abc" to type int`,
		`invalid operation: nil ^ "abc" (mismatched types <T> and string)`,
	)

}

// Test Nil ^ Nil
func TestCheckBinaryExprNilXorNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil ^ nil`, env,
		`invalid operation: nil ^ nil (operator ^ not defined on nil)`,
	)

}

// Test Nil &^ Int
func TestCheckBinaryExprNilAndNotInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil &^ 4`, env,
		`cannot convert nil to type int`,
		`invalid operation: nil &^ 4 (mismatched types <T> and int)`,
	)

}

// Test Nil &^ Rune
func TestCheckBinaryExprNilAndNotRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil &^ '@'`, env,
		`cannot convert nil to type rune`,
		`invalid operation: nil &^ '@' (mismatched types <T> and rune)`,
	)

}

// Test Nil &^ Float
func TestCheckBinaryExprNilAndNotFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil &^ 2.0`, env,
		`cannot convert nil to type float64`,
		`invalid operation: nil &^ 2 (mismatched types <T> and float64)`,
	)

}

// Test Nil &^ Complex
func TestCheckBinaryExprNilAndNotComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil &^ 8.0i`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: nil &^ 8i (mismatched types <T> and complex128)`,
	)

}

// Test Nil &^ Bool
func TestCheckBinaryExprNilAndNotBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil &^ true`, env,
		`cannot convert nil to type int`,
		`cannot convert true to type int`,
		`invalid operation: nil &^ true (mismatched types <T> and bool)`,
	)

}

// Test Nil &^ String
func TestCheckBinaryExprNilAndNotString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil &^ "abc"`, env,
		`cannot convert nil to type int`,
		`cannot convert "abc" to type int`,
		`invalid operation: nil &^ "abc" (mismatched types <T> and string)`,
	)

}

// Test Nil &^ Nil
func TestCheckBinaryExprNilAndNotNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil &^ nil`, env,
		`invalid operation: nil &^ nil (operator &^ not defined on nil)`,
	)

}

// Test Nil == Int
func TestCheckBinaryExprNilEqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil == 4`, env,
		`cannot convert nil to type int`,
		`invalid operation: nil == 4 (mismatched types <T> and int)`,
	)

}

// Test Nil == Rune
func TestCheckBinaryExprNilEqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil == '@'`, env,
		`cannot convert nil to type rune`,
		`invalid operation: nil == '@' (mismatched types <T> and rune)`,
	)

}

// Test Nil == Float
func TestCheckBinaryExprNilEqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil == 2.0`, env,
		`cannot convert nil to type float64`,
		`invalid operation: nil == 2 (mismatched types <T> and float64)`,
	)

}

// Test Nil == Complex
func TestCheckBinaryExprNilEqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil == 8.0i`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: nil == 8i (mismatched types <T> and complex128)`,
	)

}

// Test Nil == Bool
func TestCheckBinaryExprNilEqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil == true`, env,
		`cannot convert nil to type int`,
		`cannot convert true to type int`,
		`invalid operation: nil == true (mismatched types <T> and bool)`,
	)

}

// Test Nil == String
func TestCheckBinaryExprNilEqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil == "abc"`, env,
		`cannot convert nil to type int`,
		`cannot convert "abc" to type int`,
		`invalid operation: nil == "abc" (mismatched types <T> and string)`,
	)

}

// Test Nil == Nil
func TestCheckBinaryExprNilEqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil == nil`, env,
		`invalid operation: nil == nil (operator == not defined on nil)`,
	)

}

// Test Nil != Int
func TestCheckBinaryExprNilNeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil != 4`, env,
		`cannot convert nil to type int`,
		`invalid operation: nil != 4 (mismatched types <T> and int)`,
	)

}

// Test Nil != Rune
func TestCheckBinaryExprNilNeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil != '@'`, env,
		`cannot convert nil to type rune`,
		`invalid operation: nil != '@' (mismatched types <T> and rune)`,
	)

}

// Test Nil != Float
func TestCheckBinaryExprNilNeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil != 2.0`, env,
		`cannot convert nil to type float64`,
		`invalid operation: nil != 2 (mismatched types <T> and float64)`,
	)

}

// Test Nil != Complex
func TestCheckBinaryExprNilNeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil != 8.0i`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: nil != 8i (mismatched types <T> and complex128)`,
	)

}

// Test Nil != Bool
func TestCheckBinaryExprNilNeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil != true`, env,
		`cannot convert nil to type int`,
		`cannot convert true to type int`,
		`invalid operation: nil != true (mismatched types <T> and bool)`,
	)

}

// Test Nil != String
func TestCheckBinaryExprNilNeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil != "abc"`, env,
		`cannot convert nil to type int`,
		`cannot convert "abc" to type int`,
		`invalid operation: nil != "abc" (mismatched types <T> and string)`,
	)

}

// Test Nil != Nil
func TestCheckBinaryExprNilNeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil != nil`, env,
		`invalid operation: nil != nil (operator != not defined on nil)`,
	)

}

// Test Nil <= Int
func TestCheckBinaryExprNilLeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil <= 4`, env,
		`cannot convert nil to type int`,
		`invalid operation: nil <= 4 (mismatched types <T> and int)`,
	)

}

// Test Nil <= Rune
func TestCheckBinaryExprNilLeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil <= '@'`, env,
		`cannot convert nil to type rune`,
		`invalid operation: nil <= '@' (mismatched types <T> and rune)`,
	)

}

// Test Nil <= Float
func TestCheckBinaryExprNilLeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil <= 2.0`, env,
		`cannot convert nil to type float64`,
		`invalid operation: nil <= 2 (mismatched types <T> and float64)`,
	)

}

// Test Nil <= Complex
func TestCheckBinaryExprNilLeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil <= 8.0i`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: nil <= 8i (mismatched types <T> and complex128)`,
	)

}

// Test Nil <= Bool
func TestCheckBinaryExprNilLeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil <= true`, env,
		`cannot convert nil to type int`,
		`cannot convert true to type int`,
		`invalid operation: nil <= true (mismatched types <T> and bool)`,
	)

}

// Test Nil <= String
func TestCheckBinaryExprNilLeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil <= "abc"`, env,
		`cannot convert nil to type int`,
		`cannot convert "abc" to type int`,
		`invalid operation: nil <= "abc" (mismatched types <T> and string)`,
	)

}

// Test Nil <= Nil
func TestCheckBinaryExprNilLeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil <= nil`, env,
		`invalid operation: nil <= nil (operator <= not defined on nil)`,
	)

}

// Test Nil >= Int
func TestCheckBinaryExprNilGeqInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil >= 4`, env,
		`cannot convert nil to type int`,
		`invalid operation: nil >= 4 (mismatched types <T> and int)`,
	)

}

// Test Nil >= Rune
func TestCheckBinaryExprNilGeqRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil >= '@'`, env,
		`cannot convert nil to type rune`,
		`invalid operation: nil >= '@' (mismatched types <T> and rune)`,
	)

}

// Test Nil >= Float
func TestCheckBinaryExprNilGeqFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil >= 2.0`, env,
		`cannot convert nil to type float64`,
		`invalid operation: nil >= 2 (mismatched types <T> and float64)`,
	)

}

// Test Nil >= Complex
func TestCheckBinaryExprNilGeqComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil >= 8.0i`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: nil >= 8i (mismatched types <T> and complex128)`,
	)

}

// Test Nil >= Bool
func TestCheckBinaryExprNilGeqBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil >= true`, env,
		`cannot convert nil to type int`,
		`cannot convert true to type int`,
		`invalid operation: nil >= true (mismatched types <T> and bool)`,
	)

}

// Test Nil >= String
func TestCheckBinaryExprNilGeqString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil >= "abc"`, env,
		`cannot convert nil to type int`,
		`cannot convert "abc" to type int`,
		`invalid operation: nil >= "abc" (mismatched types <T> and string)`,
	)

}

// Test Nil >= Nil
func TestCheckBinaryExprNilGeqNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil >= nil`, env,
		`invalid operation: nil >= nil (operator >= not defined on nil)`,
	)

}

// Test Nil < Int
func TestCheckBinaryExprNilLssInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil < 4`, env,
		`cannot convert nil to type int`,
		`invalid operation: nil < 4 (mismatched types <T> and int)`,
	)

}

// Test Nil < Rune
func TestCheckBinaryExprNilLssRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil < '@'`, env,
		`cannot convert nil to type rune`,
		`invalid operation: nil < '@' (mismatched types <T> and rune)`,
	)

}

// Test Nil < Float
func TestCheckBinaryExprNilLssFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil < 2.0`, env,
		`cannot convert nil to type float64`,
		`invalid operation: nil < 2 (mismatched types <T> and float64)`,
	)

}

// Test Nil < Complex
func TestCheckBinaryExprNilLssComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil < 8.0i`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: nil < 8i (mismatched types <T> and complex128)`,
	)

}

// Test Nil < Bool
func TestCheckBinaryExprNilLssBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil < true`, env,
		`cannot convert nil to type int`,
		`cannot convert true to type int`,
		`invalid operation: nil < true (mismatched types <T> and bool)`,
	)

}

// Test Nil < String
func TestCheckBinaryExprNilLssString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil < "abc"`, env,
		`cannot convert nil to type int`,
		`cannot convert "abc" to type int`,
		`invalid operation: nil < "abc" (mismatched types <T> and string)`,
	)

}

// Test Nil < Nil
func TestCheckBinaryExprNilLssNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil < nil`, env,
		`invalid operation: nil < nil (operator < not defined on nil)`,
	)

}

// Test Nil > Int
func TestCheckBinaryExprNilGtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil > 4`, env,
		`cannot convert nil to type int`,
		`invalid operation: nil > 4 (mismatched types <T> and int)`,
	)

}

// Test Nil > Rune
func TestCheckBinaryExprNilGtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil > '@'`, env,
		`cannot convert nil to type rune`,
		`invalid operation: nil > '@' (mismatched types <T> and rune)`,
	)

}

// Test Nil > Float
func TestCheckBinaryExprNilGtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil > 2.0`, env,
		`cannot convert nil to type float64`,
		`invalid operation: nil > 2 (mismatched types <T> and float64)`,
	)

}

// Test Nil > Complex
func TestCheckBinaryExprNilGtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil > 8.0i`, env,
		`cannot convert nil to type complex128`,
		`invalid operation: nil > 8i (mismatched types <T> and complex128)`,
	)

}

// Test Nil > Bool
func TestCheckBinaryExprNilGtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil > true`, env,
		`cannot convert nil to type int`,
		`cannot convert true to type int`,
		`invalid operation: nil > true (mismatched types <T> and bool)`,
	)

}

// Test Nil > String
func TestCheckBinaryExprNilGtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil > "abc"`, env,
		`cannot convert nil to type int`,
		`cannot convert "abc" to type int`,
		`invalid operation: nil > "abc" (mismatched types <T> and string)`,
	)

}

// Test Nil > Nil
func TestCheckBinaryExprNilGtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil > nil`, env,
		`invalid operation: nil > nil (operator > not defined on nil)`,
	)

}

// Test Nil >> Int
func TestCheckBinaryExprNilRhlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil >> 4`, env,
		`invalid operation: nil >> 4 (shift of type nil)`,
	)

}

// Test Nil >> Rune
func TestCheckBinaryExprNilRhlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil >> '@'`, env,
		`invalid operation: nil >> 64 (shift of type nil)`,
	)

}

// Test Nil >> Float
func TestCheckBinaryExprNilRhlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil >> 2.0`, env,
		`invalid operation: nil >> 2 (shift of type nil)`,
	)

}

// Test Nil >> Complex
func TestCheckBinaryExprNilRhlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil >> 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: nil >> 0 (shift of type nil)`,
	)

}

// Test Nil >> Bool
func TestCheckBinaryExprNilRhlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil >> true`, env,
		`invalid operation: nil >> true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Nil >> String
func TestCheckBinaryExprNilRhlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil >> "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: nil >> "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Nil >> Nil
func TestCheckBinaryExprNilRhlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil >> nil`, env,
		`cannot convert nil to type uint`,
	)

}
