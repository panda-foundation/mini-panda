package expression

import (
	"github.com/panda-io/micro-panda/ast/core"
)

type Decrement struct {
	ExpressionBase
	Expression core.Expression
}

func (d *Decrement) Validate(c core.Context, expected core.Type) {
	d.Expression.Validate(c, expected)
	d.Const = false
	if d.Expression.IsConstant() {
		c.Error(d.Position, "expect variable")
	}
	if !core.IsInteger(d.Expression.Type()) {
		c.Error(d.Position, "expect integer expression")
	}
}
