package eval

import (
	"go/ast"
)

type exprVisitor interface {
	visit(expr Expr) bool
}

func walk(expr ast.Expr, visitor exprVisitor) {
	if expr == nil {
		return
	}
	switch expr := expr.(type) {
	case *BadExpr:
		visitor.visit(expr)
	case *Ident:
		visitor.visit(expr)
	case *Ellipsis:
		visitor.visit(expr)
	case *BasicLit:
		visitor.visit(expr)
	case *FuncLit:
		visitor.visit(expr)
	case *CompositeLit:
		if visitor.visit(expr) {
			for _, elt := range expr.Elts {
				walk(elt, visitor)
			}
		}
	case *ParenExpr:
		if visitor.visit(expr) {
			walk(expr.X, visitor)
		}
	case *SelectorExpr:
		if visitor.visit(expr) {
			walk(expr.X, visitor)
		}
	case *IndexExpr:
		if visitor.visit(expr) {
			walk(expr.Index, visitor)
			walk(expr.X, visitor)
		}
	case *SliceExpr:
		if visitor.visit(expr) {
			walk(expr.Low, visitor)
			walk(expr.High, visitor)
			// TODO[crc] go 1.2 introduces the [::] notation. Add after upgrade
			// walk(expr.Max, visitor)
		}
	case *TypeAssertExpr:
		if visitor.visit(expr) {
			walk(expr.X, visitor)
		}
	case *CallExpr:
		if visitor.visit(expr) {
			for _, arg := range expr.Args {
				walk(arg, visitor)
			}
		}
	case *StarExpr:
		if visitor.visit(expr) {
			walk(expr.X, visitor)
		}
	case *UnaryExpr:
		if visitor.visit(expr) {
			walk(expr.X, visitor)
		}
	case *BinaryExpr:
		if visitor.visit(expr) {
			walk(expr.X, visitor)
			walk(expr.Y, visitor)
		}
	case *KeyValueExpr:
		if visitor.visit(expr) {
			walk(expr.Key, visitor)
			walk(expr.Value, visitor)
		}
	}
}
