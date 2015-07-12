package eval

import (
	"errors"
	"reflect"
	"go/ast"
)

// Type check an ast.Expr to produce an Expr. Errors are accumulated and
// returned as a single slice. When evaluating constant expressions,
// non fatal truncation/overflow errors may be raised but type checking
// will continue. A common pattern to detect errors is
//
//  if expr, errs := CheckExpr(...); errs != nil && !expr.IsConst() {
//    fatal
//  }
//
// if expr.IsConst() is true, then the resulting Expr has been successfully
// checked, regardless of if errors are present.
func CheckExpr(expr ast.Expr, env Env) (Expr, []error) {
	if t, _, isType, _ := checkType(expr, env); isType {
		return t, []error{ErrTypeUsedAsExpression{t}}
	}

	switch expr := expr.(type) {
	case *ast.BadExpr:
		return &BadExpr{BadExpr: expr}, nil
	case *ast.Ident:
		return checkIdent(expr, env)
	case *ast.Ellipsis:
		return &Ellipsis{Ellipsis: expr}, nil
	case *ast.BasicLit:
		return checkBasicLit(expr, env)
	case *ast.FuncLit:
		return checkFuncLit(expr, env)
	case *ast.CompositeLit:
		return checkCompositeLit(expr, env)
	case *ast.ParenExpr:
		return checkParenExpr(expr, env)
	case *ast.SelectorExpr:
		return checkSelectorExpr(expr, env)
	case *ast.IndexExpr:
		return checkIndexExpr(expr, env)
	case *ast.SliceExpr:
		return checkSliceExpr(expr, env)
	case *ast.TypeAssertExpr:
		return checkTypeAssertExpr(expr, env)
	case *ast.CallExpr:
		return checkCallExpr(expr, env)
	case *ast.StarExpr:
		return checkStarExpr(expr, env)
	case *ast.UnaryExpr:
		return checkUnaryExpr(expr, env)
	case *ast.BinaryExpr:
		return checkBinaryExpr(expr, env)
	case *ast.KeyValueExpr:
		panic("eval: KeyValueExpr checked")
	default:
		panic("eval: Bad expr")
	}
}

func checkType(expr ast.Expr, env Env) (Expr, reflect.Type, bool, []error) {
	for parens, ok := expr.(*ast.ParenExpr); ok; parens, ok = expr.(*ast.ParenExpr) {
		expr = parens.X
	}
	switch node := expr.(type) {
	case *ast.Ident:
		ident := &Ident{Ident: node}
		if t := env.Type(node.Name); t != nil {
			ident.knownType = knownType{t}
			return ident, t, true, nil
		} else if t, ok := builtinTypes[node.Name]; ok {
			ident.knownType = knownType{t}
			return ident, t, true, nil
		} else {
			return ident, nil, false, []error{ErrUndefined{ident}}
		}
	case *ast.StarExpr:
		star := &StarExpr{StarExpr: node}
		elem, elemT, isType, errs := checkType(node.X, env)
		if isType {
			// Only set X if X is a type, as * can be part of an expression or type
			t := reflect.PtrTo(elemT)
			star.X = elem
			star.knownType = knownType{t}
			return star, t, isType, nil
		} else {
			return star, nil, isType, errs
		}
	case *ast.SelectorExpr:
		// TODO[crc] remove this comment after ast cleanup
		// Note that Sel SelectorExpr has its own Sel field, shadowing ast.SelectorExpr.Sel
		// This is on the cleanup list.
		sel := &SelectorExpr{SelectorExpr: node, Sel: &Ident{Ident: node.Sel}}
		if ident, ok := node.X.(*ast.Ident); !ok {
			return sel, nil, false, nil
		} else if pkg := env.Pkg(ident.Name); pkg == nil {
			return sel, nil, false, nil
		} else if t := pkg.Type(sel.Sel.Name); t == nil {
			return sel, nil, false, nil
		} else {
			// Only set X if the selector is a type, . can be part of an expression or type
			sel.X = &Ident{Ident: ident}
			sel.knownType = knownType{t}
			return sel, t, true, nil
		}
	case *ast.ArrayType:
		arrayT := &ArrayType{ArrayType: node}
		if node.Len != nil {
			return arrayT, nil, true, []error{errors.New("array types not implemented")}
		} else {
			elt, eltT, _, errs := checkType(node.Elt, env);
			arrayT.Elt = elt
			if errs != nil {
				return arrayT, nil, true, errs
			} else {
				t := reflect.SliceOf(unhackType(eltT))
				arrayT.knownType = knownType{t}
				return arrayT, t, true, nil
			}
		}
	case *ast.StructType:
		structT := &StructType{StructType: node}
		return structT, nil, true, []error{errors.New("struct types not implemented")}
	case *ast.FuncType:
		funcT := &FuncType{FuncType: node}
		return funcT, nil, true, []error{errors.New("func types not implemented")}
	case *ast.InterfaceType:
		interfaceT := &InterfaceType{InterfaceType: node}
		// Allow interface{}'s
		if node.Methods.List == nil {
			interfaceT.knownType = knownType{emptyInterface}
			return interfaceT, emptyInterface, true, nil
		}
		return interfaceT, nil, true, []error{errors.New("interface types not implemented")}
	case *ast.MapType:
		mapT := &MapType{MapType: node}
		keyT, k, _, errs := checkType(node.Key, env)
		mapT.Key = keyT
		if k != nil && !isStaticTypeComparable(k) {
			errs = append(errs, ErrUncomparableMapKey{mapT, k})
		}
		valueT, v, _, moreErrs := checkType(node.Value, env)
		mapT.Value = valueT
		if moreErrs != nil {
			errs = append(errs, moreErrs...)
		}
		if errs == nil {
			t := reflect.MapOf(unhackType(k), unhackType(v))
			mapT.knownType = knownType{t}
			return mapT, t, true, nil
		}
		return mapT, nil, true, errs
	case *ast.ChanType:
		chanT := &ChanType{ChanType: node}
		value, valueT, _, errs := checkType(node.Value, env);
		chanT.Value = value
		if errs != nil {
			return chanT, nil, true, errs
		} else {
			if node.Dir == ast.SEND {
				chanT.dir = reflect.SendDir
			} else if node.Dir == ast.RECV {
				chanT.dir = reflect.RecvDir
			} else {
				chanT.dir = reflect.BothDir
			}
			t := reflect.ChanOf(chanT.dir, unhackType(valueT))
			chanT.knownType = knownType{t}
			return chanT, t, true, nil
		}
	}
	return nil, nil, false, nil
}

func checkFieldList(list *ast.FieldList, env Env) (*FieldList, []error) {
	if list == nil {
		return nil, nil
	}
	var errs, moreErrs []error

	alist := &FieldList{FieldList: list}
	if list.List == nil {
		return alist, nil
	}
	alist.List = make([]*Field, len(list.List))
	for i, field := range list.List {
		alist.List [i], moreErrs = checkField(field, env)
		errs = append(errs, moreErrs...)
	}
	return alist, nil
}

func checkField(field *ast.Field, env Env) (*Field, []error) {
	afield := &Field{Field: field}
	if field.Names != nil {
		afield.Names = make([]*Ident, len(field.Names))
		for i, ident := range field.Names {
			afield.Names[i] = &Ident{Ident: ident}
		}
	}
	// the ellipsis is only relevant for func args
	var typ Expr
	var t reflect.Type
	var errs []error
	if ellipsis, ok := field.Type.(*ast.Ellipsis); ok {
		typ, t, _, errs = checkType(ellipsis.Elt, env)
		if t != nil {
			t = reflect.SliceOf(t)
		}
		ellipsis.Elt = typ
		typ = &Ellipsis{Ellipsis: ellipsis}
	} else {
		typ, t, _, errs = checkType(field.Type, env)
	}
	afield.Type = typ
	if t != nil {
		afield.knownType = knownType{t}
	}
	return afield, errs
}
