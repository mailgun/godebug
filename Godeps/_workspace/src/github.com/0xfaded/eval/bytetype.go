package eval

import (
	"reflect"
)

type Byte struct {
	reflect.Type
}

func (Byte) Name() string {
	return "byte"
}

func (Byte) String() string {
	return "byte"
}

var ByteType = Byte{reflect.TypeOf(byte(0))}
