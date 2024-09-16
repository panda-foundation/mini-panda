package ir

import (
	"fmt"
	"math"
	"math/big"
)

// --- [ Floating-point constants ] --------------------------------------------

// Float is an LLVM IR floating-point constant.
type Float struct {
	// Floating-point type.
	Typ *FloatType
	// Floating-point constant.
	X *big.Float
	// NaN specifies whether the floating-point constant is Not-a-Number.
	NaN bool
}

// NewFloat returns a new floating-point constant based on the given
// floating-point type and double precision floating-point value.
func NewFloat(typ *FloatType, x float64) *Float {
	if math.IsNaN(x) {
		f := &Float{Typ: typ, X: &big.Float{}, NaN: true}
		// Store sign of NaN.
		if math.Signbit(x) {
			f.X.SetFloat64(-1)
		}
		return f
	}
	return &Float{Typ: typ, X: big.NewFloat(x)}
}

// NewFloatFromString returns a new floating-point constant based on the given
// floating-point type and floating-point string.
//
// The floating-point string may be expressed in one of the following forms.
//
//    * fraction floating-point literal
//         [+-]? [0-9]+ [.] [0-9]*
func NewFloatFromString(typ *FloatType, s string) *Float {
	const base = 10
	switch typ.Kind {
	case FloatKindHalf:
		const precision = 11
		x, _, _ := big.ParseFloat(s, base, precision, big.ToNearestEven)
		return &Float{
			Typ: typ,
			X:   x,
		}
	case FloatKindFloat:
		const precision = 24
		x, _, _ := big.ParseFloat(s, base, precision, big.ToNearestEven)
		return &Float{
			Typ: typ,
			X:   x,
		}

	case FloatKindDouble:
		const precision = 53
		x, _, _ := big.ParseFloat(s, base, precision, big.ToNearestEven)
		return &Float{
			Typ: typ,
			X:   x,
		}

	default:
		panic(fmt.Errorf("support for floating-point kind %v not yet implemented", typ.Kind))
	}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *Float) String() string {
	return fmt.Sprintf("%s %s", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *Float) Type() Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *Float) Ident() string {
	f, _ := c.X.Float64()
	bits := math.Float64bits(f)
	return fmt.Sprintf("0x%X", bits)
	/*
		s := c.X.Text('e', 6)
		if !strings.ContainsRune(s, '.') {
			if pos := strings.IndexByte(s, 'e'); pos != -1 {
				s = s[:pos] + ".0" + s[pos:]
			} else {
				s += ".0"
			}
		}
		return s*/
}
