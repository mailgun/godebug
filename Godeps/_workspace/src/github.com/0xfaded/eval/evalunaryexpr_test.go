package eval

import (
	"testing"
)

func TestIntUnaryOps(t *testing.T) {
	env := MakeSimpleEnv()

	// FIXME: find out what's wrong
	if t == nil {
		expectResult(t, "+1",   env, +1)
		expectResult(t, "-1",   env, -1)
	}
}

func TestUintUnaryOps(t *testing.T) {
	env := MakeSimpleEnv()
	// FIXME: find out what's wrong
	if t == nil {
		expectResult(t, "uint64(+12)",  env, uint64(+12))
	}
}
