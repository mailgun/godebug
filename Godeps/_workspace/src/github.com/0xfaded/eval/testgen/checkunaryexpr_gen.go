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
	expectConst(t, `+"`{{ .Expr }}`"+`, env, {{ .NewConstType }}({{ .Expr }}), {{ .ResultType }}){{ end }}
`))

func (*Test) Package() string {
	return "eval"
}

func (*Test) Prefix() string {
	return "CheckUnaryExpr"
}

func (*Test) Imports() map[string]string {
	return nil
}

func (*Test) Dimensions() []testgen.Dimension {
	// All numeric values have been chosen to be a power of two, for good reason :)
	types := []testgen.Element{
		{"Int", "4"},
		{"Rune", "'@'"},
		{"Float", "2.0"},
		{"Complex", "8.0i"},
		{"Bool", "true"},
		{"String", `"abc"`},
		{"Nil", "nil"},
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
	rhs := elts[1].Name

	expr := fmt.Sprintf("%v %v", op, elts[1].Value)
	compileErrs, err := compileExpr(expr)
	if err != nil {
		return err
	}

	var newConstType string
	var resultType string

	switch rhs {
	case "Int":
		newConstType = "NewConstInt64"
		resultType = "ConstInt"
	case "Rune":
		newConstType = "NewConstRune"
		resultType = "ConstRune"
	case "Float":
		newConstType = "NewConstFloat64"
		resultType = "ConstFloat"
	case "Complex":
		newConstType = "NewConstComplex128"
		resultType = "ConstComplex"
	case "Bool":
		resultType = "ConstBool"
	}

	vars := map[string] interface{} {
		"Expr": expr,
		"Errors": compileErrs,
		"Op": elts[1],
		"NewConstType": newConstType,
		"ResultType": resultType,
	}

	return body.Execute(w, &vars)
}

