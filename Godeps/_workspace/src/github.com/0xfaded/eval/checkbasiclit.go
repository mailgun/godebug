package eval

import (
	"strconv"

	"go/ast"
	"go/token"
)

func checkBasicLit(lit *ast.BasicLit, env Env) (*BasicLit, []error) {
	aexpr := &BasicLit{BasicLit: lit}

	switch lit.Kind {
	case token.CHAR:
		if r, _, tail, err := strconv.UnquoteChar(lit.Value[1:len(lit.Value)-1], '\''); err != nil {
			return aexpr, []error{ErrBadBasicLit{aexpr}}
		} else if tail != "" {
			// parser.ParseExpr() should raise a syntax error before we get here.
			panic("go-interactive: bad char lit " + lit.Value)
		} else {
			aexpr.constValue = constValueOf(NewConstRune(r))
			aexpr.knownType = knownType{ConstRune}
			return aexpr, nil
		}
	case token.STRING:
		if str, err := strconv.Unquote(string(lit.Value)); err != nil {
			return aexpr, []error{ErrBadBasicLit{aexpr}}
		} else {
			aexpr.constValue = constValueOf(str)
			aexpr.knownType = knownType{ConstString}
			return aexpr, nil
		}
	case token.INT:
		if i, ok := NewConstInteger(lit.Value); !ok {
			return aexpr, []error{ErrBadBasicLit{aexpr}}
		} else {
			aexpr.constValue = constValueOf(i)
			aexpr.knownType = knownType{ConstInt}
			return aexpr, nil
		}
	case token.FLOAT:
		if f, ok := NewConstFloat(lit.Value); !ok {
			return aexpr, []error{ErrBadBasicLit{aexpr}}
		} else {
			aexpr.constValue = constValueOf(f)
			aexpr.knownType = knownType{ConstFloat}
			return aexpr, nil
		}
	case token.IMAG:
		if i, ok := NewConstImag(lit.Value); !ok {
			return aexpr, []error{ErrBadBasicLit{aexpr}}
		} else {
			aexpr.constValue = constValueOf(i)
			aexpr.knownType = knownType{ConstComplex}
			return aexpr, nil
		}
	default:
		return aexpr, []error{ErrBadBasicLit{aexpr}}
	}
}
