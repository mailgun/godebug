package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strconv"

	"golang.org/x/tools/go/ast/astutil"
	_ "golang.org/x/tools/go/gcimporter"
	"golang.org/x/tools/go/types"
)

// visitFn is a wrapper to make plain functions implement the ast.Visitor interface.
type visitFn func(ast.Node) ast.Visitor

// Visit is part of the ast.Visitor interface.
func (v visitFn) Visit(n ast.Node) ast.Visitor {
	return v(n)
}

var defs = make(map[*ast.Ident]types.Object)
var pkgName string
var fs *token.FileSet
var file *os.File
var pkg *types.Package

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Must pass a single *.go file.")
		os.Exit(1)
	}
	fs = token.NewFileSet()
	parsed, err := parser.ParseFile(fs, os.Args[1], nil, 0)
	if err != nil {
		log.Fatalf("error during parsing: %v", err)
	}
	pkgName = parsed.Name.Name
	pkg, err = (&types.Config{}).Check(parsed.Name.Name, fs, []*ast.File{parsed}, &types.Info{Defs: defs})
	if err != nil {
		log.Fatalf("error during type checking: %v", err)
	}
	file, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatal("error opening file:", err)
	}
	defer file.Close()
	ast.Walk(&visitor{context: parsed}, parsed)
	astutil.AddImport(fs, parsed, "github.com/jeremyschlatter/godebug")
	cfg := printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
	cfg.Fprint(os.Stdout, fs, parsed)
}

func newCallStmt(selector, fnName string) *ast.ExprStmt {
	return &ast.ExprStmt{
		X: newCall(selector, fnName),
	}
}

func newCall(selector, fnName string) *ast.CallExpr {
	return &ast.CallExpr{Fun: newSel(selector, fnName)}
}

func newSel(selector, name string) *ast.SelectorExpr {
	return &ast.SelectorExpr{
		X:   ast.NewIdent(selector),
		Sel: ast.NewIdent(name),
	}
}

func getText(start, end token.Pos) (text string) {
	startOffset, endOffset := fs.Position(start).Offset, fs.Position(end).Offset
	buf := make([]byte, 1+endOffset-startOffset)
	n, err := file.ReadAt(buf, int64(startOffset))
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
	call := newCall("godebug", "SLine")
	call.Args = append(call.Args, &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(text)})
	body.List = append(body.List, &ast.ExprStmt{X: call})
}

func ifElseCondWrap(cond ast.Expr, text string) ast.Expr {
	return &ast.CallExpr{
		Fun: &ast.FuncLit{
			Type: &ast.FuncType{
				Results: &ast.FieldList{
					List: []*ast.Field{{
						Type: ast.NewIdent("bool"),
					}},
				},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ExprStmt{
						X: &ast.CallExpr{
							Fun: newSel("godebug", "ElseIfExpr"),
							Args: []ast.Expr{
								&ast.BasicLit{
									Kind:  token.STRING,
									Value: strconv.Quote(text),
								},
							},
						},
					},
					&ast.ReturnStmt{
						Results: []ast.Expr{cond},
					},
				},
			},
		},
	}
}

func ifElseInitWrap(vars []ast.Expr, vals []ast.Expr, text string) ast.Expr {
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
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: newSel("godebug", "ElseIfSimpleStmt"),
						Args: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: strconv.Quote(text),
							},
						},
					},
				},
				&ast.ReturnStmt{
					Results: vals,
				},
			},
		},
	}}
}

func (v *visitor) finalizeNode() {
	switch i := v.context.(type) {
	case *ast.FuncDecl:
		if i.Body == nil || (pkgName == "main" && i.Name.Name == "main") {
			break
		}
		i.Body.List = append([]ast.Stmt{
			newCallStmt("godebug", "EnterFunc"),
			&ast.DeferStmt{
				Call: newCall("godebug", "ExitFunc"),
			},
		}, i.Body.List...)
	case *ast.BlockStmt:
		i.List = v.stmtBuf
	case *ast.IfStmt:
		if blk, ok := i.Else.(*ast.BlockStmt); ok {
			elseText := getText(i.Body.End()-1, blk.Lbrace)
			elseCall := newCall("godebug", "SLine")
			elseCall.Args = append(elseCall.Args, &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(elseText)})
			blk.List = append([]ast.Stmt{&ast.ExprStmt{X: elseCall}}, blk.List...)
		}
		if ifstmt, ok := i.Else.(*ast.IfStmt); ok {
			text := getText(i.Body.End()-1, ifstmt.Body.Pos())
			switch d := ifstmt.Init.(type) {
			case *ast.AssignStmt:
				d.Rhs = []ast.Expr{ifElseInitWrap(d.Lhs, d.Rhs, text)}
			case *ast.DeclStmt:
				spec := d.Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec)
				exprs := make([]ast.Expr, len(spec.Names))
				for i := range exprs {
					exprs[i] = ast.Expr(spec.Names[i])
				}
				spec.Values = []ast.Expr{ifElseInitWrap(exprs, spec.Values, text)}
			// TODO: optimize nil case
			default:
			}
			ifstmt.Cond = ifElseCondWrap(ifstmt.Cond, text)
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
					Names:  []*ast.Ident{ast.NewIdent("main_goScope")},
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
		return &visitor{context: node, blockVars: getIdents(i.Recv, i.Type.Params, i.Type.Results), scopeVar: "main_goScope"}
	case *ast.BlockStmt:
		w := &visitor{context: node, stmtBuf: make([]ast.Stmt, 0, 3*len(i.List)), scopeVar: v.scopeVar}
		if len(v.blockVars) > 0 {
			w.createScope()
			w.stmtBuf = append(w.stmtBuf, newDeclareCall(w.scopeVar, v.blockVars))
		}
		return w
	}
	if v.stmtBuf == nil {
		return &visitor{context: node, scopeVar: v.scopeVar}
	}
	if !isSetTraceCall(node) {
		v.stmtBuf = append(v.stmtBuf, newCallStmt("godebug", "Line"))
	}
	var newIdents []*ast.Ident
	switch i := node.(type) {
	case *ast.DeclStmt:
		newIdents = listNewIdentsFromDecl(i.Decl.(*ast.GenDecl))
	case *ast.AssignStmt:
		newIdents = listNewIdentsFromAssign(i)
	}
	if stmt, ok := node.(ast.Stmt); ok {
		v.stmtBuf = append(v.stmtBuf, stmt)
	}
	if len(newIdents) > 0 {
		if !v.createdExplicitScope {
			v.createScope()
		}
		v.stmtBuf = append(v.stmtBuf, newDeclareCall("", newIdents))
	}
	return &visitor{context: node}
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
	v.stmtBuf = append(v.stmtBuf, &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent(name)},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{newCall(v.scopeVar, "EnteringNewChildScope")},
	})
	v.stmtBuf = append(v.stmtBuf, &ast.DeferStmt{Call: newCall(name, "End")})
	v.scopeVar = name
	v.createdExplicitScope = true
}
