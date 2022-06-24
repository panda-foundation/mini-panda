package constant

import (
	"fmt"

	"github.com/panda-io/micro-panda/ir/core"
	"github.com/panda-io/micro-panda/ir/types"
)

type Null struct {
	Typ *types.PointerType
}

func NewNull(typ *types.PointerType) *Null {
	return &Null{Typ: typ}
}

func (c *Null) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

func (c *Null) Type() core.Type {
	return c.Typ
}

func (*Null) Ident() string {
	return "null"
}
