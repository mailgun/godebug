package eval

import (
	"testing"
	"reflect"
)

// Test Int8 + Int
func TestCheckBinaryTypedExprInt8AddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + 4`, env,
		`constant 131 overflows int8`,
	)

}

// Test Int8 + Rune
func TestCheckBinaryTypedExprInt8AddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + '@'`, env,
		`constant 191 overflows int8`,
	)

}

// Test Int8 + Float
func TestCheckBinaryTypedExprInt8AddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + 2.0`, env,
		`constant 129 overflows int8`,
	)

}

// Test Int8 + Complex
func TestCheckBinaryTypedExprInt8AddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int8 + Bool
func TestCheckBinaryTypedExprInt8AddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + true`, env,
		`cannot convert true to type int8`,
		`invalid operation: 127 + true (mismatched types int8 and bool)`,
	)

}

// Test Int8 + String
func TestCheckBinaryTypedExprInt8AddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + "abc"`, env,
		`cannot convert "abc" to type int8`,
		`invalid operation: 127 + "abc" (mismatched types int8 and string)`,
	)

}

// Test Int8 + Nil
func TestCheckBinaryTypedExprInt8AddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + nil`, env,
		`cannot convert nil to type int8`,
	)

}

// Test Int8 + Int8
func TestCheckBinaryTypedExprInt8AddInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + int8(0x7f)`, env,
		`constant 254 overflows int8`,
	)

}

// Test Int8 + Int16
func TestCheckBinaryTypedExprInt8AddInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + int16(0x7fff)`, env,
		`invalid operation: 127 + 32767 (mismatched types int8 and int16)`,
	)

}

// Test Int8 + Int32
func TestCheckBinaryTypedExprInt8AddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + int32(0x7fffffff)`, env,
		`invalid operation: 127 + 2147483647 (mismatched types int8 and int32)`,
	)

}

// Test Int8 + Int64
func TestCheckBinaryTypedExprInt8AddInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + int64(0x7fffffffffffffff)`, env,
		`invalid operation: 127 + 9223372036854775807 (mismatched types int8 and int64)`,
	)

}

// Test Int8 + Uint8
func TestCheckBinaryTypedExprInt8AddUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + uint8(0xff)`, env,
		`invalid operation: 127 + 255 (mismatched types int8 and uint8)`,
	)

}

// Test Int8 + Uint16
func TestCheckBinaryTypedExprInt8AddUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + uint16(0xffff)`, env,
		`invalid operation: 127 + 65535 (mismatched types int8 and uint16)`,
	)

}

// Test Int8 + Uint32
func TestCheckBinaryTypedExprInt8AddUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + uint32(0xffffffff)`, env,
		`invalid operation: 127 + 4294967295 (mismatched types int8 and uint32)`,
	)

}

// Test Int8 + Uint64
func TestCheckBinaryTypedExprInt8AddUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + uint64(0xffffffffffffffff)`, env,
		`invalid operation: 127 + 18446744073709551615 (mismatched types int8 and uint64)`,
	)

}

// Test Int8 + Float32
func TestCheckBinaryTypedExprInt8AddFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + float32(0xffffffff)`, env,
		`invalid operation: 127 + 4.29497e+09 (mismatched types int8 and float32)`,
	)

}

// Test Int8 + Float64
func TestCheckBinaryTypedExprInt8AddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + float64(0xffffffff)`, env,
		`invalid operation: 127 + 4.29497e+09 (mismatched types int8 and float64)`,
	)

}

// Test Int8 + Complex64
func TestCheckBinaryTypedExprInt8AddComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 127 + (4.29497e+09+4.29497e+09i) (mismatched types int8 and complex64)`,
	)

}

// Test Int8 + Complex128
func TestCheckBinaryTypedExprInt8AddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 127 + (4.29497e+09+4.29497e+09i) (mismatched types int8 and complex128)`,
	)

}

// Test Int8 + Rune32
func TestCheckBinaryTypedExprInt8AddRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + rune(0x7fffffff)`, env,
		`invalid operation: 127 + rune(2147483647) (mismatched types int8 and rune)`,
	)

}

// Test Int8 + StringT
func TestCheckBinaryTypedExprInt8AddStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + string("abc")`, env,
		`invalid operation: 127 + "abc" (mismatched types int8 and string)`,
	)

}

// Test Int8 + BoolT
func TestCheckBinaryTypedExprInt8AddBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) + bool(true)`, env,
		`invalid operation: 127 + true (mismatched types int8 and bool)`,
	)

}

// Test Int8 - Int
func TestCheckBinaryTypedExprInt8SubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) - 4`, env, int8(0x7f) - 4, reflect.TypeOf(int8(0x7f) - 4))
}

// Test Int8 - Rune
func TestCheckBinaryTypedExprInt8SubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) - '@'`, env, int8(0x7f) - '@', reflect.TypeOf(int8(0x7f) - '@'))
}

// Test Int8 - Float
func TestCheckBinaryTypedExprInt8SubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) - 2.0`, env, int8(0x7f) - 2.0, reflect.TypeOf(int8(0x7f) - 2.0))
}

// Test Int8 - Complex
func TestCheckBinaryTypedExprInt8SubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int8 - Bool
func TestCheckBinaryTypedExprInt8SubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - true`, env,
		`cannot convert true to type int8`,
		`invalid operation: 127 - true (mismatched types int8 and bool)`,
	)

}

// Test Int8 - String
func TestCheckBinaryTypedExprInt8SubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - "abc"`, env,
		`cannot convert "abc" to type int8`,
		`invalid operation: 127 - "abc" (mismatched types int8 and string)`,
	)

}

// Test Int8 - Nil
func TestCheckBinaryTypedExprInt8SubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - nil`, env,
		`cannot convert nil to type int8`,
	)

}

// Test Int8 - Int8
func TestCheckBinaryTypedExprInt8SubInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) - int8(0x7f)`, env, int8(0x7f) - int8(0x7f), reflect.TypeOf(int8(0x7f) - int8(0x7f)))
}

// Test Int8 - Int16
func TestCheckBinaryTypedExprInt8SubInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - int16(0x7fff)`, env,
		`invalid operation: 127 - 32767 (mismatched types int8 and int16)`,
	)

}

// Test Int8 - Int32
func TestCheckBinaryTypedExprInt8SubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - int32(0x7fffffff)`, env,
		`invalid operation: 127 - 2147483647 (mismatched types int8 and int32)`,
	)

}

// Test Int8 - Int64
func TestCheckBinaryTypedExprInt8SubInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - int64(0x7fffffffffffffff)`, env,
		`invalid operation: 127 - 9223372036854775807 (mismatched types int8 and int64)`,
	)

}

// Test Int8 - Uint8
func TestCheckBinaryTypedExprInt8SubUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - uint8(0xff)`, env,
		`invalid operation: 127 - 255 (mismatched types int8 and uint8)`,
	)

}

// Test Int8 - Uint16
func TestCheckBinaryTypedExprInt8SubUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - uint16(0xffff)`, env,
		`invalid operation: 127 - 65535 (mismatched types int8 and uint16)`,
	)

}

// Test Int8 - Uint32
func TestCheckBinaryTypedExprInt8SubUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - uint32(0xffffffff)`, env,
		`invalid operation: 127 - 4294967295 (mismatched types int8 and uint32)`,
	)

}

// Test Int8 - Uint64
func TestCheckBinaryTypedExprInt8SubUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - uint64(0xffffffffffffffff)`, env,
		`invalid operation: 127 - 18446744073709551615 (mismatched types int8 and uint64)`,
	)

}

// Test Int8 - Float32
func TestCheckBinaryTypedExprInt8SubFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - float32(0xffffffff)`, env,
		`invalid operation: 127 - 4.29497e+09 (mismatched types int8 and float32)`,
	)

}

// Test Int8 - Float64
func TestCheckBinaryTypedExprInt8SubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - float64(0xffffffff)`, env,
		`invalid operation: 127 - 4.29497e+09 (mismatched types int8 and float64)`,
	)

}

// Test Int8 - Complex64
func TestCheckBinaryTypedExprInt8SubComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 127 - (4.29497e+09+4.29497e+09i) (mismatched types int8 and complex64)`,
	)

}

// Test Int8 - Complex128
func TestCheckBinaryTypedExprInt8SubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 127 - (4.29497e+09+4.29497e+09i) (mismatched types int8 and complex128)`,
	)

}

// Test Int8 - Rune32
func TestCheckBinaryTypedExprInt8SubRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - rune(0x7fffffff)`, env,
		`invalid operation: 127 - rune(2147483647) (mismatched types int8 and rune)`,
	)

}

// Test Int8 - StringT
func TestCheckBinaryTypedExprInt8SubStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - string("abc")`, env,
		`invalid operation: 127 - "abc" (mismatched types int8 and string)`,
	)

}

// Test Int8 - BoolT
func TestCheckBinaryTypedExprInt8SubBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) - bool(true)`, env,
		`invalid operation: 127 - true (mismatched types int8 and bool)`,
	)

}

// Test Int8 & Int
func TestCheckBinaryTypedExprInt8AndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) & 4`, env, int8(0x7f) & 4, reflect.TypeOf(int8(0x7f) & 4))
}

// Test Int8 & Rune
func TestCheckBinaryTypedExprInt8AndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) & '@'`, env, int8(0x7f) & '@', reflect.TypeOf(int8(0x7f) & '@'))
}

// Test Int8 & Float
func TestCheckBinaryTypedExprInt8AndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) & 2.0`, env, int8(0x7f) & 2.0, reflect.TypeOf(int8(0x7f) & 2.0))
}

// Test Int8 & Complex
func TestCheckBinaryTypedExprInt8AndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int8 & Bool
func TestCheckBinaryTypedExprInt8AndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & true`, env,
		`cannot convert true to type int8`,
		`invalid operation: 127 & true (mismatched types int8 and bool)`,
	)

}

// Test Int8 & String
func TestCheckBinaryTypedExprInt8AndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & "abc"`, env,
		`cannot convert "abc" to type int8`,
		`invalid operation: 127 & "abc" (mismatched types int8 and string)`,
	)

}

// Test Int8 & Nil
func TestCheckBinaryTypedExprInt8AndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & nil`, env,
		`cannot convert nil to type int8`,
	)

}

// Test Int8 & Int8
func TestCheckBinaryTypedExprInt8AndInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) & int8(0x7f)`, env, int8(0x7f) & int8(0x7f), reflect.TypeOf(int8(0x7f) & int8(0x7f)))
}

// Test Int8 & Int16
func TestCheckBinaryTypedExprInt8AndInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & int16(0x7fff)`, env,
		`invalid operation: 127 & 32767 (mismatched types int8 and int16)`,
	)

}

// Test Int8 & Int32
func TestCheckBinaryTypedExprInt8AndInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & int32(0x7fffffff)`, env,
		`invalid operation: 127 & 2147483647 (mismatched types int8 and int32)`,
	)

}

// Test Int8 & Int64
func TestCheckBinaryTypedExprInt8AndInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & int64(0x7fffffffffffffff)`, env,
		`invalid operation: 127 & 9223372036854775807 (mismatched types int8 and int64)`,
	)

}

// Test Int8 & Uint8
func TestCheckBinaryTypedExprInt8AndUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & uint8(0xff)`, env,
		`invalid operation: 127 & 255 (mismatched types int8 and uint8)`,
	)

}

// Test Int8 & Uint16
func TestCheckBinaryTypedExprInt8AndUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & uint16(0xffff)`, env,
		`invalid operation: 127 & 65535 (mismatched types int8 and uint16)`,
	)

}

// Test Int8 & Uint32
func TestCheckBinaryTypedExprInt8AndUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & uint32(0xffffffff)`, env,
		`invalid operation: 127 & 4294967295 (mismatched types int8 and uint32)`,
	)

}

// Test Int8 & Uint64
func TestCheckBinaryTypedExprInt8AndUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & uint64(0xffffffffffffffff)`, env,
		`invalid operation: 127 & 18446744073709551615 (mismatched types int8 and uint64)`,
	)

}

// Test Int8 & Float32
func TestCheckBinaryTypedExprInt8AndFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & float32(0xffffffff)`, env,
		`invalid operation: 127 & 4.29497e+09 (mismatched types int8 and float32)`,
	)

}

// Test Int8 & Float64
func TestCheckBinaryTypedExprInt8AndFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & float64(0xffffffff)`, env,
		`invalid operation: 127 & 4.29497e+09 (mismatched types int8 and float64)`,
	)

}

// Test Int8 & Complex64
func TestCheckBinaryTypedExprInt8AndComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 127 & (4.29497e+09+4.29497e+09i) (mismatched types int8 and complex64)`,
	)

}

// Test Int8 & Complex128
func TestCheckBinaryTypedExprInt8AndComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 127 & (4.29497e+09+4.29497e+09i) (mismatched types int8 and complex128)`,
	)

}

// Test Int8 & Rune32
func TestCheckBinaryTypedExprInt8AndRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & rune(0x7fffffff)`, env,
		`invalid operation: 127 & rune(2147483647) (mismatched types int8 and rune)`,
	)

}

// Test Int8 & StringT
func TestCheckBinaryTypedExprInt8AndStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & string("abc")`, env,
		`invalid operation: 127 & "abc" (mismatched types int8 and string)`,
	)

}

// Test Int8 & BoolT
func TestCheckBinaryTypedExprInt8AndBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) & bool(true)`, env,
		`invalid operation: 127 & true (mismatched types int8 and bool)`,
	)

}

// Test Int8 % Int
func TestCheckBinaryTypedExprInt8RemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) % 4`, env, int8(0x7f) % 4, reflect.TypeOf(int8(0x7f) % 4))
}

// Test Int8 % Rune
func TestCheckBinaryTypedExprInt8RemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) % '@'`, env, int8(0x7f) % '@', reflect.TypeOf(int8(0x7f) % '@'))
}

// Test Int8 % Float
func TestCheckBinaryTypedExprInt8RemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) % 2.0`, env, int8(0x7f) % 2.0, reflect.TypeOf(int8(0x7f) % 2.0))
}

// Test Int8 % Complex
func TestCheckBinaryTypedExprInt8RemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % 8.0i`, env,
		`constant 0+8i truncated to real`,
		`division by zero`,
	)

}

// Test Int8 % Bool
func TestCheckBinaryTypedExprInt8RemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % true`, env,
		`cannot convert true to type int8`,
		`invalid operation: 127 % true (mismatched types int8 and bool)`,
	)

}

// Test Int8 % String
func TestCheckBinaryTypedExprInt8RemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % "abc"`, env,
		`cannot convert "abc" to type int8`,
		`invalid operation: 127 % "abc" (mismatched types int8 and string)`,
	)

}

// Test Int8 % Nil
func TestCheckBinaryTypedExprInt8RemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % nil`, env,
		`cannot convert nil to type int8`,
	)

}

// Test Int8 % Int8
func TestCheckBinaryTypedExprInt8RemInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) % int8(0x7f)`, env, int8(0x7f) % int8(0x7f), reflect.TypeOf(int8(0x7f) % int8(0x7f)))
}

// Test Int8 % Int16
func TestCheckBinaryTypedExprInt8RemInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % int16(0x7fff)`, env,
		`invalid operation: 127 % 32767 (mismatched types int8 and int16)`,
	)

}

// Test Int8 % Int32
func TestCheckBinaryTypedExprInt8RemInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % int32(0x7fffffff)`, env,
		`invalid operation: 127 % 2147483647 (mismatched types int8 and int32)`,
	)

}

// Test Int8 % Int64
func TestCheckBinaryTypedExprInt8RemInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % int64(0x7fffffffffffffff)`, env,
		`invalid operation: 127 % 9223372036854775807 (mismatched types int8 and int64)`,
	)

}

// Test Int8 % Uint8
func TestCheckBinaryTypedExprInt8RemUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % uint8(0xff)`, env,
		`invalid operation: 127 % 255 (mismatched types int8 and uint8)`,
	)

}

// Test Int8 % Uint16
func TestCheckBinaryTypedExprInt8RemUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % uint16(0xffff)`, env,
		`invalid operation: 127 % 65535 (mismatched types int8 and uint16)`,
	)

}

// Test Int8 % Uint32
func TestCheckBinaryTypedExprInt8RemUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % uint32(0xffffffff)`, env,
		`invalid operation: 127 % 4294967295 (mismatched types int8 and uint32)`,
	)

}

// Test Int8 % Uint64
func TestCheckBinaryTypedExprInt8RemUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % uint64(0xffffffffffffffff)`, env,
		`invalid operation: 127 % 18446744073709551615 (mismatched types int8 and uint64)`,
	)

}

// Test Int8 % Float32
func TestCheckBinaryTypedExprInt8RemFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % float32(0xffffffff)`, env,
		`invalid operation: 127 % 4.29497e+09 (mismatched types int8 and float32)`,
	)

}

// Test Int8 % Float64
func TestCheckBinaryTypedExprInt8RemFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % float64(0xffffffff)`, env,
		`invalid operation: 127 % 4.29497e+09 (mismatched types int8 and float64)`,
	)

}

// Test Int8 % Complex64
func TestCheckBinaryTypedExprInt8RemComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 127 % (4.29497e+09+4.29497e+09i) (mismatched types int8 and complex64)`,
	)

}

// Test Int8 % Complex128
func TestCheckBinaryTypedExprInt8RemComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 127 % (4.29497e+09+4.29497e+09i) (mismatched types int8 and complex128)`,
	)

}

// Test Int8 % Rune32
func TestCheckBinaryTypedExprInt8RemRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % rune(0x7fffffff)`, env,
		`invalid operation: 127 % rune(2147483647) (mismatched types int8 and rune)`,
	)

}

// Test Int8 % StringT
func TestCheckBinaryTypedExprInt8RemStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % string("abc")`, env,
		`invalid operation: 127 % "abc" (mismatched types int8 and string)`,
	)

}

// Test Int8 % BoolT
func TestCheckBinaryTypedExprInt8RemBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) % bool(true)`, env,
		`invalid operation: 127 % true (mismatched types int8 and bool)`,
	)

}

// Test Int8 == Int
func TestCheckBinaryTypedExprInt8EqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) == 4`, env, int8(0x7f) == 4, reflect.TypeOf(int8(0x7f) == 4))
}

// Test Int8 == Rune
func TestCheckBinaryTypedExprInt8EqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) == '@'`, env, int8(0x7f) == '@', reflect.TypeOf(int8(0x7f) == '@'))
}

// Test Int8 == Float
func TestCheckBinaryTypedExprInt8EqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) == 2.0`, env, int8(0x7f) == 2.0, reflect.TypeOf(int8(0x7f) == 2.0))
}

// Test Int8 == Complex
func TestCheckBinaryTypedExprInt8EqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int8 == Bool
func TestCheckBinaryTypedExprInt8EqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == true`, env,
		`cannot convert true to type int8`,
		`invalid operation: 127 == true (mismatched types int8 and bool)`,
	)

}

// Test Int8 == String
func TestCheckBinaryTypedExprInt8EqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == "abc"`, env,
		`cannot convert "abc" to type int8`,
		`invalid operation: 127 == "abc" (mismatched types int8 and string)`,
	)

}

// Test Int8 == Nil
func TestCheckBinaryTypedExprInt8EqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == nil`, env,
		`cannot convert nil to type int8`,
	)

}

// Test Int8 == Int8
func TestCheckBinaryTypedExprInt8EqlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) == int8(0x7f)`, env, int8(0x7f) == int8(0x7f), reflect.TypeOf(int8(0x7f) == int8(0x7f)))
}

// Test Int8 == Int16
func TestCheckBinaryTypedExprInt8EqlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == int16(0x7fff)`, env,
		`invalid operation: 127 == 32767 (mismatched types int8 and int16)`,
	)

}

// Test Int8 == Int32
func TestCheckBinaryTypedExprInt8EqlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == int32(0x7fffffff)`, env,
		`invalid operation: 127 == 2147483647 (mismatched types int8 and int32)`,
	)

}

// Test Int8 == Int64
func TestCheckBinaryTypedExprInt8EqlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == int64(0x7fffffffffffffff)`, env,
		`invalid operation: 127 == 9223372036854775807 (mismatched types int8 and int64)`,
	)

}

// Test Int8 == Uint8
func TestCheckBinaryTypedExprInt8EqlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == uint8(0xff)`, env,
		`invalid operation: 127 == 255 (mismatched types int8 and uint8)`,
	)

}

// Test Int8 == Uint16
func TestCheckBinaryTypedExprInt8EqlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == uint16(0xffff)`, env,
		`invalid operation: 127 == 65535 (mismatched types int8 and uint16)`,
	)

}

// Test Int8 == Uint32
func TestCheckBinaryTypedExprInt8EqlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == uint32(0xffffffff)`, env,
		`invalid operation: 127 == 4294967295 (mismatched types int8 and uint32)`,
	)

}

// Test Int8 == Uint64
func TestCheckBinaryTypedExprInt8EqlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == uint64(0xffffffffffffffff)`, env,
		`invalid operation: 127 == 18446744073709551615 (mismatched types int8 and uint64)`,
	)

}

// Test Int8 == Float32
func TestCheckBinaryTypedExprInt8EqlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == float32(0xffffffff)`, env,
		`invalid operation: 127 == 4.29497e+09 (mismatched types int8 and float32)`,
	)

}

// Test Int8 == Float64
func TestCheckBinaryTypedExprInt8EqlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == float64(0xffffffff)`, env,
		`invalid operation: 127 == 4.29497e+09 (mismatched types int8 and float64)`,
	)

}

// Test Int8 == Complex64
func TestCheckBinaryTypedExprInt8EqlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 127 == (4.29497e+09+4.29497e+09i) (mismatched types int8 and complex64)`,
	)

}

// Test Int8 == Complex128
func TestCheckBinaryTypedExprInt8EqlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 127 == (4.29497e+09+4.29497e+09i) (mismatched types int8 and complex128)`,
	)

}

// Test Int8 == Rune32
func TestCheckBinaryTypedExprInt8EqlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == rune(0x7fffffff)`, env,
		`invalid operation: 127 == rune(2147483647) (mismatched types int8 and rune)`,
	)

}

// Test Int8 == StringT
func TestCheckBinaryTypedExprInt8EqlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == string("abc")`, env,
		`invalid operation: 127 == "abc" (mismatched types int8 and string)`,
	)

}

// Test Int8 == BoolT
func TestCheckBinaryTypedExprInt8EqlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) == bool(true)`, env,
		`invalid operation: 127 == true (mismatched types int8 and bool)`,
	)

}

// Test Int8 > Int
func TestCheckBinaryTypedExprInt8GtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) > 4`, env, int8(0x7f) > 4, reflect.TypeOf(int8(0x7f) > 4))
}

// Test Int8 > Rune
func TestCheckBinaryTypedExprInt8GtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) > '@'`, env, int8(0x7f) > '@', reflect.TypeOf(int8(0x7f) > '@'))
}

// Test Int8 > Float
func TestCheckBinaryTypedExprInt8GtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) > 2.0`, env, int8(0x7f) > 2.0, reflect.TypeOf(int8(0x7f) > 2.0))
}

// Test Int8 > Complex
func TestCheckBinaryTypedExprInt8GtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int8 > Bool
func TestCheckBinaryTypedExprInt8GtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > true`, env,
		`cannot convert true to type int8`,
		`invalid operation: 127 > true (mismatched types int8 and bool)`,
	)

}

// Test Int8 > String
func TestCheckBinaryTypedExprInt8GtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > "abc"`, env,
		`cannot convert "abc" to type int8`,
		`invalid operation: 127 > "abc" (mismatched types int8 and string)`,
	)

}

// Test Int8 > Nil
func TestCheckBinaryTypedExprInt8GtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > nil`, env,
		`cannot convert nil to type int8`,
	)

}

// Test Int8 > Int8
func TestCheckBinaryTypedExprInt8GtrInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f) > int8(0x7f)`, env, int8(0x7f) > int8(0x7f), reflect.TypeOf(int8(0x7f) > int8(0x7f)))
}

// Test Int8 > Int16
func TestCheckBinaryTypedExprInt8GtrInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > int16(0x7fff)`, env,
		`invalid operation: 127 > 32767 (mismatched types int8 and int16)`,
	)

}

// Test Int8 > Int32
func TestCheckBinaryTypedExprInt8GtrInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > int32(0x7fffffff)`, env,
		`invalid operation: 127 > 2147483647 (mismatched types int8 and int32)`,
	)

}

// Test Int8 > Int64
func TestCheckBinaryTypedExprInt8GtrInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > int64(0x7fffffffffffffff)`, env,
		`invalid operation: 127 > 9223372036854775807 (mismatched types int8 and int64)`,
	)

}

// Test Int8 > Uint8
func TestCheckBinaryTypedExprInt8GtrUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > uint8(0xff)`, env,
		`invalid operation: 127 > 255 (mismatched types int8 and uint8)`,
	)

}

// Test Int8 > Uint16
func TestCheckBinaryTypedExprInt8GtrUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > uint16(0xffff)`, env,
		`invalid operation: 127 > 65535 (mismatched types int8 and uint16)`,
	)

}

// Test Int8 > Uint32
func TestCheckBinaryTypedExprInt8GtrUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > uint32(0xffffffff)`, env,
		`invalid operation: 127 > 4294967295 (mismatched types int8 and uint32)`,
	)

}

// Test Int8 > Uint64
func TestCheckBinaryTypedExprInt8GtrUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > uint64(0xffffffffffffffff)`, env,
		`invalid operation: 127 > 18446744073709551615 (mismatched types int8 and uint64)`,
	)

}

// Test Int8 > Float32
func TestCheckBinaryTypedExprInt8GtrFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > float32(0xffffffff)`, env,
		`invalid operation: 127 > 4.29497e+09 (mismatched types int8 and float32)`,
	)

}

// Test Int8 > Float64
func TestCheckBinaryTypedExprInt8GtrFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > float64(0xffffffff)`, env,
		`invalid operation: 127 > 4.29497e+09 (mismatched types int8 and float64)`,
	)

}

// Test Int8 > Complex64
func TestCheckBinaryTypedExprInt8GtrComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 127 > (4.29497e+09+4.29497e+09i) (mismatched types int8 and complex64)`,
	)

}

// Test Int8 > Complex128
func TestCheckBinaryTypedExprInt8GtrComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 127 > (4.29497e+09+4.29497e+09i) (mismatched types int8 and complex128)`,
	)

}

// Test Int8 > Rune32
func TestCheckBinaryTypedExprInt8GtrRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > rune(0x7fffffff)`, env,
		`invalid operation: 127 > rune(2147483647) (mismatched types int8 and rune)`,
	)

}

// Test Int8 > StringT
func TestCheckBinaryTypedExprInt8GtrStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > string("abc")`, env,
		`invalid operation: 127 > "abc" (mismatched types int8 and string)`,
	)

}

// Test Int8 > BoolT
func TestCheckBinaryTypedExprInt8GtrBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) > bool(true)`, env,
		`invalid operation: 127 > true (mismatched types int8 and bool)`,
	)

}

// Test Int8 << Int
func TestCheckBinaryTypedExprInt8ShlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << 4`, env,
		`constant 2032 overflows int8`,
	)

}

// Test Int8 << Rune
func TestCheckBinaryTypedExprInt8ShlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << '@'`, env,
		`constant 2342736497361113055232 overflows int8`,
	)

}

// Test Int8 << Float
func TestCheckBinaryTypedExprInt8ShlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << 2.0`, env,
		`constant 508 overflows int8`,
	)

}

// Test Int8 << Complex
func TestCheckBinaryTypedExprInt8ShlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int8 << Bool
func TestCheckBinaryTypedExprInt8ShlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << true`, env,
		`invalid operation: 127 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Int8 << String
func TestCheckBinaryTypedExprInt8ShlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: 127 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Int8 << Nil
func TestCheckBinaryTypedExprInt8ShlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Int8 << Int8
func TestCheckBinaryTypedExprInt8ShlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << int8(0x7f)`, env,
		`invalid operation: 127 << 127 (shift count type int8, must be unsigned integer)`,
	)

}

// Test Int8 << Int16
func TestCheckBinaryTypedExprInt8ShlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << int16(0x7fff)`, env,
		`invalid operation: 127 << 32767 (shift count type int16, must be unsigned integer)`,
	)

}

// Test Int8 << Int32
func TestCheckBinaryTypedExprInt8ShlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << int32(0x7fffffff)`, env,
		`invalid operation: 127 << 2147483647 (shift count type int32, must be unsigned integer)`,
	)

}

// Test Int8 << Int64
func TestCheckBinaryTypedExprInt8ShlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << int64(0x7fffffffffffffff)`, env,
		`invalid operation: 127 << 9223372036854775807 (shift count type int64, must be unsigned integer)`,
	)

}

// Test Int8 << Uint8
func TestCheckBinaryTypedExprInt8ShlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << uint8(0xff)`, env,
		`constant 7352797666569578409396757548051682148682644026268175816505556584502483732135936 overflows int8`,
	)

}

// Test Int8 << Uint16
func TestCheckBinaryTypedExprInt8ShlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << uint16(0xffff)`, env,
		`stupid shift: 65535`,
	)

}

// Test Int8 << Uint32
func TestCheckBinaryTypedExprInt8ShlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << uint32(0xffffffff)`, env,
		`stupid shift: 4294967295`,
	)

}

// Test Int8 << Uint64
func TestCheckBinaryTypedExprInt8ShlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << uint64(0xffffffffffffffff)`, env,
		`stupid shift: -1`,
	)

}

// Test Int8 << Float32
func TestCheckBinaryTypedExprInt8ShlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << float32(0xffffffff)`, env,
		`invalid operation: 127 << 4.29497e+09 (shift count type float32, must be unsigned integer)`,
	)

}

// Test Int8 << Float64
func TestCheckBinaryTypedExprInt8ShlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << float64(0xffffffff)`, env,
		`invalid operation: 127 << 4.29497e+09 (shift count type float64, must be unsigned integer)`,
	)

}

// Test Int8 << Complex64
func TestCheckBinaryTypedExprInt8ShlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 127 << (4.29497e+09+4.29497e+09i) (shift count type complex64, must be unsigned integer)`,
	)

}

// Test Int8 << Complex128
func TestCheckBinaryTypedExprInt8ShlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 127 << (4.29497e+09+4.29497e+09i) (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Int8 << Rune32
func TestCheckBinaryTypedExprInt8ShlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << rune(0x7fffffff)`, env,
		`invalid operation: 127 << rune(2147483647) (shift count type rune, must be unsigned integer)`,
	)

}

// Test Int8 << StringT
func TestCheckBinaryTypedExprInt8ShlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << string("abc")`, env,
		`invalid operation: 127 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Int8 << BoolT
func TestCheckBinaryTypedExprInt8ShlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7f) << bool(true)`, env,
		`invalid operation: 127 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Int16 + Int
func TestCheckBinaryTypedExprInt16AddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + 4`, env,
		`constant 32771 overflows int16`,
	)

}

// Test Int16 + Rune
func TestCheckBinaryTypedExprInt16AddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + '@'`, env,
		`constant 32831 overflows int16`,
	)

}

// Test Int16 + Float
func TestCheckBinaryTypedExprInt16AddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + 2.0`, env,
		`constant 32769 overflows int16`,
	)

}

// Test Int16 + Complex
func TestCheckBinaryTypedExprInt16AddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int16 + Bool
func TestCheckBinaryTypedExprInt16AddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + true`, env,
		`cannot convert true to type int16`,
		`invalid operation: 32767 + true (mismatched types int16 and bool)`,
	)

}

// Test Int16 + String
func TestCheckBinaryTypedExprInt16AddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + "abc"`, env,
		`cannot convert "abc" to type int16`,
		`invalid operation: 32767 + "abc" (mismatched types int16 and string)`,
	)

}

// Test Int16 + Nil
func TestCheckBinaryTypedExprInt16AddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + nil`, env,
		`cannot convert nil to type int16`,
	)

}

// Test Int16 + Int8
func TestCheckBinaryTypedExprInt16AddInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + int8(0x7f)`, env,
		`invalid operation: 32767 + 127 (mismatched types int16 and int8)`,
	)

}

// Test Int16 + Int16
func TestCheckBinaryTypedExprInt16AddInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + int16(0x7fff)`, env,
		`constant 65534 overflows int16`,
	)

}

// Test Int16 + Int32
func TestCheckBinaryTypedExprInt16AddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + int32(0x7fffffff)`, env,
		`invalid operation: 32767 + 2147483647 (mismatched types int16 and int32)`,
	)

}

// Test Int16 + Int64
func TestCheckBinaryTypedExprInt16AddInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + int64(0x7fffffffffffffff)`, env,
		`invalid operation: 32767 + 9223372036854775807 (mismatched types int16 and int64)`,
	)

}

// Test Int16 + Uint8
func TestCheckBinaryTypedExprInt16AddUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + uint8(0xff)`, env,
		`invalid operation: 32767 + 255 (mismatched types int16 and uint8)`,
	)

}

// Test Int16 + Uint16
func TestCheckBinaryTypedExprInt16AddUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + uint16(0xffff)`, env,
		`invalid operation: 32767 + 65535 (mismatched types int16 and uint16)`,
	)

}

// Test Int16 + Uint32
func TestCheckBinaryTypedExprInt16AddUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + uint32(0xffffffff)`, env,
		`invalid operation: 32767 + 4294967295 (mismatched types int16 and uint32)`,
	)

}

// Test Int16 + Uint64
func TestCheckBinaryTypedExprInt16AddUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + uint64(0xffffffffffffffff)`, env,
		`invalid operation: 32767 + 18446744073709551615 (mismatched types int16 and uint64)`,
	)

}

// Test Int16 + Float32
func TestCheckBinaryTypedExprInt16AddFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + float32(0xffffffff)`, env,
		`invalid operation: 32767 + 4.29497e+09 (mismatched types int16 and float32)`,
	)

}

// Test Int16 + Float64
func TestCheckBinaryTypedExprInt16AddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + float64(0xffffffff)`, env,
		`invalid operation: 32767 + 4.29497e+09 (mismatched types int16 and float64)`,
	)

}

// Test Int16 + Complex64
func TestCheckBinaryTypedExprInt16AddComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 32767 + (4.29497e+09+4.29497e+09i) (mismatched types int16 and complex64)`,
	)

}

// Test Int16 + Complex128
func TestCheckBinaryTypedExprInt16AddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 32767 + (4.29497e+09+4.29497e+09i) (mismatched types int16 and complex128)`,
	)

}

// Test Int16 + Rune32
func TestCheckBinaryTypedExprInt16AddRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + rune(0x7fffffff)`, env,
		`invalid operation: 32767 + rune(2147483647) (mismatched types int16 and rune)`,
	)

}

// Test Int16 + StringT
func TestCheckBinaryTypedExprInt16AddStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + string("abc")`, env,
		`invalid operation: 32767 + "abc" (mismatched types int16 and string)`,
	)

}

// Test Int16 + BoolT
func TestCheckBinaryTypedExprInt16AddBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) + bool(true)`, env,
		`invalid operation: 32767 + true (mismatched types int16 and bool)`,
	)

}

// Test Int16 - Int
func TestCheckBinaryTypedExprInt16SubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) - 4`, env, int16(0x7fff) - 4, reflect.TypeOf(int16(0x7fff) - 4))
}

// Test Int16 - Rune
func TestCheckBinaryTypedExprInt16SubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) - '@'`, env, int16(0x7fff) - '@', reflect.TypeOf(int16(0x7fff) - '@'))
}

// Test Int16 - Float
func TestCheckBinaryTypedExprInt16SubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) - 2.0`, env, int16(0x7fff) - 2.0, reflect.TypeOf(int16(0x7fff) - 2.0))
}

// Test Int16 - Complex
func TestCheckBinaryTypedExprInt16SubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int16 - Bool
func TestCheckBinaryTypedExprInt16SubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - true`, env,
		`cannot convert true to type int16`,
		`invalid operation: 32767 - true (mismatched types int16 and bool)`,
	)

}

// Test Int16 - String
func TestCheckBinaryTypedExprInt16SubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - "abc"`, env,
		`cannot convert "abc" to type int16`,
		`invalid operation: 32767 - "abc" (mismatched types int16 and string)`,
	)

}

// Test Int16 - Nil
func TestCheckBinaryTypedExprInt16SubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - nil`, env,
		`cannot convert nil to type int16`,
	)

}

// Test Int16 - Int8
func TestCheckBinaryTypedExprInt16SubInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - int8(0x7f)`, env,
		`invalid operation: 32767 - 127 (mismatched types int16 and int8)`,
	)

}

// Test Int16 - Int16
func TestCheckBinaryTypedExprInt16SubInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) - int16(0x7fff)`, env, int16(0x7fff) - int16(0x7fff), reflect.TypeOf(int16(0x7fff) - int16(0x7fff)))
}

// Test Int16 - Int32
func TestCheckBinaryTypedExprInt16SubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - int32(0x7fffffff)`, env,
		`invalid operation: 32767 - 2147483647 (mismatched types int16 and int32)`,
	)

}

// Test Int16 - Int64
func TestCheckBinaryTypedExprInt16SubInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - int64(0x7fffffffffffffff)`, env,
		`invalid operation: 32767 - 9223372036854775807 (mismatched types int16 and int64)`,
	)

}

// Test Int16 - Uint8
func TestCheckBinaryTypedExprInt16SubUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - uint8(0xff)`, env,
		`invalid operation: 32767 - 255 (mismatched types int16 and uint8)`,
	)

}

// Test Int16 - Uint16
func TestCheckBinaryTypedExprInt16SubUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - uint16(0xffff)`, env,
		`invalid operation: 32767 - 65535 (mismatched types int16 and uint16)`,
	)

}

// Test Int16 - Uint32
func TestCheckBinaryTypedExprInt16SubUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - uint32(0xffffffff)`, env,
		`invalid operation: 32767 - 4294967295 (mismatched types int16 and uint32)`,
	)

}

// Test Int16 - Uint64
func TestCheckBinaryTypedExprInt16SubUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - uint64(0xffffffffffffffff)`, env,
		`invalid operation: 32767 - 18446744073709551615 (mismatched types int16 and uint64)`,
	)

}

// Test Int16 - Float32
func TestCheckBinaryTypedExprInt16SubFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - float32(0xffffffff)`, env,
		`invalid operation: 32767 - 4.29497e+09 (mismatched types int16 and float32)`,
	)

}

// Test Int16 - Float64
func TestCheckBinaryTypedExprInt16SubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - float64(0xffffffff)`, env,
		`invalid operation: 32767 - 4.29497e+09 (mismatched types int16 and float64)`,
	)

}

// Test Int16 - Complex64
func TestCheckBinaryTypedExprInt16SubComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 32767 - (4.29497e+09+4.29497e+09i) (mismatched types int16 and complex64)`,
	)

}

// Test Int16 - Complex128
func TestCheckBinaryTypedExprInt16SubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 32767 - (4.29497e+09+4.29497e+09i) (mismatched types int16 and complex128)`,
	)

}

// Test Int16 - Rune32
func TestCheckBinaryTypedExprInt16SubRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - rune(0x7fffffff)`, env,
		`invalid operation: 32767 - rune(2147483647) (mismatched types int16 and rune)`,
	)

}

// Test Int16 - StringT
func TestCheckBinaryTypedExprInt16SubStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - string("abc")`, env,
		`invalid operation: 32767 - "abc" (mismatched types int16 and string)`,
	)

}

// Test Int16 - BoolT
func TestCheckBinaryTypedExprInt16SubBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) - bool(true)`, env,
		`invalid operation: 32767 - true (mismatched types int16 and bool)`,
	)

}

// Test Int16 & Int
func TestCheckBinaryTypedExprInt16AndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) & 4`, env, int16(0x7fff) & 4, reflect.TypeOf(int16(0x7fff) & 4))
}

// Test Int16 & Rune
func TestCheckBinaryTypedExprInt16AndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) & '@'`, env, int16(0x7fff) & '@', reflect.TypeOf(int16(0x7fff) & '@'))
}

// Test Int16 & Float
func TestCheckBinaryTypedExprInt16AndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) & 2.0`, env, int16(0x7fff) & 2.0, reflect.TypeOf(int16(0x7fff) & 2.0))
}

// Test Int16 & Complex
func TestCheckBinaryTypedExprInt16AndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int16 & Bool
func TestCheckBinaryTypedExprInt16AndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & true`, env,
		`cannot convert true to type int16`,
		`invalid operation: 32767 & true (mismatched types int16 and bool)`,
	)

}

// Test Int16 & String
func TestCheckBinaryTypedExprInt16AndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & "abc"`, env,
		`cannot convert "abc" to type int16`,
		`invalid operation: 32767 & "abc" (mismatched types int16 and string)`,
	)

}

// Test Int16 & Nil
func TestCheckBinaryTypedExprInt16AndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & nil`, env,
		`cannot convert nil to type int16`,
	)

}

// Test Int16 & Int8
func TestCheckBinaryTypedExprInt16AndInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & int8(0x7f)`, env,
		`invalid operation: 32767 & 127 (mismatched types int16 and int8)`,
	)

}

// Test Int16 & Int16
func TestCheckBinaryTypedExprInt16AndInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) & int16(0x7fff)`, env, int16(0x7fff) & int16(0x7fff), reflect.TypeOf(int16(0x7fff) & int16(0x7fff)))
}

// Test Int16 & Int32
func TestCheckBinaryTypedExprInt16AndInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & int32(0x7fffffff)`, env,
		`invalid operation: 32767 & 2147483647 (mismatched types int16 and int32)`,
	)

}

// Test Int16 & Int64
func TestCheckBinaryTypedExprInt16AndInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & int64(0x7fffffffffffffff)`, env,
		`invalid operation: 32767 & 9223372036854775807 (mismatched types int16 and int64)`,
	)

}

// Test Int16 & Uint8
func TestCheckBinaryTypedExprInt16AndUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & uint8(0xff)`, env,
		`invalid operation: 32767 & 255 (mismatched types int16 and uint8)`,
	)

}

// Test Int16 & Uint16
func TestCheckBinaryTypedExprInt16AndUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & uint16(0xffff)`, env,
		`invalid operation: 32767 & 65535 (mismatched types int16 and uint16)`,
	)

}

// Test Int16 & Uint32
func TestCheckBinaryTypedExprInt16AndUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & uint32(0xffffffff)`, env,
		`invalid operation: 32767 & 4294967295 (mismatched types int16 and uint32)`,
	)

}

// Test Int16 & Uint64
func TestCheckBinaryTypedExprInt16AndUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & uint64(0xffffffffffffffff)`, env,
		`invalid operation: 32767 & 18446744073709551615 (mismatched types int16 and uint64)`,
	)

}

// Test Int16 & Float32
func TestCheckBinaryTypedExprInt16AndFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & float32(0xffffffff)`, env,
		`invalid operation: 32767 & 4.29497e+09 (mismatched types int16 and float32)`,
	)

}

// Test Int16 & Float64
func TestCheckBinaryTypedExprInt16AndFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & float64(0xffffffff)`, env,
		`invalid operation: 32767 & 4.29497e+09 (mismatched types int16 and float64)`,
	)

}

// Test Int16 & Complex64
func TestCheckBinaryTypedExprInt16AndComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 32767 & (4.29497e+09+4.29497e+09i) (mismatched types int16 and complex64)`,
	)

}

// Test Int16 & Complex128
func TestCheckBinaryTypedExprInt16AndComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 32767 & (4.29497e+09+4.29497e+09i) (mismatched types int16 and complex128)`,
	)

}

// Test Int16 & Rune32
func TestCheckBinaryTypedExprInt16AndRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & rune(0x7fffffff)`, env,
		`invalid operation: 32767 & rune(2147483647) (mismatched types int16 and rune)`,
	)

}

// Test Int16 & StringT
func TestCheckBinaryTypedExprInt16AndStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & string("abc")`, env,
		`invalid operation: 32767 & "abc" (mismatched types int16 and string)`,
	)

}

// Test Int16 & BoolT
func TestCheckBinaryTypedExprInt16AndBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) & bool(true)`, env,
		`invalid operation: 32767 & true (mismatched types int16 and bool)`,
	)

}

// Test Int16 % Int
func TestCheckBinaryTypedExprInt16RemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) % 4`, env, int16(0x7fff) % 4, reflect.TypeOf(int16(0x7fff) % 4))
}

// Test Int16 % Rune
func TestCheckBinaryTypedExprInt16RemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) % '@'`, env, int16(0x7fff) % '@', reflect.TypeOf(int16(0x7fff) % '@'))
}

// Test Int16 % Float
func TestCheckBinaryTypedExprInt16RemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) % 2.0`, env, int16(0x7fff) % 2.0, reflect.TypeOf(int16(0x7fff) % 2.0))
}

// Test Int16 % Complex
func TestCheckBinaryTypedExprInt16RemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % 8.0i`, env,
		`constant 0+8i truncated to real`,
		`division by zero`,
	)

}

// Test Int16 % Bool
func TestCheckBinaryTypedExprInt16RemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % true`, env,
		`cannot convert true to type int16`,
		`invalid operation: 32767 % true (mismatched types int16 and bool)`,
	)

}

// Test Int16 % String
func TestCheckBinaryTypedExprInt16RemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % "abc"`, env,
		`cannot convert "abc" to type int16`,
		`invalid operation: 32767 % "abc" (mismatched types int16 and string)`,
	)

}

// Test Int16 % Nil
func TestCheckBinaryTypedExprInt16RemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % nil`, env,
		`cannot convert nil to type int16`,
	)

}

// Test Int16 % Int8
func TestCheckBinaryTypedExprInt16RemInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % int8(0x7f)`, env,
		`invalid operation: 32767 % 127 (mismatched types int16 and int8)`,
	)

}

// Test Int16 % Int16
func TestCheckBinaryTypedExprInt16RemInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) % int16(0x7fff)`, env, int16(0x7fff) % int16(0x7fff), reflect.TypeOf(int16(0x7fff) % int16(0x7fff)))
}

// Test Int16 % Int32
func TestCheckBinaryTypedExprInt16RemInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % int32(0x7fffffff)`, env,
		`invalid operation: 32767 % 2147483647 (mismatched types int16 and int32)`,
	)

}

// Test Int16 % Int64
func TestCheckBinaryTypedExprInt16RemInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % int64(0x7fffffffffffffff)`, env,
		`invalid operation: 32767 % 9223372036854775807 (mismatched types int16 and int64)`,
	)

}

// Test Int16 % Uint8
func TestCheckBinaryTypedExprInt16RemUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % uint8(0xff)`, env,
		`invalid operation: 32767 % 255 (mismatched types int16 and uint8)`,
	)

}

// Test Int16 % Uint16
func TestCheckBinaryTypedExprInt16RemUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % uint16(0xffff)`, env,
		`invalid operation: 32767 % 65535 (mismatched types int16 and uint16)`,
	)

}

// Test Int16 % Uint32
func TestCheckBinaryTypedExprInt16RemUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % uint32(0xffffffff)`, env,
		`invalid operation: 32767 % 4294967295 (mismatched types int16 and uint32)`,
	)

}

// Test Int16 % Uint64
func TestCheckBinaryTypedExprInt16RemUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % uint64(0xffffffffffffffff)`, env,
		`invalid operation: 32767 % 18446744073709551615 (mismatched types int16 and uint64)`,
	)

}

// Test Int16 % Float32
func TestCheckBinaryTypedExprInt16RemFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % float32(0xffffffff)`, env,
		`invalid operation: 32767 % 4.29497e+09 (mismatched types int16 and float32)`,
	)

}

// Test Int16 % Float64
func TestCheckBinaryTypedExprInt16RemFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % float64(0xffffffff)`, env,
		`invalid operation: 32767 % 4.29497e+09 (mismatched types int16 and float64)`,
	)

}

// Test Int16 % Complex64
func TestCheckBinaryTypedExprInt16RemComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 32767 % (4.29497e+09+4.29497e+09i) (mismatched types int16 and complex64)`,
	)

}

// Test Int16 % Complex128
func TestCheckBinaryTypedExprInt16RemComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 32767 % (4.29497e+09+4.29497e+09i) (mismatched types int16 and complex128)`,
	)

}

// Test Int16 % Rune32
func TestCheckBinaryTypedExprInt16RemRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % rune(0x7fffffff)`, env,
		`invalid operation: 32767 % rune(2147483647) (mismatched types int16 and rune)`,
	)

}

// Test Int16 % StringT
func TestCheckBinaryTypedExprInt16RemStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % string("abc")`, env,
		`invalid operation: 32767 % "abc" (mismatched types int16 and string)`,
	)

}

// Test Int16 % BoolT
func TestCheckBinaryTypedExprInt16RemBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) % bool(true)`, env,
		`invalid operation: 32767 % true (mismatched types int16 and bool)`,
	)

}

// Test Int16 == Int
func TestCheckBinaryTypedExprInt16EqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) == 4`, env, int16(0x7fff) == 4, reflect.TypeOf(int16(0x7fff) == 4))
}

// Test Int16 == Rune
func TestCheckBinaryTypedExprInt16EqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) == '@'`, env, int16(0x7fff) == '@', reflect.TypeOf(int16(0x7fff) == '@'))
}

// Test Int16 == Float
func TestCheckBinaryTypedExprInt16EqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) == 2.0`, env, int16(0x7fff) == 2.0, reflect.TypeOf(int16(0x7fff) == 2.0))
}

// Test Int16 == Complex
func TestCheckBinaryTypedExprInt16EqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int16 == Bool
func TestCheckBinaryTypedExprInt16EqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == true`, env,
		`cannot convert true to type int16`,
		`invalid operation: 32767 == true (mismatched types int16 and bool)`,
	)

}

// Test Int16 == String
func TestCheckBinaryTypedExprInt16EqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == "abc"`, env,
		`cannot convert "abc" to type int16`,
		`invalid operation: 32767 == "abc" (mismatched types int16 and string)`,
	)

}

// Test Int16 == Nil
func TestCheckBinaryTypedExprInt16EqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == nil`, env,
		`cannot convert nil to type int16`,
	)

}

// Test Int16 == Int8
func TestCheckBinaryTypedExprInt16EqlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == int8(0x7f)`, env,
		`invalid operation: 32767 == 127 (mismatched types int16 and int8)`,
	)

}

// Test Int16 == Int16
func TestCheckBinaryTypedExprInt16EqlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) == int16(0x7fff)`, env, int16(0x7fff) == int16(0x7fff), reflect.TypeOf(int16(0x7fff) == int16(0x7fff)))
}

// Test Int16 == Int32
func TestCheckBinaryTypedExprInt16EqlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == int32(0x7fffffff)`, env,
		`invalid operation: 32767 == 2147483647 (mismatched types int16 and int32)`,
	)

}

// Test Int16 == Int64
func TestCheckBinaryTypedExprInt16EqlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == int64(0x7fffffffffffffff)`, env,
		`invalid operation: 32767 == 9223372036854775807 (mismatched types int16 and int64)`,
	)

}

// Test Int16 == Uint8
func TestCheckBinaryTypedExprInt16EqlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == uint8(0xff)`, env,
		`invalid operation: 32767 == 255 (mismatched types int16 and uint8)`,
	)

}

// Test Int16 == Uint16
func TestCheckBinaryTypedExprInt16EqlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == uint16(0xffff)`, env,
		`invalid operation: 32767 == 65535 (mismatched types int16 and uint16)`,
	)

}

// Test Int16 == Uint32
func TestCheckBinaryTypedExprInt16EqlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == uint32(0xffffffff)`, env,
		`invalid operation: 32767 == 4294967295 (mismatched types int16 and uint32)`,
	)

}

// Test Int16 == Uint64
func TestCheckBinaryTypedExprInt16EqlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == uint64(0xffffffffffffffff)`, env,
		`invalid operation: 32767 == 18446744073709551615 (mismatched types int16 and uint64)`,
	)

}

// Test Int16 == Float32
func TestCheckBinaryTypedExprInt16EqlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == float32(0xffffffff)`, env,
		`invalid operation: 32767 == 4.29497e+09 (mismatched types int16 and float32)`,
	)

}

// Test Int16 == Float64
func TestCheckBinaryTypedExprInt16EqlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == float64(0xffffffff)`, env,
		`invalid operation: 32767 == 4.29497e+09 (mismatched types int16 and float64)`,
	)

}

// Test Int16 == Complex64
func TestCheckBinaryTypedExprInt16EqlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 32767 == (4.29497e+09+4.29497e+09i) (mismatched types int16 and complex64)`,
	)

}

// Test Int16 == Complex128
func TestCheckBinaryTypedExprInt16EqlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 32767 == (4.29497e+09+4.29497e+09i) (mismatched types int16 and complex128)`,
	)

}

// Test Int16 == Rune32
func TestCheckBinaryTypedExprInt16EqlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == rune(0x7fffffff)`, env,
		`invalid operation: 32767 == rune(2147483647) (mismatched types int16 and rune)`,
	)

}

// Test Int16 == StringT
func TestCheckBinaryTypedExprInt16EqlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == string("abc")`, env,
		`invalid operation: 32767 == "abc" (mismatched types int16 and string)`,
	)

}

// Test Int16 == BoolT
func TestCheckBinaryTypedExprInt16EqlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) == bool(true)`, env,
		`invalid operation: 32767 == true (mismatched types int16 and bool)`,
	)

}

// Test Int16 > Int
func TestCheckBinaryTypedExprInt16GtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) > 4`, env, int16(0x7fff) > 4, reflect.TypeOf(int16(0x7fff) > 4))
}

// Test Int16 > Rune
func TestCheckBinaryTypedExprInt16GtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) > '@'`, env, int16(0x7fff) > '@', reflect.TypeOf(int16(0x7fff) > '@'))
}

// Test Int16 > Float
func TestCheckBinaryTypedExprInt16GtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) > 2.0`, env, int16(0x7fff) > 2.0, reflect.TypeOf(int16(0x7fff) > 2.0))
}

// Test Int16 > Complex
func TestCheckBinaryTypedExprInt16GtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int16 > Bool
func TestCheckBinaryTypedExprInt16GtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > true`, env,
		`cannot convert true to type int16`,
		`invalid operation: 32767 > true (mismatched types int16 and bool)`,
	)

}

// Test Int16 > String
func TestCheckBinaryTypedExprInt16GtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > "abc"`, env,
		`cannot convert "abc" to type int16`,
		`invalid operation: 32767 > "abc" (mismatched types int16 and string)`,
	)

}

// Test Int16 > Nil
func TestCheckBinaryTypedExprInt16GtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > nil`, env,
		`cannot convert nil to type int16`,
	)

}

// Test Int16 > Int8
func TestCheckBinaryTypedExprInt16GtrInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > int8(0x7f)`, env,
		`invalid operation: 32767 > 127 (mismatched types int16 and int8)`,
	)

}

// Test Int16 > Int16
func TestCheckBinaryTypedExprInt16GtrInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff) > int16(0x7fff)`, env, int16(0x7fff) > int16(0x7fff), reflect.TypeOf(int16(0x7fff) > int16(0x7fff)))
}

// Test Int16 > Int32
func TestCheckBinaryTypedExprInt16GtrInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > int32(0x7fffffff)`, env,
		`invalid operation: 32767 > 2147483647 (mismatched types int16 and int32)`,
	)

}

// Test Int16 > Int64
func TestCheckBinaryTypedExprInt16GtrInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > int64(0x7fffffffffffffff)`, env,
		`invalid operation: 32767 > 9223372036854775807 (mismatched types int16 and int64)`,
	)

}

// Test Int16 > Uint8
func TestCheckBinaryTypedExprInt16GtrUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > uint8(0xff)`, env,
		`invalid operation: 32767 > 255 (mismatched types int16 and uint8)`,
	)

}

// Test Int16 > Uint16
func TestCheckBinaryTypedExprInt16GtrUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > uint16(0xffff)`, env,
		`invalid operation: 32767 > 65535 (mismatched types int16 and uint16)`,
	)

}

// Test Int16 > Uint32
func TestCheckBinaryTypedExprInt16GtrUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > uint32(0xffffffff)`, env,
		`invalid operation: 32767 > 4294967295 (mismatched types int16 and uint32)`,
	)

}

// Test Int16 > Uint64
func TestCheckBinaryTypedExprInt16GtrUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > uint64(0xffffffffffffffff)`, env,
		`invalid operation: 32767 > 18446744073709551615 (mismatched types int16 and uint64)`,
	)

}

// Test Int16 > Float32
func TestCheckBinaryTypedExprInt16GtrFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > float32(0xffffffff)`, env,
		`invalid operation: 32767 > 4.29497e+09 (mismatched types int16 and float32)`,
	)

}

// Test Int16 > Float64
func TestCheckBinaryTypedExprInt16GtrFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > float64(0xffffffff)`, env,
		`invalid operation: 32767 > 4.29497e+09 (mismatched types int16 and float64)`,
	)

}

// Test Int16 > Complex64
func TestCheckBinaryTypedExprInt16GtrComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 32767 > (4.29497e+09+4.29497e+09i) (mismatched types int16 and complex64)`,
	)

}

// Test Int16 > Complex128
func TestCheckBinaryTypedExprInt16GtrComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 32767 > (4.29497e+09+4.29497e+09i) (mismatched types int16 and complex128)`,
	)

}

// Test Int16 > Rune32
func TestCheckBinaryTypedExprInt16GtrRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > rune(0x7fffffff)`, env,
		`invalid operation: 32767 > rune(2147483647) (mismatched types int16 and rune)`,
	)

}

// Test Int16 > StringT
func TestCheckBinaryTypedExprInt16GtrStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > string("abc")`, env,
		`invalid operation: 32767 > "abc" (mismatched types int16 and string)`,
	)

}

// Test Int16 > BoolT
func TestCheckBinaryTypedExprInt16GtrBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) > bool(true)`, env,
		`invalid operation: 32767 > true (mismatched types int16 and bool)`,
	)

}

// Test Int16 << Int
func TestCheckBinaryTypedExprInt16ShlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << 4`, env,
		`constant 524272 overflows int16`,
	)

}

// Test Int16 << Rune
func TestCheckBinaryTypedExprInt16ShlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << '@'`, env,
		`constant 604444463063240877801472 overflows int16`,
	)

}

// Test Int16 << Float
func TestCheckBinaryTypedExprInt16ShlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << 2.0`, env,
		`constant 131068 overflows int16`,
	)

}

// Test Int16 << Complex
func TestCheckBinaryTypedExprInt16ShlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int16 << Bool
func TestCheckBinaryTypedExprInt16ShlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << true`, env,
		`invalid operation: 32767 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Int16 << String
func TestCheckBinaryTypedExprInt16ShlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: 32767 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Int16 << Nil
func TestCheckBinaryTypedExprInt16ShlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Int16 << Int8
func TestCheckBinaryTypedExprInt16ShlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << int8(0x7f)`, env,
		`invalid operation: 32767 << 127 (shift count type int8, must be unsigned integer)`,
	)

}

// Test Int16 << Int16
func TestCheckBinaryTypedExprInt16ShlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << int16(0x7fff)`, env,
		`invalid operation: 32767 << 32767 (shift count type int16, must be unsigned integer)`,
	)

}

// Test Int16 << Int32
func TestCheckBinaryTypedExprInt16ShlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << int32(0x7fffffff)`, env,
		`invalid operation: 32767 << 2147483647 (shift count type int32, must be unsigned integer)`,
	)

}

// Test Int16 << Int64
func TestCheckBinaryTypedExprInt16ShlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << int64(0x7fffffffffffffff)`, env,
		`invalid operation: 32767 << 9223372036854775807 (shift count type int64, must be unsigned integer)`,
	)

}

// Test Int16 << Uint8
func TestCheckBinaryTypedExprInt16ShlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << uint8(0xff)`, env,
		`constant 1897079694019569887722075232889838338314048793769522180940453327593644759455891456 overflows int16`,
	)

}

// Test Int16 << Uint16
func TestCheckBinaryTypedExprInt16ShlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << uint16(0xffff)`, env,
		`stupid shift: 65535`,
	)

}

// Test Int16 << Uint32
func TestCheckBinaryTypedExprInt16ShlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << uint32(0xffffffff)`, env,
		`stupid shift: 4294967295`,
	)

}

// Test Int16 << Uint64
func TestCheckBinaryTypedExprInt16ShlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << uint64(0xffffffffffffffff)`, env,
		`stupid shift: -1`,
	)

}

// Test Int16 << Float32
func TestCheckBinaryTypedExprInt16ShlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << float32(0xffffffff)`, env,
		`invalid operation: 32767 << 4.29497e+09 (shift count type float32, must be unsigned integer)`,
	)

}

// Test Int16 << Float64
func TestCheckBinaryTypedExprInt16ShlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << float64(0xffffffff)`, env,
		`invalid operation: 32767 << 4.29497e+09 (shift count type float64, must be unsigned integer)`,
	)

}

// Test Int16 << Complex64
func TestCheckBinaryTypedExprInt16ShlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 32767 << (4.29497e+09+4.29497e+09i) (shift count type complex64, must be unsigned integer)`,
	)

}

// Test Int16 << Complex128
func TestCheckBinaryTypedExprInt16ShlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 32767 << (4.29497e+09+4.29497e+09i) (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Int16 << Rune32
func TestCheckBinaryTypedExprInt16ShlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << rune(0x7fffffff)`, env,
		`invalid operation: 32767 << rune(2147483647) (shift count type rune, must be unsigned integer)`,
	)

}

// Test Int16 << StringT
func TestCheckBinaryTypedExprInt16ShlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << string("abc")`, env,
		`invalid operation: 32767 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Int16 << BoolT
func TestCheckBinaryTypedExprInt16ShlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fff) << bool(true)`, env,
		`invalid operation: 32767 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Int32 + Int
func TestCheckBinaryTypedExprInt32AddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + 4`, env,
		`constant 2147483651 overflows int32`,
	)

}

// Test Int32 + Rune
func TestCheckBinaryTypedExprInt32AddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + '@'`, env,
		`constant 2147483711 overflows int32`,
	)

}

// Test Int32 + Float
func TestCheckBinaryTypedExprInt32AddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + 2.0`, env,
		`constant 2147483649 overflows int32`,
	)

}

// Test Int32 + Complex
func TestCheckBinaryTypedExprInt32AddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int32 + Bool
func TestCheckBinaryTypedExprInt32AddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + true`, env,
		`cannot convert true to type int32`,
		`invalid operation: 2147483647 + true (mismatched types int32 and bool)`,
	)

}

// Test Int32 + String
func TestCheckBinaryTypedExprInt32AddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + "abc"`, env,
		`cannot convert "abc" to type int32`,
		`invalid operation: 2147483647 + "abc" (mismatched types int32 and string)`,
	)

}

// Test Int32 + Nil
func TestCheckBinaryTypedExprInt32AddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + nil`, env,
		`cannot convert nil to type int32`,
	)

}

// Test Int32 + Int8
func TestCheckBinaryTypedExprInt32AddInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + int8(0x7f)`, env,
		`invalid operation: 2147483647 + 127 (mismatched types int32 and int8)`,
	)

}

// Test Int32 + Int16
func TestCheckBinaryTypedExprInt32AddInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + int16(0x7fff)`, env,
		`invalid operation: 2147483647 + 32767 (mismatched types int32 and int16)`,
	)

}

// Test Int32 + Int32
func TestCheckBinaryTypedExprInt32AddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + int32(0x7fffffff)`, env,
		`constant 4294967294 overflows int32`,
	)

}

// Test Int32 + Int64
func TestCheckBinaryTypedExprInt32AddInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + int64(0x7fffffffffffffff)`, env,
		`invalid operation: 2147483647 + 9223372036854775807 (mismatched types int32 and int64)`,
	)

}

// Test Int32 + Uint8
func TestCheckBinaryTypedExprInt32AddUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + uint8(0xff)`, env,
		`invalid operation: 2147483647 + 255 (mismatched types int32 and uint8)`,
	)

}

// Test Int32 + Uint16
func TestCheckBinaryTypedExprInt32AddUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + uint16(0xffff)`, env,
		`invalid operation: 2147483647 + 65535 (mismatched types int32 and uint16)`,
	)

}

// Test Int32 + Uint32
func TestCheckBinaryTypedExprInt32AddUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + uint32(0xffffffff)`, env,
		`invalid operation: 2147483647 + 4294967295 (mismatched types int32 and uint32)`,
	)

}

// Test Int32 + Uint64
func TestCheckBinaryTypedExprInt32AddUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + uint64(0xffffffffffffffff)`, env,
		`invalid operation: 2147483647 + 18446744073709551615 (mismatched types int32 and uint64)`,
	)

}

// Test Int32 + Float32
func TestCheckBinaryTypedExprInt32AddFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + float32(0xffffffff)`, env,
		`invalid operation: 2147483647 + 4.29497e+09 (mismatched types int32 and float32)`,
	)

}

// Test Int32 + Float64
func TestCheckBinaryTypedExprInt32AddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + float64(0xffffffff)`, env,
		`invalid operation: 2147483647 + 4.29497e+09 (mismatched types int32 and float64)`,
	)

}

// Test Int32 + Complex64
func TestCheckBinaryTypedExprInt32AddComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 2147483647 + (4.29497e+09+4.29497e+09i) (mismatched types int32 and complex64)`,
	)

}

// Test Int32 + Complex128
func TestCheckBinaryTypedExprInt32AddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 2147483647 + (4.29497e+09+4.29497e+09i) (mismatched types int32 and complex128)`,
	)

}

// Test Int32 + Rune32
func TestCheckBinaryTypedExprInt32AddRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + rune(0x7fffffff)`, env,
		`constant 4294967294 overflows int32`,
	)

}

// Test Int32 + StringT
func TestCheckBinaryTypedExprInt32AddStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + string("abc")`, env,
		`invalid operation: 2147483647 + "abc" (mismatched types int32 and string)`,
	)

}

// Test Int32 + BoolT
func TestCheckBinaryTypedExprInt32AddBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) + bool(true)`, env,
		`invalid operation: 2147483647 + true (mismatched types int32 and bool)`,
	)

}

// Test Int32 - Int
func TestCheckBinaryTypedExprInt32SubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) - 4`, env, int32(0x7fffffff) - 4, reflect.TypeOf(int32(0x7fffffff) - 4))
}

// Test Int32 - Rune
func TestCheckBinaryTypedExprInt32SubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) - '@'`, env, int32(0x7fffffff) - '@', reflect.TypeOf(int32(0x7fffffff) - '@'))
}

// Test Int32 - Float
func TestCheckBinaryTypedExprInt32SubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) - 2.0`, env, int32(0x7fffffff) - 2.0, reflect.TypeOf(int32(0x7fffffff) - 2.0))
}

// Test Int32 - Complex
func TestCheckBinaryTypedExprInt32SubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int32 - Bool
func TestCheckBinaryTypedExprInt32SubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - true`, env,
		`cannot convert true to type int32`,
		`invalid operation: 2147483647 - true (mismatched types int32 and bool)`,
	)

}

// Test Int32 - String
func TestCheckBinaryTypedExprInt32SubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - "abc"`, env,
		`cannot convert "abc" to type int32`,
		`invalid operation: 2147483647 - "abc" (mismatched types int32 and string)`,
	)

}

// Test Int32 - Nil
func TestCheckBinaryTypedExprInt32SubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - nil`, env,
		`cannot convert nil to type int32`,
	)

}

// Test Int32 - Int8
func TestCheckBinaryTypedExprInt32SubInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - int8(0x7f)`, env,
		`invalid operation: 2147483647 - 127 (mismatched types int32 and int8)`,
	)

}

// Test Int32 - Int16
func TestCheckBinaryTypedExprInt32SubInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - int16(0x7fff)`, env,
		`invalid operation: 2147483647 - 32767 (mismatched types int32 and int16)`,
	)

}

// Test Int32 - Int32
func TestCheckBinaryTypedExprInt32SubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) - int32(0x7fffffff)`, env, int32(0x7fffffff) - int32(0x7fffffff), reflect.TypeOf(int32(0x7fffffff) - int32(0x7fffffff)))
}

// Test Int32 - Int64
func TestCheckBinaryTypedExprInt32SubInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - int64(0x7fffffffffffffff)`, env,
		`invalid operation: 2147483647 - 9223372036854775807 (mismatched types int32 and int64)`,
	)

}

// Test Int32 - Uint8
func TestCheckBinaryTypedExprInt32SubUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - uint8(0xff)`, env,
		`invalid operation: 2147483647 - 255 (mismatched types int32 and uint8)`,
	)

}

// Test Int32 - Uint16
func TestCheckBinaryTypedExprInt32SubUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - uint16(0xffff)`, env,
		`invalid operation: 2147483647 - 65535 (mismatched types int32 and uint16)`,
	)

}

// Test Int32 - Uint32
func TestCheckBinaryTypedExprInt32SubUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - uint32(0xffffffff)`, env,
		`invalid operation: 2147483647 - 4294967295 (mismatched types int32 and uint32)`,
	)

}

// Test Int32 - Uint64
func TestCheckBinaryTypedExprInt32SubUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - uint64(0xffffffffffffffff)`, env,
		`invalid operation: 2147483647 - 18446744073709551615 (mismatched types int32 and uint64)`,
	)

}

// Test Int32 - Float32
func TestCheckBinaryTypedExprInt32SubFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - float32(0xffffffff)`, env,
		`invalid operation: 2147483647 - 4.29497e+09 (mismatched types int32 and float32)`,
	)

}

// Test Int32 - Float64
func TestCheckBinaryTypedExprInt32SubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - float64(0xffffffff)`, env,
		`invalid operation: 2147483647 - 4.29497e+09 (mismatched types int32 and float64)`,
	)

}

// Test Int32 - Complex64
func TestCheckBinaryTypedExprInt32SubComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 2147483647 - (4.29497e+09+4.29497e+09i) (mismatched types int32 and complex64)`,
	)

}

// Test Int32 - Complex128
func TestCheckBinaryTypedExprInt32SubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 2147483647 - (4.29497e+09+4.29497e+09i) (mismatched types int32 and complex128)`,
	)

}

// Test Int32 - Rune32
func TestCheckBinaryTypedExprInt32SubRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) - rune(0x7fffffff)`, env, int32(0x7fffffff) - rune(0x7fffffff), reflect.TypeOf(int32(0x7fffffff) - rune(0x7fffffff)))
}

// Test Int32 - StringT
func TestCheckBinaryTypedExprInt32SubStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - string("abc")`, env,
		`invalid operation: 2147483647 - "abc" (mismatched types int32 and string)`,
	)

}

// Test Int32 - BoolT
func TestCheckBinaryTypedExprInt32SubBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) - bool(true)`, env,
		`invalid operation: 2147483647 - true (mismatched types int32 and bool)`,
	)

}

// Test Int32 & Int
func TestCheckBinaryTypedExprInt32AndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) & 4`, env, int32(0x7fffffff) & 4, reflect.TypeOf(int32(0x7fffffff) & 4))
}

// Test Int32 & Rune
func TestCheckBinaryTypedExprInt32AndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) & '@'`, env, int32(0x7fffffff) & '@', reflect.TypeOf(int32(0x7fffffff) & '@'))
}

// Test Int32 & Float
func TestCheckBinaryTypedExprInt32AndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) & 2.0`, env, int32(0x7fffffff) & 2.0, reflect.TypeOf(int32(0x7fffffff) & 2.0))
}

// Test Int32 & Complex
func TestCheckBinaryTypedExprInt32AndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int32 & Bool
func TestCheckBinaryTypedExprInt32AndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & true`, env,
		`cannot convert true to type int32`,
		`invalid operation: 2147483647 & true (mismatched types int32 and bool)`,
	)

}

// Test Int32 & String
func TestCheckBinaryTypedExprInt32AndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & "abc"`, env,
		`cannot convert "abc" to type int32`,
		`invalid operation: 2147483647 & "abc" (mismatched types int32 and string)`,
	)

}

// Test Int32 & Nil
func TestCheckBinaryTypedExprInt32AndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & nil`, env,
		`cannot convert nil to type int32`,
	)

}

// Test Int32 & Int8
func TestCheckBinaryTypedExprInt32AndInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & int8(0x7f)`, env,
		`invalid operation: 2147483647 & 127 (mismatched types int32 and int8)`,
	)

}

// Test Int32 & Int16
func TestCheckBinaryTypedExprInt32AndInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & int16(0x7fff)`, env,
		`invalid operation: 2147483647 & 32767 (mismatched types int32 and int16)`,
	)

}

// Test Int32 & Int32
func TestCheckBinaryTypedExprInt32AndInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) & int32(0x7fffffff)`, env, int32(0x7fffffff) & int32(0x7fffffff), reflect.TypeOf(int32(0x7fffffff) & int32(0x7fffffff)))
}

// Test Int32 & Int64
func TestCheckBinaryTypedExprInt32AndInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & int64(0x7fffffffffffffff)`, env,
		`invalid operation: 2147483647 & 9223372036854775807 (mismatched types int32 and int64)`,
	)

}

// Test Int32 & Uint8
func TestCheckBinaryTypedExprInt32AndUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & uint8(0xff)`, env,
		`invalid operation: 2147483647 & 255 (mismatched types int32 and uint8)`,
	)

}

// Test Int32 & Uint16
func TestCheckBinaryTypedExprInt32AndUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & uint16(0xffff)`, env,
		`invalid operation: 2147483647 & 65535 (mismatched types int32 and uint16)`,
	)

}

// Test Int32 & Uint32
func TestCheckBinaryTypedExprInt32AndUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & uint32(0xffffffff)`, env,
		`invalid operation: 2147483647 & 4294967295 (mismatched types int32 and uint32)`,
	)

}

// Test Int32 & Uint64
func TestCheckBinaryTypedExprInt32AndUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & uint64(0xffffffffffffffff)`, env,
		`invalid operation: 2147483647 & 18446744073709551615 (mismatched types int32 and uint64)`,
	)

}

// Test Int32 & Float32
func TestCheckBinaryTypedExprInt32AndFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & float32(0xffffffff)`, env,
		`invalid operation: 2147483647 & 4.29497e+09 (mismatched types int32 and float32)`,
	)

}

// Test Int32 & Float64
func TestCheckBinaryTypedExprInt32AndFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & float64(0xffffffff)`, env,
		`invalid operation: 2147483647 & 4.29497e+09 (mismatched types int32 and float64)`,
	)

}

// Test Int32 & Complex64
func TestCheckBinaryTypedExprInt32AndComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 2147483647 & (4.29497e+09+4.29497e+09i) (mismatched types int32 and complex64)`,
	)

}

// Test Int32 & Complex128
func TestCheckBinaryTypedExprInt32AndComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 2147483647 & (4.29497e+09+4.29497e+09i) (mismatched types int32 and complex128)`,
	)

}

// Test Int32 & Rune32
func TestCheckBinaryTypedExprInt32AndRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) & rune(0x7fffffff)`, env, int32(0x7fffffff) & rune(0x7fffffff), reflect.TypeOf(int32(0x7fffffff) & rune(0x7fffffff)))
}

// Test Int32 & StringT
func TestCheckBinaryTypedExprInt32AndStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & string("abc")`, env,
		`invalid operation: 2147483647 & "abc" (mismatched types int32 and string)`,
	)

}

// Test Int32 & BoolT
func TestCheckBinaryTypedExprInt32AndBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) & bool(true)`, env,
		`invalid operation: 2147483647 & true (mismatched types int32 and bool)`,
	)

}

// Test Int32 % Int
func TestCheckBinaryTypedExprInt32RemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) % 4`, env, int32(0x7fffffff) % 4, reflect.TypeOf(int32(0x7fffffff) % 4))
}

// Test Int32 % Rune
func TestCheckBinaryTypedExprInt32RemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) % '@'`, env, int32(0x7fffffff) % '@', reflect.TypeOf(int32(0x7fffffff) % '@'))
}

// Test Int32 % Float
func TestCheckBinaryTypedExprInt32RemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) % 2.0`, env, int32(0x7fffffff) % 2.0, reflect.TypeOf(int32(0x7fffffff) % 2.0))
}

// Test Int32 % Complex
func TestCheckBinaryTypedExprInt32RemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % 8.0i`, env,
		`constant 0+8i truncated to real`,
		`division by zero`,
	)

}

// Test Int32 % Bool
func TestCheckBinaryTypedExprInt32RemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % true`, env,
		`cannot convert true to type int32`,
		`invalid operation: 2147483647 % true (mismatched types int32 and bool)`,
	)

}

// Test Int32 % String
func TestCheckBinaryTypedExprInt32RemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % "abc"`, env,
		`cannot convert "abc" to type int32`,
		`invalid operation: 2147483647 % "abc" (mismatched types int32 and string)`,
	)

}

// Test Int32 % Nil
func TestCheckBinaryTypedExprInt32RemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % nil`, env,
		`cannot convert nil to type int32`,
	)

}

// Test Int32 % Int8
func TestCheckBinaryTypedExprInt32RemInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % int8(0x7f)`, env,
		`invalid operation: 2147483647 % 127 (mismatched types int32 and int8)`,
	)

}

// Test Int32 % Int16
func TestCheckBinaryTypedExprInt32RemInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % int16(0x7fff)`, env,
		`invalid operation: 2147483647 % 32767 (mismatched types int32 and int16)`,
	)

}

// Test Int32 % Int32
func TestCheckBinaryTypedExprInt32RemInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) % int32(0x7fffffff)`, env, int32(0x7fffffff) % int32(0x7fffffff), reflect.TypeOf(int32(0x7fffffff) % int32(0x7fffffff)))
}

// Test Int32 % Int64
func TestCheckBinaryTypedExprInt32RemInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % int64(0x7fffffffffffffff)`, env,
		`invalid operation: 2147483647 % 9223372036854775807 (mismatched types int32 and int64)`,
	)

}

// Test Int32 % Uint8
func TestCheckBinaryTypedExprInt32RemUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % uint8(0xff)`, env,
		`invalid operation: 2147483647 % 255 (mismatched types int32 and uint8)`,
	)

}

// Test Int32 % Uint16
func TestCheckBinaryTypedExprInt32RemUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % uint16(0xffff)`, env,
		`invalid operation: 2147483647 % 65535 (mismatched types int32 and uint16)`,
	)

}

// Test Int32 % Uint32
func TestCheckBinaryTypedExprInt32RemUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % uint32(0xffffffff)`, env,
		`invalid operation: 2147483647 % 4294967295 (mismatched types int32 and uint32)`,
	)

}

// Test Int32 % Uint64
func TestCheckBinaryTypedExprInt32RemUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % uint64(0xffffffffffffffff)`, env,
		`invalid operation: 2147483647 % 18446744073709551615 (mismatched types int32 and uint64)`,
	)

}

// Test Int32 % Float32
func TestCheckBinaryTypedExprInt32RemFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % float32(0xffffffff)`, env,
		`invalid operation: 2147483647 % 4.29497e+09 (mismatched types int32 and float32)`,
	)

}

// Test Int32 % Float64
func TestCheckBinaryTypedExprInt32RemFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % float64(0xffffffff)`, env,
		`invalid operation: 2147483647 % 4.29497e+09 (mismatched types int32 and float64)`,
	)

}

// Test Int32 % Complex64
func TestCheckBinaryTypedExprInt32RemComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 2147483647 % (4.29497e+09+4.29497e+09i) (mismatched types int32 and complex64)`,
	)

}

// Test Int32 % Complex128
func TestCheckBinaryTypedExprInt32RemComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 2147483647 % (4.29497e+09+4.29497e+09i) (mismatched types int32 and complex128)`,
	)

}

// Test Int32 % Rune32
func TestCheckBinaryTypedExprInt32RemRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) % rune(0x7fffffff)`, env, int32(0x7fffffff) % rune(0x7fffffff), reflect.TypeOf(int32(0x7fffffff) % rune(0x7fffffff)))
}

// Test Int32 % StringT
func TestCheckBinaryTypedExprInt32RemStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % string("abc")`, env,
		`invalid operation: 2147483647 % "abc" (mismatched types int32 and string)`,
	)

}

// Test Int32 % BoolT
func TestCheckBinaryTypedExprInt32RemBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) % bool(true)`, env,
		`invalid operation: 2147483647 % true (mismatched types int32 and bool)`,
	)

}

// Test Int32 == Int
func TestCheckBinaryTypedExprInt32EqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) == 4`, env, int32(0x7fffffff) == 4, reflect.TypeOf(int32(0x7fffffff) == 4))
}

// Test Int32 == Rune
func TestCheckBinaryTypedExprInt32EqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) == '@'`, env, int32(0x7fffffff) == '@', reflect.TypeOf(int32(0x7fffffff) == '@'))
}

// Test Int32 == Float
func TestCheckBinaryTypedExprInt32EqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) == 2.0`, env, int32(0x7fffffff) == 2.0, reflect.TypeOf(int32(0x7fffffff) == 2.0))
}

// Test Int32 == Complex
func TestCheckBinaryTypedExprInt32EqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int32 == Bool
func TestCheckBinaryTypedExprInt32EqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == true`, env,
		`cannot convert true to type int32`,
		`invalid operation: 2147483647 == true (mismatched types int32 and bool)`,
	)

}

// Test Int32 == String
func TestCheckBinaryTypedExprInt32EqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == "abc"`, env,
		`cannot convert "abc" to type int32`,
		`invalid operation: 2147483647 == "abc" (mismatched types int32 and string)`,
	)

}

// Test Int32 == Nil
func TestCheckBinaryTypedExprInt32EqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == nil`, env,
		`cannot convert nil to type int32`,
	)

}

// Test Int32 == Int8
func TestCheckBinaryTypedExprInt32EqlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == int8(0x7f)`, env,
		`invalid operation: 2147483647 == 127 (mismatched types int32 and int8)`,
	)

}

// Test Int32 == Int16
func TestCheckBinaryTypedExprInt32EqlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == int16(0x7fff)`, env,
		`invalid operation: 2147483647 == 32767 (mismatched types int32 and int16)`,
	)

}

// Test Int32 == Int32
func TestCheckBinaryTypedExprInt32EqlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) == int32(0x7fffffff)`, env, int32(0x7fffffff) == int32(0x7fffffff), reflect.TypeOf(int32(0x7fffffff) == int32(0x7fffffff)))
}

// Test Int32 == Int64
func TestCheckBinaryTypedExprInt32EqlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == int64(0x7fffffffffffffff)`, env,
		`invalid operation: 2147483647 == 9223372036854775807 (mismatched types int32 and int64)`,
	)

}

// Test Int32 == Uint8
func TestCheckBinaryTypedExprInt32EqlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == uint8(0xff)`, env,
		`invalid operation: 2147483647 == 255 (mismatched types int32 and uint8)`,
	)

}

// Test Int32 == Uint16
func TestCheckBinaryTypedExprInt32EqlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == uint16(0xffff)`, env,
		`invalid operation: 2147483647 == 65535 (mismatched types int32 and uint16)`,
	)

}

// Test Int32 == Uint32
func TestCheckBinaryTypedExprInt32EqlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == uint32(0xffffffff)`, env,
		`invalid operation: 2147483647 == 4294967295 (mismatched types int32 and uint32)`,
	)

}

// Test Int32 == Uint64
func TestCheckBinaryTypedExprInt32EqlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == uint64(0xffffffffffffffff)`, env,
		`invalid operation: 2147483647 == 18446744073709551615 (mismatched types int32 and uint64)`,
	)

}

// Test Int32 == Float32
func TestCheckBinaryTypedExprInt32EqlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == float32(0xffffffff)`, env,
		`invalid operation: 2147483647 == 4.29497e+09 (mismatched types int32 and float32)`,
	)

}

// Test Int32 == Float64
func TestCheckBinaryTypedExprInt32EqlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == float64(0xffffffff)`, env,
		`invalid operation: 2147483647 == 4.29497e+09 (mismatched types int32 and float64)`,
	)

}

// Test Int32 == Complex64
func TestCheckBinaryTypedExprInt32EqlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 2147483647 == (4.29497e+09+4.29497e+09i) (mismatched types int32 and complex64)`,
	)

}

// Test Int32 == Complex128
func TestCheckBinaryTypedExprInt32EqlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 2147483647 == (4.29497e+09+4.29497e+09i) (mismatched types int32 and complex128)`,
	)

}

// Test Int32 == Rune32
func TestCheckBinaryTypedExprInt32EqlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) == rune(0x7fffffff)`, env, int32(0x7fffffff) == rune(0x7fffffff), reflect.TypeOf(int32(0x7fffffff) == rune(0x7fffffff)))
}

// Test Int32 == StringT
func TestCheckBinaryTypedExprInt32EqlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == string("abc")`, env,
		`invalid operation: 2147483647 == "abc" (mismatched types int32 and string)`,
	)

}

// Test Int32 == BoolT
func TestCheckBinaryTypedExprInt32EqlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) == bool(true)`, env,
		`invalid operation: 2147483647 == true (mismatched types int32 and bool)`,
	)

}

// Test Int32 > Int
func TestCheckBinaryTypedExprInt32GtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) > 4`, env, int32(0x7fffffff) > 4, reflect.TypeOf(int32(0x7fffffff) > 4))
}

// Test Int32 > Rune
func TestCheckBinaryTypedExprInt32GtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) > '@'`, env, int32(0x7fffffff) > '@', reflect.TypeOf(int32(0x7fffffff) > '@'))
}

// Test Int32 > Float
func TestCheckBinaryTypedExprInt32GtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) > 2.0`, env, int32(0x7fffffff) > 2.0, reflect.TypeOf(int32(0x7fffffff) > 2.0))
}

// Test Int32 > Complex
func TestCheckBinaryTypedExprInt32GtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int32 > Bool
func TestCheckBinaryTypedExprInt32GtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > true`, env,
		`cannot convert true to type int32`,
		`invalid operation: 2147483647 > true (mismatched types int32 and bool)`,
	)

}

// Test Int32 > String
func TestCheckBinaryTypedExprInt32GtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > "abc"`, env,
		`cannot convert "abc" to type int32`,
		`invalid operation: 2147483647 > "abc" (mismatched types int32 and string)`,
	)

}

// Test Int32 > Nil
func TestCheckBinaryTypedExprInt32GtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > nil`, env,
		`cannot convert nil to type int32`,
	)

}

// Test Int32 > Int8
func TestCheckBinaryTypedExprInt32GtrInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > int8(0x7f)`, env,
		`invalid operation: 2147483647 > 127 (mismatched types int32 and int8)`,
	)

}

// Test Int32 > Int16
func TestCheckBinaryTypedExprInt32GtrInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > int16(0x7fff)`, env,
		`invalid operation: 2147483647 > 32767 (mismatched types int32 and int16)`,
	)

}

// Test Int32 > Int32
func TestCheckBinaryTypedExprInt32GtrInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) > int32(0x7fffffff)`, env, int32(0x7fffffff) > int32(0x7fffffff), reflect.TypeOf(int32(0x7fffffff) > int32(0x7fffffff)))
}

// Test Int32 > Int64
func TestCheckBinaryTypedExprInt32GtrInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > int64(0x7fffffffffffffff)`, env,
		`invalid operation: 2147483647 > 9223372036854775807 (mismatched types int32 and int64)`,
	)

}

// Test Int32 > Uint8
func TestCheckBinaryTypedExprInt32GtrUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > uint8(0xff)`, env,
		`invalid operation: 2147483647 > 255 (mismatched types int32 and uint8)`,
	)

}

// Test Int32 > Uint16
func TestCheckBinaryTypedExprInt32GtrUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > uint16(0xffff)`, env,
		`invalid operation: 2147483647 > 65535 (mismatched types int32 and uint16)`,
	)

}

// Test Int32 > Uint32
func TestCheckBinaryTypedExprInt32GtrUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > uint32(0xffffffff)`, env,
		`invalid operation: 2147483647 > 4294967295 (mismatched types int32 and uint32)`,
	)

}

// Test Int32 > Uint64
func TestCheckBinaryTypedExprInt32GtrUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > uint64(0xffffffffffffffff)`, env,
		`invalid operation: 2147483647 > 18446744073709551615 (mismatched types int32 and uint64)`,
	)

}

// Test Int32 > Float32
func TestCheckBinaryTypedExprInt32GtrFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > float32(0xffffffff)`, env,
		`invalid operation: 2147483647 > 4.29497e+09 (mismatched types int32 and float32)`,
	)

}

// Test Int32 > Float64
func TestCheckBinaryTypedExprInt32GtrFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > float64(0xffffffff)`, env,
		`invalid operation: 2147483647 > 4.29497e+09 (mismatched types int32 and float64)`,
	)

}

// Test Int32 > Complex64
func TestCheckBinaryTypedExprInt32GtrComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 2147483647 > (4.29497e+09+4.29497e+09i) (mismatched types int32 and complex64)`,
	)

}

// Test Int32 > Complex128
func TestCheckBinaryTypedExprInt32GtrComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 2147483647 > (4.29497e+09+4.29497e+09i) (mismatched types int32 and complex128)`,
	)

}

// Test Int32 > Rune32
func TestCheckBinaryTypedExprInt32GtrRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff) > rune(0x7fffffff)`, env, int32(0x7fffffff) > rune(0x7fffffff), reflect.TypeOf(int32(0x7fffffff) > rune(0x7fffffff)))
}

// Test Int32 > StringT
func TestCheckBinaryTypedExprInt32GtrStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > string("abc")`, env,
		`invalid operation: 2147483647 > "abc" (mismatched types int32 and string)`,
	)

}

// Test Int32 > BoolT
func TestCheckBinaryTypedExprInt32GtrBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) > bool(true)`, env,
		`invalid operation: 2147483647 > true (mismatched types int32 and bool)`,
	)

}

// Test Int32 << Int
func TestCheckBinaryTypedExprInt32ShlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << 4`, env,
		`constant 34359738352 overflows int32`,
	)

}

// Test Int32 << Rune
func TestCheckBinaryTypedExprInt32ShlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << '@'`, env,
		`constant 39614081238685424723062423552 overflows int32`,
	)

}

// Test Int32 << Float
func TestCheckBinaryTypedExprInt32ShlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << 2.0`, env,
		`constant 8589934588 overflows int32`,
	)

}

// Test Int32 << Complex
func TestCheckBinaryTypedExprInt32ShlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int32 << Bool
func TestCheckBinaryTypedExprInt32ShlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << true`, env,
		`invalid operation: 2147483647 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Int32 << String
func TestCheckBinaryTypedExprInt32ShlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: 2147483647 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Int32 << Nil
func TestCheckBinaryTypedExprInt32ShlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Int32 << Int8
func TestCheckBinaryTypedExprInt32ShlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << int8(0x7f)`, env,
		`invalid operation: 2147483647 << 127 (shift count type int8, must be unsigned integer)`,
	)

}

// Test Int32 << Int16
func TestCheckBinaryTypedExprInt32ShlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << int16(0x7fff)`, env,
		`invalid operation: 2147483647 << 32767 (shift count type int16, must be unsigned integer)`,
	)

}

// Test Int32 << Int32
func TestCheckBinaryTypedExprInt32ShlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << int32(0x7fffffff)`, env,
		`invalid operation: 2147483647 << 2147483647 (shift count type int32, must be unsigned integer)`,
	)

}

// Test Int32 << Int64
func TestCheckBinaryTypedExprInt32ShlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << int64(0x7fffffffffffffff)`, env,
		`invalid operation: 2147483647 << 9223372036854775807 (shift count type int64, must be unsigned integer)`,
	)

}

// Test Int32 << Uint8
func TestCheckBinaryTypedExprInt32ShlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << uint8(0xff)`, env,
		`constant 124330809044550615920187464324919717520770083772701937027295712203561082249176779063296 overflows int32`,
	)

}

// Test Int32 << Uint16
func TestCheckBinaryTypedExprInt32ShlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << uint16(0xffff)`, env,
		`stupid shift: 65535`,
	)

}

// Test Int32 << Uint32
func TestCheckBinaryTypedExprInt32ShlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << uint32(0xffffffff)`, env,
		`stupid shift: 4294967295`,
	)

}

// Test Int32 << Uint64
func TestCheckBinaryTypedExprInt32ShlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << uint64(0xffffffffffffffff)`, env,
		`stupid shift: -1`,
	)

}

// Test Int32 << Float32
func TestCheckBinaryTypedExprInt32ShlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << float32(0xffffffff)`, env,
		`invalid operation: 2147483647 << 4.29497e+09 (shift count type float32, must be unsigned integer)`,
	)

}

// Test Int32 << Float64
func TestCheckBinaryTypedExprInt32ShlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << float64(0xffffffff)`, env,
		`invalid operation: 2147483647 << 4.29497e+09 (shift count type float64, must be unsigned integer)`,
	)

}

// Test Int32 << Complex64
func TestCheckBinaryTypedExprInt32ShlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 2147483647 << (4.29497e+09+4.29497e+09i) (shift count type complex64, must be unsigned integer)`,
	)

}

// Test Int32 << Complex128
func TestCheckBinaryTypedExprInt32ShlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 2147483647 << (4.29497e+09+4.29497e+09i) (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Int32 << Rune32
func TestCheckBinaryTypedExprInt32ShlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << rune(0x7fffffff)`, env,
		`invalid operation: 2147483647 << rune(2147483647) (shift count type rune, must be unsigned integer)`,
	)

}

// Test Int32 << StringT
func TestCheckBinaryTypedExprInt32ShlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << string("abc")`, env,
		`invalid operation: 2147483647 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Int32 << BoolT
func TestCheckBinaryTypedExprInt32ShlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffff) << bool(true)`, env,
		`invalid operation: 2147483647 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Int64 + Int
func TestCheckBinaryTypedExprInt64AddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + 4`, env,
		`constant 9223372036854775811 overflows int64`,
	)

}

// Test Int64 + Rune
func TestCheckBinaryTypedExprInt64AddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + '@'`, env,
		`constant 9223372036854775871 overflows int64`,
	)

}

// Test Int64 + Float
func TestCheckBinaryTypedExprInt64AddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + 2.0`, env,
		`constant 9223372036854775809 overflows int64`,
	)

}

// Test Int64 + Complex
func TestCheckBinaryTypedExprInt64AddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int64 + Bool
func TestCheckBinaryTypedExprInt64AddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + true`, env,
		`cannot convert true to type int64`,
		`invalid operation: 9223372036854775807 + true (mismatched types int64 and bool)`,
	)

}

// Test Int64 + String
func TestCheckBinaryTypedExprInt64AddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + "abc"`, env,
		`cannot convert "abc" to type int64`,
		`invalid operation: 9223372036854775807 + "abc" (mismatched types int64 and string)`,
	)

}

// Test Int64 + Nil
func TestCheckBinaryTypedExprInt64AddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + nil`, env,
		`cannot convert nil to type int64`,
	)

}

// Test Int64 + Int8
func TestCheckBinaryTypedExprInt64AddInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + int8(0x7f)`, env,
		`invalid operation: 9223372036854775807 + 127 (mismatched types int64 and int8)`,
	)

}

// Test Int64 + Int16
func TestCheckBinaryTypedExprInt64AddInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + int16(0x7fff)`, env,
		`invalid operation: 9223372036854775807 + 32767 (mismatched types int64 and int16)`,
	)

}

// Test Int64 + Int32
func TestCheckBinaryTypedExprInt64AddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + int32(0x7fffffff)`, env,
		`invalid operation: 9223372036854775807 + 2147483647 (mismatched types int64 and int32)`,
	)

}

// Test Int64 + Int64
func TestCheckBinaryTypedExprInt64AddInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + int64(0x7fffffffffffffff)`, env,
		`constant 18446744073709551614 overflows int64`,
	)

}

// Test Int64 + Uint8
func TestCheckBinaryTypedExprInt64AddUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + uint8(0xff)`, env,
		`invalid operation: 9223372036854775807 + 255 (mismatched types int64 and uint8)`,
	)

}

// Test Int64 + Uint16
func TestCheckBinaryTypedExprInt64AddUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + uint16(0xffff)`, env,
		`invalid operation: 9223372036854775807 + 65535 (mismatched types int64 and uint16)`,
	)

}

// Test Int64 + Uint32
func TestCheckBinaryTypedExprInt64AddUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + uint32(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 + 4294967295 (mismatched types int64 and uint32)`,
	)

}

// Test Int64 + Uint64
func TestCheckBinaryTypedExprInt64AddUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + uint64(0xffffffffffffffff)`, env,
		`invalid operation: 9223372036854775807 + 18446744073709551615 (mismatched types int64 and uint64)`,
	)

}

// Test Int64 + Float32
func TestCheckBinaryTypedExprInt64AddFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + float32(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 + 4.29497e+09 (mismatched types int64 and float32)`,
	)

}

// Test Int64 + Float64
func TestCheckBinaryTypedExprInt64AddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + float64(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 + 4.29497e+09 (mismatched types int64 and float64)`,
	)

}

// Test Int64 + Complex64
func TestCheckBinaryTypedExprInt64AddComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 9223372036854775807 + (4.29497e+09+4.29497e+09i) (mismatched types int64 and complex64)`,
	)

}

// Test Int64 + Complex128
func TestCheckBinaryTypedExprInt64AddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 9223372036854775807 + (4.29497e+09+4.29497e+09i) (mismatched types int64 and complex128)`,
	)

}

// Test Int64 + Rune32
func TestCheckBinaryTypedExprInt64AddRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + rune(0x7fffffff)`, env,
		`invalid operation: 9223372036854775807 + rune(2147483647) (mismatched types int64 and rune)`,
	)

}

// Test Int64 + StringT
func TestCheckBinaryTypedExprInt64AddStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + string("abc")`, env,
		`invalid operation: 9223372036854775807 + "abc" (mismatched types int64 and string)`,
	)

}

// Test Int64 + BoolT
func TestCheckBinaryTypedExprInt64AddBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) + bool(true)`, env,
		`invalid operation: 9223372036854775807 + true (mismatched types int64 and bool)`,
	)

}

// Test Int64 - Int
func TestCheckBinaryTypedExprInt64SubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) - 4`, env, int64(0x7fffffffffffffff) - 4, reflect.TypeOf(int64(0x7fffffffffffffff) - 4))
}

// Test Int64 - Rune
func TestCheckBinaryTypedExprInt64SubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) - '@'`, env, int64(0x7fffffffffffffff) - '@', reflect.TypeOf(int64(0x7fffffffffffffff) - '@'))
}

// Test Int64 - Float
func TestCheckBinaryTypedExprInt64SubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) - 2.0`, env, int64(0x7fffffffffffffff) - 2.0, reflect.TypeOf(int64(0x7fffffffffffffff) - 2.0))
}

// Test Int64 - Complex
func TestCheckBinaryTypedExprInt64SubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int64 - Bool
func TestCheckBinaryTypedExprInt64SubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - true`, env,
		`cannot convert true to type int64`,
		`invalid operation: 9223372036854775807 - true (mismatched types int64 and bool)`,
	)

}

// Test Int64 - String
func TestCheckBinaryTypedExprInt64SubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - "abc"`, env,
		`cannot convert "abc" to type int64`,
		`invalid operation: 9223372036854775807 - "abc" (mismatched types int64 and string)`,
	)

}

// Test Int64 - Nil
func TestCheckBinaryTypedExprInt64SubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - nil`, env,
		`cannot convert nil to type int64`,
	)

}

// Test Int64 - Int8
func TestCheckBinaryTypedExprInt64SubInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - int8(0x7f)`, env,
		`invalid operation: 9223372036854775807 - 127 (mismatched types int64 and int8)`,
	)

}

// Test Int64 - Int16
func TestCheckBinaryTypedExprInt64SubInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - int16(0x7fff)`, env,
		`invalid operation: 9223372036854775807 - 32767 (mismatched types int64 and int16)`,
	)

}

// Test Int64 - Int32
func TestCheckBinaryTypedExprInt64SubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - int32(0x7fffffff)`, env,
		`invalid operation: 9223372036854775807 - 2147483647 (mismatched types int64 and int32)`,
	)

}

// Test Int64 - Int64
func TestCheckBinaryTypedExprInt64SubInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) - int64(0x7fffffffffffffff)`, env, int64(0x7fffffffffffffff) - int64(0x7fffffffffffffff), reflect.TypeOf(int64(0x7fffffffffffffff) - int64(0x7fffffffffffffff)))
}

// Test Int64 - Uint8
func TestCheckBinaryTypedExprInt64SubUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - uint8(0xff)`, env,
		`invalid operation: 9223372036854775807 - 255 (mismatched types int64 and uint8)`,
	)

}

// Test Int64 - Uint16
func TestCheckBinaryTypedExprInt64SubUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - uint16(0xffff)`, env,
		`invalid operation: 9223372036854775807 - 65535 (mismatched types int64 and uint16)`,
	)

}

// Test Int64 - Uint32
func TestCheckBinaryTypedExprInt64SubUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - uint32(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 - 4294967295 (mismatched types int64 and uint32)`,
	)

}

// Test Int64 - Uint64
func TestCheckBinaryTypedExprInt64SubUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - uint64(0xffffffffffffffff)`, env,
		`invalid operation: 9223372036854775807 - 18446744073709551615 (mismatched types int64 and uint64)`,
	)

}

// Test Int64 - Float32
func TestCheckBinaryTypedExprInt64SubFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - float32(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 - 4.29497e+09 (mismatched types int64 and float32)`,
	)

}

// Test Int64 - Float64
func TestCheckBinaryTypedExprInt64SubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - float64(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 - 4.29497e+09 (mismatched types int64 and float64)`,
	)

}

// Test Int64 - Complex64
func TestCheckBinaryTypedExprInt64SubComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 9223372036854775807 - (4.29497e+09+4.29497e+09i) (mismatched types int64 and complex64)`,
	)

}

// Test Int64 - Complex128
func TestCheckBinaryTypedExprInt64SubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 9223372036854775807 - (4.29497e+09+4.29497e+09i) (mismatched types int64 and complex128)`,
	)

}

// Test Int64 - Rune32
func TestCheckBinaryTypedExprInt64SubRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - rune(0x7fffffff)`, env,
		`invalid operation: 9223372036854775807 - rune(2147483647) (mismatched types int64 and rune)`,
	)

}

// Test Int64 - StringT
func TestCheckBinaryTypedExprInt64SubStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - string("abc")`, env,
		`invalid operation: 9223372036854775807 - "abc" (mismatched types int64 and string)`,
	)

}

// Test Int64 - BoolT
func TestCheckBinaryTypedExprInt64SubBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) - bool(true)`, env,
		`invalid operation: 9223372036854775807 - true (mismatched types int64 and bool)`,
	)

}

// Test Int64 & Int
func TestCheckBinaryTypedExprInt64AndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) & 4`, env, int64(0x7fffffffffffffff) & 4, reflect.TypeOf(int64(0x7fffffffffffffff) & 4))
}

// Test Int64 & Rune
func TestCheckBinaryTypedExprInt64AndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) & '@'`, env, int64(0x7fffffffffffffff) & '@', reflect.TypeOf(int64(0x7fffffffffffffff) & '@'))
}

// Test Int64 & Float
func TestCheckBinaryTypedExprInt64AndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) & 2.0`, env, int64(0x7fffffffffffffff) & 2.0, reflect.TypeOf(int64(0x7fffffffffffffff) & 2.0))
}

// Test Int64 & Complex
func TestCheckBinaryTypedExprInt64AndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int64 & Bool
func TestCheckBinaryTypedExprInt64AndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & true`, env,
		`cannot convert true to type int64`,
		`invalid operation: 9223372036854775807 & true (mismatched types int64 and bool)`,
	)

}

// Test Int64 & String
func TestCheckBinaryTypedExprInt64AndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & "abc"`, env,
		`cannot convert "abc" to type int64`,
		`invalid operation: 9223372036854775807 & "abc" (mismatched types int64 and string)`,
	)

}

// Test Int64 & Nil
func TestCheckBinaryTypedExprInt64AndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & nil`, env,
		`cannot convert nil to type int64`,
	)

}

// Test Int64 & Int8
func TestCheckBinaryTypedExprInt64AndInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & int8(0x7f)`, env,
		`invalid operation: 9223372036854775807 & 127 (mismatched types int64 and int8)`,
	)

}

// Test Int64 & Int16
func TestCheckBinaryTypedExprInt64AndInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & int16(0x7fff)`, env,
		`invalid operation: 9223372036854775807 & 32767 (mismatched types int64 and int16)`,
	)

}

// Test Int64 & Int32
func TestCheckBinaryTypedExprInt64AndInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & int32(0x7fffffff)`, env,
		`invalid operation: 9223372036854775807 & 2147483647 (mismatched types int64 and int32)`,
	)

}

// Test Int64 & Int64
func TestCheckBinaryTypedExprInt64AndInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) & int64(0x7fffffffffffffff)`, env, int64(0x7fffffffffffffff) & int64(0x7fffffffffffffff), reflect.TypeOf(int64(0x7fffffffffffffff) & int64(0x7fffffffffffffff)))
}

// Test Int64 & Uint8
func TestCheckBinaryTypedExprInt64AndUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & uint8(0xff)`, env,
		`invalid operation: 9223372036854775807 & 255 (mismatched types int64 and uint8)`,
	)

}

// Test Int64 & Uint16
func TestCheckBinaryTypedExprInt64AndUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & uint16(0xffff)`, env,
		`invalid operation: 9223372036854775807 & 65535 (mismatched types int64 and uint16)`,
	)

}

// Test Int64 & Uint32
func TestCheckBinaryTypedExprInt64AndUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & uint32(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 & 4294967295 (mismatched types int64 and uint32)`,
	)

}

// Test Int64 & Uint64
func TestCheckBinaryTypedExprInt64AndUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & uint64(0xffffffffffffffff)`, env,
		`invalid operation: 9223372036854775807 & 18446744073709551615 (mismatched types int64 and uint64)`,
	)

}

// Test Int64 & Float32
func TestCheckBinaryTypedExprInt64AndFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & float32(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 & 4.29497e+09 (mismatched types int64 and float32)`,
	)

}

// Test Int64 & Float64
func TestCheckBinaryTypedExprInt64AndFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & float64(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 & 4.29497e+09 (mismatched types int64 and float64)`,
	)

}

// Test Int64 & Complex64
func TestCheckBinaryTypedExprInt64AndComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 9223372036854775807 & (4.29497e+09+4.29497e+09i) (mismatched types int64 and complex64)`,
	)

}

// Test Int64 & Complex128
func TestCheckBinaryTypedExprInt64AndComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 9223372036854775807 & (4.29497e+09+4.29497e+09i) (mismatched types int64 and complex128)`,
	)

}

// Test Int64 & Rune32
func TestCheckBinaryTypedExprInt64AndRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & rune(0x7fffffff)`, env,
		`invalid operation: 9223372036854775807 & rune(2147483647) (mismatched types int64 and rune)`,
	)

}

// Test Int64 & StringT
func TestCheckBinaryTypedExprInt64AndStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & string("abc")`, env,
		`invalid operation: 9223372036854775807 & "abc" (mismatched types int64 and string)`,
	)

}

// Test Int64 & BoolT
func TestCheckBinaryTypedExprInt64AndBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) & bool(true)`, env,
		`invalid operation: 9223372036854775807 & true (mismatched types int64 and bool)`,
	)

}

// Test Int64 % Int
func TestCheckBinaryTypedExprInt64RemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) % 4`, env, int64(0x7fffffffffffffff) % 4, reflect.TypeOf(int64(0x7fffffffffffffff) % 4))
}

// Test Int64 % Rune
func TestCheckBinaryTypedExprInt64RemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) % '@'`, env, int64(0x7fffffffffffffff) % '@', reflect.TypeOf(int64(0x7fffffffffffffff) % '@'))
}

// Test Int64 % Float
func TestCheckBinaryTypedExprInt64RemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) % 2.0`, env, int64(0x7fffffffffffffff) % 2.0, reflect.TypeOf(int64(0x7fffffffffffffff) % 2.0))
}

// Test Int64 % Complex
func TestCheckBinaryTypedExprInt64RemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % 8.0i`, env,
		`constant 0+8i truncated to real`,
		`division by zero`,
	)

}

// Test Int64 % Bool
func TestCheckBinaryTypedExprInt64RemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % true`, env,
		`cannot convert true to type int64`,
		`invalid operation: 9223372036854775807 % true (mismatched types int64 and bool)`,
	)

}

// Test Int64 % String
func TestCheckBinaryTypedExprInt64RemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % "abc"`, env,
		`cannot convert "abc" to type int64`,
		`invalid operation: 9223372036854775807 % "abc" (mismatched types int64 and string)`,
	)

}

// Test Int64 % Nil
func TestCheckBinaryTypedExprInt64RemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % nil`, env,
		`cannot convert nil to type int64`,
	)

}

// Test Int64 % Int8
func TestCheckBinaryTypedExprInt64RemInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % int8(0x7f)`, env,
		`invalid operation: 9223372036854775807 % 127 (mismatched types int64 and int8)`,
	)

}

// Test Int64 % Int16
func TestCheckBinaryTypedExprInt64RemInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % int16(0x7fff)`, env,
		`invalid operation: 9223372036854775807 % 32767 (mismatched types int64 and int16)`,
	)

}

// Test Int64 % Int32
func TestCheckBinaryTypedExprInt64RemInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % int32(0x7fffffff)`, env,
		`invalid operation: 9223372036854775807 % 2147483647 (mismatched types int64 and int32)`,
	)

}

// Test Int64 % Int64
func TestCheckBinaryTypedExprInt64RemInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) % int64(0x7fffffffffffffff)`, env, int64(0x7fffffffffffffff) % int64(0x7fffffffffffffff), reflect.TypeOf(int64(0x7fffffffffffffff) % int64(0x7fffffffffffffff)))
}

// Test Int64 % Uint8
func TestCheckBinaryTypedExprInt64RemUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % uint8(0xff)`, env,
		`invalid operation: 9223372036854775807 % 255 (mismatched types int64 and uint8)`,
	)

}

// Test Int64 % Uint16
func TestCheckBinaryTypedExprInt64RemUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % uint16(0xffff)`, env,
		`invalid operation: 9223372036854775807 % 65535 (mismatched types int64 and uint16)`,
	)

}

// Test Int64 % Uint32
func TestCheckBinaryTypedExprInt64RemUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % uint32(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 % 4294967295 (mismatched types int64 and uint32)`,
	)

}

// Test Int64 % Uint64
func TestCheckBinaryTypedExprInt64RemUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % uint64(0xffffffffffffffff)`, env,
		`invalid operation: 9223372036854775807 % 18446744073709551615 (mismatched types int64 and uint64)`,
	)

}

// Test Int64 % Float32
func TestCheckBinaryTypedExprInt64RemFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % float32(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 % 4.29497e+09 (mismatched types int64 and float32)`,
	)

}

// Test Int64 % Float64
func TestCheckBinaryTypedExprInt64RemFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % float64(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 % 4.29497e+09 (mismatched types int64 and float64)`,
	)

}

// Test Int64 % Complex64
func TestCheckBinaryTypedExprInt64RemComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 9223372036854775807 % (4.29497e+09+4.29497e+09i) (mismatched types int64 and complex64)`,
	)

}

// Test Int64 % Complex128
func TestCheckBinaryTypedExprInt64RemComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 9223372036854775807 % (4.29497e+09+4.29497e+09i) (mismatched types int64 and complex128)`,
	)

}

// Test Int64 % Rune32
func TestCheckBinaryTypedExprInt64RemRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % rune(0x7fffffff)`, env,
		`invalid operation: 9223372036854775807 % rune(2147483647) (mismatched types int64 and rune)`,
	)

}

// Test Int64 % StringT
func TestCheckBinaryTypedExprInt64RemStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % string("abc")`, env,
		`invalid operation: 9223372036854775807 % "abc" (mismatched types int64 and string)`,
	)

}

// Test Int64 % BoolT
func TestCheckBinaryTypedExprInt64RemBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) % bool(true)`, env,
		`invalid operation: 9223372036854775807 % true (mismatched types int64 and bool)`,
	)

}

// Test Int64 == Int
func TestCheckBinaryTypedExprInt64EqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) == 4`, env, int64(0x7fffffffffffffff) == 4, reflect.TypeOf(int64(0x7fffffffffffffff) == 4))
}

// Test Int64 == Rune
func TestCheckBinaryTypedExprInt64EqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) == '@'`, env, int64(0x7fffffffffffffff) == '@', reflect.TypeOf(int64(0x7fffffffffffffff) == '@'))
}

// Test Int64 == Float
func TestCheckBinaryTypedExprInt64EqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) == 2.0`, env, int64(0x7fffffffffffffff) == 2.0, reflect.TypeOf(int64(0x7fffffffffffffff) == 2.0))
}

// Test Int64 == Complex
func TestCheckBinaryTypedExprInt64EqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int64 == Bool
func TestCheckBinaryTypedExprInt64EqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == true`, env,
		`cannot convert true to type int64`,
		`invalid operation: 9223372036854775807 == true (mismatched types int64 and bool)`,
	)

}

// Test Int64 == String
func TestCheckBinaryTypedExprInt64EqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == "abc"`, env,
		`cannot convert "abc" to type int64`,
		`invalid operation: 9223372036854775807 == "abc" (mismatched types int64 and string)`,
	)

}

// Test Int64 == Nil
func TestCheckBinaryTypedExprInt64EqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == nil`, env,
		`cannot convert nil to type int64`,
	)

}

// Test Int64 == Int8
func TestCheckBinaryTypedExprInt64EqlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == int8(0x7f)`, env,
		`invalid operation: 9223372036854775807 == 127 (mismatched types int64 and int8)`,
	)

}

// Test Int64 == Int16
func TestCheckBinaryTypedExprInt64EqlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == int16(0x7fff)`, env,
		`invalid operation: 9223372036854775807 == 32767 (mismatched types int64 and int16)`,
	)

}

// Test Int64 == Int32
func TestCheckBinaryTypedExprInt64EqlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == int32(0x7fffffff)`, env,
		`invalid operation: 9223372036854775807 == 2147483647 (mismatched types int64 and int32)`,
	)

}

// Test Int64 == Int64
func TestCheckBinaryTypedExprInt64EqlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) == int64(0x7fffffffffffffff)`, env, int64(0x7fffffffffffffff) == int64(0x7fffffffffffffff), reflect.TypeOf(int64(0x7fffffffffffffff) == int64(0x7fffffffffffffff)))
}

// Test Int64 == Uint8
func TestCheckBinaryTypedExprInt64EqlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == uint8(0xff)`, env,
		`invalid operation: 9223372036854775807 == 255 (mismatched types int64 and uint8)`,
	)

}

// Test Int64 == Uint16
func TestCheckBinaryTypedExprInt64EqlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == uint16(0xffff)`, env,
		`invalid operation: 9223372036854775807 == 65535 (mismatched types int64 and uint16)`,
	)

}

// Test Int64 == Uint32
func TestCheckBinaryTypedExprInt64EqlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == uint32(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 == 4294967295 (mismatched types int64 and uint32)`,
	)

}

// Test Int64 == Uint64
func TestCheckBinaryTypedExprInt64EqlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == uint64(0xffffffffffffffff)`, env,
		`invalid operation: 9223372036854775807 == 18446744073709551615 (mismatched types int64 and uint64)`,
	)

}

// Test Int64 == Float32
func TestCheckBinaryTypedExprInt64EqlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == float32(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 == 4.29497e+09 (mismatched types int64 and float32)`,
	)

}

// Test Int64 == Float64
func TestCheckBinaryTypedExprInt64EqlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == float64(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 == 4.29497e+09 (mismatched types int64 and float64)`,
	)

}

// Test Int64 == Complex64
func TestCheckBinaryTypedExprInt64EqlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 9223372036854775807 == (4.29497e+09+4.29497e+09i) (mismatched types int64 and complex64)`,
	)

}

// Test Int64 == Complex128
func TestCheckBinaryTypedExprInt64EqlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 9223372036854775807 == (4.29497e+09+4.29497e+09i) (mismatched types int64 and complex128)`,
	)

}

// Test Int64 == Rune32
func TestCheckBinaryTypedExprInt64EqlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == rune(0x7fffffff)`, env,
		`invalid operation: 9223372036854775807 == rune(2147483647) (mismatched types int64 and rune)`,
	)

}

// Test Int64 == StringT
func TestCheckBinaryTypedExprInt64EqlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == string("abc")`, env,
		`invalid operation: 9223372036854775807 == "abc" (mismatched types int64 and string)`,
	)

}

// Test Int64 == BoolT
func TestCheckBinaryTypedExprInt64EqlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) == bool(true)`, env,
		`invalid operation: 9223372036854775807 == true (mismatched types int64 and bool)`,
	)

}

// Test Int64 > Int
func TestCheckBinaryTypedExprInt64GtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) > 4`, env, int64(0x7fffffffffffffff) > 4, reflect.TypeOf(int64(0x7fffffffffffffff) > 4))
}

// Test Int64 > Rune
func TestCheckBinaryTypedExprInt64GtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) > '@'`, env, int64(0x7fffffffffffffff) > '@', reflect.TypeOf(int64(0x7fffffffffffffff) > '@'))
}

// Test Int64 > Float
func TestCheckBinaryTypedExprInt64GtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) > 2.0`, env, int64(0x7fffffffffffffff) > 2.0, reflect.TypeOf(int64(0x7fffffffffffffff) > 2.0))
}

// Test Int64 > Complex
func TestCheckBinaryTypedExprInt64GtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int64 > Bool
func TestCheckBinaryTypedExprInt64GtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > true`, env,
		`cannot convert true to type int64`,
		`invalid operation: 9223372036854775807 > true (mismatched types int64 and bool)`,
	)

}

// Test Int64 > String
func TestCheckBinaryTypedExprInt64GtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > "abc"`, env,
		`cannot convert "abc" to type int64`,
		`invalid operation: 9223372036854775807 > "abc" (mismatched types int64 and string)`,
	)

}

// Test Int64 > Nil
func TestCheckBinaryTypedExprInt64GtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > nil`, env,
		`cannot convert nil to type int64`,
	)

}

// Test Int64 > Int8
func TestCheckBinaryTypedExprInt64GtrInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > int8(0x7f)`, env,
		`invalid operation: 9223372036854775807 > 127 (mismatched types int64 and int8)`,
	)

}

// Test Int64 > Int16
func TestCheckBinaryTypedExprInt64GtrInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > int16(0x7fff)`, env,
		`invalid operation: 9223372036854775807 > 32767 (mismatched types int64 and int16)`,
	)

}

// Test Int64 > Int32
func TestCheckBinaryTypedExprInt64GtrInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > int32(0x7fffffff)`, env,
		`invalid operation: 9223372036854775807 > 2147483647 (mismatched types int64 and int32)`,
	)

}

// Test Int64 > Int64
func TestCheckBinaryTypedExprInt64GtrInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff) > int64(0x7fffffffffffffff)`, env, int64(0x7fffffffffffffff) > int64(0x7fffffffffffffff), reflect.TypeOf(int64(0x7fffffffffffffff) > int64(0x7fffffffffffffff)))
}

// Test Int64 > Uint8
func TestCheckBinaryTypedExprInt64GtrUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > uint8(0xff)`, env,
		`invalid operation: 9223372036854775807 > 255 (mismatched types int64 and uint8)`,
	)

}

// Test Int64 > Uint16
func TestCheckBinaryTypedExprInt64GtrUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > uint16(0xffff)`, env,
		`invalid operation: 9223372036854775807 > 65535 (mismatched types int64 and uint16)`,
	)

}

// Test Int64 > Uint32
func TestCheckBinaryTypedExprInt64GtrUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > uint32(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 > 4294967295 (mismatched types int64 and uint32)`,
	)

}

// Test Int64 > Uint64
func TestCheckBinaryTypedExprInt64GtrUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > uint64(0xffffffffffffffff)`, env,
		`invalid operation: 9223372036854775807 > 18446744073709551615 (mismatched types int64 and uint64)`,
	)

}

// Test Int64 > Float32
func TestCheckBinaryTypedExprInt64GtrFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > float32(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 > 4.29497e+09 (mismatched types int64 and float32)`,
	)

}

// Test Int64 > Float64
func TestCheckBinaryTypedExprInt64GtrFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > float64(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 > 4.29497e+09 (mismatched types int64 and float64)`,
	)

}

// Test Int64 > Complex64
func TestCheckBinaryTypedExprInt64GtrComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 9223372036854775807 > (4.29497e+09+4.29497e+09i) (mismatched types int64 and complex64)`,
	)

}

// Test Int64 > Complex128
func TestCheckBinaryTypedExprInt64GtrComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 9223372036854775807 > (4.29497e+09+4.29497e+09i) (mismatched types int64 and complex128)`,
	)

}

// Test Int64 > Rune32
func TestCheckBinaryTypedExprInt64GtrRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > rune(0x7fffffff)`, env,
		`invalid operation: 9223372036854775807 > rune(2147483647) (mismatched types int64 and rune)`,
	)

}

// Test Int64 > StringT
func TestCheckBinaryTypedExprInt64GtrStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > string("abc")`, env,
		`invalid operation: 9223372036854775807 > "abc" (mismatched types int64 and string)`,
	)

}

// Test Int64 > BoolT
func TestCheckBinaryTypedExprInt64GtrBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) > bool(true)`, env,
		`invalid operation: 9223372036854775807 > true (mismatched types int64 and bool)`,
	)

}

// Test Int64 << Int
func TestCheckBinaryTypedExprInt64ShlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << 4`, env,
		`constant 147573952589676412912 overflows int64`,
	)

}

// Test Int64 << Rune
func TestCheckBinaryTypedExprInt64ShlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << '@'`, env,
		`constant 170141183460469231713240559642174554112 overflows int64`,
	)

}

// Test Int64 << Float
func TestCheckBinaryTypedExprInt64ShlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << 2.0`, env,
		`constant 36893488147419103228 overflows int64`,
	)

}

// Test Int64 << Complex
func TestCheckBinaryTypedExprInt64ShlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int64 << Bool
func TestCheckBinaryTypedExprInt64ShlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << true`, env,
		`invalid operation: 9223372036854775807 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Int64 << String
func TestCheckBinaryTypedExprInt64ShlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: 9223372036854775807 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Int64 << Nil
func TestCheckBinaryTypedExprInt64ShlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Int64 << Int8
func TestCheckBinaryTypedExprInt64ShlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << int8(0x7f)`, env,
		`invalid operation: 9223372036854775807 << 127 (shift count type int8, must be unsigned integer)`,
	)

}

// Test Int64 << Int16
func TestCheckBinaryTypedExprInt64ShlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << int16(0x7fff)`, env,
		`invalid operation: 9223372036854775807 << 32767 (shift count type int16, must be unsigned integer)`,
	)

}

// Test Int64 << Int32
func TestCheckBinaryTypedExprInt64ShlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << int32(0x7fffffff)`, env,
		`invalid operation: 9223372036854775807 << 2147483647 (shift count type int32, must be unsigned integer)`,
	)

}

// Test Int64 << Int64
func TestCheckBinaryTypedExprInt64ShlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << int64(0x7fffffffffffffff)`, env,
		`invalid operation: 9223372036854775807 << 9223372036854775807 (shift count type int64, must be unsigned integer)`,
	)

}

// Test Int64 << Uint8
func TestCheckBinaryTypedExprInt64ShlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << uint8(0xff)`, env,
		`constant 533996758980227520540859381923729930938890638084819238560125409622734649425431345501783956914176 overflows int64`,
	)

}

// Test Int64 << Uint16
func TestCheckBinaryTypedExprInt64ShlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << uint16(0xffff)`, env,
		`stupid shift: 65535`,
	)

}

// Test Int64 << Uint32
func TestCheckBinaryTypedExprInt64ShlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << uint32(0xffffffff)`, env,
		`stupid shift: 4294967295`,
	)

}

// Test Int64 << Uint64
func TestCheckBinaryTypedExprInt64ShlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << uint64(0xffffffffffffffff)`, env,
		`stupid shift: -1`,
	)

}

// Test Int64 << Float32
func TestCheckBinaryTypedExprInt64ShlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << float32(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 << 4.29497e+09 (shift count type float32, must be unsigned integer)`,
	)

}

// Test Int64 << Float64
func TestCheckBinaryTypedExprInt64ShlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << float64(0xffffffff)`, env,
		`invalid operation: 9223372036854775807 << 4.29497e+09 (shift count type float64, must be unsigned integer)`,
	)

}

// Test Int64 << Complex64
func TestCheckBinaryTypedExprInt64ShlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 9223372036854775807 << (4.29497e+09+4.29497e+09i) (shift count type complex64, must be unsigned integer)`,
	)

}

// Test Int64 << Complex128
func TestCheckBinaryTypedExprInt64ShlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 9223372036854775807 << (4.29497e+09+4.29497e+09i) (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Int64 << Rune32
func TestCheckBinaryTypedExprInt64ShlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << rune(0x7fffffff)`, env,
		`invalid operation: 9223372036854775807 << rune(2147483647) (shift count type rune, must be unsigned integer)`,
	)

}

// Test Int64 << StringT
func TestCheckBinaryTypedExprInt64ShlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << string("abc")`, env,
		`invalid operation: 9223372036854775807 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Int64 << BoolT
func TestCheckBinaryTypedExprInt64ShlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0x7fffffffffffffff) << bool(true)`, env,
		`invalid operation: 9223372036854775807 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Uint8 + Int
func TestCheckBinaryTypedExprUint8AddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + 4`, env,
		`constant 259 overflows uint8`,
	)

}

// Test Uint8 + Rune
func TestCheckBinaryTypedExprUint8AddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + '@'`, env,
		`constant 319 overflows uint8`,
	)

}

// Test Uint8 + Float
func TestCheckBinaryTypedExprUint8AddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + 2.0`, env,
		`constant 257 overflows uint8`,
	)

}

// Test Uint8 + Complex
func TestCheckBinaryTypedExprUint8AddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint8 + Bool
func TestCheckBinaryTypedExprUint8AddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + true`, env,
		`cannot convert true to type uint8`,
		`invalid operation: 255 + true (mismatched types uint8 and bool)`,
	)

}

// Test Uint8 + String
func TestCheckBinaryTypedExprUint8AddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + "abc"`, env,
		`cannot convert "abc" to type uint8`,
		`invalid operation: 255 + "abc" (mismatched types uint8 and string)`,
	)

}

// Test Uint8 + Nil
func TestCheckBinaryTypedExprUint8AddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + nil`, env,
		`cannot convert nil to type uint8`,
	)

}

// Test Uint8 + Int8
func TestCheckBinaryTypedExprUint8AddInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + int8(0x7f)`, env,
		`invalid operation: 255 + 127 (mismatched types uint8 and int8)`,
	)

}

// Test Uint8 + Int16
func TestCheckBinaryTypedExprUint8AddInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + int16(0x7fff)`, env,
		`invalid operation: 255 + 32767 (mismatched types uint8 and int16)`,
	)

}

// Test Uint8 + Int32
func TestCheckBinaryTypedExprUint8AddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + int32(0x7fffffff)`, env,
		`invalid operation: 255 + 2147483647 (mismatched types uint8 and int32)`,
	)

}

// Test Uint8 + Int64
func TestCheckBinaryTypedExprUint8AddInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + int64(0x7fffffffffffffff)`, env,
		`invalid operation: 255 + 9223372036854775807 (mismatched types uint8 and int64)`,
	)

}

// Test Uint8 + Uint8
func TestCheckBinaryTypedExprUint8AddUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + uint8(0xff)`, env,
		`constant 510 overflows uint8`,
	)

}

// Test Uint8 + Uint16
func TestCheckBinaryTypedExprUint8AddUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + uint16(0xffff)`, env,
		`invalid operation: 255 + 65535 (mismatched types uint8 and uint16)`,
	)

}

// Test Uint8 + Uint32
func TestCheckBinaryTypedExprUint8AddUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + uint32(0xffffffff)`, env,
		`invalid operation: 255 + 4294967295 (mismatched types uint8 and uint32)`,
	)

}

// Test Uint8 + Uint64
func TestCheckBinaryTypedExprUint8AddUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + uint64(0xffffffffffffffff)`, env,
		`invalid operation: 255 + 18446744073709551615 (mismatched types uint8 and uint64)`,
	)

}

// Test Uint8 + Float32
func TestCheckBinaryTypedExprUint8AddFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + float32(0xffffffff)`, env,
		`invalid operation: 255 + 4.29497e+09 (mismatched types uint8 and float32)`,
	)

}

// Test Uint8 + Float64
func TestCheckBinaryTypedExprUint8AddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + float64(0xffffffff)`, env,
		`invalid operation: 255 + 4.29497e+09 (mismatched types uint8 and float64)`,
	)

}

// Test Uint8 + Complex64
func TestCheckBinaryTypedExprUint8AddComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 255 + (4.29497e+09+4.29497e+09i) (mismatched types uint8 and complex64)`,
	)

}

// Test Uint8 + Complex128
func TestCheckBinaryTypedExprUint8AddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 255 + (4.29497e+09+4.29497e+09i) (mismatched types uint8 and complex128)`,
	)

}

// Test Uint8 + Rune32
func TestCheckBinaryTypedExprUint8AddRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + rune(0x7fffffff)`, env,
		`invalid operation: 255 + rune(2147483647) (mismatched types uint8 and rune)`,
	)

}

// Test Uint8 + StringT
func TestCheckBinaryTypedExprUint8AddStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + string("abc")`, env,
		`invalid operation: 255 + "abc" (mismatched types uint8 and string)`,
	)

}

// Test Uint8 + BoolT
func TestCheckBinaryTypedExprUint8AddBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) + bool(true)`, env,
		`invalid operation: 255 + true (mismatched types uint8 and bool)`,
	)

}

// Test Uint8 - Int
func TestCheckBinaryTypedExprUint8SubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) - 4`, env, uint8(0xff) - 4, reflect.TypeOf(uint8(0xff) - 4))
}

// Test Uint8 - Rune
func TestCheckBinaryTypedExprUint8SubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) - '@'`, env, uint8(0xff) - '@', reflect.TypeOf(uint8(0xff) - '@'))
}

// Test Uint8 - Float
func TestCheckBinaryTypedExprUint8SubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) - 2.0`, env, uint8(0xff) - 2.0, reflect.TypeOf(uint8(0xff) - 2.0))
}

// Test Uint8 - Complex
func TestCheckBinaryTypedExprUint8SubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint8 - Bool
func TestCheckBinaryTypedExprUint8SubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - true`, env,
		`cannot convert true to type uint8`,
		`invalid operation: 255 - true (mismatched types uint8 and bool)`,
	)

}

// Test Uint8 - String
func TestCheckBinaryTypedExprUint8SubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - "abc"`, env,
		`cannot convert "abc" to type uint8`,
		`invalid operation: 255 - "abc" (mismatched types uint8 and string)`,
	)

}

// Test Uint8 - Nil
func TestCheckBinaryTypedExprUint8SubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - nil`, env,
		`cannot convert nil to type uint8`,
	)

}

// Test Uint8 - Int8
func TestCheckBinaryTypedExprUint8SubInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - int8(0x7f)`, env,
		`invalid operation: 255 - 127 (mismatched types uint8 and int8)`,
	)

}

// Test Uint8 - Int16
func TestCheckBinaryTypedExprUint8SubInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - int16(0x7fff)`, env,
		`invalid operation: 255 - 32767 (mismatched types uint8 and int16)`,
	)

}

// Test Uint8 - Int32
func TestCheckBinaryTypedExprUint8SubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - int32(0x7fffffff)`, env,
		`invalid operation: 255 - 2147483647 (mismatched types uint8 and int32)`,
	)

}

// Test Uint8 - Int64
func TestCheckBinaryTypedExprUint8SubInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - int64(0x7fffffffffffffff)`, env,
		`invalid operation: 255 - 9223372036854775807 (mismatched types uint8 and int64)`,
	)

}

// Test Uint8 - Uint8
func TestCheckBinaryTypedExprUint8SubUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) - uint8(0xff)`, env, uint8(0xff) - uint8(0xff), reflect.TypeOf(uint8(0xff) - uint8(0xff)))
}

// Test Uint8 - Uint16
func TestCheckBinaryTypedExprUint8SubUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - uint16(0xffff)`, env,
		`invalid operation: 255 - 65535 (mismatched types uint8 and uint16)`,
	)

}

// Test Uint8 - Uint32
func TestCheckBinaryTypedExprUint8SubUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - uint32(0xffffffff)`, env,
		`invalid operation: 255 - 4294967295 (mismatched types uint8 and uint32)`,
	)

}

// Test Uint8 - Uint64
func TestCheckBinaryTypedExprUint8SubUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - uint64(0xffffffffffffffff)`, env,
		`invalid operation: 255 - 18446744073709551615 (mismatched types uint8 and uint64)`,
	)

}

// Test Uint8 - Float32
func TestCheckBinaryTypedExprUint8SubFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - float32(0xffffffff)`, env,
		`invalid operation: 255 - 4.29497e+09 (mismatched types uint8 and float32)`,
	)

}

// Test Uint8 - Float64
func TestCheckBinaryTypedExprUint8SubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - float64(0xffffffff)`, env,
		`invalid operation: 255 - 4.29497e+09 (mismatched types uint8 and float64)`,
	)

}

// Test Uint8 - Complex64
func TestCheckBinaryTypedExprUint8SubComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 255 - (4.29497e+09+4.29497e+09i) (mismatched types uint8 and complex64)`,
	)

}

// Test Uint8 - Complex128
func TestCheckBinaryTypedExprUint8SubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 255 - (4.29497e+09+4.29497e+09i) (mismatched types uint8 and complex128)`,
	)

}

// Test Uint8 - Rune32
func TestCheckBinaryTypedExprUint8SubRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - rune(0x7fffffff)`, env,
		`invalid operation: 255 - rune(2147483647) (mismatched types uint8 and rune)`,
	)

}

// Test Uint8 - StringT
func TestCheckBinaryTypedExprUint8SubStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - string("abc")`, env,
		`invalid operation: 255 - "abc" (mismatched types uint8 and string)`,
	)

}

// Test Uint8 - BoolT
func TestCheckBinaryTypedExprUint8SubBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) - bool(true)`, env,
		`invalid operation: 255 - true (mismatched types uint8 and bool)`,
	)

}

// Test Uint8 & Int
func TestCheckBinaryTypedExprUint8AndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) & 4`, env, uint8(0xff) & 4, reflect.TypeOf(uint8(0xff) & 4))
}

// Test Uint8 & Rune
func TestCheckBinaryTypedExprUint8AndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) & '@'`, env, uint8(0xff) & '@', reflect.TypeOf(uint8(0xff) & '@'))
}

// Test Uint8 & Float
func TestCheckBinaryTypedExprUint8AndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) & 2.0`, env, uint8(0xff) & 2.0, reflect.TypeOf(uint8(0xff) & 2.0))
}

// Test Uint8 & Complex
func TestCheckBinaryTypedExprUint8AndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint8 & Bool
func TestCheckBinaryTypedExprUint8AndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & true`, env,
		`cannot convert true to type uint8`,
		`invalid operation: 255 & true (mismatched types uint8 and bool)`,
	)

}

// Test Uint8 & String
func TestCheckBinaryTypedExprUint8AndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & "abc"`, env,
		`cannot convert "abc" to type uint8`,
		`invalid operation: 255 & "abc" (mismatched types uint8 and string)`,
	)

}

// Test Uint8 & Nil
func TestCheckBinaryTypedExprUint8AndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & nil`, env,
		`cannot convert nil to type uint8`,
	)

}

// Test Uint8 & Int8
func TestCheckBinaryTypedExprUint8AndInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & int8(0x7f)`, env,
		`invalid operation: 255 & 127 (mismatched types uint8 and int8)`,
	)

}

// Test Uint8 & Int16
func TestCheckBinaryTypedExprUint8AndInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & int16(0x7fff)`, env,
		`invalid operation: 255 & 32767 (mismatched types uint8 and int16)`,
	)

}

// Test Uint8 & Int32
func TestCheckBinaryTypedExprUint8AndInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & int32(0x7fffffff)`, env,
		`invalid operation: 255 & 2147483647 (mismatched types uint8 and int32)`,
	)

}

// Test Uint8 & Int64
func TestCheckBinaryTypedExprUint8AndInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & int64(0x7fffffffffffffff)`, env,
		`invalid operation: 255 & 9223372036854775807 (mismatched types uint8 and int64)`,
	)

}

// Test Uint8 & Uint8
func TestCheckBinaryTypedExprUint8AndUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) & uint8(0xff)`, env, uint8(0xff) & uint8(0xff), reflect.TypeOf(uint8(0xff) & uint8(0xff)))
}

// Test Uint8 & Uint16
func TestCheckBinaryTypedExprUint8AndUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & uint16(0xffff)`, env,
		`invalid operation: 255 & 65535 (mismatched types uint8 and uint16)`,
	)

}

// Test Uint8 & Uint32
func TestCheckBinaryTypedExprUint8AndUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & uint32(0xffffffff)`, env,
		`invalid operation: 255 & 4294967295 (mismatched types uint8 and uint32)`,
	)

}

// Test Uint8 & Uint64
func TestCheckBinaryTypedExprUint8AndUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & uint64(0xffffffffffffffff)`, env,
		`invalid operation: 255 & 18446744073709551615 (mismatched types uint8 and uint64)`,
	)

}

// Test Uint8 & Float32
func TestCheckBinaryTypedExprUint8AndFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & float32(0xffffffff)`, env,
		`invalid operation: 255 & 4.29497e+09 (mismatched types uint8 and float32)`,
	)

}

// Test Uint8 & Float64
func TestCheckBinaryTypedExprUint8AndFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & float64(0xffffffff)`, env,
		`invalid operation: 255 & 4.29497e+09 (mismatched types uint8 and float64)`,
	)

}

// Test Uint8 & Complex64
func TestCheckBinaryTypedExprUint8AndComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 255 & (4.29497e+09+4.29497e+09i) (mismatched types uint8 and complex64)`,
	)

}

// Test Uint8 & Complex128
func TestCheckBinaryTypedExprUint8AndComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 255 & (4.29497e+09+4.29497e+09i) (mismatched types uint8 and complex128)`,
	)

}

// Test Uint8 & Rune32
func TestCheckBinaryTypedExprUint8AndRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & rune(0x7fffffff)`, env,
		`invalid operation: 255 & rune(2147483647) (mismatched types uint8 and rune)`,
	)

}

// Test Uint8 & StringT
func TestCheckBinaryTypedExprUint8AndStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & string("abc")`, env,
		`invalid operation: 255 & "abc" (mismatched types uint8 and string)`,
	)

}

// Test Uint8 & BoolT
func TestCheckBinaryTypedExprUint8AndBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) & bool(true)`, env,
		`invalid operation: 255 & true (mismatched types uint8 and bool)`,
	)

}

// Test Uint8 % Int
func TestCheckBinaryTypedExprUint8RemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) % 4`, env, uint8(0xff) % 4, reflect.TypeOf(uint8(0xff) % 4))
}

// Test Uint8 % Rune
func TestCheckBinaryTypedExprUint8RemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) % '@'`, env, uint8(0xff) % '@', reflect.TypeOf(uint8(0xff) % '@'))
}

// Test Uint8 % Float
func TestCheckBinaryTypedExprUint8RemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) % 2.0`, env, uint8(0xff) % 2.0, reflect.TypeOf(uint8(0xff) % 2.0))
}

// Test Uint8 % Complex
func TestCheckBinaryTypedExprUint8RemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % 8.0i`, env,
		`constant 0+8i truncated to real`,
		`division by zero`,
	)

}

// Test Uint8 % Bool
func TestCheckBinaryTypedExprUint8RemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % true`, env,
		`cannot convert true to type uint8`,
		`invalid operation: 255 % true (mismatched types uint8 and bool)`,
	)

}

// Test Uint8 % String
func TestCheckBinaryTypedExprUint8RemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % "abc"`, env,
		`cannot convert "abc" to type uint8`,
		`invalid operation: 255 % "abc" (mismatched types uint8 and string)`,
	)

}

// Test Uint8 % Nil
func TestCheckBinaryTypedExprUint8RemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % nil`, env,
		`cannot convert nil to type uint8`,
	)

}

// Test Uint8 % Int8
func TestCheckBinaryTypedExprUint8RemInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % int8(0x7f)`, env,
		`invalid operation: 255 % 127 (mismatched types uint8 and int8)`,
	)

}

// Test Uint8 % Int16
func TestCheckBinaryTypedExprUint8RemInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % int16(0x7fff)`, env,
		`invalid operation: 255 % 32767 (mismatched types uint8 and int16)`,
	)

}

// Test Uint8 % Int32
func TestCheckBinaryTypedExprUint8RemInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % int32(0x7fffffff)`, env,
		`invalid operation: 255 % 2147483647 (mismatched types uint8 and int32)`,
	)

}

// Test Uint8 % Int64
func TestCheckBinaryTypedExprUint8RemInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % int64(0x7fffffffffffffff)`, env,
		`invalid operation: 255 % 9223372036854775807 (mismatched types uint8 and int64)`,
	)

}

// Test Uint8 % Uint8
func TestCheckBinaryTypedExprUint8RemUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) % uint8(0xff)`, env, uint8(0xff) % uint8(0xff), reflect.TypeOf(uint8(0xff) % uint8(0xff)))
}

// Test Uint8 % Uint16
func TestCheckBinaryTypedExprUint8RemUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % uint16(0xffff)`, env,
		`invalid operation: 255 % 65535 (mismatched types uint8 and uint16)`,
	)

}

// Test Uint8 % Uint32
func TestCheckBinaryTypedExprUint8RemUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % uint32(0xffffffff)`, env,
		`invalid operation: 255 % 4294967295 (mismatched types uint8 and uint32)`,
	)

}

// Test Uint8 % Uint64
func TestCheckBinaryTypedExprUint8RemUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % uint64(0xffffffffffffffff)`, env,
		`invalid operation: 255 % 18446744073709551615 (mismatched types uint8 and uint64)`,
	)

}

// Test Uint8 % Float32
func TestCheckBinaryTypedExprUint8RemFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % float32(0xffffffff)`, env,
		`invalid operation: 255 % 4.29497e+09 (mismatched types uint8 and float32)`,
	)

}

// Test Uint8 % Float64
func TestCheckBinaryTypedExprUint8RemFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % float64(0xffffffff)`, env,
		`invalid operation: 255 % 4.29497e+09 (mismatched types uint8 and float64)`,
	)

}

// Test Uint8 % Complex64
func TestCheckBinaryTypedExprUint8RemComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 255 % (4.29497e+09+4.29497e+09i) (mismatched types uint8 and complex64)`,
	)

}

// Test Uint8 % Complex128
func TestCheckBinaryTypedExprUint8RemComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 255 % (4.29497e+09+4.29497e+09i) (mismatched types uint8 and complex128)`,
	)

}

// Test Uint8 % Rune32
func TestCheckBinaryTypedExprUint8RemRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % rune(0x7fffffff)`, env,
		`invalid operation: 255 % rune(2147483647) (mismatched types uint8 and rune)`,
	)

}

// Test Uint8 % StringT
func TestCheckBinaryTypedExprUint8RemStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % string("abc")`, env,
		`invalid operation: 255 % "abc" (mismatched types uint8 and string)`,
	)

}

// Test Uint8 % BoolT
func TestCheckBinaryTypedExprUint8RemBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) % bool(true)`, env,
		`invalid operation: 255 % true (mismatched types uint8 and bool)`,
	)

}

// Test Uint8 == Int
func TestCheckBinaryTypedExprUint8EqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) == 4`, env, uint8(0xff) == 4, reflect.TypeOf(uint8(0xff) == 4))
}

// Test Uint8 == Rune
func TestCheckBinaryTypedExprUint8EqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) == '@'`, env, uint8(0xff) == '@', reflect.TypeOf(uint8(0xff) == '@'))
}

// Test Uint8 == Float
func TestCheckBinaryTypedExprUint8EqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) == 2.0`, env, uint8(0xff) == 2.0, reflect.TypeOf(uint8(0xff) == 2.0))
}

// Test Uint8 == Complex
func TestCheckBinaryTypedExprUint8EqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint8 == Bool
func TestCheckBinaryTypedExprUint8EqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == true`, env,
		`cannot convert true to type uint8`,
		`invalid operation: 255 == true (mismatched types uint8 and bool)`,
	)

}

// Test Uint8 == String
func TestCheckBinaryTypedExprUint8EqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == "abc"`, env,
		`cannot convert "abc" to type uint8`,
		`invalid operation: 255 == "abc" (mismatched types uint8 and string)`,
	)

}

// Test Uint8 == Nil
func TestCheckBinaryTypedExprUint8EqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == nil`, env,
		`cannot convert nil to type uint8`,
	)

}

// Test Uint8 == Int8
func TestCheckBinaryTypedExprUint8EqlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == int8(0x7f)`, env,
		`invalid operation: 255 == 127 (mismatched types uint8 and int8)`,
	)

}

// Test Uint8 == Int16
func TestCheckBinaryTypedExprUint8EqlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == int16(0x7fff)`, env,
		`invalid operation: 255 == 32767 (mismatched types uint8 and int16)`,
	)

}

// Test Uint8 == Int32
func TestCheckBinaryTypedExprUint8EqlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == int32(0x7fffffff)`, env,
		`invalid operation: 255 == 2147483647 (mismatched types uint8 and int32)`,
	)

}

// Test Uint8 == Int64
func TestCheckBinaryTypedExprUint8EqlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == int64(0x7fffffffffffffff)`, env,
		`invalid operation: 255 == 9223372036854775807 (mismatched types uint8 and int64)`,
	)

}

// Test Uint8 == Uint8
func TestCheckBinaryTypedExprUint8EqlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) == uint8(0xff)`, env, uint8(0xff) == uint8(0xff), reflect.TypeOf(uint8(0xff) == uint8(0xff)))
}

// Test Uint8 == Uint16
func TestCheckBinaryTypedExprUint8EqlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == uint16(0xffff)`, env,
		`invalid operation: 255 == 65535 (mismatched types uint8 and uint16)`,
	)

}

// Test Uint8 == Uint32
func TestCheckBinaryTypedExprUint8EqlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == uint32(0xffffffff)`, env,
		`invalid operation: 255 == 4294967295 (mismatched types uint8 and uint32)`,
	)

}

// Test Uint8 == Uint64
func TestCheckBinaryTypedExprUint8EqlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == uint64(0xffffffffffffffff)`, env,
		`invalid operation: 255 == 18446744073709551615 (mismatched types uint8 and uint64)`,
	)

}

// Test Uint8 == Float32
func TestCheckBinaryTypedExprUint8EqlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == float32(0xffffffff)`, env,
		`invalid operation: 255 == 4.29497e+09 (mismatched types uint8 and float32)`,
	)

}

// Test Uint8 == Float64
func TestCheckBinaryTypedExprUint8EqlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == float64(0xffffffff)`, env,
		`invalid operation: 255 == 4.29497e+09 (mismatched types uint8 and float64)`,
	)

}

// Test Uint8 == Complex64
func TestCheckBinaryTypedExprUint8EqlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 255 == (4.29497e+09+4.29497e+09i) (mismatched types uint8 and complex64)`,
	)

}

// Test Uint8 == Complex128
func TestCheckBinaryTypedExprUint8EqlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 255 == (4.29497e+09+4.29497e+09i) (mismatched types uint8 and complex128)`,
	)

}

// Test Uint8 == Rune32
func TestCheckBinaryTypedExprUint8EqlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == rune(0x7fffffff)`, env,
		`invalid operation: 255 == rune(2147483647) (mismatched types uint8 and rune)`,
	)

}

// Test Uint8 == StringT
func TestCheckBinaryTypedExprUint8EqlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == string("abc")`, env,
		`invalid operation: 255 == "abc" (mismatched types uint8 and string)`,
	)

}

// Test Uint8 == BoolT
func TestCheckBinaryTypedExprUint8EqlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) == bool(true)`, env,
		`invalid operation: 255 == true (mismatched types uint8 and bool)`,
	)

}

// Test Uint8 > Int
func TestCheckBinaryTypedExprUint8GtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) > 4`, env, uint8(0xff) > 4, reflect.TypeOf(uint8(0xff) > 4))
}

// Test Uint8 > Rune
func TestCheckBinaryTypedExprUint8GtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) > '@'`, env, uint8(0xff) > '@', reflect.TypeOf(uint8(0xff) > '@'))
}

// Test Uint8 > Float
func TestCheckBinaryTypedExprUint8GtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) > 2.0`, env, uint8(0xff) > 2.0, reflect.TypeOf(uint8(0xff) > 2.0))
}

// Test Uint8 > Complex
func TestCheckBinaryTypedExprUint8GtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint8 > Bool
func TestCheckBinaryTypedExprUint8GtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > true`, env,
		`cannot convert true to type uint8`,
		`invalid operation: 255 > true (mismatched types uint8 and bool)`,
	)

}

// Test Uint8 > String
func TestCheckBinaryTypedExprUint8GtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > "abc"`, env,
		`cannot convert "abc" to type uint8`,
		`invalid operation: 255 > "abc" (mismatched types uint8 and string)`,
	)

}

// Test Uint8 > Nil
func TestCheckBinaryTypedExprUint8GtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > nil`, env,
		`cannot convert nil to type uint8`,
	)

}

// Test Uint8 > Int8
func TestCheckBinaryTypedExprUint8GtrInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > int8(0x7f)`, env,
		`invalid operation: 255 > 127 (mismatched types uint8 and int8)`,
	)

}

// Test Uint8 > Int16
func TestCheckBinaryTypedExprUint8GtrInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > int16(0x7fff)`, env,
		`invalid operation: 255 > 32767 (mismatched types uint8 and int16)`,
	)

}

// Test Uint8 > Int32
func TestCheckBinaryTypedExprUint8GtrInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > int32(0x7fffffff)`, env,
		`invalid operation: 255 > 2147483647 (mismatched types uint8 and int32)`,
	)

}

// Test Uint8 > Int64
func TestCheckBinaryTypedExprUint8GtrInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > int64(0x7fffffffffffffff)`, env,
		`invalid operation: 255 > 9223372036854775807 (mismatched types uint8 and int64)`,
	)

}

// Test Uint8 > Uint8
func TestCheckBinaryTypedExprUint8GtrUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff) > uint8(0xff)`, env, uint8(0xff) > uint8(0xff), reflect.TypeOf(uint8(0xff) > uint8(0xff)))
}

// Test Uint8 > Uint16
func TestCheckBinaryTypedExprUint8GtrUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > uint16(0xffff)`, env,
		`invalid operation: 255 > 65535 (mismatched types uint8 and uint16)`,
	)

}

// Test Uint8 > Uint32
func TestCheckBinaryTypedExprUint8GtrUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > uint32(0xffffffff)`, env,
		`invalid operation: 255 > 4294967295 (mismatched types uint8 and uint32)`,
	)

}

// Test Uint8 > Uint64
func TestCheckBinaryTypedExprUint8GtrUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > uint64(0xffffffffffffffff)`, env,
		`invalid operation: 255 > 18446744073709551615 (mismatched types uint8 and uint64)`,
	)

}

// Test Uint8 > Float32
func TestCheckBinaryTypedExprUint8GtrFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > float32(0xffffffff)`, env,
		`invalid operation: 255 > 4.29497e+09 (mismatched types uint8 and float32)`,
	)

}

// Test Uint8 > Float64
func TestCheckBinaryTypedExprUint8GtrFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > float64(0xffffffff)`, env,
		`invalid operation: 255 > 4.29497e+09 (mismatched types uint8 and float64)`,
	)

}

// Test Uint8 > Complex64
func TestCheckBinaryTypedExprUint8GtrComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 255 > (4.29497e+09+4.29497e+09i) (mismatched types uint8 and complex64)`,
	)

}

// Test Uint8 > Complex128
func TestCheckBinaryTypedExprUint8GtrComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 255 > (4.29497e+09+4.29497e+09i) (mismatched types uint8 and complex128)`,
	)

}

// Test Uint8 > Rune32
func TestCheckBinaryTypedExprUint8GtrRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > rune(0x7fffffff)`, env,
		`invalid operation: 255 > rune(2147483647) (mismatched types uint8 and rune)`,
	)

}

// Test Uint8 > StringT
func TestCheckBinaryTypedExprUint8GtrStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > string("abc")`, env,
		`invalid operation: 255 > "abc" (mismatched types uint8 and string)`,
	)

}

// Test Uint8 > BoolT
func TestCheckBinaryTypedExprUint8GtrBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) > bool(true)`, env,
		`invalid operation: 255 > true (mismatched types uint8 and bool)`,
	)

}

// Test Uint8 << Int
func TestCheckBinaryTypedExprUint8ShlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << 4`, env,
		`constant 4080 overflows uint8`,
	)

}

// Test Uint8 << Rune
func TestCheckBinaryTypedExprUint8ShlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << '@'`, env,
		`constant 4703919738795935662080 overflows uint8`,
	)

}

// Test Uint8 << Float
func TestCheckBinaryTypedExprUint8ShlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << 2.0`, env,
		`constant 1020 overflows uint8`,
	)

}

// Test Uint8 << Complex
func TestCheckBinaryTypedExprUint8ShlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint8 << Bool
func TestCheckBinaryTypedExprUint8ShlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << true`, env,
		`invalid operation: 255 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Uint8 << String
func TestCheckBinaryTypedExprUint8ShlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: 255 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Uint8 << Nil
func TestCheckBinaryTypedExprUint8ShlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Uint8 << Int8
func TestCheckBinaryTypedExprUint8ShlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << int8(0x7f)`, env,
		`invalid operation: 255 << 127 (shift count type int8, must be unsigned integer)`,
	)

}

// Test Uint8 << Int16
func TestCheckBinaryTypedExprUint8ShlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << int16(0x7fff)`, env,
		`invalid operation: 255 << 32767 (shift count type int16, must be unsigned integer)`,
	)

}

// Test Uint8 << Int32
func TestCheckBinaryTypedExprUint8ShlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << int32(0x7fffffff)`, env,
		`invalid operation: 255 << 2147483647 (shift count type int32, must be unsigned integer)`,
	)

}

// Test Uint8 << Int64
func TestCheckBinaryTypedExprUint8ShlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << int64(0x7fffffffffffffff)`, env,
		`invalid operation: 255 << 9223372036854775807 (shift count type int64, must be unsigned integer)`,
	)

}

// Test Uint8 << Uint8
func TestCheckBinaryTypedExprUint8ShlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << uint8(0xff)`, env,
		`constant 14763491377757814916505300588607708251291923044869171915030841961008924029091840 overflows uint8`,
	)

}

// Test Uint8 << Uint16
func TestCheckBinaryTypedExprUint8ShlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << uint16(0xffff)`, env,
		`stupid shift: 65535`,
	)

}

// Test Uint8 << Uint32
func TestCheckBinaryTypedExprUint8ShlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << uint32(0xffffffff)`, env,
		`stupid shift: 4294967295`,
	)

}

// Test Uint8 << Uint64
func TestCheckBinaryTypedExprUint8ShlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << uint64(0xffffffffffffffff)`, env,
		`stupid shift: -1`,
	)

}

// Test Uint8 << Float32
func TestCheckBinaryTypedExprUint8ShlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << float32(0xffffffff)`, env,
		`invalid operation: 255 << 4.29497e+09 (shift count type float32, must be unsigned integer)`,
	)

}

// Test Uint8 << Float64
func TestCheckBinaryTypedExprUint8ShlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << float64(0xffffffff)`, env,
		`invalid operation: 255 << 4.29497e+09 (shift count type float64, must be unsigned integer)`,
	)

}

// Test Uint8 << Complex64
func TestCheckBinaryTypedExprUint8ShlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 255 << (4.29497e+09+4.29497e+09i) (shift count type complex64, must be unsigned integer)`,
	)

}

// Test Uint8 << Complex128
func TestCheckBinaryTypedExprUint8ShlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 255 << (4.29497e+09+4.29497e+09i) (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Uint8 << Rune32
func TestCheckBinaryTypedExprUint8ShlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << rune(0x7fffffff)`, env,
		`invalid operation: 255 << rune(2147483647) (shift count type rune, must be unsigned integer)`,
	)

}

// Test Uint8 << StringT
func TestCheckBinaryTypedExprUint8ShlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << string("abc")`, env,
		`invalid operation: 255 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Uint8 << BoolT
func TestCheckBinaryTypedExprUint8ShlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xff) << bool(true)`, env,
		`invalid operation: 255 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Uint16 + Int
func TestCheckBinaryTypedExprUint16AddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + 4`, env,
		`constant 65539 overflows uint16`,
	)

}

// Test Uint16 + Rune
func TestCheckBinaryTypedExprUint16AddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + '@'`, env,
		`constant 65599 overflows uint16`,
	)

}

// Test Uint16 + Float
func TestCheckBinaryTypedExprUint16AddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + 2.0`, env,
		`constant 65537 overflows uint16`,
	)

}

// Test Uint16 + Complex
func TestCheckBinaryTypedExprUint16AddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint16 + Bool
func TestCheckBinaryTypedExprUint16AddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + true`, env,
		`cannot convert true to type uint16`,
		`invalid operation: 65535 + true (mismatched types uint16 and bool)`,
	)

}

// Test Uint16 + String
func TestCheckBinaryTypedExprUint16AddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + "abc"`, env,
		`cannot convert "abc" to type uint16`,
		`invalid operation: 65535 + "abc" (mismatched types uint16 and string)`,
	)

}

// Test Uint16 + Nil
func TestCheckBinaryTypedExprUint16AddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + nil`, env,
		`cannot convert nil to type uint16`,
	)

}

// Test Uint16 + Int8
func TestCheckBinaryTypedExprUint16AddInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + int8(0x7f)`, env,
		`invalid operation: 65535 + 127 (mismatched types uint16 and int8)`,
	)

}

// Test Uint16 + Int16
func TestCheckBinaryTypedExprUint16AddInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + int16(0x7fff)`, env,
		`invalid operation: 65535 + 32767 (mismatched types uint16 and int16)`,
	)

}

// Test Uint16 + Int32
func TestCheckBinaryTypedExprUint16AddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + int32(0x7fffffff)`, env,
		`invalid operation: 65535 + 2147483647 (mismatched types uint16 and int32)`,
	)

}

// Test Uint16 + Int64
func TestCheckBinaryTypedExprUint16AddInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + int64(0x7fffffffffffffff)`, env,
		`invalid operation: 65535 + 9223372036854775807 (mismatched types uint16 and int64)`,
	)

}

// Test Uint16 + Uint8
func TestCheckBinaryTypedExprUint16AddUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + uint8(0xff)`, env,
		`invalid operation: 65535 + 255 (mismatched types uint16 and uint8)`,
	)

}

// Test Uint16 + Uint16
func TestCheckBinaryTypedExprUint16AddUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + uint16(0xffff)`, env,
		`constant 131070 overflows uint16`,
	)

}

// Test Uint16 + Uint32
func TestCheckBinaryTypedExprUint16AddUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + uint32(0xffffffff)`, env,
		`invalid operation: 65535 + 4294967295 (mismatched types uint16 and uint32)`,
	)

}

// Test Uint16 + Uint64
func TestCheckBinaryTypedExprUint16AddUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + uint64(0xffffffffffffffff)`, env,
		`invalid operation: 65535 + 18446744073709551615 (mismatched types uint16 and uint64)`,
	)

}

// Test Uint16 + Float32
func TestCheckBinaryTypedExprUint16AddFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + float32(0xffffffff)`, env,
		`invalid operation: 65535 + 4.29497e+09 (mismatched types uint16 and float32)`,
	)

}

// Test Uint16 + Float64
func TestCheckBinaryTypedExprUint16AddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + float64(0xffffffff)`, env,
		`invalid operation: 65535 + 4.29497e+09 (mismatched types uint16 and float64)`,
	)

}

// Test Uint16 + Complex64
func TestCheckBinaryTypedExprUint16AddComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 65535 + (4.29497e+09+4.29497e+09i) (mismatched types uint16 and complex64)`,
	)

}

// Test Uint16 + Complex128
func TestCheckBinaryTypedExprUint16AddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 65535 + (4.29497e+09+4.29497e+09i) (mismatched types uint16 and complex128)`,
	)

}

// Test Uint16 + Rune32
func TestCheckBinaryTypedExprUint16AddRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + rune(0x7fffffff)`, env,
		`invalid operation: 65535 + rune(2147483647) (mismatched types uint16 and rune)`,
	)

}

// Test Uint16 + StringT
func TestCheckBinaryTypedExprUint16AddStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + string("abc")`, env,
		`invalid operation: 65535 + "abc" (mismatched types uint16 and string)`,
	)

}

// Test Uint16 + BoolT
func TestCheckBinaryTypedExprUint16AddBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) + bool(true)`, env,
		`invalid operation: 65535 + true (mismatched types uint16 and bool)`,
	)

}

// Test Uint16 - Int
func TestCheckBinaryTypedExprUint16SubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) - 4`, env, uint16(0xffff) - 4, reflect.TypeOf(uint16(0xffff) - 4))
}

// Test Uint16 - Rune
func TestCheckBinaryTypedExprUint16SubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) - '@'`, env, uint16(0xffff) - '@', reflect.TypeOf(uint16(0xffff) - '@'))
}

// Test Uint16 - Float
func TestCheckBinaryTypedExprUint16SubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) - 2.0`, env, uint16(0xffff) - 2.0, reflect.TypeOf(uint16(0xffff) - 2.0))
}

// Test Uint16 - Complex
func TestCheckBinaryTypedExprUint16SubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint16 - Bool
func TestCheckBinaryTypedExprUint16SubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - true`, env,
		`cannot convert true to type uint16`,
		`invalid operation: 65535 - true (mismatched types uint16 and bool)`,
	)

}

// Test Uint16 - String
func TestCheckBinaryTypedExprUint16SubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - "abc"`, env,
		`cannot convert "abc" to type uint16`,
		`invalid operation: 65535 - "abc" (mismatched types uint16 and string)`,
	)

}

// Test Uint16 - Nil
func TestCheckBinaryTypedExprUint16SubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - nil`, env,
		`cannot convert nil to type uint16`,
	)

}

// Test Uint16 - Int8
func TestCheckBinaryTypedExprUint16SubInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - int8(0x7f)`, env,
		`invalid operation: 65535 - 127 (mismatched types uint16 and int8)`,
	)

}

// Test Uint16 - Int16
func TestCheckBinaryTypedExprUint16SubInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - int16(0x7fff)`, env,
		`invalid operation: 65535 - 32767 (mismatched types uint16 and int16)`,
	)

}

// Test Uint16 - Int32
func TestCheckBinaryTypedExprUint16SubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - int32(0x7fffffff)`, env,
		`invalid operation: 65535 - 2147483647 (mismatched types uint16 and int32)`,
	)

}

// Test Uint16 - Int64
func TestCheckBinaryTypedExprUint16SubInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - int64(0x7fffffffffffffff)`, env,
		`invalid operation: 65535 - 9223372036854775807 (mismatched types uint16 and int64)`,
	)

}

// Test Uint16 - Uint8
func TestCheckBinaryTypedExprUint16SubUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - uint8(0xff)`, env,
		`invalid operation: 65535 - 255 (mismatched types uint16 and uint8)`,
	)

}

// Test Uint16 - Uint16
func TestCheckBinaryTypedExprUint16SubUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) - uint16(0xffff)`, env, uint16(0xffff) - uint16(0xffff), reflect.TypeOf(uint16(0xffff) - uint16(0xffff)))
}

// Test Uint16 - Uint32
func TestCheckBinaryTypedExprUint16SubUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - uint32(0xffffffff)`, env,
		`invalid operation: 65535 - 4294967295 (mismatched types uint16 and uint32)`,
	)

}

// Test Uint16 - Uint64
func TestCheckBinaryTypedExprUint16SubUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - uint64(0xffffffffffffffff)`, env,
		`invalid operation: 65535 - 18446744073709551615 (mismatched types uint16 and uint64)`,
	)

}

// Test Uint16 - Float32
func TestCheckBinaryTypedExprUint16SubFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - float32(0xffffffff)`, env,
		`invalid operation: 65535 - 4.29497e+09 (mismatched types uint16 and float32)`,
	)

}

// Test Uint16 - Float64
func TestCheckBinaryTypedExprUint16SubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - float64(0xffffffff)`, env,
		`invalid operation: 65535 - 4.29497e+09 (mismatched types uint16 and float64)`,
	)

}

// Test Uint16 - Complex64
func TestCheckBinaryTypedExprUint16SubComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 65535 - (4.29497e+09+4.29497e+09i) (mismatched types uint16 and complex64)`,
	)

}

// Test Uint16 - Complex128
func TestCheckBinaryTypedExprUint16SubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 65535 - (4.29497e+09+4.29497e+09i) (mismatched types uint16 and complex128)`,
	)

}

// Test Uint16 - Rune32
func TestCheckBinaryTypedExprUint16SubRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - rune(0x7fffffff)`, env,
		`invalid operation: 65535 - rune(2147483647) (mismatched types uint16 and rune)`,
	)

}

// Test Uint16 - StringT
func TestCheckBinaryTypedExprUint16SubStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - string("abc")`, env,
		`invalid operation: 65535 - "abc" (mismatched types uint16 and string)`,
	)

}

// Test Uint16 - BoolT
func TestCheckBinaryTypedExprUint16SubBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) - bool(true)`, env,
		`invalid operation: 65535 - true (mismatched types uint16 and bool)`,
	)

}

// Test Uint16 & Int
func TestCheckBinaryTypedExprUint16AndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) & 4`, env, uint16(0xffff) & 4, reflect.TypeOf(uint16(0xffff) & 4))
}

// Test Uint16 & Rune
func TestCheckBinaryTypedExprUint16AndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) & '@'`, env, uint16(0xffff) & '@', reflect.TypeOf(uint16(0xffff) & '@'))
}

// Test Uint16 & Float
func TestCheckBinaryTypedExprUint16AndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) & 2.0`, env, uint16(0xffff) & 2.0, reflect.TypeOf(uint16(0xffff) & 2.0))
}

// Test Uint16 & Complex
func TestCheckBinaryTypedExprUint16AndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint16 & Bool
func TestCheckBinaryTypedExprUint16AndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & true`, env,
		`cannot convert true to type uint16`,
		`invalid operation: 65535 & true (mismatched types uint16 and bool)`,
	)

}

// Test Uint16 & String
func TestCheckBinaryTypedExprUint16AndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & "abc"`, env,
		`cannot convert "abc" to type uint16`,
		`invalid operation: 65535 & "abc" (mismatched types uint16 and string)`,
	)

}

// Test Uint16 & Nil
func TestCheckBinaryTypedExprUint16AndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & nil`, env,
		`cannot convert nil to type uint16`,
	)

}

// Test Uint16 & Int8
func TestCheckBinaryTypedExprUint16AndInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & int8(0x7f)`, env,
		`invalid operation: 65535 & 127 (mismatched types uint16 and int8)`,
	)

}

// Test Uint16 & Int16
func TestCheckBinaryTypedExprUint16AndInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & int16(0x7fff)`, env,
		`invalid operation: 65535 & 32767 (mismatched types uint16 and int16)`,
	)

}

// Test Uint16 & Int32
func TestCheckBinaryTypedExprUint16AndInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & int32(0x7fffffff)`, env,
		`invalid operation: 65535 & 2147483647 (mismatched types uint16 and int32)`,
	)

}

// Test Uint16 & Int64
func TestCheckBinaryTypedExprUint16AndInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & int64(0x7fffffffffffffff)`, env,
		`invalid operation: 65535 & 9223372036854775807 (mismatched types uint16 and int64)`,
	)

}

// Test Uint16 & Uint8
func TestCheckBinaryTypedExprUint16AndUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & uint8(0xff)`, env,
		`invalid operation: 65535 & 255 (mismatched types uint16 and uint8)`,
	)

}

// Test Uint16 & Uint16
func TestCheckBinaryTypedExprUint16AndUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) & uint16(0xffff)`, env, uint16(0xffff) & uint16(0xffff), reflect.TypeOf(uint16(0xffff) & uint16(0xffff)))
}

// Test Uint16 & Uint32
func TestCheckBinaryTypedExprUint16AndUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & uint32(0xffffffff)`, env,
		`invalid operation: 65535 & 4294967295 (mismatched types uint16 and uint32)`,
	)

}

// Test Uint16 & Uint64
func TestCheckBinaryTypedExprUint16AndUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & uint64(0xffffffffffffffff)`, env,
		`invalid operation: 65535 & 18446744073709551615 (mismatched types uint16 and uint64)`,
	)

}

// Test Uint16 & Float32
func TestCheckBinaryTypedExprUint16AndFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & float32(0xffffffff)`, env,
		`invalid operation: 65535 & 4.29497e+09 (mismatched types uint16 and float32)`,
	)

}

// Test Uint16 & Float64
func TestCheckBinaryTypedExprUint16AndFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & float64(0xffffffff)`, env,
		`invalid operation: 65535 & 4.29497e+09 (mismatched types uint16 and float64)`,
	)

}

// Test Uint16 & Complex64
func TestCheckBinaryTypedExprUint16AndComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 65535 & (4.29497e+09+4.29497e+09i) (mismatched types uint16 and complex64)`,
	)

}

// Test Uint16 & Complex128
func TestCheckBinaryTypedExprUint16AndComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 65535 & (4.29497e+09+4.29497e+09i) (mismatched types uint16 and complex128)`,
	)

}

// Test Uint16 & Rune32
func TestCheckBinaryTypedExprUint16AndRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & rune(0x7fffffff)`, env,
		`invalid operation: 65535 & rune(2147483647) (mismatched types uint16 and rune)`,
	)

}

// Test Uint16 & StringT
func TestCheckBinaryTypedExprUint16AndStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & string("abc")`, env,
		`invalid operation: 65535 & "abc" (mismatched types uint16 and string)`,
	)

}

// Test Uint16 & BoolT
func TestCheckBinaryTypedExprUint16AndBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) & bool(true)`, env,
		`invalid operation: 65535 & true (mismatched types uint16 and bool)`,
	)

}

// Test Uint16 % Int
func TestCheckBinaryTypedExprUint16RemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) % 4`, env, uint16(0xffff) % 4, reflect.TypeOf(uint16(0xffff) % 4))
}

// Test Uint16 % Rune
func TestCheckBinaryTypedExprUint16RemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) % '@'`, env, uint16(0xffff) % '@', reflect.TypeOf(uint16(0xffff) % '@'))
}

// Test Uint16 % Float
func TestCheckBinaryTypedExprUint16RemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) % 2.0`, env, uint16(0xffff) % 2.0, reflect.TypeOf(uint16(0xffff) % 2.0))
}

// Test Uint16 % Complex
func TestCheckBinaryTypedExprUint16RemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % 8.0i`, env,
		`constant 0+8i truncated to real`,
		`division by zero`,
	)

}

// Test Uint16 % Bool
func TestCheckBinaryTypedExprUint16RemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % true`, env,
		`cannot convert true to type uint16`,
		`invalid operation: 65535 % true (mismatched types uint16 and bool)`,
	)

}

// Test Uint16 % String
func TestCheckBinaryTypedExprUint16RemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % "abc"`, env,
		`cannot convert "abc" to type uint16`,
		`invalid operation: 65535 % "abc" (mismatched types uint16 and string)`,
	)

}

// Test Uint16 % Nil
func TestCheckBinaryTypedExprUint16RemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % nil`, env,
		`cannot convert nil to type uint16`,
	)

}

// Test Uint16 % Int8
func TestCheckBinaryTypedExprUint16RemInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % int8(0x7f)`, env,
		`invalid operation: 65535 % 127 (mismatched types uint16 and int8)`,
	)

}

// Test Uint16 % Int16
func TestCheckBinaryTypedExprUint16RemInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % int16(0x7fff)`, env,
		`invalid operation: 65535 % 32767 (mismatched types uint16 and int16)`,
	)

}

// Test Uint16 % Int32
func TestCheckBinaryTypedExprUint16RemInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % int32(0x7fffffff)`, env,
		`invalid operation: 65535 % 2147483647 (mismatched types uint16 and int32)`,
	)

}

// Test Uint16 % Int64
func TestCheckBinaryTypedExprUint16RemInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % int64(0x7fffffffffffffff)`, env,
		`invalid operation: 65535 % 9223372036854775807 (mismatched types uint16 and int64)`,
	)

}

// Test Uint16 % Uint8
func TestCheckBinaryTypedExprUint16RemUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % uint8(0xff)`, env,
		`invalid operation: 65535 % 255 (mismatched types uint16 and uint8)`,
	)

}

// Test Uint16 % Uint16
func TestCheckBinaryTypedExprUint16RemUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) % uint16(0xffff)`, env, uint16(0xffff) % uint16(0xffff), reflect.TypeOf(uint16(0xffff) % uint16(0xffff)))
}

// Test Uint16 % Uint32
func TestCheckBinaryTypedExprUint16RemUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % uint32(0xffffffff)`, env,
		`invalid operation: 65535 % 4294967295 (mismatched types uint16 and uint32)`,
	)

}

// Test Uint16 % Uint64
func TestCheckBinaryTypedExprUint16RemUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % uint64(0xffffffffffffffff)`, env,
		`invalid operation: 65535 % 18446744073709551615 (mismatched types uint16 and uint64)`,
	)

}

// Test Uint16 % Float32
func TestCheckBinaryTypedExprUint16RemFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % float32(0xffffffff)`, env,
		`invalid operation: 65535 % 4.29497e+09 (mismatched types uint16 and float32)`,
	)

}

// Test Uint16 % Float64
func TestCheckBinaryTypedExprUint16RemFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % float64(0xffffffff)`, env,
		`invalid operation: 65535 % 4.29497e+09 (mismatched types uint16 and float64)`,
	)

}

// Test Uint16 % Complex64
func TestCheckBinaryTypedExprUint16RemComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 65535 % (4.29497e+09+4.29497e+09i) (mismatched types uint16 and complex64)`,
	)

}

// Test Uint16 % Complex128
func TestCheckBinaryTypedExprUint16RemComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 65535 % (4.29497e+09+4.29497e+09i) (mismatched types uint16 and complex128)`,
	)

}

// Test Uint16 % Rune32
func TestCheckBinaryTypedExprUint16RemRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % rune(0x7fffffff)`, env,
		`invalid operation: 65535 % rune(2147483647) (mismatched types uint16 and rune)`,
	)

}

// Test Uint16 % StringT
func TestCheckBinaryTypedExprUint16RemStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % string("abc")`, env,
		`invalid operation: 65535 % "abc" (mismatched types uint16 and string)`,
	)

}

// Test Uint16 % BoolT
func TestCheckBinaryTypedExprUint16RemBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) % bool(true)`, env,
		`invalid operation: 65535 % true (mismatched types uint16 and bool)`,
	)

}

// Test Uint16 == Int
func TestCheckBinaryTypedExprUint16EqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) == 4`, env, uint16(0xffff) == 4, reflect.TypeOf(uint16(0xffff) == 4))
}

// Test Uint16 == Rune
func TestCheckBinaryTypedExprUint16EqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) == '@'`, env, uint16(0xffff) == '@', reflect.TypeOf(uint16(0xffff) == '@'))
}

// Test Uint16 == Float
func TestCheckBinaryTypedExprUint16EqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) == 2.0`, env, uint16(0xffff) == 2.0, reflect.TypeOf(uint16(0xffff) == 2.0))
}

// Test Uint16 == Complex
func TestCheckBinaryTypedExprUint16EqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint16 == Bool
func TestCheckBinaryTypedExprUint16EqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == true`, env,
		`cannot convert true to type uint16`,
		`invalid operation: 65535 == true (mismatched types uint16 and bool)`,
	)

}

// Test Uint16 == String
func TestCheckBinaryTypedExprUint16EqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == "abc"`, env,
		`cannot convert "abc" to type uint16`,
		`invalid operation: 65535 == "abc" (mismatched types uint16 and string)`,
	)

}

// Test Uint16 == Nil
func TestCheckBinaryTypedExprUint16EqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == nil`, env,
		`cannot convert nil to type uint16`,
	)

}

// Test Uint16 == Int8
func TestCheckBinaryTypedExprUint16EqlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == int8(0x7f)`, env,
		`invalid operation: 65535 == 127 (mismatched types uint16 and int8)`,
	)

}

// Test Uint16 == Int16
func TestCheckBinaryTypedExprUint16EqlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == int16(0x7fff)`, env,
		`invalid operation: 65535 == 32767 (mismatched types uint16 and int16)`,
	)

}

// Test Uint16 == Int32
func TestCheckBinaryTypedExprUint16EqlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == int32(0x7fffffff)`, env,
		`invalid operation: 65535 == 2147483647 (mismatched types uint16 and int32)`,
	)

}

// Test Uint16 == Int64
func TestCheckBinaryTypedExprUint16EqlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == int64(0x7fffffffffffffff)`, env,
		`invalid operation: 65535 == 9223372036854775807 (mismatched types uint16 and int64)`,
	)

}

// Test Uint16 == Uint8
func TestCheckBinaryTypedExprUint16EqlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == uint8(0xff)`, env,
		`invalid operation: 65535 == 255 (mismatched types uint16 and uint8)`,
	)

}

// Test Uint16 == Uint16
func TestCheckBinaryTypedExprUint16EqlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) == uint16(0xffff)`, env, uint16(0xffff) == uint16(0xffff), reflect.TypeOf(uint16(0xffff) == uint16(0xffff)))
}

// Test Uint16 == Uint32
func TestCheckBinaryTypedExprUint16EqlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == uint32(0xffffffff)`, env,
		`invalid operation: 65535 == 4294967295 (mismatched types uint16 and uint32)`,
	)

}

// Test Uint16 == Uint64
func TestCheckBinaryTypedExprUint16EqlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == uint64(0xffffffffffffffff)`, env,
		`invalid operation: 65535 == 18446744073709551615 (mismatched types uint16 and uint64)`,
	)

}

// Test Uint16 == Float32
func TestCheckBinaryTypedExprUint16EqlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == float32(0xffffffff)`, env,
		`invalid operation: 65535 == 4.29497e+09 (mismatched types uint16 and float32)`,
	)

}

// Test Uint16 == Float64
func TestCheckBinaryTypedExprUint16EqlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == float64(0xffffffff)`, env,
		`invalid operation: 65535 == 4.29497e+09 (mismatched types uint16 and float64)`,
	)

}

// Test Uint16 == Complex64
func TestCheckBinaryTypedExprUint16EqlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 65535 == (4.29497e+09+4.29497e+09i) (mismatched types uint16 and complex64)`,
	)

}

// Test Uint16 == Complex128
func TestCheckBinaryTypedExprUint16EqlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 65535 == (4.29497e+09+4.29497e+09i) (mismatched types uint16 and complex128)`,
	)

}

// Test Uint16 == Rune32
func TestCheckBinaryTypedExprUint16EqlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == rune(0x7fffffff)`, env,
		`invalid operation: 65535 == rune(2147483647) (mismatched types uint16 and rune)`,
	)

}

// Test Uint16 == StringT
func TestCheckBinaryTypedExprUint16EqlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == string("abc")`, env,
		`invalid operation: 65535 == "abc" (mismatched types uint16 and string)`,
	)

}

// Test Uint16 == BoolT
func TestCheckBinaryTypedExprUint16EqlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) == bool(true)`, env,
		`invalid operation: 65535 == true (mismatched types uint16 and bool)`,
	)

}

// Test Uint16 > Int
func TestCheckBinaryTypedExprUint16GtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) > 4`, env, uint16(0xffff) > 4, reflect.TypeOf(uint16(0xffff) > 4))
}

// Test Uint16 > Rune
func TestCheckBinaryTypedExprUint16GtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) > '@'`, env, uint16(0xffff) > '@', reflect.TypeOf(uint16(0xffff) > '@'))
}

// Test Uint16 > Float
func TestCheckBinaryTypedExprUint16GtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) > 2.0`, env, uint16(0xffff) > 2.0, reflect.TypeOf(uint16(0xffff) > 2.0))
}

// Test Uint16 > Complex
func TestCheckBinaryTypedExprUint16GtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint16 > Bool
func TestCheckBinaryTypedExprUint16GtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > true`, env,
		`cannot convert true to type uint16`,
		`invalid operation: 65535 > true (mismatched types uint16 and bool)`,
	)

}

// Test Uint16 > String
func TestCheckBinaryTypedExprUint16GtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > "abc"`, env,
		`cannot convert "abc" to type uint16`,
		`invalid operation: 65535 > "abc" (mismatched types uint16 and string)`,
	)

}

// Test Uint16 > Nil
func TestCheckBinaryTypedExprUint16GtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > nil`, env,
		`cannot convert nil to type uint16`,
	)

}

// Test Uint16 > Int8
func TestCheckBinaryTypedExprUint16GtrInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > int8(0x7f)`, env,
		`invalid operation: 65535 > 127 (mismatched types uint16 and int8)`,
	)

}

// Test Uint16 > Int16
func TestCheckBinaryTypedExprUint16GtrInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > int16(0x7fff)`, env,
		`invalid operation: 65535 > 32767 (mismatched types uint16 and int16)`,
	)

}

// Test Uint16 > Int32
func TestCheckBinaryTypedExprUint16GtrInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > int32(0x7fffffff)`, env,
		`invalid operation: 65535 > 2147483647 (mismatched types uint16 and int32)`,
	)

}

// Test Uint16 > Int64
func TestCheckBinaryTypedExprUint16GtrInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > int64(0x7fffffffffffffff)`, env,
		`invalid operation: 65535 > 9223372036854775807 (mismatched types uint16 and int64)`,
	)

}

// Test Uint16 > Uint8
func TestCheckBinaryTypedExprUint16GtrUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > uint8(0xff)`, env,
		`invalid operation: 65535 > 255 (mismatched types uint16 and uint8)`,
	)

}

// Test Uint16 > Uint16
func TestCheckBinaryTypedExprUint16GtrUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff) > uint16(0xffff)`, env, uint16(0xffff) > uint16(0xffff), reflect.TypeOf(uint16(0xffff) > uint16(0xffff)))
}

// Test Uint16 > Uint32
func TestCheckBinaryTypedExprUint16GtrUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > uint32(0xffffffff)`, env,
		`invalid operation: 65535 > 4294967295 (mismatched types uint16 and uint32)`,
	)

}

// Test Uint16 > Uint64
func TestCheckBinaryTypedExprUint16GtrUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > uint64(0xffffffffffffffff)`, env,
		`invalid operation: 65535 > 18446744073709551615 (mismatched types uint16 and uint64)`,
	)

}

// Test Uint16 > Float32
func TestCheckBinaryTypedExprUint16GtrFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > float32(0xffffffff)`, env,
		`invalid operation: 65535 > 4.29497e+09 (mismatched types uint16 and float32)`,
	)

}

// Test Uint16 > Float64
func TestCheckBinaryTypedExprUint16GtrFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > float64(0xffffffff)`, env,
		`invalid operation: 65535 > 4.29497e+09 (mismatched types uint16 and float64)`,
	)

}

// Test Uint16 > Complex64
func TestCheckBinaryTypedExprUint16GtrComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 65535 > (4.29497e+09+4.29497e+09i) (mismatched types uint16 and complex64)`,
	)

}

// Test Uint16 > Complex128
func TestCheckBinaryTypedExprUint16GtrComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 65535 > (4.29497e+09+4.29497e+09i) (mismatched types uint16 and complex128)`,
	)

}

// Test Uint16 > Rune32
func TestCheckBinaryTypedExprUint16GtrRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > rune(0x7fffffff)`, env,
		`invalid operation: 65535 > rune(2147483647) (mismatched types uint16 and rune)`,
	)

}

// Test Uint16 > StringT
func TestCheckBinaryTypedExprUint16GtrStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > string("abc")`, env,
		`invalid operation: 65535 > "abc" (mismatched types uint16 and string)`,
	)

}

// Test Uint16 > BoolT
func TestCheckBinaryTypedExprUint16GtrBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) > bool(true)`, env,
		`invalid operation: 65535 > true (mismatched types uint16 and bool)`,
	)

}

// Test Uint16 << Int
func TestCheckBinaryTypedExprUint16ShlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << 4`, env,
		`constant 1048560 overflows uint16`,
	)

}

// Test Uint16 << Rune
func TestCheckBinaryTypedExprUint16ShlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << '@'`, env,
		`constant 1208907372870555465154560 overflows uint16`,
	)

}

// Test Uint16 << Float
func TestCheckBinaryTypedExprUint16ShlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << 2.0`, env,
		`constant 262140 overflows uint16`,
	)

}

// Test Uint16 << Complex
func TestCheckBinaryTypedExprUint16ShlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint16 << Bool
func TestCheckBinaryTypedExprUint16ShlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << true`, env,
		`invalid operation: 65535 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Uint16 << String
func TestCheckBinaryTypedExprUint16ShlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: 65535 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Uint16 << Nil
func TestCheckBinaryTypedExprUint16ShlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Uint16 << Int8
func TestCheckBinaryTypedExprUint16ShlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << int8(0x7f)`, env,
		`invalid operation: 65535 << 127 (shift count type int8, must be unsigned integer)`,
	)

}

// Test Uint16 << Int16
func TestCheckBinaryTypedExprUint16ShlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << int16(0x7fff)`, env,
		`invalid operation: 65535 << 32767 (shift count type int16, must be unsigned integer)`,
	)

}

// Test Uint16 << Int32
func TestCheckBinaryTypedExprUint16ShlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << int32(0x7fffffff)`, env,
		`invalid operation: 65535 << 2147483647 (shift count type int32, must be unsigned integer)`,
	)

}

// Test Uint16 << Int64
func TestCheckBinaryTypedExprUint16ShlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << int64(0x7fffffffffffffff)`, env,
		`invalid operation: 65535 << 9223372036854775807 (shift count type int64, must be unsigned integer)`,
	)

}

// Test Uint16 << Uint8
func TestCheckBinaryTypedExprUint16ShlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << uint8(0xff)`, env,
		`constant 3794217284083758433541862251272181020582024222531377182162926383979293475476602880 overflows uint16`,
	)

}

// Test Uint16 << Uint16
func TestCheckBinaryTypedExprUint16ShlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << uint16(0xffff)`, env,
		`stupid shift: 65535`,
	)

}

// Test Uint16 << Uint32
func TestCheckBinaryTypedExprUint16ShlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << uint32(0xffffffff)`, env,
		`stupid shift: 4294967295`,
	)

}

// Test Uint16 << Uint64
func TestCheckBinaryTypedExprUint16ShlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << uint64(0xffffffffffffffff)`, env,
		`stupid shift: -1`,
	)

}

// Test Uint16 << Float32
func TestCheckBinaryTypedExprUint16ShlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << float32(0xffffffff)`, env,
		`invalid operation: 65535 << 4.29497e+09 (shift count type float32, must be unsigned integer)`,
	)

}

// Test Uint16 << Float64
func TestCheckBinaryTypedExprUint16ShlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << float64(0xffffffff)`, env,
		`invalid operation: 65535 << 4.29497e+09 (shift count type float64, must be unsigned integer)`,
	)

}

// Test Uint16 << Complex64
func TestCheckBinaryTypedExprUint16ShlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 65535 << (4.29497e+09+4.29497e+09i) (shift count type complex64, must be unsigned integer)`,
	)

}

// Test Uint16 << Complex128
func TestCheckBinaryTypedExprUint16ShlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 65535 << (4.29497e+09+4.29497e+09i) (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Uint16 << Rune32
func TestCheckBinaryTypedExprUint16ShlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << rune(0x7fffffff)`, env,
		`invalid operation: 65535 << rune(2147483647) (shift count type rune, must be unsigned integer)`,
	)

}

// Test Uint16 << StringT
func TestCheckBinaryTypedExprUint16ShlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << string("abc")`, env,
		`invalid operation: 65535 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Uint16 << BoolT
func TestCheckBinaryTypedExprUint16ShlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffff) << bool(true)`, env,
		`invalid operation: 65535 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Uint32 + Int
func TestCheckBinaryTypedExprUint32AddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + 4`, env,
		`constant 4294967299 overflows uint32`,
	)

}

// Test Uint32 + Rune
func TestCheckBinaryTypedExprUint32AddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + '@'`, env,
		`constant 4294967359 overflows uint32`,
	)

}

// Test Uint32 + Float
func TestCheckBinaryTypedExprUint32AddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + 2.0`, env,
		`constant 4294967297 overflows uint32`,
	)

}

// Test Uint32 + Complex
func TestCheckBinaryTypedExprUint32AddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint32 + Bool
func TestCheckBinaryTypedExprUint32AddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + true`, env,
		`cannot convert true to type uint32`,
		`invalid operation: 4294967295 + true (mismatched types uint32 and bool)`,
	)

}

// Test Uint32 + String
func TestCheckBinaryTypedExprUint32AddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + "abc"`, env,
		`cannot convert "abc" to type uint32`,
		`invalid operation: 4294967295 + "abc" (mismatched types uint32 and string)`,
	)

}

// Test Uint32 + Nil
func TestCheckBinaryTypedExprUint32AddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + nil`, env,
		`cannot convert nil to type uint32`,
	)

}

// Test Uint32 + Int8
func TestCheckBinaryTypedExprUint32AddInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + int8(0x7f)`, env,
		`invalid operation: 4294967295 + 127 (mismatched types uint32 and int8)`,
	)

}

// Test Uint32 + Int16
func TestCheckBinaryTypedExprUint32AddInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + int16(0x7fff)`, env,
		`invalid operation: 4294967295 + 32767 (mismatched types uint32 and int16)`,
	)

}

// Test Uint32 + Int32
func TestCheckBinaryTypedExprUint32AddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + int32(0x7fffffff)`, env,
		`invalid operation: 4294967295 + 2147483647 (mismatched types uint32 and int32)`,
	)

}

// Test Uint32 + Int64
func TestCheckBinaryTypedExprUint32AddInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4294967295 + 9223372036854775807 (mismatched types uint32 and int64)`,
	)

}

// Test Uint32 + Uint8
func TestCheckBinaryTypedExprUint32AddUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + uint8(0xff)`, env,
		`invalid operation: 4294967295 + 255 (mismatched types uint32 and uint8)`,
	)

}

// Test Uint32 + Uint16
func TestCheckBinaryTypedExprUint32AddUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + uint16(0xffff)`, env,
		`invalid operation: 4294967295 + 65535 (mismatched types uint32 and uint16)`,
	)

}

// Test Uint32 + Uint32
func TestCheckBinaryTypedExprUint32AddUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + uint32(0xffffffff)`, env,
		`constant 8589934590 overflows uint32`,
	)

}

// Test Uint32 + Uint64
func TestCheckBinaryTypedExprUint32AddUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4294967295 + 18446744073709551615 (mismatched types uint32 and uint64)`,
	)

}

// Test Uint32 + Float32
func TestCheckBinaryTypedExprUint32AddFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + float32(0xffffffff)`, env,
		`invalid operation: 4294967295 + 4.29497e+09 (mismatched types uint32 and float32)`,
	)

}

// Test Uint32 + Float64
func TestCheckBinaryTypedExprUint32AddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + float64(0xffffffff)`, env,
		`invalid operation: 4294967295 + 4.29497e+09 (mismatched types uint32 and float64)`,
	)

}

// Test Uint32 + Complex64
func TestCheckBinaryTypedExprUint32AddComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4294967295 + (4.29497e+09+4.29497e+09i) (mismatched types uint32 and complex64)`,
	)

}

// Test Uint32 + Complex128
func TestCheckBinaryTypedExprUint32AddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4294967295 + (4.29497e+09+4.29497e+09i) (mismatched types uint32 and complex128)`,
	)

}

// Test Uint32 + Rune32
func TestCheckBinaryTypedExprUint32AddRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + rune(0x7fffffff)`, env,
		`invalid operation: 4294967295 + rune(2147483647) (mismatched types uint32 and rune)`,
	)

}

// Test Uint32 + StringT
func TestCheckBinaryTypedExprUint32AddStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + string("abc")`, env,
		`invalid operation: 4294967295 + "abc" (mismatched types uint32 and string)`,
	)

}

// Test Uint32 + BoolT
func TestCheckBinaryTypedExprUint32AddBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) + bool(true)`, env,
		`invalid operation: 4294967295 + true (mismatched types uint32 and bool)`,
	)

}

// Test Uint32 - Int
func TestCheckBinaryTypedExprUint32SubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) - 4`, env, uint32(0xffffffff) - 4, reflect.TypeOf(uint32(0xffffffff) - 4))
}

// Test Uint32 - Rune
func TestCheckBinaryTypedExprUint32SubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) - '@'`, env, uint32(0xffffffff) - '@', reflect.TypeOf(uint32(0xffffffff) - '@'))
}

// Test Uint32 - Float
func TestCheckBinaryTypedExprUint32SubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) - 2.0`, env, uint32(0xffffffff) - 2.0, reflect.TypeOf(uint32(0xffffffff) - 2.0))
}

// Test Uint32 - Complex
func TestCheckBinaryTypedExprUint32SubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint32 - Bool
func TestCheckBinaryTypedExprUint32SubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - true`, env,
		`cannot convert true to type uint32`,
		`invalid operation: 4294967295 - true (mismatched types uint32 and bool)`,
	)

}

// Test Uint32 - String
func TestCheckBinaryTypedExprUint32SubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - "abc"`, env,
		`cannot convert "abc" to type uint32`,
		`invalid operation: 4294967295 - "abc" (mismatched types uint32 and string)`,
	)

}

// Test Uint32 - Nil
func TestCheckBinaryTypedExprUint32SubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - nil`, env,
		`cannot convert nil to type uint32`,
	)

}

// Test Uint32 - Int8
func TestCheckBinaryTypedExprUint32SubInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - int8(0x7f)`, env,
		`invalid operation: 4294967295 - 127 (mismatched types uint32 and int8)`,
	)

}

// Test Uint32 - Int16
func TestCheckBinaryTypedExprUint32SubInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - int16(0x7fff)`, env,
		`invalid operation: 4294967295 - 32767 (mismatched types uint32 and int16)`,
	)

}

// Test Uint32 - Int32
func TestCheckBinaryTypedExprUint32SubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - int32(0x7fffffff)`, env,
		`invalid operation: 4294967295 - 2147483647 (mismatched types uint32 and int32)`,
	)

}

// Test Uint32 - Int64
func TestCheckBinaryTypedExprUint32SubInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4294967295 - 9223372036854775807 (mismatched types uint32 and int64)`,
	)

}

// Test Uint32 - Uint8
func TestCheckBinaryTypedExprUint32SubUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - uint8(0xff)`, env,
		`invalid operation: 4294967295 - 255 (mismatched types uint32 and uint8)`,
	)

}

// Test Uint32 - Uint16
func TestCheckBinaryTypedExprUint32SubUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - uint16(0xffff)`, env,
		`invalid operation: 4294967295 - 65535 (mismatched types uint32 and uint16)`,
	)

}

// Test Uint32 - Uint32
func TestCheckBinaryTypedExprUint32SubUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) - uint32(0xffffffff)`, env, uint32(0xffffffff) - uint32(0xffffffff), reflect.TypeOf(uint32(0xffffffff) - uint32(0xffffffff)))
}

// Test Uint32 - Uint64
func TestCheckBinaryTypedExprUint32SubUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4294967295 - 18446744073709551615 (mismatched types uint32 and uint64)`,
	)

}

// Test Uint32 - Float32
func TestCheckBinaryTypedExprUint32SubFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - float32(0xffffffff)`, env,
		`invalid operation: 4294967295 - 4.29497e+09 (mismatched types uint32 and float32)`,
	)

}

// Test Uint32 - Float64
func TestCheckBinaryTypedExprUint32SubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - float64(0xffffffff)`, env,
		`invalid operation: 4294967295 - 4.29497e+09 (mismatched types uint32 and float64)`,
	)

}

// Test Uint32 - Complex64
func TestCheckBinaryTypedExprUint32SubComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4294967295 - (4.29497e+09+4.29497e+09i) (mismatched types uint32 and complex64)`,
	)

}

// Test Uint32 - Complex128
func TestCheckBinaryTypedExprUint32SubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4294967295 - (4.29497e+09+4.29497e+09i) (mismatched types uint32 and complex128)`,
	)

}

// Test Uint32 - Rune32
func TestCheckBinaryTypedExprUint32SubRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - rune(0x7fffffff)`, env,
		`invalid operation: 4294967295 - rune(2147483647) (mismatched types uint32 and rune)`,
	)

}

// Test Uint32 - StringT
func TestCheckBinaryTypedExprUint32SubStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - string("abc")`, env,
		`invalid operation: 4294967295 - "abc" (mismatched types uint32 and string)`,
	)

}

// Test Uint32 - BoolT
func TestCheckBinaryTypedExprUint32SubBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) - bool(true)`, env,
		`invalid operation: 4294967295 - true (mismatched types uint32 and bool)`,
	)

}

// Test Uint32 & Int
func TestCheckBinaryTypedExprUint32AndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) & 4`, env, uint32(0xffffffff) & 4, reflect.TypeOf(uint32(0xffffffff) & 4))
}

// Test Uint32 & Rune
func TestCheckBinaryTypedExprUint32AndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) & '@'`, env, uint32(0xffffffff) & '@', reflect.TypeOf(uint32(0xffffffff) & '@'))
}

// Test Uint32 & Float
func TestCheckBinaryTypedExprUint32AndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) & 2.0`, env, uint32(0xffffffff) & 2.0, reflect.TypeOf(uint32(0xffffffff) & 2.0))
}

// Test Uint32 & Complex
func TestCheckBinaryTypedExprUint32AndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint32 & Bool
func TestCheckBinaryTypedExprUint32AndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & true`, env,
		`cannot convert true to type uint32`,
		`invalid operation: 4294967295 & true (mismatched types uint32 and bool)`,
	)

}

// Test Uint32 & String
func TestCheckBinaryTypedExprUint32AndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & "abc"`, env,
		`cannot convert "abc" to type uint32`,
		`invalid operation: 4294967295 & "abc" (mismatched types uint32 and string)`,
	)

}

// Test Uint32 & Nil
func TestCheckBinaryTypedExprUint32AndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & nil`, env,
		`cannot convert nil to type uint32`,
	)

}

// Test Uint32 & Int8
func TestCheckBinaryTypedExprUint32AndInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & int8(0x7f)`, env,
		`invalid operation: 4294967295 & 127 (mismatched types uint32 and int8)`,
	)

}

// Test Uint32 & Int16
func TestCheckBinaryTypedExprUint32AndInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & int16(0x7fff)`, env,
		`invalid operation: 4294967295 & 32767 (mismatched types uint32 and int16)`,
	)

}

// Test Uint32 & Int32
func TestCheckBinaryTypedExprUint32AndInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & int32(0x7fffffff)`, env,
		`invalid operation: 4294967295 & 2147483647 (mismatched types uint32 and int32)`,
	)

}

// Test Uint32 & Int64
func TestCheckBinaryTypedExprUint32AndInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4294967295 & 9223372036854775807 (mismatched types uint32 and int64)`,
	)

}

// Test Uint32 & Uint8
func TestCheckBinaryTypedExprUint32AndUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & uint8(0xff)`, env,
		`invalid operation: 4294967295 & 255 (mismatched types uint32 and uint8)`,
	)

}

// Test Uint32 & Uint16
func TestCheckBinaryTypedExprUint32AndUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & uint16(0xffff)`, env,
		`invalid operation: 4294967295 & 65535 (mismatched types uint32 and uint16)`,
	)

}

// Test Uint32 & Uint32
func TestCheckBinaryTypedExprUint32AndUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) & uint32(0xffffffff)`, env, uint32(0xffffffff) & uint32(0xffffffff), reflect.TypeOf(uint32(0xffffffff) & uint32(0xffffffff)))
}

// Test Uint32 & Uint64
func TestCheckBinaryTypedExprUint32AndUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4294967295 & 18446744073709551615 (mismatched types uint32 and uint64)`,
	)

}

// Test Uint32 & Float32
func TestCheckBinaryTypedExprUint32AndFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & float32(0xffffffff)`, env,
		`invalid operation: 4294967295 & 4.29497e+09 (mismatched types uint32 and float32)`,
	)

}

// Test Uint32 & Float64
func TestCheckBinaryTypedExprUint32AndFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & float64(0xffffffff)`, env,
		`invalid operation: 4294967295 & 4.29497e+09 (mismatched types uint32 and float64)`,
	)

}

// Test Uint32 & Complex64
func TestCheckBinaryTypedExprUint32AndComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4294967295 & (4.29497e+09+4.29497e+09i) (mismatched types uint32 and complex64)`,
	)

}

// Test Uint32 & Complex128
func TestCheckBinaryTypedExprUint32AndComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4294967295 & (4.29497e+09+4.29497e+09i) (mismatched types uint32 and complex128)`,
	)

}

// Test Uint32 & Rune32
func TestCheckBinaryTypedExprUint32AndRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & rune(0x7fffffff)`, env,
		`invalid operation: 4294967295 & rune(2147483647) (mismatched types uint32 and rune)`,
	)

}

// Test Uint32 & StringT
func TestCheckBinaryTypedExprUint32AndStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & string("abc")`, env,
		`invalid operation: 4294967295 & "abc" (mismatched types uint32 and string)`,
	)

}

// Test Uint32 & BoolT
func TestCheckBinaryTypedExprUint32AndBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) & bool(true)`, env,
		`invalid operation: 4294967295 & true (mismatched types uint32 and bool)`,
	)

}

// Test Uint32 % Int
func TestCheckBinaryTypedExprUint32RemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) % 4`, env, uint32(0xffffffff) % 4, reflect.TypeOf(uint32(0xffffffff) % 4))
}

// Test Uint32 % Rune
func TestCheckBinaryTypedExprUint32RemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) % '@'`, env, uint32(0xffffffff) % '@', reflect.TypeOf(uint32(0xffffffff) % '@'))
}

// Test Uint32 % Float
func TestCheckBinaryTypedExprUint32RemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) % 2.0`, env, uint32(0xffffffff) % 2.0, reflect.TypeOf(uint32(0xffffffff) % 2.0))
}

// Test Uint32 % Complex
func TestCheckBinaryTypedExprUint32RemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % 8.0i`, env,
		`constant 0+8i truncated to real`,
		`division by zero`,
	)

}

// Test Uint32 % Bool
func TestCheckBinaryTypedExprUint32RemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % true`, env,
		`cannot convert true to type uint32`,
		`invalid operation: 4294967295 % true (mismatched types uint32 and bool)`,
	)

}

// Test Uint32 % String
func TestCheckBinaryTypedExprUint32RemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % "abc"`, env,
		`cannot convert "abc" to type uint32`,
		`invalid operation: 4294967295 % "abc" (mismatched types uint32 and string)`,
	)

}

// Test Uint32 % Nil
func TestCheckBinaryTypedExprUint32RemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % nil`, env,
		`cannot convert nil to type uint32`,
	)

}

// Test Uint32 % Int8
func TestCheckBinaryTypedExprUint32RemInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % int8(0x7f)`, env,
		`invalid operation: 4294967295 % 127 (mismatched types uint32 and int8)`,
	)

}

// Test Uint32 % Int16
func TestCheckBinaryTypedExprUint32RemInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % int16(0x7fff)`, env,
		`invalid operation: 4294967295 % 32767 (mismatched types uint32 and int16)`,
	)

}

// Test Uint32 % Int32
func TestCheckBinaryTypedExprUint32RemInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % int32(0x7fffffff)`, env,
		`invalid operation: 4294967295 % 2147483647 (mismatched types uint32 and int32)`,
	)

}

// Test Uint32 % Int64
func TestCheckBinaryTypedExprUint32RemInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4294967295 % 9223372036854775807 (mismatched types uint32 and int64)`,
	)

}

// Test Uint32 % Uint8
func TestCheckBinaryTypedExprUint32RemUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % uint8(0xff)`, env,
		`invalid operation: 4294967295 % 255 (mismatched types uint32 and uint8)`,
	)

}

// Test Uint32 % Uint16
func TestCheckBinaryTypedExprUint32RemUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % uint16(0xffff)`, env,
		`invalid operation: 4294967295 % 65535 (mismatched types uint32 and uint16)`,
	)

}

// Test Uint32 % Uint32
func TestCheckBinaryTypedExprUint32RemUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) % uint32(0xffffffff)`, env, uint32(0xffffffff) % uint32(0xffffffff), reflect.TypeOf(uint32(0xffffffff) % uint32(0xffffffff)))
}

// Test Uint32 % Uint64
func TestCheckBinaryTypedExprUint32RemUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4294967295 % 18446744073709551615 (mismatched types uint32 and uint64)`,
	)

}

// Test Uint32 % Float32
func TestCheckBinaryTypedExprUint32RemFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % float32(0xffffffff)`, env,
		`invalid operation: 4294967295 % 4.29497e+09 (mismatched types uint32 and float32)`,
	)

}

// Test Uint32 % Float64
func TestCheckBinaryTypedExprUint32RemFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % float64(0xffffffff)`, env,
		`invalid operation: 4294967295 % 4.29497e+09 (mismatched types uint32 and float64)`,
	)

}

// Test Uint32 % Complex64
func TestCheckBinaryTypedExprUint32RemComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4294967295 % (4.29497e+09+4.29497e+09i) (mismatched types uint32 and complex64)`,
	)

}

// Test Uint32 % Complex128
func TestCheckBinaryTypedExprUint32RemComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4294967295 % (4.29497e+09+4.29497e+09i) (mismatched types uint32 and complex128)`,
	)

}

// Test Uint32 % Rune32
func TestCheckBinaryTypedExprUint32RemRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % rune(0x7fffffff)`, env,
		`invalid operation: 4294967295 % rune(2147483647) (mismatched types uint32 and rune)`,
	)

}

// Test Uint32 % StringT
func TestCheckBinaryTypedExprUint32RemStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % string("abc")`, env,
		`invalid operation: 4294967295 % "abc" (mismatched types uint32 and string)`,
	)

}

// Test Uint32 % BoolT
func TestCheckBinaryTypedExprUint32RemBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) % bool(true)`, env,
		`invalid operation: 4294967295 % true (mismatched types uint32 and bool)`,
	)

}

// Test Uint32 == Int
func TestCheckBinaryTypedExprUint32EqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) == 4`, env, uint32(0xffffffff) == 4, reflect.TypeOf(uint32(0xffffffff) == 4))
}

// Test Uint32 == Rune
func TestCheckBinaryTypedExprUint32EqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) == '@'`, env, uint32(0xffffffff) == '@', reflect.TypeOf(uint32(0xffffffff) == '@'))
}

// Test Uint32 == Float
func TestCheckBinaryTypedExprUint32EqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) == 2.0`, env, uint32(0xffffffff) == 2.0, reflect.TypeOf(uint32(0xffffffff) == 2.0))
}

// Test Uint32 == Complex
func TestCheckBinaryTypedExprUint32EqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint32 == Bool
func TestCheckBinaryTypedExprUint32EqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == true`, env,
		`cannot convert true to type uint32`,
		`invalid operation: 4294967295 == true (mismatched types uint32 and bool)`,
	)

}

// Test Uint32 == String
func TestCheckBinaryTypedExprUint32EqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == "abc"`, env,
		`cannot convert "abc" to type uint32`,
		`invalid operation: 4294967295 == "abc" (mismatched types uint32 and string)`,
	)

}

// Test Uint32 == Nil
func TestCheckBinaryTypedExprUint32EqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == nil`, env,
		`cannot convert nil to type uint32`,
	)

}

// Test Uint32 == Int8
func TestCheckBinaryTypedExprUint32EqlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == int8(0x7f)`, env,
		`invalid operation: 4294967295 == 127 (mismatched types uint32 and int8)`,
	)

}

// Test Uint32 == Int16
func TestCheckBinaryTypedExprUint32EqlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == int16(0x7fff)`, env,
		`invalid operation: 4294967295 == 32767 (mismatched types uint32 and int16)`,
	)

}

// Test Uint32 == Int32
func TestCheckBinaryTypedExprUint32EqlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == int32(0x7fffffff)`, env,
		`invalid operation: 4294967295 == 2147483647 (mismatched types uint32 and int32)`,
	)

}

// Test Uint32 == Int64
func TestCheckBinaryTypedExprUint32EqlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4294967295 == 9223372036854775807 (mismatched types uint32 and int64)`,
	)

}

// Test Uint32 == Uint8
func TestCheckBinaryTypedExprUint32EqlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == uint8(0xff)`, env,
		`invalid operation: 4294967295 == 255 (mismatched types uint32 and uint8)`,
	)

}

// Test Uint32 == Uint16
func TestCheckBinaryTypedExprUint32EqlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == uint16(0xffff)`, env,
		`invalid operation: 4294967295 == 65535 (mismatched types uint32 and uint16)`,
	)

}

// Test Uint32 == Uint32
func TestCheckBinaryTypedExprUint32EqlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) == uint32(0xffffffff)`, env, uint32(0xffffffff) == uint32(0xffffffff), reflect.TypeOf(uint32(0xffffffff) == uint32(0xffffffff)))
}

// Test Uint32 == Uint64
func TestCheckBinaryTypedExprUint32EqlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4294967295 == 18446744073709551615 (mismatched types uint32 and uint64)`,
	)

}

// Test Uint32 == Float32
func TestCheckBinaryTypedExprUint32EqlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == float32(0xffffffff)`, env,
		`invalid operation: 4294967295 == 4.29497e+09 (mismatched types uint32 and float32)`,
	)

}

// Test Uint32 == Float64
func TestCheckBinaryTypedExprUint32EqlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == float64(0xffffffff)`, env,
		`invalid operation: 4294967295 == 4.29497e+09 (mismatched types uint32 and float64)`,
	)

}

// Test Uint32 == Complex64
func TestCheckBinaryTypedExprUint32EqlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4294967295 == (4.29497e+09+4.29497e+09i) (mismatched types uint32 and complex64)`,
	)

}

// Test Uint32 == Complex128
func TestCheckBinaryTypedExprUint32EqlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4294967295 == (4.29497e+09+4.29497e+09i) (mismatched types uint32 and complex128)`,
	)

}

// Test Uint32 == Rune32
func TestCheckBinaryTypedExprUint32EqlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == rune(0x7fffffff)`, env,
		`invalid operation: 4294967295 == rune(2147483647) (mismatched types uint32 and rune)`,
	)

}

// Test Uint32 == StringT
func TestCheckBinaryTypedExprUint32EqlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == string("abc")`, env,
		`invalid operation: 4294967295 == "abc" (mismatched types uint32 and string)`,
	)

}

// Test Uint32 == BoolT
func TestCheckBinaryTypedExprUint32EqlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) == bool(true)`, env,
		`invalid operation: 4294967295 == true (mismatched types uint32 and bool)`,
	)

}

// Test Uint32 > Int
func TestCheckBinaryTypedExprUint32GtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) > 4`, env, uint32(0xffffffff) > 4, reflect.TypeOf(uint32(0xffffffff) > 4))
}

// Test Uint32 > Rune
func TestCheckBinaryTypedExprUint32GtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) > '@'`, env, uint32(0xffffffff) > '@', reflect.TypeOf(uint32(0xffffffff) > '@'))
}

// Test Uint32 > Float
func TestCheckBinaryTypedExprUint32GtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) > 2.0`, env, uint32(0xffffffff) > 2.0, reflect.TypeOf(uint32(0xffffffff) > 2.0))
}

// Test Uint32 > Complex
func TestCheckBinaryTypedExprUint32GtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint32 > Bool
func TestCheckBinaryTypedExprUint32GtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > true`, env,
		`cannot convert true to type uint32`,
		`invalid operation: 4294967295 > true (mismatched types uint32 and bool)`,
	)

}

// Test Uint32 > String
func TestCheckBinaryTypedExprUint32GtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > "abc"`, env,
		`cannot convert "abc" to type uint32`,
		`invalid operation: 4294967295 > "abc" (mismatched types uint32 and string)`,
	)

}

// Test Uint32 > Nil
func TestCheckBinaryTypedExprUint32GtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > nil`, env,
		`cannot convert nil to type uint32`,
	)

}

// Test Uint32 > Int8
func TestCheckBinaryTypedExprUint32GtrInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > int8(0x7f)`, env,
		`invalid operation: 4294967295 > 127 (mismatched types uint32 and int8)`,
	)

}

// Test Uint32 > Int16
func TestCheckBinaryTypedExprUint32GtrInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > int16(0x7fff)`, env,
		`invalid operation: 4294967295 > 32767 (mismatched types uint32 and int16)`,
	)

}

// Test Uint32 > Int32
func TestCheckBinaryTypedExprUint32GtrInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > int32(0x7fffffff)`, env,
		`invalid operation: 4294967295 > 2147483647 (mismatched types uint32 and int32)`,
	)

}

// Test Uint32 > Int64
func TestCheckBinaryTypedExprUint32GtrInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4294967295 > 9223372036854775807 (mismatched types uint32 and int64)`,
	)

}

// Test Uint32 > Uint8
func TestCheckBinaryTypedExprUint32GtrUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > uint8(0xff)`, env,
		`invalid operation: 4294967295 > 255 (mismatched types uint32 and uint8)`,
	)

}

// Test Uint32 > Uint16
func TestCheckBinaryTypedExprUint32GtrUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > uint16(0xffff)`, env,
		`invalid operation: 4294967295 > 65535 (mismatched types uint32 and uint16)`,
	)

}

// Test Uint32 > Uint32
func TestCheckBinaryTypedExprUint32GtrUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff) > uint32(0xffffffff)`, env, uint32(0xffffffff) > uint32(0xffffffff), reflect.TypeOf(uint32(0xffffffff) > uint32(0xffffffff)))
}

// Test Uint32 > Uint64
func TestCheckBinaryTypedExprUint32GtrUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4294967295 > 18446744073709551615 (mismatched types uint32 and uint64)`,
	)

}

// Test Uint32 > Float32
func TestCheckBinaryTypedExprUint32GtrFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > float32(0xffffffff)`, env,
		`invalid operation: 4294967295 > 4.29497e+09 (mismatched types uint32 and float32)`,
	)

}

// Test Uint32 > Float64
func TestCheckBinaryTypedExprUint32GtrFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > float64(0xffffffff)`, env,
		`invalid operation: 4294967295 > 4.29497e+09 (mismatched types uint32 and float64)`,
	)

}

// Test Uint32 > Complex64
func TestCheckBinaryTypedExprUint32GtrComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4294967295 > (4.29497e+09+4.29497e+09i) (mismatched types uint32 and complex64)`,
	)

}

// Test Uint32 > Complex128
func TestCheckBinaryTypedExprUint32GtrComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4294967295 > (4.29497e+09+4.29497e+09i) (mismatched types uint32 and complex128)`,
	)

}

// Test Uint32 > Rune32
func TestCheckBinaryTypedExprUint32GtrRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > rune(0x7fffffff)`, env,
		`invalid operation: 4294967295 > rune(2147483647) (mismatched types uint32 and rune)`,
	)

}

// Test Uint32 > StringT
func TestCheckBinaryTypedExprUint32GtrStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > string("abc")`, env,
		`invalid operation: 4294967295 > "abc" (mismatched types uint32 and string)`,
	)

}

// Test Uint32 > BoolT
func TestCheckBinaryTypedExprUint32GtrBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) > bool(true)`, env,
		`invalid operation: 4294967295 > true (mismatched types uint32 and bool)`,
	)

}

// Test Uint32 << Int
func TestCheckBinaryTypedExprUint32ShlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << 4`, env,
		`constant 68719476720 overflows uint32`,
	)

}

// Test Uint32 << Rune
func TestCheckBinaryTypedExprUint32ShlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << '@'`, env,
		`constant 79228162495817593519834398720 overflows uint32`,
	)

}

// Test Uint32 << Float
func TestCheckBinaryTypedExprUint32ShlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << 2.0`, env,
		`constant 17179869180 overflows uint32`,
	)

}

// Test Uint32 << Complex
func TestCheckBinaryTypedExprUint32ShlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint32 << Bool
func TestCheckBinaryTypedExprUint32ShlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << true`, env,
		`invalid operation: 4294967295 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Uint32 << String
func TestCheckBinaryTypedExprUint32ShlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: 4294967295 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Uint32 << Nil
func TestCheckBinaryTypedExprUint32ShlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Uint32 << Int8
func TestCheckBinaryTypedExprUint32ShlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << int8(0x7f)`, env,
		`invalid operation: 4294967295 << 127 (shift count type int8, must be unsigned integer)`,
	)

}

// Test Uint32 << Int16
func TestCheckBinaryTypedExprUint32ShlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << int16(0x7fff)`, env,
		`invalid operation: 4294967295 << 32767 (shift count type int16, must be unsigned integer)`,
	)

}

// Test Uint32 << Int32
func TestCheckBinaryTypedExprUint32ShlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << int32(0x7fffffff)`, env,
		`invalid operation: 4294967295 << 2147483647 (shift count type int32, must be unsigned integer)`,
	)

}

// Test Uint32 << Int64
func TestCheckBinaryTypedExprUint32ShlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4294967295 << 9223372036854775807 (shift count type int64, must be unsigned integer)`,
	)

}

// Test Uint32 << Uint8
func TestCheckBinaryTypedExprUint32ShlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << uint8(0xff)`, env,
		`constant 248661618146997276459033026361624927545884121472038866387411706426850956502310122946560 overflows uint32`,
	)

}

// Test Uint32 << Uint16
func TestCheckBinaryTypedExprUint32ShlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << uint16(0xffff)`, env,
		`stupid shift: 65535`,
	)

}

// Test Uint32 << Uint32
func TestCheckBinaryTypedExprUint32ShlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << uint32(0xffffffff)`, env,
		`stupid shift: 4294967295`,
	)

}

// Test Uint32 << Uint64
func TestCheckBinaryTypedExprUint32ShlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << uint64(0xffffffffffffffff)`, env,
		`stupid shift: -1`,
	)

}

// Test Uint32 << Float32
func TestCheckBinaryTypedExprUint32ShlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << float32(0xffffffff)`, env,
		`invalid operation: 4294967295 << 4.29497e+09 (shift count type float32, must be unsigned integer)`,
	)

}

// Test Uint32 << Float64
func TestCheckBinaryTypedExprUint32ShlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << float64(0xffffffff)`, env,
		`invalid operation: 4294967295 << 4.29497e+09 (shift count type float64, must be unsigned integer)`,
	)

}

// Test Uint32 << Complex64
func TestCheckBinaryTypedExprUint32ShlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4294967295 << (4.29497e+09+4.29497e+09i) (shift count type complex64, must be unsigned integer)`,
	)

}

// Test Uint32 << Complex128
func TestCheckBinaryTypedExprUint32ShlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4294967295 << (4.29497e+09+4.29497e+09i) (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Uint32 << Rune32
func TestCheckBinaryTypedExprUint32ShlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << rune(0x7fffffff)`, env,
		`invalid operation: 4294967295 << rune(2147483647) (shift count type rune, must be unsigned integer)`,
	)

}

// Test Uint32 << StringT
func TestCheckBinaryTypedExprUint32ShlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << string("abc")`, env,
		`invalid operation: 4294967295 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Uint32 << BoolT
func TestCheckBinaryTypedExprUint32ShlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffff) << bool(true)`, env,
		`invalid operation: 4294967295 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Uint64 + Int
func TestCheckBinaryTypedExprUint64AddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + 4`, env,
		`constant 18446744073709551619 overflows uint64`,
	)

}

// Test Uint64 + Rune
func TestCheckBinaryTypedExprUint64AddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + '@'`, env,
		`constant 18446744073709551679 overflows uint64`,
	)

}

// Test Uint64 + Float
func TestCheckBinaryTypedExprUint64AddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + 2.0`, env,
		`constant 18446744073709551617 overflows uint64`,
	)

}

// Test Uint64 + Complex
func TestCheckBinaryTypedExprUint64AddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint64 + Bool
func TestCheckBinaryTypedExprUint64AddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + true`, env,
		`cannot convert true to type uint64`,
		`invalid operation: 18446744073709551615 + true (mismatched types uint64 and bool)`,
	)

}

// Test Uint64 + String
func TestCheckBinaryTypedExprUint64AddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + "abc"`, env,
		`cannot convert "abc" to type uint64`,
		`invalid operation: 18446744073709551615 + "abc" (mismatched types uint64 and string)`,
	)

}

// Test Uint64 + Nil
func TestCheckBinaryTypedExprUint64AddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + nil`, env,
		`cannot convert nil to type uint64`,
	)

}

// Test Uint64 + Int8
func TestCheckBinaryTypedExprUint64AddInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + int8(0x7f)`, env,
		`invalid operation: 18446744073709551615 + 127 (mismatched types uint64 and int8)`,
	)

}

// Test Uint64 + Int16
func TestCheckBinaryTypedExprUint64AddInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + int16(0x7fff)`, env,
		`invalid operation: 18446744073709551615 + 32767 (mismatched types uint64 and int16)`,
	)

}

// Test Uint64 + Int32
func TestCheckBinaryTypedExprUint64AddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + int32(0x7fffffff)`, env,
		`invalid operation: 18446744073709551615 + 2147483647 (mismatched types uint64 and int32)`,
	)

}

// Test Uint64 + Int64
func TestCheckBinaryTypedExprUint64AddInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + int64(0x7fffffffffffffff)`, env,
		`invalid operation: 18446744073709551615 + 9223372036854775807 (mismatched types uint64 and int64)`,
	)

}

// Test Uint64 + Uint8
func TestCheckBinaryTypedExprUint64AddUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + uint8(0xff)`, env,
		`invalid operation: 18446744073709551615 + 255 (mismatched types uint64 and uint8)`,
	)

}

// Test Uint64 + Uint16
func TestCheckBinaryTypedExprUint64AddUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + uint16(0xffff)`, env,
		`invalid operation: 18446744073709551615 + 65535 (mismatched types uint64 and uint16)`,
	)

}

// Test Uint64 + Uint32
func TestCheckBinaryTypedExprUint64AddUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + uint32(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 + 4294967295 (mismatched types uint64 and uint32)`,
	)

}

// Test Uint64 + Uint64
func TestCheckBinaryTypedExprUint64AddUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + uint64(0xffffffffffffffff)`, env,
		`constant 36893488147419103230 overflows uint64`,
	)

}

// Test Uint64 + Float32
func TestCheckBinaryTypedExprUint64AddFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + float32(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 + 4.29497e+09 (mismatched types uint64 and float32)`,
	)

}

// Test Uint64 + Float64
func TestCheckBinaryTypedExprUint64AddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + float64(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 + 4.29497e+09 (mismatched types uint64 and float64)`,
	)

}

// Test Uint64 + Complex64
func TestCheckBinaryTypedExprUint64AddComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 18446744073709551615 + (4.29497e+09+4.29497e+09i) (mismatched types uint64 and complex64)`,
	)

}

// Test Uint64 + Complex128
func TestCheckBinaryTypedExprUint64AddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 18446744073709551615 + (4.29497e+09+4.29497e+09i) (mismatched types uint64 and complex128)`,
	)

}

// Test Uint64 + Rune32
func TestCheckBinaryTypedExprUint64AddRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + rune(0x7fffffff)`, env,
		`invalid operation: 18446744073709551615 + rune(2147483647) (mismatched types uint64 and rune)`,
	)

}

// Test Uint64 + StringT
func TestCheckBinaryTypedExprUint64AddStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + string("abc")`, env,
		`invalid operation: 18446744073709551615 + "abc" (mismatched types uint64 and string)`,
	)

}

// Test Uint64 + BoolT
func TestCheckBinaryTypedExprUint64AddBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) + bool(true)`, env,
		`invalid operation: 18446744073709551615 + true (mismatched types uint64 and bool)`,
	)

}

// Test Uint64 - Int
func TestCheckBinaryTypedExprUint64SubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) - 4`, env, uint64(0xffffffffffffffff) - 4, reflect.TypeOf(uint64(0xffffffffffffffff) - 4))
}

// Test Uint64 - Rune
func TestCheckBinaryTypedExprUint64SubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) - '@'`, env, uint64(0xffffffffffffffff) - '@', reflect.TypeOf(uint64(0xffffffffffffffff) - '@'))
}

// Test Uint64 - Float
func TestCheckBinaryTypedExprUint64SubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) - 2.0`, env, uint64(0xffffffffffffffff) - 2.0, reflect.TypeOf(uint64(0xffffffffffffffff) - 2.0))
}

// Test Uint64 - Complex
func TestCheckBinaryTypedExprUint64SubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint64 - Bool
func TestCheckBinaryTypedExprUint64SubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - true`, env,
		`cannot convert true to type uint64`,
		`invalid operation: 18446744073709551615 - true (mismatched types uint64 and bool)`,
	)

}

// Test Uint64 - String
func TestCheckBinaryTypedExprUint64SubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - "abc"`, env,
		`cannot convert "abc" to type uint64`,
		`invalid operation: 18446744073709551615 - "abc" (mismatched types uint64 and string)`,
	)

}

// Test Uint64 - Nil
func TestCheckBinaryTypedExprUint64SubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - nil`, env,
		`cannot convert nil to type uint64`,
	)

}

// Test Uint64 - Int8
func TestCheckBinaryTypedExprUint64SubInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - int8(0x7f)`, env,
		`invalid operation: 18446744073709551615 - 127 (mismatched types uint64 and int8)`,
	)

}

// Test Uint64 - Int16
func TestCheckBinaryTypedExprUint64SubInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - int16(0x7fff)`, env,
		`invalid operation: 18446744073709551615 - 32767 (mismatched types uint64 and int16)`,
	)

}

// Test Uint64 - Int32
func TestCheckBinaryTypedExprUint64SubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - int32(0x7fffffff)`, env,
		`invalid operation: 18446744073709551615 - 2147483647 (mismatched types uint64 and int32)`,
	)

}

// Test Uint64 - Int64
func TestCheckBinaryTypedExprUint64SubInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - int64(0x7fffffffffffffff)`, env,
		`invalid operation: 18446744073709551615 - 9223372036854775807 (mismatched types uint64 and int64)`,
	)

}

// Test Uint64 - Uint8
func TestCheckBinaryTypedExprUint64SubUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - uint8(0xff)`, env,
		`invalid operation: 18446744073709551615 - 255 (mismatched types uint64 and uint8)`,
	)

}

// Test Uint64 - Uint16
func TestCheckBinaryTypedExprUint64SubUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - uint16(0xffff)`, env,
		`invalid operation: 18446744073709551615 - 65535 (mismatched types uint64 and uint16)`,
	)

}

// Test Uint64 - Uint32
func TestCheckBinaryTypedExprUint64SubUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - uint32(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 - 4294967295 (mismatched types uint64 and uint32)`,
	)

}

// Test Uint64 - Uint64
func TestCheckBinaryTypedExprUint64SubUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) - uint64(0xffffffffffffffff)`, env, uint64(0xffffffffffffffff) - uint64(0xffffffffffffffff), reflect.TypeOf(uint64(0xffffffffffffffff) - uint64(0xffffffffffffffff)))
}

// Test Uint64 - Float32
func TestCheckBinaryTypedExprUint64SubFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - float32(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 - 4.29497e+09 (mismatched types uint64 and float32)`,
	)

}

// Test Uint64 - Float64
func TestCheckBinaryTypedExprUint64SubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - float64(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 - 4.29497e+09 (mismatched types uint64 and float64)`,
	)

}

// Test Uint64 - Complex64
func TestCheckBinaryTypedExprUint64SubComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 18446744073709551615 - (4.29497e+09+4.29497e+09i) (mismatched types uint64 and complex64)`,
	)

}

// Test Uint64 - Complex128
func TestCheckBinaryTypedExprUint64SubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 18446744073709551615 - (4.29497e+09+4.29497e+09i) (mismatched types uint64 and complex128)`,
	)

}

// Test Uint64 - Rune32
func TestCheckBinaryTypedExprUint64SubRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - rune(0x7fffffff)`, env,
		`invalid operation: 18446744073709551615 - rune(2147483647) (mismatched types uint64 and rune)`,
	)

}

// Test Uint64 - StringT
func TestCheckBinaryTypedExprUint64SubStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - string("abc")`, env,
		`invalid operation: 18446744073709551615 - "abc" (mismatched types uint64 and string)`,
	)

}

// Test Uint64 - BoolT
func TestCheckBinaryTypedExprUint64SubBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) - bool(true)`, env,
		`invalid operation: 18446744073709551615 - true (mismatched types uint64 and bool)`,
	)

}

// Test Uint64 & Int
func TestCheckBinaryTypedExprUint64AndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) & 4`, env, uint64(0xffffffffffffffff) & 4, reflect.TypeOf(uint64(0xffffffffffffffff) & 4))
}

// Test Uint64 & Rune
func TestCheckBinaryTypedExprUint64AndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) & '@'`, env, uint64(0xffffffffffffffff) & '@', reflect.TypeOf(uint64(0xffffffffffffffff) & '@'))
}

// Test Uint64 & Float
func TestCheckBinaryTypedExprUint64AndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) & 2.0`, env, uint64(0xffffffffffffffff) & 2.0, reflect.TypeOf(uint64(0xffffffffffffffff) & 2.0))
}

// Test Uint64 & Complex
func TestCheckBinaryTypedExprUint64AndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint64 & Bool
func TestCheckBinaryTypedExprUint64AndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & true`, env,
		`cannot convert true to type uint64`,
		`invalid operation: 18446744073709551615 & true (mismatched types uint64 and bool)`,
	)

}

// Test Uint64 & String
func TestCheckBinaryTypedExprUint64AndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & "abc"`, env,
		`cannot convert "abc" to type uint64`,
		`invalid operation: 18446744073709551615 & "abc" (mismatched types uint64 and string)`,
	)

}

// Test Uint64 & Nil
func TestCheckBinaryTypedExprUint64AndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & nil`, env,
		`cannot convert nil to type uint64`,
	)

}

// Test Uint64 & Int8
func TestCheckBinaryTypedExprUint64AndInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & int8(0x7f)`, env,
		`invalid operation: 18446744073709551615 & 127 (mismatched types uint64 and int8)`,
	)

}

// Test Uint64 & Int16
func TestCheckBinaryTypedExprUint64AndInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & int16(0x7fff)`, env,
		`invalid operation: 18446744073709551615 & 32767 (mismatched types uint64 and int16)`,
	)

}

// Test Uint64 & Int32
func TestCheckBinaryTypedExprUint64AndInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & int32(0x7fffffff)`, env,
		`invalid operation: 18446744073709551615 & 2147483647 (mismatched types uint64 and int32)`,
	)

}

// Test Uint64 & Int64
func TestCheckBinaryTypedExprUint64AndInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & int64(0x7fffffffffffffff)`, env,
		`invalid operation: 18446744073709551615 & 9223372036854775807 (mismatched types uint64 and int64)`,
	)

}

// Test Uint64 & Uint8
func TestCheckBinaryTypedExprUint64AndUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & uint8(0xff)`, env,
		`invalid operation: 18446744073709551615 & 255 (mismatched types uint64 and uint8)`,
	)

}

// Test Uint64 & Uint16
func TestCheckBinaryTypedExprUint64AndUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & uint16(0xffff)`, env,
		`invalid operation: 18446744073709551615 & 65535 (mismatched types uint64 and uint16)`,
	)

}

// Test Uint64 & Uint32
func TestCheckBinaryTypedExprUint64AndUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & uint32(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 & 4294967295 (mismatched types uint64 and uint32)`,
	)

}

// Test Uint64 & Uint64
func TestCheckBinaryTypedExprUint64AndUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) & uint64(0xffffffffffffffff)`, env, uint64(0xffffffffffffffff) & uint64(0xffffffffffffffff), reflect.TypeOf(uint64(0xffffffffffffffff) & uint64(0xffffffffffffffff)))
}

// Test Uint64 & Float32
func TestCheckBinaryTypedExprUint64AndFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & float32(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 & 4.29497e+09 (mismatched types uint64 and float32)`,
	)

}

// Test Uint64 & Float64
func TestCheckBinaryTypedExprUint64AndFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & float64(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 & 4.29497e+09 (mismatched types uint64 and float64)`,
	)

}

// Test Uint64 & Complex64
func TestCheckBinaryTypedExprUint64AndComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 18446744073709551615 & (4.29497e+09+4.29497e+09i) (mismatched types uint64 and complex64)`,
	)

}

// Test Uint64 & Complex128
func TestCheckBinaryTypedExprUint64AndComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 18446744073709551615 & (4.29497e+09+4.29497e+09i) (mismatched types uint64 and complex128)`,
	)

}

// Test Uint64 & Rune32
func TestCheckBinaryTypedExprUint64AndRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & rune(0x7fffffff)`, env,
		`invalid operation: 18446744073709551615 & rune(2147483647) (mismatched types uint64 and rune)`,
	)

}

// Test Uint64 & StringT
func TestCheckBinaryTypedExprUint64AndStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & string("abc")`, env,
		`invalid operation: 18446744073709551615 & "abc" (mismatched types uint64 and string)`,
	)

}

// Test Uint64 & BoolT
func TestCheckBinaryTypedExprUint64AndBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) & bool(true)`, env,
		`invalid operation: 18446744073709551615 & true (mismatched types uint64 and bool)`,
	)

}

// Test Uint64 % Int
func TestCheckBinaryTypedExprUint64RemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) % 4`, env, uint64(0xffffffffffffffff) % 4, reflect.TypeOf(uint64(0xffffffffffffffff) % 4))
}

// Test Uint64 % Rune
func TestCheckBinaryTypedExprUint64RemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) % '@'`, env, uint64(0xffffffffffffffff) % '@', reflect.TypeOf(uint64(0xffffffffffffffff) % '@'))
}

// Test Uint64 % Float
func TestCheckBinaryTypedExprUint64RemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) % 2.0`, env, uint64(0xffffffffffffffff) % 2.0, reflect.TypeOf(uint64(0xffffffffffffffff) % 2.0))
}

// Test Uint64 % Complex
func TestCheckBinaryTypedExprUint64RemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % 8.0i`, env,
		`constant 0+8i truncated to real`,
		`division by zero`,
	)

}

// Test Uint64 % Bool
func TestCheckBinaryTypedExprUint64RemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % true`, env,
		`cannot convert true to type uint64`,
		`invalid operation: 18446744073709551615 % true (mismatched types uint64 and bool)`,
	)

}

// Test Uint64 % String
func TestCheckBinaryTypedExprUint64RemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % "abc"`, env,
		`cannot convert "abc" to type uint64`,
		`invalid operation: 18446744073709551615 % "abc" (mismatched types uint64 and string)`,
	)

}

// Test Uint64 % Nil
func TestCheckBinaryTypedExprUint64RemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % nil`, env,
		`cannot convert nil to type uint64`,
	)

}

// Test Uint64 % Int8
func TestCheckBinaryTypedExprUint64RemInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % int8(0x7f)`, env,
		`invalid operation: 18446744073709551615 % 127 (mismatched types uint64 and int8)`,
	)

}

// Test Uint64 % Int16
func TestCheckBinaryTypedExprUint64RemInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % int16(0x7fff)`, env,
		`invalid operation: 18446744073709551615 % 32767 (mismatched types uint64 and int16)`,
	)

}

// Test Uint64 % Int32
func TestCheckBinaryTypedExprUint64RemInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % int32(0x7fffffff)`, env,
		`invalid operation: 18446744073709551615 % 2147483647 (mismatched types uint64 and int32)`,
	)

}

// Test Uint64 % Int64
func TestCheckBinaryTypedExprUint64RemInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % int64(0x7fffffffffffffff)`, env,
		`invalid operation: 18446744073709551615 % 9223372036854775807 (mismatched types uint64 and int64)`,
	)

}

// Test Uint64 % Uint8
func TestCheckBinaryTypedExprUint64RemUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % uint8(0xff)`, env,
		`invalid operation: 18446744073709551615 % 255 (mismatched types uint64 and uint8)`,
	)

}

// Test Uint64 % Uint16
func TestCheckBinaryTypedExprUint64RemUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % uint16(0xffff)`, env,
		`invalid operation: 18446744073709551615 % 65535 (mismatched types uint64 and uint16)`,
	)

}

// Test Uint64 % Uint32
func TestCheckBinaryTypedExprUint64RemUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % uint32(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 % 4294967295 (mismatched types uint64 and uint32)`,
	)

}

// Test Uint64 % Uint64
func TestCheckBinaryTypedExprUint64RemUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) % uint64(0xffffffffffffffff)`, env, uint64(0xffffffffffffffff) % uint64(0xffffffffffffffff), reflect.TypeOf(uint64(0xffffffffffffffff) % uint64(0xffffffffffffffff)))
}

// Test Uint64 % Float32
func TestCheckBinaryTypedExprUint64RemFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % float32(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 % 4.29497e+09 (mismatched types uint64 and float32)`,
	)

}

// Test Uint64 % Float64
func TestCheckBinaryTypedExprUint64RemFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % float64(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 % 4.29497e+09 (mismatched types uint64 and float64)`,
	)

}

// Test Uint64 % Complex64
func TestCheckBinaryTypedExprUint64RemComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 18446744073709551615 % (4.29497e+09+4.29497e+09i) (mismatched types uint64 and complex64)`,
	)

}

// Test Uint64 % Complex128
func TestCheckBinaryTypedExprUint64RemComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 18446744073709551615 % (4.29497e+09+4.29497e+09i) (mismatched types uint64 and complex128)`,
	)

}

// Test Uint64 % Rune32
func TestCheckBinaryTypedExprUint64RemRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % rune(0x7fffffff)`, env,
		`invalid operation: 18446744073709551615 % rune(2147483647) (mismatched types uint64 and rune)`,
	)

}

// Test Uint64 % StringT
func TestCheckBinaryTypedExprUint64RemStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % string("abc")`, env,
		`invalid operation: 18446744073709551615 % "abc" (mismatched types uint64 and string)`,
	)

}

// Test Uint64 % BoolT
func TestCheckBinaryTypedExprUint64RemBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) % bool(true)`, env,
		`invalid operation: 18446744073709551615 % true (mismatched types uint64 and bool)`,
	)

}

// Test Uint64 == Int
func TestCheckBinaryTypedExprUint64EqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) == 4`, env, uint64(0xffffffffffffffff) == 4, reflect.TypeOf(uint64(0xffffffffffffffff) == 4))
}

// Test Uint64 == Rune
func TestCheckBinaryTypedExprUint64EqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) == '@'`, env, uint64(0xffffffffffffffff) == '@', reflect.TypeOf(uint64(0xffffffffffffffff) == '@'))
}

// Test Uint64 == Float
func TestCheckBinaryTypedExprUint64EqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) == 2.0`, env, uint64(0xffffffffffffffff) == 2.0, reflect.TypeOf(uint64(0xffffffffffffffff) == 2.0))
}

// Test Uint64 == Complex
func TestCheckBinaryTypedExprUint64EqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint64 == Bool
func TestCheckBinaryTypedExprUint64EqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == true`, env,
		`cannot convert true to type uint64`,
		`invalid operation: 18446744073709551615 == true (mismatched types uint64 and bool)`,
	)

}

// Test Uint64 == String
func TestCheckBinaryTypedExprUint64EqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == "abc"`, env,
		`cannot convert "abc" to type uint64`,
		`invalid operation: 18446744073709551615 == "abc" (mismatched types uint64 and string)`,
	)

}

// Test Uint64 == Nil
func TestCheckBinaryTypedExprUint64EqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == nil`, env,
		`cannot convert nil to type uint64`,
	)

}

// Test Uint64 == Int8
func TestCheckBinaryTypedExprUint64EqlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == int8(0x7f)`, env,
		`invalid operation: 18446744073709551615 == 127 (mismatched types uint64 and int8)`,
	)

}

// Test Uint64 == Int16
func TestCheckBinaryTypedExprUint64EqlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == int16(0x7fff)`, env,
		`invalid operation: 18446744073709551615 == 32767 (mismatched types uint64 and int16)`,
	)

}

// Test Uint64 == Int32
func TestCheckBinaryTypedExprUint64EqlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == int32(0x7fffffff)`, env,
		`invalid operation: 18446744073709551615 == 2147483647 (mismatched types uint64 and int32)`,
	)

}

// Test Uint64 == Int64
func TestCheckBinaryTypedExprUint64EqlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == int64(0x7fffffffffffffff)`, env,
		`invalid operation: 18446744073709551615 == 9223372036854775807 (mismatched types uint64 and int64)`,
	)

}

// Test Uint64 == Uint8
func TestCheckBinaryTypedExprUint64EqlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == uint8(0xff)`, env,
		`invalid operation: 18446744073709551615 == 255 (mismatched types uint64 and uint8)`,
	)

}

// Test Uint64 == Uint16
func TestCheckBinaryTypedExprUint64EqlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == uint16(0xffff)`, env,
		`invalid operation: 18446744073709551615 == 65535 (mismatched types uint64 and uint16)`,
	)

}

// Test Uint64 == Uint32
func TestCheckBinaryTypedExprUint64EqlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == uint32(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 == 4294967295 (mismatched types uint64 and uint32)`,
	)

}

// Test Uint64 == Uint64
func TestCheckBinaryTypedExprUint64EqlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) == uint64(0xffffffffffffffff)`, env, uint64(0xffffffffffffffff) == uint64(0xffffffffffffffff), reflect.TypeOf(uint64(0xffffffffffffffff) == uint64(0xffffffffffffffff)))
}

// Test Uint64 == Float32
func TestCheckBinaryTypedExprUint64EqlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == float32(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 == 4.29497e+09 (mismatched types uint64 and float32)`,
	)

}

// Test Uint64 == Float64
func TestCheckBinaryTypedExprUint64EqlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == float64(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 == 4.29497e+09 (mismatched types uint64 and float64)`,
	)

}

// Test Uint64 == Complex64
func TestCheckBinaryTypedExprUint64EqlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 18446744073709551615 == (4.29497e+09+4.29497e+09i) (mismatched types uint64 and complex64)`,
	)

}

// Test Uint64 == Complex128
func TestCheckBinaryTypedExprUint64EqlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 18446744073709551615 == (4.29497e+09+4.29497e+09i) (mismatched types uint64 and complex128)`,
	)

}

// Test Uint64 == Rune32
func TestCheckBinaryTypedExprUint64EqlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == rune(0x7fffffff)`, env,
		`invalid operation: 18446744073709551615 == rune(2147483647) (mismatched types uint64 and rune)`,
	)

}

// Test Uint64 == StringT
func TestCheckBinaryTypedExprUint64EqlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == string("abc")`, env,
		`invalid operation: 18446744073709551615 == "abc" (mismatched types uint64 and string)`,
	)

}

// Test Uint64 == BoolT
func TestCheckBinaryTypedExprUint64EqlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) == bool(true)`, env,
		`invalid operation: 18446744073709551615 == true (mismatched types uint64 and bool)`,
	)

}

// Test Uint64 > Int
func TestCheckBinaryTypedExprUint64GtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) > 4`, env, uint64(0xffffffffffffffff) > 4, reflect.TypeOf(uint64(0xffffffffffffffff) > 4))
}

// Test Uint64 > Rune
func TestCheckBinaryTypedExprUint64GtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) > '@'`, env, uint64(0xffffffffffffffff) > '@', reflect.TypeOf(uint64(0xffffffffffffffff) > '@'))
}

// Test Uint64 > Float
func TestCheckBinaryTypedExprUint64GtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) > 2.0`, env, uint64(0xffffffffffffffff) > 2.0, reflect.TypeOf(uint64(0xffffffffffffffff) > 2.0))
}

// Test Uint64 > Complex
func TestCheckBinaryTypedExprUint64GtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint64 > Bool
func TestCheckBinaryTypedExprUint64GtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > true`, env,
		`cannot convert true to type uint64`,
		`invalid operation: 18446744073709551615 > true (mismatched types uint64 and bool)`,
	)

}

// Test Uint64 > String
func TestCheckBinaryTypedExprUint64GtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > "abc"`, env,
		`cannot convert "abc" to type uint64`,
		`invalid operation: 18446744073709551615 > "abc" (mismatched types uint64 and string)`,
	)

}

// Test Uint64 > Nil
func TestCheckBinaryTypedExprUint64GtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > nil`, env,
		`cannot convert nil to type uint64`,
	)

}

// Test Uint64 > Int8
func TestCheckBinaryTypedExprUint64GtrInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > int8(0x7f)`, env,
		`invalid operation: 18446744073709551615 > 127 (mismatched types uint64 and int8)`,
	)

}

// Test Uint64 > Int16
func TestCheckBinaryTypedExprUint64GtrInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > int16(0x7fff)`, env,
		`invalid operation: 18446744073709551615 > 32767 (mismatched types uint64 and int16)`,
	)

}

// Test Uint64 > Int32
func TestCheckBinaryTypedExprUint64GtrInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > int32(0x7fffffff)`, env,
		`invalid operation: 18446744073709551615 > 2147483647 (mismatched types uint64 and int32)`,
	)

}

// Test Uint64 > Int64
func TestCheckBinaryTypedExprUint64GtrInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > int64(0x7fffffffffffffff)`, env,
		`invalid operation: 18446744073709551615 > 9223372036854775807 (mismatched types uint64 and int64)`,
	)

}

// Test Uint64 > Uint8
func TestCheckBinaryTypedExprUint64GtrUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > uint8(0xff)`, env,
		`invalid operation: 18446744073709551615 > 255 (mismatched types uint64 and uint8)`,
	)

}

// Test Uint64 > Uint16
func TestCheckBinaryTypedExprUint64GtrUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > uint16(0xffff)`, env,
		`invalid operation: 18446744073709551615 > 65535 (mismatched types uint64 and uint16)`,
	)

}

// Test Uint64 > Uint32
func TestCheckBinaryTypedExprUint64GtrUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > uint32(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 > 4294967295 (mismatched types uint64 and uint32)`,
	)

}

// Test Uint64 > Uint64
func TestCheckBinaryTypedExprUint64GtrUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff) > uint64(0xffffffffffffffff)`, env, uint64(0xffffffffffffffff) > uint64(0xffffffffffffffff), reflect.TypeOf(uint64(0xffffffffffffffff) > uint64(0xffffffffffffffff)))
}

// Test Uint64 > Float32
func TestCheckBinaryTypedExprUint64GtrFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > float32(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 > 4.29497e+09 (mismatched types uint64 and float32)`,
	)

}

// Test Uint64 > Float64
func TestCheckBinaryTypedExprUint64GtrFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > float64(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 > 4.29497e+09 (mismatched types uint64 and float64)`,
	)

}

// Test Uint64 > Complex64
func TestCheckBinaryTypedExprUint64GtrComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 18446744073709551615 > (4.29497e+09+4.29497e+09i) (mismatched types uint64 and complex64)`,
	)

}

// Test Uint64 > Complex128
func TestCheckBinaryTypedExprUint64GtrComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 18446744073709551615 > (4.29497e+09+4.29497e+09i) (mismatched types uint64 and complex128)`,
	)

}

// Test Uint64 > Rune32
func TestCheckBinaryTypedExprUint64GtrRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > rune(0x7fffffff)`, env,
		`invalid operation: 18446744073709551615 > rune(2147483647) (mismatched types uint64 and rune)`,
	)

}

// Test Uint64 > StringT
func TestCheckBinaryTypedExprUint64GtrStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > string("abc")`, env,
		`invalid operation: 18446744073709551615 > "abc" (mismatched types uint64 and string)`,
	)

}

// Test Uint64 > BoolT
func TestCheckBinaryTypedExprUint64GtrBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) > bool(true)`, env,
		`invalid operation: 18446744073709551615 > true (mismatched types uint64 and bool)`,
	)

}

// Test Uint64 << Int
func TestCheckBinaryTypedExprUint64ShlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << 4`, env,
		`constant 295147905179352825840 overflows uint64`,
	)

}

// Test Uint64 << Rune
func TestCheckBinaryTypedExprUint64ShlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << '@'`, env,
		`constant 340282366920938463444927863358058659840 overflows uint64`,
	)

}

// Test Uint64 << Float
func TestCheckBinaryTypedExprUint64ShlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << 2.0`, env,
		`constant 73786976294838206460 overflows uint64`,
	)

}

// Test Uint64 << Complex
func TestCheckBinaryTypedExprUint64ShlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Uint64 << Bool
func TestCheckBinaryTypedExprUint64ShlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << true`, env,
		`invalid operation: 18446744073709551615 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Uint64 << String
func TestCheckBinaryTypedExprUint64ShlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: 18446744073709551615 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Uint64 << Nil
func TestCheckBinaryTypedExprUint64ShlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Uint64 << Int8
func TestCheckBinaryTypedExprUint64ShlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << int8(0x7f)`, env,
		`invalid operation: 18446744073709551615 << 127 (shift count type int8, must be unsigned integer)`,
	)

}

// Test Uint64 << Int16
func TestCheckBinaryTypedExprUint64ShlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << int16(0x7fff)`, env,
		`invalid operation: 18446744073709551615 << 32767 (shift count type int16, must be unsigned integer)`,
	)

}

// Test Uint64 << Int32
func TestCheckBinaryTypedExprUint64ShlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << int32(0x7fffffff)`, env,
		`invalid operation: 18446744073709551615 << 2147483647 (shift count type int32, must be unsigned integer)`,
	)

}

// Test Uint64 << Int64
func TestCheckBinaryTypedExprUint64ShlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << int64(0x7fffffffffffffff)`, env,
		`invalid operation: 18446744073709551615 << 9223372036854775807 (shift count type int64, must be unsigned integer)`,
	)

}

// Test Uint64 << Uint8
func TestCheckBinaryTypedExprUint64ShlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << uint8(0xff)`, env,
		`constant 1067993517960455041139614808466117959589566768673982431046885811578289580870591483007524478648320 overflows uint64`,
	)

}

// Test Uint64 << Uint16
func TestCheckBinaryTypedExprUint64ShlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << uint16(0xffff)`, env,
		`stupid shift: 65535`,
	)

}

// Test Uint64 << Uint32
func TestCheckBinaryTypedExprUint64ShlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << uint32(0xffffffff)`, env,
		`stupid shift: 4294967295`,
	)

}

// Test Uint64 << Uint64
func TestCheckBinaryTypedExprUint64ShlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << uint64(0xffffffffffffffff)`, env,
		`stupid shift: -1`,
	)

}

// Test Uint64 << Float32
func TestCheckBinaryTypedExprUint64ShlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << float32(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 << 4.29497e+09 (shift count type float32, must be unsigned integer)`,
	)

}

// Test Uint64 << Float64
func TestCheckBinaryTypedExprUint64ShlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << float64(0xffffffff)`, env,
		`invalid operation: 18446744073709551615 << 4.29497e+09 (shift count type float64, must be unsigned integer)`,
	)

}

// Test Uint64 << Complex64
func TestCheckBinaryTypedExprUint64ShlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 18446744073709551615 << (4.29497e+09+4.29497e+09i) (shift count type complex64, must be unsigned integer)`,
	)

}

// Test Uint64 << Complex128
func TestCheckBinaryTypedExprUint64ShlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 18446744073709551615 << (4.29497e+09+4.29497e+09i) (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Uint64 << Rune32
func TestCheckBinaryTypedExprUint64ShlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << rune(0x7fffffff)`, env,
		`invalid operation: 18446744073709551615 << rune(2147483647) (shift count type rune, must be unsigned integer)`,
	)

}

// Test Uint64 << StringT
func TestCheckBinaryTypedExprUint64ShlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << string("abc")`, env,
		`invalid operation: 18446744073709551615 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Uint64 << BoolT
func TestCheckBinaryTypedExprUint64ShlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(0xffffffffffffffff) << bool(true)`, env,
		`invalid operation: 18446744073709551615 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Float32 + Int
func TestCheckBinaryTypedExprFloat32AddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) + 4`, env, float32(0xffffffff) + 4, reflect.TypeOf(float32(0xffffffff) + 4))
}

// Test Float32 + Rune
func TestCheckBinaryTypedExprFloat32AddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) + '@'`, env, float32(0xffffffff) + '@', reflect.TypeOf(float32(0xffffffff) + '@'))
}

// Test Float32 + Float
func TestCheckBinaryTypedExprFloat32AddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) + 2.0`, env, float32(0xffffffff) + 2.0, reflect.TypeOf(float32(0xffffffff) + 2.0))
}

// Test Float32 + Complex
func TestCheckBinaryTypedExprFloat32AddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Float32 + Bool
func TestCheckBinaryTypedExprFloat32AddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + true`, env,
		`cannot convert true to type float32`,
		`invalid operation: 4.29497e+09 + true (mismatched types float32 and bool)`,
	)

}

// Test Float32 + String
func TestCheckBinaryTypedExprFloat32AddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + "abc"`, env,
		`cannot convert "abc" to type float32`,
		`invalid operation: 4.29497e+09 + "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 + Nil
func TestCheckBinaryTypedExprFloat32AddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + nil`, env,
		`cannot convert nil to type float32`,
	)

}

// Test Float32 + Int8
func TestCheckBinaryTypedExprFloat32AddInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + int8(0x7f)`, env,
		`invalid operation: 4.29497e+09 + 127 (mismatched types float32 and int8)`,
	)

}

// Test Float32 + Int16
func TestCheckBinaryTypedExprFloat32AddInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + int16(0x7fff)`, env,
		`invalid operation: 4.29497e+09 + 32767 (mismatched types float32 and int16)`,
	)

}

// Test Float32 + Int32
func TestCheckBinaryTypedExprFloat32AddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + int32(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 + 2147483647 (mismatched types float32 and int32)`,
	)

}

// Test Float32 + Int64
func TestCheckBinaryTypedExprFloat32AddInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 + 9223372036854775807 (mismatched types float32 and int64)`,
	)

}

// Test Float32 + Uint8
func TestCheckBinaryTypedExprFloat32AddUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + uint8(0xff)`, env,
		`invalid operation: 4.29497e+09 + 255 (mismatched types float32 and uint8)`,
	)

}

// Test Float32 + Uint16
func TestCheckBinaryTypedExprFloat32AddUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + uint16(0xffff)`, env,
		`invalid operation: 4.29497e+09 + 65535 (mismatched types float32 and uint16)`,
	)

}

// Test Float32 + Uint32
func TestCheckBinaryTypedExprFloat32AddUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + uint32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 + 4294967295 (mismatched types float32 and uint32)`,
	)

}

// Test Float32 + Uint64
func TestCheckBinaryTypedExprFloat32AddUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 + 18446744073709551615 (mismatched types float32 and uint64)`,
	)

}

// Test Float32 + Float32
func TestCheckBinaryTypedExprFloat32AddFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) + float32(0xffffffff)`, env, float32(0xffffffff) + float32(0xffffffff), reflect.TypeOf(float32(0xffffffff) + float32(0xffffffff)))
}

// Test Float32 + Float64
func TestCheckBinaryTypedExprFloat32AddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + float64(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 + 4.29497e+09 (mismatched types float32 and float64)`,
	)

}

// Test Float32 + Complex64
func TestCheckBinaryTypedExprFloat32AddComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 + (4.29497e+09+4.29497e+09i) (mismatched types float32 and complex64)`,
	)

}

// Test Float32 + Complex128
func TestCheckBinaryTypedExprFloat32AddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 + (4.29497e+09+4.29497e+09i) (mismatched types float32 and complex128)`,
	)

}

// Test Float32 + Rune32
func TestCheckBinaryTypedExprFloat32AddRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + rune(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 + rune(2147483647) (mismatched types float32 and rune)`,
	)

}

// Test Float32 + StringT
func TestCheckBinaryTypedExprFloat32AddStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + string("abc")`, env,
		`invalid operation: 4.29497e+09 + "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 + BoolT
func TestCheckBinaryTypedExprFloat32AddBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) + bool(true)`, env,
		`invalid operation: 4.29497e+09 + true (mismatched types float32 and bool)`,
	)

}

// Test Float32 - Int
func TestCheckBinaryTypedExprFloat32SubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) - 4`, env, float32(0xffffffff) - 4, reflect.TypeOf(float32(0xffffffff) - 4))
}

// Test Float32 - Rune
func TestCheckBinaryTypedExprFloat32SubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) - '@'`, env, float32(0xffffffff) - '@', reflect.TypeOf(float32(0xffffffff) - '@'))
}

// Test Float32 - Float
func TestCheckBinaryTypedExprFloat32SubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) - 2.0`, env, float32(0xffffffff) - 2.0, reflect.TypeOf(float32(0xffffffff) - 2.0))
}

// Test Float32 - Complex
func TestCheckBinaryTypedExprFloat32SubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Float32 - Bool
func TestCheckBinaryTypedExprFloat32SubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - true`, env,
		`cannot convert true to type float32`,
		`invalid operation: 4.29497e+09 - true (mismatched types float32 and bool)`,
	)

}

// Test Float32 - String
func TestCheckBinaryTypedExprFloat32SubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - "abc"`, env,
		`cannot convert "abc" to type float32`,
		`invalid operation: 4.29497e+09 - "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 - Nil
func TestCheckBinaryTypedExprFloat32SubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - nil`, env,
		`cannot convert nil to type float32`,
	)

}

// Test Float32 - Int8
func TestCheckBinaryTypedExprFloat32SubInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - int8(0x7f)`, env,
		`invalid operation: 4.29497e+09 - 127 (mismatched types float32 and int8)`,
	)

}

// Test Float32 - Int16
func TestCheckBinaryTypedExprFloat32SubInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - int16(0x7fff)`, env,
		`invalid operation: 4.29497e+09 - 32767 (mismatched types float32 and int16)`,
	)

}

// Test Float32 - Int32
func TestCheckBinaryTypedExprFloat32SubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - int32(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 - 2147483647 (mismatched types float32 and int32)`,
	)

}

// Test Float32 - Int64
func TestCheckBinaryTypedExprFloat32SubInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 - 9223372036854775807 (mismatched types float32 and int64)`,
	)

}

// Test Float32 - Uint8
func TestCheckBinaryTypedExprFloat32SubUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - uint8(0xff)`, env,
		`invalid operation: 4.29497e+09 - 255 (mismatched types float32 and uint8)`,
	)

}

// Test Float32 - Uint16
func TestCheckBinaryTypedExprFloat32SubUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - uint16(0xffff)`, env,
		`invalid operation: 4.29497e+09 - 65535 (mismatched types float32 and uint16)`,
	)

}

// Test Float32 - Uint32
func TestCheckBinaryTypedExprFloat32SubUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - uint32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 - 4294967295 (mismatched types float32 and uint32)`,
	)

}

// Test Float32 - Uint64
func TestCheckBinaryTypedExprFloat32SubUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 - 18446744073709551615 (mismatched types float32 and uint64)`,
	)

}

// Test Float32 - Float32
func TestCheckBinaryTypedExprFloat32SubFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) - float32(0xffffffff)`, env, float32(0xffffffff) - float32(0xffffffff), reflect.TypeOf(float32(0xffffffff) - float32(0xffffffff)))
}

// Test Float32 - Float64
func TestCheckBinaryTypedExprFloat32SubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - float64(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 - 4.29497e+09 (mismatched types float32 and float64)`,
	)

}

// Test Float32 - Complex64
func TestCheckBinaryTypedExprFloat32SubComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 - (4.29497e+09+4.29497e+09i) (mismatched types float32 and complex64)`,
	)

}

// Test Float32 - Complex128
func TestCheckBinaryTypedExprFloat32SubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 - (4.29497e+09+4.29497e+09i) (mismatched types float32 and complex128)`,
	)

}

// Test Float32 - Rune32
func TestCheckBinaryTypedExprFloat32SubRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - rune(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 - rune(2147483647) (mismatched types float32 and rune)`,
	)

}

// Test Float32 - StringT
func TestCheckBinaryTypedExprFloat32SubStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - string("abc")`, env,
		`invalid operation: 4.29497e+09 - "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 - BoolT
func TestCheckBinaryTypedExprFloat32SubBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) - bool(true)`, env,
		`invalid operation: 4.29497e+09 - true (mismatched types float32 and bool)`,
	)

}

// Test Float32 & Int
func TestCheckBinaryTypedExprFloat32AndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & 4`, env,
		`invalid operation: 4.29497e+09 & 4 (operator & not defined on float32)`,
	)

}

// Test Float32 & Rune
func TestCheckBinaryTypedExprFloat32AndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & '@'`, env,
		`invalid operation: 4.29497e+09 & 64 (operator & not defined on float32)`,
	)

}

// Test Float32 & Float
func TestCheckBinaryTypedExprFloat32AndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & 2.0`, env,
		`invalid operation: 4.29497e+09 & 2 (operator & not defined on float32)`,
	)

}

// Test Float32 & Complex
func TestCheckBinaryTypedExprFloat32AndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: 4.29497e+09 & 0 (operator & not defined on float32)`,
	)

}

// Test Float32 & Bool
func TestCheckBinaryTypedExprFloat32AndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & true`, env,
		`cannot convert true to type float32`,
		`invalid operation: 4.29497e+09 & true (mismatched types float32 and bool)`,
	)

}

// Test Float32 & String
func TestCheckBinaryTypedExprFloat32AndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & "abc"`, env,
		`cannot convert "abc" to type float32`,
		`invalid operation: 4.29497e+09 & "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 & Nil
func TestCheckBinaryTypedExprFloat32AndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & nil`, env,
		`cannot convert nil to type float32`,
	)

}

// Test Float32 & Int8
func TestCheckBinaryTypedExprFloat32AndInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & int8(0x7f)`, env,
		`invalid operation: 4.29497e+09 & 127 (mismatched types float32 and int8)`,
	)

}

// Test Float32 & Int16
func TestCheckBinaryTypedExprFloat32AndInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & int16(0x7fff)`, env,
		`invalid operation: 4.29497e+09 & 32767 (mismatched types float32 and int16)`,
	)

}

// Test Float32 & Int32
func TestCheckBinaryTypedExprFloat32AndInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & int32(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 & 2147483647 (mismatched types float32 and int32)`,
	)

}

// Test Float32 & Int64
func TestCheckBinaryTypedExprFloat32AndInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 & 9223372036854775807 (mismatched types float32 and int64)`,
	)

}

// Test Float32 & Uint8
func TestCheckBinaryTypedExprFloat32AndUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & uint8(0xff)`, env,
		`invalid operation: 4.29497e+09 & 255 (mismatched types float32 and uint8)`,
	)

}

// Test Float32 & Uint16
func TestCheckBinaryTypedExprFloat32AndUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & uint16(0xffff)`, env,
		`invalid operation: 4.29497e+09 & 65535 (mismatched types float32 and uint16)`,
	)

}

// Test Float32 & Uint32
func TestCheckBinaryTypedExprFloat32AndUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & uint32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 & 4294967295 (mismatched types float32 and uint32)`,
	)

}

// Test Float32 & Uint64
func TestCheckBinaryTypedExprFloat32AndUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 & 18446744073709551615 (mismatched types float32 and uint64)`,
	)

}

// Test Float32 & Float32
func TestCheckBinaryTypedExprFloat32AndFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & float32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 & 4.29497e+09 (operator & not defined on float32)`,
	)

}

// Test Float32 & Float64
func TestCheckBinaryTypedExprFloat32AndFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & float64(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 & 4.29497e+09 (mismatched types float32 and float64)`,
	)

}

// Test Float32 & Complex64
func TestCheckBinaryTypedExprFloat32AndComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 & (4.29497e+09+4.29497e+09i) (mismatched types float32 and complex64)`,
	)

}

// Test Float32 & Complex128
func TestCheckBinaryTypedExprFloat32AndComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 & (4.29497e+09+4.29497e+09i) (mismatched types float32 and complex128)`,
	)

}

// Test Float32 & Rune32
func TestCheckBinaryTypedExprFloat32AndRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & rune(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 & rune(2147483647) (mismatched types float32 and rune)`,
	)

}

// Test Float32 & StringT
func TestCheckBinaryTypedExprFloat32AndStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & string("abc")`, env,
		`invalid operation: 4.29497e+09 & "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 & BoolT
func TestCheckBinaryTypedExprFloat32AndBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) & bool(true)`, env,
		`invalid operation: 4.29497e+09 & true (mismatched types float32 and bool)`,
	)

}

// Test Float32 % Int
func TestCheckBinaryTypedExprFloat32RemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % 4`, env,
		`invalid operation: 4.29497e+09 % 4 (operator % not defined on float32)`,
	)

}

// Test Float32 % Rune
func TestCheckBinaryTypedExprFloat32RemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % '@'`, env,
		`invalid operation: 4.29497e+09 % 64 (operator % not defined on float32)`,
	)

}

// Test Float32 % Float
func TestCheckBinaryTypedExprFloat32RemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % 2.0`, env,
		`invalid operation: 4.29497e+09 % 2 (operator % not defined on float32)`,
	)

}

// Test Float32 % Complex
func TestCheckBinaryTypedExprFloat32RemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: 4.29497e+09 % 0 (operator % not defined on float32)`,
	)

}

// Test Float32 % Bool
func TestCheckBinaryTypedExprFloat32RemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % true`, env,
		`cannot convert true to type float32`,
		`invalid operation: 4.29497e+09 % true (mismatched types float32 and bool)`,
	)

}

// Test Float32 % String
func TestCheckBinaryTypedExprFloat32RemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % "abc"`, env,
		`cannot convert "abc" to type float32`,
		`invalid operation: 4.29497e+09 % "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 % Nil
func TestCheckBinaryTypedExprFloat32RemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % nil`, env,
		`cannot convert nil to type float32`,
	)

}

// Test Float32 % Int8
func TestCheckBinaryTypedExprFloat32RemInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % int8(0x7f)`, env,
		`invalid operation: 4.29497e+09 % 127 (mismatched types float32 and int8)`,
	)

}

// Test Float32 % Int16
func TestCheckBinaryTypedExprFloat32RemInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % int16(0x7fff)`, env,
		`invalid operation: 4.29497e+09 % 32767 (mismatched types float32 and int16)`,
	)

}

// Test Float32 % Int32
func TestCheckBinaryTypedExprFloat32RemInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % int32(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 % 2147483647 (mismatched types float32 and int32)`,
	)

}

// Test Float32 % Int64
func TestCheckBinaryTypedExprFloat32RemInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 % 9223372036854775807 (mismatched types float32 and int64)`,
	)

}

// Test Float32 % Uint8
func TestCheckBinaryTypedExprFloat32RemUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % uint8(0xff)`, env,
		`invalid operation: 4.29497e+09 % 255 (mismatched types float32 and uint8)`,
	)

}

// Test Float32 % Uint16
func TestCheckBinaryTypedExprFloat32RemUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % uint16(0xffff)`, env,
		`invalid operation: 4.29497e+09 % 65535 (mismatched types float32 and uint16)`,
	)

}

// Test Float32 % Uint32
func TestCheckBinaryTypedExprFloat32RemUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % uint32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 % 4294967295 (mismatched types float32 and uint32)`,
	)

}

// Test Float32 % Uint64
func TestCheckBinaryTypedExprFloat32RemUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 % 18446744073709551615 (mismatched types float32 and uint64)`,
	)

}

// Test Float32 % Float32
func TestCheckBinaryTypedExprFloat32RemFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % float32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 % 4.29497e+09 (operator % not defined on float32)`,
	)

}

// Test Float32 % Float64
func TestCheckBinaryTypedExprFloat32RemFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % float64(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 % 4.29497e+09 (mismatched types float32 and float64)`,
	)

}

// Test Float32 % Complex64
func TestCheckBinaryTypedExprFloat32RemComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 % (4.29497e+09+4.29497e+09i) (mismatched types float32 and complex64)`,
	)

}

// Test Float32 % Complex128
func TestCheckBinaryTypedExprFloat32RemComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 % (4.29497e+09+4.29497e+09i) (mismatched types float32 and complex128)`,
	)

}

// Test Float32 % Rune32
func TestCheckBinaryTypedExprFloat32RemRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % rune(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 % rune(2147483647) (mismatched types float32 and rune)`,
	)

}

// Test Float32 % StringT
func TestCheckBinaryTypedExprFloat32RemStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % string("abc")`, env,
		`invalid operation: 4.29497e+09 % "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 % BoolT
func TestCheckBinaryTypedExprFloat32RemBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) % bool(true)`, env,
		`invalid operation: 4.29497e+09 % true (mismatched types float32 and bool)`,
	)

}

// Test Float32 == Int
func TestCheckBinaryTypedExprFloat32EqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) == 4`, env, float32(0xffffffff) == 4, reflect.TypeOf(float32(0xffffffff) == 4))
}

// Test Float32 == Rune
func TestCheckBinaryTypedExprFloat32EqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) == '@'`, env, float32(0xffffffff) == '@', reflect.TypeOf(float32(0xffffffff) == '@'))
}

// Test Float32 == Float
func TestCheckBinaryTypedExprFloat32EqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) == 2.0`, env, float32(0xffffffff) == 2.0, reflect.TypeOf(float32(0xffffffff) == 2.0))
}

// Test Float32 == Complex
func TestCheckBinaryTypedExprFloat32EqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Float32 == Bool
func TestCheckBinaryTypedExprFloat32EqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == true`, env,
		`cannot convert true to type float32`,
		`invalid operation: 4.29497e+09 == true (mismatched types float32 and bool)`,
	)

}

// Test Float32 == String
func TestCheckBinaryTypedExprFloat32EqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == "abc"`, env,
		`cannot convert "abc" to type float32`,
		`invalid operation: 4.29497e+09 == "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 == Nil
func TestCheckBinaryTypedExprFloat32EqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == nil`, env,
		`cannot convert nil to type float32`,
	)

}

// Test Float32 == Int8
func TestCheckBinaryTypedExprFloat32EqlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == int8(0x7f)`, env,
		`invalid operation: 4.29497e+09 == 127 (mismatched types float32 and int8)`,
	)

}

// Test Float32 == Int16
func TestCheckBinaryTypedExprFloat32EqlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == int16(0x7fff)`, env,
		`invalid operation: 4.29497e+09 == 32767 (mismatched types float32 and int16)`,
	)

}

// Test Float32 == Int32
func TestCheckBinaryTypedExprFloat32EqlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == int32(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 == 2147483647 (mismatched types float32 and int32)`,
	)

}

// Test Float32 == Int64
func TestCheckBinaryTypedExprFloat32EqlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 == 9223372036854775807 (mismatched types float32 and int64)`,
	)

}

// Test Float32 == Uint8
func TestCheckBinaryTypedExprFloat32EqlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == uint8(0xff)`, env,
		`invalid operation: 4.29497e+09 == 255 (mismatched types float32 and uint8)`,
	)

}

// Test Float32 == Uint16
func TestCheckBinaryTypedExprFloat32EqlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == uint16(0xffff)`, env,
		`invalid operation: 4.29497e+09 == 65535 (mismatched types float32 and uint16)`,
	)

}

// Test Float32 == Uint32
func TestCheckBinaryTypedExprFloat32EqlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == uint32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 == 4294967295 (mismatched types float32 and uint32)`,
	)

}

// Test Float32 == Uint64
func TestCheckBinaryTypedExprFloat32EqlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 == 18446744073709551615 (mismatched types float32 and uint64)`,
	)

}

// Test Float32 == Float32
func TestCheckBinaryTypedExprFloat32EqlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) == float32(0xffffffff)`, env, float32(0xffffffff) == float32(0xffffffff), reflect.TypeOf(float32(0xffffffff) == float32(0xffffffff)))
}

// Test Float32 == Float64
func TestCheckBinaryTypedExprFloat32EqlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == float64(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 == 4.29497e+09 (mismatched types float32 and float64)`,
	)

}

// Test Float32 == Complex64
func TestCheckBinaryTypedExprFloat32EqlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 == (4.29497e+09+4.29497e+09i) (mismatched types float32 and complex64)`,
	)

}

// Test Float32 == Complex128
func TestCheckBinaryTypedExprFloat32EqlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 == (4.29497e+09+4.29497e+09i) (mismatched types float32 and complex128)`,
	)

}

// Test Float32 == Rune32
func TestCheckBinaryTypedExprFloat32EqlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == rune(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 == rune(2147483647) (mismatched types float32 and rune)`,
	)

}

// Test Float32 == StringT
func TestCheckBinaryTypedExprFloat32EqlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == string("abc")`, env,
		`invalid operation: 4.29497e+09 == "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 == BoolT
func TestCheckBinaryTypedExprFloat32EqlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) == bool(true)`, env,
		`invalid operation: 4.29497e+09 == true (mismatched types float32 and bool)`,
	)

}

// Test Float32 > Int
func TestCheckBinaryTypedExprFloat32GtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) > 4`, env, float32(0xffffffff) > 4, reflect.TypeOf(float32(0xffffffff) > 4))
}

// Test Float32 > Rune
func TestCheckBinaryTypedExprFloat32GtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) > '@'`, env, float32(0xffffffff) > '@', reflect.TypeOf(float32(0xffffffff) > '@'))
}

// Test Float32 > Float
func TestCheckBinaryTypedExprFloat32GtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) > 2.0`, env, float32(0xffffffff) > 2.0, reflect.TypeOf(float32(0xffffffff) > 2.0))
}

// Test Float32 > Complex
func TestCheckBinaryTypedExprFloat32GtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Float32 > Bool
func TestCheckBinaryTypedExprFloat32GtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > true`, env,
		`cannot convert true to type float32`,
		`invalid operation: 4.29497e+09 > true (mismatched types float32 and bool)`,
	)

}

// Test Float32 > String
func TestCheckBinaryTypedExprFloat32GtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > "abc"`, env,
		`cannot convert "abc" to type float32`,
		`invalid operation: 4.29497e+09 > "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 > Nil
func TestCheckBinaryTypedExprFloat32GtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > nil`, env,
		`cannot convert nil to type float32`,
	)

}

// Test Float32 > Int8
func TestCheckBinaryTypedExprFloat32GtrInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > int8(0x7f)`, env,
		`invalid operation: 4.29497e+09 > 127 (mismatched types float32 and int8)`,
	)

}

// Test Float32 > Int16
func TestCheckBinaryTypedExprFloat32GtrInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > int16(0x7fff)`, env,
		`invalid operation: 4.29497e+09 > 32767 (mismatched types float32 and int16)`,
	)

}

// Test Float32 > Int32
func TestCheckBinaryTypedExprFloat32GtrInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > int32(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 > 2147483647 (mismatched types float32 and int32)`,
	)

}

// Test Float32 > Int64
func TestCheckBinaryTypedExprFloat32GtrInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 > 9223372036854775807 (mismatched types float32 and int64)`,
	)

}

// Test Float32 > Uint8
func TestCheckBinaryTypedExprFloat32GtrUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > uint8(0xff)`, env,
		`invalid operation: 4.29497e+09 > 255 (mismatched types float32 and uint8)`,
	)

}

// Test Float32 > Uint16
func TestCheckBinaryTypedExprFloat32GtrUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > uint16(0xffff)`, env,
		`invalid operation: 4.29497e+09 > 65535 (mismatched types float32 and uint16)`,
	)

}

// Test Float32 > Uint32
func TestCheckBinaryTypedExprFloat32GtrUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > uint32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 > 4294967295 (mismatched types float32 and uint32)`,
	)

}

// Test Float32 > Uint64
func TestCheckBinaryTypedExprFloat32GtrUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 > 18446744073709551615 (mismatched types float32 and uint64)`,
	)

}

// Test Float32 > Float32
func TestCheckBinaryTypedExprFloat32GtrFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff) > float32(0xffffffff)`, env, float32(0xffffffff) > float32(0xffffffff), reflect.TypeOf(float32(0xffffffff) > float32(0xffffffff)))
}

// Test Float32 > Float64
func TestCheckBinaryTypedExprFloat32GtrFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > float64(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 > 4.29497e+09 (mismatched types float32 and float64)`,
	)

}

// Test Float32 > Complex64
func TestCheckBinaryTypedExprFloat32GtrComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 > (4.29497e+09+4.29497e+09i) (mismatched types float32 and complex64)`,
	)

}

// Test Float32 > Complex128
func TestCheckBinaryTypedExprFloat32GtrComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 > (4.29497e+09+4.29497e+09i) (mismatched types float32 and complex128)`,
	)

}

// Test Float32 > Rune32
func TestCheckBinaryTypedExprFloat32GtrRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > rune(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 > rune(2147483647) (mismatched types float32 and rune)`,
	)

}

// Test Float32 > StringT
func TestCheckBinaryTypedExprFloat32GtrStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > string("abc")`, env,
		`invalid operation: 4.29497e+09 > "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 > BoolT
func TestCheckBinaryTypedExprFloat32GtrBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) > bool(true)`, env,
		`invalid operation: 4.29497e+09 > true (mismatched types float32 and bool)`,
	)

}

// Test Float32 << Int
func TestCheckBinaryTypedExprFloat32ShlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << 4`, env,
		`invalid operation: 4.29497e+09 << 4 (shift of type float32)`,
	)

}

// Test Float32 << Rune
func TestCheckBinaryTypedExprFloat32ShlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << '@'`, env,
		`invalid operation: 4.29497e+09 << 64 (shift of type float32)`,
	)

}

// Test Float32 << Float
func TestCheckBinaryTypedExprFloat32ShlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << 2.0`, env,
		`invalid operation: 4.29497e+09 << 2 (shift of type float32)`,
	)

}

// Test Float32 << Complex
func TestCheckBinaryTypedExprFloat32ShlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: 4.29497e+09 << 0 (shift of type float32)`,
	)

}

// Test Float32 << Bool
func TestCheckBinaryTypedExprFloat32ShlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << true`, env,
		`invalid operation: 4.29497e+09 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Float32 << String
func TestCheckBinaryTypedExprFloat32ShlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: 4.29497e+09 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Float32 << Nil
func TestCheckBinaryTypedExprFloat32ShlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Float32 << Int8
func TestCheckBinaryTypedExprFloat32ShlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << int8(0x7f)`, env,
		`invalid operation: 4.29497e+09 << 127 (shift count type int8, must be unsigned integer)`,
	)

}

// Test Float32 << Int16
func TestCheckBinaryTypedExprFloat32ShlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << int16(0x7fff)`, env,
		`invalid operation: 4.29497e+09 << 32767 (shift count type int16, must be unsigned integer)`,
	)

}

// Test Float32 << Int32
func TestCheckBinaryTypedExprFloat32ShlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << int32(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 << 2147483647 (shift count type int32, must be unsigned integer)`,
	)

}

// Test Float32 << Int64
func TestCheckBinaryTypedExprFloat32ShlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 << 9223372036854775807 (shift count type int64, must be unsigned integer)`,
	)

}

// Test Float32 << Uint8
func TestCheckBinaryTypedExprFloat32ShlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << uint8(0xff)`, env,
		`invalid operation: 4.29497e+09 << 255 (shift of type float32)`,
	)

}

// Test Float32 << Uint16
func TestCheckBinaryTypedExprFloat32ShlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << uint16(0xffff)`, env,
		`invalid operation: 4.29497e+09 << 65535 (shift of type float32)`,
	)

}

// Test Float32 << Uint32
func TestCheckBinaryTypedExprFloat32ShlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << uint32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 << 4294967295 (shift of type float32)`,
	)

}

// Test Float32 << Uint64
func TestCheckBinaryTypedExprFloat32ShlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 << 18446744073709551615 (shift of type float32)`,
	)

}

// Test Float32 << Float32
func TestCheckBinaryTypedExprFloat32ShlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << float32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 << 4.29497e+09 (shift count type float32, must be unsigned integer)`,
	)

}

// Test Float32 << Float64
func TestCheckBinaryTypedExprFloat32ShlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << float64(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 << 4.29497e+09 (shift count type float64, must be unsigned integer)`,
	)

}

// Test Float32 << Complex64
func TestCheckBinaryTypedExprFloat32ShlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 << (4.29497e+09+4.29497e+09i) (shift count type complex64, must be unsigned integer)`,
	)

}

// Test Float32 << Complex128
func TestCheckBinaryTypedExprFloat32ShlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 << (4.29497e+09+4.29497e+09i) (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Float32 << Rune32
func TestCheckBinaryTypedExprFloat32ShlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << rune(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 << rune(2147483647) (shift count type rune, must be unsigned integer)`,
	)

}

// Test Float32 << StringT
func TestCheckBinaryTypedExprFloat32ShlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << string("abc")`, env,
		`invalid operation: 4.29497e+09 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Float32 << BoolT
func TestCheckBinaryTypedExprFloat32ShlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(0xffffffff) << bool(true)`, env,
		`invalid operation: 4.29497e+09 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Float64 + Int
func TestCheckBinaryTypedExprFloat64AddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) + 4`, env, float64(0xffffffff) + 4, reflect.TypeOf(float64(0xffffffff) + 4))
}

// Test Float64 + Rune
func TestCheckBinaryTypedExprFloat64AddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) + '@'`, env, float64(0xffffffff) + '@', reflect.TypeOf(float64(0xffffffff) + '@'))
}

// Test Float64 + Float
func TestCheckBinaryTypedExprFloat64AddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) + 2.0`, env, float64(0xffffffff) + 2.0, reflect.TypeOf(float64(0xffffffff) + 2.0))
}

// Test Float64 + Complex
func TestCheckBinaryTypedExprFloat64AddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Float64 + Bool
func TestCheckBinaryTypedExprFloat64AddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 4.29497e+09 + true (mismatched types float64 and bool)`,
	)

}

// Test Float64 + String
func TestCheckBinaryTypedExprFloat64AddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 4.29497e+09 + "abc" (mismatched types float64 and string)`,
	)

}

// Test Float64 + Nil
func TestCheckBinaryTypedExprFloat64AddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + nil`, env,
		`cannot convert nil to type float64`,
	)

}

// Test Float64 + Int8
func TestCheckBinaryTypedExprFloat64AddInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + int8(0x7f)`, env,
		`invalid operation: 4.29497e+09 + 127 (mismatched types float64 and int8)`,
	)

}

// Test Float64 + Int16
func TestCheckBinaryTypedExprFloat64AddInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + int16(0x7fff)`, env,
		`invalid operation: 4.29497e+09 + 32767 (mismatched types float64 and int16)`,
	)

}

// Test Float64 + Int32
func TestCheckBinaryTypedExprFloat64AddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + int32(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 + 2147483647 (mismatched types float64 and int32)`,
	)

}

// Test Float64 + Int64
func TestCheckBinaryTypedExprFloat64AddInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 + 9223372036854775807 (mismatched types float64 and int64)`,
	)

}

// Test Float64 + Uint8
func TestCheckBinaryTypedExprFloat64AddUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + uint8(0xff)`, env,
		`invalid operation: 4.29497e+09 + 255 (mismatched types float64 and uint8)`,
	)

}

// Test Float64 + Uint16
func TestCheckBinaryTypedExprFloat64AddUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + uint16(0xffff)`, env,
		`invalid operation: 4.29497e+09 + 65535 (mismatched types float64 and uint16)`,
	)

}

// Test Float64 + Uint32
func TestCheckBinaryTypedExprFloat64AddUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + uint32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 + 4294967295 (mismatched types float64 and uint32)`,
	)

}

// Test Float64 + Uint64
func TestCheckBinaryTypedExprFloat64AddUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 + 18446744073709551615 (mismatched types float64 and uint64)`,
	)

}

// Test Float64 + Float32
func TestCheckBinaryTypedExprFloat64AddFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + float32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 + 4.29497e+09 (mismatched types float64 and float32)`,
	)

}

// Test Float64 + Float64
func TestCheckBinaryTypedExprFloat64AddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) + float64(0xffffffff)`, env, float64(0xffffffff) + float64(0xffffffff), reflect.TypeOf(float64(0xffffffff) + float64(0xffffffff)))
}

// Test Float64 + Complex64
func TestCheckBinaryTypedExprFloat64AddComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 + (4.29497e+09+4.29497e+09i) (mismatched types float64 and complex64)`,
	)

}

// Test Float64 + Complex128
func TestCheckBinaryTypedExprFloat64AddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 + (4.29497e+09+4.29497e+09i) (mismatched types float64 and complex128)`,
	)

}

// Test Float64 + Rune32
func TestCheckBinaryTypedExprFloat64AddRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + rune(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 + rune(2147483647) (mismatched types float64 and rune)`,
	)

}

// Test Float64 + StringT
func TestCheckBinaryTypedExprFloat64AddStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + string("abc")`, env,
		`invalid operation: 4.29497e+09 + "abc" (mismatched types float64 and string)`,
	)

}

// Test Float64 + BoolT
func TestCheckBinaryTypedExprFloat64AddBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) + bool(true)`, env,
		`invalid operation: 4.29497e+09 + true (mismatched types float64 and bool)`,
	)

}

// Test Float64 - Int
func TestCheckBinaryTypedExprFloat64SubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) - 4`, env, float64(0xffffffff) - 4, reflect.TypeOf(float64(0xffffffff) - 4))
}

// Test Float64 - Rune
func TestCheckBinaryTypedExprFloat64SubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) - '@'`, env, float64(0xffffffff) - '@', reflect.TypeOf(float64(0xffffffff) - '@'))
}

// Test Float64 - Float
func TestCheckBinaryTypedExprFloat64SubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) - 2.0`, env, float64(0xffffffff) - 2.0, reflect.TypeOf(float64(0xffffffff) - 2.0))
}

// Test Float64 - Complex
func TestCheckBinaryTypedExprFloat64SubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Float64 - Bool
func TestCheckBinaryTypedExprFloat64SubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 4.29497e+09 - true (mismatched types float64 and bool)`,
	)

}

// Test Float64 - String
func TestCheckBinaryTypedExprFloat64SubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 4.29497e+09 - "abc" (mismatched types float64 and string)`,
	)

}

// Test Float64 - Nil
func TestCheckBinaryTypedExprFloat64SubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - nil`, env,
		`cannot convert nil to type float64`,
	)

}

// Test Float64 - Int8
func TestCheckBinaryTypedExprFloat64SubInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - int8(0x7f)`, env,
		`invalid operation: 4.29497e+09 - 127 (mismatched types float64 and int8)`,
	)

}

// Test Float64 - Int16
func TestCheckBinaryTypedExprFloat64SubInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - int16(0x7fff)`, env,
		`invalid operation: 4.29497e+09 - 32767 (mismatched types float64 and int16)`,
	)

}

// Test Float64 - Int32
func TestCheckBinaryTypedExprFloat64SubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - int32(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 - 2147483647 (mismatched types float64 and int32)`,
	)

}

// Test Float64 - Int64
func TestCheckBinaryTypedExprFloat64SubInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 - 9223372036854775807 (mismatched types float64 and int64)`,
	)

}

// Test Float64 - Uint8
func TestCheckBinaryTypedExprFloat64SubUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - uint8(0xff)`, env,
		`invalid operation: 4.29497e+09 - 255 (mismatched types float64 and uint8)`,
	)

}

// Test Float64 - Uint16
func TestCheckBinaryTypedExprFloat64SubUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - uint16(0xffff)`, env,
		`invalid operation: 4.29497e+09 - 65535 (mismatched types float64 and uint16)`,
	)

}

// Test Float64 - Uint32
func TestCheckBinaryTypedExprFloat64SubUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - uint32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 - 4294967295 (mismatched types float64 and uint32)`,
	)

}

// Test Float64 - Uint64
func TestCheckBinaryTypedExprFloat64SubUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 - 18446744073709551615 (mismatched types float64 and uint64)`,
	)

}

// Test Float64 - Float32
func TestCheckBinaryTypedExprFloat64SubFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - float32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 - 4.29497e+09 (mismatched types float64 and float32)`,
	)

}

// Test Float64 - Float64
func TestCheckBinaryTypedExprFloat64SubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) - float64(0xffffffff)`, env, float64(0xffffffff) - float64(0xffffffff), reflect.TypeOf(float64(0xffffffff) - float64(0xffffffff)))
}

// Test Float64 - Complex64
func TestCheckBinaryTypedExprFloat64SubComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 - (4.29497e+09+4.29497e+09i) (mismatched types float64 and complex64)`,
	)

}

// Test Float64 - Complex128
func TestCheckBinaryTypedExprFloat64SubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 - (4.29497e+09+4.29497e+09i) (mismatched types float64 and complex128)`,
	)

}

// Test Float64 - Rune32
func TestCheckBinaryTypedExprFloat64SubRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - rune(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 - rune(2147483647) (mismatched types float64 and rune)`,
	)

}

// Test Float64 - StringT
func TestCheckBinaryTypedExprFloat64SubStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - string("abc")`, env,
		`invalid operation: 4.29497e+09 - "abc" (mismatched types float64 and string)`,
	)

}

// Test Float64 - BoolT
func TestCheckBinaryTypedExprFloat64SubBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) - bool(true)`, env,
		`invalid operation: 4.29497e+09 - true (mismatched types float64 and bool)`,
	)

}

// Test Float64 & Int
func TestCheckBinaryTypedExprFloat64AndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & 4`, env,
		`invalid operation: 4.29497e+09 & 4 (operator & not defined on float64)`,
	)

}

// Test Float64 & Rune
func TestCheckBinaryTypedExprFloat64AndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & '@'`, env,
		`invalid operation: 4.29497e+09 & 64 (operator & not defined on float64)`,
	)

}

// Test Float64 & Float
func TestCheckBinaryTypedExprFloat64AndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & 2.0`, env,
		`invalid operation: 4.29497e+09 & 2 (operator & not defined on float64)`,
	)

}

// Test Float64 & Complex
func TestCheckBinaryTypedExprFloat64AndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: 4.29497e+09 & 0 (operator & not defined on float64)`,
	)

}

// Test Float64 & Bool
func TestCheckBinaryTypedExprFloat64AndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 4.29497e+09 & true (mismatched types float64 and bool)`,
	)

}

// Test Float64 & String
func TestCheckBinaryTypedExprFloat64AndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 4.29497e+09 & "abc" (mismatched types float64 and string)`,
	)

}

// Test Float64 & Nil
func TestCheckBinaryTypedExprFloat64AndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & nil`, env,
		`cannot convert nil to type float64`,
	)

}

// Test Float64 & Int8
func TestCheckBinaryTypedExprFloat64AndInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & int8(0x7f)`, env,
		`invalid operation: 4.29497e+09 & 127 (mismatched types float64 and int8)`,
	)

}

// Test Float64 & Int16
func TestCheckBinaryTypedExprFloat64AndInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & int16(0x7fff)`, env,
		`invalid operation: 4.29497e+09 & 32767 (mismatched types float64 and int16)`,
	)

}

// Test Float64 & Int32
func TestCheckBinaryTypedExprFloat64AndInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & int32(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 & 2147483647 (mismatched types float64 and int32)`,
	)

}

// Test Float64 & Int64
func TestCheckBinaryTypedExprFloat64AndInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 & 9223372036854775807 (mismatched types float64 and int64)`,
	)

}

// Test Float64 & Uint8
func TestCheckBinaryTypedExprFloat64AndUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & uint8(0xff)`, env,
		`invalid operation: 4.29497e+09 & 255 (mismatched types float64 and uint8)`,
	)

}

// Test Float64 & Uint16
func TestCheckBinaryTypedExprFloat64AndUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & uint16(0xffff)`, env,
		`invalid operation: 4.29497e+09 & 65535 (mismatched types float64 and uint16)`,
	)

}

// Test Float64 & Uint32
func TestCheckBinaryTypedExprFloat64AndUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & uint32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 & 4294967295 (mismatched types float64 and uint32)`,
	)

}

// Test Float64 & Uint64
func TestCheckBinaryTypedExprFloat64AndUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 & 18446744073709551615 (mismatched types float64 and uint64)`,
	)

}

// Test Float64 & Float32
func TestCheckBinaryTypedExprFloat64AndFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & float32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 & 4.29497e+09 (mismatched types float64 and float32)`,
	)

}

// Test Float64 & Float64
func TestCheckBinaryTypedExprFloat64AndFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & float64(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 & 4.29497e+09 (operator & not defined on float64)`,
	)

}

// Test Float64 & Complex64
func TestCheckBinaryTypedExprFloat64AndComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 & (4.29497e+09+4.29497e+09i) (mismatched types float64 and complex64)`,
	)

}

// Test Float64 & Complex128
func TestCheckBinaryTypedExprFloat64AndComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 & (4.29497e+09+4.29497e+09i) (mismatched types float64 and complex128)`,
	)

}

// Test Float64 & Rune32
func TestCheckBinaryTypedExprFloat64AndRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & rune(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 & rune(2147483647) (mismatched types float64 and rune)`,
	)

}

// Test Float64 & StringT
func TestCheckBinaryTypedExprFloat64AndStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & string("abc")`, env,
		`invalid operation: 4.29497e+09 & "abc" (mismatched types float64 and string)`,
	)

}

// Test Float64 & BoolT
func TestCheckBinaryTypedExprFloat64AndBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) & bool(true)`, env,
		`invalid operation: 4.29497e+09 & true (mismatched types float64 and bool)`,
	)

}

// Test Float64 % Int
func TestCheckBinaryTypedExprFloat64RemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % 4`, env,
		`invalid operation: 4.29497e+09 % 4 (operator % not defined on float64)`,
	)

}

// Test Float64 % Rune
func TestCheckBinaryTypedExprFloat64RemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % '@'`, env,
		`invalid operation: 4.29497e+09 % 64 (operator % not defined on float64)`,
	)

}

// Test Float64 % Float
func TestCheckBinaryTypedExprFloat64RemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % 2.0`, env,
		`invalid operation: 4.29497e+09 % 2 (operator % not defined on float64)`,
	)

}

// Test Float64 % Complex
func TestCheckBinaryTypedExprFloat64RemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: 4.29497e+09 % 0 (operator % not defined on float64)`,
	)

}

// Test Float64 % Bool
func TestCheckBinaryTypedExprFloat64RemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 4.29497e+09 % true (mismatched types float64 and bool)`,
	)

}

// Test Float64 % String
func TestCheckBinaryTypedExprFloat64RemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 4.29497e+09 % "abc" (mismatched types float64 and string)`,
	)

}

// Test Float64 % Nil
func TestCheckBinaryTypedExprFloat64RemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % nil`, env,
		`cannot convert nil to type float64`,
	)

}

// Test Float64 % Int8
func TestCheckBinaryTypedExprFloat64RemInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % int8(0x7f)`, env,
		`invalid operation: 4.29497e+09 % 127 (mismatched types float64 and int8)`,
	)

}

// Test Float64 % Int16
func TestCheckBinaryTypedExprFloat64RemInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % int16(0x7fff)`, env,
		`invalid operation: 4.29497e+09 % 32767 (mismatched types float64 and int16)`,
	)

}

// Test Float64 % Int32
func TestCheckBinaryTypedExprFloat64RemInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % int32(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 % 2147483647 (mismatched types float64 and int32)`,
	)

}

// Test Float64 % Int64
func TestCheckBinaryTypedExprFloat64RemInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 % 9223372036854775807 (mismatched types float64 and int64)`,
	)

}

// Test Float64 % Uint8
func TestCheckBinaryTypedExprFloat64RemUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % uint8(0xff)`, env,
		`invalid operation: 4.29497e+09 % 255 (mismatched types float64 and uint8)`,
	)

}

// Test Float64 % Uint16
func TestCheckBinaryTypedExprFloat64RemUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % uint16(0xffff)`, env,
		`invalid operation: 4.29497e+09 % 65535 (mismatched types float64 and uint16)`,
	)

}

// Test Float64 % Uint32
func TestCheckBinaryTypedExprFloat64RemUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % uint32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 % 4294967295 (mismatched types float64 and uint32)`,
	)

}

// Test Float64 % Uint64
func TestCheckBinaryTypedExprFloat64RemUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 % 18446744073709551615 (mismatched types float64 and uint64)`,
	)

}

// Test Float64 % Float32
func TestCheckBinaryTypedExprFloat64RemFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % float32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 % 4.29497e+09 (mismatched types float64 and float32)`,
	)

}

// Test Float64 % Float64
func TestCheckBinaryTypedExprFloat64RemFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % float64(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 % 4.29497e+09 (operator % not defined on float64)`,
	)

}

// Test Float64 % Complex64
func TestCheckBinaryTypedExprFloat64RemComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 % (4.29497e+09+4.29497e+09i) (mismatched types float64 and complex64)`,
	)

}

// Test Float64 % Complex128
func TestCheckBinaryTypedExprFloat64RemComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 % (4.29497e+09+4.29497e+09i) (mismatched types float64 and complex128)`,
	)

}

// Test Float64 % Rune32
func TestCheckBinaryTypedExprFloat64RemRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % rune(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 % rune(2147483647) (mismatched types float64 and rune)`,
	)

}

// Test Float64 % StringT
func TestCheckBinaryTypedExprFloat64RemStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % string("abc")`, env,
		`invalid operation: 4.29497e+09 % "abc" (mismatched types float64 and string)`,
	)

}

// Test Float64 % BoolT
func TestCheckBinaryTypedExprFloat64RemBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) % bool(true)`, env,
		`invalid operation: 4.29497e+09 % true (mismatched types float64 and bool)`,
	)

}

// Test Float64 == Int
func TestCheckBinaryTypedExprFloat64EqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) == 4`, env, float64(0xffffffff) == 4, reflect.TypeOf(float64(0xffffffff) == 4))
}

// Test Float64 == Rune
func TestCheckBinaryTypedExprFloat64EqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) == '@'`, env, float64(0xffffffff) == '@', reflect.TypeOf(float64(0xffffffff) == '@'))
}

// Test Float64 == Float
func TestCheckBinaryTypedExprFloat64EqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) == 2.0`, env, float64(0xffffffff) == 2.0, reflect.TypeOf(float64(0xffffffff) == 2.0))
}

// Test Float64 == Complex
func TestCheckBinaryTypedExprFloat64EqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Float64 == Bool
func TestCheckBinaryTypedExprFloat64EqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 4.29497e+09 == true (mismatched types float64 and bool)`,
	)

}

// Test Float64 == String
func TestCheckBinaryTypedExprFloat64EqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 4.29497e+09 == "abc" (mismatched types float64 and string)`,
	)

}

// Test Float64 == Nil
func TestCheckBinaryTypedExprFloat64EqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == nil`, env,
		`cannot convert nil to type float64`,
	)

}

// Test Float64 == Int8
func TestCheckBinaryTypedExprFloat64EqlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == int8(0x7f)`, env,
		`invalid operation: 4.29497e+09 == 127 (mismatched types float64 and int8)`,
	)

}

// Test Float64 == Int16
func TestCheckBinaryTypedExprFloat64EqlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == int16(0x7fff)`, env,
		`invalid operation: 4.29497e+09 == 32767 (mismatched types float64 and int16)`,
	)

}

// Test Float64 == Int32
func TestCheckBinaryTypedExprFloat64EqlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == int32(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 == 2147483647 (mismatched types float64 and int32)`,
	)

}

// Test Float64 == Int64
func TestCheckBinaryTypedExprFloat64EqlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 == 9223372036854775807 (mismatched types float64 and int64)`,
	)

}

// Test Float64 == Uint8
func TestCheckBinaryTypedExprFloat64EqlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == uint8(0xff)`, env,
		`invalid operation: 4.29497e+09 == 255 (mismatched types float64 and uint8)`,
	)

}

// Test Float64 == Uint16
func TestCheckBinaryTypedExprFloat64EqlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == uint16(0xffff)`, env,
		`invalid operation: 4.29497e+09 == 65535 (mismatched types float64 and uint16)`,
	)

}

// Test Float64 == Uint32
func TestCheckBinaryTypedExprFloat64EqlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == uint32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 == 4294967295 (mismatched types float64 and uint32)`,
	)

}

// Test Float64 == Uint64
func TestCheckBinaryTypedExprFloat64EqlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 == 18446744073709551615 (mismatched types float64 and uint64)`,
	)

}

// Test Float64 == Float32
func TestCheckBinaryTypedExprFloat64EqlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == float32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 == 4.29497e+09 (mismatched types float64 and float32)`,
	)

}

// Test Float64 == Float64
func TestCheckBinaryTypedExprFloat64EqlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) == float64(0xffffffff)`, env, float64(0xffffffff) == float64(0xffffffff), reflect.TypeOf(float64(0xffffffff) == float64(0xffffffff)))
}

// Test Float64 == Complex64
func TestCheckBinaryTypedExprFloat64EqlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 == (4.29497e+09+4.29497e+09i) (mismatched types float64 and complex64)`,
	)

}

// Test Float64 == Complex128
func TestCheckBinaryTypedExprFloat64EqlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 == (4.29497e+09+4.29497e+09i) (mismatched types float64 and complex128)`,
	)

}

// Test Float64 == Rune32
func TestCheckBinaryTypedExprFloat64EqlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == rune(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 == rune(2147483647) (mismatched types float64 and rune)`,
	)

}

// Test Float64 == StringT
func TestCheckBinaryTypedExprFloat64EqlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == string("abc")`, env,
		`invalid operation: 4.29497e+09 == "abc" (mismatched types float64 and string)`,
	)

}

// Test Float64 == BoolT
func TestCheckBinaryTypedExprFloat64EqlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) == bool(true)`, env,
		`invalid operation: 4.29497e+09 == true (mismatched types float64 and bool)`,
	)

}

// Test Float64 > Int
func TestCheckBinaryTypedExprFloat64GtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) > 4`, env, float64(0xffffffff) > 4, reflect.TypeOf(float64(0xffffffff) > 4))
}

// Test Float64 > Rune
func TestCheckBinaryTypedExprFloat64GtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) > '@'`, env, float64(0xffffffff) > '@', reflect.TypeOf(float64(0xffffffff) > '@'))
}

// Test Float64 > Float
func TestCheckBinaryTypedExprFloat64GtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) > 2.0`, env, float64(0xffffffff) > 2.0, reflect.TypeOf(float64(0xffffffff) > 2.0))
}

// Test Float64 > Complex
func TestCheckBinaryTypedExprFloat64GtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Float64 > Bool
func TestCheckBinaryTypedExprFloat64GtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > true`, env,
		`cannot convert true to type float64`,
		`invalid operation: 4.29497e+09 > true (mismatched types float64 and bool)`,
	)

}

// Test Float64 > String
func TestCheckBinaryTypedExprFloat64GtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > "abc"`, env,
		`cannot convert "abc" to type float64`,
		`invalid operation: 4.29497e+09 > "abc" (mismatched types float64 and string)`,
	)

}

// Test Float64 > Nil
func TestCheckBinaryTypedExprFloat64GtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > nil`, env,
		`cannot convert nil to type float64`,
	)

}

// Test Float64 > Int8
func TestCheckBinaryTypedExprFloat64GtrInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > int8(0x7f)`, env,
		`invalid operation: 4.29497e+09 > 127 (mismatched types float64 and int8)`,
	)

}

// Test Float64 > Int16
func TestCheckBinaryTypedExprFloat64GtrInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > int16(0x7fff)`, env,
		`invalid operation: 4.29497e+09 > 32767 (mismatched types float64 and int16)`,
	)

}

// Test Float64 > Int32
func TestCheckBinaryTypedExprFloat64GtrInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > int32(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 > 2147483647 (mismatched types float64 and int32)`,
	)

}

// Test Float64 > Int64
func TestCheckBinaryTypedExprFloat64GtrInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 > 9223372036854775807 (mismatched types float64 and int64)`,
	)

}

// Test Float64 > Uint8
func TestCheckBinaryTypedExprFloat64GtrUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > uint8(0xff)`, env,
		`invalid operation: 4.29497e+09 > 255 (mismatched types float64 and uint8)`,
	)

}

// Test Float64 > Uint16
func TestCheckBinaryTypedExprFloat64GtrUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > uint16(0xffff)`, env,
		`invalid operation: 4.29497e+09 > 65535 (mismatched types float64 and uint16)`,
	)

}

// Test Float64 > Uint32
func TestCheckBinaryTypedExprFloat64GtrUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > uint32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 > 4294967295 (mismatched types float64 and uint32)`,
	)

}

// Test Float64 > Uint64
func TestCheckBinaryTypedExprFloat64GtrUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 > 18446744073709551615 (mismatched types float64 and uint64)`,
	)

}

// Test Float64 > Float32
func TestCheckBinaryTypedExprFloat64GtrFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > float32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 > 4.29497e+09 (mismatched types float64 and float32)`,
	)

}

// Test Float64 > Float64
func TestCheckBinaryTypedExprFloat64GtrFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff) > float64(0xffffffff)`, env, float64(0xffffffff) > float64(0xffffffff), reflect.TypeOf(float64(0xffffffff) > float64(0xffffffff)))
}

// Test Float64 > Complex64
func TestCheckBinaryTypedExprFloat64GtrComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 > (4.29497e+09+4.29497e+09i) (mismatched types float64 and complex64)`,
	)

}

// Test Float64 > Complex128
func TestCheckBinaryTypedExprFloat64GtrComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 > (4.29497e+09+4.29497e+09i) (mismatched types float64 and complex128)`,
	)

}

// Test Float64 > Rune32
func TestCheckBinaryTypedExprFloat64GtrRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > rune(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 > rune(2147483647) (mismatched types float64 and rune)`,
	)

}

// Test Float64 > StringT
func TestCheckBinaryTypedExprFloat64GtrStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > string("abc")`, env,
		`invalid operation: 4.29497e+09 > "abc" (mismatched types float64 and string)`,
	)

}

// Test Float64 > BoolT
func TestCheckBinaryTypedExprFloat64GtrBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) > bool(true)`, env,
		`invalid operation: 4.29497e+09 > true (mismatched types float64 and bool)`,
	)

}

// Test Float64 << Int
func TestCheckBinaryTypedExprFloat64ShlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << 4`, env,
		`invalid operation: 4.29497e+09 << 4 (shift of type float64)`,
	)

}

// Test Float64 << Rune
func TestCheckBinaryTypedExprFloat64ShlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << '@'`, env,
		`invalid operation: 4.29497e+09 << 64 (shift of type float64)`,
	)

}

// Test Float64 << Float
func TestCheckBinaryTypedExprFloat64ShlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << 2.0`, env,
		`invalid operation: 4.29497e+09 << 2 (shift of type float64)`,
	)

}

// Test Float64 << Complex
func TestCheckBinaryTypedExprFloat64ShlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: 4.29497e+09 << 0 (shift of type float64)`,
	)

}

// Test Float64 << Bool
func TestCheckBinaryTypedExprFloat64ShlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << true`, env,
		`invalid operation: 4.29497e+09 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Float64 << String
func TestCheckBinaryTypedExprFloat64ShlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: 4.29497e+09 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Float64 << Nil
func TestCheckBinaryTypedExprFloat64ShlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Float64 << Int8
func TestCheckBinaryTypedExprFloat64ShlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << int8(0x7f)`, env,
		`invalid operation: 4.29497e+09 << 127 (shift count type int8, must be unsigned integer)`,
	)

}

// Test Float64 << Int16
func TestCheckBinaryTypedExprFloat64ShlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << int16(0x7fff)`, env,
		`invalid operation: 4.29497e+09 << 32767 (shift count type int16, must be unsigned integer)`,
	)

}

// Test Float64 << Int32
func TestCheckBinaryTypedExprFloat64ShlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << int32(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 << 2147483647 (shift count type int32, must be unsigned integer)`,
	)

}

// Test Float64 << Int64
func TestCheckBinaryTypedExprFloat64ShlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << int64(0x7fffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 << 9223372036854775807 (shift count type int64, must be unsigned integer)`,
	)

}

// Test Float64 << Uint8
func TestCheckBinaryTypedExprFloat64ShlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << uint8(0xff)`, env,
		`invalid operation: 4.29497e+09 << 255 (shift of type float64)`,
	)

}

// Test Float64 << Uint16
func TestCheckBinaryTypedExprFloat64ShlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << uint16(0xffff)`, env,
		`invalid operation: 4.29497e+09 << 65535 (shift of type float64)`,
	)

}

// Test Float64 << Uint32
func TestCheckBinaryTypedExprFloat64ShlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << uint32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 << 4294967295 (shift of type float64)`,
	)

}

// Test Float64 << Uint64
func TestCheckBinaryTypedExprFloat64ShlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << uint64(0xffffffffffffffff)`, env,
		`invalid operation: 4.29497e+09 << 18446744073709551615 (shift of type float64)`,
	)

}

// Test Float64 << Float32
func TestCheckBinaryTypedExprFloat64ShlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << float32(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 << 4.29497e+09 (shift count type float32, must be unsigned integer)`,
	)

}

// Test Float64 << Float64
func TestCheckBinaryTypedExprFloat64ShlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << float64(0xffffffff)`, env,
		`invalid operation: 4.29497e+09 << 4.29497e+09 (shift count type float64, must be unsigned integer)`,
	)

}

// Test Float64 << Complex64
func TestCheckBinaryTypedExprFloat64ShlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 << (4.29497e+09+4.29497e+09i) (shift count type complex64, must be unsigned integer)`,
	)

}

// Test Float64 << Complex128
func TestCheckBinaryTypedExprFloat64ShlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: 4.29497e+09 << (4.29497e+09+4.29497e+09i) (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Float64 << Rune32
func TestCheckBinaryTypedExprFloat64ShlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << rune(0x7fffffff)`, env,
		`invalid operation: 4.29497e+09 << rune(2147483647) (shift count type rune, must be unsigned integer)`,
	)

}

// Test Float64 << StringT
func TestCheckBinaryTypedExprFloat64ShlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << string("abc")`, env,
		`invalid operation: 4.29497e+09 << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Float64 << BoolT
func TestCheckBinaryTypedExprFloat64ShlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(0xffffffff) << bool(true)`, env,
		`invalid operation: 4.29497e+09 << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Complex64 + Int
func TestCheckBinaryTypedExprComplex64AddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff + 0xffffffff * 1i) + 4`, env, complex64(0xffffffff + 0xffffffff * 1i) + 4, reflect.TypeOf(complex64(0xffffffff + 0xffffffff * 1i) + 4))
}

// Test Complex64 + Rune
func TestCheckBinaryTypedExprComplex64AddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff + 0xffffffff * 1i) + '@'`, env, complex64(0xffffffff + 0xffffffff * 1i) + '@', reflect.TypeOf(complex64(0xffffffff + 0xffffffff * 1i) + '@'))
}

// Test Complex64 + Float
func TestCheckBinaryTypedExprComplex64AddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff + 0xffffffff * 1i) + 2.0`, env, complex64(0xffffffff + 0xffffffff * 1i) + 2.0, reflect.TypeOf(complex64(0xffffffff + 0xffffffff * 1i) + 2.0))
}

// Test Complex64 + Complex
func TestCheckBinaryTypedExprComplex64AddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff + 0xffffffff * 1i) + 8.0i`, env, complex64(0xffffffff + 0xffffffff * 1i) + 8.0i, reflect.TypeOf(complex64(0xffffffff + 0xffffffff * 1i) + 8.0i))
}

// Test Complex64 + Bool
func TestCheckBinaryTypedExprComplex64AddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + true`, env,
		`cannot convert true to type complex64`,
		`invalid operation: (4.29497e+09+4.29497e+09i) + true (mismatched types complex64 and bool)`,
	)

}

// Test Complex64 + String
func TestCheckBinaryTypedExprComplex64AddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + "abc"`, env,
		`cannot convert "abc" to type complex64`,
		`invalid operation: (4.29497e+09+4.29497e+09i) + "abc" (mismatched types complex64 and string)`,
	)

}

// Test Complex64 + Nil
func TestCheckBinaryTypedExprComplex64AddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + nil`, env,
		`cannot convert nil to type complex64`,
	)

}

// Test Complex64 + Int8
func TestCheckBinaryTypedExprComplex64AddInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + int8(0x7f)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 127 (mismatched types complex64 and int8)`,
	)

}

// Test Complex64 + Int16
func TestCheckBinaryTypedExprComplex64AddInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + int16(0x7fff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 32767 (mismatched types complex64 and int16)`,
	)

}

// Test Complex64 + Int32
func TestCheckBinaryTypedExprComplex64AddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + int32(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 2147483647 (mismatched types complex64 and int32)`,
	)

}

// Test Complex64 + Int64
func TestCheckBinaryTypedExprComplex64AddInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + int64(0x7fffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 9223372036854775807 (mismatched types complex64 and int64)`,
	)

}

// Test Complex64 + Uint8
func TestCheckBinaryTypedExprComplex64AddUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + uint8(0xff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 255 (mismatched types complex64 and uint8)`,
	)

}

// Test Complex64 + Uint16
func TestCheckBinaryTypedExprComplex64AddUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + uint16(0xffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 65535 (mismatched types complex64 and uint16)`,
	)

}

// Test Complex64 + Uint32
func TestCheckBinaryTypedExprComplex64AddUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + uint32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 4294967295 (mismatched types complex64 and uint32)`,
	)

}

// Test Complex64 + Uint64
func TestCheckBinaryTypedExprComplex64AddUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + uint64(0xffffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 18446744073709551615 (mismatched types complex64 and uint64)`,
	)

}

// Test Complex64 + Float32
func TestCheckBinaryTypedExprComplex64AddFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + float32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 4.29497e+09 (mismatched types complex64 and float32)`,
	)

}

// Test Complex64 + Float64
func TestCheckBinaryTypedExprComplex64AddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + float64(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 4.29497e+09 (mismatched types complex64 and float64)`,
	)

}

// Test Complex64 + Complex64
func TestCheckBinaryTypedExprComplex64AddComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff + 0xffffffff * 1i) + complex64(0xffffffff + 0xffffffff * 1i)`, env, complex64(0xffffffff + 0xffffffff * 1i) + complex64(0xffffffff + 0xffffffff * 1i), reflect.TypeOf(complex64(0xffffffff + 0xffffffff * 1i) + complex64(0xffffffff + 0xffffffff * 1i)))
}

// Test Complex64 + Complex128
func TestCheckBinaryTypedExprComplex64AddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + (4.29497e+09+4.29497e+09i) (mismatched types complex64 and complex128)`,
	)

}

// Test Complex64 + Rune32
func TestCheckBinaryTypedExprComplex64AddRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + rune(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + rune(2147483647) (mismatched types complex64 and rune)`,
	)

}

// Test Complex64 + StringT
func TestCheckBinaryTypedExprComplex64AddStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + string("abc")`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + "abc" (mismatched types complex64 and string)`,
	)

}

// Test Complex64 + BoolT
func TestCheckBinaryTypedExprComplex64AddBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) + bool(true)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + true (mismatched types complex64 and bool)`,
	)

}

// Test Complex64 - Int
func TestCheckBinaryTypedExprComplex64SubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff + 0xffffffff * 1i) - 4`, env, complex64(0xffffffff + 0xffffffff * 1i) - 4, reflect.TypeOf(complex64(0xffffffff + 0xffffffff * 1i) - 4))
}

// Test Complex64 - Rune
func TestCheckBinaryTypedExprComplex64SubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff + 0xffffffff * 1i) - '@'`, env, complex64(0xffffffff + 0xffffffff * 1i) - '@', reflect.TypeOf(complex64(0xffffffff + 0xffffffff * 1i) - '@'))
}

// Test Complex64 - Float
func TestCheckBinaryTypedExprComplex64SubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff + 0xffffffff * 1i) - 2.0`, env, complex64(0xffffffff + 0xffffffff * 1i) - 2.0, reflect.TypeOf(complex64(0xffffffff + 0xffffffff * 1i) - 2.0))
}

// Test Complex64 - Complex
func TestCheckBinaryTypedExprComplex64SubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff + 0xffffffff * 1i) - 8.0i`, env, complex64(0xffffffff + 0xffffffff * 1i) - 8.0i, reflect.TypeOf(complex64(0xffffffff + 0xffffffff * 1i) - 8.0i))
}

// Test Complex64 - Bool
func TestCheckBinaryTypedExprComplex64SubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - true`, env,
		`cannot convert true to type complex64`,
		`invalid operation: (4.29497e+09+4.29497e+09i) - true (mismatched types complex64 and bool)`,
	)

}

// Test Complex64 - String
func TestCheckBinaryTypedExprComplex64SubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - "abc"`, env,
		`cannot convert "abc" to type complex64`,
		`invalid operation: (4.29497e+09+4.29497e+09i) - "abc" (mismatched types complex64 and string)`,
	)

}

// Test Complex64 - Nil
func TestCheckBinaryTypedExprComplex64SubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - nil`, env,
		`cannot convert nil to type complex64`,
	)

}

// Test Complex64 - Int8
func TestCheckBinaryTypedExprComplex64SubInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - int8(0x7f)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 127 (mismatched types complex64 and int8)`,
	)

}

// Test Complex64 - Int16
func TestCheckBinaryTypedExprComplex64SubInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - int16(0x7fff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 32767 (mismatched types complex64 and int16)`,
	)

}

// Test Complex64 - Int32
func TestCheckBinaryTypedExprComplex64SubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - int32(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 2147483647 (mismatched types complex64 and int32)`,
	)

}

// Test Complex64 - Int64
func TestCheckBinaryTypedExprComplex64SubInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - int64(0x7fffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 9223372036854775807 (mismatched types complex64 and int64)`,
	)

}

// Test Complex64 - Uint8
func TestCheckBinaryTypedExprComplex64SubUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - uint8(0xff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 255 (mismatched types complex64 and uint8)`,
	)

}

// Test Complex64 - Uint16
func TestCheckBinaryTypedExprComplex64SubUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - uint16(0xffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 65535 (mismatched types complex64 and uint16)`,
	)

}

// Test Complex64 - Uint32
func TestCheckBinaryTypedExprComplex64SubUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - uint32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 4294967295 (mismatched types complex64 and uint32)`,
	)

}

// Test Complex64 - Uint64
func TestCheckBinaryTypedExprComplex64SubUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - uint64(0xffffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 18446744073709551615 (mismatched types complex64 and uint64)`,
	)

}

// Test Complex64 - Float32
func TestCheckBinaryTypedExprComplex64SubFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - float32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 4.29497e+09 (mismatched types complex64 and float32)`,
	)

}

// Test Complex64 - Float64
func TestCheckBinaryTypedExprComplex64SubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - float64(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 4.29497e+09 (mismatched types complex64 and float64)`,
	)

}

// Test Complex64 - Complex64
func TestCheckBinaryTypedExprComplex64SubComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff + 0xffffffff * 1i) - complex64(0xffffffff + 0xffffffff * 1i)`, env, complex64(0xffffffff + 0xffffffff * 1i) - complex64(0xffffffff + 0xffffffff * 1i), reflect.TypeOf(complex64(0xffffffff + 0xffffffff * 1i) - complex64(0xffffffff + 0xffffffff * 1i)))
}

// Test Complex64 - Complex128
func TestCheckBinaryTypedExprComplex64SubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - (4.29497e+09+4.29497e+09i) (mismatched types complex64 and complex128)`,
	)

}

// Test Complex64 - Rune32
func TestCheckBinaryTypedExprComplex64SubRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - rune(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - rune(2147483647) (mismatched types complex64 and rune)`,
	)

}

// Test Complex64 - StringT
func TestCheckBinaryTypedExprComplex64SubStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - string("abc")`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - "abc" (mismatched types complex64 and string)`,
	)

}

// Test Complex64 - BoolT
func TestCheckBinaryTypedExprComplex64SubBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) - bool(true)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - true (mismatched types complex64 and bool)`,
	)

}

// Test Complex64 & Int
func TestCheckBinaryTypedExprComplex64AndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & 4`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 4 (operator & not defined on complex64)`,
	)

}

// Test Complex64 & Rune
func TestCheckBinaryTypedExprComplex64AndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & '@'`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 64 (operator & not defined on complex64)`,
	)

}

// Test Complex64 & Float
func TestCheckBinaryTypedExprComplex64AndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & 2.0`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 2 (operator & not defined on complex64)`,
	)

}

// Test Complex64 & Complex
func TestCheckBinaryTypedExprComplex64AndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & 8.0i`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 8i (operator & not defined on complex64)`,
	)

}

// Test Complex64 & Bool
func TestCheckBinaryTypedExprComplex64AndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & true`, env,
		`cannot convert true to type complex64`,
		`invalid operation: (4.29497e+09+4.29497e+09i) & true (mismatched types complex64 and bool)`,
	)

}

// Test Complex64 & String
func TestCheckBinaryTypedExprComplex64AndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & "abc"`, env,
		`cannot convert "abc" to type complex64`,
		`invalid operation: (4.29497e+09+4.29497e+09i) & "abc" (mismatched types complex64 and string)`,
	)

}

// Test Complex64 & Nil
func TestCheckBinaryTypedExprComplex64AndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & nil`, env,
		`cannot convert nil to type complex64`,
	)

}

// Test Complex64 & Int8
func TestCheckBinaryTypedExprComplex64AndInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & int8(0x7f)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 127 (mismatched types complex64 and int8)`,
	)

}

// Test Complex64 & Int16
func TestCheckBinaryTypedExprComplex64AndInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & int16(0x7fff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 32767 (mismatched types complex64 and int16)`,
	)

}

// Test Complex64 & Int32
func TestCheckBinaryTypedExprComplex64AndInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & int32(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 2147483647 (mismatched types complex64 and int32)`,
	)

}

// Test Complex64 & Int64
func TestCheckBinaryTypedExprComplex64AndInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & int64(0x7fffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 9223372036854775807 (mismatched types complex64 and int64)`,
	)

}

// Test Complex64 & Uint8
func TestCheckBinaryTypedExprComplex64AndUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & uint8(0xff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 255 (mismatched types complex64 and uint8)`,
	)

}

// Test Complex64 & Uint16
func TestCheckBinaryTypedExprComplex64AndUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & uint16(0xffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 65535 (mismatched types complex64 and uint16)`,
	)

}

// Test Complex64 & Uint32
func TestCheckBinaryTypedExprComplex64AndUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & uint32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 4294967295 (mismatched types complex64 and uint32)`,
	)

}

// Test Complex64 & Uint64
func TestCheckBinaryTypedExprComplex64AndUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & uint64(0xffffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 18446744073709551615 (mismatched types complex64 and uint64)`,
	)

}

// Test Complex64 & Float32
func TestCheckBinaryTypedExprComplex64AndFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & float32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 4.29497e+09 (mismatched types complex64 and float32)`,
	)

}

// Test Complex64 & Float64
func TestCheckBinaryTypedExprComplex64AndFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & float64(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 4.29497e+09 (mismatched types complex64 and float64)`,
	)

}

// Test Complex64 & Complex64
func TestCheckBinaryTypedExprComplex64AndComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & (4.29497e+09+4.29497e+09i) (operator & not defined on complex64)`,
	)

}

// Test Complex64 & Complex128
func TestCheckBinaryTypedExprComplex64AndComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & (4.29497e+09+4.29497e+09i) (mismatched types complex64 and complex128)`,
	)

}

// Test Complex64 & Rune32
func TestCheckBinaryTypedExprComplex64AndRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & rune(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & rune(2147483647) (mismatched types complex64 and rune)`,
	)

}

// Test Complex64 & StringT
func TestCheckBinaryTypedExprComplex64AndStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & string("abc")`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & "abc" (mismatched types complex64 and string)`,
	)

}

// Test Complex64 & BoolT
func TestCheckBinaryTypedExprComplex64AndBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) & bool(true)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & true (mismatched types complex64 and bool)`,
	)

}

// Test Complex64 % Int
func TestCheckBinaryTypedExprComplex64RemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % 4`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 4 (operator % not defined on complex64)`,
	)

}

// Test Complex64 % Rune
func TestCheckBinaryTypedExprComplex64RemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % '@'`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 64 (operator % not defined on complex64)`,
	)

}

// Test Complex64 % Float
func TestCheckBinaryTypedExprComplex64RemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % 2.0`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 2 (operator % not defined on complex64)`,
	)

}

// Test Complex64 % Complex
func TestCheckBinaryTypedExprComplex64RemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % 8.0i`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 8i (operator % not defined on complex64)`,
	)

}

// Test Complex64 % Bool
func TestCheckBinaryTypedExprComplex64RemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % true`, env,
		`cannot convert true to type complex64`,
		`invalid operation: (4.29497e+09+4.29497e+09i) % true (mismatched types complex64 and bool)`,
	)

}

// Test Complex64 % String
func TestCheckBinaryTypedExprComplex64RemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % "abc"`, env,
		`cannot convert "abc" to type complex64`,
		`invalid operation: (4.29497e+09+4.29497e+09i) % "abc" (mismatched types complex64 and string)`,
	)

}

// Test Complex64 % Nil
func TestCheckBinaryTypedExprComplex64RemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % nil`, env,
		`cannot convert nil to type complex64`,
	)

}

// Test Complex64 % Int8
func TestCheckBinaryTypedExprComplex64RemInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % int8(0x7f)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 127 (mismatched types complex64 and int8)`,
	)

}

// Test Complex64 % Int16
func TestCheckBinaryTypedExprComplex64RemInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % int16(0x7fff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 32767 (mismatched types complex64 and int16)`,
	)

}

// Test Complex64 % Int32
func TestCheckBinaryTypedExprComplex64RemInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % int32(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 2147483647 (mismatched types complex64 and int32)`,
	)

}

// Test Complex64 % Int64
func TestCheckBinaryTypedExprComplex64RemInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % int64(0x7fffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 9223372036854775807 (mismatched types complex64 and int64)`,
	)

}

// Test Complex64 % Uint8
func TestCheckBinaryTypedExprComplex64RemUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % uint8(0xff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 255 (mismatched types complex64 and uint8)`,
	)

}

// Test Complex64 % Uint16
func TestCheckBinaryTypedExprComplex64RemUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % uint16(0xffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 65535 (mismatched types complex64 and uint16)`,
	)

}

// Test Complex64 % Uint32
func TestCheckBinaryTypedExprComplex64RemUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % uint32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 4294967295 (mismatched types complex64 and uint32)`,
	)

}

// Test Complex64 % Uint64
func TestCheckBinaryTypedExprComplex64RemUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % uint64(0xffffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 18446744073709551615 (mismatched types complex64 and uint64)`,
	)

}

// Test Complex64 % Float32
func TestCheckBinaryTypedExprComplex64RemFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % float32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 4.29497e+09 (mismatched types complex64 and float32)`,
	)

}

// Test Complex64 % Float64
func TestCheckBinaryTypedExprComplex64RemFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % float64(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 4.29497e+09 (mismatched types complex64 and float64)`,
	)

}

// Test Complex64 % Complex64
func TestCheckBinaryTypedExprComplex64RemComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % (4.29497e+09+4.29497e+09i) (operator % not defined on complex64)`,
	)

}

// Test Complex64 % Complex128
func TestCheckBinaryTypedExprComplex64RemComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % (4.29497e+09+4.29497e+09i) (mismatched types complex64 and complex128)`,
	)

}

// Test Complex64 % Rune32
func TestCheckBinaryTypedExprComplex64RemRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % rune(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % rune(2147483647) (mismatched types complex64 and rune)`,
	)

}

// Test Complex64 % StringT
func TestCheckBinaryTypedExprComplex64RemStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % string("abc")`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % "abc" (mismatched types complex64 and string)`,
	)

}

// Test Complex64 % BoolT
func TestCheckBinaryTypedExprComplex64RemBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) % bool(true)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % true (mismatched types complex64 and bool)`,
	)

}

// Test Complex64 == Int
func TestCheckBinaryTypedExprComplex64EqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff + 0xffffffff * 1i) == 4`, env, complex64(0xffffffff + 0xffffffff * 1i) == 4, reflect.TypeOf(complex64(0xffffffff + 0xffffffff * 1i) == 4))
}

// Test Complex64 == Rune
func TestCheckBinaryTypedExprComplex64EqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff + 0xffffffff * 1i) == '@'`, env, complex64(0xffffffff + 0xffffffff * 1i) == '@', reflect.TypeOf(complex64(0xffffffff + 0xffffffff * 1i) == '@'))
}

// Test Complex64 == Float
func TestCheckBinaryTypedExprComplex64EqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff + 0xffffffff * 1i) == 2.0`, env, complex64(0xffffffff + 0xffffffff * 1i) == 2.0, reflect.TypeOf(complex64(0xffffffff + 0xffffffff * 1i) == 2.0))
}

// Test Complex64 == Complex
func TestCheckBinaryTypedExprComplex64EqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff + 0xffffffff * 1i) == 8.0i`, env, complex64(0xffffffff + 0xffffffff * 1i) == 8.0i, reflect.TypeOf(complex64(0xffffffff + 0xffffffff * 1i) == 8.0i))
}

// Test Complex64 == Bool
func TestCheckBinaryTypedExprComplex64EqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == true`, env,
		`cannot convert true to type complex64`,
		`invalid operation: (4.29497e+09+4.29497e+09i) == true (mismatched types complex64 and bool)`,
	)

}

// Test Complex64 == String
func TestCheckBinaryTypedExprComplex64EqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == "abc"`, env,
		`cannot convert "abc" to type complex64`,
		`invalid operation: (4.29497e+09+4.29497e+09i) == "abc" (mismatched types complex64 and string)`,
	)

}

// Test Complex64 == Nil
func TestCheckBinaryTypedExprComplex64EqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == nil`, env,
		`cannot convert nil to type complex64`,
	)

}

// Test Complex64 == Int8
func TestCheckBinaryTypedExprComplex64EqlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == int8(0x7f)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 127 (mismatched types complex64 and int8)`,
	)

}

// Test Complex64 == Int16
func TestCheckBinaryTypedExprComplex64EqlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == int16(0x7fff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 32767 (mismatched types complex64 and int16)`,
	)

}

// Test Complex64 == Int32
func TestCheckBinaryTypedExprComplex64EqlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == int32(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 2147483647 (mismatched types complex64 and int32)`,
	)

}

// Test Complex64 == Int64
func TestCheckBinaryTypedExprComplex64EqlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == int64(0x7fffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 9223372036854775807 (mismatched types complex64 and int64)`,
	)

}

// Test Complex64 == Uint8
func TestCheckBinaryTypedExprComplex64EqlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == uint8(0xff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 255 (mismatched types complex64 and uint8)`,
	)

}

// Test Complex64 == Uint16
func TestCheckBinaryTypedExprComplex64EqlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == uint16(0xffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 65535 (mismatched types complex64 and uint16)`,
	)

}

// Test Complex64 == Uint32
func TestCheckBinaryTypedExprComplex64EqlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == uint32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 4294967295 (mismatched types complex64 and uint32)`,
	)

}

// Test Complex64 == Uint64
func TestCheckBinaryTypedExprComplex64EqlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == uint64(0xffffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 18446744073709551615 (mismatched types complex64 and uint64)`,
	)

}

// Test Complex64 == Float32
func TestCheckBinaryTypedExprComplex64EqlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == float32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 4.29497e+09 (mismatched types complex64 and float32)`,
	)

}

// Test Complex64 == Float64
func TestCheckBinaryTypedExprComplex64EqlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == float64(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 4.29497e+09 (mismatched types complex64 and float64)`,
	)

}

// Test Complex64 == Complex64
func TestCheckBinaryTypedExprComplex64EqlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff + 0xffffffff * 1i) == complex64(0xffffffff + 0xffffffff * 1i)`, env, complex64(0xffffffff + 0xffffffff * 1i) == complex64(0xffffffff + 0xffffffff * 1i), reflect.TypeOf(complex64(0xffffffff + 0xffffffff * 1i) == complex64(0xffffffff + 0xffffffff * 1i)))
}

// Test Complex64 == Complex128
func TestCheckBinaryTypedExprComplex64EqlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == (4.29497e+09+4.29497e+09i) (mismatched types complex64 and complex128)`,
	)

}

// Test Complex64 == Rune32
func TestCheckBinaryTypedExprComplex64EqlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == rune(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == rune(2147483647) (mismatched types complex64 and rune)`,
	)

}

// Test Complex64 == StringT
func TestCheckBinaryTypedExprComplex64EqlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == string("abc")`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == "abc" (mismatched types complex64 and string)`,
	)

}

// Test Complex64 == BoolT
func TestCheckBinaryTypedExprComplex64EqlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) == bool(true)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == true (mismatched types complex64 and bool)`,
	)

}

// Test Complex64 > Int
func TestCheckBinaryTypedExprComplex64GtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > 4`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 4 (operator > not defined on complex64)`,
	)

}

// Test Complex64 > Rune
func TestCheckBinaryTypedExprComplex64GtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > '@'`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 64 (operator > not defined on complex64)`,
	)

}

// Test Complex64 > Float
func TestCheckBinaryTypedExprComplex64GtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > 2.0`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 2 (operator > not defined on complex64)`,
	)

}

// Test Complex64 > Complex
func TestCheckBinaryTypedExprComplex64GtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > 8.0i`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 8i (operator > not defined on complex64)`,
	)

}

// Test Complex64 > Bool
func TestCheckBinaryTypedExprComplex64GtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > true`, env,
		`cannot convert true to type complex64`,
		`invalid operation: (4.29497e+09+4.29497e+09i) > true (mismatched types complex64 and bool)`,
	)

}

// Test Complex64 > String
func TestCheckBinaryTypedExprComplex64GtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > "abc"`, env,
		`cannot convert "abc" to type complex64`,
		`invalid operation: (4.29497e+09+4.29497e+09i) > "abc" (mismatched types complex64 and string)`,
	)

}

// Test Complex64 > Nil
func TestCheckBinaryTypedExprComplex64GtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > nil`, env,
		`cannot convert nil to type complex64`,
	)

}

// Test Complex64 > Int8
func TestCheckBinaryTypedExprComplex64GtrInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > int8(0x7f)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 127 (mismatched types complex64 and int8)`,
	)

}

// Test Complex64 > Int16
func TestCheckBinaryTypedExprComplex64GtrInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > int16(0x7fff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 32767 (mismatched types complex64 and int16)`,
	)

}

// Test Complex64 > Int32
func TestCheckBinaryTypedExprComplex64GtrInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > int32(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 2147483647 (mismatched types complex64 and int32)`,
	)

}

// Test Complex64 > Int64
func TestCheckBinaryTypedExprComplex64GtrInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > int64(0x7fffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 9223372036854775807 (mismatched types complex64 and int64)`,
	)

}

// Test Complex64 > Uint8
func TestCheckBinaryTypedExprComplex64GtrUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > uint8(0xff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 255 (mismatched types complex64 and uint8)`,
	)

}

// Test Complex64 > Uint16
func TestCheckBinaryTypedExprComplex64GtrUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > uint16(0xffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 65535 (mismatched types complex64 and uint16)`,
	)

}

// Test Complex64 > Uint32
func TestCheckBinaryTypedExprComplex64GtrUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > uint32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 4294967295 (mismatched types complex64 and uint32)`,
	)

}

// Test Complex64 > Uint64
func TestCheckBinaryTypedExprComplex64GtrUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > uint64(0xffffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 18446744073709551615 (mismatched types complex64 and uint64)`,
	)

}

// Test Complex64 > Float32
func TestCheckBinaryTypedExprComplex64GtrFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > float32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 4.29497e+09 (mismatched types complex64 and float32)`,
	)

}

// Test Complex64 > Float64
func TestCheckBinaryTypedExprComplex64GtrFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > float64(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 4.29497e+09 (mismatched types complex64 and float64)`,
	)

}

// Test Complex64 > Complex64
func TestCheckBinaryTypedExprComplex64GtrComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > (4.29497e+09+4.29497e+09i) (operator > not defined on complex64)`,
	)

}

// Test Complex64 > Complex128
func TestCheckBinaryTypedExprComplex64GtrComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > (4.29497e+09+4.29497e+09i) (mismatched types complex64 and complex128)`,
	)

}

// Test Complex64 > Rune32
func TestCheckBinaryTypedExprComplex64GtrRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > rune(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > rune(2147483647) (mismatched types complex64 and rune)`,
	)

}

// Test Complex64 > StringT
func TestCheckBinaryTypedExprComplex64GtrStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > string("abc")`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > "abc" (mismatched types complex64 and string)`,
	)

}

// Test Complex64 > BoolT
func TestCheckBinaryTypedExprComplex64GtrBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) > bool(true)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > true (mismatched types complex64 and bool)`,
	)

}

// Test Complex64 << Int
func TestCheckBinaryTypedExprComplex64ShlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << 4`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 4 (shift of type complex64)`,
	)

}

// Test Complex64 << Rune
func TestCheckBinaryTypedExprComplex64ShlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << '@'`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 64 (shift of type complex64)`,
	)

}

// Test Complex64 << Float
func TestCheckBinaryTypedExprComplex64ShlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << 2.0`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 2 (shift of type complex64)`,
	)

}

// Test Complex64 << Complex
func TestCheckBinaryTypedExprComplex64ShlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 0 (shift of type complex64)`,
	)

}

// Test Complex64 << Bool
func TestCheckBinaryTypedExprComplex64ShlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << true`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Complex64 << String
func TestCheckBinaryTypedExprComplex64ShlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: (4.29497e+09+4.29497e+09i) << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Complex64 << Nil
func TestCheckBinaryTypedExprComplex64ShlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Complex64 << Int8
func TestCheckBinaryTypedExprComplex64ShlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << int8(0x7f)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 127 (shift count type int8, must be unsigned integer)`,
	)

}

// Test Complex64 << Int16
func TestCheckBinaryTypedExprComplex64ShlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << int16(0x7fff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 32767 (shift count type int16, must be unsigned integer)`,
	)

}

// Test Complex64 << Int32
func TestCheckBinaryTypedExprComplex64ShlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << int32(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 2147483647 (shift count type int32, must be unsigned integer)`,
	)

}

// Test Complex64 << Int64
func TestCheckBinaryTypedExprComplex64ShlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << int64(0x7fffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 9223372036854775807 (shift count type int64, must be unsigned integer)`,
	)

}

// Test Complex64 << Uint8
func TestCheckBinaryTypedExprComplex64ShlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << uint8(0xff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 255 (shift of type complex64)`,
	)

}

// Test Complex64 << Uint16
func TestCheckBinaryTypedExprComplex64ShlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << uint16(0xffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 65535 (shift of type complex64)`,
	)

}

// Test Complex64 << Uint32
func TestCheckBinaryTypedExprComplex64ShlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << uint32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 4294967295 (shift of type complex64)`,
	)

}

// Test Complex64 << Uint64
func TestCheckBinaryTypedExprComplex64ShlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << uint64(0xffffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 18446744073709551615 (shift of type complex64)`,
	)

}

// Test Complex64 << Float32
func TestCheckBinaryTypedExprComplex64ShlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << float32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 4.29497e+09 (shift count type float32, must be unsigned integer)`,
	)

}

// Test Complex64 << Float64
func TestCheckBinaryTypedExprComplex64ShlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << float64(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 4.29497e+09 (shift count type float64, must be unsigned integer)`,
	)

}

// Test Complex64 << Complex64
func TestCheckBinaryTypedExprComplex64ShlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << (4.29497e+09+4.29497e+09i) (shift count type complex64, must be unsigned integer)`,
	)

}

// Test Complex64 << Complex128
func TestCheckBinaryTypedExprComplex64ShlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << (4.29497e+09+4.29497e+09i) (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Complex64 << Rune32
func TestCheckBinaryTypedExprComplex64ShlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << rune(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << rune(2147483647) (shift count type rune, must be unsigned integer)`,
	)

}

// Test Complex64 << StringT
func TestCheckBinaryTypedExprComplex64ShlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << string("abc")`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Complex64 << BoolT
func TestCheckBinaryTypedExprComplex64ShlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(0xffffffff + 0xffffffff * 1i) << bool(true)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Complex128 + Int
func TestCheckBinaryTypedExprComplex128AddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff + 0xffffffff * 1i) + 4`, env, complex128(0xffffffff + 0xffffffff * 1i) + 4, reflect.TypeOf(complex128(0xffffffff + 0xffffffff * 1i) + 4))
}

// Test Complex128 + Rune
func TestCheckBinaryTypedExprComplex128AddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff + 0xffffffff * 1i) + '@'`, env, complex128(0xffffffff + 0xffffffff * 1i) + '@', reflect.TypeOf(complex128(0xffffffff + 0xffffffff * 1i) + '@'))
}

// Test Complex128 + Float
func TestCheckBinaryTypedExprComplex128AddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff + 0xffffffff * 1i) + 2.0`, env, complex128(0xffffffff + 0xffffffff * 1i) + 2.0, reflect.TypeOf(complex128(0xffffffff + 0xffffffff * 1i) + 2.0))
}

// Test Complex128 + Complex
func TestCheckBinaryTypedExprComplex128AddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff + 0xffffffff * 1i) + 8.0i`, env, complex128(0xffffffff + 0xffffffff * 1i) + 8.0i, reflect.TypeOf(complex128(0xffffffff + 0xffffffff * 1i) + 8.0i))
}

// Test Complex128 + Bool
func TestCheckBinaryTypedExprComplex128AddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: (4.29497e+09+4.29497e+09i) + true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 + String
func TestCheckBinaryTypedExprComplex128AddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: (4.29497e+09+4.29497e+09i) + "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 + Nil
func TestCheckBinaryTypedExprComplex128AddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + nil`, env,
		`cannot convert nil to type complex128`,
	)

}

// Test Complex128 + Int8
func TestCheckBinaryTypedExprComplex128AddInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + int8(0x7f)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 127 (mismatched types complex128 and int8)`,
	)

}

// Test Complex128 + Int16
func TestCheckBinaryTypedExprComplex128AddInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + int16(0x7fff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 32767 (mismatched types complex128 and int16)`,
	)

}

// Test Complex128 + Int32
func TestCheckBinaryTypedExprComplex128AddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + int32(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 2147483647 (mismatched types complex128 and int32)`,
	)

}

// Test Complex128 + Int64
func TestCheckBinaryTypedExprComplex128AddInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + int64(0x7fffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 9223372036854775807 (mismatched types complex128 and int64)`,
	)

}

// Test Complex128 + Uint8
func TestCheckBinaryTypedExprComplex128AddUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + uint8(0xff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 255 (mismatched types complex128 and uint8)`,
	)

}

// Test Complex128 + Uint16
func TestCheckBinaryTypedExprComplex128AddUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + uint16(0xffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 65535 (mismatched types complex128 and uint16)`,
	)

}

// Test Complex128 + Uint32
func TestCheckBinaryTypedExprComplex128AddUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + uint32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 4294967295 (mismatched types complex128 and uint32)`,
	)

}

// Test Complex128 + Uint64
func TestCheckBinaryTypedExprComplex128AddUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + uint64(0xffffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 18446744073709551615 (mismatched types complex128 and uint64)`,
	)

}

// Test Complex128 + Float32
func TestCheckBinaryTypedExprComplex128AddFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + float32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 4.29497e+09 (mismatched types complex128 and float32)`,
	)

}

// Test Complex128 + Float64
func TestCheckBinaryTypedExprComplex128AddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + float64(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + 4.29497e+09 (mismatched types complex128 and float64)`,
	)

}

// Test Complex128 + Complex64
func TestCheckBinaryTypedExprComplex128AddComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + (4.29497e+09+4.29497e+09i) (mismatched types complex128 and complex64)`,
	)

}

// Test Complex128 + Complex128
func TestCheckBinaryTypedExprComplex128AddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff + 0xffffffff * 1i) + complex128(0xffffffff + 0xffffffff * 1i)`, env, complex128(0xffffffff + 0xffffffff * 1i) + complex128(0xffffffff + 0xffffffff * 1i), reflect.TypeOf(complex128(0xffffffff + 0xffffffff * 1i) + complex128(0xffffffff + 0xffffffff * 1i)))
}

// Test Complex128 + Rune32
func TestCheckBinaryTypedExprComplex128AddRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + rune(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + rune(2147483647) (mismatched types complex128 and rune)`,
	)

}

// Test Complex128 + StringT
func TestCheckBinaryTypedExprComplex128AddStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + string("abc")`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 + BoolT
func TestCheckBinaryTypedExprComplex128AddBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) + bool(true)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) + true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 - Int
func TestCheckBinaryTypedExprComplex128SubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff + 0xffffffff * 1i) - 4`, env, complex128(0xffffffff + 0xffffffff * 1i) - 4, reflect.TypeOf(complex128(0xffffffff + 0xffffffff * 1i) - 4))
}

// Test Complex128 - Rune
func TestCheckBinaryTypedExprComplex128SubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff + 0xffffffff * 1i) - '@'`, env, complex128(0xffffffff + 0xffffffff * 1i) - '@', reflect.TypeOf(complex128(0xffffffff + 0xffffffff * 1i) - '@'))
}

// Test Complex128 - Float
func TestCheckBinaryTypedExprComplex128SubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff + 0xffffffff * 1i) - 2.0`, env, complex128(0xffffffff + 0xffffffff * 1i) - 2.0, reflect.TypeOf(complex128(0xffffffff + 0xffffffff * 1i) - 2.0))
}

// Test Complex128 - Complex
func TestCheckBinaryTypedExprComplex128SubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff + 0xffffffff * 1i) - 8.0i`, env, complex128(0xffffffff + 0xffffffff * 1i) - 8.0i, reflect.TypeOf(complex128(0xffffffff + 0xffffffff * 1i) - 8.0i))
}

// Test Complex128 - Bool
func TestCheckBinaryTypedExprComplex128SubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: (4.29497e+09+4.29497e+09i) - true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 - String
func TestCheckBinaryTypedExprComplex128SubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: (4.29497e+09+4.29497e+09i) - "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 - Nil
func TestCheckBinaryTypedExprComplex128SubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - nil`, env,
		`cannot convert nil to type complex128`,
	)

}

// Test Complex128 - Int8
func TestCheckBinaryTypedExprComplex128SubInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - int8(0x7f)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 127 (mismatched types complex128 and int8)`,
	)

}

// Test Complex128 - Int16
func TestCheckBinaryTypedExprComplex128SubInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - int16(0x7fff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 32767 (mismatched types complex128 and int16)`,
	)

}

// Test Complex128 - Int32
func TestCheckBinaryTypedExprComplex128SubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - int32(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 2147483647 (mismatched types complex128 and int32)`,
	)

}

// Test Complex128 - Int64
func TestCheckBinaryTypedExprComplex128SubInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - int64(0x7fffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 9223372036854775807 (mismatched types complex128 and int64)`,
	)

}

// Test Complex128 - Uint8
func TestCheckBinaryTypedExprComplex128SubUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - uint8(0xff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 255 (mismatched types complex128 and uint8)`,
	)

}

// Test Complex128 - Uint16
func TestCheckBinaryTypedExprComplex128SubUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - uint16(0xffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 65535 (mismatched types complex128 and uint16)`,
	)

}

// Test Complex128 - Uint32
func TestCheckBinaryTypedExprComplex128SubUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - uint32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 4294967295 (mismatched types complex128 and uint32)`,
	)

}

// Test Complex128 - Uint64
func TestCheckBinaryTypedExprComplex128SubUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - uint64(0xffffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 18446744073709551615 (mismatched types complex128 and uint64)`,
	)

}

// Test Complex128 - Float32
func TestCheckBinaryTypedExprComplex128SubFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - float32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 4.29497e+09 (mismatched types complex128 and float32)`,
	)

}

// Test Complex128 - Float64
func TestCheckBinaryTypedExprComplex128SubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - float64(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - 4.29497e+09 (mismatched types complex128 and float64)`,
	)

}

// Test Complex128 - Complex64
func TestCheckBinaryTypedExprComplex128SubComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - (4.29497e+09+4.29497e+09i) (mismatched types complex128 and complex64)`,
	)

}

// Test Complex128 - Complex128
func TestCheckBinaryTypedExprComplex128SubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff + 0xffffffff * 1i) - complex128(0xffffffff + 0xffffffff * 1i)`, env, complex128(0xffffffff + 0xffffffff * 1i) - complex128(0xffffffff + 0xffffffff * 1i), reflect.TypeOf(complex128(0xffffffff + 0xffffffff * 1i) - complex128(0xffffffff + 0xffffffff * 1i)))
}

// Test Complex128 - Rune32
func TestCheckBinaryTypedExprComplex128SubRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - rune(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - rune(2147483647) (mismatched types complex128 and rune)`,
	)

}

// Test Complex128 - StringT
func TestCheckBinaryTypedExprComplex128SubStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - string("abc")`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 - BoolT
func TestCheckBinaryTypedExprComplex128SubBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) - bool(true)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) - true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 & Int
func TestCheckBinaryTypedExprComplex128AndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & 4`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 4 (operator & not defined on complex128)`,
	)

}

// Test Complex128 & Rune
func TestCheckBinaryTypedExprComplex128AndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & '@'`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 64 (operator & not defined on complex128)`,
	)

}

// Test Complex128 & Float
func TestCheckBinaryTypedExprComplex128AndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & 2.0`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 2 (operator & not defined on complex128)`,
	)

}

// Test Complex128 & Complex
func TestCheckBinaryTypedExprComplex128AndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & 8.0i`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 8i (operator & not defined on complex128)`,
	)

}

// Test Complex128 & Bool
func TestCheckBinaryTypedExprComplex128AndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: (4.29497e+09+4.29497e+09i) & true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 & String
func TestCheckBinaryTypedExprComplex128AndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: (4.29497e+09+4.29497e+09i) & "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 & Nil
func TestCheckBinaryTypedExprComplex128AndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & nil`, env,
		`cannot convert nil to type complex128`,
	)

}

// Test Complex128 & Int8
func TestCheckBinaryTypedExprComplex128AndInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & int8(0x7f)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 127 (mismatched types complex128 and int8)`,
	)

}

// Test Complex128 & Int16
func TestCheckBinaryTypedExprComplex128AndInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & int16(0x7fff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 32767 (mismatched types complex128 and int16)`,
	)

}

// Test Complex128 & Int32
func TestCheckBinaryTypedExprComplex128AndInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & int32(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 2147483647 (mismatched types complex128 and int32)`,
	)

}

// Test Complex128 & Int64
func TestCheckBinaryTypedExprComplex128AndInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & int64(0x7fffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 9223372036854775807 (mismatched types complex128 and int64)`,
	)

}

// Test Complex128 & Uint8
func TestCheckBinaryTypedExprComplex128AndUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & uint8(0xff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 255 (mismatched types complex128 and uint8)`,
	)

}

// Test Complex128 & Uint16
func TestCheckBinaryTypedExprComplex128AndUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & uint16(0xffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 65535 (mismatched types complex128 and uint16)`,
	)

}

// Test Complex128 & Uint32
func TestCheckBinaryTypedExprComplex128AndUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & uint32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 4294967295 (mismatched types complex128 and uint32)`,
	)

}

// Test Complex128 & Uint64
func TestCheckBinaryTypedExprComplex128AndUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & uint64(0xffffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 18446744073709551615 (mismatched types complex128 and uint64)`,
	)

}

// Test Complex128 & Float32
func TestCheckBinaryTypedExprComplex128AndFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & float32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 4.29497e+09 (mismatched types complex128 and float32)`,
	)

}

// Test Complex128 & Float64
func TestCheckBinaryTypedExprComplex128AndFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & float64(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & 4.29497e+09 (mismatched types complex128 and float64)`,
	)

}

// Test Complex128 & Complex64
func TestCheckBinaryTypedExprComplex128AndComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & (4.29497e+09+4.29497e+09i) (mismatched types complex128 and complex64)`,
	)

}

// Test Complex128 & Complex128
func TestCheckBinaryTypedExprComplex128AndComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & (4.29497e+09+4.29497e+09i) (operator & not defined on complex128)`,
	)

}

// Test Complex128 & Rune32
func TestCheckBinaryTypedExprComplex128AndRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & rune(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & rune(2147483647) (mismatched types complex128 and rune)`,
	)

}

// Test Complex128 & StringT
func TestCheckBinaryTypedExprComplex128AndStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & string("abc")`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 & BoolT
func TestCheckBinaryTypedExprComplex128AndBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) & bool(true)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) & true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 % Int
func TestCheckBinaryTypedExprComplex128RemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % 4`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 4 (operator % not defined on complex128)`,
	)

}

// Test Complex128 % Rune
func TestCheckBinaryTypedExprComplex128RemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % '@'`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 64 (operator % not defined on complex128)`,
	)

}

// Test Complex128 % Float
func TestCheckBinaryTypedExprComplex128RemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % 2.0`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 2 (operator % not defined on complex128)`,
	)

}

// Test Complex128 % Complex
func TestCheckBinaryTypedExprComplex128RemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % 8.0i`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 8i (operator % not defined on complex128)`,
	)

}

// Test Complex128 % Bool
func TestCheckBinaryTypedExprComplex128RemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: (4.29497e+09+4.29497e+09i) % true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 % String
func TestCheckBinaryTypedExprComplex128RemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: (4.29497e+09+4.29497e+09i) % "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 % Nil
func TestCheckBinaryTypedExprComplex128RemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % nil`, env,
		`cannot convert nil to type complex128`,
	)

}

// Test Complex128 % Int8
func TestCheckBinaryTypedExprComplex128RemInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % int8(0x7f)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 127 (mismatched types complex128 and int8)`,
	)

}

// Test Complex128 % Int16
func TestCheckBinaryTypedExprComplex128RemInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % int16(0x7fff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 32767 (mismatched types complex128 and int16)`,
	)

}

// Test Complex128 % Int32
func TestCheckBinaryTypedExprComplex128RemInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % int32(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 2147483647 (mismatched types complex128 and int32)`,
	)

}

// Test Complex128 % Int64
func TestCheckBinaryTypedExprComplex128RemInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % int64(0x7fffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 9223372036854775807 (mismatched types complex128 and int64)`,
	)

}

// Test Complex128 % Uint8
func TestCheckBinaryTypedExprComplex128RemUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % uint8(0xff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 255 (mismatched types complex128 and uint8)`,
	)

}

// Test Complex128 % Uint16
func TestCheckBinaryTypedExprComplex128RemUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % uint16(0xffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 65535 (mismatched types complex128 and uint16)`,
	)

}

// Test Complex128 % Uint32
func TestCheckBinaryTypedExprComplex128RemUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % uint32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 4294967295 (mismatched types complex128 and uint32)`,
	)

}

// Test Complex128 % Uint64
func TestCheckBinaryTypedExprComplex128RemUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % uint64(0xffffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 18446744073709551615 (mismatched types complex128 and uint64)`,
	)

}

// Test Complex128 % Float32
func TestCheckBinaryTypedExprComplex128RemFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % float32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 4.29497e+09 (mismatched types complex128 and float32)`,
	)

}

// Test Complex128 % Float64
func TestCheckBinaryTypedExprComplex128RemFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % float64(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % 4.29497e+09 (mismatched types complex128 and float64)`,
	)

}

// Test Complex128 % Complex64
func TestCheckBinaryTypedExprComplex128RemComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % (4.29497e+09+4.29497e+09i) (mismatched types complex128 and complex64)`,
	)

}

// Test Complex128 % Complex128
func TestCheckBinaryTypedExprComplex128RemComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % (4.29497e+09+4.29497e+09i) (operator % not defined on complex128)`,
	)

}

// Test Complex128 % Rune32
func TestCheckBinaryTypedExprComplex128RemRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % rune(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % rune(2147483647) (mismatched types complex128 and rune)`,
	)

}

// Test Complex128 % StringT
func TestCheckBinaryTypedExprComplex128RemStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % string("abc")`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 % BoolT
func TestCheckBinaryTypedExprComplex128RemBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) % bool(true)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) % true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 == Int
func TestCheckBinaryTypedExprComplex128EqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff + 0xffffffff * 1i) == 4`, env, complex128(0xffffffff + 0xffffffff * 1i) == 4, reflect.TypeOf(complex128(0xffffffff + 0xffffffff * 1i) == 4))
}

// Test Complex128 == Rune
func TestCheckBinaryTypedExprComplex128EqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff + 0xffffffff * 1i) == '@'`, env, complex128(0xffffffff + 0xffffffff * 1i) == '@', reflect.TypeOf(complex128(0xffffffff + 0xffffffff * 1i) == '@'))
}

// Test Complex128 == Float
func TestCheckBinaryTypedExprComplex128EqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff + 0xffffffff * 1i) == 2.0`, env, complex128(0xffffffff + 0xffffffff * 1i) == 2.0, reflect.TypeOf(complex128(0xffffffff + 0xffffffff * 1i) == 2.0))
}

// Test Complex128 == Complex
func TestCheckBinaryTypedExprComplex128EqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff + 0xffffffff * 1i) == 8.0i`, env, complex128(0xffffffff + 0xffffffff * 1i) == 8.0i, reflect.TypeOf(complex128(0xffffffff + 0xffffffff * 1i) == 8.0i))
}

// Test Complex128 == Bool
func TestCheckBinaryTypedExprComplex128EqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: (4.29497e+09+4.29497e+09i) == true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 == String
func TestCheckBinaryTypedExprComplex128EqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: (4.29497e+09+4.29497e+09i) == "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 == Nil
func TestCheckBinaryTypedExprComplex128EqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == nil`, env,
		`cannot convert nil to type complex128`,
	)

}

// Test Complex128 == Int8
func TestCheckBinaryTypedExprComplex128EqlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == int8(0x7f)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 127 (mismatched types complex128 and int8)`,
	)

}

// Test Complex128 == Int16
func TestCheckBinaryTypedExprComplex128EqlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == int16(0x7fff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 32767 (mismatched types complex128 and int16)`,
	)

}

// Test Complex128 == Int32
func TestCheckBinaryTypedExprComplex128EqlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == int32(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 2147483647 (mismatched types complex128 and int32)`,
	)

}

// Test Complex128 == Int64
func TestCheckBinaryTypedExprComplex128EqlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == int64(0x7fffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 9223372036854775807 (mismatched types complex128 and int64)`,
	)

}

// Test Complex128 == Uint8
func TestCheckBinaryTypedExprComplex128EqlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == uint8(0xff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 255 (mismatched types complex128 and uint8)`,
	)

}

// Test Complex128 == Uint16
func TestCheckBinaryTypedExprComplex128EqlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == uint16(0xffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 65535 (mismatched types complex128 and uint16)`,
	)

}

// Test Complex128 == Uint32
func TestCheckBinaryTypedExprComplex128EqlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == uint32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 4294967295 (mismatched types complex128 and uint32)`,
	)

}

// Test Complex128 == Uint64
func TestCheckBinaryTypedExprComplex128EqlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == uint64(0xffffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 18446744073709551615 (mismatched types complex128 and uint64)`,
	)

}

// Test Complex128 == Float32
func TestCheckBinaryTypedExprComplex128EqlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == float32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 4.29497e+09 (mismatched types complex128 and float32)`,
	)

}

// Test Complex128 == Float64
func TestCheckBinaryTypedExprComplex128EqlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == float64(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == 4.29497e+09 (mismatched types complex128 and float64)`,
	)

}

// Test Complex128 == Complex64
func TestCheckBinaryTypedExprComplex128EqlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == (4.29497e+09+4.29497e+09i) (mismatched types complex128 and complex64)`,
	)

}

// Test Complex128 == Complex128
func TestCheckBinaryTypedExprComplex128EqlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff + 0xffffffff * 1i) == complex128(0xffffffff + 0xffffffff * 1i)`, env, complex128(0xffffffff + 0xffffffff * 1i) == complex128(0xffffffff + 0xffffffff * 1i), reflect.TypeOf(complex128(0xffffffff + 0xffffffff * 1i) == complex128(0xffffffff + 0xffffffff * 1i)))
}

// Test Complex128 == Rune32
func TestCheckBinaryTypedExprComplex128EqlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == rune(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == rune(2147483647) (mismatched types complex128 and rune)`,
	)

}

// Test Complex128 == StringT
func TestCheckBinaryTypedExprComplex128EqlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == string("abc")`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 == BoolT
func TestCheckBinaryTypedExprComplex128EqlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) == bool(true)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) == true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 > Int
func TestCheckBinaryTypedExprComplex128GtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > 4`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 4 (operator > not defined on complex128)`,
	)

}

// Test Complex128 > Rune
func TestCheckBinaryTypedExprComplex128GtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > '@'`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 64 (operator > not defined on complex128)`,
	)

}

// Test Complex128 > Float
func TestCheckBinaryTypedExprComplex128GtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > 2.0`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 2 (operator > not defined on complex128)`,
	)

}

// Test Complex128 > Complex
func TestCheckBinaryTypedExprComplex128GtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > 8.0i`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 8i (operator > not defined on complex128)`,
	)

}

// Test Complex128 > Bool
func TestCheckBinaryTypedExprComplex128GtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: (4.29497e+09+4.29497e+09i) > true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 > String
func TestCheckBinaryTypedExprComplex128GtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: (4.29497e+09+4.29497e+09i) > "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 > Nil
func TestCheckBinaryTypedExprComplex128GtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > nil`, env,
		`cannot convert nil to type complex128`,
	)

}

// Test Complex128 > Int8
func TestCheckBinaryTypedExprComplex128GtrInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > int8(0x7f)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 127 (mismatched types complex128 and int8)`,
	)

}

// Test Complex128 > Int16
func TestCheckBinaryTypedExprComplex128GtrInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > int16(0x7fff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 32767 (mismatched types complex128 and int16)`,
	)

}

// Test Complex128 > Int32
func TestCheckBinaryTypedExprComplex128GtrInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > int32(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 2147483647 (mismatched types complex128 and int32)`,
	)

}

// Test Complex128 > Int64
func TestCheckBinaryTypedExprComplex128GtrInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > int64(0x7fffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 9223372036854775807 (mismatched types complex128 and int64)`,
	)

}

// Test Complex128 > Uint8
func TestCheckBinaryTypedExprComplex128GtrUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > uint8(0xff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 255 (mismatched types complex128 and uint8)`,
	)

}

// Test Complex128 > Uint16
func TestCheckBinaryTypedExprComplex128GtrUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > uint16(0xffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 65535 (mismatched types complex128 and uint16)`,
	)

}

// Test Complex128 > Uint32
func TestCheckBinaryTypedExprComplex128GtrUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > uint32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 4294967295 (mismatched types complex128 and uint32)`,
	)

}

// Test Complex128 > Uint64
func TestCheckBinaryTypedExprComplex128GtrUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > uint64(0xffffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 18446744073709551615 (mismatched types complex128 and uint64)`,
	)

}

// Test Complex128 > Float32
func TestCheckBinaryTypedExprComplex128GtrFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > float32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 4.29497e+09 (mismatched types complex128 and float32)`,
	)

}

// Test Complex128 > Float64
func TestCheckBinaryTypedExprComplex128GtrFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > float64(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > 4.29497e+09 (mismatched types complex128 and float64)`,
	)

}

// Test Complex128 > Complex64
func TestCheckBinaryTypedExprComplex128GtrComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > (4.29497e+09+4.29497e+09i) (mismatched types complex128 and complex64)`,
	)

}

// Test Complex128 > Complex128
func TestCheckBinaryTypedExprComplex128GtrComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > (4.29497e+09+4.29497e+09i) (operator > not defined on complex128)`,
	)

}

// Test Complex128 > Rune32
func TestCheckBinaryTypedExprComplex128GtrRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > rune(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > rune(2147483647) (mismatched types complex128 and rune)`,
	)

}

// Test Complex128 > StringT
func TestCheckBinaryTypedExprComplex128GtrStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > string("abc")`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 > BoolT
func TestCheckBinaryTypedExprComplex128GtrBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) > bool(true)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) > true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 << Int
func TestCheckBinaryTypedExprComplex128ShlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << 4`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 4 (shift of type complex128)`,
	)

}

// Test Complex128 << Rune
func TestCheckBinaryTypedExprComplex128ShlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << '@'`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 64 (shift of type complex128)`,
	)

}

// Test Complex128 << Float
func TestCheckBinaryTypedExprComplex128ShlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << 2.0`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 2 (shift of type complex128)`,
	)

}

// Test Complex128 << Complex
func TestCheckBinaryTypedExprComplex128ShlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 0 (shift of type complex128)`,
	)

}

// Test Complex128 << Bool
func TestCheckBinaryTypedExprComplex128ShlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << true`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Complex128 << String
func TestCheckBinaryTypedExprComplex128ShlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: (4.29497e+09+4.29497e+09i) << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Complex128 << Nil
func TestCheckBinaryTypedExprComplex128ShlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Complex128 << Int8
func TestCheckBinaryTypedExprComplex128ShlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << int8(0x7f)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 127 (shift count type int8, must be unsigned integer)`,
	)

}

// Test Complex128 << Int16
func TestCheckBinaryTypedExprComplex128ShlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << int16(0x7fff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 32767 (shift count type int16, must be unsigned integer)`,
	)

}

// Test Complex128 << Int32
func TestCheckBinaryTypedExprComplex128ShlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << int32(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 2147483647 (shift count type int32, must be unsigned integer)`,
	)

}

// Test Complex128 << Int64
func TestCheckBinaryTypedExprComplex128ShlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << int64(0x7fffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 9223372036854775807 (shift count type int64, must be unsigned integer)`,
	)

}

// Test Complex128 << Uint8
func TestCheckBinaryTypedExprComplex128ShlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << uint8(0xff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 255 (shift of type complex128)`,
	)

}

// Test Complex128 << Uint16
func TestCheckBinaryTypedExprComplex128ShlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << uint16(0xffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 65535 (shift of type complex128)`,
	)

}

// Test Complex128 << Uint32
func TestCheckBinaryTypedExprComplex128ShlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << uint32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 4294967295 (shift of type complex128)`,
	)

}

// Test Complex128 << Uint64
func TestCheckBinaryTypedExprComplex128ShlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << uint64(0xffffffffffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 18446744073709551615 (shift of type complex128)`,
	)

}

// Test Complex128 << Float32
func TestCheckBinaryTypedExprComplex128ShlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << float32(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 4.29497e+09 (shift count type float32, must be unsigned integer)`,
	)

}

// Test Complex128 << Float64
func TestCheckBinaryTypedExprComplex128ShlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << float64(0xffffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << 4.29497e+09 (shift count type float64, must be unsigned integer)`,
	)

}

// Test Complex128 << Complex64
func TestCheckBinaryTypedExprComplex128ShlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << (4.29497e+09+4.29497e+09i) (shift count type complex64, must be unsigned integer)`,
	)

}

// Test Complex128 << Complex128
func TestCheckBinaryTypedExprComplex128ShlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << (4.29497e+09+4.29497e+09i) (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Complex128 << Rune32
func TestCheckBinaryTypedExprComplex128ShlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << rune(0x7fffffff)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << rune(2147483647) (shift count type rune, must be unsigned integer)`,
	)

}

// Test Complex128 << StringT
func TestCheckBinaryTypedExprComplex128ShlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << string("abc")`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Complex128 << BoolT
func TestCheckBinaryTypedExprComplex128ShlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(0xffffffff + 0xffffffff * 1i) << bool(true)`, env,
		`invalid operation: (4.29497e+09+4.29497e+09i) << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Rune32 + Int
func TestCheckBinaryTypedExprRune32AddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + 4`, env,
		`constant 2147483651 overflows rune`,
	)

}

// Test Rune32 + Rune
func TestCheckBinaryTypedExprRune32AddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + '@'`, env,
		`constant 2147483711 overflows rune`,
	)

}

// Test Rune32 + Float
func TestCheckBinaryTypedExprRune32AddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + 2.0`, env,
		`constant 2147483649 overflows rune`,
	)

}

// Test Rune32 + Complex
func TestCheckBinaryTypedExprRune32AddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Rune32 + Bool
func TestCheckBinaryTypedExprRune32AddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + true`, env,
		`cannot convert true to type rune`,
		`invalid operation: rune(2147483647) + true (mismatched types rune and bool)`,
	)

}

// Test Rune32 + String
func TestCheckBinaryTypedExprRune32AddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: rune(2147483647) + "abc" (mismatched types rune and string)`,
	)

}

// Test Rune32 + Nil
func TestCheckBinaryTypedExprRune32AddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + nil`, env,
		`cannot convert nil to type rune`,
	)

}

// Test Rune32 + Int8
func TestCheckBinaryTypedExprRune32AddInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + int8(0x7f)`, env,
		`invalid operation: rune(2147483647) + 127 (mismatched types rune and int8)`,
	)

}

// Test Rune32 + Int16
func TestCheckBinaryTypedExprRune32AddInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + int16(0x7fff)`, env,
		`invalid operation: rune(2147483647) + 32767 (mismatched types rune and int16)`,
	)

}

// Test Rune32 + Int32
func TestCheckBinaryTypedExprRune32AddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + int32(0x7fffffff)`, env,
		`constant 4294967294 overflows rune`,
	)

}

// Test Rune32 + Int64
func TestCheckBinaryTypedExprRune32AddInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + int64(0x7fffffffffffffff)`, env,
		`invalid operation: rune(2147483647) + 9223372036854775807 (mismatched types rune and int64)`,
	)

}

// Test Rune32 + Uint8
func TestCheckBinaryTypedExprRune32AddUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + uint8(0xff)`, env,
		`invalid operation: rune(2147483647) + 255 (mismatched types rune and uint8)`,
	)

}

// Test Rune32 + Uint16
func TestCheckBinaryTypedExprRune32AddUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + uint16(0xffff)`, env,
		`invalid operation: rune(2147483647) + 65535 (mismatched types rune and uint16)`,
	)

}

// Test Rune32 + Uint32
func TestCheckBinaryTypedExprRune32AddUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + uint32(0xffffffff)`, env,
		`invalid operation: rune(2147483647) + 4294967295 (mismatched types rune and uint32)`,
	)

}

// Test Rune32 + Uint64
func TestCheckBinaryTypedExprRune32AddUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + uint64(0xffffffffffffffff)`, env,
		`invalid operation: rune(2147483647) + 18446744073709551615 (mismatched types rune and uint64)`,
	)

}

// Test Rune32 + Float32
func TestCheckBinaryTypedExprRune32AddFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + float32(0xffffffff)`, env,
		`invalid operation: rune(2147483647) + 4.29497e+09 (mismatched types rune and float32)`,
	)

}

// Test Rune32 + Float64
func TestCheckBinaryTypedExprRune32AddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + float64(0xffffffff)`, env,
		`invalid operation: rune(2147483647) + 4.29497e+09 (mismatched types rune and float64)`,
	)

}

// Test Rune32 + Complex64
func TestCheckBinaryTypedExprRune32AddComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: rune(2147483647) + (4.29497e+09+4.29497e+09i) (mismatched types rune and complex64)`,
	)

}

// Test Rune32 + Complex128
func TestCheckBinaryTypedExprRune32AddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: rune(2147483647) + (4.29497e+09+4.29497e+09i) (mismatched types rune and complex128)`,
	)

}

// Test Rune32 + Rune32
func TestCheckBinaryTypedExprRune32AddRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + rune(0x7fffffff)`, env,
		`constant 4294967294 overflows rune`,
	)

}

// Test Rune32 + StringT
func TestCheckBinaryTypedExprRune32AddStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + string("abc")`, env,
		`invalid operation: rune(2147483647) + "abc" (mismatched types rune and string)`,
	)

}

// Test Rune32 + BoolT
func TestCheckBinaryTypedExprRune32AddBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) + bool(true)`, env,
		`invalid operation: rune(2147483647) + true (mismatched types rune and bool)`,
	)

}

// Test Rune32 - Int
func TestCheckBinaryTypedExprRune32SubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) - 4`, env, rune(0x7fffffff) - 4, reflect.TypeOf(rune(0x7fffffff) - 4))
}

// Test Rune32 - Rune
func TestCheckBinaryTypedExprRune32SubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) - '@'`, env, rune(0x7fffffff) - '@', reflect.TypeOf(rune(0x7fffffff) - '@'))
}

// Test Rune32 - Float
func TestCheckBinaryTypedExprRune32SubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) - 2.0`, env, rune(0x7fffffff) - 2.0, reflect.TypeOf(rune(0x7fffffff) - 2.0))
}

// Test Rune32 - Complex
func TestCheckBinaryTypedExprRune32SubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Rune32 - Bool
func TestCheckBinaryTypedExprRune32SubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - true`, env,
		`cannot convert true to type rune`,
		`invalid operation: rune(2147483647) - true (mismatched types rune and bool)`,
	)

}

// Test Rune32 - String
func TestCheckBinaryTypedExprRune32SubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: rune(2147483647) - "abc" (mismatched types rune and string)`,
	)

}

// Test Rune32 - Nil
func TestCheckBinaryTypedExprRune32SubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - nil`, env,
		`cannot convert nil to type rune`,
	)

}

// Test Rune32 - Int8
func TestCheckBinaryTypedExprRune32SubInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - int8(0x7f)`, env,
		`invalid operation: rune(2147483647) - 127 (mismatched types rune and int8)`,
	)

}

// Test Rune32 - Int16
func TestCheckBinaryTypedExprRune32SubInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - int16(0x7fff)`, env,
		`invalid operation: rune(2147483647) - 32767 (mismatched types rune and int16)`,
	)

}

// Test Rune32 - Int32
func TestCheckBinaryTypedExprRune32SubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) - int32(0x7fffffff)`, env, rune(0x7fffffff) - int32(0x7fffffff), reflect.TypeOf(rune(0x7fffffff) - int32(0x7fffffff)))
}

// Test Rune32 - Int64
func TestCheckBinaryTypedExprRune32SubInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - int64(0x7fffffffffffffff)`, env,
		`invalid operation: rune(2147483647) - 9223372036854775807 (mismatched types rune and int64)`,
	)

}

// Test Rune32 - Uint8
func TestCheckBinaryTypedExprRune32SubUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - uint8(0xff)`, env,
		`invalid operation: rune(2147483647) - 255 (mismatched types rune and uint8)`,
	)

}

// Test Rune32 - Uint16
func TestCheckBinaryTypedExprRune32SubUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - uint16(0xffff)`, env,
		`invalid operation: rune(2147483647) - 65535 (mismatched types rune and uint16)`,
	)

}

// Test Rune32 - Uint32
func TestCheckBinaryTypedExprRune32SubUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - uint32(0xffffffff)`, env,
		`invalid operation: rune(2147483647) - 4294967295 (mismatched types rune and uint32)`,
	)

}

// Test Rune32 - Uint64
func TestCheckBinaryTypedExprRune32SubUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - uint64(0xffffffffffffffff)`, env,
		`invalid operation: rune(2147483647) - 18446744073709551615 (mismatched types rune and uint64)`,
	)

}

// Test Rune32 - Float32
func TestCheckBinaryTypedExprRune32SubFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - float32(0xffffffff)`, env,
		`invalid operation: rune(2147483647) - 4.29497e+09 (mismatched types rune and float32)`,
	)

}

// Test Rune32 - Float64
func TestCheckBinaryTypedExprRune32SubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - float64(0xffffffff)`, env,
		`invalid operation: rune(2147483647) - 4.29497e+09 (mismatched types rune and float64)`,
	)

}

// Test Rune32 - Complex64
func TestCheckBinaryTypedExprRune32SubComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: rune(2147483647) - (4.29497e+09+4.29497e+09i) (mismatched types rune and complex64)`,
	)

}

// Test Rune32 - Complex128
func TestCheckBinaryTypedExprRune32SubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: rune(2147483647) - (4.29497e+09+4.29497e+09i) (mismatched types rune and complex128)`,
	)

}

// Test Rune32 - Rune32
func TestCheckBinaryTypedExprRune32SubRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) - rune(0x7fffffff)`, env, rune(0x7fffffff) - rune(0x7fffffff), reflect.TypeOf(rune(0x7fffffff) - rune(0x7fffffff)))
}

// Test Rune32 - StringT
func TestCheckBinaryTypedExprRune32SubStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - string("abc")`, env,
		`invalid operation: rune(2147483647) - "abc" (mismatched types rune and string)`,
	)

}

// Test Rune32 - BoolT
func TestCheckBinaryTypedExprRune32SubBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) - bool(true)`, env,
		`invalid operation: rune(2147483647) - true (mismatched types rune and bool)`,
	)

}

// Test Rune32 & Int
func TestCheckBinaryTypedExprRune32AndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) & 4`, env, rune(0x7fffffff) & 4, reflect.TypeOf(rune(0x7fffffff) & 4))
}

// Test Rune32 & Rune
func TestCheckBinaryTypedExprRune32AndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) & '@'`, env, rune(0x7fffffff) & '@', reflect.TypeOf(rune(0x7fffffff) & '@'))
}

// Test Rune32 & Float
func TestCheckBinaryTypedExprRune32AndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) & 2.0`, env, rune(0x7fffffff) & 2.0, reflect.TypeOf(rune(0x7fffffff) & 2.0))
}

// Test Rune32 & Complex
func TestCheckBinaryTypedExprRune32AndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Rune32 & Bool
func TestCheckBinaryTypedExprRune32AndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & true`, env,
		`cannot convert true to type rune`,
		`invalid operation: rune(2147483647) & true (mismatched types rune and bool)`,
	)

}

// Test Rune32 & String
func TestCheckBinaryTypedExprRune32AndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: rune(2147483647) & "abc" (mismatched types rune and string)`,
	)

}

// Test Rune32 & Nil
func TestCheckBinaryTypedExprRune32AndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & nil`, env,
		`cannot convert nil to type rune`,
	)

}

// Test Rune32 & Int8
func TestCheckBinaryTypedExprRune32AndInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & int8(0x7f)`, env,
		`invalid operation: rune(2147483647) & 127 (mismatched types rune and int8)`,
	)

}

// Test Rune32 & Int16
func TestCheckBinaryTypedExprRune32AndInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & int16(0x7fff)`, env,
		`invalid operation: rune(2147483647) & 32767 (mismatched types rune and int16)`,
	)

}

// Test Rune32 & Int32
func TestCheckBinaryTypedExprRune32AndInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) & int32(0x7fffffff)`, env, rune(0x7fffffff) & int32(0x7fffffff), reflect.TypeOf(rune(0x7fffffff) & int32(0x7fffffff)))
}

// Test Rune32 & Int64
func TestCheckBinaryTypedExprRune32AndInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & int64(0x7fffffffffffffff)`, env,
		`invalid operation: rune(2147483647) & 9223372036854775807 (mismatched types rune and int64)`,
	)

}

// Test Rune32 & Uint8
func TestCheckBinaryTypedExprRune32AndUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & uint8(0xff)`, env,
		`invalid operation: rune(2147483647) & 255 (mismatched types rune and uint8)`,
	)

}

// Test Rune32 & Uint16
func TestCheckBinaryTypedExprRune32AndUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & uint16(0xffff)`, env,
		`invalid operation: rune(2147483647) & 65535 (mismatched types rune and uint16)`,
	)

}

// Test Rune32 & Uint32
func TestCheckBinaryTypedExprRune32AndUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & uint32(0xffffffff)`, env,
		`invalid operation: rune(2147483647) & 4294967295 (mismatched types rune and uint32)`,
	)

}

// Test Rune32 & Uint64
func TestCheckBinaryTypedExprRune32AndUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & uint64(0xffffffffffffffff)`, env,
		`invalid operation: rune(2147483647) & 18446744073709551615 (mismatched types rune and uint64)`,
	)

}

// Test Rune32 & Float32
func TestCheckBinaryTypedExprRune32AndFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & float32(0xffffffff)`, env,
		`invalid operation: rune(2147483647) & 4.29497e+09 (mismatched types rune and float32)`,
	)

}

// Test Rune32 & Float64
func TestCheckBinaryTypedExprRune32AndFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & float64(0xffffffff)`, env,
		`invalid operation: rune(2147483647) & 4.29497e+09 (mismatched types rune and float64)`,
	)

}

// Test Rune32 & Complex64
func TestCheckBinaryTypedExprRune32AndComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: rune(2147483647) & (4.29497e+09+4.29497e+09i) (mismatched types rune and complex64)`,
	)

}

// Test Rune32 & Complex128
func TestCheckBinaryTypedExprRune32AndComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: rune(2147483647) & (4.29497e+09+4.29497e+09i) (mismatched types rune and complex128)`,
	)

}

// Test Rune32 & Rune32
func TestCheckBinaryTypedExprRune32AndRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) & rune(0x7fffffff)`, env, rune(0x7fffffff) & rune(0x7fffffff), reflect.TypeOf(rune(0x7fffffff) & rune(0x7fffffff)))
}

// Test Rune32 & StringT
func TestCheckBinaryTypedExprRune32AndStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & string("abc")`, env,
		`invalid operation: rune(2147483647) & "abc" (mismatched types rune and string)`,
	)

}

// Test Rune32 & BoolT
func TestCheckBinaryTypedExprRune32AndBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) & bool(true)`, env,
		`invalid operation: rune(2147483647) & true (mismatched types rune and bool)`,
	)

}

// Test Rune32 % Int
func TestCheckBinaryTypedExprRune32RemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) % 4`, env, rune(0x7fffffff) % 4, reflect.TypeOf(rune(0x7fffffff) % 4))
}

// Test Rune32 % Rune
func TestCheckBinaryTypedExprRune32RemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) % '@'`, env, rune(0x7fffffff) % '@', reflect.TypeOf(rune(0x7fffffff) % '@'))
}

// Test Rune32 % Float
func TestCheckBinaryTypedExprRune32RemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) % 2.0`, env, rune(0x7fffffff) % 2.0, reflect.TypeOf(rune(0x7fffffff) % 2.0))
}

// Test Rune32 % Complex
func TestCheckBinaryTypedExprRune32RemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % 8.0i`, env,
		`constant 0+8i truncated to real`,
		`division by zero`,
	)

}

// Test Rune32 % Bool
func TestCheckBinaryTypedExprRune32RemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % true`, env,
		`cannot convert true to type rune`,
		`invalid operation: rune(2147483647) % true (mismatched types rune and bool)`,
	)

}

// Test Rune32 % String
func TestCheckBinaryTypedExprRune32RemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: rune(2147483647) % "abc" (mismatched types rune and string)`,
	)

}

// Test Rune32 % Nil
func TestCheckBinaryTypedExprRune32RemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % nil`, env,
		`cannot convert nil to type rune`,
	)

}

// Test Rune32 % Int8
func TestCheckBinaryTypedExprRune32RemInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % int8(0x7f)`, env,
		`invalid operation: rune(2147483647) % 127 (mismatched types rune and int8)`,
	)

}

// Test Rune32 % Int16
func TestCheckBinaryTypedExprRune32RemInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % int16(0x7fff)`, env,
		`invalid operation: rune(2147483647) % 32767 (mismatched types rune and int16)`,
	)

}

// Test Rune32 % Int32
func TestCheckBinaryTypedExprRune32RemInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) % int32(0x7fffffff)`, env, rune(0x7fffffff) % int32(0x7fffffff), reflect.TypeOf(rune(0x7fffffff) % int32(0x7fffffff)))
}

// Test Rune32 % Int64
func TestCheckBinaryTypedExprRune32RemInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % int64(0x7fffffffffffffff)`, env,
		`invalid operation: rune(2147483647) % 9223372036854775807 (mismatched types rune and int64)`,
	)

}

// Test Rune32 % Uint8
func TestCheckBinaryTypedExprRune32RemUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % uint8(0xff)`, env,
		`invalid operation: rune(2147483647) % 255 (mismatched types rune and uint8)`,
	)

}

// Test Rune32 % Uint16
func TestCheckBinaryTypedExprRune32RemUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % uint16(0xffff)`, env,
		`invalid operation: rune(2147483647) % 65535 (mismatched types rune and uint16)`,
	)

}

// Test Rune32 % Uint32
func TestCheckBinaryTypedExprRune32RemUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % uint32(0xffffffff)`, env,
		`invalid operation: rune(2147483647) % 4294967295 (mismatched types rune and uint32)`,
	)

}

// Test Rune32 % Uint64
func TestCheckBinaryTypedExprRune32RemUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % uint64(0xffffffffffffffff)`, env,
		`invalid operation: rune(2147483647) % 18446744073709551615 (mismatched types rune and uint64)`,
	)

}

// Test Rune32 % Float32
func TestCheckBinaryTypedExprRune32RemFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % float32(0xffffffff)`, env,
		`invalid operation: rune(2147483647) % 4.29497e+09 (mismatched types rune and float32)`,
	)

}

// Test Rune32 % Float64
func TestCheckBinaryTypedExprRune32RemFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % float64(0xffffffff)`, env,
		`invalid operation: rune(2147483647) % 4.29497e+09 (mismatched types rune and float64)`,
	)

}

// Test Rune32 % Complex64
func TestCheckBinaryTypedExprRune32RemComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: rune(2147483647) % (4.29497e+09+4.29497e+09i) (mismatched types rune and complex64)`,
	)

}

// Test Rune32 % Complex128
func TestCheckBinaryTypedExprRune32RemComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: rune(2147483647) % (4.29497e+09+4.29497e+09i) (mismatched types rune and complex128)`,
	)

}

// Test Rune32 % Rune32
func TestCheckBinaryTypedExprRune32RemRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) % rune(0x7fffffff)`, env, rune(0x7fffffff) % rune(0x7fffffff), reflect.TypeOf(rune(0x7fffffff) % rune(0x7fffffff)))
}

// Test Rune32 % StringT
func TestCheckBinaryTypedExprRune32RemStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % string("abc")`, env,
		`invalid operation: rune(2147483647) % "abc" (mismatched types rune and string)`,
	)

}

// Test Rune32 % BoolT
func TestCheckBinaryTypedExprRune32RemBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) % bool(true)`, env,
		`invalid operation: rune(2147483647) % true (mismatched types rune and bool)`,
	)

}

// Test Rune32 == Int
func TestCheckBinaryTypedExprRune32EqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) == 4`, env, rune(0x7fffffff) == 4, reflect.TypeOf(rune(0x7fffffff) == 4))
}

// Test Rune32 == Rune
func TestCheckBinaryTypedExprRune32EqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) == '@'`, env, rune(0x7fffffff) == '@', reflect.TypeOf(rune(0x7fffffff) == '@'))
}

// Test Rune32 == Float
func TestCheckBinaryTypedExprRune32EqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) == 2.0`, env, rune(0x7fffffff) == 2.0, reflect.TypeOf(rune(0x7fffffff) == 2.0))
}

// Test Rune32 == Complex
func TestCheckBinaryTypedExprRune32EqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Rune32 == Bool
func TestCheckBinaryTypedExprRune32EqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == true`, env,
		`cannot convert true to type rune`,
		`invalid operation: rune(2147483647) == true (mismatched types rune and bool)`,
	)

}

// Test Rune32 == String
func TestCheckBinaryTypedExprRune32EqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: rune(2147483647) == "abc" (mismatched types rune and string)`,
	)

}

// Test Rune32 == Nil
func TestCheckBinaryTypedExprRune32EqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == nil`, env,
		`cannot convert nil to type rune`,
	)

}

// Test Rune32 == Int8
func TestCheckBinaryTypedExprRune32EqlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == int8(0x7f)`, env,
		`invalid operation: rune(2147483647) == 127 (mismatched types rune and int8)`,
	)

}

// Test Rune32 == Int16
func TestCheckBinaryTypedExprRune32EqlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == int16(0x7fff)`, env,
		`invalid operation: rune(2147483647) == 32767 (mismatched types rune and int16)`,
	)

}

// Test Rune32 == Int32
func TestCheckBinaryTypedExprRune32EqlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) == int32(0x7fffffff)`, env, rune(0x7fffffff) == int32(0x7fffffff), reflect.TypeOf(rune(0x7fffffff) == int32(0x7fffffff)))
}

// Test Rune32 == Int64
func TestCheckBinaryTypedExprRune32EqlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == int64(0x7fffffffffffffff)`, env,
		`invalid operation: rune(2147483647) == 9223372036854775807 (mismatched types rune and int64)`,
	)

}

// Test Rune32 == Uint8
func TestCheckBinaryTypedExprRune32EqlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == uint8(0xff)`, env,
		`invalid operation: rune(2147483647) == 255 (mismatched types rune and uint8)`,
	)

}

// Test Rune32 == Uint16
func TestCheckBinaryTypedExprRune32EqlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == uint16(0xffff)`, env,
		`invalid operation: rune(2147483647) == 65535 (mismatched types rune and uint16)`,
	)

}

// Test Rune32 == Uint32
func TestCheckBinaryTypedExprRune32EqlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == uint32(0xffffffff)`, env,
		`invalid operation: rune(2147483647) == 4294967295 (mismatched types rune and uint32)`,
	)

}

// Test Rune32 == Uint64
func TestCheckBinaryTypedExprRune32EqlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == uint64(0xffffffffffffffff)`, env,
		`invalid operation: rune(2147483647) == 18446744073709551615 (mismatched types rune and uint64)`,
	)

}

// Test Rune32 == Float32
func TestCheckBinaryTypedExprRune32EqlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == float32(0xffffffff)`, env,
		`invalid operation: rune(2147483647) == 4.29497e+09 (mismatched types rune and float32)`,
	)

}

// Test Rune32 == Float64
func TestCheckBinaryTypedExprRune32EqlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == float64(0xffffffff)`, env,
		`invalid operation: rune(2147483647) == 4.29497e+09 (mismatched types rune and float64)`,
	)

}

// Test Rune32 == Complex64
func TestCheckBinaryTypedExprRune32EqlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: rune(2147483647) == (4.29497e+09+4.29497e+09i) (mismatched types rune and complex64)`,
	)

}

// Test Rune32 == Complex128
func TestCheckBinaryTypedExprRune32EqlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: rune(2147483647) == (4.29497e+09+4.29497e+09i) (mismatched types rune and complex128)`,
	)

}

// Test Rune32 == Rune32
func TestCheckBinaryTypedExprRune32EqlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) == rune(0x7fffffff)`, env, rune(0x7fffffff) == rune(0x7fffffff), reflect.TypeOf(rune(0x7fffffff) == rune(0x7fffffff)))
}

// Test Rune32 == StringT
func TestCheckBinaryTypedExprRune32EqlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == string("abc")`, env,
		`invalid operation: rune(2147483647) == "abc" (mismatched types rune and string)`,
	)

}

// Test Rune32 == BoolT
func TestCheckBinaryTypedExprRune32EqlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) == bool(true)`, env,
		`invalid operation: rune(2147483647) == true (mismatched types rune and bool)`,
	)

}

// Test Rune32 > Int
func TestCheckBinaryTypedExprRune32GtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) > 4`, env, rune(0x7fffffff) > 4, reflect.TypeOf(rune(0x7fffffff) > 4))
}

// Test Rune32 > Rune
func TestCheckBinaryTypedExprRune32GtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) > '@'`, env, rune(0x7fffffff) > '@', reflect.TypeOf(rune(0x7fffffff) > '@'))
}

// Test Rune32 > Float
func TestCheckBinaryTypedExprRune32GtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) > 2.0`, env, rune(0x7fffffff) > 2.0, reflect.TypeOf(rune(0x7fffffff) > 2.0))
}

// Test Rune32 > Complex
func TestCheckBinaryTypedExprRune32GtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Rune32 > Bool
func TestCheckBinaryTypedExprRune32GtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > true`, env,
		`cannot convert true to type rune`,
		`invalid operation: rune(2147483647) > true (mismatched types rune and bool)`,
	)

}

// Test Rune32 > String
func TestCheckBinaryTypedExprRune32GtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > "abc"`, env,
		`cannot convert "abc" to type rune`,
		`invalid operation: rune(2147483647) > "abc" (mismatched types rune and string)`,
	)

}

// Test Rune32 > Nil
func TestCheckBinaryTypedExprRune32GtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > nil`, env,
		`cannot convert nil to type rune`,
	)

}

// Test Rune32 > Int8
func TestCheckBinaryTypedExprRune32GtrInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > int8(0x7f)`, env,
		`invalid operation: rune(2147483647) > 127 (mismatched types rune and int8)`,
	)

}

// Test Rune32 > Int16
func TestCheckBinaryTypedExprRune32GtrInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > int16(0x7fff)`, env,
		`invalid operation: rune(2147483647) > 32767 (mismatched types rune and int16)`,
	)

}

// Test Rune32 > Int32
func TestCheckBinaryTypedExprRune32GtrInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) > int32(0x7fffffff)`, env, rune(0x7fffffff) > int32(0x7fffffff), reflect.TypeOf(rune(0x7fffffff) > int32(0x7fffffff)))
}

// Test Rune32 > Int64
func TestCheckBinaryTypedExprRune32GtrInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > int64(0x7fffffffffffffff)`, env,
		`invalid operation: rune(2147483647) > 9223372036854775807 (mismatched types rune and int64)`,
	)

}

// Test Rune32 > Uint8
func TestCheckBinaryTypedExprRune32GtrUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > uint8(0xff)`, env,
		`invalid operation: rune(2147483647) > 255 (mismatched types rune and uint8)`,
	)

}

// Test Rune32 > Uint16
func TestCheckBinaryTypedExprRune32GtrUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > uint16(0xffff)`, env,
		`invalid operation: rune(2147483647) > 65535 (mismatched types rune and uint16)`,
	)

}

// Test Rune32 > Uint32
func TestCheckBinaryTypedExprRune32GtrUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > uint32(0xffffffff)`, env,
		`invalid operation: rune(2147483647) > 4294967295 (mismatched types rune and uint32)`,
	)

}

// Test Rune32 > Uint64
func TestCheckBinaryTypedExprRune32GtrUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > uint64(0xffffffffffffffff)`, env,
		`invalid operation: rune(2147483647) > 18446744073709551615 (mismatched types rune and uint64)`,
	)

}

// Test Rune32 > Float32
func TestCheckBinaryTypedExprRune32GtrFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > float32(0xffffffff)`, env,
		`invalid operation: rune(2147483647) > 4.29497e+09 (mismatched types rune and float32)`,
	)

}

// Test Rune32 > Float64
func TestCheckBinaryTypedExprRune32GtrFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > float64(0xffffffff)`, env,
		`invalid operation: rune(2147483647) > 4.29497e+09 (mismatched types rune and float64)`,
	)

}

// Test Rune32 > Complex64
func TestCheckBinaryTypedExprRune32GtrComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: rune(2147483647) > (4.29497e+09+4.29497e+09i) (mismatched types rune and complex64)`,
	)

}

// Test Rune32 > Complex128
func TestCheckBinaryTypedExprRune32GtrComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: rune(2147483647) > (4.29497e+09+4.29497e+09i) (mismatched types rune and complex128)`,
	)

}

// Test Rune32 > Rune32
func TestCheckBinaryTypedExprRune32GtrRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff) > rune(0x7fffffff)`, env, rune(0x7fffffff) > rune(0x7fffffff), reflect.TypeOf(rune(0x7fffffff) > rune(0x7fffffff)))
}

// Test Rune32 > StringT
func TestCheckBinaryTypedExprRune32GtrStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > string("abc")`, env,
		`invalid operation: rune(2147483647) > "abc" (mismatched types rune and string)`,
	)

}

// Test Rune32 > BoolT
func TestCheckBinaryTypedExprRune32GtrBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) > bool(true)`, env,
		`invalid operation: rune(2147483647) > true (mismatched types rune and bool)`,
	)

}

// Test Rune32 << Int
func TestCheckBinaryTypedExprRune32ShlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << 4`, env,
		`constant 34359738352 overflows rune`,
	)

}

// Test Rune32 << Rune
func TestCheckBinaryTypedExprRune32ShlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << '@'`, env,
		`constant 39614081238685424723062423552 overflows rune`,
	)

}

// Test Rune32 << Float
func TestCheckBinaryTypedExprRune32ShlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << 2.0`, env,
		`constant 8589934588 overflows rune`,
	)

}

// Test Rune32 << Complex
func TestCheckBinaryTypedExprRune32ShlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Rune32 << Bool
func TestCheckBinaryTypedExprRune32ShlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << true`, env,
		`invalid operation: rune(2147483647) << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Rune32 << String
func TestCheckBinaryTypedExprRune32ShlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: rune(2147483647) << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Rune32 << Nil
func TestCheckBinaryTypedExprRune32ShlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Rune32 << Int8
func TestCheckBinaryTypedExprRune32ShlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << int8(0x7f)`, env,
		`invalid operation: rune(2147483647) << 127 (shift count type int8, must be unsigned integer)`,
	)

}

// Test Rune32 << Int16
func TestCheckBinaryTypedExprRune32ShlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << int16(0x7fff)`, env,
		`invalid operation: rune(2147483647) << 32767 (shift count type int16, must be unsigned integer)`,
	)

}

// Test Rune32 << Int32
func TestCheckBinaryTypedExprRune32ShlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << int32(0x7fffffff)`, env,
		`invalid operation: rune(2147483647) << 2147483647 (shift count type int32, must be unsigned integer)`,
	)

}

// Test Rune32 << Int64
func TestCheckBinaryTypedExprRune32ShlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << int64(0x7fffffffffffffff)`, env,
		`invalid operation: rune(2147483647) << 9223372036854775807 (shift count type int64, must be unsigned integer)`,
	)

}

// Test Rune32 << Uint8
func TestCheckBinaryTypedExprRune32ShlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << uint8(0xff)`, env,
		`constant 124330809044550615920187464324919717520770083772701937027295712203561082249176779063296 overflows rune`,
	)

}

// Test Rune32 << Uint16
func TestCheckBinaryTypedExprRune32ShlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << uint16(0xffff)`, env,
		`stupid shift: 65535`,
	)

}

// Test Rune32 << Uint32
func TestCheckBinaryTypedExprRune32ShlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << uint32(0xffffffff)`, env,
		`stupid shift: 4294967295`,
	)

}

// Test Rune32 << Uint64
func TestCheckBinaryTypedExprRune32ShlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << uint64(0xffffffffffffffff)`, env,
		`stupid shift: -1`,
	)

}

// Test Rune32 << Float32
func TestCheckBinaryTypedExprRune32ShlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << float32(0xffffffff)`, env,
		`invalid operation: rune(2147483647) << 4.29497e+09 (shift count type float32, must be unsigned integer)`,
	)

}

// Test Rune32 << Float64
func TestCheckBinaryTypedExprRune32ShlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << float64(0xffffffff)`, env,
		`invalid operation: rune(2147483647) << 4.29497e+09 (shift count type float64, must be unsigned integer)`,
	)

}

// Test Rune32 << Complex64
func TestCheckBinaryTypedExprRune32ShlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: rune(2147483647) << (4.29497e+09+4.29497e+09i) (shift count type complex64, must be unsigned integer)`,
	)

}

// Test Rune32 << Complex128
func TestCheckBinaryTypedExprRune32ShlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: rune(2147483647) << (4.29497e+09+4.29497e+09i) (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Rune32 << Rune32
func TestCheckBinaryTypedExprRune32ShlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << rune(0x7fffffff)`, env,
		`invalid operation: rune(2147483647) << rune(2147483647) (shift count type rune, must be unsigned integer)`,
	)

}

// Test Rune32 << StringT
func TestCheckBinaryTypedExprRune32ShlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << string("abc")`, env,
		`invalid operation: rune(2147483647) << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Rune32 << BoolT
func TestCheckBinaryTypedExprRune32ShlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffff) << bool(true)`, env,
		`invalid operation: rune(2147483647) << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test StringT + Int
func TestCheckBinaryTypedExprStringTAddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + 4`, env,
		`cannot convert 4 to type string`,
		`invalid operation: "abc" + 4 (mismatched types string and int)`,
	)

}

// Test StringT + Rune
func TestCheckBinaryTypedExprStringTAddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + '@'`, env,
		`cannot convert '@' to type string`,
		`invalid operation: "abc" + rune(64) (mismatched types string and rune)`,
	)

}

// Test StringT + Float
func TestCheckBinaryTypedExprStringTAddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + 2.0`, env,
		`cannot convert 2 to type string`,
		`invalid operation: "abc" + 2 (mismatched types string and float64)`,
	)

}

// Test StringT + Complex
func TestCheckBinaryTypedExprStringTAddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + 8.0i`, env,
		`cannot convert 8i to type string`,
		`invalid operation: "abc" + 8i (mismatched types string and complex128)`,
	)

}

// Test StringT + Bool
func TestCheckBinaryTypedExprStringTAddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + true`, env,
		`cannot convert true to type string`,
		`invalid operation: "abc" + true (mismatched types string and bool)`,
	)

}

// Test StringT + String
func TestCheckBinaryTypedExprStringTAddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `string("abc") + "abc"`, env, string("abc") + "abc", reflect.TypeOf(string("abc") + "abc"))
}

// Test StringT + Nil
func TestCheckBinaryTypedExprStringTAddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + nil`, env,
		`invalid operation: "abc" + nil (mismatched types string and nil)`,
	)

}

// Test StringT + Int8
func TestCheckBinaryTypedExprStringTAddInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + int8(0x7f)`, env,
		`invalid operation: "abc" + 127 (mismatched types string and int8)`,
	)

}

// Test StringT + Int16
func TestCheckBinaryTypedExprStringTAddInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + int16(0x7fff)`, env,
		`invalid operation: "abc" + 32767 (mismatched types string and int16)`,
	)

}

// Test StringT + Int32
func TestCheckBinaryTypedExprStringTAddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + int32(0x7fffffff)`, env,
		`invalid operation: "abc" + 2147483647 (mismatched types string and int32)`,
	)

}

// Test StringT + Int64
func TestCheckBinaryTypedExprStringTAddInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + int64(0x7fffffffffffffff)`, env,
		`invalid operation: "abc" + 9223372036854775807 (mismatched types string and int64)`,
	)

}

// Test StringT + Uint8
func TestCheckBinaryTypedExprStringTAddUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + uint8(0xff)`, env,
		`invalid operation: "abc" + 255 (mismatched types string and uint8)`,
	)

}

// Test StringT + Uint16
func TestCheckBinaryTypedExprStringTAddUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + uint16(0xffff)`, env,
		`invalid operation: "abc" + 65535 (mismatched types string and uint16)`,
	)

}

// Test StringT + Uint32
func TestCheckBinaryTypedExprStringTAddUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + uint32(0xffffffff)`, env,
		`invalid operation: "abc" + 4294967295 (mismatched types string and uint32)`,
	)

}

// Test StringT + Uint64
func TestCheckBinaryTypedExprStringTAddUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + uint64(0xffffffffffffffff)`, env,
		`invalid operation: "abc" + 18446744073709551615 (mismatched types string and uint64)`,
	)

}

// Test StringT + Float32
func TestCheckBinaryTypedExprStringTAddFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + float32(0xffffffff)`, env,
		`invalid operation: "abc" + 4.29497e+09 (mismatched types string and float32)`,
	)

}

// Test StringT + Float64
func TestCheckBinaryTypedExprStringTAddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + float64(0xffffffff)`, env,
		`invalid operation: "abc" + 4.29497e+09 (mismatched types string and float64)`,
	)

}

// Test StringT + Complex64
func TestCheckBinaryTypedExprStringTAddComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: "abc" + (4.29497e+09+4.29497e+09i) (mismatched types string and complex64)`,
	)

}

// Test StringT + Complex128
func TestCheckBinaryTypedExprStringTAddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: "abc" + (4.29497e+09+4.29497e+09i) (mismatched types string and complex128)`,
	)

}

// Test StringT + Rune32
func TestCheckBinaryTypedExprStringTAddRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + rune(0x7fffffff)`, env,
		`invalid operation: "abc" + rune(2147483647) (mismatched types string and rune)`,
	)

}

// Test StringT + StringT
func TestCheckBinaryTypedExprStringTAddStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `string("abc") + string("abc")`, env, string("abc") + string("abc"), reflect.TypeOf(string("abc") + string("abc")))
}

// Test StringT + BoolT
func TestCheckBinaryTypedExprStringTAddBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") + bool(true)`, env,
		`invalid operation: "abc" + true (mismatched types string and bool)`,
	)

}

// Test StringT - Int
func TestCheckBinaryTypedExprStringTSubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - 4`, env,
		`cannot convert 4 to type string`,
		`invalid operation: "abc" - 4 (mismatched types string and int)`,
	)

}

// Test StringT - Rune
func TestCheckBinaryTypedExprStringTSubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - '@'`, env,
		`cannot convert '@' to type string`,
		`invalid operation: "abc" - rune(64) (mismatched types string and rune)`,
	)

}

// Test StringT - Float
func TestCheckBinaryTypedExprStringTSubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - 2.0`, env,
		`cannot convert 2 to type string`,
		`invalid operation: "abc" - 2 (mismatched types string and float64)`,
	)

}

// Test StringT - Complex
func TestCheckBinaryTypedExprStringTSubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - 8.0i`, env,
		`cannot convert 8i to type string`,
		`invalid operation: "abc" - 8i (mismatched types string and complex128)`,
	)

}

// Test StringT - Bool
func TestCheckBinaryTypedExprStringTSubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - true`, env,
		`cannot convert true to type string`,
		`invalid operation: "abc" - true (mismatched types string and bool)`,
	)

}

// Test StringT - String
func TestCheckBinaryTypedExprStringTSubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - "abc"`, env,
		`invalid operation: "abc" - "abc" (operator - not defined on string)`,
	)

}

// Test StringT - Nil
func TestCheckBinaryTypedExprStringTSubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - nil`, env,
		`invalid operation: "abc" - nil (mismatched types string and nil)`,
	)

}

// Test StringT - Int8
func TestCheckBinaryTypedExprStringTSubInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - int8(0x7f)`, env,
		`invalid operation: "abc" - 127 (mismatched types string and int8)`,
	)

}

// Test StringT - Int16
func TestCheckBinaryTypedExprStringTSubInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - int16(0x7fff)`, env,
		`invalid operation: "abc" - 32767 (mismatched types string and int16)`,
	)

}

// Test StringT - Int32
func TestCheckBinaryTypedExprStringTSubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - int32(0x7fffffff)`, env,
		`invalid operation: "abc" - 2147483647 (mismatched types string and int32)`,
	)

}

// Test StringT - Int64
func TestCheckBinaryTypedExprStringTSubInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - int64(0x7fffffffffffffff)`, env,
		`invalid operation: "abc" - 9223372036854775807 (mismatched types string and int64)`,
	)

}

// Test StringT - Uint8
func TestCheckBinaryTypedExprStringTSubUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - uint8(0xff)`, env,
		`invalid operation: "abc" - 255 (mismatched types string and uint8)`,
	)

}

// Test StringT - Uint16
func TestCheckBinaryTypedExprStringTSubUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - uint16(0xffff)`, env,
		`invalid operation: "abc" - 65535 (mismatched types string and uint16)`,
	)

}

// Test StringT - Uint32
func TestCheckBinaryTypedExprStringTSubUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - uint32(0xffffffff)`, env,
		`invalid operation: "abc" - 4294967295 (mismatched types string and uint32)`,
	)

}

// Test StringT - Uint64
func TestCheckBinaryTypedExprStringTSubUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - uint64(0xffffffffffffffff)`, env,
		`invalid operation: "abc" - 18446744073709551615 (mismatched types string and uint64)`,
	)

}

// Test StringT - Float32
func TestCheckBinaryTypedExprStringTSubFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - float32(0xffffffff)`, env,
		`invalid operation: "abc" - 4.29497e+09 (mismatched types string and float32)`,
	)

}

// Test StringT - Float64
func TestCheckBinaryTypedExprStringTSubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - float64(0xffffffff)`, env,
		`invalid operation: "abc" - 4.29497e+09 (mismatched types string and float64)`,
	)

}

// Test StringT - Complex64
func TestCheckBinaryTypedExprStringTSubComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: "abc" - (4.29497e+09+4.29497e+09i) (mismatched types string and complex64)`,
	)

}

// Test StringT - Complex128
func TestCheckBinaryTypedExprStringTSubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: "abc" - (4.29497e+09+4.29497e+09i) (mismatched types string and complex128)`,
	)

}

// Test StringT - Rune32
func TestCheckBinaryTypedExprStringTSubRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - rune(0x7fffffff)`, env,
		`invalid operation: "abc" - rune(2147483647) (mismatched types string and rune)`,
	)

}

// Test StringT - StringT
func TestCheckBinaryTypedExprStringTSubStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - string("abc")`, env,
		`invalid operation: "abc" - "abc" (operator - not defined on string)`,
	)

}

// Test StringT - BoolT
func TestCheckBinaryTypedExprStringTSubBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") - bool(true)`, env,
		`invalid operation: "abc" - true (mismatched types string and bool)`,
	)

}

// Test StringT & Int
func TestCheckBinaryTypedExprStringTAndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & 4`, env,
		`cannot convert 4 to type string`,
		`invalid operation: "abc" & 4 (mismatched types string and int)`,
	)

}

// Test StringT & Rune
func TestCheckBinaryTypedExprStringTAndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & '@'`, env,
		`cannot convert '@' to type string`,
		`invalid operation: "abc" & rune(64) (mismatched types string and rune)`,
	)

}

// Test StringT & Float
func TestCheckBinaryTypedExprStringTAndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & 2.0`, env,
		`cannot convert 2 to type string`,
		`invalid operation: "abc" & 2 (mismatched types string and float64)`,
	)

}

// Test StringT & Complex
func TestCheckBinaryTypedExprStringTAndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & 8.0i`, env,
		`cannot convert 8i to type string`,
		`invalid operation: "abc" & 8i (mismatched types string and complex128)`,
	)

}

// Test StringT & Bool
func TestCheckBinaryTypedExprStringTAndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & true`, env,
		`cannot convert true to type string`,
		`invalid operation: "abc" & true (mismatched types string and bool)`,
	)

}

// Test StringT & String
func TestCheckBinaryTypedExprStringTAndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & "abc"`, env,
		`invalid operation: "abc" & "abc" (operator & not defined on string)`,
	)

}

// Test StringT & Nil
func TestCheckBinaryTypedExprStringTAndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & nil`, env,
		`invalid operation: "abc" & nil (mismatched types string and nil)`,
	)

}

// Test StringT & Int8
func TestCheckBinaryTypedExprStringTAndInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & int8(0x7f)`, env,
		`invalid operation: "abc" & 127 (mismatched types string and int8)`,
	)

}

// Test StringT & Int16
func TestCheckBinaryTypedExprStringTAndInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & int16(0x7fff)`, env,
		`invalid operation: "abc" & 32767 (mismatched types string and int16)`,
	)

}

// Test StringT & Int32
func TestCheckBinaryTypedExprStringTAndInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & int32(0x7fffffff)`, env,
		`invalid operation: "abc" & 2147483647 (mismatched types string and int32)`,
	)

}

// Test StringT & Int64
func TestCheckBinaryTypedExprStringTAndInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & int64(0x7fffffffffffffff)`, env,
		`invalid operation: "abc" & 9223372036854775807 (mismatched types string and int64)`,
	)

}

// Test StringT & Uint8
func TestCheckBinaryTypedExprStringTAndUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & uint8(0xff)`, env,
		`invalid operation: "abc" & 255 (mismatched types string and uint8)`,
	)

}

// Test StringT & Uint16
func TestCheckBinaryTypedExprStringTAndUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & uint16(0xffff)`, env,
		`invalid operation: "abc" & 65535 (mismatched types string and uint16)`,
	)

}

// Test StringT & Uint32
func TestCheckBinaryTypedExprStringTAndUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & uint32(0xffffffff)`, env,
		`invalid operation: "abc" & 4294967295 (mismatched types string and uint32)`,
	)

}

// Test StringT & Uint64
func TestCheckBinaryTypedExprStringTAndUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & uint64(0xffffffffffffffff)`, env,
		`invalid operation: "abc" & 18446744073709551615 (mismatched types string and uint64)`,
	)

}

// Test StringT & Float32
func TestCheckBinaryTypedExprStringTAndFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & float32(0xffffffff)`, env,
		`invalid operation: "abc" & 4.29497e+09 (mismatched types string and float32)`,
	)

}

// Test StringT & Float64
func TestCheckBinaryTypedExprStringTAndFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & float64(0xffffffff)`, env,
		`invalid operation: "abc" & 4.29497e+09 (mismatched types string and float64)`,
	)

}

// Test StringT & Complex64
func TestCheckBinaryTypedExprStringTAndComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: "abc" & (4.29497e+09+4.29497e+09i) (mismatched types string and complex64)`,
	)

}

// Test StringT & Complex128
func TestCheckBinaryTypedExprStringTAndComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: "abc" & (4.29497e+09+4.29497e+09i) (mismatched types string and complex128)`,
	)

}

// Test StringT & Rune32
func TestCheckBinaryTypedExprStringTAndRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & rune(0x7fffffff)`, env,
		`invalid operation: "abc" & rune(2147483647) (mismatched types string and rune)`,
	)

}

// Test StringT & StringT
func TestCheckBinaryTypedExprStringTAndStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & string("abc")`, env,
		`invalid operation: "abc" & "abc" (operator & not defined on string)`,
	)

}

// Test StringT & BoolT
func TestCheckBinaryTypedExprStringTAndBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") & bool(true)`, env,
		`invalid operation: "abc" & true (mismatched types string and bool)`,
	)

}

// Test StringT % Int
func TestCheckBinaryTypedExprStringTRemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % 4`, env,
		`cannot convert 4 to type string`,
		`invalid operation: "abc" % 4 (mismatched types string and int)`,
	)

}

// Test StringT % Rune
func TestCheckBinaryTypedExprStringTRemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % '@'`, env,
		`cannot convert '@' to type string`,
		`invalid operation: "abc" % rune(64) (mismatched types string and rune)`,
	)

}

// Test StringT % Float
func TestCheckBinaryTypedExprStringTRemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % 2.0`, env,
		`cannot convert 2 to type string`,
		`invalid operation: "abc" % 2 (mismatched types string and float64)`,
	)

}

// Test StringT % Complex
func TestCheckBinaryTypedExprStringTRemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % 8.0i`, env,
		`cannot convert 8i to type string`,
		`invalid operation: "abc" % 8i (mismatched types string and complex128)`,
	)

}

// Test StringT % Bool
func TestCheckBinaryTypedExprStringTRemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % true`, env,
		`cannot convert true to type string`,
		`invalid operation: "abc" % true (mismatched types string and bool)`,
	)

}

// Test StringT % String
func TestCheckBinaryTypedExprStringTRemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % "abc"`, env,
		`invalid operation: "abc" % "abc" (operator % not defined on string)`,
	)

}

// Test StringT % Nil
func TestCheckBinaryTypedExprStringTRemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % nil`, env,
		`invalid operation: "abc" % nil (mismatched types string and nil)`,
	)

}

// Test StringT % Int8
func TestCheckBinaryTypedExprStringTRemInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % int8(0x7f)`, env,
		`invalid operation: "abc" % 127 (mismatched types string and int8)`,
	)

}

// Test StringT % Int16
func TestCheckBinaryTypedExprStringTRemInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % int16(0x7fff)`, env,
		`invalid operation: "abc" % 32767 (mismatched types string and int16)`,
	)

}

// Test StringT % Int32
func TestCheckBinaryTypedExprStringTRemInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % int32(0x7fffffff)`, env,
		`invalid operation: "abc" % 2147483647 (mismatched types string and int32)`,
	)

}

// Test StringT % Int64
func TestCheckBinaryTypedExprStringTRemInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % int64(0x7fffffffffffffff)`, env,
		`invalid operation: "abc" % 9223372036854775807 (mismatched types string and int64)`,
	)

}

// Test StringT % Uint8
func TestCheckBinaryTypedExprStringTRemUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % uint8(0xff)`, env,
		`invalid operation: "abc" % 255 (mismatched types string and uint8)`,
	)

}

// Test StringT % Uint16
func TestCheckBinaryTypedExprStringTRemUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % uint16(0xffff)`, env,
		`invalid operation: "abc" % 65535 (mismatched types string and uint16)`,
	)

}

// Test StringT % Uint32
func TestCheckBinaryTypedExprStringTRemUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % uint32(0xffffffff)`, env,
		`invalid operation: "abc" % 4294967295 (mismatched types string and uint32)`,
	)

}

// Test StringT % Uint64
func TestCheckBinaryTypedExprStringTRemUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % uint64(0xffffffffffffffff)`, env,
		`invalid operation: "abc" % 18446744073709551615 (mismatched types string and uint64)`,
	)

}

// Test StringT % Float32
func TestCheckBinaryTypedExprStringTRemFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % float32(0xffffffff)`, env,
		`invalid operation: "abc" % 4.29497e+09 (mismatched types string and float32)`,
	)

}

// Test StringT % Float64
func TestCheckBinaryTypedExprStringTRemFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % float64(0xffffffff)`, env,
		`invalid operation: "abc" % 4.29497e+09 (mismatched types string and float64)`,
	)

}

// Test StringT % Complex64
func TestCheckBinaryTypedExprStringTRemComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: "abc" % (4.29497e+09+4.29497e+09i) (mismatched types string and complex64)`,
	)

}

// Test StringT % Complex128
func TestCheckBinaryTypedExprStringTRemComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: "abc" % (4.29497e+09+4.29497e+09i) (mismatched types string and complex128)`,
	)

}

// Test StringT % Rune32
func TestCheckBinaryTypedExprStringTRemRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % rune(0x7fffffff)`, env,
		`invalid operation: "abc" % rune(2147483647) (mismatched types string and rune)`,
	)

}

// Test StringT % StringT
func TestCheckBinaryTypedExprStringTRemStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % string("abc")`, env,
		`invalid operation: "abc" % "abc" (operator % not defined on string)`,
	)

}

// Test StringT % BoolT
func TestCheckBinaryTypedExprStringTRemBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") % bool(true)`, env,
		`invalid operation: "abc" % true (mismatched types string and bool)`,
	)

}

// Test StringT == Int
func TestCheckBinaryTypedExprStringTEqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == 4`, env,
		`cannot convert 4 to type string`,
		`invalid operation: "abc" == 4 (mismatched types string and int)`,
	)

}

// Test StringT == Rune
func TestCheckBinaryTypedExprStringTEqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == '@'`, env,
		`cannot convert '@' to type string`,
		`invalid operation: "abc" == rune(64) (mismatched types string and rune)`,
	)

}

// Test StringT == Float
func TestCheckBinaryTypedExprStringTEqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == 2.0`, env,
		`cannot convert 2 to type string`,
		`invalid operation: "abc" == 2 (mismatched types string and float64)`,
	)

}

// Test StringT == Complex
func TestCheckBinaryTypedExprStringTEqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == 8.0i`, env,
		`cannot convert 8i to type string`,
		`invalid operation: "abc" == 8i (mismatched types string and complex128)`,
	)

}

// Test StringT == Bool
func TestCheckBinaryTypedExprStringTEqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == true`, env,
		`cannot convert true to type string`,
		`invalid operation: "abc" == true (mismatched types string and bool)`,
	)

}

// Test StringT == String
func TestCheckBinaryTypedExprStringTEqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `string("abc") == "abc"`, env, string("abc") == "abc", reflect.TypeOf(string("abc") == "abc"))
}

// Test StringT == Nil
func TestCheckBinaryTypedExprStringTEqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == nil`, env,
		`invalid operation: "abc" == nil (mismatched types string and nil)`,
	)

}

// Test StringT == Int8
func TestCheckBinaryTypedExprStringTEqlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == int8(0x7f)`, env,
		`invalid operation: "abc" == 127 (mismatched types string and int8)`,
	)

}

// Test StringT == Int16
func TestCheckBinaryTypedExprStringTEqlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == int16(0x7fff)`, env,
		`invalid operation: "abc" == 32767 (mismatched types string and int16)`,
	)

}

// Test StringT == Int32
func TestCheckBinaryTypedExprStringTEqlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == int32(0x7fffffff)`, env,
		`invalid operation: "abc" == 2147483647 (mismatched types string and int32)`,
	)

}

// Test StringT == Int64
func TestCheckBinaryTypedExprStringTEqlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == int64(0x7fffffffffffffff)`, env,
		`invalid operation: "abc" == 9223372036854775807 (mismatched types string and int64)`,
	)

}

// Test StringT == Uint8
func TestCheckBinaryTypedExprStringTEqlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == uint8(0xff)`, env,
		`invalid operation: "abc" == 255 (mismatched types string and uint8)`,
	)

}

// Test StringT == Uint16
func TestCheckBinaryTypedExprStringTEqlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == uint16(0xffff)`, env,
		`invalid operation: "abc" == 65535 (mismatched types string and uint16)`,
	)

}

// Test StringT == Uint32
func TestCheckBinaryTypedExprStringTEqlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == uint32(0xffffffff)`, env,
		`invalid operation: "abc" == 4294967295 (mismatched types string and uint32)`,
	)

}

// Test StringT == Uint64
func TestCheckBinaryTypedExprStringTEqlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == uint64(0xffffffffffffffff)`, env,
		`invalid operation: "abc" == 18446744073709551615 (mismatched types string and uint64)`,
	)

}

// Test StringT == Float32
func TestCheckBinaryTypedExprStringTEqlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == float32(0xffffffff)`, env,
		`invalid operation: "abc" == 4.29497e+09 (mismatched types string and float32)`,
	)

}

// Test StringT == Float64
func TestCheckBinaryTypedExprStringTEqlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == float64(0xffffffff)`, env,
		`invalid operation: "abc" == 4.29497e+09 (mismatched types string and float64)`,
	)

}

// Test StringT == Complex64
func TestCheckBinaryTypedExprStringTEqlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: "abc" == (4.29497e+09+4.29497e+09i) (mismatched types string and complex64)`,
	)

}

// Test StringT == Complex128
func TestCheckBinaryTypedExprStringTEqlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: "abc" == (4.29497e+09+4.29497e+09i) (mismatched types string and complex128)`,
	)

}

// Test StringT == Rune32
func TestCheckBinaryTypedExprStringTEqlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == rune(0x7fffffff)`, env,
		`invalid operation: "abc" == rune(2147483647) (mismatched types string and rune)`,
	)

}

// Test StringT == StringT
func TestCheckBinaryTypedExprStringTEqlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `string("abc") == string("abc")`, env, string("abc") == string("abc"), reflect.TypeOf(string("abc") == string("abc")))
}

// Test StringT == BoolT
func TestCheckBinaryTypedExprStringTEqlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") == bool(true)`, env,
		`invalid operation: "abc" == true (mismatched types string and bool)`,
	)

}

// Test StringT > Int
func TestCheckBinaryTypedExprStringTGtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > 4`, env,
		`cannot convert 4 to type string`,
		`invalid operation: "abc" > 4 (mismatched types string and int)`,
	)

}

// Test StringT > Rune
func TestCheckBinaryTypedExprStringTGtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > '@'`, env,
		`cannot convert '@' to type string`,
		`invalid operation: "abc" > rune(64) (mismatched types string and rune)`,
	)

}

// Test StringT > Float
func TestCheckBinaryTypedExprStringTGtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > 2.0`, env,
		`cannot convert 2 to type string`,
		`invalid operation: "abc" > 2 (mismatched types string and float64)`,
	)

}

// Test StringT > Complex
func TestCheckBinaryTypedExprStringTGtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > 8.0i`, env,
		`cannot convert 8i to type string`,
		`invalid operation: "abc" > 8i (mismatched types string and complex128)`,
	)

}

// Test StringT > Bool
func TestCheckBinaryTypedExprStringTGtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > true`, env,
		`cannot convert true to type string`,
		`invalid operation: "abc" > true (mismatched types string and bool)`,
	)

}

// Test StringT > String
func TestCheckBinaryTypedExprStringTGtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `string("abc") > "abc"`, env, string("abc") > "abc", reflect.TypeOf(string("abc") > "abc"))
}

// Test StringT > Nil
func TestCheckBinaryTypedExprStringTGtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > nil`, env,
		`invalid operation: "abc" > nil (mismatched types string and nil)`,
	)

}

// Test StringT > Int8
func TestCheckBinaryTypedExprStringTGtrInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > int8(0x7f)`, env,
		`invalid operation: "abc" > 127 (mismatched types string and int8)`,
	)

}

// Test StringT > Int16
func TestCheckBinaryTypedExprStringTGtrInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > int16(0x7fff)`, env,
		`invalid operation: "abc" > 32767 (mismatched types string and int16)`,
	)

}

// Test StringT > Int32
func TestCheckBinaryTypedExprStringTGtrInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > int32(0x7fffffff)`, env,
		`invalid operation: "abc" > 2147483647 (mismatched types string and int32)`,
	)

}

// Test StringT > Int64
func TestCheckBinaryTypedExprStringTGtrInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > int64(0x7fffffffffffffff)`, env,
		`invalid operation: "abc" > 9223372036854775807 (mismatched types string and int64)`,
	)

}

// Test StringT > Uint8
func TestCheckBinaryTypedExprStringTGtrUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > uint8(0xff)`, env,
		`invalid operation: "abc" > 255 (mismatched types string and uint8)`,
	)

}

// Test StringT > Uint16
func TestCheckBinaryTypedExprStringTGtrUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > uint16(0xffff)`, env,
		`invalid operation: "abc" > 65535 (mismatched types string and uint16)`,
	)

}

// Test StringT > Uint32
func TestCheckBinaryTypedExprStringTGtrUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > uint32(0xffffffff)`, env,
		`invalid operation: "abc" > 4294967295 (mismatched types string and uint32)`,
	)

}

// Test StringT > Uint64
func TestCheckBinaryTypedExprStringTGtrUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > uint64(0xffffffffffffffff)`, env,
		`invalid operation: "abc" > 18446744073709551615 (mismatched types string and uint64)`,
	)

}

// Test StringT > Float32
func TestCheckBinaryTypedExprStringTGtrFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > float32(0xffffffff)`, env,
		`invalid operation: "abc" > 4.29497e+09 (mismatched types string and float32)`,
	)

}

// Test StringT > Float64
func TestCheckBinaryTypedExprStringTGtrFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > float64(0xffffffff)`, env,
		`invalid operation: "abc" > 4.29497e+09 (mismatched types string and float64)`,
	)

}

// Test StringT > Complex64
func TestCheckBinaryTypedExprStringTGtrComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: "abc" > (4.29497e+09+4.29497e+09i) (mismatched types string and complex64)`,
	)

}

// Test StringT > Complex128
func TestCheckBinaryTypedExprStringTGtrComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: "abc" > (4.29497e+09+4.29497e+09i) (mismatched types string and complex128)`,
	)

}

// Test StringT > Rune32
func TestCheckBinaryTypedExprStringTGtrRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > rune(0x7fffffff)`, env,
		`invalid operation: "abc" > rune(2147483647) (mismatched types string and rune)`,
	)

}

// Test StringT > StringT
func TestCheckBinaryTypedExprStringTGtrStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `string("abc") > string("abc")`, env, string("abc") > string("abc"), reflect.TypeOf(string("abc") > string("abc")))
}

// Test StringT > BoolT
func TestCheckBinaryTypedExprStringTGtrBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") > bool(true)`, env,
		`invalid operation: "abc" > true (mismatched types string and bool)`,
	)

}

// Test StringT << Int
func TestCheckBinaryTypedExprStringTShlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << 4`, env,
		`invalid operation: "abc" << 4 (shift of type string)`,
	)

}

// Test StringT << Rune
func TestCheckBinaryTypedExprStringTShlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << '@'`, env,
		`invalid operation: "abc" << 64 (shift of type string)`,
	)

}

// Test StringT << Float
func TestCheckBinaryTypedExprStringTShlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << 2.0`, env,
		`invalid operation: "abc" << 2 (shift of type string)`,
	)

}

// Test StringT << Complex
func TestCheckBinaryTypedExprStringTShlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: "abc" << 0 (shift of type string)`,
	)

}

// Test StringT << Bool
func TestCheckBinaryTypedExprStringTShlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << true`, env,
		`invalid operation: "abc" << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test StringT << String
func TestCheckBinaryTypedExprStringTShlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: "abc" << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test StringT << Nil
func TestCheckBinaryTypedExprStringTShlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test StringT << Int8
func TestCheckBinaryTypedExprStringTShlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << int8(0x7f)`, env,
		`invalid operation: "abc" << 127 (shift count type int8, must be unsigned integer)`,
	)

}

// Test StringT << Int16
func TestCheckBinaryTypedExprStringTShlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << int16(0x7fff)`, env,
		`invalid operation: "abc" << 32767 (shift count type int16, must be unsigned integer)`,
	)

}

// Test StringT << Int32
func TestCheckBinaryTypedExprStringTShlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << int32(0x7fffffff)`, env,
		`invalid operation: "abc" << 2147483647 (shift count type int32, must be unsigned integer)`,
	)

}

// Test StringT << Int64
func TestCheckBinaryTypedExprStringTShlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << int64(0x7fffffffffffffff)`, env,
		`invalid operation: "abc" << 9223372036854775807 (shift count type int64, must be unsigned integer)`,
	)

}

// Test StringT << Uint8
func TestCheckBinaryTypedExprStringTShlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << uint8(0xff)`, env,
		`invalid operation: "abc" << 255 (shift of type string)`,
	)

}

// Test StringT << Uint16
func TestCheckBinaryTypedExprStringTShlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << uint16(0xffff)`, env,
		`invalid operation: "abc" << 65535 (shift of type string)`,
	)

}

// Test StringT << Uint32
func TestCheckBinaryTypedExprStringTShlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << uint32(0xffffffff)`, env,
		`invalid operation: "abc" << 4294967295 (shift of type string)`,
	)

}

// Test StringT << Uint64
func TestCheckBinaryTypedExprStringTShlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << uint64(0xffffffffffffffff)`, env,
		`invalid operation: "abc" << 18446744073709551615 (shift of type string)`,
	)

}

// Test StringT << Float32
func TestCheckBinaryTypedExprStringTShlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << float32(0xffffffff)`, env,
		`invalid operation: "abc" << 4.29497e+09 (shift count type float32, must be unsigned integer)`,
	)

}

// Test StringT << Float64
func TestCheckBinaryTypedExprStringTShlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << float64(0xffffffff)`, env,
		`invalid operation: "abc" << 4.29497e+09 (shift count type float64, must be unsigned integer)`,
	)

}

// Test StringT << Complex64
func TestCheckBinaryTypedExprStringTShlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: "abc" << (4.29497e+09+4.29497e+09i) (shift count type complex64, must be unsigned integer)`,
	)

}

// Test StringT << Complex128
func TestCheckBinaryTypedExprStringTShlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: "abc" << (4.29497e+09+4.29497e+09i) (shift count type complex128, must be unsigned integer)`,
	)

}

// Test StringT << Rune32
func TestCheckBinaryTypedExprStringTShlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << rune(0x7fffffff)`, env,
		`invalid operation: "abc" << rune(2147483647) (shift count type rune, must be unsigned integer)`,
	)

}

// Test StringT << StringT
func TestCheckBinaryTypedExprStringTShlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << string("abc")`, env,
		`invalid operation: "abc" << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test StringT << BoolT
func TestCheckBinaryTypedExprStringTShlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string("abc") << bool(true)`, env,
		`invalid operation: "abc" << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test BoolT + Int
func TestCheckBinaryTypedExprBoolTAddInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true + 4 (mismatched types bool and int)`,
	)

}

// Test BoolT + Rune
func TestCheckBinaryTypedExprBoolTAddRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true + rune(64) (mismatched types bool and rune)`,
	)

}

// Test BoolT + Float
func TestCheckBinaryTypedExprBoolTAddFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true + 2 (mismatched types bool and float64)`,
	)

}

// Test BoolT + Complex
func TestCheckBinaryTypedExprBoolTAddComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true + 8i (mismatched types bool and complex128)`,
	)

}

// Test BoolT + Bool
func TestCheckBinaryTypedExprBoolTAddBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + true`, env,
		`invalid operation: true + true (operator + not defined on bool)`,
	)

}

// Test BoolT + String
func TestCheckBinaryTypedExprBoolTAddString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true + "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT + Nil
func TestCheckBinaryTypedExprBoolTAddNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + nil`, env,
		`cannot convert nil to type bool`,
	)

}

// Test BoolT + Int8
func TestCheckBinaryTypedExprBoolTAddInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + int8(0x7f)`, env,
		`invalid operation: true + 127 (mismatched types bool and int8)`,
	)

}

// Test BoolT + Int16
func TestCheckBinaryTypedExprBoolTAddInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + int16(0x7fff)`, env,
		`invalid operation: true + 32767 (mismatched types bool and int16)`,
	)

}

// Test BoolT + Int32
func TestCheckBinaryTypedExprBoolTAddInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + int32(0x7fffffff)`, env,
		`invalid operation: true + 2147483647 (mismatched types bool and int32)`,
	)

}

// Test BoolT + Int64
func TestCheckBinaryTypedExprBoolTAddInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + int64(0x7fffffffffffffff)`, env,
		`invalid operation: true + 9223372036854775807 (mismatched types bool and int64)`,
	)

}

// Test BoolT + Uint8
func TestCheckBinaryTypedExprBoolTAddUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + uint8(0xff)`, env,
		`invalid operation: true + 255 (mismatched types bool and uint8)`,
	)

}

// Test BoolT + Uint16
func TestCheckBinaryTypedExprBoolTAddUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + uint16(0xffff)`, env,
		`invalid operation: true + 65535 (mismatched types bool and uint16)`,
	)

}

// Test BoolT + Uint32
func TestCheckBinaryTypedExprBoolTAddUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + uint32(0xffffffff)`, env,
		`invalid operation: true + 4294967295 (mismatched types bool and uint32)`,
	)

}

// Test BoolT + Uint64
func TestCheckBinaryTypedExprBoolTAddUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + uint64(0xffffffffffffffff)`, env,
		`invalid operation: true + 18446744073709551615 (mismatched types bool and uint64)`,
	)

}

// Test BoolT + Float32
func TestCheckBinaryTypedExprBoolTAddFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + float32(0xffffffff)`, env,
		`invalid operation: true + 4.29497e+09 (mismatched types bool and float32)`,
	)

}

// Test BoolT + Float64
func TestCheckBinaryTypedExprBoolTAddFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + float64(0xffffffff)`, env,
		`invalid operation: true + 4.29497e+09 (mismatched types bool and float64)`,
	)

}

// Test BoolT + Complex64
func TestCheckBinaryTypedExprBoolTAddComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: true + (4.29497e+09+4.29497e+09i) (mismatched types bool and complex64)`,
	)

}

// Test BoolT + Complex128
func TestCheckBinaryTypedExprBoolTAddComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: true + (4.29497e+09+4.29497e+09i) (mismatched types bool and complex128)`,
	)

}

// Test BoolT + Rune32
func TestCheckBinaryTypedExprBoolTAddRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + rune(0x7fffffff)`, env,
		`invalid operation: true + rune(2147483647) (mismatched types bool and rune)`,
	)

}

// Test BoolT + StringT
func TestCheckBinaryTypedExprBoolTAddStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + string("abc")`, env,
		`invalid operation: true + "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT + BoolT
func TestCheckBinaryTypedExprBoolTAddBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) + bool(true)`, env,
		`invalid operation: true + true (operator + not defined on bool)`,
	)

}

// Test BoolT - Int
func TestCheckBinaryTypedExprBoolTSubInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true - 4 (mismatched types bool and int)`,
	)

}

// Test BoolT - Rune
func TestCheckBinaryTypedExprBoolTSubRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true - rune(64) (mismatched types bool and rune)`,
	)

}

// Test BoolT - Float
func TestCheckBinaryTypedExprBoolTSubFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true - 2 (mismatched types bool and float64)`,
	)

}

// Test BoolT - Complex
func TestCheckBinaryTypedExprBoolTSubComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true - 8i (mismatched types bool and complex128)`,
	)

}

// Test BoolT - Bool
func TestCheckBinaryTypedExprBoolTSubBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - true`, env,
		`invalid operation: true - true (operator - not defined on bool)`,
	)

}

// Test BoolT - String
func TestCheckBinaryTypedExprBoolTSubString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true - "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT - Nil
func TestCheckBinaryTypedExprBoolTSubNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - nil`, env,
		`cannot convert nil to type bool`,
	)

}

// Test BoolT - Int8
func TestCheckBinaryTypedExprBoolTSubInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - int8(0x7f)`, env,
		`invalid operation: true - 127 (mismatched types bool and int8)`,
	)

}

// Test BoolT - Int16
func TestCheckBinaryTypedExprBoolTSubInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - int16(0x7fff)`, env,
		`invalid operation: true - 32767 (mismatched types bool and int16)`,
	)

}

// Test BoolT - Int32
func TestCheckBinaryTypedExprBoolTSubInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - int32(0x7fffffff)`, env,
		`invalid operation: true - 2147483647 (mismatched types bool and int32)`,
	)

}

// Test BoolT - Int64
func TestCheckBinaryTypedExprBoolTSubInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - int64(0x7fffffffffffffff)`, env,
		`invalid operation: true - 9223372036854775807 (mismatched types bool and int64)`,
	)

}

// Test BoolT - Uint8
func TestCheckBinaryTypedExprBoolTSubUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - uint8(0xff)`, env,
		`invalid operation: true - 255 (mismatched types bool and uint8)`,
	)

}

// Test BoolT - Uint16
func TestCheckBinaryTypedExprBoolTSubUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - uint16(0xffff)`, env,
		`invalid operation: true - 65535 (mismatched types bool and uint16)`,
	)

}

// Test BoolT - Uint32
func TestCheckBinaryTypedExprBoolTSubUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - uint32(0xffffffff)`, env,
		`invalid operation: true - 4294967295 (mismatched types bool and uint32)`,
	)

}

// Test BoolT - Uint64
func TestCheckBinaryTypedExprBoolTSubUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - uint64(0xffffffffffffffff)`, env,
		`invalid operation: true - 18446744073709551615 (mismatched types bool and uint64)`,
	)

}

// Test BoolT - Float32
func TestCheckBinaryTypedExprBoolTSubFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - float32(0xffffffff)`, env,
		`invalid operation: true - 4.29497e+09 (mismatched types bool and float32)`,
	)

}

// Test BoolT - Float64
func TestCheckBinaryTypedExprBoolTSubFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - float64(0xffffffff)`, env,
		`invalid operation: true - 4.29497e+09 (mismatched types bool and float64)`,
	)

}

// Test BoolT - Complex64
func TestCheckBinaryTypedExprBoolTSubComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: true - (4.29497e+09+4.29497e+09i) (mismatched types bool and complex64)`,
	)

}

// Test BoolT - Complex128
func TestCheckBinaryTypedExprBoolTSubComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: true - (4.29497e+09+4.29497e+09i) (mismatched types bool and complex128)`,
	)

}

// Test BoolT - Rune32
func TestCheckBinaryTypedExprBoolTSubRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - rune(0x7fffffff)`, env,
		`invalid operation: true - rune(2147483647) (mismatched types bool and rune)`,
	)

}

// Test BoolT - StringT
func TestCheckBinaryTypedExprBoolTSubStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - string("abc")`, env,
		`invalid operation: true - "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT - BoolT
func TestCheckBinaryTypedExprBoolTSubBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) - bool(true)`, env,
		`invalid operation: true - true (operator - not defined on bool)`,
	)

}

// Test BoolT & Int
func TestCheckBinaryTypedExprBoolTAndInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true & 4 (mismatched types bool and int)`,
	)

}

// Test BoolT & Rune
func TestCheckBinaryTypedExprBoolTAndRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true & rune(64) (mismatched types bool and rune)`,
	)

}

// Test BoolT & Float
func TestCheckBinaryTypedExprBoolTAndFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true & 2 (mismatched types bool and float64)`,
	)

}

// Test BoolT & Complex
func TestCheckBinaryTypedExprBoolTAndComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true & 8i (mismatched types bool and complex128)`,
	)

}

// Test BoolT & Bool
func TestCheckBinaryTypedExprBoolTAndBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & true`, env,
		`invalid operation: true & true (operator & not defined on bool)`,
	)

}

// Test BoolT & String
func TestCheckBinaryTypedExprBoolTAndString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true & "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT & Nil
func TestCheckBinaryTypedExprBoolTAndNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & nil`, env,
		`cannot convert nil to type bool`,
	)

}

// Test BoolT & Int8
func TestCheckBinaryTypedExprBoolTAndInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & int8(0x7f)`, env,
		`invalid operation: true & 127 (mismatched types bool and int8)`,
	)

}

// Test BoolT & Int16
func TestCheckBinaryTypedExprBoolTAndInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & int16(0x7fff)`, env,
		`invalid operation: true & 32767 (mismatched types bool and int16)`,
	)

}

// Test BoolT & Int32
func TestCheckBinaryTypedExprBoolTAndInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & int32(0x7fffffff)`, env,
		`invalid operation: true & 2147483647 (mismatched types bool and int32)`,
	)

}

// Test BoolT & Int64
func TestCheckBinaryTypedExprBoolTAndInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & int64(0x7fffffffffffffff)`, env,
		`invalid operation: true & 9223372036854775807 (mismatched types bool and int64)`,
	)

}

// Test BoolT & Uint8
func TestCheckBinaryTypedExprBoolTAndUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & uint8(0xff)`, env,
		`invalid operation: true & 255 (mismatched types bool and uint8)`,
	)

}

// Test BoolT & Uint16
func TestCheckBinaryTypedExprBoolTAndUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & uint16(0xffff)`, env,
		`invalid operation: true & 65535 (mismatched types bool and uint16)`,
	)

}

// Test BoolT & Uint32
func TestCheckBinaryTypedExprBoolTAndUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & uint32(0xffffffff)`, env,
		`invalid operation: true & 4294967295 (mismatched types bool and uint32)`,
	)

}

// Test BoolT & Uint64
func TestCheckBinaryTypedExprBoolTAndUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & uint64(0xffffffffffffffff)`, env,
		`invalid operation: true & 18446744073709551615 (mismatched types bool and uint64)`,
	)

}

// Test BoolT & Float32
func TestCheckBinaryTypedExprBoolTAndFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & float32(0xffffffff)`, env,
		`invalid operation: true & 4.29497e+09 (mismatched types bool and float32)`,
	)

}

// Test BoolT & Float64
func TestCheckBinaryTypedExprBoolTAndFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & float64(0xffffffff)`, env,
		`invalid operation: true & 4.29497e+09 (mismatched types bool and float64)`,
	)

}

// Test BoolT & Complex64
func TestCheckBinaryTypedExprBoolTAndComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: true & (4.29497e+09+4.29497e+09i) (mismatched types bool and complex64)`,
	)

}

// Test BoolT & Complex128
func TestCheckBinaryTypedExprBoolTAndComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: true & (4.29497e+09+4.29497e+09i) (mismatched types bool and complex128)`,
	)

}

// Test BoolT & Rune32
func TestCheckBinaryTypedExprBoolTAndRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & rune(0x7fffffff)`, env,
		`invalid operation: true & rune(2147483647) (mismatched types bool and rune)`,
	)

}

// Test BoolT & StringT
func TestCheckBinaryTypedExprBoolTAndStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & string("abc")`, env,
		`invalid operation: true & "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT & BoolT
func TestCheckBinaryTypedExprBoolTAndBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) & bool(true)`, env,
		`invalid operation: true & true (operator & not defined on bool)`,
	)

}

// Test BoolT % Int
func TestCheckBinaryTypedExprBoolTRemInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true % 4 (mismatched types bool and int)`,
	)

}

// Test BoolT % Rune
func TestCheckBinaryTypedExprBoolTRemRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true % rune(64) (mismatched types bool and rune)`,
	)

}

// Test BoolT % Float
func TestCheckBinaryTypedExprBoolTRemFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true % 2 (mismatched types bool and float64)`,
	)

}

// Test BoolT % Complex
func TestCheckBinaryTypedExprBoolTRemComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true % 8i (mismatched types bool and complex128)`,
	)

}

// Test BoolT % Bool
func TestCheckBinaryTypedExprBoolTRemBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % true`, env,
		`invalid operation: true % true (operator % not defined on bool)`,
	)

}

// Test BoolT % String
func TestCheckBinaryTypedExprBoolTRemString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true % "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT % Nil
func TestCheckBinaryTypedExprBoolTRemNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % nil`, env,
		`cannot convert nil to type bool`,
	)

}

// Test BoolT % Int8
func TestCheckBinaryTypedExprBoolTRemInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % int8(0x7f)`, env,
		`invalid operation: true % 127 (mismatched types bool and int8)`,
	)

}

// Test BoolT % Int16
func TestCheckBinaryTypedExprBoolTRemInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % int16(0x7fff)`, env,
		`invalid operation: true % 32767 (mismatched types bool and int16)`,
	)

}

// Test BoolT % Int32
func TestCheckBinaryTypedExprBoolTRemInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % int32(0x7fffffff)`, env,
		`invalid operation: true % 2147483647 (mismatched types bool and int32)`,
	)

}

// Test BoolT % Int64
func TestCheckBinaryTypedExprBoolTRemInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % int64(0x7fffffffffffffff)`, env,
		`invalid operation: true % 9223372036854775807 (mismatched types bool and int64)`,
	)

}

// Test BoolT % Uint8
func TestCheckBinaryTypedExprBoolTRemUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % uint8(0xff)`, env,
		`invalid operation: true % 255 (mismatched types bool and uint8)`,
	)

}

// Test BoolT % Uint16
func TestCheckBinaryTypedExprBoolTRemUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % uint16(0xffff)`, env,
		`invalid operation: true % 65535 (mismatched types bool and uint16)`,
	)

}

// Test BoolT % Uint32
func TestCheckBinaryTypedExprBoolTRemUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % uint32(0xffffffff)`, env,
		`invalid operation: true % 4294967295 (mismatched types bool and uint32)`,
	)

}

// Test BoolT % Uint64
func TestCheckBinaryTypedExprBoolTRemUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % uint64(0xffffffffffffffff)`, env,
		`invalid operation: true % 18446744073709551615 (mismatched types bool and uint64)`,
	)

}

// Test BoolT % Float32
func TestCheckBinaryTypedExprBoolTRemFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % float32(0xffffffff)`, env,
		`invalid operation: true % 4.29497e+09 (mismatched types bool and float32)`,
	)

}

// Test BoolT % Float64
func TestCheckBinaryTypedExprBoolTRemFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % float64(0xffffffff)`, env,
		`invalid operation: true % 4.29497e+09 (mismatched types bool and float64)`,
	)

}

// Test BoolT % Complex64
func TestCheckBinaryTypedExprBoolTRemComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: true % (4.29497e+09+4.29497e+09i) (mismatched types bool and complex64)`,
	)

}

// Test BoolT % Complex128
func TestCheckBinaryTypedExprBoolTRemComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: true % (4.29497e+09+4.29497e+09i) (mismatched types bool and complex128)`,
	)

}

// Test BoolT % Rune32
func TestCheckBinaryTypedExprBoolTRemRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % rune(0x7fffffff)`, env,
		`invalid operation: true % rune(2147483647) (mismatched types bool and rune)`,
	)

}

// Test BoolT % StringT
func TestCheckBinaryTypedExprBoolTRemStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % string("abc")`, env,
		`invalid operation: true % "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT % BoolT
func TestCheckBinaryTypedExprBoolTRemBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) % bool(true)`, env,
		`invalid operation: true % true (operator % not defined on bool)`,
	)

}

// Test BoolT == Int
func TestCheckBinaryTypedExprBoolTEqlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true == 4 (mismatched types bool and int)`,
	)

}

// Test BoolT == Rune
func TestCheckBinaryTypedExprBoolTEqlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true == rune(64) (mismatched types bool and rune)`,
	)

}

// Test BoolT == Float
func TestCheckBinaryTypedExprBoolTEqlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true == 2 (mismatched types bool and float64)`,
	)

}

// Test BoolT == Complex
func TestCheckBinaryTypedExprBoolTEqlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true == 8i (mismatched types bool and complex128)`,
	)

}

// Test BoolT == Bool
func TestCheckBinaryTypedExprBoolTEqlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `bool(true) == true`, env, bool(true) == true, reflect.TypeOf(bool(true) == true))
}

// Test BoolT == String
func TestCheckBinaryTypedExprBoolTEqlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true == "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT == Nil
func TestCheckBinaryTypedExprBoolTEqlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == nil`, env,
		`cannot convert nil to type bool`,
	)

}

// Test BoolT == Int8
func TestCheckBinaryTypedExprBoolTEqlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == int8(0x7f)`, env,
		`invalid operation: true == 127 (mismatched types bool and int8)`,
	)

}

// Test BoolT == Int16
func TestCheckBinaryTypedExprBoolTEqlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == int16(0x7fff)`, env,
		`invalid operation: true == 32767 (mismatched types bool and int16)`,
	)

}

// Test BoolT == Int32
func TestCheckBinaryTypedExprBoolTEqlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == int32(0x7fffffff)`, env,
		`invalid operation: true == 2147483647 (mismatched types bool and int32)`,
	)

}

// Test BoolT == Int64
func TestCheckBinaryTypedExprBoolTEqlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == int64(0x7fffffffffffffff)`, env,
		`invalid operation: true == 9223372036854775807 (mismatched types bool and int64)`,
	)

}

// Test BoolT == Uint8
func TestCheckBinaryTypedExprBoolTEqlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == uint8(0xff)`, env,
		`invalid operation: true == 255 (mismatched types bool and uint8)`,
	)

}

// Test BoolT == Uint16
func TestCheckBinaryTypedExprBoolTEqlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == uint16(0xffff)`, env,
		`invalid operation: true == 65535 (mismatched types bool and uint16)`,
	)

}

// Test BoolT == Uint32
func TestCheckBinaryTypedExprBoolTEqlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == uint32(0xffffffff)`, env,
		`invalid operation: true == 4294967295 (mismatched types bool and uint32)`,
	)

}

// Test BoolT == Uint64
func TestCheckBinaryTypedExprBoolTEqlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == uint64(0xffffffffffffffff)`, env,
		`invalid operation: true == 18446744073709551615 (mismatched types bool and uint64)`,
	)

}

// Test BoolT == Float32
func TestCheckBinaryTypedExprBoolTEqlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == float32(0xffffffff)`, env,
		`invalid operation: true == 4.29497e+09 (mismatched types bool and float32)`,
	)

}

// Test BoolT == Float64
func TestCheckBinaryTypedExprBoolTEqlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == float64(0xffffffff)`, env,
		`invalid operation: true == 4.29497e+09 (mismatched types bool and float64)`,
	)

}

// Test BoolT == Complex64
func TestCheckBinaryTypedExprBoolTEqlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: true == (4.29497e+09+4.29497e+09i) (mismatched types bool and complex64)`,
	)

}

// Test BoolT == Complex128
func TestCheckBinaryTypedExprBoolTEqlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: true == (4.29497e+09+4.29497e+09i) (mismatched types bool and complex128)`,
	)

}

// Test BoolT == Rune32
func TestCheckBinaryTypedExprBoolTEqlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == rune(0x7fffffff)`, env,
		`invalid operation: true == rune(2147483647) (mismatched types bool and rune)`,
	)

}

// Test BoolT == StringT
func TestCheckBinaryTypedExprBoolTEqlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) == string("abc")`, env,
		`invalid operation: true == "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT == BoolT
func TestCheckBinaryTypedExprBoolTEqlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `bool(true) == bool(true)`, env, bool(true) == bool(true), reflect.TypeOf(bool(true) == bool(true)))
}

// Test BoolT > Int
func TestCheckBinaryTypedExprBoolTGtrInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: true > 4 (mismatched types bool and int)`,
	)

}

// Test BoolT > Rune
func TestCheckBinaryTypedExprBoolTGtrRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: true > rune(64) (mismatched types bool and rune)`,
	)

}

// Test BoolT > Float
func TestCheckBinaryTypedExprBoolTGtrFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: true > 2 (mismatched types bool and float64)`,
	)

}

// Test BoolT > Complex
func TestCheckBinaryTypedExprBoolTGtrComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: true > 8i (mismatched types bool and complex128)`,
	)

}

// Test BoolT > Bool
func TestCheckBinaryTypedExprBoolTGtrBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > true`, env,
		`invalid operation: true > true (operator > not defined on bool)`,
	)

}

// Test BoolT > String
func TestCheckBinaryTypedExprBoolTGtrString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: true > "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT > Nil
func TestCheckBinaryTypedExprBoolTGtrNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > nil`, env,
		`cannot convert nil to type bool`,
	)

}

// Test BoolT > Int8
func TestCheckBinaryTypedExprBoolTGtrInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > int8(0x7f)`, env,
		`invalid operation: true > 127 (mismatched types bool and int8)`,
	)

}

// Test BoolT > Int16
func TestCheckBinaryTypedExprBoolTGtrInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > int16(0x7fff)`, env,
		`invalid operation: true > 32767 (mismatched types bool and int16)`,
	)

}

// Test BoolT > Int32
func TestCheckBinaryTypedExprBoolTGtrInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > int32(0x7fffffff)`, env,
		`invalid operation: true > 2147483647 (mismatched types bool and int32)`,
	)

}

// Test BoolT > Int64
func TestCheckBinaryTypedExprBoolTGtrInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > int64(0x7fffffffffffffff)`, env,
		`invalid operation: true > 9223372036854775807 (mismatched types bool and int64)`,
	)

}

// Test BoolT > Uint8
func TestCheckBinaryTypedExprBoolTGtrUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > uint8(0xff)`, env,
		`invalid operation: true > 255 (mismatched types bool and uint8)`,
	)

}

// Test BoolT > Uint16
func TestCheckBinaryTypedExprBoolTGtrUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > uint16(0xffff)`, env,
		`invalid operation: true > 65535 (mismatched types bool and uint16)`,
	)

}

// Test BoolT > Uint32
func TestCheckBinaryTypedExprBoolTGtrUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > uint32(0xffffffff)`, env,
		`invalid operation: true > 4294967295 (mismatched types bool and uint32)`,
	)

}

// Test BoolT > Uint64
func TestCheckBinaryTypedExprBoolTGtrUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > uint64(0xffffffffffffffff)`, env,
		`invalid operation: true > 18446744073709551615 (mismatched types bool and uint64)`,
	)

}

// Test BoolT > Float32
func TestCheckBinaryTypedExprBoolTGtrFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > float32(0xffffffff)`, env,
		`invalid operation: true > 4.29497e+09 (mismatched types bool and float32)`,
	)

}

// Test BoolT > Float64
func TestCheckBinaryTypedExprBoolTGtrFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > float64(0xffffffff)`, env,
		`invalid operation: true > 4.29497e+09 (mismatched types bool and float64)`,
	)

}

// Test BoolT > Complex64
func TestCheckBinaryTypedExprBoolTGtrComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: true > (4.29497e+09+4.29497e+09i) (mismatched types bool and complex64)`,
	)

}

// Test BoolT > Complex128
func TestCheckBinaryTypedExprBoolTGtrComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: true > (4.29497e+09+4.29497e+09i) (mismatched types bool and complex128)`,
	)

}

// Test BoolT > Rune32
func TestCheckBinaryTypedExprBoolTGtrRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > rune(0x7fffffff)`, env,
		`invalid operation: true > rune(2147483647) (mismatched types bool and rune)`,
	)

}

// Test BoolT > StringT
func TestCheckBinaryTypedExprBoolTGtrStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > string("abc")`, env,
		`invalid operation: true > "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT > BoolT
func TestCheckBinaryTypedExprBoolTGtrBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) > bool(true)`, env,
		`invalid operation: true > true (operator > not defined on bool)`,
	)

}

// Test BoolT << Int
func TestCheckBinaryTypedExprBoolTShlInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << 4`, env,
		`invalid operation: true << 4 (shift of type bool)`,
	)

}

// Test BoolT << Rune
func TestCheckBinaryTypedExprBoolTShlRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << '@'`, env,
		`invalid operation: true << 64 (shift of type bool)`,
	)

}

// Test BoolT << Float
func TestCheckBinaryTypedExprBoolTShlFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << 2.0`, env,
		`invalid operation: true << 2 (shift of type bool)`,
	)

}

// Test BoolT << Complex
func TestCheckBinaryTypedExprBoolTShlComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: true << 0 (shift of type bool)`,
	)

}

// Test BoolT << Bool
func TestCheckBinaryTypedExprBoolTShlBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << true`, env,
		`invalid operation: true << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test BoolT << String
func TestCheckBinaryTypedExprBoolTShlString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: true << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test BoolT << Nil
func TestCheckBinaryTypedExprBoolTShlNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test BoolT << Int8
func TestCheckBinaryTypedExprBoolTShlInt8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << int8(0x7f)`, env,
		`invalid operation: true << 127 (shift count type int8, must be unsigned integer)`,
	)

}

// Test BoolT << Int16
func TestCheckBinaryTypedExprBoolTShlInt16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << int16(0x7fff)`, env,
		`invalid operation: true << 32767 (shift count type int16, must be unsigned integer)`,
	)

}

// Test BoolT << Int32
func TestCheckBinaryTypedExprBoolTShlInt32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << int32(0x7fffffff)`, env,
		`invalid operation: true << 2147483647 (shift count type int32, must be unsigned integer)`,
	)

}

// Test BoolT << Int64
func TestCheckBinaryTypedExprBoolTShlInt64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << int64(0x7fffffffffffffff)`, env,
		`invalid operation: true << 9223372036854775807 (shift count type int64, must be unsigned integer)`,
	)

}

// Test BoolT << Uint8
func TestCheckBinaryTypedExprBoolTShlUint8(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << uint8(0xff)`, env,
		`invalid operation: true << 255 (shift of type bool)`,
	)

}

// Test BoolT << Uint16
func TestCheckBinaryTypedExprBoolTShlUint16(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << uint16(0xffff)`, env,
		`invalid operation: true << 65535 (shift of type bool)`,
	)

}

// Test BoolT << Uint32
func TestCheckBinaryTypedExprBoolTShlUint32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << uint32(0xffffffff)`, env,
		`invalid operation: true << 4294967295 (shift of type bool)`,
	)

}

// Test BoolT << Uint64
func TestCheckBinaryTypedExprBoolTShlUint64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << uint64(0xffffffffffffffff)`, env,
		`invalid operation: true << 18446744073709551615 (shift of type bool)`,
	)

}

// Test BoolT << Float32
func TestCheckBinaryTypedExprBoolTShlFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << float32(0xffffffff)`, env,
		`invalid operation: true << 4.29497e+09 (shift count type float32, must be unsigned integer)`,
	)

}

// Test BoolT << Float64
func TestCheckBinaryTypedExprBoolTShlFloat64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << float64(0xffffffff)`, env,
		`invalid operation: true << 4.29497e+09 (shift count type float64, must be unsigned integer)`,
	)

}

// Test BoolT << Complex64
func TestCheckBinaryTypedExprBoolTShlComplex64(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << complex64(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: true << (4.29497e+09+4.29497e+09i) (shift count type complex64, must be unsigned integer)`,
	)

}

// Test BoolT << Complex128
func TestCheckBinaryTypedExprBoolTShlComplex128(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << complex128(0xffffffff + 0xffffffff * 1i)`, env,
		`invalid operation: true << (4.29497e+09+4.29497e+09i) (shift count type complex128, must be unsigned integer)`,
	)

}

// Test BoolT << Rune32
func TestCheckBinaryTypedExprBoolTShlRune32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << rune(0x7fffffff)`, env,
		`invalid operation: true << rune(2147483647) (shift count type rune, must be unsigned integer)`,
	)

}

// Test BoolT << StringT
func TestCheckBinaryTypedExprBoolTShlStringT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << string("abc")`, env,
		`invalid operation: true << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test BoolT << BoolT
func TestCheckBinaryTypedExprBoolTShlBoolT(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `bool(true) << bool(true)`, env,
		`invalid operation: true << true (shift count type bool, must be unsigned integer)`,
	)

}
