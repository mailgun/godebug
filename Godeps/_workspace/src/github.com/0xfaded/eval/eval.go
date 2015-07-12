package eval

import (
	"fmt"
	"reflect"

	"go/ast"
	"go/parser"
	"go/scanner"
)

var emptyEnv Env = MakeSimpleEnv()

func Eval(expr string) (result []reflect.Value, panik error, compileErrors []error) {
	return EvalEnv(expr, emptyEnv)
}

func EvalEnv(expr string, env Env) (result []reflect.Value, panik error, compileErrors []error) {
	if e, err := parser.ParseExpr(expr); err != nil {
		errs := err.(scanner.ErrorList)
		for i := range errs {
			compileErrors = append(compileErrors, errs[i])
		}
	} else if cexpr, errs := CheckExpr(e, env); errs != nil {
		compileErrors = errs
	} else {
		result, panik = EvalExpr(cexpr, env)
	}
	return
}

func Interpret(stmt string, env Env) (result []reflect.Value, panik error, compileErrors []error) {
	if s, err := ParseStmt(stmt); err != nil {
		if errs, ok := err.(scanner.ErrorList); ok {
			for i := range errs {
				compileErrors = append(compileErrors, errs[i])
			}
		} else {
			compileErrors = append(compileErrors, err)
		}
	} else if e, ok := s.(*ast.ExprStmt); ok {
		if cexpr, errs := CheckExpr(e.X, env); errs != nil {
			compileErrors = errs
		} else {
			result, panik = EvalExpr(cexpr, env)
		}
	} else {
		if cstmt, errs := CheckStmt(s, env); errs != nil {
			compileErrors = errs
		} else {
			_, panik = InterpStmt(cstmt, env)
		}
	}
	return
}

// A convenience function for parsing a Stmt by wrapping it in a func literal
func ParseStmt(stmt string) (ast.Stmt, error) {
	// parser.ParseExpr has some quirks, for example it won't parse unused map literals as
	// ExprStmts. Unused map literals are technically illegal, but it would be
	// nice to check them at a later stage. Therefore, we want to parse expressions
	// as expressions, and stmts and stmts. We try both.
	// However, there is a bug in parser.ParseExpr that it does not detect excess input.
	// Therefore, the _ of _ = 1 will be parsed as an expression. To avoid this, attempt
	// to parse the input as a statement first, and fall back to an expression
	expr := "func(){" + stmt + ";}"
	if e, err := parser.ParseExpr(expr); err != nil {
		if e, err := parser.ParseExpr(stmt); err == nil {
			return &ast.ExprStmt{X: e}, nil
		}
		errs := err.(scanner.ErrorList)
		for i := range errs {
			errs[i].Pos.Offset -= 7
			errs[i].Pos.Column -= 7
		}
		return nil, errs
	} else {
		node := e.(*ast.FuncLit).Body.List[0]
		if stmt, ok := node.(Stmt); !ok {
			return nil, fmt.Errorf("%T not supported", node)
		} else {
			return stmt, nil
		}
	}
}

