package main

import (
	"io"
	"text/template"
	"github.com/0xfaded/go-testgen"
)

type Test struct{}

var comment = template.Must(template.New("Comment").Parse(
`// Test {{ .Stmt.Name }}
`))

var body = template.Must(template.New("Body").Parse(
`	env := MakeSimpleEnv()
	f := func() (int, int) { return 1, 1 }
	env.Vars["f"] = reflect.ValueOf(&f)
	expectCheckError(t, `+"`{{ .Stmt.Value }}`"+`, env,{{ range .Errors }}
		`+"`{{ . }}`"+`,{{ end }}
	)
`))

func (*Test) Package() string {
	return "eval"
}

func (*Test) Prefix() string {
	return "CheckAssignStmt"
}

func (*Test) Imports() map[string]string {
	return map[string]string{"reflect": ""}
}

func (*Test) Dimensions() []testgen.Dimension {
	stmts := []testgen.Element{
		{"TooMany", "_, _ = 1, 2, 3"},
		{"TooFew", "_, _ = 1"},
		{"NoNewIdents1", "f := nil"},
		{"NoNewIdents2", "f, _ := nil, 1"},
		{"UnderscoreNil", "_ = nil"},
		// this case is actually detected by the parser
		//{"NonName", "a.b, _ := 1, 1"},
		{"Unaddressable", "1 = 1"},
		{"Unaddressable2", "1, 2 = 1, 2"},
		{"ToNil", "nil = 1"},
		{"Mistyped1", "f = true"},
		{"Mistyped2", "f, f = true, false"},
	}

	return []testgen.Dimension{stmts}
}

func (*Test) Globals(w io.Writer) error {
	return nil
}

func (*Test) Comment(w io.Writer, elts ...testgen.Element) error {
	vars := map[string] interface{} {
		"Stmt": elts[0],
	}

	return comment.Execute(w, vars)
}

func (*Test) Body(w io.Writer, elts ...testgen.Element) error {
	stmt := elts[0]
	compileErrs, err := compileVoidExprWithDefs(stmt.Value.(string), "f := func() (int, int) { return 1, 1 }")
	if err != nil {
		return err
	}

	vars := map[string] interface{} {
		"Stmt": stmt,
		"Errors": compileErrs,
	}

	return body.Execute(w, &vars)
}

