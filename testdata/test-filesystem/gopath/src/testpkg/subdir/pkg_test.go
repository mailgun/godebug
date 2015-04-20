package subdir

import (
	"foo"
	"testing"
)

func TestIt(t *testing.T) {
	_ = "breakpoint"
	foo.Foo()
	_ = "finishing test"
}
