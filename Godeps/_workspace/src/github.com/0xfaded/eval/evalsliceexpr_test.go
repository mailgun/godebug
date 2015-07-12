package eval

import (
	"testing"
	"reflect"
)

func TestSliceArray(t *testing.T) {
	a := [3]int{1, 2, 3}

	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expected := a[1:1]
	expr := "a[1:1]"

	expectResult(t, expr, env, expected)

	expected = a[0:2]
	expr = "a[0:2]"

	expectResult(t, expr, env, expected)
}

func TestSliceSlice(t *testing.T) {
	a := []int{1, 2, 3}

	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expected := a[0:2]
	expr := "a[0:2]"

	expectResult(t, expr, env, expected)

	expected = a[1:1]
	expr = "a[1:1]"

	expectResult(t, expr, env, expected)

}

func TestSliceString(t *testing.T) {
	a := "abc"

	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expected := a[1:1]
	expr := "a[1:1]"

	expectResult(t, expr, env, expected)

	expected = a[0:2]
	expr = "a[0:2]"

	expectResult(t, expr, env, expected)

}

