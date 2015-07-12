package eval

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"unicode"
)

// Inspect prints a reflect.Value the way you would enter it.
// Some like this should really be part of the reflect package.
func Inspect(val reflect.Value) string {

	if val.CanInterface() {
		if s, ok := val.Interface().(fmt.Stringer); ok {
			return s.String()
		}
	}
	switch val.Kind() {
	case reflect.String:
		return strconv.QuoteToASCII(val.String())
	case reflect.Slice, reflect.Array:
		if val.Len() == 0 {
			return "[]"
		}
		sep := "{"
		str := ""
		for i:=0; i < val.Len(); i++ {
			str += sep
			sep = ", "
			str += InspectShort(val.Index(i))
		}
		str += "}"
		return str

	case reflect.Struct:
		if val.Type() == untypedNilType {
			return "nil"
		}
		n  := val.NumField()
		str := val.Type().String()
		if n == 0 {
			return str + "{}"
		}
		sep := "{\n\t"
		unexported := false
		for i := 0; i < n; i++ {
			name  := val.Type().Field(i).Name
			if unicode.IsLower([]rune(name)[0]) {
				unexported = true
				continue
			}
			field := val.Field(i)
			str += fmt.Sprintf("%s%s: %s", sep, name, InspectShort(field))
			sep = ",\n\t"
		}
		if unexported {
			str += fmt.Sprintf("\n\t// unexported fields")
		}
		str += "\n}"
		return str

	case reflect.Map:
		str := val.Type().String()
		if val.Len() == 0 {
			return str + "{}"
		}
		sep := " {\n\t"
		keys := sortMapKeys(val.MapKeys())
		for _, k := range keys {
			v := val.MapIndex(k)
			str += fmt.Sprintf("%s%s: %s", sep, InspectShort(k), InspectShort(v))
			sep = ",\n\t"
		}
		str += "\n}"
		return str

	case reflect.Ptr:
		// Internal const numbers
		i := val.Interface()
		if cn, ok := i.(*ConstNumber); ok {
			return fmt.Sprint(cn)
		} else if val.IsNil() {
			return "nil"
		} else {
			return "&" + InspectShort(reflect.Indirect(val))
		}
	case reflect.Interface:
		if val.IsNil() {
			return "nil"
		} else {
			return fmt.Sprintf("%s.(%v)", InspectShort(val.Elem()), val.Type())
		}
	default:
		// FIXME: add more Kinds as folks are annoyed with the output of
		// the below:
		if val.CanInterface() {
			return fmt.Sprintf("%v", val.Interface())
		} else {
			return fmt.Sprintf("<%v>", val.Type())
		}
	}
}

// Returns type{...} for composite lits
func InspectShort(val reflect.Value) string {
	switch val.Kind() {
	case reflect.Slice, reflect.Array, reflect.Struct, reflect.Map:
		return fmt.Sprintf("%v{...}", val.Type())
	default:
		return Inspect(val)
	}
}

type mapKeys []reflect.Value
func sortMapKeys(keys []reflect.Value) []reflect.Value{
	ks := mapKeys(keys)
	sort.Sort(ks)
	return ks
}

func (keys mapKeys) Len() int {
	return len(keys)
}

func (keys mapKeys) Swap(i, j int) {
	keys[i], keys[j] = keys[j], keys[i]
}

func (keys mapKeys) Less(i, j int) bool {
	x, y := keys[i], keys[j]
	switch x.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return x.Int() < y.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return x.Uint() < y.Uint()
	case reflect.Float32, reflect.Float64:
		return x.Float() < y.Float()
	case reflect.Complex64, reflect.Complex128:
		return real(x.Complex()) < real(y.Complex())
	case reflect.String:
		return x.String() < y.String()
	case reflect.Bool:
		return x.Bool()
	default:
		return i < j
	}
}
