package main

import (
	"fmt"
	"io"
	"text/template"
	"github.com/0xfaded/go-testgen"
)

type Test struct{}

var comment = template.Must(template.New("Comment").Parse(
`// Test {{ .Type.Value }}({{ .Value.Value }})
`))

var body = template.Must(template.New("Body").Parse(
`	env := MakeSimpleEnv()
{{ if .Errors }}{{ if .TestErrs }}
	expectCheckError(t, `+"`{{ .Expr }}`"+`, env,{{ range .Errors }}
		`+"`{{ . }}`"+`,{{ end }}
	){{ else }}	_ = env{{ end }}
{{ else }}
	expectConst(t, `+"`{{ .Expr }}`"+`, env, {{ .Expr }}, reflect.TypeOf({{ .Expr }})){{ end }}
`))

func (*Test) Package() string {
	return "eval"
}

func (*Test) Prefix() string {
	return "CheckCallExpr"
}

func (*Test) Imports() map[string]string {
	return map[string]string { "reflect": "" }
}

func (*Test) Dimensions() []testgen.Dimension {
	types := []testgen.Element{
		{"Int8", "int8"},
		{"Int16", "int16"},
		{"Int32", "int32"},
		{"Int64", "int64"},
		{"Uint8", "uint8"},
		{"Uint16", "uint16"},
		{"Uint32", "uint32"},
		{"Uint64", "uint64"},
		{"Float32", "float32"},
		{"Float64", "float64"},
		{"Complex64", "complex64"},
		{"Complex128", "complex128"},
		{"Rune", "rune"},
		{"Bool", "bool"},
		{"String", "string"},
		{"Nil", "nil"},
	}
	values := []testgen.Element{
		{"From7bits", "0x7f"},
		{"From8bits", "0xff"},
		{"From15bits", "0x7fff"},
		{"From16bits", "0xffff"},
		{"From31bits", "0x7fffffff"},
		{"From32bits", "0xffffffff"},
		{"From63bits", "0x7fffffffffffffff"},
		{"From64bits", "0xffffffffffffffff"},
		{"FromRune", "'d'"},
		{"FromWideRune", "'日'"},
		{"FromFloatingInt", "1.0"},
		{"FromFloating", "1.5"},
		{"FromComplexInt", "1.0+0i"},
		{"FromComplexFloat", "1.5+0i"},
		{"FromComplex", "1.5+1.5i"},
		{"FromBool", "true"},
		{"FromString", `"abc"`},
		{"FromNil", `nil`},
	}
	return []testgen.Dimension{
		types,
		values,
	}
}

func (*Test) Globals(w io.Writer) error {
	return nil
}

func (*Test) Comment(w io.Writer, elts ...testgen.Element) error {
	vars := map[string] interface{} {
		"Type": elts[0],
		"Value": elts[1],
	}

	return comment.Execute(w, vars)
}

func (*Test) Body(w io.Writer, elts ...testgen.Element) error {
	expr := fmt.Sprintf("%v(%v)", elts[0].Value, elts[1].Value)

	compileErrs, err := compileExpr(expr)
	if err != nil {
		return err
	}

	// https://github.com/0xfaded/go-interactive/issues/20
	testErrs := true
	if elts[0].Name == "Bool" {
		testErrs = false
	}

	vars := map[string] interface{} {
		"Expr": expr,
		"Errors": compileErrs,
		"TestErrs": testErrs,
	}

	return body.Execute(w, &vars)
}

