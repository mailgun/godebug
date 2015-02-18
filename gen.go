package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/ast/astutil"
	_ "github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/gcimporter"
	"github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/loader"
	"github.com/mailgun/godebug/Godeps/_workspace/src/golang.org/x/tools/go/types"
	_ "github.com/mailgun/godebug/lib" // so the library is also installed whenever this package is
	_ "golang.org/x/tools/go/gccgoimporter"
)

var w = flag.Bool("w", false, "write result to (source) file instead of stdout")

var (
	defs   map[*ast.Ident]types.Object
	_types map[ast.Expr]types.TypeAndValue
	fs     *token.FileSet
	pkg    *types.Package
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

type blankLineStripper struct {
	io.Writer
	lastWasNewline bool
	buf            bytes.Buffer
}

func (w *blankLineStripper) Write(p []byte) (n int, err error) {
	w.buf.Reset()
	for _, b := range p {
		if !(w.lastWasNewline && b == '\n') {
			w.buf.WriteByte(b)
		}
		w.lastWasNewline = b == '\n'
	}
	n = len(p) - len(w.buf.Bytes())
	nn, err := w.Writer.Write(w.buf.Bytes())
	return n + nn, err
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
	for _, pkgInfo := range prog.InitialPackages() {
		defs = pkgInfo.Defs
		_types = pkgInfo.Types
		pkg = pkgInfo.Pkg
		for _, f := range pkgInfo.Files {
			fs = prog.Fset
			fname := fs.Position(f.Pos()).Filename
			if strings.HasSuffix(fname, "/C") {
				continue
			}
			b, err := ioutil.ReadFile(fname)
			if err != nil {
				fmt.Fprint(os.Stderr, "Error reading file:", err)
				os.Exit(1)
			}
			quotedContents := rawQuote(string(b))
			ast1, fs1 := parseCgoFile(fname)
			if ast1 != nil {
				f = ast1
				fs = fs1
			}
			generateGodebugIdentifiers(f)
			for _, imp := range f.Imports {
				if imp.Path.Value == `"github.com/mailgun/godebug/lib"` {
					idents.godebug = "godebug"
					if imp.Name != nil {
						idents.godebug = imp.Name.Name
					}
					break
				}
			}
			if idents.godebug == "" {
				idents.godebug = createConflictFreeName("godebug", f, false)
			}
			ast.Walk(&visitor{context: f, scopeVar: idents.fileScope}, f)
			importName := idents.godebug
			if importName == "godebug" {
				importName = ""
			}
			astutil.AddNamedImport(fs, f, importName, "github.com/mailgun/godebug/lib")
			cfg := printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
			var out io.Writer = os.Stdout
			if *w {
				file, err := os.Create(fs.Position(f.Pos()).Filename)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				defer file.Close()
				out = file
			}
			cfg.Fprint(&blankLineStripper{Writer: out}, fs, f)
			fmt.Fprintln(out, "\nvar", idents.fileContents, "=", quotedContents)
		}
	}
}

func varDecl(specs ...ast.Spec) ast.Decl {
	return &ast.GenDecl{Tok: token.VAR, Specs: specs}
}

func newInt(n int) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.INT,
		Value: strconv.Itoa(n),
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
	// TODO: Either change this entrypoint or make it compatible with the user importing godebug by another name.
	return sel.X.(*ast.Ident).Name == "godebug" && sel.Sel.Name == "SetTrace"
}

type visitor struct {
	context              ast.Node
	stmtBuf              []ast.Stmt
	scopeVar             string
	blockVars            []*ast.Ident
	createdExplicitScope bool
	hasRecovers          bool
	loopState
}

type loopState struct {
	newIdents []*ast.Ident
}

func rewriteFnWithRecovers(body *ast.BlockStmt) (wrapped *ast.FuncLit) {
	// The formatting of the channel declarations is ugly, but it's presented this way here to show how it will look in the actual output.
	// As far as I know, I would need to set the token.Pos values for the left and right braces of the struct and interface type literals
	// in order to get them on one line, but I don't think I can do that without computing all of the other token.Pos values for everything
	// else I generate.
	newBody := astPrintf(`
		quit := make(chan struct {
		})
		_godebug_recover_chan_ := make(chan chan interface {
		})
		rr := make(chan interface {
		})
		godebug.Go(func() {
		})
		for {
			select {
			case <-quit:
				return
			case _godebug_recover_chan_ <- rr:
				rr <- recover()
			}
		}`)
	wrapped = newBody[3].(*ast.ExprStmt).X.(*ast.CallExpr).Args[0].(*ast.FuncLit)
	wrapped.Body.List = body.List
	body.List = newBody
	body.Rbrace = token.NoPos // without this I was getting extra whitespace at the end of the function
	return wrapped
}

func (v *visitor) finalizeLoop(pos token.Pos, body *ast.BlockStmt) {
	if body == nil {
		return
	}
	line := fs.Position(pos).Line
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

func (v *visitor) ifElseCondWrap(cond ast.Expr, pos token.Pos) ast.Expr {
	return astPrintfExpr(`func() bool {
		godebug.ElseIfExpr(ctx, %s, %s)
		return %s
	}()
	`, v.scopeVar, strconv.Itoa(fs.Position(pos).Line), cond)
}

func (v *visitor) ifElseInitWrap(vars []ast.Expr, vals []ast.Expr, pos token.Pos) ast.Expr {
	results := make([]string, len(vars))
	for i, expr := range vars {
		ident, ok := expr.(*ast.Ident)
		if !ok {
			panic(fmt.Sprintf("Unsupported type in if statement initializer: %T. Sorry! Let me (jeremy@mailgunhq.com) know about this and I'll fix it.", expr))
		}
		results[i] = types.TypeString(pkg, defs[ident].Type())
	}
	return astPrintfExpr(`func() (%s) {
		godebug.ElseIfSimpleStmt(ctx, %s, %s)
		return %s
	}()`, strings.Join(results, ", "), v.scopeVar, strconv.Itoa(fs.Position(pos).Line), vals)
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
		if len(fn.Recv.List[0].Names) == 0 {
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

func (v *visitor) wrapLoop(node ast.Stmt, body *ast.BlockStmt) (block *ast.BlockStmt, loop ast.Stmt) {
	block = astPrintf(`
		{
			scope := %s.EnteringNewChildScope()
			_ = scope // placeholder
			godebug.Line(ctx, scope, %s)
		}`, v.scopeVar, strconv.Itoa(fs.Position(node.Pos()).Line))[0].(*ast.BlockStmt)
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
			i.Body = genEnterFuncLit(i.Type, i.Body, true)
			rewriteFnWithRecovers(i.Body)
			break
		}
		declOuts, outputs := inputsOrOutputs(i.Type.Results, idents.result)
		declIns, inputs := inputsOrOutputs(i.Type.Params, idents.input)
		prepend := append(declIns, declOuts...)
		prepend = append(prepend, genEnterFunc(i, inputs, outputs)...)
		if !(pkg.Name() == "main" && i.Name.Name == "main") {
			prepend = append(prepend, &ast.DeferStmt{
				Call: newCall(idents.godebug, "ExitFunc", ast.NewIdent(idents.ctx)),
			})
		}

		i.Body.List = append(prepend, i.Body.List...)

	case *ast.FuncLit:
		i.Body = genEnterFuncLit(i.Type, i.Body, v.hasRecovers)
		if v.hasRecovers {
			rewriteFnWithRecovers(i.Body)
		}

	case *ast.BlockStmt:
		i.List = v.stmtBuf

	case *ast.IfStmt:
		if blk, ok := i.Else.(*ast.BlockStmt); ok {
			elseCall := newCall(idents.godebug, "Line", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar), newInt(fs.Position(i.Else.Pos()).Line))
			blk.List = append([]ast.Stmt{&ast.ExprStmt{X: elseCall}}, blk.List...)
		}
		if ifstmt, ok := i.Else.(*ast.IfStmt); ok {
			switch d := ifstmt.Init.(type) {
			case *ast.AssignStmt:
				d.Rhs = []ast.Expr{v.ifElseInitWrap(d.Lhs, d.Rhs, i.Else.Pos())}
			case *ast.DeclStmt:
				spec := d.Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec)
				exprs := make([]ast.Expr, len(spec.Names))
				for i := range exprs {
					exprs[i] = ast.Expr(spec.Names[i])
				}
				spec.Values = []ast.Expr{v.ifElseInitWrap(exprs, spec.Values, i.Else.Pos())}
			// TODO: optimize nil case
			default:
			}
			ifstmt.Cond = v.ifElseCondWrap(ifstmt.Cond, i.Else.Pos())
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
		newDecls = append(newDecls, varDecl(&ast.ValueSpec{
			Names:  []*ast.Ident{ast.NewIdent(idents.fileScope)},
			Values: []ast.Expr{newCall(idents.godebug, "EnteringNewScope", ast.NewIdent(idents.fileContents))},
		}))
		i.Decls = append(newDecls, i.Decls...)
	}
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	switch i := node.(type) {

	case nil:
		v.finalizeNode()
		return nil

	case *ast.FuncDecl:
		// Don't output debugging calls for init functions or empty functions.
		if i.Name.Name == "init" && i.Recv == nil || i.Body == nil || len(i.Body.List) == 0 {
			return nil
		}
		// If there is a call to recover() anywhere in this function, it needs some fairly elaborate treatment.
		didRewrite := rewriteRecoversIn(i.Body)
		return &visitor{context: node, blockVars: getIdents(i.Recv, i.Type.Params, i.Type.Results), scopeVar: idents.fileScope, hasRecovers: didRewrite}

	case *ast.FuncLit:
		// If there is a call to recover() anywhere in this function, it needs some fairly elaborate treatment.
		didRewrite := rewriteRecoversIn(i.Body)
		return &visitor{context: node, blockVars: getIdents(i.Type.Params, i.Type.Results), scopeVar: v.scopeVar, hasRecovers: didRewrite}

	case *ast.BlockStmt:
		if v.stmtBuf != nil {
			v.stmtBuf = append(v.stmtBuf, newCallStmt(idents.godebug, "Line", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar), newInt(fs.Position(node.Pos()).Line)))
			v.stmtBuf = append(v.stmtBuf, i)
		}
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

	// The code below is about inserting debug calls. Skip it if we're not in a context where that makes sense.
	if v.stmtBuf == nil {
		return &visitor{context: node, scopeVar: v.scopeVar}
	}

	// If this is a loop that declares identifiers, wrap it in a block statement so we can declare identifiers at the right time.
	switch i := node.(type) {
	case *ast.ForStmt:
		newIdents := newIdentsInSimpleStmt(i.Init)
		if len(newIdents) > 0 {
			block, loop := v.wrapLoop(i, i.Body)
			v.stmtBuf = append(v.stmtBuf, block)
			return &visitor{context: loop, scopeVar: v.scopeVar, loopState: loopState{newIdents: newIdents}}
		}
	case *ast.RangeStmt:
		newIdents := newIdentsInRange(i)
		if len(newIdents) > 0 {
			block, loop := v.wrapLoop(i, i.Body)
			v.stmtBuf = append(v.stmtBuf, block)
			return &visitor{context: loop, scopeVar: v.scopeVar, loopState: loopState{newIdents: newIdents}}
		}
	}

	if !isSetTraceCall(node) {
		v.stmtBuf = append(v.stmtBuf, newCallStmt(idents.godebug, "Line", ast.NewIdent(idents.ctx), ast.NewIdent(v.scopeVar), newInt(fs.Position(node.Pos()).Line)))
	}

	// Copy the statement into the new block we are building.
	if stmt, ok := node.(ast.Stmt); ok {
		if isSetTraceCall(node) {
			// Rewrite godebug.SetTrace() as godebug.SetTraceGen(ctx)
			call := stmt.(*ast.ExprStmt).X.(*ast.CallExpr)
			call.Args = []ast.Expr{ast.NewIdent(idents.ctx)}
			call.Fun.(*ast.SelectorExpr).Sel.Name = "SetTraceGen"
		}
		v.stmtBuf = append(v.stmtBuf, stmt)
	}

	// Create the next visitor now, because v.scopeVar may change below.
	w := &visitor{context: node, scopeVar: v.scopeVar}

	// If this statement declared new identifiers, output a Declare call.
	var newIdents []*ast.Ident
	switch i := node.(type) {
	case *ast.DeclStmt:
		newIdents = listNewIdentsFromDecl(i.Decl.(*ast.GenDecl))
	case *ast.AssignStmt:
		newIdents = listNewIdentsFromAssign(i)
	}
	if len(newIdents) > 0 {
		if !v.createdExplicitScope {
			v.createScope()
		}
		v.stmtBuf = append(v.stmtBuf, newDeclareCall("", newIdents))
	}

	// If this is a defer statement, defer another function right after it that will let the user step into it if they wish.
	if _, isDefer := node.(*ast.DeferStmt); isDefer {
		v.stmtBuf = append(v.stmtBuf, astPrintf(`defer godebug.Defer(ctx, %s, %s)`, v.scopeVar, strconv.Itoa(fs.Position(node.Pos()).Line))[0])
	}

	if _if, ok := node.(*ast.IfStmt); ok {
		w.blockVars = newIdentsInSimpleStmt(_if.Init)
	}
	return w
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
	if scopeVar == "" {
		scopeVar = idents.scope
	}
	expr := newCallStmt(scopeVar, "Declare")
	call := expr.X.(*ast.CallExpr)
	call.Args = make([]ast.Expr, 2*len(newVars))
	for i, _var := range newVars {
		call.Args[2*i] = &ast.BasicLit{
			Kind:  token.STRING,
			Value: strconv.Quote(_var.Name),
		}
		call.Args[2*i+1] = &ast.UnaryExpr{
			Op: token.AND,
			X:  _var,
		}
	}
	return expr
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
	ctx, ok, scope, receiver, fileScope, fileContents, godebug, result, input string
}

func generateGodebugIdentifiers(f *ast.File) {
	// Variables that won't have suffixes.
	idents.ctx = createConflictFreeName("ctx", f, false)
	idents.ok = createConflictFreeName("ok", f, false)
	idents.scope = createConflictFreeName("scope", f, false)
	idents.receiver = createConflictFreeName("receiver", f, false)

	// godebug is set elsewhere.
	//idents.godebug = createConflictFreeName("godebug", f, false)

	// Variables that will have suffixes.
	idents.result = createConflictFreeName("result", f, true)
	idents.input = createConflictFreeName("input", f, true)

	// Variables with names derived from the filename.
	base := strings.Map(func(r rune) rune {
		if !unicode.In(r, unicode.Digit, unicode.Letter) {
			return '_'
		}
		return r
	}, path.Base(fs.Position(f.Pos()).Filename))
	idents.fileScope = createConflictFreeName(base+"_scope", f, false)
	idents.fileContents = createConflictFreeName(base+"_contents", f, false)
}

func createConflictFreeName(name string, parent ast.Node, hasSuffix bool) string {
	// Visit all descendants of parent and check for usage of name. Prepend underscores until there are no conflicts, then return.
	v := &nameVisitor{base: name, suffix: hasSuffix, conflicts: make(map[string]bool)}
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
		if ident, ok := x.Fun.(*ast.Ident); ok {
			if ident.Name == "recover" && _types[x.Fun].IsBuiltin() {
				rewriteRecoverCall(v.parent, node)
				*v.didRewrite = true
				return nil
			}
		}
	case *ast.FuncLit:
		// Ignore recover calls in nested function literals.
		return nil
	}
	return &recoverVisitor{node, v.didRewrite}
}

func rewriteRecoverCall(parent, _recover ast.Node) {
	rewritten := astPrintfExpr("<-(<-_godebug_recover_chan_)")
	rewritten.(*ast.UnaryExpr).OpPos = _recover.Pos()
	v := reflect.ValueOf(parent).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Interface() == _recover {
			f.Set(reflect.ValueOf(rewritten))
			return
		}
		if f.Kind() == reflect.Slice {
			for i := 0; i < f.Len(); i++ {
				if f.Index(i).Interface() == _recover {
					f.Index(i).Set(reflect.ValueOf(rewritten))
				}
			}
		}
	}
}

func parseCgoFile(filename string) (*ast.File, *token.FileSet) {
	fs := token.NewFileSet()
	ast1, err := parser.ParseFile(fs, filename, nil, parser.ParseComments)
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
