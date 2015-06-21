package gen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("%s|{{[^{}]*}}")

func init() {
	re.Longest()
}

func isEmpty(i interface{}) bool {
	switch x := i.(type) {
	case string:
		return x == ""
	case []ast.Expr:
		return len(x) == 0
	case []ast.Stmt:
		return len(x) == 0
	case nil:
		return true
	case ast.Node:
		return false
	}
	panic(fmt.Sprintf("unrecognized type: %T", i))
}

func parseOptionalSections(format string, a ...interface{}) (string, []interface{}) {
	out := make([]interface{}, 0, len(a))
	var next interface{}
	b := re.ReplaceAllFunc([]byte(format), func(match []byte) []byte {
		next, a = a[0], a[1:]
		if bytes.Equal(match, []byte("%s")) {
			out = append(out, next)
			return match
		}
		if isEmpty(next) {
			return nil
		}
		out = append(out, next)
		return match[2 : len(match)-2]
	})
	return string(b), out
}

func replaceIdents(format string) string {
	r := strings.NewReplacer(
		"ctx", idents.ctx,
		"ok", idents.ok,
		"scope", idents.scope,
		"receiver", idents.receiver,
		"recoverChan", idents.recoverChan,
		"_r", idents.recoverChanChan,
		"recovers", idents.recovers,
		"panicVal", idents.panicVal,
		"panicChan", idents.panicChan,
		"godebug", idents.godebug)
	return r.Replace(format)
}

// not thread safe.
func astPrintf(format string, a ...interface{}) []ast.Stmt {
	format = replaceIdents(format)
	format, a = parseOptionalSections(format, a...)
	var buf bytes.Buffer
	r := io.TeeReader(makeReader(format, a), &buf)
	f, err := parser.ParseFile(token.NewFileSet(), "", r, 0)
	if err != nil {
		_, _ = io.Copy(os.Stdout, &buf)
		panic(err)
	}
	body := f.Decls[0].(*ast.FuncDecl).Body
	ast.Walk(posEraser{}, body)
	return body.List
}

var (
	progBeginning = `package main
func main() {`
	progEnd = "}"
)

func fprintFieldList(w io.Writer, lst *ast.FieldList) error {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, token.NewFileSet(), &ast.FuncLit{Type: &ast.FuncType{Params: lst}}); err != nil {
		return err
	}
	b := buf.Bytes()
	_, err := w.Write(b[len("func(") : len(b)-len(")")])
	return err
}

func fprint(w io.Writer, a interface{}) error {
	switch x := a.(type) {
	case *ast.FieldList:
		// printer.Fprint does not support this type. Hack around it.
		return fprintFieldList(w, x)
	case ast.Node, []ast.Decl, []ast.Stmt:
		return printer.Fprint(w, token.NewFileSet(), x)
	case []ast.Expr:
		i := 0
		for ; i < len(x)-1; i++ {
			if err := printer.Fprint(w, token.NewFileSet(), x[i]); err != nil {
				return err
			}
			if _, err := w.Write([]byte(", ")); err != nil {
				return err
			}
		}
		if len(x) != 0 {
			return printer.Fprint(w, token.NewFileSet(), x[i])
		}
		return nil
	case string:
		_, err := io.WriteString(w, x)
		return err
	default:
		panic(fmt.Sprintf("unsupported value: %v", x))
	}
}

func makeReader(format string, as []interface{}) io.Reader {
	parts := bytes.Split([]byte(format), []byte("%s"))
	if len(parts) != len(as) && len(parts) != len(as)+1 {
		panic("Wrong number of arguments")
	}
	pipeReaders := make([]*io.PipeReader, len(as))
	pipeWriters := make([]*io.PipeWriter, len(as))
	for i, a := range as {
		pipeReaders[i], pipeWriters[i] = io.Pipe()
		go func(i int, a interface{}) {
			if err := fprint(pipeWriters[i], a); err != nil {
				panic(err)
			}
			pipeWriters[i].Close()
		}(i, a)
		//go printer.Fprint(pipeWriters[i], token.NewFileSet(), node)
	}
	readers := make([]io.Reader, 2+len(parts)+len(pipeReaders))
	readers[0] = strings.NewReader(progBeginning)
	for i := 0; i < len(parts)+len(pipeReaders); i++ {
		if i%2 == 0 {
			readers[i+1] = bytes.NewReader(parts[i/2])
		} else {
			readers[i+1] = pipeReaders[(i-1)/2]
		}
	}
	readers[len(readers)-1] = strings.NewReader(progEnd)
	return io.MultiReader(readers...)
}

type posEraser struct{}

func (eraser posEraser) Visit(node ast.Node) ast.Visitor {
	switch i := node.(type) {
	// Special case for ellipsis: replace nonzero token.Pos values with token.Pos(1).
	case *ast.CallExpr:
		i.Lparen = token.NoPos
		if i.Ellipsis.IsValid() {
			i.Ellipsis = token.Pos(1)
		}
		i.Rparen = token.NoPos
	case *ast.Ellipsis:
		if i.Ellipsis.IsValid() {
			i.Ellipsis = token.Pos(1)
		}
	// Special case for *ast.GenDecl: the printer assumes that this is a single declaration
	// if Lparen == token.NoPos. So use token.Pos(1) instead.
	case *ast.GenDecl:
		i.TokPos = token.NoPos
		if i.Lparen.IsValid() {
			i.Lparen = token.Pos(1)
		}
		i.Rparen = token.NoPos

	// Zero out all other token.Pos values.
	case *ast.ArrayType:
		i.Lbrack = token.NoPos
	case *ast.AssignStmt:
		i.TokPos = token.NoPos
	case *ast.BasicLit:
		i.ValuePos = token.NoPos
	case *ast.BinaryExpr:
		i.OpPos = token.NoPos
	case *ast.BlockStmt:
		i.Lbrace = token.NoPos
		i.Rbrace = token.NoPos
	case *ast.BranchStmt:
		i.TokPos = token.NoPos
	case *ast.CaseClause:
		i.Case = token.NoPos
		i.Colon = token.NoPos
	case *ast.ChanType:
		i.Begin = token.NoPos
		i.Arrow = token.NoPos
	case *ast.CommClause:
		i.Case = token.NoPos
		i.Colon = token.NoPos
	case *ast.Comment:
		i.Slash = token.NoPos
	case *ast.CompositeLit:
		i.Lbrace = token.NoPos
		i.Rbrace = token.NoPos
	case *ast.DeferStmt:
		i.Defer = token.NoPos
	case *ast.EmptyStmt:
		i.Semicolon = token.NoPos
	case *ast.FieldList:
		i.Opening = token.NoPos
		i.Closing = token.NoPos
	case *ast.File:
		i.Package = token.NoPos
	case *ast.ForStmt:
		i.For = token.NoPos
	case *ast.FuncType:
		i.Func = token.NoPos
	case *ast.GoStmt:
		i.Go = token.NoPos
	case *ast.Ident:
		i.NamePos = token.NoPos
	case *ast.IfStmt:
		i.If = token.NoPos
	case *ast.ImportSpec:
		i.EndPos = token.NoPos
	case *ast.IncDecStmt:
		i.TokPos = token.NoPos
	case *ast.IndexExpr:
		i.Lbrack = token.NoPos
		i.Rbrack = token.NoPos
	case *ast.InterfaceType:
		i.Interface = token.NoPos
	case *ast.KeyValueExpr:
		i.Colon = token.NoPos
	case *ast.LabeledStmt:
		i.Colon = token.NoPos
	case *ast.MapType:
		i.Map = token.NoPos
	case *ast.ParenExpr:
		i.Lparen = token.NoPos
		i.Rparen = token.NoPos
	case *ast.RangeStmt:
		i.For = token.NoPos
		i.TokPos = token.NoPos
	case *ast.ReturnStmt:
		i.Return = token.NoPos
	case *ast.SelectStmt:
		i.Select = token.NoPos
	case *ast.SendStmt:
		i.Arrow = token.NoPos
	case *ast.SliceExpr:
		i.Lbrack = token.NoPos
		i.Rbrack = token.NoPos
	case *ast.StarExpr:
		i.Star = token.NoPos
	case *ast.StructType:
		i.Struct = token.NoPos
	case *ast.SwitchStmt:
		i.Switch = token.NoPos
	case *ast.TypeAssertExpr:
		i.Lparen = token.NoPos
		i.Rparen = token.NoPos
	case *ast.TypeSwitchStmt:
		i.Switch = token.NoPos
	case *ast.UnaryExpr:
		i.OpPos = token.NoPos
	}
	return eraser
}
