package eval

import (
	"testing"
	"reflect"
)

type interfaceX interface { x() }
type interfaceY interface { x() }
type interfaceZ interface { z() }
type XinterfaceX int
func (XinterfaceX) x() {}
type arrayT [2]int
type mapT map[int]int
type sliceT []int
type structT struct {
	a int
	_ []int
}
type structUncompT struct {
	a int
	b []int
}

func makeCheckBinaryNonConstExprEnv() *SimpleEnv {
	env := MakeSimpleEnv()
	env.Types["interfaceX"] = reflect.TypeOf(new(interfaceX)).Elem()
	env.Types["interfaceY"] = reflect.TypeOf(new(interfaceY)).Elem()
	env.Types["interfaceZ"] = reflect.TypeOf(new(interfaceZ)).Elem()
	env.Types["XinterfaceX"] = reflect.TypeOf(XinterfaceX(0))
	env.Types["arrayT"] = reflect.TypeOf(arrayT{})
	env.Types["mapT"] = reflect.TypeOf(mapT{})
	env.Types["sliceT"] = reflect.TypeOf(sliceT{})
	env.Types["structT"] = reflect.TypeOf(structT{})
	env.Types["structUncompT"] = reflect.TypeOf(structUncompT{})
	return env
}

// Test Int + ConstInt
func TestCheckBinaryNonConstExprIntAddConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a + 4`, env, reflect.TypeOf(a + 4))
}

// Test Int + ConstRune
func TestCheckBinaryNonConstExprIntAddConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a + '@'`, env, reflect.TypeOf(a + '@'))
}

// Test Int + ConstFloat
func TestCheckBinaryNonConstExprIntAddConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a + 2.0`, env, reflect.TypeOf(a + 2.0))
}

// Test Int + ConstComplex
func TestCheckBinaryNonConstExprIntAddConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int + ConstBool
func TestCheckBinaryNonConstExprIntAddConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + true`, env,
		`cannot convert true to type int`,
		`invalid operation: a + true (mismatched types int and bool)`,
	)

}

// Test Int + ConstString
func TestCheckBinaryNonConstExprIntAddConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: a + "abc" (mismatched types int and string)`,
	)

}

// Test Int + ConstNil
func TestCheckBinaryNonConstExprIntAddConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + nil`, env,
		`cannot convert nil to type int`,
	)

}

// Test Int + Int
func TestCheckBinaryNonConstExprIntAddInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a + b`, env, reflect.TypeOf(a + b))
}

// Test Int + Float32
func TestCheckBinaryNonConstExprIntAddFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types int and float32)`,
	)

}

// Test Int + Complex128
func TestCheckBinaryNonConstExprIntAddComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types int and complex128)`,
	)

}

// Test Int + String
func TestCheckBinaryNonConstExprIntAddString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types int and string)`,
	)

}

// Test Int + BoolT
func TestCheckBinaryNonConstExprIntAddBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types int and bool)`,
	)

}

// Test Int + Slice
func TestCheckBinaryNonConstExprIntAddSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types int and eval.sliceT)`,
	)

}

// Test Int + Array
func TestCheckBinaryNonConstExprIntAddArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types int and eval.arrayT)`,
	)

}

// Test Int + Map
func TestCheckBinaryNonConstExprIntAddMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types int and eval.mapT)`,
	)

}

// Test Int + XinterfaceX
func TestCheckBinaryNonConstExprIntAddXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types int and eval.XinterfaceX)`,
	)

}

// Test Int + InterfaceX
func TestCheckBinaryNonConstExprIntAddInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types int and eval.interfaceX)`,
	)

}

// Test Int + InterfaceY
func TestCheckBinaryNonConstExprIntAddInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types int and eval.interfaceY)`,
	)

}

// Test Int + InterfaceZ
func TestCheckBinaryNonConstExprIntAddInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types int and eval.interfaceZ)`,
	)

}

// Test Int + Ptr
func TestCheckBinaryNonConstExprIntAddPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types int and *int)`,
	)

}

// Test Int + Struct
func TestCheckBinaryNonConstExprIntAddStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types int and eval.structT)`,
	)

}

// Test Int + StructUncomp
func TestCheckBinaryNonConstExprIntAddStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types int and eval.structUncompT)`,
	)

}

// Test Int & ConstInt
func TestCheckBinaryNonConstExprIntAndConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a & 4`, env, reflect.TypeOf(a & 4))
}

// Test Int & ConstRune
func TestCheckBinaryNonConstExprIntAndConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a & '@'`, env, reflect.TypeOf(a & '@'))
}

// Test Int & ConstFloat
func TestCheckBinaryNonConstExprIntAndConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a & 2.0`, env, reflect.TypeOf(a & 2.0))
}

// Test Int & ConstComplex
func TestCheckBinaryNonConstExprIntAndConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int & ConstBool
func TestCheckBinaryNonConstExprIntAndConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & true`, env,
		`cannot convert true to type int`,
		`invalid operation: a & true (mismatched types int and bool)`,
	)

}

// Test Int & ConstString
func TestCheckBinaryNonConstExprIntAndConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: a & "abc" (mismatched types int and string)`,
	)

}

// Test Int & ConstNil
func TestCheckBinaryNonConstExprIntAndConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & nil`, env,
		`cannot convert nil to type int`,
	)

}

// Test Int & Int
func TestCheckBinaryNonConstExprIntAndInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a & b`, env, reflect.TypeOf(a & b))
}

// Test Int & Float32
func TestCheckBinaryNonConstExprIntAndFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types int and float32)`,
	)

}

// Test Int & Complex128
func TestCheckBinaryNonConstExprIntAndComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types int and complex128)`,
	)

}

// Test Int & String
func TestCheckBinaryNonConstExprIntAndString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types int and string)`,
	)

}

// Test Int & BoolT
func TestCheckBinaryNonConstExprIntAndBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types int and bool)`,
	)

}

// Test Int & Slice
func TestCheckBinaryNonConstExprIntAndSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types int and eval.sliceT)`,
	)

}

// Test Int & Array
func TestCheckBinaryNonConstExprIntAndArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types int and eval.arrayT)`,
	)

}

// Test Int & Map
func TestCheckBinaryNonConstExprIntAndMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types int and eval.mapT)`,
	)

}

// Test Int & XinterfaceX
func TestCheckBinaryNonConstExprIntAndXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types int and eval.XinterfaceX)`,
	)

}

// Test Int & InterfaceX
func TestCheckBinaryNonConstExprIntAndInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types int and eval.interfaceX)`,
	)

}

// Test Int & InterfaceY
func TestCheckBinaryNonConstExprIntAndInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types int and eval.interfaceY)`,
	)

}

// Test Int & InterfaceZ
func TestCheckBinaryNonConstExprIntAndInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types int and eval.interfaceZ)`,
	)

}

// Test Int & Ptr
func TestCheckBinaryNonConstExprIntAndPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types int and *int)`,
	)

}

// Test Int & Struct
func TestCheckBinaryNonConstExprIntAndStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types int and eval.structT)`,
	)

}

// Test Int & StructUncomp
func TestCheckBinaryNonConstExprIntAndStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types int and eval.structUncompT)`,
	)

}

// Test Int % ConstInt
func TestCheckBinaryNonConstExprIntRemConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a % 4`, env, reflect.TypeOf(a % 4))
}

// Test Int % ConstRune
func TestCheckBinaryNonConstExprIntRemConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a % '@'`, env, reflect.TypeOf(a % '@'))
}

// Test Int % ConstFloat
func TestCheckBinaryNonConstExprIntRemConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a % 2.0`, env, reflect.TypeOf(a % 2.0))
}

// Test Int % ConstComplex
func TestCheckBinaryNonConstExprIntRemConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 8.0i`, env,
		`constant 0+8i truncated to real`,
		`division by zero`,
	)

}

// Test Int % ConstBool
func TestCheckBinaryNonConstExprIntRemConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % true`, env,
		`cannot convert true to type int`,
		`invalid operation: a % true (mismatched types int and bool)`,
	)

}

// Test Int % ConstString
func TestCheckBinaryNonConstExprIntRemConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: a % "abc" (mismatched types int and string)`,
	)

}

// Test Int % ConstNil
func TestCheckBinaryNonConstExprIntRemConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % nil`, env,
		`cannot convert nil to type int`,
	)

}

// Test Int % Int
func TestCheckBinaryNonConstExprIntRemInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a % b`, env, reflect.TypeOf(a % b))
}

// Test Int % Float32
func TestCheckBinaryNonConstExprIntRemFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types int and float32)`,
	)

}

// Test Int % Complex128
func TestCheckBinaryNonConstExprIntRemComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types int and complex128)`,
	)

}

// Test Int % String
func TestCheckBinaryNonConstExprIntRemString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types int and string)`,
	)

}

// Test Int % BoolT
func TestCheckBinaryNonConstExprIntRemBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types int and bool)`,
	)

}

// Test Int % Slice
func TestCheckBinaryNonConstExprIntRemSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types int and eval.sliceT)`,
	)

}

// Test Int % Array
func TestCheckBinaryNonConstExprIntRemArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types int and eval.arrayT)`,
	)

}

// Test Int % Map
func TestCheckBinaryNonConstExprIntRemMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types int and eval.mapT)`,
	)

}

// Test Int % XinterfaceX
func TestCheckBinaryNonConstExprIntRemXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types int and eval.XinterfaceX)`,
	)

}

// Test Int % InterfaceX
func TestCheckBinaryNonConstExprIntRemInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types int and eval.interfaceX)`,
	)

}

// Test Int % InterfaceY
func TestCheckBinaryNonConstExprIntRemInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types int and eval.interfaceY)`,
	)

}

// Test Int % InterfaceZ
func TestCheckBinaryNonConstExprIntRemInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types int and eval.interfaceZ)`,
	)

}

// Test Int % Ptr
func TestCheckBinaryNonConstExprIntRemPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types int and *int)`,
	)

}

// Test Int % Struct
func TestCheckBinaryNonConstExprIntRemStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types int and eval.structT)`,
	)

}

// Test Int % StructUncomp
func TestCheckBinaryNonConstExprIntRemStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types int and eval.structUncompT)`,
	)

}

// Test Int == ConstInt
func TestCheckBinaryNonConstExprIntEqlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == 4`, env, reflect.TypeOf(a == 4))
}

// Test Int == ConstRune
func TestCheckBinaryNonConstExprIntEqlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == '@'`, env, reflect.TypeOf(a == '@'))
}

// Test Int == ConstFloat
func TestCheckBinaryNonConstExprIntEqlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == 2.0`, env, reflect.TypeOf(a == 2.0))
}

// Test Int == ConstComplex
func TestCheckBinaryNonConstExprIntEqlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int == ConstBool
func TestCheckBinaryNonConstExprIntEqlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == true`, env,
		`cannot convert true to type int`,
		`invalid operation: a == true (mismatched types int and bool)`,
	)

}

// Test Int == ConstString
func TestCheckBinaryNonConstExprIntEqlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: a == "abc" (mismatched types int and string)`,
	)

}

// Test Int == ConstNil
func TestCheckBinaryNonConstExprIntEqlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == nil`, env,
		`cannot convert nil to type int`,
	)

}

// Test Int == Int
func TestCheckBinaryNonConstExprIntEqlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test Int == Float32
func TestCheckBinaryNonConstExprIntEqlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types int and float32)`,
	)

}

// Test Int == Complex128
func TestCheckBinaryNonConstExprIntEqlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types int and complex128)`,
	)

}

// Test Int == String
func TestCheckBinaryNonConstExprIntEqlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types int and string)`,
	)

}

// Test Int == BoolT
func TestCheckBinaryNonConstExprIntEqlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types int and bool)`,
	)

}

// Test Int == Slice
func TestCheckBinaryNonConstExprIntEqlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types int and eval.sliceT)`,
	)

}

// Test Int == Array
func TestCheckBinaryNonConstExprIntEqlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types int and eval.arrayT)`,
	)

}

// Test Int == Map
func TestCheckBinaryNonConstExprIntEqlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types int and eval.mapT)`,
	)

}

// Test Int == XinterfaceX
func TestCheckBinaryNonConstExprIntEqlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types int and eval.XinterfaceX)`,
	)

}

// Test Int == InterfaceX
func TestCheckBinaryNonConstExprIntEqlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types int and eval.interfaceX)`,
	)

}

// Test Int == InterfaceY
func TestCheckBinaryNonConstExprIntEqlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types int and eval.interfaceY)`,
	)

}

// Test Int == InterfaceZ
func TestCheckBinaryNonConstExprIntEqlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types int and eval.interfaceZ)`,
	)

}

// Test Int == Ptr
func TestCheckBinaryNonConstExprIntEqlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types int and *int)`,
	)

}

// Test Int == Struct
func TestCheckBinaryNonConstExprIntEqlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types int and eval.structT)`,
	)

}

// Test Int == StructUncomp
func TestCheckBinaryNonConstExprIntEqlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types int and eval.structUncompT)`,
	)

}

// Test Int > ConstInt
func TestCheckBinaryNonConstExprIntGtrConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a > 4`, env, reflect.TypeOf(a > 4))
}

// Test Int > ConstRune
func TestCheckBinaryNonConstExprIntGtrConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a > '@'`, env, reflect.TypeOf(a > '@'))
}

// Test Int > ConstFloat
func TestCheckBinaryNonConstExprIntGtrConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a > 2.0`, env, reflect.TypeOf(a > 2.0))
}

// Test Int > ConstComplex
func TestCheckBinaryNonConstExprIntGtrConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int > ConstBool
func TestCheckBinaryNonConstExprIntGtrConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > true`, env,
		`cannot convert true to type int`,
		`invalid operation: a > true (mismatched types int and bool)`,
	)

}

// Test Int > ConstString
func TestCheckBinaryNonConstExprIntGtrConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > "abc"`, env,
		`cannot convert "abc" to type int`,
		`invalid operation: a > "abc" (mismatched types int and string)`,
	)

}

// Test Int > ConstNil
func TestCheckBinaryNonConstExprIntGtrConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > nil`, env,
		`cannot convert nil to type int`,
	)

}

// Test Int > Int
func TestCheckBinaryNonConstExprIntGtrInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a > b`, env, reflect.TypeOf(a > b))
}

// Test Int > Float32
func TestCheckBinaryNonConstExprIntGtrFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types int and float32)`,
	)

}

// Test Int > Complex128
func TestCheckBinaryNonConstExprIntGtrComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types int and complex128)`,
	)

}

// Test Int > String
func TestCheckBinaryNonConstExprIntGtrString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types int and string)`,
	)

}

// Test Int > BoolT
func TestCheckBinaryNonConstExprIntGtrBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types int and bool)`,
	)

}

// Test Int > Slice
func TestCheckBinaryNonConstExprIntGtrSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types int and eval.sliceT)`,
	)

}

// Test Int > Array
func TestCheckBinaryNonConstExprIntGtrArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types int and eval.arrayT)`,
	)

}

// Test Int > Map
func TestCheckBinaryNonConstExprIntGtrMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types int and eval.mapT)`,
	)

}

// Test Int > XinterfaceX
func TestCheckBinaryNonConstExprIntGtrXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types int and eval.XinterfaceX)`,
	)

}

// Test Int > InterfaceX
func TestCheckBinaryNonConstExprIntGtrInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types int and eval.interfaceX)`,
	)

}

// Test Int > InterfaceY
func TestCheckBinaryNonConstExprIntGtrInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types int and eval.interfaceY)`,
	)

}

// Test Int > InterfaceZ
func TestCheckBinaryNonConstExprIntGtrInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types int and eval.interfaceZ)`,
	)

}

// Test Int > Ptr
func TestCheckBinaryNonConstExprIntGtrPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types int and *int)`,
	)

}

// Test Int > Struct
func TestCheckBinaryNonConstExprIntGtrStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types int and eval.structT)`,
	)

}

// Test Int > StructUncomp
func TestCheckBinaryNonConstExprIntGtrStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types int and eval.structUncompT)`,
	)

}

// Test Int << ConstInt
func TestCheckBinaryNonConstExprIntShlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a << 4`, env, reflect.TypeOf(a << 4))
}

// Test Int << ConstRune
func TestCheckBinaryNonConstExprIntShlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a << '@'`, env, reflect.TypeOf(a << '@'))
}

// Test Int << ConstFloat
func TestCheckBinaryNonConstExprIntShlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a << 2.0`, env, reflect.TypeOf(a << 2.0))
}

// Test Int << ConstComplex
func TestCheckBinaryNonConstExprIntShlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Int << ConstBool
func TestCheckBinaryNonConstExprIntShlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << true`, env,
		`invalid operation: a << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Int << ConstString
func TestCheckBinaryNonConstExprIntShlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: a << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Int << ConstNil
func TestCheckBinaryNonConstExprIntShlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Int << Int
func TestCheckBinaryNonConstExprIntShlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type int, must be unsigned integer)`,
	)

}

// Test Int << Float32
func TestCheckBinaryNonConstExprIntShlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type float32, must be unsigned integer)`,
	)

}

// Test Int << Complex128
func TestCheckBinaryNonConstExprIntShlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Int << String
func TestCheckBinaryNonConstExprIntShlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type string, must be unsigned integer)`,
	)

}

// Test Int << BoolT
func TestCheckBinaryNonConstExprIntShlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type bool, must be unsigned integer)`,
	)

}

// Test Int << Slice
func TestCheckBinaryNonConstExprIntShlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.sliceT, must be unsigned integer)`,
	)

}

// Test Int << Array
func TestCheckBinaryNonConstExprIntShlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.arrayT, must be unsigned integer)`,
	)

}

// Test Int << Map
func TestCheckBinaryNonConstExprIntShlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.mapT, must be unsigned integer)`,
	)

}

// Test Int << XinterfaceX
func TestCheckBinaryNonConstExprIntShlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.XinterfaceX, must be unsigned integer)`,
	)

}

// Test Int << InterfaceX
func TestCheckBinaryNonConstExprIntShlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceX, must be unsigned integer)`,
	)

}

// Test Int << InterfaceY
func TestCheckBinaryNonConstExprIntShlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceY, must be unsigned integer)`,
	)

}

// Test Int << InterfaceZ
func TestCheckBinaryNonConstExprIntShlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceZ, must be unsigned integer)`,
	)

}

// Test Int << Ptr
func TestCheckBinaryNonConstExprIntShlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type *int, must be unsigned integer)`,
	)

}

// Test Int << Struct
func TestCheckBinaryNonConstExprIntShlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structT, must be unsigned integer)`,
	)

}

// Test Int << StructUncomp
func TestCheckBinaryNonConstExprIntShlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := int(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structUncompT, must be unsigned integer)`,
	)

}

// Test Float32 + ConstInt
func TestCheckBinaryNonConstExprFloat32AddConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a + 4`, env, reflect.TypeOf(a + 4))
}

// Test Float32 + ConstRune
func TestCheckBinaryNonConstExprFloat32AddConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a + '@'`, env, reflect.TypeOf(a + '@'))
}

// Test Float32 + ConstFloat
func TestCheckBinaryNonConstExprFloat32AddConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a + 2.0`, env, reflect.TypeOf(a + 2.0))
}

// Test Float32 + ConstComplex
func TestCheckBinaryNonConstExprFloat32AddConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Float32 + ConstBool
func TestCheckBinaryNonConstExprFloat32AddConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + true`, env,
		`cannot convert true to type float32`,
		`invalid operation: a + true (mismatched types float32 and bool)`,
	)

}

// Test Float32 + ConstString
func TestCheckBinaryNonConstExprFloat32AddConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + "abc"`, env,
		`cannot convert "abc" to type float32`,
		`invalid operation: a + "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 + ConstNil
func TestCheckBinaryNonConstExprFloat32AddConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + nil`, env,
		`cannot convert nil to type float32`,
	)

}

// Test Float32 + Int
func TestCheckBinaryNonConstExprFloat32AddInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types float32 and int)`,
	)

}

// Test Float32 + Float32
func TestCheckBinaryNonConstExprFloat32AddFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a + b`, env, reflect.TypeOf(a + b))
}

// Test Float32 + Complex128
func TestCheckBinaryNonConstExprFloat32AddComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types float32 and complex128)`,
	)

}

// Test Float32 + String
func TestCheckBinaryNonConstExprFloat32AddString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types float32 and string)`,
	)

}

// Test Float32 + BoolT
func TestCheckBinaryNonConstExprFloat32AddBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types float32 and bool)`,
	)

}

// Test Float32 + Slice
func TestCheckBinaryNonConstExprFloat32AddSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types float32 and eval.sliceT)`,
	)

}

// Test Float32 + Array
func TestCheckBinaryNonConstExprFloat32AddArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types float32 and eval.arrayT)`,
	)

}

// Test Float32 + Map
func TestCheckBinaryNonConstExprFloat32AddMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types float32 and eval.mapT)`,
	)

}

// Test Float32 + XinterfaceX
func TestCheckBinaryNonConstExprFloat32AddXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types float32 and eval.XinterfaceX)`,
	)

}

// Test Float32 + InterfaceX
func TestCheckBinaryNonConstExprFloat32AddInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types float32 and eval.interfaceX)`,
	)

}

// Test Float32 + InterfaceY
func TestCheckBinaryNonConstExprFloat32AddInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types float32 and eval.interfaceY)`,
	)

}

// Test Float32 + InterfaceZ
func TestCheckBinaryNonConstExprFloat32AddInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types float32 and eval.interfaceZ)`,
	)

}

// Test Float32 + Ptr
func TestCheckBinaryNonConstExprFloat32AddPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types float32 and *int)`,
	)

}

// Test Float32 + Struct
func TestCheckBinaryNonConstExprFloat32AddStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types float32 and eval.structT)`,
	)

}

// Test Float32 + StructUncomp
func TestCheckBinaryNonConstExprFloat32AddStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types float32 and eval.structUncompT)`,
	)

}

// Test Float32 & ConstInt
func TestCheckBinaryNonConstExprFloat32AndConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 4`, env,
		`invalid operation: a & 4 (operator & not defined on float32)`,
	)

}

// Test Float32 & ConstRune
func TestCheckBinaryNonConstExprFloat32AndConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & '@'`, env,
		`invalid operation: a & 64 (operator & not defined on float32)`,
	)

}

// Test Float32 & ConstFloat
func TestCheckBinaryNonConstExprFloat32AndConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 2.0`, env,
		`invalid operation: a & 2 (operator & not defined on float32)`,
	)

}

// Test Float32 & ConstComplex
func TestCheckBinaryNonConstExprFloat32AndConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: a & 0 (operator & not defined on float32)`,
	)

}

// Test Float32 & ConstBool
func TestCheckBinaryNonConstExprFloat32AndConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & true`, env,
		`cannot convert true to type float32`,
		`invalid operation: a & true (mismatched types float32 and bool)`,
	)

}

// Test Float32 & ConstString
func TestCheckBinaryNonConstExprFloat32AndConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & "abc"`, env,
		`cannot convert "abc" to type float32`,
		`invalid operation: a & "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 & ConstNil
func TestCheckBinaryNonConstExprFloat32AndConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & nil`, env,
		`cannot convert nil to type float32`,
	)

}

// Test Float32 & Int
func TestCheckBinaryNonConstExprFloat32AndInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types float32 and int)`,
	)

}

// Test Float32 & Float32
func TestCheckBinaryNonConstExprFloat32AndFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (operator & not defined on float32)`,
	)

}

// Test Float32 & Complex128
func TestCheckBinaryNonConstExprFloat32AndComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types float32 and complex128)`,
	)

}

// Test Float32 & String
func TestCheckBinaryNonConstExprFloat32AndString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types float32 and string)`,
	)

}

// Test Float32 & BoolT
func TestCheckBinaryNonConstExprFloat32AndBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types float32 and bool)`,
	)

}

// Test Float32 & Slice
func TestCheckBinaryNonConstExprFloat32AndSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types float32 and eval.sliceT)`,
	)

}

// Test Float32 & Array
func TestCheckBinaryNonConstExprFloat32AndArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types float32 and eval.arrayT)`,
	)

}

// Test Float32 & Map
func TestCheckBinaryNonConstExprFloat32AndMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types float32 and eval.mapT)`,
	)

}

// Test Float32 & XinterfaceX
func TestCheckBinaryNonConstExprFloat32AndXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types float32 and eval.XinterfaceX)`,
	)

}

// Test Float32 & InterfaceX
func TestCheckBinaryNonConstExprFloat32AndInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types float32 and eval.interfaceX)`,
	)

}

// Test Float32 & InterfaceY
func TestCheckBinaryNonConstExprFloat32AndInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types float32 and eval.interfaceY)`,
	)

}

// Test Float32 & InterfaceZ
func TestCheckBinaryNonConstExprFloat32AndInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types float32 and eval.interfaceZ)`,
	)

}

// Test Float32 & Ptr
func TestCheckBinaryNonConstExprFloat32AndPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types float32 and *int)`,
	)

}

// Test Float32 & Struct
func TestCheckBinaryNonConstExprFloat32AndStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types float32 and eval.structT)`,
	)

}

// Test Float32 & StructUncomp
func TestCheckBinaryNonConstExprFloat32AndStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types float32 and eval.structUncompT)`,
	)

}

// Test Float32 % ConstInt
func TestCheckBinaryNonConstExprFloat32RemConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 4`, env,
		`invalid operation: a % 4 (operator % not defined on float32)`,
	)

}

// Test Float32 % ConstRune
func TestCheckBinaryNonConstExprFloat32RemConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % '@'`, env,
		`invalid operation: a % 64 (operator % not defined on float32)`,
	)

}

// Test Float32 % ConstFloat
func TestCheckBinaryNonConstExprFloat32RemConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 2.0`, env,
		`invalid operation: a % 2 (operator % not defined on float32)`,
	)

}

// Test Float32 % ConstComplex
func TestCheckBinaryNonConstExprFloat32RemConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: a % 0 (operator % not defined on float32)`,
	)

}

// Test Float32 % ConstBool
func TestCheckBinaryNonConstExprFloat32RemConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % true`, env,
		`cannot convert true to type float32`,
		`invalid operation: a % true (mismatched types float32 and bool)`,
	)

}

// Test Float32 % ConstString
func TestCheckBinaryNonConstExprFloat32RemConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % "abc"`, env,
		`cannot convert "abc" to type float32`,
		`invalid operation: a % "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 % ConstNil
func TestCheckBinaryNonConstExprFloat32RemConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % nil`, env,
		`cannot convert nil to type float32`,
	)

}

// Test Float32 % Int
func TestCheckBinaryNonConstExprFloat32RemInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types float32 and int)`,
	)

}

// Test Float32 % Float32
func TestCheckBinaryNonConstExprFloat32RemFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (operator % not defined on float32)`,
	)

}

// Test Float32 % Complex128
func TestCheckBinaryNonConstExprFloat32RemComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types float32 and complex128)`,
	)

}

// Test Float32 % String
func TestCheckBinaryNonConstExprFloat32RemString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types float32 and string)`,
	)

}

// Test Float32 % BoolT
func TestCheckBinaryNonConstExprFloat32RemBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types float32 and bool)`,
	)

}

// Test Float32 % Slice
func TestCheckBinaryNonConstExprFloat32RemSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types float32 and eval.sliceT)`,
	)

}

// Test Float32 % Array
func TestCheckBinaryNonConstExprFloat32RemArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types float32 and eval.arrayT)`,
	)

}

// Test Float32 % Map
func TestCheckBinaryNonConstExprFloat32RemMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types float32 and eval.mapT)`,
	)

}

// Test Float32 % XinterfaceX
func TestCheckBinaryNonConstExprFloat32RemXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types float32 and eval.XinterfaceX)`,
	)

}

// Test Float32 % InterfaceX
func TestCheckBinaryNonConstExprFloat32RemInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types float32 and eval.interfaceX)`,
	)

}

// Test Float32 % InterfaceY
func TestCheckBinaryNonConstExprFloat32RemInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types float32 and eval.interfaceY)`,
	)

}

// Test Float32 % InterfaceZ
func TestCheckBinaryNonConstExprFloat32RemInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types float32 and eval.interfaceZ)`,
	)

}

// Test Float32 % Ptr
func TestCheckBinaryNonConstExprFloat32RemPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types float32 and *int)`,
	)

}

// Test Float32 % Struct
func TestCheckBinaryNonConstExprFloat32RemStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types float32 and eval.structT)`,
	)

}

// Test Float32 % StructUncomp
func TestCheckBinaryNonConstExprFloat32RemStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types float32 and eval.structUncompT)`,
	)

}

// Test Float32 == ConstInt
func TestCheckBinaryNonConstExprFloat32EqlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == 4`, env, reflect.TypeOf(a == 4))
}

// Test Float32 == ConstRune
func TestCheckBinaryNonConstExprFloat32EqlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == '@'`, env, reflect.TypeOf(a == '@'))
}

// Test Float32 == ConstFloat
func TestCheckBinaryNonConstExprFloat32EqlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == 2.0`, env, reflect.TypeOf(a == 2.0))
}

// Test Float32 == ConstComplex
func TestCheckBinaryNonConstExprFloat32EqlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Float32 == ConstBool
func TestCheckBinaryNonConstExprFloat32EqlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == true`, env,
		`cannot convert true to type float32`,
		`invalid operation: a == true (mismatched types float32 and bool)`,
	)

}

// Test Float32 == ConstString
func TestCheckBinaryNonConstExprFloat32EqlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == "abc"`, env,
		`cannot convert "abc" to type float32`,
		`invalid operation: a == "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 == ConstNil
func TestCheckBinaryNonConstExprFloat32EqlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == nil`, env,
		`cannot convert nil to type float32`,
	)

}

// Test Float32 == Int
func TestCheckBinaryNonConstExprFloat32EqlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types float32 and int)`,
	)

}

// Test Float32 == Float32
func TestCheckBinaryNonConstExprFloat32EqlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test Float32 == Complex128
func TestCheckBinaryNonConstExprFloat32EqlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types float32 and complex128)`,
	)

}

// Test Float32 == String
func TestCheckBinaryNonConstExprFloat32EqlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types float32 and string)`,
	)

}

// Test Float32 == BoolT
func TestCheckBinaryNonConstExprFloat32EqlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types float32 and bool)`,
	)

}

// Test Float32 == Slice
func TestCheckBinaryNonConstExprFloat32EqlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types float32 and eval.sliceT)`,
	)

}

// Test Float32 == Array
func TestCheckBinaryNonConstExprFloat32EqlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types float32 and eval.arrayT)`,
	)

}

// Test Float32 == Map
func TestCheckBinaryNonConstExprFloat32EqlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types float32 and eval.mapT)`,
	)

}

// Test Float32 == XinterfaceX
func TestCheckBinaryNonConstExprFloat32EqlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types float32 and eval.XinterfaceX)`,
	)

}

// Test Float32 == InterfaceX
func TestCheckBinaryNonConstExprFloat32EqlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types float32 and eval.interfaceX)`,
	)

}

// Test Float32 == InterfaceY
func TestCheckBinaryNonConstExprFloat32EqlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types float32 and eval.interfaceY)`,
	)

}

// Test Float32 == InterfaceZ
func TestCheckBinaryNonConstExprFloat32EqlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types float32 and eval.interfaceZ)`,
	)

}

// Test Float32 == Ptr
func TestCheckBinaryNonConstExprFloat32EqlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types float32 and *int)`,
	)

}

// Test Float32 == Struct
func TestCheckBinaryNonConstExprFloat32EqlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types float32 and eval.structT)`,
	)

}

// Test Float32 == StructUncomp
func TestCheckBinaryNonConstExprFloat32EqlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types float32 and eval.structUncompT)`,
	)

}

// Test Float32 > ConstInt
func TestCheckBinaryNonConstExprFloat32GtrConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a > 4`, env, reflect.TypeOf(a > 4))
}

// Test Float32 > ConstRune
func TestCheckBinaryNonConstExprFloat32GtrConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a > '@'`, env, reflect.TypeOf(a > '@'))
}

// Test Float32 > ConstFloat
func TestCheckBinaryNonConstExprFloat32GtrConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a > 2.0`, env, reflect.TypeOf(a > 2.0))
}

// Test Float32 > ConstComplex
func TestCheckBinaryNonConstExprFloat32GtrConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test Float32 > ConstBool
func TestCheckBinaryNonConstExprFloat32GtrConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > true`, env,
		`cannot convert true to type float32`,
		`invalid operation: a > true (mismatched types float32 and bool)`,
	)

}

// Test Float32 > ConstString
func TestCheckBinaryNonConstExprFloat32GtrConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > "abc"`, env,
		`cannot convert "abc" to type float32`,
		`invalid operation: a > "abc" (mismatched types float32 and string)`,
	)

}

// Test Float32 > ConstNil
func TestCheckBinaryNonConstExprFloat32GtrConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > nil`, env,
		`cannot convert nil to type float32`,
	)

}

// Test Float32 > Int
func TestCheckBinaryNonConstExprFloat32GtrInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types float32 and int)`,
	)

}

// Test Float32 > Float32
func TestCheckBinaryNonConstExprFloat32GtrFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a > b`, env, reflect.TypeOf(a > b))
}

// Test Float32 > Complex128
func TestCheckBinaryNonConstExprFloat32GtrComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types float32 and complex128)`,
	)

}

// Test Float32 > String
func TestCheckBinaryNonConstExprFloat32GtrString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types float32 and string)`,
	)

}

// Test Float32 > BoolT
func TestCheckBinaryNonConstExprFloat32GtrBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types float32 and bool)`,
	)

}

// Test Float32 > Slice
func TestCheckBinaryNonConstExprFloat32GtrSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types float32 and eval.sliceT)`,
	)

}

// Test Float32 > Array
func TestCheckBinaryNonConstExprFloat32GtrArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types float32 and eval.arrayT)`,
	)

}

// Test Float32 > Map
func TestCheckBinaryNonConstExprFloat32GtrMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types float32 and eval.mapT)`,
	)

}

// Test Float32 > XinterfaceX
func TestCheckBinaryNonConstExprFloat32GtrXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types float32 and eval.XinterfaceX)`,
	)

}

// Test Float32 > InterfaceX
func TestCheckBinaryNonConstExprFloat32GtrInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types float32 and eval.interfaceX)`,
	)

}

// Test Float32 > InterfaceY
func TestCheckBinaryNonConstExprFloat32GtrInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types float32 and eval.interfaceY)`,
	)

}

// Test Float32 > InterfaceZ
func TestCheckBinaryNonConstExprFloat32GtrInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types float32 and eval.interfaceZ)`,
	)

}

// Test Float32 > Ptr
func TestCheckBinaryNonConstExprFloat32GtrPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types float32 and *int)`,
	)

}

// Test Float32 > Struct
func TestCheckBinaryNonConstExprFloat32GtrStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types float32 and eval.structT)`,
	)

}

// Test Float32 > StructUncomp
func TestCheckBinaryNonConstExprFloat32GtrStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types float32 and eval.structUncompT)`,
	)

}

// Test Float32 << ConstInt
func TestCheckBinaryNonConstExprFloat32ShlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 4`, env,
		`invalid operation: a << 4 (shift of type float32)`,
	)

}

// Test Float32 << ConstRune
func TestCheckBinaryNonConstExprFloat32ShlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << '@'`, env,
		`invalid operation: a << 64 (shift of type float32)`,
	)

}

// Test Float32 << ConstFloat
func TestCheckBinaryNonConstExprFloat32ShlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 2.0`, env,
		`invalid operation: a << 2 (shift of type float32)`,
	)

}

// Test Float32 << ConstComplex
func TestCheckBinaryNonConstExprFloat32ShlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: a << 0 (shift of type float32)`,
	)

}

// Test Float32 << ConstBool
func TestCheckBinaryNonConstExprFloat32ShlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << true`, env,
		`invalid operation: a << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Float32 << ConstString
func TestCheckBinaryNonConstExprFloat32ShlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: a << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Float32 << ConstNil
func TestCheckBinaryNonConstExprFloat32ShlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Float32 << Int
func TestCheckBinaryNonConstExprFloat32ShlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type int, must be unsigned integer)`,
	)

}

// Test Float32 << Float32
func TestCheckBinaryNonConstExprFloat32ShlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type float32, must be unsigned integer)`,
	)

}

// Test Float32 << Complex128
func TestCheckBinaryNonConstExprFloat32ShlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Float32 << String
func TestCheckBinaryNonConstExprFloat32ShlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type string, must be unsigned integer)`,
	)

}

// Test Float32 << BoolT
func TestCheckBinaryNonConstExprFloat32ShlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type bool, must be unsigned integer)`,
	)

}

// Test Float32 << Slice
func TestCheckBinaryNonConstExprFloat32ShlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.sliceT, must be unsigned integer)`,
	)

}

// Test Float32 << Array
func TestCheckBinaryNonConstExprFloat32ShlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.arrayT, must be unsigned integer)`,
	)

}

// Test Float32 << Map
func TestCheckBinaryNonConstExprFloat32ShlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.mapT, must be unsigned integer)`,
	)

}

// Test Float32 << XinterfaceX
func TestCheckBinaryNonConstExprFloat32ShlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.XinterfaceX, must be unsigned integer)`,
	)

}

// Test Float32 << InterfaceX
func TestCheckBinaryNonConstExprFloat32ShlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceX, must be unsigned integer)`,
	)

}

// Test Float32 << InterfaceY
func TestCheckBinaryNonConstExprFloat32ShlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceY, must be unsigned integer)`,
	)

}

// Test Float32 << InterfaceZ
func TestCheckBinaryNonConstExprFloat32ShlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceZ, must be unsigned integer)`,
	)

}

// Test Float32 << Ptr
func TestCheckBinaryNonConstExprFloat32ShlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type *int, must be unsigned integer)`,
	)

}

// Test Float32 << Struct
func TestCheckBinaryNonConstExprFloat32ShlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structT, must be unsigned integer)`,
	)

}

// Test Float32 << StructUncomp
func TestCheckBinaryNonConstExprFloat32ShlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := float32(1.5); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structUncompT, must be unsigned integer)`,
	)

}

// Test Complex128 + ConstInt
func TestCheckBinaryNonConstExprComplex128AddConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a + 4`, env, reflect.TypeOf(a + 4))
}

// Test Complex128 + ConstRune
func TestCheckBinaryNonConstExprComplex128AddConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a + '@'`, env, reflect.TypeOf(a + '@'))
}

// Test Complex128 + ConstFloat
func TestCheckBinaryNonConstExprComplex128AddConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a + 2.0`, env, reflect.TypeOf(a + 2.0))
}

// Test Complex128 + ConstComplex
func TestCheckBinaryNonConstExprComplex128AddConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a + 8.0i`, env, reflect.TypeOf(a + 8.0i))
}

// Test Complex128 + ConstBool
func TestCheckBinaryNonConstExprComplex128AddConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: a + true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 + ConstString
func TestCheckBinaryNonConstExprComplex128AddConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: a + "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 + ConstNil
func TestCheckBinaryNonConstExprComplex128AddConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + nil`, env,
		`cannot convert nil to type complex128`,
	)

}

// Test Complex128 + Int
func TestCheckBinaryNonConstExprComplex128AddInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types complex128 and int)`,
	)

}

// Test Complex128 + Float32
func TestCheckBinaryNonConstExprComplex128AddFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types complex128 and float32)`,
	)

}

// Test Complex128 + Complex128
func TestCheckBinaryNonConstExprComplex128AddComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a + b`, env, reflect.TypeOf(a + b))
}

// Test Complex128 + String
func TestCheckBinaryNonConstExprComplex128AddString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types complex128 and string)`,
	)

}

// Test Complex128 + BoolT
func TestCheckBinaryNonConstExprComplex128AddBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 + Slice
func TestCheckBinaryNonConstExprComplex128AddSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types complex128 and eval.sliceT)`,
	)

}

// Test Complex128 + Array
func TestCheckBinaryNonConstExprComplex128AddArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types complex128 and eval.arrayT)`,
	)

}

// Test Complex128 + Map
func TestCheckBinaryNonConstExprComplex128AddMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types complex128 and eval.mapT)`,
	)

}

// Test Complex128 + XinterfaceX
func TestCheckBinaryNonConstExprComplex128AddXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types complex128 and eval.XinterfaceX)`,
	)

}

// Test Complex128 + InterfaceX
func TestCheckBinaryNonConstExprComplex128AddInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types complex128 and eval.interfaceX)`,
	)

}

// Test Complex128 + InterfaceY
func TestCheckBinaryNonConstExprComplex128AddInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types complex128 and eval.interfaceY)`,
	)

}

// Test Complex128 + InterfaceZ
func TestCheckBinaryNonConstExprComplex128AddInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types complex128 and eval.interfaceZ)`,
	)

}

// Test Complex128 + Ptr
func TestCheckBinaryNonConstExprComplex128AddPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types complex128 and *int)`,
	)

}

// Test Complex128 + Struct
func TestCheckBinaryNonConstExprComplex128AddStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types complex128 and eval.structT)`,
	)

}

// Test Complex128 + StructUncomp
func TestCheckBinaryNonConstExprComplex128AddStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types complex128 and eval.structUncompT)`,
	)

}

// Test Complex128 & ConstInt
func TestCheckBinaryNonConstExprComplex128AndConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 4`, env,
		`invalid operation: a & 4 (operator & not defined on complex128)`,
	)

}

// Test Complex128 & ConstRune
func TestCheckBinaryNonConstExprComplex128AndConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & '@'`, env,
		`invalid operation: a & 64 (operator & not defined on complex128)`,
	)

}

// Test Complex128 & ConstFloat
func TestCheckBinaryNonConstExprComplex128AndConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 2.0`, env,
		`invalid operation: a & 2 (operator & not defined on complex128)`,
	)

}

// Test Complex128 & ConstComplex
func TestCheckBinaryNonConstExprComplex128AndConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 8.0i`, env,
		`invalid operation: a & 8i (operator & not defined on complex128)`,
	)

}

// Test Complex128 & ConstBool
func TestCheckBinaryNonConstExprComplex128AndConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: a & true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 & ConstString
func TestCheckBinaryNonConstExprComplex128AndConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: a & "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 & ConstNil
func TestCheckBinaryNonConstExprComplex128AndConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & nil`, env,
		`cannot convert nil to type complex128`,
	)

}

// Test Complex128 & Int
func TestCheckBinaryNonConstExprComplex128AndInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types complex128 and int)`,
	)

}

// Test Complex128 & Float32
func TestCheckBinaryNonConstExprComplex128AndFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types complex128 and float32)`,
	)

}

// Test Complex128 & Complex128
func TestCheckBinaryNonConstExprComplex128AndComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (operator & not defined on complex128)`,
	)

}

// Test Complex128 & String
func TestCheckBinaryNonConstExprComplex128AndString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types complex128 and string)`,
	)

}

// Test Complex128 & BoolT
func TestCheckBinaryNonConstExprComplex128AndBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 & Slice
func TestCheckBinaryNonConstExprComplex128AndSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types complex128 and eval.sliceT)`,
	)

}

// Test Complex128 & Array
func TestCheckBinaryNonConstExprComplex128AndArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types complex128 and eval.arrayT)`,
	)

}

// Test Complex128 & Map
func TestCheckBinaryNonConstExprComplex128AndMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types complex128 and eval.mapT)`,
	)

}

// Test Complex128 & XinterfaceX
func TestCheckBinaryNonConstExprComplex128AndXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types complex128 and eval.XinterfaceX)`,
	)

}

// Test Complex128 & InterfaceX
func TestCheckBinaryNonConstExprComplex128AndInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types complex128 and eval.interfaceX)`,
	)

}

// Test Complex128 & InterfaceY
func TestCheckBinaryNonConstExprComplex128AndInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types complex128 and eval.interfaceY)`,
	)

}

// Test Complex128 & InterfaceZ
func TestCheckBinaryNonConstExprComplex128AndInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types complex128 and eval.interfaceZ)`,
	)

}

// Test Complex128 & Ptr
func TestCheckBinaryNonConstExprComplex128AndPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types complex128 and *int)`,
	)

}

// Test Complex128 & Struct
func TestCheckBinaryNonConstExprComplex128AndStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types complex128 and eval.structT)`,
	)

}

// Test Complex128 & StructUncomp
func TestCheckBinaryNonConstExprComplex128AndStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types complex128 and eval.structUncompT)`,
	)

}

// Test Complex128 % ConstInt
func TestCheckBinaryNonConstExprComplex128RemConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 4`, env,
		`invalid operation: a % 4 (operator % not defined on complex128)`,
	)

}

// Test Complex128 % ConstRune
func TestCheckBinaryNonConstExprComplex128RemConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % '@'`, env,
		`invalid operation: a % 64 (operator % not defined on complex128)`,
	)

}

// Test Complex128 % ConstFloat
func TestCheckBinaryNonConstExprComplex128RemConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 2.0`, env,
		`invalid operation: a % 2 (operator % not defined on complex128)`,
	)

}

// Test Complex128 % ConstComplex
func TestCheckBinaryNonConstExprComplex128RemConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 8.0i`, env,
		`invalid operation: a % 8i (operator % not defined on complex128)`,
	)

}

// Test Complex128 % ConstBool
func TestCheckBinaryNonConstExprComplex128RemConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: a % true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 % ConstString
func TestCheckBinaryNonConstExprComplex128RemConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: a % "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 % ConstNil
func TestCheckBinaryNonConstExprComplex128RemConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % nil`, env,
		`cannot convert nil to type complex128`,
	)

}

// Test Complex128 % Int
func TestCheckBinaryNonConstExprComplex128RemInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types complex128 and int)`,
	)

}

// Test Complex128 % Float32
func TestCheckBinaryNonConstExprComplex128RemFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types complex128 and float32)`,
	)

}

// Test Complex128 % Complex128
func TestCheckBinaryNonConstExprComplex128RemComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (operator % not defined on complex128)`,
	)

}

// Test Complex128 % String
func TestCheckBinaryNonConstExprComplex128RemString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types complex128 and string)`,
	)

}

// Test Complex128 % BoolT
func TestCheckBinaryNonConstExprComplex128RemBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 % Slice
func TestCheckBinaryNonConstExprComplex128RemSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types complex128 and eval.sliceT)`,
	)

}

// Test Complex128 % Array
func TestCheckBinaryNonConstExprComplex128RemArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types complex128 and eval.arrayT)`,
	)

}

// Test Complex128 % Map
func TestCheckBinaryNonConstExprComplex128RemMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types complex128 and eval.mapT)`,
	)

}

// Test Complex128 % XinterfaceX
func TestCheckBinaryNonConstExprComplex128RemXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types complex128 and eval.XinterfaceX)`,
	)

}

// Test Complex128 % InterfaceX
func TestCheckBinaryNonConstExprComplex128RemInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types complex128 and eval.interfaceX)`,
	)

}

// Test Complex128 % InterfaceY
func TestCheckBinaryNonConstExprComplex128RemInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types complex128 and eval.interfaceY)`,
	)

}

// Test Complex128 % InterfaceZ
func TestCheckBinaryNonConstExprComplex128RemInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types complex128 and eval.interfaceZ)`,
	)

}

// Test Complex128 % Ptr
func TestCheckBinaryNonConstExprComplex128RemPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types complex128 and *int)`,
	)

}

// Test Complex128 % Struct
func TestCheckBinaryNonConstExprComplex128RemStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types complex128 and eval.structT)`,
	)

}

// Test Complex128 % StructUncomp
func TestCheckBinaryNonConstExprComplex128RemStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types complex128 and eval.structUncompT)`,
	)

}

// Test Complex128 == ConstInt
func TestCheckBinaryNonConstExprComplex128EqlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == 4`, env, reflect.TypeOf(a == 4))
}

// Test Complex128 == ConstRune
func TestCheckBinaryNonConstExprComplex128EqlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == '@'`, env, reflect.TypeOf(a == '@'))
}

// Test Complex128 == ConstFloat
func TestCheckBinaryNonConstExprComplex128EqlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == 2.0`, env, reflect.TypeOf(a == 2.0))
}

// Test Complex128 == ConstComplex
func TestCheckBinaryNonConstExprComplex128EqlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == 8.0i`, env, reflect.TypeOf(a == 8.0i))
}

// Test Complex128 == ConstBool
func TestCheckBinaryNonConstExprComplex128EqlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: a == true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 == ConstString
func TestCheckBinaryNonConstExprComplex128EqlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: a == "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 == ConstNil
func TestCheckBinaryNonConstExprComplex128EqlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == nil`, env,
		`cannot convert nil to type complex128`,
	)

}

// Test Complex128 == Int
func TestCheckBinaryNonConstExprComplex128EqlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types complex128 and int)`,
	)

}

// Test Complex128 == Float32
func TestCheckBinaryNonConstExprComplex128EqlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types complex128 and float32)`,
	)

}

// Test Complex128 == Complex128
func TestCheckBinaryNonConstExprComplex128EqlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test Complex128 == String
func TestCheckBinaryNonConstExprComplex128EqlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types complex128 and string)`,
	)

}

// Test Complex128 == BoolT
func TestCheckBinaryNonConstExprComplex128EqlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 == Slice
func TestCheckBinaryNonConstExprComplex128EqlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types complex128 and eval.sliceT)`,
	)

}

// Test Complex128 == Array
func TestCheckBinaryNonConstExprComplex128EqlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types complex128 and eval.arrayT)`,
	)

}

// Test Complex128 == Map
func TestCheckBinaryNonConstExprComplex128EqlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types complex128 and eval.mapT)`,
	)

}

// Test Complex128 == XinterfaceX
func TestCheckBinaryNonConstExprComplex128EqlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types complex128 and eval.XinterfaceX)`,
	)

}

// Test Complex128 == InterfaceX
func TestCheckBinaryNonConstExprComplex128EqlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types complex128 and eval.interfaceX)`,
	)

}

// Test Complex128 == InterfaceY
func TestCheckBinaryNonConstExprComplex128EqlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types complex128 and eval.interfaceY)`,
	)

}

// Test Complex128 == InterfaceZ
func TestCheckBinaryNonConstExprComplex128EqlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types complex128 and eval.interfaceZ)`,
	)

}

// Test Complex128 == Ptr
func TestCheckBinaryNonConstExprComplex128EqlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types complex128 and *int)`,
	)

}

// Test Complex128 == Struct
func TestCheckBinaryNonConstExprComplex128EqlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types complex128 and eval.structT)`,
	)

}

// Test Complex128 == StructUncomp
func TestCheckBinaryNonConstExprComplex128EqlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types complex128 and eval.structUncompT)`,
	)

}

// Test Complex128 > ConstInt
func TestCheckBinaryNonConstExprComplex128GtrConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 4`, env,
		`invalid operation: a > 4 (operator > not defined on complex128)`,
	)

}

// Test Complex128 > ConstRune
func TestCheckBinaryNonConstExprComplex128GtrConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > '@'`, env,
		`invalid operation: a > 64 (operator > not defined on complex128)`,
	)

}

// Test Complex128 > ConstFloat
func TestCheckBinaryNonConstExprComplex128GtrConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 2.0`, env,
		`invalid operation: a > 2 (operator > not defined on complex128)`,
	)

}

// Test Complex128 > ConstComplex
func TestCheckBinaryNonConstExprComplex128GtrConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 8.0i`, env,
		`invalid operation: a > 8i (operator > not defined on complex128)`,
	)

}

// Test Complex128 > ConstBool
func TestCheckBinaryNonConstExprComplex128GtrConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > true`, env,
		`cannot convert true to type complex128`,
		`invalid operation: a > true (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 > ConstString
func TestCheckBinaryNonConstExprComplex128GtrConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > "abc"`, env,
		`cannot convert "abc" to type complex128`,
		`invalid operation: a > "abc" (mismatched types complex128 and string)`,
	)

}

// Test Complex128 > ConstNil
func TestCheckBinaryNonConstExprComplex128GtrConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > nil`, env,
		`cannot convert nil to type complex128`,
	)

}

// Test Complex128 > Int
func TestCheckBinaryNonConstExprComplex128GtrInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types complex128 and int)`,
	)

}

// Test Complex128 > Float32
func TestCheckBinaryNonConstExprComplex128GtrFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types complex128 and float32)`,
	)

}

// Test Complex128 > Complex128
func TestCheckBinaryNonConstExprComplex128GtrComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (operator > not defined on complex128)`,
	)

}

// Test Complex128 > String
func TestCheckBinaryNonConstExprComplex128GtrString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types complex128 and string)`,
	)

}

// Test Complex128 > BoolT
func TestCheckBinaryNonConstExprComplex128GtrBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types complex128 and bool)`,
	)

}

// Test Complex128 > Slice
func TestCheckBinaryNonConstExprComplex128GtrSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types complex128 and eval.sliceT)`,
	)

}

// Test Complex128 > Array
func TestCheckBinaryNonConstExprComplex128GtrArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types complex128 and eval.arrayT)`,
	)

}

// Test Complex128 > Map
func TestCheckBinaryNonConstExprComplex128GtrMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types complex128 and eval.mapT)`,
	)

}

// Test Complex128 > XinterfaceX
func TestCheckBinaryNonConstExprComplex128GtrXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types complex128 and eval.XinterfaceX)`,
	)

}

// Test Complex128 > InterfaceX
func TestCheckBinaryNonConstExprComplex128GtrInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types complex128 and eval.interfaceX)`,
	)

}

// Test Complex128 > InterfaceY
func TestCheckBinaryNonConstExprComplex128GtrInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types complex128 and eval.interfaceY)`,
	)

}

// Test Complex128 > InterfaceZ
func TestCheckBinaryNonConstExprComplex128GtrInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types complex128 and eval.interfaceZ)`,
	)

}

// Test Complex128 > Ptr
func TestCheckBinaryNonConstExprComplex128GtrPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types complex128 and *int)`,
	)

}

// Test Complex128 > Struct
func TestCheckBinaryNonConstExprComplex128GtrStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types complex128 and eval.structT)`,
	)

}

// Test Complex128 > StructUncomp
func TestCheckBinaryNonConstExprComplex128GtrStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types complex128 and eval.structUncompT)`,
	)

}

// Test Complex128 << ConstInt
func TestCheckBinaryNonConstExprComplex128ShlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 4`, env,
		`invalid operation: a << 4 (shift of type complex128)`,
	)

}

// Test Complex128 << ConstRune
func TestCheckBinaryNonConstExprComplex128ShlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << '@'`, env,
		`invalid operation: a << 64 (shift of type complex128)`,
	)

}

// Test Complex128 << ConstFloat
func TestCheckBinaryNonConstExprComplex128ShlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 2.0`, env,
		`invalid operation: a << 2 (shift of type complex128)`,
	)

}

// Test Complex128 << ConstComplex
func TestCheckBinaryNonConstExprComplex128ShlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: a << 0 (shift of type complex128)`,
	)

}

// Test Complex128 << ConstBool
func TestCheckBinaryNonConstExprComplex128ShlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << true`, env,
		`invalid operation: a << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Complex128 << ConstString
func TestCheckBinaryNonConstExprComplex128ShlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: a << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Complex128 << ConstNil
func TestCheckBinaryNonConstExprComplex128ShlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Complex128 << Int
func TestCheckBinaryNonConstExprComplex128ShlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type int, must be unsigned integer)`,
	)

}

// Test Complex128 << Float32
func TestCheckBinaryNonConstExprComplex128ShlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type float32, must be unsigned integer)`,
	)

}

// Test Complex128 << Complex128
func TestCheckBinaryNonConstExprComplex128ShlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Complex128 << String
func TestCheckBinaryNonConstExprComplex128ShlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type string, must be unsigned integer)`,
	)

}

// Test Complex128 << BoolT
func TestCheckBinaryNonConstExprComplex128ShlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type bool, must be unsigned integer)`,
	)

}

// Test Complex128 << Slice
func TestCheckBinaryNonConstExprComplex128ShlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.sliceT, must be unsigned integer)`,
	)

}

// Test Complex128 << Array
func TestCheckBinaryNonConstExprComplex128ShlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.arrayT, must be unsigned integer)`,
	)

}

// Test Complex128 << Map
func TestCheckBinaryNonConstExprComplex128ShlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.mapT, must be unsigned integer)`,
	)

}

// Test Complex128 << XinterfaceX
func TestCheckBinaryNonConstExprComplex128ShlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.XinterfaceX, must be unsigned integer)`,
	)

}

// Test Complex128 << InterfaceX
func TestCheckBinaryNonConstExprComplex128ShlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceX, must be unsigned integer)`,
	)

}

// Test Complex128 << InterfaceY
func TestCheckBinaryNonConstExprComplex128ShlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceY, must be unsigned integer)`,
	)

}

// Test Complex128 << InterfaceZ
func TestCheckBinaryNonConstExprComplex128ShlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceZ, must be unsigned integer)`,
	)

}

// Test Complex128 << Ptr
func TestCheckBinaryNonConstExprComplex128ShlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type *int, must be unsigned integer)`,
	)

}

// Test Complex128 << Struct
func TestCheckBinaryNonConstExprComplex128ShlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structT, must be unsigned integer)`,
	)

}

// Test Complex128 << StructUncomp
func TestCheckBinaryNonConstExprComplex128ShlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := complex128(1i); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structUncompT, must be unsigned integer)`,
	)

}

// Test String + ConstInt
func TestCheckBinaryNonConstExprStringAddConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 4`, env,
		`cannot convert 4 to type string`,
		`invalid operation: a + 4 (mismatched types string and int)`,
	)

}

// Test String + ConstRune
func TestCheckBinaryNonConstExprStringAddConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + '@'`, env,
		`cannot convert '@' to type string`,
		`invalid operation: a + rune(64) (mismatched types string and rune)`,
	)

}

// Test String + ConstFloat
func TestCheckBinaryNonConstExprStringAddConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 2.0`, env,
		`cannot convert 2 to type string`,
		`invalid operation: a + 2 (mismatched types string and float64)`,
	)

}

// Test String + ConstComplex
func TestCheckBinaryNonConstExprStringAddConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 8.0i`, env,
		`cannot convert 8i to type string`,
		`invalid operation: a + 8i (mismatched types string and complex128)`,
	)

}

// Test String + ConstBool
func TestCheckBinaryNonConstExprStringAddConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + true`, env,
		`cannot convert true to type string`,
		`invalid operation: a + true (mismatched types string and bool)`,
	)

}

// Test String + ConstString
func TestCheckBinaryNonConstExprStringAddConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a + "abc"`, env, reflect.TypeOf(a + "abc"))
}

// Test String + ConstNil
func TestCheckBinaryNonConstExprStringAddConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + nil`, env,
		`invalid operation: a + nil (mismatched types string and nil)`,
	)

}

// Test String + Int
func TestCheckBinaryNonConstExprStringAddInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types string and int)`,
	)

}

// Test String + Float32
func TestCheckBinaryNonConstExprStringAddFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types string and float32)`,
	)

}

// Test String + Complex128
func TestCheckBinaryNonConstExprStringAddComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types string and complex128)`,
	)

}

// Test String + String
func TestCheckBinaryNonConstExprStringAddString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a + b`, env, reflect.TypeOf(a + b))
}

// Test String + BoolT
func TestCheckBinaryNonConstExprStringAddBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types string and bool)`,
	)

}

// Test String + Slice
func TestCheckBinaryNonConstExprStringAddSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types string and eval.sliceT)`,
	)

}

// Test String + Array
func TestCheckBinaryNonConstExprStringAddArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types string and eval.arrayT)`,
	)

}

// Test String + Map
func TestCheckBinaryNonConstExprStringAddMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types string and eval.mapT)`,
	)

}

// Test String + XinterfaceX
func TestCheckBinaryNonConstExprStringAddXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types string and eval.XinterfaceX)`,
	)

}

// Test String + InterfaceX
func TestCheckBinaryNonConstExprStringAddInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types string and eval.interfaceX)`,
	)

}

// Test String + InterfaceY
func TestCheckBinaryNonConstExprStringAddInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types string and eval.interfaceY)`,
	)

}

// Test String + InterfaceZ
func TestCheckBinaryNonConstExprStringAddInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types string and eval.interfaceZ)`,
	)

}

// Test String + Ptr
func TestCheckBinaryNonConstExprStringAddPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types string and *int)`,
	)

}

// Test String + Struct
func TestCheckBinaryNonConstExprStringAddStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types string and eval.structT)`,
	)

}

// Test String + StructUncomp
func TestCheckBinaryNonConstExprStringAddStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types string and eval.structUncompT)`,
	)

}

// Test String & ConstInt
func TestCheckBinaryNonConstExprStringAndConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 4`, env,
		`cannot convert 4 to type string`,
		`invalid operation: a & 4 (mismatched types string and int)`,
	)

}

// Test String & ConstRune
func TestCheckBinaryNonConstExprStringAndConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & '@'`, env,
		`cannot convert '@' to type string`,
		`invalid operation: a & rune(64) (mismatched types string and rune)`,
	)

}

// Test String & ConstFloat
func TestCheckBinaryNonConstExprStringAndConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 2.0`, env,
		`cannot convert 2 to type string`,
		`invalid operation: a & 2 (mismatched types string and float64)`,
	)

}

// Test String & ConstComplex
func TestCheckBinaryNonConstExprStringAndConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 8.0i`, env,
		`cannot convert 8i to type string`,
		`invalid operation: a & 8i (mismatched types string and complex128)`,
	)

}

// Test String & ConstBool
func TestCheckBinaryNonConstExprStringAndConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & true`, env,
		`cannot convert true to type string`,
		`invalid operation: a & true (mismatched types string and bool)`,
	)

}

// Test String & ConstString
func TestCheckBinaryNonConstExprStringAndConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & "abc"`, env,
		`invalid operation: a & "abc" (operator & not defined on string)`,
	)

}

// Test String & ConstNil
func TestCheckBinaryNonConstExprStringAndConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & nil`, env,
		`invalid operation: a & nil (mismatched types string and nil)`,
	)

}

// Test String & Int
func TestCheckBinaryNonConstExprStringAndInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types string and int)`,
	)

}

// Test String & Float32
func TestCheckBinaryNonConstExprStringAndFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types string and float32)`,
	)

}

// Test String & Complex128
func TestCheckBinaryNonConstExprStringAndComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types string and complex128)`,
	)

}

// Test String & String
func TestCheckBinaryNonConstExprStringAndString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (operator & not defined on string)`,
	)

}

// Test String & BoolT
func TestCheckBinaryNonConstExprStringAndBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types string and bool)`,
	)

}

// Test String & Slice
func TestCheckBinaryNonConstExprStringAndSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types string and eval.sliceT)`,
	)

}

// Test String & Array
func TestCheckBinaryNonConstExprStringAndArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types string and eval.arrayT)`,
	)

}

// Test String & Map
func TestCheckBinaryNonConstExprStringAndMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types string and eval.mapT)`,
	)

}

// Test String & XinterfaceX
func TestCheckBinaryNonConstExprStringAndXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types string and eval.XinterfaceX)`,
	)

}

// Test String & InterfaceX
func TestCheckBinaryNonConstExprStringAndInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types string and eval.interfaceX)`,
	)

}

// Test String & InterfaceY
func TestCheckBinaryNonConstExprStringAndInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types string and eval.interfaceY)`,
	)

}

// Test String & InterfaceZ
func TestCheckBinaryNonConstExprStringAndInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types string and eval.interfaceZ)`,
	)

}

// Test String & Ptr
func TestCheckBinaryNonConstExprStringAndPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types string and *int)`,
	)

}

// Test String & Struct
func TestCheckBinaryNonConstExprStringAndStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types string and eval.structT)`,
	)

}

// Test String & StructUncomp
func TestCheckBinaryNonConstExprStringAndStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types string and eval.structUncompT)`,
	)

}

// Test String % ConstInt
func TestCheckBinaryNonConstExprStringRemConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 4`, env,
		`cannot convert 4 to type string`,
		`invalid operation: a % 4 (mismatched types string and int)`,
	)

}

// Test String % ConstRune
func TestCheckBinaryNonConstExprStringRemConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % '@'`, env,
		`cannot convert '@' to type string`,
		`invalid operation: a % rune(64) (mismatched types string and rune)`,
	)

}

// Test String % ConstFloat
func TestCheckBinaryNonConstExprStringRemConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 2.0`, env,
		`cannot convert 2 to type string`,
		`invalid operation: a % 2 (mismatched types string and float64)`,
	)

}

// Test String % ConstComplex
func TestCheckBinaryNonConstExprStringRemConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 8.0i`, env,
		`cannot convert 8i to type string`,
		`invalid operation: a % 8i (mismatched types string and complex128)`,
	)

}

// Test String % ConstBool
func TestCheckBinaryNonConstExprStringRemConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % true`, env,
		`cannot convert true to type string`,
		`invalid operation: a % true (mismatched types string and bool)`,
	)

}

// Test String % ConstString
func TestCheckBinaryNonConstExprStringRemConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % "abc"`, env,
		`invalid operation: a % "abc" (operator % not defined on string)`,
	)

}

// Test String % ConstNil
func TestCheckBinaryNonConstExprStringRemConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % nil`, env,
		`invalid operation: a % nil (mismatched types string and nil)`,
	)

}

// Test String % Int
func TestCheckBinaryNonConstExprStringRemInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types string and int)`,
	)

}

// Test String % Float32
func TestCheckBinaryNonConstExprStringRemFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types string and float32)`,
	)

}

// Test String % Complex128
func TestCheckBinaryNonConstExprStringRemComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types string and complex128)`,
	)

}

// Test String % String
func TestCheckBinaryNonConstExprStringRemString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (operator % not defined on string)`,
	)

}

// Test String % BoolT
func TestCheckBinaryNonConstExprStringRemBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types string and bool)`,
	)

}

// Test String % Slice
func TestCheckBinaryNonConstExprStringRemSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types string and eval.sliceT)`,
	)

}

// Test String % Array
func TestCheckBinaryNonConstExprStringRemArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types string and eval.arrayT)`,
	)

}

// Test String % Map
func TestCheckBinaryNonConstExprStringRemMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types string and eval.mapT)`,
	)

}

// Test String % XinterfaceX
func TestCheckBinaryNonConstExprStringRemXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types string and eval.XinterfaceX)`,
	)

}

// Test String % InterfaceX
func TestCheckBinaryNonConstExprStringRemInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types string and eval.interfaceX)`,
	)

}

// Test String % InterfaceY
func TestCheckBinaryNonConstExprStringRemInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types string and eval.interfaceY)`,
	)

}

// Test String % InterfaceZ
func TestCheckBinaryNonConstExprStringRemInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types string and eval.interfaceZ)`,
	)

}

// Test String % Ptr
func TestCheckBinaryNonConstExprStringRemPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types string and *int)`,
	)

}

// Test String % Struct
func TestCheckBinaryNonConstExprStringRemStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types string and eval.structT)`,
	)

}

// Test String % StructUncomp
func TestCheckBinaryNonConstExprStringRemStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types string and eval.structUncompT)`,
	)

}

// Test String == ConstInt
func TestCheckBinaryNonConstExprStringEqlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 4`, env,
		`cannot convert 4 to type string`,
		`invalid operation: a == 4 (mismatched types string and int)`,
	)

}

// Test String == ConstRune
func TestCheckBinaryNonConstExprStringEqlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == '@'`, env,
		`cannot convert '@' to type string`,
		`invalid operation: a == rune(64) (mismatched types string and rune)`,
	)

}

// Test String == ConstFloat
func TestCheckBinaryNonConstExprStringEqlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 2.0`, env,
		`cannot convert 2 to type string`,
		`invalid operation: a == 2 (mismatched types string and float64)`,
	)

}

// Test String == ConstComplex
func TestCheckBinaryNonConstExprStringEqlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 8.0i`, env,
		`cannot convert 8i to type string`,
		`invalid operation: a == 8i (mismatched types string and complex128)`,
	)

}

// Test String == ConstBool
func TestCheckBinaryNonConstExprStringEqlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == true`, env,
		`cannot convert true to type string`,
		`invalid operation: a == true (mismatched types string and bool)`,
	)

}

// Test String == ConstString
func TestCheckBinaryNonConstExprStringEqlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == "abc"`, env, reflect.TypeOf(a == "abc"))
}

// Test String == ConstNil
func TestCheckBinaryNonConstExprStringEqlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == nil`, env,
		`invalid operation: a == nil (mismatched types string and nil)`,
	)

}

// Test String == Int
func TestCheckBinaryNonConstExprStringEqlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types string and int)`,
	)

}

// Test String == Float32
func TestCheckBinaryNonConstExprStringEqlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types string and float32)`,
	)

}

// Test String == Complex128
func TestCheckBinaryNonConstExprStringEqlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types string and complex128)`,
	)

}

// Test String == String
func TestCheckBinaryNonConstExprStringEqlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test String == BoolT
func TestCheckBinaryNonConstExprStringEqlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types string and bool)`,
	)

}

// Test String == Slice
func TestCheckBinaryNonConstExprStringEqlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types string and eval.sliceT)`,
	)

}

// Test String == Array
func TestCheckBinaryNonConstExprStringEqlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types string and eval.arrayT)`,
	)

}

// Test String == Map
func TestCheckBinaryNonConstExprStringEqlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types string and eval.mapT)`,
	)

}

// Test String == XinterfaceX
func TestCheckBinaryNonConstExprStringEqlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types string and eval.XinterfaceX)`,
	)

}

// Test String == InterfaceX
func TestCheckBinaryNonConstExprStringEqlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types string and eval.interfaceX)`,
	)

}

// Test String == InterfaceY
func TestCheckBinaryNonConstExprStringEqlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types string and eval.interfaceY)`,
	)

}

// Test String == InterfaceZ
func TestCheckBinaryNonConstExprStringEqlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types string and eval.interfaceZ)`,
	)

}

// Test String == Ptr
func TestCheckBinaryNonConstExprStringEqlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types string and *int)`,
	)

}

// Test String == Struct
func TestCheckBinaryNonConstExprStringEqlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types string and eval.structT)`,
	)

}

// Test String == StructUncomp
func TestCheckBinaryNonConstExprStringEqlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types string and eval.structUncompT)`,
	)

}

// Test String > ConstInt
func TestCheckBinaryNonConstExprStringGtrConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 4`, env,
		`cannot convert 4 to type string`,
		`invalid operation: a > 4 (mismatched types string and int)`,
	)

}

// Test String > ConstRune
func TestCheckBinaryNonConstExprStringGtrConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > '@'`, env,
		`cannot convert '@' to type string`,
		`invalid operation: a > rune(64) (mismatched types string and rune)`,
	)

}

// Test String > ConstFloat
func TestCheckBinaryNonConstExprStringGtrConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 2.0`, env,
		`cannot convert 2 to type string`,
		`invalid operation: a > 2 (mismatched types string and float64)`,
	)

}

// Test String > ConstComplex
func TestCheckBinaryNonConstExprStringGtrConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 8.0i`, env,
		`cannot convert 8i to type string`,
		`invalid operation: a > 8i (mismatched types string and complex128)`,
	)

}

// Test String > ConstBool
func TestCheckBinaryNonConstExprStringGtrConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > true`, env,
		`cannot convert true to type string`,
		`invalid operation: a > true (mismatched types string and bool)`,
	)

}

// Test String > ConstString
func TestCheckBinaryNonConstExprStringGtrConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a > "abc"`, env, reflect.TypeOf(a > "abc"))
}

// Test String > ConstNil
func TestCheckBinaryNonConstExprStringGtrConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > nil`, env,
		`invalid operation: a > nil (mismatched types string and nil)`,
	)

}

// Test String > Int
func TestCheckBinaryNonConstExprStringGtrInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types string and int)`,
	)

}

// Test String > Float32
func TestCheckBinaryNonConstExprStringGtrFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types string and float32)`,
	)

}

// Test String > Complex128
func TestCheckBinaryNonConstExprStringGtrComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types string and complex128)`,
	)

}

// Test String > String
func TestCheckBinaryNonConstExprStringGtrString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a > b`, env, reflect.TypeOf(a > b))
}

// Test String > BoolT
func TestCheckBinaryNonConstExprStringGtrBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types string and bool)`,
	)

}

// Test String > Slice
func TestCheckBinaryNonConstExprStringGtrSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types string and eval.sliceT)`,
	)

}

// Test String > Array
func TestCheckBinaryNonConstExprStringGtrArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types string and eval.arrayT)`,
	)

}

// Test String > Map
func TestCheckBinaryNonConstExprStringGtrMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types string and eval.mapT)`,
	)

}

// Test String > XinterfaceX
func TestCheckBinaryNonConstExprStringGtrXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types string and eval.XinterfaceX)`,
	)

}

// Test String > InterfaceX
func TestCheckBinaryNonConstExprStringGtrInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types string and eval.interfaceX)`,
	)

}

// Test String > InterfaceY
func TestCheckBinaryNonConstExprStringGtrInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types string and eval.interfaceY)`,
	)

}

// Test String > InterfaceZ
func TestCheckBinaryNonConstExprStringGtrInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types string and eval.interfaceZ)`,
	)

}

// Test String > Ptr
func TestCheckBinaryNonConstExprStringGtrPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types string and *int)`,
	)

}

// Test String > Struct
func TestCheckBinaryNonConstExprStringGtrStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types string and eval.structT)`,
	)

}

// Test String > StructUncomp
func TestCheckBinaryNonConstExprStringGtrStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types string and eval.structUncompT)`,
	)

}

// Test String << ConstInt
func TestCheckBinaryNonConstExprStringShlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 4`, env,
		`invalid operation: a << 4 (shift of type string)`,
	)

}

// Test String << ConstRune
func TestCheckBinaryNonConstExprStringShlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << '@'`, env,
		`invalid operation: a << 64 (shift of type string)`,
	)

}

// Test String << ConstFloat
func TestCheckBinaryNonConstExprStringShlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 2.0`, env,
		`invalid operation: a << 2 (shift of type string)`,
	)

}

// Test String << ConstComplex
func TestCheckBinaryNonConstExprStringShlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: a << 0 (shift of type string)`,
	)

}

// Test String << ConstBool
func TestCheckBinaryNonConstExprStringShlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << true`, env,
		`invalid operation: a << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test String << ConstString
func TestCheckBinaryNonConstExprStringShlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: a << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test String << ConstNil
func TestCheckBinaryNonConstExprStringShlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test String << Int
func TestCheckBinaryNonConstExprStringShlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type int, must be unsigned integer)`,
	)

}

// Test String << Float32
func TestCheckBinaryNonConstExprStringShlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type float32, must be unsigned integer)`,
	)

}

// Test String << Complex128
func TestCheckBinaryNonConstExprStringShlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type complex128, must be unsigned integer)`,
	)

}

// Test String << String
func TestCheckBinaryNonConstExprStringShlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type string, must be unsigned integer)`,
	)

}

// Test String << BoolT
func TestCheckBinaryNonConstExprStringShlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type bool, must be unsigned integer)`,
	)

}

// Test String << Slice
func TestCheckBinaryNonConstExprStringShlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.sliceT, must be unsigned integer)`,
	)

}

// Test String << Array
func TestCheckBinaryNonConstExprStringShlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.arrayT, must be unsigned integer)`,
	)

}

// Test String << Map
func TestCheckBinaryNonConstExprStringShlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.mapT, must be unsigned integer)`,
	)

}

// Test String << XinterfaceX
func TestCheckBinaryNonConstExprStringShlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.XinterfaceX, must be unsigned integer)`,
	)

}

// Test String << InterfaceX
func TestCheckBinaryNonConstExprStringShlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceX, must be unsigned integer)`,
	)

}

// Test String << InterfaceY
func TestCheckBinaryNonConstExprStringShlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceY, must be unsigned integer)`,
	)

}

// Test String << InterfaceZ
func TestCheckBinaryNonConstExprStringShlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceZ, must be unsigned integer)`,
	)

}

// Test String << Ptr
func TestCheckBinaryNonConstExprStringShlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type *int, must be unsigned integer)`,
	)

}

// Test String << Struct
func TestCheckBinaryNonConstExprStringShlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structT, must be unsigned integer)`,
	)

}

// Test String << StructUncomp
func TestCheckBinaryNonConstExprStringShlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := string("abc"); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structUncompT, must be unsigned integer)`,
	)

}

// Test BoolT + ConstInt
func TestCheckBinaryNonConstExprBoolTAddConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: a + 4 (mismatched types bool and int)`,
	)

}

// Test BoolT + ConstRune
func TestCheckBinaryNonConstExprBoolTAddConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: a + rune(64) (mismatched types bool and rune)`,
	)

}

// Test BoolT + ConstFloat
func TestCheckBinaryNonConstExprBoolTAddConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: a + 2 (mismatched types bool and float64)`,
	)

}

// Test BoolT + ConstComplex
func TestCheckBinaryNonConstExprBoolTAddConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: a + 8i (mismatched types bool and complex128)`,
	)

}

// Test BoolT + ConstBool
func TestCheckBinaryNonConstExprBoolTAddConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + true`, env,
		`invalid operation: a + true (operator + not defined on bool)`,
	)

}

// Test BoolT + ConstString
func TestCheckBinaryNonConstExprBoolTAddConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: a + "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT + ConstNil
func TestCheckBinaryNonConstExprBoolTAddConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + nil`, env,
		`cannot convert nil to type bool`,
	)

}

// Test BoolT + Int
func TestCheckBinaryNonConstExprBoolTAddInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types bool and int)`,
	)

}

// Test BoolT + Float32
func TestCheckBinaryNonConstExprBoolTAddFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types bool and float32)`,
	)

}

// Test BoolT + Complex128
func TestCheckBinaryNonConstExprBoolTAddComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types bool and complex128)`,
	)

}

// Test BoolT + String
func TestCheckBinaryNonConstExprBoolTAddString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types bool and string)`,
	)

}

// Test BoolT + BoolT
func TestCheckBinaryNonConstExprBoolTAddBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (operator + not defined on bool)`,
	)

}

// Test BoolT + Slice
func TestCheckBinaryNonConstExprBoolTAddSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types bool and eval.sliceT)`,
	)

}

// Test BoolT + Array
func TestCheckBinaryNonConstExprBoolTAddArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types bool and eval.arrayT)`,
	)

}

// Test BoolT + Map
func TestCheckBinaryNonConstExprBoolTAddMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types bool and eval.mapT)`,
	)

}

// Test BoolT + XinterfaceX
func TestCheckBinaryNonConstExprBoolTAddXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types bool and eval.XinterfaceX)`,
	)

}

// Test BoolT + InterfaceX
func TestCheckBinaryNonConstExprBoolTAddInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types bool and eval.interfaceX)`,
	)

}

// Test BoolT + InterfaceY
func TestCheckBinaryNonConstExprBoolTAddInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types bool and eval.interfaceY)`,
	)

}

// Test BoolT + InterfaceZ
func TestCheckBinaryNonConstExprBoolTAddInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types bool and eval.interfaceZ)`,
	)

}

// Test BoolT + Ptr
func TestCheckBinaryNonConstExprBoolTAddPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types bool and *int)`,
	)

}

// Test BoolT + Struct
func TestCheckBinaryNonConstExprBoolTAddStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types bool and eval.structT)`,
	)

}

// Test BoolT + StructUncomp
func TestCheckBinaryNonConstExprBoolTAddStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types bool and eval.structUncompT)`,
	)

}

// Test BoolT & ConstInt
func TestCheckBinaryNonConstExprBoolTAndConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: a & 4 (mismatched types bool and int)`,
	)

}

// Test BoolT & ConstRune
func TestCheckBinaryNonConstExprBoolTAndConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: a & rune(64) (mismatched types bool and rune)`,
	)

}

// Test BoolT & ConstFloat
func TestCheckBinaryNonConstExprBoolTAndConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: a & 2 (mismatched types bool and float64)`,
	)

}

// Test BoolT & ConstComplex
func TestCheckBinaryNonConstExprBoolTAndConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: a & 8i (mismatched types bool and complex128)`,
	)

}

// Test BoolT & ConstBool
func TestCheckBinaryNonConstExprBoolTAndConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & true`, env,
		`invalid operation: a & true (operator & not defined on bool)`,
	)

}

// Test BoolT & ConstString
func TestCheckBinaryNonConstExprBoolTAndConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: a & "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT & ConstNil
func TestCheckBinaryNonConstExprBoolTAndConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & nil`, env,
		`cannot convert nil to type bool`,
	)

}

// Test BoolT & Int
func TestCheckBinaryNonConstExprBoolTAndInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types bool and int)`,
	)

}

// Test BoolT & Float32
func TestCheckBinaryNonConstExprBoolTAndFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types bool and float32)`,
	)

}

// Test BoolT & Complex128
func TestCheckBinaryNonConstExprBoolTAndComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types bool and complex128)`,
	)

}

// Test BoolT & String
func TestCheckBinaryNonConstExprBoolTAndString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types bool and string)`,
	)

}

// Test BoolT & BoolT
func TestCheckBinaryNonConstExprBoolTAndBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (operator & not defined on bool)`,
	)

}

// Test BoolT & Slice
func TestCheckBinaryNonConstExprBoolTAndSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types bool and eval.sliceT)`,
	)

}

// Test BoolT & Array
func TestCheckBinaryNonConstExprBoolTAndArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types bool and eval.arrayT)`,
	)

}

// Test BoolT & Map
func TestCheckBinaryNonConstExprBoolTAndMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types bool and eval.mapT)`,
	)

}

// Test BoolT & XinterfaceX
func TestCheckBinaryNonConstExprBoolTAndXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types bool and eval.XinterfaceX)`,
	)

}

// Test BoolT & InterfaceX
func TestCheckBinaryNonConstExprBoolTAndInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types bool and eval.interfaceX)`,
	)

}

// Test BoolT & InterfaceY
func TestCheckBinaryNonConstExprBoolTAndInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types bool and eval.interfaceY)`,
	)

}

// Test BoolT & InterfaceZ
func TestCheckBinaryNonConstExprBoolTAndInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types bool and eval.interfaceZ)`,
	)

}

// Test BoolT & Ptr
func TestCheckBinaryNonConstExprBoolTAndPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types bool and *int)`,
	)

}

// Test BoolT & Struct
func TestCheckBinaryNonConstExprBoolTAndStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types bool and eval.structT)`,
	)

}

// Test BoolT & StructUncomp
func TestCheckBinaryNonConstExprBoolTAndStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types bool and eval.structUncompT)`,
	)

}

// Test BoolT % ConstInt
func TestCheckBinaryNonConstExprBoolTRemConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: a % 4 (mismatched types bool and int)`,
	)

}

// Test BoolT % ConstRune
func TestCheckBinaryNonConstExprBoolTRemConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: a % rune(64) (mismatched types bool and rune)`,
	)

}

// Test BoolT % ConstFloat
func TestCheckBinaryNonConstExprBoolTRemConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: a % 2 (mismatched types bool and float64)`,
	)

}

// Test BoolT % ConstComplex
func TestCheckBinaryNonConstExprBoolTRemConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: a % 8i (mismatched types bool and complex128)`,
	)

}

// Test BoolT % ConstBool
func TestCheckBinaryNonConstExprBoolTRemConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % true`, env,
		`invalid operation: a % true (operator % not defined on bool)`,
	)

}

// Test BoolT % ConstString
func TestCheckBinaryNonConstExprBoolTRemConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: a % "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT % ConstNil
func TestCheckBinaryNonConstExprBoolTRemConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % nil`, env,
		`cannot convert nil to type bool`,
	)

}

// Test BoolT % Int
func TestCheckBinaryNonConstExprBoolTRemInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types bool and int)`,
	)

}

// Test BoolT % Float32
func TestCheckBinaryNonConstExprBoolTRemFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types bool and float32)`,
	)

}

// Test BoolT % Complex128
func TestCheckBinaryNonConstExprBoolTRemComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types bool and complex128)`,
	)

}

// Test BoolT % String
func TestCheckBinaryNonConstExprBoolTRemString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types bool and string)`,
	)

}

// Test BoolT % BoolT
func TestCheckBinaryNonConstExprBoolTRemBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (operator % not defined on bool)`,
	)

}

// Test BoolT % Slice
func TestCheckBinaryNonConstExprBoolTRemSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types bool and eval.sliceT)`,
	)

}

// Test BoolT % Array
func TestCheckBinaryNonConstExprBoolTRemArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types bool and eval.arrayT)`,
	)

}

// Test BoolT % Map
func TestCheckBinaryNonConstExprBoolTRemMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types bool and eval.mapT)`,
	)

}

// Test BoolT % XinterfaceX
func TestCheckBinaryNonConstExprBoolTRemXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types bool and eval.XinterfaceX)`,
	)

}

// Test BoolT % InterfaceX
func TestCheckBinaryNonConstExprBoolTRemInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types bool and eval.interfaceX)`,
	)

}

// Test BoolT % InterfaceY
func TestCheckBinaryNonConstExprBoolTRemInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types bool and eval.interfaceY)`,
	)

}

// Test BoolT % InterfaceZ
func TestCheckBinaryNonConstExprBoolTRemInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types bool and eval.interfaceZ)`,
	)

}

// Test BoolT % Ptr
func TestCheckBinaryNonConstExprBoolTRemPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types bool and *int)`,
	)

}

// Test BoolT % Struct
func TestCheckBinaryNonConstExprBoolTRemStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types bool and eval.structT)`,
	)

}

// Test BoolT % StructUncomp
func TestCheckBinaryNonConstExprBoolTRemStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types bool and eval.structUncompT)`,
	)

}

// Test BoolT == ConstInt
func TestCheckBinaryNonConstExprBoolTEqlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: a == 4 (mismatched types bool and int)`,
	)

}

// Test BoolT == ConstRune
func TestCheckBinaryNonConstExprBoolTEqlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: a == rune(64) (mismatched types bool and rune)`,
	)

}

// Test BoolT == ConstFloat
func TestCheckBinaryNonConstExprBoolTEqlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: a == 2 (mismatched types bool and float64)`,
	)

}

// Test BoolT == ConstComplex
func TestCheckBinaryNonConstExprBoolTEqlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: a == 8i (mismatched types bool and complex128)`,
	)

}

// Test BoolT == ConstBool
func TestCheckBinaryNonConstExprBoolTEqlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == true`, env, reflect.TypeOf(a == true))
}

// Test BoolT == ConstString
func TestCheckBinaryNonConstExprBoolTEqlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: a == "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT == ConstNil
func TestCheckBinaryNonConstExprBoolTEqlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == nil`, env,
		`cannot convert nil to type bool`,
	)

}

// Test BoolT == Int
func TestCheckBinaryNonConstExprBoolTEqlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types bool and int)`,
	)

}

// Test BoolT == Float32
func TestCheckBinaryNonConstExprBoolTEqlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types bool and float32)`,
	)

}

// Test BoolT == Complex128
func TestCheckBinaryNonConstExprBoolTEqlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types bool and complex128)`,
	)

}

// Test BoolT == String
func TestCheckBinaryNonConstExprBoolTEqlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types bool and string)`,
	)

}

// Test BoolT == BoolT
func TestCheckBinaryNonConstExprBoolTEqlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test BoolT == Slice
func TestCheckBinaryNonConstExprBoolTEqlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types bool and eval.sliceT)`,
	)

}

// Test BoolT == Array
func TestCheckBinaryNonConstExprBoolTEqlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types bool and eval.arrayT)`,
	)

}

// Test BoolT == Map
func TestCheckBinaryNonConstExprBoolTEqlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types bool and eval.mapT)`,
	)

}

// Test BoolT == XinterfaceX
func TestCheckBinaryNonConstExprBoolTEqlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types bool and eval.XinterfaceX)`,
	)

}

// Test BoolT == InterfaceX
func TestCheckBinaryNonConstExprBoolTEqlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types bool and eval.interfaceX)`,
	)

}

// Test BoolT == InterfaceY
func TestCheckBinaryNonConstExprBoolTEqlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types bool and eval.interfaceY)`,
	)

}

// Test BoolT == InterfaceZ
func TestCheckBinaryNonConstExprBoolTEqlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types bool and eval.interfaceZ)`,
	)

}

// Test BoolT == Ptr
func TestCheckBinaryNonConstExprBoolTEqlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types bool and *int)`,
	)

}

// Test BoolT == Struct
func TestCheckBinaryNonConstExprBoolTEqlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types bool and eval.structT)`,
	)

}

// Test BoolT == StructUncomp
func TestCheckBinaryNonConstExprBoolTEqlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types bool and eval.structUncompT)`,
	)

}

// Test BoolT > ConstInt
func TestCheckBinaryNonConstExprBoolTGtrConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 4`, env,
		`cannot convert 4 to type bool`,
		`invalid operation: a > 4 (mismatched types bool and int)`,
	)

}

// Test BoolT > ConstRune
func TestCheckBinaryNonConstExprBoolTGtrConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > '@'`, env,
		`cannot convert '@' to type bool`,
		`invalid operation: a > rune(64) (mismatched types bool and rune)`,
	)

}

// Test BoolT > ConstFloat
func TestCheckBinaryNonConstExprBoolTGtrConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 2.0`, env,
		`cannot convert 2 to type bool`,
		`invalid operation: a > 2 (mismatched types bool and float64)`,
	)

}

// Test BoolT > ConstComplex
func TestCheckBinaryNonConstExprBoolTGtrConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 8.0i`, env,
		`cannot convert 8i to type bool`,
		`invalid operation: a > 8i (mismatched types bool and complex128)`,
	)

}

// Test BoolT > ConstBool
func TestCheckBinaryNonConstExprBoolTGtrConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > true`, env,
		`invalid operation: a > true (operator > not defined on bool)`,
	)

}

// Test BoolT > ConstString
func TestCheckBinaryNonConstExprBoolTGtrConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > "abc"`, env,
		`cannot convert "abc" to type bool`,
		`invalid operation: a > "abc" (mismatched types bool and string)`,
	)

}

// Test BoolT > ConstNil
func TestCheckBinaryNonConstExprBoolTGtrConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > nil`, env,
		`cannot convert nil to type bool`,
	)

}

// Test BoolT > Int
func TestCheckBinaryNonConstExprBoolTGtrInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types bool and int)`,
	)

}

// Test BoolT > Float32
func TestCheckBinaryNonConstExprBoolTGtrFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types bool and float32)`,
	)

}

// Test BoolT > Complex128
func TestCheckBinaryNonConstExprBoolTGtrComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types bool and complex128)`,
	)

}

// Test BoolT > String
func TestCheckBinaryNonConstExprBoolTGtrString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types bool and string)`,
	)

}

// Test BoolT > BoolT
func TestCheckBinaryNonConstExprBoolTGtrBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (operator > not defined on bool)`,
	)

}

// Test BoolT > Slice
func TestCheckBinaryNonConstExprBoolTGtrSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types bool and eval.sliceT)`,
	)

}

// Test BoolT > Array
func TestCheckBinaryNonConstExprBoolTGtrArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types bool and eval.arrayT)`,
	)

}

// Test BoolT > Map
func TestCheckBinaryNonConstExprBoolTGtrMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types bool and eval.mapT)`,
	)

}

// Test BoolT > XinterfaceX
func TestCheckBinaryNonConstExprBoolTGtrXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types bool and eval.XinterfaceX)`,
	)

}

// Test BoolT > InterfaceX
func TestCheckBinaryNonConstExprBoolTGtrInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types bool and eval.interfaceX)`,
	)

}

// Test BoolT > InterfaceY
func TestCheckBinaryNonConstExprBoolTGtrInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types bool and eval.interfaceY)`,
	)

}

// Test BoolT > InterfaceZ
func TestCheckBinaryNonConstExprBoolTGtrInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types bool and eval.interfaceZ)`,
	)

}

// Test BoolT > Ptr
func TestCheckBinaryNonConstExprBoolTGtrPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types bool and *int)`,
	)

}

// Test BoolT > Struct
func TestCheckBinaryNonConstExprBoolTGtrStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types bool and eval.structT)`,
	)

}

// Test BoolT > StructUncomp
func TestCheckBinaryNonConstExprBoolTGtrStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types bool and eval.structUncompT)`,
	)

}

// Test BoolT << ConstInt
func TestCheckBinaryNonConstExprBoolTShlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 4`, env,
		`invalid operation: a << 4 (shift of type bool)`,
	)

}

// Test BoolT << ConstRune
func TestCheckBinaryNonConstExprBoolTShlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << '@'`, env,
		`invalid operation: a << 64 (shift of type bool)`,
	)

}

// Test BoolT << ConstFloat
func TestCheckBinaryNonConstExprBoolTShlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 2.0`, env,
		`invalid operation: a << 2 (shift of type bool)`,
	)

}

// Test BoolT << ConstComplex
func TestCheckBinaryNonConstExprBoolTShlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: a << 0 (shift of type bool)`,
	)

}

// Test BoolT << ConstBool
func TestCheckBinaryNonConstExprBoolTShlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << true`, env,
		`invalid operation: a << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test BoolT << ConstString
func TestCheckBinaryNonConstExprBoolTShlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: a << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test BoolT << ConstNil
func TestCheckBinaryNonConstExprBoolTShlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test BoolT << Int
func TestCheckBinaryNonConstExprBoolTShlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type int, must be unsigned integer)`,
	)

}

// Test BoolT << Float32
func TestCheckBinaryNonConstExprBoolTShlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type float32, must be unsigned integer)`,
	)

}

// Test BoolT << Complex128
func TestCheckBinaryNonConstExprBoolTShlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type complex128, must be unsigned integer)`,
	)

}

// Test BoolT << String
func TestCheckBinaryNonConstExprBoolTShlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type string, must be unsigned integer)`,
	)

}

// Test BoolT << BoolT
func TestCheckBinaryNonConstExprBoolTShlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type bool, must be unsigned integer)`,
	)

}

// Test BoolT << Slice
func TestCheckBinaryNonConstExprBoolTShlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.sliceT, must be unsigned integer)`,
	)

}

// Test BoolT << Array
func TestCheckBinaryNonConstExprBoolTShlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.arrayT, must be unsigned integer)`,
	)

}

// Test BoolT << Map
func TestCheckBinaryNonConstExprBoolTShlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.mapT, must be unsigned integer)`,
	)

}

// Test BoolT << XinterfaceX
func TestCheckBinaryNonConstExprBoolTShlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.XinterfaceX, must be unsigned integer)`,
	)

}

// Test BoolT << InterfaceX
func TestCheckBinaryNonConstExprBoolTShlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceX, must be unsigned integer)`,
	)

}

// Test BoolT << InterfaceY
func TestCheckBinaryNonConstExprBoolTShlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceY, must be unsigned integer)`,
	)

}

// Test BoolT << InterfaceZ
func TestCheckBinaryNonConstExprBoolTShlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceZ, must be unsigned integer)`,
	)

}

// Test BoolT << Ptr
func TestCheckBinaryNonConstExprBoolTShlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type *int, must be unsigned integer)`,
	)

}

// Test BoolT << Struct
func TestCheckBinaryNonConstExprBoolTShlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structT, must be unsigned integer)`,
	)

}

// Test BoolT << StructUncomp
func TestCheckBinaryNonConstExprBoolTShlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := bool(true); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structUncompT, must be unsigned integer)`,
	)

}

// Test Slice + ConstInt
func TestCheckBinaryNonConstExprSliceAddConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 4`, env,
		`invalid operation: a + 4 (mismatched types eval.sliceT and int)`,
	)

}

// Test Slice + ConstRune
func TestCheckBinaryNonConstExprSliceAddConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + '@'`, env,
		`invalid operation: a + rune(64) (mismatched types eval.sliceT and rune)`,
	)

}

// Test Slice + ConstFloat
func TestCheckBinaryNonConstExprSliceAddConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 2.0`, env,
		`invalid operation: a + 2 (mismatched types eval.sliceT and float64)`,
	)

}

// Test Slice + ConstComplex
func TestCheckBinaryNonConstExprSliceAddConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 8.0i`, env,
		`invalid operation: a + 8i (mismatched types eval.sliceT and complex128)`,
	)

}

// Test Slice + ConstBool
func TestCheckBinaryNonConstExprSliceAddConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + true`, env,
		`invalid operation: a + true (mismatched types eval.sliceT and bool)`,
	)

}

// Test Slice + ConstString
func TestCheckBinaryNonConstExprSliceAddConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + "abc"`, env,
		`invalid operation: a + "abc" (mismatched types eval.sliceT and string)`,
	)

}

// Test Slice + ConstNil
func TestCheckBinaryNonConstExprSliceAddConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + nil`, env,
		`invalid operation: a + nil (operator + not defined on slice)`,
	)

}

// Test Slice + Int
func TestCheckBinaryNonConstExprSliceAddInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.sliceT and int)`,
	)

}

// Test Slice + Float32
func TestCheckBinaryNonConstExprSliceAddFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.sliceT and float32)`,
	)

}

// Test Slice + Complex128
func TestCheckBinaryNonConstExprSliceAddComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.sliceT and complex128)`,
	)

}

// Test Slice + String
func TestCheckBinaryNonConstExprSliceAddString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.sliceT and string)`,
	)

}

// Test Slice + BoolT
func TestCheckBinaryNonConstExprSliceAddBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.sliceT and bool)`,
	)

}

// Test Slice + Slice
func TestCheckBinaryNonConstExprSliceAddSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (operator + not defined on slice)`,
	)

}

// Test Slice + Array
func TestCheckBinaryNonConstExprSliceAddArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.sliceT and eval.arrayT)`,
	)

}

// Test Slice + Map
func TestCheckBinaryNonConstExprSliceAddMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.sliceT and eval.mapT)`,
	)

}

// Test Slice + XinterfaceX
func TestCheckBinaryNonConstExprSliceAddXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.sliceT and eval.XinterfaceX)`,
	)

}

// Test Slice + InterfaceX
func TestCheckBinaryNonConstExprSliceAddInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.sliceT and eval.interfaceX)`,
	)

}

// Test Slice + InterfaceY
func TestCheckBinaryNonConstExprSliceAddInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.sliceT and eval.interfaceY)`,
	)

}

// Test Slice + InterfaceZ
func TestCheckBinaryNonConstExprSliceAddInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.sliceT and eval.interfaceZ)`,
	)

}

// Test Slice + Ptr
func TestCheckBinaryNonConstExprSliceAddPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.sliceT and *int)`,
	)

}

// Test Slice + Struct
func TestCheckBinaryNonConstExprSliceAddStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.sliceT and eval.structT)`,
	)

}

// Test Slice + StructUncomp
func TestCheckBinaryNonConstExprSliceAddStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.sliceT and eval.structUncompT)`,
	)

}

// Test Slice & ConstInt
func TestCheckBinaryNonConstExprSliceAndConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 4`, env,
		`invalid operation: a & 4 (mismatched types eval.sliceT and int)`,
	)

}

// Test Slice & ConstRune
func TestCheckBinaryNonConstExprSliceAndConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & '@'`, env,
		`invalid operation: a & rune(64) (mismatched types eval.sliceT and rune)`,
	)

}

// Test Slice & ConstFloat
func TestCheckBinaryNonConstExprSliceAndConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 2.0`, env,
		`invalid operation: a & 2 (mismatched types eval.sliceT and float64)`,
	)

}

// Test Slice & ConstComplex
func TestCheckBinaryNonConstExprSliceAndConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 8.0i`, env,
		`invalid operation: a & 8i (mismatched types eval.sliceT and complex128)`,
	)

}

// Test Slice & ConstBool
func TestCheckBinaryNonConstExprSliceAndConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & true`, env,
		`invalid operation: a & true (mismatched types eval.sliceT and bool)`,
	)

}

// Test Slice & ConstString
func TestCheckBinaryNonConstExprSliceAndConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & "abc"`, env,
		`invalid operation: a & "abc" (mismatched types eval.sliceT and string)`,
	)

}

// Test Slice & ConstNil
func TestCheckBinaryNonConstExprSliceAndConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & nil`, env,
		`invalid operation: a & nil (operator & not defined on slice)`,
	)

}

// Test Slice & Int
func TestCheckBinaryNonConstExprSliceAndInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.sliceT and int)`,
	)

}

// Test Slice & Float32
func TestCheckBinaryNonConstExprSliceAndFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.sliceT and float32)`,
	)

}

// Test Slice & Complex128
func TestCheckBinaryNonConstExprSliceAndComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.sliceT and complex128)`,
	)

}

// Test Slice & String
func TestCheckBinaryNonConstExprSliceAndString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.sliceT and string)`,
	)

}

// Test Slice & BoolT
func TestCheckBinaryNonConstExprSliceAndBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.sliceT and bool)`,
	)

}

// Test Slice & Slice
func TestCheckBinaryNonConstExprSliceAndSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (operator & not defined on slice)`,
	)

}

// Test Slice & Array
func TestCheckBinaryNonConstExprSliceAndArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.sliceT and eval.arrayT)`,
	)

}

// Test Slice & Map
func TestCheckBinaryNonConstExprSliceAndMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.sliceT and eval.mapT)`,
	)

}

// Test Slice & XinterfaceX
func TestCheckBinaryNonConstExprSliceAndXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.sliceT and eval.XinterfaceX)`,
	)

}

// Test Slice & InterfaceX
func TestCheckBinaryNonConstExprSliceAndInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.sliceT and eval.interfaceX)`,
	)

}

// Test Slice & InterfaceY
func TestCheckBinaryNonConstExprSliceAndInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.sliceT and eval.interfaceY)`,
	)

}

// Test Slice & InterfaceZ
func TestCheckBinaryNonConstExprSliceAndInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.sliceT and eval.interfaceZ)`,
	)

}

// Test Slice & Ptr
func TestCheckBinaryNonConstExprSliceAndPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.sliceT and *int)`,
	)

}

// Test Slice & Struct
func TestCheckBinaryNonConstExprSliceAndStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.sliceT and eval.structT)`,
	)

}

// Test Slice & StructUncomp
func TestCheckBinaryNonConstExprSliceAndStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.sliceT and eval.structUncompT)`,
	)

}

// Test Slice % ConstInt
func TestCheckBinaryNonConstExprSliceRemConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 4`, env,
		`invalid operation: a % 4 (mismatched types eval.sliceT and int)`,
	)

}

// Test Slice % ConstRune
func TestCheckBinaryNonConstExprSliceRemConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % '@'`, env,
		`invalid operation: a % rune(64) (mismatched types eval.sliceT and rune)`,
	)

}

// Test Slice % ConstFloat
func TestCheckBinaryNonConstExprSliceRemConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 2.0`, env,
		`invalid operation: a % 2 (mismatched types eval.sliceT and float64)`,
	)

}

// Test Slice % ConstComplex
func TestCheckBinaryNonConstExprSliceRemConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 8.0i`, env,
		`invalid operation: a % 8i (mismatched types eval.sliceT and complex128)`,
	)

}

// Test Slice % ConstBool
func TestCheckBinaryNonConstExprSliceRemConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % true`, env,
		`invalid operation: a % true (mismatched types eval.sliceT and bool)`,
	)

}

// Test Slice % ConstString
func TestCheckBinaryNonConstExprSliceRemConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % "abc"`, env,
		`invalid operation: a % "abc" (mismatched types eval.sliceT and string)`,
	)

}

// Test Slice % ConstNil
func TestCheckBinaryNonConstExprSliceRemConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % nil`, env,
		`invalid operation: a % nil (operator % not defined on slice)`,
	)

}

// Test Slice % Int
func TestCheckBinaryNonConstExprSliceRemInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.sliceT and int)`,
	)

}

// Test Slice % Float32
func TestCheckBinaryNonConstExprSliceRemFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.sliceT and float32)`,
	)

}

// Test Slice % Complex128
func TestCheckBinaryNonConstExprSliceRemComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.sliceT and complex128)`,
	)

}

// Test Slice % String
func TestCheckBinaryNonConstExprSliceRemString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.sliceT and string)`,
	)

}

// Test Slice % BoolT
func TestCheckBinaryNonConstExprSliceRemBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.sliceT and bool)`,
	)

}

// Test Slice % Slice
func TestCheckBinaryNonConstExprSliceRemSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (operator % not defined on slice)`,
	)

}

// Test Slice % Array
func TestCheckBinaryNonConstExprSliceRemArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.sliceT and eval.arrayT)`,
	)

}

// Test Slice % Map
func TestCheckBinaryNonConstExprSliceRemMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.sliceT and eval.mapT)`,
	)

}

// Test Slice % XinterfaceX
func TestCheckBinaryNonConstExprSliceRemXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.sliceT and eval.XinterfaceX)`,
	)

}

// Test Slice % InterfaceX
func TestCheckBinaryNonConstExprSliceRemInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.sliceT and eval.interfaceX)`,
	)

}

// Test Slice % InterfaceY
func TestCheckBinaryNonConstExprSliceRemInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.sliceT and eval.interfaceY)`,
	)

}

// Test Slice % InterfaceZ
func TestCheckBinaryNonConstExprSliceRemInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.sliceT and eval.interfaceZ)`,
	)

}

// Test Slice % Ptr
func TestCheckBinaryNonConstExprSliceRemPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.sliceT and *int)`,
	)

}

// Test Slice % Struct
func TestCheckBinaryNonConstExprSliceRemStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.sliceT and eval.structT)`,
	)

}

// Test Slice % StructUncomp
func TestCheckBinaryNonConstExprSliceRemStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.sliceT and eval.structUncompT)`,
	)

}

// Test Slice == ConstInt
func TestCheckBinaryNonConstExprSliceEqlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 4`, env,
		`invalid operation: a == 4 (mismatched types eval.sliceT and int)`,
	)

}

// Test Slice == ConstRune
func TestCheckBinaryNonConstExprSliceEqlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == '@'`, env,
		`invalid operation: a == rune(64) (mismatched types eval.sliceT and rune)`,
	)

}

// Test Slice == ConstFloat
func TestCheckBinaryNonConstExprSliceEqlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 2.0`, env,
		`invalid operation: a == 2 (mismatched types eval.sliceT and float64)`,
	)

}

// Test Slice == ConstComplex
func TestCheckBinaryNonConstExprSliceEqlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 8.0i`, env,
		`invalid operation: a == 8i (mismatched types eval.sliceT and complex128)`,
	)

}

// Test Slice == ConstBool
func TestCheckBinaryNonConstExprSliceEqlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == true`, env,
		`invalid operation: a == true (mismatched types eval.sliceT and bool)`,
	)

}

// Test Slice == ConstString
func TestCheckBinaryNonConstExprSliceEqlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == "abc"`, env,
		`invalid operation: a == "abc" (mismatched types eval.sliceT and string)`,
	)

}

// Test Slice == ConstNil
func TestCheckBinaryNonConstExprSliceEqlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == nil`, env, reflect.TypeOf(a == nil))
}

// Test Slice == Int
func TestCheckBinaryNonConstExprSliceEqlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.sliceT and int)`,
	)

}

// Test Slice == Float32
func TestCheckBinaryNonConstExprSliceEqlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.sliceT and float32)`,
	)

}

// Test Slice == Complex128
func TestCheckBinaryNonConstExprSliceEqlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.sliceT and complex128)`,
	)

}

// Test Slice == String
func TestCheckBinaryNonConstExprSliceEqlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.sliceT and string)`,
	)

}

// Test Slice == BoolT
func TestCheckBinaryNonConstExprSliceEqlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.sliceT and bool)`,
	)

}

// Test Slice == Slice
func TestCheckBinaryNonConstExprSliceEqlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (slice can only be compared to nil)`,
	)

}

// Test Slice == Array
func TestCheckBinaryNonConstExprSliceEqlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.sliceT and eval.arrayT)`,
	)

}

// Test Slice == Map
func TestCheckBinaryNonConstExprSliceEqlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.sliceT and eval.mapT)`,
	)

}

// Test Slice == XinterfaceX
func TestCheckBinaryNonConstExprSliceEqlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.sliceT and eval.XinterfaceX)`,
	)

}

// Test Slice == InterfaceX
func TestCheckBinaryNonConstExprSliceEqlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.sliceT and eval.interfaceX)`,
	)

}

// Test Slice == InterfaceY
func TestCheckBinaryNonConstExprSliceEqlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.sliceT and eval.interfaceY)`,
	)

}

// Test Slice == InterfaceZ
func TestCheckBinaryNonConstExprSliceEqlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.sliceT and eval.interfaceZ)`,
	)

}

// Test Slice == Ptr
func TestCheckBinaryNonConstExprSliceEqlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.sliceT and *int)`,
	)

}

// Test Slice == Struct
func TestCheckBinaryNonConstExprSliceEqlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.sliceT and eval.structT)`,
	)

}

// Test Slice == StructUncomp
func TestCheckBinaryNonConstExprSliceEqlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.sliceT and eval.structUncompT)`,
	)

}

// Test Slice > ConstInt
func TestCheckBinaryNonConstExprSliceGtrConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 4`, env,
		`invalid operation: a > 4 (mismatched types eval.sliceT and int)`,
	)

}

// Test Slice > ConstRune
func TestCheckBinaryNonConstExprSliceGtrConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > '@'`, env,
		`invalid operation: a > rune(64) (mismatched types eval.sliceT and rune)`,
	)

}

// Test Slice > ConstFloat
func TestCheckBinaryNonConstExprSliceGtrConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 2.0`, env,
		`invalid operation: a > 2 (mismatched types eval.sliceT and float64)`,
	)

}

// Test Slice > ConstComplex
func TestCheckBinaryNonConstExprSliceGtrConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 8.0i`, env,
		`invalid operation: a > 8i (mismatched types eval.sliceT and complex128)`,
	)

}

// Test Slice > ConstBool
func TestCheckBinaryNonConstExprSliceGtrConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > true`, env,
		`invalid operation: a > true (mismatched types eval.sliceT and bool)`,
	)

}

// Test Slice > ConstString
func TestCheckBinaryNonConstExprSliceGtrConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > "abc"`, env,
		`invalid operation: a > "abc" (mismatched types eval.sliceT and string)`,
	)

}

// Test Slice > ConstNil
func TestCheckBinaryNonConstExprSliceGtrConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > nil`, env,
		`invalid operation: a > nil (operator > not defined on slice)`,
	)

}

// Test Slice > Int
func TestCheckBinaryNonConstExprSliceGtrInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.sliceT and int)`,
	)

}

// Test Slice > Float32
func TestCheckBinaryNonConstExprSliceGtrFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.sliceT and float32)`,
	)

}

// Test Slice > Complex128
func TestCheckBinaryNonConstExprSliceGtrComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.sliceT and complex128)`,
	)

}

// Test Slice > String
func TestCheckBinaryNonConstExprSliceGtrString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.sliceT and string)`,
	)

}

// Test Slice > BoolT
func TestCheckBinaryNonConstExprSliceGtrBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.sliceT and bool)`,
	)

}

// Test Slice > Slice
func TestCheckBinaryNonConstExprSliceGtrSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (operator > not defined on slice)`,
	)

}

// Test Slice > Array
func TestCheckBinaryNonConstExprSliceGtrArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.sliceT and eval.arrayT)`,
	)

}

// Test Slice > Map
func TestCheckBinaryNonConstExprSliceGtrMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.sliceT and eval.mapT)`,
	)

}

// Test Slice > XinterfaceX
func TestCheckBinaryNonConstExprSliceGtrXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.sliceT and eval.XinterfaceX)`,
	)

}

// Test Slice > InterfaceX
func TestCheckBinaryNonConstExprSliceGtrInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.sliceT and eval.interfaceX)`,
	)

}

// Test Slice > InterfaceY
func TestCheckBinaryNonConstExprSliceGtrInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.sliceT and eval.interfaceY)`,
	)

}

// Test Slice > InterfaceZ
func TestCheckBinaryNonConstExprSliceGtrInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.sliceT and eval.interfaceZ)`,
	)

}

// Test Slice > Ptr
func TestCheckBinaryNonConstExprSliceGtrPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.sliceT and *int)`,
	)

}

// Test Slice > Struct
func TestCheckBinaryNonConstExprSliceGtrStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.sliceT and eval.structT)`,
	)

}

// Test Slice > StructUncomp
func TestCheckBinaryNonConstExprSliceGtrStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.sliceT and eval.structUncompT)`,
	)

}

// Test Slice << ConstInt
func TestCheckBinaryNonConstExprSliceShlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 4`, env,
		`invalid operation: a << 4 (shift of type eval.sliceT)`,
	)

}

// Test Slice << ConstRune
func TestCheckBinaryNonConstExprSliceShlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << '@'`, env,
		`invalid operation: a << 64 (shift of type eval.sliceT)`,
	)

}

// Test Slice << ConstFloat
func TestCheckBinaryNonConstExprSliceShlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 2.0`, env,
		`invalid operation: a << 2 (shift of type eval.sliceT)`,
	)

}

// Test Slice << ConstComplex
func TestCheckBinaryNonConstExprSliceShlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: a << 0 (shift of type eval.sliceT)`,
	)

}

// Test Slice << ConstBool
func TestCheckBinaryNonConstExprSliceShlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << true`, env,
		`invalid operation: a << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Slice << ConstString
func TestCheckBinaryNonConstExprSliceShlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: a << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Slice << ConstNil
func TestCheckBinaryNonConstExprSliceShlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Slice << Int
func TestCheckBinaryNonConstExprSliceShlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type int, must be unsigned integer)`,
	)

}

// Test Slice << Float32
func TestCheckBinaryNonConstExprSliceShlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type float32, must be unsigned integer)`,
	)

}

// Test Slice << Complex128
func TestCheckBinaryNonConstExprSliceShlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Slice << String
func TestCheckBinaryNonConstExprSliceShlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type string, must be unsigned integer)`,
	)

}

// Test Slice << BoolT
func TestCheckBinaryNonConstExprSliceShlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type bool, must be unsigned integer)`,
	)

}

// Test Slice << Slice
func TestCheckBinaryNonConstExprSliceShlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.sliceT, must be unsigned integer)`,
	)

}

// Test Slice << Array
func TestCheckBinaryNonConstExprSliceShlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.arrayT, must be unsigned integer)`,
	)

}

// Test Slice << Map
func TestCheckBinaryNonConstExprSliceShlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.mapT, must be unsigned integer)`,
	)

}

// Test Slice << XinterfaceX
func TestCheckBinaryNonConstExprSliceShlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.XinterfaceX, must be unsigned integer)`,
	)

}

// Test Slice << InterfaceX
func TestCheckBinaryNonConstExprSliceShlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceX, must be unsigned integer)`,
	)

}

// Test Slice << InterfaceY
func TestCheckBinaryNonConstExprSliceShlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceY, must be unsigned integer)`,
	)

}

// Test Slice << InterfaceZ
func TestCheckBinaryNonConstExprSliceShlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceZ, must be unsigned integer)`,
	)

}

// Test Slice << Ptr
func TestCheckBinaryNonConstExprSliceShlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type *int, must be unsigned integer)`,
	)

}

// Test Slice << Struct
func TestCheckBinaryNonConstExprSliceShlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structT, must be unsigned integer)`,
	)

}

// Test Slice << StructUncomp
func TestCheckBinaryNonConstExprSliceShlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := sliceT(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structUncompT, must be unsigned integer)`,
	)

}

// Test Array + ConstInt
func TestCheckBinaryNonConstExprArrayAddConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 4`, env,
		`invalid operation: a + 4 (mismatched types eval.arrayT and int)`,
	)

}

// Test Array + ConstRune
func TestCheckBinaryNonConstExprArrayAddConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + '@'`, env,
		`invalid operation: a + rune(64) (mismatched types eval.arrayT and rune)`,
	)

}

// Test Array + ConstFloat
func TestCheckBinaryNonConstExprArrayAddConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 2.0`, env,
		`invalid operation: a + 2 (mismatched types eval.arrayT and float64)`,
	)

}

// Test Array + ConstComplex
func TestCheckBinaryNonConstExprArrayAddConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 8.0i`, env,
		`invalid operation: a + 8i (mismatched types eval.arrayT and complex128)`,
	)

}

// Test Array + ConstBool
func TestCheckBinaryNonConstExprArrayAddConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + true`, env,
		`invalid operation: a + true (mismatched types eval.arrayT and bool)`,
	)

}

// Test Array + ConstString
func TestCheckBinaryNonConstExprArrayAddConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + "abc"`, env,
		`invalid operation: a + "abc" (mismatched types eval.arrayT and string)`,
	)

}

// Test Array + ConstNil
func TestCheckBinaryNonConstExprArrayAddConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + nil`, env,
		`cannot convert nil to type eval.arrayT`,
		`use of untyped nil`,
	)

}

// Test Array + Int
func TestCheckBinaryNonConstExprArrayAddInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.arrayT and int)`,
	)

}

// Test Array + Float32
func TestCheckBinaryNonConstExprArrayAddFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.arrayT and float32)`,
	)

}

// Test Array + Complex128
func TestCheckBinaryNonConstExprArrayAddComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.arrayT and complex128)`,
	)

}

// Test Array + String
func TestCheckBinaryNonConstExprArrayAddString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.arrayT and string)`,
	)

}

// Test Array + BoolT
func TestCheckBinaryNonConstExprArrayAddBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.arrayT and bool)`,
	)

}

// Test Array + Slice
func TestCheckBinaryNonConstExprArrayAddSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.arrayT and eval.sliceT)`,
	)

}

// Test Array + Array
func TestCheckBinaryNonConstExprArrayAddArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (operator + not defined on array)`,
	)

}

// Test Array + Map
func TestCheckBinaryNonConstExprArrayAddMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.arrayT and eval.mapT)`,
	)

}

// Test Array + XinterfaceX
func TestCheckBinaryNonConstExprArrayAddXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.arrayT and eval.XinterfaceX)`,
	)

}

// Test Array + InterfaceX
func TestCheckBinaryNonConstExprArrayAddInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.arrayT and eval.interfaceX)`,
	)

}

// Test Array + InterfaceY
func TestCheckBinaryNonConstExprArrayAddInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.arrayT and eval.interfaceY)`,
	)

}

// Test Array + InterfaceZ
func TestCheckBinaryNonConstExprArrayAddInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.arrayT and eval.interfaceZ)`,
	)

}

// Test Array + Ptr
func TestCheckBinaryNonConstExprArrayAddPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.arrayT and *int)`,
	)

}

// Test Array + Struct
func TestCheckBinaryNonConstExprArrayAddStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.arrayT and eval.structT)`,
	)

}

// Test Array + StructUncomp
func TestCheckBinaryNonConstExprArrayAddStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.arrayT and eval.structUncompT)`,
	)

}

// Test Array & ConstInt
func TestCheckBinaryNonConstExprArrayAndConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 4`, env,
		`invalid operation: a & 4 (mismatched types eval.arrayT and int)`,
	)

}

// Test Array & ConstRune
func TestCheckBinaryNonConstExprArrayAndConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & '@'`, env,
		`invalid operation: a & rune(64) (mismatched types eval.arrayT and rune)`,
	)

}

// Test Array & ConstFloat
func TestCheckBinaryNonConstExprArrayAndConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 2.0`, env,
		`invalid operation: a & 2 (mismatched types eval.arrayT and float64)`,
	)

}

// Test Array & ConstComplex
func TestCheckBinaryNonConstExprArrayAndConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 8.0i`, env,
		`invalid operation: a & 8i (mismatched types eval.arrayT and complex128)`,
	)

}

// Test Array & ConstBool
func TestCheckBinaryNonConstExprArrayAndConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & true`, env,
		`invalid operation: a & true (mismatched types eval.arrayT and bool)`,
	)

}

// Test Array & ConstString
func TestCheckBinaryNonConstExprArrayAndConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & "abc"`, env,
		`invalid operation: a & "abc" (mismatched types eval.arrayT and string)`,
	)

}

// Test Array & ConstNil
func TestCheckBinaryNonConstExprArrayAndConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & nil`, env,
		`cannot convert nil to type eval.arrayT`,
		`use of untyped nil`,
	)

}

// Test Array & Int
func TestCheckBinaryNonConstExprArrayAndInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.arrayT and int)`,
	)

}

// Test Array & Float32
func TestCheckBinaryNonConstExprArrayAndFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.arrayT and float32)`,
	)

}

// Test Array & Complex128
func TestCheckBinaryNonConstExprArrayAndComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.arrayT and complex128)`,
	)

}

// Test Array & String
func TestCheckBinaryNonConstExprArrayAndString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.arrayT and string)`,
	)

}

// Test Array & BoolT
func TestCheckBinaryNonConstExprArrayAndBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.arrayT and bool)`,
	)

}

// Test Array & Slice
func TestCheckBinaryNonConstExprArrayAndSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.arrayT and eval.sliceT)`,
	)

}

// Test Array & Array
func TestCheckBinaryNonConstExprArrayAndArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (operator & not defined on array)`,
	)

}

// Test Array & Map
func TestCheckBinaryNonConstExprArrayAndMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.arrayT and eval.mapT)`,
	)

}

// Test Array & XinterfaceX
func TestCheckBinaryNonConstExprArrayAndXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.arrayT and eval.XinterfaceX)`,
	)

}

// Test Array & InterfaceX
func TestCheckBinaryNonConstExprArrayAndInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.arrayT and eval.interfaceX)`,
	)

}

// Test Array & InterfaceY
func TestCheckBinaryNonConstExprArrayAndInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.arrayT and eval.interfaceY)`,
	)

}

// Test Array & InterfaceZ
func TestCheckBinaryNonConstExprArrayAndInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.arrayT and eval.interfaceZ)`,
	)

}

// Test Array & Ptr
func TestCheckBinaryNonConstExprArrayAndPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.arrayT and *int)`,
	)

}

// Test Array & Struct
func TestCheckBinaryNonConstExprArrayAndStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.arrayT and eval.structT)`,
	)

}

// Test Array & StructUncomp
func TestCheckBinaryNonConstExprArrayAndStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.arrayT and eval.structUncompT)`,
	)

}

// Test Array % ConstInt
func TestCheckBinaryNonConstExprArrayRemConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 4`, env,
		`invalid operation: a % 4 (mismatched types eval.arrayT and int)`,
	)

}

// Test Array % ConstRune
func TestCheckBinaryNonConstExprArrayRemConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % '@'`, env,
		`invalid operation: a % rune(64) (mismatched types eval.arrayT and rune)`,
	)

}

// Test Array % ConstFloat
func TestCheckBinaryNonConstExprArrayRemConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 2.0`, env,
		`invalid operation: a % 2 (mismatched types eval.arrayT and float64)`,
	)

}

// Test Array % ConstComplex
func TestCheckBinaryNonConstExprArrayRemConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 8.0i`, env,
		`invalid operation: a % 8i (mismatched types eval.arrayT and complex128)`,
	)

}

// Test Array % ConstBool
func TestCheckBinaryNonConstExprArrayRemConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % true`, env,
		`invalid operation: a % true (mismatched types eval.arrayT and bool)`,
	)

}

// Test Array % ConstString
func TestCheckBinaryNonConstExprArrayRemConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % "abc"`, env,
		`invalid operation: a % "abc" (mismatched types eval.arrayT and string)`,
	)

}

// Test Array % ConstNil
func TestCheckBinaryNonConstExprArrayRemConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % nil`, env,
		`cannot convert nil to type eval.arrayT`,
		`use of untyped nil`,
	)

}

// Test Array % Int
func TestCheckBinaryNonConstExprArrayRemInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.arrayT and int)`,
	)

}

// Test Array % Float32
func TestCheckBinaryNonConstExprArrayRemFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.arrayT and float32)`,
	)

}

// Test Array % Complex128
func TestCheckBinaryNonConstExprArrayRemComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.arrayT and complex128)`,
	)

}

// Test Array % String
func TestCheckBinaryNonConstExprArrayRemString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.arrayT and string)`,
	)

}

// Test Array % BoolT
func TestCheckBinaryNonConstExprArrayRemBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.arrayT and bool)`,
	)

}

// Test Array % Slice
func TestCheckBinaryNonConstExprArrayRemSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.arrayT and eval.sliceT)`,
	)

}

// Test Array % Array
func TestCheckBinaryNonConstExprArrayRemArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (operator % not defined on array)`,
	)

}

// Test Array % Map
func TestCheckBinaryNonConstExprArrayRemMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.arrayT and eval.mapT)`,
	)

}

// Test Array % XinterfaceX
func TestCheckBinaryNonConstExprArrayRemXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.arrayT and eval.XinterfaceX)`,
	)

}

// Test Array % InterfaceX
func TestCheckBinaryNonConstExprArrayRemInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.arrayT and eval.interfaceX)`,
	)

}

// Test Array % InterfaceY
func TestCheckBinaryNonConstExprArrayRemInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.arrayT and eval.interfaceY)`,
	)

}

// Test Array % InterfaceZ
func TestCheckBinaryNonConstExprArrayRemInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.arrayT and eval.interfaceZ)`,
	)

}

// Test Array % Ptr
func TestCheckBinaryNonConstExprArrayRemPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.arrayT and *int)`,
	)

}

// Test Array % Struct
func TestCheckBinaryNonConstExprArrayRemStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.arrayT and eval.structT)`,
	)

}

// Test Array % StructUncomp
func TestCheckBinaryNonConstExprArrayRemStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.arrayT and eval.structUncompT)`,
	)

}

// Test Array == ConstInt
func TestCheckBinaryNonConstExprArrayEqlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 4`, env,
		`invalid operation: a == 4 (mismatched types eval.arrayT and int)`,
	)

}

// Test Array == ConstRune
func TestCheckBinaryNonConstExprArrayEqlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == '@'`, env,
		`invalid operation: a == rune(64) (mismatched types eval.arrayT and rune)`,
	)

}

// Test Array == ConstFloat
func TestCheckBinaryNonConstExprArrayEqlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 2.0`, env,
		`invalid operation: a == 2 (mismatched types eval.arrayT and float64)`,
	)

}

// Test Array == ConstComplex
func TestCheckBinaryNonConstExprArrayEqlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 8.0i`, env,
		`invalid operation: a == 8i (mismatched types eval.arrayT and complex128)`,
	)

}

// Test Array == ConstBool
func TestCheckBinaryNonConstExprArrayEqlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == true`, env,
		`invalid operation: a == true (mismatched types eval.arrayT and bool)`,
	)

}

// Test Array == ConstString
func TestCheckBinaryNonConstExprArrayEqlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == "abc"`, env,
		`invalid operation: a == "abc" (mismatched types eval.arrayT and string)`,
	)

}

// Test Array == ConstNil
func TestCheckBinaryNonConstExprArrayEqlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == nil`, env,
		`cannot convert nil to type eval.arrayT`,
		`use of untyped nil`,
	)

}

// Test Array == Int
func TestCheckBinaryNonConstExprArrayEqlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.arrayT and int)`,
	)

}

// Test Array == Float32
func TestCheckBinaryNonConstExprArrayEqlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.arrayT and float32)`,
	)

}

// Test Array == Complex128
func TestCheckBinaryNonConstExprArrayEqlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.arrayT and complex128)`,
	)

}

// Test Array == String
func TestCheckBinaryNonConstExprArrayEqlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.arrayT and string)`,
	)

}

// Test Array == BoolT
func TestCheckBinaryNonConstExprArrayEqlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.arrayT and bool)`,
	)

}

// Test Array == Slice
func TestCheckBinaryNonConstExprArrayEqlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.arrayT and eval.sliceT)`,
	)

}

// Test Array == Array
func TestCheckBinaryNonConstExprArrayEqlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test Array == Map
func TestCheckBinaryNonConstExprArrayEqlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.arrayT and eval.mapT)`,
	)

}

// Test Array == XinterfaceX
func TestCheckBinaryNonConstExprArrayEqlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.arrayT and eval.XinterfaceX)`,
	)

}

// Test Array == InterfaceX
func TestCheckBinaryNonConstExprArrayEqlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.arrayT and eval.interfaceX)`,
	)

}

// Test Array == InterfaceY
func TestCheckBinaryNonConstExprArrayEqlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.arrayT and eval.interfaceY)`,
	)

}

// Test Array == InterfaceZ
func TestCheckBinaryNonConstExprArrayEqlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.arrayT and eval.interfaceZ)`,
	)

}

// Test Array == Ptr
func TestCheckBinaryNonConstExprArrayEqlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.arrayT and *int)`,
	)

}

// Test Array == Struct
func TestCheckBinaryNonConstExprArrayEqlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.arrayT and eval.structT)`,
	)

}

// Test Array == StructUncomp
func TestCheckBinaryNonConstExprArrayEqlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.arrayT and eval.structUncompT)`,
	)

}

// Test Array > ConstInt
func TestCheckBinaryNonConstExprArrayGtrConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 4`, env,
		`invalid operation: a > 4 (mismatched types eval.arrayT and int)`,
	)

}

// Test Array > ConstRune
func TestCheckBinaryNonConstExprArrayGtrConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > '@'`, env,
		`invalid operation: a > rune(64) (mismatched types eval.arrayT and rune)`,
	)

}

// Test Array > ConstFloat
func TestCheckBinaryNonConstExprArrayGtrConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 2.0`, env,
		`invalid operation: a > 2 (mismatched types eval.arrayT and float64)`,
	)

}

// Test Array > ConstComplex
func TestCheckBinaryNonConstExprArrayGtrConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 8.0i`, env,
		`invalid operation: a > 8i (mismatched types eval.arrayT and complex128)`,
	)

}

// Test Array > ConstBool
func TestCheckBinaryNonConstExprArrayGtrConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > true`, env,
		`invalid operation: a > true (mismatched types eval.arrayT and bool)`,
	)

}

// Test Array > ConstString
func TestCheckBinaryNonConstExprArrayGtrConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > "abc"`, env,
		`invalid operation: a > "abc" (mismatched types eval.arrayT and string)`,
	)

}

// Test Array > ConstNil
func TestCheckBinaryNonConstExprArrayGtrConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > nil`, env,
		`cannot convert nil to type eval.arrayT`,
		`use of untyped nil`,
	)

}

// Test Array > Int
func TestCheckBinaryNonConstExprArrayGtrInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.arrayT and int)`,
	)

}

// Test Array > Float32
func TestCheckBinaryNonConstExprArrayGtrFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.arrayT and float32)`,
	)

}

// Test Array > Complex128
func TestCheckBinaryNonConstExprArrayGtrComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.arrayT and complex128)`,
	)

}

// Test Array > String
func TestCheckBinaryNonConstExprArrayGtrString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.arrayT and string)`,
	)

}

// Test Array > BoolT
func TestCheckBinaryNonConstExprArrayGtrBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.arrayT and bool)`,
	)

}

// Test Array > Slice
func TestCheckBinaryNonConstExprArrayGtrSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.arrayT and eval.sliceT)`,
	)

}

// Test Array > Array
func TestCheckBinaryNonConstExprArrayGtrArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (operator > not defined on array)`,
	)

}

// Test Array > Map
func TestCheckBinaryNonConstExprArrayGtrMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.arrayT and eval.mapT)`,
	)

}

// Test Array > XinterfaceX
func TestCheckBinaryNonConstExprArrayGtrXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.arrayT and eval.XinterfaceX)`,
	)

}

// Test Array > InterfaceX
func TestCheckBinaryNonConstExprArrayGtrInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.arrayT and eval.interfaceX)`,
	)

}

// Test Array > InterfaceY
func TestCheckBinaryNonConstExprArrayGtrInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.arrayT and eval.interfaceY)`,
	)

}

// Test Array > InterfaceZ
func TestCheckBinaryNonConstExprArrayGtrInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.arrayT and eval.interfaceZ)`,
	)

}

// Test Array > Ptr
func TestCheckBinaryNonConstExprArrayGtrPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.arrayT and *int)`,
	)

}

// Test Array > Struct
func TestCheckBinaryNonConstExprArrayGtrStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.arrayT and eval.structT)`,
	)

}

// Test Array > StructUncomp
func TestCheckBinaryNonConstExprArrayGtrStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.arrayT and eval.structUncompT)`,
	)

}

// Test Array << ConstInt
func TestCheckBinaryNonConstExprArrayShlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 4`, env,
		`invalid operation: a << 4 (shift of type eval.arrayT)`,
	)

}

// Test Array << ConstRune
func TestCheckBinaryNonConstExprArrayShlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << '@'`, env,
		`invalid operation: a << 64 (shift of type eval.arrayT)`,
	)

}

// Test Array << ConstFloat
func TestCheckBinaryNonConstExprArrayShlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 2.0`, env,
		`invalid operation: a << 2 (shift of type eval.arrayT)`,
	)

}

// Test Array << ConstComplex
func TestCheckBinaryNonConstExprArrayShlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: a << 0 (shift of type eval.arrayT)`,
	)

}

// Test Array << ConstBool
func TestCheckBinaryNonConstExprArrayShlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << true`, env,
		`invalid operation: a << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Array << ConstString
func TestCheckBinaryNonConstExprArrayShlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: a << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Array << ConstNil
func TestCheckBinaryNonConstExprArrayShlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Array << Int
func TestCheckBinaryNonConstExprArrayShlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type int, must be unsigned integer)`,
	)

}

// Test Array << Float32
func TestCheckBinaryNonConstExprArrayShlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type float32, must be unsigned integer)`,
	)

}

// Test Array << Complex128
func TestCheckBinaryNonConstExprArrayShlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Array << String
func TestCheckBinaryNonConstExprArrayShlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type string, must be unsigned integer)`,
	)

}

// Test Array << BoolT
func TestCheckBinaryNonConstExprArrayShlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type bool, must be unsigned integer)`,
	)

}

// Test Array << Slice
func TestCheckBinaryNonConstExprArrayShlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.sliceT, must be unsigned integer)`,
	)

}

// Test Array << Array
func TestCheckBinaryNonConstExprArrayShlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.arrayT, must be unsigned integer)`,
	)

}

// Test Array << Map
func TestCheckBinaryNonConstExprArrayShlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.mapT, must be unsigned integer)`,
	)

}

// Test Array << XinterfaceX
func TestCheckBinaryNonConstExprArrayShlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.XinterfaceX, must be unsigned integer)`,
	)

}

// Test Array << InterfaceX
func TestCheckBinaryNonConstExprArrayShlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceX, must be unsigned integer)`,
	)

}

// Test Array << InterfaceY
func TestCheckBinaryNonConstExprArrayShlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceY, must be unsigned integer)`,
	)

}

// Test Array << InterfaceZ
func TestCheckBinaryNonConstExprArrayShlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceZ, must be unsigned integer)`,
	)

}

// Test Array << Ptr
func TestCheckBinaryNonConstExprArrayShlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type *int, must be unsigned integer)`,
	)

}

// Test Array << Struct
func TestCheckBinaryNonConstExprArrayShlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structT, must be unsigned integer)`,
	)

}

// Test Array << StructUncomp
func TestCheckBinaryNonConstExprArrayShlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := arrayT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structUncompT, must be unsigned integer)`,
	)

}

// Test Map + ConstInt
func TestCheckBinaryNonConstExprMapAddConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 4`, env,
		`invalid operation: a + 4 (mismatched types eval.mapT and int)`,
	)

}

// Test Map + ConstRune
func TestCheckBinaryNonConstExprMapAddConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + '@'`, env,
		`invalid operation: a + rune(64) (mismatched types eval.mapT and rune)`,
	)

}

// Test Map + ConstFloat
func TestCheckBinaryNonConstExprMapAddConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 2.0`, env,
		`invalid operation: a + 2 (mismatched types eval.mapT and float64)`,
	)

}

// Test Map + ConstComplex
func TestCheckBinaryNonConstExprMapAddConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 8.0i`, env,
		`invalid operation: a + 8i (mismatched types eval.mapT and complex128)`,
	)

}

// Test Map + ConstBool
func TestCheckBinaryNonConstExprMapAddConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + true`, env,
		`invalid operation: a + true (mismatched types eval.mapT and bool)`,
	)

}

// Test Map + ConstString
func TestCheckBinaryNonConstExprMapAddConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + "abc"`, env,
		`invalid operation: a + "abc" (mismatched types eval.mapT and string)`,
	)

}

// Test Map + ConstNil
func TestCheckBinaryNonConstExprMapAddConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + nil`, env,
		`invalid operation: a + nil (operator + not defined on map)`,
	)

}

// Test Map + Int
func TestCheckBinaryNonConstExprMapAddInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.mapT and int)`,
	)

}

// Test Map + Float32
func TestCheckBinaryNonConstExprMapAddFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.mapT and float32)`,
	)

}

// Test Map + Complex128
func TestCheckBinaryNonConstExprMapAddComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.mapT and complex128)`,
	)

}

// Test Map + String
func TestCheckBinaryNonConstExprMapAddString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.mapT and string)`,
	)

}

// Test Map + BoolT
func TestCheckBinaryNonConstExprMapAddBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.mapT and bool)`,
	)

}

// Test Map + Slice
func TestCheckBinaryNonConstExprMapAddSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.mapT and eval.sliceT)`,
	)

}

// Test Map + Array
func TestCheckBinaryNonConstExprMapAddArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.mapT and eval.arrayT)`,
	)

}

// Test Map + Map
func TestCheckBinaryNonConstExprMapAddMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (operator + not defined on map)`,
	)

}

// Test Map + XinterfaceX
func TestCheckBinaryNonConstExprMapAddXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.mapT and eval.XinterfaceX)`,
	)

}

// Test Map + InterfaceX
func TestCheckBinaryNonConstExprMapAddInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.mapT and eval.interfaceX)`,
	)

}

// Test Map + InterfaceY
func TestCheckBinaryNonConstExprMapAddInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.mapT and eval.interfaceY)`,
	)

}

// Test Map + InterfaceZ
func TestCheckBinaryNonConstExprMapAddInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.mapT and eval.interfaceZ)`,
	)

}

// Test Map + Ptr
func TestCheckBinaryNonConstExprMapAddPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.mapT and *int)`,
	)

}

// Test Map + Struct
func TestCheckBinaryNonConstExprMapAddStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.mapT and eval.structT)`,
	)

}

// Test Map + StructUncomp
func TestCheckBinaryNonConstExprMapAddStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.mapT and eval.structUncompT)`,
	)

}

// Test Map & ConstInt
func TestCheckBinaryNonConstExprMapAndConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 4`, env,
		`invalid operation: a & 4 (mismatched types eval.mapT and int)`,
	)

}

// Test Map & ConstRune
func TestCheckBinaryNonConstExprMapAndConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & '@'`, env,
		`invalid operation: a & rune(64) (mismatched types eval.mapT and rune)`,
	)

}

// Test Map & ConstFloat
func TestCheckBinaryNonConstExprMapAndConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 2.0`, env,
		`invalid operation: a & 2 (mismatched types eval.mapT and float64)`,
	)

}

// Test Map & ConstComplex
func TestCheckBinaryNonConstExprMapAndConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 8.0i`, env,
		`invalid operation: a & 8i (mismatched types eval.mapT and complex128)`,
	)

}

// Test Map & ConstBool
func TestCheckBinaryNonConstExprMapAndConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & true`, env,
		`invalid operation: a & true (mismatched types eval.mapT and bool)`,
	)

}

// Test Map & ConstString
func TestCheckBinaryNonConstExprMapAndConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & "abc"`, env,
		`invalid operation: a & "abc" (mismatched types eval.mapT and string)`,
	)

}

// Test Map & ConstNil
func TestCheckBinaryNonConstExprMapAndConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & nil`, env,
		`invalid operation: a & nil (operator & not defined on map)`,
	)

}

// Test Map & Int
func TestCheckBinaryNonConstExprMapAndInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.mapT and int)`,
	)

}

// Test Map & Float32
func TestCheckBinaryNonConstExprMapAndFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.mapT and float32)`,
	)

}

// Test Map & Complex128
func TestCheckBinaryNonConstExprMapAndComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.mapT and complex128)`,
	)

}

// Test Map & String
func TestCheckBinaryNonConstExprMapAndString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.mapT and string)`,
	)

}

// Test Map & BoolT
func TestCheckBinaryNonConstExprMapAndBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.mapT and bool)`,
	)

}

// Test Map & Slice
func TestCheckBinaryNonConstExprMapAndSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.mapT and eval.sliceT)`,
	)

}

// Test Map & Array
func TestCheckBinaryNonConstExprMapAndArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.mapT and eval.arrayT)`,
	)

}

// Test Map & Map
func TestCheckBinaryNonConstExprMapAndMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (operator & not defined on map)`,
	)

}

// Test Map & XinterfaceX
func TestCheckBinaryNonConstExprMapAndXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.mapT and eval.XinterfaceX)`,
	)

}

// Test Map & InterfaceX
func TestCheckBinaryNonConstExprMapAndInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.mapT and eval.interfaceX)`,
	)

}

// Test Map & InterfaceY
func TestCheckBinaryNonConstExprMapAndInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.mapT and eval.interfaceY)`,
	)

}

// Test Map & InterfaceZ
func TestCheckBinaryNonConstExprMapAndInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.mapT and eval.interfaceZ)`,
	)

}

// Test Map & Ptr
func TestCheckBinaryNonConstExprMapAndPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.mapT and *int)`,
	)

}

// Test Map & Struct
func TestCheckBinaryNonConstExprMapAndStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.mapT and eval.structT)`,
	)

}

// Test Map & StructUncomp
func TestCheckBinaryNonConstExprMapAndStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.mapT and eval.structUncompT)`,
	)

}

// Test Map % ConstInt
func TestCheckBinaryNonConstExprMapRemConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 4`, env,
		`invalid operation: a % 4 (mismatched types eval.mapT and int)`,
	)

}

// Test Map % ConstRune
func TestCheckBinaryNonConstExprMapRemConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % '@'`, env,
		`invalid operation: a % rune(64) (mismatched types eval.mapT and rune)`,
	)

}

// Test Map % ConstFloat
func TestCheckBinaryNonConstExprMapRemConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 2.0`, env,
		`invalid operation: a % 2 (mismatched types eval.mapT and float64)`,
	)

}

// Test Map % ConstComplex
func TestCheckBinaryNonConstExprMapRemConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 8.0i`, env,
		`invalid operation: a % 8i (mismatched types eval.mapT and complex128)`,
	)

}

// Test Map % ConstBool
func TestCheckBinaryNonConstExprMapRemConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % true`, env,
		`invalid operation: a % true (mismatched types eval.mapT and bool)`,
	)

}

// Test Map % ConstString
func TestCheckBinaryNonConstExprMapRemConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % "abc"`, env,
		`invalid operation: a % "abc" (mismatched types eval.mapT and string)`,
	)

}

// Test Map % ConstNil
func TestCheckBinaryNonConstExprMapRemConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % nil`, env,
		`invalid operation: a % nil (operator % not defined on map)`,
	)

}

// Test Map % Int
func TestCheckBinaryNonConstExprMapRemInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.mapT and int)`,
	)

}

// Test Map % Float32
func TestCheckBinaryNonConstExprMapRemFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.mapT and float32)`,
	)

}

// Test Map % Complex128
func TestCheckBinaryNonConstExprMapRemComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.mapT and complex128)`,
	)

}

// Test Map % String
func TestCheckBinaryNonConstExprMapRemString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.mapT and string)`,
	)

}

// Test Map % BoolT
func TestCheckBinaryNonConstExprMapRemBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.mapT and bool)`,
	)

}

// Test Map % Slice
func TestCheckBinaryNonConstExprMapRemSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.mapT and eval.sliceT)`,
	)

}

// Test Map % Array
func TestCheckBinaryNonConstExprMapRemArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.mapT and eval.arrayT)`,
	)

}

// Test Map % Map
func TestCheckBinaryNonConstExprMapRemMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (operator % not defined on map)`,
	)

}

// Test Map % XinterfaceX
func TestCheckBinaryNonConstExprMapRemXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.mapT and eval.XinterfaceX)`,
	)

}

// Test Map % InterfaceX
func TestCheckBinaryNonConstExprMapRemInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.mapT and eval.interfaceX)`,
	)

}

// Test Map % InterfaceY
func TestCheckBinaryNonConstExprMapRemInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.mapT and eval.interfaceY)`,
	)

}

// Test Map % InterfaceZ
func TestCheckBinaryNonConstExprMapRemInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.mapT and eval.interfaceZ)`,
	)

}

// Test Map % Ptr
func TestCheckBinaryNonConstExprMapRemPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.mapT and *int)`,
	)

}

// Test Map % Struct
func TestCheckBinaryNonConstExprMapRemStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.mapT and eval.structT)`,
	)

}

// Test Map % StructUncomp
func TestCheckBinaryNonConstExprMapRemStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.mapT and eval.structUncompT)`,
	)

}

// Test Map == ConstInt
func TestCheckBinaryNonConstExprMapEqlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 4`, env,
		`invalid operation: a == 4 (mismatched types eval.mapT and int)`,
	)

}

// Test Map == ConstRune
func TestCheckBinaryNonConstExprMapEqlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == '@'`, env,
		`invalid operation: a == rune(64) (mismatched types eval.mapT and rune)`,
	)

}

// Test Map == ConstFloat
func TestCheckBinaryNonConstExprMapEqlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 2.0`, env,
		`invalid operation: a == 2 (mismatched types eval.mapT and float64)`,
	)

}

// Test Map == ConstComplex
func TestCheckBinaryNonConstExprMapEqlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 8.0i`, env,
		`invalid operation: a == 8i (mismatched types eval.mapT and complex128)`,
	)

}

// Test Map == ConstBool
func TestCheckBinaryNonConstExprMapEqlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == true`, env,
		`invalid operation: a == true (mismatched types eval.mapT and bool)`,
	)

}

// Test Map == ConstString
func TestCheckBinaryNonConstExprMapEqlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == "abc"`, env,
		`invalid operation: a == "abc" (mismatched types eval.mapT and string)`,
	)

}

// Test Map == ConstNil
func TestCheckBinaryNonConstExprMapEqlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == nil`, env, reflect.TypeOf(a == nil))
}

// Test Map == Int
func TestCheckBinaryNonConstExprMapEqlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.mapT and int)`,
	)

}

// Test Map == Float32
func TestCheckBinaryNonConstExprMapEqlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.mapT and float32)`,
	)

}

// Test Map == Complex128
func TestCheckBinaryNonConstExprMapEqlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.mapT and complex128)`,
	)

}

// Test Map == String
func TestCheckBinaryNonConstExprMapEqlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.mapT and string)`,
	)

}

// Test Map == BoolT
func TestCheckBinaryNonConstExprMapEqlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.mapT and bool)`,
	)

}

// Test Map == Slice
func TestCheckBinaryNonConstExprMapEqlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.mapT and eval.sliceT)`,
	)

}

// Test Map == Array
func TestCheckBinaryNonConstExprMapEqlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.mapT and eval.arrayT)`,
	)

}

// Test Map == Map
func TestCheckBinaryNonConstExprMapEqlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (map can only be compared to nil)`,
	)

}

// Test Map == XinterfaceX
func TestCheckBinaryNonConstExprMapEqlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.mapT and eval.XinterfaceX)`,
	)

}

// Test Map == InterfaceX
func TestCheckBinaryNonConstExprMapEqlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.mapT and eval.interfaceX)`,
	)

}

// Test Map == InterfaceY
func TestCheckBinaryNonConstExprMapEqlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.mapT and eval.interfaceY)`,
	)

}

// Test Map == InterfaceZ
func TestCheckBinaryNonConstExprMapEqlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.mapT and eval.interfaceZ)`,
	)

}

// Test Map == Ptr
func TestCheckBinaryNonConstExprMapEqlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.mapT and *int)`,
	)

}

// Test Map == Struct
func TestCheckBinaryNonConstExprMapEqlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.mapT and eval.structT)`,
	)

}

// Test Map == StructUncomp
func TestCheckBinaryNonConstExprMapEqlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.mapT and eval.structUncompT)`,
	)

}

// Test Map > ConstInt
func TestCheckBinaryNonConstExprMapGtrConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 4`, env,
		`invalid operation: a > 4 (mismatched types eval.mapT and int)`,
	)

}

// Test Map > ConstRune
func TestCheckBinaryNonConstExprMapGtrConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > '@'`, env,
		`invalid operation: a > rune(64) (mismatched types eval.mapT and rune)`,
	)

}

// Test Map > ConstFloat
func TestCheckBinaryNonConstExprMapGtrConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 2.0`, env,
		`invalid operation: a > 2 (mismatched types eval.mapT and float64)`,
	)

}

// Test Map > ConstComplex
func TestCheckBinaryNonConstExprMapGtrConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 8.0i`, env,
		`invalid operation: a > 8i (mismatched types eval.mapT and complex128)`,
	)

}

// Test Map > ConstBool
func TestCheckBinaryNonConstExprMapGtrConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > true`, env,
		`invalid operation: a > true (mismatched types eval.mapT and bool)`,
	)

}

// Test Map > ConstString
func TestCheckBinaryNonConstExprMapGtrConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > "abc"`, env,
		`invalid operation: a > "abc" (mismatched types eval.mapT and string)`,
	)

}

// Test Map > ConstNil
func TestCheckBinaryNonConstExprMapGtrConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > nil`, env,
		`invalid operation: a > nil (operator > not defined on map)`,
	)

}

// Test Map > Int
func TestCheckBinaryNonConstExprMapGtrInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.mapT and int)`,
	)

}

// Test Map > Float32
func TestCheckBinaryNonConstExprMapGtrFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.mapT and float32)`,
	)

}

// Test Map > Complex128
func TestCheckBinaryNonConstExprMapGtrComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.mapT and complex128)`,
	)

}

// Test Map > String
func TestCheckBinaryNonConstExprMapGtrString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.mapT and string)`,
	)

}

// Test Map > BoolT
func TestCheckBinaryNonConstExprMapGtrBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.mapT and bool)`,
	)

}

// Test Map > Slice
func TestCheckBinaryNonConstExprMapGtrSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.mapT and eval.sliceT)`,
	)

}

// Test Map > Array
func TestCheckBinaryNonConstExprMapGtrArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.mapT and eval.arrayT)`,
	)

}

// Test Map > Map
func TestCheckBinaryNonConstExprMapGtrMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (operator > not defined on map)`,
	)

}

// Test Map > XinterfaceX
func TestCheckBinaryNonConstExprMapGtrXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.mapT and eval.XinterfaceX)`,
	)

}

// Test Map > InterfaceX
func TestCheckBinaryNonConstExprMapGtrInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.mapT and eval.interfaceX)`,
	)

}

// Test Map > InterfaceY
func TestCheckBinaryNonConstExprMapGtrInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.mapT and eval.interfaceY)`,
	)

}

// Test Map > InterfaceZ
func TestCheckBinaryNonConstExprMapGtrInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.mapT and eval.interfaceZ)`,
	)

}

// Test Map > Ptr
func TestCheckBinaryNonConstExprMapGtrPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.mapT and *int)`,
	)

}

// Test Map > Struct
func TestCheckBinaryNonConstExprMapGtrStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.mapT and eval.structT)`,
	)

}

// Test Map > StructUncomp
func TestCheckBinaryNonConstExprMapGtrStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.mapT and eval.structUncompT)`,
	)

}

// Test Map << ConstInt
func TestCheckBinaryNonConstExprMapShlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 4`, env,
		`invalid operation: a << 4 (shift of type eval.mapT)`,
	)

}

// Test Map << ConstRune
func TestCheckBinaryNonConstExprMapShlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << '@'`, env,
		`invalid operation: a << 64 (shift of type eval.mapT)`,
	)

}

// Test Map << ConstFloat
func TestCheckBinaryNonConstExprMapShlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 2.0`, env,
		`invalid operation: a << 2 (shift of type eval.mapT)`,
	)

}

// Test Map << ConstComplex
func TestCheckBinaryNonConstExprMapShlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: a << 0 (shift of type eval.mapT)`,
	)

}

// Test Map << ConstBool
func TestCheckBinaryNonConstExprMapShlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << true`, env,
		`invalid operation: a << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Map << ConstString
func TestCheckBinaryNonConstExprMapShlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: a << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Map << ConstNil
func TestCheckBinaryNonConstExprMapShlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Map << Int
func TestCheckBinaryNonConstExprMapShlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type int, must be unsigned integer)`,
	)

}

// Test Map << Float32
func TestCheckBinaryNonConstExprMapShlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type float32, must be unsigned integer)`,
	)

}

// Test Map << Complex128
func TestCheckBinaryNonConstExprMapShlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Map << String
func TestCheckBinaryNonConstExprMapShlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type string, must be unsigned integer)`,
	)

}

// Test Map << BoolT
func TestCheckBinaryNonConstExprMapShlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type bool, must be unsigned integer)`,
	)

}

// Test Map << Slice
func TestCheckBinaryNonConstExprMapShlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.sliceT, must be unsigned integer)`,
	)

}

// Test Map << Array
func TestCheckBinaryNonConstExprMapShlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.arrayT, must be unsigned integer)`,
	)

}

// Test Map << Map
func TestCheckBinaryNonConstExprMapShlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.mapT, must be unsigned integer)`,
	)

}

// Test Map << XinterfaceX
func TestCheckBinaryNonConstExprMapShlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.XinterfaceX, must be unsigned integer)`,
	)

}

// Test Map << InterfaceX
func TestCheckBinaryNonConstExprMapShlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceX, must be unsigned integer)`,
	)

}

// Test Map << InterfaceY
func TestCheckBinaryNonConstExprMapShlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceY, must be unsigned integer)`,
	)

}

// Test Map << InterfaceZ
func TestCheckBinaryNonConstExprMapShlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceZ, must be unsigned integer)`,
	)

}

// Test Map << Ptr
func TestCheckBinaryNonConstExprMapShlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type *int, must be unsigned integer)`,
	)

}

// Test Map << Struct
func TestCheckBinaryNonConstExprMapShlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structT, must be unsigned integer)`,
	)

}

// Test Map << StructUncomp
func TestCheckBinaryNonConstExprMapShlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := mapT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structUncompT, must be unsigned integer)`,
	)

}

// Test XinterfaceX + ConstInt
func TestCheckBinaryNonConstExprXinterfaceXAddConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a + 4`, env, reflect.TypeOf(a + 4))
}

// Test XinterfaceX + ConstRune
func TestCheckBinaryNonConstExprXinterfaceXAddConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a + '@'`, env, reflect.TypeOf(a + '@'))
}

// Test XinterfaceX + ConstFloat
func TestCheckBinaryNonConstExprXinterfaceXAddConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a + 2.0`, env, reflect.TypeOf(a + 2.0))
}

// Test XinterfaceX + ConstComplex
func TestCheckBinaryNonConstExprXinterfaceXAddConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test XinterfaceX + ConstBool
func TestCheckBinaryNonConstExprXinterfaceXAddConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + true`, env,
		`cannot convert true to type eval.XinterfaceX`,
		`invalid operation: a + true (mismatched types eval.XinterfaceX and bool)`,
	)

}

// Test XinterfaceX + ConstString
func TestCheckBinaryNonConstExprXinterfaceXAddConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + "abc"`, env,
		`cannot convert "abc" to type eval.XinterfaceX`,
		`invalid operation: a + "abc" (mismatched types eval.XinterfaceX and string)`,
	)

}

// Test XinterfaceX + ConstNil
func TestCheckBinaryNonConstExprXinterfaceXAddConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + nil`, env,
		`cannot convert nil to type eval.XinterfaceX`,
	)

}

// Test XinterfaceX + Int
func TestCheckBinaryNonConstExprXinterfaceXAddInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.XinterfaceX and int)`,
	)

}

// Test XinterfaceX + Float32
func TestCheckBinaryNonConstExprXinterfaceXAddFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.XinterfaceX and float32)`,
	)

}

// Test XinterfaceX + Complex128
func TestCheckBinaryNonConstExprXinterfaceXAddComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.XinterfaceX and complex128)`,
	)

}

// Test XinterfaceX + String
func TestCheckBinaryNonConstExprXinterfaceXAddString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.XinterfaceX and string)`,
	)

}

// Test XinterfaceX + BoolT
func TestCheckBinaryNonConstExprXinterfaceXAddBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.XinterfaceX and bool)`,
	)

}

// Test XinterfaceX + Slice
func TestCheckBinaryNonConstExprXinterfaceXAddSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.XinterfaceX and eval.sliceT)`,
	)

}

// Test XinterfaceX + Array
func TestCheckBinaryNonConstExprXinterfaceXAddArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.XinterfaceX and eval.arrayT)`,
	)

}

// Test XinterfaceX + Map
func TestCheckBinaryNonConstExprXinterfaceXAddMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.XinterfaceX and eval.mapT)`,
	)

}

// Test XinterfaceX + XinterfaceX
func TestCheckBinaryNonConstExprXinterfaceXAddXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a + b`, env, reflect.TypeOf(a + b))
}

// Test XinterfaceX + InterfaceX
func TestCheckBinaryNonConstExprXinterfaceXAddInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.XinterfaceX and eval.interfaceX)`,
	)

}

// Test XinterfaceX + InterfaceY
func TestCheckBinaryNonConstExprXinterfaceXAddInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.XinterfaceX and eval.interfaceY)`,
	)

}

// Test XinterfaceX + InterfaceZ
func TestCheckBinaryNonConstExprXinterfaceXAddInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.XinterfaceX and eval.interfaceZ)`,
	)

}

// Test XinterfaceX + Ptr
func TestCheckBinaryNonConstExprXinterfaceXAddPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.XinterfaceX and *int)`,
	)

}

// Test XinterfaceX + Struct
func TestCheckBinaryNonConstExprXinterfaceXAddStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.XinterfaceX and eval.structT)`,
	)

}

// Test XinterfaceX + StructUncomp
func TestCheckBinaryNonConstExprXinterfaceXAddStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.XinterfaceX and eval.structUncompT)`,
	)

}

// Test XinterfaceX & ConstInt
func TestCheckBinaryNonConstExprXinterfaceXAndConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a & 4`, env, reflect.TypeOf(a & 4))
}

// Test XinterfaceX & ConstRune
func TestCheckBinaryNonConstExprXinterfaceXAndConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a & '@'`, env, reflect.TypeOf(a & '@'))
}

// Test XinterfaceX & ConstFloat
func TestCheckBinaryNonConstExprXinterfaceXAndConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a & 2.0`, env, reflect.TypeOf(a & 2.0))
}

// Test XinterfaceX & ConstComplex
func TestCheckBinaryNonConstExprXinterfaceXAndConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test XinterfaceX & ConstBool
func TestCheckBinaryNonConstExprXinterfaceXAndConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & true`, env,
		`cannot convert true to type eval.XinterfaceX`,
		`invalid operation: a & true (mismatched types eval.XinterfaceX and bool)`,
	)

}

// Test XinterfaceX & ConstString
func TestCheckBinaryNonConstExprXinterfaceXAndConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & "abc"`, env,
		`cannot convert "abc" to type eval.XinterfaceX`,
		`invalid operation: a & "abc" (mismatched types eval.XinterfaceX and string)`,
	)

}

// Test XinterfaceX & ConstNil
func TestCheckBinaryNonConstExprXinterfaceXAndConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & nil`, env,
		`cannot convert nil to type eval.XinterfaceX`,
	)

}

// Test XinterfaceX & Int
func TestCheckBinaryNonConstExprXinterfaceXAndInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.XinterfaceX and int)`,
	)

}

// Test XinterfaceX & Float32
func TestCheckBinaryNonConstExprXinterfaceXAndFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.XinterfaceX and float32)`,
	)

}

// Test XinterfaceX & Complex128
func TestCheckBinaryNonConstExprXinterfaceXAndComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.XinterfaceX and complex128)`,
	)

}

// Test XinterfaceX & String
func TestCheckBinaryNonConstExprXinterfaceXAndString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.XinterfaceX and string)`,
	)

}

// Test XinterfaceX & BoolT
func TestCheckBinaryNonConstExprXinterfaceXAndBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.XinterfaceX and bool)`,
	)

}

// Test XinterfaceX & Slice
func TestCheckBinaryNonConstExprXinterfaceXAndSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.XinterfaceX and eval.sliceT)`,
	)

}

// Test XinterfaceX & Array
func TestCheckBinaryNonConstExprXinterfaceXAndArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.XinterfaceX and eval.arrayT)`,
	)

}

// Test XinterfaceX & Map
func TestCheckBinaryNonConstExprXinterfaceXAndMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.XinterfaceX and eval.mapT)`,
	)

}

// Test XinterfaceX & XinterfaceX
func TestCheckBinaryNonConstExprXinterfaceXAndXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a & b`, env, reflect.TypeOf(a & b))
}

// Test XinterfaceX & InterfaceX
func TestCheckBinaryNonConstExprXinterfaceXAndInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.XinterfaceX and eval.interfaceX)`,
	)

}

// Test XinterfaceX & InterfaceY
func TestCheckBinaryNonConstExprXinterfaceXAndInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.XinterfaceX and eval.interfaceY)`,
	)

}

// Test XinterfaceX & InterfaceZ
func TestCheckBinaryNonConstExprXinterfaceXAndInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.XinterfaceX and eval.interfaceZ)`,
	)

}

// Test XinterfaceX & Ptr
func TestCheckBinaryNonConstExprXinterfaceXAndPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.XinterfaceX and *int)`,
	)

}

// Test XinterfaceX & Struct
func TestCheckBinaryNonConstExprXinterfaceXAndStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.XinterfaceX and eval.structT)`,
	)

}

// Test XinterfaceX & StructUncomp
func TestCheckBinaryNonConstExprXinterfaceXAndStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.XinterfaceX and eval.structUncompT)`,
	)

}

// Test XinterfaceX % ConstInt
func TestCheckBinaryNonConstExprXinterfaceXRemConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a % 4`, env, reflect.TypeOf(a % 4))
}

// Test XinterfaceX % ConstRune
func TestCheckBinaryNonConstExprXinterfaceXRemConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a % '@'`, env, reflect.TypeOf(a % '@'))
}

// Test XinterfaceX % ConstFloat
func TestCheckBinaryNonConstExprXinterfaceXRemConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a % 2.0`, env, reflect.TypeOf(a % 2.0))
}

// Test XinterfaceX % ConstComplex
func TestCheckBinaryNonConstExprXinterfaceXRemConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 8.0i`, env,
		`constant 0+8i truncated to real`,
		`division by zero`,
	)

}

// Test XinterfaceX % ConstBool
func TestCheckBinaryNonConstExprXinterfaceXRemConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % true`, env,
		`cannot convert true to type eval.XinterfaceX`,
		`invalid operation: a % true (mismatched types eval.XinterfaceX and bool)`,
	)

}

// Test XinterfaceX % ConstString
func TestCheckBinaryNonConstExprXinterfaceXRemConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % "abc"`, env,
		`cannot convert "abc" to type eval.XinterfaceX`,
		`invalid operation: a % "abc" (mismatched types eval.XinterfaceX and string)`,
	)

}

// Test XinterfaceX % ConstNil
func TestCheckBinaryNonConstExprXinterfaceXRemConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % nil`, env,
		`cannot convert nil to type eval.XinterfaceX`,
	)

}

// Test XinterfaceX % Int
func TestCheckBinaryNonConstExprXinterfaceXRemInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.XinterfaceX and int)`,
	)

}

// Test XinterfaceX % Float32
func TestCheckBinaryNonConstExprXinterfaceXRemFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.XinterfaceX and float32)`,
	)

}

// Test XinterfaceX % Complex128
func TestCheckBinaryNonConstExprXinterfaceXRemComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.XinterfaceX and complex128)`,
	)

}

// Test XinterfaceX % String
func TestCheckBinaryNonConstExprXinterfaceXRemString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.XinterfaceX and string)`,
	)

}

// Test XinterfaceX % BoolT
func TestCheckBinaryNonConstExprXinterfaceXRemBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.XinterfaceX and bool)`,
	)

}

// Test XinterfaceX % Slice
func TestCheckBinaryNonConstExprXinterfaceXRemSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.XinterfaceX and eval.sliceT)`,
	)

}

// Test XinterfaceX % Array
func TestCheckBinaryNonConstExprXinterfaceXRemArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.XinterfaceX and eval.arrayT)`,
	)

}

// Test XinterfaceX % Map
func TestCheckBinaryNonConstExprXinterfaceXRemMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.XinterfaceX and eval.mapT)`,
	)

}

// Test XinterfaceX % XinterfaceX
func TestCheckBinaryNonConstExprXinterfaceXRemXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a % b`, env, reflect.TypeOf(a % b))
}

// Test XinterfaceX % InterfaceX
func TestCheckBinaryNonConstExprXinterfaceXRemInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.XinterfaceX and eval.interfaceX)`,
	)

}

// Test XinterfaceX % InterfaceY
func TestCheckBinaryNonConstExprXinterfaceXRemInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.XinterfaceX and eval.interfaceY)`,
	)

}

// Test XinterfaceX % InterfaceZ
func TestCheckBinaryNonConstExprXinterfaceXRemInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.XinterfaceX and eval.interfaceZ)`,
	)

}

// Test XinterfaceX % Ptr
func TestCheckBinaryNonConstExprXinterfaceXRemPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.XinterfaceX and *int)`,
	)

}

// Test XinterfaceX % Struct
func TestCheckBinaryNonConstExprXinterfaceXRemStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.XinterfaceX and eval.structT)`,
	)

}

// Test XinterfaceX % StructUncomp
func TestCheckBinaryNonConstExprXinterfaceXRemStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.XinterfaceX and eval.structUncompT)`,
	)

}

// Test XinterfaceX == ConstInt
func TestCheckBinaryNonConstExprXinterfaceXEqlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == 4`, env, reflect.TypeOf(a == 4))
}

// Test XinterfaceX == ConstRune
func TestCheckBinaryNonConstExprXinterfaceXEqlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == '@'`, env, reflect.TypeOf(a == '@'))
}

// Test XinterfaceX == ConstFloat
func TestCheckBinaryNonConstExprXinterfaceXEqlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == 2.0`, env, reflect.TypeOf(a == 2.0))
}

// Test XinterfaceX == ConstComplex
func TestCheckBinaryNonConstExprXinterfaceXEqlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test XinterfaceX == ConstBool
func TestCheckBinaryNonConstExprXinterfaceXEqlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == true`, env,
		`cannot convert true to type eval.XinterfaceX`,
		`invalid operation: a == true (mismatched types eval.XinterfaceX and bool)`,
	)

}

// Test XinterfaceX == ConstString
func TestCheckBinaryNonConstExprXinterfaceXEqlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == "abc"`, env,
		`cannot convert "abc" to type eval.XinterfaceX`,
		`invalid operation: a == "abc" (mismatched types eval.XinterfaceX and string)`,
	)

}

// Test XinterfaceX == ConstNil
func TestCheckBinaryNonConstExprXinterfaceXEqlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == nil`, env,
		`cannot convert nil to type eval.XinterfaceX`,
	)

}

// Test XinterfaceX == Int
func TestCheckBinaryNonConstExprXinterfaceXEqlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.XinterfaceX and int)`,
	)

}

// Test XinterfaceX == Float32
func TestCheckBinaryNonConstExprXinterfaceXEqlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.XinterfaceX and float32)`,
	)

}

// Test XinterfaceX == Complex128
func TestCheckBinaryNonConstExprXinterfaceXEqlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.XinterfaceX and complex128)`,
	)

}

// Test XinterfaceX == String
func TestCheckBinaryNonConstExprXinterfaceXEqlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.XinterfaceX and string)`,
	)

}

// Test XinterfaceX == BoolT
func TestCheckBinaryNonConstExprXinterfaceXEqlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.XinterfaceX and bool)`,
	)

}

// Test XinterfaceX == Slice
func TestCheckBinaryNonConstExprXinterfaceXEqlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.XinterfaceX and eval.sliceT)`,
	)

}

// Test XinterfaceX == Array
func TestCheckBinaryNonConstExprXinterfaceXEqlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.XinterfaceX and eval.arrayT)`,
	)

}

// Test XinterfaceX == Map
func TestCheckBinaryNonConstExprXinterfaceXEqlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.XinterfaceX and eval.mapT)`,
	)

}

// Test XinterfaceX == XinterfaceX
func TestCheckBinaryNonConstExprXinterfaceXEqlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test XinterfaceX == InterfaceX
func TestCheckBinaryNonConstExprXinterfaceXEqlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test XinterfaceX == InterfaceY
func TestCheckBinaryNonConstExprXinterfaceXEqlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test XinterfaceX == InterfaceZ
func TestCheckBinaryNonConstExprXinterfaceXEqlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.XinterfaceX and eval.interfaceZ)`,
	)

}

// Test XinterfaceX == Ptr
func TestCheckBinaryNonConstExprXinterfaceXEqlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.XinterfaceX and *int)`,
	)

}

// Test XinterfaceX == Struct
func TestCheckBinaryNonConstExprXinterfaceXEqlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.XinterfaceX and eval.structT)`,
	)

}

// Test XinterfaceX == StructUncomp
func TestCheckBinaryNonConstExprXinterfaceXEqlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.XinterfaceX and eval.structUncompT)`,
	)

}

// Test XinterfaceX > ConstInt
func TestCheckBinaryNonConstExprXinterfaceXGtrConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a > 4`, env, reflect.TypeOf(a > 4))
}

// Test XinterfaceX > ConstRune
func TestCheckBinaryNonConstExprXinterfaceXGtrConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a > '@'`, env, reflect.TypeOf(a > '@'))
}

// Test XinterfaceX > ConstFloat
func TestCheckBinaryNonConstExprXinterfaceXGtrConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a > 2.0`, env, reflect.TypeOf(a > 2.0))
}

// Test XinterfaceX > ConstComplex
func TestCheckBinaryNonConstExprXinterfaceXGtrConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test XinterfaceX > ConstBool
func TestCheckBinaryNonConstExprXinterfaceXGtrConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > true`, env,
		`cannot convert true to type eval.XinterfaceX`,
		`invalid operation: a > true (mismatched types eval.XinterfaceX and bool)`,
	)

}

// Test XinterfaceX > ConstString
func TestCheckBinaryNonConstExprXinterfaceXGtrConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > "abc"`, env,
		`cannot convert "abc" to type eval.XinterfaceX`,
		`invalid operation: a > "abc" (mismatched types eval.XinterfaceX and string)`,
	)

}

// Test XinterfaceX > ConstNil
func TestCheckBinaryNonConstExprXinterfaceXGtrConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > nil`, env,
		`cannot convert nil to type eval.XinterfaceX`,
	)

}

// Test XinterfaceX > Int
func TestCheckBinaryNonConstExprXinterfaceXGtrInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.XinterfaceX and int)`,
	)

}

// Test XinterfaceX > Float32
func TestCheckBinaryNonConstExprXinterfaceXGtrFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.XinterfaceX and float32)`,
	)

}

// Test XinterfaceX > Complex128
func TestCheckBinaryNonConstExprXinterfaceXGtrComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.XinterfaceX and complex128)`,
	)

}

// Test XinterfaceX > String
func TestCheckBinaryNonConstExprXinterfaceXGtrString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.XinterfaceX and string)`,
	)

}

// Test XinterfaceX > BoolT
func TestCheckBinaryNonConstExprXinterfaceXGtrBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.XinterfaceX and bool)`,
	)

}

// Test XinterfaceX > Slice
func TestCheckBinaryNonConstExprXinterfaceXGtrSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.XinterfaceX and eval.sliceT)`,
	)

}

// Test XinterfaceX > Array
func TestCheckBinaryNonConstExprXinterfaceXGtrArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.XinterfaceX and eval.arrayT)`,
	)

}

// Test XinterfaceX > Map
func TestCheckBinaryNonConstExprXinterfaceXGtrMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.XinterfaceX and eval.mapT)`,
	)

}

// Test XinterfaceX > XinterfaceX
func TestCheckBinaryNonConstExprXinterfaceXGtrXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a > b`, env, reflect.TypeOf(a > b))
}

// Test XinterfaceX > InterfaceX
func TestCheckBinaryNonConstExprXinterfaceXGtrInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: eval.interfaceX(a) > b (operator > not defined on interface)`,
	)

}

// Test XinterfaceX > InterfaceY
func TestCheckBinaryNonConstExprXinterfaceXGtrInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: eval.interfaceY(a) > b (operator > not defined on interface)`,
	)

}

// Test XinterfaceX > InterfaceZ
func TestCheckBinaryNonConstExprXinterfaceXGtrInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.XinterfaceX and eval.interfaceZ)`,
	)

}

// Test XinterfaceX > Ptr
func TestCheckBinaryNonConstExprXinterfaceXGtrPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.XinterfaceX and *int)`,
	)

}

// Test XinterfaceX > Struct
func TestCheckBinaryNonConstExprXinterfaceXGtrStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.XinterfaceX and eval.structT)`,
	)

}

// Test XinterfaceX > StructUncomp
func TestCheckBinaryNonConstExprXinterfaceXGtrStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.XinterfaceX and eval.structUncompT)`,
	)

}

// Test XinterfaceX << ConstInt
func TestCheckBinaryNonConstExprXinterfaceXShlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a << 4`, env, reflect.TypeOf(a << 4))
}

// Test XinterfaceX << ConstRune
func TestCheckBinaryNonConstExprXinterfaceXShlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a << '@'`, env, reflect.TypeOf(a << '@'))
}

// Test XinterfaceX << ConstFloat
func TestCheckBinaryNonConstExprXinterfaceXShlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a << 2.0`, env, reflect.TypeOf(a << 2.0))
}

// Test XinterfaceX << ConstComplex
func TestCheckBinaryNonConstExprXinterfaceXShlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 8.0i`, env,
		`constant 0+8i truncated to real`,
	)

}

// Test XinterfaceX << ConstBool
func TestCheckBinaryNonConstExprXinterfaceXShlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << true`, env,
		`invalid operation: a << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test XinterfaceX << ConstString
func TestCheckBinaryNonConstExprXinterfaceXShlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: a << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test XinterfaceX << ConstNil
func TestCheckBinaryNonConstExprXinterfaceXShlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test XinterfaceX << Int
func TestCheckBinaryNonConstExprXinterfaceXShlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type int, must be unsigned integer)`,
	)

}

// Test XinterfaceX << Float32
func TestCheckBinaryNonConstExprXinterfaceXShlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type float32, must be unsigned integer)`,
	)

}

// Test XinterfaceX << Complex128
func TestCheckBinaryNonConstExprXinterfaceXShlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type complex128, must be unsigned integer)`,
	)

}

// Test XinterfaceX << String
func TestCheckBinaryNonConstExprXinterfaceXShlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type string, must be unsigned integer)`,
	)

}

// Test XinterfaceX << BoolT
func TestCheckBinaryNonConstExprXinterfaceXShlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type bool, must be unsigned integer)`,
	)

}

// Test XinterfaceX << Slice
func TestCheckBinaryNonConstExprXinterfaceXShlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.sliceT, must be unsigned integer)`,
	)

}

// Test XinterfaceX << Array
func TestCheckBinaryNonConstExprXinterfaceXShlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.arrayT, must be unsigned integer)`,
	)

}

// Test XinterfaceX << Map
func TestCheckBinaryNonConstExprXinterfaceXShlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.mapT, must be unsigned integer)`,
	)

}

// Test XinterfaceX << XinterfaceX
func TestCheckBinaryNonConstExprXinterfaceXShlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.XinterfaceX, must be unsigned integer)`,
	)

}

// Test XinterfaceX << InterfaceX
func TestCheckBinaryNonConstExprXinterfaceXShlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceX, must be unsigned integer)`,
	)

}

// Test XinterfaceX << InterfaceY
func TestCheckBinaryNonConstExprXinterfaceXShlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceY, must be unsigned integer)`,
	)

}

// Test XinterfaceX << InterfaceZ
func TestCheckBinaryNonConstExprXinterfaceXShlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceZ, must be unsigned integer)`,
	)

}

// Test XinterfaceX << Ptr
func TestCheckBinaryNonConstExprXinterfaceXShlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type *int, must be unsigned integer)`,
	)

}

// Test XinterfaceX << Struct
func TestCheckBinaryNonConstExprXinterfaceXShlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structT, must be unsigned integer)`,
	)

}

// Test XinterfaceX << StructUncomp
func TestCheckBinaryNonConstExprXinterfaceXShlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := XinterfaceX(1); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structUncompT, must be unsigned integer)`,
	)

}

// Test InterfaceX + ConstInt
func TestCheckBinaryNonConstExprInterfaceXAddConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 4`, env,
		`invalid operation: a + 4 (mismatched types eval.interfaceX and int)`,
	)

}

// Test InterfaceX + ConstRune
func TestCheckBinaryNonConstExprInterfaceXAddConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + '@'`, env,
		`invalid operation: a + rune(64) (mismatched types eval.interfaceX and rune)`,
	)

}

// Test InterfaceX + ConstFloat
func TestCheckBinaryNonConstExprInterfaceXAddConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 2.0`, env,
		`invalid operation: a + 2 (mismatched types eval.interfaceX and float64)`,
	)

}

// Test InterfaceX + ConstComplex
func TestCheckBinaryNonConstExprInterfaceXAddConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 8.0i`, env,
		`invalid operation: a + 8i (mismatched types eval.interfaceX and complex128)`,
	)

}

// Test InterfaceX + ConstBool
func TestCheckBinaryNonConstExprInterfaceXAddConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + true`, env,
		`invalid operation: a + true (mismatched types eval.interfaceX and bool)`,
	)

}

// Test InterfaceX + ConstString
func TestCheckBinaryNonConstExprInterfaceXAddConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + "abc"`, env,
		`invalid operation: a + "abc" (mismatched types eval.interfaceX and string)`,
	)

}

// Test InterfaceX + ConstNil
func TestCheckBinaryNonConstExprInterfaceXAddConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + nil`, env,
		`invalid operation: a + nil (operator + not defined on interface)`,
	)

}

// Test InterfaceX + Int
func TestCheckBinaryNonConstExprInterfaceXAddInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceX and int)`,
	)

}

// Test InterfaceX + Float32
func TestCheckBinaryNonConstExprInterfaceXAddFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceX and float32)`,
	)

}

// Test InterfaceX + Complex128
func TestCheckBinaryNonConstExprInterfaceXAddComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceX and complex128)`,
	)

}

// Test InterfaceX + String
func TestCheckBinaryNonConstExprInterfaceXAddString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceX and string)`,
	)

}

// Test InterfaceX + BoolT
func TestCheckBinaryNonConstExprInterfaceXAddBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceX and bool)`,
	)

}

// Test InterfaceX + Slice
func TestCheckBinaryNonConstExprInterfaceXAddSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceX and eval.sliceT)`,
	)

}

// Test InterfaceX + Array
func TestCheckBinaryNonConstExprInterfaceXAddArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceX and eval.arrayT)`,
	)

}

// Test InterfaceX + Map
func TestCheckBinaryNonConstExprInterfaceXAddMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceX and eval.mapT)`,
	)

}

// Test InterfaceX + XinterfaceX
func TestCheckBinaryNonConstExprInterfaceXAddXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceX and eval.XinterfaceX)`,
	)

}

// Test InterfaceX + InterfaceX
func TestCheckBinaryNonConstExprInterfaceXAddInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (operator + not defined on interface)`,
	)

}

// Test InterfaceX + InterfaceY
func TestCheckBinaryNonConstExprInterfaceXAddInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceX and eval.interfaceY)`,
	)

}

// Test InterfaceX + InterfaceZ
func TestCheckBinaryNonConstExprInterfaceXAddInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceX and eval.interfaceZ)`,
	)

}

// Test InterfaceX + Ptr
func TestCheckBinaryNonConstExprInterfaceXAddPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceX and *int)`,
	)

}

// Test InterfaceX + Struct
func TestCheckBinaryNonConstExprInterfaceXAddStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceX and eval.structT)`,
	)

}

// Test InterfaceX + StructUncomp
func TestCheckBinaryNonConstExprInterfaceXAddStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceX and eval.structUncompT)`,
	)

}

// Test InterfaceX & ConstInt
func TestCheckBinaryNonConstExprInterfaceXAndConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 4`, env,
		`invalid operation: a & 4 (mismatched types eval.interfaceX and int)`,
	)

}

// Test InterfaceX & ConstRune
func TestCheckBinaryNonConstExprInterfaceXAndConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & '@'`, env,
		`invalid operation: a & rune(64) (mismatched types eval.interfaceX and rune)`,
	)

}

// Test InterfaceX & ConstFloat
func TestCheckBinaryNonConstExprInterfaceXAndConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 2.0`, env,
		`invalid operation: a & 2 (mismatched types eval.interfaceX and float64)`,
	)

}

// Test InterfaceX & ConstComplex
func TestCheckBinaryNonConstExprInterfaceXAndConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 8.0i`, env,
		`invalid operation: a & 8i (mismatched types eval.interfaceX and complex128)`,
	)

}

// Test InterfaceX & ConstBool
func TestCheckBinaryNonConstExprInterfaceXAndConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & true`, env,
		`invalid operation: a & true (mismatched types eval.interfaceX and bool)`,
	)

}

// Test InterfaceX & ConstString
func TestCheckBinaryNonConstExprInterfaceXAndConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & "abc"`, env,
		`invalid operation: a & "abc" (mismatched types eval.interfaceX and string)`,
	)

}

// Test InterfaceX & ConstNil
func TestCheckBinaryNonConstExprInterfaceXAndConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & nil`, env,
		`invalid operation: a & nil (operator & not defined on interface)`,
	)

}

// Test InterfaceX & Int
func TestCheckBinaryNonConstExprInterfaceXAndInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceX and int)`,
	)

}

// Test InterfaceX & Float32
func TestCheckBinaryNonConstExprInterfaceXAndFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceX and float32)`,
	)

}

// Test InterfaceX & Complex128
func TestCheckBinaryNonConstExprInterfaceXAndComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceX and complex128)`,
	)

}

// Test InterfaceX & String
func TestCheckBinaryNonConstExprInterfaceXAndString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceX and string)`,
	)

}

// Test InterfaceX & BoolT
func TestCheckBinaryNonConstExprInterfaceXAndBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceX and bool)`,
	)

}

// Test InterfaceX & Slice
func TestCheckBinaryNonConstExprInterfaceXAndSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceX and eval.sliceT)`,
	)

}

// Test InterfaceX & Array
func TestCheckBinaryNonConstExprInterfaceXAndArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceX and eval.arrayT)`,
	)

}

// Test InterfaceX & Map
func TestCheckBinaryNonConstExprInterfaceXAndMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceX and eval.mapT)`,
	)

}

// Test InterfaceX & XinterfaceX
func TestCheckBinaryNonConstExprInterfaceXAndXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceX and eval.XinterfaceX)`,
	)

}

// Test InterfaceX & InterfaceX
func TestCheckBinaryNonConstExprInterfaceXAndInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (operator & not defined on interface)`,
	)

}

// Test InterfaceX & InterfaceY
func TestCheckBinaryNonConstExprInterfaceXAndInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceX and eval.interfaceY)`,
	)

}

// Test InterfaceX & InterfaceZ
func TestCheckBinaryNonConstExprInterfaceXAndInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceX and eval.interfaceZ)`,
	)

}

// Test InterfaceX & Ptr
func TestCheckBinaryNonConstExprInterfaceXAndPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceX and *int)`,
	)

}

// Test InterfaceX & Struct
func TestCheckBinaryNonConstExprInterfaceXAndStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceX and eval.structT)`,
	)

}

// Test InterfaceX & StructUncomp
func TestCheckBinaryNonConstExprInterfaceXAndStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceX and eval.structUncompT)`,
	)

}

// Test InterfaceX % ConstInt
func TestCheckBinaryNonConstExprInterfaceXRemConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 4`, env,
		`invalid operation: a % 4 (mismatched types eval.interfaceX and int)`,
	)

}

// Test InterfaceX % ConstRune
func TestCheckBinaryNonConstExprInterfaceXRemConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % '@'`, env,
		`invalid operation: a % rune(64) (mismatched types eval.interfaceX and rune)`,
	)

}

// Test InterfaceX % ConstFloat
func TestCheckBinaryNonConstExprInterfaceXRemConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 2.0`, env,
		`invalid operation: a % 2 (mismatched types eval.interfaceX and float64)`,
	)

}

// Test InterfaceX % ConstComplex
func TestCheckBinaryNonConstExprInterfaceXRemConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 8.0i`, env,
		`invalid operation: a % 8i (mismatched types eval.interfaceX and complex128)`,
	)

}

// Test InterfaceX % ConstBool
func TestCheckBinaryNonConstExprInterfaceXRemConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % true`, env,
		`invalid operation: a % true (mismatched types eval.interfaceX and bool)`,
	)

}

// Test InterfaceX % ConstString
func TestCheckBinaryNonConstExprInterfaceXRemConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % "abc"`, env,
		`invalid operation: a % "abc" (mismatched types eval.interfaceX and string)`,
	)

}

// Test InterfaceX % ConstNil
func TestCheckBinaryNonConstExprInterfaceXRemConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % nil`, env,
		`invalid operation: a % nil (operator % not defined on interface)`,
	)

}

// Test InterfaceX % Int
func TestCheckBinaryNonConstExprInterfaceXRemInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceX and int)`,
	)

}

// Test InterfaceX % Float32
func TestCheckBinaryNonConstExprInterfaceXRemFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceX and float32)`,
	)

}

// Test InterfaceX % Complex128
func TestCheckBinaryNonConstExprInterfaceXRemComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceX and complex128)`,
	)

}

// Test InterfaceX % String
func TestCheckBinaryNonConstExprInterfaceXRemString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceX and string)`,
	)

}

// Test InterfaceX % BoolT
func TestCheckBinaryNonConstExprInterfaceXRemBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceX and bool)`,
	)

}

// Test InterfaceX % Slice
func TestCheckBinaryNonConstExprInterfaceXRemSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceX and eval.sliceT)`,
	)

}

// Test InterfaceX % Array
func TestCheckBinaryNonConstExprInterfaceXRemArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceX and eval.arrayT)`,
	)

}

// Test InterfaceX % Map
func TestCheckBinaryNonConstExprInterfaceXRemMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceX and eval.mapT)`,
	)

}

// Test InterfaceX % XinterfaceX
func TestCheckBinaryNonConstExprInterfaceXRemXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceX and eval.XinterfaceX)`,
	)

}

// Test InterfaceX % InterfaceX
func TestCheckBinaryNonConstExprInterfaceXRemInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (operator % not defined on interface)`,
	)

}

// Test InterfaceX % InterfaceY
func TestCheckBinaryNonConstExprInterfaceXRemInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceX and eval.interfaceY)`,
	)

}

// Test InterfaceX % InterfaceZ
func TestCheckBinaryNonConstExprInterfaceXRemInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceX and eval.interfaceZ)`,
	)

}

// Test InterfaceX % Ptr
func TestCheckBinaryNonConstExprInterfaceXRemPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceX and *int)`,
	)

}

// Test InterfaceX % Struct
func TestCheckBinaryNonConstExprInterfaceXRemStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceX and eval.structT)`,
	)

}

// Test InterfaceX % StructUncomp
func TestCheckBinaryNonConstExprInterfaceXRemStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceX and eval.structUncompT)`,
	)

}

// Test InterfaceX == ConstInt
func TestCheckBinaryNonConstExprInterfaceXEqlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 4`, env,
		`invalid operation: a == 4 (mismatched types eval.interfaceX and int)`,
	)

}

// Test InterfaceX == ConstRune
func TestCheckBinaryNonConstExprInterfaceXEqlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == '@'`, env,
		`invalid operation: a == rune(64) (mismatched types eval.interfaceX and rune)`,
	)

}

// Test InterfaceX == ConstFloat
func TestCheckBinaryNonConstExprInterfaceXEqlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 2.0`, env,
		`invalid operation: a == 2 (mismatched types eval.interfaceX and float64)`,
	)

}

// Test InterfaceX == ConstComplex
func TestCheckBinaryNonConstExprInterfaceXEqlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 8.0i`, env,
		`invalid operation: a == 8i (mismatched types eval.interfaceX and complex128)`,
	)

}

// Test InterfaceX == ConstBool
func TestCheckBinaryNonConstExprInterfaceXEqlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == true`, env,
		`invalid operation: a == true (mismatched types eval.interfaceX and bool)`,
	)

}

// Test InterfaceX == ConstString
func TestCheckBinaryNonConstExprInterfaceXEqlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == "abc"`, env,
		`invalid operation: a == "abc" (mismatched types eval.interfaceX and string)`,
	)

}

// Test InterfaceX == ConstNil
func TestCheckBinaryNonConstExprInterfaceXEqlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == nil`, env, reflect.TypeOf(a == nil))
}

// Test InterfaceX == Int
func TestCheckBinaryNonConstExprInterfaceXEqlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceX and int)`,
	)

}

// Test InterfaceX == Float32
func TestCheckBinaryNonConstExprInterfaceXEqlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceX and float32)`,
	)

}

// Test InterfaceX == Complex128
func TestCheckBinaryNonConstExprInterfaceXEqlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceX and complex128)`,
	)

}

// Test InterfaceX == String
func TestCheckBinaryNonConstExprInterfaceXEqlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceX and string)`,
	)

}

// Test InterfaceX == BoolT
func TestCheckBinaryNonConstExprInterfaceXEqlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceX and bool)`,
	)

}

// Test InterfaceX == Slice
func TestCheckBinaryNonConstExprInterfaceXEqlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceX and eval.sliceT)`,
	)

}

// Test InterfaceX == Array
func TestCheckBinaryNonConstExprInterfaceXEqlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceX and eval.arrayT)`,
	)

}

// Test InterfaceX == Map
func TestCheckBinaryNonConstExprInterfaceXEqlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceX and eval.mapT)`,
	)

}

// Test InterfaceX == XinterfaceX
func TestCheckBinaryNonConstExprInterfaceXEqlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test InterfaceX == InterfaceX
func TestCheckBinaryNonConstExprInterfaceXEqlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test InterfaceX == InterfaceY
func TestCheckBinaryNonConstExprInterfaceXEqlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test InterfaceX == InterfaceZ
func TestCheckBinaryNonConstExprInterfaceXEqlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceX and eval.interfaceZ)`,
	)

}

// Test InterfaceX == Ptr
func TestCheckBinaryNonConstExprInterfaceXEqlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceX and *int)`,
	)

}

// Test InterfaceX == Struct
func TestCheckBinaryNonConstExprInterfaceXEqlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceX and eval.structT)`,
	)

}

// Test InterfaceX == StructUncomp
func TestCheckBinaryNonConstExprInterfaceXEqlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceX and eval.structUncompT)`,
	)

}

// Test InterfaceX > ConstInt
func TestCheckBinaryNonConstExprInterfaceXGtrConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 4`, env,
		`invalid operation: a > 4 (mismatched types eval.interfaceX and int)`,
	)

}

// Test InterfaceX > ConstRune
func TestCheckBinaryNonConstExprInterfaceXGtrConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > '@'`, env,
		`invalid operation: a > rune(64) (mismatched types eval.interfaceX and rune)`,
	)

}

// Test InterfaceX > ConstFloat
func TestCheckBinaryNonConstExprInterfaceXGtrConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 2.0`, env,
		`invalid operation: a > 2 (mismatched types eval.interfaceX and float64)`,
	)

}

// Test InterfaceX > ConstComplex
func TestCheckBinaryNonConstExprInterfaceXGtrConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 8.0i`, env,
		`invalid operation: a > 8i (mismatched types eval.interfaceX and complex128)`,
	)

}

// Test InterfaceX > ConstBool
func TestCheckBinaryNonConstExprInterfaceXGtrConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > true`, env,
		`invalid operation: a > true (mismatched types eval.interfaceX and bool)`,
	)

}

// Test InterfaceX > ConstString
func TestCheckBinaryNonConstExprInterfaceXGtrConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > "abc"`, env,
		`invalid operation: a > "abc" (mismatched types eval.interfaceX and string)`,
	)

}

// Test InterfaceX > ConstNil
func TestCheckBinaryNonConstExprInterfaceXGtrConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > nil`, env,
		`invalid operation: a > nil (operator > not defined on interface)`,
	)

}

// Test InterfaceX > Int
func TestCheckBinaryNonConstExprInterfaceXGtrInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceX and int)`,
	)

}

// Test InterfaceX > Float32
func TestCheckBinaryNonConstExprInterfaceXGtrFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceX and float32)`,
	)

}

// Test InterfaceX > Complex128
func TestCheckBinaryNonConstExprInterfaceXGtrComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceX and complex128)`,
	)

}

// Test InterfaceX > String
func TestCheckBinaryNonConstExprInterfaceXGtrString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceX and string)`,
	)

}

// Test InterfaceX > BoolT
func TestCheckBinaryNonConstExprInterfaceXGtrBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceX and bool)`,
	)

}

// Test InterfaceX > Slice
func TestCheckBinaryNonConstExprInterfaceXGtrSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceX and eval.sliceT)`,
	)

}

// Test InterfaceX > Array
func TestCheckBinaryNonConstExprInterfaceXGtrArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceX and eval.arrayT)`,
	)

}

// Test InterfaceX > Map
func TestCheckBinaryNonConstExprInterfaceXGtrMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceX and eval.mapT)`,
	)

}

// Test InterfaceX > XinterfaceX
func TestCheckBinaryNonConstExprInterfaceXGtrXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > eval.interfaceX(b) (operator > not defined on interface)`,
	)

}

// Test InterfaceX > InterfaceX
func TestCheckBinaryNonConstExprInterfaceXGtrInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (operator > not defined on interface)`,
	)

}

// Test InterfaceX > InterfaceY
func TestCheckBinaryNonConstExprInterfaceXGtrInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: eval.interfaceY(a) > b (operator > not defined on interface)`,
	)

}

// Test InterfaceX > InterfaceZ
func TestCheckBinaryNonConstExprInterfaceXGtrInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceX and eval.interfaceZ)`,
	)

}

// Test InterfaceX > Ptr
func TestCheckBinaryNonConstExprInterfaceXGtrPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceX and *int)`,
	)

}

// Test InterfaceX > Struct
func TestCheckBinaryNonConstExprInterfaceXGtrStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceX and eval.structT)`,
	)

}

// Test InterfaceX > StructUncomp
func TestCheckBinaryNonConstExprInterfaceXGtrStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceX and eval.structUncompT)`,
	)

}

// Test InterfaceX << ConstInt
func TestCheckBinaryNonConstExprInterfaceXShlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 4`, env,
		`invalid operation: a << 4 (shift of type eval.interfaceX)`,
	)

}

// Test InterfaceX << ConstRune
func TestCheckBinaryNonConstExprInterfaceXShlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << '@'`, env,
		`invalid operation: a << 64 (shift of type eval.interfaceX)`,
	)

}

// Test InterfaceX << ConstFloat
func TestCheckBinaryNonConstExprInterfaceXShlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 2.0`, env,
		`invalid operation: a << 2 (shift of type eval.interfaceX)`,
	)

}

// Test InterfaceX << ConstComplex
func TestCheckBinaryNonConstExprInterfaceXShlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: a << 0 (shift of type eval.interfaceX)`,
	)

}

// Test InterfaceX << ConstBool
func TestCheckBinaryNonConstExprInterfaceXShlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << true`, env,
		`invalid operation: a << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test InterfaceX << ConstString
func TestCheckBinaryNonConstExprInterfaceXShlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: a << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test InterfaceX << ConstNil
func TestCheckBinaryNonConstExprInterfaceXShlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test InterfaceX << Int
func TestCheckBinaryNonConstExprInterfaceXShlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type int, must be unsigned integer)`,
	)

}

// Test InterfaceX << Float32
func TestCheckBinaryNonConstExprInterfaceXShlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type float32, must be unsigned integer)`,
	)

}

// Test InterfaceX << Complex128
func TestCheckBinaryNonConstExprInterfaceXShlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type complex128, must be unsigned integer)`,
	)

}

// Test InterfaceX << String
func TestCheckBinaryNonConstExprInterfaceXShlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type string, must be unsigned integer)`,
	)

}

// Test InterfaceX << BoolT
func TestCheckBinaryNonConstExprInterfaceXShlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type bool, must be unsigned integer)`,
	)

}

// Test InterfaceX << Slice
func TestCheckBinaryNonConstExprInterfaceXShlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.sliceT, must be unsigned integer)`,
	)

}

// Test InterfaceX << Array
func TestCheckBinaryNonConstExprInterfaceXShlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.arrayT, must be unsigned integer)`,
	)

}

// Test InterfaceX << Map
func TestCheckBinaryNonConstExprInterfaceXShlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.mapT, must be unsigned integer)`,
	)

}

// Test InterfaceX << XinterfaceX
func TestCheckBinaryNonConstExprInterfaceXShlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.XinterfaceX, must be unsigned integer)`,
	)

}

// Test InterfaceX << InterfaceX
func TestCheckBinaryNonConstExprInterfaceXShlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceX, must be unsigned integer)`,
	)

}

// Test InterfaceX << InterfaceY
func TestCheckBinaryNonConstExprInterfaceXShlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceY, must be unsigned integer)`,
	)

}

// Test InterfaceX << InterfaceZ
func TestCheckBinaryNonConstExprInterfaceXShlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceZ, must be unsigned integer)`,
	)

}

// Test InterfaceX << Ptr
func TestCheckBinaryNonConstExprInterfaceXShlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type *int, must be unsigned integer)`,
	)

}

// Test InterfaceX << Struct
func TestCheckBinaryNonConstExprInterfaceXShlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structT, must be unsigned integer)`,
	)

}

// Test InterfaceX << StructUncomp
func TestCheckBinaryNonConstExprInterfaceXShlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceX(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structUncompT, must be unsigned integer)`,
	)

}

// Test InterfaceY + ConstInt
func TestCheckBinaryNonConstExprInterfaceYAddConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 4`, env,
		`invalid operation: a + 4 (mismatched types eval.interfaceY and int)`,
	)

}

// Test InterfaceY + ConstRune
func TestCheckBinaryNonConstExprInterfaceYAddConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + '@'`, env,
		`invalid operation: a + rune(64) (mismatched types eval.interfaceY and rune)`,
	)

}

// Test InterfaceY + ConstFloat
func TestCheckBinaryNonConstExprInterfaceYAddConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 2.0`, env,
		`invalid operation: a + 2 (mismatched types eval.interfaceY and float64)`,
	)

}

// Test InterfaceY + ConstComplex
func TestCheckBinaryNonConstExprInterfaceYAddConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 8.0i`, env,
		`invalid operation: a + 8i (mismatched types eval.interfaceY and complex128)`,
	)

}

// Test InterfaceY + ConstBool
func TestCheckBinaryNonConstExprInterfaceYAddConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + true`, env,
		`invalid operation: a + true (mismatched types eval.interfaceY and bool)`,
	)

}

// Test InterfaceY + ConstString
func TestCheckBinaryNonConstExprInterfaceYAddConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + "abc"`, env,
		`invalid operation: a + "abc" (mismatched types eval.interfaceY and string)`,
	)

}

// Test InterfaceY + ConstNil
func TestCheckBinaryNonConstExprInterfaceYAddConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + nil`, env,
		`invalid operation: a + nil (operator + not defined on interface)`,
	)

}

// Test InterfaceY + Int
func TestCheckBinaryNonConstExprInterfaceYAddInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceY and int)`,
	)

}

// Test InterfaceY + Float32
func TestCheckBinaryNonConstExprInterfaceYAddFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceY and float32)`,
	)

}

// Test InterfaceY + Complex128
func TestCheckBinaryNonConstExprInterfaceYAddComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceY and complex128)`,
	)

}

// Test InterfaceY + String
func TestCheckBinaryNonConstExprInterfaceYAddString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceY and string)`,
	)

}

// Test InterfaceY + BoolT
func TestCheckBinaryNonConstExprInterfaceYAddBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceY and bool)`,
	)

}

// Test InterfaceY + Slice
func TestCheckBinaryNonConstExprInterfaceYAddSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceY and eval.sliceT)`,
	)

}

// Test InterfaceY + Array
func TestCheckBinaryNonConstExprInterfaceYAddArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceY and eval.arrayT)`,
	)

}

// Test InterfaceY + Map
func TestCheckBinaryNonConstExprInterfaceYAddMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceY and eval.mapT)`,
	)

}

// Test InterfaceY + XinterfaceX
func TestCheckBinaryNonConstExprInterfaceYAddXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceY and eval.XinterfaceX)`,
	)

}

// Test InterfaceY + InterfaceX
func TestCheckBinaryNonConstExprInterfaceYAddInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceY and eval.interfaceX)`,
	)

}

// Test InterfaceY + InterfaceY
func TestCheckBinaryNonConstExprInterfaceYAddInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (operator + not defined on interface)`,
	)

}

// Test InterfaceY + InterfaceZ
func TestCheckBinaryNonConstExprInterfaceYAddInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceY and eval.interfaceZ)`,
	)

}

// Test InterfaceY + Ptr
func TestCheckBinaryNonConstExprInterfaceYAddPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceY and *int)`,
	)

}

// Test InterfaceY + Struct
func TestCheckBinaryNonConstExprInterfaceYAddStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceY and eval.structT)`,
	)

}

// Test InterfaceY + StructUncomp
func TestCheckBinaryNonConstExprInterfaceYAddStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceY and eval.structUncompT)`,
	)

}

// Test InterfaceY & ConstInt
func TestCheckBinaryNonConstExprInterfaceYAndConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 4`, env,
		`invalid operation: a & 4 (mismatched types eval.interfaceY and int)`,
	)

}

// Test InterfaceY & ConstRune
func TestCheckBinaryNonConstExprInterfaceYAndConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & '@'`, env,
		`invalid operation: a & rune(64) (mismatched types eval.interfaceY and rune)`,
	)

}

// Test InterfaceY & ConstFloat
func TestCheckBinaryNonConstExprInterfaceYAndConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 2.0`, env,
		`invalid operation: a & 2 (mismatched types eval.interfaceY and float64)`,
	)

}

// Test InterfaceY & ConstComplex
func TestCheckBinaryNonConstExprInterfaceYAndConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 8.0i`, env,
		`invalid operation: a & 8i (mismatched types eval.interfaceY and complex128)`,
	)

}

// Test InterfaceY & ConstBool
func TestCheckBinaryNonConstExprInterfaceYAndConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & true`, env,
		`invalid operation: a & true (mismatched types eval.interfaceY and bool)`,
	)

}

// Test InterfaceY & ConstString
func TestCheckBinaryNonConstExprInterfaceYAndConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & "abc"`, env,
		`invalid operation: a & "abc" (mismatched types eval.interfaceY and string)`,
	)

}

// Test InterfaceY & ConstNil
func TestCheckBinaryNonConstExprInterfaceYAndConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & nil`, env,
		`invalid operation: a & nil (operator & not defined on interface)`,
	)

}

// Test InterfaceY & Int
func TestCheckBinaryNonConstExprInterfaceYAndInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceY and int)`,
	)

}

// Test InterfaceY & Float32
func TestCheckBinaryNonConstExprInterfaceYAndFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceY and float32)`,
	)

}

// Test InterfaceY & Complex128
func TestCheckBinaryNonConstExprInterfaceYAndComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceY and complex128)`,
	)

}

// Test InterfaceY & String
func TestCheckBinaryNonConstExprInterfaceYAndString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceY and string)`,
	)

}

// Test InterfaceY & BoolT
func TestCheckBinaryNonConstExprInterfaceYAndBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceY and bool)`,
	)

}

// Test InterfaceY & Slice
func TestCheckBinaryNonConstExprInterfaceYAndSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceY and eval.sliceT)`,
	)

}

// Test InterfaceY & Array
func TestCheckBinaryNonConstExprInterfaceYAndArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceY and eval.arrayT)`,
	)

}

// Test InterfaceY & Map
func TestCheckBinaryNonConstExprInterfaceYAndMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceY and eval.mapT)`,
	)

}

// Test InterfaceY & XinterfaceX
func TestCheckBinaryNonConstExprInterfaceYAndXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceY and eval.XinterfaceX)`,
	)

}

// Test InterfaceY & InterfaceX
func TestCheckBinaryNonConstExprInterfaceYAndInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceY and eval.interfaceX)`,
	)

}

// Test InterfaceY & InterfaceY
func TestCheckBinaryNonConstExprInterfaceYAndInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (operator & not defined on interface)`,
	)

}

// Test InterfaceY & InterfaceZ
func TestCheckBinaryNonConstExprInterfaceYAndInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceY and eval.interfaceZ)`,
	)

}

// Test InterfaceY & Ptr
func TestCheckBinaryNonConstExprInterfaceYAndPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceY and *int)`,
	)

}

// Test InterfaceY & Struct
func TestCheckBinaryNonConstExprInterfaceYAndStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceY and eval.structT)`,
	)

}

// Test InterfaceY & StructUncomp
func TestCheckBinaryNonConstExprInterfaceYAndStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceY and eval.structUncompT)`,
	)

}

// Test InterfaceY % ConstInt
func TestCheckBinaryNonConstExprInterfaceYRemConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 4`, env,
		`invalid operation: a % 4 (mismatched types eval.interfaceY and int)`,
	)

}

// Test InterfaceY % ConstRune
func TestCheckBinaryNonConstExprInterfaceYRemConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % '@'`, env,
		`invalid operation: a % rune(64) (mismatched types eval.interfaceY and rune)`,
	)

}

// Test InterfaceY % ConstFloat
func TestCheckBinaryNonConstExprInterfaceYRemConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 2.0`, env,
		`invalid operation: a % 2 (mismatched types eval.interfaceY and float64)`,
	)

}

// Test InterfaceY % ConstComplex
func TestCheckBinaryNonConstExprInterfaceYRemConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 8.0i`, env,
		`invalid operation: a % 8i (mismatched types eval.interfaceY and complex128)`,
	)

}

// Test InterfaceY % ConstBool
func TestCheckBinaryNonConstExprInterfaceYRemConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % true`, env,
		`invalid operation: a % true (mismatched types eval.interfaceY and bool)`,
	)

}

// Test InterfaceY % ConstString
func TestCheckBinaryNonConstExprInterfaceYRemConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % "abc"`, env,
		`invalid operation: a % "abc" (mismatched types eval.interfaceY and string)`,
	)

}

// Test InterfaceY % ConstNil
func TestCheckBinaryNonConstExprInterfaceYRemConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % nil`, env,
		`invalid operation: a % nil (operator % not defined on interface)`,
	)

}

// Test InterfaceY % Int
func TestCheckBinaryNonConstExprInterfaceYRemInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceY and int)`,
	)

}

// Test InterfaceY % Float32
func TestCheckBinaryNonConstExprInterfaceYRemFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceY and float32)`,
	)

}

// Test InterfaceY % Complex128
func TestCheckBinaryNonConstExprInterfaceYRemComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceY and complex128)`,
	)

}

// Test InterfaceY % String
func TestCheckBinaryNonConstExprInterfaceYRemString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceY and string)`,
	)

}

// Test InterfaceY % BoolT
func TestCheckBinaryNonConstExprInterfaceYRemBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceY and bool)`,
	)

}

// Test InterfaceY % Slice
func TestCheckBinaryNonConstExprInterfaceYRemSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceY and eval.sliceT)`,
	)

}

// Test InterfaceY % Array
func TestCheckBinaryNonConstExprInterfaceYRemArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceY and eval.arrayT)`,
	)

}

// Test InterfaceY % Map
func TestCheckBinaryNonConstExprInterfaceYRemMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceY and eval.mapT)`,
	)

}

// Test InterfaceY % XinterfaceX
func TestCheckBinaryNonConstExprInterfaceYRemXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceY and eval.XinterfaceX)`,
	)

}

// Test InterfaceY % InterfaceX
func TestCheckBinaryNonConstExprInterfaceYRemInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceY and eval.interfaceX)`,
	)

}

// Test InterfaceY % InterfaceY
func TestCheckBinaryNonConstExprInterfaceYRemInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (operator % not defined on interface)`,
	)

}

// Test InterfaceY % InterfaceZ
func TestCheckBinaryNonConstExprInterfaceYRemInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceY and eval.interfaceZ)`,
	)

}

// Test InterfaceY % Ptr
func TestCheckBinaryNonConstExprInterfaceYRemPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceY and *int)`,
	)

}

// Test InterfaceY % Struct
func TestCheckBinaryNonConstExprInterfaceYRemStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceY and eval.structT)`,
	)

}

// Test InterfaceY % StructUncomp
func TestCheckBinaryNonConstExprInterfaceYRemStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceY and eval.structUncompT)`,
	)

}

// Test InterfaceY == ConstInt
func TestCheckBinaryNonConstExprInterfaceYEqlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 4`, env,
		`invalid operation: a == 4 (mismatched types eval.interfaceY and int)`,
	)

}

// Test InterfaceY == ConstRune
func TestCheckBinaryNonConstExprInterfaceYEqlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == '@'`, env,
		`invalid operation: a == rune(64) (mismatched types eval.interfaceY and rune)`,
	)

}

// Test InterfaceY == ConstFloat
func TestCheckBinaryNonConstExprInterfaceYEqlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 2.0`, env,
		`invalid operation: a == 2 (mismatched types eval.interfaceY and float64)`,
	)

}

// Test InterfaceY == ConstComplex
func TestCheckBinaryNonConstExprInterfaceYEqlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 8.0i`, env,
		`invalid operation: a == 8i (mismatched types eval.interfaceY and complex128)`,
	)

}

// Test InterfaceY == ConstBool
func TestCheckBinaryNonConstExprInterfaceYEqlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == true`, env,
		`invalid operation: a == true (mismatched types eval.interfaceY and bool)`,
	)

}

// Test InterfaceY == ConstString
func TestCheckBinaryNonConstExprInterfaceYEqlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == "abc"`, env,
		`invalid operation: a == "abc" (mismatched types eval.interfaceY and string)`,
	)

}

// Test InterfaceY == ConstNil
func TestCheckBinaryNonConstExprInterfaceYEqlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == nil`, env, reflect.TypeOf(a == nil))
}

// Test InterfaceY == Int
func TestCheckBinaryNonConstExprInterfaceYEqlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceY and int)`,
	)

}

// Test InterfaceY == Float32
func TestCheckBinaryNonConstExprInterfaceYEqlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceY and float32)`,
	)

}

// Test InterfaceY == Complex128
func TestCheckBinaryNonConstExprInterfaceYEqlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceY and complex128)`,
	)

}

// Test InterfaceY == String
func TestCheckBinaryNonConstExprInterfaceYEqlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceY and string)`,
	)

}

// Test InterfaceY == BoolT
func TestCheckBinaryNonConstExprInterfaceYEqlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceY and bool)`,
	)

}

// Test InterfaceY == Slice
func TestCheckBinaryNonConstExprInterfaceYEqlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceY and eval.sliceT)`,
	)

}

// Test InterfaceY == Array
func TestCheckBinaryNonConstExprInterfaceYEqlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceY and eval.arrayT)`,
	)

}

// Test InterfaceY == Map
func TestCheckBinaryNonConstExprInterfaceYEqlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceY and eval.mapT)`,
	)

}

// Test InterfaceY == XinterfaceX
func TestCheckBinaryNonConstExprInterfaceYEqlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test InterfaceY == InterfaceX
func TestCheckBinaryNonConstExprInterfaceYEqlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test InterfaceY == InterfaceY
func TestCheckBinaryNonConstExprInterfaceYEqlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test InterfaceY == InterfaceZ
func TestCheckBinaryNonConstExprInterfaceYEqlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceY and eval.interfaceZ)`,
	)

}

// Test InterfaceY == Ptr
func TestCheckBinaryNonConstExprInterfaceYEqlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceY and *int)`,
	)

}

// Test InterfaceY == Struct
func TestCheckBinaryNonConstExprInterfaceYEqlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceY and eval.structT)`,
	)

}

// Test InterfaceY == StructUncomp
func TestCheckBinaryNonConstExprInterfaceYEqlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceY and eval.structUncompT)`,
	)

}

// Test InterfaceY > ConstInt
func TestCheckBinaryNonConstExprInterfaceYGtrConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 4`, env,
		`invalid operation: a > 4 (mismatched types eval.interfaceY and int)`,
	)

}

// Test InterfaceY > ConstRune
func TestCheckBinaryNonConstExprInterfaceYGtrConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > '@'`, env,
		`invalid operation: a > rune(64) (mismatched types eval.interfaceY and rune)`,
	)

}

// Test InterfaceY > ConstFloat
func TestCheckBinaryNonConstExprInterfaceYGtrConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 2.0`, env,
		`invalid operation: a > 2 (mismatched types eval.interfaceY and float64)`,
	)

}

// Test InterfaceY > ConstComplex
func TestCheckBinaryNonConstExprInterfaceYGtrConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 8.0i`, env,
		`invalid operation: a > 8i (mismatched types eval.interfaceY and complex128)`,
	)

}

// Test InterfaceY > ConstBool
func TestCheckBinaryNonConstExprInterfaceYGtrConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > true`, env,
		`invalid operation: a > true (mismatched types eval.interfaceY and bool)`,
	)

}

// Test InterfaceY > ConstString
func TestCheckBinaryNonConstExprInterfaceYGtrConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > "abc"`, env,
		`invalid operation: a > "abc" (mismatched types eval.interfaceY and string)`,
	)

}

// Test InterfaceY > ConstNil
func TestCheckBinaryNonConstExprInterfaceYGtrConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > nil`, env,
		`invalid operation: a > nil (operator > not defined on interface)`,
	)

}

// Test InterfaceY > Int
func TestCheckBinaryNonConstExprInterfaceYGtrInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceY and int)`,
	)

}

// Test InterfaceY > Float32
func TestCheckBinaryNonConstExprInterfaceYGtrFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceY and float32)`,
	)

}

// Test InterfaceY > Complex128
func TestCheckBinaryNonConstExprInterfaceYGtrComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceY and complex128)`,
	)

}

// Test InterfaceY > String
func TestCheckBinaryNonConstExprInterfaceYGtrString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceY and string)`,
	)

}

// Test InterfaceY > BoolT
func TestCheckBinaryNonConstExprInterfaceYGtrBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceY and bool)`,
	)

}

// Test InterfaceY > Slice
func TestCheckBinaryNonConstExprInterfaceYGtrSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceY and eval.sliceT)`,
	)

}

// Test InterfaceY > Array
func TestCheckBinaryNonConstExprInterfaceYGtrArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceY and eval.arrayT)`,
	)

}

// Test InterfaceY > Map
func TestCheckBinaryNonConstExprInterfaceYGtrMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceY and eval.mapT)`,
	)

}

// Test InterfaceY > XinterfaceX
func TestCheckBinaryNonConstExprInterfaceYGtrXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > eval.interfaceY(b) (operator > not defined on interface)`,
	)

}

// Test InterfaceY > InterfaceX
func TestCheckBinaryNonConstExprInterfaceYGtrInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: eval.interfaceX(a) > b (operator > not defined on interface)`,
	)

}

// Test InterfaceY > InterfaceY
func TestCheckBinaryNonConstExprInterfaceYGtrInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (operator > not defined on interface)`,
	)

}

// Test InterfaceY > InterfaceZ
func TestCheckBinaryNonConstExprInterfaceYGtrInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceY and eval.interfaceZ)`,
	)

}

// Test InterfaceY > Ptr
func TestCheckBinaryNonConstExprInterfaceYGtrPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceY and *int)`,
	)

}

// Test InterfaceY > Struct
func TestCheckBinaryNonConstExprInterfaceYGtrStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceY and eval.structT)`,
	)

}

// Test InterfaceY > StructUncomp
func TestCheckBinaryNonConstExprInterfaceYGtrStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceY and eval.structUncompT)`,
	)

}

// Test InterfaceY << ConstInt
func TestCheckBinaryNonConstExprInterfaceYShlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 4`, env,
		`invalid operation: a << 4 (shift of type eval.interfaceY)`,
	)

}

// Test InterfaceY << ConstRune
func TestCheckBinaryNonConstExprInterfaceYShlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << '@'`, env,
		`invalid operation: a << 64 (shift of type eval.interfaceY)`,
	)

}

// Test InterfaceY << ConstFloat
func TestCheckBinaryNonConstExprInterfaceYShlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 2.0`, env,
		`invalid operation: a << 2 (shift of type eval.interfaceY)`,
	)

}

// Test InterfaceY << ConstComplex
func TestCheckBinaryNonConstExprInterfaceYShlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: a << 0 (shift of type eval.interfaceY)`,
	)

}

// Test InterfaceY << ConstBool
func TestCheckBinaryNonConstExprInterfaceYShlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << true`, env,
		`invalid operation: a << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test InterfaceY << ConstString
func TestCheckBinaryNonConstExprInterfaceYShlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: a << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test InterfaceY << ConstNil
func TestCheckBinaryNonConstExprInterfaceYShlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test InterfaceY << Int
func TestCheckBinaryNonConstExprInterfaceYShlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type int, must be unsigned integer)`,
	)

}

// Test InterfaceY << Float32
func TestCheckBinaryNonConstExprInterfaceYShlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type float32, must be unsigned integer)`,
	)

}

// Test InterfaceY << Complex128
func TestCheckBinaryNonConstExprInterfaceYShlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type complex128, must be unsigned integer)`,
	)

}

// Test InterfaceY << String
func TestCheckBinaryNonConstExprInterfaceYShlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type string, must be unsigned integer)`,
	)

}

// Test InterfaceY << BoolT
func TestCheckBinaryNonConstExprInterfaceYShlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type bool, must be unsigned integer)`,
	)

}

// Test InterfaceY << Slice
func TestCheckBinaryNonConstExprInterfaceYShlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.sliceT, must be unsigned integer)`,
	)

}

// Test InterfaceY << Array
func TestCheckBinaryNonConstExprInterfaceYShlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.arrayT, must be unsigned integer)`,
	)

}

// Test InterfaceY << Map
func TestCheckBinaryNonConstExprInterfaceYShlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.mapT, must be unsigned integer)`,
	)

}

// Test InterfaceY << XinterfaceX
func TestCheckBinaryNonConstExprInterfaceYShlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.XinterfaceX, must be unsigned integer)`,
	)

}

// Test InterfaceY << InterfaceX
func TestCheckBinaryNonConstExprInterfaceYShlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceX, must be unsigned integer)`,
	)

}

// Test InterfaceY << InterfaceY
func TestCheckBinaryNonConstExprInterfaceYShlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceY, must be unsigned integer)`,
	)

}

// Test InterfaceY << InterfaceZ
func TestCheckBinaryNonConstExprInterfaceYShlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceZ, must be unsigned integer)`,
	)

}

// Test InterfaceY << Ptr
func TestCheckBinaryNonConstExprInterfaceYShlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type *int, must be unsigned integer)`,
	)

}

// Test InterfaceY << Struct
func TestCheckBinaryNonConstExprInterfaceYShlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structT, must be unsigned integer)`,
	)

}

// Test InterfaceY << StructUncomp
func TestCheckBinaryNonConstExprInterfaceYShlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceY(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structUncompT, must be unsigned integer)`,
	)

}

// Test InterfaceZ + ConstInt
func TestCheckBinaryNonConstExprInterfaceZAddConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 4`, env,
		`invalid operation: a + 4 (mismatched types eval.interfaceZ and int)`,
	)

}

// Test InterfaceZ + ConstRune
func TestCheckBinaryNonConstExprInterfaceZAddConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + '@'`, env,
		`invalid operation: a + rune(64) (mismatched types eval.interfaceZ and rune)`,
	)

}

// Test InterfaceZ + ConstFloat
func TestCheckBinaryNonConstExprInterfaceZAddConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 2.0`, env,
		`invalid operation: a + 2 (mismatched types eval.interfaceZ and float64)`,
	)

}

// Test InterfaceZ + ConstComplex
func TestCheckBinaryNonConstExprInterfaceZAddConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 8.0i`, env,
		`invalid operation: a + 8i (mismatched types eval.interfaceZ and complex128)`,
	)

}

// Test InterfaceZ + ConstBool
func TestCheckBinaryNonConstExprInterfaceZAddConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + true`, env,
		`invalid operation: a + true (mismatched types eval.interfaceZ and bool)`,
	)

}

// Test InterfaceZ + ConstString
func TestCheckBinaryNonConstExprInterfaceZAddConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + "abc"`, env,
		`invalid operation: a + "abc" (mismatched types eval.interfaceZ and string)`,
	)

}

// Test InterfaceZ + ConstNil
func TestCheckBinaryNonConstExprInterfaceZAddConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + nil`, env,
		`invalid operation: a + nil (operator + not defined on interface)`,
	)

}

// Test InterfaceZ + Int
func TestCheckBinaryNonConstExprInterfaceZAddInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceZ and int)`,
	)

}

// Test InterfaceZ + Float32
func TestCheckBinaryNonConstExprInterfaceZAddFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceZ and float32)`,
	)

}

// Test InterfaceZ + Complex128
func TestCheckBinaryNonConstExprInterfaceZAddComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceZ and complex128)`,
	)

}

// Test InterfaceZ + String
func TestCheckBinaryNonConstExprInterfaceZAddString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceZ and string)`,
	)

}

// Test InterfaceZ + BoolT
func TestCheckBinaryNonConstExprInterfaceZAddBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceZ and bool)`,
	)

}

// Test InterfaceZ + Slice
func TestCheckBinaryNonConstExprInterfaceZAddSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceZ and eval.sliceT)`,
	)

}

// Test InterfaceZ + Array
func TestCheckBinaryNonConstExprInterfaceZAddArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceZ and eval.arrayT)`,
	)

}

// Test InterfaceZ + Map
func TestCheckBinaryNonConstExprInterfaceZAddMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceZ and eval.mapT)`,
	)

}

// Test InterfaceZ + XinterfaceX
func TestCheckBinaryNonConstExprInterfaceZAddXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceZ and eval.XinterfaceX)`,
	)

}

// Test InterfaceZ + InterfaceX
func TestCheckBinaryNonConstExprInterfaceZAddInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceZ and eval.interfaceX)`,
	)

}

// Test InterfaceZ + InterfaceY
func TestCheckBinaryNonConstExprInterfaceZAddInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceZ and eval.interfaceY)`,
	)

}

// Test InterfaceZ + InterfaceZ
func TestCheckBinaryNonConstExprInterfaceZAddInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (operator + not defined on interface)`,
	)

}

// Test InterfaceZ + Ptr
func TestCheckBinaryNonConstExprInterfaceZAddPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceZ and *int)`,
	)

}

// Test InterfaceZ + Struct
func TestCheckBinaryNonConstExprInterfaceZAddStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceZ and eval.structT)`,
	)

}

// Test InterfaceZ + StructUncomp
func TestCheckBinaryNonConstExprInterfaceZAddStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.interfaceZ and eval.structUncompT)`,
	)

}

// Test InterfaceZ & ConstInt
func TestCheckBinaryNonConstExprInterfaceZAndConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 4`, env,
		`invalid operation: a & 4 (mismatched types eval.interfaceZ and int)`,
	)

}

// Test InterfaceZ & ConstRune
func TestCheckBinaryNonConstExprInterfaceZAndConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & '@'`, env,
		`invalid operation: a & rune(64) (mismatched types eval.interfaceZ and rune)`,
	)

}

// Test InterfaceZ & ConstFloat
func TestCheckBinaryNonConstExprInterfaceZAndConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 2.0`, env,
		`invalid operation: a & 2 (mismatched types eval.interfaceZ and float64)`,
	)

}

// Test InterfaceZ & ConstComplex
func TestCheckBinaryNonConstExprInterfaceZAndConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 8.0i`, env,
		`invalid operation: a & 8i (mismatched types eval.interfaceZ and complex128)`,
	)

}

// Test InterfaceZ & ConstBool
func TestCheckBinaryNonConstExprInterfaceZAndConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & true`, env,
		`invalid operation: a & true (mismatched types eval.interfaceZ and bool)`,
	)

}

// Test InterfaceZ & ConstString
func TestCheckBinaryNonConstExprInterfaceZAndConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & "abc"`, env,
		`invalid operation: a & "abc" (mismatched types eval.interfaceZ and string)`,
	)

}

// Test InterfaceZ & ConstNil
func TestCheckBinaryNonConstExprInterfaceZAndConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & nil`, env,
		`invalid operation: a & nil (operator & not defined on interface)`,
	)

}

// Test InterfaceZ & Int
func TestCheckBinaryNonConstExprInterfaceZAndInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceZ and int)`,
	)

}

// Test InterfaceZ & Float32
func TestCheckBinaryNonConstExprInterfaceZAndFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceZ and float32)`,
	)

}

// Test InterfaceZ & Complex128
func TestCheckBinaryNonConstExprInterfaceZAndComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceZ and complex128)`,
	)

}

// Test InterfaceZ & String
func TestCheckBinaryNonConstExprInterfaceZAndString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceZ and string)`,
	)

}

// Test InterfaceZ & BoolT
func TestCheckBinaryNonConstExprInterfaceZAndBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceZ and bool)`,
	)

}

// Test InterfaceZ & Slice
func TestCheckBinaryNonConstExprInterfaceZAndSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceZ and eval.sliceT)`,
	)

}

// Test InterfaceZ & Array
func TestCheckBinaryNonConstExprInterfaceZAndArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceZ and eval.arrayT)`,
	)

}

// Test InterfaceZ & Map
func TestCheckBinaryNonConstExprInterfaceZAndMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceZ and eval.mapT)`,
	)

}

// Test InterfaceZ & XinterfaceX
func TestCheckBinaryNonConstExprInterfaceZAndXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceZ and eval.XinterfaceX)`,
	)

}

// Test InterfaceZ & InterfaceX
func TestCheckBinaryNonConstExprInterfaceZAndInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceZ and eval.interfaceX)`,
	)

}

// Test InterfaceZ & InterfaceY
func TestCheckBinaryNonConstExprInterfaceZAndInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceZ and eval.interfaceY)`,
	)

}

// Test InterfaceZ & InterfaceZ
func TestCheckBinaryNonConstExprInterfaceZAndInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (operator & not defined on interface)`,
	)

}

// Test InterfaceZ & Ptr
func TestCheckBinaryNonConstExprInterfaceZAndPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceZ and *int)`,
	)

}

// Test InterfaceZ & Struct
func TestCheckBinaryNonConstExprInterfaceZAndStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceZ and eval.structT)`,
	)

}

// Test InterfaceZ & StructUncomp
func TestCheckBinaryNonConstExprInterfaceZAndStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.interfaceZ and eval.structUncompT)`,
	)

}

// Test InterfaceZ % ConstInt
func TestCheckBinaryNonConstExprInterfaceZRemConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 4`, env,
		`invalid operation: a % 4 (mismatched types eval.interfaceZ and int)`,
	)

}

// Test InterfaceZ % ConstRune
func TestCheckBinaryNonConstExprInterfaceZRemConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % '@'`, env,
		`invalid operation: a % rune(64) (mismatched types eval.interfaceZ and rune)`,
	)

}

// Test InterfaceZ % ConstFloat
func TestCheckBinaryNonConstExprInterfaceZRemConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 2.0`, env,
		`invalid operation: a % 2 (mismatched types eval.interfaceZ and float64)`,
	)

}

// Test InterfaceZ % ConstComplex
func TestCheckBinaryNonConstExprInterfaceZRemConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 8.0i`, env,
		`invalid operation: a % 8i (mismatched types eval.interfaceZ and complex128)`,
	)

}

// Test InterfaceZ % ConstBool
func TestCheckBinaryNonConstExprInterfaceZRemConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % true`, env,
		`invalid operation: a % true (mismatched types eval.interfaceZ and bool)`,
	)

}

// Test InterfaceZ % ConstString
func TestCheckBinaryNonConstExprInterfaceZRemConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % "abc"`, env,
		`invalid operation: a % "abc" (mismatched types eval.interfaceZ and string)`,
	)

}

// Test InterfaceZ % ConstNil
func TestCheckBinaryNonConstExprInterfaceZRemConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % nil`, env,
		`invalid operation: a % nil (operator % not defined on interface)`,
	)

}

// Test InterfaceZ % Int
func TestCheckBinaryNonConstExprInterfaceZRemInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceZ and int)`,
	)

}

// Test InterfaceZ % Float32
func TestCheckBinaryNonConstExprInterfaceZRemFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceZ and float32)`,
	)

}

// Test InterfaceZ % Complex128
func TestCheckBinaryNonConstExprInterfaceZRemComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceZ and complex128)`,
	)

}

// Test InterfaceZ % String
func TestCheckBinaryNonConstExprInterfaceZRemString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceZ and string)`,
	)

}

// Test InterfaceZ % BoolT
func TestCheckBinaryNonConstExprInterfaceZRemBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceZ and bool)`,
	)

}

// Test InterfaceZ % Slice
func TestCheckBinaryNonConstExprInterfaceZRemSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceZ and eval.sliceT)`,
	)

}

// Test InterfaceZ % Array
func TestCheckBinaryNonConstExprInterfaceZRemArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceZ and eval.arrayT)`,
	)

}

// Test InterfaceZ % Map
func TestCheckBinaryNonConstExprInterfaceZRemMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceZ and eval.mapT)`,
	)

}

// Test InterfaceZ % XinterfaceX
func TestCheckBinaryNonConstExprInterfaceZRemXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceZ and eval.XinterfaceX)`,
	)

}

// Test InterfaceZ % InterfaceX
func TestCheckBinaryNonConstExprInterfaceZRemInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceZ and eval.interfaceX)`,
	)

}

// Test InterfaceZ % InterfaceY
func TestCheckBinaryNonConstExprInterfaceZRemInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceZ and eval.interfaceY)`,
	)

}

// Test InterfaceZ % InterfaceZ
func TestCheckBinaryNonConstExprInterfaceZRemInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (operator % not defined on interface)`,
	)

}

// Test InterfaceZ % Ptr
func TestCheckBinaryNonConstExprInterfaceZRemPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceZ and *int)`,
	)

}

// Test InterfaceZ % Struct
func TestCheckBinaryNonConstExprInterfaceZRemStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceZ and eval.structT)`,
	)

}

// Test InterfaceZ % StructUncomp
func TestCheckBinaryNonConstExprInterfaceZRemStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.interfaceZ and eval.structUncompT)`,
	)

}

// Test InterfaceZ == ConstInt
func TestCheckBinaryNonConstExprInterfaceZEqlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 4`, env,
		`invalid operation: a == 4 (mismatched types eval.interfaceZ and int)`,
	)

}

// Test InterfaceZ == ConstRune
func TestCheckBinaryNonConstExprInterfaceZEqlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == '@'`, env,
		`invalid operation: a == rune(64) (mismatched types eval.interfaceZ and rune)`,
	)

}

// Test InterfaceZ == ConstFloat
func TestCheckBinaryNonConstExprInterfaceZEqlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 2.0`, env,
		`invalid operation: a == 2 (mismatched types eval.interfaceZ and float64)`,
	)

}

// Test InterfaceZ == ConstComplex
func TestCheckBinaryNonConstExprInterfaceZEqlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 8.0i`, env,
		`invalid operation: a == 8i (mismatched types eval.interfaceZ and complex128)`,
	)

}

// Test InterfaceZ == ConstBool
func TestCheckBinaryNonConstExprInterfaceZEqlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == true`, env,
		`invalid operation: a == true (mismatched types eval.interfaceZ and bool)`,
	)

}

// Test InterfaceZ == ConstString
func TestCheckBinaryNonConstExprInterfaceZEqlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == "abc"`, env,
		`invalid operation: a == "abc" (mismatched types eval.interfaceZ and string)`,
	)

}

// Test InterfaceZ == ConstNil
func TestCheckBinaryNonConstExprInterfaceZEqlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == nil`, env, reflect.TypeOf(a == nil))
}

// Test InterfaceZ == Int
func TestCheckBinaryNonConstExprInterfaceZEqlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceZ and int)`,
	)

}

// Test InterfaceZ == Float32
func TestCheckBinaryNonConstExprInterfaceZEqlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceZ and float32)`,
	)

}

// Test InterfaceZ == Complex128
func TestCheckBinaryNonConstExprInterfaceZEqlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceZ and complex128)`,
	)

}

// Test InterfaceZ == String
func TestCheckBinaryNonConstExprInterfaceZEqlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceZ and string)`,
	)

}

// Test InterfaceZ == BoolT
func TestCheckBinaryNonConstExprInterfaceZEqlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceZ and bool)`,
	)

}

// Test InterfaceZ == Slice
func TestCheckBinaryNonConstExprInterfaceZEqlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceZ and eval.sliceT)`,
	)

}

// Test InterfaceZ == Array
func TestCheckBinaryNonConstExprInterfaceZEqlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceZ and eval.arrayT)`,
	)

}

// Test InterfaceZ == Map
func TestCheckBinaryNonConstExprInterfaceZEqlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceZ and eval.mapT)`,
	)

}

// Test InterfaceZ == XinterfaceX
func TestCheckBinaryNonConstExprInterfaceZEqlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceZ and eval.XinterfaceX)`,
	)

}

// Test InterfaceZ == InterfaceX
func TestCheckBinaryNonConstExprInterfaceZEqlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceZ and eval.interfaceX)`,
	)

}

// Test InterfaceZ == InterfaceY
func TestCheckBinaryNonConstExprInterfaceZEqlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceZ and eval.interfaceY)`,
	)

}

// Test InterfaceZ == InterfaceZ
func TestCheckBinaryNonConstExprInterfaceZEqlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test InterfaceZ == Ptr
func TestCheckBinaryNonConstExprInterfaceZEqlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceZ and *int)`,
	)

}

// Test InterfaceZ == Struct
func TestCheckBinaryNonConstExprInterfaceZEqlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceZ and eval.structT)`,
	)

}

// Test InterfaceZ == StructUncomp
func TestCheckBinaryNonConstExprInterfaceZEqlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.interfaceZ and eval.structUncompT)`,
	)

}

// Test InterfaceZ > ConstInt
func TestCheckBinaryNonConstExprInterfaceZGtrConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 4`, env,
		`invalid operation: a > 4 (mismatched types eval.interfaceZ and int)`,
	)

}

// Test InterfaceZ > ConstRune
func TestCheckBinaryNonConstExprInterfaceZGtrConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > '@'`, env,
		`invalid operation: a > rune(64) (mismatched types eval.interfaceZ and rune)`,
	)

}

// Test InterfaceZ > ConstFloat
func TestCheckBinaryNonConstExprInterfaceZGtrConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 2.0`, env,
		`invalid operation: a > 2 (mismatched types eval.interfaceZ and float64)`,
	)

}

// Test InterfaceZ > ConstComplex
func TestCheckBinaryNonConstExprInterfaceZGtrConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 8.0i`, env,
		`invalid operation: a > 8i (mismatched types eval.interfaceZ and complex128)`,
	)

}

// Test InterfaceZ > ConstBool
func TestCheckBinaryNonConstExprInterfaceZGtrConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > true`, env,
		`invalid operation: a > true (mismatched types eval.interfaceZ and bool)`,
	)

}

// Test InterfaceZ > ConstString
func TestCheckBinaryNonConstExprInterfaceZGtrConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > "abc"`, env,
		`invalid operation: a > "abc" (mismatched types eval.interfaceZ and string)`,
	)

}

// Test InterfaceZ > ConstNil
func TestCheckBinaryNonConstExprInterfaceZGtrConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > nil`, env,
		`invalid operation: a > nil (operator > not defined on interface)`,
	)

}

// Test InterfaceZ > Int
func TestCheckBinaryNonConstExprInterfaceZGtrInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceZ and int)`,
	)

}

// Test InterfaceZ > Float32
func TestCheckBinaryNonConstExprInterfaceZGtrFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceZ and float32)`,
	)

}

// Test InterfaceZ > Complex128
func TestCheckBinaryNonConstExprInterfaceZGtrComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceZ and complex128)`,
	)

}

// Test InterfaceZ > String
func TestCheckBinaryNonConstExprInterfaceZGtrString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceZ and string)`,
	)

}

// Test InterfaceZ > BoolT
func TestCheckBinaryNonConstExprInterfaceZGtrBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceZ and bool)`,
	)

}

// Test InterfaceZ > Slice
func TestCheckBinaryNonConstExprInterfaceZGtrSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceZ and eval.sliceT)`,
	)

}

// Test InterfaceZ > Array
func TestCheckBinaryNonConstExprInterfaceZGtrArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceZ and eval.arrayT)`,
	)

}

// Test InterfaceZ > Map
func TestCheckBinaryNonConstExprInterfaceZGtrMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceZ and eval.mapT)`,
	)

}

// Test InterfaceZ > XinterfaceX
func TestCheckBinaryNonConstExprInterfaceZGtrXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceZ and eval.XinterfaceX)`,
	)

}

// Test InterfaceZ > InterfaceX
func TestCheckBinaryNonConstExprInterfaceZGtrInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceZ and eval.interfaceX)`,
	)

}

// Test InterfaceZ > InterfaceY
func TestCheckBinaryNonConstExprInterfaceZGtrInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceZ and eval.interfaceY)`,
	)

}

// Test InterfaceZ > InterfaceZ
func TestCheckBinaryNonConstExprInterfaceZGtrInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (operator > not defined on interface)`,
	)

}

// Test InterfaceZ > Ptr
func TestCheckBinaryNonConstExprInterfaceZGtrPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceZ and *int)`,
	)

}

// Test InterfaceZ > Struct
func TestCheckBinaryNonConstExprInterfaceZGtrStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceZ and eval.structT)`,
	)

}

// Test InterfaceZ > StructUncomp
func TestCheckBinaryNonConstExprInterfaceZGtrStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.interfaceZ and eval.structUncompT)`,
	)

}

// Test InterfaceZ << ConstInt
func TestCheckBinaryNonConstExprInterfaceZShlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 4`, env,
		`invalid operation: a << 4 (shift of type eval.interfaceZ)`,
	)

}

// Test InterfaceZ << ConstRune
func TestCheckBinaryNonConstExprInterfaceZShlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << '@'`, env,
		`invalid operation: a << 64 (shift of type eval.interfaceZ)`,
	)

}

// Test InterfaceZ << ConstFloat
func TestCheckBinaryNonConstExprInterfaceZShlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 2.0`, env,
		`invalid operation: a << 2 (shift of type eval.interfaceZ)`,
	)

}

// Test InterfaceZ << ConstComplex
func TestCheckBinaryNonConstExprInterfaceZShlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: a << 0 (shift of type eval.interfaceZ)`,
	)

}

// Test InterfaceZ << ConstBool
func TestCheckBinaryNonConstExprInterfaceZShlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << true`, env,
		`invalid operation: a << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test InterfaceZ << ConstString
func TestCheckBinaryNonConstExprInterfaceZShlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: a << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test InterfaceZ << ConstNil
func TestCheckBinaryNonConstExprInterfaceZShlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test InterfaceZ << Int
func TestCheckBinaryNonConstExprInterfaceZShlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type int, must be unsigned integer)`,
	)

}

// Test InterfaceZ << Float32
func TestCheckBinaryNonConstExprInterfaceZShlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type float32, must be unsigned integer)`,
	)

}

// Test InterfaceZ << Complex128
func TestCheckBinaryNonConstExprInterfaceZShlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type complex128, must be unsigned integer)`,
	)

}

// Test InterfaceZ << String
func TestCheckBinaryNonConstExprInterfaceZShlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type string, must be unsigned integer)`,
	)

}

// Test InterfaceZ << BoolT
func TestCheckBinaryNonConstExprInterfaceZShlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type bool, must be unsigned integer)`,
	)

}

// Test InterfaceZ << Slice
func TestCheckBinaryNonConstExprInterfaceZShlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.sliceT, must be unsigned integer)`,
	)

}

// Test InterfaceZ << Array
func TestCheckBinaryNonConstExprInterfaceZShlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.arrayT, must be unsigned integer)`,
	)

}

// Test InterfaceZ << Map
func TestCheckBinaryNonConstExprInterfaceZShlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.mapT, must be unsigned integer)`,
	)

}

// Test InterfaceZ << XinterfaceX
func TestCheckBinaryNonConstExprInterfaceZShlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.XinterfaceX, must be unsigned integer)`,
	)

}

// Test InterfaceZ << InterfaceX
func TestCheckBinaryNonConstExprInterfaceZShlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceX, must be unsigned integer)`,
	)

}

// Test InterfaceZ << InterfaceY
func TestCheckBinaryNonConstExprInterfaceZShlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceY, must be unsigned integer)`,
	)

}

// Test InterfaceZ << InterfaceZ
func TestCheckBinaryNonConstExprInterfaceZShlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceZ, must be unsigned integer)`,
	)

}

// Test InterfaceZ << Ptr
func TestCheckBinaryNonConstExprInterfaceZShlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type *int, must be unsigned integer)`,
	)

}

// Test InterfaceZ << Struct
func TestCheckBinaryNonConstExprInterfaceZShlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structT, must be unsigned integer)`,
	)

}

// Test InterfaceZ << StructUncomp
func TestCheckBinaryNonConstExprInterfaceZShlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := interfaceZ(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structUncompT, must be unsigned integer)`,
	)

}

// Test Ptr + ConstInt
func TestCheckBinaryNonConstExprPtrAddConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 4`, env,
		`invalid operation: a + 4 (mismatched types *int and int)`,
	)

}

// Test Ptr + ConstRune
func TestCheckBinaryNonConstExprPtrAddConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + '@'`, env,
		`invalid operation: a + rune(64) (mismatched types *int and rune)`,
	)

}

// Test Ptr + ConstFloat
func TestCheckBinaryNonConstExprPtrAddConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 2.0`, env,
		`invalid operation: a + 2 (mismatched types *int and float64)`,
	)

}

// Test Ptr + ConstComplex
func TestCheckBinaryNonConstExprPtrAddConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 8.0i`, env,
		`invalid operation: a + 8i (mismatched types *int and complex128)`,
	)

}

// Test Ptr + ConstBool
func TestCheckBinaryNonConstExprPtrAddConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + true`, env,
		`invalid operation: a + true (mismatched types *int and bool)`,
	)

}

// Test Ptr + ConstString
func TestCheckBinaryNonConstExprPtrAddConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + "abc"`, env,
		`invalid operation: a + "abc" (mismatched types *int and string)`,
	)

}

// Test Ptr + ConstNil
func TestCheckBinaryNonConstExprPtrAddConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + nil`, env,
		`invalid operation: a + nil (operator + not defined on pointer)`,
	)

}

// Test Ptr + Int
func TestCheckBinaryNonConstExprPtrAddInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types *int and int)`,
	)

}

// Test Ptr + Float32
func TestCheckBinaryNonConstExprPtrAddFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types *int and float32)`,
	)

}

// Test Ptr + Complex128
func TestCheckBinaryNonConstExprPtrAddComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types *int and complex128)`,
	)

}

// Test Ptr + String
func TestCheckBinaryNonConstExprPtrAddString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types *int and string)`,
	)

}

// Test Ptr + BoolT
func TestCheckBinaryNonConstExprPtrAddBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types *int and bool)`,
	)

}

// Test Ptr + Slice
func TestCheckBinaryNonConstExprPtrAddSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types *int and eval.sliceT)`,
	)

}

// Test Ptr + Array
func TestCheckBinaryNonConstExprPtrAddArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types *int and eval.arrayT)`,
	)

}

// Test Ptr + Map
func TestCheckBinaryNonConstExprPtrAddMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types *int and eval.mapT)`,
	)

}

// Test Ptr + XinterfaceX
func TestCheckBinaryNonConstExprPtrAddXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types *int and eval.XinterfaceX)`,
	)

}

// Test Ptr + InterfaceX
func TestCheckBinaryNonConstExprPtrAddInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types *int and eval.interfaceX)`,
	)

}

// Test Ptr + InterfaceY
func TestCheckBinaryNonConstExprPtrAddInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types *int and eval.interfaceY)`,
	)

}

// Test Ptr + InterfaceZ
func TestCheckBinaryNonConstExprPtrAddInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types *int and eval.interfaceZ)`,
	)

}

// Test Ptr + Ptr
func TestCheckBinaryNonConstExprPtrAddPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (operator + not defined on pointer)`,
	)

}

// Test Ptr + Struct
func TestCheckBinaryNonConstExprPtrAddStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types *int and eval.structT)`,
	)

}

// Test Ptr + StructUncomp
func TestCheckBinaryNonConstExprPtrAddStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types *int and eval.structUncompT)`,
	)

}

// Test Ptr & ConstInt
func TestCheckBinaryNonConstExprPtrAndConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 4`, env,
		`invalid operation: a & 4 (mismatched types *int and int)`,
	)

}

// Test Ptr & ConstRune
func TestCheckBinaryNonConstExprPtrAndConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & '@'`, env,
		`invalid operation: a & rune(64) (mismatched types *int and rune)`,
	)

}

// Test Ptr & ConstFloat
func TestCheckBinaryNonConstExprPtrAndConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 2.0`, env,
		`invalid operation: a & 2 (mismatched types *int and float64)`,
	)

}

// Test Ptr & ConstComplex
func TestCheckBinaryNonConstExprPtrAndConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 8.0i`, env,
		`invalid operation: a & 8i (mismatched types *int and complex128)`,
	)

}

// Test Ptr & ConstBool
func TestCheckBinaryNonConstExprPtrAndConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & true`, env,
		`invalid operation: a & true (mismatched types *int and bool)`,
	)

}

// Test Ptr & ConstString
func TestCheckBinaryNonConstExprPtrAndConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & "abc"`, env,
		`invalid operation: a & "abc" (mismatched types *int and string)`,
	)

}

// Test Ptr & ConstNil
func TestCheckBinaryNonConstExprPtrAndConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & nil`, env,
		`invalid operation: a & nil (operator & not defined on pointer)`,
	)

}

// Test Ptr & Int
func TestCheckBinaryNonConstExprPtrAndInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types *int and int)`,
	)

}

// Test Ptr & Float32
func TestCheckBinaryNonConstExprPtrAndFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types *int and float32)`,
	)

}

// Test Ptr & Complex128
func TestCheckBinaryNonConstExprPtrAndComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types *int and complex128)`,
	)

}

// Test Ptr & String
func TestCheckBinaryNonConstExprPtrAndString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types *int and string)`,
	)

}

// Test Ptr & BoolT
func TestCheckBinaryNonConstExprPtrAndBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types *int and bool)`,
	)

}

// Test Ptr & Slice
func TestCheckBinaryNonConstExprPtrAndSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types *int and eval.sliceT)`,
	)

}

// Test Ptr & Array
func TestCheckBinaryNonConstExprPtrAndArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types *int and eval.arrayT)`,
	)

}

// Test Ptr & Map
func TestCheckBinaryNonConstExprPtrAndMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types *int and eval.mapT)`,
	)

}

// Test Ptr & XinterfaceX
func TestCheckBinaryNonConstExprPtrAndXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types *int and eval.XinterfaceX)`,
	)

}

// Test Ptr & InterfaceX
func TestCheckBinaryNonConstExprPtrAndInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types *int and eval.interfaceX)`,
	)

}

// Test Ptr & InterfaceY
func TestCheckBinaryNonConstExprPtrAndInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types *int and eval.interfaceY)`,
	)

}

// Test Ptr & InterfaceZ
func TestCheckBinaryNonConstExprPtrAndInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types *int and eval.interfaceZ)`,
	)

}

// Test Ptr & Ptr
func TestCheckBinaryNonConstExprPtrAndPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (operator & not defined on pointer)`,
	)

}

// Test Ptr & Struct
func TestCheckBinaryNonConstExprPtrAndStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types *int and eval.structT)`,
	)

}

// Test Ptr & StructUncomp
func TestCheckBinaryNonConstExprPtrAndStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types *int and eval.structUncompT)`,
	)

}

// Test Ptr % ConstInt
func TestCheckBinaryNonConstExprPtrRemConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 4`, env,
		`invalid operation: a % 4 (mismatched types *int and int)`,
	)

}

// Test Ptr % ConstRune
func TestCheckBinaryNonConstExprPtrRemConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % '@'`, env,
		`invalid operation: a % rune(64) (mismatched types *int and rune)`,
	)

}

// Test Ptr % ConstFloat
func TestCheckBinaryNonConstExprPtrRemConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 2.0`, env,
		`invalid operation: a % 2 (mismatched types *int and float64)`,
	)

}

// Test Ptr % ConstComplex
func TestCheckBinaryNonConstExprPtrRemConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 8.0i`, env,
		`invalid operation: a % 8i (mismatched types *int and complex128)`,
	)

}

// Test Ptr % ConstBool
func TestCheckBinaryNonConstExprPtrRemConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % true`, env,
		`invalid operation: a % true (mismatched types *int and bool)`,
	)

}

// Test Ptr % ConstString
func TestCheckBinaryNonConstExprPtrRemConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % "abc"`, env,
		`invalid operation: a % "abc" (mismatched types *int and string)`,
	)

}

// Test Ptr % ConstNil
func TestCheckBinaryNonConstExprPtrRemConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % nil`, env,
		`invalid operation: a % nil (operator % not defined on pointer)`,
	)

}

// Test Ptr % Int
func TestCheckBinaryNonConstExprPtrRemInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types *int and int)`,
	)

}

// Test Ptr % Float32
func TestCheckBinaryNonConstExprPtrRemFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types *int and float32)`,
	)

}

// Test Ptr % Complex128
func TestCheckBinaryNonConstExprPtrRemComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types *int and complex128)`,
	)

}

// Test Ptr % String
func TestCheckBinaryNonConstExprPtrRemString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types *int and string)`,
	)

}

// Test Ptr % BoolT
func TestCheckBinaryNonConstExprPtrRemBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types *int and bool)`,
	)

}

// Test Ptr % Slice
func TestCheckBinaryNonConstExprPtrRemSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types *int and eval.sliceT)`,
	)

}

// Test Ptr % Array
func TestCheckBinaryNonConstExprPtrRemArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types *int and eval.arrayT)`,
	)

}

// Test Ptr % Map
func TestCheckBinaryNonConstExprPtrRemMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types *int and eval.mapT)`,
	)

}

// Test Ptr % XinterfaceX
func TestCheckBinaryNonConstExprPtrRemXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types *int and eval.XinterfaceX)`,
	)

}

// Test Ptr % InterfaceX
func TestCheckBinaryNonConstExprPtrRemInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types *int and eval.interfaceX)`,
	)

}

// Test Ptr % InterfaceY
func TestCheckBinaryNonConstExprPtrRemInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types *int and eval.interfaceY)`,
	)

}

// Test Ptr % InterfaceZ
func TestCheckBinaryNonConstExprPtrRemInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types *int and eval.interfaceZ)`,
	)

}

// Test Ptr % Ptr
func TestCheckBinaryNonConstExprPtrRemPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (operator % not defined on pointer)`,
	)

}

// Test Ptr % Struct
func TestCheckBinaryNonConstExprPtrRemStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types *int and eval.structT)`,
	)

}

// Test Ptr % StructUncomp
func TestCheckBinaryNonConstExprPtrRemStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types *int and eval.structUncompT)`,
	)

}

// Test Ptr == ConstInt
func TestCheckBinaryNonConstExprPtrEqlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 4`, env,
		`invalid operation: a == 4 (mismatched types *int and int)`,
	)

}

// Test Ptr == ConstRune
func TestCheckBinaryNonConstExprPtrEqlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == '@'`, env,
		`invalid operation: a == rune(64) (mismatched types *int and rune)`,
	)

}

// Test Ptr == ConstFloat
func TestCheckBinaryNonConstExprPtrEqlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 2.0`, env,
		`invalid operation: a == 2 (mismatched types *int and float64)`,
	)

}

// Test Ptr == ConstComplex
func TestCheckBinaryNonConstExprPtrEqlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 8.0i`, env,
		`invalid operation: a == 8i (mismatched types *int and complex128)`,
	)

}

// Test Ptr == ConstBool
func TestCheckBinaryNonConstExprPtrEqlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == true`, env,
		`invalid operation: a == true (mismatched types *int and bool)`,
	)

}

// Test Ptr == ConstString
func TestCheckBinaryNonConstExprPtrEqlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == "abc"`, env,
		`invalid operation: a == "abc" (mismatched types *int and string)`,
	)

}

// Test Ptr == ConstNil
func TestCheckBinaryNonConstExprPtrEqlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectType(t, `a == nil`, env, reflect.TypeOf(a == nil))
}

// Test Ptr == Int
func TestCheckBinaryNonConstExprPtrEqlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types *int and int)`,
	)

}

// Test Ptr == Float32
func TestCheckBinaryNonConstExprPtrEqlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types *int and float32)`,
	)

}

// Test Ptr == Complex128
func TestCheckBinaryNonConstExprPtrEqlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types *int and complex128)`,
	)

}

// Test Ptr == String
func TestCheckBinaryNonConstExprPtrEqlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types *int and string)`,
	)

}

// Test Ptr == BoolT
func TestCheckBinaryNonConstExprPtrEqlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types *int and bool)`,
	)

}

// Test Ptr == Slice
func TestCheckBinaryNonConstExprPtrEqlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types *int and eval.sliceT)`,
	)

}

// Test Ptr == Array
func TestCheckBinaryNonConstExprPtrEqlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types *int and eval.arrayT)`,
	)

}

// Test Ptr == Map
func TestCheckBinaryNonConstExprPtrEqlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types *int and eval.mapT)`,
	)

}

// Test Ptr == XinterfaceX
func TestCheckBinaryNonConstExprPtrEqlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types *int and eval.XinterfaceX)`,
	)

}

// Test Ptr == InterfaceX
func TestCheckBinaryNonConstExprPtrEqlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types *int and eval.interfaceX)`,
	)

}

// Test Ptr == InterfaceY
func TestCheckBinaryNonConstExprPtrEqlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types *int and eval.interfaceY)`,
	)

}

// Test Ptr == InterfaceZ
func TestCheckBinaryNonConstExprPtrEqlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types *int and eval.interfaceZ)`,
	)

}

// Test Ptr == Ptr
func TestCheckBinaryNonConstExprPtrEqlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectType(t, `a == b`, env, reflect.TypeOf(a == b))
}

// Test Ptr == Struct
func TestCheckBinaryNonConstExprPtrEqlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types *int and eval.structT)`,
	)

}

// Test Ptr == StructUncomp
func TestCheckBinaryNonConstExprPtrEqlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types *int and eval.structUncompT)`,
	)

}

// Test Ptr > ConstInt
func TestCheckBinaryNonConstExprPtrGtrConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 4`, env,
		`invalid operation: a > 4 (mismatched types *int and int)`,
	)

}

// Test Ptr > ConstRune
func TestCheckBinaryNonConstExprPtrGtrConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > '@'`, env,
		`invalid operation: a > rune(64) (mismatched types *int and rune)`,
	)

}

// Test Ptr > ConstFloat
func TestCheckBinaryNonConstExprPtrGtrConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 2.0`, env,
		`invalid operation: a > 2 (mismatched types *int and float64)`,
	)

}

// Test Ptr > ConstComplex
func TestCheckBinaryNonConstExprPtrGtrConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 8.0i`, env,
		`invalid operation: a > 8i (mismatched types *int and complex128)`,
	)

}

// Test Ptr > ConstBool
func TestCheckBinaryNonConstExprPtrGtrConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > true`, env,
		`invalid operation: a > true (mismatched types *int and bool)`,
	)

}

// Test Ptr > ConstString
func TestCheckBinaryNonConstExprPtrGtrConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > "abc"`, env,
		`invalid operation: a > "abc" (mismatched types *int and string)`,
	)

}

// Test Ptr > ConstNil
func TestCheckBinaryNonConstExprPtrGtrConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > nil`, env,
		`invalid operation: a > nil (operator > not defined on pointer)`,
	)

}

// Test Ptr > Int
func TestCheckBinaryNonConstExprPtrGtrInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types *int and int)`,
	)

}

// Test Ptr > Float32
func TestCheckBinaryNonConstExprPtrGtrFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types *int and float32)`,
	)

}

// Test Ptr > Complex128
func TestCheckBinaryNonConstExprPtrGtrComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types *int and complex128)`,
	)

}

// Test Ptr > String
func TestCheckBinaryNonConstExprPtrGtrString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types *int and string)`,
	)

}

// Test Ptr > BoolT
func TestCheckBinaryNonConstExprPtrGtrBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types *int and bool)`,
	)

}

// Test Ptr > Slice
func TestCheckBinaryNonConstExprPtrGtrSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types *int and eval.sliceT)`,
	)

}

// Test Ptr > Array
func TestCheckBinaryNonConstExprPtrGtrArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types *int and eval.arrayT)`,
	)

}

// Test Ptr > Map
func TestCheckBinaryNonConstExprPtrGtrMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types *int and eval.mapT)`,
	)

}

// Test Ptr > XinterfaceX
func TestCheckBinaryNonConstExprPtrGtrXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types *int and eval.XinterfaceX)`,
	)

}

// Test Ptr > InterfaceX
func TestCheckBinaryNonConstExprPtrGtrInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types *int and eval.interfaceX)`,
	)

}

// Test Ptr > InterfaceY
func TestCheckBinaryNonConstExprPtrGtrInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types *int and eval.interfaceY)`,
	)

}

// Test Ptr > InterfaceZ
func TestCheckBinaryNonConstExprPtrGtrInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types *int and eval.interfaceZ)`,
	)

}

// Test Ptr > Ptr
func TestCheckBinaryNonConstExprPtrGtrPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (operator > not defined on pointer)`,
	)

}

// Test Ptr > Struct
func TestCheckBinaryNonConstExprPtrGtrStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types *int and eval.structT)`,
	)

}

// Test Ptr > StructUncomp
func TestCheckBinaryNonConstExprPtrGtrStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types *int and eval.structUncompT)`,
	)

}

// Test Ptr << ConstInt
func TestCheckBinaryNonConstExprPtrShlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 4`, env,
		`invalid operation: a << 4 (shift of type *int)`,
	)

}

// Test Ptr << ConstRune
func TestCheckBinaryNonConstExprPtrShlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << '@'`, env,
		`invalid operation: a << 64 (shift of type *int)`,
	)

}

// Test Ptr << ConstFloat
func TestCheckBinaryNonConstExprPtrShlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 2.0`, env,
		`invalid operation: a << 2 (shift of type *int)`,
	)

}

// Test Ptr << ConstComplex
func TestCheckBinaryNonConstExprPtrShlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: a << 0 (shift of type *int)`,
	)

}

// Test Ptr << ConstBool
func TestCheckBinaryNonConstExprPtrShlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << true`, env,
		`invalid operation: a << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Ptr << ConstString
func TestCheckBinaryNonConstExprPtrShlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: a << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Ptr << ConstNil
func TestCheckBinaryNonConstExprPtrShlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Ptr << Int
func TestCheckBinaryNonConstExprPtrShlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type int, must be unsigned integer)`,
	)

}

// Test Ptr << Float32
func TestCheckBinaryNonConstExprPtrShlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type float32, must be unsigned integer)`,
	)

}

// Test Ptr << Complex128
func TestCheckBinaryNonConstExprPtrShlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Ptr << String
func TestCheckBinaryNonConstExprPtrShlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type string, must be unsigned integer)`,
	)

}

// Test Ptr << BoolT
func TestCheckBinaryNonConstExprPtrShlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type bool, must be unsigned integer)`,
	)

}

// Test Ptr << Slice
func TestCheckBinaryNonConstExprPtrShlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.sliceT, must be unsigned integer)`,
	)

}

// Test Ptr << Array
func TestCheckBinaryNonConstExprPtrShlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.arrayT, must be unsigned integer)`,
	)

}

// Test Ptr << Map
func TestCheckBinaryNonConstExprPtrShlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.mapT, must be unsigned integer)`,
	)

}

// Test Ptr << XinterfaceX
func TestCheckBinaryNonConstExprPtrShlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.XinterfaceX, must be unsigned integer)`,
	)

}

// Test Ptr << InterfaceX
func TestCheckBinaryNonConstExprPtrShlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceX, must be unsigned integer)`,
	)

}

// Test Ptr << InterfaceY
func TestCheckBinaryNonConstExprPtrShlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceY, must be unsigned integer)`,
	)

}

// Test Ptr << InterfaceZ
func TestCheckBinaryNonConstExprPtrShlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceZ, must be unsigned integer)`,
	)

}

// Test Ptr << Ptr
func TestCheckBinaryNonConstExprPtrShlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type *int, must be unsigned integer)`,
	)

}

// Test Ptr << Struct
func TestCheckBinaryNonConstExprPtrShlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structT, must be unsigned integer)`,
	)

}

// Test Ptr << StructUncomp
func TestCheckBinaryNonConstExprPtrShlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := (*int)(nil); env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structUncompT, must be unsigned integer)`,
	)

}

// Test Struct + ConstInt
func TestCheckBinaryNonConstExprStructAddConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 4`, env,
		`invalid operation: a + 4 (mismatched types eval.structT and int)`,
	)

}

// Test Struct + ConstRune
func TestCheckBinaryNonConstExprStructAddConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + '@'`, env,
		`invalid operation: a + rune(64) (mismatched types eval.structT and rune)`,
	)

}

// Test Struct + ConstFloat
func TestCheckBinaryNonConstExprStructAddConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 2.0`, env,
		`invalid operation: a + 2 (mismatched types eval.structT and float64)`,
	)

}

// Test Struct + ConstComplex
func TestCheckBinaryNonConstExprStructAddConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 8.0i`, env,
		`invalid operation: a + 8i (mismatched types eval.structT and complex128)`,
	)

}

// Test Struct + ConstBool
func TestCheckBinaryNonConstExprStructAddConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + true`, env,
		`invalid operation: a + true (mismatched types eval.structT and bool)`,
	)

}

// Test Struct + ConstString
func TestCheckBinaryNonConstExprStructAddConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + "abc"`, env,
		`invalid operation: a + "abc" (mismatched types eval.structT and string)`,
	)

}

// Test Struct + ConstNil
func TestCheckBinaryNonConstExprStructAddConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + nil`, env,
		`cannot convert nil to type eval.structT`,
	)

}

// Test Struct + Int
func TestCheckBinaryNonConstExprStructAddInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structT and int)`,
	)

}

// Test Struct + Float32
func TestCheckBinaryNonConstExprStructAddFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structT and float32)`,
	)

}

// Test Struct + Complex128
func TestCheckBinaryNonConstExprStructAddComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structT and complex128)`,
	)

}

// Test Struct + String
func TestCheckBinaryNonConstExprStructAddString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structT and string)`,
	)

}

// Test Struct + BoolT
func TestCheckBinaryNonConstExprStructAddBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structT and bool)`,
	)

}

// Test Struct + Slice
func TestCheckBinaryNonConstExprStructAddSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structT and eval.sliceT)`,
	)

}

// Test Struct + Array
func TestCheckBinaryNonConstExprStructAddArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structT and eval.arrayT)`,
	)

}

// Test Struct + Map
func TestCheckBinaryNonConstExprStructAddMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structT and eval.mapT)`,
	)

}

// Test Struct + XinterfaceX
func TestCheckBinaryNonConstExprStructAddXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structT and eval.XinterfaceX)`,
	)

}

// Test Struct + InterfaceX
func TestCheckBinaryNonConstExprStructAddInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structT and eval.interfaceX)`,
	)

}

// Test Struct + InterfaceY
func TestCheckBinaryNonConstExprStructAddInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structT and eval.interfaceY)`,
	)

}

// Test Struct + InterfaceZ
func TestCheckBinaryNonConstExprStructAddInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structT and eval.interfaceZ)`,
	)

}

// Test Struct + Ptr
func TestCheckBinaryNonConstExprStructAddPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structT and *int)`,
	)

}

// Test Struct + Struct
func TestCheckBinaryNonConstExprStructAddStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (operator + not defined on struct)`,
	)

}

// Test Struct + StructUncomp
func TestCheckBinaryNonConstExprStructAddStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structT and eval.structUncompT)`,
	)

}

// Test Struct & ConstInt
func TestCheckBinaryNonConstExprStructAndConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 4`, env,
		`invalid operation: a & 4 (mismatched types eval.structT and int)`,
	)

}

// Test Struct & ConstRune
func TestCheckBinaryNonConstExprStructAndConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & '@'`, env,
		`invalid operation: a & rune(64) (mismatched types eval.structT and rune)`,
	)

}

// Test Struct & ConstFloat
func TestCheckBinaryNonConstExprStructAndConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 2.0`, env,
		`invalid operation: a & 2 (mismatched types eval.structT and float64)`,
	)

}

// Test Struct & ConstComplex
func TestCheckBinaryNonConstExprStructAndConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 8.0i`, env,
		`invalid operation: a & 8i (mismatched types eval.structT and complex128)`,
	)

}

// Test Struct & ConstBool
func TestCheckBinaryNonConstExprStructAndConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & true`, env,
		`invalid operation: a & true (mismatched types eval.structT and bool)`,
	)

}

// Test Struct & ConstString
func TestCheckBinaryNonConstExprStructAndConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & "abc"`, env,
		`invalid operation: a & "abc" (mismatched types eval.structT and string)`,
	)

}

// Test Struct & ConstNil
func TestCheckBinaryNonConstExprStructAndConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & nil`, env,
		`cannot convert nil to type eval.structT`,
	)

}

// Test Struct & Int
func TestCheckBinaryNonConstExprStructAndInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structT and int)`,
	)

}

// Test Struct & Float32
func TestCheckBinaryNonConstExprStructAndFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structT and float32)`,
	)

}

// Test Struct & Complex128
func TestCheckBinaryNonConstExprStructAndComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structT and complex128)`,
	)

}

// Test Struct & String
func TestCheckBinaryNonConstExprStructAndString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structT and string)`,
	)

}

// Test Struct & BoolT
func TestCheckBinaryNonConstExprStructAndBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structT and bool)`,
	)

}

// Test Struct & Slice
func TestCheckBinaryNonConstExprStructAndSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structT and eval.sliceT)`,
	)

}

// Test Struct & Array
func TestCheckBinaryNonConstExprStructAndArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structT and eval.arrayT)`,
	)

}

// Test Struct & Map
func TestCheckBinaryNonConstExprStructAndMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structT and eval.mapT)`,
	)

}

// Test Struct & XinterfaceX
func TestCheckBinaryNonConstExprStructAndXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structT and eval.XinterfaceX)`,
	)

}

// Test Struct & InterfaceX
func TestCheckBinaryNonConstExprStructAndInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structT and eval.interfaceX)`,
	)

}

// Test Struct & InterfaceY
func TestCheckBinaryNonConstExprStructAndInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structT and eval.interfaceY)`,
	)

}

// Test Struct & InterfaceZ
func TestCheckBinaryNonConstExprStructAndInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structT and eval.interfaceZ)`,
	)

}

// Test Struct & Ptr
func TestCheckBinaryNonConstExprStructAndPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structT and *int)`,
	)

}

// Test Struct & Struct
func TestCheckBinaryNonConstExprStructAndStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (operator & not defined on struct)`,
	)

}

// Test Struct & StructUncomp
func TestCheckBinaryNonConstExprStructAndStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structT and eval.structUncompT)`,
	)

}

// Test Struct % ConstInt
func TestCheckBinaryNonConstExprStructRemConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 4`, env,
		`invalid operation: a % 4 (mismatched types eval.structT and int)`,
	)

}

// Test Struct % ConstRune
func TestCheckBinaryNonConstExprStructRemConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % '@'`, env,
		`invalid operation: a % rune(64) (mismatched types eval.structT and rune)`,
	)

}

// Test Struct % ConstFloat
func TestCheckBinaryNonConstExprStructRemConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 2.0`, env,
		`invalid operation: a % 2 (mismatched types eval.structT and float64)`,
	)

}

// Test Struct % ConstComplex
func TestCheckBinaryNonConstExprStructRemConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 8.0i`, env,
		`invalid operation: a % 8i (mismatched types eval.structT and complex128)`,
	)

}

// Test Struct % ConstBool
func TestCheckBinaryNonConstExprStructRemConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % true`, env,
		`invalid operation: a % true (mismatched types eval.structT and bool)`,
	)

}

// Test Struct % ConstString
func TestCheckBinaryNonConstExprStructRemConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % "abc"`, env,
		`invalid operation: a % "abc" (mismatched types eval.structT and string)`,
	)

}

// Test Struct % ConstNil
func TestCheckBinaryNonConstExprStructRemConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % nil`, env,
		`cannot convert nil to type eval.structT`,
	)

}

// Test Struct % Int
func TestCheckBinaryNonConstExprStructRemInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structT and int)`,
	)

}

// Test Struct % Float32
func TestCheckBinaryNonConstExprStructRemFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structT and float32)`,
	)

}

// Test Struct % Complex128
func TestCheckBinaryNonConstExprStructRemComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structT and complex128)`,
	)

}

// Test Struct % String
func TestCheckBinaryNonConstExprStructRemString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structT and string)`,
	)

}

// Test Struct % BoolT
func TestCheckBinaryNonConstExprStructRemBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structT and bool)`,
	)

}

// Test Struct % Slice
func TestCheckBinaryNonConstExprStructRemSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structT and eval.sliceT)`,
	)

}

// Test Struct % Array
func TestCheckBinaryNonConstExprStructRemArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structT and eval.arrayT)`,
	)

}

// Test Struct % Map
func TestCheckBinaryNonConstExprStructRemMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structT and eval.mapT)`,
	)

}

// Test Struct % XinterfaceX
func TestCheckBinaryNonConstExprStructRemXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structT and eval.XinterfaceX)`,
	)

}

// Test Struct % InterfaceX
func TestCheckBinaryNonConstExprStructRemInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structT and eval.interfaceX)`,
	)

}

// Test Struct % InterfaceY
func TestCheckBinaryNonConstExprStructRemInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structT and eval.interfaceY)`,
	)

}

// Test Struct % InterfaceZ
func TestCheckBinaryNonConstExprStructRemInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structT and eval.interfaceZ)`,
	)

}

// Test Struct % Ptr
func TestCheckBinaryNonConstExprStructRemPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structT and *int)`,
	)

}

// Test Struct % Struct
func TestCheckBinaryNonConstExprStructRemStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (operator % not defined on struct)`,
	)

}

// Test Struct % StructUncomp
func TestCheckBinaryNonConstExprStructRemStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structT and eval.structUncompT)`,
	)

}

// Test Struct == ConstInt
func TestCheckBinaryNonConstExprStructEqlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 4`, env,
		`invalid operation: a == 4 (mismatched types eval.structT and int)`,
	)

}

// Test Struct == ConstRune
func TestCheckBinaryNonConstExprStructEqlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == '@'`, env,
		`invalid operation: a == rune(64) (mismatched types eval.structT and rune)`,
	)

}

// Test Struct == ConstFloat
func TestCheckBinaryNonConstExprStructEqlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 2.0`, env,
		`invalid operation: a == 2 (mismatched types eval.structT and float64)`,
	)

}

// Test Struct == ConstComplex
func TestCheckBinaryNonConstExprStructEqlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 8.0i`, env,
		`invalid operation: a == 8i (mismatched types eval.structT and complex128)`,
	)

}

// Test Struct == ConstBool
func TestCheckBinaryNonConstExprStructEqlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == true`, env,
		`invalid operation: a == true (mismatched types eval.structT and bool)`,
	)

}

// Test Struct == ConstString
func TestCheckBinaryNonConstExprStructEqlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == "abc"`, env,
		`invalid operation: a == "abc" (mismatched types eval.structT and string)`,
	)

}

// Test Struct == ConstNil
func TestCheckBinaryNonConstExprStructEqlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == nil`, env,
		`cannot convert nil to type eval.structT`,
	)

}

// Test Struct == Int
func TestCheckBinaryNonConstExprStructEqlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structT and int)`,
	)

}

// Test Struct == Float32
func TestCheckBinaryNonConstExprStructEqlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structT and float32)`,
	)

}

// Test Struct == Complex128
func TestCheckBinaryNonConstExprStructEqlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structT and complex128)`,
	)

}

// Test Struct == String
func TestCheckBinaryNonConstExprStructEqlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structT and string)`,
	)

}

// Test Struct == BoolT
func TestCheckBinaryNonConstExprStructEqlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structT and bool)`,
	)

}

// Test Struct == Slice
func TestCheckBinaryNonConstExprStructEqlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structT and eval.sliceT)`,
	)

}

// Test Struct == Array
func TestCheckBinaryNonConstExprStructEqlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structT and eval.arrayT)`,
	)

}

// Test Struct == Map
func TestCheckBinaryNonConstExprStructEqlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structT and eval.mapT)`,
	)

}

// Test Struct == XinterfaceX
func TestCheckBinaryNonConstExprStructEqlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structT and eval.XinterfaceX)`,
	)

}

// Test Struct == InterfaceX
func TestCheckBinaryNonConstExprStructEqlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structT and eval.interfaceX)`,
	)

}

// Test Struct == InterfaceY
func TestCheckBinaryNonConstExprStructEqlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structT and eval.interfaceY)`,
	)

}

// Test Struct == InterfaceZ
func TestCheckBinaryNonConstExprStructEqlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structT and eval.interfaceZ)`,
	)

}

// Test Struct == Ptr
func TestCheckBinaryNonConstExprStructEqlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structT and *int)`,
	)

}

// Test Struct == Struct
func TestCheckBinaryNonConstExprStructEqlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (struct containing []int cannot be compared)`,
	)

}

// Test Struct == StructUncomp
func TestCheckBinaryNonConstExprStructEqlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structT and eval.structUncompT)`,
	)

}

// Test Struct > ConstInt
func TestCheckBinaryNonConstExprStructGtrConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 4`, env,
		`invalid operation: a > 4 (mismatched types eval.structT and int)`,
	)

}

// Test Struct > ConstRune
func TestCheckBinaryNonConstExprStructGtrConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > '@'`, env,
		`invalid operation: a > rune(64) (mismatched types eval.structT and rune)`,
	)

}

// Test Struct > ConstFloat
func TestCheckBinaryNonConstExprStructGtrConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 2.0`, env,
		`invalid operation: a > 2 (mismatched types eval.structT and float64)`,
	)

}

// Test Struct > ConstComplex
func TestCheckBinaryNonConstExprStructGtrConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 8.0i`, env,
		`invalid operation: a > 8i (mismatched types eval.structT and complex128)`,
	)

}

// Test Struct > ConstBool
func TestCheckBinaryNonConstExprStructGtrConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > true`, env,
		`invalid operation: a > true (mismatched types eval.structT and bool)`,
	)

}

// Test Struct > ConstString
func TestCheckBinaryNonConstExprStructGtrConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > "abc"`, env,
		`invalid operation: a > "abc" (mismatched types eval.structT and string)`,
	)

}

// Test Struct > ConstNil
func TestCheckBinaryNonConstExprStructGtrConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > nil`, env,
		`cannot convert nil to type eval.structT`,
	)

}

// Test Struct > Int
func TestCheckBinaryNonConstExprStructGtrInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structT and int)`,
	)

}

// Test Struct > Float32
func TestCheckBinaryNonConstExprStructGtrFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structT and float32)`,
	)

}

// Test Struct > Complex128
func TestCheckBinaryNonConstExprStructGtrComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structT and complex128)`,
	)

}

// Test Struct > String
func TestCheckBinaryNonConstExprStructGtrString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structT and string)`,
	)

}

// Test Struct > BoolT
func TestCheckBinaryNonConstExprStructGtrBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structT and bool)`,
	)

}

// Test Struct > Slice
func TestCheckBinaryNonConstExprStructGtrSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structT and eval.sliceT)`,
	)

}

// Test Struct > Array
func TestCheckBinaryNonConstExprStructGtrArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structT and eval.arrayT)`,
	)

}

// Test Struct > Map
func TestCheckBinaryNonConstExprStructGtrMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structT and eval.mapT)`,
	)

}

// Test Struct > XinterfaceX
func TestCheckBinaryNonConstExprStructGtrXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structT and eval.XinterfaceX)`,
	)

}

// Test Struct > InterfaceX
func TestCheckBinaryNonConstExprStructGtrInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structT and eval.interfaceX)`,
	)

}

// Test Struct > InterfaceY
func TestCheckBinaryNonConstExprStructGtrInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structT and eval.interfaceY)`,
	)

}

// Test Struct > InterfaceZ
func TestCheckBinaryNonConstExprStructGtrInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structT and eval.interfaceZ)`,
	)

}

// Test Struct > Ptr
func TestCheckBinaryNonConstExprStructGtrPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structT and *int)`,
	)

}

// Test Struct > Struct
func TestCheckBinaryNonConstExprStructGtrStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (operator > not defined on struct)`,
	)

}

// Test Struct > StructUncomp
func TestCheckBinaryNonConstExprStructGtrStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structT and eval.structUncompT)`,
	)

}

// Test Struct << ConstInt
func TestCheckBinaryNonConstExprStructShlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 4`, env,
		`invalid operation: a << 4 (shift of type eval.structT)`,
	)

}

// Test Struct << ConstRune
func TestCheckBinaryNonConstExprStructShlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << '@'`, env,
		`invalid operation: a << 64 (shift of type eval.structT)`,
	)

}

// Test Struct << ConstFloat
func TestCheckBinaryNonConstExprStructShlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 2.0`, env,
		`invalid operation: a << 2 (shift of type eval.structT)`,
	)

}

// Test Struct << ConstComplex
func TestCheckBinaryNonConstExprStructShlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: a << 0 (shift of type eval.structT)`,
	)

}

// Test Struct << ConstBool
func TestCheckBinaryNonConstExprStructShlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << true`, env,
		`invalid operation: a << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test Struct << ConstString
func TestCheckBinaryNonConstExprStructShlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: a << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test Struct << ConstNil
func TestCheckBinaryNonConstExprStructShlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test Struct << Int
func TestCheckBinaryNonConstExprStructShlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type int, must be unsigned integer)`,
	)

}

// Test Struct << Float32
func TestCheckBinaryNonConstExprStructShlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type float32, must be unsigned integer)`,
	)

}

// Test Struct << Complex128
func TestCheckBinaryNonConstExprStructShlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type complex128, must be unsigned integer)`,
	)

}

// Test Struct << String
func TestCheckBinaryNonConstExprStructShlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type string, must be unsigned integer)`,
	)

}

// Test Struct << BoolT
func TestCheckBinaryNonConstExprStructShlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type bool, must be unsigned integer)`,
	)

}

// Test Struct << Slice
func TestCheckBinaryNonConstExprStructShlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.sliceT, must be unsigned integer)`,
	)

}

// Test Struct << Array
func TestCheckBinaryNonConstExprStructShlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.arrayT, must be unsigned integer)`,
	)

}

// Test Struct << Map
func TestCheckBinaryNonConstExprStructShlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.mapT, must be unsigned integer)`,
	)

}

// Test Struct << XinterfaceX
func TestCheckBinaryNonConstExprStructShlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.XinterfaceX, must be unsigned integer)`,
	)

}

// Test Struct << InterfaceX
func TestCheckBinaryNonConstExprStructShlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceX, must be unsigned integer)`,
	)

}

// Test Struct << InterfaceY
func TestCheckBinaryNonConstExprStructShlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceY, must be unsigned integer)`,
	)

}

// Test Struct << InterfaceZ
func TestCheckBinaryNonConstExprStructShlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceZ, must be unsigned integer)`,
	)

}

// Test Struct << Ptr
func TestCheckBinaryNonConstExprStructShlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type *int, must be unsigned integer)`,
	)

}

// Test Struct << Struct
func TestCheckBinaryNonConstExprStructShlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structT, must be unsigned integer)`,
	)

}

// Test Struct << StructUncomp
func TestCheckBinaryNonConstExprStructShlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structUncompT, must be unsigned integer)`,
	)

}

// Test StructUncomp + ConstInt
func TestCheckBinaryNonConstExprStructUncompAddConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 4`, env,
		`invalid operation: a + 4 (mismatched types eval.structUncompT and int)`,
	)

}

// Test StructUncomp + ConstRune
func TestCheckBinaryNonConstExprStructUncompAddConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + '@'`, env,
		`invalid operation: a + rune(64) (mismatched types eval.structUncompT and rune)`,
	)

}

// Test StructUncomp + ConstFloat
func TestCheckBinaryNonConstExprStructUncompAddConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 2.0`, env,
		`invalid operation: a + 2 (mismatched types eval.structUncompT and float64)`,
	)

}

// Test StructUncomp + ConstComplex
func TestCheckBinaryNonConstExprStructUncompAddConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + 8.0i`, env,
		`invalid operation: a + 8i (mismatched types eval.structUncompT and complex128)`,
	)

}

// Test StructUncomp + ConstBool
func TestCheckBinaryNonConstExprStructUncompAddConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + true`, env,
		`invalid operation: a + true (mismatched types eval.structUncompT and bool)`,
	)

}

// Test StructUncomp + ConstString
func TestCheckBinaryNonConstExprStructUncompAddConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + "abc"`, env,
		`invalid operation: a + "abc" (mismatched types eval.structUncompT and string)`,
	)

}

// Test StructUncomp + ConstNil
func TestCheckBinaryNonConstExprStructUncompAddConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a + nil`, env,
		`cannot convert nil to type eval.structUncompT`,
	)

}

// Test StructUncomp + Int
func TestCheckBinaryNonConstExprStructUncompAddInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structUncompT and int)`,
	)

}

// Test StructUncomp + Float32
func TestCheckBinaryNonConstExprStructUncompAddFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structUncompT and float32)`,
	)

}

// Test StructUncomp + Complex128
func TestCheckBinaryNonConstExprStructUncompAddComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structUncompT and complex128)`,
	)

}

// Test StructUncomp + String
func TestCheckBinaryNonConstExprStructUncompAddString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structUncompT and string)`,
	)

}

// Test StructUncomp + BoolT
func TestCheckBinaryNonConstExprStructUncompAddBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structUncompT and bool)`,
	)

}

// Test StructUncomp + Slice
func TestCheckBinaryNonConstExprStructUncompAddSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structUncompT and eval.sliceT)`,
	)

}

// Test StructUncomp + Array
func TestCheckBinaryNonConstExprStructUncompAddArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structUncompT and eval.arrayT)`,
	)

}

// Test StructUncomp + Map
func TestCheckBinaryNonConstExprStructUncompAddMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structUncompT and eval.mapT)`,
	)

}

// Test StructUncomp + XinterfaceX
func TestCheckBinaryNonConstExprStructUncompAddXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structUncompT and eval.XinterfaceX)`,
	)

}

// Test StructUncomp + InterfaceX
func TestCheckBinaryNonConstExprStructUncompAddInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structUncompT and eval.interfaceX)`,
	)

}

// Test StructUncomp + InterfaceY
func TestCheckBinaryNonConstExprStructUncompAddInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structUncompT and eval.interfaceY)`,
	)

}

// Test StructUncomp + InterfaceZ
func TestCheckBinaryNonConstExprStructUncompAddInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structUncompT and eval.interfaceZ)`,
	)

}

// Test StructUncomp + Ptr
func TestCheckBinaryNonConstExprStructUncompAddPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structUncompT and *int)`,
	)

}

// Test StructUncomp + Struct
func TestCheckBinaryNonConstExprStructUncompAddStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (mismatched types eval.structUncompT and eval.structT)`,
	)

}

// Test StructUncomp + StructUncomp
func TestCheckBinaryNonConstExprStructUncompAddStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a + b`, env,
		`invalid operation: a + b (operator + not defined on struct)`,
	)

}

// Test StructUncomp & ConstInt
func TestCheckBinaryNonConstExprStructUncompAndConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 4`, env,
		`invalid operation: a & 4 (mismatched types eval.structUncompT and int)`,
	)

}

// Test StructUncomp & ConstRune
func TestCheckBinaryNonConstExprStructUncompAndConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & '@'`, env,
		`invalid operation: a & rune(64) (mismatched types eval.structUncompT and rune)`,
	)

}

// Test StructUncomp & ConstFloat
func TestCheckBinaryNonConstExprStructUncompAndConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 2.0`, env,
		`invalid operation: a & 2 (mismatched types eval.structUncompT and float64)`,
	)

}

// Test StructUncomp & ConstComplex
func TestCheckBinaryNonConstExprStructUncompAndConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & 8.0i`, env,
		`invalid operation: a & 8i (mismatched types eval.structUncompT and complex128)`,
	)

}

// Test StructUncomp & ConstBool
func TestCheckBinaryNonConstExprStructUncompAndConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & true`, env,
		`invalid operation: a & true (mismatched types eval.structUncompT and bool)`,
	)

}

// Test StructUncomp & ConstString
func TestCheckBinaryNonConstExprStructUncompAndConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & "abc"`, env,
		`invalid operation: a & "abc" (mismatched types eval.structUncompT and string)`,
	)

}

// Test StructUncomp & ConstNil
func TestCheckBinaryNonConstExprStructUncompAndConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a & nil`, env,
		`cannot convert nil to type eval.structUncompT`,
	)

}

// Test StructUncomp & Int
func TestCheckBinaryNonConstExprStructUncompAndInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structUncompT and int)`,
	)

}

// Test StructUncomp & Float32
func TestCheckBinaryNonConstExprStructUncompAndFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structUncompT and float32)`,
	)

}

// Test StructUncomp & Complex128
func TestCheckBinaryNonConstExprStructUncompAndComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structUncompT and complex128)`,
	)

}

// Test StructUncomp & String
func TestCheckBinaryNonConstExprStructUncompAndString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structUncompT and string)`,
	)

}

// Test StructUncomp & BoolT
func TestCheckBinaryNonConstExprStructUncompAndBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structUncompT and bool)`,
	)

}

// Test StructUncomp & Slice
func TestCheckBinaryNonConstExprStructUncompAndSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structUncompT and eval.sliceT)`,
	)

}

// Test StructUncomp & Array
func TestCheckBinaryNonConstExprStructUncompAndArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structUncompT and eval.arrayT)`,
	)

}

// Test StructUncomp & Map
func TestCheckBinaryNonConstExprStructUncompAndMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structUncompT and eval.mapT)`,
	)

}

// Test StructUncomp & XinterfaceX
func TestCheckBinaryNonConstExprStructUncompAndXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structUncompT and eval.XinterfaceX)`,
	)

}

// Test StructUncomp & InterfaceX
func TestCheckBinaryNonConstExprStructUncompAndInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structUncompT and eval.interfaceX)`,
	)

}

// Test StructUncomp & InterfaceY
func TestCheckBinaryNonConstExprStructUncompAndInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structUncompT and eval.interfaceY)`,
	)

}

// Test StructUncomp & InterfaceZ
func TestCheckBinaryNonConstExprStructUncompAndInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structUncompT and eval.interfaceZ)`,
	)

}

// Test StructUncomp & Ptr
func TestCheckBinaryNonConstExprStructUncompAndPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structUncompT and *int)`,
	)

}

// Test StructUncomp & Struct
func TestCheckBinaryNonConstExprStructUncompAndStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (mismatched types eval.structUncompT and eval.structT)`,
	)

}

// Test StructUncomp & StructUncomp
func TestCheckBinaryNonConstExprStructUncompAndStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a & b`, env,
		`invalid operation: a & b (operator & not defined on struct)`,
	)

}

// Test StructUncomp % ConstInt
func TestCheckBinaryNonConstExprStructUncompRemConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 4`, env,
		`invalid operation: a % 4 (mismatched types eval.structUncompT and int)`,
	)

}

// Test StructUncomp % ConstRune
func TestCheckBinaryNonConstExprStructUncompRemConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % '@'`, env,
		`invalid operation: a % rune(64) (mismatched types eval.structUncompT and rune)`,
	)

}

// Test StructUncomp % ConstFloat
func TestCheckBinaryNonConstExprStructUncompRemConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 2.0`, env,
		`invalid operation: a % 2 (mismatched types eval.structUncompT and float64)`,
	)

}

// Test StructUncomp % ConstComplex
func TestCheckBinaryNonConstExprStructUncompRemConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % 8.0i`, env,
		`invalid operation: a % 8i (mismatched types eval.structUncompT and complex128)`,
	)

}

// Test StructUncomp % ConstBool
func TestCheckBinaryNonConstExprStructUncompRemConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % true`, env,
		`invalid operation: a % true (mismatched types eval.structUncompT and bool)`,
	)

}

// Test StructUncomp % ConstString
func TestCheckBinaryNonConstExprStructUncompRemConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % "abc"`, env,
		`invalid operation: a % "abc" (mismatched types eval.structUncompT and string)`,
	)

}

// Test StructUncomp % ConstNil
func TestCheckBinaryNonConstExprStructUncompRemConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a % nil`, env,
		`cannot convert nil to type eval.structUncompT`,
	)

}

// Test StructUncomp % Int
func TestCheckBinaryNonConstExprStructUncompRemInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structUncompT and int)`,
	)

}

// Test StructUncomp % Float32
func TestCheckBinaryNonConstExprStructUncompRemFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structUncompT and float32)`,
	)

}

// Test StructUncomp % Complex128
func TestCheckBinaryNonConstExprStructUncompRemComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structUncompT and complex128)`,
	)

}

// Test StructUncomp % String
func TestCheckBinaryNonConstExprStructUncompRemString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structUncompT and string)`,
	)

}

// Test StructUncomp % BoolT
func TestCheckBinaryNonConstExprStructUncompRemBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structUncompT and bool)`,
	)

}

// Test StructUncomp % Slice
func TestCheckBinaryNonConstExprStructUncompRemSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structUncompT and eval.sliceT)`,
	)

}

// Test StructUncomp % Array
func TestCheckBinaryNonConstExprStructUncompRemArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structUncompT and eval.arrayT)`,
	)

}

// Test StructUncomp % Map
func TestCheckBinaryNonConstExprStructUncompRemMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structUncompT and eval.mapT)`,
	)

}

// Test StructUncomp % XinterfaceX
func TestCheckBinaryNonConstExprStructUncompRemXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structUncompT and eval.XinterfaceX)`,
	)

}

// Test StructUncomp % InterfaceX
func TestCheckBinaryNonConstExprStructUncompRemInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structUncompT and eval.interfaceX)`,
	)

}

// Test StructUncomp % InterfaceY
func TestCheckBinaryNonConstExprStructUncompRemInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structUncompT and eval.interfaceY)`,
	)

}

// Test StructUncomp % InterfaceZ
func TestCheckBinaryNonConstExprStructUncompRemInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structUncompT and eval.interfaceZ)`,
	)

}

// Test StructUncomp % Ptr
func TestCheckBinaryNonConstExprStructUncompRemPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structUncompT and *int)`,
	)

}

// Test StructUncomp % Struct
func TestCheckBinaryNonConstExprStructUncompRemStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (mismatched types eval.structUncompT and eval.structT)`,
	)

}

// Test StructUncomp % StructUncomp
func TestCheckBinaryNonConstExprStructUncompRemStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a % b`, env,
		`invalid operation: a % b (operator % not defined on struct)`,
	)

}

// Test StructUncomp == ConstInt
func TestCheckBinaryNonConstExprStructUncompEqlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 4`, env,
		`invalid operation: a == 4 (mismatched types eval.structUncompT and int)`,
	)

}

// Test StructUncomp == ConstRune
func TestCheckBinaryNonConstExprStructUncompEqlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == '@'`, env,
		`invalid operation: a == rune(64) (mismatched types eval.structUncompT and rune)`,
	)

}

// Test StructUncomp == ConstFloat
func TestCheckBinaryNonConstExprStructUncompEqlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 2.0`, env,
		`invalid operation: a == 2 (mismatched types eval.structUncompT and float64)`,
	)

}

// Test StructUncomp == ConstComplex
func TestCheckBinaryNonConstExprStructUncompEqlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == 8.0i`, env,
		`invalid operation: a == 8i (mismatched types eval.structUncompT and complex128)`,
	)

}

// Test StructUncomp == ConstBool
func TestCheckBinaryNonConstExprStructUncompEqlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == true`, env,
		`invalid operation: a == true (mismatched types eval.structUncompT and bool)`,
	)

}

// Test StructUncomp == ConstString
func TestCheckBinaryNonConstExprStructUncompEqlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == "abc"`, env,
		`invalid operation: a == "abc" (mismatched types eval.structUncompT and string)`,
	)

}

// Test StructUncomp == ConstNil
func TestCheckBinaryNonConstExprStructUncompEqlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a == nil`, env,
		`cannot convert nil to type eval.structUncompT`,
	)

}

// Test StructUncomp == Int
func TestCheckBinaryNonConstExprStructUncompEqlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structUncompT and int)`,
	)

}

// Test StructUncomp == Float32
func TestCheckBinaryNonConstExprStructUncompEqlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structUncompT and float32)`,
	)

}

// Test StructUncomp == Complex128
func TestCheckBinaryNonConstExprStructUncompEqlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structUncompT and complex128)`,
	)

}

// Test StructUncomp == String
func TestCheckBinaryNonConstExprStructUncompEqlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structUncompT and string)`,
	)

}

// Test StructUncomp == BoolT
func TestCheckBinaryNonConstExprStructUncompEqlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structUncompT and bool)`,
	)

}

// Test StructUncomp == Slice
func TestCheckBinaryNonConstExprStructUncompEqlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structUncompT and eval.sliceT)`,
	)

}

// Test StructUncomp == Array
func TestCheckBinaryNonConstExprStructUncompEqlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structUncompT and eval.arrayT)`,
	)

}

// Test StructUncomp == Map
func TestCheckBinaryNonConstExprStructUncompEqlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structUncompT and eval.mapT)`,
	)

}

// Test StructUncomp == XinterfaceX
func TestCheckBinaryNonConstExprStructUncompEqlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structUncompT and eval.XinterfaceX)`,
	)

}

// Test StructUncomp == InterfaceX
func TestCheckBinaryNonConstExprStructUncompEqlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structUncompT and eval.interfaceX)`,
	)

}

// Test StructUncomp == InterfaceY
func TestCheckBinaryNonConstExprStructUncompEqlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structUncompT and eval.interfaceY)`,
	)

}

// Test StructUncomp == InterfaceZ
func TestCheckBinaryNonConstExprStructUncompEqlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structUncompT and eval.interfaceZ)`,
	)

}

// Test StructUncomp == Ptr
func TestCheckBinaryNonConstExprStructUncompEqlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structUncompT and *int)`,
	)

}

// Test StructUncomp == Struct
func TestCheckBinaryNonConstExprStructUncompEqlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (mismatched types eval.structUncompT and eval.structT)`,
	)

}

// Test StructUncomp == StructUncomp
func TestCheckBinaryNonConstExprStructUncompEqlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a == b`, env,
		`invalid operation: a == b (struct containing []int cannot be compared)`,
	)

}

// Test StructUncomp > ConstInt
func TestCheckBinaryNonConstExprStructUncompGtrConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 4`, env,
		`invalid operation: a > 4 (mismatched types eval.structUncompT and int)`,
	)

}

// Test StructUncomp > ConstRune
func TestCheckBinaryNonConstExprStructUncompGtrConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > '@'`, env,
		`invalid operation: a > rune(64) (mismatched types eval.structUncompT and rune)`,
	)

}

// Test StructUncomp > ConstFloat
func TestCheckBinaryNonConstExprStructUncompGtrConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 2.0`, env,
		`invalid operation: a > 2 (mismatched types eval.structUncompT and float64)`,
	)

}

// Test StructUncomp > ConstComplex
func TestCheckBinaryNonConstExprStructUncompGtrConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > 8.0i`, env,
		`invalid operation: a > 8i (mismatched types eval.structUncompT and complex128)`,
	)

}

// Test StructUncomp > ConstBool
func TestCheckBinaryNonConstExprStructUncompGtrConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > true`, env,
		`invalid operation: a > true (mismatched types eval.structUncompT and bool)`,
	)

}

// Test StructUncomp > ConstString
func TestCheckBinaryNonConstExprStructUncompGtrConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > "abc"`, env,
		`invalid operation: a > "abc" (mismatched types eval.structUncompT and string)`,
	)

}

// Test StructUncomp > ConstNil
func TestCheckBinaryNonConstExprStructUncompGtrConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a > nil`, env,
		`cannot convert nil to type eval.structUncompT`,
	)

}

// Test StructUncomp > Int
func TestCheckBinaryNonConstExprStructUncompGtrInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structUncompT and int)`,
	)

}

// Test StructUncomp > Float32
func TestCheckBinaryNonConstExprStructUncompGtrFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structUncompT and float32)`,
	)

}

// Test StructUncomp > Complex128
func TestCheckBinaryNonConstExprStructUncompGtrComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structUncompT and complex128)`,
	)

}

// Test StructUncomp > String
func TestCheckBinaryNonConstExprStructUncompGtrString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structUncompT and string)`,
	)

}

// Test StructUncomp > BoolT
func TestCheckBinaryNonConstExprStructUncompGtrBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structUncompT and bool)`,
	)

}

// Test StructUncomp > Slice
func TestCheckBinaryNonConstExprStructUncompGtrSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structUncompT and eval.sliceT)`,
	)

}

// Test StructUncomp > Array
func TestCheckBinaryNonConstExprStructUncompGtrArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structUncompT and eval.arrayT)`,
	)

}

// Test StructUncomp > Map
func TestCheckBinaryNonConstExprStructUncompGtrMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structUncompT and eval.mapT)`,
	)

}

// Test StructUncomp > XinterfaceX
func TestCheckBinaryNonConstExprStructUncompGtrXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structUncompT and eval.XinterfaceX)`,
	)

}

// Test StructUncomp > InterfaceX
func TestCheckBinaryNonConstExprStructUncompGtrInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structUncompT and eval.interfaceX)`,
	)

}

// Test StructUncomp > InterfaceY
func TestCheckBinaryNonConstExprStructUncompGtrInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structUncompT and eval.interfaceY)`,
	)

}

// Test StructUncomp > InterfaceZ
func TestCheckBinaryNonConstExprStructUncompGtrInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structUncompT and eval.interfaceZ)`,
	)

}

// Test StructUncomp > Ptr
func TestCheckBinaryNonConstExprStructUncompGtrPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structUncompT and *int)`,
	)

}

// Test StructUncomp > Struct
func TestCheckBinaryNonConstExprStructUncompGtrStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (mismatched types eval.structUncompT and eval.structT)`,
	)

}

// Test StructUncomp > StructUncomp
func TestCheckBinaryNonConstExprStructUncompGtrStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a > b`, env,
		`invalid operation: a > b (operator > not defined on struct)`,
	)

}

// Test StructUncomp << ConstInt
func TestCheckBinaryNonConstExprStructUncompShlConstInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 4`, env,
		`invalid operation: a << 4 (shift of type eval.structUncompT)`,
	)

}

// Test StructUncomp << ConstRune
func TestCheckBinaryNonConstExprStructUncompShlConstRune(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << '@'`, env,
		`invalid operation: a << 64 (shift of type eval.structUncompT)`,
	)

}

// Test StructUncomp << ConstFloat
func TestCheckBinaryNonConstExprStructUncompShlConstFloat(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 2.0`, env,
		`invalid operation: a << 2 (shift of type eval.structUncompT)`,
	)

}

// Test StructUncomp << ConstComplex
func TestCheckBinaryNonConstExprStructUncompShlConstComplex(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << 8.0i`, env,
		`constant 0+8i truncated to real`,
		`invalid operation: a << 0 (shift of type eval.structUncompT)`,
	)

}

// Test StructUncomp << ConstBool
func TestCheckBinaryNonConstExprStructUncompShlConstBool(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << true`, env,
		`invalid operation: a << true (shift count type bool, must be unsigned integer)`,
	)

}

// Test StructUncomp << ConstString
func TestCheckBinaryNonConstExprStructUncompShlConstString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << "abc"`, env,
		`cannot convert "abc" to type uint`,
		`invalid operation: a << "abc" (shift count type string, must be unsigned integer)`,
	)

}

// Test StructUncomp << ConstNil
func TestCheckBinaryNonConstExprStructUncompShlConstNil(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	
	expectCheckError(t, `a << nil`, env,
		`cannot convert nil to type uint`,
	)

}

// Test StructUncomp << Int
func TestCheckBinaryNonConstExprStructUncompShlInt(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := int(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type int, must be unsigned integer)`,
	)

}

// Test StructUncomp << Float32
func TestCheckBinaryNonConstExprStructUncompShlFloat32(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := float32(1.5); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type float32, must be unsigned integer)`,
	)

}

// Test StructUncomp << Complex128
func TestCheckBinaryNonConstExprStructUncompShlComplex128(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := complex128(1i); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type complex128, must be unsigned integer)`,
	)

}

// Test StructUncomp << String
func TestCheckBinaryNonConstExprStructUncompShlString(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := string("abc"); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type string, must be unsigned integer)`,
	)

}

// Test StructUncomp << BoolT
func TestCheckBinaryNonConstExprStructUncompShlBoolT(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := bool(true); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type bool, must be unsigned integer)`,
	)

}

// Test StructUncomp << Slice
func TestCheckBinaryNonConstExprStructUncompShlSlice(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := sliceT(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.sliceT, must be unsigned integer)`,
	)

}

// Test StructUncomp << Array
func TestCheckBinaryNonConstExprStructUncompShlArray(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := arrayT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.arrayT, must be unsigned integer)`,
	)

}

// Test StructUncomp << Map
func TestCheckBinaryNonConstExprStructUncompShlMap(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := mapT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.mapT, must be unsigned integer)`,
	)

}

// Test StructUncomp << XinterfaceX
func TestCheckBinaryNonConstExprStructUncompShlXinterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := XinterfaceX(1); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.XinterfaceX, must be unsigned integer)`,
	)

}

// Test StructUncomp << InterfaceX
func TestCheckBinaryNonConstExprStructUncompShlInterfaceX(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceX(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceX, must be unsigned integer)`,
	)

}

// Test StructUncomp << InterfaceY
func TestCheckBinaryNonConstExprStructUncompShlInterfaceY(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceY(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceY, must be unsigned integer)`,
	)

}

// Test StructUncomp << InterfaceZ
func TestCheckBinaryNonConstExprStructUncompShlInterfaceZ(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := interfaceZ(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.interfaceZ, must be unsigned integer)`,
	)

}

// Test StructUncomp << Ptr
func TestCheckBinaryNonConstExprStructUncompShlPtr(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := (*int)(nil); env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type *int, must be unsigned integer)`,
	)

}

// Test StructUncomp << Struct
func TestCheckBinaryNonConstExprStructUncompShlStruct(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structT, must be unsigned integer)`,
	)

}

// Test StructUncomp << StructUncomp
func TestCheckBinaryNonConstExprStructUncompShlStructUncomp(t *testing.T) {
	env := makeCheckBinaryNonConstExprEnv()
	a := structUncompT{}; env.Vars["a"] = reflect.ValueOf(&a)
	b := structUncompT{}; env.Vars["b"] = reflect.ValueOf(&b)
	expectCheckError(t, `a << b`, env,
		`invalid operation: a << b (shift count type eval.structUncompT, must be unsigned integer)`,
	)

}
