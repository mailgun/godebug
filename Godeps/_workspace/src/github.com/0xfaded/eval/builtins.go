package eval

import (
	"errors"
	"reflect"
)

var (
	intType reflect.Type = reflect.TypeOf(int(0))
	i8 reflect.Type = reflect.TypeOf(int8(0))
	i16 reflect.Type = reflect.TypeOf(int16(0))
	i32 reflect.Type = reflect.TypeOf(int32(0))
	i64 reflect.Type = reflect.TypeOf(int64(0))

	uintType reflect.Type = reflect.TypeOf(uint(0))
	u8 reflect.Type = reflect.TypeOf(uint8(0))
	u16 reflect.Type = reflect.TypeOf(uint16(0))
	u32 reflect.Type = reflect.TypeOf(uint32(0))
	u64 reflect.Type = reflect.TypeOf(uint64(0))

	f32 reflect.Type = reflect.TypeOf(float32(0))
	f64 reflect.Type = reflect.TypeOf(float64(0))
	c64 reflect.Type = reflect.TypeOf(complex64(0))
	c128 reflect.Type = reflect.TypeOf(complex128(0))

	boolType reflect.Type = reflect.TypeOf(bool(false))
	stringType reflect.Type = reflect.TypeOf(string(""))

	emptyInterface reflect.Type = reflect.TypeOf(new(interface{})).Elem()

	byteSlice reflect.Type = reflect.SliceOf(u8)
	runeSlice reflect.Type = reflect.SliceOf(i32)
	untypedNilType reflect.Type = reflect.TypeOf(UntypedNil{})
)

var builtinTypes = map[string] reflect.Type{
	"int": intType,
	"int8": i8,
	"int16": i16,
	"int32": i32,
	"int64": i64,

	"uint": uintType,
	"uint8": u8,
	"uint16": u16,
	"uint32": u32,
	"uint64": u64,

	"float32": f32,
	"float64": f64,

	"complex64": c64,
	"complex128": c128,

	"bool": boolType,
	"byte": ByteType,
	"rune": RuneType,
	"string": stringType,

	"error": reflect.TypeOf(errors.New("")),
}

var builtinFuncs = map[string] reflect.Value{
	"complex": reflect.ValueOf(builtinComplex),
	"real": reflect.ValueOf(builtinReal),
	"imag": reflect.ValueOf(builtinImag),
	"append": reflect.ValueOf(builtinAppend),
	"cap": reflect.ValueOf(builtinCap),
	"len": reflect.ValueOf(builtinLen),
	"new": reflect.ValueOf(builtinNew),
	"copy": reflect.ValueOf(builtinCopy),
	"delete": reflect.ValueOf(builtinDelete),
	"panic": reflect.ValueOf(builtinPanic),
}

func builtinComplex(re, im reflect.Value) reflect.Value {
	if re.Type() == f64 {
		return reflect.ValueOf(complex128(complex(re.Float(), im.Float())))
	} else {
		return reflect.ValueOf(complex64(complex(re.Float(), im.Float())))
	}
}

func builtinReal(cplx reflect.Value) reflect.Value {
	if cplx.Type() == c128 {
		return reflect.ValueOf(float64(real(cplx.Complex())))
	} else {
		return reflect.ValueOf(float32(real(cplx.Complex())))
	}
}

func builtinImag(cplx reflect.Value) reflect.Value {
	if cplx.Type() == c128 {
		return reflect.ValueOf(float64(imag(cplx.Complex())))
	} else {
		return reflect.ValueOf(float32(imag(cplx.Complex())))
	}
}

func builtinAppend(s, t reflect.Value) reflect.Value {
	if s.Type() == byteSlice && t.Type().Kind() == reflect.String {
		t = reflect.ValueOf([]byte(t.String()))
	}
	return reflect.AppendSlice(s, t)
}

func builtinLen(v reflect.Value) reflect.Value {
	return reflect.ValueOf(v.Len())
}

func builtinCap(v reflect.Value) reflect.Value {
	return reflect.ValueOf(v.Cap())
}

func builtinNew(t reflect.Type) reflect.Value {
	return reflect.New(t)
}

func builtinCopy(s, t reflect.Value) reflect.Value {
	n := reflect.Copy(s, t)
	return reflect.ValueOf(n)
}

func builtinDelete(m, k reflect.Value) reflect.Value {
	m.SetMapIndex(k, reflect.Value{})
	return reflect.Value{}
}

func builtinPanic(i reflect.Value) error {
	return PanicUser(i)
}
