package eval

import (
	"testing"
	"reflect"
)

// Test A1{}
func TestCheckCompositeLitExprA1XX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `a1{}`, env, reflect.TypeOf(a1{}))
}

// Test A1{1}
func TestCheckCompositeLitExprA1XInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `a1{1}`, env, reflect.TypeOf(a1{1}))
}

// Test A1{"a"}
func TestCheckCompositeLitExprA1XString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{"a"}`, env,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{nil}
func TestCheckCompositeLitExprA1XNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{nil}`, env,
		`cannot convert nil to type int`,
	)

}

// Test A1{1: 1}
func TestCheckCompositeLitExprA1XIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: 1}`, env,
		`array index 2 out of bounds [0:1]`,
	)

}

// Test A1{1: "a"}
func TestCheckCompositeLitExprA1XIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: "a"}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{float32(1.5): 1}
func TestCheckCompositeLitExprA1XFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{float32(1.5): 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{a: 1}
func TestCheckCompositeLitExprA1XAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{a: "a"}
func TestCheckCompositeLitExprA1XAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{b: 1}
func TestCheckCompositeLitExprA1XBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{b: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{c: 1}
func TestCheckCompositeLitExprA1XCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{c: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{1}
func TestCheckCompositeLitExprA1IntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `a1{1}`, env, reflect.TypeOf(a1{1}))
}

// Test A1{1, 1}
func TestCheckCompositeLitExprA1IntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1, 1}`, env,
		`array index 2 out of bounds [0:1]`,
	)

}

// Test A1{1, "a"}
func TestCheckCompositeLitExprA1IntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1, "a"}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{1, nil}
func TestCheckCompositeLitExprA1IntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1, nil}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot convert nil to type int`,
	)

}

// Test A1{1, 1: 1}
func TestCheckCompositeLitExprA1IntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1, 1: 1}`, env,
		`array index 2 out of bounds [0:1]`,
	)

}

// Test A1{1, 1: "a"}
func TestCheckCompositeLitExprA1IntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1, 1: "a"}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{1, float32(1.5): 1}
func TestCheckCompositeLitExprA1IntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1, float32(1.5): 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{1, a: 1}
func TestCheckCompositeLitExprA1IntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1, a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{1, a: "a"}
func TestCheckCompositeLitExprA1IntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1, a: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{1, b: 1}
func TestCheckCompositeLitExprA1IntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1, b: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{1, c: 1}
func TestCheckCompositeLitExprA1IntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1, c: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{nil}
func TestCheckCompositeLitExprA1NilX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{nil}`, env,
		`cannot convert nil to type int`,
	)

}

// Test A1{nil, 1}
func TestCheckCompositeLitExprA1NilInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{nil, 1}`, env,
		`cannot convert nil to type int`,
		`array index 2 out of bounds [0:1]`,
	)

}

// Test A1{nil, "a"}
func TestCheckCompositeLitExprA1NilString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{nil, "a"}`, env,
		`cannot convert nil to type int`,
		`array index 2 out of bounds [0:1]`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{nil, nil}
func TestCheckCompositeLitExprA1NilNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{nil, nil}`, env,
		`cannot convert nil to type int`,
		`array index 2 out of bounds [0:1]`,
		`cannot convert nil to type int`,
	)

}

// Test A1{nil, 1: 1}
func TestCheckCompositeLitExprA1NilIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{nil, 1: 1}`, env,
		`cannot convert nil to type int`,
		`array index 2 out of bounds [0:1]`,
	)

}

// Test A1{nil, 1: "a"}
func TestCheckCompositeLitExprA1NilIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{nil, 1: "a"}`, env,
		`cannot convert nil to type int`,
		`array index 2 out of bounds [0:1]`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{nil, float32(1.5): 1}
func TestCheckCompositeLitExprA1NilFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{nil, float32(1.5): 1}`, env,
		`cannot convert nil to type int`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{nil, a: 1}
func TestCheckCompositeLitExprA1NilAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{nil, a: 1}`, env,
		`cannot convert nil to type int`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{nil, a: "a"}
func TestCheckCompositeLitExprA1NilAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{nil, a: "a"}`, env,
		`cannot convert nil to type int`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{nil, b: 1}
func TestCheckCompositeLitExprA1NilBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{nil, b: 1}`, env,
		`cannot convert nil to type int`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{nil, c: 1}
func TestCheckCompositeLitExprA1NilCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{nil, c: 1}`, env,
		`cannot convert nil to type int`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{true}
func TestCheckCompositeLitExprA1BoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{true}`, env,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test A1{true, 1}
func TestCheckCompositeLitExprA1BoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{true, 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`array index 2 out of bounds [0:1]`,
	)

}

// Test A1{true, "a"}
func TestCheckCompositeLitExprA1BoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{true, "a"}`, env,
		`cannot use true (type bool) as type int in array element`,
		`array index 2 out of bounds [0:1]`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{true, nil}
func TestCheckCompositeLitExprA1BoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{true, nil}`, env,
		`cannot use true (type bool) as type int in array element`,
		`array index 2 out of bounds [0:1]`,
		`cannot convert nil to type int`,
	)

}

// Test A1{true, 1: 1}
func TestCheckCompositeLitExprA1BoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{true, 1: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`array index 2 out of bounds [0:1]`,
	)

}

// Test A1{true, 1: "a"}
func TestCheckCompositeLitExprA1BoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{true, 1: "a"}`, env,
		`cannot use true (type bool) as type int in array element`,
		`array index 2 out of bounds [0:1]`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{true, float32(1.5): 1}
func TestCheckCompositeLitExprA1BoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{true, float32(1.5): 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{true, a: 1}
func TestCheckCompositeLitExprA1BoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{true, a: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{true, a: "a"}
func TestCheckCompositeLitExprA1BoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{true, a: "a"}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{true, b: 1}
func TestCheckCompositeLitExprA1BoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{true, b: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{true, c: 1}
func TestCheckCompositeLitExprA1BoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{true, c: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{1: 1}
func TestCheckCompositeLitExprA1IntKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: 1}`, env,
		`array index 2 out of bounds [0:1]`,
	)

}

// Test A1{1: 1, 1}
func TestCheckCompositeLitExprA1IntKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: 1, 1}`, env,
		`array index 2 out of bounds [0:1]`,
	)

}

// Test A1{1: 1, "a"}
func TestCheckCompositeLitExprA1IntKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: 1, "a"}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{1: 1, nil}
func TestCheckCompositeLitExprA1IntKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: 1, nil}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot convert nil to type int`,
	)

}

// Test A1{1: 1, 1: 1}
func TestCheckCompositeLitExprA1IntKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: 1, 1: 1}`, env,
		`array index 2 out of bounds [0:1]`,
		`duplicate index in array literal: 1`,
	)

}

// Test A1{1: 1, 1: "a"}
func TestCheckCompositeLitExprA1IntKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: 1, 1: "a"}`, env,
		`array index 2 out of bounds [0:1]`,
		`duplicate index in array literal: 1`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{1: 1, float32(1.5): 1}
func TestCheckCompositeLitExprA1IntKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: 1, float32(1.5): 1}`, env,
		`array index 2 out of bounds [0:1]`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{1: 1, a: 1}
func TestCheckCompositeLitExprA1IntKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: 1, a: 1}`, env,
		`array index 2 out of bounds [0:1]`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{1: 1, a: "a"}
func TestCheckCompositeLitExprA1IntKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: 1, a: "a"}`, env,
		`array index 2 out of bounds [0:1]`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{1: 1, b: 1}
func TestCheckCompositeLitExprA1IntKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: 1, b: 1}`, env,
		`array index 2 out of bounds [0:1]`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{1: 1, c: 1}
func TestCheckCompositeLitExprA1IntKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: 1, c: 1}`, env,
		`array index 2 out of bounds [0:1]`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{1: true}
func TestCheckCompositeLitExprA1IntKBoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: true}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test A1{1: true, 1}
func TestCheckCompositeLitExprA1IntKBoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: true, 1}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test A1{1: true, "a"}
func TestCheckCompositeLitExprA1IntKBoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: true, "a"}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{1: true, nil}
func TestCheckCompositeLitExprA1IntKBoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: true, nil}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert nil to type int`,
	)

}

// Test A1{1: true, 1: 1}
func TestCheckCompositeLitExprA1IntKBoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: true, 1: 1}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot use true (type bool) as type int in array element`,
		`duplicate index in array literal: 1`,
	)

}

// Test A1{1: true, 1: "a"}
func TestCheckCompositeLitExprA1IntKBoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: true, 1: "a"}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot use true (type bool) as type int in array element`,
		`duplicate index in array literal: 1`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{1: true, float32(1.5): 1}
func TestCheckCompositeLitExprA1IntKBoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: true, float32(1.5): 1}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot use true (type bool) as type int in array element`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{1: true, a: 1}
func TestCheckCompositeLitExprA1IntKBoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: true, a: 1}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{1: true, a: "a"}
func TestCheckCompositeLitExprA1IntKBoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: true, a: "a"}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{1: true, b: 1}
func TestCheckCompositeLitExprA1IntKBoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: true, b: 1}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{1: true, c: 1}
func TestCheckCompositeLitExprA1IntKBoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{1: true, c: 1}`, env,
		`array index 2 out of bounds [0:1]`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{float32(1.4): 1}
func TestCheckCompositeLitExprA1FloatKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{float32(1.4): 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{float32(1.4): 1, 1}
func TestCheckCompositeLitExprA1FloatKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{float32(1.4): 1, 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{float32(1.4): 1, "a"}
func TestCheckCompositeLitExprA1FloatKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{float32(1.4): 1, "a"}`, env,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{float32(1.4): 1, nil}
func TestCheckCompositeLitExprA1FloatKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{float32(1.4): 1, nil}`, env,
		`array index must be non-negative integer constant`,
		`cannot convert nil to type int`,
	)

}

// Test A1{float32(1.4): 1, 1: 1}
func TestCheckCompositeLitExprA1FloatKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{float32(1.4): 1, 1: 1}`, env,
		`array index must be non-negative integer constant`,
		`array index 2 out of bounds [0:1]`,
	)

}

// Test A1{float32(1.4): 1, 1: "a"}
func TestCheckCompositeLitExprA1FloatKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{float32(1.4): 1, 1: "a"}`, env,
		`array index must be non-negative integer constant`,
		`array index 2 out of bounds [0:1]`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{float32(1.4): 1, float32(1.5): 1}
func TestCheckCompositeLitExprA1FloatKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{float32(1.4): 1, float32(1.5): 1}`, env,
		`array index must be non-negative integer constant`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{float32(1.4): 1, a: 1}
func TestCheckCompositeLitExprA1FloatKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{float32(1.4): 1, a: 1}`, env,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{float32(1.4): 1, a: "a"}
func TestCheckCompositeLitExprA1FloatKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{float32(1.4): 1, a: "a"}`, env,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{float32(1.4): 1, b: 1}
func TestCheckCompositeLitExprA1FloatKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{float32(1.4): 1, b: 1}`, env,
		`array index must be non-negative integer constant`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{float32(1.4): 1, c: 1}
func TestCheckCompositeLitExprA1FloatKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{float32(1.4): 1, c: 1}`, env,
		`array index must be non-negative integer constant`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{a: 1}
func TestCheckCompositeLitExprA1AKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{a: 1, 1}
func TestCheckCompositeLitExprA1AKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: 1, 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{a: 1, "a"}
func TestCheckCompositeLitExprA1AKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: 1, "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{a: 1, nil}
func TestCheckCompositeLitExprA1AKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: 1, nil}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert nil to type int`,
	)

}

// Test A1{a: 1, 1: 1}
func TestCheckCompositeLitExprA1AKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: 1, 1: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`array index 2 out of bounds [0:1]`,
	)

}

// Test A1{a: 1, 1: "a"}
func TestCheckCompositeLitExprA1AKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: 1, 1: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`array index 2 out of bounds [0:1]`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{a: 1, float32(1.5): 1}
func TestCheckCompositeLitExprA1AKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: 1, float32(1.5): 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{a: 1, a: 1}
func TestCheckCompositeLitExprA1AKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: 1, a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{a: 1, a: "a"}
func TestCheckCompositeLitExprA1AKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: 1, a: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{a: 1, b: 1}
func TestCheckCompositeLitExprA1AKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: 1, b: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{a: 1, c: 1}
func TestCheckCompositeLitExprA1AKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: 1, c: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{a: true}
func TestCheckCompositeLitExprA1AKBoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: true}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test A1{a: true, 1}
func TestCheckCompositeLitExprA1AKBoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: true, 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test A1{a: true, "a"}
func TestCheckCompositeLitExprA1AKBoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: true, "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{a: true, nil}
func TestCheckCompositeLitExprA1AKBoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: true, nil}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert nil to type int`,
	)

}

// Test A1{a: true, 1: 1}
func TestCheckCompositeLitExprA1AKBoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: true, 1: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`array index 2 out of bounds [0:1]`,
	)

}

// Test A1{a: true, 1: "a"}
func TestCheckCompositeLitExprA1AKBoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: true, 1: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`array index 2 out of bounds [0:1]`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{a: true, float32(1.5): 1}
func TestCheckCompositeLitExprA1AKBoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: true, float32(1.5): 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{a: true, a: 1}
func TestCheckCompositeLitExprA1AKBoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: true, a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{a: true, a: "a"}
func TestCheckCompositeLitExprA1AKBoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: true, a: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{a: true, b: 1}
func TestCheckCompositeLitExprA1AKBoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: true, b: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{a: true, c: 1}
func TestCheckCompositeLitExprA1AKBoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{a: true, c: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{b: 1}
func TestCheckCompositeLitExprA1BKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{b: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{b: 1, 1}
func TestCheckCompositeLitExprA1BKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{b: 1, 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{b: 1, "a"}
func TestCheckCompositeLitExprA1BKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{b: 1, "a"}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{b: 1, nil}
func TestCheckCompositeLitExprA1BKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{b: 1, nil}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`cannot convert nil to type int`,
	)

}

// Test A1{b: 1, 1: 1}
func TestCheckCompositeLitExprA1BKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{b: 1, 1: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`array index 2 out of bounds [0:1]`,
	)

}

// Test A1{b: 1, 1: "a"}
func TestCheckCompositeLitExprA1BKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{b: 1, 1: "a"}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`array index 2 out of bounds [0:1]`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{b: 1, float32(1.5): 1}
func TestCheckCompositeLitExprA1BKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{b: 1, float32(1.5): 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{b: 1, a: 1}
func TestCheckCompositeLitExprA1BKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{b: 1, a: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{b: 1, a: "a"}
func TestCheckCompositeLitExprA1BKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{b: 1, a: "a"}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{b: 1, b: 1}
func TestCheckCompositeLitExprA1BKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{b: 1, b: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{b: 1, c: 1}
func TestCheckCompositeLitExprA1BKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{b: 1, c: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{c: 1}
func TestCheckCompositeLitExprA1CKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{c: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{c: 1, 1}
func TestCheckCompositeLitExprA1CKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{c: 1, 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{c: 1, "a"}
func TestCheckCompositeLitExprA1CKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{c: 1, "a"}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{c: 1, nil}
func TestCheckCompositeLitExprA1CKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{c: 1, nil}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`cannot convert nil to type int`,
	)

}

// Test A1{c: 1, 1: 1}
func TestCheckCompositeLitExprA1CKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{c: 1, 1: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`array index 2 out of bounds [0:1]`,
	)

}

// Test A1{c: 1, 1: "a"}
func TestCheckCompositeLitExprA1CKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{c: 1, 1: "a"}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`array index 2 out of bounds [0:1]`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{c: 1, float32(1.5): 1}
func TestCheckCompositeLitExprA1CKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{c: 1, float32(1.5): 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{c: 1, a: 1}
func TestCheckCompositeLitExprA1CKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{c: 1, a: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{c: 1, a: "a"}
func TestCheckCompositeLitExprA1CKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{c: 1, a: "a"}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A1{c: 1, b: 1}
func TestCheckCompositeLitExprA1CKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{c: 1, b: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A1{c: 1, c: 1}
func TestCheckCompositeLitExprA1CKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a1{c: 1, c: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{}
func TestCheckCompositeLitExprA2XX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `a2{}`, env, reflect.TypeOf(a2{}))
}

// Test A2{1}
func TestCheckCompositeLitExprA2XInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `a2{1}`, env, reflect.TypeOf(a2{1}))
}

// Test A2{"a"}
func TestCheckCompositeLitExprA2XString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{"a"}`, env,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{nil}
func TestCheckCompositeLitExprA2XNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{nil}`, env,
		`cannot convert nil to type int`,
	)

}

// Test A2{1: 1}
func TestCheckCompositeLitExprA2XIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `a2{1: 1}`, env, reflect.TypeOf(a2{1: 1}))
}

// Test A2{1: "a"}
func TestCheckCompositeLitExprA2XIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: "a"}`, env,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{float32(1.5): 1}
func TestCheckCompositeLitExprA2XFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{float32(1.5): 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{a: 1}
func TestCheckCompositeLitExprA2XAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{a: "a"}
func TestCheckCompositeLitExprA2XAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{b: 1}
func TestCheckCompositeLitExprA2XBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{b: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{c: 1}
func TestCheckCompositeLitExprA2XCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{c: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{1}
func TestCheckCompositeLitExprA2IntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `a2{1}`, env, reflect.TypeOf(a2{1}))
}

// Test A2{1, 1}
func TestCheckCompositeLitExprA2IntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `a2{1, 1}`, env, reflect.TypeOf(a2{1, 1}))
}

// Test A2{1, "a"}
func TestCheckCompositeLitExprA2IntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1, "a"}`, env,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{1, nil}
func TestCheckCompositeLitExprA2IntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1, nil}`, env,
		`cannot convert nil to type int`,
	)

}

// Test A2{1, 1: 1}
func TestCheckCompositeLitExprA2IntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `a2{1, 1: 1}`, env, reflect.TypeOf(a2{1, 1: 1}))
}

// Test A2{1, 1: "a"}
func TestCheckCompositeLitExprA2IntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1, 1: "a"}`, env,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{1, float32(1.5): 1}
func TestCheckCompositeLitExprA2IntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1, float32(1.5): 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{1, a: 1}
func TestCheckCompositeLitExprA2IntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1, a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{1, a: "a"}
func TestCheckCompositeLitExprA2IntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1, a: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{1, b: 1}
func TestCheckCompositeLitExprA2IntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1, b: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{1, c: 1}
func TestCheckCompositeLitExprA2IntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1, c: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{nil}
func TestCheckCompositeLitExprA2NilX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{nil}`, env,
		`cannot convert nil to type int`,
	)

}

// Test A2{nil, 1}
func TestCheckCompositeLitExprA2NilInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{nil, 1}`, env,
		`cannot convert nil to type int`,
	)

}

// Test A2{nil, "a"}
func TestCheckCompositeLitExprA2NilString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{nil, "a"}`, env,
		`cannot convert nil to type int`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{nil, nil}
func TestCheckCompositeLitExprA2NilNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{nil, nil}`, env,
		`cannot convert nil to type int`,
		`cannot convert nil to type int`,
	)

}

// Test A2{nil, 1: 1}
func TestCheckCompositeLitExprA2NilIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{nil, 1: 1}`, env,
		`cannot convert nil to type int`,
	)

}

// Test A2{nil, 1: "a"}
func TestCheckCompositeLitExprA2NilIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{nil, 1: "a"}`, env,
		`cannot convert nil to type int`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{nil, float32(1.5): 1}
func TestCheckCompositeLitExprA2NilFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{nil, float32(1.5): 1}`, env,
		`cannot convert nil to type int`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{nil, a: 1}
func TestCheckCompositeLitExprA2NilAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{nil, a: 1}`, env,
		`cannot convert nil to type int`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{nil, a: "a"}
func TestCheckCompositeLitExprA2NilAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{nil, a: "a"}`, env,
		`cannot convert nil to type int`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{nil, b: 1}
func TestCheckCompositeLitExprA2NilBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{nil, b: 1}`, env,
		`cannot convert nil to type int`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{nil, c: 1}
func TestCheckCompositeLitExprA2NilCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{nil, c: 1}`, env,
		`cannot convert nil to type int`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{true}
func TestCheckCompositeLitExprA2BoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{true}`, env,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test A2{true, 1}
func TestCheckCompositeLitExprA2BoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{true, 1}`, env,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test A2{true, "a"}
func TestCheckCompositeLitExprA2BoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{true, "a"}`, env,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{true, nil}
func TestCheckCompositeLitExprA2BoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{true, nil}`, env,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert nil to type int`,
	)

}

// Test A2{true, 1: 1}
func TestCheckCompositeLitExprA2BoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{true, 1: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test A2{true, 1: "a"}
func TestCheckCompositeLitExprA2BoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{true, 1: "a"}`, env,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{true, float32(1.5): 1}
func TestCheckCompositeLitExprA2BoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{true, float32(1.5): 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{true, a: 1}
func TestCheckCompositeLitExprA2BoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{true, a: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{true, a: "a"}
func TestCheckCompositeLitExprA2BoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{true, a: "a"}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{true, b: 1}
func TestCheckCompositeLitExprA2BoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{true, b: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{true, c: 1}
func TestCheckCompositeLitExprA2BoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{true, c: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{1: 1}
func TestCheckCompositeLitExprA2IntKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `a2{1: 1}`, env, reflect.TypeOf(a2{1: 1}))
}

// Test A2{1: 1, 1}
func TestCheckCompositeLitExprA2IntKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: 1, 1}`, env,
		`array index 3 out of bounds [0:2]`,
	)

}

// Test A2{1: 1, "a"}
func TestCheckCompositeLitExprA2IntKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: 1, "a"}`, env,
		`array index 3 out of bounds [0:2]`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{1: 1, nil}
func TestCheckCompositeLitExprA2IntKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: 1, nil}`, env,
		`array index 3 out of bounds [0:2]`,
		`cannot convert nil to type int`,
	)

}

// Test A2{1: 1, 1: 1}
func TestCheckCompositeLitExprA2IntKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: 1, 1: 1}`, env,
		`duplicate index in array literal: 1`,
	)

}

// Test A2{1: 1, 1: "a"}
func TestCheckCompositeLitExprA2IntKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: 1, 1: "a"}`, env,
		`duplicate index in array literal: 1`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{1: 1, float32(1.5): 1}
func TestCheckCompositeLitExprA2IntKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: 1, float32(1.5): 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{1: 1, a: 1}
func TestCheckCompositeLitExprA2IntKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: 1, a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{1: 1, a: "a"}
func TestCheckCompositeLitExprA2IntKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: 1, a: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{1: 1, b: 1}
func TestCheckCompositeLitExprA2IntKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: 1, b: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{1: 1, c: 1}
func TestCheckCompositeLitExprA2IntKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: 1, c: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{1: true}
func TestCheckCompositeLitExprA2IntKBoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: true}`, env,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test A2{1: true, 1}
func TestCheckCompositeLitExprA2IntKBoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: true, 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`array index 3 out of bounds [0:2]`,
	)

}

// Test A2{1: true, "a"}
func TestCheckCompositeLitExprA2IntKBoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: true, "a"}`, env,
		`cannot use true (type bool) as type int in array element`,
		`array index 3 out of bounds [0:2]`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{1: true, nil}
func TestCheckCompositeLitExprA2IntKBoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: true, nil}`, env,
		`cannot use true (type bool) as type int in array element`,
		`array index 3 out of bounds [0:2]`,
		`cannot convert nil to type int`,
	)

}

// Test A2{1: true, 1: 1}
func TestCheckCompositeLitExprA2IntKBoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: true, 1: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`duplicate index in array literal: 1`,
	)

}

// Test A2{1: true, 1: "a"}
func TestCheckCompositeLitExprA2IntKBoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: true, 1: "a"}`, env,
		`cannot use true (type bool) as type int in array element`,
		`duplicate index in array literal: 1`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{1: true, float32(1.5): 1}
func TestCheckCompositeLitExprA2IntKBoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: true, float32(1.5): 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{1: true, a: 1}
func TestCheckCompositeLitExprA2IntKBoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: true, a: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{1: true, a: "a"}
func TestCheckCompositeLitExprA2IntKBoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: true, a: "a"}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{1: true, b: 1}
func TestCheckCompositeLitExprA2IntKBoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: true, b: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{1: true, c: 1}
func TestCheckCompositeLitExprA2IntKBoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{1: true, c: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{float32(1.4): 1}
func TestCheckCompositeLitExprA2FloatKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{float32(1.4): 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{float32(1.4): 1, 1}
func TestCheckCompositeLitExprA2FloatKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{float32(1.4): 1, 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{float32(1.4): 1, "a"}
func TestCheckCompositeLitExprA2FloatKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{float32(1.4): 1, "a"}`, env,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{float32(1.4): 1, nil}
func TestCheckCompositeLitExprA2FloatKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{float32(1.4): 1, nil}`, env,
		`array index must be non-negative integer constant`,
		`cannot convert nil to type int`,
	)

}

// Test A2{float32(1.4): 1, 1: 1}
func TestCheckCompositeLitExprA2FloatKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{float32(1.4): 1, 1: 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{float32(1.4): 1, 1: "a"}
func TestCheckCompositeLitExprA2FloatKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{float32(1.4): 1, 1: "a"}`, env,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{float32(1.4): 1, float32(1.5): 1}
func TestCheckCompositeLitExprA2FloatKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{float32(1.4): 1, float32(1.5): 1}`, env,
		`array index must be non-negative integer constant`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{float32(1.4): 1, a: 1}
func TestCheckCompositeLitExprA2FloatKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{float32(1.4): 1, a: 1}`, env,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{float32(1.4): 1, a: "a"}
func TestCheckCompositeLitExprA2FloatKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{float32(1.4): 1, a: "a"}`, env,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{float32(1.4): 1, b: 1}
func TestCheckCompositeLitExprA2FloatKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{float32(1.4): 1, b: 1}`, env,
		`array index must be non-negative integer constant`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{float32(1.4): 1, c: 1}
func TestCheckCompositeLitExprA2FloatKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{float32(1.4): 1, c: 1}`, env,
		`array index must be non-negative integer constant`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{a: 1}
func TestCheckCompositeLitExprA2AKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{a: 1, 1}
func TestCheckCompositeLitExprA2AKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: 1, 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{a: 1, "a"}
func TestCheckCompositeLitExprA2AKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: 1, "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{a: 1, nil}
func TestCheckCompositeLitExprA2AKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: 1, nil}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert nil to type int`,
	)

}

// Test A2{a: 1, 1: 1}
func TestCheckCompositeLitExprA2AKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: 1, 1: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{a: 1, 1: "a"}
func TestCheckCompositeLitExprA2AKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: 1, 1: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{a: 1, float32(1.5): 1}
func TestCheckCompositeLitExprA2AKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: 1, float32(1.5): 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{a: 1, a: 1}
func TestCheckCompositeLitExprA2AKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: 1, a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{a: 1, a: "a"}
func TestCheckCompositeLitExprA2AKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: 1, a: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{a: 1, b: 1}
func TestCheckCompositeLitExprA2AKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: 1, b: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{a: 1, c: 1}
func TestCheckCompositeLitExprA2AKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: 1, c: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{a: true}
func TestCheckCompositeLitExprA2AKBoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: true}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test A2{a: true, 1}
func TestCheckCompositeLitExprA2AKBoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: true, 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test A2{a: true, "a"}
func TestCheckCompositeLitExprA2AKBoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: true, "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{a: true, nil}
func TestCheckCompositeLitExprA2AKBoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: true, nil}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert nil to type int`,
	)

}

// Test A2{a: true, 1: 1}
func TestCheckCompositeLitExprA2AKBoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: true, 1: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test A2{a: true, 1: "a"}
func TestCheckCompositeLitExprA2AKBoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: true, 1: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{a: true, float32(1.5): 1}
func TestCheckCompositeLitExprA2AKBoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: true, float32(1.5): 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{a: true, a: 1}
func TestCheckCompositeLitExprA2AKBoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: true, a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{a: true, a: "a"}
func TestCheckCompositeLitExprA2AKBoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: true, a: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{a: true, b: 1}
func TestCheckCompositeLitExprA2AKBoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: true, b: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{a: true, c: 1}
func TestCheckCompositeLitExprA2AKBoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{a: true, c: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{b: 1}
func TestCheckCompositeLitExprA2BKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{b: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{b: 1, 1}
func TestCheckCompositeLitExprA2BKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{b: 1, 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{b: 1, "a"}
func TestCheckCompositeLitExprA2BKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{b: 1, "a"}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{b: 1, nil}
func TestCheckCompositeLitExprA2BKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{b: 1, nil}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`cannot convert nil to type int`,
	)

}

// Test A2{b: 1, 1: 1}
func TestCheckCompositeLitExprA2BKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{b: 1, 1: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{b: 1, 1: "a"}
func TestCheckCompositeLitExprA2BKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{b: 1, 1: "a"}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{b: 1, float32(1.5): 1}
func TestCheckCompositeLitExprA2BKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{b: 1, float32(1.5): 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{b: 1, a: 1}
func TestCheckCompositeLitExprA2BKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{b: 1, a: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{b: 1, a: "a"}
func TestCheckCompositeLitExprA2BKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{b: 1, a: "a"}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{b: 1, b: 1}
func TestCheckCompositeLitExprA2BKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{b: 1, b: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{b: 1, c: 1}
func TestCheckCompositeLitExprA2BKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{b: 1, c: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{c: 1}
func TestCheckCompositeLitExprA2CKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{c: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{c: 1, 1}
func TestCheckCompositeLitExprA2CKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{c: 1, 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{c: 1, "a"}
func TestCheckCompositeLitExprA2CKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{c: 1, "a"}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{c: 1, nil}
func TestCheckCompositeLitExprA2CKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{c: 1, nil}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`cannot convert nil to type int`,
	)

}

// Test A2{c: 1, 1: 1}
func TestCheckCompositeLitExprA2CKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{c: 1, 1: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{c: 1, 1: "a"}
func TestCheckCompositeLitExprA2CKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{c: 1, 1: "a"}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{c: 1, float32(1.5): 1}
func TestCheckCompositeLitExprA2CKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{c: 1, float32(1.5): 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{c: 1, a: 1}
func TestCheckCompositeLitExprA2CKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{c: 1, a: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{c: 1, a: "a"}
func TestCheckCompositeLitExprA2CKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{c: 1, a: "a"}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test A2{c: 1, b: 1}
func TestCheckCompositeLitExprA2CKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{c: 1, b: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test A2{c: 1, c: 1}
func TestCheckCompositeLitExprA2CKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `a2{c: 1, c: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test S1{}
func TestCheckCompositeLitExprS1XX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `s1{}`, env, reflect.TypeOf(s1{}))
}

// Test S1{1}
func TestCheckCompositeLitExprS1XInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `s1{1}`, env, reflect.TypeOf(s1{1}))
}

// Test S1{"a"}
func TestCheckCompositeLitExprS1XString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{"a"}`, env,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S1{nil}
func TestCheckCompositeLitExprS1XNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{nil}`, env,
		`cannot use nil as type int in field value`,
	)

}

// Test S1{1: 1}
func TestCheckCompositeLitExprS1XIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: 1}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{1: "a"}
func TestCheckCompositeLitExprS1XIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: "a"}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{float32(1.5): 1}
func TestCheckCompositeLitExprS1XFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{float32(1.5): 1}`, env,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S1{a: 1}
func TestCheckCompositeLitExprS1XAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `s1{a: 1}`, env, reflect.TypeOf(s1{a: 1}))
}

// Test S1{a: "a"}
func TestCheckCompositeLitExprS1XAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: "a"}`, env,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S1{b: 1}
func TestCheckCompositeLitExprS1XBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{b: 1}`, env,
		`unknown eval.s1 field 'b' in struct literal`,
	)

}

// Test S1{c: 1}
func TestCheckCompositeLitExprS1XCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{c: 1}`, env,
		`unknown eval.s1 field 'c' in struct literal`,
	)

}

// Test S1{1}
func TestCheckCompositeLitExprS1IntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `s1{1}`, env, reflect.TypeOf(s1{1}))
}

// Test S1{1, 1}
func TestCheckCompositeLitExprS1IntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1, 1}`, env,
		`too many values in struct initializer`,
	)

}

// Test S1{1, "a"}
func TestCheckCompositeLitExprS1IntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1, "a"}`, env,
		`too many values in struct initializer`,
	)

}

// Test S1{1, nil}
func TestCheckCompositeLitExprS1IntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1, nil}`, env,
		`too many values in struct initializer`,
	)

}

// Test S1{1, 1: 1}
func TestCheckCompositeLitExprS1IntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1, 1: 1}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{1, 1: "a"}
func TestCheckCompositeLitExprS1IntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1, 1: "a"}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{1, float32(1.5): 1}
func TestCheckCompositeLitExprS1IntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1, float32(1.5): 1}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S1{1, a: 1}
func TestCheckCompositeLitExprS1IntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1, a: 1}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{1, a: "a"}
func TestCheckCompositeLitExprS1IntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1, a: "a"}`, env,
		`mixture of field:value and value initializers`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S1{1, b: 1}
func TestCheckCompositeLitExprS1IntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1, b: 1}`, env,
		`mixture of field:value and value initializers`,
		`unknown eval.s1 field 'b' in struct literal`,
	)

}

// Test S1{1, c: 1}
func TestCheckCompositeLitExprS1IntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1, c: 1}`, env,
		`mixture of field:value and value initializers`,
		`unknown eval.s1 field 'c' in struct literal`,
	)

}

// Test S1{nil}
func TestCheckCompositeLitExprS1NilX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{nil}`, env,
		`cannot use nil as type int in field value`,
	)

}

// Test S1{nil, 1}
func TestCheckCompositeLitExprS1NilInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{nil, 1}`, env,
		`cannot use nil as type int in field value`,
		`too many values in struct initializer`,
	)

}

// Test S1{nil, "a"}
func TestCheckCompositeLitExprS1NilString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{nil, "a"}`, env,
		`cannot use nil as type int in field value`,
		`too many values in struct initializer`,
	)

}

// Test S1{nil, nil}
func TestCheckCompositeLitExprS1NilNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{nil, nil}`, env,
		`cannot use nil as type int in field value`,
		`too many values in struct initializer`,
	)

}

// Test S1{nil, 1: 1}
func TestCheckCompositeLitExprS1NilIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{nil, 1: 1}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{nil, 1: "a"}
func TestCheckCompositeLitExprS1NilIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{nil, 1: "a"}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{nil, float32(1.5): 1}
func TestCheckCompositeLitExprS1NilFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{nil, float32(1.5): 1}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S1{nil, a: 1}
func TestCheckCompositeLitExprS1NilAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{nil, a: 1}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{nil, a: "a"}
func TestCheckCompositeLitExprS1NilAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{nil, a: "a"}`, env,
		`mixture of field:value and value initializers`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S1{nil, b: 1}
func TestCheckCompositeLitExprS1NilBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{nil, b: 1}`, env,
		`mixture of field:value and value initializers`,
		`unknown eval.s1 field 'b' in struct literal`,
	)

}

// Test S1{nil, c: 1}
func TestCheckCompositeLitExprS1NilCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{nil, c: 1}`, env,
		`mixture of field:value and value initializers`,
		`unknown eval.s1 field 'c' in struct literal`,
	)

}

// Test S1{true}
func TestCheckCompositeLitExprS1BoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{true}`, env,
		`cannot use true (type bool) as type int in field value`,
	)

}

// Test S1{true, 1}
func TestCheckCompositeLitExprS1BoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{true, 1}`, env,
		`cannot use true (type bool) as type int in field value`,
		`too many values in struct initializer`,
	)

}

// Test S1{true, "a"}
func TestCheckCompositeLitExprS1BoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{true, "a"}`, env,
		`cannot use true (type bool) as type int in field value`,
		`too many values in struct initializer`,
	)

}

// Test S1{true, nil}
func TestCheckCompositeLitExprS1BoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{true, nil}`, env,
		`cannot use true (type bool) as type int in field value`,
		`too many values in struct initializer`,
	)

}

// Test S1{true, 1: 1}
func TestCheckCompositeLitExprS1BoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{true, 1: 1}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{true, 1: "a"}
func TestCheckCompositeLitExprS1BoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{true, 1: "a"}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{true, float32(1.5): 1}
func TestCheckCompositeLitExprS1BoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{true, float32(1.5): 1}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S1{true, a: 1}
func TestCheckCompositeLitExprS1BoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{true, a: 1}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{true, a: "a"}
func TestCheckCompositeLitExprS1BoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{true, a: "a"}`, env,
		`mixture of field:value and value initializers`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S1{true, b: 1}
func TestCheckCompositeLitExprS1BoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{true, b: 1}`, env,
		`mixture of field:value and value initializers`,
		`unknown eval.s1 field 'b' in struct literal`,
	)

}

// Test S1{true, c: 1}
func TestCheckCompositeLitExprS1BoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{true, c: 1}`, env,
		`mixture of field:value and value initializers`,
		`unknown eval.s1 field 'c' in struct literal`,
	)

}

// Test S1{1: 1}
func TestCheckCompositeLitExprS1IntKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: 1}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{1: 1, 1}
func TestCheckCompositeLitExprS1IntKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: 1, 1}`, env,
		`invalid field name 1 in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{1: 1, "a"}
func TestCheckCompositeLitExprS1IntKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: 1, "a"}`, env,
		`invalid field name 1 in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{1: 1, nil}
func TestCheckCompositeLitExprS1IntKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: 1, nil}`, env,
		`invalid field name 1 in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{1: 1, 1: 1}
func TestCheckCompositeLitExprS1IntKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: 1, 1: 1}`, env,
		`invalid field name 1 in struct initializer`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{1: 1, 1: "a"}
func TestCheckCompositeLitExprS1IntKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: 1, 1: "a"}`, env,
		`invalid field name 1 in struct initializer`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{1: 1, float32(1.5): 1}
func TestCheckCompositeLitExprS1IntKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: 1, float32(1.5): 1}`, env,
		`invalid field name 1 in struct initializer`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S1{1: 1, a: 1}
func TestCheckCompositeLitExprS1IntKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: 1, a: 1}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{1: 1, a: "a"}
func TestCheckCompositeLitExprS1IntKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: 1, a: "a"}`, env,
		`invalid field name 1 in struct initializer`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S1{1: 1, b: 1}
func TestCheckCompositeLitExprS1IntKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: 1, b: 1}`, env,
		`invalid field name 1 in struct initializer`,
		`unknown eval.s1 field 'b' in struct literal`,
	)

}

// Test S1{1: 1, c: 1}
func TestCheckCompositeLitExprS1IntKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: 1, c: 1}`, env,
		`invalid field name 1 in struct initializer`,
		`unknown eval.s1 field 'c' in struct literal`,
	)

}

// Test S1{1: true}
func TestCheckCompositeLitExprS1IntKBoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: true}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{1: true, 1}
func TestCheckCompositeLitExprS1IntKBoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: true, 1}`, env,
		`invalid field name 1 in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{1: true, "a"}
func TestCheckCompositeLitExprS1IntKBoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: true, "a"}`, env,
		`invalid field name 1 in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{1: true, nil}
func TestCheckCompositeLitExprS1IntKBoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: true, nil}`, env,
		`invalid field name 1 in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{1: true, 1: 1}
func TestCheckCompositeLitExprS1IntKBoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: true, 1: 1}`, env,
		`invalid field name 1 in struct initializer`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{1: true, 1: "a"}
func TestCheckCompositeLitExprS1IntKBoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: true, 1: "a"}`, env,
		`invalid field name 1 in struct initializer`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{1: true, float32(1.5): 1}
func TestCheckCompositeLitExprS1IntKBoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: true, float32(1.5): 1}`, env,
		`invalid field name 1 in struct initializer`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S1{1: true, a: 1}
func TestCheckCompositeLitExprS1IntKBoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: true, a: 1}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{1: true, a: "a"}
func TestCheckCompositeLitExprS1IntKBoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: true, a: "a"}`, env,
		`invalid field name 1 in struct initializer`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S1{1: true, b: 1}
func TestCheckCompositeLitExprS1IntKBoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: true, b: 1}`, env,
		`invalid field name 1 in struct initializer`,
		`unknown eval.s1 field 'b' in struct literal`,
	)

}

// Test S1{1: true, c: 1}
func TestCheckCompositeLitExprS1IntKBoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{1: true, c: 1}`, env,
		`invalid field name 1 in struct initializer`,
		`unknown eval.s1 field 'c' in struct literal`,
	)

}

// Test S1{float32(1.4): 1}
func TestCheckCompositeLitExprS1FloatKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{float32(1.4): 1}`, env,
		`invalid field name float32(1.4) in struct initializer`,
	)

}

// Test S1{float32(1.4): 1, 1}
func TestCheckCompositeLitExprS1FloatKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{float32(1.4): 1, 1}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{float32(1.4): 1, "a"}
func TestCheckCompositeLitExprS1FloatKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{float32(1.4): 1, "a"}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{float32(1.4): 1, nil}
func TestCheckCompositeLitExprS1FloatKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{float32(1.4): 1, nil}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{float32(1.4): 1, 1: 1}
func TestCheckCompositeLitExprS1FloatKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{float32(1.4): 1, 1: 1}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{float32(1.4): 1, 1: "a"}
func TestCheckCompositeLitExprS1FloatKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{float32(1.4): 1, 1: "a"}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{float32(1.4): 1, float32(1.5): 1}
func TestCheckCompositeLitExprS1FloatKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{float32(1.4): 1, float32(1.5): 1}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S1{float32(1.4): 1, a: 1}
func TestCheckCompositeLitExprS1FloatKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{float32(1.4): 1, a: 1}`, env,
		`invalid field name float32(1.4) in struct initializer`,
	)

}

// Test S1{float32(1.4): 1, a: "a"}
func TestCheckCompositeLitExprS1FloatKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{float32(1.4): 1, a: "a"}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S1{float32(1.4): 1, b: 1}
func TestCheckCompositeLitExprS1FloatKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{float32(1.4): 1, b: 1}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`unknown eval.s1 field 'b' in struct literal`,
	)

}

// Test S1{float32(1.4): 1, c: 1}
func TestCheckCompositeLitExprS1FloatKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{float32(1.4): 1, c: 1}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`unknown eval.s1 field 'c' in struct literal`,
	)

}

// Test S1{a: 1}
func TestCheckCompositeLitExprS1AKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `s1{a: 1}`, env, reflect.TypeOf(s1{a: 1}))
}

// Test S1{a: 1, 1}
func TestCheckCompositeLitExprS1AKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: 1, 1}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{a: 1, "a"}
func TestCheckCompositeLitExprS1AKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: 1, "a"}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{a: 1, nil}
func TestCheckCompositeLitExprS1AKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: 1, nil}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{a: 1, 1: 1}
func TestCheckCompositeLitExprS1AKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: 1, 1: 1}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{a: 1, 1: "a"}
func TestCheckCompositeLitExprS1AKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: 1, 1: "a"}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{a: 1, float32(1.5): 1}
func TestCheckCompositeLitExprS1AKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: 1, float32(1.5): 1}`, env,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S1{a: 1, a: 1}
func TestCheckCompositeLitExprS1AKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: 1, a: 1}`, env,
		`duplicate field name in struct literal: a`,
	)

}

// Test S1{a: 1, a: "a"}
func TestCheckCompositeLitExprS1AKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: 1, a: "a"}`, env,
		`duplicate field name in struct literal: a`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S1{a: 1, b: 1}
func TestCheckCompositeLitExprS1AKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: 1, b: 1}`, env,
		`unknown eval.s1 field 'b' in struct literal`,
	)

}

// Test S1{a: 1, c: 1}
func TestCheckCompositeLitExprS1AKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: 1, c: 1}`, env,
		`unknown eval.s1 field 'c' in struct literal`,
	)

}

// Test S1{a: true}
func TestCheckCompositeLitExprS1AKBoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: true}`, env,
		`cannot use true (type bool) as type int in field value`,
	)

}

// Test S1{a: true, 1}
func TestCheckCompositeLitExprS1AKBoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: true, 1}`, env,
		`cannot use true (type bool) as type int in field value`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{a: true, "a"}
func TestCheckCompositeLitExprS1AKBoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: true, "a"}`, env,
		`cannot use true (type bool) as type int in field value`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{a: true, nil}
func TestCheckCompositeLitExprS1AKBoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: true, nil}`, env,
		`cannot use true (type bool) as type int in field value`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{a: true, 1: 1}
func TestCheckCompositeLitExprS1AKBoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: true, 1: 1}`, env,
		`cannot use true (type bool) as type int in field value`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{a: true, 1: "a"}
func TestCheckCompositeLitExprS1AKBoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: true, 1: "a"}`, env,
		`cannot use true (type bool) as type int in field value`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{a: true, float32(1.5): 1}
func TestCheckCompositeLitExprS1AKBoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: true, float32(1.5): 1}`, env,
		`cannot use true (type bool) as type int in field value`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S1{a: true, a: 1}
func TestCheckCompositeLitExprS1AKBoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: true, a: 1}`, env,
		`cannot use true (type bool) as type int in field value`,
		`duplicate field name in struct literal: a`,
	)

}

// Test S1{a: true, a: "a"}
func TestCheckCompositeLitExprS1AKBoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: true, a: "a"}`, env,
		`cannot use true (type bool) as type int in field value`,
		`duplicate field name in struct literal: a`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S1{a: true, b: 1}
func TestCheckCompositeLitExprS1AKBoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: true, b: 1}`, env,
		`cannot use true (type bool) as type int in field value`,
		`unknown eval.s1 field 'b' in struct literal`,
	)

}

// Test S1{a: true, c: 1}
func TestCheckCompositeLitExprS1AKBoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{a: true, c: 1}`, env,
		`cannot use true (type bool) as type int in field value`,
		`unknown eval.s1 field 'c' in struct literal`,
	)

}

// Test S1{b: 1}
func TestCheckCompositeLitExprS1BKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{b: 1}`, env,
		`unknown eval.s1 field 'b' in struct literal`,
	)

}

// Test S1{b: 1, 1}
func TestCheckCompositeLitExprS1BKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{b: 1, 1}`, env,
		`unknown eval.s1 field 'b' in struct literal`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{b: 1, "a"}
func TestCheckCompositeLitExprS1BKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{b: 1, "a"}`, env,
		`unknown eval.s1 field 'b' in struct literal`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{b: 1, nil}
func TestCheckCompositeLitExprS1BKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{b: 1, nil}`, env,
		`unknown eval.s1 field 'b' in struct literal`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{b: 1, 1: 1}
func TestCheckCompositeLitExprS1BKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{b: 1, 1: 1}`, env,
		`unknown eval.s1 field 'b' in struct literal`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{b: 1, 1: "a"}
func TestCheckCompositeLitExprS1BKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{b: 1, 1: "a"}`, env,
		`unknown eval.s1 field 'b' in struct literal`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{b: 1, float32(1.5): 1}
func TestCheckCompositeLitExprS1BKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{b: 1, float32(1.5): 1}`, env,
		`unknown eval.s1 field 'b' in struct literal`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S1{b: 1, a: 1}
func TestCheckCompositeLitExprS1BKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{b: 1, a: 1}`, env,
		`unknown eval.s1 field 'b' in struct literal`,
	)

}

// Test S1{b: 1, a: "a"}
func TestCheckCompositeLitExprS1BKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{b: 1, a: "a"}`, env,
		`unknown eval.s1 field 'b' in struct literal`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S1{b: 1, b: 1}
func TestCheckCompositeLitExprS1BKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{b: 1, b: 1}`, env,
		`unknown eval.s1 field 'b' in struct literal`,
		`unknown eval.s1 field 'b' in struct literal`,
	)

}

// Test S1{b: 1, c: 1}
func TestCheckCompositeLitExprS1BKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{b: 1, c: 1}`, env,
		`unknown eval.s1 field 'b' in struct literal`,
		`unknown eval.s1 field 'c' in struct literal`,
	)

}

// Test S1{c: 1}
func TestCheckCompositeLitExprS1CKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{c: 1}`, env,
		`unknown eval.s1 field 'c' in struct literal`,
	)

}

// Test S1{c: 1, 1}
func TestCheckCompositeLitExprS1CKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{c: 1, 1}`, env,
		`unknown eval.s1 field 'c' in struct literal`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{c: 1, "a"}
func TestCheckCompositeLitExprS1CKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{c: 1, "a"}`, env,
		`unknown eval.s1 field 'c' in struct literal`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{c: 1, nil}
func TestCheckCompositeLitExprS1CKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{c: 1, nil}`, env,
		`unknown eval.s1 field 'c' in struct literal`,
		`mixture of field:value and value initializers`,
	)

}

// Test S1{c: 1, 1: 1}
func TestCheckCompositeLitExprS1CKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{c: 1, 1: 1}`, env,
		`unknown eval.s1 field 'c' in struct literal`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{c: 1, 1: "a"}
func TestCheckCompositeLitExprS1CKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{c: 1, 1: "a"}`, env,
		`unknown eval.s1 field 'c' in struct literal`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S1{c: 1, float32(1.5): 1}
func TestCheckCompositeLitExprS1CKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{c: 1, float32(1.5): 1}`, env,
		`unknown eval.s1 field 'c' in struct literal`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S1{c: 1, a: 1}
func TestCheckCompositeLitExprS1CKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{c: 1, a: 1}`, env,
		`unknown eval.s1 field 'c' in struct literal`,
	)

}

// Test S1{c: 1, a: "a"}
func TestCheckCompositeLitExprS1CKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{c: 1, a: "a"}`, env,
		`unknown eval.s1 field 'c' in struct literal`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S1{c: 1, b: 1}
func TestCheckCompositeLitExprS1CKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{c: 1, b: 1}`, env,
		`unknown eval.s1 field 'c' in struct literal`,
		`unknown eval.s1 field 'b' in struct literal`,
	)

}

// Test S1{c: 1, c: 1}
func TestCheckCompositeLitExprS1CKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s1{c: 1, c: 1}`, env,
		`unknown eval.s1 field 'c' in struct literal`,
		`unknown eval.s1 field 'c' in struct literal`,
	)

}

// Test S2{}
func TestCheckCompositeLitExprS2XX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `s2{}`, env, reflect.TypeOf(s2{}))
}

// Test S2{1}
func TestCheckCompositeLitExprS2XInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1}`, env,
		`too few values in struct initializer`,
	)

}

// Test S2{"a"}
func TestCheckCompositeLitExprS2XString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{"a"}`, env,
		`cannot use "a" (type string) as type int in field value`,
		`too few values in struct initializer`,
	)

}

// Test S2{nil}
func TestCheckCompositeLitExprS2XNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{nil}`, env,
		`cannot use nil as type int in field value`,
		`too few values in struct initializer`,
	)

}

// Test S2{1: 1}
func TestCheckCompositeLitExprS2XIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: 1}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{1: "a"}
func TestCheckCompositeLitExprS2XIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: "a"}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{float32(1.5): 1}
func TestCheckCompositeLitExprS2XFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{float32(1.5): 1}`, env,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S2{a: 1}
func TestCheckCompositeLitExprS2XAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `s2{a: 1}`, env, reflect.TypeOf(s2{a: 1}))
}

// Test S2{a: "a"}
func TestCheckCompositeLitExprS2XAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: "a"}`, env,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S2{b: 1}
func TestCheckCompositeLitExprS2XBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `s2{b: 1}`, env, reflect.TypeOf(s2{b: 1}))
}

// Test S2{c: 1}
func TestCheckCompositeLitExprS2XCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{c: 1}`, env,
		`unknown eval.s2 field 'c' in struct literal`,
	)

}

// Test S2{1}
func TestCheckCompositeLitExprS2IntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1}`, env,
		`too few values in struct initializer`,
	)

}

// Test S2{1, 1}
func TestCheckCompositeLitExprS2IntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `s2{1, 1}`, env, reflect.TypeOf(s2{1, 1}))
}

// Test S2{1, "a"}
func TestCheckCompositeLitExprS2IntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1, "a"}`, env,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S2{1, nil}
func TestCheckCompositeLitExprS2IntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1, nil}`, env,
		`cannot use nil as type int in field value`,
	)

}

// Test S2{1, 1: 1}
func TestCheckCompositeLitExprS2IntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1, 1: 1}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{1, 1: "a"}
func TestCheckCompositeLitExprS2IntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1, 1: "a"}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{1, float32(1.5): 1}
func TestCheckCompositeLitExprS2IntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1, float32(1.5): 1}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S2{1, a: 1}
func TestCheckCompositeLitExprS2IntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1, a: 1}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{1, a: "a"}
func TestCheckCompositeLitExprS2IntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1, a: "a"}`, env,
		`mixture of field:value and value initializers`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S2{1, b: 1}
func TestCheckCompositeLitExprS2IntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1, b: 1}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{1, c: 1}
func TestCheckCompositeLitExprS2IntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1, c: 1}`, env,
		`mixture of field:value and value initializers`,
		`unknown eval.s2 field 'c' in struct literal`,
	)

}

// Test S2{nil}
func TestCheckCompositeLitExprS2NilX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{nil}`, env,
		`cannot use nil as type int in field value`,
		`too few values in struct initializer`,
	)

}

// Test S2{nil, 1}
func TestCheckCompositeLitExprS2NilInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{nil, 1}`, env,
		`cannot use nil as type int in field value`,
	)

}

// Test S2{nil, "a"}
func TestCheckCompositeLitExprS2NilString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{nil, "a"}`, env,
		`cannot use nil as type int in field value`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S2{nil, nil}
func TestCheckCompositeLitExprS2NilNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{nil, nil}`, env,
		`cannot use nil as type int in field value`,
		`cannot use nil as type int in field value`,
	)

}

// Test S2{nil, 1: 1}
func TestCheckCompositeLitExprS2NilIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{nil, 1: 1}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{nil, 1: "a"}
func TestCheckCompositeLitExprS2NilIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{nil, 1: "a"}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{nil, float32(1.5): 1}
func TestCheckCompositeLitExprS2NilFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{nil, float32(1.5): 1}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S2{nil, a: 1}
func TestCheckCompositeLitExprS2NilAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{nil, a: 1}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{nil, a: "a"}
func TestCheckCompositeLitExprS2NilAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{nil, a: "a"}`, env,
		`mixture of field:value and value initializers`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S2{nil, b: 1}
func TestCheckCompositeLitExprS2NilBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{nil, b: 1}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{nil, c: 1}
func TestCheckCompositeLitExprS2NilCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{nil, c: 1}`, env,
		`mixture of field:value and value initializers`,
		`unknown eval.s2 field 'c' in struct literal`,
	)

}

// Test S2{true}
func TestCheckCompositeLitExprS2BoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{true}`, env,
		`cannot use true (type bool) as type int in field value`,
		`too few values in struct initializer`,
	)

}

// Test S2{true, 1}
func TestCheckCompositeLitExprS2BoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{true, 1}`, env,
		`cannot use true (type bool) as type int in field value`,
	)

}

// Test S2{true, "a"}
func TestCheckCompositeLitExprS2BoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{true, "a"}`, env,
		`cannot use true (type bool) as type int in field value`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S2{true, nil}
func TestCheckCompositeLitExprS2BoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{true, nil}`, env,
		`cannot use true (type bool) as type int in field value`,
		`cannot use nil as type int in field value`,
	)

}

// Test S2{true, 1: 1}
func TestCheckCompositeLitExprS2BoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{true, 1: 1}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{true, 1: "a"}
func TestCheckCompositeLitExprS2BoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{true, 1: "a"}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{true, float32(1.5): 1}
func TestCheckCompositeLitExprS2BoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{true, float32(1.5): 1}`, env,
		`mixture of field:value and value initializers`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S2{true, a: 1}
func TestCheckCompositeLitExprS2BoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{true, a: 1}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{true, a: "a"}
func TestCheckCompositeLitExprS2BoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{true, a: "a"}`, env,
		`mixture of field:value and value initializers`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S2{true, b: 1}
func TestCheckCompositeLitExprS2BoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{true, b: 1}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{true, c: 1}
func TestCheckCompositeLitExprS2BoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{true, c: 1}`, env,
		`mixture of field:value and value initializers`,
		`unknown eval.s2 field 'c' in struct literal`,
	)

}

// Test S2{1: 1}
func TestCheckCompositeLitExprS2IntKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: 1}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{1: 1, 1}
func TestCheckCompositeLitExprS2IntKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: 1, 1}`, env,
		`invalid field name 1 in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{1: 1, "a"}
func TestCheckCompositeLitExprS2IntKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: 1, "a"}`, env,
		`invalid field name 1 in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{1: 1, nil}
func TestCheckCompositeLitExprS2IntKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: 1, nil}`, env,
		`invalid field name 1 in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{1: 1, 1: 1}
func TestCheckCompositeLitExprS2IntKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: 1, 1: 1}`, env,
		`invalid field name 1 in struct initializer`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{1: 1, 1: "a"}
func TestCheckCompositeLitExprS2IntKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: 1, 1: "a"}`, env,
		`invalid field name 1 in struct initializer`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{1: 1, float32(1.5): 1}
func TestCheckCompositeLitExprS2IntKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: 1, float32(1.5): 1}`, env,
		`invalid field name 1 in struct initializer`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S2{1: 1, a: 1}
func TestCheckCompositeLitExprS2IntKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: 1, a: 1}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{1: 1, a: "a"}
func TestCheckCompositeLitExprS2IntKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: 1, a: "a"}`, env,
		`invalid field name 1 in struct initializer`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S2{1: 1, b: 1}
func TestCheckCompositeLitExprS2IntKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: 1, b: 1}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{1: 1, c: 1}
func TestCheckCompositeLitExprS2IntKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: 1, c: 1}`, env,
		`invalid field name 1 in struct initializer`,
		`unknown eval.s2 field 'c' in struct literal`,
	)

}

// Test S2{1: true}
func TestCheckCompositeLitExprS2IntKBoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: true}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{1: true, 1}
func TestCheckCompositeLitExprS2IntKBoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: true, 1}`, env,
		`invalid field name 1 in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{1: true, "a"}
func TestCheckCompositeLitExprS2IntKBoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: true, "a"}`, env,
		`invalid field name 1 in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{1: true, nil}
func TestCheckCompositeLitExprS2IntKBoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: true, nil}`, env,
		`invalid field name 1 in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{1: true, 1: 1}
func TestCheckCompositeLitExprS2IntKBoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: true, 1: 1}`, env,
		`invalid field name 1 in struct initializer`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{1: true, 1: "a"}
func TestCheckCompositeLitExprS2IntKBoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: true, 1: "a"}`, env,
		`invalid field name 1 in struct initializer`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{1: true, float32(1.5): 1}
func TestCheckCompositeLitExprS2IntKBoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: true, float32(1.5): 1}`, env,
		`invalid field name 1 in struct initializer`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S2{1: true, a: 1}
func TestCheckCompositeLitExprS2IntKBoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: true, a: 1}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{1: true, a: "a"}
func TestCheckCompositeLitExprS2IntKBoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: true, a: "a"}`, env,
		`invalid field name 1 in struct initializer`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S2{1: true, b: 1}
func TestCheckCompositeLitExprS2IntKBoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: true, b: 1}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{1: true, c: 1}
func TestCheckCompositeLitExprS2IntKBoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{1: true, c: 1}`, env,
		`invalid field name 1 in struct initializer`,
		`unknown eval.s2 field 'c' in struct literal`,
	)

}

// Test S2{float32(1.4): 1}
func TestCheckCompositeLitExprS2FloatKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{float32(1.4): 1}`, env,
		`invalid field name float32(1.4) in struct initializer`,
	)

}

// Test S2{float32(1.4): 1, 1}
func TestCheckCompositeLitExprS2FloatKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{float32(1.4): 1, 1}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{float32(1.4): 1, "a"}
func TestCheckCompositeLitExprS2FloatKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{float32(1.4): 1, "a"}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{float32(1.4): 1, nil}
func TestCheckCompositeLitExprS2FloatKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{float32(1.4): 1, nil}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{float32(1.4): 1, 1: 1}
func TestCheckCompositeLitExprS2FloatKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{float32(1.4): 1, 1: 1}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{float32(1.4): 1, 1: "a"}
func TestCheckCompositeLitExprS2FloatKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{float32(1.4): 1, 1: "a"}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{float32(1.4): 1, float32(1.5): 1}
func TestCheckCompositeLitExprS2FloatKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{float32(1.4): 1, float32(1.5): 1}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S2{float32(1.4): 1, a: 1}
func TestCheckCompositeLitExprS2FloatKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{float32(1.4): 1, a: 1}`, env,
		`invalid field name float32(1.4) in struct initializer`,
	)

}

// Test S2{float32(1.4): 1, a: "a"}
func TestCheckCompositeLitExprS2FloatKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{float32(1.4): 1, a: "a"}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S2{float32(1.4): 1, b: 1}
func TestCheckCompositeLitExprS2FloatKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{float32(1.4): 1, b: 1}`, env,
		`invalid field name float32(1.4) in struct initializer`,
	)

}

// Test S2{float32(1.4): 1, c: 1}
func TestCheckCompositeLitExprS2FloatKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{float32(1.4): 1, c: 1}`, env,
		`invalid field name float32(1.4) in struct initializer`,
		`unknown eval.s2 field 'c' in struct literal`,
	)

}

// Test S2{a: 1}
func TestCheckCompositeLitExprS2AKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `s2{a: 1}`, env, reflect.TypeOf(s2{a: 1}))
}

// Test S2{a: 1, 1}
func TestCheckCompositeLitExprS2AKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: 1, 1}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{a: 1, "a"}
func TestCheckCompositeLitExprS2AKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: 1, "a"}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{a: 1, nil}
func TestCheckCompositeLitExprS2AKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: 1, nil}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{a: 1, 1: 1}
func TestCheckCompositeLitExprS2AKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: 1, 1: 1}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{a: 1, 1: "a"}
func TestCheckCompositeLitExprS2AKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: 1, 1: "a"}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{a: 1, float32(1.5): 1}
func TestCheckCompositeLitExprS2AKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: 1, float32(1.5): 1}`, env,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S2{a: 1, a: 1}
func TestCheckCompositeLitExprS2AKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: 1, a: 1}`, env,
		`duplicate field name in struct literal: a`,
	)

}

// Test S2{a: 1, a: "a"}
func TestCheckCompositeLitExprS2AKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: 1, a: "a"}`, env,
		`duplicate field name in struct literal: a`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S2{a: 1, b: 1}
func TestCheckCompositeLitExprS2AKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `s2{a: 1, b: 1}`, env, reflect.TypeOf(s2{a: 1, b: 1}))
}

// Test S2{a: 1, c: 1}
func TestCheckCompositeLitExprS2AKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: 1, c: 1}`, env,
		`unknown eval.s2 field 'c' in struct literal`,
	)

}

// Test S2{a: true}
func TestCheckCompositeLitExprS2AKBoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: true}`, env,
		`cannot use true (type bool) as type int in field value`,
	)

}

// Test S2{a: true, 1}
func TestCheckCompositeLitExprS2AKBoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: true, 1}`, env,
		`cannot use true (type bool) as type int in field value`,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{a: true, "a"}
func TestCheckCompositeLitExprS2AKBoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: true, "a"}`, env,
		`cannot use true (type bool) as type int in field value`,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{a: true, nil}
func TestCheckCompositeLitExprS2AKBoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: true, nil}`, env,
		`cannot use true (type bool) as type int in field value`,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{a: true, 1: 1}
func TestCheckCompositeLitExprS2AKBoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: true, 1: 1}`, env,
		`cannot use true (type bool) as type int in field value`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{a: true, 1: "a"}
func TestCheckCompositeLitExprS2AKBoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: true, 1: "a"}`, env,
		`cannot use true (type bool) as type int in field value`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{a: true, float32(1.5): 1}
func TestCheckCompositeLitExprS2AKBoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: true, float32(1.5): 1}`, env,
		`cannot use true (type bool) as type int in field value`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S2{a: true, a: 1}
func TestCheckCompositeLitExprS2AKBoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: true, a: 1}`, env,
		`cannot use true (type bool) as type int in field value`,
		`duplicate field name in struct literal: a`,
	)

}

// Test S2{a: true, a: "a"}
func TestCheckCompositeLitExprS2AKBoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: true, a: "a"}`, env,
		`cannot use true (type bool) as type int in field value`,
		`duplicate field name in struct literal: a`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S2{a: true, b: 1}
func TestCheckCompositeLitExprS2AKBoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: true, b: 1}`, env,
		`cannot use true (type bool) as type int in field value`,
	)

}

// Test S2{a: true, c: 1}
func TestCheckCompositeLitExprS2AKBoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{a: true, c: 1}`, env,
		`cannot use true (type bool) as type int in field value`,
		`unknown eval.s2 field 'c' in struct literal`,
	)

}

// Test S2{b: 1}
func TestCheckCompositeLitExprS2BKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `s2{b: 1}`, env, reflect.TypeOf(s2{b: 1}))
}

// Test S2{b: 1, 1}
func TestCheckCompositeLitExprS2BKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{b: 1, 1}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{b: 1, "a"}
func TestCheckCompositeLitExprS2BKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{b: 1, "a"}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{b: 1, nil}
func TestCheckCompositeLitExprS2BKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{b: 1, nil}`, env,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{b: 1, 1: 1}
func TestCheckCompositeLitExprS2BKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{b: 1, 1: 1}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{b: 1, 1: "a"}
func TestCheckCompositeLitExprS2BKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{b: 1, 1: "a"}`, env,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{b: 1, float32(1.5): 1}
func TestCheckCompositeLitExprS2BKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{b: 1, float32(1.5): 1}`, env,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S2{b: 1, a: 1}
func TestCheckCompositeLitExprS2BKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `s2{b: 1, a: 1}`, env, reflect.TypeOf(s2{b: 1, a: 1}))
}

// Test S2{b: 1, a: "a"}
func TestCheckCompositeLitExprS2BKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{b: 1, a: "a"}`, env,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S2{b: 1, b: 1}
func TestCheckCompositeLitExprS2BKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{b: 1, b: 1}`, env,
		`duplicate field name in struct literal: b`,
	)

}

// Test S2{b: 1, c: 1}
func TestCheckCompositeLitExprS2BKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{b: 1, c: 1}`, env,
		`unknown eval.s2 field 'c' in struct literal`,
	)

}

// Test S2{c: 1}
func TestCheckCompositeLitExprS2CKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{c: 1}`, env,
		`unknown eval.s2 field 'c' in struct literal`,
	)

}

// Test S2{c: 1, 1}
func TestCheckCompositeLitExprS2CKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{c: 1, 1}`, env,
		`unknown eval.s2 field 'c' in struct literal`,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{c: 1, "a"}
func TestCheckCompositeLitExprS2CKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{c: 1, "a"}`, env,
		`unknown eval.s2 field 'c' in struct literal`,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{c: 1, nil}
func TestCheckCompositeLitExprS2CKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{c: 1, nil}`, env,
		`unknown eval.s2 field 'c' in struct literal`,
		`mixture of field:value and value initializers`,
	)

}

// Test S2{c: 1, 1: 1}
func TestCheckCompositeLitExprS2CKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{c: 1, 1: 1}`, env,
		`unknown eval.s2 field 'c' in struct literal`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{c: 1, 1: "a"}
func TestCheckCompositeLitExprS2CKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{c: 1, 1: "a"}`, env,
		`unknown eval.s2 field 'c' in struct literal`,
		`invalid field name 1 in struct initializer`,
	)

}

// Test S2{c: 1, float32(1.5): 1}
func TestCheckCompositeLitExprS2CKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{c: 1, float32(1.5): 1}`, env,
		`unknown eval.s2 field 'c' in struct literal`,
		`invalid field name float32(1.5) in struct initializer`,
	)

}

// Test S2{c: 1, a: 1}
func TestCheckCompositeLitExprS2CKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{c: 1, a: 1}`, env,
		`unknown eval.s2 field 'c' in struct literal`,
	)

}

// Test S2{c: 1, a: "a"}
func TestCheckCompositeLitExprS2CKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{c: 1, a: "a"}`, env,
		`unknown eval.s2 field 'c' in struct literal`,
		`cannot use "a" (type string) as type int in field value`,
	)

}

// Test S2{c: 1, b: 1}
func TestCheckCompositeLitExprS2CKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{c: 1, b: 1}`, env,
		`unknown eval.s2 field 'c' in struct literal`,
	)

}

// Test S2{c: 1, c: 1}
func TestCheckCompositeLitExprS2CKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `s2{c: 1, c: 1}`, env,
		`unknown eval.s2 field 'c' in struct literal`,
		`unknown eval.s2 field 'c' in struct literal`,
	)

}

// Test Slice{}
func TestCheckCompositeLitExprSliceXX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `[]int{}`, env, reflect.TypeOf([]int{}))
}

// Test Slice{1}
func TestCheckCompositeLitExprSliceXInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `[]int{1}`, env, reflect.TypeOf([]int{1}))
}

// Test Slice{"a"}
func TestCheckCompositeLitExprSliceXString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{"a"}`, env,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{nil}
func TestCheckCompositeLitExprSliceXNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{nil}`, env,
		`cannot convert nil to type int`,
	)

}

// Test Slice{1: 1}
func TestCheckCompositeLitExprSliceXIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `[]int{1: 1}`, env, reflect.TypeOf([]int{1: 1}))
}

// Test Slice{1: "a"}
func TestCheckCompositeLitExprSliceXIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: "a"}`, env,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{float32(1.5): 1}
func TestCheckCompositeLitExprSliceXFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{float32(1.5): 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{a: 1}
func TestCheckCompositeLitExprSliceXAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{a: "a"}
func TestCheckCompositeLitExprSliceXAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{b: 1}
func TestCheckCompositeLitExprSliceXBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{b: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{c: 1}
func TestCheckCompositeLitExprSliceXCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{c: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{1}
func TestCheckCompositeLitExprSliceIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `[]int{1}`, env, reflect.TypeOf([]int{1}))
}

// Test Slice{1, 1}
func TestCheckCompositeLitExprSliceIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `[]int{1, 1}`, env, reflect.TypeOf([]int{1, 1}))
}

// Test Slice{1, "a"}
func TestCheckCompositeLitExprSliceIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1, "a"}`, env,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{1, nil}
func TestCheckCompositeLitExprSliceIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1, nil}`, env,
		`cannot convert nil to type int`,
	)

}

// Test Slice{1, 1: 1}
func TestCheckCompositeLitExprSliceIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `[]int{1, 1: 1}`, env, reflect.TypeOf([]int{1, 1: 1}))
}

// Test Slice{1, 1: "a"}
func TestCheckCompositeLitExprSliceIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1, 1: "a"}`, env,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{1, float32(1.5): 1}
func TestCheckCompositeLitExprSliceIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1, float32(1.5): 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{1, a: 1}
func TestCheckCompositeLitExprSliceIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1, a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{1, a: "a"}
func TestCheckCompositeLitExprSliceIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1, a: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{1, b: 1}
func TestCheckCompositeLitExprSliceIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1, b: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{1, c: 1}
func TestCheckCompositeLitExprSliceIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1, c: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{nil}
func TestCheckCompositeLitExprSliceNilX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{nil}`, env,
		`cannot convert nil to type int`,
	)

}

// Test Slice{nil, 1}
func TestCheckCompositeLitExprSliceNilInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{nil, 1}`, env,
		`cannot convert nil to type int`,
	)

}

// Test Slice{nil, "a"}
func TestCheckCompositeLitExprSliceNilString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{nil, "a"}`, env,
		`cannot convert nil to type int`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{nil, nil}
func TestCheckCompositeLitExprSliceNilNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{nil, nil}`, env,
		`cannot convert nil to type int`,
		`cannot convert nil to type int`,
	)

}

// Test Slice{nil, 1: 1}
func TestCheckCompositeLitExprSliceNilIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{nil, 1: 1}`, env,
		`cannot convert nil to type int`,
	)

}

// Test Slice{nil, 1: "a"}
func TestCheckCompositeLitExprSliceNilIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{nil, 1: "a"}`, env,
		`cannot convert nil to type int`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{nil, float32(1.5): 1}
func TestCheckCompositeLitExprSliceNilFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{nil, float32(1.5): 1}`, env,
		`cannot convert nil to type int`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{nil, a: 1}
func TestCheckCompositeLitExprSliceNilAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{nil, a: 1}`, env,
		`cannot convert nil to type int`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{nil, a: "a"}
func TestCheckCompositeLitExprSliceNilAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{nil, a: "a"}`, env,
		`cannot convert nil to type int`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{nil, b: 1}
func TestCheckCompositeLitExprSliceNilBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{nil, b: 1}`, env,
		`cannot convert nil to type int`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{nil, c: 1}
func TestCheckCompositeLitExprSliceNilCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{nil, c: 1}`, env,
		`cannot convert nil to type int`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{true}
func TestCheckCompositeLitExprSliceBoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{true}`, env,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test Slice{true, 1}
func TestCheckCompositeLitExprSliceBoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{true, 1}`, env,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test Slice{true, "a"}
func TestCheckCompositeLitExprSliceBoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{true, "a"}`, env,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{true, nil}
func TestCheckCompositeLitExprSliceBoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{true, nil}`, env,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert nil to type int`,
	)

}

// Test Slice{true, 1: 1}
func TestCheckCompositeLitExprSliceBoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{true, 1: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test Slice{true, 1: "a"}
func TestCheckCompositeLitExprSliceBoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{true, 1: "a"}`, env,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{true, float32(1.5): 1}
func TestCheckCompositeLitExprSliceBoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{true, float32(1.5): 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{true, a: 1}
func TestCheckCompositeLitExprSliceBoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{true, a: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{true, a: "a"}
func TestCheckCompositeLitExprSliceBoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{true, a: "a"}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{true, b: 1}
func TestCheckCompositeLitExprSliceBoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{true, b: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{true, c: 1}
func TestCheckCompositeLitExprSliceBoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{true, c: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{1: 1}
func TestCheckCompositeLitExprSliceIntKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `[]int{1: 1}`, env, reflect.TypeOf([]int{1: 1}))
}

// Test Slice{1: 1, 1}
func TestCheckCompositeLitExprSliceIntKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `[]int{1: 1, 1}`, env, reflect.TypeOf([]int{1: 1, 1}))
}

// Test Slice{1: 1, "a"}
func TestCheckCompositeLitExprSliceIntKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: 1, "a"}`, env,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{1: 1, nil}
func TestCheckCompositeLitExprSliceIntKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: 1, nil}`, env,
		`cannot convert nil to type int`,
	)

}

// Test Slice{1: 1, 1: 1}
func TestCheckCompositeLitExprSliceIntKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: 1, 1: 1}`, env,
		`duplicate index in array literal: 1`,
	)

}

// Test Slice{1: 1, 1: "a"}
func TestCheckCompositeLitExprSliceIntKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: 1, 1: "a"}`, env,
		`duplicate index in array literal: 1`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{1: 1, float32(1.5): 1}
func TestCheckCompositeLitExprSliceIntKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: 1, float32(1.5): 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{1: 1, a: 1}
func TestCheckCompositeLitExprSliceIntKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: 1, a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{1: 1, a: "a"}
func TestCheckCompositeLitExprSliceIntKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: 1, a: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{1: 1, b: 1}
func TestCheckCompositeLitExprSliceIntKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: 1, b: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{1: 1, c: 1}
func TestCheckCompositeLitExprSliceIntKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: 1, c: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{1: true}
func TestCheckCompositeLitExprSliceIntKBoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: true}`, env,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test Slice{1: true, 1}
func TestCheckCompositeLitExprSliceIntKBoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: true, 1}`, env,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test Slice{1: true, "a"}
func TestCheckCompositeLitExprSliceIntKBoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: true, "a"}`, env,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{1: true, nil}
func TestCheckCompositeLitExprSliceIntKBoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: true, nil}`, env,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert nil to type int`,
	)

}

// Test Slice{1: true, 1: 1}
func TestCheckCompositeLitExprSliceIntKBoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: true, 1: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`duplicate index in array literal: 1`,
	)

}

// Test Slice{1: true, 1: "a"}
func TestCheckCompositeLitExprSliceIntKBoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: true, 1: "a"}`, env,
		`cannot use true (type bool) as type int in array element`,
		`duplicate index in array literal: 1`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{1: true, float32(1.5): 1}
func TestCheckCompositeLitExprSliceIntKBoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: true, float32(1.5): 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{1: true, a: 1}
func TestCheckCompositeLitExprSliceIntKBoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: true, a: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{1: true, a: "a"}
func TestCheckCompositeLitExprSliceIntKBoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: true, a: "a"}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{1: true, b: 1}
func TestCheckCompositeLitExprSliceIntKBoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: true, b: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{1: true, c: 1}
func TestCheckCompositeLitExprSliceIntKBoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{1: true, c: 1}`, env,
		`cannot use true (type bool) as type int in array element`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{float32(1.4): 1}
func TestCheckCompositeLitExprSliceFloatKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{float32(1.4): 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{float32(1.4): 1, 1}
func TestCheckCompositeLitExprSliceFloatKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{float32(1.4): 1, 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{float32(1.4): 1, "a"}
func TestCheckCompositeLitExprSliceFloatKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{float32(1.4): 1, "a"}`, env,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{float32(1.4): 1, nil}
func TestCheckCompositeLitExprSliceFloatKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{float32(1.4): 1, nil}`, env,
		`array index must be non-negative integer constant`,
		`cannot convert nil to type int`,
	)

}

// Test Slice{float32(1.4): 1, 1: 1}
func TestCheckCompositeLitExprSliceFloatKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{float32(1.4): 1, 1: 1}`, env,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{float32(1.4): 1, 1: "a"}
func TestCheckCompositeLitExprSliceFloatKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{float32(1.4): 1, 1: "a"}`, env,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{float32(1.4): 1, float32(1.5): 1}
func TestCheckCompositeLitExprSliceFloatKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{float32(1.4): 1, float32(1.5): 1}`, env,
		`array index must be non-negative integer constant`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{float32(1.4): 1, a: 1}
func TestCheckCompositeLitExprSliceFloatKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{float32(1.4): 1, a: 1}`, env,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{float32(1.4): 1, a: "a"}
func TestCheckCompositeLitExprSliceFloatKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{float32(1.4): 1, a: "a"}`, env,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{float32(1.4): 1, b: 1}
func TestCheckCompositeLitExprSliceFloatKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{float32(1.4): 1, b: 1}`, env,
		`array index must be non-negative integer constant`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{float32(1.4): 1, c: 1}
func TestCheckCompositeLitExprSliceFloatKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{float32(1.4): 1, c: 1}`, env,
		`array index must be non-negative integer constant`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{a: 1}
func TestCheckCompositeLitExprSliceAKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{a: 1, 1}
func TestCheckCompositeLitExprSliceAKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: 1, 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{a: 1, "a"}
func TestCheckCompositeLitExprSliceAKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: 1, "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{a: 1, nil}
func TestCheckCompositeLitExprSliceAKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: 1, nil}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert nil to type int`,
	)

}

// Test Slice{a: 1, 1: 1}
func TestCheckCompositeLitExprSliceAKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: 1, 1: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{a: 1, 1: "a"}
func TestCheckCompositeLitExprSliceAKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: 1, 1: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{a: 1, float32(1.5): 1}
func TestCheckCompositeLitExprSliceAKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: 1, float32(1.5): 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{a: 1, a: 1}
func TestCheckCompositeLitExprSliceAKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: 1, a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{a: 1, a: "a"}
func TestCheckCompositeLitExprSliceAKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: 1, a: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{a: 1, b: 1}
func TestCheckCompositeLitExprSliceAKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: 1, b: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{a: 1, c: 1}
func TestCheckCompositeLitExprSliceAKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: 1, c: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{a: true}
func TestCheckCompositeLitExprSliceAKBoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: true}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test Slice{a: true, 1}
func TestCheckCompositeLitExprSliceAKBoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: true, 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test Slice{a: true, "a"}
func TestCheckCompositeLitExprSliceAKBoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: true, "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{a: true, nil}
func TestCheckCompositeLitExprSliceAKBoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: true, nil}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert nil to type int`,
	)

}

// Test Slice{a: true, 1: 1}
func TestCheckCompositeLitExprSliceAKBoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: true, 1: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
	)

}

// Test Slice{a: true, 1: "a"}
func TestCheckCompositeLitExprSliceAKBoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: true, 1: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{a: true, float32(1.5): 1}
func TestCheckCompositeLitExprSliceAKBoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: true, float32(1.5): 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{a: true, a: 1}
func TestCheckCompositeLitExprSliceAKBoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: true, a: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{a: true, a: "a"}
func TestCheckCompositeLitExprSliceAKBoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: true, a: "a"}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{a: true, b: 1}
func TestCheckCompositeLitExprSliceAKBoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: true, b: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{a: true, c: 1}
func TestCheckCompositeLitExprSliceAKBoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{a: true, c: 1}`, env,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot use true (type bool) as type int in array element`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{b: 1}
func TestCheckCompositeLitExprSliceBKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{b: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{b: 1, 1}
func TestCheckCompositeLitExprSliceBKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{b: 1, 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{b: 1, "a"}
func TestCheckCompositeLitExprSliceBKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{b: 1, "a"}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{b: 1, nil}
func TestCheckCompositeLitExprSliceBKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{b: 1, nil}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`cannot convert nil to type int`,
	)

}

// Test Slice{b: 1, 1: 1}
func TestCheckCompositeLitExprSliceBKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{b: 1, 1: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{b: 1, 1: "a"}
func TestCheckCompositeLitExprSliceBKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{b: 1, 1: "a"}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{b: 1, float32(1.5): 1}
func TestCheckCompositeLitExprSliceBKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{b: 1, float32(1.5): 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{b: 1, a: 1}
func TestCheckCompositeLitExprSliceBKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{b: 1, a: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{b: 1, a: "a"}
func TestCheckCompositeLitExprSliceBKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{b: 1, a: "a"}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{b: 1, b: 1}
func TestCheckCompositeLitExprSliceBKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{b: 1, b: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{b: 1, c: 1}
func TestCheckCompositeLitExprSliceBKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{b: 1, c: 1}`, env,
		`undefined: b`,
		`array index must be non-negative integer constant`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{c: 1}
func TestCheckCompositeLitExprSliceCKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{c: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{c: 1, 1}
func TestCheckCompositeLitExprSliceCKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{c: 1, 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{c: 1, "a"}
func TestCheckCompositeLitExprSliceCKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{c: 1, "a"}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{c: 1, nil}
func TestCheckCompositeLitExprSliceCKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{c: 1, nil}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`cannot convert nil to type int`,
	)

}

// Test Slice{c: 1, 1: 1}
func TestCheckCompositeLitExprSliceCKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{c: 1, 1: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{c: 1, 1: "a"}
func TestCheckCompositeLitExprSliceCKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{c: 1, 1: "a"}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{c: 1, float32(1.5): 1}
func TestCheckCompositeLitExprSliceCKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{c: 1, float32(1.5): 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{c: 1, a: 1}
func TestCheckCompositeLitExprSliceCKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{c: 1, a: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{c: 1, a: "a"}
func TestCheckCompositeLitExprSliceCKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{c: 1, a: "a"}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`undefined: a`,
		`array index must be non-negative integer constant`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in array element`,
	)

}

// Test Slice{c: 1, b: 1}
func TestCheckCompositeLitExprSliceCKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{c: 1, b: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`undefined: b`,
		`array index must be non-negative integer constant`,
	)

}

// Test Slice{c: 1, c: 1}
func TestCheckCompositeLitExprSliceCKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `[]int{c: 1, c: 1}`, env,
		`undefined: c`,
		`array index must be non-negative integer constant`,
		`undefined: c`,
		`array index must be non-negative integer constant`,
	)

}

// Test Map{}
func TestCheckCompositeLitExprMapXX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `map[int] int{}`, env, reflect.TypeOf(map[int] int{}))
}

// Test Map{1}
func TestCheckCompositeLitExprMapXInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1}`, env,
		`missing key in map literal`,
	)

}

// Test Map{"a"}
func TestCheckCompositeLitExprMapXString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{"a"}`, env,
		`missing key in map literal`,
	)

}

// Test Map{nil}
func TestCheckCompositeLitExprMapXNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{nil}`, env,
		`missing key in map literal`,
	)

}

// Test Map{1: 1}
func TestCheckCompositeLitExprMapXIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `map[int] int{1: 1}`, env, reflect.TypeOf(map[int] int{1: 1}))
}

// Test Map{1: "a"}
func TestCheckCompositeLitExprMapXIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: "a"}`, env,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{float32(1.5): 1}
func TestCheckCompositeLitExprMapXFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{float32(1.5): 1}`, env,
		`cannot use float32(1.5) (type float32) as type int in map key`,
	)

}

// Test Map{a: 1}
func TestCheckCompositeLitExprMapXAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: 1}`, env,
		`undefined: a`,
	)

}

// Test Map{a: "a"}
func TestCheckCompositeLitExprMapXAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: "a"}`, env,
		`undefined: a`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{b: 1}
func TestCheckCompositeLitExprMapXBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{b: 1}`, env,
		`undefined: b`,
	)

}

// Test Map{c: 1}
func TestCheckCompositeLitExprMapXCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{c: 1}`, env,
		`undefined: c`,
	)

}

// Test Map{1}
func TestCheckCompositeLitExprMapIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1}`, env,
		`missing key in map literal`,
	)

}

// Test Map{1, 1}
func TestCheckCompositeLitExprMapIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1, 1}`, env,
		`missing key in map literal`,
		`missing key in map literal`,
	)

}

// Test Map{1, "a"}
func TestCheckCompositeLitExprMapIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1, "a"}`, env,
		`missing key in map literal`,
		`missing key in map literal`,
	)

}

// Test Map{1, nil}
func TestCheckCompositeLitExprMapIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1, nil}`, env,
		`missing key in map literal`,
		`missing key in map literal`,
	)

}

// Test Map{1, 1: 1}
func TestCheckCompositeLitExprMapIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1, 1: 1}`, env,
		`missing key in map literal`,
	)

}

// Test Map{1, 1: "a"}
func TestCheckCompositeLitExprMapIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1, 1: "a"}`, env,
		`missing key in map literal`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{1, float32(1.5): 1}
func TestCheckCompositeLitExprMapIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1, float32(1.5): 1}`, env,
		`missing key in map literal`,
		`cannot use float32(1.5) (type float32) as type int in map key`,
	)

}

// Test Map{1, a: 1}
func TestCheckCompositeLitExprMapIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1, a: 1}`, env,
		`missing key in map literal`,
		`undefined: a`,
	)

}

// Test Map{1, a: "a"}
func TestCheckCompositeLitExprMapIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1, a: "a"}`, env,
		`missing key in map literal`,
		`undefined: a`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{1, b: 1}
func TestCheckCompositeLitExprMapIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1, b: 1}`, env,
		`missing key in map literal`,
		`undefined: b`,
	)

}

// Test Map{1, c: 1}
func TestCheckCompositeLitExprMapIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1, c: 1}`, env,
		`missing key in map literal`,
		`undefined: c`,
	)

}

// Test Map{nil}
func TestCheckCompositeLitExprMapNilX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{nil}`, env,
		`missing key in map literal`,
	)

}

// Test Map{nil, 1}
func TestCheckCompositeLitExprMapNilInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{nil, 1}`, env,
		`missing key in map literal`,
		`missing key in map literal`,
	)

}

// Test Map{nil, "a"}
func TestCheckCompositeLitExprMapNilString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{nil, "a"}`, env,
		`missing key in map literal`,
		`missing key in map literal`,
	)

}

// Test Map{nil, nil}
func TestCheckCompositeLitExprMapNilNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{nil, nil}`, env,
		`missing key in map literal`,
		`missing key in map literal`,
	)

}

// Test Map{nil, 1: 1}
func TestCheckCompositeLitExprMapNilIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{nil, 1: 1}`, env,
		`missing key in map literal`,
	)

}

// Test Map{nil, 1: "a"}
func TestCheckCompositeLitExprMapNilIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{nil, 1: "a"}`, env,
		`missing key in map literal`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{nil, float32(1.5): 1}
func TestCheckCompositeLitExprMapNilFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{nil, float32(1.5): 1}`, env,
		`missing key in map literal`,
		`cannot use float32(1.5) (type float32) as type int in map key`,
	)

}

// Test Map{nil, a: 1}
func TestCheckCompositeLitExprMapNilAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{nil, a: 1}`, env,
		`missing key in map literal`,
		`undefined: a`,
	)

}

// Test Map{nil, a: "a"}
func TestCheckCompositeLitExprMapNilAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{nil, a: "a"}`, env,
		`missing key in map literal`,
		`undefined: a`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{nil, b: 1}
func TestCheckCompositeLitExprMapNilBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{nil, b: 1}`, env,
		`missing key in map literal`,
		`undefined: b`,
	)

}

// Test Map{nil, c: 1}
func TestCheckCompositeLitExprMapNilCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{nil, c: 1}`, env,
		`missing key in map literal`,
		`undefined: c`,
	)

}

// Test Map{true}
func TestCheckCompositeLitExprMapBoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{true}`, env,
		`missing key in map literal`,
	)

}

// Test Map{true, 1}
func TestCheckCompositeLitExprMapBoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{true, 1}`, env,
		`missing key in map literal`,
		`missing key in map literal`,
	)

}

// Test Map{true, "a"}
func TestCheckCompositeLitExprMapBoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{true, "a"}`, env,
		`missing key in map literal`,
		`missing key in map literal`,
	)

}

// Test Map{true, nil}
func TestCheckCompositeLitExprMapBoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{true, nil}`, env,
		`missing key in map literal`,
		`missing key in map literal`,
	)

}

// Test Map{true, 1: 1}
func TestCheckCompositeLitExprMapBoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{true, 1: 1}`, env,
		`missing key in map literal`,
	)

}

// Test Map{true, 1: "a"}
func TestCheckCompositeLitExprMapBoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{true, 1: "a"}`, env,
		`missing key in map literal`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{true, float32(1.5): 1}
func TestCheckCompositeLitExprMapBoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{true, float32(1.5): 1}`, env,
		`missing key in map literal`,
		`cannot use float32(1.5) (type float32) as type int in map key`,
	)

}

// Test Map{true, a: 1}
func TestCheckCompositeLitExprMapBoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{true, a: 1}`, env,
		`missing key in map literal`,
		`undefined: a`,
	)

}

// Test Map{true, a: "a"}
func TestCheckCompositeLitExprMapBoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{true, a: "a"}`, env,
		`missing key in map literal`,
		`undefined: a`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{true, b: 1}
func TestCheckCompositeLitExprMapBoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{true, b: 1}`, env,
		`missing key in map literal`,
		`undefined: b`,
	)

}

// Test Map{true, c: 1}
func TestCheckCompositeLitExprMapBoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{true, c: 1}`, env,
		`missing key in map literal`,
		`undefined: c`,
	)

}

// Test Map{1: 1}
func TestCheckCompositeLitExprMapIntKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectType(t, `map[int] int{1: 1}`, env, reflect.TypeOf(map[int] int{1: 1}))
}

// Test Map{1: 1, 1}
func TestCheckCompositeLitExprMapIntKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: 1, 1}`, env,
		`missing key in map literal`,
	)

}

// Test Map{1: 1, "a"}
func TestCheckCompositeLitExprMapIntKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: 1, "a"}`, env,
		`missing key in map literal`,
	)

}

// Test Map{1: 1, nil}
func TestCheckCompositeLitExprMapIntKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: 1, nil}`, env,
		`missing key in map literal`,
	)

}

// Test Map{1: 1, 1: 1}
func TestCheckCompositeLitExprMapIntKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: 1, 1: 1}`, env,
		`duplicate key 1 in map literal`,
	)

}

// Test Map{1: 1, 1: "a"}
func TestCheckCompositeLitExprMapIntKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: 1, 1: "a"}`, env,
		`duplicate key 1 in map literal`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{1: 1, float32(1.5): 1}
func TestCheckCompositeLitExprMapIntKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: 1, float32(1.5): 1}`, env,
		`cannot use float32(1.5) (type float32) as type int in map key`,
	)

}

// Test Map{1: 1, a: 1}
func TestCheckCompositeLitExprMapIntKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: 1, a: 1}`, env,
		`undefined: a`,
	)

}

// Test Map{1: 1, a: "a"}
func TestCheckCompositeLitExprMapIntKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: 1, a: "a"}`, env,
		`undefined: a`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{1: 1, b: 1}
func TestCheckCompositeLitExprMapIntKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: 1, b: 1}`, env,
		`undefined: b`,
	)

}

// Test Map{1: 1, c: 1}
func TestCheckCompositeLitExprMapIntKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: 1, c: 1}`, env,
		`undefined: c`,
	)

}

// Test Map{1: true}
func TestCheckCompositeLitExprMapIntKBoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: true}`, env,
		`cannot use true (type bool) as type int in map value`,
	)

}

// Test Map{1: true, 1}
func TestCheckCompositeLitExprMapIntKBoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: true, 1}`, env,
		`cannot use true (type bool) as type int in map value`,
		`missing key in map literal`,
	)

}

// Test Map{1: true, "a"}
func TestCheckCompositeLitExprMapIntKBoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: true, "a"}`, env,
		`cannot use true (type bool) as type int in map value`,
		`missing key in map literal`,
	)

}

// Test Map{1: true, nil}
func TestCheckCompositeLitExprMapIntKBoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: true, nil}`, env,
		`cannot use true (type bool) as type int in map value`,
		`missing key in map literal`,
	)

}

// Test Map{1: true, 1: 1}
func TestCheckCompositeLitExprMapIntKBoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: true, 1: 1}`, env,
		`cannot use true (type bool) as type int in map value`,
		`duplicate key 1 in map literal`,
	)

}

// Test Map{1: true, 1: "a"}
func TestCheckCompositeLitExprMapIntKBoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: true, 1: "a"}`, env,
		`cannot use true (type bool) as type int in map value`,
		`duplicate key 1 in map literal`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{1: true, float32(1.5): 1}
func TestCheckCompositeLitExprMapIntKBoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: true, float32(1.5): 1}`, env,
		`cannot use true (type bool) as type int in map value`,
		`cannot use float32(1.5) (type float32) as type int in map key`,
	)

}

// Test Map{1: true, a: 1}
func TestCheckCompositeLitExprMapIntKBoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: true, a: 1}`, env,
		`cannot use true (type bool) as type int in map value`,
		`undefined: a`,
	)

}

// Test Map{1: true, a: "a"}
func TestCheckCompositeLitExprMapIntKBoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: true, a: "a"}`, env,
		`cannot use true (type bool) as type int in map value`,
		`undefined: a`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{1: true, b: 1}
func TestCheckCompositeLitExprMapIntKBoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: true, b: 1}`, env,
		`cannot use true (type bool) as type int in map value`,
		`undefined: b`,
	)

}

// Test Map{1: true, c: 1}
func TestCheckCompositeLitExprMapIntKBoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{1: true, c: 1}`, env,
		`cannot use true (type bool) as type int in map value`,
		`undefined: c`,
	)

}

// Test Map{float32(1.4): 1}
func TestCheckCompositeLitExprMapFloatKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{float32(1.4): 1}`, env,
		`cannot use float32(1.4) (type float32) as type int in map key`,
	)

}

// Test Map{float32(1.4): 1, 1}
func TestCheckCompositeLitExprMapFloatKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{float32(1.4): 1, 1}`, env,
		`cannot use float32(1.4) (type float32) as type int in map key`,
		`missing key in map literal`,
	)

}

// Test Map{float32(1.4): 1, "a"}
func TestCheckCompositeLitExprMapFloatKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{float32(1.4): 1, "a"}`, env,
		`cannot use float32(1.4) (type float32) as type int in map key`,
		`missing key in map literal`,
	)

}

// Test Map{float32(1.4): 1, nil}
func TestCheckCompositeLitExprMapFloatKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{float32(1.4): 1, nil}`, env,
		`cannot use float32(1.4) (type float32) as type int in map key`,
		`missing key in map literal`,
	)

}

// Test Map{float32(1.4): 1, 1: 1}
func TestCheckCompositeLitExprMapFloatKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{float32(1.4): 1, 1: 1}`, env,
		`cannot use float32(1.4) (type float32) as type int in map key`,
	)

}

// Test Map{float32(1.4): 1, 1: "a"}
func TestCheckCompositeLitExprMapFloatKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{float32(1.4): 1, 1: "a"}`, env,
		`cannot use float32(1.4) (type float32) as type int in map key`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{float32(1.4): 1, float32(1.5): 1}
func TestCheckCompositeLitExprMapFloatKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{float32(1.4): 1, float32(1.5): 1}`, env,
		`cannot use float32(1.4) (type float32) as type int in map key`,
		`cannot use float32(1.5) (type float32) as type int in map key`,
	)

}

// Test Map{float32(1.4): 1, a: 1}
func TestCheckCompositeLitExprMapFloatKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{float32(1.4): 1, a: 1}`, env,
		`cannot use float32(1.4) (type float32) as type int in map key`,
		`undefined: a`,
	)

}

// Test Map{float32(1.4): 1, a: "a"}
func TestCheckCompositeLitExprMapFloatKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{float32(1.4): 1, a: "a"}`, env,
		`cannot use float32(1.4) (type float32) as type int in map key`,
		`undefined: a`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{float32(1.4): 1, b: 1}
func TestCheckCompositeLitExprMapFloatKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{float32(1.4): 1, b: 1}`, env,
		`cannot use float32(1.4) (type float32) as type int in map key`,
		`undefined: b`,
	)

}

// Test Map{float32(1.4): 1, c: 1}
func TestCheckCompositeLitExprMapFloatKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{float32(1.4): 1, c: 1}`, env,
		`cannot use float32(1.4) (type float32) as type int in map key`,
		`undefined: c`,
	)

}

// Test Map{a: 1}
func TestCheckCompositeLitExprMapAKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: 1}`, env,
		`undefined: a`,
	)

}

// Test Map{a: 1, 1}
func TestCheckCompositeLitExprMapAKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: 1, 1}`, env,
		`undefined: a`,
		`missing key in map literal`,
	)

}

// Test Map{a: 1, "a"}
func TestCheckCompositeLitExprMapAKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: 1, "a"}`, env,
		`undefined: a`,
		`missing key in map literal`,
	)

}

// Test Map{a: 1, nil}
func TestCheckCompositeLitExprMapAKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: 1, nil}`, env,
		`undefined: a`,
		`missing key in map literal`,
	)

}

// Test Map{a: 1, 1: 1}
func TestCheckCompositeLitExprMapAKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: 1, 1: 1}`, env,
		`undefined: a`,
	)

}

// Test Map{a: 1, 1: "a"}
func TestCheckCompositeLitExprMapAKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: 1, 1: "a"}`, env,
		`undefined: a`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{a: 1, float32(1.5): 1}
func TestCheckCompositeLitExprMapAKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: 1, float32(1.5): 1}`, env,
		`undefined: a`,
		`cannot use float32(1.5) (type float32) as type int in map key`,
	)

}

// Test Map{a: 1, a: 1}
func TestCheckCompositeLitExprMapAKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: 1, a: 1}`, env,
		`undefined: a`,
		`undefined: a`,
	)

}

// Test Map{a: 1, a: "a"}
func TestCheckCompositeLitExprMapAKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: 1, a: "a"}`, env,
		`undefined: a`,
		`undefined: a`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{a: 1, b: 1}
func TestCheckCompositeLitExprMapAKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: 1, b: 1}`, env,
		`undefined: a`,
		`undefined: b`,
	)

}

// Test Map{a: 1, c: 1}
func TestCheckCompositeLitExprMapAKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: 1, c: 1}`, env,
		`undefined: a`,
		`undefined: c`,
	)

}

// Test Map{a: true}
func TestCheckCompositeLitExprMapAKBoolX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: true}`, env,
		`undefined: a`,
		`cannot use true (type bool) as type int in map value`,
	)

}

// Test Map{a: true, 1}
func TestCheckCompositeLitExprMapAKBoolInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: true, 1}`, env,
		`undefined: a`,
		`cannot use true (type bool) as type int in map value`,
		`missing key in map literal`,
	)

}

// Test Map{a: true, "a"}
func TestCheckCompositeLitExprMapAKBoolString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: true, "a"}`, env,
		`undefined: a`,
		`cannot use true (type bool) as type int in map value`,
		`missing key in map literal`,
	)

}

// Test Map{a: true, nil}
func TestCheckCompositeLitExprMapAKBoolNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: true, nil}`, env,
		`undefined: a`,
		`cannot use true (type bool) as type int in map value`,
		`missing key in map literal`,
	)

}

// Test Map{a: true, 1: 1}
func TestCheckCompositeLitExprMapAKBoolIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: true, 1: 1}`, env,
		`undefined: a`,
		`cannot use true (type bool) as type int in map value`,
	)

}

// Test Map{a: true, 1: "a"}
func TestCheckCompositeLitExprMapAKBoolIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: true, 1: "a"}`, env,
		`undefined: a`,
		`cannot use true (type bool) as type int in map value`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{a: true, float32(1.5): 1}
func TestCheckCompositeLitExprMapAKBoolFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: true, float32(1.5): 1}`, env,
		`undefined: a`,
		`cannot use true (type bool) as type int in map value`,
		`cannot use float32(1.5) (type float32) as type int in map key`,
	)

}

// Test Map{a: true, a: 1}
func TestCheckCompositeLitExprMapAKBoolAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: true, a: 1}`, env,
		`undefined: a`,
		`cannot use true (type bool) as type int in map value`,
		`undefined: a`,
	)

}

// Test Map{a: true, a: "a"}
func TestCheckCompositeLitExprMapAKBoolAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: true, a: "a"}`, env,
		`undefined: a`,
		`cannot use true (type bool) as type int in map value`,
		`undefined: a`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{a: true, b: 1}
func TestCheckCompositeLitExprMapAKBoolBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: true, b: 1}`, env,
		`undefined: a`,
		`cannot use true (type bool) as type int in map value`,
		`undefined: b`,
	)

}

// Test Map{a: true, c: 1}
func TestCheckCompositeLitExprMapAKBoolCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{a: true, c: 1}`, env,
		`undefined: a`,
		`cannot use true (type bool) as type int in map value`,
		`undefined: c`,
	)

}

// Test Map{b: 1}
func TestCheckCompositeLitExprMapBKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{b: 1}`, env,
		`undefined: b`,
	)

}

// Test Map{b: 1, 1}
func TestCheckCompositeLitExprMapBKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{b: 1, 1}`, env,
		`undefined: b`,
		`missing key in map literal`,
	)

}

// Test Map{b: 1, "a"}
func TestCheckCompositeLitExprMapBKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{b: 1, "a"}`, env,
		`undefined: b`,
		`missing key in map literal`,
	)

}

// Test Map{b: 1, nil}
func TestCheckCompositeLitExprMapBKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{b: 1, nil}`, env,
		`undefined: b`,
		`missing key in map literal`,
	)

}

// Test Map{b: 1, 1: 1}
func TestCheckCompositeLitExprMapBKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{b: 1, 1: 1}`, env,
		`undefined: b`,
	)

}

// Test Map{b: 1, 1: "a"}
func TestCheckCompositeLitExprMapBKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{b: 1, 1: "a"}`, env,
		`undefined: b`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{b: 1, float32(1.5): 1}
func TestCheckCompositeLitExprMapBKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{b: 1, float32(1.5): 1}`, env,
		`undefined: b`,
		`cannot use float32(1.5) (type float32) as type int in map key`,
	)

}

// Test Map{b: 1, a: 1}
func TestCheckCompositeLitExprMapBKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{b: 1, a: 1}`, env,
		`undefined: b`,
		`undefined: a`,
	)

}

// Test Map{b: 1, a: "a"}
func TestCheckCompositeLitExprMapBKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{b: 1, a: "a"}`, env,
		`undefined: b`,
		`undefined: a`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{b: 1, b: 1}
func TestCheckCompositeLitExprMapBKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{b: 1, b: 1}`, env,
		`undefined: b`,
		`undefined: b`,
	)

}

// Test Map{b: 1, c: 1}
func TestCheckCompositeLitExprMapBKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{b: 1, c: 1}`, env,
		`undefined: b`,
		`undefined: c`,
	)

}

// Test Map{c: 1}
func TestCheckCompositeLitExprMapCKIntX(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{c: 1}`, env,
		`undefined: c`,
	)

}

// Test Map{c: 1, 1}
func TestCheckCompositeLitExprMapCKIntInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{c: 1, 1}`, env,
		`undefined: c`,
		`missing key in map literal`,
	)

}

// Test Map{c: 1, "a"}
func TestCheckCompositeLitExprMapCKIntString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{c: 1, "a"}`, env,
		`undefined: c`,
		`missing key in map literal`,
	)

}

// Test Map{c: 1, nil}
func TestCheckCompositeLitExprMapCKIntNil(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{c: 1, nil}`, env,
		`undefined: c`,
		`missing key in map literal`,
	)

}

// Test Map{c: 1, 1: 1}
func TestCheckCompositeLitExprMapCKIntIntKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{c: 1, 1: 1}`, env,
		`undefined: c`,
	)

}

// Test Map{c: 1, 1: "a"}
func TestCheckCompositeLitExprMapCKIntIntKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{c: 1, 1: "a"}`, env,
		`undefined: c`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{c: 1, float32(1.5): 1}
func TestCheckCompositeLitExprMapCKIntFloatKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{c: 1, float32(1.5): 1}`, env,
		`undefined: c`,
		`cannot use float32(1.5) (type float32) as type int in map key`,
	)

}

// Test Map{c: 1, a: 1}
func TestCheckCompositeLitExprMapCKIntAKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{c: 1, a: 1}`, env,
		`undefined: c`,
		`undefined: a`,
	)

}

// Test Map{c: 1, a: "a"}
func TestCheckCompositeLitExprMapCKIntAKString(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{c: 1, a: "a"}`, env,
		`undefined: c`,
		`undefined: a`,
		`cannot convert "a" to type int`,
		`cannot use "a" (type string) as type int in map value`,
	)

}

// Test Map{c: 1, b: 1}
func TestCheckCompositeLitExprMapCKIntBKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{c: 1, b: 1}`, env,
		`undefined: c`,
		`undefined: b`,
	)

}

// Test Map{c: 1, c: 1}
func TestCheckCompositeLitExprMapCKIntCKInt(t *testing.T) {
	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})


	expectCheckError(t, `map[int] int{c: 1, c: 1}`, env,
		`undefined: c`,
		`undefined: c`,
	)

}
