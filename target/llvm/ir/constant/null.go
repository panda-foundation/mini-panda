package constant

import (
	"fmt"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

type Null struct {
	Typ *ir_types.PointerType
}

func NewNull(typ *ir_types.PointerType) *Null {
	return &Null{Typ: typ}
}

func (c *Null) String() string {
	return fmt.Sprintf("%s %s", c.Type().String(), c.Ident())
}

func (c *Null) Type() ir.Type {
	return c.Typ
}

func (*Null) Ident() string {
	return "null"
}
