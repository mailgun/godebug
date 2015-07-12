package eval

import (
	"errors"
	"reflect"

	"go/ast"
	"go/token"
)

type checkCtx struct {
	outerFunc reflect.Type
	emptyReturnOk bool
	stack []Stmt
}

// Place holder for something more substantial
func CheckStmt(stmt ast.Stmt, env Env) (Stmt, []error) {
	// Create a dummy env where variables can be added without affecting the global env
	return checkStmt(stmt, env.PushScope(), checkCtx{})
}

func checkBlock(block *ast.BlockStmt, env Env, ctx checkCtx) (*BlockStmt, []error) {
	var errs, moreErrs []error
	if block == nil {
		return nil, nil
	}
	ablock := &BlockStmt{BlockStmt: block}
	if block.List != nil {
		ablock.List = make([]Stmt, len(block.List))
		for i, stmt := range block.List {
			ablock.List[i], moreErrs = checkStmt(stmt, env, ctx)
			errs = append(errs, moreErrs...)
		}
	}
	return ablock, errs
}

func checkStmt(stmt ast.Stmt, env Env, ctx checkCtx) (Stmt, []error) {
	var errs, moreErrs []error
	ctx.stack = append(ctx.stack, stmt)
	switch s := stmt.(type) {
	case nil:
		// AST often has nil nodes for optional elements.
		return nil, nil
	case *ast.AssignStmt:
		a := &AssignStmt{
			AssignStmt: s,
			Lhs: make([]Expr, len(s.Lhs)),
			Rhs: make([]Expr, len(s.Rhs)),
		}

		names := map[int]string{}
		// Identify names
		if s.Tok == token.DEFINE {
			newName := false
			for i, l := range s.Lhs {
				if ident, ok := l.(*ast.Ident); ok {
					if ident.Name != "_" && !inTopEnv(ident.Name, env) {
						newName = true
						names[i] = ident.Name
					}
				} else {
					errs = append(errs, ErrNonNameInDeclaration{fakeCheckExpr(l, env)})
				}
			}
			if !newName {
				errs = append(errs, ErrNoNewNamesInDeclaration{a})
			}
		} else if s.Tok != token.ASSIGN {
			// Could probably check and return here as an optimisation, but duplicates some logic
			binary := &ast.BinaryExpr{X: s.Lhs[0], OpPos: s.TokPos, Op: s.Tok, Y: s.Rhs[0]}
			s.Rhs[0] = binary
		}

		// Check lhs
		lhsChecked := true
		for i, l := range s.Lhs {
			if isBlankIdentifier(l) {
				names[i] = "_"
				a.Lhs[i] = fakeCheckExpr(l, env)
				continue
			} else if _, ok := names[i]; ok {
				a.Lhs[i] = fakeCheckExpr(l, env)
				continue
			}
			a.Lhs[i], moreErrs = CheckExpr(l, env)
			if moreErrs != nil && !a.Lhs[i].IsConst() {
				lhsChecked = false
				errs = append(errs, moreErrs...)
				continue
			}
			// Must be addressable or map index expr
			ll := skipSuperfluousParens(a.Lhs[i])
			if index, ok := ll.(*IndexExpr); ok {
				k := index.X.KnownType()[0].Kind()
				if k == reflect.Map || k == reflect.Slice {
					continue
				}
			} else {
				if _, err := expectSingleType(a.Lhs[i]); err != nil {
					errs = append(errs, err)
				}
			}
			if !isAddressable(ll) {
				errs = append(errs, ErrCannotAssignToUnaddressable{a.Lhs[i]})
			}
		}

		isMulti := false
		var types []reflect.Type
		if len(a.Rhs) == 1 {
			a.Rhs[0], moreErrs = CheckExpr(s.Rhs[0], env)
			errs = append(errs, moreErrs...)
			if errs != nil && !a.Rhs[0].IsConst() {
				goto done
			}
			types = make([]reflect.Type, len(a.Rhs[0].KnownType()))
			copy(types, a.Rhs[0].KnownType())
			if len(a.Lhs) == 2 && len(a.Rhs) == 1 && multivalueOk(a.Rhs[0]) {
				types = append(types, boolType)
			}
			isMulti = len(a.Lhs) > 1
		} else {
			types = make([]reflect.Type, len(a.Rhs))
			for i, r := range s.Rhs {
				a.Rhs[i], moreErrs = CheckExpr(r, env)
				errs = append(errs, moreErrs...)
				if moreErrs != nil && !a.Rhs[i].IsConst() {
					continue
				}
				if t, err := expectSingleType(a.Rhs[i]); err != nil {
					errs = append(errs, err)
				} else {
					types[i] = t
				}
			}
		}

		// Check rhs
		if len(a.Lhs) != len(types) {
			errs = append(errs, ErrAssignCountMismatch{a, len(a.Lhs), len(types)})
			goto done
		}
		// Check for assignability
		if !lhsChecked {
			goto done
		}

		for i := range a.Rhs {
			if types[i] == nil {
				// new variable or typecheck failed
				continue
			}
			kt := a.Lhs[i].KnownType()
			assignable := true
			if kt == nil {
				// _ or new name
				if ct, ok := types[i].(ConstType); ok {
					if ct == ConstNil {
						errs = append(errs, ErrUntypedNil{a.Rhs[i]})
						continue
					} else {
						types[i] = ct.DefaultPromotion()
					}
				}
			} else if _, ok := kt[0].(ConstType); ok {
				// Corner case for assigning to const basic lits. e.g. 1 = 1
				_, assignable = names[i]
			} else {
				// expect the left type
				types[i] = kt[0]
				var convErrs []error
				assignable, convErrs = exprAssignableTo(a.Rhs[i], types[i])
				if assignable {
					errs = append(errs, convErrs...)
				}
			}
			if !assignable {
				if isMulti {
					errs = append(errs, ErrCannotAssignToType{a.Lhs[i], a.Rhs[0], i})
				} else {
					errs = append(errs, ErrCannotAssignToType{a.Lhs[i], a.Rhs[i], -1})
				}
			}
		}

		for i, name := range names {
			if name != "_" {
				env.AddVar(name, reflect.New(types[i]))
			}
		}
done:
		a.newNames = names
		a.types = types
		return a, errs

	case *ast.BranchStmt:
		astmt := &BranchStmt{BranchStmt: s}
		if s.Label != nil {
			astmt.Label = &Ident{Ident: s.Label}
		}
		return astmt, nil

	case *ast.BlockStmt:
		return checkBlock(s, env, ctx)

	case *ast.EmptyStmt:
		return &EmptyStmt{EmptyStmt: s}, nil

	case *ast.ExprStmt:
		x, errs := CheckExpr(s.X, env)
		return &ExprStmt{ExprStmt: s, X: x}, errs

	case *ast.IfStmt:
		astmt := &IfStmt{IfStmt: s}
		env = env.PushScope() // Env for the if block

		astmt.Init, moreErrs = checkStmt(s.Init, env, ctx)
		errs = append(errs, moreErrs...)

		astmt.Cond, moreErrs = checkCond(s.Cond, astmt, env, ctx)
		errs = append(errs, moreErrs...)

		astmt.Body, moreErrs = checkBlock(s.Body, env, ctx)
		errs = append(errs, moreErrs...)

		astmt.Else, moreErrs = checkStmt(s.Else, env, ctx)
		errs = append(errs, moreErrs...)
		return astmt, errs

	case *ast.IncDecStmt:
		var tok token.Token
		if s.Tok == token.INC {
			tok = token.ADD_ASSIGN
		} else {
			tok = token.SUB_ASSIGN
		}

		// one.End() is after the ++/--
		one := &ast.BasicLit{
			ValuePos: s.TokPos + 1,  // Pos of second +/- in ++/--
			Kind: token.INT,
			Value: "1",
		}
		assign := &ast.AssignStmt{
			Lhs: []ast.Expr{s.X},
			TokPos: s.TokPos,
			Tok: tok,
			Rhs: []ast.Expr{one},
		}
		return checkStmt(assign, env, ctx)

	case *ast.LabeledStmt:
		astmt := &LabeledStmt{LabeledStmt: s}
		astmt.Label = &Ident{Ident: s.Label}
		astmt.Stmt, moreErrs = checkStmt(s.Stmt, env, ctx)
		switch loop := astmt.Stmt.(type) {
		case *ForStmt:
			loop.label = astmt.Label.Name
		}
		return astmt, errs

	case *ast.ForStmt:
		astmt := &ForStmt{ForStmt: s}
		env = env.PushScope() // Env for the for block

		astmt.Init, moreErrs = checkStmt(s.Init, env, ctx)
		errs = append(errs, moreErrs...)

		astmt.Cond, moreErrs = checkCond(s.Cond, astmt, env, ctx)
		errs = append(errs, moreErrs...)

		astmt.Post, moreErrs = checkStmt(s.Post, env, ctx)
		errs = append(errs, moreErrs...)

		astmt.Body, moreErrs = checkBlock(s.Body, env, ctx)
		errs = append(errs, moreErrs...)
		return astmt, errs

	case *ast.ReturnStmt:
		astmt := &ReturnStmt{ReturnStmt: s}
		if s.Results != nil {
			astmt.Results = make([]Expr, len(s.Results))
		}
		var i, numOut, numResults int
		var ok bool
		if ctx.outerFunc != nil {
			numOut = ctx.outerFunc.NumOut()
		}
		if len(s.Results) == 1 {
			r, _ := CheckExpr(s.Results[0], env)
			if kt := r.KnownType(); len(kt) > 1 {
				numResults = len(kt)
				astmt.Results[0] = r
				for i = 0; i < numOut && i < numResults ; i += 1 {
					t := ctx.outerFunc.Out(i)
					if !typeAssignableTo(kt[i], ctx.outerFunc.Out(1)) {
						errs = append(errs, ErrBadReturnValue{astmt.Results[0], t, i})
					}
				}
				goto checkcount
			}
		}
		numResults = len(s.Results)
		for i = 0; i < numResults && i < numOut; i += 1 {
			t := ctx.outerFunc.Out(i)
			astmt.Results[i], ok, moreErrs = checkExprAssignableTo(s.Results[i], t, env)
			if ok {
				errs = append(errs, moreErrs...)
			} else {
				errs = append(errs, ErrBadReturnValue{astmt.Results[i], t, -1})
			}
		}
		for ; i < numResults; i += 1 {
			astmt.Results[i], moreErrs = CheckExpr(s.Results[i], env)
			errs = append(errs, moreErrs...)
		}
checkcount:
		if ctx.outerFunc != nil && numResults != numOut && !(numOut == 0 && ctx.emptyReturnOk) {
			errs = append(errs, ErrWrongNumberOfReturnValues{astmt, ctx.outerFunc})
		}
		return astmt, errs

	case *ast.SwitchStmt:
		body := &BlockStmt{BlockStmt: s.Body, List: make([]Stmt, len(s.Body.List))}
		astmt := &SwitchStmt{SwitchStmt: s, Body: body}
		env = env.PushScope() // Env for the switch

		astmt.Init, moreErrs = checkStmt(s.Init, env, ctx)
		errs = append(errs, moreErrs...)

		tag := s.Tag
		if tag == nil {
			tag = &ast.Ident{Name: "true", NamePos: s.Body.Lbrace - 1}
		}

		astmt.Tag, moreErrs = CheckExpr(tag, env)
		errs = append(errs, moreErrs...)

		if moreErrs == nil || astmt.Tag.IsConst() {
			if t, err := expectSingleType(astmt.Tag); err != nil {
				errs = append(errs, err)
			} else if ct, ok := t.(ConstType); ok {
				if ct == ConstNil {
					errs = append(errs, ErrUntypedNil{astmt.Tag})
				} else {
					astmt.tagT = ct.DefaultPromotion()
				}
			} else {
				astmt.tagT = t
			}
		}

		var ok bool
		t := astmt.tagT
		for i, stmt := range s.Body.List {
			clause := stmt.(*ast.CaseClause)
			aclause := &CaseClause{CaseClause: clause}
			if clause.List == nil {
				astmt.def = aclause
			} else {
				aclause.List = make([]Expr, len(clause.List))
			}
			for j, expr := range clause.List {
				if t != nil {
					aclause.List[j], ok, moreErrs = checkExprAssignableTo(expr, t, env)
					errs = append(errs, moreErrs...)
					if !ok {
						errs = append(errs, ErrInvalidCase{aclause.List[j], astmt.Tag})
					}
				} else {
					aclause.List[j], moreErrs = CheckExpr(expr, env)
					errs = append(errs, moreErrs...)
				}
			}
			astmt.Body.List[i], moreErrs = checkCaseClauseBody(aclause, env, ctx)
		}
		return astmt, errs

	case *ast.TypeSwitchStmt:
		body := &BlockStmt{BlockStmt: s.Body, List: make([]Stmt, len(s.Body.List))}
		astmt := &TypeSwitchStmt{TypeSwitchStmt: s, Body: body}
		env = env.PushScope() // Env for the switch

		astmt.Init, moreErrs = checkStmt(s.Init, env, ctx)
		errs = append(errs, moreErrs...)

		// Env for the case clause
		caseEnv := env.PushScope()

		var t reflect.Type
		var name string
		var tag Expr

		if assign, ok := s.Assign.(*ast.AssignStmt); ok {
			assign.Rhs[0] = assign.Rhs[0].(*ast.TypeAssertExpr).X
		} else if exprstmt, ok := s.Assign.(*ast.ExprStmt); ok {
			exprstmt.X = exprstmt.X.(*ast.TypeAssertExpr).X
		} else {
			panic("TypeSwitchStmt.Assign is not (Assign|Expr)Stmt ")
		}

		astmt.Assign, moreErrs = checkStmt(s.Assign, caseEnv, ctx)
		errs = append(errs, moreErrs...)

		if assign, ok := astmt.Assign.(*AssignStmt); ok {
			name = assign.Lhs[0].(*Ident).Name
			tag = assign.Rhs[0]

			if moreErrs == nil || tag.IsConst() {
				t = caseEnv.Var(name).Elem().Type()
			}
		} else if exprstmt, ok := astmt.Assign.(*ExprStmt); ok {
			tag = exprstmt.X

			if moreErrs == nil || tag.IsConst() {
				var err error
				if t, err = expectSingleType(tag); err != nil {
					errs = append(errs, err)
				}
			}
		} else {
			panic("TypeSwitchStmt.Assign is not (Assign|Expr)Stmt ")
		}

		if t != nil && (t == ConstNil || t.Kind() != reflect.Interface) {
			errs = append(errs, ErrNonInterfaceTypeSwitch{tag})
			t = nil
		}

		for i, stmt := range s.Body.List {
			caseEnv := env.PushScope()
			clause := stmt.(*ast.CaseClause)
			aclause := &CaseClause{CaseClause: clause}
			if clause.List == nil {
				astmt.def = aclause
			} else {
				aclause.List = make([]Expr, len(clause.List))
			}
			for j, expr := range clause.List {
				aexpr, tt, isType, moreErrs := checkType(expr, env)
				errs = append(errs, moreErrs...)
				if !isType {
					aexpr, moreErrs = CheckExpr(expr, env)
					errs = append(errs, moreErrs...)
					if moreErrs == nil || aexpr.IsConst() {
						errs = append(errs, ErrBuiltinNonTypeArg{aexpr})
					}
				// isType == true && tt == nil for unimplemented types
				} else if t != nil && tt != nil {
					if tt.Kind() != reflect.Interface && !tt.Implements(t) {
						errs = append(errs, ErrImpossibleTypeCase{aexpr, tag})
					} else if len(clause.List) == 1 {
						caseEnv.AddVar(name, reflect.New(tt))
					}
				}
				aclause.List[j] = aexpr
			}
			astmt.Body.List[i], moreErrs = checkCaseClauseBody(aclause, caseEnv, ctx)
			errs = append(errs, moreErrs...)
		}
		return astmt, errs

	default:
		return nil, []error{errors.New("Only assign statements are currently supported")}
	}
}

func checkCond(cond ast.Expr, parent Stmt, env Env, ctx checkCtx) (Expr, []error) {
	if cond == nil {
		return nil, nil
	}
	acond, errs := CheckExpr(cond, env)
	if errs == nil || acond.IsConst() {
		if t, err := expectSingleType(acond); err != nil {
			errs = append(errs, err)
		} else if t.Kind() != reflect.Bool {
			errs = append(errs, ErrNonBoolCondition{acond, parent})
		}
	}
	return acond, errs
}

func checkCaseClauseBody(clause *CaseClause, env Env, ctx checkCtx) (*CaseClause, []error) {
	var errs, moreErrs []error
	if clause.CaseClause.Body != nil {
		clause.Body = make([]Stmt, len(clause.CaseClause.Body))
	}
	for i, stmt := range clause.CaseClause.Body {
		clause.Body[i], moreErrs = checkStmt(stmt, env, ctx)
		errs = append(errs, moreErrs...)
	}
	return clause, errs
}

// Find labeled stmt referenced by branch. Return jump, which contains
// the stack of statements from the root of the checkCtx to the
// found label, or nil if not found. Statements common to both
// ctx.stack and the found label are filtered.
func findLabel(branch *BranchStmt, ctx checkCtx) (jump []Stmt) {
	target := branch.Label.Name
	// break/continue must refer to a labelled for/range stmt in the stack
	if branch.Tok == token.BREAK || branch.Tok == token.CONTINUE {
		for i := len(ctx.stack) - 1; i >= 0; i -= 1 {
			if l, ok := ctx.stack[i].(*LabeledStmt); ok && l.Label.Name == target {
				return []Stmt{l}
			}
		}
		return nil
	}

	// depth first search starting from the branch stmt, backing up the stack
	closed := map[Stmt]bool{}
	open := append([]Stmt(nil), ctx.stack...)
	s := len(open) - 1
	for s >= 0 {
		top := open[s]
		s -= 1
		if closed[top] {
			continue
		}
		closed[top] = true
		open = open[:s+2]

		switch s := top.(type) {
		case *BlockStmt:
			open = append(open, s.List...)
		case *CaseClause:
			open = append(open, s.Body...)
		case *IfStmt:
			if s.Else != nil {
				open = append(open, s.Else)
			}
			open = append(open, s.Body)
		case *LabeledStmt:
			if s.Label.Name == target {
				// Currently ctx.stack is a trace to the goto stmt,
				// and open is a trace to the target label.
				// In addition to parent nodes, open also contains
				// unexpanded sibling nodes.
				//
				// Compute the portion of open that is not common
				// to ctx.stack, which is used to detect skipped
				// declarations and jumps into foreign blocks.
				//
				// We always return the LabeledStmt as part of the
				// jump. Iterating to len(open)-1 removes a corner
				// case where the labeled stmt is part of the ctx.stack
				// trace and is hence removed.
				common := 0
				for j, k := 0, 0; j < len(open)-1 && k < len(ctx.stack); j += 1 {
					if open[j] == ctx.stack[k] {
						common = j+1
						k += 1
					}
				}
				return open[common:]
			}
			if s.Stmt != nil {
				open = append(open, s.Stmt)
			}
		case *ForStmt:
			open = append(open, s.Body)
		case *SwitchStmt:
			open = append(open, s.Body)
		case *TypeSwitchStmt:
			open = append(open, s.Body)
		}
		s = len(open) - 1
	}
	return nil
}

