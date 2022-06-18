package expression

import (
	"github.com/panda-io/micro-panda/ast/core"
)

type This struct {
	ExpressionBase
}

func (t *This) Validate(c core.Context, expected core.Type) {
	t.Const = false
	t.Typ = c.FindObject("this")
	if t.Typ == nil {
		c.Error(t.Position, "undefined 'this'")
	}
}
