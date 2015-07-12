package eval

import (
	"testing"
	"reflect"
)

// Test Complex()
func TestCheckBuiltinComplexXX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex()`, env,
		`missing argument to complex - complex(<N>, <N>)`,
	)

}

// Test Complex(1)
func TestCheckBuiltinComplexXInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1)`, env,
		`missing argument to complex - complex(1, <N>)`,
	)

}

// Test Complex(float32(1))
func TestCheckBuiltinComplexXFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(float32(1))`, env,
		`missing argument to complex - complex(float32(1), <N>)`,
	)

}

// Test Complex("abc")
func TestCheckBuiltinComplexXString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex("abc")`, env,
		`missing argument to complex - complex("abc", <N>)`,
	)

}

// Test Complex(nil)
func TestCheckBuiltinComplexXNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(nil)`, env,
		`missing argument to complex - complex(nil, <N>)`,
	)

}

// Test Complex(1.5)
func TestCheckBuiltinComplexXFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1.5)`, env,
		`missing argument to complex - complex(1.5, <N>)`,
	)

}

// Test Complex([]int{})
func TestCheckBuiltinComplexXSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex([]int{})`, env,
		`missing argument to complex - complex(composite literal, <N>)`,
	)

}

// Test Complex(map[int]int{})
func TestCheckBuiltinComplexXMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int{})`, env,
		`missing argument to complex - complex(composite literal, <N>)`,
	)

}

// Test Complex(int)
func TestCheckBuiltinComplexXType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(int)`, env,
		`missing argument to complex - complex(int, <N>)`,
	)

}

// Test Complex(map[int]int)
func TestCheckBuiltinComplexXMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int)`, env,
		`missing argument to complex - complex(map[int]int, <N>)`,
	)

}

// Test Complex(1, 1)
func TestCheckBuiltinComplexXDouble(t *testing.T) {
	env := MakeSimpleEnv()
	expectConst(t, `complex(1, 1)`, env, complex(1, 1), reflect.TypeOf(complex(1, 1)))

}

// Test Complex([]int{1,2}...)
func TestCheckBuiltinComplexXEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex([]int{1,2}...)`, env,
		`invalid use of ... with builtin complex`,
		`missing argument to complex - complex(composite literal, <N>)`,
	)

}

// Test Complex(1)
func TestCheckBuiltinComplexIntX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1)`, env,
		`missing argument to complex - complex(1, <N>)`,
	)

}

// Test Complex(1, 1)
func TestCheckBuiltinComplexIntInt(t *testing.T) {
	env := MakeSimpleEnv()
	expectConst(t, `complex(1, 1)`, env, complex(1, 1), reflect.TypeOf(complex(1, 1)))

}

// Test Complex(1, float32(1))
func TestCheckBuiltinComplexIntFloat32(t *testing.T) {
	env := MakeSimpleEnv()
	expectConst(t, `complex(1, float32(1))`, env, complex(1, float32(1)), reflect.TypeOf(complex(1, float32(1))))

}

// Test Complex(1, "abc")
func TestCheckBuiltinComplexIntString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1, "abc")`, env,
		`invalid operation: complex(1, "abc") (mismatched types untyped number and untyped string)`,
	)

}

// Test Complex(1, nil)
func TestCheckBuiltinComplexIntNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1, nil)`, env,
		`invalid operation: complex(1, nil) (mismatched types untyped number and nil)`,
	)

}

// Test Complex(1, 1.5)
func TestCheckBuiltinComplexIntFloat(t *testing.T) {
	env := MakeSimpleEnv()
	expectConst(t, `complex(1, 1.5)`, env, complex(1, 1.5), reflect.TypeOf(complex(1, 1.5)))

}

// Test Complex(1, []int{})
func TestCheckBuiltinComplexIntSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1, []int{})`, env,
		`invalid operation: complex(1, []int literal) (mismatched types int and []int)`,
	)

}

// Test Complex(1, map[int]int{})
func TestCheckBuiltinComplexIntMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1, map[int]int{})`, env,
		`invalid operation: complex(1, map[int]int literal) (mismatched types int and map[int]int)`,
	)

}

// Test Complex(1, int)
func TestCheckBuiltinComplexIntType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1, int)`, env,
		`type int is not an expression`,
	)

}

// Test Complex(1, map[int]int)
func TestCheckBuiltinComplexIntMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1, map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Complex(1, 1, 1)
func TestCheckBuiltinComplexIntDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1, 1, 1)`, env,
		`too many arguments to complex - complex(1, <N>)`,
	)

}

// Test Complex(1, []int{1,2}...)
func TestCheckBuiltinComplexIntEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1, []int{1,2}...)`, env,
		`invalid use of ... with builtin complex`,
		`invalid operation: complex(1, []int literal) (mismatched types int and []int)`,
	)

}

// Test Complex(float32(1))
func TestCheckBuiltinComplexFloat32X(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(float32(1))`, env,
		`missing argument to complex - complex(float32(1), <N>)`,
	)

}

// Test Complex(float32(1), 1)
func TestCheckBuiltinComplexFloat32Int(t *testing.T) {
	env := MakeSimpleEnv()
	expectConst(t, `complex(float32(1), 1)`, env, complex(float32(1), 1), reflect.TypeOf(complex(float32(1), 1)))

}

// Test Complex(float32(1), float32(1))
func TestCheckBuiltinComplexFloat32Float32(t *testing.T) {
	env := MakeSimpleEnv()
	expectConst(t, `complex(float32(1), float32(1))`, env, complex(float32(1), float32(1)), reflect.TypeOf(complex(float32(1), float32(1))))

}

// Test Complex(float32(1), "abc")
func TestCheckBuiltinComplexFloat32String(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(float32(1), "abc")`, env,
		`cannot convert "abc" to type float32`,
		`invalid operation: complex(float32(1), "abc") (mismatched types float32 and string)`,
	)

}

// Test Complex(float32(1), nil)
func TestCheckBuiltinComplexFloat32Nil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(float32(1), nil)`, env,
		`cannot convert nil to type float32`,
	)

}

// Test Complex(float32(1), 1.5)
func TestCheckBuiltinComplexFloat32Float(t *testing.T) {
	env := MakeSimpleEnv()
	expectConst(t, `complex(float32(1), 1.5)`, env, complex(float32(1), 1.5), reflect.TypeOf(complex(float32(1), 1.5)))

}

// Test Complex(float32(1), []int{})
func TestCheckBuiltinComplexFloat32Slice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(float32(1), []int{})`, env,
		`invalid operation: complex(float32(1), []int literal) (mismatched types float32 and []int)`,
	)

}

// Test Complex(float32(1), map[int]int{})
func TestCheckBuiltinComplexFloat32Map(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(float32(1), map[int]int{})`, env,
		`invalid operation: complex(float32(1), map[int]int literal) (mismatched types float32 and map[int]int)`,
	)

}

// Test Complex(float32(1), int)
func TestCheckBuiltinComplexFloat32Type(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(float32(1), int)`, env,
		`type int is not an expression`,
	)

}

// Test Complex(float32(1), map[int]int)
func TestCheckBuiltinComplexFloat32MakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(float32(1), map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Complex(float32(1), 1, 1)
func TestCheckBuiltinComplexFloat32Double(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(float32(1), 1, 1)`, env,
		`too many arguments to complex - complex(float32(1), <N>)`,
	)

}

// Test Complex(float32(1), []int{1,2}...)
func TestCheckBuiltinComplexFloat32Ellipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(float32(1), []int{1,2}...)`, env,
		`invalid use of ... with builtin complex`,
		`invalid operation: complex(float32(1), []int literal) (mismatched types float32 and []int)`,
	)

}

// Test Complex("abc")
func TestCheckBuiltinComplexStringX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex("abc")`, env,
		`missing argument to complex - complex("abc", <N>)`,
	)

}

// Test Complex("abc", 1)
func TestCheckBuiltinComplexStringInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex("abc", 1)`, env,
		`invalid operation: complex("abc", 1) (mismatched types untyped string and untyped number)`,
	)

}

// Test Complex("abc", float32(1))
func TestCheckBuiltinComplexStringFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex("abc", float32(1))`, env,
		`cannot convert "abc" to type float32`,
		`invalid operation: complex("abc", float32(1)) (mismatched types string and float32)`,
	)

}

// Test Complex("abc", "abc")
func TestCheckBuiltinComplexStringString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex("abc", "abc")`, env,
		`invalid operation: complex("abc", "abc") (arguments have type untyped string, expected floating-point)`,
	)

}

// Test Complex("abc", nil)
func TestCheckBuiltinComplexStringNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex("abc", nil)`, env,
		`invalid operation: complex("abc", nil) (mismatched types untyped string and nil)`,
	)

}

// Test Complex("abc", 1.5)
func TestCheckBuiltinComplexStringFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex("abc", 1.5)`, env,
		`invalid operation: complex("abc", 1.5) (mismatched types untyped string and untyped number)`,
	)

}

// Test Complex("abc", []int{})
func TestCheckBuiltinComplexStringSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex("abc", []int{})`, env,
		`invalid operation: complex("abc", []int literal) (mismatched types string and []int)`,
	)

}

// Test Complex("abc", map[int]int{})
func TestCheckBuiltinComplexStringMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex("abc", map[int]int{})`, env,
		`invalid operation: complex("abc", map[int]int literal) (mismatched types string and map[int]int)`,
	)

}

// Test Complex("abc", int)
func TestCheckBuiltinComplexStringType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex("abc", int)`, env,
		`type int is not an expression`,
	)

}

// Test Complex("abc", map[int]int)
func TestCheckBuiltinComplexStringMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex("abc", map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Complex("abc", 1, 1)
func TestCheckBuiltinComplexStringDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex("abc", 1, 1)`, env,
		`too many arguments to complex - complex("abc", <N>)`,
	)

}

// Test Complex("abc", []int{1,2}...)
func TestCheckBuiltinComplexStringEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex("abc", []int{1,2}...)`, env,
		`invalid use of ... with builtin complex`,
		`invalid operation: complex("abc", []int literal) (mismatched types string and []int)`,
	)

}

// Test Complex(nil)
func TestCheckBuiltinComplexNilX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(nil)`, env,
		`missing argument to complex - complex(nil, <N>)`,
	)

}

// Test Complex(nil, 1)
func TestCheckBuiltinComplexNilInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(nil, 1)`, env,
		`invalid operation: complex(nil, 1) (mismatched types nil and untyped number)`,
	)

}

// Test Complex(nil, float32(1))
func TestCheckBuiltinComplexNilFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(nil, float32(1))`, env,
		`cannot convert nil to type float32`,
	)

}

// Test Complex(nil, "abc")
func TestCheckBuiltinComplexNilString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(nil, "abc")`, env,
		`invalid operation: complex(nil, "abc") (mismatched types nil and untyped string)`,
	)

}

// Test Complex(nil, nil)
func TestCheckBuiltinComplexNilNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(nil, nil)`, env,
		`invalid operation: complex(nil, nil) (arguments have type nil, expected floating-point)`,
	)

}

// Test Complex(nil, 1.5)
func TestCheckBuiltinComplexNilFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(nil, 1.5)`, env,
		`invalid operation: complex(nil, 1.5) (mismatched types nil and untyped number)`,
	)

}

// Test Complex(nil, []int{})
func TestCheckBuiltinComplexNilSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(nil, []int{})`, env,
		`invalid operation: complex(nil, []int literal) (arguments have type []int, expected floating-point)`,
	)

}

// Test Complex(nil, map[int]int{})
func TestCheckBuiltinComplexNilMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(nil, map[int]int{})`, env,
		`invalid operation: complex(nil, map[int]int literal) (arguments have type map[int]int, expected floating-point)`,
	)

}

// Test Complex(nil, int)
func TestCheckBuiltinComplexNilType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(nil, int)`, env,
		`type int is not an expression`,
	)

}

// Test Complex(nil, map[int]int)
func TestCheckBuiltinComplexNilMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(nil, map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Complex(nil, 1, 1)
func TestCheckBuiltinComplexNilDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(nil, 1, 1)`, env,
		`too many arguments to complex - complex(nil, <N>)`,
	)

}

// Test Complex(nil, []int{1,2}...)
func TestCheckBuiltinComplexNilEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(nil, []int{1,2}...)`, env,
		`invalid use of ... with builtin complex`,
		`invalid operation: complex(nil, []int literal) (arguments have type []int, expected floating-point)`,
	)

}

// Test Complex(1.5)
func TestCheckBuiltinComplexFloatX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1.5)`, env,
		`missing argument to complex - complex(1.5, <N>)`,
	)

}

// Test Complex(1.5, 1)
func TestCheckBuiltinComplexFloatInt(t *testing.T) {
	env := MakeSimpleEnv()
	expectConst(t, `complex(1.5, 1)`, env, complex(1.5, 1), reflect.TypeOf(complex(1.5, 1)))

}

// Test Complex(1.5, float32(1))
func TestCheckBuiltinComplexFloatFloat32(t *testing.T) {
	env := MakeSimpleEnv()
	expectConst(t, `complex(1.5, float32(1))`, env, complex(1.5, float32(1)), reflect.TypeOf(complex(1.5, float32(1))))

}

// Test Complex(1.5, "abc")
func TestCheckBuiltinComplexFloatString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1.5, "abc")`, env,
		`invalid operation: complex(1.5, "abc") (mismatched types untyped number and untyped string)`,
	)

}

// Test Complex(1.5, nil)
func TestCheckBuiltinComplexFloatNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1.5, nil)`, env,
		`invalid operation: complex(1.5, nil) (mismatched types untyped number and nil)`,
	)

}

// Test Complex(1.5, 1.5)
func TestCheckBuiltinComplexFloatFloat(t *testing.T) {
	env := MakeSimpleEnv()
	expectConst(t, `complex(1.5, 1.5)`, env, complex(1.5, 1.5), reflect.TypeOf(complex(1.5, 1.5)))

}

// Test Complex(1.5, []int{})
func TestCheckBuiltinComplexFloatSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1.5, []int{})`, env,
		`invalid operation: complex(1.5, []int literal) (mismatched types float64 and []int)`,
	)

}

// Test Complex(1.5, map[int]int{})
func TestCheckBuiltinComplexFloatMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1.5, map[int]int{})`, env,
		`invalid operation: complex(1.5, map[int]int literal) (mismatched types float64 and map[int]int)`,
	)

}

// Test Complex(1.5, int)
func TestCheckBuiltinComplexFloatType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1.5, int)`, env,
		`type int is not an expression`,
	)

}

// Test Complex(1.5, map[int]int)
func TestCheckBuiltinComplexFloatMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1.5, map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Complex(1.5, 1, 1)
func TestCheckBuiltinComplexFloatDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1.5, 1, 1)`, env,
		`too many arguments to complex - complex(1.5, <N>)`,
	)

}

// Test Complex(1.5, []int{1,2}...)
func TestCheckBuiltinComplexFloatEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(1.5, []int{1,2}...)`, env,
		`invalid use of ... with builtin complex`,
		`invalid operation: complex(1.5, []int literal) (mismatched types float64 and []int)`,
	)

}

// Test Complex([]int{})
func TestCheckBuiltinComplexSliceX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex([]int{})`, env,
		`missing argument to complex - complex(composite literal, <N>)`,
	)

}

// Test Complex([]int{}, 1)
func TestCheckBuiltinComplexSliceInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex([]int{}, 1)`, env,
		`invalid operation: complex([]int literal, 1) (mismatched types []int and int)`,
	)

}

// Test Complex([]int{}, float32(1))
func TestCheckBuiltinComplexSliceFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex([]int{}, float32(1))`, env,
		`invalid operation: complex([]int literal, float32(1)) (mismatched types []int and float32)`,
	)

}

// Test Complex([]int{}, "abc")
func TestCheckBuiltinComplexSliceString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex([]int{}, "abc")`, env,
		`invalid operation: complex([]int literal, "abc") (mismatched types []int and string)`,
	)

}

// Test Complex([]int{}, nil)
func TestCheckBuiltinComplexSliceNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex([]int{}, nil)`, env,
		`invalid operation: complex([]int literal, nil) (arguments have type []int, expected floating-point)`,
	)

}

// Test Complex([]int{}, 1.5)
func TestCheckBuiltinComplexSliceFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex([]int{}, 1.5)`, env,
		`invalid operation: complex([]int literal, 1.5) (mismatched types []int and float64)`,
	)

}

// Test Complex([]int{}, []int{})
func TestCheckBuiltinComplexSliceSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex([]int{}, []int{})`, env,
		`invalid operation: complex([]int literal, []int literal) (arguments have type []int, expected floating-point)`,
	)

}

// Test Complex([]int{}, map[int]int{})
func TestCheckBuiltinComplexSliceMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex([]int{}, map[int]int{})`, env,
		`invalid operation: complex([]int literal, map[int]int literal) (mismatched types []int and map[int]int)`,
	)

}

// Test Complex([]int{}, int)
func TestCheckBuiltinComplexSliceType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex([]int{}, int)`, env,
		`type int is not an expression`,
	)

}

// Test Complex([]int{}, map[int]int)
func TestCheckBuiltinComplexSliceMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex([]int{}, map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Complex([]int{}, 1, 1)
func TestCheckBuiltinComplexSliceDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex([]int{}, 1, 1)`, env,
		`too many arguments to complex - complex(composite literal, <N>)`,
	)

}

// Test Complex([]int{}, []int{1,2}...)
func TestCheckBuiltinComplexSliceEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex([]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin complex`,
		`invalid operation: complex([]int literal, []int literal) (arguments have type []int, expected floating-point)`,
	)

}

// Test Complex(map[int]int{})
func TestCheckBuiltinComplexMapX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int{})`, env,
		`missing argument to complex - complex(composite literal, <N>)`,
	)

}

// Test Complex(map[int]int{}, 1)
func TestCheckBuiltinComplexMapInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int{}, 1)`, env,
		`invalid operation: complex(map[int]int literal, 1) (mismatched types map[int]int and int)`,
	)

}

// Test Complex(map[int]int{}, float32(1))
func TestCheckBuiltinComplexMapFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int{}, float32(1))`, env,
		`invalid operation: complex(map[int]int literal, float32(1)) (mismatched types map[int]int and float32)`,
	)

}

// Test Complex(map[int]int{}, "abc")
func TestCheckBuiltinComplexMapString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int{}, "abc")`, env,
		`invalid operation: complex(map[int]int literal, "abc") (mismatched types map[int]int and string)`,
	)

}

// Test Complex(map[int]int{}, nil)
func TestCheckBuiltinComplexMapNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int{}, nil)`, env,
		`invalid operation: complex(map[int]int literal, nil) (arguments have type map[int]int, expected floating-point)`,
	)

}

// Test Complex(map[int]int{}, 1.5)
func TestCheckBuiltinComplexMapFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int{}, 1.5)`, env,
		`invalid operation: complex(map[int]int literal, 1.5) (mismatched types map[int]int and float64)`,
	)

}

// Test Complex(map[int]int{}, []int{})
func TestCheckBuiltinComplexMapSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int{}, []int{})`, env,
		`invalid operation: complex(map[int]int literal, []int literal) (mismatched types map[int]int and []int)`,
	)

}

// Test Complex(map[int]int{}, map[int]int{})
func TestCheckBuiltinComplexMapMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int{}, map[int]int{})`, env,
		`invalid operation: complex(map[int]int literal, map[int]int literal) (arguments have type map[int]int, expected floating-point)`,
	)

}

// Test Complex(map[int]int{}, int)
func TestCheckBuiltinComplexMapType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int{}, int)`, env,
		`type int is not an expression`,
	)

}

// Test Complex(map[int]int{}, map[int]int)
func TestCheckBuiltinComplexMapMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int{}, map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Complex(map[int]int{}, 1, 1)
func TestCheckBuiltinComplexMapDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int{}, 1, 1)`, env,
		`too many arguments to complex - complex(composite literal, <N>)`,
	)

}

// Test Complex(map[int]int{}, []int{1,2}...)
func TestCheckBuiltinComplexMapEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin complex`,
		`invalid operation: complex(map[int]int literal, []int literal) (mismatched types map[int]int and []int)`,
	)

}

// Test Complex(int)
func TestCheckBuiltinComplexTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(int)`, env,
		`missing argument to complex - complex(int, <N>)`,
	)

}

// Test Complex(int, 1)
func TestCheckBuiltinComplexTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(int, 1)`, env,
		`type int is not an expression`,
	)

}

// Test Complex(int, float32(1))
func TestCheckBuiltinComplexTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(int, float32(1))`, env,
		`type int is not an expression`,
	)

}

// Test Complex(int, "abc")
func TestCheckBuiltinComplexTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(int, "abc")`, env,
		`type int is not an expression`,
	)

}

// Test Complex(int, nil)
func TestCheckBuiltinComplexTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(int, nil)`, env,
		`type int is not an expression`,
	)

}

// Test Complex(int, 1.5)
func TestCheckBuiltinComplexTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(int, 1.5)`, env,
		`type int is not an expression`,
	)

}

// Test Complex(int, []int{})
func TestCheckBuiltinComplexTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(int, []int{})`, env,
		`type int is not an expression`,
	)

}

// Test Complex(int, map[int]int{})
func TestCheckBuiltinComplexTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(int, map[int]int{})`, env,
		`type int is not an expression`,
	)

}

// Test Complex(int, int)
func TestCheckBuiltinComplexTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(int, int)`, env,
		`type int is not an expression`,
		`type int is not an expression`,
	)

}

// Test Complex(int, map[int]int)
func TestCheckBuiltinComplexTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(int, map[int]int)`, env,
		`type int is not an expression`,
		`type map[int]int is not an expression`,
	)

}

// Test Complex(int, 1, 1)
func TestCheckBuiltinComplexTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(int, 1, 1)`, env,
		`too many arguments to complex - complex(int, <N>)`,
	)

}

// Test Complex(int, []int{1,2}...)
func TestCheckBuiltinComplexTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(int, []int{1,2}...)`, env,
		`invalid use of ... with builtin complex`,
		`type int is not an expression`,
	)

}

// Test Complex(map[int]int)
func TestCheckBuiltinComplexMakeTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int)`, env,
		`missing argument to complex - complex(map[int]int, <N>)`,
	)

}

// Test Complex(map[int]int, 1)
func TestCheckBuiltinComplexMakeTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int, 1)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Complex(map[int]int, float32(1))
func TestCheckBuiltinComplexMakeTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int, float32(1))`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Complex(map[int]int, "abc")
func TestCheckBuiltinComplexMakeTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int, "abc")`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Complex(map[int]int, nil)
func TestCheckBuiltinComplexMakeTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int, nil)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Complex(map[int]int, 1.5)
func TestCheckBuiltinComplexMakeTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int, 1.5)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Complex(map[int]int, []int{})
func TestCheckBuiltinComplexMakeTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int, []int{})`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Complex(map[int]int, map[int]int{})
func TestCheckBuiltinComplexMakeTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int, map[int]int{})`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Complex(map[int]int, int)
func TestCheckBuiltinComplexMakeTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int, int)`, env,
		`type map[int]int is not an expression`,
		`type int is not an expression`,
	)

}

// Test Complex(map[int]int, map[int]int)
func TestCheckBuiltinComplexMakeTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int, map[int]int)`, env,
		`type map[int]int is not an expression`,
		`type map[int]int is not an expression`,
	)

}

// Test Complex(map[int]int, 1, 1)
func TestCheckBuiltinComplexMakeTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int, 1, 1)`, env,
		`too many arguments to complex - complex(map[int]int, <N>)`,
	)

}

// Test Complex(map[int]int, []int{1,2}...)
func TestCheckBuiltinComplexMakeTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `complex(map[int]int, []int{1,2}...)`, env,
		`invalid use of ... with builtin complex`,
		`type map[int]int is not an expression`,
	)

}

// Test Real()
func TestCheckBuiltinRealXX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real()`, env,
		`missing argument to real: real()`,
	)

}

// Test Real(1)
func TestCheckBuiltinRealXInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1)`, env,
		`invalid argument 1 (type int) for real`,
	)

}

// Test Real(float32(1))
func TestCheckBuiltinRealXFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(float32(1))`, env,
		`invalid argument float32(1) (type float32) for real`,
	)

}

// Test Real("abc")
func TestCheckBuiltinRealXString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real("abc")`, env,
		`invalid argument "abc" (type string) for real`,
	)

}

// Test Real(nil)
func TestCheckBuiltinRealXNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(nil)`, env,
		`use of untyped nil`,
	)

}

// Test Real(1.5)
func TestCheckBuiltinRealXFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1.5)`, env,
		`invalid argument 1.5 (type float64) for real`,
	)

}

// Test Real([]int{})
func TestCheckBuiltinRealXSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real([]int{})`, env,
		`invalid argument []int literal (type []int) for real`,
	)

}

// Test Real(map[int]int{})
func TestCheckBuiltinRealXMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int{})`, env,
		`invalid argument map[int]int literal (type map[int]int) for real`,
	)

}

// Test Real(int)
func TestCheckBuiltinRealXType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(int)`, env,
		`type int is not an expression`,
	)

}

// Test Real(map[int]int)
func TestCheckBuiltinRealXMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Real(1, 1)
func TestCheckBuiltinRealXDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1, 1)`, env,
		`too many arguments to real: real(1, 1)`,
	)

}

// Test Real([]int{1,2}...)
func TestCheckBuiltinRealXEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real([]int{1,2}...)`, env,
		`invalid use of ... with builtin real`,
		`invalid argument []int literal (type []int) for real`,
	)

}

// Test Real(1)
func TestCheckBuiltinRealIntX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1)`, env,
		`invalid argument 1 (type int) for real`,
	)

}

// Test Real(1, 1)
func TestCheckBuiltinRealIntInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1, 1)`, env,
		`too many arguments to real: real(1, 1)`,
	)

}

// Test Real(1, float32(1))
func TestCheckBuiltinRealIntFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1, float32(1))`, env,
		`too many arguments to real: real(1, float32(1))`,
	)

}

// Test Real(1, "abc")
func TestCheckBuiltinRealIntString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1, "abc")`, env,
		`too many arguments to real: real(1, "abc")`,
	)

}

// Test Real(1, nil)
func TestCheckBuiltinRealIntNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1, nil)`, env,
		`too many arguments to real: real(1, nil)`,
	)

}

// Test Real(1, 1.5)
func TestCheckBuiltinRealIntFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1, 1.5)`, env,
		`too many arguments to real: real(1, 1.5)`,
	)

}

// Test Real(1, []int{})
func TestCheckBuiltinRealIntSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1, []int{})`, env,
		`too many arguments to real: real(1, composite literal)`,
	)

}

// Test Real(1, map[int]int{})
func TestCheckBuiltinRealIntMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1, map[int]int{})`, env,
		`too many arguments to real: real(1, composite literal)`,
	)

}

// Test Real(1, int)
func TestCheckBuiltinRealIntType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1, int)`, env,
		`too many arguments to real: real(1, int)`,
	)

}

// Test Real(1, map[int]int)
func TestCheckBuiltinRealIntMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1, map[int]int)`, env,
		`too many arguments to real: real(1, map[int]int)`,
	)

}

// Test Real(1, 1, 1)
func TestCheckBuiltinRealIntDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1, 1, 1)`, env,
		`too many arguments to real: real(1, 1, 1)`,
	)

}

// Test Real(1, []int{1,2}...)
func TestCheckBuiltinRealIntEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1, []int{1,2}...)`, env,
		`invalid use of ... with builtin real`,
		`too many arguments to real: real(1, composite literal...)`,
	)

}

// Test Real(float32(1))
func TestCheckBuiltinRealFloat32X(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(float32(1))`, env,
		`invalid argument float32(1) (type float32) for real`,
	)

}

// Test Real(float32(1), 1)
func TestCheckBuiltinRealFloat32Int(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(float32(1), 1)`, env,
		`too many arguments to real: real(float32(1), 1)`,
	)

}

// Test Real(float32(1), float32(1))
func TestCheckBuiltinRealFloat32Float32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(float32(1), float32(1))`, env,
		`too many arguments to real: real(float32(1), float32(1))`,
	)

}

// Test Real(float32(1), "abc")
func TestCheckBuiltinRealFloat32String(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(float32(1), "abc")`, env,
		`too many arguments to real: real(float32(1), "abc")`,
	)

}

// Test Real(float32(1), nil)
func TestCheckBuiltinRealFloat32Nil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(float32(1), nil)`, env,
		`too many arguments to real: real(float32(1), nil)`,
	)

}

// Test Real(float32(1), 1.5)
func TestCheckBuiltinRealFloat32Float(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(float32(1), 1.5)`, env,
		`too many arguments to real: real(float32(1), 1.5)`,
	)

}

// Test Real(float32(1), []int{})
func TestCheckBuiltinRealFloat32Slice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(float32(1), []int{})`, env,
		`too many arguments to real: real(float32(1), composite literal)`,
	)

}

// Test Real(float32(1), map[int]int{})
func TestCheckBuiltinRealFloat32Map(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(float32(1), map[int]int{})`, env,
		`too many arguments to real: real(float32(1), composite literal)`,
	)

}

// Test Real(float32(1), int)
func TestCheckBuiltinRealFloat32Type(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(float32(1), int)`, env,
		`too many arguments to real: real(float32(1), int)`,
	)

}

// Test Real(float32(1), map[int]int)
func TestCheckBuiltinRealFloat32MakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(float32(1), map[int]int)`, env,
		`too many arguments to real: real(float32(1), map[int]int)`,
	)

}

// Test Real(float32(1), 1, 1)
func TestCheckBuiltinRealFloat32Double(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(float32(1), 1, 1)`, env,
		`too many arguments to real: real(float32(1), 1, 1)`,
	)

}

// Test Real(float32(1), []int{1,2}...)
func TestCheckBuiltinRealFloat32Ellipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(float32(1), []int{1,2}...)`, env,
		`invalid use of ... with builtin real`,
		`too many arguments to real: real(float32(1), composite literal...)`,
	)

}

// Test Real("abc")
func TestCheckBuiltinRealStringX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real("abc")`, env,
		`invalid argument "abc" (type string) for real`,
	)

}

// Test Real("abc", 1)
func TestCheckBuiltinRealStringInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real("abc", 1)`, env,
		`too many arguments to real: real("abc", 1)`,
	)

}

// Test Real("abc", float32(1))
func TestCheckBuiltinRealStringFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real("abc", float32(1))`, env,
		`too many arguments to real: real("abc", float32(1))`,
	)

}

// Test Real("abc", "abc")
func TestCheckBuiltinRealStringString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real("abc", "abc")`, env,
		`too many arguments to real: real("abc", "abc")`,
	)

}

// Test Real("abc", nil)
func TestCheckBuiltinRealStringNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real("abc", nil)`, env,
		`too many arguments to real: real("abc", nil)`,
	)

}

// Test Real("abc", 1.5)
func TestCheckBuiltinRealStringFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real("abc", 1.5)`, env,
		`too many arguments to real: real("abc", 1.5)`,
	)

}

// Test Real("abc", []int{})
func TestCheckBuiltinRealStringSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real("abc", []int{})`, env,
		`too many arguments to real: real("abc", composite literal)`,
	)

}

// Test Real("abc", map[int]int{})
func TestCheckBuiltinRealStringMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real("abc", map[int]int{})`, env,
		`too many arguments to real: real("abc", composite literal)`,
	)

}

// Test Real("abc", int)
func TestCheckBuiltinRealStringType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real("abc", int)`, env,
		`too many arguments to real: real("abc", int)`,
	)

}

// Test Real("abc", map[int]int)
func TestCheckBuiltinRealStringMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real("abc", map[int]int)`, env,
		`too many arguments to real: real("abc", map[int]int)`,
	)

}

// Test Real("abc", 1, 1)
func TestCheckBuiltinRealStringDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real("abc", 1, 1)`, env,
		`too many arguments to real: real("abc", 1, 1)`,
	)

}

// Test Real("abc", []int{1,2}...)
func TestCheckBuiltinRealStringEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real("abc", []int{1,2}...)`, env,
		`invalid use of ... with builtin real`,
		`too many arguments to real: real("abc", composite literal...)`,
	)

}

// Test Real(nil)
func TestCheckBuiltinRealNilX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(nil)`, env,
		`use of untyped nil`,
	)

}

// Test Real(nil, 1)
func TestCheckBuiltinRealNilInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(nil, 1)`, env,
		`too many arguments to real: real(nil, 1)`,
	)

}

// Test Real(nil, float32(1))
func TestCheckBuiltinRealNilFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(nil, float32(1))`, env,
		`too many arguments to real: real(nil, float32(1))`,
	)

}

// Test Real(nil, "abc")
func TestCheckBuiltinRealNilString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(nil, "abc")`, env,
		`too many arguments to real: real(nil, "abc")`,
	)

}

// Test Real(nil, nil)
func TestCheckBuiltinRealNilNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(nil, nil)`, env,
		`too many arguments to real: real(nil, nil)`,
	)

}

// Test Real(nil, 1.5)
func TestCheckBuiltinRealNilFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(nil, 1.5)`, env,
		`too many arguments to real: real(nil, 1.5)`,
	)

}

// Test Real(nil, []int{})
func TestCheckBuiltinRealNilSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(nil, []int{})`, env,
		`too many arguments to real: real(nil, composite literal)`,
	)

}

// Test Real(nil, map[int]int{})
func TestCheckBuiltinRealNilMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(nil, map[int]int{})`, env,
		`too many arguments to real: real(nil, composite literal)`,
	)

}

// Test Real(nil, int)
func TestCheckBuiltinRealNilType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(nil, int)`, env,
		`too many arguments to real: real(nil, int)`,
	)

}

// Test Real(nil, map[int]int)
func TestCheckBuiltinRealNilMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(nil, map[int]int)`, env,
		`too many arguments to real: real(nil, map[int]int)`,
	)

}

// Test Real(nil, 1, 1)
func TestCheckBuiltinRealNilDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(nil, 1, 1)`, env,
		`too many arguments to real: real(nil, 1, 1)`,
	)

}

// Test Real(nil, []int{1,2}...)
func TestCheckBuiltinRealNilEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(nil, []int{1,2}...)`, env,
		`invalid use of ... with builtin real`,
		`too many arguments to real: real(nil, composite literal...)`,
	)

}

// Test Real(1.5)
func TestCheckBuiltinRealFloatX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1.5)`, env,
		`invalid argument 1.5 (type float64) for real`,
	)

}

// Test Real(1.5, 1)
func TestCheckBuiltinRealFloatInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1.5, 1)`, env,
		`too many arguments to real: real(1.5, 1)`,
	)

}

// Test Real(1.5, float32(1))
func TestCheckBuiltinRealFloatFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1.5, float32(1))`, env,
		`too many arguments to real: real(1.5, float32(1))`,
	)

}

// Test Real(1.5, "abc")
func TestCheckBuiltinRealFloatString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1.5, "abc")`, env,
		`too many arguments to real: real(1.5, "abc")`,
	)

}

// Test Real(1.5, nil)
func TestCheckBuiltinRealFloatNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1.5, nil)`, env,
		`too many arguments to real: real(1.5, nil)`,
	)

}

// Test Real(1.5, 1.5)
func TestCheckBuiltinRealFloatFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1.5, 1.5)`, env,
		`too many arguments to real: real(1.5, 1.5)`,
	)

}

// Test Real(1.5, []int{})
func TestCheckBuiltinRealFloatSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1.5, []int{})`, env,
		`too many arguments to real: real(1.5, composite literal)`,
	)

}

// Test Real(1.5, map[int]int{})
func TestCheckBuiltinRealFloatMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1.5, map[int]int{})`, env,
		`too many arguments to real: real(1.5, composite literal)`,
	)

}

// Test Real(1.5, int)
func TestCheckBuiltinRealFloatType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1.5, int)`, env,
		`too many arguments to real: real(1.5, int)`,
	)

}

// Test Real(1.5, map[int]int)
func TestCheckBuiltinRealFloatMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1.5, map[int]int)`, env,
		`too many arguments to real: real(1.5, map[int]int)`,
	)

}

// Test Real(1.5, 1, 1)
func TestCheckBuiltinRealFloatDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1.5, 1, 1)`, env,
		`too many arguments to real: real(1.5, 1, 1)`,
	)

}

// Test Real(1.5, []int{1,2}...)
func TestCheckBuiltinRealFloatEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(1.5, []int{1,2}...)`, env,
		`invalid use of ... with builtin real`,
		`too many arguments to real: real(1.5, composite literal...)`,
	)

}

// Test Real([]int{})
func TestCheckBuiltinRealSliceX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real([]int{})`, env,
		`invalid argument []int literal (type []int) for real`,
	)

}

// Test Real([]int{}, 1)
func TestCheckBuiltinRealSliceInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real([]int{}, 1)`, env,
		`too many arguments to real: real(composite literal, 1)`,
	)

}

// Test Real([]int{}, float32(1))
func TestCheckBuiltinRealSliceFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real([]int{}, float32(1))`, env,
		`too many arguments to real: real(composite literal, float32(1))`,
	)

}

// Test Real([]int{}, "abc")
func TestCheckBuiltinRealSliceString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real([]int{}, "abc")`, env,
		`too many arguments to real: real(composite literal, "abc")`,
	)

}

// Test Real([]int{}, nil)
func TestCheckBuiltinRealSliceNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real([]int{}, nil)`, env,
		`too many arguments to real: real(composite literal, nil)`,
	)

}

// Test Real([]int{}, 1.5)
func TestCheckBuiltinRealSliceFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real([]int{}, 1.5)`, env,
		`too many arguments to real: real(composite literal, 1.5)`,
	)

}

// Test Real([]int{}, []int{})
func TestCheckBuiltinRealSliceSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real([]int{}, []int{})`, env,
		`too many arguments to real: real(composite literal, composite literal)`,
	)

}

// Test Real([]int{}, map[int]int{})
func TestCheckBuiltinRealSliceMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real([]int{}, map[int]int{})`, env,
		`too many arguments to real: real(composite literal, composite literal)`,
	)

}

// Test Real([]int{}, int)
func TestCheckBuiltinRealSliceType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real([]int{}, int)`, env,
		`too many arguments to real: real(composite literal, int)`,
	)

}

// Test Real([]int{}, map[int]int)
func TestCheckBuiltinRealSliceMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real([]int{}, map[int]int)`, env,
		`too many arguments to real: real(composite literal, map[int]int)`,
	)

}

// Test Real([]int{}, 1, 1)
func TestCheckBuiltinRealSliceDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real([]int{}, 1, 1)`, env,
		`too many arguments to real: real(composite literal, 1, 1)`,
	)

}

// Test Real([]int{}, []int{1,2}...)
func TestCheckBuiltinRealSliceEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real([]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin real`,
		`too many arguments to real: real(composite literal, composite literal...)`,
	)

}

// Test Real(map[int]int{})
func TestCheckBuiltinRealMapX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int{})`, env,
		`invalid argument map[int]int literal (type map[int]int) for real`,
	)

}

// Test Real(map[int]int{}, 1)
func TestCheckBuiltinRealMapInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int{}, 1)`, env,
		`too many arguments to real: real(composite literal, 1)`,
	)

}

// Test Real(map[int]int{}, float32(1))
func TestCheckBuiltinRealMapFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int{}, float32(1))`, env,
		`too many arguments to real: real(composite literal, float32(1))`,
	)

}

// Test Real(map[int]int{}, "abc")
func TestCheckBuiltinRealMapString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int{}, "abc")`, env,
		`too many arguments to real: real(composite literal, "abc")`,
	)

}

// Test Real(map[int]int{}, nil)
func TestCheckBuiltinRealMapNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int{}, nil)`, env,
		`too many arguments to real: real(composite literal, nil)`,
	)

}

// Test Real(map[int]int{}, 1.5)
func TestCheckBuiltinRealMapFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int{}, 1.5)`, env,
		`too many arguments to real: real(composite literal, 1.5)`,
	)

}

// Test Real(map[int]int{}, []int{})
func TestCheckBuiltinRealMapSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int{}, []int{})`, env,
		`too many arguments to real: real(composite literal, composite literal)`,
	)

}

// Test Real(map[int]int{}, map[int]int{})
func TestCheckBuiltinRealMapMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int{}, map[int]int{})`, env,
		`too many arguments to real: real(composite literal, composite literal)`,
	)

}

// Test Real(map[int]int{}, int)
func TestCheckBuiltinRealMapType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int{}, int)`, env,
		`too many arguments to real: real(composite literal, int)`,
	)

}

// Test Real(map[int]int{}, map[int]int)
func TestCheckBuiltinRealMapMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int{}, map[int]int)`, env,
		`too many arguments to real: real(composite literal, map[int]int)`,
	)

}

// Test Real(map[int]int{}, 1, 1)
func TestCheckBuiltinRealMapDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int{}, 1, 1)`, env,
		`too many arguments to real: real(composite literal, 1, 1)`,
	)

}

// Test Real(map[int]int{}, []int{1,2}...)
func TestCheckBuiltinRealMapEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin real`,
		`too many arguments to real: real(composite literal, composite literal...)`,
	)

}

// Test Real(int)
func TestCheckBuiltinRealTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(int)`, env,
		`type int is not an expression`,
	)

}

// Test Real(int, 1)
func TestCheckBuiltinRealTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(int, 1)`, env,
		`too many arguments to real: real(int, 1)`,
	)

}

// Test Real(int, float32(1))
func TestCheckBuiltinRealTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(int, float32(1))`, env,
		`too many arguments to real: real(int, float32(1))`,
	)

}

// Test Real(int, "abc")
func TestCheckBuiltinRealTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(int, "abc")`, env,
		`too many arguments to real: real(int, "abc")`,
	)

}

// Test Real(int, nil)
func TestCheckBuiltinRealTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(int, nil)`, env,
		`too many arguments to real: real(int, nil)`,
	)

}

// Test Real(int, 1.5)
func TestCheckBuiltinRealTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(int, 1.5)`, env,
		`too many arguments to real: real(int, 1.5)`,
	)

}

// Test Real(int, []int{})
func TestCheckBuiltinRealTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(int, []int{})`, env,
		`too many arguments to real: real(int, composite literal)`,
	)

}

// Test Real(int, map[int]int{})
func TestCheckBuiltinRealTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(int, map[int]int{})`, env,
		`too many arguments to real: real(int, composite literal)`,
	)

}

// Test Real(int, int)
func TestCheckBuiltinRealTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(int, int)`, env,
		`too many arguments to real: real(int, int)`,
	)

}

// Test Real(int, map[int]int)
func TestCheckBuiltinRealTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(int, map[int]int)`, env,
		`too many arguments to real: real(int, map[int]int)`,
	)

}

// Test Real(int, 1, 1)
func TestCheckBuiltinRealTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(int, 1, 1)`, env,
		`too many arguments to real: real(int, 1, 1)`,
	)

}

// Test Real(int, []int{1,2}...)
func TestCheckBuiltinRealTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(int, []int{1,2}...)`, env,
		`invalid use of ... with builtin real`,
		`too many arguments to real: real(int, composite literal...)`,
	)

}

// Test Real(map[int]int)
func TestCheckBuiltinRealMakeTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Real(map[int]int, 1)
func TestCheckBuiltinRealMakeTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int, 1)`, env,
		`too many arguments to real: real(map[int]int, 1)`,
	)

}

// Test Real(map[int]int, float32(1))
func TestCheckBuiltinRealMakeTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int, float32(1))`, env,
		`too many arguments to real: real(map[int]int, float32(1))`,
	)

}

// Test Real(map[int]int, "abc")
func TestCheckBuiltinRealMakeTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int, "abc")`, env,
		`too many arguments to real: real(map[int]int, "abc")`,
	)

}

// Test Real(map[int]int, nil)
func TestCheckBuiltinRealMakeTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int, nil)`, env,
		`too many arguments to real: real(map[int]int, nil)`,
	)

}

// Test Real(map[int]int, 1.5)
func TestCheckBuiltinRealMakeTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int, 1.5)`, env,
		`too many arguments to real: real(map[int]int, 1.5)`,
	)

}

// Test Real(map[int]int, []int{})
func TestCheckBuiltinRealMakeTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int, []int{})`, env,
		`too many arguments to real: real(map[int]int, composite literal)`,
	)

}

// Test Real(map[int]int, map[int]int{})
func TestCheckBuiltinRealMakeTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int, map[int]int{})`, env,
		`too many arguments to real: real(map[int]int, composite literal)`,
	)

}

// Test Real(map[int]int, int)
func TestCheckBuiltinRealMakeTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int, int)`, env,
		`too many arguments to real: real(map[int]int, int)`,
	)

}

// Test Real(map[int]int, map[int]int)
func TestCheckBuiltinRealMakeTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int, map[int]int)`, env,
		`too many arguments to real: real(map[int]int, map[int]int)`,
	)

}

// Test Real(map[int]int, 1, 1)
func TestCheckBuiltinRealMakeTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int, 1, 1)`, env,
		`too many arguments to real: real(map[int]int, 1, 1)`,
	)

}

// Test Real(map[int]int, []int{1,2}...)
func TestCheckBuiltinRealMakeTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `real(map[int]int, []int{1,2}...)`, env,
		`invalid use of ... with builtin real`,
		`too many arguments to real: real(map[int]int, composite literal...)`,
	)

}

// Test Imag()
func TestCheckBuiltinImagXX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag()`, env,
		`missing argument to imag: imag()`,
	)

}

// Test Imag(1)
func TestCheckBuiltinImagXInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1)`, env,
		`invalid argument 1 (type int) for imag`,
	)

}

// Test Imag(float32(1))
func TestCheckBuiltinImagXFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(float32(1))`, env,
		`invalid argument float32(1) (type float32) for imag`,
	)

}

// Test Imag("abc")
func TestCheckBuiltinImagXString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag("abc")`, env,
		`invalid argument "abc" (type string) for imag`,
	)

}

// Test Imag(nil)
func TestCheckBuiltinImagXNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(nil)`, env,
		`use of untyped nil`,
	)

}

// Test Imag(1.5)
func TestCheckBuiltinImagXFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1.5)`, env,
		`invalid argument 1.5 (type float64) for imag`,
	)

}

// Test Imag([]int{})
func TestCheckBuiltinImagXSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag([]int{})`, env,
		`invalid argument []int literal (type []int) for imag`,
	)

}

// Test Imag(map[int]int{})
func TestCheckBuiltinImagXMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int{})`, env,
		`invalid argument map[int]int literal (type map[int]int) for imag`,
	)

}

// Test Imag(int)
func TestCheckBuiltinImagXType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(int)`, env,
		`type int is not an expression`,
	)

}

// Test Imag(map[int]int)
func TestCheckBuiltinImagXMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Imag(1, 1)
func TestCheckBuiltinImagXDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1, 1)`, env,
		`too many arguments to imag: imag(1, 1)`,
	)

}

// Test Imag([]int{1,2}...)
func TestCheckBuiltinImagXEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag([]int{1,2}...)`, env,
		`invalid use of ... with builtin imag`,
		`invalid argument []int literal (type []int) for imag`,
	)

}

// Test Imag(1)
func TestCheckBuiltinImagIntX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1)`, env,
		`invalid argument 1 (type int) for imag`,
	)

}

// Test Imag(1, 1)
func TestCheckBuiltinImagIntInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1, 1)`, env,
		`too many arguments to imag: imag(1, 1)`,
	)

}

// Test Imag(1, float32(1))
func TestCheckBuiltinImagIntFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1, float32(1))`, env,
		`too many arguments to imag: imag(1, float32(1))`,
	)

}

// Test Imag(1, "abc")
func TestCheckBuiltinImagIntString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1, "abc")`, env,
		`too many arguments to imag: imag(1, "abc")`,
	)

}

// Test Imag(1, nil)
func TestCheckBuiltinImagIntNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1, nil)`, env,
		`too many arguments to imag: imag(1, nil)`,
	)

}

// Test Imag(1, 1.5)
func TestCheckBuiltinImagIntFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1, 1.5)`, env,
		`too many arguments to imag: imag(1, 1.5)`,
	)

}

// Test Imag(1, []int{})
func TestCheckBuiltinImagIntSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1, []int{})`, env,
		`too many arguments to imag: imag(1, composite literal)`,
	)

}

// Test Imag(1, map[int]int{})
func TestCheckBuiltinImagIntMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1, map[int]int{})`, env,
		`too many arguments to imag: imag(1, composite literal)`,
	)

}

// Test Imag(1, int)
func TestCheckBuiltinImagIntType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1, int)`, env,
		`too many arguments to imag: imag(1, int)`,
	)

}

// Test Imag(1, map[int]int)
func TestCheckBuiltinImagIntMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1, map[int]int)`, env,
		`too many arguments to imag: imag(1, map[int]int)`,
	)

}

// Test Imag(1, 1, 1)
func TestCheckBuiltinImagIntDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1, 1, 1)`, env,
		`too many arguments to imag: imag(1, 1, 1)`,
	)

}

// Test Imag(1, []int{1,2}...)
func TestCheckBuiltinImagIntEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1, []int{1,2}...)`, env,
		`invalid use of ... with builtin imag`,
		`too many arguments to imag: imag(1, composite literal...)`,
	)

}

// Test Imag(float32(1))
func TestCheckBuiltinImagFloat32X(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(float32(1))`, env,
		`invalid argument float32(1) (type float32) for imag`,
	)

}

// Test Imag(float32(1), 1)
func TestCheckBuiltinImagFloat32Int(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(float32(1), 1)`, env,
		`too many arguments to imag: imag(float32(1), 1)`,
	)

}

// Test Imag(float32(1), float32(1))
func TestCheckBuiltinImagFloat32Float32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(float32(1), float32(1))`, env,
		`too many arguments to imag: imag(float32(1), float32(1))`,
	)

}

// Test Imag(float32(1), "abc")
func TestCheckBuiltinImagFloat32String(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(float32(1), "abc")`, env,
		`too many arguments to imag: imag(float32(1), "abc")`,
	)

}

// Test Imag(float32(1), nil)
func TestCheckBuiltinImagFloat32Nil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(float32(1), nil)`, env,
		`too many arguments to imag: imag(float32(1), nil)`,
	)

}

// Test Imag(float32(1), 1.5)
func TestCheckBuiltinImagFloat32Float(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(float32(1), 1.5)`, env,
		`too many arguments to imag: imag(float32(1), 1.5)`,
	)

}

// Test Imag(float32(1), []int{})
func TestCheckBuiltinImagFloat32Slice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(float32(1), []int{})`, env,
		`too many arguments to imag: imag(float32(1), composite literal)`,
	)

}

// Test Imag(float32(1), map[int]int{})
func TestCheckBuiltinImagFloat32Map(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(float32(1), map[int]int{})`, env,
		`too many arguments to imag: imag(float32(1), composite literal)`,
	)

}

// Test Imag(float32(1), int)
func TestCheckBuiltinImagFloat32Type(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(float32(1), int)`, env,
		`too many arguments to imag: imag(float32(1), int)`,
	)

}

// Test Imag(float32(1), map[int]int)
func TestCheckBuiltinImagFloat32MakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(float32(1), map[int]int)`, env,
		`too many arguments to imag: imag(float32(1), map[int]int)`,
	)

}

// Test Imag(float32(1), 1, 1)
func TestCheckBuiltinImagFloat32Double(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(float32(1), 1, 1)`, env,
		`too many arguments to imag: imag(float32(1), 1, 1)`,
	)

}

// Test Imag(float32(1), []int{1,2}...)
func TestCheckBuiltinImagFloat32Ellipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(float32(1), []int{1,2}...)`, env,
		`invalid use of ... with builtin imag`,
		`too many arguments to imag: imag(float32(1), composite literal...)`,
	)

}

// Test Imag("abc")
func TestCheckBuiltinImagStringX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag("abc")`, env,
		`invalid argument "abc" (type string) for imag`,
	)

}

// Test Imag("abc", 1)
func TestCheckBuiltinImagStringInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag("abc", 1)`, env,
		`too many arguments to imag: imag("abc", 1)`,
	)

}

// Test Imag("abc", float32(1))
func TestCheckBuiltinImagStringFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag("abc", float32(1))`, env,
		`too many arguments to imag: imag("abc", float32(1))`,
	)

}

// Test Imag("abc", "abc")
func TestCheckBuiltinImagStringString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag("abc", "abc")`, env,
		`too many arguments to imag: imag("abc", "abc")`,
	)

}

// Test Imag("abc", nil)
func TestCheckBuiltinImagStringNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag("abc", nil)`, env,
		`too many arguments to imag: imag("abc", nil)`,
	)

}

// Test Imag("abc", 1.5)
func TestCheckBuiltinImagStringFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag("abc", 1.5)`, env,
		`too many arguments to imag: imag("abc", 1.5)`,
	)

}

// Test Imag("abc", []int{})
func TestCheckBuiltinImagStringSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag("abc", []int{})`, env,
		`too many arguments to imag: imag("abc", composite literal)`,
	)

}

// Test Imag("abc", map[int]int{})
func TestCheckBuiltinImagStringMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag("abc", map[int]int{})`, env,
		`too many arguments to imag: imag("abc", composite literal)`,
	)

}

// Test Imag("abc", int)
func TestCheckBuiltinImagStringType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag("abc", int)`, env,
		`too many arguments to imag: imag("abc", int)`,
	)

}

// Test Imag("abc", map[int]int)
func TestCheckBuiltinImagStringMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag("abc", map[int]int)`, env,
		`too many arguments to imag: imag("abc", map[int]int)`,
	)

}

// Test Imag("abc", 1, 1)
func TestCheckBuiltinImagStringDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag("abc", 1, 1)`, env,
		`too many arguments to imag: imag("abc", 1, 1)`,
	)

}

// Test Imag("abc", []int{1,2}...)
func TestCheckBuiltinImagStringEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag("abc", []int{1,2}...)`, env,
		`invalid use of ... with builtin imag`,
		`too many arguments to imag: imag("abc", composite literal...)`,
	)

}

// Test Imag(nil)
func TestCheckBuiltinImagNilX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(nil)`, env,
		`use of untyped nil`,
	)

}

// Test Imag(nil, 1)
func TestCheckBuiltinImagNilInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(nil, 1)`, env,
		`too many arguments to imag: imag(nil, 1)`,
	)

}

// Test Imag(nil, float32(1))
func TestCheckBuiltinImagNilFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(nil, float32(1))`, env,
		`too many arguments to imag: imag(nil, float32(1))`,
	)

}

// Test Imag(nil, "abc")
func TestCheckBuiltinImagNilString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(nil, "abc")`, env,
		`too many arguments to imag: imag(nil, "abc")`,
	)

}

// Test Imag(nil, nil)
func TestCheckBuiltinImagNilNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(nil, nil)`, env,
		`too many arguments to imag: imag(nil, nil)`,
	)

}

// Test Imag(nil, 1.5)
func TestCheckBuiltinImagNilFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(nil, 1.5)`, env,
		`too many arguments to imag: imag(nil, 1.5)`,
	)

}

// Test Imag(nil, []int{})
func TestCheckBuiltinImagNilSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(nil, []int{})`, env,
		`too many arguments to imag: imag(nil, composite literal)`,
	)

}

// Test Imag(nil, map[int]int{})
func TestCheckBuiltinImagNilMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(nil, map[int]int{})`, env,
		`too many arguments to imag: imag(nil, composite literal)`,
	)

}

// Test Imag(nil, int)
func TestCheckBuiltinImagNilType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(nil, int)`, env,
		`too many arguments to imag: imag(nil, int)`,
	)

}

// Test Imag(nil, map[int]int)
func TestCheckBuiltinImagNilMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(nil, map[int]int)`, env,
		`too many arguments to imag: imag(nil, map[int]int)`,
	)

}

// Test Imag(nil, 1, 1)
func TestCheckBuiltinImagNilDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(nil, 1, 1)`, env,
		`too many arguments to imag: imag(nil, 1, 1)`,
	)

}

// Test Imag(nil, []int{1,2}...)
func TestCheckBuiltinImagNilEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(nil, []int{1,2}...)`, env,
		`invalid use of ... with builtin imag`,
		`too many arguments to imag: imag(nil, composite literal...)`,
	)

}

// Test Imag(1.5)
func TestCheckBuiltinImagFloatX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1.5)`, env,
		`invalid argument 1.5 (type float64) for imag`,
	)

}

// Test Imag(1.5, 1)
func TestCheckBuiltinImagFloatInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1.5, 1)`, env,
		`too many arguments to imag: imag(1.5, 1)`,
	)

}

// Test Imag(1.5, float32(1))
func TestCheckBuiltinImagFloatFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1.5, float32(1))`, env,
		`too many arguments to imag: imag(1.5, float32(1))`,
	)

}

// Test Imag(1.5, "abc")
func TestCheckBuiltinImagFloatString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1.5, "abc")`, env,
		`too many arguments to imag: imag(1.5, "abc")`,
	)

}

// Test Imag(1.5, nil)
func TestCheckBuiltinImagFloatNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1.5, nil)`, env,
		`too many arguments to imag: imag(1.5, nil)`,
	)

}

// Test Imag(1.5, 1.5)
func TestCheckBuiltinImagFloatFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1.5, 1.5)`, env,
		`too many arguments to imag: imag(1.5, 1.5)`,
	)

}

// Test Imag(1.5, []int{})
func TestCheckBuiltinImagFloatSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1.5, []int{})`, env,
		`too many arguments to imag: imag(1.5, composite literal)`,
	)

}

// Test Imag(1.5, map[int]int{})
func TestCheckBuiltinImagFloatMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1.5, map[int]int{})`, env,
		`too many arguments to imag: imag(1.5, composite literal)`,
	)

}

// Test Imag(1.5, int)
func TestCheckBuiltinImagFloatType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1.5, int)`, env,
		`too many arguments to imag: imag(1.5, int)`,
	)

}

// Test Imag(1.5, map[int]int)
func TestCheckBuiltinImagFloatMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1.5, map[int]int)`, env,
		`too many arguments to imag: imag(1.5, map[int]int)`,
	)

}

// Test Imag(1.5, 1, 1)
func TestCheckBuiltinImagFloatDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1.5, 1, 1)`, env,
		`too many arguments to imag: imag(1.5, 1, 1)`,
	)

}

// Test Imag(1.5, []int{1,2}...)
func TestCheckBuiltinImagFloatEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(1.5, []int{1,2}...)`, env,
		`invalid use of ... with builtin imag`,
		`too many arguments to imag: imag(1.5, composite literal...)`,
	)

}

// Test Imag([]int{})
func TestCheckBuiltinImagSliceX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag([]int{})`, env,
		`invalid argument []int literal (type []int) for imag`,
	)

}

// Test Imag([]int{}, 1)
func TestCheckBuiltinImagSliceInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag([]int{}, 1)`, env,
		`too many arguments to imag: imag(composite literal, 1)`,
	)

}

// Test Imag([]int{}, float32(1))
func TestCheckBuiltinImagSliceFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag([]int{}, float32(1))`, env,
		`too many arguments to imag: imag(composite literal, float32(1))`,
	)

}

// Test Imag([]int{}, "abc")
func TestCheckBuiltinImagSliceString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag([]int{}, "abc")`, env,
		`too many arguments to imag: imag(composite literal, "abc")`,
	)

}

// Test Imag([]int{}, nil)
func TestCheckBuiltinImagSliceNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag([]int{}, nil)`, env,
		`too many arguments to imag: imag(composite literal, nil)`,
	)

}

// Test Imag([]int{}, 1.5)
func TestCheckBuiltinImagSliceFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag([]int{}, 1.5)`, env,
		`too many arguments to imag: imag(composite literal, 1.5)`,
	)

}

// Test Imag([]int{}, []int{})
func TestCheckBuiltinImagSliceSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag([]int{}, []int{})`, env,
		`too many arguments to imag: imag(composite literal, composite literal)`,
	)

}

// Test Imag([]int{}, map[int]int{})
func TestCheckBuiltinImagSliceMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag([]int{}, map[int]int{})`, env,
		`too many arguments to imag: imag(composite literal, composite literal)`,
	)

}

// Test Imag([]int{}, int)
func TestCheckBuiltinImagSliceType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag([]int{}, int)`, env,
		`too many arguments to imag: imag(composite literal, int)`,
	)

}

// Test Imag([]int{}, map[int]int)
func TestCheckBuiltinImagSliceMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag([]int{}, map[int]int)`, env,
		`too many arguments to imag: imag(composite literal, map[int]int)`,
	)

}

// Test Imag([]int{}, 1, 1)
func TestCheckBuiltinImagSliceDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag([]int{}, 1, 1)`, env,
		`too many arguments to imag: imag(composite literal, 1, 1)`,
	)

}

// Test Imag([]int{}, []int{1,2}...)
func TestCheckBuiltinImagSliceEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag([]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin imag`,
		`too many arguments to imag: imag(composite literal, composite literal...)`,
	)

}

// Test Imag(map[int]int{})
func TestCheckBuiltinImagMapX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int{})`, env,
		`invalid argument map[int]int literal (type map[int]int) for imag`,
	)

}

// Test Imag(map[int]int{}, 1)
func TestCheckBuiltinImagMapInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int{}, 1)`, env,
		`too many arguments to imag: imag(composite literal, 1)`,
	)

}

// Test Imag(map[int]int{}, float32(1))
func TestCheckBuiltinImagMapFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int{}, float32(1))`, env,
		`too many arguments to imag: imag(composite literal, float32(1))`,
	)

}

// Test Imag(map[int]int{}, "abc")
func TestCheckBuiltinImagMapString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int{}, "abc")`, env,
		`too many arguments to imag: imag(composite literal, "abc")`,
	)

}

// Test Imag(map[int]int{}, nil)
func TestCheckBuiltinImagMapNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int{}, nil)`, env,
		`too many arguments to imag: imag(composite literal, nil)`,
	)

}

// Test Imag(map[int]int{}, 1.5)
func TestCheckBuiltinImagMapFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int{}, 1.5)`, env,
		`too many arguments to imag: imag(composite literal, 1.5)`,
	)

}

// Test Imag(map[int]int{}, []int{})
func TestCheckBuiltinImagMapSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int{}, []int{})`, env,
		`too many arguments to imag: imag(composite literal, composite literal)`,
	)

}

// Test Imag(map[int]int{}, map[int]int{})
func TestCheckBuiltinImagMapMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int{}, map[int]int{})`, env,
		`too many arguments to imag: imag(composite literal, composite literal)`,
	)

}

// Test Imag(map[int]int{}, int)
func TestCheckBuiltinImagMapType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int{}, int)`, env,
		`too many arguments to imag: imag(composite literal, int)`,
	)

}

// Test Imag(map[int]int{}, map[int]int)
func TestCheckBuiltinImagMapMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int{}, map[int]int)`, env,
		`too many arguments to imag: imag(composite literal, map[int]int)`,
	)

}

// Test Imag(map[int]int{}, 1, 1)
func TestCheckBuiltinImagMapDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int{}, 1, 1)`, env,
		`too many arguments to imag: imag(composite literal, 1, 1)`,
	)

}

// Test Imag(map[int]int{}, []int{1,2}...)
func TestCheckBuiltinImagMapEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin imag`,
		`too many arguments to imag: imag(composite literal, composite literal...)`,
	)

}

// Test Imag(int)
func TestCheckBuiltinImagTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(int)`, env,
		`type int is not an expression`,
	)

}

// Test Imag(int, 1)
func TestCheckBuiltinImagTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(int, 1)`, env,
		`too many arguments to imag: imag(int, 1)`,
	)

}

// Test Imag(int, float32(1))
func TestCheckBuiltinImagTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(int, float32(1))`, env,
		`too many arguments to imag: imag(int, float32(1))`,
	)

}

// Test Imag(int, "abc")
func TestCheckBuiltinImagTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(int, "abc")`, env,
		`too many arguments to imag: imag(int, "abc")`,
	)

}

// Test Imag(int, nil)
func TestCheckBuiltinImagTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(int, nil)`, env,
		`too many arguments to imag: imag(int, nil)`,
	)

}

// Test Imag(int, 1.5)
func TestCheckBuiltinImagTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(int, 1.5)`, env,
		`too many arguments to imag: imag(int, 1.5)`,
	)

}

// Test Imag(int, []int{})
func TestCheckBuiltinImagTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(int, []int{})`, env,
		`too many arguments to imag: imag(int, composite literal)`,
	)

}

// Test Imag(int, map[int]int{})
func TestCheckBuiltinImagTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(int, map[int]int{})`, env,
		`too many arguments to imag: imag(int, composite literal)`,
	)

}

// Test Imag(int, int)
func TestCheckBuiltinImagTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(int, int)`, env,
		`too many arguments to imag: imag(int, int)`,
	)

}

// Test Imag(int, map[int]int)
func TestCheckBuiltinImagTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(int, map[int]int)`, env,
		`too many arguments to imag: imag(int, map[int]int)`,
	)

}

// Test Imag(int, 1, 1)
func TestCheckBuiltinImagTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(int, 1, 1)`, env,
		`too many arguments to imag: imag(int, 1, 1)`,
	)

}

// Test Imag(int, []int{1,2}...)
func TestCheckBuiltinImagTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(int, []int{1,2}...)`, env,
		`invalid use of ... with builtin imag`,
		`too many arguments to imag: imag(int, composite literal...)`,
	)

}

// Test Imag(map[int]int)
func TestCheckBuiltinImagMakeTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Imag(map[int]int, 1)
func TestCheckBuiltinImagMakeTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int, 1)`, env,
		`too many arguments to imag: imag(map[int]int, 1)`,
	)

}

// Test Imag(map[int]int, float32(1))
func TestCheckBuiltinImagMakeTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int, float32(1))`, env,
		`too many arguments to imag: imag(map[int]int, float32(1))`,
	)

}

// Test Imag(map[int]int, "abc")
func TestCheckBuiltinImagMakeTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int, "abc")`, env,
		`too many arguments to imag: imag(map[int]int, "abc")`,
	)

}

// Test Imag(map[int]int, nil)
func TestCheckBuiltinImagMakeTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int, nil)`, env,
		`too many arguments to imag: imag(map[int]int, nil)`,
	)

}

// Test Imag(map[int]int, 1.5)
func TestCheckBuiltinImagMakeTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int, 1.5)`, env,
		`too many arguments to imag: imag(map[int]int, 1.5)`,
	)

}

// Test Imag(map[int]int, []int{})
func TestCheckBuiltinImagMakeTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int, []int{})`, env,
		`too many arguments to imag: imag(map[int]int, composite literal)`,
	)

}

// Test Imag(map[int]int, map[int]int{})
func TestCheckBuiltinImagMakeTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int, map[int]int{})`, env,
		`too many arguments to imag: imag(map[int]int, composite literal)`,
	)

}

// Test Imag(map[int]int, int)
func TestCheckBuiltinImagMakeTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int, int)`, env,
		`too many arguments to imag: imag(map[int]int, int)`,
	)

}

// Test Imag(map[int]int, map[int]int)
func TestCheckBuiltinImagMakeTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int, map[int]int)`, env,
		`too many arguments to imag: imag(map[int]int, map[int]int)`,
	)

}

// Test Imag(map[int]int, 1, 1)
func TestCheckBuiltinImagMakeTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int, 1, 1)`, env,
		`too many arguments to imag: imag(map[int]int, 1, 1)`,
	)

}

// Test Imag(map[int]int, []int{1,2}...)
func TestCheckBuiltinImagMakeTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `imag(map[int]int, []int{1,2}...)`, env,
		`invalid use of ... with builtin imag`,
		`too many arguments to imag: imag(map[int]int, composite literal...)`,
	)

}

// Test New()
func TestCheckBuiltinNewXX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new()`, env,
		`missing argument to new`,
	)

}

// Test New(1)
func TestCheckBuiltinNewXInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1)`, env,
		`1 is not a type`,
	)

}

// Test New(float32(1))
func TestCheckBuiltinNewXFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(float32(1))`, env,
		`float32(1) is not a type`,
	)

}

// Test New("abc")
func TestCheckBuiltinNewXString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new("abc")`, env,
		`"abc" is not a type`,
	)

}

// Test New(nil)
func TestCheckBuiltinNewXNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(nil)`, env,
		`nil is not a type`,
	)

}

// Test New(1.5)
func TestCheckBuiltinNewXFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1.5)`, env,
		`1.5 is not a type`,
	)

}

// Test New([]int{})
func TestCheckBuiltinNewXSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new([]int{})`, env,
		`[]int literal is not a type`,
	)

}

// Test New(map[int]int{})
func TestCheckBuiltinNewXMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int{})`, env,
		`map[int]int literal is not a type`,
	)

}

// Test New(int)
func TestCheckBuiltinNewXType(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `new(int)`, env, reflect.TypeOf(new(int)))
}

// Test New(map[int]int)
func TestCheckBuiltinNewXMakeType(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `new(map[int]int)`, env, reflect.TypeOf(new(map[int]int)))
}

// Test New(1, 1)
func TestCheckBuiltinNewXDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1, 1)`, env,
		`1 is not a type`,
	)

}

// Test New([]int{1,2}...)
func TestCheckBuiltinNewXEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new([]int{1,2}...)`, env,
		`invalid use of ... with builtin new`,
		`[]int literal is not a type`,
	)

}

// Test New(1)
func TestCheckBuiltinNewIntX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1)`, env,
		`1 is not a type`,
	)

}

// Test New(1, 1)
func TestCheckBuiltinNewIntInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1, 1)`, env,
		`1 is not a type`,
	)

}

// Test New(1, float32(1))
func TestCheckBuiltinNewIntFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1, float32(1))`, env,
		`1 is not a type`,
	)

}

// Test New(1, "abc")
func TestCheckBuiltinNewIntString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1, "abc")`, env,
		`1 is not a type`,
	)

}

// Test New(1, nil)
func TestCheckBuiltinNewIntNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1, nil)`, env,
		`1 is not a type`,
	)

}

// Test New(1, 1.5)
func TestCheckBuiltinNewIntFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1, 1.5)`, env,
		`1 is not a type`,
	)

}

// Test New(1, []int{})
func TestCheckBuiltinNewIntSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1, []int{})`, env,
		`1 is not a type`,
	)

}

// Test New(1, map[int]int{})
func TestCheckBuiltinNewIntMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1, map[int]int{})`, env,
		`1 is not a type`,
	)

}

// Test New(1, int)
func TestCheckBuiltinNewIntType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1, int)`, env,
		`1 is not a type`,
	)

}

// Test New(1, map[int]int)
func TestCheckBuiltinNewIntMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1, map[int]int)`, env,
		`1 is not a type`,
	)

}

// Test New(1, 1, 1)
func TestCheckBuiltinNewIntDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1, 1, 1)`, env,
		`1 is not a type`,
	)

}

// Test New(1, []int{1,2}...)
func TestCheckBuiltinNewIntEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1, []int{1,2}...)`, env,
		`invalid use of ... with builtin new`,
		`1 is not a type`,
	)

}

// Test New(float32(1))
func TestCheckBuiltinNewFloat32X(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(float32(1))`, env,
		`float32(1) is not a type`,
	)

}

// Test New(float32(1), 1)
func TestCheckBuiltinNewFloat32Int(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(float32(1), 1)`, env,
		`float32(1) is not a type`,
	)

}

// Test New(float32(1), float32(1))
func TestCheckBuiltinNewFloat32Float32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(float32(1), float32(1))`, env,
		`float32(1) is not a type`,
	)

}

// Test New(float32(1), "abc")
func TestCheckBuiltinNewFloat32String(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(float32(1), "abc")`, env,
		`float32(1) is not a type`,
	)

}

// Test New(float32(1), nil)
func TestCheckBuiltinNewFloat32Nil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(float32(1), nil)`, env,
		`float32(1) is not a type`,
	)

}

// Test New(float32(1), 1.5)
func TestCheckBuiltinNewFloat32Float(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(float32(1), 1.5)`, env,
		`float32(1) is not a type`,
	)

}

// Test New(float32(1), []int{})
func TestCheckBuiltinNewFloat32Slice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(float32(1), []int{})`, env,
		`float32(1) is not a type`,
	)

}

// Test New(float32(1), map[int]int{})
func TestCheckBuiltinNewFloat32Map(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(float32(1), map[int]int{})`, env,
		`float32(1) is not a type`,
	)

}

// Test New(float32(1), int)
func TestCheckBuiltinNewFloat32Type(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(float32(1), int)`, env,
		`float32(1) is not a type`,
	)

}

// Test New(float32(1), map[int]int)
func TestCheckBuiltinNewFloat32MakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(float32(1), map[int]int)`, env,
		`float32(1) is not a type`,
	)

}

// Test New(float32(1), 1, 1)
func TestCheckBuiltinNewFloat32Double(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(float32(1), 1, 1)`, env,
		`float32(1) is not a type`,
	)

}

// Test New(float32(1), []int{1,2}...)
func TestCheckBuiltinNewFloat32Ellipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(float32(1), []int{1,2}...)`, env,
		`invalid use of ... with builtin new`,
		`float32(1) is not a type`,
	)

}

// Test New("abc")
func TestCheckBuiltinNewStringX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new("abc")`, env,
		`"abc" is not a type`,
	)

}

// Test New("abc", 1)
func TestCheckBuiltinNewStringInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new("abc", 1)`, env,
		`"abc" is not a type`,
	)

}

// Test New("abc", float32(1))
func TestCheckBuiltinNewStringFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new("abc", float32(1))`, env,
		`"abc" is not a type`,
	)

}

// Test New("abc", "abc")
func TestCheckBuiltinNewStringString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new("abc", "abc")`, env,
		`"abc" is not a type`,
	)

}

// Test New("abc", nil)
func TestCheckBuiltinNewStringNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new("abc", nil)`, env,
		`"abc" is not a type`,
	)

}

// Test New("abc", 1.5)
func TestCheckBuiltinNewStringFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new("abc", 1.5)`, env,
		`"abc" is not a type`,
	)

}

// Test New("abc", []int{})
func TestCheckBuiltinNewStringSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new("abc", []int{})`, env,
		`"abc" is not a type`,
	)

}

// Test New("abc", map[int]int{})
func TestCheckBuiltinNewStringMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new("abc", map[int]int{})`, env,
		`"abc" is not a type`,
	)

}

// Test New("abc", int)
func TestCheckBuiltinNewStringType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new("abc", int)`, env,
		`"abc" is not a type`,
	)

}

// Test New("abc", map[int]int)
func TestCheckBuiltinNewStringMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new("abc", map[int]int)`, env,
		`"abc" is not a type`,
	)

}

// Test New("abc", 1, 1)
func TestCheckBuiltinNewStringDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new("abc", 1, 1)`, env,
		`"abc" is not a type`,
	)

}

// Test New("abc", []int{1,2}...)
func TestCheckBuiltinNewStringEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new("abc", []int{1,2}...)`, env,
		`invalid use of ... with builtin new`,
		`"abc" is not a type`,
	)

}

// Test New(nil)
func TestCheckBuiltinNewNilX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(nil)`, env,
		`nil is not a type`,
	)

}

// Test New(nil, 1)
func TestCheckBuiltinNewNilInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(nil, 1)`, env,
		`nil is not a type`,
	)

}

// Test New(nil, float32(1))
func TestCheckBuiltinNewNilFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(nil, float32(1))`, env,
		`nil is not a type`,
	)

}

// Test New(nil, "abc")
func TestCheckBuiltinNewNilString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(nil, "abc")`, env,
		`nil is not a type`,
	)

}

// Test New(nil, nil)
func TestCheckBuiltinNewNilNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(nil, nil)`, env,
		`nil is not a type`,
	)

}

// Test New(nil, 1.5)
func TestCheckBuiltinNewNilFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(nil, 1.5)`, env,
		`nil is not a type`,
	)

}

// Test New(nil, []int{})
func TestCheckBuiltinNewNilSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(nil, []int{})`, env,
		`nil is not a type`,
	)

}

// Test New(nil, map[int]int{})
func TestCheckBuiltinNewNilMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(nil, map[int]int{})`, env,
		`nil is not a type`,
	)

}

// Test New(nil, int)
func TestCheckBuiltinNewNilType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(nil, int)`, env,
		`nil is not a type`,
	)

}

// Test New(nil, map[int]int)
func TestCheckBuiltinNewNilMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(nil, map[int]int)`, env,
		`nil is not a type`,
	)

}

// Test New(nil, 1, 1)
func TestCheckBuiltinNewNilDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(nil, 1, 1)`, env,
		`nil is not a type`,
	)

}

// Test New(nil, []int{1,2}...)
func TestCheckBuiltinNewNilEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(nil, []int{1,2}...)`, env,
		`invalid use of ... with builtin new`,
		`nil is not a type`,
	)

}

// Test New(1.5)
func TestCheckBuiltinNewFloatX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1.5)`, env,
		`1.5 is not a type`,
	)

}

// Test New(1.5, 1)
func TestCheckBuiltinNewFloatInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1.5, 1)`, env,
		`1.5 is not a type`,
	)

}

// Test New(1.5, float32(1))
func TestCheckBuiltinNewFloatFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1.5, float32(1))`, env,
		`1.5 is not a type`,
	)

}

// Test New(1.5, "abc")
func TestCheckBuiltinNewFloatString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1.5, "abc")`, env,
		`1.5 is not a type`,
	)

}

// Test New(1.5, nil)
func TestCheckBuiltinNewFloatNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1.5, nil)`, env,
		`1.5 is not a type`,
	)

}

// Test New(1.5, 1.5)
func TestCheckBuiltinNewFloatFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1.5, 1.5)`, env,
		`1.5 is not a type`,
	)

}

// Test New(1.5, []int{})
func TestCheckBuiltinNewFloatSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1.5, []int{})`, env,
		`1.5 is not a type`,
	)

}

// Test New(1.5, map[int]int{})
func TestCheckBuiltinNewFloatMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1.5, map[int]int{})`, env,
		`1.5 is not a type`,
	)

}

// Test New(1.5, int)
func TestCheckBuiltinNewFloatType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1.5, int)`, env,
		`1.5 is not a type`,
	)

}

// Test New(1.5, map[int]int)
func TestCheckBuiltinNewFloatMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1.5, map[int]int)`, env,
		`1.5 is not a type`,
	)

}

// Test New(1.5, 1, 1)
func TestCheckBuiltinNewFloatDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1.5, 1, 1)`, env,
		`1.5 is not a type`,
	)

}

// Test New(1.5, []int{1,2}...)
func TestCheckBuiltinNewFloatEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(1.5, []int{1,2}...)`, env,
		`invalid use of ... with builtin new`,
		`1.5 is not a type`,
	)

}

// Test New([]int{})
func TestCheckBuiltinNewSliceX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new([]int{})`, env,
		`[]int literal is not a type`,
	)

}

// Test New([]int{}, 1)
func TestCheckBuiltinNewSliceInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new([]int{}, 1)`, env,
		`[]int literal is not a type`,
	)

}

// Test New([]int{}, float32(1))
func TestCheckBuiltinNewSliceFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new([]int{}, float32(1))`, env,
		`[]int literal is not a type`,
	)

}

// Test New([]int{}, "abc")
func TestCheckBuiltinNewSliceString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new([]int{}, "abc")`, env,
		`[]int literal is not a type`,
	)

}

// Test New([]int{}, nil)
func TestCheckBuiltinNewSliceNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new([]int{}, nil)`, env,
		`[]int literal is not a type`,
	)

}

// Test New([]int{}, 1.5)
func TestCheckBuiltinNewSliceFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new([]int{}, 1.5)`, env,
		`[]int literal is not a type`,
	)

}

// Test New([]int{}, []int{})
func TestCheckBuiltinNewSliceSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new([]int{}, []int{})`, env,
		`[]int literal is not a type`,
	)

}

// Test New([]int{}, map[int]int{})
func TestCheckBuiltinNewSliceMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new([]int{}, map[int]int{})`, env,
		`[]int literal is not a type`,
	)

}

// Test New([]int{}, int)
func TestCheckBuiltinNewSliceType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new([]int{}, int)`, env,
		`[]int literal is not a type`,
	)

}

// Test New([]int{}, map[int]int)
func TestCheckBuiltinNewSliceMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new([]int{}, map[int]int)`, env,
		`[]int literal is not a type`,
	)

}

// Test New([]int{}, 1, 1)
func TestCheckBuiltinNewSliceDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new([]int{}, 1, 1)`, env,
		`[]int literal is not a type`,
	)

}

// Test New([]int{}, []int{1,2}...)
func TestCheckBuiltinNewSliceEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new([]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin new`,
		`[]int literal is not a type`,
	)

}

// Test New(map[int]int{})
func TestCheckBuiltinNewMapX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int{})`, env,
		`map[int]int literal is not a type`,
	)

}

// Test New(map[int]int{}, 1)
func TestCheckBuiltinNewMapInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int{}, 1)`, env,
		`map[int]int literal is not a type`,
	)

}

// Test New(map[int]int{}, float32(1))
func TestCheckBuiltinNewMapFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int{}, float32(1))`, env,
		`map[int]int literal is not a type`,
	)

}

// Test New(map[int]int{}, "abc")
func TestCheckBuiltinNewMapString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int{}, "abc")`, env,
		`map[int]int literal is not a type`,
	)

}

// Test New(map[int]int{}, nil)
func TestCheckBuiltinNewMapNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int{}, nil)`, env,
		`map[int]int literal is not a type`,
	)

}

// Test New(map[int]int{}, 1.5)
func TestCheckBuiltinNewMapFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int{}, 1.5)`, env,
		`map[int]int literal is not a type`,
	)

}

// Test New(map[int]int{}, []int{})
func TestCheckBuiltinNewMapSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int{}, []int{})`, env,
		`map[int]int literal is not a type`,
	)

}

// Test New(map[int]int{}, map[int]int{})
func TestCheckBuiltinNewMapMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int{}, map[int]int{})`, env,
		`map[int]int literal is not a type`,
	)

}

// Test New(map[int]int{}, int)
func TestCheckBuiltinNewMapType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int{}, int)`, env,
		`map[int]int literal is not a type`,
	)

}

// Test New(map[int]int{}, map[int]int)
func TestCheckBuiltinNewMapMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int{}, map[int]int)`, env,
		`map[int]int literal is not a type`,
	)

}

// Test New(map[int]int{}, 1, 1)
func TestCheckBuiltinNewMapDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int{}, 1, 1)`, env,
		`map[int]int literal is not a type`,
	)

}

// Test New(map[int]int{}, []int{1,2}...)
func TestCheckBuiltinNewMapEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin new`,
		`map[int]int literal is not a type`,
	)

}

// Test New(int)
func TestCheckBuiltinNewTypeX(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `new(int)`, env, reflect.TypeOf(new(int)))
}

// Test New(int, 1)
func TestCheckBuiltinNewTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(int, 1)`, env,
		`too many arguments to new(int)`,
	)

}

// Test New(int, float32(1))
func TestCheckBuiltinNewTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(int, float32(1))`, env,
		`too many arguments to new(int)`,
	)

}

// Test New(int, "abc")
func TestCheckBuiltinNewTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(int, "abc")`, env,
		`too many arguments to new(int)`,
	)

}

// Test New(int, nil)
func TestCheckBuiltinNewTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(int, nil)`, env,
		`too many arguments to new(int)`,
	)

}

// Test New(int, 1.5)
func TestCheckBuiltinNewTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(int, 1.5)`, env,
		`too many arguments to new(int)`,
	)

}

// Test New(int, []int{})
func TestCheckBuiltinNewTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(int, []int{})`, env,
		`too many arguments to new(int)`,
	)

}

// Test New(int, map[int]int{})
func TestCheckBuiltinNewTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(int, map[int]int{})`, env,
		`too many arguments to new(int)`,
	)

}

// Test New(int, int)
func TestCheckBuiltinNewTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(int, int)`, env,
		`too many arguments to new(int)`,
	)

}

// Test New(int, map[int]int)
func TestCheckBuiltinNewTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(int, map[int]int)`, env,
		`too many arguments to new(int)`,
	)

}

// Test New(int, 1, 1)
func TestCheckBuiltinNewTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(int, 1, 1)`, env,
		`too many arguments to new(int)`,
	)

}

// Test New(int, []int{1,2}...)
func TestCheckBuiltinNewTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(int, []int{1,2}...)`, env,
		`invalid use of ... with builtin new`,
		`too many arguments to new(int)`,
	)

}

// Test New(map[int]int)
func TestCheckBuiltinNewMakeTypeX(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `new(map[int]int)`, env, reflect.TypeOf(new(map[int]int)))
}

// Test New(map[int]int, 1)
func TestCheckBuiltinNewMakeTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int, 1)`, env,
		`too many arguments to new(map[int]int)`,
	)

}

// Test New(map[int]int, float32(1))
func TestCheckBuiltinNewMakeTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int, float32(1))`, env,
		`too many arguments to new(map[int]int)`,
	)

}

// Test New(map[int]int, "abc")
func TestCheckBuiltinNewMakeTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int, "abc")`, env,
		`too many arguments to new(map[int]int)`,
	)

}

// Test New(map[int]int, nil)
func TestCheckBuiltinNewMakeTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int, nil)`, env,
		`too many arguments to new(map[int]int)`,
	)

}

// Test New(map[int]int, 1.5)
func TestCheckBuiltinNewMakeTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int, 1.5)`, env,
		`too many arguments to new(map[int]int)`,
	)

}

// Test New(map[int]int, []int{})
func TestCheckBuiltinNewMakeTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int, []int{})`, env,
		`too many arguments to new(map[int]int)`,
	)

}

// Test New(map[int]int, map[int]int{})
func TestCheckBuiltinNewMakeTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int, map[int]int{})`, env,
		`too many arguments to new(map[int]int)`,
	)

}

// Test New(map[int]int, int)
func TestCheckBuiltinNewMakeTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int, int)`, env,
		`too many arguments to new(map[int]int)`,
	)

}

// Test New(map[int]int, map[int]int)
func TestCheckBuiltinNewMakeTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int, map[int]int)`, env,
		`too many arguments to new(map[int]int)`,
	)

}

// Test New(map[int]int, 1, 1)
func TestCheckBuiltinNewMakeTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int, 1, 1)`, env,
		`too many arguments to new(map[int]int)`,
	)

}

// Test New(map[int]int, []int{1,2}...)
func TestCheckBuiltinNewMakeTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `new(map[int]int, []int{1,2}...)`, env,
		`invalid use of ... with builtin new`,
		`too many arguments to new(map[int]int)`,
	)

}

// Test Len()
func TestCheckBuiltinLenXX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len()`, env,
		`missing argument to len: len()`,
	)

}

// Test Len(1)
func TestCheckBuiltinLenXInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1)`, env,
		`invalid argument 1 (type int) for len`,
	)

}

// Test Len(float32(1))
func TestCheckBuiltinLenXFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(float32(1))`, env,
		`invalid argument float32(1) (type float32) for len`,
	)

}

// Test Len("abc")
func TestCheckBuiltinLenXString(t *testing.T) {
	env := MakeSimpleEnv()
	expectConst(t, `len("abc")`, env, len("abc"), reflect.TypeOf(len("abc")))

}

// Test Len(nil)
func TestCheckBuiltinLenXNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(nil)`, env,
		`use of untyped nil`,
	)

}

// Test Len(1.5)
func TestCheckBuiltinLenXFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1.5)`, env,
		`invalid argument 1.5 (type float64) for len`,
	)

}

// Test Len([]int{})
func TestCheckBuiltinLenXSlice(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `len([]int{})`, env, reflect.TypeOf(len([]int{})))
}

// Test Len(map[int]int{})
func TestCheckBuiltinLenXMap(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `len(map[int]int{})`, env, reflect.TypeOf(len(map[int]int{})))
}

// Test Len(int)
func TestCheckBuiltinLenXType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(int)`, env,
		`type int is not an expression`,
	)

}

// Test Len(map[int]int)
func TestCheckBuiltinLenXMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Len(1, 1)
func TestCheckBuiltinLenXDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1, 1)`, env,
		`too many arguments to len: len(1, 1)`,
	)

}

// Test Len([]int{1,2}...)
func TestCheckBuiltinLenXEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len([]int{1,2}...)`, env,
		`invalid use of ... with builtin len`,
	)

}

// Test Len(1)
func TestCheckBuiltinLenIntX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1)`, env,
		`invalid argument 1 (type int) for len`,
	)

}

// Test Len(1, 1)
func TestCheckBuiltinLenIntInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1, 1)`, env,
		`too many arguments to len: len(1, 1)`,
	)

}

// Test Len(1, float32(1))
func TestCheckBuiltinLenIntFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1, float32(1))`, env,
		`too many arguments to len: len(1, float32(1))`,
	)

}

// Test Len(1, "abc")
func TestCheckBuiltinLenIntString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1, "abc")`, env,
		`too many arguments to len: len(1, "abc")`,
	)

}

// Test Len(1, nil)
func TestCheckBuiltinLenIntNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1, nil)`, env,
		`too many arguments to len: len(1, nil)`,
	)

}

// Test Len(1, 1.5)
func TestCheckBuiltinLenIntFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1, 1.5)`, env,
		`too many arguments to len: len(1, 1.5)`,
	)

}

// Test Len(1, []int{})
func TestCheckBuiltinLenIntSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1, []int{})`, env,
		`too many arguments to len: len(1, composite literal)`,
	)

}

// Test Len(1, map[int]int{})
func TestCheckBuiltinLenIntMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1, map[int]int{})`, env,
		`too many arguments to len: len(1, composite literal)`,
	)

}

// Test Len(1, int)
func TestCheckBuiltinLenIntType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1, int)`, env,
		`too many arguments to len: len(1, int)`,
	)

}

// Test Len(1, map[int]int)
func TestCheckBuiltinLenIntMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1, map[int]int)`, env,
		`too many arguments to len: len(1, map[int]int)`,
	)

}

// Test Len(1, 1, 1)
func TestCheckBuiltinLenIntDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1, 1, 1)`, env,
		`too many arguments to len: len(1, 1, 1)`,
	)

}

// Test Len(1, []int{1,2}...)
func TestCheckBuiltinLenIntEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1, []int{1,2}...)`, env,
		`invalid use of ... with builtin len`,
		`too many arguments to len: len(1, composite literal...)`,
	)

}

// Test Len(float32(1))
func TestCheckBuiltinLenFloat32X(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(float32(1))`, env,
		`invalid argument float32(1) (type float32) for len`,
	)

}

// Test Len(float32(1), 1)
func TestCheckBuiltinLenFloat32Int(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(float32(1), 1)`, env,
		`too many arguments to len: len(float32(1), 1)`,
	)

}

// Test Len(float32(1), float32(1))
func TestCheckBuiltinLenFloat32Float32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(float32(1), float32(1))`, env,
		`too many arguments to len: len(float32(1), float32(1))`,
	)

}

// Test Len(float32(1), "abc")
func TestCheckBuiltinLenFloat32String(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(float32(1), "abc")`, env,
		`too many arguments to len: len(float32(1), "abc")`,
	)

}

// Test Len(float32(1), nil)
func TestCheckBuiltinLenFloat32Nil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(float32(1), nil)`, env,
		`too many arguments to len: len(float32(1), nil)`,
	)

}

// Test Len(float32(1), 1.5)
func TestCheckBuiltinLenFloat32Float(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(float32(1), 1.5)`, env,
		`too many arguments to len: len(float32(1), 1.5)`,
	)

}

// Test Len(float32(1), []int{})
func TestCheckBuiltinLenFloat32Slice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(float32(1), []int{})`, env,
		`too many arguments to len: len(float32(1), composite literal)`,
	)

}

// Test Len(float32(1), map[int]int{})
func TestCheckBuiltinLenFloat32Map(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(float32(1), map[int]int{})`, env,
		`too many arguments to len: len(float32(1), composite literal)`,
	)

}

// Test Len(float32(1), int)
func TestCheckBuiltinLenFloat32Type(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(float32(1), int)`, env,
		`too many arguments to len: len(float32(1), int)`,
	)

}

// Test Len(float32(1), map[int]int)
func TestCheckBuiltinLenFloat32MakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(float32(1), map[int]int)`, env,
		`too many arguments to len: len(float32(1), map[int]int)`,
	)

}

// Test Len(float32(1), 1, 1)
func TestCheckBuiltinLenFloat32Double(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(float32(1), 1, 1)`, env,
		`too many arguments to len: len(float32(1), 1, 1)`,
	)

}

// Test Len(float32(1), []int{1,2}...)
func TestCheckBuiltinLenFloat32Ellipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(float32(1), []int{1,2}...)`, env,
		`invalid use of ... with builtin len`,
		`too many arguments to len: len(float32(1), composite literal...)`,
	)

}

// Test Len("abc")
func TestCheckBuiltinLenStringX(t *testing.T) {
	env := MakeSimpleEnv()
	expectConst(t, `len("abc")`, env, len("abc"), reflect.TypeOf(len("abc")))

}

// Test Len("abc", 1)
func TestCheckBuiltinLenStringInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len("abc", 1)`, env,
		`too many arguments to len: len("abc", 1)`,
	)

}

// Test Len("abc", float32(1))
func TestCheckBuiltinLenStringFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len("abc", float32(1))`, env,
		`too many arguments to len: len("abc", float32(1))`,
	)

}

// Test Len("abc", "abc")
func TestCheckBuiltinLenStringString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len("abc", "abc")`, env,
		`too many arguments to len: len("abc", "abc")`,
	)

}

// Test Len("abc", nil)
func TestCheckBuiltinLenStringNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len("abc", nil)`, env,
		`too many arguments to len: len("abc", nil)`,
	)

}

// Test Len("abc", 1.5)
func TestCheckBuiltinLenStringFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len("abc", 1.5)`, env,
		`too many arguments to len: len("abc", 1.5)`,
	)

}

// Test Len("abc", []int{})
func TestCheckBuiltinLenStringSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len("abc", []int{})`, env,
		`too many arguments to len: len("abc", composite literal)`,
	)

}

// Test Len("abc", map[int]int{})
func TestCheckBuiltinLenStringMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len("abc", map[int]int{})`, env,
		`too many arguments to len: len("abc", composite literal)`,
	)

}

// Test Len("abc", int)
func TestCheckBuiltinLenStringType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len("abc", int)`, env,
		`too many arguments to len: len("abc", int)`,
	)

}

// Test Len("abc", map[int]int)
func TestCheckBuiltinLenStringMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len("abc", map[int]int)`, env,
		`too many arguments to len: len("abc", map[int]int)`,
	)

}

// Test Len("abc", 1, 1)
func TestCheckBuiltinLenStringDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len("abc", 1, 1)`, env,
		`too many arguments to len: len("abc", 1, 1)`,
	)

}

// Test Len("abc", []int{1,2}...)
func TestCheckBuiltinLenStringEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len("abc", []int{1,2}...)`, env,
		`invalid use of ... with builtin len`,
		`too many arguments to len: len("abc", composite literal...)`,
	)

}

// Test Len(nil)
func TestCheckBuiltinLenNilX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(nil)`, env,
		`use of untyped nil`,
	)

}

// Test Len(nil, 1)
func TestCheckBuiltinLenNilInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(nil, 1)`, env,
		`too many arguments to len: len(nil, 1)`,
	)

}

// Test Len(nil, float32(1))
func TestCheckBuiltinLenNilFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(nil, float32(1))`, env,
		`too many arguments to len: len(nil, float32(1))`,
	)

}

// Test Len(nil, "abc")
func TestCheckBuiltinLenNilString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(nil, "abc")`, env,
		`too many arguments to len: len(nil, "abc")`,
	)

}

// Test Len(nil, nil)
func TestCheckBuiltinLenNilNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(nil, nil)`, env,
		`too many arguments to len: len(nil, nil)`,
	)

}

// Test Len(nil, 1.5)
func TestCheckBuiltinLenNilFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(nil, 1.5)`, env,
		`too many arguments to len: len(nil, 1.5)`,
	)

}

// Test Len(nil, []int{})
func TestCheckBuiltinLenNilSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(nil, []int{})`, env,
		`too many arguments to len: len(nil, composite literal)`,
	)

}

// Test Len(nil, map[int]int{})
func TestCheckBuiltinLenNilMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(nil, map[int]int{})`, env,
		`too many arguments to len: len(nil, composite literal)`,
	)

}

// Test Len(nil, int)
func TestCheckBuiltinLenNilType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(nil, int)`, env,
		`too many arguments to len: len(nil, int)`,
	)

}

// Test Len(nil, map[int]int)
func TestCheckBuiltinLenNilMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(nil, map[int]int)`, env,
		`too many arguments to len: len(nil, map[int]int)`,
	)

}

// Test Len(nil, 1, 1)
func TestCheckBuiltinLenNilDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(nil, 1, 1)`, env,
		`too many arguments to len: len(nil, 1, 1)`,
	)

}

// Test Len(nil, []int{1,2}...)
func TestCheckBuiltinLenNilEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(nil, []int{1,2}...)`, env,
		`invalid use of ... with builtin len`,
		`too many arguments to len: len(nil, composite literal...)`,
	)

}

// Test Len(1.5)
func TestCheckBuiltinLenFloatX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1.5)`, env,
		`invalid argument 1.5 (type float64) for len`,
	)

}

// Test Len(1.5, 1)
func TestCheckBuiltinLenFloatInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1.5, 1)`, env,
		`too many arguments to len: len(1.5, 1)`,
	)

}

// Test Len(1.5, float32(1))
func TestCheckBuiltinLenFloatFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1.5, float32(1))`, env,
		`too many arguments to len: len(1.5, float32(1))`,
	)

}

// Test Len(1.5, "abc")
func TestCheckBuiltinLenFloatString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1.5, "abc")`, env,
		`too many arguments to len: len(1.5, "abc")`,
	)

}

// Test Len(1.5, nil)
func TestCheckBuiltinLenFloatNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1.5, nil)`, env,
		`too many arguments to len: len(1.5, nil)`,
	)

}

// Test Len(1.5, 1.5)
func TestCheckBuiltinLenFloatFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1.5, 1.5)`, env,
		`too many arguments to len: len(1.5, 1.5)`,
	)

}

// Test Len(1.5, []int{})
func TestCheckBuiltinLenFloatSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1.5, []int{})`, env,
		`too many arguments to len: len(1.5, composite literal)`,
	)

}

// Test Len(1.5, map[int]int{})
func TestCheckBuiltinLenFloatMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1.5, map[int]int{})`, env,
		`too many arguments to len: len(1.5, composite literal)`,
	)

}

// Test Len(1.5, int)
func TestCheckBuiltinLenFloatType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1.5, int)`, env,
		`too many arguments to len: len(1.5, int)`,
	)

}

// Test Len(1.5, map[int]int)
func TestCheckBuiltinLenFloatMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1.5, map[int]int)`, env,
		`too many arguments to len: len(1.5, map[int]int)`,
	)

}

// Test Len(1.5, 1, 1)
func TestCheckBuiltinLenFloatDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1.5, 1, 1)`, env,
		`too many arguments to len: len(1.5, 1, 1)`,
	)

}

// Test Len(1.5, []int{1,2}...)
func TestCheckBuiltinLenFloatEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(1.5, []int{1,2}...)`, env,
		`invalid use of ... with builtin len`,
		`too many arguments to len: len(1.5, composite literal...)`,
	)

}

// Test Len([]int{})
func TestCheckBuiltinLenSliceX(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `len([]int{})`, env, reflect.TypeOf(len([]int{})))
}

// Test Len([]int{}, 1)
func TestCheckBuiltinLenSliceInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len([]int{}, 1)`, env,
		`too many arguments to len: len(composite literal, 1)`,
	)

}

// Test Len([]int{}, float32(1))
func TestCheckBuiltinLenSliceFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len([]int{}, float32(1))`, env,
		`too many arguments to len: len(composite literal, float32(1))`,
	)

}

// Test Len([]int{}, "abc")
func TestCheckBuiltinLenSliceString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len([]int{}, "abc")`, env,
		`too many arguments to len: len(composite literal, "abc")`,
	)

}

// Test Len([]int{}, nil)
func TestCheckBuiltinLenSliceNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len([]int{}, nil)`, env,
		`too many arguments to len: len(composite literal, nil)`,
	)

}

// Test Len([]int{}, 1.5)
func TestCheckBuiltinLenSliceFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len([]int{}, 1.5)`, env,
		`too many arguments to len: len(composite literal, 1.5)`,
	)

}

// Test Len([]int{}, []int{})
func TestCheckBuiltinLenSliceSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len([]int{}, []int{})`, env,
		`too many arguments to len: len(composite literal, composite literal)`,
	)

}

// Test Len([]int{}, map[int]int{})
func TestCheckBuiltinLenSliceMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len([]int{}, map[int]int{})`, env,
		`too many arguments to len: len(composite literal, composite literal)`,
	)

}

// Test Len([]int{}, int)
func TestCheckBuiltinLenSliceType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len([]int{}, int)`, env,
		`too many arguments to len: len(composite literal, int)`,
	)

}

// Test Len([]int{}, map[int]int)
func TestCheckBuiltinLenSliceMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len([]int{}, map[int]int)`, env,
		`too many arguments to len: len(composite literal, map[int]int)`,
	)

}

// Test Len([]int{}, 1, 1)
func TestCheckBuiltinLenSliceDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len([]int{}, 1, 1)`, env,
		`too many arguments to len: len(composite literal, 1, 1)`,
	)

}

// Test Len([]int{}, []int{1,2}...)
func TestCheckBuiltinLenSliceEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len([]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin len`,
		`too many arguments to len: len(composite literal, composite literal...)`,
	)

}

// Test Len(map[int]int{})
func TestCheckBuiltinLenMapX(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `len(map[int]int{})`, env, reflect.TypeOf(len(map[int]int{})))
}

// Test Len(map[int]int{}, 1)
func TestCheckBuiltinLenMapInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int{}, 1)`, env,
		`too many arguments to len: len(composite literal, 1)`,
	)

}

// Test Len(map[int]int{}, float32(1))
func TestCheckBuiltinLenMapFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int{}, float32(1))`, env,
		`too many arguments to len: len(composite literal, float32(1))`,
	)

}

// Test Len(map[int]int{}, "abc")
func TestCheckBuiltinLenMapString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int{}, "abc")`, env,
		`too many arguments to len: len(composite literal, "abc")`,
	)

}

// Test Len(map[int]int{}, nil)
func TestCheckBuiltinLenMapNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int{}, nil)`, env,
		`too many arguments to len: len(composite literal, nil)`,
	)

}

// Test Len(map[int]int{}, 1.5)
func TestCheckBuiltinLenMapFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int{}, 1.5)`, env,
		`too many arguments to len: len(composite literal, 1.5)`,
	)

}

// Test Len(map[int]int{}, []int{})
func TestCheckBuiltinLenMapSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int{}, []int{})`, env,
		`too many arguments to len: len(composite literal, composite literal)`,
	)

}

// Test Len(map[int]int{}, map[int]int{})
func TestCheckBuiltinLenMapMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int{}, map[int]int{})`, env,
		`too many arguments to len: len(composite literal, composite literal)`,
	)

}

// Test Len(map[int]int{}, int)
func TestCheckBuiltinLenMapType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int{}, int)`, env,
		`too many arguments to len: len(composite literal, int)`,
	)

}

// Test Len(map[int]int{}, map[int]int)
func TestCheckBuiltinLenMapMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int{}, map[int]int)`, env,
		`too many arguments to len: len(composite literal, map[int]int)`,
	)

}

// Test Len(map[int]int{}, 1, 1)
func TestCheckBuiltinLenMapDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int{}, 1, 1)`, env,
		`too many arguments to len: len(composite literal, 1, 1)`,
	)

}

// Test Len(map[int]int{}, []int{1,2}...)
func TestCheckBuiltinLenMapEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin len`,
		`too many arguments to len: len(composite literal, composite literal...)`,
	)

}

// Test Len(int)
func TestCheckBuiltinLenTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(int)`, env,
		`type int is not an expression`,
	)

}

// Test Len(int, 1)
func TestCheckBuiltinLenTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(int, 1)`, env,
		`too many arguments to len: len(int, 1)`,
	)

}

// Test Len(int, float32(1))
func TestCheckBuiltinLenTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(int, float32(1))`, env,
		`too many arguments to len: len(int, float32(1))`,
	)

}

// Test Len(int, "abc")
func TestCheckBuiltinLenTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(int, "abc")`, env,
		`too many arguments to len: len(int, "abc")`,
	)

}

// Test Len(int, nil)
func TestCheckBuiltinLenTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(int, nil)`, env,
		`too many arguments to len: len(int, nil)`,
	)

}

// Test Len(int, 1.5)
func TestCheckBuiltinLenTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(int, 1.5)`, env,
		`too many arguments to len: len(int, 1.5)`,
	)

}

// Test Len(int, []int{})
func TestCheckBuiltinLenTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(int, []int{})`, env,
		`too many arguments to len: len(int, composite literal)`,
	)

}

// Test Len(int, map[int]int{})
func TestCheckBuiltinLenTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(int, map[int]int{})`, env,
		`too many arguments to len: len(int, composite literal)`,
	)

}

// Test Len(int, int)
func TestCheckBuiltinLenTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(int, int)`, env,
		`too many arguments to len: len(int, int)`,
	)

}

// Test Len(int, map[int]int)
func TestCheckBuiltinLenTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(int, map[int]int)`, env,
		`too many arguments to len: len(int, map[int]int)`,
	)

}

// Test Len(int, 1, 1)
func TestCheckBuiltinLenTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(int, 1, 1)`, env,
		`too many arguments to len: len(int, 1, 1)`,
	)

}

// Test Len(int, []int{1,2}...)
func TestCheckBuiltinLenTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(int, []int{1,2}...)`, env,
		`invalid use of ... with builtin len`,
		`too many arguments to len: len(int, composite literal...)`,
	)

}

// Test Len(map[int]int)
func TestCheckBuiltinLenMakeTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Len(map[int]int, 1)
func TestCheckBuiltinLenMakeTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int, 1)`, env,
		`too many arguments to len: len(map[int]int, 1)`,
	)

}

// Test Len(map[int]int, float32(1))
func TestCheckBuiltinLenMakeTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int, float32(1))`, env,
		`too many arguments to len: len(map[int]int, float32(1))`,
	)

}

// Test Len(map[int]int, "abc")
func TestCheckBuiltinLenMakeTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int, "abc")`, env,
		`too many arguments to len: len(map[int]int, "abc")`,
	)

}

// Test Len(map[int]int, nil)
func TestCheckBuiltinLenMakeTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int, nil)`, env,
		`too many arguments to len: len(map[int]int, nil)`,
	)

}

// Test Len(map[int]int, 1.5)
func TestCheckBuiltinLenMakeTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int, 1.5)`, env,
		`too many arguments to len: len(map[int]int, 1.5)`,
	)

}

// Test Len(map[int]int, []int{})
func TestCheckBuiltinLenMakeTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int, []int{})`, env,
		`too many arguments to len: len(map[int]int, composite literal)`,
	)

}

// Test Len(map[int]int, map[int]int{})
func TestCheckBuiltinLenMakeTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int, map[int]int{})`, env,
		`too many arguments to len: len(map[int]int, composite literal)`,
	)

}

// Test Len(map[int]int, int)
func TestCheckBuiltinLenMakeTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int, int)`, env,
		`too many arguments to len: len(map[int]int, int)`,
	)

}

// Test Len(map[int]int, map[int]int)
func TestCheckBuiltinLenMakeTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int, map[int]int)`, env,
		`too many arguments to len: len(map[int]int, map[int]int)`,
	)

}

// Test Len(map[int]int, 1, 1)
func TestCheckBuiltinLenMakeTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int, 1, 1)`, env,
		`too many arguments to len: len(map[int]int, 1, 1)`,
	)

}

// Test Len(map[int]int, []int{1,2}...)
func TestCheckBuiltinLenMakeTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `len(map[int]int, []int{1,2}...)`, env,
		`invalid use of ... with builtin len`,
		`too many arguments to len: len(map[int]int, composite literal...)`,
	)

}

// Test Cap()
func TestCheckBuiltinCapXX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap()`, env,
		`missing argument to cap: cap()`,
	)

}

// Test Cap(1)
func TestCheckBuiltinCapXInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1)`, env,
		`invalid argument 1 (type int) for cap`,
	)

}

// Test Cap(float32(1))
func TestCheckBuiltinCapXFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(float32(1))`, env,
		`invalid argument float32(1) (type float32) for cap`,
	)

}

// Test Cap("abc")
func TestCheckBuiltinCapXString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap("abc")`, env,
		`invalid argument "abc" (type string) for cap`,
	)

}

// Test Cap(nil)
func TestCheckBuiltinCapXNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(nil)`, env,
		`use of untyped nil`,
	)

}

// Test Cap(1.5)
func TestCheckBuiltinCapXFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1.5)`, env,
		`invalid argument 1.5 (type float64) for cap`,
	)

}

// Test Cap([]int{})
func TestCheckBuiltinCapXSlice(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `cap([]int{})`, env, reflect.TypeOf(cap([]int{})))
}

// Test Cap(map[int]int{})
func TestCheckBuiltinCapXMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int{})`, env,
		`invalid argument map[int]int literal (type map[int]int) for cap`,
	)

}

// Test Cap(int)
func TestCheckBuiltinCapXType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(int)`, env,
		`type int is not an expression`,
	)

}

// Test Cap(map[int]int)
func TestCheckBuiltinCapXMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Cap(1, 1)
func TestCheckBuiltinCapXDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1, 1)`, env,
		`too many arguments to cap: cap(1, 1)`,
	)

}

// Test Cap([]int{1,2}...)
func TestCheckBuiltinCapXEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap([]int{1,2}...)`, env,
		`invalid use of ... with builtin cap`,
	)

}

// Test Cap(1)
func TestCheckBuiltinCapIntX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1)`, env,
		`invalid argument 1 (type int) for cap`,
	)

}

// Test Cap(1, 1)
func TestCheckBuiltinCapIntInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1, 1)`, env,
		`too many arguments to cap: cap(1, 1)`,
	)

}

// Test Cap(1, float32(1))
func TestCheckBuiltinCapIntFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1, float32(1))`, env,
		`too many arguments to cap: cap(1, float32(1))`,
	)

}

// Test Cap(1, "abc")
func TestCheckBuiltinCapIntString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1, "abc")`, env,
		`too many arguments to cap: cap(1, "abc")`,
	)

}

// Test Cap(1, nil)
func TestCheckBuiltinCapIntNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1, nil)`, env,
		`too many arguments to cap: cap(1, nil)`,
	)

}

// Test Cap(1, 1.5)
func TestCheckBuiltinCapIntFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1, 1.5)`, env,
		`too many arguments to cap: cap(1, 1.5)`,
	)

}

// Test Cap(1, []int{})
func TestCheckBuiltinCapIntSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1, []int{})`, env,
		`too many arguments to cap: cap(1, composite literal)`,
	)

}

// Test Cap(1, map[int]int{})
func TestCheckBuiltinCapIntMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1, map[int]int{})`, env,
		`too many arguments to cap: cap(1, composite literal)`,
	)

}

// Test Cap(1, int)
func TestCheckBuiltinCapIntType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1, int)`, env,
		`too many arguments to cap: cap(1, int)`,
	)

}

// Test Cap(1, map[int]int)
func TestCheckBuiltinCapIntMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1, map[int]int)`, env,
		`too many arguments to cap: cap(1, map[int]int)`,
	)

}

// Test Cap(1, 1, 1)
func TestCheckBuiltinCapIntDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1, 1, 1)`, env,
		`too many arguments to cap: cap(1, 1, 1)`,
	)

}

// Test Cap(1, []int{1,2}...)
func TestCheckBuiltinCapIntEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1, []int{1,2}...)`, env,
		`invalid use of ... with builtin cap`,
		`too many arguments to cap: cap(1, composite literal...)`,
	)

}

// Test Cap(float32(1))
func TestCheckBuiltinCapFloat32X(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(float32(1))`, env,
		`invalid argument float32(1) (type float32) for cap`,
	)

}

// Test Cap(float32(1), 1)
func TestCheckBuiltinCapFloat32Int(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(float32(1), 1)`, env,
		`too many arguments to cap: cap(float32(1), 1)`,
	)

}

// Test Cap(float32(1), float32(1))
func TestCheckBuiltinCapFloat32Float32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(float32(1), float32(1))`, env,
		`too many arguments to cap: cap(float32(1), float32(1))`,
	)

}

// Test Cap(float32(1), "abc")
func TestCheckBuiltinCapFloat32String(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(float32(1), "abc")`, env,
		`too many arguments to cap: cap(float32(1), "abc")`,
	)

}

// Test Cap(float32(1), nil)
func TestCheckBuiltinCapFloat32Nil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(float32(1), nil)`, env,
		`too many arguments to cap: cap(float32(1), nil)`,
	)

}

// Test Cap(float32(1), 1.5)
func TestCheckBuiltinCapFloat32Float(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(float32(1), 1.5)`, env,
		`too many arguments to cap: cap(float32(1), 1.5)`,
	)

}

// Test Cap(float32(1), []int{})
func TestCheckBuiltinCapFloat32Slice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(float32(1), []int{})`, env,
		`too many arguments to cap: cap(float32(1), composite literal)`,
	)

}

// Test Cap(float32(1), map[int]int{})
func TestCheckBuiltinCapFloat32Map(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(float32(1), map[int]int{})`, env,
		`too many arguments to cap: cap(float32(1), composite literal)`,
	)

}

// Test Cap(float32(1), int)
func TestCheckBuiltinCapFloat32Type(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(float32(1), int)`, env,
		`too many arguments to cap: cap(float32(1), int)`,
	)

}

// Test Cap(float32(1), map[int]int)
func TestCheckBuiltinCapFloat32MakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(float32(1), map[int]int)`, env,
		`too many arguments to cap: cap(float32(1), map[int]int)`,
	)

}

// Test Cap(float32(1), 1, 1)
func TestCheckBuiltinCapFloat32Double(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(float32(1), 1, 1)`, env,
		`too many arguments to cap: cap(float32(1), 1, 1)`,
	)

}

// Test Cap(float32(1), []int{1,2}...)
func TestCheckBuiltinCapFloat32Ellipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(float32(1), []int{1,2}...)`, env,
		`invalid use of ... with builtin cap`,
		`too many arguments to cap: cap(float32(1), composite literal...)`,
	)

}

// Test Cap("abc")
func TestCheckBuiltinCapStringX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap("abc")`, env,
		`invalid argument "abc" (type string) for cap`,
	)

}

// Test Cap("abc", 1)
func TestCheckBuiltinCapStringInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap("abc", 1)`, env,
		`too many arguments to cap: cap("abc", 1)`,
	)

}

// Test Cap("abc", float32(1))
func TestCheckBuiltinCapStringFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap("abc", float32(1))`, env,
		`too many arguments to cap: cap("abc", float32(1))`,
	)

}

// Test Cap("abc", "abc")
func TestCheckBuiltinCapStringString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap("abc", "abc")`, env,
		`too many arguments to cap: cap("abc", "abc")`,
	)

}

// Test Cap("abc", nil)
func TestCheckBuiltinCapStringNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap("abc", nil)`, env,
		`too many arguments to cap: cap("abc", nil)`,
	)

}

// Test Cap("abc", 1.5)
func TestCheckBuiltinCapStringFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap("abc", 1.5)`, env,
		`too many arguments to cap: cap("abc", 1.5)`,
	)

}

// Test Cap("abc", []int{})
func TestCheckBuiltinCapStringSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap("abc", []int{})`, env,
		`too many arguments to cap: cap("abc", composite literal)`,
	)

}

// Test Cap("abc", map[int]int{})
func TestCheckBuiltinCapStringMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap("abc", map[int]int{})`, env,
		`too many arguments to cap: cap("abc", composite literal)`,
	)

}

// Test Cap("abc", int)
func TestCheckBuiltinCapStringType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap("abc", int)`, env,
		`too many arguments to cap: cap("abc", int)`,
	)

}

// Test Cap("abc", map[int]int)
func TestCheckBuiltinCapStringMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap("abc", map[int]int)`, env,
		`too many arguments to cap: cap("abc", map[int]int)`,
	)

}

// Test Cap("abc", 1, 1)
func TestCheckBuiltinCapStringDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap("abc", 1, 1)`, env,
		`too many arguments to cap: cap("abc", 1, 1)`,
	)

}

// Test Cap("abc", []int{1,2}...)
func TestCheckBuiltinCapStringEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap("abc", []int{1,2}...)`, env,
		`invalid use of ... with builtin cap`,
		`too many arguments to cap: cap("abc", composite literal...)`,
	)

}

// Test Cap(nil)
func TestCheckBuiltinCapNilX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(nil)`, env,
		`use of untyped nil`,
	)

}

// Test Cap(nil, 1)
func TestCheckBuiltinCapNilInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(nil, 1)`, env,
		`too many arguments to cap: cap(nil, 1)`,
	)

}

// Test Cap(nil, float32(1))
func TestCheckBuiltinCapNilFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(nil, float32(1))`, env,
		`too many arguments to cap: cap(nil, float32(1))`,
	)

}

// Test Cap(nil, "abc")
func TestCheckBuiltinCapNilString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(nil, "abc")`, env,
		`too many arguments to cap: cap(nil, "abc")`,
	)

}

// Test Cap(nil, nil)
func TestCheckBuiltinCapNilNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(nil, nil)`, env,
		`too many arguments to cap: cap(nil, nil)`,
	)

}

// Test Cap(nil, 1.5)
func TestCheckBuiltinCapNilFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(nil, 1.5)`, env,
		`too many arguments to cap: cap(nil, 1.5)`,
	)

}

// Test Cap(nil, []int{})
func TestCheckBuiltinCapNilSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(nil, []int{})`, env,
		`too many arguments to cap: cap(nil, composite literal)`,
	)

}

// Test Cap(nil, map[int]int{})
func TestCheckBuiltinCapNilMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(nil, map[int]int{})`, env,
		`too many arguments to cap: cap(nil, composite literal)`,
	)

}

// Test Cap(nil, int)
func TestCheckBuiltinCapNilType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(nil, int)`, env,
		`too many arguments to cap: cap(nil, int)`,
	)

}

// Test Cap(nil, map[int]int)
func TestCheckBuiltinCapNilMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(nil, map[int]int)`, env,
		`too many arguments to cap: cap(nil, map[int]int)`,
	)

}

// Test Cap(nil, 1, 1)
func TestCheckBuiltinCapNilDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(nil, 1, 1)`, env,
		`too many arguments to cap: cap(nil, 1, 1)`,
	)

}

// Test Cap(nil, []int{1,2}...)
func TestCheckBuiltinCapNilEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(nil, []int{1,2}...)`, env,
		`invalid use of ... with builtin cap`,
		`too many arguments to cap: cap(nil, composite literal...)`,
	)

}

// Test Cap(1.5)
func TestCheckBuiltinCapFloatX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1.5)`, env,
		`invalid argument 1.5 (type float64) for cap`,
	)

}

// Test Cap(1.5, 1)
func TestCheckBuiltinCapFloatInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1.5, 1)`, env,
		`too many arguments to cap: cap(1.5, 1)`,
	)

}

// Test Cap(1.5, float32(1))
func TestCheckBuiltinCapFloatFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1.5, float32(1))`, env,
		`too many arguments to cap: cap(1.5, float32(1))`,
	)

}

// Test Cap(1.5, "abc")
func TestCheckBuiltinCapFloatString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1.5, "abc")`, env,
		`too many arguments to cap: cap(1.5, "abc")`,
	)

}

// Test Cap(1.5, nil)
func TestCheckBuiltinCapFloatNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1.5, nil)`, env,
		`too many arguments to cap: cap(1.5, nil)`,
	)

}

// Test Cap(1.5, 1.5)
func TestCheckBuiltinCapFloatFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1.5, 1.5)`, env,
		`too many arguments to cap: cap(1.5, 1.5)`,
	)

}

// Test Cap(1.5, []int{})
func TestCheckBuiltinCapFloatSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1.5, []int{})`, env,
		`too many arguments to cap: cap(1.5, composite literal)`,
	)

}

// Test Cap(1.5, map[int]int{})
func TestCheckBuiltinCapFloatMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1.5, map[int]int{})`, env,
		`too many arguments to cap: cap(1.5, composite literal)`,
	)

}

// Test Cap(1.5, int)
func TestCheckBuiltinCapFloatType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1.5, int)`, env,
		`too many arguments to cap: cap(1.5, int)`,
	)

}

// Test Cap(1.5, map[int]int)
func TestCheckBuiltinCapFloatMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1.5, map[int]int)`, env,
		`too many arguments to cap: cap(1.5, map[int]int)`,
	)

}

// Test Cap(1.5, 1, 1)
func TestCheckBuiltinCapFloatDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1.5, 1, 1)`, env,
		`too many arguments to cap: cap(1.5, 1, 1)`,
	)

}

// Test Cap(1.5, []int{1,2}...)
func TestCheckBuiltinCapFloatEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(1.5, []int{1,2}...)`, env,
		`invalid use of ... with builtin cap`,
		`too many arguments to cap: cap(1.5, composite literal...)`,
	)

}

// Test Cap([]int{})
func TestCheckBuiltinCapSliceX(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `cap([]int{})`, env, reflect.TypeOf(cap([]int{})))
}

// Test Cap([]int{}, 1)
func TestCheckBuiltinCapSliceInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap([]int{}, 1)`, env,
		`too many arguments to cap: cap(composite literal, 1)`,
	)

}

// Test Cap([]int{}, float32(1))
func TestCheckBuiltinCapSliceFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap([]int{}, float32(1))`, env,
		`too many arguments to cap: cap(composite literal, float32(1))`,
	)

}

// Test Cap([]int{}, "abc")
func TestCheckBuiltinCapSliceString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap([]int{}, "abc")`, env,
		`too many arguments to cap: cap(composite literal, "abc")`,
	)

}

// Test Cap([]int{}, nil)
func TestCheckBuiltinCapSliceNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap([]int{}, nil)`, env,
		`too many arguments to cap: cap(composite literal, nil)`,
	)

}

// Test Cap([]int{}, 1.5)
func TestCheckBuiltinCapSliceFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap([]int{}, 1.5)`, env,
		`too many arguments to cap: cap(composite literal, 1.5)`,
	)

}

// Test Cap([]int{}, []int{})
func TestCheckBuiltinCapSliceSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap([]int{}, []int{})`, env,
		`too many arguments to cap: cap(composite literal, composite literal)`,
	)

}

// Test Cap([]int{}, map[int]int{})
func TestCheckBuiltinCapSliceMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap([]int{}, map[int]int{})`, env,
		`too many arguments to cap: cap(composite literal, composite literal)`,
	)

}

// Test Cap([]int{}, int)
func TestCheckBuiltinCapSliceType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap([]int{}, int)`, env,
		`too many arguments to cap: cap(composite literal, int)`,
	)

}

// Test Cap([]int{}, map[int]int)
func TestCheckBuiltinCapSliceMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap([]int{}, map[int]int)`, env,
		`too many arguments to cap: cap(composite literal, map[int]int)`,
	)

}

// Test Cap([]int{}, 1, 1)
func TestCheckBuiltinCapSliceDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap([]int{}, 1, 1)`, env,
		`too many arguments to cap: cap(composite literal, 1, 1)`,
	)

}

// Test Cap([]int{}, []int{1,2}...)
func TestCheckBuiltinCapSliceEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap([]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin cap`,
		`too many arguments to cap: cap(composite literal, composite literal...)`,
	)

}

// Test Cap(map[int]int{})
func TestCheckBuiltinCapMapX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int{})`, env,
		`invalid argument map[int]int literal (type map[int]int) for cap`,
	)

}

// Test Cap(map[int]int{}, 1)
func TestCheckBuiltinCapMapInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int{}, 1)`, env,
		`too many arguments to cap: cap(composite literal, 1)`,
	)

}

// Test Cap(map[int]int{}, float32(1))
func TestCheckBuiltinCapMapFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int{}, float32(1))`, env,
		`too many arguments to cap: cap(composite literal, float32(1))`,
	)

}

// Test Cap(map[int]int{}, "abc")
func TestCheckBuiltinCapMapString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int{}, "abc")`, env,
		`too many arguments to cap: cap(composite literal, "abc")`,
	)

}

// Test Cap(map[int]int{}, nil)
func TestCheckBuiltinCapMapNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int{}, nil)`, env,
		`too many arguments to cap: cap(composite literal, nil)`,
	)

}

// Test Cap(map[int]int{}, 1.5)
func TestCheckBuiltinCapMapFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int{}, 1.5)`, env,
		`too many arguments to cap: cap(composite literal, 1.5)`,
	)

}

// Test Cap(map[int]int{}, []int{})
func TestCheckBuiltinCapMapSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int{}, []int{})`, env,
		`too many arguments to cap: cap(composite literal, composite literal)`,
	)

}

// Test Cap(map[int]int{}, map[int]int{})
func TestCheckBuiltinCapMapMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int{}, map[int]int{})`, env,
		`too many arguments to cap: cap(composite literal, composite literal)`,
	)

}

// Test Cap(map[int]int{}, int)
func TestCheckBuiltinCapMapType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int{}, int)`, env,
		`too many arguments to cap: cap(composite literal, int)`,
	)

}

// Test Cap(map[int]int{}, map[int]int)
func TestCheckBuiltinCapMapMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int{}, map[int]int)`, env,
		`too many arguments to cap: cap(composite literal, map[int]int)`,
	)

}

// Test Cap(map[int]int{}, 1, 1)
func TestCheckBuiltinCapMapDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int{}, 1, 1)`, env,
		`too many arguments to cap: cap(composite literal, 1, 1)`,
	)

}

// Test Cap(map[int]int{}, []int{1,2}...)
func TestCheckBuiltinCapMapEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin cap`,
		`too many arguments to cap: cap(composite literal, composite literal...)`,
	)

}

// Test Cap(int)
func TestCheckBuiltinCapTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(int)`, env,
		`type int is not an expression`,
	)

}

// Test Cap(int, 1)
func TestCheckBuiltinCapTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(int, 1)`, env,
		`too many arguments to cap: cap(int, 1)`,
	)

}

// Test Cap(int, float32(1))
func TestCheckBuiltinCapTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(int, float32(1))`, env,
		`too many arguments to cap: cap(int, float32(1))`,
	)

}

// Test Cap(int, "abc")
func TestCheckBuiltinCapTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(int, "abc")`, env,
		`too many arguments to cap: cap(int, "abc")`,
	)

}

// Test Cap(int, nil)
func TestCheckBuiltinCapTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(int, nil)`, env,
		`too many arguments to cap: cap(int, nil)`,
	)

}

// Test Cap(int, 1.5)
func TestCheckBuiltinCapTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(int, 1.5)`, env,
		`too many arguments to cap: cap(int, 1.5)`,
	)

}

// Test Cap(int, []int{})
func TestCheckBuiltinCapTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(int, []int{})`, env,
		`too many arguments to cap: cap(int, composite literal)`,
	)

}

// Test Cap(int, map[int]int{})
func TestCheckBuiltinCapTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(int, map[int]int{})`, env,
		`too many arguments to cap: cap(int, composite literal)`,
	)

}

// Test Cap(int, int)
func TestCheckBuiltinCapTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(int, int)`, env,
		`too many arguments to cap: cap(int, int)`,
	)

}

// Test Cap(int, map[int]int)
func TestCheckBuiltinCapTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(int, map[int]int)`, env,
		`too many arguments to cap: cap(int, map[int]int)`,
	)

}

// Test Cap(int, 1, 1)
func TestCheckBuiltinCapTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(int, 1, 1)`, env,
		`too many arguments to cap: cap(int, 1, 1)`,
	)

}

// Test Cap(int, []int{1,2}...)
func TestCheckBuiltinCapTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(int, []int{1,2}...)`, env,
		`invalid use of ... with builtin cap`,
		`too many arguments to cap: cap(int, composite literal...)`,
	)

}

// Test Cap(map[int]int)
func TestCheckBuiltinCapMakeTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Cap(map[int]int, 1)
func TestCheckBuiltinCapMakeTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int, 1)`, env,
		`too many arguments to cap: cap(map[int]int, 1)`,
	)

}

// Test Cap(map[int]int, float32(1))
func TestCheckBuiltinCapMakeTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int, float32(1))`, env,
		`too many arguments to cap: cap(map[int]int, float32(1))`,
	)

}

// Test Cap(map[int]int, "abc")
func TestCheckBuiltinCapMakeTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int, "abc")`, env,
		`too many arguments to cap: cap(map[int]int, "abc")`,
	)

}

// Test Cap(map[int]int, nil)
func TestCheckBuiltinCapMakeTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int, nil)`, env,
		`too many arguments to cap: cap(map[int]int, nil)`,
	)

}

// Test Cap(map[int]int, 1.5)
func TestCheckBuiltinCapMakeTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int, 1.5)`, env,
		`too many arguments to cap: cap(map[int]int, 1.5)`,
	)

}

// Test Cap(map[int]int, []int{})
func TestCheckBuiltinCapMakeTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int, []int{})`, env,
		`too many arguments to cap: cap(map[int]int, composite literal)`,
	)

}

// Test Cap(map[int]int, map[int]int{})
func TestCheckBuiltinCapMakeTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int, map[int]int{})`, env,
		`too many arguments to cap: cap(map[int]int, composite literal)`,
	)

}

// Test Cap(map[int]int, int)
func TestCheckBuiltinCapMakeTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int, int)`, env,
		`too many arguments to cap: cap(map[int]int, int)`,
	)

}

// Test Cap(map[int]int, map[int]int)
func TestCheckBuiltinCapMakeTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int, map[int]int)`, env,
		`too many arguments to cap: cap(map[int]int, map[int]int)`,
	)

}

// Test Cap(map[int]int, 1, 1)
func TestCheckBuiltinCapMakeTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int, 1, 1)`, env,
		`too many arguments to cap: cap(map[int]int, 1, 1)`,
	)

}

// Test Cap(map[int]int, []int{1,2}...)
func TestCheckBuiltinCapMakeTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `cap(map[int]int, []int{1,2}...)`, env,
		`invalid use of ... with builtin cap`,
		`too many arguments to cap: cap(map[int]int, composite literal...)`,
	)

}

// Test Append()
func TestCheckBuiltinAppendXX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append()`, env,
		`missing arguments to append`,
	)

}

// Test Append(1)
func TestCheckBuiltinAppendXInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1)`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(float32(1))
func TestCheckBuiltinAppendXFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(float32(1))`, env,
		`first argument to append must be slice; have float32`,
	)

}

// Test Append("abc")
func TestCheckBuiltinAppendXString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append("abc")`, env,
		`first argument to append must be slice; have untyped string`,
	)

}

// Test Append(nil)
func TestCheckBuiltinAppendXNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(nil)`, env,
		`first argument to append must be typed slice; have untyped nil`,
	)

}

// Test Append(1.5)
func TestCheckBuiltinAppendXFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1.5)`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append([]int{})
func TestCheckBuiltinAppendXSlice(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `append([]int{})`, env, reflect.TypeOf(append([]int{})))
}

// Test Append(map[int]int{})
func TestCheckBuiltinAppendXMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int{})`, env,
		`first argument to append must be slice; have map[int]int`,
	)

}

// Test Append(int)
func TestCheckBuiltinAppendXType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(int)`, env,
		`type int is not an expression`,
	)

}

// Test Append(map[int]int)
func TestCheckBuiltinAppendXMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Append(1, 1)
func TestCheckBuiltinAppendXDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1, 1)`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append([]int{1,2}...)
func TestCheckBuiltinAppendXEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append([]int{1,2}...)`, env,
		`cannot use ... on first argument to append`,
	)

}

// Test Append(1)
func TestCheckBuiltinAppendIntX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1)`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1, 1)
func TestCheckBuiltinAppendIntInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1, 1)`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1, float32(1))
func TestCheckBuiltinAppendIntFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1, float32(1))`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1, "abc")
func TestCheckBuiltinAppendIntString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1, "abc")`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1, nil)
func TestCheckBuiltinAppendIntNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1, nil)`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1, 1.5)
func TestCheckBuiltinAppendIntFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1, 1.5)`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1, []int{})
func TestCheckBuiltinAppendIntSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1, []int{})`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1, map[int]int{})
func TestCheckBuiltinAppendIntMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1, map[int]int{})`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1, int)
func TestCheckBuiltinAppendIntType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1, int)`, env,
		`type int is not an expression`,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1, map[int]int)
func TestCheckBuiltinAppendIntMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1, map[int]int)`, env,
		`type map[int]int is not an expression`,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1, 1, 1)
func TestCheckBuiltinAppendIntDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1, 1, 1)`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1, []int{1,2}...)
func TestCheckBuiltinAppendIntEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1, []int{1,2}...)`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(float32(1))
func TestCheckBuiltinAppendFloat32X(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(float32(1))`, env,
		`first argument to append must be slice; have float32`,
	)

}

// Test Append(float32(1), 1)
func TestCheckBuiltinAppendFloat32Int(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(float32(1), 1)`, env,
		`first argument to append must be slice; have float32`,
	)

}

// Test Append(float32(1), float32(1))
func TestCheckBuiltinAppendFloat32Float32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(float32(1), float32(1))`, env,
		`first argument to append must be slice; have float32`,
	)

}

// Test Append(float32(1), "abc")
func TestCheckBuiltinAppendFloat32String(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(float32(1), "abc")`, env,
		`first argument to append must be slice; have float32`,
	)

}

// Test Append(float32(1), nil)
func TestCheckBuiltinAppendFloat32Nil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(float32(1), nil)`, env,
		`first argument to append must be slice; have float32`,
	)

}

// Test Append(float32(1), 1.5)
func TestCheckBuiltinAppendFloat32Float(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(float32(1), 1.5)`, env,
		`first argument to append must be slice; have float32`,
	)

}

// Test Append(float32(1), []int{})
func TestCheckBuiltinAppendFloat32Slice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(float32(1), []int{})`, env,
		`first argument to append must be slice; have float32`,
	)

}

// Test Append(float32(1), map[int]int{})
func TestCheckBuiltinAppendFloat32Map(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(float32(1), map[int]int{})`, env,
		`first argument to append must be slice; have float32`,
	)

}

// Test Append(float32(1), int)
func TestCheckBuiltinAppendFloat32Type(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(float32(1), int)`, env,
		`type int is not an expression`,
		`first argument to append must be slice; have float32`,
	)

}

// Test Append(float32(1), map[int]int)
func TestCheckBuiltinAppendFloat32MakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(float32(1), map[int]int)`, env,
		`type map[int]int is not an expression`,
		`first argument to append must be slice; have float32`,
	)

}

// Test Append(float32(1), 1, 1)
func TestCheckBuiltinAppendFloat32Double(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(float32(1), 1, 1)`, env,
		`first argument to append must be slice; have float32`,
	)

}

// Test Append(float32(1), []int{1,2}...)
func TestCheckBuiltinAppendFloat32Ellipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(float32(1), []int{1,2}...)`, env,
		`first argument to append must be slice; have float32`,
	)

}

// Test Append("abc")
func TestCheckBuiltinAppendStringX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append("abc")`, env,
		`first argument to append must be slice; have untyped string`,
	)

}

// Test Append("abc", 1)
func TestCheckBuiltinAppendStringInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append("abc", 1)`, env,
		`first argument to append must be slice; have untyped string`,
	)

}

// Test Append("abc", float32(1))
func TestCheckBuiltinAppendStringFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append("abc", float32(1))`, env,
		`first argument to append must be slice; have untyped string`,
	)

}

// Test Append("abc", "abc")
func TestCheckBuiltinAppendStringString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append("abc", "abc")`, env,
		`first argument to append must be slice; have untyped string`,
	)

}

// Test Append("abc", nil)
func TestCheckBuiltinAppendStringNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append("abc", nil)`, env,
		`first argument to append must be slice; have untyped string`,
	)

}

// Test Append("abc", 1.5)
func TestCheckBuiltinAppendStringFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append("abc", 1.5)`, env,
		`first argument to append must be slice; have untyped string`,
	)

}

// Test Append("abc", []int{})
func TestCheckBuiltinAppendStringSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append("abc", []int{})`, env,
		`first argument to append must be slice; have untyped string`,
	)

}

// Test Append("abc", map[int]int{})
func TestCheckBuiltinAppendStringMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append("abc", map[int]int{})`, env,
		`first argument to append must be slice; have untyped string`,
	)

}

// Test Append("abc", int)
func TestCheckBuiltinAppendStringType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append("abc", int)`, env,
		`type int is not an expression`,
		`first argument to append must be slice; have untyped string`,
	)

}

// Test Append("abc", map[int]int)
func TestCheckBuiltinAppendStringMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append("abc", map[int]int)`, env,
		`type map[int]int is not an expression`,
		`first argument to append must be slice; have untyped string`,
	)

}

// Test Append("abc", 1, 1)
func TestCheckBuiltinAppendStringDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append("abc", 1, 1)`, env,
		`first argument to append must be slice; have untyped string`,
	)

}

// Test Append("abc", []int{1,2}...)
func TestCheckBuiltinAppendStringEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append("abc", []int{1,2}...)`, env,
		`first argument to append must be slice; have untyped string`,
	)

}

// Test Append(nil)
func TestCheckBuiltinAppendNilX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(nil)`, env,
		`first argument to append must be typed slice; have untyped nil`,
	)

}

// Test Append(nil, 1)
func TestCheckBuiltinAppendNilInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(nil, 1)`, env,
		`first argument to append must be typed slice; have untyped nil`,
	)

}

// Test Append(nil, float32(1))
func TestCheckBuiltinAppendNilFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(nil, float32(1))`, env,
		`first argument to append must be typed slice; have untyped nil`,
	)

}

// Test Append(nil, "abc")
func TestCheckBuiltinAppendNilString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(nil, "abc")`, env,
		`first argument to append must be typed slice; have untyped nil`,
	)

}

// Test Append(nil, nil)
func TestCheckBuiltinAppendNilNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(nil, nil)`, env,
		`first argument to append must be typed slice; have untyped nil`,
	)

}

// Test Append(nil, 1.5)
func TestCheckBuiltinAppendNilFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(nil, 1.5)`, env,
		`first argument to append must be typed slice; have untyped nil`,
	)

}

// Test Append(nil, []int{})
func TestCheckBuiltinAppendNilSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(nil, []int{})`, env,
		`first argument to append must be typed slice; have untyped nil`,
	)

}

// Test Append(nil, map[int]int{})
func TestCheckBuiltinAppendNilMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(nil, map[int]int{})`, env,
		`first argument to append must be typed slice; have untyped nil`,
	)

}

// Test Append(nil, int)
func TestCheckBuiltinAppendNilType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(nil, int)`, env,
		`type int is not an expression`,
		`first argument to append must be typed slice; have untyped nil`,
	)

}

// Test Append(nil, map[int]int)
func TestCheckBuiltinAppendNilMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(nil, map[int]int)`, env,
		`type map[int]int is not an expression`,
		`first argument to append must be typed slice; have untyped nil`,
	)

}

// Test Append(nil, 1, 1)
func TestCheckBuiltinAppendNilDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(nil, 1, 1)`, env,
		`first argument to append must be typed slice; have untyped nil`,
	)

}

// Test Append(nil, []int{1,2}...)
func TestCheckBuiltinAppendNilEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(nil, []int{1,2}...)`, env,
		`first argument to append must be typed slice; have untyped nil`,
	)

}

// Test Append(1.5)
func TestCheckBuiltinAppendFloatX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1.5)`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1.5, 1)
func TestCheckBuiltinAppendFloatInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1.5, 1)`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1.5, float32(1))
func TestCheckBuiltinAppendFloatFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1.5, float32(1))`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1.5, "abc")
func TestCheckBuiltinAppendFloatString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1.5, "abc")`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1.5, nil)
func TestCheckBuiltinAppendFloatNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1.5, nil)`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1.5, 1.5)
func TestCheckBuiltinAppendFloatFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1.5, 1.5)`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1.5, []int{})
func TestCheckBuiltinAppendFloatSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1.5, []int{})`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1.5, map[int]int{})
func TestCheckBuiltinAppendFloatMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1.5, map[int]int{})`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1.5, int)
func TestCheckBuiltinAppendFloatType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1.5, int)`, env,
		`type int is not an expression`,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1.5, map[int]int)
func TestCheckBuiltinAppendFloatMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1.5, map[int]int)`, env,
		`type map[int]int is not an expression`,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1.5, 1, 1)
func TestCheckBuiltinAppendFloatDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1.5, 1, 1)`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append(1.5, []int{1,2}...)
func TestCheckBuiltinAppendFloatEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(1.5, []int{1,2}...)`, env,
		`first argument to append must be slice; have untyped number`,
	)

}

// Test Append([]int{})
func TestCheckBuiltinAppendSliceX(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `append([]int{})`, env, reflect.TypeOf(append([]int{})))
}

// Test Append([]int{}, 1)
func TestCheckBuiltinAppendSliceInt(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `append([]int{}, 1)`, env, reflect.TypeOf(append([]int{}, 1)))
}

// Test Append([]int{}, float32(1))
func TestCheckBuiltinAppendSliceFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append([]int{}, float32(1))`, env,
		`cannot use float32(1) (type float32) as type int in append`,
	)

}

// Test Append([]int{}, "abc")
func TestCheckBuiltinAppendSliceString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append([]int{}, "abc")`, env,
		`cannot use "abc" (type string) as type int in append`,
	)

}

// Test Append([]int{}, nil)
func TestCheckBuiltinAppendSliceNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append([]int{}, nil)`, env,
		`cannot use nil as type int in append`,
	)

}

// Test Append([]int{}, 1.5)
func TestCheckBuiltinAppendSliceFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append([]int{}, 1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Append([]int{}, []int{})
func TestCheckBuiltinAppendSliceSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append([]int{}, []int{})`, env,
		`cannot use []int literal (type []int) as type int in append`,
	)

}

// Test Append([]int{}, map[int]int{})
func TestCheckBuiltinAppendSliceMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append([]int{}, map[int]int{})`, env,
		`cannot use map[int]int literal (type map[int]int) as type int in append`,
	)

}

// Test Append([]int{}, int)
func TestCheckBuiltinAppendSliceType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append([]int{}, int)`, env,
		`type int is not an expression`,
	)

}

// Test Append([]int{}, map[int]int)
func TestCheckBuiltinAppendSliceMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append([]int{}, map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Append([]int{}, 1, 1)
func TestCheckBuiltinAppendSliceDouble(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `append([]int{}, 1, 1)`, env, reflect.TypeOf(append([]int{}, 1, 1)))
}

// Test Append([]int{}, []int{1,2}...)
func TestCheckBuiltinAppendSliceEllipsis(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `append([]int{}, []int{1,2}...)`, env, reflect.TypeOf(append([]int{}, []int{1,2}...)))
}

// Test Append(map[int]int{})
func TestCheckBuiltinAppendMapX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int{})`, env,
		`first argument to append must be slice; have map[int]int`,
	)

}

// Test Append(map[int]int{}, 1)
func TestCheckBuiltinAppendMapInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int{}, 1)`, env,
		`first argument to append must be slice; have map[int]int`,
	)

}

// Test Append(map[int]int{}, float32(1))
func TestCheckBuiltinAppendMapFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int{}, float32(1))`, env,
		`first argument to append must be slice; have map[int]int`,
	)

}

// Test Append(map[int]int{}, "abc")
func TestCheckBuiltinAppendMapString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int{}, "abc")`, env,
		`first argument to append must be slice; have map[int]int`,
	)

}

// Test Append(map[int]int{}, nil)
func TestCheckBuiltinAppendMapNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int{}, nil)`, env,
		`first argument to append must be slice; have map[int]int`,
	)

}

// Test Append(map[int]int{}, 1.5)
func TestCheckBuiltinAppendMapFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int{}, 1.5)`, env,
		`first argument to append must be slice; have map[int]int`,
	)

}

// Test Append(map[int]int{}, []int{})
func TestCheckBuiltinAppendMapSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int{}, []int{})`, env,
		`first argument to append must be slice; have map[int]int`,
	)

}

// Test Append(map[int]int{}, map[int]int{})
func TestCheckBuiltinAppendMapMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int{}, map[int]int{})`, env,
		`first argument to append must be slice; have map[int]int`,
	)

}

// Test Append(map[int]int{}, int)
func TestCheckBuiltinAppendMapType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int{}, int)`, env,
		`type int is not an expression`,
		`first argument to append must be slice; have map[int]int`,
	)

}

// Test Append(map[int]int{}, map[int]int)
func TestCheckBuiltinAppendMapMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int{}, map[int]int)`, env,
		`type map[int]int is not an expression`,
		`first argument to append must be slice; have map[int]int`,
	)

}

// Test Append(map[int]int{}, 1, 1)
func TestCheckBuiltinAppendMapDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int{}, 1, 1)`, env,
		`first argument to append must be slice; have map[int]int`,
	)

}

// Test Append(map[int]int{}, []int{1,2}...)
func TestCheckBuiltinAppendMapEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int{}, []int{1,2}...)`, env,
		`first argument to append must be slice; have map[int]int`,
	)

}

// Test Append(int)
func TestCheckBuiltinAppendTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(int)`, env,
		`type int is not an expression`,
	)

}

// Test Append(int, 1)
func TestCheckBuiltinAppendTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(int, 1)`, env,
		`type int is not an expression`,
	)

}

// Test Append(int, float32(1))
func TestCheckBuiltinAppendTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(int, float32(1))`, env,
		`type int is not an expression`,
	)

}

// Test Append(int, "abc")
func TestCheckBuiltinAppendTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(int, "abc")`, env,
		`type int is not an expression`,
	)

}

// Test Append(int, nil)
func TestCheckBuiltinAppendTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(int, nil)`, env,
		`type int is not an expression`,
	)

}

// Test Append(int, 1.5)
func TestCheckBuiltinAppendTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(int, 1.5)`, env,
		`type int is not an expression`,
	)

}

// Test Append(int, []int{})
func TestCheckBuiltinAppendTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(int, []int{})`, env,
		`type int is not an expression`,
	)

}

// Test Append(int, map[int]int{})
func TestCheckBuiltinAppendTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(int, map[int]int{})`, env,
		`type int is not an expression`,
	)

}

// Test Append(int, int)
func TestCheckBuiltinAppendTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(int, int)`, env,
		`type int is not an expression`,
		`type int is not an expression`,
	)

}

// Test Append(int, map[int]int)
func TestCheckBuiltinAppendTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(int, map[int]int)`, env,
		`type int is not an expression`,
		`type map[int]int is not an expression`,
	)

}

// Test Append(int, 1, 1)
func TestCheckBuiltinAppendTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(int, 1, 1)`, env,
		`type int is not an expression`,
	)

}

// Test Append(int, []int{1,2}...)
func TestCheckBuiltinAppendTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(int, []int{1,2}...)`, env,
		`type int is not an expression`,
	)

}

// Test Append(map[int]int)
func TestCheckBuiltinAppendMakeTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Append(map[int]int, 1)
func TestCheckBuiltinAppendMakeTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int, 1)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Append(map[int]int, float32(1))
func TestCheckBuiltinAppendMakeTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int, float32(1))`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Append(map[int]int, "abc")
func TestCheckBuiltinAppendMakeTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int, "abc")`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Append(map[int]int, nil)
func TestCheckBuiltinAppendMakeTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int, nil)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Append(map[int]int, 1.5)
func TestCheckBuiltinAppendMakeTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int, 1.5)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Append(map[int]int, []int{})
func TestCheckBuiltinAppendMakeTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int, []int{})`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Append(map[int]int, map[int]int{})
func TestCheckBuiltinAppendMakeTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int, map[int]int{})`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Append(map[int]int, int)
func TestCheckBuiltinAppendMakeTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int, int)`, env,
		`type map[int]int is not an expression`,
		`type int is not an expression`,
	)

}

// Test Append(map[int]int, map[int]int)
func TestCheckBuiltinAppendMakeTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int, map[int]int)`, env,
		`type map[int]int is not an expression`,
		`type map[int]int is not an expression`,
	)

}

// Test Append(map[int]int, 1, 1)
func TestCheckBuiltinAppendMakeTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int, 1, 1)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Append(map[int]int, []int{1,2}...)
func TestCheckBuiltinAppendMakeTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `append(map[int]int, []int{1,2}...)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Copy()
func TestCheckBuiltinCopyXX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy()`, env,
		`missing arguments to copy`,
	)

}

// Test Copy(1)
func TestCheckBuiltinCopyXInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1)`, env,
		`missing arguments to copy`,
	)

}

// Test Copy(float32(1))
func TestCheckBuiltinCopyXFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(float32(1))`, env,
		`missing arguments to copy`,
	)

}

// Test Copy("abc")
func TestCheckBuiltinCopyXString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy("abc")`, env,
		`missing arguments to copy`,
	)

}

// Test Copy(nil)
func TestCheckBuiltinCopyXNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(nil)`, env,
		`missing arguments to copy`,
	)

}

// Test Copy(1.5)
func TestCheckBuiltinCopyXFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1.5)`, env,
		`missing arguments to copy`,
	)

}

// Test Copy([]int{})
func TestCheckBuiltinCopyXSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy([]int{})`, env,
		`missing arguments to copy`,
	)

}

// Test Copy(map[int]int{})
func TestCheckBuiltinCopyXMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int{})`, env,
		`missing arguments to copy`,
	)

}

// Test Copy(int)
func TestCheckBuiltinCopyXType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(int)`, env,
		`missing arguments to copy`,
	)

}

// Test Copy(map[int]int)
func TestCheckBuiltinCopyXMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int)`, env,
		`missing arguments to copy`,
	)

}

// Test Copy(1, 1)
func TestCheckBuiltinCopyXDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1, 1)`, env,
		`arguments to copy must be slices; have int, int`,
	)

}

// Test Copy([]int{1,2}...)
func TestCheckBuiltinCopyXEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy([]int{1,2}...)`, env,
		`invalid use of ... with builtin copy`,
		`missing arguments to copy`,
	)

}

// Test Copy(1)
func TestCheckBuiltinCopyIntX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1)`, env,
		`missing arguments to copy`,
	)

}

// Test Copy(1, 1)
func TestCheckBuiltinCopyIntInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1, 1)`, env,
		`arguments to copy must be slices; have int, int`,
	)

}

// Test Copy(1, float32(1))
func TestCheckBuiltinCopyIntFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1, float32(1))`, env,
		`arguments to copy must be slices; have int, float32`,
	)

}

// Test Copy(1, "abc")
func TestCheckBuiltinCopyIntString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1, "abc")`, env,
		`arguments to copy must be slices; have int, string`,
	)

}

// Test Copy(1, nil)
func TestCheckBuiltinCopyIntNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1, nil)`, env,
		`use of untyped nil`,
		`arguments to copy must be slices; have int, <T>`,
	)

}

// Test Copy(1, 1.5)
func TestCheckBuiltinCopyIntFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1, 1.5)`, env,
		`arguments to copy must be slices; have int, float64`,
	)

}

// Test Copy(1, []int{})
func TestCheckBuiltinCopyIntSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1, []int{})`, env,
		`first argument to copy should be slice; have int`,
	)

}

// Test Copy(1, map[int]int{})
func TestCheckBuiltinCopyIntMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1, map[int]int{})`, env,
		`arguments to copy must be slices; have int, map[int]int`,
	)

}

// Test Copy(1, int)
func TestCheckBuiltinCopyIntType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1, int)`, env,
		`type int is not an expression`,
	)

}

// Test Copy(1, map[int]int)
func TestCheckBuiltinCopyIntMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1, map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Copy(1, 1, 1)
func TestCheckBuiltinCopyIntDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1, 1, 1)`, env,
		`too many arguments to copy`,
	)

}

// Test Copy(1, []int{1,2}...)
func TestCheckBuiltinCopyIntEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1, []int{1,2}...)`, env,
		`invalid use of ... with builtin copy`,
		`first argument to copy should be slice; have int`,
	)

}

// Test Copy(float32(1))
func TestCheckBuiltinCopyFloat32X(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(float32(1))`, env,
		`missing arguments to copy`,
	)

}

// Test Copy(float32(1), 1)
func TestCheckBuiltinCopyFloat32Int(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(float32(1), 1)`, env,
		`arguments to copy must be slices; have float32, int`,
	)

}

// Test Copy(float32(1), float32(1))
func TestCheckBuiltinCopyFloat32Float32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(float32(1), float32(1))`, env,
		`arguments to copy must be slices; have float32, float32`,
	)

}

// Test Copy(float32(1), "abc")
func TestCheckBuiltinCopyFloat32String(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(float32(1), "abc")`, env,
		`arguments to copy must be slices; have float32, string`,
	)

}

// Test Copy(float32(1), nil)
func TestCheckBuiltinCopyFloat32Nil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(float32(1), nil)`, env,
		`use of untyped nil`,
		`arguments to copy must be slices; have float32, <T>`,
	)

}

// Test Copy(float32(1), 1.5)
func TestCheckBuiltinCopyFloat32Float(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(float32(1), 1.5)`, env,
		`arguments to copy must be slices; have float32, float64`,
	)

}

// Test Copy(float32(1), []int{})
func TestCheckBuiltinCopyFloat32Slice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(float32(1), []int{})`, env,
		`first argument to copy should be slice; have float32`,
	)

}

// Test Copy(float32(1), map[int]int{})
func TestCheckBuiltinCopyFloat32Map(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(float32(1), map[int]int{})`, env,
		`arguments to copy must be slices; have float32, map[int]int`,
	)

}

// Test Copy(float32(1), int)
func TestCheckBuiltinCopyFloat32Type(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(float32(1), int)`, env,
		`type int is not an expression`,
	)

}

// Test Copy(float32(1), map[int]int)
func TestCheckBuiltinCopyFloat32MakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(float32(1), map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Copy(float32(1), 1, 1)
func TestCheckBuiltinCopyFloat32Double(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(float32(1), 1, 1)`, env,
		`too many arguments to copy`,
	)

}

// Test Copy(float32(1), []int{1,2}...)
func TestCheckBuiltinCopyFloat32Ellipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(float32(1), []int{1,2}...)`, env,
		`invalid use of ... with builtin copy`,
		`first argument to copy should be slice; have float32`,
	)

}

// Test Copy("abc")
func TestCheckBuiltinCopyStringX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy("abc")`, env,
		`missing arguments to copy`,
	)

}

// Test Copy("abc", 1)
func TestCheckBuiltinCopyStringInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy("abc", 1)`, env,
		`arguments to copy must be slices; have string, int`,
	)

}

// Test Copy("abc", float32(1))
func TestCheckBuiltinCopyStringFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy("abc", float32(1))`, env,
		`arguments to copy must be slices; have string, float32`,
	)

}

// Test Copy("abc", "abc")
func TestCheckBuiltinCopyStringString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy("abc", "abc")`, env,
		`arguments to copy must be slices; have string, string`,
	)

}

// Test Copy("abc", nil)
func TestCheckBuiltinCopyStringNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy("abc", nil)`, env,
		`use of untyped nil`,
		`arguments to copy must be slices; have string, <T>`,
	)

}

// Test Copy("abc", 1.5)
func TestCheckBuiltinCopyStringFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy("abc", 1.5)`, env,
		`arguments to copy must be slices; have string, float64`,
	)

}

// Test Copy("abc", []int{})
func TestCheckBuiltinCopyStringSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy("abc", []int{})`, env,
		`first argument to copy should be slice; have string`,
	)

}

// Test Copy("abc", map[int]int{})
func TestCheckBuiltinCopyStringMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy("abc", map[int]int{})`, env,
		`arguments to copy must be slices; have string, map[int]int`,
	)

}

// Test Copy("abc", int)
func TestCheckBuiltinCopyStringType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy("abc", int)`, env,
		`type int is not an expression`,
	)

}

// Test Copy("abc", map[int]int)
func TestCheckBuiltinCopyStringMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy("abc", map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Copy("abc", 1, 1)
func TestCheckBuiltinCopyStringDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy("abc", 1, 1)`, env,
		`too many arguments to copy`,
	)

}

// Test Copy("abc", []int{1,2}...)
func TestCheckBuiltinCopyStringEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy("abc", []int{1,2}...)`, env,
		`invalid use of ... with builtin copy`,
		`first argument to copy should be slice; have string`,
	)

}

// Test Copy(nil)
func TestCheckBuiltinCopyNilX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(nil)`, env,
		`missing arguments to copy`,
	)

}

// Test Copy(nil, 1)
func TestCheckBuiltinCopyNilInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(nil, 1)`, env,
		`use of untyped nil`,
		`arguments to copy must be slices; have <T>, int`,
	)

}

// Test Copy(nil, float32(1))
func TestCheckBuiltinCopyNilFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(nil, float32(1))`, env,
		`use of untyped nil`,
		`arguments to copy must be slices; have <T>, float32`,
	)

}

// Test Copy(nil, "abc")
func TestCheckBuiltinCopyNilString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(nil, "abc")`, env,
		`use of untyped nil`,
		`arguments to copy must be slices; have <T>, string`,
	)

}

// Test Copy(nil, nil)
func TestCheckBuiltinCopyNilNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(nil, nil)`, env,
		`use of untyped nil`,
		`use of untyped nil`,
		`arguments to copy must be slices; have <T>, <T>`,
	)

}

// Test Copy(nil, 1.5)
func TestCheckBuiltinCopyNilFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(nil, 1.5)`, env,
		`use of untyped nil`,
		`arguments to copy must be slices; have <T>, float64`,
	)

}

// Test Copy(nil, []int{})
func TestCheckBuiltinCopyNilSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(nil, []int{})`, env,
		`use of untyped nil`,
		`first argument to copy should be slice; have <T>`,
	)

}

// Test Copy(nil, map[int]int{})
func TestCheckBuiltinCopyNilMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(nil, map[int]int{})`, env,
		`use of untyped nil`,
		`arguments to copy must be slices; have <T>, map[int]int`,
	)

}

// Test Copy(nil, int)
func TestCheckBuiltinCopyNilType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(nil, int)`, env,
		`type int is not an expression`,
	)

}

// Test Copy(nil, map[int]int)
func TestCheckBuiltinCopyNilMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(nil, map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Copy(nil, 1, 1)
func TestCheckBuiltinCopyNilDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(nil, 1, 1)`, env,
		`too many arguments to copy`,
	)

}

// Test Copy(nil, []int{1,2}...)
func TestCheckBuiltinCopyNilEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(nil, []int{1,2}...)`, env,
		`invalid use of ... with builtin copy`,
		`use of untyped nil`,
		`first argument to copy should be slice; have <T>`,
	)

}

// Test Copy(1.5)
func TestCheckBuiltinCopyFloatX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1.5)`, env,
		`missing arguments to copy`,
	)

}

// Test Copy(1.5, 1)
func TestCheckBuiltinCopyFloatInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1.5, 1)`, env,
		`arguments to copy must be slices; have float64, int`,
	)

}

// Test Copy(1.5, float32(1))
func TestCheckBuiltinCopyFloatFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1.5, float32(1))`, env,
		`arguments to copy must be slices; have float64, float32`,
	)

}

// Test Copy(1.5, "abc")
func TestCheckBuiltinCopyFloatString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1.5, "abc")`, env,
		`arguments to copy must be slices; have float64, string`,
	)

}

// Test Copy(1.5, nil)
func TestCheckBuiltinCopyFloatNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1.5, nil)`, env,
		`use of untyped nil`,
		`arguments to copy must be slices; have float64, <T>`,
	)

}

// Test Copy(1.5, 1.5)
func TestCheckBuiltinCopyFloatFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1.5, 1.5)`, env,
		`arguments to copy must be slices; have float64, float64`,
	)

}

// Test Copy(1.5, []int{})
func TestCheckBuiltinCopyFloatSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1.5, []int{})`, env,
		`first argument to copy should be slice; have float64`,
	)

}

// Test Copy(1.5, map[int]int{})
func TestCheckBuiltinCopyFloatMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1.5, map[int]int{})`, env,
		`arguments to copy must be slices; have float64, map[int]int`,
	)

}

// Test Copy(1.5, int)
func TestCheckBuiltinCopyFloatType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1.5, int)`, env,
		`type int is not an expression`,
	)

}

// Test Copy(1.5, map[int]int)
func TestCheckBuiltinCopyFloatMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1.5, map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Copy(1.5, 1, 1)
func TestCheckBuiltinCopyFloatDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1.5, 1, 1)`, env,
		`too many arguments to copy`,
	)

}

// Test Copy(1.5, []int{1,2}...)
func TestCheckBuiltinCopyFloatEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(1.5, []int{1,2}...)`, env,
		`invalid use of ... with builtin copy`,
		`first argument to copy should be slice; have float64`,
	)

}

// Test Copy([]int{})
func TestCheckBuiltinCopySliceX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy([]int{})`, env,
		`missing arguments to copy`,
	)

}

// Test Copy([]int{}, 1)
func TestCheckBuiltinCopySliceInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy([]int{}, 1)`, env,
		`second argument to copy should be slice or string; have int`,
	)

}

// Test Copy([]int{}, float32(1))
func TestCheckBuiltinCopySliceFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy([]int{}, float32(1))`, env,
		`second argument to copy should be slice or string; have float32`,
	)

}

// Test Copy([]int{}, "abc")
func TestCheckBuiltinCopySliceString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy([]int{}, "abc")`, env,
		`arguments to copy have different element types: []int and string`,
	)

}

// Test Copy([]int{}, nil)
func TestCheckBuiltinCopySliceNil(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env

}

// Test Copy([]int{}, 1.5)
func TestCheckBuiltinCopySliceFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy([]int{}, 1.5)`, env,
		`second argument to copy should be slice or string; have float64`,
	)

}

// Test Copy([]int{}, []int{})
func TestCheckBuiltinCopySliceSlice(t *testing.T) {
	env := MakeSimpleEnv()
	expectType(t, `copy([]int{}, []int{})`, env, reflect.TypeOf(copy([]int{}, []int{})))
}

// Test Copy([]int{}, map[int]int{})
func TestCheckBuiltinCopySliceMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy([]int{}, map[int]int{})`, env,
		`second argument to copy should be slice or string; have map[int]int`,
	)

}

// Test Copy([]int{}, int)
func TestCheckBuiltinCopySliceType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy([]int{}, int)`, env,
		`type int is not an expression`,
	)

}

// Test Copy([]int{}, map[int]int)
func TestCheckBuiltinCopySliceMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy([]int{}, map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Copy([]int{}, 1, 1)
func TestCheckBuiltinCopySliceDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy([]int{}, 1, 1)`, env,
		`too many arguments to copy`,
	)

}

// Test Copy([]int{}, []int{1,2}...)
func TestCheckBuiltinCopySliceEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy([]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin copy`,
	)

}

// Test Copy(map[int]int{})
func TestCheckBuiltinCopyMapX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int{})`, env,
		`missing arguments to copy`,
	)

}

// Test Copy(map[int]int{}, 1)
func TestCheckBuiltinCopyMapInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int{}, 1)`, env,
		`arguments to copy must be slices; have map[int]int, int`,
	)

}

// Test Copy(map[int]int{}, float32(1))
func TestCheckBuiltinCopyMapFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int{}, float32(1))`, env,
		`arguments to copy must be slices; have map[int]int, float32`,
	)

}

// Test Copy(map[int]int{}, "abc")
func TestCheckBuiltinCopyMapString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int{}, "abc")`, env,
		`arguments to copy must be slices; have map[int]int, string`,
	)

}

// Test Copy(map[int]int{}, nil)
func TestCheckBuiltinCopyMapNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int{}, nil)`, env,
		`use of untyped nil`,
		`arguments to copy must be slices; have map[int]int, <T>`,
	)

}

// Test Copy(map[int]int{}, 1.5)
func TestCheckBuiltinCopyMapFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int{}, 1.5)`, env,
		`arguments to copy must be slices; have map[int]int, float64`,
	)

}

// Test Copy(map[int]int{}, []int{})
func TestCheckBuiltinCopyMapSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int{}, []int{})`, env,
		`first argument to copy should be slice; have map[int]int`,
	)

}

// Test Copy(map[int]int{}, map[int]int{})
func TestCheckBuiltinCopyMapMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int{}, map[int]int{})`, env,
		`arguments to copy must be slices; have map[int]int, map[int]int`,
	)

}

// Test Copy(map[int]int{}, int)
func TestCheckBuiltinCopyMapType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int{}, int)`, env,
		`type int is not an expression`,
	)

}

// Test Copy(map[int]int{}, map[int]int)
func TestCheckBuiltinCopyMapMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int{}, map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Copy(map[int]int{}, 1, 1)
func TestCheckBuiltinCopyMapDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int{}, 1, 1)`, env,
		`too many arguments to copy`,
	)

}

// Test Copy(map[int]int{}, []int{1,2}...)
func TestCheckBuiltinCopyMapEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin copy`,
		`first argument to copy should be slice; have map[int]int`,
	)

}

// Test Copy(int)
func TestCheckBuiltinCopyTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(int)`, env,
		`missing arguments to copy`,
	)

}

// Test Copy(int, 1)
func TestCheckBuiltinCopyTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(int, 1)`, env,
		`type int is not an expression`,
	)

}

// Test Copy(int, float32(1))
func TestCheckBuiltinCopyTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(int, float32(1))`, env,
		`type int is not an expression`,
	)

}

// Test Copy(int, "abc")
func TestCheckBuiltinCopyTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(int, "abc")`, env,
		`type int is not an expression`,
	)

}

// Test Copy(int, nil)
func TestCheckBuiltinCopyTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(int, nil)`, env,
		`type int is not an expression`,
	)

}

// Test Copy(int, 1.5)
func TestCheckBuiltinCopyTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(int, 1.5)`, env,
		`type int is not an expression`,
	)

}

// Test Copy(int, []int{})
func TestCheckBuiltinCopyTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(int, []int{})`, env,
		`type int is not an expression`,
	)

}

// Test Copy(int, map[int]int{})
func TestCheckBuiltinCopyTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(int, map[int]int{})`, env,
		`type int is not an expression`,
	)

}

// Test Copy(int, int)
func TestCheckBuiltinCopyTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(int, int)`, env,
		`type int is not an expression`,
		`type int is not an expression`,
	)

}

// Test Copy(int, map[int]int)
func TestCheckBuiltinCopyTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(int, map[int]int)`, env,
		`type int is not an expression`,
		`type map[int]int is not an expression`,
	)

}

// Test Copy(int, 1, 1)
func TestCheckBuiltinCopyTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(int, 1, 1)`, env,
		`too many arguments to copy`,
	)

}

// Test Copy(int, []int{1,2}...)
func TestCheckBuiltinCopyTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(int, []int{1,2}...)`, env,
		`invalid use of ... with builtin copy`,
		`type int is not an expression`,
	)

}

// Test Copy(map[int]int)
func TestCheckBuiltinCopyMakeTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int)`, env,
		`missing arguments to copy`,
	)

}

// Test Copy(map[int]int, 1)
func TestCheckBuiltinCopyMakeTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int, 1)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Copy(map[int]int, float32(1))
func TestCheckBuiltinCopyMakeTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int, float32(1))`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Copy(map[int]int, "abc")
func TestCheckBuiltinCopyMakeTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int, "abc")`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Copy(map[int]int, nil)
func TestCheckBuiltinCopyMakeTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int, nil)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Copy(map[int]int, 1.5)
func TestCheckBuiltinCopyMakeTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int, 1.5)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Copy(map[int]int, []int{})
func TestCheckBuiltinCopyMakeTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int, []int{})`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Copy(map[int]int, map[int]int{})
func TestCheckBuiltinCopyMakeTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int, map[int]int{})`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Copy(map[int]int, int)
func TestCheckBuiltinCopyMakeTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int, int)`, env,
		`type map[int]int is not an expression`,
		`type int is not an expression`,
	)

}

// Test Copy(map[int]int, map[int]int)
func TestCheckBuiltinCopyMakeTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int, map[int]int)`, env,
		`type map[int]int is not an expression`,
		`type map[int]int is not an expression`,
	)

}

// Test Copy(map[int]int, 1, 1)
func TestCheckBuiltinCopyMakeTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int, 1, 1)`, env,
		`too many arguments to copy`,
	)

}

// Test Copy(map[int]int, []int{1,2}...)
func TestCheckBuiltinCopyMakeTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `copy(map[int]int, []int{1,2}...)`, env,
		`invalid use of ... with builtin copy`,
		`type map[int]int is not an expression`,
	)

}

// Test Delete()
func TestCheckBuiltinDeleteXX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete()`, env,
		`missing arguments to delete`,
	)

}

// Test Delete(1)
func TestCheckBuiltinDeleteXInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1)`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete(float32(1))
func TestCheckBuiltinDeleteXFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(float32(1))`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete("abc")
func TestCheckBuiltinDeleteXString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete("abc")`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete(nil)
func TestCheckBuiltinDeleteXNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(nil)`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete(1.5)
func TestCheckBuiltinDeleteXFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1.5)`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete([]int{})
func TestCheckBuiltinDeleteXSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete([]int{})`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete(map[int]int{})
func TestCheckBuiltinDeleteXMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int{})`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete(int)
func TestCheckBuiltinDeleteXType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(int)`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete(map[int]int)
func TestCheckBuiltinDeleteXMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int)`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete(1, 1)
func TestCheckBuiltinDeleteXDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1, 1)`, env,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete([]int{1,2}...)
func TestCheckBuiltinDeleteXEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete([]int{1,2}...)`, env,
		`invalid use of ... with builtin delete`,
		`missing second (key) argument to delete`,
	)

}

// Test Delete(1)
func TestCheckBuiltinDeleteIntX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1)`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete(1, 1)
func TestCheckBuiltinDeleteIntInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1, 1)`, env,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1, float32(1))
func TestCheckBuiltinDeleteIntFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1, float32(1))`, env,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1, "abc")
func TestCheckBuiltinDeleteIntString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1, "abc")`, env,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1, nil)
func TestCheckBuiltinDeleteIntNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1, nil)`, env,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1, 1.5)
func TestCheckBuiltinDeleteIntFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1, 1.5)`, env,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1, []int{})
func TestCheckBuiltinDeleteIntSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1, []int{})`, env,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1, map[int]int{})
func TestCheckBuiltinDeleteIntMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1, map[int]int{})`, env,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1, int)
func TestCheckBuiltinDeleteIntType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1, int)`, env,
		`type int is not an expression`,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1, map[int]int)
func TestCheckBuiltinDeleteIntMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1, map[int]int)`, env,
		`type map[int]int is not an expression`,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1, 1, 1)
func TestCheckBuiltinDeleteIntDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1, 1, 1)`, env,
		`too many arguments to delete`,
	)

}

// Test Delete(1, []int{1,2}...)
func TestCheckBuiltinDeleteIntEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1, []int{1,2}...)`, env,
		`invalid use of ... with builtin delete`,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(float32(1))
func TestCheckBuiltinDeleteFloat32X(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(float32(1))`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete(float32(1), 1)
func TestCheckBuiltinDeleteFloat32Int(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(float32(1), 1)`, env,
		`first argument to delete must be map; have float32`,
	)

}

// Test Delete(float32(1), float32(1))
func TestCheckBuiltinDeleteFloat32Float32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(float32(1), float32(1))`, env,
		`first argument to delete must be map; have float32`,
	)

}

// Test Delete(float32(1), "abc")
func TestCheckBuiltinDeleteFloat32String(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(float32(1), "abc")`, env,
		`first argument to delete must be map; have float32`,
	)

}

// Test Delete(float32(1), nil)
func TestCheckBuiltinDeleteFloat32Nil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(float32(1), nil)`, env,
		`first argument to delete must be map; have float32`,
	)

}

// Test Delete(float32(1), 1.5)
func TestCheckBuiltinDeleteFloat32Float(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(float32(1), 1.5)`, env,
		`first argument to delete must be map; have float32`,
	)

}

// Test Delete(float32(1), []int{})
func TestCheckBuiltinDeleteFloat32Slice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(float32(1), []int{})`, env,
		`first argument to delete must be map; have float32`,
	)

}

// Test Delete(float32(1), map[int]int{})
func TestCheckBuiltinDeleteFloat32Map(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(float32(1), map[int]int{})`, env,
		`first argument to delete must be map; have float32`,
	)

}

// Test Delete(float32(1), int)
func TestCheckBuiltinDeleteFloat32Type(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(float32(1), int)`, env,
		`type int is not an expression`,
		`first argument to delete must be map; have float32`,
	)

}

// Test Delete(float32(1), map[int]int)
func TestCheckBuiltinDeleteFloat32MakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(float32(1), map[int]int)`, env,
		`type map[int]int is not an expression`,
		`first argument to delete must be map; have float32`,
	)

}

// Test Delete(float32(1), 1, 1)
func TestCheckBuiltinDeleteFloat32Double(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(float32(1), 1, 1)`, env,
		`too many arguments to delete`,
	)

}

// Test Delete(float32(1), []int{1,2}...)
func TestCheckBuiltinDeleteFloat32Ellipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(float32(1), []int{1,2}...)`, env,
		`invalid use of ... with builtin delete`,
		`first argument to delete must be map; have float32`,
	)

}

// Test Delete("abc")
func TestCheckBuiltinDeleteStringX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete("abc")`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete("abc", 1)
func TestCheckBuiltinDeleteStringInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete("abc", 1)`, env,
		`first argument to delete must be map; have untyped string`,
	)

}

// Test Delete("abc", float32(1))
func TestCheckBuiltinDeleteStringFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete("abc", float32(1))`, env,
		`first argument to delete must be map; have untyped string`,
	)

}

// Test Delete("abc", "abc")
func TestCheckBuiltinDeleteStringString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete("abc", "abc")`, env,
		`first argument to delete must be map; have untyped string`,
	)

}

// Test Delete("abc", nil)
func TestCheckBuiltinDeleteStringNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete("abc", nil)`, env,
		`first argument to delete must be map; have untyped string`,
	)

}

// Test Delete("abc", 1.5)
func TestCheckBuiltinDeleteStringFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete("abc", 1.5)`, env,
		`first argument to delete must be map; have untyped string`,
	)

}

// Test Delete("abc", []int{})
func TestCheckBuiltinDeleteStringSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete("abc", []int{})`, env,
		`first argument to delete must be map; have untyped string`,
	)

}

// Test Delete("abc", map[int]int{})
func TestCheckBuiltinDeleteStringMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete("abc", map[int]int{})`, env,
		`first argument to delete must be map; have untyped string`,
	)

}

// Test Delete("abc", int)
func TestCheckBuiltinDeleteStringType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete("abc", int)`, env,
		`type int is not an expression`,
		`first argument to delete must be map; have untyped string`,
	)

}

// Test Delete("abc", map[int]int)
func TestCheckBuiltinDeleteStringMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete("abc", map[int]int)`, env,
		`type map[int]int is not an expression`,
		`first argument to delete must be map; have untyped string`,
	)

}

// Test Delete("abc", 1, 1)
func TestCheckBuiltinDeleteStringDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete("abc", 1, 1)`, env,
		`too many arguments to delete`,
	)

}

// Test Delete("abc", []int{1,2}...)
func TestCheckBuiltinDeleteStringEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete("abc", []int{1,2}...)`, env,
		`invalid use of ... with builtin delete`,
		`first argument to delete must be map; have untyped string`,
	)

}

// Test Delete(nil)
func TestCheckBuiltinDeleteNilX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(nil)`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete(nil, 1)
func TestCheckBuiltinDeleteNilInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(nil, 1)`, env,
		`first argument to delete must be map; have nil`,
	)

}

// Test Delete(nil, float32(1))
func TestCheckBuiltinDeleteNilFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(nil, float32(1))`, env,
		`first argument to delete must be map; have nil`,
	)

}

// Test Delete(nil, "abc")
func TestCheckBuiltinDeleteNilString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(nil, "abc")`, env,
		`first argument to delete must be map; have nil`,
	)

}

// Test Delete(nil, nil)
func TestCheckBuiltinDeleteNilNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(nil, nil)`, env,
		`first argument to delete must be map; have nil`,
	)

}

// Test Delete(nil, 1.5)
func TestCheckBuiltinDeleteNilFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(nil, 1.5)`, env,
		`first argument to delete must be map; have nil`,
	)

}

// Test Delete(nil, []int{})
func TestCheckBuiltinDeleteNilSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(nil, []int{})`, env,
		`first argument to delete must be map; have nil`,
	)

}

// Test Delete(nil, map[int]int{})
func TestCheckBuiltinDeleteNilMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(nil, map[int]int{})`, env,
		`first argument to delete must be map; have nil`,
	)

}

// Test Delete(nil, int)
func TestCheckBuiltinDeleteNilType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(nil, int)`, env,
		`type int is not an expression`,
		`first argument to delete must be map; have nil`,
	)

}

// Test Delete(nil, map[int]int)
func TestCheckBuiltinDeleteNilMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(nil, map[int]int)`, env,
		`type map[int]int is not an expression`,
		`first argument to delete must be map; have nil`,
	)

}

// Test Delete(nil, 1, 1)
func TestCheckBuiltinDeleteNilDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(nil, 1, 1)`, env,
		`too many arguments to delete`,
	)

}

// Test Delete(nil, []int{1,2}...)
func TestCheckBuiltinDeleteNilEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(nil, []int{1,2}...)`, env,
		`invalid use of ... with builtin delete`,
		`first argument to delete must be map; have nil`,
	)

}

// Test Delete(1.5)
func TestCheckBuiltinDeleteFloatX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1.5)`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete(1.5, 1)
func TestCheckBuiltinDeleteFloatInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1.5, 1)`, env,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1.5, float32(1))
func TestCheckBuiltinDeleteFloatFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1.5, float32(1))`, env,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1.5, "abc")
func TestCheckBuiltinDeleteFloatString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1.5, "abc")`, env,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1.5, nil)
func TestCheckBuiltinDeleteFloatNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1.5, nil)`, env,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1.5, 1.5)
func TestCheckBuiltinDeleteFloatFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1.5, 1.5)`, env,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1.5, []int{})
func TestCheckBuiltinDeleteFloatSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1.5, []int{})`, env,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1.5, map[int]int{})
func TestCheckBuiltinDeleteFloatMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1.5, map[int]int{})`, env,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1.5, int)
func TestCheckBuiltinDeleteFloatType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1.5, int)`, env,
		`type int is not an expression`,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1.5, map[int]int)
func TestCheckBuiltinDeleteFloatMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1.5, map[int]int)`, env,
		`type map[int]int is not an expression`,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete(1.5, 1, 1)
func TestCheckBuiltinDeleteFloatDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1.5, 1, 1)`, env,
		`too many arguments to delete`,
	)

}

// Test Delete(1.5, []int{1,2}...)
func TestCheckBuiltinDeleteFloatEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(1.5, []int{1,2}...)`, env,
		`invalid use of ... with builtin delete`,
		`first argument to delete must be map; have untyped number`,
	)

}

// Test Delete([]int{})
func TestCheckBuiltinDeleteSliceX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete([]int{})`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete([]int{}, 1)
func TestCheckBuiltinDeleteSliceInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete([]int{}, 1)`, env,
		`first argument to delete must be map; have []int`,
	)

}

// Test Delete([]int{}, float32(1))
func TestCheckBuiltinDeleteSliceFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete([]int{}, float32(1))`, env,
		`first argument to delete must be map; have []int`,
	)

}

// Test Delete([]int{}, "abc")
func TestCheckBuiltinDeleteSliceString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete([]int{}, "abc")`, env,
		`first argument to delete must be map; have []int`,
	)

}

// Test Delete([]int{}, nil)
func TestCheckBuiltinDeleteSliceNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete([]int{}, nil)`, env,
		`first argument to delete must be map; have []int`,
	)

}

// Test Delete([]int{}, 1.5)
func TestCheckBuiltinDeleteSliceFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete([]int{}, 1.5)`, env,
		`first argument to delete must be map; have []int`,
	)

}

// Test Delete([]int{}, []int{})
func TestCheckBuiltinDeleteSliceSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete([]int{}, []int{})`, env,
		`first argument to delete must be map; have []int`,
	)

}

// Test Delete([]int{}, map[int]int{})
func TestCheckBuiltinDeleteSliceMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete([]int{}, map[int]int{})`, env,
		`first argument to delete must be map; have []int`,
	)

}

// Test Delete([]int{}, int)
func TestCheckBuiltinDeleteSliceType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete([]int{}, int)`, env,
		`type int is not an expression`,
		`first argument to delete must be map; have []int`,
	)

}

// Test Delete([]int{}, map[int]int)
func TestCheckBuiltinDeleteSliceMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete([]int{}, map[int]int)`, env,
		`type map[int]int is not an expression`,
		`first argument to delete must be map; have []int`,
	)

}

// Test Delete([]int{}, 1, 1)
func TestCheckBuiltinDeleteSliceDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete([]int{}, 1, 1)`, env,
		`too many arguments to delete`,
	)

}

// Test Delete([]int{}, []int{1,2}...)
func TestCheckBuiltinDeleteSliceEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete([]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin delete`,
		`first argument to delete must be map; have []int`,
	)

}

// Test Delete(map[int]int{})
func TestCheckBuiltinDeleteMapX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int{})`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete(map[int]int{}, 1)
func TestCheckBuiltinDeleteMapInt(t *testing.T) {
	env := MakeSimpleEnv()
	_ = env
}

// Test Delete(map[int]int{}, float32(1))
func TestCheckBuiltinDeleteMapFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int{}, float32(1))`, env,
		`cannot use float32(1) (type float32) as type int in delete`,
	)

}

// Test Delete(map[int]int{}, "abc")
func TestCheckBuiltinDeleteMapString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int{}, "abc")`, env,
		`cannot use "abc" (type string) as type int in delete`,
	)

}

// Test Delete(map[int]int{}, nil)
func TestCheckBuiltinDeleteMapNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int{}, nil)`, env,
		`cannot use nil as type int in delete`,
	)

}

// Test Delete(map[int]int{}, 1.5)
func TestCheckBuiltinDeleteMapFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int{}, 1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Delete(map[int]int{}, []int{})
func TestCheckBuiltinDeleteMapSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int{}, []int{})`, env,
		`cannot use []int literal (type []int) as type int in delete`,
	)

}

// Test Delete(map[int]int{}, map[int]int{})
func TestCheckBuiltinDeleteMapMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int{}, map[int]int{})`, env,
		`cannot use map[int]int literal (type map[int]int) as type int in delete`,
	)

}

// Test Delete(map[int]int{}, int)
func TestCheckBuiltinDeleteMapType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int{}, int)`, env,
		`type int is not an expression`,
	)

}

// Test Delete(map[int]int{}, map[int]int)
func TestCheckBuiltinDeleteMapMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int{}, map[int]int)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Delete(map[int]int{}, 1, 1)
func TestCheckBuiltinDeleteMapDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int{}, 1, 1)`, env,
		`too many arguments to delete`,
	)

}

// Test Delete(map[int]int{}, []int{1,2}...)
func TestCheckBuiltinDeleteMapEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int{}, []int{1,2}...)`, env,
		`invalid use of ... with builtin delete`,
		`cannot use []int literal (type []int) as type int in delete`,
	)

}

// Test Delete(int)
func TestCheckBuiltinDeleteTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(int)`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete(int, 1)
func TestCheckBuiltinDeleteTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(int, 1)`, env,
		`type int is not an expression`,
	)

}

// Test Delete(int, float32(1))
func TestCheckBuiltinDeleteTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(int, float32(1))`, env,
		`type int is not an expression`,
	)

}

// Test Delete(int, "abc")
func TestCheckBuiltinDeleteTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(int, "abc")`, env,
		`type int is not an expression`,
	)

}

// Test Delete(int, nil)
func TestCheckBuiltinDeleteTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(int, nil)`, env,
		`type int is not an expression`,
	)

}

// Test Delete(int, 1.5)
func TestCheckBuiltinDeleteTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(int, 1.5)`, env,
		`type int is not an expression`,
	)

}

// Test Delete(int, []int{})
func TestCheckBuiltinDeleteTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(int, []int{})`, env,
		`type int is not an expression`,
	)

}

// Test Delete(int, map[int]int{})
func TestCheckBuiltinDeleteTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(int, map[int]int{})`, env,
		`type int is not an expression`,
	)

}

// Test Delete(int, int)
func TestCheckBuiltinDeleteTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(int, int)`, env,
		`type int is not an expression`,
		`type int is not an expression`,
	)

}

// Test Delete(int, map[int]int)
func TestCheckBuiltinDeleteTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(int, map[int]int)`, env,
		`type int is not an expression`,
		`type map[int]int is not an expression`,
	)

}

// Test Delete(int, 1, 1)
func TestCheckBuiltinDeleteTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(int, 1, 1)`, env,
		`too many arguments to delete`,
	)

}

// Test Delete(int, []int{1,2}...)
func TestCheckBuiltinDeleteTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(int, []int{1,2}...)`, env,
		`invalid use of ... with builtin delete`,
		`type int is not an expression`,
	)

}

// Test Delete(map[int]int)
func TestCheckBuiltinDeleteMakeTypeX(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int)`, env,
		`missing second (key) argument to delete`,
	)

}

// Test Delete(map[int]int, 1)
func TestCheckBuiltinDeleteMakeTypeInt(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int, 1)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Delete(map[int]int, float32(1))
func TestCheckBuiltinDeleteMakeTypeFloat32(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int, float32(1))`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Delete(map[int]int, "abc")
func TestCheckBuiltinDeleteMakeTypeString(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int, "abc")`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Delete(map[int]int, nil)
func TestCheckBuiltinDeleteMakeTypeNil(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int, nil)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Delete(map[int]int, 1.5)
func TestCheckBuiltinDeleteMakeTypeFloat(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int, 1.5)`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Delete(map[int]int, []int{})
func TestCheckBuiltinDeleteMakeTypeSlice(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int, []int{})`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Delete(map[int]int, map[int]int{})
func TestCheckBuiltinDeleteMakeTypeMap(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int, map[int]int{})`, env,
		`type map[int]int is not an expression`,
	)

}

// Test Delete(map[int]int, int)
func TestCheckBuiltinDeleteMakeTypeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int, int)`, env,
		`type map[int]int is not an expression`,
		`type int is not an expression`,
	)

}

// Test Delete(map[int]int, map[int]int)
func TestCheckBuiltinDeleteMakeTypeMakeType(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int, map[int]int)`, env,
		`type map[int]int is not an expression`,
		`type map[int]int is not an expression`,
	)

}

// Test Delete(map[int]int, 1, 1)
func TestCheckBuiltinDeleteMakeTypeDouble(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int, 1, 1)`, env,
		`too many arguments to delete`,
	)

}

// Test Delete(map[int]int, []int{1,2}...)
func TestCheckBuiltinDeleteMakeTypeEllipsis(t *testing.T) {
	env := MakeSimpleEnv()

	expectCheckError(t, `delete(map[int]int, []int{1,2}...)`, env,
		`invalid use of ... with builtin delete`,
		`type map[int]int is not an expression`,
	)

}
