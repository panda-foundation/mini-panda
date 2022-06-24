package constant

import (
	"fmt"
	"math/big"

	"github.com/panda-io/micro-panda/ir/core"
	"github.com/panda-io/micro-panda/ir/types"
)

type Int struct {
	Typ *types.IntType
	X   *big.Int
}

func NewInt(typ *types.IntType, x int64) *Int {
	return &Int{Typ: typ, X: big.NewInt(x)}
}

func NewBool(x bool) *Int {
	if x {
		return True
	}
	return False
}

func NewIntFromString(typ *types.IntType, s string) *Int {
	switch s {
	case "true":
		if !typ.Equal(types.I1) {
			panic(fmt.Errorf("invalid boolean type; expected i1, got %T", typ))
		}
		return True
	case "false":
		if !typ.Equal(types.I1) {
			panic(fmt.Errorf("invalid boolean type; expected i1, got %T", typ))
		}
		return False
	}
	x, _ := (&big.Int{}).SetString(s, 0)
	return &Int{Typ: typ, X: x}
}

func (c *Int) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

func (c *Int) Type() core.Type {
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
