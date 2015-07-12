package eval

import (
	"reflect"
	"testing"
)

type XI interface {
	x()
}

type YI interface {
	y()
}

type ZI interface {
	x()
}

type X int
type Y int
type Z int

func (X) x() {}
func (Y) y() {}
func (Z) x() {}

func TestTypeAssertInterfaceToInterface(t *testing.T) {
	env := makeTypeAssertEnv()
	a := XI(X(0))
	env.Vars["a"] = reflect.ValueOf(&a)
	expectResults(t, "a.(ZI)", env, a.(ZI), true)
}

func TestTypeAssertInterfaceToDynamic(t *testing.T) {
	env := makeTypeAssertEnv()
	a := XI(X(0))
	env.Vars["a"] = reflect.ValueOf(&a)
	expectResults(t, "a.(X)", env, a.(ZI), true)
}

func TestTypeAssertInterfaceToInterfacePanic(t *testing.T) {
	env := makeTypeAssertEnv()
	a := XI(X(0))
	env.Vars["a"] = reflect.ValueOf(&a)
	expectPanic(t, "a.(YI)", env, "interface conversion: eval.XI is not eval.YI: missing method y")
}

func TestTypeAssertInterfaceToDynamicPanic(t *testing.T) {
	env := makeTypeAssertEnv()
	a := XI(X(0))
	env.Vars["a"] = reflect.ValueOf(&a)
	expectPanic(t, "a.(Z)", env, "interface conversion: eval.XI is eval.X, not eval.Z")
}

func TestTypeAssertNilToInterfacePanic(t *testing.T) {
	env := makeTypeAssertEnv()
	a := XI(nil)
	env.Vars["a"] = reflect.ValueOf(&a)
	expectPanic(t, "a.(XI)", env, "interface conversion: nil is not eval.XI")
}

func TestTypeAssertNilToDynamicPanic(t *testing.T) {
	env := makeTypeAssertEnv()
	a := XI(nil)
	env.Vars["a"] = reflect.ValueOf(&a)
	expectPanic(t, "a.(X)", env, "interface conversion: nil is not eval.X")
}

func makeTypeAssertEnv() *SimpleEnv {
	// This little pointer dance is required to prevent
	// reflect.TypeOf(typ interface{}) from interpreting
	// typ as value which has been promoted to interface{}
	//
	// Basically, there is no way for reflect.TypeOf to
	// distinguish between
	// reflect.TypeOf(X(0)) - and -
	// reflect.TypeOf(XI(X(0)))
	// as XI is demoted to interface{} without any change to
	// the interal interface data.

	env := MakeSimpleEnv()
	var xi *XI = new(XI)
	var yi *YI = new(YI)
	var zi *ZI = new(ZI)
	*xi = XI(X(0))
	*yi = YI(Y(0))
	*zi = ZI(Z(0))
	env.Types["XI"] = reflect.TypeOf(xi).Elem()
	env.Types["YI"] = reflect.TypeOf(yi).Elem()
	env.Types["ZI"] = reflect.TypeOf(zi).Elem()
	env.Types["X"] = reflect.TypeOf(X(0))
	env.Types["Y"] = reflect.TypeOf(Y(0))
	env.Types["Z"] = reflect.TypeOf(Z(0))

	return env
}
