package eval

import (
	"testing"
	"reflect"

	"go/parser"
)

func errorPosEqual(a, b []string) bool {
	if len(a) != len(b) { return false }
	for i, aVal := range a {
		if aVal != b[i] { return false }
	}
	return true
}

func TestFormatErrorPos(t *testing.T) {
	source  := `split(os.Args ", )")`
	errmsg  := `1:15: expected ')', found 'STRING' ", "`
	results := FormatErrorPos(source, errmsg)
	expect  := []string { source,  "--------------^" }
	if !errorPosEqual(expect, results) {
		t.Fatalf("Expected %v, got %v", expect, results)
	}

	source  = "`"
	errmsg  = `1:1: string not terminated`
	results = FormatErrorPos(source, errmsg)
	expect  = []string { source,  "^" }

	source  = "y("
	errmsg  = `1:3: expected ')', found 'EOF'`
	results = FormatErrorPos(source, errmsg)
	expect  = []string { source,  "--^" }

}

func TestIsPkgAddressable(t *testing.T) {
	env := MakeSimpleEnv()
	env.Pkgs["a"] = MakeSimpleEnv()
	env.Pkgs["a"].(*SimpleEnv).Vars["v"] = reflect.ValueOf(new(int))
	env.Pkgs["a"].(*SimpleEnv).Funcs["f"] = reflect.ValueOf(func() {})
	v, _ := parser.ParseExpr("a.v")
	vv, _ := CheckExpr(v, env)
	if !isAddressable(vv) {
		t.Fatalf("expected package var 'a.v' to be addressible")
	}

	f, _ := parser.ParseExpr("a.f")
	ff, _ := CheckExpr(f, env)
	if isAddressable(ff) {
		t.Fatalf("expected package func 'a.f' to be unaddressible")
	}
}
