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
)

// visitFn is a wrapper to make plain functions implement the ast.Visitor interface.
type visitFn func(ast.Node) ast.Visitor

// Visit is part of the ast.Visitor interface.
func (v visitFn) Visit(n ast.Node) ast.Visitor {
	return v(n)
}

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
	processBlock(forstmt.Body)
}

func processRange(rangestmt *ast.RangeStmt) {
	processBlock(rangestmt.Body)
}

// appendIdentToRecordVars adds a new identifier to a godebug.RecordVars call.
// Given an *ast.Ident that corresponds to a variable "foo" and an *ast.CallExpr
// that corresponds to a call `godebug.RecordVars(&x, "x")`, this function
// modifies the call to the ast that corresponds to the new call
// `godebug.RecordVars(&x, "x", &foo, "foo")`.
func appendIdentToRecordVars(call *ast.CallExpr, ident *ast.Ident) {
	// Don't record the blank identifier.
	if ident.Name == "_" {
		return
	}
	call.Args = append(call.Args, []ast.Expr{
		&ast.UnaryExpr{
			Op: token.AND,
			X:  ident,
		},
		&ast.BasicLit{
			Kind:  token.STRING,
			Value: strconv.Quote(ident.Name),
		},
	}...)
}

func recordVarsIfNeeded(stmt ast.Stmt) *ast.ExprStmt {
	switch i := stmt.(type) {
	case *ast.DeclStmt:
		return recordVarsFromDecl(i.Decl.(*ast.GenDecl))
	case *ast.AssignStmt:
		return recordVarsFromAssign(i)
	default:
		return nil
	}
}

// recordVarsFromDecl is for declarations using the keyword "var"
func recordVarsFromDecl(decl *ast.GenDecl) *ast.ExprStmt {
	if decl.Tok != token.VAR {
		return nil
	}
	call := newGodebugCall("RecordVars")
	for _, specs := range decl.Specs {
		for _, ident := range specs.(*ast.ValueSpec).Names {
			appendIdentToRecordVars(call, ident)
		}
	}
	if len(call.Args) == 0 {
		return nil
	}
	return &ast.ExprStmt{
		X: call,
	}
}

// recordVarsFromAssign is for short variable declarations
func recordVarsFromAssign(assign *ast.AssignStmt) *ast.ExprStmt {
	// Ignore plain assignments. We're just looking for declarations.
	if assign.Tok != token.DEFINE {
		return nil
	}
	call := newGodebugCall("RecordVars")
	for _, expr := range assign.Lhs {
		if ident, ok := expr.(*ast.Ident); ok {
			appendIdentToRecordVars(call, ident)
		}
	}
	if len(call.Args) == 0 {
		return nil
	}
	return &ast.ExprStmt{
		X: call,
	}
}

func processBlock(blk *ast.BlockStmt) {
	if blk == nil {
		return
	}
	newBody := make([]ast.Stmt, 0, 2*len(blk.List))
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
		if call := recordVarsIfNeeded(stmt); call != nil {
			newBody = append(newBody, call)
		}
	}
	blk.List = newBody
}

func process(node ast.Node) ast.Visitor {
	if _, ok := node.(*ast.File); ok {
		return visitFn(process)
	}
	fn, ok := node.(*ast.FuncDecl)
	if !ok {
		return nil
	}
	processBlock(fn.Body)
	if fn.Body != nil {
		fn.Body.List = append([]ast.Stmt{
			newGodebugExpr("EnterFunc"),
			&ast.DeferStmt{
				Call: newGodebugCall("ExitFunc"),
			},
		}, fn.Body.List...)
	}
	return nil
}
