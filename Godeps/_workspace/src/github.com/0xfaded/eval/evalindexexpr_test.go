package eval

import (
	"testing"
	"reflect"
)

func TestIndexArray(t *testing.T) {
	a := [2]int{1, 2}

	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expected := a[1]
	expr := "a[1]"

	expectResult(t, expr, env, expected)
}

func TestIndexArrayPtr(t *testing.T) {
	a := &[2]int{1, 2}

	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expected := a[1]
	expr := "a[1]"

	expectResult(t, expr, env, expected)
}

func TestIndexSlice(t *testing.T) {
	a := []int{1, 2}

	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expected := a[1]
	expr := "a[1]"

	expectResult(t, expr, env, expected)
}

func TestIndexString(t *testing.T) {
	a := "ab"

	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)

	expected := a[1]
	expr := "a[1]"

	expectResult(t, expr, env, expected)
}

