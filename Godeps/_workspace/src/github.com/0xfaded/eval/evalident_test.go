package eval

import (
	"reflect"
	"testing"
)

func TestStringVar(t *testing.T) {
	env := MakeSimpleEnv()
        s := "abc"
	env.Vars["arg0"] = reflect.ValueOf(&s)
	expectResult(t, "arg0", env, "abc")
}
