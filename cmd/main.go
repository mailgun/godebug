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
	ast.Walk(visitFn(addImport), parsed)
	/*
		//ast.Inspect(parsed, inspect)
		//cfg := printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
		//cfg.Fprint(os.Stdout, fs, parsed)
		for fn := range fnsCalled {
			fmt.Println(fn.Name)
		}
	*/
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

func addImport(node ast.Node) ast.Visitor {
	if _, ok := node.(*ast.File); ok {
		return visitFn(addImport)
	}
	genDecl, ok := node.(*ast.GenDecl)
	if !ok || genDecl.Tok != token.IMPORT {
		return nil
	}
	genDecl.Specs = append(genDecl.Specs, &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: `"github.com/jeremyschlatter/godebug"`,
		},
	})
	return nil
}

/*
var fnsCalled map[*ast.Object]bool
func findMain(node ast.Node) ast.Visitor {
	fn, ok := node.(*ast.FuncDecl)
	if ok && fn.Name.Name == "main" {
		return visitFn(traceMain)
	}
	return visitFn(findMain)
}

func traceMain(node ast.Node) ast.Visitor {
	callExpr, ok := node.(*ast.CallExpr)
	if !ok {
		return visitFn(traceMain)
	}
	var fn *ast.Ident
	switch i := callExpr.Fun.(type) {
	case *ast.Ident:
		fn = i
	case *ast.SelectorExpr:
		fmt.Printf("%#v\n", i.X.(*ast.Ident).Obj)
		fmt.Printf("%#v\n", i.Sel)
		fn = i.Sel
	default:
		return visitFn(traceMain)
	}
	if fn.Obj == nil || fn.Obj.Name == "Hello" {
		fmt.Println("hello")
		return visitFn(traceMain)
	}
	if fnsCalled[fn.Obj] {
		return visitFn(traceMain)
	}
	fnsCalled[fn.Obj] = true
	ast.Walk(visitFn(traceMain), fn.Obj.Decl.(*ast.FuncDecl))
	return visitFn(traceMain)
}
*/

/*
func inspect(node ast.Node) bool {
	fn, ok := node.(*ast.FuncDecl)
	if !ok {
		return true
	}
	switch fn.Name.Name {
	case "foo":
	default:
		return true
	}
	ctx := ast.Field{
		Names: []*ast.Ident{ast.NewIdent("ctx")},
		Type:  ast.NewIdent("Context"),
	}
	fn.Type.Params.List = append([]*ast.Field{&ctx}, fn.Type.Params.List...)
	return true
}
*/
