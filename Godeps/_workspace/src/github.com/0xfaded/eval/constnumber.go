package eval

import "strconv"

type ConstNumber struct {
	Value BigComplex
	Type ConstType
}

// Use with token.INT ast.BasicLit
func NewConstInteger(i string) (*ConstNumber, bool) {
	z := new(ConstNumber)
	z.Type = ConstInt
	z.Value.Re.Denom().SetInt64(1)
	_, ok := z.Value.Re.Num().SetString(i, 0)
	return z, ok
}

// Use with token.FLOAT ast.BasicLit
func NewConstFloat(r string) (*ConstNumber, bool) {
	z := new(ConstNumber)
	z.Type = ConstFloat
	_, ok := z.Value.Re.SetString(r)
	return z, ok
}

// Use with token.IMAG ast.BasicLit
func NewConstImag(i string) (*ConstNumber, bool) {
	z := new(ConstNumber)
	z.Type = ConstComplex
	ok := i[len(i)-1] == 'i'
	if ok {
		_, ok = z.Value.Im.SetString(i[:len(i)-1])
	}
	return z, ok
}

// Use with token.CHAR ast.BasicLit
func NewConstRune(n rune) *ConstNumber {
	z := new(ConstNumber)
	z.Type = ConstRune
	z.Value.Re.Denom().SetInt64(1)
	z.Value.Re.Num().SetInt64(int64(n))
	return z
}

func NewConstInt64(i int64) *ConstNumber {
	z := new(ConstNumber)
	z.Type = ConstInt
	z.Value.Re.Denom().SetInt64(1)
	z.Value.Re.Num().SetInt64(i)
	return z
}

func NewConstUint64(u uint64) *ConstNumber {
	z := new(ConstNumber)
	z.Type = ConstInt
	z.Value.Re.Denom().SetInt64(1)
	z.Value.Re.Num().SetUint64(u)
	return z
}

func NewConstFloat64(f float64) *ConstNumber {
	z := new(ConstNumber)
	z.Type = ConstFloat
	z.Value.Re.SetFloat64(f)
	return z
}

func NewConstComplex128(c complex128) *ConstNumber {
	z := new(ConstNumber)
	z.Type = ConstComplex
	z.Value.Re.SetFloat64(real(c))
	z.Value.Im.SetFloat64(imag(c))
	return z
}

func (z *ConstNumber) String() string {
	return z.StringShow0i(true)
}

func (z *ConstNumber) StringShow0i(show0i bool) string {
	if z.Type == ConstRune && z.Value.Re.Num().BitLen() <= 32 {
		r, _, _ := z.Value.Int(32)
		return strconv.QuoteRuneToASCII(rune(r))
	} else if z.Type == ConstComplex {
		// Hack to show 0i instead of 0
		if z.Value.IsZero() {
			return "0i"
		} else {
			return z.Value.StringShow0i(show0i)
		}
	} else {
		return z.Value.StringShow0i(false)
	}
}

// Add two ConstNumbers, promoting the type automatically.
func (z *ConstNumber) Add(x, y *ConstNumber) *ConstNumber {
	z.Type = promoteConstNumbers(x.Type, y.Type)
	z.Value.Add(&x.Value, &y.Value)
	return z
}

// z.Sub() subtracts two ConstNumbers, promoting the type
// automatically.
func (z *ConstNumber) Sub(x, y *ConstNumber) *ConstNumber {
	z.Type = promoteConstNumbers(x.Type, y.Type)
	z.Value.Sub(&x.Value, &y.Value)
	return z
}

// z.Mul multiplies two ConstNumbers, promoting the type
// automatically.
func (z *ConstNumber) Mul(x, y *ConstNumber) *ConstNumber {
	z.Type = promoteConstNumbers(x.Type, y.Type)
	z.Value.Mul(&x.Value, &y.Value)
	return z
}

// z.Quo divides two ConstNumbers, promoting the type
// automatically. If both operands are of ConstInt, then integer
// division is performed.
func (z *ConstNumber) Quo(x, y *ConstNumber) *ConstNumber {
	z.Type = promoteConstNumbers(x.Type, y.Type)
	if z.Type.IsIntegral() {
		z.Value.Re.Num().Div(x.Value.Re.Num(), y.Value.Re.Num())
	} else {
		z.Value.Quo(&x.Value, &y.Value)
	}
	return z
}

// z.Rem computes remainder of two ConstNumbers, promoting the type
// automatically. The result is undefined if both x and y are not
// integral types.
func (z *ConstNumber) Rem(x, y *ConstNumber) *ConstNumber {
	z.Type = promoteConstNumbers(x.Type, y.Type)
	z.Value.Re.Num().Rem(x.Value.Re.Num(), y.Value.Re.Num())
	return z
}

// z.And compute logical "and" of two ConstNumbers, promoting the type
// automatically.  The result is undefined if both x and y are not
// integral types.
func (z *ConstNumber) And(x, y *ConstNumber) *ConstNumber {
	z.Type = promoteConstNumbers(x.Type, y.Type)
	z.Value.Re.Num().And(x.Value.Re.Num(), y.Value.Re.Num())
	return z
}

// z.Or computes the logical "o"r of two ConstNumbers, promoting the
// type automatically. The result is undefined if both x and y are not
// integral types.
func (z *ConstNumber) Or(x, y *ConstNumber) *ConstNumber {
	z.Type = promoteConstNumbers(x.Type, y.Type)
	z.Value.Re.Num().Or(x.Value.Re.Num(), y.Value.Re.Num())
	return z
}

// z.Xor computes the exclusive "or" of two ConstNumbers, promoting
// the type automatically. The result is undefined if both x and y are
// not integral types.
func (z *ConstNumber) Xor(x, y *ConstNumber) *ConstNumber {
	z.Type = promoteConstNumbers(x.Type, y.Type)
	z.Value.Re.Num().Xor(x.Value.Re.Num(), y.Value.Re.Num())
	return z
}

// z.AndNot computes "and not" of two ConstNumbers, promoting the type
// automatically.  The result is undefined if both x and y are not
// integral types.
func (z *ConstNumber) AndNot(x, y *ConstNumber) *ConstNumber {
	z.Type = promoteConstNumbers(x.Type, y.Type)
	z.Value.Re.Num().AndNot(x.Value.Re.Num(), y.Value.Re.Num())
	return z
}

// z.Shl() shifts x left by count. Type is set to ConstShiftedInt
func (z *ConstNumber) Lsh(x *ConstNumber, count uint) *ConstNumber {
	z.Type = ConstShiftedInt
	z.Value.Lsh(&x.Value, count)
	return z
}

// z.Shl() shifts x right by count. Type is set to ConstShiftedInt
func (z *ConstNumber) Rsh(x *ConstNumber, count uint) *ConstNumber {
	z.Type = ConstShiftedInt
	z.Value.Rsh(&x.Value, count)
	return z
}

