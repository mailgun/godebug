package eval

import (
	"testing"
)

func TestIntOverflows(t *testing.T) {
	expectIntOverflow(t, 8, newBigInt("0x000000fe"), 0xfe)
	expectIntOverflow(t, 8, newBigInt("-0x000000fe"), 0x02)
	expectIntOverflow(t, 8, newBigInt("0xfffffffe"), 0xfe)
	expectIntOverflow(t, 8, newBigInt("-0xfffffffe"), 0x02)

	expectIntOverflow(t, 16, newBigInt("0x0000fffe"), 0xfffe)
	expectIntOverflow(t, 16, newBigInt("-0x0000fffe"), 0x0002)
	expectIntOverflow(t, 16, newBigInt("0xfffffffe"), 0xfffe)
	expectIntOverflow(t, 16, newBigInt("-0xfffffffe"), 0x0002)

	expectIntOverflow(t, 32, newBigInt("0x00fffffffe"), 0xfffffffe)
	expectIntOverflow(t, 32, newBigInt("-0x00fffffffe"), 0x00000002)
	expectIntOverflow(t, 32, newBigInt("0xfffffffffe"), 0xfffffffe)
	expectIntOverflow(t, 32, newBigInt("-0xfffffffffe"), 0x00000002)

	expectIntOverflow(t, 64, newBigInt("0x00fffffffffffffffe"), -0x0000000000000002)
	expectIntOverflow(t, 64, newBigInt("-0x00fffffffffffffffe"), 0x0000000000000002)
	expectIntOverflow(t, 64, newBigInt("0xfffffffffffffffffe"), -0x0000000000000002)
	expectIntOverflow(t, 64, newBigInt("-0xfffffffffffffffffe"), 0x0000000000000002)
}

func TestUintOverflows(t *testing.T) {
	expectUintOverflow(t, 8, newBigInt("0x000001fe"), 0xfe)
	expectUintOverflow(t, 8, newBigInt("0xfffffffe"), 0xfe)
	expectUintOverflow(t, 8, newBigInt("-0x000000fe"), 0x02)
	expectUintOverflow(t, 8, newBigInt("-0x000001fe"), 0x02)
	expectUintOverflow(t, 8, newBigInt("-0xfffffffe"), 0x02)

	expectUintOverflow(t, 16, newBigInt("0x0001fffe"), 0xfffe)
	expectUintOverflow(t, 16, newBigInt("0xfffffffe"), 0xfffe)
	expectUintOverflow(t, 16, newBigInt("-0x0000fffe"), 0x0002)
	expectUintOverflow(t, 16, newBigInt("-0x0001fffe"), 0x0002)
	expectUintOverflow(t, 16, newBigInt("-0xfffffffe"), 0x0002)

	expectUintOverflow(t, 32, newBigInt("0x01fffffffe"), 0xfffffffe)
	expectUintOverflow(t, 32, newBigInt("0xfffffffffe"), 0xfffffffe)
	expectUintOverflow(t, 32, newBigInt("-0x00fffffffe"), 0x00000002)
	expectUintOverflow(t, 32, newBigInt("-0x01fffffffe"), 0x00000002)
	expectUintOverflow(t, 32, newBigInt("-0xfffffffffe"), 0x00000002)

	expectUintOverflow(t, 64, newBigInt("0x01fffffffffffffffe"), 0xfffffffffffffffe)
	expectUintOverflow(t, 64, newBigInt("0xfffffffffffffffffe"), 0xfffffffffffffffe)
	expectUintOverflow(t, 64, newBigInt("-0x00fffffffffffffffe"), 0x0000000000000002)
	expectUintOverflow(t, 64, newBigInt("-0x01fffffffffffffffe"), 0x0000000000000002)
	expectUintOverflow(t, 64, newBigInt("-0xfffffffffffffffffe"), 0x0000000000000002)
}

func expectIntOverflow(t *testing.T, bits int, c *BigComplex, expected int64) {
	if result, truncation, overflow := c.Int(bits); truncation {
		t.Fatalf("Unexpected truncation")
	} else if !overflow {
		t.Fatalf("Expected overflow")
	} else if result != expected {
		t.Fatalf("Expected %v, got %v\n", expected, result)
	}
}

func expectUintOverflow(t *testing.T, bits int, c *BigComplex, expected uint64) {
	if result, truncation, overflow := c.Uint(bits); truncation {
		t.Fatalf("Unexpected truncation")
	} else if !overflow {
		t.Fatalf("Expected overflow")
	} else if result != expected {
		t.Fatalf("Expected %v, got %v\n", expected, result)
	}
}

func newBigInt(i string) *BigComplex {
	integer := new(BigComplex)
	integer.Re.Denom().SetInt64(1)
	if _, ok := integer.Re.Num().SetString(i, 0); !ok {
		panic("Invalid BigInt string '" + i + "'")
	} else {
		return integer
	}
}
