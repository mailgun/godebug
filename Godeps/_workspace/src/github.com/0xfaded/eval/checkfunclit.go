package eval

import (
	"go/ast"

	"reflect"
)

func checkFuncLit(lit *ast.FuncLit, env Env) (*FuncLit, []error) {
	alit := &FuncLit{FuncLit: lit}
	atype, t, _, errs := checkType(lit.Type, env)
	alit.Type = atype.(*FuncType)
	if t != nil {
		alit.knownType = knownType{t}
	}

	// Create a new environment containing all valid params and results.
	seen := map[string]bool{}
	invalid := map[string]bool{}
	env = env.PushScope()
	fields := alit.Type.Params
	// Type.Results may be nil
	for i := 0; i < 2 && fields != nil; i += 1 {
		for _, field := range fields.List {
			var z reflect.Value
			if len(field.KnownType()) != 0 {
				z = reflect.New(field.KnownType()[0])
			}
			for _, name := range field.Names {
				if name.Name == "_" {
					continue
				}
				if !z.IsValid() {
					invalid[name.Name] = true
				} else {
					env.AddVar(name.Name, z)
				}
				if seen[name.Name] {
					errs = append(errs, ErrDuplicateArg{name})
				}
				seen[name.Name] = true
			}
		}
		fields = alit.Type.Results
	}

	// t may be nil, in this case return statements won't be checked
	block, moreErrs := checkBlock(lit.Body, env, checkCtx{outerFunc: t})
	alit.Body = block

	// Filter out undefined errors caused by invalid params
	var filtered []error
	for _, err := range moreErrs {
		if undef, ok := err.(ErrUndefined); ok {
			if ident, ok := undef.Expr.(*Ident); ok && invalid[ident.Name] {
				continue
			}
		}
		filtered = append(filtered, err)
	}
	return alit, append(errs, filtered...)
}
