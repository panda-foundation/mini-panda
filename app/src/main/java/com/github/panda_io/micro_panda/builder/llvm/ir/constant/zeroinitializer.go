package constant

import (
	"fmt"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type ZeroInitializer struct {
	Typ ir.Type
}

func NewZeroInitializer(typ ir.Type) *ZeroInitializer {
	return &ZeroInitializer{Typ: typ}
}

func (c *ZeroInitializer) String() string {
	return fmt.Sprintf("%s %s", c.Type().String(), c.Ident())
}

func (c *ZeroInitializer) Type() ir.Type {
	return c.Typ
}

func (c *ZeroInitializer) Ident() string {
	return "zeroinitializer"
}
