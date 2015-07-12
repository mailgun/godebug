package eval

import (
	"testing"
	"reflect"
)

// Test NoArg()
func TestCheckCallExprNoArgXX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `s()`, env, reflect.TypeOf(s()))
}

// Test NoArg(1)
func TestCheckCallExprNoArgXInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1.5)
func TestCheckCallExprNoArgXFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1.5)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(true)
func TestCheckCallExprNoArgXBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(true)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(I(1))
func TestCheckCallExprNoArgXIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(I(1))`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(is...)
func TestCheckCallExprNoArgXInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(is...)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(e())
func TestCheckCallExprNoArgXEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(e())`, env,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(s())
func TestCheckCallExprNoArgXSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(s())`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(m())
func TestCheckCallExprNoArgXMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(m())`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(mt())
func TestCheckCallExprNoArgXMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(mt())`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1)
func TestCheckCallExprNoArgIntX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1, 1)
func TestCheckCallExprNoArgIntInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1, 1)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1, 1.5)
func TestCheckCallExprNoArgIntFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1, 1.5)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1, true)
func TestCheckCallExprNoArgIntBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1, true)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1, I(1))
func TestCheckCallExprNoArgIntIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1, I(1))`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1, is...)
func TestCheckCallExprNoArgIntInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1, is...)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1, e())
func TestCheckCallExprNoArgIntEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1, e())`, env,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1, s())
func TestCheckCallExprNoArgIntSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1, s())`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1, m())
func TestCheckCallExprNoArgIntMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1, m())`, env,
		`multiple-value m() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1, mt())
func TestCheckCallExprNoArgIntMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1, mt())`, env,
		`multiple-value mt() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1.5)
func TestCheckCallExprNoArgFloatX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1.5)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1.5, 1)
func TestCheckCallExprNoArgFloatInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1.5, 1)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1.5, 1.5)
func TestCheckCallExprNoArgFloatFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1.5, 1.5)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1.5, true)
func TestCheckCallExprNoArgFloatBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1.5, true)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1.5, I(1))
func TestCheckCallExprNoArgFloatIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1.5, I(1))`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1.5, is...)
func TestCheckCallExprNoArgFloatInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1.5, is...)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1.5, e())
func TestCheckCallExprNoArgFloatEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1.5, e())`, env,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1.5, s())
func TestCheckCallExprNoArgFloatSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1.5, s())`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1.5, m())
func TestCheckCallExprNoArgFloatMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1.5, m())`, env,
		`multiple-value m() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(1.5, mt())
func TestCheckCallExprNoArgFloatMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(1.5, mt())`, env,
		`multiple-value mt() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(true)
func TestCheckCallExprNoArgBoolX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(true)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(true, 1)
func TestCheckCallExprNoArgBoolInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(true, 1)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(true, 1.5)
func TestCheckCallExprNoArgBoolFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(true, 1.5)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(true, true)
func TestCheckCallExprNoArgBoolBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(true, true)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(true, I(1))
func TestCheckCallExprNoArgBoolIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(true, I(1))`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(true, is...)
func TestCheckCallExprNoArgBoolInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(true, is...)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(true, e())
func TestCheckCallExprNoArgBoolEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(true, e())`, env,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(true, s())
func TestCheckCallExprNoArgBoolSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(true, s())`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(true, m())
func TestCheckCallExprNoArgBoolMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(true, m())`, env,
		`multiple-value m() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(true, mt())
func TestCheckCallExprNoArgBoolMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(true, mt())`, env,
		`multiple-value mt() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(I(1))
func TestCheckCallExprNoArgIntTypedX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(I(1))`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(I(1), 1)
func TestCheckCallExprNoArgIntTypedInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(I(1), 1)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(I(1), 1.5)
func TestCheckCallExprNoArgIntTypedFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(I(1), 1.5)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(I(1), true)
func TestCheckCallExprNoArgIntTypedBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(I(1), true)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(I(1), I(1))
func TestCheckCallExprNoArgIntTypedIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(I(1), I(1))`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(I(1), is...)
func TestCheckCallExprNoArgIntTypedInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(I(1), is...)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(I(1), e())
func TestCheckCallExprNoArgIntTypedEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(I(1), e())`, env,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(I(1), s())
func TestCheckCallExprNoArgIntTypedSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(I(1), s())`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(I(1), m())
func TestCheckCallExprNoArgIntTypedMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(I(1), m())`, env,
		`multiple-value m() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(I(1), mt())
func TestCheckCallExprNoArgIntTypedMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(I(1), mt())`, env,
		`multiple-value mt() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(is...)
func TestCheckCallExprNoArgIntsX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(is...)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(is..., 1)
func TestCheckCallExprNoArgIntsInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test NoArg(is..., 1.5)
func TestCheckCallExprNoArgIntsFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test NoArg(is..., true)
func TestCheckCallExprNoArgIntsBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test NoArg(is..., I(1))
func TestCheckCallExprNoArgIntsIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test NoArg(is..., is...)
func TestCheckCallExprNoArgIntsInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test NoArg(is..., e())
func TestCheckCallExprNoArgIntsEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test NoArg(is..., s())
func TestCheckCallExprNoArgIntsSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test NoArg(is..., m())
func TestCheckCallExprNoArgIntsMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test NoArg(is..., mt())
func TestCheckCallExprNoArgIntsMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test NoArg(e())
func TestCheckCallExprNoArgEmptyFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(e())`, env,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(e(), 1)
func TestCheckCallExprNoArgEmptyFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(e(), 1)`, env,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(e(), 1.5)
func TestCheckCallExprNoArgEmptyFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(e(), 1.5)`, env,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(e(), true)
func TestCheckCallExprNoArgEmptyFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(e(), true)`, env,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(e(), I(1))
func TestCheckCallExprNoArgEmptyFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(e(), I(1))`, env,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(e(), is...)
func TestCheckCallExprNoArgEmptyFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(e(), is...)`, env,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(e(), e())
func TestCheckCallExprNoArgEmptyFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(e(), e())`, env,
		`e() used as value`,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(e(), s())
func TestCheckCallExprNoArgEmptyFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(e(), s())`, env,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(e(), m())
func TestCheckCallExprNoArgEmptyFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(e(), m())`, env,
		`e() used as value`,
		`multiple-value m() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(e(), mt())
func TestCheckCallExprNoArgEmptyFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(e(), mt())`, env,
		`e() used as value`,
		`multiple-value mt() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(s())
func TestCheckCallExprNoArgSingleFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(s())`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(s(), 1)
func TestCheckCallExprNoArgSingleFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(s(), 1)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(s(), 1.5)
func TestCheckCallExprNoArgSingleFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(s(), 1.5)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(s(), true)
func TestCheckCallExprNoArgSingleFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(s(), true)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(s(), I(1))
func TestCheckCallExprNoArgSingleFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(s(), I(1))`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(s(), is...)
func TestCheckCallExprNoArgSingleFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(s(), is...)`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(s(), e())
func TestCheckCallExprNoArgSingleFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(s(), e())`, env,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(s(), s())
func TestCheckCallExprNoArgSingleFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(s(), s())`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(s(), m())
func TestCheckCallExprNoArgSingleFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(s(), m())`, env,
		`multiple-value m() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(s(), mt())
func TestCheckCallExprNoArgSingleFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(s(), mt())`, env,
		`multiple-value mt() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(m())
func TestCheckCallExprNoArgMultiFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(m())`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(m(), 1)
func TestCheckCallExprNoArgMultiFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(m(), 1)`, env,
		`multiple-value m() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(m(), 1.5)
func TestCheckCallExprNoArgMultiFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(m(), 1.5)`, env,
		`multiple-value m() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(m(), true)
func TestCheckCallExprNoArgMultiFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(m(), true)`, env,
		`multiple-value m() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(m(), I(1))
func TestCheckCallExprNoArgMultiFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(m(), I(1))`, env,
		`multiple-value m() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(m(), is...)
func TestCheckCallExprNoArgMultiFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(m(), is...)`, env,
		`multiple-value m() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(m(), e())
func TestCheckCallExprNoArgMultiFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(m(), e())`, env,
		`multiple-value m() in single-value context`,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(m(), s())
func TestCheckCallExprNoArgMultiFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(m(), s())`, env,
		`multiple-value m() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(m(), m())
func TestCheckCallExprNoArgMultiFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(m(), m())`, env,
		`multiple-value m() in single-value context`,
		`multiple-value m() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(m(), mt())
func TestCheckCallExprNoArgMultiFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(m(), mt())`, env,
		`multiple-value m() in single-value context`,
		`multiple-value mt() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(mt())
func TestCheckCallExprNoArgMultiFuncMixedTypesX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(mt())`, env,
		`too many arguments in call to s`,
	)

}

// Test NoArg(mt(), 1)
func TestCheckCallExprNoArgMultiFuncMixedTypesInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(mt(), 1)`, env,
		`multiple-value mt() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(mt(), 1.5)
func TestCheckCallExprNoArgMultiFuncMixedTypesFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(mt(), 1.5)`, env,
		`multiple-value mt() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(mt(), true)
func TestCheckCallExprNoArgMultiFuncMixedTypesBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(mt(), true)`, env,
		`multiple-value mt() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(mt(), I(1))
func TestCheckCallExprNoArgMultiFuncMixedTypesIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(mt(), I(1))`, env,
		`multiple-value mt() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(mt(), is...)
func TestCheckCallExprNoArgMultiFuncMixedTypesInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(mt(), is...)`, env,
		`multiple-value mt() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(mt(), e())
func TestCheckCallExprNoArgMultiFuncMixedTypesEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(mt(), e())`, env,
		`multiple-value mt() in single-value context`,
		`e() used as value`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(mt(), s())
func TestCheckCallExprNoArgMultiFuncMixedTypesSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(mt(), s())`, env,
		`multiple-value mt() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(mt(), m())
func TestCheckCallExprNoArgMultiFuncMixedTypesMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(mt(), m())`, env,
		`multiple-value mt() in single-value context`,
		`multiple-value m() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test NoArg(mt(), mt())
func TestCheckCallExprNoArgMultiFuncMixedTypesMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `s(mt(), mt())`, env,
		`multiple-value mt() in single-value context`,
		`multiple-value mt() in single-value context`,
		`too many arguments in call to s`,
	)

}

// Test Fixed()
func TestCheckCallExprFixedXX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f()`, env,
		`not enough arguments in call to f`,
	)

}

// Test Fixed(1)
func TestCheckCallExprFixedXInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1)`, env,
		`not enough arguments in call to f`,
	)

}

// Test Fixed(1.5)
func TestCheckCallExprFixedXFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1.5)`, env,
		`constant 1.5 truncated to integer`,
		`not enough arguments in call to f`,
	)

}

// Test Fixed(true)
func TestCheckCallExprFixedXBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(true)`, env,
		`cannot use true (type bool) as type int in function argument`,
		`not enough arguments in call to f`,
	)

}

// Test Fixed(I(1))
func TestCheckCallExprFixedXIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
		`not enough arguments in call to f`,
	)

}

// Test Fixed(is...)
func TestCheckCallExprFixedXInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(is...)`, env,
		`cannot use is (type []int) as type int in function argument`,
		`not enough arguments in call to f`,
	)

}

// Test Fixed(e())
func TestCheckCallExprFixedXEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(e())`, env,
		`e() used as value`,
		`not enough arguments in call to f`,
	)

}

// Test Fixed(s())
func TestCheckCallExprFixedXSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(s())`, env,
		`not enough arguments in call to f`,
	)

}

// Test Fixed(m())
func TestCheckCallExprFixedXMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(m())`, env,
		`cannot use int as type bool in argument to f`,
	)

}

// Test Fixed(mt())
func TestCheckCallExprFixedXMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(mt())`, env,
		`cannot use eval.I as type bool in argument to f`,
	)

}

// Test Fixed(1)
func TestCheckCallExprFixedIntX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1)`, env,
		`not enough arguments in call to f`,
	)

}

// Test Fixed(1, 1)
func TestCheckCallExprFixedIntInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1, 1)`, env,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test Fixed(1, 1.5)
func TestCheckCallExprFixedIntFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1, 1.5)`, env,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test Fixed(1, true)
func TestCheckCallExprFixedIntBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `f(1, true)`, env, reflect.TypeOf(f(1, true)))
}

// Test Fixed(1, I(1))
func TestCheckCallExprFixedIntIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1, I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test Fixed(1, is...)
func TestCheckCallExprFixedIntInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1, is...)`, env,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to f`,
	)

}

// Test Fixed(1, e())
func TestCheckCallExprFixedIntEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1, e())`, env,
		`e() used as value`,
	)

}

// Test Fixed(1, s())
func TestCheckCallExprFixedIntSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1, s())`, env,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test Fixed(1, m())
func TestCheckCallExprFixedIntMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1, m())`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test Fixed(1, mt())
func TestCheckCallExprFixedIntMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1, mt())`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test Fixed(1.5)
func TestCheckCallExprFixedFloatX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1.5)`, env,
		`constant 1.5 truncated to integer`,
		`not enough arguments in call to f`,
	)

}

// Test Fixed(1.5, 1)
func TestCheckCallExprFixedFloatInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1.5, 1)`, env,
		`constant 1.5 truncated to integer`,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test Fixed(1.5, 1.5)
func TestCheckCallExprFixedFloatFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1.5, 1.5)`, env,
		`constant 1.5 truncated to integer`,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test Fixed(1.5, true)
func TestCheckCallExprFixedFloatBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1.5, true)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Fixed(1.5, I(1))
func TestCheckCallExprFixedFloatIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1.5, I(1))`, env,
		`constant 1.5 truncated to integer`,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test Fixed(1.5, is...)
func TestCheckCallExprFixedFloatInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1.5, is...)`, env,
		`constant 1.5 truncated to integer`,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to f`,
	)

}

// Test Fixed(1.5, e())
func TestCheckCallExprFixedFloatEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1.5, e())`, env,
		`e() used as value`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Fixed(1.5, s())
func TestCheckCallExprFixedFloatSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1.5, s())`, env,
		`constant 1.5 truncated to integer`,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test Fixed(1.5, m())
func TestCheckCallExprFixedFloatMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1.5, m())`, env,
		`multiple-value m() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Fixed(1.5, mt())
func TestCheckCallExprFixedFloatMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(1.5, mt())`, env,
		`multiple-value mt() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Fixed(true)
func TestCheckCallExprFixedBoolX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(true)`, env,
		`cannot use true (type bool) as type int in function argument`,
		`not enough arguments in call to f`,
	)

}

// Test Fixed(true, 1)
func TestCheckCallExprFixedBoolInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(true, 1)`, env,
		`cannot use true (type bool) as type int in function argument`,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test Fixed(true, 1.5)
func TestCheckCallExprFixedBoolFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(true, 1.5)`, env,
		`cannot use true (type bool) as type int in function argument`,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test Fixed(true, true)
func TestCheckCallExprFixedBoolBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(true, true)`, env,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Fixed(true, I(1))
func TestCheckCallExprFixedBoolIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(true, I(1))`, env,
		`cannot use true (type bool) as type int in function argument`,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test Fixed(true, is...)
func TestCheckCallExprFixedBoolInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(true, is...)`, env,
		`cannot use true (type bool) as type int in function argument`,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to f`,
	)

}

// Test Fixed(true, e())
func TestCheckCallExprFixedBoolEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(true, e())`, env,
		`e() used as value`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Fixed(true, s())
func TestCheckCallExprFixedBoolSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(true, s())`, env,
		`cannot use true (type bool) as type int in function argument`,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test Fixed(true, m())
func TestCheckCallExprFixedBoolMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(true, m())`, env,
		`multiple-value m() in single-value context`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Fixed(true, mt())
func TestCheckCallExprFixedBoolMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(true, mt())`, env,
		`multiple-value mt() in single-value context`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Fixed(I(1))
func TestCheckCallExprFixedIntTypedX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
		`not enough arguments in call to f`,
	)

}

// Test Fixed(I(1), 1)
func TestCheckCallExprFixedIntTypedInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(I(1), 1)`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test Fixed(I(1), 1.5)
func TestCheckCallExprFixedIntTypedFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(I(1), 1.5)`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test Fixed(I(1), true)
func TestCheckCallExprFixedIntTypedBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(I(1), true)`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Fixed(I(1), I(1))
func TestCheckCallExprFixedIntTypedIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(I(1), I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test Fixed(I(1), is...)
func TestCheckCallExprFixedIntTypedInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(I(1), is...)`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to f`,
	)

}

// Test Fixed(I(1), e())
func TestCheckCallExprFixedIntTypedEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(I(1), e())`, env,
		`e() used as value`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Fixed(I(1), s())
func TestCheckCallExprFixedIntTypedSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(I(1), s())`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test Fixed(I(1), m())
func TestCheckCallExprFixedIntTypedMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(I(1), m())`, env,
		`multiple-value m() in single-value context`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Fixed(I(1), mt())
func TestCheckCallExprFixedIntTypedMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(I(1), mt())`, env,
		`multiple-value mt() in single-value context`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Fixed(is...)
func TestCheckCallExprFixedIntsX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(is...)`, env,
		`cannot use is (type []int) as type int in function argument`,
		`not enough arguments in call to f`,
	)

}

// Test Fixed(is..., 1)
func TestCheckCallExprFixedIntsInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Fixed(is..., 1.5)
func TestCheckCallExprFixedIntsFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Fixed(is..., true)
func TestCheckCallExprFixedIntsBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Fixed(is..., I(1))
func TestCheckCallExprFixedIntsIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Fixed(is..., is...)
func TestCheckCallExprFixedIntsInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Fixed(is..., e())
func TestCheckCallExprFixedIntsEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Fixed(is..., s())
func TestCheckCallExprFixedIntsSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Fixed(is..., m())
func TestCheckCallExprFixedIntsMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Fixed(is..., mt())
func TestCheckCallExprFixedIntsMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Fixed(e())
func TestCheckCallExprFixedEmptyFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(e())`, env,
		`e() used as value`,
		`not enough arguments in call to f`,
	)

}

// Test Fixed(e(), 1)
func TestCheckCallExprFixedEmptyFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(e(), 1)`, env,
		`e() used as value`,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test Fixed(e(), 1.5)
func TestCheckCallExprFixedEmptyFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(e(), 1.5)`, env,
		`e() used as value`,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test Fixed(e(), true)
func TestCheckCallExprFixedEmptyFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(e(), true)`, env,
		`e() used as value`,
	)

}

// Test Fixed(e(), I(1))
func TestCheckCallExprFixedEmptyFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(e(), I(1))`, env,
		`e() used as value`,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test Fixed(e(), is...)
func TestCheckCallExprFixedEmptyFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(e(), is...)`, env,
		`e() used as value`,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to f`,
	)

}

// Test Fixed(e(), e())
func TestCheckCallExprFixedEmptyFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(e(), e())`, env,
		`e() used as value`,
		`e() used as value`,
	)

}

// Test Fixed(e(), s())
func TestCheckCallExprFixedEmptyFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(e(), s())`, env,
		`e() used as value`,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test Fixed(e(), m())
func TestCheckCallExprFixedEmptyFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(e(), m())`, env,
		`e() used as value`,
		`multiple-value m() in single-value context`,
	)

}

// Test Fixed(e(), mt())
func TestCheckCallExprFixedEmptyFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(e(), mt())`, env,
		`e() used as value`,
		`multiple-value mt() in single-value context`,
	)

}

// Test Fixed(s())
func TestCheckCallExprFixedSingleFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(s())`, env,
		`not enough arguments in call to f`,
	)

}

// Test Fixed(s(), 1)
func TestCheckCallExprFixedSingleFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(s(), 1)`, env,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test Fixed(s(), 1.5)
func TestCheckCallExprFixedSingleFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(s(), 1.5)`, env,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test Fixed(s(), true)
func TestCheckCallExprFixedSingleFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `f(s(), true)`, env, reflect.TypeOf(f(s(), true)))
}

// Test Fixed(s(), I(1))
func TestCheckCallExprFixedSingleFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(s(), I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test Fixed(s(), is...)
func TestCheckCallExprFixedSingleFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(s(), is...)`, env,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to f`,
	)

}

// Test Fixed(s(), e())
func TestCheckCallExprFixedSingleFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(s(), e())`, env,
		`e() used as value`,
	)

}

// Test Fixed(s(), s())
func TestCheckCallExprFixedSingleFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(s(), s())`, env,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test Fixed(s(), m())
func TestCheckCallExprFixedSingleFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(s(), m())`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test Fixed(s(), mt())
func TestCheckCallExprFixedSingleFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(s(), mt())`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test Fixed(m())
func TestCheckCallExprFixedMultiFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(m())`, env,
		`cannot use int as type bool in argument to f`,
	)

}

// Test Fixed(m(), 1)
func TestCheckCallExprFixedMultiFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(m(), 1)`, env,
		`multiple-value m() in single-value context`,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test Fixed(m(), 1.5)
func TestCheckCallExprFixedMultiFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(m(), 1.5)`, env,
		`multiple-value m() in single-value context`,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test Fixed(m(), true)
func TestCheckCallExprFixedMultiFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(m(), true)`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test Fixed(m(), I(1))
func TestCheckCallExprFixedMultiFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(m(), I(1))`, env,
		`multiple-value m() in single-value context`,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test Fixed(m(), is...)
func TestCheckCallExprFixedMultiFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(m(), is...)`, env,
		`multiple-value m() in single-value context`,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to f`,
	)

}

// Test Fixed(m(), e())
func TestCheckCallExprFixedMultiFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(m(), e())`, env,
		`multiple-value m() in single-value context`,
		`e() used as value`,
	)

}

// Test Fixed(m(), s())
func TestCheckCallExprFixedMultiFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(m(), s())`, env,
		`multiple-value m() in single-value context`,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test Fixed(m(), m())
func TestCheckCallExprFixedMultiFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(m(), m())`, env,
		`multiple-value m() in single-value context`,
		`multiple-value m() in single-value context`,
	)

}

// Test Fixed(m(), mt())
func TestCheckCallExprFixedMultiFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(m(), mt())`, env,
		`multiple-value m() in single-value context`,
		`multiple-value mt() in single-value context`,
	)

}

// Test Fixed(mt())
func TestCheckCallExprFixedMultiFuncMixedTypesX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(mt())`, env,
		`cannot use eval.I as type bool in argument to f`,
	)

}

// Test Fixed(mt(), 1)
func TestCheckCallExprFixedMultiFuncMixedTypesInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(mt(), 1)`, env,
		`multiple-value mt() in single-value context`,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test Fixed(mt(), 1.5)
func TestCheckCallExprFixedMultiFuncMixedTypesFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(mt(), 1.5)`, env,
		`multiple-value mt() in single-value context`,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test Fixed(mt(), true)
func TestCheckCallExprFixedMultiFuncMixedTypesBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(mt(), true)`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test Fixed(mt(), I(1))
func TestCheckCallExprFixedMultiFuncMixedTypesIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(mt(), I(1))`, env,
		`multiple-value mt() in single-value context`,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test Fixed(mt(), is...)
func TestCheckCallExprFixedMultiFuncMixedTypesInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(mt(), is...)`, env,
		`multiple-value mt() in single-value context`,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to f`,
	)

}

// Test Fixed(mt(), e())
func TestCheckCallExprFixedMultiFuncMixedTypesEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(mt(), e())`, env,
		`multiple-value mt() in single-value context`,
		`e() used as value`,
	)

}

// Test Fixed(mt(), s())
func TestCheckCallExprFixedMultiFuncMixedTypesSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(mt(), s())`, env,
		`multiple-value mt() in single-value context`,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test Fixed(mt(), m())
func TestCheckCallExprFixedMultiFuncMixedTypesMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(mt(), m())`, env,
		`multiple-value mt() in single-value context`,
		`multiple-value m() in single-value context`,
	)

}

// Test Fixed(mt(), mt())
func TestCheckCallExprFixedMultiFuncMixedTypesMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `f(mt(), mt())`, env,
		`multiple-value mt() in single-value context`,
		`multiple-value mt() in single-value context`,
	)

}

// Test FixedTyped()
func TestCheckCallExprFixedTypedXX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft()`, env,
		`not enough arguments in call to ft`,
	)

}

// Test FixedTyped(1)
func TestCheckCallExprFixedTypedXInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1)`, env,
		`not enough arguments in call to ft`,
	)

}

// Test FixedTyped(1.5)
func TestCheckCallExprFixedTypedXFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1.5)`, env,
		`constant 1.5 truncated to integer`,
		`not enough arguments in call to ft`,
	)

}

// Test FixedTyped(true)
func TestCheckCallExprFixedTypedXBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(true)`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
		`not enough arguments in call to ft`,
	)

}

// Test FixedTyped(I(1))
func TestCheckCallExprFixedTypedXIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(I(1))`, env,
		`not enough arguments in call to ft`,
	)

}

// Test FixedTyped(is...)
func TestCheckCallExprFixedTypedXInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(is...)`, env,
		`cannot use is (type []int) as type eval.I in function argument`,
		`not enough arguments in call to ft`,
	)

}

// Test FixedTyped(e())
func TestCheckCallExprFixedTypedXEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(e())`, env,
		`e() used as value`,
		`not enough arguments in call to ft`,
	)

}

// Test FixedTyped(s())
func TestCheckCallExprFixedTypedXSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(s())`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
		`not enough arguments in call to ft`,
	)

}

// Test FixedTyped(m())
func TestCheckCallExprFixedTypedXMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(m())`, env,
		`cannot use int as type eval.I in argument to ft`,
		`cannot use int as type bool in argument to ft`,
	)

}

// Test FixedTyped(mt())
func TestCheckCallExprFixedTypedXMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(mt())`, env,
		`cannot use int as type eval.I in argument to ft`,
		`cannot use eval.I as type bool in argument to ft`,
	)

}

// Test FixedTyped(1)
func TestCheckCallExprFixedTypedIntX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1)`, env,
		`not enough arguments in call to ft`,
	)

}

// Test FixedTyped(1, 1)
func TestCheckCallExprFixedTypedIntInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1, 1)`, env,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(1, 1.5)
func TestCheckCallExprFixedTypedIntFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1, 1.5)`, env,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test FixedTyped(1, true)
func TestCheckCallExprFixedTypedIntBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `ft(1, true)`, env, reflect.TypeOf(ft(1, true)))
}

// Test FixedTyped(1, I(1))
func TestCheckCallExprFixedTypedIntIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1, I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test FixedTyped(1, is...)
func TestCheckCallExprFixedTypedIntInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1, is...)`, env,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to ft`,
	)

}

// Test FixedTyped(1, e())
func TestCheckCallExprFixedTypedIntEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1, e())`, env,
		`e() used as value`,
	)

}

// Test FixedTyped(1, s())
func TestCheckCallExprFixedTypedIntSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1, s())`, env,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(1, m())
func TestCheckCallExprFixedTypedIntMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1, m())`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test FixedTyped(1, mt())
func TestCheckCallExprFixedTypedIntMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1, mt())`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test FixedTyped(1.5)
func TestCheckCallExprFixedTypedFloatX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1.5)`, env,
		`constant 1.5 truncated to integer`,
		`not enough arguments in call to ft`,
	)

}

// Test FixedTyped(1.5, 1)
func TestCheckCallExprFixedTypedFloatInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1.5, 1)`, env,
		`constant 1.5 truncated to integer`,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(1.5, 1.5)
func TestCheckCallExprFixedTypedFloatFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1.5, 1.5)`, env,
		`constant 1.5 truncated to integer`,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test FixedTyped(1.5, true)
func TestCheckCallExprFixedTypedFloatBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1.5, true)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test FixedTyped(1.5, I(1))
func TestCheckCallExprFixedTypedFloatIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1.5, I(1))`, env,
		`constant 1.5 truncated to integer`,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test FixedTyped(1.5, is...)
func TestCheckCallExprFixedTypedFloatInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1.5, is...)`, env,
		`constant 1.5 truncated to integer`,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to ft`,
	)

}

// Test FixedTyped(1.5, e())
func TestCheckCallExprFixedTypedFloatEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1.5, e())`, env,
		`e() used as value`,
		`constant 1.5 truncated to integer`,
	)

}

// Test FixedTyped(1.5, s())
func TestCheckCallExprFixedTypedFloatSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1.5, s())`, env,
		`constant 1.5 truncated to integer`,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(1.5, m())
func TestCheckCallExprFixedTypedFloatMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1.5, m())`, env,
		`multiple-value m() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test FixedTyped(1.5, mt())
func TestCheckCallExprFixedTypedFloatMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(1.5, mt())`, env,
		`multiple-value mt() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test FixedTyped(true)
func TestCheckCallExprFixedTypedBoolX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(true)`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
		`not enough arguments in call to ft`,
	)

}

// Test FixedTyped(true, 1)
func TestCheckCallExprFixedTypedBoolInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(true, 1)`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(true, 1.5)
func TestCheckCallExprFixedTypedBoolFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(true, 1.5)`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test FixedTyped(true, true)
func TestCheckCallExprFixedTypedBoolBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(true, true)`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test FixedTyped(true, I(1))
func TestCheckCallExprFixedTypedBoolIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(true, I(1))`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test FixedTyped(true, is...)
func TestCheckCallExprFixedTypedBoolInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(true, is...)`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to ft`,
	)

}

// Test FixedTyped(true, e())
func TestCheckCallExprFixedTypedBoolEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(true, e())`, env,
		`e() used as value`,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test FixedTyped(true, s())
func TestCheckCallExprFixedTypedBoolSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(true, s())`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(true, m())
func TestCheckCallExprFixedTypedBoolMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(true, m())`, env,
		`multiple-value m() in single-value context`,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test FixedTyped(true, mt())
func TestCheckCallExprFixedTypedBoolMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(true, mt())`, env,
		`multiple-value mt() in single-value context`,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test FixedTyped(I(1))
func TestCheckCallExprFixedTypedIntTypedX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(I(1))`, env,
		`not enough arguments in call to ft`,
	)

}

// Test FixedTyped(I(1), 1)
func TestCheckCallExprFixedTypedIntTypedInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(I(1), 1)`, env,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(I(1), 1.5)
func TestCheckCallExprFixedTypedIntTypedFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(I(1), 1.5)`, env,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test FixedTyped(I(1), true)
func TestCheckCallExprFixedTypedIntTypedBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `ft(I(1), true)`, env, reflect.TypeOf(ft(I(1), true)))
}

// Test FixedTyped(I(1), I(1))
func TestCheckCallExprFixedTypedIntTypedIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(I(1), I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test FixedTyped(I(1), is...)
func TestCheckCallExprFixedTypedIntTypedInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(I(1), is...)`, env,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to ft`,
	)

}

// Test FixedTyped(I(1), e())
func TestCheckCallExprFixedTypedIntTypedEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(I(1), e())`, env,
		`e() used as value`,
	)

}

// Test FixedTyped(I(1), s())
func TestCheckCallExprFixedTypedIntTypedSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(I(1), s())`, env,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(I(1), m())
func TestCheckCallExprFixedTypedIntTypedMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(I(1), m())`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test FixedTyped(I(1), mt())
func TestCheckCallExprFixedTypedIntTypedMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(I(1), mt())`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test FixedTyped(is...)
func TestCheckCallExprFixedTypedIntsX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(is...)`, env,
		`cannot use is (type []int) as type eval.I in function argument`,
		`not enough arguments in call to ft`,
	)

}

// Test FixedTyped(is..., 1)
func TestCheckCallExprFixedTypedIntsInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test FixedTyped(is..., 1.5)
func TestCheckCallExprFixedTypedIntsFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test FixedTyped(is..., true)
func TestCheckCallExprFixedTypedIntsBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test FixedTyped(is..., I(1))
func TestCheckCallExprFixedTypedIntsIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test FixedTyped(is..., is...)
func TestCheckCallExprFixedTypedIntsInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test FixedTyped(is..., e())
func TestCheckCallExprFixedTypedIntsEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test FixedTyped(is..., s())
func TestCheckCallExprFixedTypedIntsSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test FixedTyped(is..., m())
func TestCheckCallExprFixedTypedIntsMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test FixedTyped(is..., mt())
func TestCheckCallExprFixedTypedIntsMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test FixedTyped(e())
func TestCheckCallExprFixedTypedEmptyFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(e())`, env,
		`e() used as value`,
		`not enough arguments in call to ft`,
	)

}

// Test FixedTyped(e(), 1)
func TestCheckCallExprFixedTypedEmptyFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(e(), 1)`, env,
		`e() used as value`,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(e(), 1.5)
func TestCheckCallExprFixedTypedEmptyFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(e(), 1.5)`, env,
		`e() used as value`,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test FixedTyped(e(), true)
func TestCheckCallExprFixedTypedEmptyFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(e(), true)`, env,
		`e() used as value`,
	)

}

// Test FixedTyped(e(), I(1))
func TestCheckCallExprFixedTypedEmptyFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(e(), I(1))`, env,
		`e() used as value`,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test FixedTyped(e(), is...)
func TestCheckCallExprFixedTypedEmptyFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(e(), is...)`, env,
		`e() used as value`,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to ft`,
	)

}

// Test FixedTyped(e(), e())
func TestCheckCallExprFixedTypedEmptyFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(e(), e())`, env,
		`e() used as value`,
		`e() used as value`,
	)

}

// Test FixedTyped(e(), s())
func TestCheckCallExprFixedTypedEmptyFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(e(), s())`, env,
		`e() used as value`,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(e(), m())
func TestCheckCallExprFixedTypedEmptyFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(e(), m())`, env,
		`e() used as value`,
		`multiple-value m() in single-value context`,
	)

}

// Test FixedTyped(e(), mt())
func TestCheckCallExprFixedTypedEmptyFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(e(), mt())`, env,
		`e() used as value`,
		`multiple-value mt() in single-value context`,
	)

}

// Test FixedTyped(s())
func TestCheckCallExprFixedTypedSingleFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(s())`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
		`not enough arguments in call to ft`,
	)

}

// Test FixedTyped(s(), 1)
func TestCheckCallExprFixedTypedSingleFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(s(), 1)`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(s(), 1.5)
func TestCheckCallExprFixedTypedSingleFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(s(), 1.5)`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test FixedTyped(s(), true)
func TestCheckCallExprFixedTypedSingleFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(s(), true)`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test FixedTyped(s(), I(1))
func TestCheckCallExprFixedTypedSingleFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(s(), I(1))`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test FixedTyped(s(), is...)
func TestCheckCallExprFixedTypedSingleFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(s(), is...)`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to ft`,
	)

}

// Test FixedTyped(s(), e())
func TestCheckCallExprFixedTypedSingleFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(s(), e())`, env,
		`e() used as value`,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test FixedTyped(s(), s())
func TestCheckCallExprFixedTypedSingleFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(s(), s())`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(s(), m())
func TestCheckCallExprFixedTypedSingleFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(s(), m())`, env,
		`multiple-value m() in single-value context`,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test FixedTyped(s(), mt())
func TestCheckCallExprFixedTypedSingleFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(s(), mt())`, env,
		`multiple-value mt() in single-value context`,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test FixedTyped(m())
func TestCheckCallExprFixedTypedMultiFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(m())`, env,
		`cannot use int as type eval.I in argument to ft`,
		`cannot use int as type bool in argument to ft`,
	)

}

// Test FixedTyped(m(), 1)
func TestCheckCallExprFixedTypedMultiFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(m(), 1)`, env,
		`multiple-value m() in single-value context`,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(m(), 1.5)
func TestCheckCallExprFixedTypedMultiFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(m(), 1.5)`, env,
		`multiple-value m() in single-value context`,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test FixedTyped(m(), true)
func TestCheckCallExprFixedTypedMultiFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(m(), true)`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test FixedTyped(m(), I(1))
func TestCheckCallExprFixedTypedMultiFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(m(), I(1))`, env,
		`multiple-value m() in single-value context`,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test FixedTyped(m(), is...)
func TestCheckCallExprFixedTypedMultiFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(m(), is...)`, env,
		`multiple-value m() in single-value context`,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to ft`,
	)

}

// Test FixedTyped(m(), e())
func TestCheckCallExprFixedTypedMultiFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(m(), e())`, env,
		`multiple-value m() in single-value context`,
		`e() used as value`,
	)

}

// Test FixedTyped(m(), s())
func TestCheckCallExprFixedTypedMultiFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(m(), s())`, env,
		`multiple-value m() in single-value context`,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(m(), m())
func TestCheckCallExprFixedTypedMultiFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(m(), m())`, env,
		`multiple-value m() in single-value context`,
		`multiple-value m() in single-value context`,
	)

}

// Test FixedTyped(m(), mt())
func TestCheckCallExprFixedTypedMultiFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(m(), mt())`, env,
		`multiple-value m() in single-value context`,
		`multiple-value mt() in single-value context`,
	)

}

// Test FixedTyped(mt())
func TestCheckCallExprFixedTypedMultiFuncMixedTypesX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(mt())`, env,
		`cannot use int as type eval.I in argument to ft`,
		`cannot use eval.I as type bool in argument to ft`,
	)

}

// Test FixedTyped(mt(), 1)
func TestCheckCallExprFixedTypedMultiFuncMixedTypesInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(mt(), 1)`, env,
		`multiple-value mt() in single-value context`,
		`cannot use 1 (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(mt(), 1.5)
func TestCheckCallExprFixedTypedMultiFuncMixedTypesFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(mt(), 1.5)`, env,
		`multiple-value mt() in single-value context`,
		`cannot use 1.5 (type float64) as type bool in function argument`,
	)

}

// Test FixedTyped(mt(), true)
func TestCheckCallExprFixedTypedMultiFuncMixedTypesBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(mt(), true)`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test FixedTyped(mt(), I(1))
func TestCheckCallExprFixedTypedMultiFuncMixedTypesIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(mt(), I(1))`, env,
		`multiple-value mt() in single-value context`,
		`cannot use eval.I(1) (type eval.I) as type bool in function argument`,
	)

}

// Test FixedTyped(mt(), is...)
func TestCheckCallExprFixedTypedMultiFuncMixedTypesInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(mt(), is...)`, env,
		`multiple-value mt() in single-value context`,
		`cannot use is (type []int) as type bool in function argument`,
		`invalid use of ... in call to ft`,
	)

}

// Test FixedTyped(mt(), e())
func TestCheckCallExprFixedTypedMultiFuncMixedTypesEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(mt(), e())`, env,
		`multiple-value mt() in single-value context`,
		`e() used as value`,
	)

}

// Test FixedTyped(mt(), s())
func TestCheckCallExprFixedTypedMultiFuncMixedTypesSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(mt(), s())`, env,
		`multiple-value mt() in single-value context`,
		`cannot use s() (type int) as type bool in function argument`,
	)

}

// Test FixedTyped(mt(), m())
func TestCheckCallExprFixedTypedMultiFuncMixedTypesMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(mt(), m())`, env,
		`multiple-value mt() in single-value context`,
		`multiple-value m() in single-value context`,
	)

}

// Test FixedTyped(mt(), mt())
func TestCheckCallExprFixedTypedMultiFuncMixedTypesMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `ft(mt(), mt())`, env,
		`multiple-value mt() in single-value context`,
		`multiple-value mt() in single-value context`,
	)

}

// Test Variadic1()
func TestCheckCallExprVariadic1XX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v1()`, env, reflect.TypeOf(v1()))
}

// Test Variadic1(1)
func TestCheckCallExprVariadic1XInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v1(1)`, env, reflect.TypeOf(v1(1)))
}

// Test Variadic1(1.5)
func TestCheckCallExprVariadic1XFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic1(true)
func TestCheckCallExprVariadic1XBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(true)`, env,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic1(I(1))
func TestCheckCallExprVariadic1XIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic1(is...)
func TestCheckCallExprVariadic1XInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v1(is...)`, env, reflect.TypeOf(v1(is...)))
}

// Test Variadic1(e())
func TestCheckCallExprVariadic1XEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(e())`, env,
		`e() used as value`,
	)

}

// Test Variadic1(s())
func TestCheckCallExprVariadic1XSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v1(s())`, env, reflect.TypeOf(v1(s())))
}

// Test Variadic1(m())
func TestCheckCallExprVariadic1XMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v1(m())`, env, reflect.TypeOf(v1(m())))
}

// Test Variadic1(mt())
func TestCheckCallExprVariadic1XMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(mt())`, env,
		`cannot use eval.I as type int in argument to v1`,
	)

}

// Test Variadic1(1)
func TestCheckCallExprVariadic1IntX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v1(1)`, env, reflect.TypeOf(v1(1)))
}

// Test Variadic1(1, 1)
func TestCheckCallExprVariadic1IntInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v1(1, 1)`, env, reflect.TypeOf(v1(1, 1)))
}

// Test Variadic1(1, 1.5)
func TestCheckCallExprVariadic1IntFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1, 1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic1(1, true)
func TestCheckCallExprVariadic1IntBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1, true)`, env,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic1(1, I(1))
func TestCheckCallExprVariadic1IntIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1, I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic1(1, is...)
func TestCheckCallExprVariadic1IntInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1, is...)`, env,
		`too many arguments in call to v1`,
	)

}

// Test Variadic1(1, e())
func TestCheckCallExprVariadic1IntEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1, e())`, env,
		`e() used as value`,
	)

}

// Test Variadic1(1, s())
func TestCheckCallExprVariadic1IntSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v1(1, s())`, env, reflect.TypeOf(v1(1, s())))
}

// Test Variadic1(1, m())
func TestCheckCallExprVariadic1IntMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1, m())`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test Variadic1(1, mt())
func TestCheckCallExprVariadic1IntMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1, mt())`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test Variadic1(1.5)
func TestCheckCallExprVariadic1FloatX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic1(1.5, 1)
func TestCheckCallExprVariadic1FloatInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1.5, 1)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic1(1.5, 1.5)
func TestCheckCallExprVariadic1FloatFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1.5, 1.5)`, env,
		`constant 1.5 truncated to integer`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic1(1.5, true)
func TestCheckCallExprVariadic1FloatBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1.5, true)`, env,
		`constant 1.5 truncated to integer`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic1(1.5, I(1))
func TestCheckCallExprVariadic1FloatIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1.5, I(1))`, env,
		`constant 1.5 truncated to integer`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic1(1.5, is...)
func TestCheckCallExprVariadic1FloatInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1.5, is...)`, env,
		`too many arguments in call to v1`,
	)

}

// Test Variadic1(1.5, e())
func TestCheckCallExprVariadic1FloatEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1.5, e())`, env,
		`e() used as value`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic1(1.5, s())
func TestCheckCallExprVariadic1FloatSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1.5, s())`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic1(1.5, m())
func TestCheckCallExprVariadic1FloatMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1.5, m())`, env,
		`multiple-value m() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic1(1.5, mt())
func TestCheckCallExprVariadic1FloatMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(1.5, mt())`, env,
		`multiple-value mt() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic1(true)
func TestCheckCallExprVariadic1BoolX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(true)`, env,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic1(true, 1)
func TestCheckCallExprVariadic1BoolInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(true, 1)`, env,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic1(true, 1.5)
func TestCheckCallExprVariadic1BoolFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(true, 1.5)`, env,
		`cannot use true (type bool) as type int in function argument`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic1(true, true)
func TestCheckCallExprVariadic1BoolBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(true, true)`, env,
		`cannot use true (type bool) as type int in function argument`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic1(true, I(1))
func TestCheckCallExprVariadic1BoolIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(true, I(1))`, env,
		`cannot use true (type bool) as type int in function argument`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic1(true, is...)
func TestCheckCallExprVariadic1BoolInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(true, is...)`, env,
		`too many arguments in call to v1`,
	)

}

// Test Variadic1(true, e())
func TestCheckCallExprVariadic1BoolEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(true, e())`, env,
		`e() used as value`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic1(true, s())
func TestCheckCallExprVariadic1BoolSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(true, s())`, env,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic1(true, m())
func TestCheckCallExprVariadic1BoolMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(true, m())`, env,
		`multiple-value m() in single-value context`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic1(true, mt())
func TestCheckCallExprVariadic1BoolMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(true, mt())`, env,
		`multiple-value mt() in single-value context`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic1(I(1))
func TestCheckCallExprVariadic1IntTypedX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic1(I(1), 1)
func TestCheckCallExprVariadic1IntTypedInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(I(1), 1)`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic1(I(1), 1.5)
func TestCheckCallExprVariadic1IntTypedFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(I(1), 1.5)`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic1(I(1), true)
func TestCheckCallExprVariadic1IntTypedBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(I(1), true)`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic1(I(1), I(1))
func TestCheckCallExprVariadic1IntTypedIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(I(1), I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic1(I(1), is...)
func TestCheckCallExprVariadic1IntTypedInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(I(1), is...)`, env,
		`too many arguments in call to v1`,
	)

}

// Test Variadic1(I(1), e())
func TestCheckCallExprVariadic1IntTypedEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(I(1), e())`, env,
		`e() used as value`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic1(I(1), s())
func TestCheckCallExprVariadic1IntTypedSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(I(1), s())`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic1(I(1), m())
func TestCheckCallExprVariadic1IntTypedMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(I(1), m())`, env,
		`multiple-value m() in single-value context`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic1(I(1), mt())
func TestCheckCallExprVariadic1IntTypedMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(I(1), mt())`, env,
		`multiple-value mt() in single-value context`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic1(is...)
func TestCheckCallExprVariadic1IntsX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v1(is...)`, env, reflect.TypeOf(v1(is...)))
}

// Test Variadic1(is..., 1)
func TestCheckCallExprVariadic1IntsInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic1(is..., 1.5)
func TestCheckCallExprVariadic1IntsFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic1(is..., true)
func TestCheckCallExprVariadic1IntsBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic1(is..., I(1))
func TestCheckCallExprVariadic1IntsIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic1(is..., is...)
func TestCheckCallExprVariadic1IntsInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic1(is..., e())
func TestCheckCallExprVariadic1IntsEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic1(is..., s())
func TestCheckCallExprVariadic1IntsSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic1(is..., m())
func TestCheckCallExprVariadic1IntsMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic1(is..., mt())
func TestCheckCallExprVariadic1IntsMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic1(e())
func TestCheckCallExprVariadic1EmptyFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(e())`, env,
		`e() used as value`,
	)

}

// Test Variadic1(e(), 1)
func TestCheckCallExprVariadic1EmptyFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(e(), 1)`, env,
		`e() used as value`,
	)

}

// Test Variadic1(e(), 1.5)
func TestCheckCallExprVariadic1EmptyFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(e(), 1.5)`, env,
		`e() used as value`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic1(e(), true)
func TestCheckCallExprVariadic1EmptyFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(e(), true)`, env,
		`e() used as value`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic1(e(), I(1))
func TestCheckCallExprVariadic1EmptyFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(e(), I(1))`, env,
		`e() used as value`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic1(e(), is...)
func TestCheckCallExprVariadic1EmptyFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(e(), is...)`, env,
		`e() used as value`,
		`too many arguments in call to v1`,
	)

}

// Test Variadic1(e(), e())
func TestCheckCallExprVariadic1EmptyFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(e(), e())`, env,
		`e() used as value`,
		`e() used as value`,
	)

}

// Test Variadic1(e(), s())
func TestCheckCallExprVariadic1EmptyFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(e(), s())`, env,
		`e() used as value`,
	)

}

// Test Variadic1(e(), m())
func TestCheckCallExprVariadic1EmptyFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(e(), m())`, env,
		`e() used as value`,
		`multiple-value m() in single-value context`,
	)

}

// Test Variadic1(e(), mt())
func TestCheckCallExprVariadic1EmptyFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(e(), mt())`, env,
		`e() used as value`,
		`multiple-value mt() in single-value context`,
	)

}

// Test Variadic1(s())
func TestCheckCallExprVariadic1SingleFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v1(s())`, env, reflect.TypeOf(v1(s())))
}

// Test Variadic1(s(), 1)
func TestCheckCallExprVariadic1SingleFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v1(s(), 1)`, env, reflect.TypeOf(v1(s(), 1)))
}

// Test Variadic1(s(), 1.5)
func TestCheckCallExprVariadic1SingleFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(s(), 1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic1(s(), true)
func TestCheckCallExprVariadic1SingleFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(s(), true)`, env,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic1(s(), I(1))
func TestCheckCallExprVariadic1SingleFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(s(), I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic1(s(), is...)
func TestCheckCallExprVariadic1SingleFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(s(), is...)`, env,
		`too many arguments in call to v1`,
	)

}

// Test Variadic1(s(), e())
func TestCheckCallExprVariadic1SingleFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(s(), e())`, env,
		`e() used as value`,
	)

}

// Test Variadic1(s(), s())
func TestCheckCallExprVariadic1SingleFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v1(s(), s())`, env, reflect.TypeOf(v1(s(), s())))
}

// Test Variadic1(s(), m())
func TestCheckCallExprVariadic1SingleFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(s(), m())`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test Variadic1(s(), mt())
func TestCheckCallExprVariadic1SingleFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(s(), mt())`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test Variadic1(m())
func TestCheckCallExprVariadic1MultiFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v1(m())`, env, reflect.TypeOf(v1(m())))
}

// Test Variadic1(m(), 1)
func TestCheckCallExprVariadic1MultiFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(m(), 1)`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test Variadic1(m(), 1.5)
func TestCheckCallExprVariadic1MultiFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(m(), 1.5)`, env,
		`multiple-value m() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic1(m(), true)
func TestCheckCallExprVariadic1MultiFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(m(), true)`, env,
		`multiple-value m() in single-value context`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic1(m(), I(1))
func TestCheckCallExprVariadic1MultiFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(m(), I(1))`, env,
		`multiple-value m() in single-value context`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic1(m(), is...)
func TestCheckCallExprVariadic1MultiFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(m(), is...)`, env,
		`multiple-value m() in single-value context`,
		`too many arguments in call to v1`,
	)

}

// Test Variadic1(m(), e())
func TestCheckCallExprVariadic1MultiFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(m(), e())`, env,
		`multiple-value m() in single-value context`,
		`e() used as value`,
	)

}

// Test Variadic1(m(), s())
func TestCheckCallExprVariadic1MultiFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(m(), s())`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test Variadic1(m(), m())
func TestCheckCallExprVariadic1MultiFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(m(), m())`, env,
		`multiple-value m() in single-value context`,
		`multiple-value m() in single-value context`,
	)

}

// Test Variadic1(m(), mt())
func TestCheckCallExprVariadic1MultiFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(m(), mt())`, env,
		`multiple-value m() in single-value context`,
		`multiple-value mt() in single-value context`,
	)

}

// Test Variadic1(mt())
func TestCheckCallExprVariadic1MultiFuncMixedTypesX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(mt())`, env,
		`cannot use eval.I as type int in argument to v1`,
	)

}

// Test Variadic1(mt(), 1)
func TestCheckCallExprVariadic1MultiFuncMixedTypesInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(mt(), 1)`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test Variadic1(mt(), 1.5)
func TestCheckCallExprVariadic1MultiFuncMixedTypesFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(mt(), 1.5)`, env,
		`multiple-value mt() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic1(mt(), true)
func TestCheckCallExprVariadic1MultiFuncMixedTypesBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(mt(), true)`, env,
		`multiple-value mt() in single-value context`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic1(mt(), I(1))
func TestCheckCallExprVariadic1MultiFuncMixedTypesIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(mt(), I(1))`, env,
		`multiple-value mt() in single-value context`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic1(mt(), is...)
func TestCheckCallExprVariadic1MultiFuncMixedTypesInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(mt(), is...)`, env,
		`multiple-value mt() in single-value context`,
		`too many arguments in call to v1`,
	)

}

// Test Variadic1(mt(), e())
func TestCheckCallExprVariadic1MultiFuncMixedTypesEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(mt(), e())`, env,
		`multiple-value mt() in single-value context`,
		`e() used as value`,
	)

}

// Test Variadic1(mt(), s())
func TestCheckCallExprVariadic1MultiFuncMixedTypesSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(mt(), s())`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test Variadic1(mt(), m())
func TestCheckCallExprVariadic1MultiFuncMixedTypesMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(mt(), m())`, env,
		`multiple-value mt() in single-value context`,
		`multiple-value m() in single-value context`,
	)

}

// Test Variadic1(mt(), mt())
func TestCheckCallExprVariadic1MultiFuncMixedTypesMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v1(mt(), mt())`, env,
		`multiple-value mt() in single-value context`,
		`multiple-value mt() in single-value context`,
	)

}

// Test Variadic2()
func TestCheckCallExprVariadic2XX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2()`, env,
		`not enough arguments in call to v2`,
	)

}

// Test Variadic2(1)
func TestCheckCallExprVariadic2XInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v2(1)`, env, reflect.TypeOf(v2(1)))
}

// Test Variadic2(1.5)
func TestCheckCallExprVariadic2XFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(true)
func TestCheckCallExprVariadic2XBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(true)`, env,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(I(1))
func TestCheckCallExprVariadic2XIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(is...)
func TestCheckCallExprVariadic2XInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(is...)`, env,
		`cannot use is (type []int) as type int in function argument`,
		`not enough arguments in call to v2`,
	)

}

// Test Variadic2(e())
func TestCheckCallExprVariadic2XEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(e())`, env,
		`e() used as value`,
	)

}

// Test Variadic2(s())
func TestCheckCallExprVariadic2XSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v2(s())`, env, reflect.TypeOf(v2(s())))
}

// Test Variadic2(m())
func TestCheckCallExprVariadic2XMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v2(m())`, env, reflect.TypeOf(v2(m())))
}

// Test Variadic2(mt())
func TestCheckCallExprVariadic2XMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(mt())`, env,
		`cannot use eval.I as type int in argument to v2`,
	)

}

// Test Variadic2(1)
func TestCheckCallExprVariadic2IntX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v2(1)`, env, reflect.TypeOf(v2(1)))
}

// Test Variadic2(1, 1)
func TestCheckCallExprVariadic2IntInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v2(1, 1)`, env, reflect.TypeOf(v2(1, 1)))
}

// Test Variadic2(1, 1.5)
func TestCheckCallExprVariadic2IntFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1, 1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(1, true)
func TestCheckCallExprVariadic2IntBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1, true)`, env,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(1, I(1))
func TestCheckCallExprVariadic2IntIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1, I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(1, is...)
func TestCheckCallExprVariadic2IntInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v2(1, is...)`, env, reflect.TypeOf(v2(1, is...)))
}

// Test Variadic2(1, e())
func TestCheckCallExprVariadic2IntEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1, e())`, env,
		`e() used as value`,
	)

}

// Test Variadic2(1, s())
func TestCheckCallExprVariadic2IntSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v2(1, s())`, env, reflect.TypeOf(v2(1, s())))
}

// Test Variadic2(1, m())
func TestCheckCallExprVariadic2IntMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1, m())`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test Variadic2(1, mt())
func TestCheckCallExprVariadic2IntMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1, mt())`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test Variadic2(1.5)
func TestCheckCallExprVariadic2FloatX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(1.5, 1)
func TestCheckCallExprVariadic2FloatInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1.5, 1)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(1.5, 1.5)
func TestCheckCallExprVariadic2FloatFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1.5, 1.5)`, env,
		`constant 1.5 truncated to integer`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(1.5, true)
func TestCheckCallExprVariadic2FloatBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1.5, true)`, env,
		`constant 1.5 truncated to integer`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(1.5, I(1))
func TestCheckCallExprVariadic2FloatIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1.5, I(1))`, env,
		`constant 1.5 truncated to integer`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(1.5, is...)
func TestCheckCallExprVariadic2FloatInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1.5, is...)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(1.5, e())
func TestCheckCallExprVariadic2FloatEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1.5, e())`, env,
		`e() used as value`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(1.5, s())
func TestCheckCallExprVariadic2FloatSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1.5, s())`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(1.5, m())
func TestCheckCallExprVariadic2FloatMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1.5, m())`, env,
		`multiple-value m() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(1.5, mt())
func TestCheckCallExprVariadic2FloatMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(1.5, mt())`, env,
		`multiple-value mt() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(true)
func TestCheckCallExprVariadic2BoolX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(true)`, env,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(true, 1)
func TestCheckCallExprVariadic2BoolInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(true, 1)`, env,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(true, 1.5)
func TestCheckCallExprVariadic2BoolFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(true, 1.5)`, env,
		`cannot use true (type bool) as type int in function argument`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(true, true)
func TestCheckCallExprVariadic2BoolBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(true, true)`, env,
		`cannot use true (type bool) as type int in function argument`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(true, I(1))
func TestCheckCallExprVariadic2BoolIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(true, I(1))`, env,
		`cannot use true (type bool) as type int in function argument`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(true, is...)
func TestCheckCallExprVariadic2BoolInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(true, is...)`, env,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(true, e())
func TestCheckCallExprVariadic2BoolEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(true, e())`, env,
		`e() used as value`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(true, s())
func TestCheckCallExprVariadic2BoolSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(true, s())`, env,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(true, m())
func TestCheckCallExprVariadic2BoolMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(true, m())`, env,
		`multiple-value m() in single-value context`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(true, mt())
func TestCheckCallExprVariadic2BoolMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(true, mt())`, env,
		`multiple-value mt() in single-value context`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(I(1))
func TestCheckCallExprVariadic2IntTypedX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(I(1), 1)
func TestCheckCallExprVariadic2IntTypedInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(I(1), 1)`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(I(1), 1.5)
func TestCheckCallExprVariadic2IntTypedFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(I(1), 1.5)`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(I(1), true)
func TestCheckCallExprVariadic2IntTypedBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(I(1), true)`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(I(1), I(1))
func TestCheckCallExprVariadic2IntTypedIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(I(1), I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(I(1), is...)
func TestCheckCallExprVariadic2IntTypedInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(I(1), is...)`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(I(1), e())
func TestCheckCallExprVariadic2IntTypedEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(I(1), e())`, env,
		`e() used as value`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(I(1), s())
func TestCheckCallExprVariadic2IntTypedSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(I(1), s())`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(I(1), m())
func TestCheckCallExprVariadic2IntTypedMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(I(1), m())`, env,
		`multiple-value m() in single-value context`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(I(1), mt())
func TestCheckCallExprVariadic2IntTypedMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(I(1), mt())`, env,
		`multiple-value mt() in single-value context`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(is...)
func TestCheckCallExprVariadic2IntsX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(is...)`, env,
		`cannot use is (type []int) as type int in function argument`,
		`not enough arguments in call to v2`,
	)

}

// Test Variadic2(is..., 1)
func TestCheckCallExprVariadic2IntsInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic2(is..., 1.5)
func TestCheckCallExprVariadic2IntsFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic2(is..., true)
func TestCheckCallExprVariadic2IntsBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic2(is..., I(1))
func TestCheckCallExprVariadic2IntsIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic2(is..., is...)
func TestCheckCallExprVariadic2IntsInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic2(is..., e())
func TestCheckCallExprVariadic2IntsEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic2(is..., s())
func TestCheckCallExprVariadic2IntsSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic2(is..., m())
func TestCheckCallExprVariadic2IntsMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic2(is..., mt())
func TestCheckCallExprVariadic2IntsMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test Variadic2(e())
func TestCheckCallExprVariadic2EmptyFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(e())`, env,
		`e() used as value`,
	)

}

// Test Variadic2(e(), 1)
func TestCheckCallExprVariadic2EmptyFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(e(), 1)`, env,
		`e() used as value`,
	)

}

// Test Variadic2(e(), 1.5)
func TestCheckCallExprVariadic2EmptyFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(e(), 1.5)`, env,
		`e() used as value`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(e(), true)
func TestCheckCallExprVariadic2EmptyFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(e(), true)`, env,
		`e() used as value`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(e(), I(1))
func TestCheckCallExprVariadic2EmptyFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(e(), I(1))`, env,
		`e() used as value`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(e(), is...)
func TestCheckCallExprVariadic2EmptyFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(e(), is...)`, env,
		`e() used as value`,
	)

}

// Test Variadic2(e(), e())
func TestCheckCallExprVariadic2EmptyFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(e(), e())`, env,
		`e() used as value`,
		`e() used as value`,
	)

}

// Test Variadic2(e(), s())
func TestCheckCallExprVariadic2EmptyFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(e(), s())`, env,
		`e() used as value`,
	)

}

// Test Variadic2(e(), m())
func TestCheckCallExprVariadic2EmptyFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(e(), m())`, env,
		`e() used as value`,
		`multiple-value m() in single-value context`,
	)

}

// Test Variadic2(e(), mt())
func TestCheckCallExprVariadic2EmptyFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(e(), mt())`, env,
		`e() used as value`,
		`multiple-value mt() in single-value context`,
	)

}

// Test Variadic2(s())
func TestCheckCallExprVariadic2SingleFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v2(s())`, env, reflect.TypeOf(v2(s())))
}

// Test Variadic2(s(), 1)
func TestCheckCallExprVariadic2SingleFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v2(s(), 1)`, env, reflect.TypeOf(v2(s(), 1)))
}

// Test Variadic2(s(), 1.5)
func TestCheckCallExprVariadic2SingleFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(s(), 1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(s(), true)
func TestCheckCallExprVariadic2SingleFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(s(), true)`, env,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(s(), I(1))
func TestCheckCallExprVariadic2SingleFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(s(), I(1))`, env,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(s(), is...)
func TestCheckCallExprVariadic2SingleFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v2(s(), is...)`, env, reflect.TypeOf(v2(s(), is...)))
}

// Test Variadic2(s(), e())
func TestCheckCallExprVariadic2SingleFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(s(), e())`, env,
		`e() used as value`,
	)

}

// Test Variadic2(s(), s())
func TestCheckCallExprVariadic2SingleFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v2(s(), s())`, env, reflect.TypeOf(v2(s(), s())))
}

// Test Variadic2(s(), m())
func TestCheckCallExprVariadic2SingleFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(s(), m())`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test Variadic2(s(), mt())
func TestCheckCallExprVariadic2SingleFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(s(), mt())`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test Variadic2(m())
func TestCheckCallExprVariadic2MultiFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `v2(m())`, env, reflect.TypeOf(v2(m())))
}

// Test Variadic2(m(), 1)
func TestCheckCallExprVariadic2MultiFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(m(), 1)`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test Variadic2(m(), 1.5)
func TestCheckCallExprVariadic2MultiFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(m(), 1.5)`, env,
		`multiple-value m() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(m(), true)
func TestCheckCallExprVariadic2MultiFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(m(), true)`, env,
		`multiple-value m() in single-value context`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(m(), I(1))
func TestCheckCallExprVariadic2MultiFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(m(), I(1))`, env,
		`multiple-value m() in single-value context`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(m(), is...)
func TestCheckCallExprVariadic2MultiFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(m(), is...)`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test Variadic2(m(), e())
func TestCheckCallExprVariadic2MultiFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(m(), e())`, env,
		`multiple-value m() in single-value context`,
		`e() used as value`,
	)

}

// Test Variadic2(m(), s())
func TestCheckCallExprVariadic2MultiFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(m(), s())`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test Variadic2(m(), m())
func TestCheckCallExprVariadic2MultiFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(m(), m())`, env,
		`multiple-value m() in single-value context`,
		`multiple-value m() in single-value context`,
	)

}

// Test Variadic2(m(), mt())
func TestCheckCallExprVariadic2MultiFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(m(), mt())`, env,
		`multiple-value m() in single-value context`,
		`multiple-value mt() in single-value context`,
	)

}

// Test Variadic2(mt())
func TestCheckCallExprVariadic2MultiFuncMixedTypesX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(mt())`, env,
		`cannot use eval.I as type int in argument to v2`,
	)

}

// Test Variadic2(mt(), 1)
func TestCheckCallExprVariadic2MultiFuncMixedTypesInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(mt(), 1)`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test Variadic2(mt(), 1.5)
func TestCheckCallExprVariadic2MultiFuncMixedTypesFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(mt(), 1.5)`, env,
		`multiple-value mt() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test Variadic2(mt(), true)
func TestCheckCallExprVariadic2MultiFuncMixedTypesBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(mt(), true)`, env,
		`multiple-value mt() in single-value context`,
		`cannot use true (type bool) as type int in function argument`,
	)

}

// Test Variadic2(mt(), I(1))
func TestCheckCallExprVariadic2MultiFuncMixedTypesIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(mt(), I(1))`, env,
		`multiple-value mt() in single-value context`,
		`cannot use eval.I(1) (type eval.I) as type int in function argument`,
	)

}

// Test Variadic2(mt(), is...)
func TestCheckCallExprVariadic2MultiFuncMixedTypesInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(mt(), is...)`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test Variadic2(mt(), e())
func TestCheckCallExprVariadic2MultiFuncMixedTypesEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(mt(), e())`, env,
		`multiple-value mt() in single-value context`,
		`e() used as value`,
	)

}

// Test Variadic2(mt(), s())
func TestCheckCallExprVariadic2MultiFuncMixedTypesSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(mt(), s())`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test Variadic2(mt(), m())
func TestCheckCallExprVariadic2MultiFuncMixedTypesMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(mt(), m())`, env,
		`multiple-value mt() in single-value context`,
		`multiple-value m() in single-value context`,
	)

}

// Test Variadic2(mt(), mt())
func TestCheckCallExprVariadic2MultiFuncMixedTypesMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `v2(mt(), mt())`, env,
		`multiple-value mt() in single-value context`,
		`multiple-value mt() in single-value context`,
	)

}

// Test VariadicTyped()
func TestCheckCallExprVariadicTypedXX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt()`, env,
		`not enough arguments in call to vt`,
	)

}

// Test VariadicTyped(1)
func TestCheckCallExprVariadicTypedXInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `vt(1)`, env, reflect.TypeOf(vt(1)))
}

// Test VariadicTyped(1.5)
func TestCheckCallExprVariadicTypedXFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test VariadicTyped(true)
func TestCheckCallExprVariadicTypedXBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(true)`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(I(1))
func TestCheckCallExprVariadicTypedXIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `vt(I(1))`, env, reflect.TypeOf(vt(I(1))))
}

// Test VariadicTyped(is...)
func TestCheckCallExprVariadicTypedXInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(is...)`, env,
		`cannot use is (type []int) as type eval.I in function argument`,
		`not enough arguments in call to vt`,
	)

}

// Test VariadicTyped(e())
func TestCheckCallExprVariadicTypedXEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(e())`, env,
		`e() used as value`,
	)

}

// Test VariadicTyped(s())
func TestCheckCallExprVariadicTypedXSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(s())`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(m())
func TestCheckCallExprVariadicTypedXMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(m())`, env,
		`cannot use int as type eval.I in argument to vt`,
		`cannot use int as type eval.I in argument to vt`,
	)

}

// Test VariadicTyped(mt())
func TestCheckCallExprVariadicTypedXMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(mt())`, env,
		`cannot use int as type eval.I in argument to vt`,
	)

}

// Test VariadicTyped(1)
func TestCheckCallExprVariadicTypedIntX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `vt(1)`, env, reflect.TypeOf(vt(1)))
}

// Test VariadicTyped(1, 1)
func TestCheckCallExprVariadicTypedIntInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `vt(1, 1)`, env, reflect.TypeOf(vt(1, 1)))
}

// Test VariadicTyped(1, 1.5)
func TestCheckCallExprVariadicTypedIntFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1, 1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test VariadicTyped(1, true)
func TestCheckCallExprVariadicTypedIntBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1, true)`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(1, I(1))
func TestCheckCallExprVariadicTypedIntIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `vt(1, I(1))`, env, reflect.TypeOf(vt(1, I(1))))
}

// Test VariadicTyped(1, is...)
func TestCheckCallExprVariadicTypedIntInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1, is...)`, env,
		`cannot use is (type []int) as type []eval.I in function argument`,
	)

}

// Test VariadicTyped(1, e())
func TestCheckCallExprVariadicTypedIntEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1, e())`, env,
		`e() used as value`,
	)

}

// Test VariadicTyped(1, s())
func TestCheckCallExprVariadicTypedIntSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1, s())`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(1, m())
func TestCheckCallExprVariadicTypedIntMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1, m())`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test VariadicTyped(1, mt())
func TestCheckCallExprVariadicTypedIntMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1, mt())`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test VariadicTyped(1.5)
func TestCheckCallExprVariadicTypedFloatX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test VariadicTyped(1.5, 1)
func TestCheckCallExprVariadicTypedFloatInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1.5, 1)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test VariadicTyped(1.5, 1.5)
func TestCheckCallExprVariadicTypedFloatFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1.5, 1.5)`, env,
		`constant 1.5 truncated to integer`,
		`constant 1.5 truncated to integer`,
	)

}

// Test VariadicTyped(1.5, true)
func TestCheckCallExprVariadicTypedFloatBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1.5, true)`, env,
		`constant 1.5 truncated to integer`,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(1.5, I(1))
func TestCheckCallExprVariadicTypedFloatIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1.5, I(1))`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test VariadicTyped(1.5, is...)
func TestCheckCallExprVariadicTypedFloatInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1.5, is...)`, env,
		`constant 1.5 truncated to integer`,
		`cannot use is (type []int) as type []eval.I in function argument`,
	)

}

// Test VariadicTyped(1.5, e())
func TestCheckCallExprVariadicTypedFloatEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1.5, e())`, env,
		`e() used as value`,
		`constant 1.5 truncated to integer`,
	)

}

// Test VariadicTyped(1.5, s())
func TestCheckCallExprVariadicTypedFloatSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1.5, s())`, env,
		`constant 1.5 truncated to integer`,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(1.5, m())
func TestCheckCallExprVariadicTypedFloatMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1.5, m())`, env,
		`multiple-value m() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test VariadicTyped(1.5, mt())
func TestCheckCallExprVariadicTypedFloatMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(1.5, mt())`, env,
		`multiple-value mt() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test VariadicTyped(true)
func TestCheckCallExprVariadicTypedBoolX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(true)`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(true, 1)
func TestCheckCallExprVariadicTypedBoolInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(true, 1)`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(true, 1.5)
func TestCheckCallExprVariadicTypedBoolFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(true, 1.5)`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
		`constant 1.5 truncated to integer`,
	)

}

// Test VariadicTyped(true, true)
func TestCheckCallExprVariadicTypedBoolBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(true, true)`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(true, I(1))
func TestCheckCallExprVariadicTypedBoolIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(true, I(1))`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(true, is...)
func TestCheckCallExprVariadicTypedBoolInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(true, is...)`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
		`cannot use is (type []int) as type []eval.I in function argument`,
	)

}

// Test VariadicTyped(true, e())
func TestCheckCallExprVariadicTypedBoolEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(true, e())`, env,
		`e() used as value`,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(true, s())
func TestCheckCallExprVariadicTypedBoolSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(true, s())`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(true, m())
func TestCheckCallExprVariadicTypedBoolMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(true, m())`, env,
		`multiple-value m() in single-value context`,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(true, mt())
func TestCheckCallExprVariadicTypedBoolMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(true, mt())`, env,
		`multiple-value mt() in single-value context`,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(I(1))
func TestCheckCallExprVariadicTypedIntTypedX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `vt(I(1))`, env, reflect.TypeOf(vt(I(1))))
}

// Test VariadicTyped(I(1), 1)
func TestCheckCallExprVariadicTypedIntTypedInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `vt(I(1), 1)`, env, reflect.TypeOf(vt(I(1), 1)))
}

// Test VariadicTyped(I(1), 1.5)
func TestCheckCallExprVariadicTypedIntTypedFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(I(1), 1.5)`, env,
		`constant 1.5 truncated to integer`,
	)

}

// Test VariadicTyped(I(1), true)
func TestCheckCallExprVariadicTypedIntTypedBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(I(1), true)`, env,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(I(1), I(1))
func TestCheckCallExprVariadicTypedIntTypedIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectType(t, `vt(I(1), I(1))`, env, reflect.TypeOf(vt(I(1), I(1))))
}

// Test VariadicTyped(I(1), is...)
func TestCheckCallExprVariadicTypedIntTypedInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(I(1), is...)`, env,
		`cannot use is (type []int) as type []eval.I in function argument`,
	)

}

// Test VariadicTyped(I(1), e())
func TestCheckCallExprVariadicTypedIntTypedEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(I(1), e())`, env,
		`e() used as value`,
	)

}

// Test VariadicTyped(I(1), s())
func TestCheckCallExprVariadicTypedIntTypedSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(I(1), s())`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(I(1), m())
func TestCheckCallExprVariadicTypedIntTypedMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(I(1), m())`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test VariadicTyped(I(1), mt())
func TestCheckCallExprVariadicTypedIntTypedMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(I(1), mt())`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test VariadicTyped(is...)
func TestCheckCallExprVariadicTypedIntsX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(is...)`, env,
		`cannot use is (type []int) as type eval.I in function argument`,
		`not enough arguments in call to vt`,
	)

}

// Test VariadicTyped(is..., 1)
func TestCheckCallExprVariadicTypedIntsInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test VariadicTyped(is..., 1.5)
func TestCheckCallExprVariadicTypedIntsFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test VariadicTyped(is..., true)
func TestCheckCallExprVariadicTypedIntsBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test VariadicTyped(is..., I(1))
func TestCheckCallExprVariadicTypedIntsIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test VariadicTyped(is..., is...)
func TestCheckCallExprVariadicTypedIntsInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test VariadicTyped(is..., e())
func TestCheckCallExprVariadicTypedIntsEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test VariadicTyped(is..., s())
func TestCheckCallExprVariadicTypedIntsSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test VariadicTyped(is..., m())
func TestCheckCallExprVariadicTypedIntsMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test VariadicTyped(is..., mt())
func TestCheckCallExprVariadicTypedIntsMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)

	_ = env

}

// Test VariadicTyped(e())
func TestCheckCallExprVariadicTypedEmptyFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(e())`, env,
		`e() used as value`,
	)

}

// Test VariadicTyped(e(), 1)
func TestCheckCallExprVariadicTypedEmptyFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(e(), 1)`, env,
		`e() used as value`,
	)

}

// Test VariadicTyped(e(), 1.5)
func TestCheckCallExprVariadicTypedEmptyFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(e(), 1.5)`, env,
		`e() used as value`,
		`constant 1.5 truncated to integer`,
	)

}

// Test VariadicTyped(e(), true)
func TestCheckCallExprVariadicTypedEmptyFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(e(), true)`, env,
		`e() used as value`,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(e(), I(1))
func TestCheckCallExprVariadicTypedEmptyFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(e(), I(1))`, env,
		`e() used as value`,
	)

}

// Test VariadicTyped(e(), is...)
func TestCheckCallExprVariadicTypedEmptyFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(e(), is...)`, env,
		`e() used as value`,
		`cannot use is (type []int) as type []eval.I in function argument`,
	)

}

// Test VariadicTyped(e(), e())
func TestCheckCallExprVariadicTypedEmptyFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(e(), e())`, env,
		`e() used as value`,
		`e() used as value`,
	)

}

// Test VariadicTyped(e(), s())
func TestCheckCallExprVariadicTypedEmptyFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(e(), s())`, env,
		`e() used as value`,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(e(), m())
func TestCheckCallExprVariadicTypedEmptyFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(e(), m())`, env,
		`e() used as value`,
		`multiple-value m() in single-value context`,
	)

}

// Test VariadicTyped(e(), mt())
func TestCheckCallExprVariadicTypedEmptyFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(e(), mt())`, env,
		`e() used as value`,
		`multiple-value mt() in single-value context`,
	)

}

// Test VariadicTyped(s())
func TestCheckCallExprVariadicTypedSingleFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(s())`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(s(), 1)
func TestCheckCallExprVariadicTypedSingleFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(s(), 1)`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(s(), 1.5)
func TestCheckCallExprVariadicTypedSingleFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(s(), 1.5)`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
		`constant 1.5 truncated to integer`,
	)

}

// Test VariadicTyped(s(), true)
func TestCheckCallExprVariadicTypedSingleFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(s(), true)`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(s(), I(1))
func TestCheckCallExprVariadicTypedSingleFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(s(), I(1))`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(s(), is...)
func TestCheckCallExprVariadicTypedSingleFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(s(), is...)`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
		`cannot use is (type []int) as type []eval.I in function argument`,
	)

}

// Test VariadicTyped(s(), e())
func TestCheckCallExprVariadicTypedSingleFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(s(), e())`, env,
		`e() used as value`,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(s(), s())
func TestCheckCallExprVariadicTypedSingleFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(s(), s())`, env,
		`cannot use s() (type int) as type eval.I in function argument`,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(s(), m())
func TestCheckCallExprVariadicTypedSingleFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(s(), m())`, env,
		`multiple-value m() in single-value context`,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(s(), mt())
func TestCheckCallExprVariadicTypedSingleFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(s(), mt())`, env,
		`multiple-value mt() in single-value context`,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(m())
func TestCheckCallExprVariadicTypedMultiFuncX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(m())`, env,
		`cannot use int as type eval.I in argument to vt`,
		`cannot use int as type eval.I in argument to vt`,
	)

}

// Test VariadicTyped(m(), 1)
func TestCheckCallExprVariadicTypedMultiFuncInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(m(), 1)`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test VariadicTyped(m(), 1.5)
func TestCheckCallExprVariadicTypedMultiFuncFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(m(), 1.5)`, env,
		`multiple-value m() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test VariadicTyped(m(), true)
func TestCheckCallExprVariadicTypedMultiFuncBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(m(), true)`, env,
		`multiple-value m() in single-value context`,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(m(), I(1))
func TestCheckCallExprVariadicTypedMultiFuncIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(m(), I(1))`, env,
		`multiple-value m() in single-value context`,
	)

}

// Test VariadicTyped(m(), is...)
func TestCheckCallExprVariadicTypedMultiFuncInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(m(), is...)`, env,
		`multiple-value m() in single-value context`,
		`cannot use is (type []int) as type []eval.I in function argument`,
	)

}

// Test VariadicTyped(m(), e())
func TestCheckCallExprVariadicTypedMultiFuncEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(m(), e())`, env,
		`multiple-value m() in single-value context`,
		`e() used as value`,
	)

}

// Test VariadicTyped(m(), s())
func TestCheckCallExprVariadicTypedMultiFuncSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(m(), s())`, env,
		`multiple-value m() in single-value context`,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(m(), m())
func TestCheckCallExprVariadicTypedMultiFuncMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(m(), m())`, env,
		`multiple-value m() in single-value context`,
		`multiple-value m() in single-value context`,
	)

}

// Test VariadicTyped(m(), mt())
func TestCheckCallExprVariadicTypedMultiFuncMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(m(), mt())`, env,
		`multiple-value m() in single-value context`,
		`multiple-value mt() in single-value context`,
	)

}

// Test VariadicTyped(mt())
func TestCheckCallExprVariadicTypedMultiFuncMixedTypesX(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(mt())`, env,
		`cannot use int as type eval.I in argument to vt`,
	)

}

// Test VariadicTyped(mt(), 1)
func TestCheckCallExprVariadicTypedMultiFuncMixedTypesInt(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(mt(), 1)`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test VariadicTyped(mt(), 1.5)
func TestCheckCallExprVariadicTypedMultiFuncMixedTypesFloat(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(mt(), 1.5)`, env,
		`multiple-value mt() in single-value context`,
		`constant 1.5 truncated to integer`,
	)

}

// Test VariadicTyped(mt(), true)
func TestCheckCallExprVariadicTypedMultiFuncMixedTypesBool(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(mt(), true)`, env,
		`multiple-value mt() in single-value context`,
		`cannot use true (type bool) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(mt(), I(1))
func TestCheckCallExprVariadicTypedMultiFuncMixedTypesIntTyped(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(mt(), I(1))`, env,
		`multiple-value mt() in single-value context`,
	)

}

// Test VariadicTyped(mt(), is...)
func TestCheckCallExprVariadicTypedMultiFuncMixedTypesInts(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(mt(), is...)`, env,
		`multiple-value mt() in single-value context`,
		`cannot use is (type []int) as type []eval.I in function argument`,
	)

}

// Test VariadicTyped(mt(), e())
func TestCheckCallExprVariadicTypedMultiFuncMixedTypesEmptyFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(mt(), e())`, env,
		`multiple-value mt() in single-value context`,
		`e() used as value`,
	)

}

// Test VariadicTyped(mt(), s())
func TestCheckCallExprVariadicTypedMultiFuncMixedTypesSingleFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(mt(), s())`, env,
		`multiple-value mt() in single-value context`,
		`cannot use s() (type int) as type eval.I in function argument`,
	)

}

// Test VariadicTyped(mt(), m())
func TestCheckCallExprVariadicTypedMultiFuncMixedTypesMultiFunc(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(mt(), m())`, env,
		`multiple-value mt() in single-value context`,
		`multiple-value m() in single-value context`,
	)

}

// Test VariadicTyped(mt(), mt())
func TestCheckCallExprVariadicTypedMultiFuncMixedTypesMultiFuncMixedTypes(t *testing.T) {
	type I int
	is := []int{1, 2, 3}
	f := func(int, bool) int { return 1 }
	ft := func(I, bool) int { return 1 }
	v1 := func(...int) int { return 1 }
	v2 := func(int, ...int) int { return 1 }
	vt := func(I, ...I) int { return 1 }
	e := func() {}
	s := func() int { return 1}
	m := func() (int, int) { return 1, 1 }
	mt := func() (int, I) { return 1, 1 }


	env := MakeSimpleEnv()
	env.Types["I"] = reflect.TypeOf(I(0))
	env.Vars["is"] = reflect.ValueOf(&is)
	env.Funcs["f"] = reflect.ValueOf(f)
	env.Funcs["ft"] = reflect.ValueOf(ft)
	env.Funcs["v1"] = reflect.ValueOf(v1)
	env.Funcs["v2"] = reflect.ValueOf(v2)
	env.Funcs["vt"] = reflect.ValueOf(vt)
	env.Funcs["e"] = reflect.ValueOf(e)
	env.Funcs["s"] = reflect.ValueOf(s)
	env.Funcs["m"] = reflect.ValueOf(m)
	env.Funcs["mt"] = reflect.ValueOf(mt)


	expectCheckError(t, `vt(mt(), mt())`, env,
		`multiple-value mt() in single-value context`,
		`multiple-value mt() in single-value context`,
	)

}
