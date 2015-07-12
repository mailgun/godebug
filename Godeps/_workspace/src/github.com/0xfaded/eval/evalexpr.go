package eval

import (
	"reflect"
)

// EvalExpr is the main function to call to evaluate an ast-parsed
// expression, expr. 
// Parameter env, contains an evaluation environment from
// which to get reflect.Values from. Note however that env can be
// subverted somewhat by supplying callback hooks routines which
// access variables and by supplying user-defined conversion routines.
func EvalExpr(expr Expr, env Env) ([]reflect.Value, error) {
	switch node := expr.(type) {
	case *Ident:
		v, err := evalIdent(node, env)
		return []reflect.Value{v}, err
	case *BasicLit:
		v, err := evalBasicLit(node)
		return []reflect.Value{v}, err
	case *FuncLit:
		v, err := evalFuncLit(node, env)
		return []reflect.Value{v}, err
	case *CompositeLit:
		v, err := evalCompositeLit(node, env)
		return []reflect.Value{v}, err
	case *ParenExpr:
		return EvalExpr(node.X, env)
	case *SelectorExpr:
		v, err := evalSelectorExpr(node, env)
		return []reflect.Value{v}, err
	case *IndexExpr:
		return evalIndexExpr(node, env)
	case *SliceExpr:
		v, err := evalSliceExpr(node, env)
		return []reflect.Value{v}, err
	case *TypeAssertExpr:
		return evalTypeAssertExpr(node, env)
	case *CallExpr:
		return evalCallExpr(node, env)
	case *StarExpr:
		v, err := evalStarExpr(node, env)
		return []reflect.Value{v}, err
	case *UnaryExpr:
		return evalUnaryExpr(node, env)
	case *BinaryExpr:
		v, err := evalBinaryExpr(node, env)
		return []reflect.Value{v}, err
	default:
		panic(dytc("unevalutable node"))
	}
}
