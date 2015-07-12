package eval

import (
        "reflect"
        "testing"
)

// Test Int32 - Int32
func TestCheckUnaryTypedExpr(t *testing.T) {
	env := MakeSimpleEnv()
        x := int32(1)
        env.Vars["x"] = reflect.ValueOf(&x)

	expectType(t, `-x`, env, reflect.TypeOf(-x))
}
