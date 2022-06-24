package constant

import (
	"fmt"

	"github.com/panda-io/micro-panda/ir/core"
)

type ZeroInitializer struct {
	Typ core.Type
}

func NewZeroInitializer(typ core.Type) *ZeroInitializer {
	return &ZeroInitializer{Typ: typ}
}

func (c *ZeroInitializer) String() string {
	return fmt.Sprintf("%s %s", c.Type(), c.Ident())
}

func (c *ZeroInitializer) Type() core.Type {
	return c.Typ
}

func (c *ZeroInitializer) Ident() string {
	return "zeroinitializer"
}
