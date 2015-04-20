package testpkg2

import (
	"testing"

	"foo"
	"foo/subfoo"
)

func TestFoos(t *testing.T) {
	_ = "breakpoint"
	foo.Foo()
	subfoo.SubFoo()
}
