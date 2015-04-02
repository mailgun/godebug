package subdir

import (
	"foo"
	"testing"

	"github.com/mailgun/godebug/lib"
)

func TestIt(t *testing.T) {
	godebug.SetTrace()
	foo.Foo()
	_ = "finishing test"
}
