package constant

import (
	"fmt"
	"math/big"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

type Int struct {
	Typ *ir_types.IntType
	X   *big.Int
}

func NewInt(typ *ir_types.IntType, x int64) *Int {
	return &Int{Typ: typ, X: big.NewInt(x)}
}

func NewBool(x bool) *Int {
	if x {
		return True
	}
	return False
}

func NewIntFromString(typ *ir_types.IntType, s string) *Int {
	switch s {
	case "true":
		if !typ.Equal(ir_types.I1) {
			panic(fmt.Errorf("invalid boolean type; expected i1, got %T", typ))
		}
		return True
	case "false":
		if !typ.Equal(ir_types.I1) {
			panic(fmt.Errorf("invalid boolean type; expected i1, got %T", typ))
		}
		return False
	}
	x, _ := (&big.Int{}).SetString(s, 0)
	return &Int{Typ: typ, X: x}
}

func (c *Int) String() string {
	return fmt.Sprintf("%s %s", c.Type().String(), c.Ident())
}

func (c *Int) Type() ir.Type {
	return c.Typ
}

func (c *Int) Ident() string {
	if c.Typ.BitSize == 1 {
		switch c.X.Int64() {
		case 0:
			return "false"
		case 1:
			return "true"
		default:
			panic(fmt.Errorf("invalid integer value of boolean type; expected 0 or 1, got %d", c.X.Int64()))
		}
	}
	return c.X.String()
}
