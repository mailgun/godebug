package eval


import (
	"reflect"
	"testing"
)

func TestBuiltinComplex(t *testing.T) {
	env := MakeSimpleEnv()

	expectResult(t, "complex(1, 2)", env, complex(1, 2))
	expectResult(t, "complex(float64(1), 2)", env, complex(float64(1), 2))
	expectResult(t, "complex(1, float64(2))", env, complex(1, float64(2)))
	expectResult(t, "complex(float64(1), float64(2))", env, complex(float64(1), float64(2)))
}

func TestBuiltinReal(t *testing.T) {
	env := MakeSimpleEnv()

	expectResult(t, "real(complex(1, 2))", env, real(complex(1, 2)))
	expectResult(t, "real(complex(float64(1), 2))", env, real(complex(float64(1), 2)))
	expectResult(t, "real(complex(1, float64(2)))", env, real(complex(1, float64(2))))
	expectResult(t, "real(complex(float64(1), float64(2)))", env, real(complex(float64(1), float64(2))))
}

func TestBuiltinImag(t *testing.T) {
	env := MakeSimpleEnv()

	expectResult(t, "imag(complex(1, 2))", env, imag(complex(1, 2)))
	expectResult(t, "imag(complex(float64(1), 2))", env, imag(complex(float64(1), 2)))
	expectResult(t, "imag(complex(1, float64(2)))", env, imag(complex(1, float64(2))))
	expectResult(t, "imag(complex(float64(1), float64(2)))", env, imag(complex(float64(1), float64(2))))
}

func TestBuiltinAppend(t *testing.T) {
	env := MakeSimpleEnv()
	strings := []string {"one", "two"}
	ints := []int{1, 2}
	env.Vars["strings"] = reflect.ValueOf(&strings)
	env.Vars["ints"] = reflect.ValueOf(&ints)

	expectResult(t, "append(strings, \"three\")", env, append(strings, "three"))
	expectResult(t, "append(ints, 3, 4)", env, append(ints, 3, 4))
}

func TestBuiltinAppendSlice(t *testing.T) {
	env := MakeSimpleEnv()
	a := []string {"one", "two"}
	b := []string {"three", "four"}
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)

	expectResult(t, "append(a, b...)", env, append(a, b...))
}

func TestBuiltinCap(t *testing.T) {
	env := MakeSimpleEnv()
	slice := []int {1, 2}
	env.Vars["slice"] = reflect.ValueOf(&slice)

	expectResult(t, "cap(slice)", env, cap(slice))
}

func TestBuiltinLen(t *testing.T) {
	env := MakeSimpleEnv()
	slice := []int {1, 2}
	env.Vars["slice"] = reflect.ValueOf(&slice)

	expectResult(t, "len(\"abc\")", env, len("abc"))
	expectResult(t, "len(slice)", env, len(slice))
}

func TestBuiltinNew(t *testing.T) {
	env := MakeSimpleEnv()
	expr := "new(int)"
	results := getResults(t, expr, env)
	returnKind := results[0].Kind().String()
	if returnKind != "ptr" {
		t.Fatalf("Error Expecting `%s' return Kind to be `ptr' is `%s`", expr, returnKind)
	}
}

func TestBuiltinCopy(t *testing.T) {
	env := MakeSimpleEnv()
	a := []int{1,2,3}
	b := []int{4,5}
	env.Vars["a"] = reflect.ValueOf(&a)
	env.Vars["b"] = reflect.ValueOf(&b)
	expectResult(t, "copy(a, b)", env, copy(a, b))
	expectResult(t, "copy(b, a)", env, copy(b, a))
}

func TestBuiltinDelete(t *testing.T) {
	env := MakeSimpleEnv()
	a := map[int]int{1: 2}
	env.Vars["a"] = reflect.ValueOf(&a)
	getResults(t, "delete(a, 1)", env)
	getResults(t, "delete(a, 1)", env)
	if _, ok := a[1]; ok {
		t.Fatalf("Failed to delete(a, 1)`")
	}
}
