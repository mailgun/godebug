package eval

import (
	"fmt"
	"reflect"
)

type PanicUser reflect.Value
type PanicDivideByZero struct {}
type PanicInvalidDereference struct {}
type PanicIndexOutOfBounds struct {}
type PanicSliceOutOfBounds struct {}
type PanicInterfaceConversion struct {
	// type of type assert operand
	xT reflect.Type

	// target assertion type
	aT reflect.Type

	// the dynamic type of operand. nil for interface to interface assertions
	dynamicT reflect.Type
}
type PanicUncomparableType struct {
	dynamicT reflect.Type
}
type PanicUnhashableType struct {
	dynamicT reflect.Type
}

func (p PanicUser) Error() string {
	return fmt.Sprint(reflect.Value(p).Interface())
}

func (err PanicDivideByZero) Error() string {
        return "runtime error: integer divide by zero"
}

func (err PanicInvalidDereference) Error() string {
	return "runtime error: invalid memory address or nil pointer dereference"
}

func (err PanicIndexOutOfBounds) Error() string {
        return "runtime error: index out of range"
}

func (err PanicSliceOutOfBounds) Error() string {
        return "runtime error: slice bounds out of range"
}

func (err PanicInterfaceConversion) Error() string {
	if err.xT == nil {
		return fmt.Sprintf("interface conversion: nil is not %v", err.aT)
	} else if err.dynamicT == nil {
		var missingMethod string
		numMethod := err.aT.NumMethod()
		for i := 0; i < numMethod; i += 1 {
			missingMethod = err.aT.Method(i).Name
			if _, ok := err.xT.MethodByName(missingMethod); !ok {
				break
			}
		}

		return fmt.Sprintf("interface conversion: %v is not %v: missing method %s",
			err.xT, err.aT, missingMethod)
	} else {
		return fmt.Sprintf("interface conversion: %v is %v, not %v",
			err.xT, err.dynamicT, err.aT)
	}
}

func (err PanicUncomparableType) Error() string {
        return fmt.Sprintf("runtime error: comparing uncomparable type %v", err.dynamicT)
}

func (err PanicUnhashableType) Error() string {
        return fmt.Sprintf("runtime error: hash of unhashable type %v", err.dynamicT)
}
