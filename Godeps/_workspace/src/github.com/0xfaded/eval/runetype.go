package eval

import (
	"reflect"
)

type Rune struct {
	reflect.Type
}

func (Rune) Name() string {
	return "rune"
}

func (Rune) String() string {
	return "rune"
}

var RuneType = Rune{reflect.TypeOf(rune(0))}
