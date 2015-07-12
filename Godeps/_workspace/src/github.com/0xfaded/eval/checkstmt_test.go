package eval

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestGotoBackward(t *testing.T) {
	_, ctx:= checkString(`{
target:
	_ = 1+1
	goto target
}`)
	jump := findLabel(gothere("target"), ctx)
	if len(jump) != 1 {
		t.Errorf("Wrong number of jump stmts\n%#v", jump)
	}
}

func TestGotoForward(t *testing.T) {
	_, ctx:= checkString(`{
	goto target
	_ = 1+1
target:
}`)
	jump := findLabel(gothere("target"), ctx)
	if len(jump) != 2 {
		t.Errorf("Wrong number of jump stmts\n%#v", jump)
	}
}

func TestGotoEscapeBlockBackward(t *testing.T) {
	_, ctx:= checkString(`{
target:
	_ = 1+1
	{
		goto target
	}
}`)
	jump := findLabel(gothere("target"), ctx)
	if len(jump) != 1 {
		t.Errorf("Wrong number of jump stmts\n%#v", jump)
	}
}

func TestGotoEscapeBlockForwards(t *testing.T) {
	_, ctx:= checkString(`{
	{
		goto target
	}
	_ = 1+1
target:
}`)
	jump := findLabel(gothere("target"), ctx)
	if len(jump) != 2 {
		t.Errorf("Wrong number of jump stmts\n%#v", jump)
	}
}

func TestGotoIntoBlockBackwards(t *testing.T) {
	_, ctx:= checkString(`{
	{
		target:
	}
	_ = 1+1
	goto target
}`)
	jump := findLabel(gothere("target"), ctx)
	if len(jump) != 2 {
		t.Errorf("Wrong number of jump stmts\n%#v", jump)
	}
}

func TestGotoIntoBlockForward(t *testing.T) {
	_, ctx:= checkString(`{
	goto target
	_ = 1+1
	{
		target:
	}
}`)
	jump := findLabel(gothere("target"), ctx)
	if len(jump) != 3 {
		t.Errorf("Wrong number of jump stmts\n%#v", jump)
	}
}

func TestGotoBetweenBlocksBackwards(t *testing.T) {
	_, ctx:= checkString(`{
	{
		target:
	}
	_ = 1+1
	{
		goto target
	}
}`)
	jump := findLabel(gothere("target"), ctx)
	if len(jump) != 2 {
		t.Errorf("Wrong number of jump stmts\n%#v", jump)
	}
}

func TestGotoBetweenBlocksForward(t *testing.T) {
	_, ctx:= checkString(`{
	{
		goto target
	}
	_ = 1+1
	{
		target:
	}
}`)
	jump := findLabel(gothere("target"), ctx)
	if len(jump) != 3 {
		t.Errorf("Wrong number of jump stmts\n%#v", jump)
	}
}

func checkString(stmt string) (Stmt, checkCtx) {
	env := MakeSimpleEnv()
	s, _ := ParseStmt(stmt)
	c, _ := CheckStmt(s, env)
	ctx := checkCtx{stack: findGoto(c, []Stmt{})}
	return c, ctx
}

func findGoto(stmt Stmt, stack []Stmt) []Stmt {
	stack = append(stack, stmt)
	switch s := stmt.(type) {
	case *BlockStmt:
		for _, x := range s.List {
			if found := findGoto(x, stack); found != nil {
				return found
			}
		}
	case *CaseClause:
		for _, x := range s.Body {
			if found := findGoto(x, stack); found != nil {
				return found
			}
		}
	case *IfStmt:
		if found := findGoto(s.Body, stack); found != nil {
			return found
		}
		if s.Else != nil {
			return findGoto(s.Else, stack)
		}
	case *BranchStmt:
		if s.Tok == token.GOTO {
			return stack
		}
	case *LabeledStmt:
		return findGoto(s.Stmt, stack)
	case *ForStmt:
		return findGoto(s.Body, stack)
	case *SwitchStmt:
		return findGoto(s.Body, stack)
	case *TypeSwitchStmt:
		return findGoto(s.Body, stack)
	}
	return nil
}

func gothere(label string) *BranchStmt {
	return &BranchStmt{
		BranchStmt: &ast.BranchStmt{
			Tok: token.GOTO,
		},
		Label: &Ident{
			Ident: &ast.Ident{
				Name: label,
			},
		},
	}
}
