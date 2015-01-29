package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"

	"github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/ast/astutil"
	_ "github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/gcimporter"
	"github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/loader"
	"github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/types"
	_ "github.com/mailgun/godebug/lib" // so the library is also installed whenever this package is
)

var w = flag.Bool("w", false, "write result to (source) file instead of stdout")

var (
	defs map[*ast.Ident]types.Object
	fs   *token.FileSet
	pkg  *types.Package
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "godebug <flags> <args>\n\n")
	fmt.Fprintf(os.Stderr, "flags:\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, loader.FromArgsUsage)
	os.Exit(2)
}

func main() {
	flag.Parse()
	flag.Usage = Usage
	var conf loader.Config
	rest, err := conf.FromArgs(flag.Args(), true)
	if len(rest) > 0 {
		fmt.Fprintf(os.Stderr, "Unrecognized arguments:\n%v\n\n", strings.Join(rest, "\n"))
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error identifying packages: %v\n\n", err)
	}
	if len(rest) > 0 || err != nil {
		flag.Usage()
	}
	conf.SourceImports = true
	prog, err := conf.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading packages: %v\n\n", err)
		flag.Usage()
	}
	fs = prog.Fset
	for _, pkgInfo := range prog.InitialPackages() {
		defs = pkgInfo.Defs
		pkg = pkgInfo.Pkg
		for _, f := range pkgInfo.Files {
			ast.Walk(&visitor{context: f, scopeVar: fileScope(f.Pos())}, f)
			astutil.AddImport(fs, f, "github.com/mailgun/godebug/lib")
			cfg := printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
			out := os.Stdout
			if *w {
				file, err := os.Create(fs.Position(f.Pos()).Filename)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				defer file.Close()
				out = file
			}
			cfg.Fprint(out, fs, f)
		}
	}
}

func newCallStmt(selector, fnName string, args ...ast.Expr) *ast.ExprStmt {
	return &ast.ExprStmt{
		X: newCall(selector, fnName, args...),
	}
}

func newCall(selector, fnName string, args ...ast.Expr) *ast.CallExpr {
	return &ast.CallExpr{Fun: newSel(selector, fnName), Args: args}
}

func newSel(selector, name string) *ast.SelectorExpr {
	return &ast.SelectorExpr{
		X:   ast.NewIdent(selector),
		Sel: ast.NewIdent(name),
	}
}

func fileScope(pos token.Pos) string {
	return strings.Map(func(r rune) rune {
		if !unicode.In(r, unicode.Digit, unicode.Letter) {
			return '_'
		}
		return r
	}, path.Base(fs.Position(pos).Filename)) + "Scope"
}

func getText(start, end token.Pos) (text string) {
	f, err := os.Open(fs.Position(start).Filename)
	if err != nil {
		return "<< Error reading source >>"
	}
	defer f.Close()
	startOffset, endOffset := fs.Position(start).Offset, fs.Position(end).Offset
	buf := make([]byte, 1+endOffset-startOffset)
	n, err := f.ReadAt(buf, int64(startOffset))
	text = string(buf[:n])
	if err != nil {
		text += "<< Error reading source >>"
	}
	return
}

func isNewIdent(ident *ast.Ident) bool {
	return ident.Name != "_" && defs[ident] != nil
}

// listNewIdentsFromDecl is for declarations using the keyword "var"
func listNewIdentsFromDecl(decl *ast.GenDecl) (idents []*ast.Ident) {
	if decl.Tok != token.VAR {
		return
	}
	for _, specs := range decl.Specs {
		for _, ident := range specs.(*ast.ValueSpec).Names {
			if isNewIdent(ident) {
				idents = append(idents, ident)
			}
		}
	}
	return
}

// listNewIdentsFromAssign is for short variable declarations
func listNewIdentsFromAssign(assign *ast.AssignStmt) (idents []*ast.Ident) {
	for _, expr := range assign.Lhs {
		if ident, ok := expr.(*ast.Ident); ok && isNewIdent(ident) {
			idents = append(idents, ident)
		}
	}
	return
}

func isSetTraceCall(node ast.Node) (b bool) {
	defer func() {
		if r := recover(); r != nil {
			b = false
		}
	}()
	sel := node.(*ast.ExprStmt).X.(*ast.CallExpr).Fun.(*ast.SelectorExpr)
	return sel.X.(*ast.Ident).Name == "godebug" && sel.Sel.Name == "SetTrace"
}

type visitor struct {
	context              ast.Node
	stmtBuf              []ast.Stmt
	scopeVar             string
	blockVars            []*ast.Ident
	createdExplicitScope bool
}

func (v *visitor) finalizeLoop(pos token.Pos, body *ast.BlockStmt) {
	if body == nil {
		return
	}
	text := getText(pos, body.Lbrace)
	call := newCall("godebug", "SLine", ast.NewIdent("ctx"), ast.NewIdent(v.scopeVar), &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(text)})
	body.List = append(body.List, &ast.ExprStmt{X: call})
}

func (v *visitor) ifElseCondWrap(cond ast.Expr, text string) ast.Expr {
	return &ast.CallExpr{
		Fun: &ast.FuncLit{
			Type: &ast.FuncType{
				Results: &ast.FieldList{
					List: []*ast.Field{{
						Type: ast.NewIdent("bool")}}}},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					newCallStmt("godebug", "ElseIfExpr", ast.NewIdent("ctx"), ast.NewIdent(v.scopeVar), &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(text)}),
					&ast.ReturnStmt{Results: []ast.Expr{cond}}}}}}
}

func (v *visitor) ifElseInitWrap(vars []ast.Expr, vals []ast.Expr, text string) ast.Expr {
	results := &ast.FieldList{List: make([]*ast.Field, len(vars))}
	for i, expr := range vars {
		ident, ok := expr.(*ast.Ident)
		if !ok {
			panic(fmt.Sprintf("Unsupported type in if statement initializer: %T. Sorry! Let me (jeremy@mailgunhq.com) know about this and I'll fix it.", expr))
		}
		results.List[i] = &ast.Field{Type: ast.NewIdent(types.TypeString(pkg, defs[ident].Type()))}
	}
	return &ast.CallExpr{Fun: &ast.FuncLit{
		Type: &ast.FuncType{
			Results: results,
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				newCallStmt("godebug", "ElseIfSimpleStmt", ast.NewIdent("ctx"), ast.NewIdent(v.scopeVar), &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(text)}),
				&ast.ReturnStmt{Results: vals}}}}}
}

var blank = ast.NewIdent("_")

func inputsOrOutputs(fieldList *ast.FieldList, prefix string) (decl []ast.Stmt, all []ast.Expr) {
	if fieldList == nil {
		return
	}
	count := 1
	var specs []ast.Spec
	for _, field := range fieldList.List {
		names := field.Names
		if names == nil {
			names = []*ast.Ident{blank}
		}
		spec := &ast.ValueSpec{Type: field.Type}
		for _, name := range names {
			if name.Name == "_" {
				name = ast.NewIdent(prefix + strconv.Itoa(count))
				spec.Names = append(spec.Names, name)
			}
			count++
			all = append(all, name)
		}
		if len(spec.Names) > 0 {
			specs = append(specs, spec)
		}
	}
	if len(specs) > 0 {
		decl = []ast.Stmt{&ast.DeclStmt{Decl: &ast.GenDecl{Tok: token.VAR, Specs: specs}}}
	}
	return decl, all
}

func genEnterFunc(fn *ast.FuncDecl, inputs, outputs []ast.Expr) []ast.Stmt {
	// Generates this structure:
	//   ctx, ok := godebug.EnterFunc(func) {
	//       <outputs> = <fn.Name>(inputs)
	//   })
	//   if !ok {
	//       return <outputs>
	//   }
	//
	//   if !godebug.EnterFunc(func() {
	//       <outputs> = <fn.Name>(inputs)
	//   }) {
	//       return <outputs>
	//   }
	var innerCall ast.Stmt = &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun:  fn.Name,
			Args: inputs,
		},
	}
	if len(outputs) > 0 {
		innerCall = &ast.AssignStmt{
			Lhs: outputs,
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{innerCall.(*ast.ExprStmt).X},
		}
	}
	return []ast.Stmt{
		&ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent("ctx"), ast.NewIdent("ok")},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				newCall("godebug", "EnterFunc", &ast.FuncLit{
					Type: &ast.FuncType{Params: &ast.FieldList{}},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							innerCall}}})}},
		&ast.IfStmt{
			Cond: &ast.UnaryExpr{
				Op: token.NOT,
				X:  ast.NewIdent("ok"),
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ReturnStmt{
						Results: outputs}}}}}
}

func (v *visitor) finalizeNode() {
	switch i := v.context.(type) {
	case *ast.FuncDecl:
		if i.Body == nil {
			break
		}
		declOuts, outputs := inputsOrOutputs(i.Type.Results, "godebugResult")
		declIns, inputs := inputsOrOutputs(i.Type.Params, "godebugInput")
		prepend := append(declIns, declOuts...)
		prepend = append(prepend, genEnterFunc(i, inputs, outputs)...)
		if !(pkg.Name() == "main" && i.Name.Name == "main") {
			prepend = append(prepend, &ast.DeferStmt{
				Call: newCall("godebug", "ExitFunc"),
			})
		}

		i.Body.List = append(prepend, i.Body.List...)
	case *ast.FuncLit:
		// Transform the function literal into this form:
		//   func(<inputs>) <outputs> {
		//       <var statement for unnamed outputs>
		//       var ctx *godebug.Context
		//       fn := func() {
		//           <outputs> = func() <outputs> {
		//               ... function body here
		//           }()
		//       }
		//       var ok bool
		//       ctx, ok = godebug.EnterFunc(fn)
		//       if ok {
		//           fn()
		//       }
		//       godebug.ExitFunc()
		//       return <outputs>
		//   }
		//
		//   OR
		//
		//   func(<inputs>) {
		//       var ctx *godebug.Context
		//       fn := func() {
		//           ... function body here
		//       }
		//       var ok bool
		//       ctx, ok = godebug.EnterFunc(fn)
		//       if ok {
		//           fn()
		//       }
		//       godebug.ExitFunc()
		//   }
		decl, outputs := inputsOrOutputs(i.Type.Results, "godebugResult")
		wrappedFuncLit := &ast.FuncLit{
			Type: &ast.FuncType{
				Params:  &ast.FieldList{},
				Results: i.Type.Results,
			},
			Body: i.Body,
		}
		if len(outputs) > 0 {
			wrappedFuncLit = &ast.FuncLit{
				Type: &ast.FuncType{Params: &ast.FieldList{}},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Lhs: outputs,
							Tok: token.ASSIGN,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: wrappedFuncLit}}}}}}
		}
		newBody := &ast.BlockStmt{}
		newBody.List = append(decl,
			&ast.DeclStmt{
				Decl: &ast.GenDecl{
					Tok: token.VAR,
					Specs: []ast.Spec{
						&ast.ValueSpec{
							Names: []*ast.Ident{ast.NewIdent("ctx")},
							Type:  &ast.StarExpr{X: newSel("godebug", "Context")}}}}},
			&ast.AssignStmt{
				Lhs: []ast.Expr{ast.NewIdent("fn")},
				Tok: token.DEFINE,
				Rhs: []ast.Expr{wrappedFuncLit}},
			&ast.DeclStmt{
				Decl: &ast.GenDecl{
					Tok: token.VAR,
					Specs: []ast.Spec{
						&ast.ValueSpec{
							Names: []*ast.Ident{ast.NewIdent("ok")},
							Type:  ast.NewIdent("bool")}}}},
			&ast.AssignStmt{
				Lhs: []ast.Expr{ast.NewIdent("ctx"), ast.NewIdent("ok")},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{newCall("godebug", "EnterFunc", ast.NewIdent("fn"))}},
			&ast.IfStmt{
				Cond: ast.NewIdent("ok"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: ast.NewIdent("fn")}}}}},
			newCallStmt("godebug", "ExitFunc"),
		)
		if len(outputs) > 0 {
			newBody.List = append(newBody.List, &ast.ReturnStmt{Results: outputs})
		}
		i.Body = newBody
	case *ast.BlockStmt:
		i.List = v.stmtBuf
	case *ast.IfStmt:
		if blk, ok := i.Else.(*ast.BlockStmt); ok {
			elseText := getText(i.Body.End()-1, blk.Lbrace)
			elseCall := newCall("godebug", "SLine", ast.NewIdent("ctx"), ast.NewIdent(v.scopeVar), &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(elseText)})
			blk.List = append([]ast.Stmt{&ast.ExprStmt{X: elseCall}}, blk.List...)
		}
		if ifstmt, ok := i.Else.(*ast.IfStmt); ok {
			text := getText(i.Body.End()-1, ifstmt.Body.Pos())
			switch d := ifstmt.Init.(type) {
			case *ast.AssignStmt:
				d.Rhs = []ast.Expr{v.ifElseInitWrap(d.Lhs, d.Rhs, text)}
			case *ast.DeclStmt:
				spec := d.Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec)
				exprs := make([]ast.Expr, len(spec.Names))
				for i := range exprs {
					exprs[i] = ast.Expr(spec.Names[i])
				}
				spec.Values = []ast.Expr{v.ifElseInitWrap(exprs, spec.Values, text)}
			// TODO: optimize nil case
			default:
			}
			ifstmt.Cond = v.ifElseCondWrap(ifstmt.Cond, text)
		}
	case *ast.RangeStmt:
		v.finalizeLoop(i.For, i.Body)
	case *ast.ForStmt:
		v.finalizeLoop(i.For, i.Body)
	case *ast.File:
		// Insert declaration of file-level godebug.Scope variable as first declaration in file.
		var newDecls []ast.Decl
		// But put it after any import declarations.
		for len(i.Decls) > 0 {
			if gd, ok := i.Decls[0].(*ast.GenDecl); !ok || gd.Tok != token.IMPORT {
				break
			}
			newDecls = append(newDecls, i.Decls[0])
			i.Decls = i.Decls[1:]
		}
		newDecls = append(newDecls, &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names:  []*ast.Ident{ast.NewIdent(fileScope(i.Pos()))},
					Values: []ast.Expr{newCall("godebug", "EnteringNewScope")},
				},
			},
		})
		i.Decls = append(newDecls, i.Decls...)
	}
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	switch i := node.(type) {
	case nil:
		v.finalizeNode()
		return nil
	case *ast.FuncDecl:
		// Add Declare() call first thing in the function for any variables bound by the function signature.
		return &visitor{context: node, blockVars: getIdents(i.Recv, i.Type.Params, i.Type.Results), scopeVar: fileScope(i.Pos())}
	case *ast.FuncLit:
		// Add Declare() call first thing in the function for any variables bound by the function signature.
		return &visitor{context: node, blockVars: getIdents(i.Type.Params, i.Type.Results), scopeVar: v.scopeVar}
	case *ast.BlockStmt:
		w := &visitor{context: node, stmtBuf: make([]ast.Stmt, 0, 3*len(i.List)), scopeVar: v.scopeVar}
		if len(v.blockVars) > 0 {
			w.createScope()
			w.stmtBuf = append(w.stmtBuf, newDeclareCall(w.scopeVar, v.blockVars))
		}
		return w
	// TODO: Wrap these clauses the same way as if-else clauses.
	case *ast.CommClause, *ast.CaseClause:
		v.stmtBuf = append(v.stmtBuf, i.(ast.Stmt))
		return &visitor{context: node, scopeVar: v.scopeVar}
	}
	if v.stmtBuf == nil {
		return &visitor{context: node, scopeVar: v.scopeVar}
	}
	if !isSetTraceCall(node) {
		v.stmtBuf = append(v.stmtBuf, newCallStmt("godebug", "Line", ast.NewIdent("ctx"), ast.NewIdent(v.scopeVar)))
	}
	var newIdents []*ast.Ident
	switch i := node.(type) {
	case *ast.DeclStmt:
		newIdents = listNewIdentsFromDecl(i.Decl.(*ast.GenDecl))
	case *ast.AssignStmt:
		newIdents = listNewIdentsFromAssign(i)
	}
	if stmt, ok := node.(ast.Stmt); ok {
		if isSetTraceCall(node) {
			// Rewrite godebug.SetTrace() as godebug.SetTraceGen(ctx)
			call := stmt.(*ast.ExprStmt).X.(*ast.CallExpr)
			call.Args = []ast.Expr{ast.NewIdent("ctx")}
			call.Fun.(*ast.SelectorExpr).Sel.Name = "SetTraceGen"
		}
		v.stmtBuf = append(v.stmtBuf, stmt)
	}
	if len(newIdents) > 0 {
		if !v.createdExplicitScope {
			v.createScope()
		}
		v.stmtBuf = append(v.stmtBuf, newDeclareCall("", newIdents))
	}
	return &visitor{context: node, scopeVar: v.scopeVar}
}

func getIdents(lists ...*ast.FieldList) (idents []*ast.Ident) {
	for _, l := range lists {
		if l == nil {
			continue
		}
		for _, fields := range l.List {
			for _, ident := range fields.Names {
				if ident.Name != "_" {
					idents = append(idents, ident)
				}
			}
		}
	}
	return
}

func newDeclareCall(scopeVar string, idents []*ast.Ident) ast.Stmt {
	if scopeVar == "" {
		scopeVar = "godebugScope"
	}
	expr := newCallStmt(scopeVar, "Declare")
	call := expr.X.(*ast.CallExpr)
	call.Args = make([]ast.Expr, 2*len(idents))
	for i, ident := range idents {
		call.Args[2*i] = &ast.BasicLit{
			Kind:  token.STRING,
			Value: strconv.Quote(ident.Name),
		}
		call.Args[2*i+1] = &ast.UnaryExpr{
			Op: token.AND,
			X:  ident,
		}
	}
	return expr
}

func (v *visitor) createScope() {
	name := "godebugScope"
	if v.scopeVar == "" {
		v.scopeVar = name
	}
	v.stmtBuf = append(v.stmtBuf, &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent(name)},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{newCall(v.scopeVar, "EnteringNewChildScope")},
	})
	v.scopeVar = name
	v.createdExplicitScope = true
}
