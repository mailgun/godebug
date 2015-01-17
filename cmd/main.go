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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Must pass a single *.go file.")
		os.Exit(1)
	}
	fs := token.NewFileSet()
	parsed, err := parser.ParseFile(fs, os.Args[1], nil, 0)
	if err != nil {
		log.Fatalf("error during parsing: %v", err)
	}
	_, err = (&types.Config{}).Check(parsed.Name.Name, fs, []*ast.File{parsed}, &types.Info{Defs: defs})
	if err != nil {
		log.Fatalf("error during type checking: %v", err)
	}
	astutil.AddImport(fs, parsed, "github.com/jeremyschlatter/godebug")
	ast.Walk(visitFn(process), parsed)
	cfg := printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
	cfg.Fprint(os.Stdout, fs, parsed)
}

func newGodebugExpr(fnName string) *ast.ExprStmt {
	return &ast.ExprStmt{
		X: newGodebugCall(fnName),
	}
}

func newGodebugCall(fnName string) *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent("godebug"),
			Sel: ast.NewIdent(fnName),
		},
	}
}

func processIf(ifstmt *ast.IfStmt) {
	processBlock(ifstmt.Body)
	switch i := ifstmt.Else.(type) {
	case *ast.IfStmt:
		processIf(i)
	case *ast.BlockStmt:
		processBlock(i)
	}
}

func processFor(forstmt *ast.ForStmt) {
	cleanup := processBlock(forstmt.Body)
	if cleanup != nil {
		forstmt.Body.List = append(forstmt.Body.List, &ast.ExprStmt{
			X: cleanup,
		})
	}
}

func processRange(rangestmt *ast.RangeStmt) {
	cleanup := processBlock(rangestmt.Body)
	if cleanup != nil {
		rangestmt.Body.List = append(rangestmt.Body.List, &ast.ExprStmt{
			X: cleanup,
		})
	}
}

func listNewIdents(stmt ast.Stmt) []*ast.Ident {
	switch i := stmt.(type) {
	case *ast.DeclStmt:
		return listNewIdentsFromDecl(i.Decl.(*ast.GenDecl))
	case *ast.AssignStmt:
		return listNewIdentsFromAssign(i)
	default:
		return nil
	}
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

func recordVars(idents []*ast.Ident) ast.Stmt {
	expr := newGodebugExpr("RecordVars")
	call := expr.X.(*ast.CallExpr)
	call.Args = make([]ast.Expr, 2*len(idents))
	for i, ident := range idents {
		call.Args[2*i] = &ast.UnaryExpr{
			Op: token.AND,
			X:  ident,
		}
		call.Args[2*i+1] = &ast.BasicLit{
			Kind:  token.STRING,
			Value: strconv.Quote(ident.Name),
		}
	}
	return expr
}

func outOfScopeVars(idents []*ast.Ident) *ast.CallExpr {
	call := newGodebugCall("OutOfScope")
	call.Args = make([]ast.Expr, len(idents))
	for i, ident := range idents {
		call.Args[i] = &ast.BasicLit{
			Kind:  token.STRING,
			Value: strconv.Quote(ident.Name),
		}
	}
	return call
}

func processBlock(blk *ast.BlockStmt) (cleanupCall *ast.CallExpr) {
	if blk == nil {
		return
	}
	newBody := make([]ast.Stmt, 0, 2*len(blk.List))
	var scopedIdents []*ast.Ident
	for _, stmt := range blk.List {
		newBody = append(newBody, newGodebugExpr("Line"))
		if ifstmt, ok := stmt.(*ast.IfStmt); ok {
			processIf(ifstmt)
		}
		if forstmt, ok := stmt.(*ast.ForStmt); ok {
			processFor(forstmt)
		}
		if forstmt, ok := stmt.(*ast.RangeStmt); ok {
			processRange(forstmt)
		}
		newBody = append(newBody, stmt)
		newIdents := listNewIdents(stmt)
		if len(newIdents) > 0 {
			newBody = append(newBody, recordVars(newIdents))
			scopedIdents = append(scopedIdents, newIdents...)
		}
	}
	blk.List = newBody
	if len(scopedIdents) > 0 {
		cleanupCall = outOfScopeVars(scopedIdents)
	}
	return cleanupCall
}

func process(node ast.Node) ast.Visitor {
	if _, ok := node.(*ast.File); ok {
		return visitFn(process)
	}
	fn, ok := node.(*ast.FuncDecl)
	if !ok {
		return nil
	}
	cleanupCall := processBlock(fn.Body)
	prepend := []ast.Stmt{
		newGodebugExpr("EnterFunc"),
		&ast.DeferStmt{
			Call: newGodebugCall("ExitFunc"),
		},
	}
	if cleanupCall != nil {
		prepend = append(prepend, &ast.DeferStmt{
			Call: cleanupCall,
		})
	}
	if fn.Body != nil {
		fn.Body.List = append(prepend, fn.Body.List...)
	}
	return nil
}
