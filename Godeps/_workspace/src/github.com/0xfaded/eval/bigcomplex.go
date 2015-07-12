package eval

import (
	"fmt"
	"math/big"
)

// BigComplex behaves like a *big.Re, but has an imaginary component
// and separate implementation for + - * /
type BigComplex struct {
	Re big.Rat
	Im big.Rat
}

func (z *BigComplex) Add(x, y *BigComplex) *BigComplex {
	z.Re.Add(&x.Re, &y.Re)
	z.Im.Add(&x.Im, &y.Im)
	return z
}

func (z *BigComplex) Sub(x, y *BigComplex) *BigComplex {
	z.Re.Sub(&x.Re, &y.Re)
	z.Im.Sub(&x.Im, &y.Im)
	return z
}

func (z *BigComplex) Mul(x, y *BigComplex) *BigComplex {
	re := new(big.Rat).Mul(&x.Re, &y.Re)
	re.Sub(re, new(big.Rat).Mul(&x.Im, &y.Im))

	im := new(big.Rat).Mul(&x.Re, &y.Im)
	im.Add(im, new(big.Rat).Mul(&x.Im, &y.Re))

	z.Re = *re
	z.Im = *im
	return z
}

func (z *BigComplex) Quo(x, y *BigComplex) *BigComplex {
	// a+bi   ac+bd   bc-ad
	// ---- = ----- + ----- i
	// c+di   cc+dd   cc+dd

	cc := new(big.Rat).Mul(&y.Re, &y.Re)
	dd := new(big.Rat).Mul(&y.Im, &y.Im)
	ccdd := new(big.Rat).Add(cc, dd)

	ac := new(big.Rat).Mul(&x.Re, &y.Re)
	ad := new(big.Rat).Mul(&x.Re, &y.Im)
	bc := new(big.Rat).Mul(&x.Im, &y.Re)
	bd := new(big.Rat).Mul(&x.Im, &y.Im)

	z.Re.Add(ac, bd)
	z.Re.Quo(&z.Re, ccdd)

	z.Im.Sub(bc, ad)
	z.Im.Quo(&z.Im, ccdd)

	return z
}

func (z *BigComplex) Lsh(x *BigComplex, count uint) *BigComplex {
	z.Re.Num().Lsh(x.Re.Num(), count)
	return z
}

func (z *BigComplex) Rsh(x *BigComplex, count uint) *BigComplex {
	z.Re.Num().Rsh(x.Re.Num(), count)
	return z
}

func (z *BigComplex) IsZero() bool {
	return z.Re.Num().BitLen() == 0 && z.Im.Num().BitLen() == 0
}

// z.Int() returns a representation of z, truncated to be an int of
// length bits.  Valid values for bits are 8, 16, 32, 64. Result is
// otherwise undefined If a truncation occurs, the decimal part is
// dropped and the conversion continues as usual. truncation will be
// true If an overflow occurs, the result is equivelant to a cast of
// the form int32(x). overflow will be true.
func (z *BigComplex) Int(bits int) (_ int64, truncation, overflow bool) {
	var integer *BigComplex
	integer, truncation = z.Integer()
	res := new(big.Int).Set(integer.Re.Num())

	// Numerator must fit in bits - 1, with 1 bit left for sign.
        // An exceptional case when only the signed bit is set.
	if overflow = res.BitLen() > bits - 1; overflow {
		var mask uint64 = ^uint64(0) >> uint(64 - bits)
                if res.BitLen() == bits && res.Sign() < 0 {
                        // To detect the edge of minus 0b1000..., add one
                        // to get 0b0ff... and recount the bits
                        plus1 := new(big.Int).Add(res, big.NewInt(1))
                        if plus1.BitLen() < bits {
	                        return res.Int64(), truncation, false
                        }
                }
		res.And(res, new(big.Int).SetUint64(mask))
	}
	return res.Int64(), truncation, overflow
}

// z.Uint() returns a representation of z truncated to be a uint of
// length bits.  Valid values for bits are 0, 8, 16, 32, 64. The
// returned result is otherwise undefined. If a truncation occurs, the
// decimal part is dropped and the conversion continues as
// usual. Return values truncation and overflow will be true if an
// overflow occurs. The result is equivelant to a cast of the form
// uint32(x).
func (z *BigComplex) Uint(bits int) (_ uint64, truncation, overflow bool) {
	var integer *BigComplex
	integer, truncation = z.Integer()
	res := new(big.Int).Set(integer.Re.Num())

	var mask uint64 = ^uint64(0) >> uint(64 - bits)
	if overflow = res.BitLen() > bits; overflow {
		res.And(res, new(big.Int).SetUint64(mask))
		res = new(big.Int).And(res, new(big.Int).SetUint64(mask))
	}

	r := res.Uint64()
	if res.Sign() < 0 {
		overflow = true
		r = (^r + 1) & mask
	}
	return r, truncation, overflow
}

// z.Float64() returns a representation of z truncated to a float64 If
// a truncation from a complex occurs. The imaginary part is dropped
// and the conversion continues as usual. return value truncation will
// be true exact will be true if the conversion was completed without
// loss of precision.
func (z *BigComplex) Float64() (f float64, truncation, exact bool) {
	f, exact = z.Re.Float64()
	return f, !z.IsReal(), exact
}

// z.Complex128() returns a complex128 representation of z. Return value
// exact will be true if the conversion was completed without loss of
// precision.
func (z *BigComplex) Complex128() (_ complex128, exact bool) {
	r, re := z.Re.Float64()
	i, ie := z.Im.Float64()
	return complex(r, i), re && ie
}

// z.Integer() returns a representation of z, a *BigComplex, truncated
// to be a integer value. The second return value is true if a
// truncation occured in the real component.
func (z *BigComplex) Integer() (_ *BigComplex, truncation bool) {
	if z.IsInteger() {
		return z, false
	} else if z.Re.IsInt() {
		re := new(BigComplex)
		re.Re.Set(&z.Re)
		return re, false
	} else {
		trunc := new(BigComplex)
		trunc.Re.SetInt(z.Re.Num())
		trunc.Re.Num().Div(trunc.Re.Num(), z.Re.Denom())
		return trunc, true
	}
}

// z.Real() returns a representation of z, truncated to a real
// value. The second return valuie is true if a truncation occured.
func (z *BigComplex) Real() (_ *BigComplex, truncation bool) {
	if z.IsReal() {
		return z, false
	} else {
		return &BigComplex{Re: z.Re}, true
	}
}

func (z *BigComplex) IsInteger() bool {
	return z.Re.IsInt() && z.Im.Num().BitLen() == 0
}

func (z *BigComplex) IsReal() bool {
	return z.Im.Num().BitLen() == 0
}

func (z *BigComplex) Equals(other *BigComplex) bool {
	return new(BigComplex).Sub(z, other).IsZero()
}

func (z *BigComplex) String() string {
	return z.StringShow0i(true)
}

func (z *BigComplex) StringShow0i(show0i bool) string {
	var s string
	if z.Re.Num().BitLen() != 0 || show0i {
		if z.Re.IsInt() {
			s += z.Re.Num().String()
		} else {
			f, _ := z.Re.Float64()
			s += fmt.Sprintf("%.5g", f)
		}
	}
	if !z.IsReal() || show0i {
		if s != "" {
			s += "+"
		}
		if z.Im.IsInt() {
			s += z.Im.Num().String()
		} else {
			f, _ := z.Im.Float64()
			s += fmt.Sprintf("%.5g", f)
		}
		s += "i"
	}
	if s == "" {
		s = "0"
	}
	return s
}
