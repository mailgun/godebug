package main

import (
	"io"
	"text/template"
	"github.com/0xfaded/go-testgen"
)

type Test struct{}

var comment = template.Must(template.New("Comment").Parse(
`// Test {{ .Comment }}
`))

var body = template.Must(template.New("Body").Parse(
`	env := MakeSimpleEnv()
{{ if .Errors }}{{ if .TestErrs }}
	expectCheckError(t, `+"`{{ .Expr }}`"+`, env,{{ range .Errors }}
		`+"`{{ . }}`"+`,{{ end }}
	){{ else }}	_ = env{{ end }}
{{ else }}	{{ if .TestSuccess }}{{ if .ExpectConst }}expectConst(t, `+"`{{ .Expr }}`"+`, env, {{ .Expr }}, reflect.TypeOf({{ .Expr}}))
{{ else }}expectType(t, `+"`{{ .Expr }}`"+`, env, reflect.TypeOf({{ .Expr }})){{ end }}{{ else }}_ = env{{ end }}{{ end }}
`))

func (*Test) Package() string {
	return "eval"
}

func (*Test) Prefix() string {
	return "CheckBuiltin"
}

func (*Test) Imports() map[string]string {
	return map[string]string { "reflect": "" }
}

func (*Test) Dimensions() []testgen.Dimension {
	builtins := []testgen.Element{
		{"Complex", "complex"},
		{"Real", "real"},
		{"Imag", "imag"},
		{"New", "new"},
		//{"Make", "make"},
		{"Len", "len"},
		{"Cap", "cap"},
		{"Append", "append"},
		{"Copy", "copy"},
		{"Delete", "delete"},
		//{"Panic", "panic"},
	}
	arg0 := []testgen.Element{
		{"X", ""},
		{"Int", "1"},
		{"Float32", "float32(1)"},
		{"String", `"abc"`},
		{"Nil", "nil"},
		{"Float", "1.5"},
		{"Slice", "[]int{}"},
		{"Map", "map[int]int{}"},
		{"Type", "int"},
		{"MakeType", "map[int]int"},
	}
	arg1 := append(arg0, testgen.Element{"Double", "1, 1"}, testgen.Element{"Ellipsis", "[]int{1,2}..."})
	return []testgen.Dimension{
		builtins,
		arg0,
		arg1,
	}
}

func (*Test) Globals(w io.Writer) error {
	return nil
}

func (*Test) Comment(w io.Writer, elts ...testgen.Element) error {
	builtin := elts[0].Name
	sep := "("
	for _, elt := range elts[1:] {
		if elt.Value != "" {
			builtin += sep + elt.Value.(string)
			sep = ", "
		}
	}
	if sep == "(" {
		builtin += "("
	}
	builtin += ")"

	vars := map[string] interface{} {
		"Comment": builtin,
	}

	return comment.Execute(w, vars)
}

func (*Test) Body(w io.Writer, elts ...testgen.Element) error {
	expr := elts[0].Value.(string)
	sep := "("
	for _, elt := range elts[1:] {
		if elt.Value != "" {
			expr += sep + elt.Value.(string)
			sep = ", "
		}
	}
	if sep == "(" {
		expr += "("
	}
	expr += ")"

	f := elts[0].Name
	x := elts[1].Name
	y := elts[2].Name

	var compileErrs []string
	var err error
	if f == "Delete" {
		compileErrs, err = compileVoidExpr(expr)
	} else {
		compileErrs, err = compileExpr(expr)
	}
	if err != nil {
		return err
	}

	testErrs := true
	if f == "Complex" || f == "Append" || f == "Copy" || f == "Delete" {
		if len(compileErrs) == 1 && (x == "Type" || x == "MakeType") {
			if y == "Type" {
				compileErrs = append(compileErrs, "type int is not an expression")
			} else if y == "MakeType" {
				compileErrs = append(compileErrs, "type map[int]int is not an expression")
			}
		}
		if f == "Copy" {
			if x == "Nil" && y == "Nil" {
				compileErrs = append(compileErrs[:1], append([]string{compileErrs[0]}, compileErrs[1:]...)...)
			}
			if x == "Slice" && y == "Nil" {
				// http://code.google.com/p/go/issues/detail?id=7310
				testErrs = false
			}
		}
	}

	expectConst := false
	switch f {
	case "Complex", "Real", "Imag":
		expectConst = true
	case "Len", "Cap":
		z := x
		if z == "X" {
			z = y
		}
		switch z {
		case "String":
			expectConst = true
		}
	}

	vars := map[string] interface{} {
		"Expr": expr,
		"Errors": compileErrs,
		"TestErrs": testErrs,
		"TestSuccess": f != "Delete",
		"ExpectConst": expectConst,
	}

	return body.Execute(w, &vars)
}

