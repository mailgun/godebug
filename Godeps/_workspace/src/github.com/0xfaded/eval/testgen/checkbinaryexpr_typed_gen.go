package main

import (
	"fmt"
	"io"
        "strings"
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
	expectConst(t, `+"`{{ .Expr }}`"+`, env, {{ .Expr }}, reflect.TypeOf({{ .Expr }})){{ end }}
`))

func (*Test) Package() string {
	return "eval"
}

func (*Test) Prefix() string {
	return "CheckBinaryTypedExpr"
}

func (*Test) Imports() map[string]string {
	return map[string]string { "reflect": "" }
}

func (*Test) Dimensions() []testgen.Dimension {
	lhs := []testgen.Element{
		{"Int8", "int8(0x7f)"},
		{"Int16", "int16(0x7fff)"},
		{"Int32", "int32(0x7fffffff)"},
		{"Int64", "int64(0x7fffffffffffffff)"},
		{"Uint8", "uint8(0xff)"},
		{"Uint16", "uint16(0xffff)"},
		{"Uint32", "uint32(0xffffffff)"},
		{"Uint64", "uint64(0xffffffffffffffff)"},
		{"Float32", "float32(0xffffffff)"},
		{"Float64", "float64(0xffffffff)"},
		{"Complex64", "complex64(0xffffffff + 0xffffffff * 1i)"},
		{"Complex128", "complex128(0xffffffff + 0xffffffff * 1i)"},
		{"Rune32", "rune(0x7fffffff)"},
		{"StringT", `string("abc")`},
		{"BoolT", "bool(true)"},
	}
	ops := []testgen.Element{
		{"Add", token.ADD},
		{"Sub", token.SUB},
		{"And", token.AND},
		{"Rem", token.REM},
		{"Eql", token.EQL},
		{"Gtr", token.GTR},
		{"Shl", token.SHL},
	}
	rhs := []testgen.Element{
		{"Int", "4"},
		{"Rune", "'@'"},
		{"Float", "2.0"},
		{"Complex", "8.0i"},
		{"Bool", "true"},
		{"String", `"abc"`},
		{"Nil", "nil"},
		{"Int8", "int8(0x7f)"},
		{"Int16", "int16(0x7fff)"},
		{"Int32", "int32(0x7fffffff)"},
		{"Int64", "int64(0x7fffffffffffffff)"},
		{"Uint8", "uint8(0xff)"},
		{"Uint16", "uint16(0xffff)"},
		{"Uint32", "uint32(0xffffffff)"},
		{"Uint64", "uint64(0xffffffffffffffff)"},
		{"Float32", "float32(0xffffffff)"},
		{"Float64", "float64(0xffffffff)"},
		{"Complex64", "complex64(0xffffffff + 0xffffffff * 1i)"},
		{"Complex128", "complex128(0xffffffff + 0xffffffff * 1i)"},
		{"Rune32", "rune(0x7fffffff)"},
		{"StringT", `string("abc")`},
		{"BoolT", "bool(true)"},
	}
	return []testgen.Dimension{
		lhs,
		ops,
		rhs,
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
	op  := elts[1].Value.(token.Token)

	expr := fmt.Sprintf("%v %v %v", elts[0].Value, op, elts[2].Value)
	compileErrs, err := compileExpr(expr)
	if err != nil {
		return err
	}

        // TODO Fix for bad complex number formatting in gc. This has been
        // addressed in gc 1.3. Remove this after release.
        buggy := "4.29497e+094.29497e+09i"
        fix := "4.29497e+09+4.29497e+09i"
        for i := range compileErrs {
                compileErrs[i] = strings.Replace(compileErrs[i], buggy, fix, -1)
        }

	vars := map[string] interface{} {
		"Expr": expr,
		"Errors": compileErrs,
		"Op": elts[1],
	}

	return body.Execute(w, &vars)
}

