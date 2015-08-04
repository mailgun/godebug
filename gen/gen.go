package gen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/ast/astutil"
	"github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/exact"
	_ "github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/gccgoimporter"
	_ "github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/gcimporter"
	"github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/loader"
	"github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/types"
	_ "github.com/mailgun/godebug/lib" // so the library is also installed whenever this package is
)

var (
	defs   map[*ast.Ident]types.Object
	_types map[ast.Expr]types.TypeAndValue
	fs     *token.FileSet
	pkg    *types.Package
)

type Config struct {
	loader.Config
}

func Generate(prog *loader.Program, getFileBytes func(string) ([]byte, error), writerFor func(importPath, filename string) io.WriteCloser) {
	for _, pkgInfo := range prog.InitialPackages() {
		defs = pkgInfo.Defs
		_types = pkgInfo.Types
		pkg = pkgInfo.Pkg
		path := pkg.Path()
		if !pkgInfo.Importable && strings.HasSuffix(pkgInfo.Pkg.Name(), "_test") {
			// EXTERNAL TEST package, strip the _test to find the path
			path = strings.TrimSuffix(pkg.Path(), "_test")
		}
		idents.pkgScope = generatePackageScopeIdent(pkgInfo)
		for i, f := range pkgInfo.Files {
			fs = prog.Fset
			fname := fs.Position(f.Pos()).Filename
			if strings.HasSuffix(fname, "/C") {
				continue
			}
			b, err := getFileBytes(fname)
			if err != nil {
				fmt.Fprint(os.Stderr, "Error reading file:", err)
				os.Exit(1)
			}
			b = normalizeCRLF(b)
			quotedContents := rawQuote(string(b))
			ast1, fs1 := parseCgoFile(fname, b)
			if ast1 != nil {
				f = ast1
				fs = fs1
			}
			generateGodebugIdentifiers(f)
			ast.Walk(&visitor{context: f, scopeVar: idents.fileScope}, f)
			importName := idents.godebug
			if importName == "godebug" {
				importName = ""
			}
			astutil.AddNamedImport(fs, f, importName, "github.com/mailgun/godebug/lib")
			cfg := printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
			out := writerFor(path, fname)
			defer out.Close()
			_ = cfg.Fprint(out, fs, f)
			fmt.Fprintln(out, "\nvar", idents.fileContents, "=", quotedContents)
			if i == 0 {
				err := generatePackageFile(idents.pkgScope, pkgInfo, out)
				if err != nil {
					fmt.Fprint(os.Stderr, "Error writing package file:", err)
					os.Exit(1)
				}
			}
		}
	}
}

var packageFileTmpl = template.Must(template.New("").Delims("<", ">").Parse(
	`

var <.Scope> = &<.Godebug>.Scope{}

func init() {
	<.Scope>.Vars = map[string]interface{}{<range .Vars>
		"<.>": &<.>,<end>
	}
	<.Scope>.Consts = map[string]interface{}{<range .Consts>
		"<.>": <.>,<end>
	}
	<.Scope>.Funcs = map[string]interface{}{<range .Funcs>
		"<.>": <.>,<end>
	}
}`))

func generatePackageFile(scope string, pkg *loader.PackageInfo, w io.Writer) error {
	data := struct {
		Package, Scope, Godebug string
		Vars, Consts, Funcs     []string
	}{
		Godebug: idents.godebug,
		Package: pkg.Pkg.Name(),
		Scope:   scope,
	}
	for _, f := range pkg.Files {
		for _, decl := range f.Decls {
			switch d := decl.(type) {
			case *ast.FuncDecl:
				if d.Recv == nil && d.Name.Name != "init" {
					data.Funcs = append(data.Funcs, d.Name.Name)
				}
			case *ast.GenDecl:
				switch d.Tok {
				case token.VAR:
					for _, ident := range listNewIdentsFromDecl(d) {
						data.Vars = append(data.Vars, ident.Name)
					}
				case token.CONST:
					for _, ident := range listNewIdentsFromDecl(d) {
						data.Consts = append(data.Consts, ident.Name)
					}
				}
			}
		}
	}
	return packageFileTmpl.Execute(w, data)
}

func varDecl(specs ...ast.Spec) ast.Decl {
	return &ast.GenDecl{Tok: token.VAR, Specs: specs}
}

func newStringLit(v string) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.STRING,
		Value: v,
	}
}

func newInt(n int) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.INT,
		Value: strconv.Itoa(n),
	}
}

func newTerminatingReceiveCase(x ast.Expr) *ast.CommClause {
	return &ast.CommClause{
		Comm: &ast.ExprStmt{
			X: &ast.UnaryExpr{
				Op: token.ARROW,
				X:  x}},
		Body: astPrintf(`panic("impossible")`),
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

func pos2line(pos token.Pos) (line int) {
	return fs.Position(pos).Line
}

func pos2lineString(pos token.Pos) string {
	return strconv.Itoa(pos2line(pos))
}

func isNewIdent(ident *ast.Ident) bool {
	return ident.Name != "_" && defs[ident] != nil
}

func listNewIdentsFromDecl(decl *ast.GenDecl) (idents []*ast.Ident) {
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

func IsBreakpoint(node ast.Node) (b bool) {
	return isOldBreakpoint(node) || isNewBreakpoint(node)
}

func isNewBreakpoint(node ast.Node) (b bool) {
	defer func() {
		if r := recover(); r != nil {
			b = false
		}
	}()
	a := node.(*ast.AssignStmt)
	return a.Lhs[0].(*ast.Ident).Name == "_" && a.Rhs[0].(*ast.BasicLit).Value == `"breakpoint"`
}

func isOldBreakpoint(node ast.Node) (b bool) {
	defer func() {
		if r := recover(); r != nil {
			b = false
		}
	}()
	sel := node.(*ast.ExprStmt).X.(*ast.CallExpr).Fun.(*ast.SelectorExpr)
	return sel.X.(*ast.Ident).Name == idents.godebug && sel.Sel.Name == "SetTrace"
}

type visitor struct {
	context              ast.Node
	stmtBuf              []ast.Stmt
	scopeVar             string
	blockVars            []*ast.Ident
	createdExplicitScope bool
	hasRecovers          bool
	parentIsExprSwitch   bool

	loopState
}

type loopState struct {
	newIdents []*ast.Ident
}

func rewriteFnWithRecovers(body *ast.BlockStmt, fnType *ast.FuncType) (wrapped *ast.FuncLit) {
	// The formatting of the channel declaration is ugly, but it's presented this way here to show how it will look in the actual output.
	// As far as I know, I would need to set the token.Pos values for the left and right braces of the struct and interface type literals
	// in order to get them on one line, but I don't think I can do that without computing all of the other token.Pos values for everything
	// else I generate.
	// TODO: These identifiers will probably conflict if there is a nested function that also has unnamed outputs. Should probably make a better gensym.
	outputDecls, outputs := inputsOrOutputs(fnType.Results, idents.result)
	if len(outputs) > 0 {
		body.List = astPrintf(`%s = func() (%s) {%s}()`, outputs, fnType.Results, body.List)
	}
	body.List = astPrintf(`
		{{%s}}
		_r := make(chan chan interface {
		})
		recovers, panicChan := godebug.EnterFuncWithRecovers(_r, func(ctx *godebug.Context) {
			%s
		})
		for recoverChan := range recovers {
			recoverChan <- recover()
		}
		if panicVal, ok := <-panicChan; ok {
			panic(panicVal)
		}
		{{return %s}}`, outputDecls, body.List, outputs)
	body.Rbrace = token.NoPos // without this I was getting extra whitespace at the end of the function
	return wrapped
}

func (v *visitor) finalizeLoop(pos token.Pos, body *ast.BlockStmt) {
	if body == nil {
		return
	}
	line := pos2line(pos)
	if len(v.newIdents) == 0 {
		call := newCall(idents.godebug, "Line", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar), newInt(line))
		body.List = append(body.List, &ast.ExprStmt{X: call})
	} else {
		body.List = append([]ast.Stmt{
			astPrintf(`godebug.Line(ctx, scope, %s)`, strconv.Itoa(line))[0],
			newDeclareCall(idents.scope, v.newIdents),
		}, body.List...)
	}
}

var blank = ast.NewIdent("_")

func inputsOrOutputs(fieldList *ast.FieldList, prefix string) (decl []ast.Stmt, all []ast.Expr) {
	if fieldList == nil {
		return
	}
	count := 1
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
			decl = append(decl, &ast.DeclStmt{Decl: varDecl(spec)})
		}
	}
	return decl, all
}

func genEnterFunc(fn *ast.FuncDecl, inputs, outputs []ast.Expr) (stmts []ast.Stmt) {
	var (
		pseudoIdent ast.Expr = fn.Name
		recvType    ast.Expr
		ellipsis    string
	)

	// Is this a method call or a function call?
	if fn.Recv != nil {
		// Is the receiver named or anonymous?
		if len(fn.Recv.List[0].Names) == 0 || fn.Recv.List[0].Names[0].Name == "_" {
			pseudoIdent = newSel(idents.receiver, fn.Name.Name)
			recvType = fn.Recv.List[0].Type
		} else {
			pseudoIdent = newSel(fn.Recv.List[0].Names[0].Name, fn.Name.Name)
		}
	}

	// Is the last argument variadic?
	if list := fn.Type.Params.List; len(list) > 0 {
		if _, ok := list[len(list)-1].Type.(*ast.Ellipsis); ok {
			ellipsis = "..."
		}
	}

	// If there are no inputs or outputs, we can skip the function literal wrapper.
	if len(inputs) == 0 && len(outputs) == 0 {
		return astPrintf(`
		{{var receiver %s}}
		ctx, ok := godebug.EnterFunc(%s)
		if !ok {
			return
		}`,
			recvType, pseudoIdent)
	}

	return astPrintf(`
			{{var receiver %s}}
			ctx, ok := godebug.EnterFunc(func() {
				{{%s =}} %s(%s%s)
			})
			if !ok {
				return %s
			}`,
		recvType, outputs, pseudoIdent, inputs, ellipsis, outputs)
}

func genEnterFuncLit(fnType *ast.FuncType, body *ast.BlockStmt, hasRecovers bool) *ast.BlockStmt {
	fn := createConflictFreeName("fn", fnType, false)
	decl, outputs := inputsOrOutputs(fnType.Results, idents.result)
	deferCloseQuit := ""
	if hasRecovers {
		deferCloseQuit = "defer close(quit)"
	}
	newBody := &ast.BlockStmt{}
	if len(outputs) > 0 {
		newBody.List = astPrintf(`
				{{%s}}
				{{%s}}
				%s := func(ctx *godebug.Context) {
					%s = func() (%s) {
						%s
					}()
				}
				if ctx, ok := godebug.EnterFuncLit(%s); ok {
					defer godebug.ExitFunc(ctx)
					%s(ctx)
				}
				return %s
			`, deferCloseQuit, decl, fn, outputs, fnType.Results, body.List, fn, fn, outputs)
	} else {
		newBody.List = astPrintf(`
				{{%s}}
				%s := func(ctx *godebug.Context) {
					%s
				}
				if ctx, ok := godebug.EnterFuncLit(%s); ok {
					defer godebug.ExitFunc(ctx)
					%s(ctx)
				}
				`, deferCloseQuit, fn, body.List, fn, fn)
	}
	return newBody
}

func newIdentsInSimpleStmt(stmt ast.Stmt) (idents []*ast.Ident) {
	if assign, ok := stmt.(*ast.AssignStmt); ok {
		idents = listNewIdentsFromAssign(assign)
	}
	return
}

func newIdentsInRange(_range *ast.RangeStmt) (idents []*ast.Ident) {
	if _range.Tok != token.DEFINE {
		return
	}
	if i, ok := _range.Key.(*ast.Ident); ok && isNewIdent(i) {
		idents = append(idents, i)
	}
	if i, ok := _range.Value.(*ast.Ident); ok && isNewIdent(i) {
		idents = append(idents, i)
	}
	return
}

func (v *visitor) wrapSwitch(_switch *ast.SwitchStmt, identList []*ast.Ident) (block *ast.BlockStmt) {
	block = astPrintf(`
		{
			godebug.Line(ctx, %s, %s)
			%s
			scope := %s.EnteringNewChildScope()
			_ = scope // placeholder
			_ = scope // placeholder
		}`, v.scopeVar, pos2lineString(_switch.Pos()), _switch.Init, v.scopeVar)[0].(*ast.BlockStmt)
	block.List[3] = newDeclareCall(idents.scope, identList)
	_switch.Init = nil
	block.List[4] = _switch
	return block
}

func (v *visitor) wrapLoop(node ast.Stmt, body *ast.BlockStmt) (block *ast.BlockStmt, loop ast.Stmt) {
	block = astPrintf(`
		{
			scope := %s.EnteringNewChildScope()
			_ = scope // placeholder
			godebug.Line(ctx, scope, %s)
		}`, v.scopeVar, pos2lineString(node.Pos()))[0].(*ast.BlockStmt)
	block.List[1] = node
	loop = node
	return
}

func (v *visitor) finalizeNode() {
	switch i := v.context.(type) {

	case *ast.FuncDecl:
		if i.Body == nil {
			break
		}
		if v.hasRecovers {
			rewriteFnWithRecovers(i.Body, i.Type)
			break
		}
		declOuts, outputs := inputsOrOutputs(i.Type.Results, idents.result)
		declIns, inputs := inputsOrOutputs(i.Type.Params, idents.input)
		prepend := append(declIns, declOuts...)
		// We will refer to this function by name when we call genEnterFunc. If any of the
		// parameters have the same name as the function, they will conflict. To get around that,
		// rename any such parameters now.
		rewriteConflictingNames(i)
		prepend = append(prepend, genEnterFunc(i, inputs, outputs)...)
		if !(pkg.Name() == "main" && i.Name.Name == "main") {
			prepend = append(prepend, &ast.DeferStmt{
				Call: newCall(idents.godebug, "ExitFunc", ast.NewIdent(idents.ctx)),
			})
		}

		i.Body.List = append(prepend, i.Body.List...)

	case *ast.FuncLit:
		if v.hasRecovers {
			rewriteFnWithRecovers(i.Body, i.Type)
		} else {
			i.Body = genEnterFuncLit(i.Type, i.Body, v.hasRecovers)
		}

	case *ast.BlockStmt:
		i.List = v.stmtBuf

	case *ast.IfStmt:
		if blk, ok := i.Else.(*ast.BlockStmt); ok {
			elseCall := newCall(idents.godebug, "Line", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar), newInt(pos2line(i.Else.Pos())))
			blk.List = append([]ast.Stmt{&ast.ExprStmt{X: elseCall}}, blk.List...)
		}
		if ifstmt, ok := i.Else.(*ast.IfStmt); ok {
			// We have something that looks like:
			//
			//   if cond {
			//       doThing()
			//   } else if cond2 {
			//       doOtherThing()
			//   }
			//
			// Change it to this:
			//
			//   if cond {
			//       doThing()
			//   } else {
			//       if cond2 {
			//           doOtherThing()
			//       }
			//   }
			//
			// (plus debugging instrumentation)
			var list []ast.Stmt

			// Handle initializer, if it exists.
			if ifstmt.Init != nil {
				list = append(list, newCallStmt(idents.godebug, "ElseIfSimpleStmt", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar), newInt(pos2line(ifstmt.Init.Pos()))))
				list = append(list, ifstmt.Init)
				ifstmt.Init = nil
			}

			// Handle expression.
			list = append(list, newCallStmt(idents.godebug, "ElseIfExpr", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar), newInt(pos2line(ifstmt.Cond.Pos()))))
			list = append(list, ifstmt)

			// Swap in the new block.
			i.Else = &ast.BlockStmt{List: list}
		}

	case *ast.RangeStmt:
		v.finalizeLoop(i.For, i.Body)

	case *ast.ForStmt:
		v.finalizeLoop(i.For, i.Body)

	case *ast.SelectStmt:
		i.Body.List = append(i.Body.List, newTerminatingReceiveCase(
			newCall(idents.godebug, "EndSelect", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar))))

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
		newDecls = append(newDecls, varDecl(&ast.ValueSpec{
			Names:  []*ast.Ident{ast.NewIdent(idents.fileScope)},
			Values: []ast.Expr{newCall(idents.godebug, "EnteringNewFile", ast.NewIdent(idents.pkgScope), ast.NewIdent(idents.fileContents))},
		}))
		i.Decls = append(newDecls, i.Decls...)
	}
}

func stopAtBlockIn(node ast.Node) bool {
	// This is intended to determine whether pausing at the beginning of a block statement will or will not produce
	// something more interesting than "{" on its own line. Hasn't been thought through very carefully.
	switch node.(type) {
	case *ast.BlockStmt:
		return false
	default:
		return true
	}
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	childVisitor := &visitor{context: node, scopeVar: v.scopeVar, parentIsExprSwitch: v.parentIsExprSwitch}

	switch i := node.(type) {

	case nil:
		v.finalizeNode()
		return nil

	case *ast.FuncDecl:
		// TODO: Warning if user tries to put a breakpoint in init.
		// Don't output debugging calls for init functions or empty functions.
		if i.Name.Name == "init" && i.Recv == nil || i.Body == nil || len(i.Body.List) == 0 {
			return nil
		}
		// If there is a call to recover() anywhere in this function, it needs some fairly elaborate treatment.
		childVisitor.blockVars = getIdents(i.Recv, i.Type.Params, i.Type.Results)
		childVisitor.hasRecovers = rewriteRecoversIn(i.Body)
		return childVisitor

	case *ast.FuncLit:
		// If there is a call to recover() anywhere in this function, it needs some fairly elaborate treatment.
		childVisitor.blockVars = getIdents(i.Type.Params, i.Type.Results)
		childVisitor.hasRecovers = rewriteRecoversIn(i.Body)
		return childVisitor

	case *ast.BlockStmt:
		if v.stmtBuf != nil {
			if stopAtBlockIn(v.context) {
				v.stmtBuf = append(v.stmtBuf, newCallStmt(idents.godebug, "Line", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar), newInt(pos2line(node.Pos()))))
			}
			v.stmtBuf = append(v.stmtBuf, i)
		}
		childVisitor.stmtBuf = make([]ast.Stmt, 0, 3*len(i.List))
		if len(v.blockVars) > 0 {
			childVisitor.createScope()
			childVisitor.stmtBuf = append(childVisitor.stmtBuf, newDeclareCall(childVisitor.scopeVar, v.blockVars))
		}
		return childVisitor

	case *ast.SwitchStmt:
		childVisitor.parentIsExprSwitch = true
		newIdents := newIdentsInSimpleStmt(i.Init)
		if len(newIdents) > 0 {
			block := v.wrapSwitch(i, newIdents)
			v.stmtBuf = append(v.stmtBuf, block)
			// wrapSwitch opened a new scope. Switch away from the file-level scope if we haven't already.
			childVisitor.scopeVar = idents.scope
			return childVisitor
		}

	case *ast.TypeSwitchStmt:
		childVisitor.parentIsExprSwitch = false

	case *ast.CaseClause:
		if v.stmtBuf != nil {
			if v.parentIsExprSwitch && len(i.List) > 0 {
				v.stmtBuf = append(v.stmtBuf,
					&ast.CaseClause{
						List: []ast.Expr{
							newCall(idents.godebug, "Case", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar), newInt(pos2line(i.Pos())))},
						// In case this switch is a terminating statement, make this clause be terminating.
						Body: []ast.Stmt{&ast.BranchStmt{Tok: token.FALLTHROUGH}}})
			}
			v.stmtBuf = append(v.stmtBuf, i)
		}
		childVisitor.stmtBuf = make([]ast.Stmt, 0, 3*len(i.List))
		for _, stmt := range i.Body {
			ast.Walk(childVisitor, stmt)
		}
		i.Body = childVisitor.stmtBuf
		if i.List == nil || !v.parentIsExprSwitch { // then this is a default clause or a type switch clause
			i.Body = append([]ast.Stmt{newCallStmt(idents.godebug, "Line", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar), newInt(pos2line(i.Pos())))}, i.Body...)
		}
		return nil

	case *ast.SelectStmt:
		v.stmtBuf = append(v.stmtBuf, newCallStmt(idents.godebug, "Select", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar), newInt(pos2line(node.Pos()))), i)
		return childVisitor

	case *ast.CommClause:
		// Mark this case with a godebug.Comm call if it's not the default.
		if i.Comm != nil { // nil means default case
			v.stmtBuf = append(v.stmtBuf, newTerminatingReceiveCase(newCall(idents.godebug, "Comm", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar), newInt(pos2line(i.Pos())))))
		}

		// Manually walk its descendants.
		v.stmtBuf = append(v.stmtBuf, i)
		childVisitor.Visit(i.Comm)
		childVisitor.stmtBuf = make([]ast.Stmt, 0, 3*len(i.Body))

		// Declare any new variables.
		if newIdents := newIdentsInSimpleStmt(i.Comm); newIdents != nil {
			childVisitor.createScope()
			childVisitor.stmtBuf = append(childVisitor.stmtBuf, newDeclareCall(childVisitor.scopeVar, newIdents))
		}

		for _, node := range i.Body {
			childVisitor.Visit(node)
		}
		i.Body = append([]ast.Stmt{newCallStmt(idents.godebug, "Line", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar), newInt(pos2line(i.Pos())))}, childVisitor.stmtBuf...)

		return nil
	}

	// The code below is about inserting debug calls. Skip it if we're not in a context where that makes sense.
	if v.stmtBuf == nil {
		return childVisitor
	}

	// If this is a loop that declares identifiers, wrap it in a block statement so we can declare identifiers at the right time.
	switch i := node.(type) {
	case *ast.ForStmt:
		newIdents := newIdentsInSimpleStmt(i.Init)
		if len(newIdents) > 0 {
			block, loop := v.wrapLoop(i, i.Body)
			v.stmtBuf = append(v.stmtBuf, block)
			childVisitor.context = loop
			childVisitor.loopState = loopState{newIdents: newIdents}
			// wrapLoop opened a new scope. Switch away from the file-level scope if we haven't already.
			childVisitor.scopeVar = idents.scope
			return childVisitor
		}
	case *ast.RangeStmt:
		newIdents := newIdentsInRange(i)
		if len(newIdents) > 0 {
			block, loop := v.wrapLoop(i, i.Body)
			v.stmtBuf = append(v.stmtBuf, block)
			childVisitor.context = loop
			childVisitor.loopState = loopState{newIdents: newIdents}
			// wrapLoop opened a new scope. Switch away from the file-level scope if we haven't already.
			childVisitor.scopeVar = idents.scope
			return childVisitor
		}
	}

	if !IsBreakpoint(node) {
		v.stmtBuf = append(v.stmtBuf, newCallStmt(idents.godebug, "Line", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar), newInt(pos2line(node.Pos()))))
	}

	// Copy the statement into the new block we are building.
	if stmt, ok := node.(ast.Stmt); ok {
		if IsBreakpoint(node) {
			// Rewrite `godebug.SetTrace()` and `_ = "breakpoint"` as `godebug.SetTraceGen(ctx)`.
			v.stmtBuf = append(v.stmtBuf, astPrintf("godebug.SetTraceGen(ctx)")[0], newCallStmt(idents.godebug, "Line", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar), newInt(pos2line(node.Pos()))))
		} else {
			v.stmtBuf = append(v.stmtBuf, stmt)
		}
	}

	// If this statement declared new variables, output a Declare call.
	var newIdents []*ast.Ident
	switch i := node.(type) {
	case *ast.DeclStmt:
		if gen := i.Decl.(*ast.GenDecl); gen.Tok == token.VAR {
			newIdents = listNewIdentsFromDecl(gen)
		}
	case *ast.AssignStmt:
		newIdents = listNewIdentsFromAssign(i)
	}
	if len(newIdents) > 0 {
		if !v.createdExplicitScope {
			v.createScope()
		}
		v.stmtBuf = append(v.stmtBuf, newDeclareCall("", newIdents))
	}

	// If this statement declared new constants, output a Constant call.
	if decl, ok := node.(*ast.DeclStmt); ok {
		if gen := decl.Decl.(*ast.GenDecl); gen.Tok == token.CONST {
			newIdents = listNewIdentsFromDecl(gen)
			if len(newIdents) > 0 {
				if !v.createdExplicitScope {
					v.createScope()
				}
				v.stmtBuf = append(v.stmtBuf, newConstantCall("", newIdents))
			}
		}
	}

	// If this is a defer statement, defer another function right after it that will let the user step into it if they wish.
	if _, isDefer := node.(*ast.DeferStmt); isDefer {
		v.stmtBuf = append(v.stmtBuf, astPrintf(`defer godebug.Defer(ctx, %s, %s)`, v.scopeVar, pos2lineString(node.Pos()))[0])
	}

	if _if, ok := node.(*ast.IfStmt); ok {
		childVisitor.blockVars = newIdentsInSimpleStmt(_if.Init)
	}
	return childVisitor
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

func newDeclareCall(scopeVar string, newVars []*ast.Ident) ast.Stmt {
	return newIdentsCall(scopeVar, newVars, false)
}

func newConstantCall(scopeVar string, newConsts []*ast.Ident) ast.Stmt {
	return newIdentsCall(scopeVar, newConsts, true)
}

func newIdentsCall(scopeVar string, newIdents []*ast.Ident, isConst bool) ast.Stmt {
	if scopeVar == "" {
		scopeVar = idents.scope
	}
	f := "Declare"
	if isConst {
		f = "Constant"
	}
	expr := newCallStmt(scopeVar, f)
	call := expr.X.(*ast.CallExpr)
	call.Args = make([]ast.Expr, 2*len(newIdents))
	for i, ident := range newIdents {
		// Quoted identifier name.
		call.Args[2*i] = newStringLit(strconv.Quote(ident.Name))

		if isConst {
			// Pass the value if constant. Cast it if it doesn't fit an int32.
			call.Args[2*i+1] = castIfOverflow(ident)
		} else {
			// Pass a pointer if variable.
			call.Args[2*i+1] = &ast.UnaryExpr{
				Op: token.AND,
				X:  ident,
			}
		}
	}
	return expr
}

var (
	minInt32 = exact.MakeInt64(math.MinInt32)
	maxInt32 = exact.MakeInt64(math.MaxInt32)
)

func castIfOverflow(ident *ast.Ident) ast.Expr {
	c, ok := defs[ident].(*types.Const)
	if !ok || c == nil {
		return ident
	}
	v := c.Val()
	switch {
	case v.Kind() != exact.Int:
		return ident
	case exact.Compare(v, token.LSS, minInt32):
		return &ast.CallExpr{Fun: ast.NewIdent("int64"), Args: []ast.Expr{ident}}
	case exact.Compare(v, token.GTR, maxInt32):
		return &ast.CallExpr{Fun: ast.NewIdent("uint64"), Args: []ast.Expr{ident}}
	default:
		return ident
	}
}

func (v *visitor) createScope() {
	name := idents.scope
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

var idents struct {
	ctx, ok, pkgScope, scope, receiver, fileScope, fileContents, godebug, result, input, recoverChan, recoverChanChan, recovers, panicVal, panicChan string
}

func generateGodebugPkgName(f *ast.File) string {
	var pkgName string
	for _, imp := range f.Imports {
		if imp.Path.Value == `"github.com/mailgun/godebug/lib"` {
			pkgName = "godebug"
			if imp.Name != nil {
				if imp.Name.Name == "_" {
					imp.Name.Name = createConflictFreeName("godebug", f, false)
				}
				pkgName = imp.Name.Name
			}
			break
		}
	}
	if pkgName == "" {
		pkgName = createConflictFreeName("godebug", f, false)
	}
	return pkgName
}

func generatePackageScopeIdent(pkg *loader.PackageInfo) string {
	v := &nameVisitor{
		base:      pkg.Pkg.Name() + "_pkg_scope",
		suffix:    false,
		conflicts: make(map[string]bool),
	}
	for _, f := range pkg.Files {
		ast.Walk(v, f)
	}
	return v.getName()
}

func generateGodebugIdentifiers(f *ast.File) {
	// Variables that won't have suffixes.
	idents.ctx = createConflictFreeName("ctx", f, false)
	idents.ok = createConflictFreeName("ok", f, false)
	idents.scope = createConflictFreeName("scope", f, false)
	idents.receiver = createConflictFreeName("receiver", f, false)
	idents.recoverChan = createConflictFreeName("rr", f, false)
	idents.recoverChanChan = createConflictFreeName("r", f, false)
	idents.recovers = createConflictFreeName("recovers", f, false)
	idents.panicVal = createConflictFreeName("v", f, false)
	idents.panicChan = createConflictFreeName("panicChan", f, false)

	idents.godebug = generateGodebugPkgName(f)

	// Variables that will have suffixes.
	idents.result = createConflictFreeName("result", f, true)
	idents.input = createConflictFreeName("input", f, true)

	// Variables with names derived from the filename.
	base := strings.Map(func(r rune) rune {
		if !unicode.In(r, unicode.Digit, unicode.Letter) {
			return '_'
		}
		return r
	}, filepath.Base(fs.Position(f.Pos()).Filename))
	if !unicode.IsLetter(rune(base[0])) {
		// identifiers must start with letters
		base = "a" + base
	}
	idents.fileScope = createConflictFreeName(base+"_scope", f, false)
	idents.fileContents = createConflictFreeName(base+"_contents", f, false)
}

func createConflictFreeName(name string, parent ast.Node, hasSuffix bool) string {
	// Visit all descendants of parent and check for usage of name. Prepend underscores until there are no conflicts, then return.
	v := &nameVisitor{base: name, suffix: hasSuffix, conflicts: make(map[string]bool)}
	ast.Walk(v, parent)
	return v.getName()
}

func createConflictFreeNameCheckIdents(name string, parent ast.Node) string {
	v := &nameVisitor{
		base:   name,
		suffix: false,
		conflicts: map[string]bool{
			idents.ctx:          true,
			idents.fileContents: true,
			idents.fileScope:    true,
			idents.godebug:      true,
			idents.input:        true,
			idents.ok:           true,
			idents.receiver:     true,
			idents.result:       true,
			idents.scope:        true,
		},
	}
	ast.Walk(v, parent)
	return v.getName()
}

type nameVisitor struct {
	base      string
	suffix    bool
	conflicts map[string]bool
}

// getName returns a name that does not conflict with any identifiers observed while visiting nodes.
func (v *nameVisitor) getName() (name string) {
	for name = v.base; v.conflicts[name]; name = "_" + name {
	}
	return
}

func (v *nameVisitor) Visit(node ast.Node) ast.Visitor {
	switch i := node.(type) {

	// Some identifiers will not cause conflicts and can be ignored:
	case *ast.SelectorExpr:
		// For a selector expression x.f, identifiers in x can cause conflicts, but the identifier f will not.
		ast.Walk(v, i.X)
		return nil
	case *ast.LabeledStmt:
		// Labels do not conflict with identifiers that are not labels. Ignore the label, but walk its statement.
		ast.Walk(v, i.Stmt)
		return nil
	case *ast.FuncDecl:
		// Method names will not cause conflicts, but other parts of a function declaration can.
		if i.Recv != nil {
			ast.Walk(v, i.Recv)
			ast.Walk(v, i.Type)
			ast.Walk(v, i.Body)
			return nil
		}
	case *ast.InterfaceType, *ast.StructType:
		// Identifiers within interfaces and structs will not cause conflicts.
		return nil

	// Any other identifier we reach should be checked for a conflicting name.
	case *ast.Ident:
		name := i.Name
		if v.suffix {
			name = strings.TrimRight(name, "0123456789")
		}
		if strings.TrimLeft(name, "_") == v.base {
			v.conflicts[name] = true
		}
	}
	return v
}

func rewriteRecoversIn(block *ast.BlockStmt) bool {
	var result bool
	visitor := recoverVisitor{nil, &result}
	ast.Walk(&visitor, block)
	return result
}

type recoverVisitor struct {
	parent     ast.Node
	didRewrite *bool
}

func (v *recoverVisitor) Visit(node ast.Node) ast.Visitor {
	switch x := node.(type) {
	case *ast.CallExpr:
		if isBuiltinFunc(x.Fun, "recover") {
			rewriteRecoverCall(v.parent, node)
			*v.didRewrite = true
			return nil
		}
	case *ast.FuncLit:
		// Ignore recover calls in nested function literals.
		return nil
	}
	return &recoverVisitor{node, v.didRewrite}
}

func rewriteRecoverCall(parent, _recover ast.Node) {
	rewritten, _ := parser.ParseExpr(fmt.Sprintf("<-(<-%s)", idents.recoverChanChan))
	rewritten.(*ast.UnaryExpr).OpPos = _recover.Pos()
	v := reflect.ValueOf(parent).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Interface() == _recover {
			f.Set(reflect.ValueOf(rewritten))
			return
		}
		if f.Kind() == reflect.Slice {
			for j := 0; j < f.Len(); j++ {
				if f.Index(j).Interface() == _recover {
					f.Index(j).Set(reflect.ValueOf(rewritten))
				}
			}
		}
	}
}

func parseCgoFile(filename string, bytes []byte) (*ast.File, *token.FileSet) {
	fs := token.NewFileSet()
	ast1, err := parser.ParseFile(fs, filename, bytes, parser.ParseComments)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		os.Exit(1)
	}
	for _, spec := range ast1.Imports {
		if spec.Path.Value == `"C"` {
			return ast1, fs
		}
	}
	return nil, nil
}

func rewriteConflictingNames(fn *ast.FuncDecl) {
	for _, fieldList := range []*ast.FieldList{fn.Recv, fn.Type.Params, fn.Type.Results} {
		if fieldList == nil {
			continue
		}
		for _, f := range fieldList.List {
			for _, name := range f.Names {
				if name.Name == fn.Name.Name {
					oldName := name.Name
					newName := createConflictFreeNameCheckIdents(name.Name, fn)
					rewriteFn := func(node ast.Node) bool {
						if ident, ok := node.(*ast.Ident); ok && ident.Name == oldName {
							ident.Name = newName
						}
						return true
					}
					// Instead of walking all of fn, walk fn.Recv, fn.Type, and fn.Body.
					// If we walked all of fn we would rewrite the name of the function itself
					// in addition to the parameter we are rewriting.
					if fn.Recv != nil && fn.Recv.List != nil {
						ast.Inspect(fn.Recv, rewriteFn)
					}
					if fn.Type != nil {
						ast.Inspect(fn.Type, rewriteFn)
					}
					if fn.Body != nil {
						ast.Inspect(fn.Body, rewriteFn)
					}
					return // at most one parameter can share a name with the function
				}
			}
		}
	}
}

func isBuiltinFunc(fn ast.Expr, name string) bool {
	ident, ok := fn.(*ast.Ident)
	return ok && ident.Name == name && _types[fn].IsBuiltin()
}

func normalizeCRLF(b []byte) []byte {
	return bytes.Replace(b, []byte("\r\n"), []byte("\n"), -1)
}
