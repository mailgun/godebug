package eval

import (
	"fmt"
	"reflect"
	"go/token"
)

type State struct {
	Last Stmt
	Env Env
}

func InterpStmt(stmt Stmt, env Env) (last *State, err error) {
	switch s := stmt.(type) {
	case nil:
	case *AssignStmt:
		if len(s.Rhs) == 1 {
			rs, err := evalTypedExpr(s.Rhs[0], s.types, env)
			if err != nil {
				if _, ok := err.(PanicInterfaceConversion); !ok || len(s.types) != 2 {
					return nil, err
				}
			}
			for i, lhs := range s.Lhs {
				if name, ok := s.newNames[i]; !ok {
					assign(lhs, rs[i], env)
				} else if name != "_" {
					v := hackedNew(s.types[i])
					v.Elem().Set(rs[i])
					env.AddVar(name, v)
				}
			}
		} else {
			for i, lhs := range s.Lhs {
				r, err := evalTypedExpr(s.Rhs[i], s.types[i:i+1], env)
				if err != nil {
					return nil, err
				}
				if name, ok := s.newNames[i]; !ok {
					assign(lhs, r[0], env)
				} else if name != "_" {
					v := hackedNew(s.types[i])
					v.Elem().Set(r[0])
					env.AddVar(name, v)
				}
			}
		}
	case *BranchStmt:
		return &State{s, env}, nil
	case *BlockStmt:
		return interpBlock(s.List, env)
	case *CaseClause:
		return interpBlock(s.Body, env)
	case *EmptyStmt:
		return nil, nil
	case *ExprStmt:
		_, err := EvalExpr(s.X, env)
		return nil, err
	case *IfStmt:
		env = env.PushScope()
		if _ , err = InterpStmt(s.Init, env); err != nil {
			return nil, err
		} else if rs, err := EvalExpr(s.Cond, env); err != nil {
			return nil, err
		} else if rs[0].Bool() {
			return InterpStmt(s.Body, env)
		} else {
			return InterpStmt(s.Else, env)
		}
	case *ForStmt:
		env = env.PushScope()
		if _, err = InterpStmt(s.Init, env); err != nil {
			return nil, err
		}
		for {
			if s.Cond != nil {
				if rs, err := EvalExpr(s.Cond, env); err != nil {
					return nil, err
				} else if !rs[0].Bool() {
					break
				}
			}
			if last, err = InterpStmt(s.Body, env); err != nil {
				return last, err
			}
			if last != nil {
				if branch, ok := last.Last.(*BranchStmt); ok {
					// Are we the target of this branch?
					if branch.Label == nil || branch.Label.Name == s.label {
						last = nil
						if branch.Tok == token.CONTINUE {
							goto cont
						}
					}
				}
				return last, nil
			}
cont:
			if _, err = InterpStmt(s.Post, env); err != nil {
				return nil, err
			}
		}
	case *LabeledStmt:
		return InterpStmt(s.Stmt, env)
	case *ReturnStmt:
		return &State{s, env}, nil
	case *SwitchStmt:
		env = env.PushScope()
		t := knownType{s.tagT}
		if _, err := InterpStmt(s.Init, env); err != nil {
			return nil, err
		}
		tag, err := evalTypedExpr(s.Tag, t, env)
		if err != nil {
			return nil, err
		}
		env = env.PushScope()
		for _, stmt := range s.Body.List {
			clause := stmt.(*CaseClause)
			for _, expr := range clause.List {
				if sw, err := evalTypedExpr(expr, t, env); err != nil {
					return nil, err
				} else if eq, err := equal(tag[0], sw[0]); err != nil {
					return nil, err
				} else if eq {
					return InterpStmt(clause, env)
				}
			}
		}
		return InterpStmt(s.def, env)

	case *TypeSwitchStmt:
		env = env.PushScope()
		if _, err = InterpStmt(s.Init, env); err != nil {
			return nil, err
		}

		x, err := EvalExpr(s.Tag(), env)
		if err != nil {
			return nil, err
		}
		// interface.elem()
		dynamicX := x[0].Elem()
		dynamicT := dynamicX.Type()

		env = env.PushScope()
		if name := s.Name(); name != "" {
			// dynamicX may not be addressable
			x := reflect.New(dynamicT)
			x.Elem().Set(dynamicX)
			env.AddVar(name, x)
		}

		for _, stmt := range s.Body.List {
			clause := stmt.(*CaseClause)
			for _, expr := range clause.List {
				t := expr.KnownType()[0]
				if t.Kind() == reflect.Interface {
					if dynamicT.Implements(t) {
						return InterpStmt(clause, env)
					}
				} else if t == dynamicT {
					return InterpStmt(clause, env)
				}
			}
		}
		return InterpStmt(s.def, env)

	default:
		panic(dytc(fmt.Sprintf("Unsupported statement %T", s)))
	}
	return nil, nil
}

func assign(lhs Expr, rhs reflect.Value, env Env) error {
	lhs = skipSuperfluousParens(lhs)
	// Always evaluate even if we are doing a map index assign. There are some nasty
	// corner cases with map index comparibility that is best left not reimplemented.
	if l, err := evalTypedExpr(lhs, lhs.KnownType(), env); err != nil {
		return err
	} else if index, ok := lhs.(*IndexExpr); ok && index.X.KnownType()[0].Kind() == reflect.Map {
		mT := index.X.KnownType()[0]
		// known to succeed from above
		m, _ := evalTypedExpr(index.X, knownType{mT}, env)
		k, _ := evalTypedExpr(index.Index, knownType{mT.Elem()}, env)
		m[0].SetMapIndex(k[0], rhs)
	} else {
		l[0].Set(rhs)
	}
	return nil
}

func interpBlock(list []Stmt, env Env) (last *State, err error) {
	for i := 0; i < len(list); i += 1 {
		if last, err = InterpStmt(list[i], env); err != nil {
			return last, err
		} else if last != nil {
			if last, i = branch(list, last, env); last != nil {
				break
			}
		}
	}
	return last, nil
}

// Really naive implementation which simply scans block for the branch target
func branch(list []Stmt, last *State, env Env) (*State, int) {
	branch, ok := last.Last.(*BranchStmt)
	if !ok {
		return last, 0
	}

	// breaks should go to the next stmt, goto should go to the labeled stmt
	brk := 0
	if branch.Tok == token.GOTO {
		brk = -1
	}
	for i, stmt := range list {
		switch s := stmt.(type) {
		case *LabeledStmt:
			if branch.Label != nil && branch.Label.Name == s.Label.Name {
				return nil, i+brk
			}
		// TODO[crc] add SelectStmt and RangeStmt here when implemented
		case *ForStmt, *SwitchStmt, *TypeSwitchStmt:
			if branch.Label == nil {
				return nil, i+brk
			}
		}
	}
	return last, 0
}
