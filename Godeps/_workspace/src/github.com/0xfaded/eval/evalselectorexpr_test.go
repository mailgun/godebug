package eval

import (
	"fmt"
	"reflect"
	"testing"
)

type SelInterface interface {
	E() int
	F() int
}

type SelInt int
type SelNested struct {
	D int
}
type SelStruct struct {
	A int
	B struct { C int }
	SelNested
}

func (SelStruct) E() int {
	return 1
}

func (*SelStruct) F() int {
	return 2
}

func (SelInt) E() int {
	return 1
}

func (*SelInt) F() int {
	return 2
}

func TestSelectStructField(t *testing.T) {
	env := MakeSimpleEnv()
	s := SelStruct{}
	env.Vars["s"] = reflect.ValueOf(&s)
	expectResult(t, "s.A", env, s.A)
}

func TestSelectStructNestedField(t *testing.T) {
	env := MakeSimpleEnv()
	s := SelStruct{}
	env.Vars["s"] = reflect.ValueOf(&s)
	expectResult(t, "s.B", env, s.B)
	expectResult(t, "s.B.C", env, s.B.C)
}

func TestSelectStructAnonField(t *testing.T) {
	env := MakeSimpleEnv()
	s := SelStruct{}
	env.Vars["s"] = reflect.ValueOf(&s)
	expectResult(t, "s.D", env, s.D)
	expectResult(t, "s.SelNested.D", env, s.SelNested.D)
}

func TestSelectStructMethod(t *testing.T) {
	env := MakeSimpleEnv()
	s := SelStruct{}
	env.Vars["s"] = reflect.ValueOf(&s)
	expectResult(t, "s.E()", env, s.E())
}

func TestSelectStructMethodP(t *testing.T) {
	env := MakeSimpleEnv()
	s := SelStruct{}
	env.Vars["s"] = reflect.ValueOf(&s)
	expectResult(t, "s.F()", env, s.F())
}

func TestSelectIntMethod(t *testing.T) {
	env := MakeSimpleEnv()
	i := SelInt(0)
	env.Vars["i"] = reflect.ValueOf(&i)
	expectResult(t, "i.E()", env, i.E())
}

func TestSelectIntMethodP(t *testing.T) {
	env := MakeSimpleEnv()
	i := SelInt(0)
	env.Vars["i"] = reflect.ValueOf(&i)
	expectResult(t, "i.F()", env, i.F())
}

func TestSelectStructPField(t *testing.T) {
	env := MakeSimpleEnv()
	s := &SelStruct{}
	env.Vars["s"] = reflect.ValueOf(&s)
	expectResult(t, "s.A", env, s.A)
}

func TestSelectStructPNestedField(t *testing.T) {
	env := MakeSimpleEnv()
	s := &SelStruct{}
	env.Vars["s"] = reflect.ValueOf(&s)
	expectResult(t, "s.B", env, s.B)
	expectResult(t, "s.B.C", env, s.B.C)
}

func TestSelectStructPAnonField(t *testing.T) {
	env := MakeSimpleEnv()
	s := &SelStruct{}
	env.Vars["s"] = reflect.ValueOf(&s)
	expectResult(t, "s.D", env, s.D)
	expectResult(t, "s.SelNested.D", env, s.SelNested.D)
}

func TestSelectStructPMethod(t *testing.T) {
	env := MakeSimpleEnv()
	s := &SelStruct{}
	env.Vars["s"] = reflect.ValueOf(&s)
	expectResult(t, "s.E()", env, s.E())
}

func TestSelectStructPMethodP(t *testing.T) {
	env := MakeSimpleEnv()
	s := &SelStruct{}
	env.Vars["s"] = reflect.ValueOf(&s)
	expectResult(t, "s.F()", env, s.F())
}

func TestSelectIntPMethod(t *testing.T) {
	env := MakeSimpleEnv()
	i := new(SelInt)
	env.Vars["i"] = reflect.ValueOf(&i)
	expectResult(t, "i.E()", env, i.E())
}

func TestSelectIntPMethodP(t *testing.T) {
	env := MakeSimpleEnv()
	i := new(SelInt)
	env.Vars["i"] = reflect.ValueOf(&i)
	expectResult(t, "i.F()", env, i.F())
}

func TestSelectInterfaceMethod(t *testing.T) {
	env := MakeSimpleEnv()
	var i SelInterface = new(SelInt)
	env.Vars["i"] = reflect.ValueOf(&i)
	expectResult(t, "i.E()", env, i.E())
}

func TestSelectInterfaceMethodP(t *testing.T) {
	env := MakeSimpleEnv()
	var i SelInterface = new(SelInt)
	env.Vars["i"] = reflect.ValueOf(&i)
	expectResult(t, "i.F()", env, i.F())
}

func TestSelectPackageFunc(t *testing.T) {
	env := MakeSimpleEnv()
	pkg := MakeSimpleEnv()
	env.Pkgs["fmt"] = pkg
	pkg.Funcs["Sprintf"] = reflect.ValueOf(fmt.Sprintf)

	expectResult(t, `fmt.Sprintf("abc")`, env, fmt.Sprintf("abc"))
}

