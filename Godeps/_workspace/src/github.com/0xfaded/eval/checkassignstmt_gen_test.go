package eval

import (
	"testing"
	"reflect"
)

// Test TooMany
func TestCheckAssignStmtTooMany(t *testing.T) {
	env := MakeSimpleEnv()
	f := func() (int, int) { return 1, 1 }
	env.Vars["f"] = reflect.ValueOf(&f)
	expectCheckError(t, `_, _ = 1, 2, 3`, env,
		`assignment count mismatch: 2 = 3`,
	)
}

// Test TooFew
func TestCheckAssignStmtTooFew(t *testing.T) {
	env := MakeSimpleEnv()
	f := func() (int, int) { return 1, 1 }
	env.Vars["f"] = reflect.ValueOf(&f)
	expectCheckError(t, `_, _ = 1`, env,
		`assignment count mismatch: 2 = 1`,
	)
}

// Test NoNewIdents1
func TestCheckAssignStmtNoNewIdents1(t *testing.T) {
	env := MakeSimpleEnv()
	f := func() (int, int) { return 1, 1 }
	env.Vars["f"] = reflect.ValueOf(&f)
	expectCheckError(t, `f := nil`, env,
		`no new variables on left side of :=`,
	)
}

// Test NoNewIdents2
func TestCheckAssignStmtNoNewIdents2(t *testing.T) {
	env := MakeSimpleEnv()
	f := func() (int, int) { return 1, 1 }
	env.Vars["f"] = reflect.ValueOf(&f)
	expectCheckError(t, `f, _ := nil, 1`, env,
		`no new variables on left side of :=`,
	)
}

// Test UnderscoreNil
func TestCheckAssignStmtUnderscoreNil(t *testing.T) {
	env := MakeSimpleEnv()
	f := func() (int, int) { return 1, 1 }
	env.Vars["f"] = reflect.ValueOf(&f)
	expectCheckError(t, `_ = nil`, env,
		`use of untyped nil`,
	)
}

// Test Unaddressable
func TestCheckAssignStmtUnaddressable(t *testing.T) {
	env := MakeSimpleEnv()
	f := func() (int, int) { return 1, 1 }
	env.Vars["f"] = reflect.ValueOf(&f)
	expectCheckError(t, `1 = 1`, env,
		`cannot assign to 1`,
		`cannot use 1 (type int) as type untyped number in assignment`,
	)
}

// Test Unaddressable2
func TestCheckAssignStmtUnaddressable2(t *testing.T) {
	env := MakeSimpleEnv()
	f := func() (int, int) { return 1, 1 }
	env.Vars["f"] = reflect.ValueOf(&f)
	expectCheckError(t, `1, 2 = 1, 2`, env,
		`cannot assign to 1`,
		`cannot assign to 2`,
		`cannot use 1 (type int) as type untyped number in assignment`,
		`cannot use 2 (type int) as type untyped number in assignment`,
	)
}

// Test ToNil
func TestCheckAssignStmtToNil(t *testing.T) {
	env := MakeSimpleEnv()
	f := func() (int, int) { return 1, 1 }
	env.Vars["f"] = reflect.ValueOf(&f)
	expectCheckError(t, `nil = 1`, env,
		`cannot assign to nil`,
		`cannot use 1 (type int) as type nil in assignment`,
	)
}

// Test Mistyped1
func TestCheckAssignStmtMistyped1(t *testing.T) {
	env := MakeSimpleEnv()
	f := func() (int, int) { return 1, 1 }
	env.Vars["f"] = reflect.ValueOf(&f)
	expectCheckError(t, `f = true`, env,
		`cannot use true (type bool) as type func() (int, int) in assignment`,
	)
}

// Test Mistyped2
func TestCheckAssignStmtMistyped2(t *testing.T) {
	env := MakeSimpleEnv()
	f := func() (int, int) { return 1, 1 }
	env.Vars["f"] = reflect.ValueOf(&f)
	expectCheckError(t, `f, f = true, false`, env,
		`cannot use true (type bool) as type func() (int, int) in assignment`,
		`cannot use false (type bool) as type func() (int, int) in assignment`,
	)
}
