package testpkg2

import (
	"testing"

	"github.com/mailgun/godebug/lib"

	"foo"
	"foo/subfoo"
)

func TestFoos(t *testing.T) {
	godebug.SetTrace()
	foo.Foo()
	subfoo.SubFoo()
}
