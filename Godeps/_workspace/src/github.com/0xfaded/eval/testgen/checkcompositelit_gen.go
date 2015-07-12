package main

import (
	"io"
	"strings"
	"text/template"
	"github.com/0xfaded/go-testgen"
)

type Test struct{}

var comment = template.Must(template.New("Comment").Parse(
`// Test {{ .Comment }}
`))

var defs =
`	type a1 [1]int
	type a2 [2]int
	type s1 struct{ a int }
	type s2 struct{ a, b int }
`
var body = template.Must(template.New("Body").Parse(defs +
`	env := MakeSimpleEnv()
	env.Types["a1"] = reflect.TypeOf(a1{})
	env.Types["a2"] = reflect.TypeOf(a2{})
	env.Types["s1"] = reflect.TypeOf(s1{})
	env.Types["s2"] = reflect.TypeOf(s2{})

{{ if .Errors }}{{ if .TestErrs }}
	expectCheckError(t, `+"`{{ .Expr }}`"+`, env,{{ range .Errors }}
		`+"`{{ . }}`"+`,{{ end }}
	){{ else }}	_ = env{{ end }}
{{ else }}
	expectType(t, `+"`{{ .Expr }}`"+`, env, reflect.TypeOf({{ .Expr }})){{ end }}
`))

func (*Test) Package() string {
	return "eval"
}

func (*Test) Prefix() string {
	return "CheckCompositeLitExpr"
}

func (*Test) Imports() map[string]string {
	return map[string]string { "reflect": "" }
}

func (*Test) Dimensions() []testgen.Dimension {
	funs := []testgen.Element{
		{"A1", "a1"},
		{"A2", "a2"},
		{"S1", "s1"},
		{"S2", "s2"},
		{"Slice", "[]int"},
		{"Map", "map[int] int"},
	}
	arg0 := []testgen.Element{
		{"X", ""},
		{"Int", "1"},
		{"Nil", "nil"},
		{"Bool", "true"},
		{"IntKInt", "1: 1"},
		{"IntKBool", "1: true"},
		{"FloatKInt", "float32(1.4): 1"},
		{"AKInt", "a: 1"},
		{"AKBool", "a: true"},
		{"BKInt", "b: 1"},
		{"CKInt", "c: 1"},
	}
	arg1 := []testgen.Element{
		{"X", ""},
		{"Int", "1"},
		{"String", "\"a\""},
		{"Nil", "nil"},
		{"IntKInt", "1: 1"},
		{"IntKString", "1: \"a\""},
		{"FloatKInt", "float32(1.5): 1"},
		{"AKInt", "a: 1"},
		{"AKString", "a: \"a\""},
		{"BKInt", "b: 1"},
		{"CKInt", "c: 1"},
	}
	return []testgen.Dimension{
		funs,
		arg0,
		arg1,
	}
}

func (*Test) Globals(w io.Writer) error {
	return nil
}

func (*Test) Comment(w io.Writer, elts ...testgen.Element) error {
	fun := elts[0].Name
	sep := "{"
	for _, elt := range elts[1:] {
		if elt.Value != "" {
			fun += sep + elt.Value.(string)
			sep = ", "
		}
	}
	if sep == "{" {
		fun += "{"
	}
	fun += "}"

	vars := map[string] interface{} {
		"Comment": fun,
	}

	return comment.Execute(w, vars)
}

func (*Test) Body(w io.Writer, elts ...testgen.Element) error {
	expr := elts[0].Value.(string)
	sep := "{"
	for _, elt := range elts[1:] {
		if elt.Value != "" {
			expr += sep + elt.Value.(string)
			sep = ", "
		}
	}
	if sep == "{" {
		expr += "{"
	}
	expr += "}"

	compileErrs, err := compileExprWithDefs(expr, defs)
	if err != nil {
		return err
	}

	n0, n1, n2 := elts[0].Name, elts[1].Name, elts[2].Name
	n1HasKey := !(n1 == "X" || n1 == "Int" || n1 == "String" || n1 == "Bool" || n1 == "Nil")
	n2HasKey := !(n2 == "X" || n2 == "Int" || n2 == "String" || n2 == "Bool" || n2 == "Nil")
	seenBadArrayIndex := false
	testErrs := true
	for i := 0; i < len(compileErrs); i += 1 {
		if strings.Index(compileErrs[i], "syntax error") != -1 {
			testErrs = false
		}
		compileErrs[i] = strings.Replace(compileErrs[i], "a1", "eval.a1", -1)
		compileErrs[i] = strings.Replace(compileErrs[i], "a2", "eval.a2", -1)
		compileErrs[i] = strings.Replace(compileErrs[i], "s1", "eval.s1", -1)
		compileErrs[i] = strings.Replace(compileErrs[i], "s2", "eval.s2", -1)

		// TODO[crc] Fix bugs for gc 1.2 bugs. Remove after upgrade to 1.3
		if n0 == "A2" {
			compileErrs[i] = strings.Replace(compileErrs[i], "[0:0]", "[0:2]", -1)
		} else if n0 == "A1" {
			compileErrs[i] = strings.Replace(compileErrs[i], "[0:0]", "[0:1]", -1)
		}

		// Fix for bug http://code.google.com/p/go/issues/detail?id=7153
		if !n2HasKey && strings.HasPrefix(compileErrs[i], "array index must be non-") {
			if seenBadArrayIndex {
				// remove the error
				compileErrs = append(compileErrs[:i], compileErrs[i+1:]...)
				i -= 1
			}
			seenBadArrayIndex = true
		}
	}

	// gc strips duplicate errors. This is understandable but we want to catch all errors.
	if len(compileErrs) > 0 && (n0 == "S1" || n0 == "S2") &&
		strings.HasPrefix(n1, "IntK") && strings.HasPrefix(n2, "IntK") {
		compileErrs = append(compileErrs[:1], append([]string{compileErrs[0]}, compileErrs[1:]...)...)
	}
	if len(compileErrs) == 1 && strings.HasPrefix(compileErrs[0], "unknown") && n1 == n2 {
		// duplicate `unknown eval.XX field 'X' in struct literal`
		compileErrs = append(compileErrs, compileErrs[0])
	}
	if len(compileErrs) == 1 && n1 == "Nil" && n2 == "Nil" {
		// duplicate `cannot use nil as type int in field value`
		compileErrs = append(compileErrs, compileErrs[0])
	}
	if len(compileErrs) == 1 &&
		compileErrs[0] == "array index must be non-negative integer constant" &&
		n1HasKey && n2HasKey &&
		n1 != "IntKInt" && n1 != "IntKBool" && n2 != "IntKInt" && n2 != "IntKBool" {
		compileErrs = append(compileErrs, compileErrs[0])
	}
	if len(compileErrs) == 2 &&
		strings.HasPrefix(compileErrs[0], "undefined") &&
		compileErrs[1] == "array index must be non-negative integer constant" &&
		n1HasKey && n2HasKey &&
		n1 != "IntKInt" && n1 != "IntKBool" && n2 != "IntKInt" && n2 != "IntKBool" {
		compileErrs = append(compileErrs, compileErrs[1])
	}
	if len(compileErrs) > 0 && n0 == "Map" && !n1HasKey && !n2HasKey &&
		n1 != "X" && n2 != "X" && !(n1 == "Nil" && n2 == "Nil") {
		compileErrs = append(compileErrs, compileErrs[0])
	}
	if len(compileErrs) == 1 && n0 == "Map" && n1HasKey && n2HasKey &&
		strings.HasPrefix(compileErrs[0], "undefined") &&
		n1[0] == n2[0] {

		compileErrs = append(compileErrs[:1], append([]string{compileErrs[0]}, compileErrs[1:]...)...)
	}
	if n0 == "Map" && n1 == "AKInt" && n2 == "AKString" {
		compileErrs = append(compileErrs[:1], append([]string{compileErrs[0]}, compileErrs[1:]...)...)
	}
	vars := map[string] interface{} {
		"Expr": expr,
		"Errors": compileErrs,
		"TestErrs": testErrs,
	}

	return body.Execute(w, &vars)
}

