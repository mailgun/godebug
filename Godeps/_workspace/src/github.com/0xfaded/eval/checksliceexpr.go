package eval

import (
	"reflect"
	"go/ast"
)

// TODO[crc] support [::] syntax after go1.2 upgrade
func checkSliceExpr(slice *ast.SliceExpr, env Env) (*SliceExpr, []error) {
	aexpr := &SliceExpr{SliceExpr: slice}
	x, errs := CheckExpr(slice.X, env)
	aexpr.X = x
	if errs != nil && !x.IsConst() {
		return aexpr, errs
	}

	t, err := expectSingleType(x)
	if err != nil {
		return aexpr, append(errs, err)
	}

	// arrays must be addressable
	if t.Kind() == reflect.Array && !isAddressable(x) {
		return aexpr, append(errs, ErrUnaddressableSliceOperand{aexpr})
	}
	// slice of array pointer is short hand for dereference and then slice
	if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Array {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.Array, reflect.Slice, reflect.String:
		var l, h int
		var low, high Expr
		var moreErrs []error
		if t == ConstString {
			// spec: ConstString[:] fields string
			aexpr.knownType = knownType{stringType}
		} else {
			aexpr.knownType = knownType(x.KnownType())
		}
		if slice.Low != nil {
			low, l, moreErrs = checkSliceVectorExpr(x, slice.Low, env)
			aexpr.Low = low
			if moreErrs != nil {
				errs = append(errs, moreErrs...)
				if !low.IsConst() {
					return aexpr, errs
				}
			}
		}
		if slice.High != nil {
			high, h, moreErrs = checkSliceVectorExpr(x, slice.High, env)
			aexpr.High = high
			if moreErrs != nil {
				errs = append(errs, moreErrs...)
				if !high.IsConst() {
					return aexpr, errs
				}
			}
			if low != nil && low.IsConst() && high.IsConst() && !(l <= h) {
				errs = append(errs, ErrInvalidSliceIndex{aexpr})
			}
		}
		return aexpr, errs
	default:
		return aexpr, append(errs, ErrInvalidSliceOperation{aexpr})
	}
}

func checkSliceVectorExpr(x Expr, index ast.Expr, env Env) (Expr, int, []error) {
	t := x.KnownType()[0]
	i, iint, ok, errs := checkInteger(index, env)
	if errs != nil && !i.IsConst() {
		// Type check of index failed
	} else if !ok {
		// Type check of index passed but this node is not an integer
		printableIndex := fakeCheckExpr(index, env)
		printableIndex.setKnownType(i.KnownType())
		errs = append(errs, ErrNonIntegerIndex{printableIndex})
	} else if i.IsConst() {
		// If we know the index at compile time, we must assert it is in bounds.
		// NOTE[crc] There is no upper bounds check on a const string. This is
		// to match gc. See issue http://code.google.com/p/go/issues/detail?id=7200
		if iint < 0 {
			errs = append(errs, ErrIndexOutOfBounds{i, x, iint})
		} else if t.Kind() == reflect.Array {
			if iint >= t.Len() {
				errs = append(errs, ErrIndexOutOfBounds{i, x, iint})
			}
		}
	}
	return i, iint, errs
}
