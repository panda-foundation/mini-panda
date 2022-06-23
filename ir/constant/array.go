package constant

import (
	"fmt"
	"strings"

	"github.com/panda-io/micro-panda/ir/core"
	"github.com/panda-io/micro-panda/ir/types"
)

type Array struct {
	Typ   *types.ArrayType
	Elems []Constant
}

func NewArray(t *types.ArrayType, elems ...Constant) *Array {
	c := &Array{
		Elems: elems,
		Typ:   t,
	}
	c.Type()
	return c
}

func (c *Array) String() string {
	return fmt.Sprintf("%s %s", c.Type(), c.Ident())
}

func (c *Array) Type() core.Type {
	if c.Typ == nil {
		elemType := c.Elems[0].Type()
		c.Typ = types.NewArrayType(uint64(len(c.Elems)), elemType)
	}
	return c.Typ
}

func (c *Array) Ident() string {
	buf := &strings.Builder{}
	buf.WriteString("[")
	for i, elem := range c.Elems {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(elem.String())
	}
	buf.WriteString("]")
	return buf.String()
}

type CharArray struct {
	Typ *types.ArrayType
	X   []byte
}

func NewCharArray(x []byte) *CharArray {
	typ := types.NewArrayType(uint64(len(x)), types.I8)
	return &CharArray{Typ: typ, X: x}
}

func NewCharArrayFromString(s string) *CharArray {
	return NewCharArray([]byte(s))
}

func (c *CharArray) String() string {
	return fmt.Sprintf("%s %s", c.Type(), c.Ident())
}

func (c *CharArray) Type() core.Type {
	return c.Typ
}

func (c *CharArray) Ident() string {
	return "c" + core.Quote(c.X)
}
