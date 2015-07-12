package main

import (
	"fmt"
	"io"
	"text/template"
	"github.com/0xfaded/go-testgen"
)

type Test struct{}

var comment = template.Must(template.New("Comment").Parse(
`// Test {{ .Amp.Value }}
`))

var defs =
`
	a := 1
	b := &a
	_ = b
`
var body = template.Must(template.New("Body").Parse(defs +
`	env := MakeSimpleEnv()
	env.Vars["a"] = reflect.ValueOf(&a)
{{ if .Errors }}
	expectCheckError(t, `+"`{{ .Expr }}`"+`, env,{{ range .Errors }}
		`+"`{{ . }}`"+`,{{ end }}
	)
{{ else }}
	expectType(t, `+"`{{ .Expr }}`"+`, env, reflect.TypeOf({{ .Expr }})){{ end }}
`))

func (*Test) Package() string {
	return "eval"
}

func (*Test) Prefix() string {
	return "CheckAddrExpr"
}

func (*Test) Imports() map[string]string {
	return map[string]string { "reflect": "" }
}

func (*Test) Dimensions() []testgen.Dimension {
	stars := []testgen.Element{
		{"A", "a"},
		{"Int", "int(1)"},
		{"Number", "1.4"},
		{"Rune", "'a'"},
		{"Bool", "true"},
		{"String", `"a"`},
		{"Nil", "nil"},
		{"AtA", " &a"},
		{"StarB", " *a"},
		{"Slice", "[]int{1}"},
		{"SliceElt", "[]int{1}[0]"},
	}
	return []testgen.Dimension{
		stars,
	}
}

func (*Test) Globals(w io.Writer) error {
	return nil
}

func (*Test) Comment(w io.Writer, elts ...testgen.Element) error {
	vars := map[string] interface{} {
		"Amp": elts[0],
	}

	return comment.Execute(w, vars)
}

func (*Test) Body(w io.Writer, elts ...testgen.Element) error {
	expr := fmt.Sprintf("&%v", elts[0].Value)

	compileErrs, err := compileExprWithDefs(expr, defs)
	if err != nil {
		return err
	}

	vars := map[string] interface{} {
		"Expr": expr,
		"Errors": compileErrs,
	}

	return body.Execute(w, &vars)
}

