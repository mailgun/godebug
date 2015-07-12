package eval

import (
	"reflect"
	"testing"
)

func TestIntBinaryOps(t *testing.T) {
	x := int32(5)
	env := MakeSimpleEnv()
	env.Vars["x"] = reflect.ValueOf(&x)

	expectResult(t, "x+2",   env, x+2)
	expectResult(t, "x-2",   env, x-2)
	expectResult(t, "x*3.0",   env, x*3.0)
	expectResult(t, "x/2",   env, x/2)
	expectResult(t, "x%2",   env, x%2)
	expectResult(t, "x&1",   env, x&1)
	expectResult(t, "x|1",   env, x|1)
	expectResult(t, "x^1",   env, x^1)
	expectResult(t, "x&^1",  env, x&^1)
	expectResult(t, "x+x",  env, x+x)
	expectResult(t, "x>>1",  env, x>>1)
	expectResult(t, "x<<1",  env, x<<1)

	expectResult(t, "x<1",   env, bool(x<1))
	expectResult(t, "-x<=3",  env, bool(-x<=3))
	expectResult(t, "x>1",   env, bool(x>1))
	expectResult(t, "x>=3",   env, bool(x>=3))
	expectResult(t, "-x==-5", env, bool(-x==-5))
	expectResult(t, "-x==3", env, bool(x==3))
	expectResult(t, "x!=1",  env, bool(x!=1))
}

func TestUintBinaryOps(t *testing.T) {
	x := uint32(5)
	env := MakeSimpleEnv()
	env.Vars["x"] = reflect.ValueOf(&x)

	expectResult(t, "x+2",   env, x+2)
	expectResult(t, "x-2",   env, x-2)
	expectResult(t, "x*3.0",   env, x*3.0)
	expectResult(t, "x/2",   env, x/2)
	expectResult(t, "x%2",   env, x%2)
	expectResult(t, "x&1",   env, x&1)
	expectResult(t, "x|1",   env, x|1)
	expectResult(t, "x^1",   env, x^1)
	expectResult(t, "x&^1",  env, x&^1)
	expectResult(t, "x+x",  env, x+x)
	expectResult(t, "x>>1",  env, x>>1)
	expectResult(t, "x<<1",  env, x<<1)

	expectResult(t, "x<1",   env, bool(x<1))
	expectResult(t, "-x<=3",  env, bool(-x<=3))
	expectResult(t, "x>1",   env, bool(x>1))
	expectResult(t, "x>=3",   env, bool(x>=3))
	expectResult(t, "-x==3", env, bool(x==3))
	expectResult(t, "x!=1",  env, bool(x!=1))
}

func TestFloatBinaryOps(t *testing.T) {
	x := float32(2.25)
	env := MakeSimpleEnv()
	env.Vars["x"] = reflect.ValueOf(&x)

	expectResult(t, "x+2",   env, x+2)
	expectResult(t, "x-2",   env, x-2)
	expectResult(t, "x*3.4",   env, x*3.4)
	expectResult(t, "x/2",   env, x/2)
	expectResult(t, "x+x",  env, x+x)

	expectResult(t, "x<1",   env, bool(x<1))
	expectResult(t, "-x<=3",  env, bool(-x<=3))
	expectResult(t, "x>1",   env, bool(x>1))
	expectResult(t, "x>=3",   env, bool(x>=3))
	expectResult(t, "-x==-2.25", env, bool(-x==-2.25))
	expectResult(t, "x!=1",  env, bool(x!=1))
}

func TestComplexBinaryOps(t *testing.T) {
	x := complex(1, 2)
	env := MakeSimpleEnv()
	env.Vars["x"] = reflect.ValueOf(&x)

	expectResult(t, "x + complex(3, 4)", env, x + complex(3, 4))
	expectResult(t, "x - complex(3, 4)", env, x - complex(3, 4))
	expectResult(t, "x * complex(3, 4)", env, x * complex(3, 4))
	expectResult(t, "x / complex(3, 4)", env, x / complex(3, 4))
        expectResult(t, `x + x`, env, x + x)

	expectResult(t, "x == complex(1, 2)", env, bool(x == complex(1, 2)))
	expectResult(t, "x != complex(1, 2)", env, bool(x != complex(1, 2)))
}

func TestStringBinaryOps(t *testing.T) {
	x := "a"
	env := MakeSimpleEnv()
	env.Vars["x"] = reflect.ValueOf(&x)

        expectResult(t, `x + "b"`, env, "a" + "b")
        expectResult(t, `x == "b"`, env, x == "b")
        expectResult(t, `x != "b"`, env, x != "b")
        expectResult(t, `x < "b"`, env, x < "b")
        expectResult(t, `x <= "b"`, env, x <= "b")
        expectResult(t, `x > "a"`, env, x > "a")
        expectResult(t, `x >= "a"`, env, x >= "a")
        expectResult(t, `x + x`, env, x + x)
}

func TestBoolBinaryOps(t *testing.T) {
	x := true
	env := MakeSimpleEnv()
	env.Vars["x"] = reflect.ValueOf(&x)

        expectResult(t, "x == true", env, x == true)
        expectResult(t, "x != true", env, x != true)
        expectResult(t, "x && false", env, x && false)
        expectResult(t, "x || false", env, x || false)
        expectResult(t, `x && x`, env, x && x)
}

func TestArrayStructBinaryOps(t *testing.T) {
	type S struct {
		i int
	}
	a := S{1}
	b := S{2}
	c := [2]int{1,2}
	d := [2]int{1,2}

	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)
	env.Vars["c"] = reflect.ValueOf(&c)
	env.Vars["d"] = reflect.ValueOf(&d)

        expectResult(t, "a == b", env, a == b)
        expectResult(t, "a != b", env, a != b)
        expectResult(t, "c == d", env, c == d)
        expectResult(t, "c != d", env, c != d)
}

func TestInterfaceBinaryOps(t *testing.T) {
	var xi XI = X(0)
	var zi0 ZI = Z(0)
	var zi1 ZI = X(0)
	var x X = X(0)
	env := MakeSimpleEnv()
	env.Vars["xi"] = reflect.ValueOf(&xi)
	env.Vars["zi0"] = reflect.ValueOf(&zi0)
	env.Vars["zi1"] = reflect.ValueOf(&zi1)
	env.Vars["x"] = reflect.ValueOf(&x)

        expectResult(t, "xi == nil", env, xi == nil)
        expectResult(t, "xi != nil", env, xi != nil)
        expectResult(t, "zi0 == xi", env, zi0 == xi)
        expectResult(t, "zi0 != xi", env, zi0 != xi)
	// gc bug http://code.google.com/p/go/issues/detail?id=7207
        expectResult(t, "zi1 == xi", env, true)
        expectResult(t, "zi1 != xi", env, false)
        expectResult(t, "zi0 == x", env, zi0 == x)
        expectResult(t, "zi0 != x", env, zi0 != x)
        expectResult(t, "zi1 == x", env, zi1 == x)
        expectResult(t, "zi1 != x", env, zi1 != x)
}

func TestInterfaceUncompBinaryOps(t *testing.T) {
	type uncompS struct {
		_ []int
	}
	type uncompNestedS struct {
		a interface{}
	}
	var a interface{} = uncompS{}
	var b interface{} = uncompNestedS{a}
	var c interface{} = [1]interface{}{b}
	var d interface{} = []int{}

	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)
	env.Vars["c"] = reflect.ValueOf(&c)
	env.Vars["d"] = reflect.ValueOf(&d)

        expectPanic(t, "a == a", env, "runtime error: comparing uncomparable type eval.uncompS")
        expectPanic(t, "b == b", env, "runtime error: comparing uncomparable type eval.uncompS")
        expectPanic(t, "c == c", env, "runtime error: comparing uncomparable type eval.uncompS")
        expectPanic(t, "d == d", env, "runtime error: comparing uncomparable type []int")
}

func TestPtrBinaryOps(t *testing.T) {
	a := new(int)
	b := a
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)

        expectResult(t, "a == nil", env, a == nil)
        expectResult(t, "a != nil", env, a != nil)
        expectResult(t, "a == b", env, a == b)
        expectResult(t, "a != b", env, a != b)
}

func TestMapSliceFuncBinaryOps(t *testing.T) {
	a := map[int]int{}
	b := []int(nil)
	c := (func())(nil)
	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)
	env.Vars["c"] = reflect.ValueOf(&c)

        expectResult(t, "a == nil", env, a == nil)
        expectResult(t, "a != nil", env, a != nil)
        expectResult(t, "b == nil", env, b == nil)
        expectResult(t, "b != nil", env, b != nil)
        expectResult(t, "c == nil", env, c == nil)
        expectResult(t, "c != nil", env, c != nil)
}

func TestTypedBinaryOps(t *testing.T) {

	type Foo int

	env := MakeSimpleEnv()
	env.Types["Foo"] = reflect.TypeOf(Foo(0))

	expectResult(t, "Foo(1)+Foo(2)", env, Foo(1)+Foo(2))
	expectResult(t, "1-Foo(2)", env, 1-Foo(2))
	expectResult(t, "Foo(1)|2", env, Foo(1)|2)
}

func TestBinaryParens(t *testing.T) {
	x := int32(5)
	env := MakeSimpleEnv()
	env.Vars["x"] = reflect.ValueOf(&x)

	expectResult(t, "(x)+(2)",   env, (x)+(2))
}

