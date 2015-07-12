package eval

import (
	"testing"
	"reflect"
)

func TestFuncCallWithConst(t *testing.T) {
	env := MakeSimpleEnv()
	env.Consts["X"] = reflect.ValueOf(int(10))
	env.Funcs["Foo"] = reflect.ValueOf(func (int) int { return 1; })

	expectResult(t, "Foo(X)", env, 1)
}

func TestFuncCallWithSplatOne(t *testing.T) {
	env := MakeSimpleEnv()

	f := func() int { return 1 }
	g := func(a int) int { return a }

	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["g"] = reflect.ValueOf(g)

	expr := "g(f())"
	expected := g(f())

	expectResult(t, expr, env, expected)
}

func TestFuncCallWithSplatTwo(t *testing.T) {
	env := MakeSimpleEnv()

	f := func() (int, int) { return 1, 2 }
	g := func(a int, b int) int { return a + b }

	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["g"] = reflect.ValueOf(g)

	expr := "g(f())"
	expected := g(f())

	expectResult(t, expr, env, expected)
}

func TestTypeConversionWithEmptyInterface(t *testing.T) {
	env := MakeSimpleEnv()
	expectResult(t, "interface{}(1)", env, interface{}(1))
	expectResult(t, "interface{}(1.5)", env, interface{}(1.5))
	expectResult(t, "interface{}(1i)", env, interface{}(1i))
	expectResult(t, "interface{}('a')", env, interface{}('a'))
	expectResult(t, "interface{}(\"a\")", env, interface{}("a"))
	expectResult(t, "interface{}(nil)", env, interface{}(nil))
	expectResult(t, "interface{}([]int{})", env, interface{}([]int{}))
	expectResult(t, "interface{}('a')", env, interface{}('a'))
}

