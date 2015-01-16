package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
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
	ast.Walk(visitFn(addLineFuncs), parsed)
	cfg := printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
	cfg.Fprint(os.Stdout, fs, parsed)
}

func newGodebugCall(fnName string) *ast.ExprStmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("godebug"),
				Sel: ast.NewIdent(fnName),
			},
		},
	}
}

func addLineFuncsToBlock(blk *ast.BlockStmt) {
	if blk == nil {
		return
	}
	newBody := make([]ast.Stmt, 0, 2*len(blk.List))
	for _, stmt := range blk.List {
		newBody = append(newBody, newGodebugCall("Line"))
		if ifstmt, ok := stmt.(*ast.IfStmt); ok {
			addLineFuncsToBlock(ifstmt.Body)
		}
		newBody = append(newBody, stmt)
	}
	blk.List = newBody
}

func addLineFuncs(node ast.Node) ast.Visitor {
	if _, ok := node.(*ast.File); ok {
		return visitFn(addLineFuncs)
	}
	fn, ok := node.(*ast.FuncDecl)
	if !ok {
		return nil
	}
	/*
		for i := range fn.Body.List {
			fmt.Printf("%T\n", fn.Body.List[i].(*ast.ExprStmt).X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).X)
		}
	*/
	addLineFuncsToBlock(fn.Body)
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
