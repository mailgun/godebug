package eval

import (
	"testing"
	"reflect"
)

// Test int8(0x7f)
func TestCheckCallExprInt8From7bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(0x7f)`, env, int8(0x7f), reflect.TypeOf(int8(0x7f)))
}

// Test int8(0xff)
func TestCheckCallExprInt8From8bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0xff)`, env,
		`constant 255 overflows int8`,
	)

}

// Test int8(0x7fff)
func TestCheckCallExprInt8From15bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7fff)`, env,
		`constant 32767 overflows int8`,
	)

}

// Test int8(0xffff)
func TestCheckCallExprInt8From16bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0xffff)`, env,
		`constant 65535 overflows int8`,
	)

}

// Test int8(0x7fffffff)
func TestCheckCallExprInt8From31bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7fffffff)`, env,
		`constant 2147483647 overflows int8`,
	)

}

// Test int8(0xffffffff)
func TestCheckCallExprInt8From32bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0xffffffff)`, env,
		`constant 4294967295 overflows int8`,
	)

}

// Test int8(0x7fffffffffffffff)
func TestCheckCallExprInt8From63bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0x7fffffffffffffff)`, env,
		`constant 9223372036854775807 overflows int8`,
	)

}

// Test int8(0xffffffffffffffff)
func TestCheckCallExprInt8From64bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(0xffffffffffffffff)`, env,
		`constant 18446744073709551615 overflows int8`,
	)

}

// Test int8('d')
func TestCheckCallExprInt8FromRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8('d')`, env, int8('d'), reflect.TypeOf(int8('d')))
}

// Test int8('日')
func TestCheckCallExprInt8FromWideRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8('日')`, env,
		`constant 26085 overflows int8`,
	)

}

// Test int8(1.0)
func TestCheckCallExprInt8FromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(1.0)`, env, int8(1.0), reflect.TypeOf(int8(1.0)))
}

// Test int8(1.5)
func TestCheckCallExprInt8FromFloating(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test int8(1.0+0i)
func TestCheckCallExprInt8FromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int8(1.0+0i)`, env, int8(1.0+0i), reflect.TypeOf(int8(1.0+0i)))
}

// Test int8(1.5+0i)
func TestCheckCallExprInt8FromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(1.5+0i)`, env,
		`constant 1.5+0i truncated to integer`,
	)

}

// Test int8(1.5+1.5i)
func TestCheckCallExprInt8FromComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(1.5+1.5i)`, env,
		`constant 1.5+1.5i truncated to integer`,
		`constant 1.5+1.5i truncated to real`,
	)

}

// Test int8(true)
func TestCheckCallExprInt8FromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(true)`, env,
		`cannot convert true to type int8`,
		`cannot convert true (type bool) to type int8`,
	)

}

// Test int8("abc")
func TestCheckCallExprInt8FromString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8("abc")`, env,
		`cannot convert "abc" to type int8`,
		`cannot convert "abc" (type string) to type int8`,
	)

}

// Test int8(nil)
func TestCheckCallExprInt8FromNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int8(nil)`, env,
		`cannot convert nil to type int8`,
	)

}

// Test int16(0x7f)
func TestCheckCallExprInt16From7bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7f)`, env, int16(0x7f), reflect.TypeOf(int16(0x7f)))
}

// Test int16(0xff)
func TestCheckCallExprInt16From8bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0xff)`, env, int16(0xff), reflect.TypeOf(int16(0xff)))
}

// Test int16(0x7fff)
func TestCheckCallExprInt16From15bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(0x7fff)`, env, int16(0x7fff), reflect.TypeOf(int16(0x7fff)))
}

// Test int16(0xffff)
func TestCheckCallExprInt16From16bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0xffff)`, env,
		`constant 65535 overflows int16`,
	)

}

// Test int16(0x7fffffff)
func TestCheckCallExprInt16From31bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fffffff)`, env,
		`constant 2147483647 overflows int16`,
	)

}

// Test int16(0xffffffff)
func TestCheckCallExprInt16From32bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0xffffffff)`, env,
		`constant 4294967295 overflows int16`,
	)

}

// Test int16(0x7fffffffffffffff)
func TestCheckCallExprInt16From63bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0x7fffffffffffffff)`, env,
		`constant 9223372036854775807 overflows int16`,
	)

}

// Test int16(0xffffffffffffffff)
func TestCheckCallExprInt16From64bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(0xffffffffffffffff)`, env,
		`constant 18446744073709551615 overflows int16`,
	)

}

// Test int16('d')
func TestCheckCallExprInt16FromRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16('d')`, env, int16('d'), reflect.TypeOf(int16('d')))
}

// Test int16('日')
func TestCheckCallExprInt16FromWideRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16('日')`, env, int16('日'), reflect.TypeOf(int16('日')))
}

// Test int16(1.0)
func TestCheckCallExprInt16FromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(1.0)`, env, int16(1.0), reflect.TypeOf(int16(1.0)))
}

// Test int16(1.5)
func TestCheckCallExprInt16FromFloating(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test int16(1.0+0i)
func TestCheckCallExprInt16FromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int16(1.0+0i)`, env, int16(1.0+0i), reflect.TypeOf(int16(1.0+0i)))
}

// Test int16(1.5+0i)
func TestCheckCallExprInt16FromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(1.5+0i)`, env,
		`constant 1.5+0i truncated to integer`,
	)

}

// Test int16(1.5+1.5i)
func TestCheckCallExprInt16FromComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(1.5+1.5i)`, env,
		`constant 1.5+1.5i truncated to integer`,
		`constant 1.5+1.5i truncated to real`,
	)

}

// Test int16(true)
func TestCheckCallExprInt16FromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(true)`, env,
		`cannot convert true to type int16`,
		`cannot convert true (type bool) to type int16`,
	)

}

// Test int16("abc")
func TestCheckCallExprInt16FromString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16("abc")`, env,
		`cannot convert "abc" to type int16`,
		`cannot convert "abc" (type string) to type int16`,
	)

}

// Test int16(nil)
func TestCheckCallExprInt16FromNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int16(nil)`, env,
		`cannot convert nil to type int16`,
	)

}

// Test int32(0x7f)
func TestCheckCallExprInt32From7bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7f)`, env, int32(0x7f), reflect.TypeOf(int32(0x7f)))
}

// Test int32(0xff)
func TestCheckCallExprInt32From8bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0xff)`, env, int32(0xff), reflect.TypeOf(int32(0xff)))
}

// Test int32(0x7fff)
func TestCheckCallExprInt32From15bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fff)`, env, int32(0x7fff), reflect.TypeOf(int32(0x7fff)))
}

// Test int32(0xffff)
func TestCheckCallExprInt32From16bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0xffff)`, env, int32(0xffff), reflect.TypeOf(int32(0xffff)))
}

// Test int32(0x7fffffff)
func TestCheckCallExprInt32From31bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(0x7fffffff)`, env, int32(0x7fffffff), reflect.TypeOf(int32(0x7fffffff)))
}

// Test int32(0xffffffff)
func TestCheckCallExprInt32From32bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0xffffffff)`, env,
		`constant 4294967295 overflows int32`,
	)

}

// Test int32(0x7fffffffffffffff)
func TestCheckCallExprInt32From63bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0x7fffffffffffffff)`, env,
		`constant 9223372036854775807 overflows int32`,
	)

}

// Test int32(0xffffffffffffffff)
func TestCheckCallExprInt32From64bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(0xffffffffffffffff)`, env,
		`constant 18446744073709551615 overflows int32`,
	)

}

// Test int32('d')
func TestCheckCallExprInt32FromRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32('d')`, env, int32('d'), reflect.TypeOf(int32('d')))
}

// Test int32('日')
func TestCheckCallExprInt32FromWideRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32('日')`, env, int32('日'), reflect.TypeOf(int32('日')))
}

// Test int32(1.0)
func TestCheckCallExprInt32FromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(1.0)`, env, int32(1.0), reflect.TypeOf(int32(1.0)))
}

// Test int32(1.5)
func TestCheckCallExprInt32FromFloating(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test int32(1.0+0i)
func TestCheckCallExprInt32FromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int32(1.0+0i)`, env, int32(1.0+0i), reflect.TypeOf(int32(1.0+0i)))
}

// Test int32(1.5+0i)
func TestCheckCallExprInt32FromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(1.5+0i)`, env,
		`constant 1.5+0i truncated to integer`,
	)

}

// Test int32(1.5+1.5i)
func TestCheckCallExprInt32FromComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(1.5+1.5i)`, env,
		`constant 1.5+1.5i truncated to integer`,
		`constant 1.5+1.5i truncated to real`,
	)

}

// Test int32(true)
func TestCheckCallExprInt32FromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(true)`, env,
		`cannot convert true to type int32`,
		`cannot convert true (type bool) to type int32`,
	)

}

// Test int32("abc")
func TestCheckCallExprInt32FromString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32("abc")`, env,
		`cannot convert "abc" to type int32`,
		`cannot convert "abc" (type string) to type int32`,
	)

}

// Test int32(nil)
func TestCheckCallExprInt32FromNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int32(nil)`, env,
		`cannot convert nil to type int32`,
	)

}

// Test int64(0x7f)
func TestCheckCallExprInt64From7bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7f)`, env, int64(0x7f), reflect.TypeOf(int64(0x7f)))
}

// Test int64(0xff)
func TestCheckCallExprInt64From8bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0xff)`, env, int64(0xff), reflect.TypeOf(int64(0xff)))
}

// Test int64(0x7fff)
func TestCheckCallExprInt64From15bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fff)`, env, int64(0x7fff), reflect.TypeOf(int64(0x7fff)))
}

// Test int64(0xffff)
func TestCheckCallExprInt64From16bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0xffff)`, env, int64(0xffff), reflect.TypeOf(int64(0xffff)))
}

// Test int64(0x7fffffff)
func TestCheckCallExprInt64From31bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffff)`, env, int64(0x7fffffff), reflect.TypeOf(int64(0x7fffffff)))
}

// Test int64(0xffffffff)
func TestCheckCallExprInt64From32bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0xffffffff)`, env, int64(0xffffffff), reflect.TypeOf(int64(0xffffffff)))
}

// Test int64(0x7fffffffffffffff)
func TestCheckCallExprInt64From63bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(0x7fffffffffffffff)`, env, int64(0x7fffffffffffffff), reflect.TypeOf(int64(0x7fffffffffffffff)))
}

// Test int64(0xffffffffffffffff)
func TestCheckCallExprInt64From64bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(0xffffffffffffffff)`, env,
		`constant 18446744073709551615 overflows int64`,
	)

}

// Test int64('d')
func TestCheckCallExprInt64FromRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64('d')`, env, int64('d'), reflect.TypeOf(int64('d')))
}

// Test int64('日')
func TestCheckCallExprInt64FromWideRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64('日')`, env, int64('日'), reflect.TypeOf(int64('日')))
}

// Test int64(1.0)
func TestCheckCallExprInt64FromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(1.0)`, env, int64(1.0), reflect.TypeOf(int64(1.0)))
}

// Test int64(1.5)
func TestCheckCallExprInt64FromFloating(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test int64(1.0+0i)
func TestCheckCallExprInt64FromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `int64(1.0+0i)`, env, int64(1.0+0i), reflect.TypeOf(int64(1.0+0i)))
}

// Test int64(1.5+0i)
func TestCheckCallExprInt64FromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(1.5+0i)`, env,
		`constant 1.5+0i truncated to integer`,
	)

}

// Test int64(1.5+1.5i)
func TestCheckCallExprInt64FromComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(1.5+1.5i)`, env,
		`constant 1.5+1.5i truncated to integer`,
		`constant 1.5+1.5i truncated to real`,
	)

}

// Test int64(true)
func TestCheckCallExprInt64FromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(true)`, env,
		`cannot convert true to type int64`,
		`cannot convert true (type bool) to type int64`,
	)

}

// Test int64("abc")
func TestCheckCallExprInt64FromString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64("abc")`, env,
		`cannot convert "abc" to type int64`,
		`cannot convert "abc" (type string) to type int64`,
	)

}

// Test int64(nil)
func TestCheckCallExprInt64FromNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `int64(nil)`, env,
		`cannot convert nil to type int64`,
	)

}

// Test uint8(0x7f)
func TestCheckCallExprUint8From7bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0x7f)`, env, uint8(0x7f), reflect.TypeOf(uint8(0x7f)))
}

// Test uint8(0xff)
func TestCheckCallExprUint8From8bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(0xff)`, env, uint8(0xff), reflect.TypeOf(uint8(0xff)))
}

// Test uint8(0x7fff)
func TestCheckCallExprUint8From15bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0x7fff)`, env,
		`constant 32767 overflows uint8`,
	)

}

// Test uint8(0xffff)
func TestCheckCallExprUint8From16bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xffff)`, env,
		`constant 65535 overflows uint8`,
	)

}

// Test uint8(0x7fffffff)
func TestCheckCallExprUint8From31bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0x7fffffff)`, env,
		`constant 2147483647 overflows uint8`,
	)

}

// Test uint8(0xffffffff)
func TestCheckCallExprUint8From32bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xffffffff)`, env,
		`constant 4294967295 overflows uint8`,
	)

}

// Test uint8(0x7fffffffffffffff)
func TestCheckCallExprUint8From63bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0x7fffffffffffffff)`, env,
		`constant 9223372036854775807 overflows uint8`,
	)

}

// Test uint8(0xffffffffffffffff)
func TestCheckCallExprUint8From64bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(0xffffffffffffffff)`, env,
		`constant 18446744073709551615 overflows uint8`,
	)

}

// Test uint8('d')
func TestCheckCallExprUint8FromRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8('d')`, env, uint8('d'), reflect.TypeOf(uint8('d')))
}

// Test uint8('日')
func TestCheckCallExprUint8FromWideRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8('日')`, env,
		`constant 26085 overflows uint8`,
	)

}

// Test uint8(1.0)
func TestCheckCallExprUint8FromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(1.0)`, env, uint8(1.0), reflect.TypeOf(uint8(1.0)))
}

// Test uint8(1.5)
func TestCheckCallExprUint8FromFloating(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test uint8(1.0+0i)
func TestCheckCallExprUint8FromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint8(1.0+0i)`, env, uint8(1.0+0i), reflect.TypeOf(uint8(1.0+0i)))
}

// Test uint8(1.5+0i)
func TestCheckCallExprUint8FromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(1.5+0i)`, env,
		`constant 1.5+0i truncated to integer`,
	)

}

// Test uint8(1.5+1.5i)
func TestCheckCallExprUint8FromComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(1.5+1.5i)`, env,
		`constant 1.5+1.5i truncated to integer`,
		`constant 1.5+1.5i truncated to real`,
	)

}

// Test uint8(true)
func TestCheckCallExprUint8FromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(true)`, env,
		`cannot convert true to type uint8`,
		`cannot convert true (type bool) to type uint8`,
	)

}

// Test uint8("abc")
func TestCheckCallExprUint8FromString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8("abc")`, env,
		`cannot convert "abc" to type uint8`,
		`cannot convert "abc" (type string) to type uint8`,
	)

}

// Test uint8(nil)
func TestCheckCallExprUint8FromNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint8(nil)`, env,
		`cannot convert nil to type uint8`,
	)

}

// Test uint16(0x7f)
func TestCheckCallExprUint16From7bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0x7f)`, env, uint16(0x7f), reflect.TypeOf(uint16(0x7f)))
}

// Test uint16(0xff)
func TestCheckCallExprUint16From8bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xff)`, env, uint16(0xff), reflect.TypeOf(uint16(0xff)))
}

// Test uint16(0x7fff)
func TestCheckCallExprUint16From15bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0x7fff)`, env, uint16(0x7fff), reflect.TypeOf(uint16(0x7fff)))
}

// Test uint16(0xffff)
func TestCheckCallExprUint16From16bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(0xffff)`, env, uint16(0xffff), reflect.TypeOf(uint16(0xffff)))
}

// Test uint16(0x7fffffff)
func TestCheckCallExprUint16From31bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0x7fffffff)`, env,
		`constant 2147483647 overflows uint16`,
	)

}

// Test uint16(0xffffffff)
func TestCheckCallExprUint16From32bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffffffff)`, env,
		`constant 4294967295 overflows uint16`,
	)

}

// Test uint16(0x7fffffffffffffff)
func TestCheckCallExprUint16From63bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0x7fffffffffffffff)`, env,
		`constant 9223372036854775807 overflows uint16`,
	)

}

// Test uint16(0xffffffffffffffff)
func TestCheckCallExprUint16From64bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(0xffffffffffffffff)`, env,
		`constant 18446744073709551615 overflows uint16`,
	)

}

// Test uint16('d')
func TestCheckCallExprUint16FromRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16('d')`, env, uint16('d'), reflect.TypeOf(uint16('d')))
}

// Test uint16('日')
func TestCheckCallExprUint16FromWideRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16('日')`, env, uint16('日'), reflect.TypeOf(uint16('日')))
}

// Test uint16(1.0)
func TestCheckCallExprUint16FromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(1.0)`, env, uint16(1.0), reflect.TypeOf(uint16(1.0)))
}

// Test uint16(1.5)
func TestCheckCallExprUint16FromFloating(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test uint16(1.0+0i)
func TestCheckCallExprUint16FromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint16(1.0+0i)`, env, uint16(1.0+0i), reflect.TypeOf(uint16(1.0+0i)))
}

// Test uint16(1.5+0i)
func TestCheckCallExprUint16FromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(1.5+0i)`, env,
		`constant 1.5+0i truncated to integer`,
	)

}

// Test uint16(1.5+1.5i)
func TestCheckCallExprUint16FromComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(1.5+1.5i)`, env,
		`constant 1.5+1.5i truncated to integer`,
		`constant 1.5+1.5i truncated to real`,
	)

}

// Test uint16(true)
func TestCheckCallExprUint16FromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(true)`, env,
		`cannot convert true to type uint16`,
		`cannot convert true (type bool) to type uint16`,
	)

}

// Test uint16("abc")
func TestCheckCallExprUint16FromString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16("abc")`, env,
		`cannot convert "abc" to type uint16`,
		`cannot convert "abc" (type string) to type uint16`,
	)

}

// Test uint16(nil)
func TestCheckCallExprUint16FromNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint16(nil)`, env,
		`cannot convert nil to type uint16`,
	)

}

// Test uint32(0x7f)
func TestCheckCallExprUint32From7bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0x7f)`, env, uint32(0x7f), reflect.TypeOf(uint32(0x7f)))
}

// Test uint32(0xff)
func TestCheckCallExprUint32From8bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xff)`, env, uint32(0xff), reflect.TypeOf(uint32(0xff)))
}

// Test uint32(0x7fff)
func TestCheckCallExprUint32From15bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0x7fff)`, env, uint32(0x7fff), reflect.TypeOf(uint32(0x7fff)))
}

// Test uint32(0xffff)
func TestCheckCallExprUint32From16bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffff)`, env, uint32(0xffff), reflect.TypeOf(uint32(0xffff)))
}

// Test uint32(0x7fffffff)
func TestCheckCallExprUint32From31bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0x7fffffff)`, env, uint32(0x7fffffff), reflect.TypeOf(uint32(0x7fffffff)))
}

// Test uint32(0xffffffff)
func TestCheckCallExprUint32From32bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(0xffffffff)`, env, uint32(0xffffffff), reflect.TypeOf(uint32(0xffffffff)))
}

// Test uint32(0x7fffffffffffffff)
func TestCheckCallExprUint32From63bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0x7fffffffffffffff)`, env,
		`constant 9223372036854775807 overflows uint32`,
	)

}

// Test uint32(0xffffffffffffffff)
func TestCheckCallExprUint32From64bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(0xffffffffffffffff)`, env,
		`constant 18446744073709551615 overflows uint32`,
	)

}

// Test uint32('d')
func TestCheckCallExprUint32FromRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32('d')`, env, uint32('d'), reflect.TypeOf(uint32('d')))
}

// Test uint32('日')
func TestCheckCallExprUint32FromWideRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32('日')`, env, uint32('日'), reflect.TypeOf(uint32('日')))
}

// Test uint32(1.0)
func TestCheckCallExprUint32FromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(1.0)`, env, uint32(1.0), reflect.TypeOf(uint32(1.0)))
}

// Test uint32(1.5)
func TestCheckCallExprUint32FromFloating(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test uint32(1.0+0i)
func TestCheckCallExprUint32FromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint32(1.0+0i)`, env, uint32(1.0+0i), reflect.TypeOf(uint32(1.0+0i)))
}

// Test uint32(1.5+0i)
func TestCheckCallExprUint32FromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(1.5+0i)`, env,
		`constant 1.5+0i truncated to integer`,
	)

}

// Test uint32(1.5+1.5i)
func TestCheckCallExprUint32FromComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(1.5+1.5i)`, env,
		`constant 1.5+1.5i truncated to integer`,
		`constant 1.5+1.5i truncated to real`,
	)

}

// Test uint32(true)
func TestCheckCallExprUint32FromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(true)`, env,
		`cannot convert true to type uint32`,
		`cannot convert true (type bool) to type uint32`,
	)

}

// Test uint32("abc")
func TestCheckCallExprUint32FromString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32("abc")`, env,
		`cannot convert "abc" to type uint32`,
		`cannot convert "abc" (type string) to type uint32`,
	)

}

// Test uint32(nil)
func TestCheckCallExprUint32FromNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint32(nil)`, env,
		`cannot convert nil to type uint32`,
	)

}

// Test uint64(0x7f)
func TestCheckCallExprUint64From7bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0x7f)`, env, uint64(0x7f), reflect.TypeOf(uint64(0x7f)))
}

// Test uint64(0xff)
func TestCheckCallExprUint64From8bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xff)`, env, uint64(0xff), reflect.TypeOf(uint64(0xff)))
}

// Test uint64(0x7fff)
func TestCheckCallExprUint64From15bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0x7fff)`, env, uint64(0x7fff), reflect.TypeOf(uint64(0x7fff)))
}

// Test uint64(0xffff)
func TestCheckCallExprUint64From16bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffff)`, env, uint64(0xffff), reflect.TypeOf(uint64(0xffff)))
}

// Test uint64(0x7fffffff)
func TestCheckCallExprUint64From31bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0x7fffffff)`, env, uint64(0x7fffffff), reflect.TypeOf(uint64(0x7fffffff)))
}

// Test uint64(0xffffffff)
func TestCheckCallExprUint64From32bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffff)`, env, uint64(0xffffffff), reflect.TypeOf(uint64(0xffffffff)))
}

// Test uint64(0x7fffffffffffffff)
func TestCheckCallExprUint64From63bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0x7fffffffffffffff)`, env, uint64(0x7fffffffffffffff), reflect.TypeOf(uint64(0x7fffffffffffffff)))
}

// Test uint64(0xffffffffffffffff)
func TestCheckCallExprUint64From64bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(0xffffffffffffffff)`, env, uint64(0xffffffffffffffff), reflect.TypeOf(uint64(0xffffffffffffffff)))
}

// Test uint64('d')
func TestCheckCallExprUint64FromRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64('d')`, env, uint64('d'), reflect.TypeOf(uint64('d')))
}

// Test uint64('日')
func TestCheckCallExprUint64FromWideRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64('日')`, env, uint64('日'), reflect.TypeOf(uint64('日')))
}

// Test uint64(1.0)
func TestCheckCallExprUint64FromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(1.0)`, env, uint64(1.0), reflect.TypeOf(uint64(1.0)))
}

// Test uint64(1.5)
func TestCheckCallExprUint64FromFloating(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test uint64(1.0+0i)
func TestCheckCallExprUint64FromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `uint64(1.0+0i)`, env, uint64(1.0+0i), reflect.TypeOf(uint64(1.0+0i)))
}

// Test uint64(1.5+0i)
func TestCheckCallExprUint64FromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(1.5+0i)`, env,
		`constant 1.5+0i truncated to integer`,
	)

}

// Test uint64(1.5+1.5i)
func TestCheckCallExprUint64FromComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(1.5+1.5i)`, env,
		`constant 1.5+1.5i truncated to integer`,
		`constant 1.5+1.5i truncated to real`,
	)

}

// Test uint64(true)
func TestCheckCallExprUint64FromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(true)`, env,
		`cannot convert true to type uint64`,
		`cannot convert true (type bool) to type uint64`,
	)

}

// Test uint64("abc")
func TestCheckCallExprUint64FromString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64("abc")`, env,
		`cannot convert "abc" to type uint64`,
		`cannot convert "abc" (type string) to type uint64`,
	)

}

// Test uint64(nil)
func TestCheckCallExprUint64FromNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `uint64(nil)`, env,
		`cannot convert nil to type uint64`,
	)

}

// Test float32(0x7f)
func TestCheckCallExprFloat32From7bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0x7f)`, env, float32(0x7f), reflect.TypeOf(float32(0x7f)))
}

// Test float32(0xff)
func TestCheckCallExprFloat32From8bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xff)`, env, float32(0xff), reflect.TypeOf(float32(0xff)))
}

// Test float32(0x7fff)
func TestCheckCallExprFloat32From15bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0x7fff)`, env, float32(0x7fff), reflect.TypeOf(float32(0x7fff)))
}

// Test float32(0xffff)
func TestCheckCallExprFloat32From16bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffff)`, env, float32(0xffff), reflect.TypeOf(float32(0xffff)))
}

// Test float32(0x7fffffff)
func TestCheckCallExprFloat32From31bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0x7fffffff)`, env, float32(0x7fffffff), reflect.TypeOf(float32(0x7fffffff)))
}

// Test float32(0xffffffff)
func TestCheckCallExprFloat32From32bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffff)`, env, float32(0xffffffff), reflect.TypeOf(float32(0xffffffff)))
}

// Test float32(0x7fffffffffffffff)
func TestCheckCallExprFloat32From63bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0x7fffffffffffffff)`, env, float32(0x7fffffffffffffff), reflect.TypeOf(float32(0x7fffffffffffffff)))
}

// Test float32(0xffffffffffffffff)
func TestCheckCallExprFloat32From64bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(0xffffffffffffffff)`, env, float32(0xffffffffffffffff), reflect.TypeOf(float32(0xffffffffffffffff)))
}

// Test float32('d')
func TestCheckCallExprFloat32FromRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32('d')`, env, float32('d'), reflect.TypeOf(float32('d')))
}

// Test float32('日')
func TestCheckCallExprFloat32FromWideRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32('日')`, env, float32('日'), reflect.TypeOf(float32('日')))
}

// Test float32(1.0)
func TestCheckCallExprFloat32FromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(1.0)`, env, float32(1.0), reflect.TypeOf(float32(1.0)))
}

// Test float32(1.5)
func TestCheckCallExprFloat32FromFloating(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(1.5)`, env, float32(1.5), reflect.TypeOf(float32(1.5)))
}

// Test float32(1.0+0i)
func TestCheckCallExprFloat32FromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(1.0+0i)`, env, float32(1.0+0i), reflect.TypeOf(float32(1.0+0i)))
}

// Test float32(1.5+0i)
func TestCheckCallExprFloat32FromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float32(1.5+0i)`, env, float32(1.5+0i), reflect.TypeOf(float32(1.5+0i)))
}

// Test float32(1.5+1.5i)
func TestCheckCallExprFloat32FromComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(1.5+1.5i)`, env,
		`constant 1.5+1.5i truncated to real`,
	)

}

// Test float32(true)
func TestCheckCallExprFloat32FromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(true)`, env,
		`cannot convert true to type float32`,
		`cannot convert true (type bool) to type float32`,
	)

}

// Test float32("abc")
func TestCheckCallExprFloat32FromString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32("abc")`, env,
		`cannot convert "abc" to type float32`,
		`cannot convert "abc" (type string) to type float32`,
	)

}

// Test float32(nil)
func TestCheckCallExprFloat32FromNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float32(nil)`, env,
		`cannot convert nil to type float32`,
	)

}

// Test float64(0x7f)
func TestCheckCallExprFloat64From7bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0x7f)`, env, float64(0x7f), reflect.TypeOf(float64(0x7f)))
}

// Test float64(0xff)
func TestCheckCallExprFloat64From8bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xff)`, env, float64(0xff), reflect.TypeOf(float64(0xff)))
}

// Test float64(0x7fff)
func TestCheckCallExprFloat64From15bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0x7fff)`, env, float64(0x7fff), reflect.TypeOf(float64(0x7fff)))
}

// Test float64(0xffff)
func TestCheckCallExprFloat64From16bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffff)`, env, float64(0xffff), reflect.TypeOf(float64(0xffff)))
}

// Test float64(0x7fffffff)
func TestCheckCallExprFloat64From31bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0x7fffffff)`, env, float64(0x7fffffff), reflect.TypeOf(float64(0x7fffffff)))
}

// Test float64(0xffffffff)
func TestCheckCallExprFloat64From32bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffff)`, env, float64(0xffffffff), reflect.TypeOf(float64(0xffffffff)))
}

// Test float64(0x7fffffffffffffff)
func TestCheckCallExprFloat64From63bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0x7fffffffffffffff)`, env, float64(0x7fffffffffffffff), reflect.TypeOf(float64(0x7fffffffffffffff)))
}

// Test float64(0xffffffffffffffff)
func TestCheckCallExprFloat64From64bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(0xffffffffffffffff)`, env, float64(0xffffffffffffffff), reflect.TypeOf(float64(0xffffffffffffffff)))
}

// Test float64('d')
func TestCheckCallExprFloat64FromRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64('d')`, env, float64('d'), reflect.TypeOf(float64('d')))
}

// Test float64('日')
func TestCheckCallExprFloat64FromWideRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64('日')`, env, float64('日'), reflect.TypeOf(float64('日')))
}

// Test float64(1.0)
func TestCheckCallExprFloat64FromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(1.0)`, env, float64(1.0), reflect.TypeOf(float64(1.0)))
}

// Test float64(1.5)
func TestCheckCallExprFloat64FromFloating(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(1.5)`, env, float64(1.5), reflect.TypeOf(float64(1.5)))
}

// Test float64(1.0+0i)
func TestCheckCallExprFloat64FromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(1.0+0i)`, env, float64(1.0+0i), reflect.TypeOf(float64(1.0+0i)))
}

// Test float64(1.5+0i)
func TestCheckCallExprFloat64FromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `float64(1.5+0i)`, env, float64(1.5+0i), reflect.TypeOf(float64(1.5+0i)))
}

// Test float64(1.5+1.5i)
func TestCheckCallExprFloat64FromComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(1.5+1.5i)`, env,
		`constant 1.5+1.5i truncated to real`,
	)

}

// Test float64(true)
func TestCheckCallExprFloat64FromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(true)`, env,
		`cannot convert true to type float64`,
		`cannot convert true (type bool) to type float64`,
	)

}

// Test float64("abc")
func TestCheckCallExprFloat64FromString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64("abc")`, env,
		`cannot convert "abc" to type float64`,
		`cannot convert "abc" (type string) to type float64`,
	)

}

// Test float64(nil)
func TestCheckCallExprFloat64FromNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `float64(nil)`, env,
		`cannot convert nil to type float64`,
	)

}

// Test complex64(0x7f)
func TestCheckCallExprComplex64From7bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0x7f)`, env, complex64(0x7f), reflect.TypeOf(complex64(0x7f)))
}

// Test complex64(0xff)
func TestCheckCallExprComplex64From8bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xff)`, env, complex64(0xff), reflect.TypeOf(complex64(0xff)))
}

// Test complex64(0x7fff)
func TestCheckCallExprComplex64From15bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0x7fff)`, env, complex64(0x7fff), reflect.TypeOf(complex64(0x7fff)))
}

// Test complex64(0xffff)
func TestCheckCallExprComplex64From16bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffff)`, env, complex64(0xffff), reflect.TypeOf(complex64(0xffff)))
}

// Test complex64(0x7fffffff)
func TestCheckCallExprComplex64From31bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0x7fffffff)`, env, complex64(0x7fffffff), reflect.TypeOf(complex64(0x7fffffff)))
}

// Test complex64(0xffffffff)
func TestCheckCallExprComplex64From32bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffff)`, env, complex64(0xffffffff), reflect.TypeOf(complex64(0xffffffff)))
}

// Test complex64(0x7fffffffffffffff)
func TestCheckCallExprComplex64From63bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0x7fffffffffffffff)`, env, complex64(0x7fffffffffffffff), reflect.TypeOf(complex64(0x7fffffffffffffff)))
}

// Test complex64(0xffffffffffffffff)
func TestCheckCallExprComplex64From64bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(0xffffffffffffffff)`, env, complex64(0xffffffffffffffff), reflect.TypeOf(complex64(0xffffffffffffffff)))
}

// Test complex64('d')
func TestCheckCallExprComplex64FromRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64('d')`, env, complex64('d'), reflect.TypeOf(complex64('d')))
}

// Test complex64('日')
func TestCheckCallExprComplex64FromWideRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64('日')`, env, complex64('日'), reflect.TypeOf(complex64('日')))
}

// Test complex64(1.0)
func TestCheckCallExprComplex64FromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(1.0)`, env, complex64(1.0), reflect.TypeOf(complex64(1.0)))
}

// Test complex64(1.5)
func TestCheckCallExprComplex64FromFloating(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(1.5)`, env, complex64(1.5), reflect.TypeOf(complex64(1.5)))
}

// Test complex64(1.0+0i)
func TestCheckCallExprComplex64FromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(1.0+0i)`, env, complex64(1.0+0i), reflect.TypeOf(complex64(1.0+0i)))
}

// Test complex64(1.5+0i)
func TestCheckCallExprComplex64FromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(1.5+0i)`, env, complex64(1.5+0i), reflect.TypeOf(complex64(1.5+0i)))
}

// Test complex64(1.5+1.5i)
func TestCheckCallExprComplex64FromComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex64(1.5+1.5i)`, env, complex64(1.5+1.5i), reflect.TypeOf(complex64(1.5+1.5i)))
}

// Test complex64(true)
func TestCheckCallExprComplex64FromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(true)`, env,
		`cannot convert true to type complex64`,
		`cannot convert true (type bool) to type complex64`,
	)

}

// Test complex64("abc")
func TestCheckCallExprComplex64FromString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64("abc")`, env,
		`cannot convert "abc" to type complex64`,
		`cannot convert "abc" (type string) to type complex64`,
	)

}

// Test complex64(nil)
func TestCheckCallExprComplex64FromNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex64(nil)`, env,
		`cannot convert nil to type complex64`,
	)

}

// Test complex128(0x7f)
func TestCheckCallExprComplex128From7bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0x7f)`, env, complex128(0x7f), reflect.TypeOf(complex128(0x7f)))
}

// Test complex128(0xff)
func TestCheckCallExprComplex128From8bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xff)`, env, complex128(0xff), reflect.TypeOf(complex128(0xff)))
}

// Test complex128(0x7fff)
func TestCheckCallExprComplex128From15bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0x7fff)`, env, complex128(0x7fff), reflect.TypeOf(complex128(0x7fff)))
}

// Test complex128(0xffff)
func TestCheckCallExprComplex128From16bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffff)`, env, complex128(0xffff), reflect.TypeOf(complex128(0xffff)))
}

// Test complex128(0x7fffffff)
func TestCheckCallExprComplex128From31bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0x7fffffff)`, env, complex128(0x7fffffff), reflect.TypeOf(complex128(0x7fffffff)))
}

// Test complex128(0xffffffff)
func TestCheckCallExprComplex128From32bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffff)`, env, complex128(0xffffffff), reflect.TypeOf(complex128(0xffffffff)))
}

// Test complex128(0x7fffffffffffffff)
func TestCheckCallExprComplex128From63bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0x7fffffffffffffff)`, env, complex128(0x7fffffffffffffff), reflect.TypeOf(complex128(0x7fffffffffffffff)))
}

// Test complex128(0xffffffffffffffff)
func TestCheckCallExprComplex128From64bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(0xffffffffffffffff)`, env, complex128(0xffffffffffffffff), reflect.TypeOf(complex128(0xffffffffffffffff)))
}

// Test complex128('d')
func TestCheckCallExprComplex128FromRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128('d')`, env, complex128('d'), reflect.TypeOf(complex128('d')))
}

// Test complex128('日')
func TestCheckCallExprComplex128FromWideRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128('日')`, env, complex128('日'), reflect.TypeOf(complex128('日')))
}

// Test complex128(1.0)
func TestCheckCallExprComplex128FromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(1.0)`, env, complex128(1.0), reflect.TypeOf(complex128(1.0)))
}

// Test complex128(1.5)
func TestCheckCallExprComplex128FromFloating(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(1.5)`, env, complex128(1.5), reflect.TypeOf(complex128(1.5)))
}

// Test complex128(1.0+0i)
func TestCheckCallExprComplex128FromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(1.0+0i)`, env, complex128(1.0+0i), reflect.TypeOf(complex128(1.0+0i)))
}

// Test complex128(1.5+0i)
func TestCheckCallExprComplex128FromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(1.5+0i)`, env, complex128(1.5+0i), reflect.TypeOf(complex128(1.5+0i)))
}

// Test complex128(1.5+1.5i)
func TestCheckCallExprComplex128FromComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `complex128(1.5+1.5i)`, env, complex128(1.5+1.5i), reflect.TypeOf(complex128(1.5+1.5i)))
}

// Test complex128(true)
func TestCheckCallExprComplex128FromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(true)`, env,
		`cannot convert true to type complex128`,
		`cannot convert true (type bool) to type complex128`,
	)

}

// Test complex128("abc")
func TestCheckCallExprComplex128FromString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128("abc")`, env,
		`cannot convert "abc" to type complex128`,
		`cannot convert "abc" (type string) to type complex128`,
	)

}

// Test complex128(nil)
func TestCheckCallExprComplex128FromNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex128(nil)`, env,
		`cannot convert nil to type complex128`,
	)

}

// Test rune(0x7f)
func TestCheckCallExprRuneFrom7bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7f)`, env, rune(0x7f), reflect.TypeOf(rune(0x7f)))
}

// Test rune(0xff)
func TestCheckCallExprRuneFrom8bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0xff)`, env, rune(0xff), reflect.TypeOf(rune(0xff)))
}

// Test rune(0x7fff)
func TestCheckCallExprRuneFrom15bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fff)`, env, rune(0x7fff), reflect.TypeOf(rune(0x7fff)))
}

// Test rune(0xffff)
func TestCheckCallExprRuneFrom16bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0xffff)`, env, rune(0xffff), reflect.TypeOf(rune(0xffff)))
}

// Test rune(0x7fffffff)
func TestCheckCallExprRuneFrom31bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(0x7fffffff)`, env, rune(0x7fffffff), reflect.TypeOf(rune(0x7fffffff)))
}

// Test rune(0xffffffff)
func TestCheckCallExprRuneFrom32bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0xffffffff)`, env,
		`constant 4294967295 overflows rune`,
	)

}

// Test rune(0x7fffffffffffffff)
func TestCheckCallExprRuneFrom63bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0x7fffffffffffffff)`, env,
		`constant 9223372036854775807 overflows rune`,
	)

}

// Test rune(0xffffffffffffffff)
func TestCheckCallExprRuneFrom64bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(0xffffffffffffffff)`, env,
		`constant 18446744073709551615 overflows rune`,
	)

}

// Test rune('d')
func TestCheckCallExprRuneFromRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune('d')`, env, rune('d'), reflect.TypeOf(rune('d')))
}

// Test rune('日')
func TestCheckCallExprRuneFromWideRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune('日')`, env, rune('日'), reflect.TypeOf(rune('日')))
}

// Test rune(1.0)
func TestCheckCallExprRuneFromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(1.0)`, env, rune(1.0), reflect.TypeOf(rune(1.0)))
}

// Test rune(1.5)
func TestCheckCallExprRuneFromFloating(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test rune(1.0+0i)
func TestCheckCallExprRuneFromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `rune(1.0+0i)`, env, rune(1.0+0i), reflect.TypeOf(rune(1.0+0i)))
}

// Test rune(1.5+0i)
func TestCheckCallExprRuneFromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(1.5+0i)`, env,
		`constant 1.5+0i truncated to integer`,
	)

}

// Test rune(1.5+1.5i)
func TestCheckCallExprRuneFromComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(1.5+1.5i)`, env,
		`constant 1.5+1.5i truncated to integer`,
		`constant 1.5+1.5i truncated to real`,
	)

}

// Test rune(true)
func TestCheckCallExprRuneFromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(true)`, env,
		`cannot convert true to type rune`,
		`cannot convert true (type bool) to type rune`,
	)

}

// Test rune("abc")
func TestCheckCallExprRuneFromString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune("abc")`, env,
		`cannot convert "abc" to type rune`,
		`cannot convert "abc" (type string) to type rune`,
	)

}

// Test rune(nil)
func TestCheckCallExprRuneFromNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `rune(nil)`, env,
		`cannot convert nil to type rune`,
	)

}

// Test bool(0x7f)
func TestCheckCallExprBoolFrom7bits(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool(0xff)
func TestCheckCallExprBoolFrom8bits(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool(0x7fff)
func TestCheckCallExprBoolFrom15bits(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool(0xffff)
func TestCheckCallExprBoolFrom16bits(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool(0x7fffffff)
func TestCheckCallExprBoolFrom31bits(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool(0xffffffff)
func TestCheckCallExprBoolFrom32bits(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool(0x7fffffffffffffff)
func TestCheckCallExprBoolFrom63bits(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool(0xffffffffffffffff)
func TestCheckCallExprBoolFrom64bits(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool('d')
func TestCheckCallExprBoolFromRune(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool('日')
func TestCheckCallExprBoolFromWideRune(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool(1.0)
func TestCheckCallExprBoolFromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool(1.5)
func TestCheckCallExprBoolFromFloating(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool(1.0+0i)
func TestCheckCallExprBoolFromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool(1.5+0i)
func TestCheckCallExprBoolFromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool(1.5+1.5i)
func TestCheckCallExprBoolFromComplex(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool(true)
func TestCheckCallExprBoolFromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `bool(true)`, env, bool(true), reflect.TypeOf(bool(true)))
}

// Test bool("abc")
func TestCheckCallExprBoolFromString(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test bool(nil)
func TestCheckCallExprBoolFromNil(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test string(0x7f)
func TestCheckCallExprStringFrom7bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `string(0x7f)`, env, string(0x7f), reflect.TypeOf(string(0x7f)))
}

// Test string(0xff)
func TestCheckCallExprStringFrom8bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `string(0xff)`, env, string(0xff), reflect.TypeOf(string(0xff)))
}

// Test string(0x7fff)
func TestCheckCallExprStringFrom15bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `string(0x7fff)`, env, string(0x7fff), reflect.TypeOf(string(0x7fff)))
}

// Test string(0xffff)
func TestCheckCallExprStringFrom16bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `string(0xffff)`, env, string(0xffff), reflect.TypeOf(string(0xffff)))
}

// Test string(0x7fffffff)
func TestCheckCallExprStringFrom31bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `string(0x7fffffff)`, env, string(0x7fffffff), reflect.TypeOf(string(0x7fffffff)))
}

// Test string(0xffffffff)
func TestCheckCallExprStringFrom32bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string(0xffffffff)`, env,
		`overflow in int -> string`,
	)

}

// Test string(0x7fffffffffffffff)
func TestCheckCallExprStringFrom63bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string(0x7fffffffffffffff)`, env,
		`overflow in int -> string`,
	)

}

// Test string(0xffffffffffffffff)
func TestCheckCallExprStringFrom64bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string(0xffffffffffffffff)`, env,
		`overflow in int -> string`,
	)

}

// Test string('d')
func TestCheckCallExprStringFromRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `string('d')`, env, string('d'), reflect.TypeOf(string('d')))
}

// Test string('日')
func TestCheckCallExprStringFromWideRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `string('日')`, env, string('日'), reflect.TypeOf(string('日')))
}

// Test string(1.0)
func TestCheckCallExprStringFromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string(1.0)`, env,
		`cannot convert 1 to type string`,
		`cannot convert 1 (type float64) to type string`,
	)

}

// Test string(1.5)
func TestCheckCallExprStringFromFloating(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string(1.5)`, env,
		`cannot convert 1.5 to type string`,
		`cannot convert 1.5 (type float64) to type string`,
	)

}

// Test string(1.0+0i)
func TestCheckCallExprStringFromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string(1.0+0i)`, env,
		`cannot convert 1 + 0i to type string`,
		`cannot convert 1 + 0i (type complex128) to type string`,
	)

}

// Test string(1.5+0i)
func TestCheckCallExprStringFromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string(1.5+0i)`, env,
		`cannot convert 1.5 + 0i to type string`,
		`cannot convert 1.5 + 0i (type complex128) to type string`,
	)

}

// Test string(1.5+1.5i)
func TestCheckCallExprStringFromComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string(1.5+1.5i)`, env,
		`cannot convert 1.5 + 1.5i to type string`,
		`cannot convert 1.5 + 1.5i (type complex128) to type string`,
	)

}

// Test string(true)
func TestCheckCallExprStringFromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string(true)`, env,
		`cannot convert true to type string`,
		`cannot convert true (type bool) to type string`,
	)

}

// Test string("abc")
func TestCheckCallExprStringFromString(t *testing.T) {
	env := MakeSimpleEnv()

	expectConst(t, `string("abc")`, env, string("abc"), reflect.TypeOf(string("abc")))
}

// Test string(nil)
func TestCheckCallExprStringFromNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `string(nil)`, env,
		`cannot convert nil to type string`,
	)

}

// Test nil(0x7f)
func TestCheckCallExprNilFrom7bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil(0x7f)`, env,
		`use of untyped nil`,
	)

}

// Test nil(0xff)
func TestCheckCallExprNilFrom8bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil(0xff)`, env,
		`use of untyped nil`,
	)

}

// Test nil(0x7fff)
func TestCheckCallExprNilFrom15bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil(0x7fff)`, env,
		`use of untyped nil`,
	)

}

// Test nil(0xffff)
func TestCheckCallExprNilFrom16bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil(0xffff)`, env,
		`use of untyped nil`,
	)

}

// Test nil(0x7fffffff)
func TestCheckCallExprNilFrom31bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil(0x7fffffff)`, env,
		`use of untyped nil`,
	)

}

// Test nil(0xffffffff)
func TestCheckCallExprNilFrom32bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil(0xffffffff)`, env,
		`use of untyped nil`,
	)

}

// Test nil(0x7fffffffffffffff)
func TestCheckCallExprNilFrom63bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil(0x7fffffffffffffff)`, env,
		`use of untyped nil`,
	)

}

// Test nil(0xffffffffffffffff)
func TestCheckCallExprNilFrom64bits(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil(0xffffffffffffffff)`, env,
		`use of untyped nil`,
	)

}

// Test nil('d')
func TestCheckCallExprNilFromRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil('d')`, env,
		`use of untyped nil`,
	)

}

// Test nil('日')
func TestCheckCallExprNilFromWideRune(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil('日')`, env,
		`use of untyped nil`,
	)

}

// Test nil(1.0)
func TestCheckCallExprNilFromFloatingInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil(1.0)`, env,
		`use of untyped nil`,
	)

}

// Test nil(1.5)
func TestCheckCallExprNilFromFloating(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil(1.5)`, env,
		`use of untyped nil`,
	)

}

// Test nil(1.0+0i)
func TestCheckCallExprNilFromComplexInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil(1.0+0i)`, env,
		`use of untyped nil`,
	)

}

// Test nil(1.5+0i)
func TestCheckCallExprNilFromComplexFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil(1.5+0i)`, env,
		`use of untyped nil`,
	)

}

// Test nil(1.5+1.5i)
func TestCheckCallExprNilFromComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil(1.5+1.5i)`, env,
		`use of untyped nil`,
	)

}

// Test nil(true)
func TestCheckCallExprNilFromBool(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil(true)`, env,
		`use of untyped nil`,
	)

}

// Test nil("abc")
func TestCheckCallExprNilFromString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil("abc")`, env,
		`use of untyped nil`,
	)

}

// Test nil(nil)
func TestCheckCallExprNilFromNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `nil(nil)`, env,
		`use of untyped nil`,
	)

}
