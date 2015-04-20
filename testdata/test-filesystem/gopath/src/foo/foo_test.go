package foo

import "testing"

func TestFoo(t *testing.T) {
	_ = "breakpoint"
	t.Fail()
}
