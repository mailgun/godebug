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
`// Test {{ .Lhs.Name }} {{ .Op.Value }} {{ .Rhs.Name }}
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
	return "CheckBinaryExpr"
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
		{"Mul", token.MUL},
		{"Quo", token.QUO},
		{"Rem", token.REM},
		{"And", token.AND},
		{"Or", token.OR},
		{"Xor", token.XOR},
		{"AndNot",  token.AND_NOT},
		{"Eql", token.EQL},
		{"Neq", token.NEQ},
		{"Leq", token.LEQ},
		{"Geq", token.GEQ},
		{"Lss", token.LSS},
		{"Gtr", token.GTR},
		{"Rhl", token.SHR},
	}
	return []testgen.Dimension{
		types,
		ops,
		types,
	}
}

func (*Test) Globals(w io.Writer) error {
	return nil
}

func (*Test) Comment(w io.Writer, elts ...testgen.Element) error {
	vars := map[string] interface{} {
		"Lhs": elts[0],
		"Op": elts[1],
		"Rhs": elts[2],
	}

	return comment.Execute(w, vars)
}

func (*Test) Body(w io.Writer, elts ...testgen.Element) error {
	lhs := elts[0].Name
	op  := elts[1].Value.(token.Token)
	rhs := elts[2].Name

	expr := fmt.Sprintf("%v %v %v", elts[0].Value, op, elts[2].Value)
	compileErrs, err := compileExpr(expr)
	if err != nil {
		return err
	}

	var newConstType string
	var resultType string

	switch lhs {
	case "Int":
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
		}
	case "Rune":
		switch rhs {
		case "Int", "Rune":
			newConstType = "NewConstRune"
			resultType = "ConstRune"
		case "Float":
			newConstType = "NewConstFloat64"
			resultType = "ConstFloat"
		case "Complex":
			newConstType = "NewConstComplex128"
			resultType = "ConstComplex"
		}
	case "Float":
		switch rhs {
		case "Int", "Rune", "Float":
			newConstType = "NewConstFloat64"
			resultType = "ConstFloat"
		case "Complex":
			newConstType = "NewConstComplex128"
			resultType = "ConstComplex"
		}
	case "Complex":
		switch rhs {
		case "Int", "Rune", "Float", "Complex":
			newConstType = "NewConstComplex128"
			resultType = "ConstComplex"
		}
	case "String":
		switch rhs {
		case "String":
			resultType = "ConstString"
		}
	}

	// If its bool op, the result is a bool. Doesn't require a constructor
	switch op {
	case token.EQL, token.NEQ, token.LEQ, token.GEQ, token.LSS, token.GTR, token.LAND, token.LOR:
		newConstType = ""
		resultType = "ConstBool"
	case token.SHR, token.SHL:
		newConstType = "NewConstInt64"
		resultType= "ConstInt"
		if lhs == rhs && lhs == "Complex" {
			// double truncation error lost
			compileErrs = append(compileErrs, compileErrs[0])
		}
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

