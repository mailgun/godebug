package main

import (
	"fmt"
	"io"
	"text/template"
	"go/token"
	"github.com/0xfaded/go-testgen"
)


type Test struct{}

var comment = template.Must(template.New("Comment").Parse(
`// Test {{ .Op.Value }} {{ .Rhs.Name }}
`))

var body = template.Must(template.New("Body").Parse(
`	env := MakeSimpleEnv()
{{ if .Errors }}
	expectCheckError(t, `+"`{{ .Expr }}`"+`, env,{{ range .Errors }}
		`+"`{{ . }}`"+`,{{ end }}
	)
{{ else }}
	expectConst(t, `+"`{{ .Expr }}`"+`, env, {{ .Expr }}, reflect.TypeOf({{ .Expr }})){{ end }}
`))

func (*Test) Package() string {
	return "eval"
}

func (*Test) Prefix() string {
	return "CheckUnaryTypedExpr"
}

func (*Test) Imports() map[string]string {
	return map[string]string {"reflect": ""}
}

func (*Test) Dimensions() []testgen.Dimension {
	// All numeric values have been chosen to be a power of two, for good reason :)
	types := []testgen.Element{
		{"Int32", "int32(4)"},
		//TODO[crc] go1.1 doesnt pick this up. Add after upgrade to 1.2
		// {"int32Overflow", "int32(-0x80000000)"},
		{"Float64", "float64(2)"},
		{"Complex128", "complex128(8i)"},
		{"Bool", "bool(true)"},
		{"String", `string("abc")`},
	}
	ops := []testgen.Element{
		{"Add", token.ADD},
		{"Sub", token.SUB},
		{"Xor", token.XOR},
		{"Not", token.NOT},
	}
	return []testgen.Dimension{
		ops,
		types,
	}
}

func (*Test) Globals(w io.Writer) error {
	return nil
}

func (*Test) Comment(w io.Writer, elts ...testgen.Element) error {
	vars := map[string] interface{} {
		"Op": elts[0],
		"Rhs": elts[1],
	}

	return comment.Execute(w, vars)
}

func (*Test) Body(w io.Writer, elts ...testgen.Element) error {
	op  := elts[0].Value.(token.Token)

	expr := fmt.Sprintf("%v %v", op, elts[1].Value)
	compileErrs, err := compileExpr(expr)
	if err != nil {
		return err
	}

	vars := map[string] interface{} {
		"Expr": expr,
		"Errors": compileErrs,
		"Op": elts[1],
	}

	return body.Execute(w, &vars)
}

