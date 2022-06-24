package constant

import (
	"fmt"
	"math"
	"math/big"

	"github.com/panda-io/micro-panda/ir/core"
	"github.com/panda-io/micro-panda/ir/types"
)

type Float struct {
	Typ *types.FloatType
	X   *big.Float
	NaN bool
}

func NewFloat(typ *types.FloatType, x float64) *Float {
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

func NewFloatFromString(typ *types.FloatType, s string) *Float {
	const base = 10
	switch typ.Kind {
	case types.FloatKindHalf:
		const precision = 11
		x, _, _ := big.ParseFloat(s, base, precision, big.ToNearestEven)
		return &Float{
			Typ: typ,
			X:   x,
		}
	case types.FloatKindFloat:
		const precision = 24
		x, _, _ := big.ParseFloat(s, base, precision, big.ToNearestEven)
		return &Float{
			Typ: typ,
			X:   x,
		}

	case types.FloatKindDouble:
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

func (c *Float) String() string {
	return fmt.Sprintf("%s %s", c.Type(), c.Ident())
}

func (c *Float) Type() core.Type {
	return c.Typ
}

func (c *Float) Ident() string {
	f, _ := c.X.Float64()
	bits := math.Float64bits(f)
	return fmt.Sprintf("0x%X", bits)
}
