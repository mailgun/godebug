package eval

// Utilities for other tests live here

import (
	"fmt"
	"testing"
	"reflect"

	"go/ast"
	"go/parser"
)

func getResults(t *testing.T, expr string, env Env) []reflect.Value {
	if e, err := parser.ParseExpr(expr); err != nil {
		t.Fatalf("Failed to parse expression '%s' (%v)", expr, err)
	} else if aexpr, errs := CheckExpr(e, env); errs != nil {
		t.Fatalf("Failed to check expression '%s' (%v)", expr, errs)
	} else if results, err := EvalExpr(aexpr, env); err != nil {
		t.Fatalf("Error evaluating expression '%s' (%v)", expr, err)
	} else {
		return results
	}
	return nil
}

func expectResult(t *testing.T, expr string, env Env, expected interface{}) {
	expectResults(t, expr, env, expected)
}

func expectResults(t *testing.T, expr string, env Env, expected ...interface{}) {
	results := getResults(t, expr, env)
	if nil == results {
		if expected != nil {
			t.Fatalf("Expression '%s' is nil but expected '%+v'", expr, expected)
		}
		return
	} else if expected == nil {
		t.Fatalf("Expression '%s'expected is '%+v', expected to be nil", expr, results)
	}
	resultsi := make([]interface{}, len(results))
	for i, result := range results {
		resultsi[i] = result.Interface()
	}
	if !reflect.DeepEqual(resultsi, expected) {
		t.Fatalf("Expression '%s' yielded '%+v', expected '%+v'", expr, resultsi, expected)
	}
}

func expectPanic(t *testing.T, expr string, env Env, panicString string) {
	if e, err := parser.ParseExpr(expr); err != nil {
		t.Fatalf("Failed to parse expression '%s' (%v)", expr, err)
	} else if aexpr, errs := CheckExpr(e, env); errs != nil {
		t.Fatalf("Failed to check expression '%s' (%v)", expr, errs)
	} else if _, err := EvalExpr(aexpr, env); err == nil {
		t.Fatalf("Expected expression '%s' to panic", expr)
	} else if err.Error() != panicString {
		t.Fatalf("Panic `%s` != Expected `%s`", err.Error(), panicString)
	}
}

func expectConst(t *testing.T, expr string, env Env, expected interface{}, expectedType reflect.Type) {
	if e, err := parser.ParseExpr(expr); err != nil {
		t.Fatalf("Failed to parse expression '%s' (%v)", expr, err)
	} else if aexpr, errs := CheckExpr(e, env); errs != nil {
		t.Fatalf("Failed to check expression '%s' (%v)", expr, errs)
	} else if !aexpr.IsConst() {
		t.Fatalf("Expression '%s' did not yield a const node(%+v)", expr, aexpr)
	} else if expectedNumber, ok := expected.(*ConstNumber); ok {
		if actual, ok2 := aexpr.Const().Interface().(*ConstNumber); !ok2 {
			t.Fatalf("Expression '%s' yielded '%v', expected '%v'", expr, aexpr.Const(), expected)
		} else if !actual.Value.Equals(&expectedNumber.Value) {
			t.Fatalf("Expression '%s' yielded '%v', expected '%v'", expr, actual, expected)
		} else if len(aexpr.KnownType()) == 0 {
			t.Fatalf("Expression '%s' expected to have type '%v'", expr, expectedType)
		} else if actual := aexpr.KnownType()[0]; !typesEqual(expectedType, actual) {
			t.Fatalf("Expression '%s' has type '%v', expected '%v'", expr, actual, expectedType)
		}
	} else {
		if actual := aexpr.Const().Interface(); !reflect.DeepEqual(actual, expected) {
			t.Fatalf("Expression '%s' yielded '%+v', expected '%+v'", expr, actual, expected)
		} else if len(aexpr.KnownType()) == 0 {
			t.Fatalf("Expression '%s' expected to have type '%v'", expr, expectedType)
		} else if actual := aexpr.KnownType()[0]; !typesEqual(expectedType, actual) {
			t.Fatalf("Expression '%s' has type '%v', expected '%v'", expr, aexpr.KnownType()[0], expectedType)
		}
	}
}

func expectType(t *testing.T, expr string, env Env, expectedType reflect.Type) {
	if e, err := parser.ParseExpr(expr); err != nil {
		t.Fatalf("Failed to parse expression '%s' (%v)", expr, err)
	} else if aexpr, errs := CheckExpr(e, env); errs != nil {
		t.Fatalf("Failed to check expression '%s' (%v)", expr, errs)
	} else if aexpr.IsConst() {
		t.Fatalf("Expression '%s' yielded a const node(%+v)", expr, aexpr)
	} else if len(aexpr.KnownType()) != 1 {
		t.Fatalf("Expression '%s' expected to have single type '%v'", expr, expectedType)
	} else if actual := aexpr.KnownType()[0]; !typesEqual(expectedType, actual) {
		t.Fatalf("Expression '%s' has type '%v', expected '%v'", expr, aexpr.KnownType()[0], expectedType)
	}
}
func expectCheckError(t *testing.T, expr string, env Env, errorString ...string) {
	var errs []error
	if s, err := ParseStmt(expr); err != nil {
		t.Fatalf("Failed to parse expression '%s' (%v)", expr, err)
	} else if e, ok := s.(*ast.ExprStmt); ok {
		_, errs = CheckExpr(e.X, env)
	} else {
		_, errs = checkStmt(s, env, checkCtx{})
	}
	if errs != nil {
		var i int
		out := "\n"
		ok := true
		for i = 0; i < len(errorString); i += 1 {
			if i >= len(errs) {
				out += fmt.Sprintf("%d. Expected `%v` missing\n", i, errorString[i])
				ok = false
			} else if errorString[i] == errs[i].Error() {
				out += fmt.Sprintf("%d. Expected `%v` == `%v`\n", i, errorString[i], errs[i])
			} else {
				out += fmt.Sprintf("%d. Expected `%v` != `%v`\n", i, errorString[i], errs[i])
				ok = false
			}
		}
		for ; i < len(errs); i += 1 {
			out += fmt.Sprintf("%d. Unexpected `%v`\n", i, errs[i])
			ok = false
		}
		if !ok {
			t.Fatalf("%sWrong check errors for expression '%s'", out, expr)
		}
	} else {
		for i, s := range errorString {
			t.Logf("%d. Expected `%v` missing\n", i, s)
		}
		t.Fatalf("Missing check errors for expression '%s'", expr )
	}
}

func expectInterp(t *testing.T, stmt string, env Env) {
	if s, err := ParseStmt(stmt); err != nil {
		t.Fatalf("Failed to parse stmt '%s' (%v)", stmt, err)
	} else if c, errs := checkStmt(s, env, checkCtx{}); errs != nil {
		t.Logf("Failed to check stmt '%s'", stmt)
		for _, err := range errs {
			t.Logf("\t%v", err)
		}
		t.FailNow()
	} else if _, panik := InterpStmt(c, env); panik != nil {
		t.Fatalf("Statement '%s' panicked with %v", stmt, panik)
	}
}

func typesEqual(expected, actual reflect.Type) bool {
	var unwrapped reflect.Type
	switch t := actual.(type) {
	case Rune:
		unwrapped = t.Type
	default:
		unwrapped = actual
	}
	return reflect.DeepEqual(expected, unwrapped)
}

