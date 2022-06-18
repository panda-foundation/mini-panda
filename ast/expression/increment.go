package expression

import (
	"github.com/panda-io/micro-panda/ast/core"
)

type Increment struct {
	ExpressionBase
	Expression core.Expression
}

func (i *Increment) Validate(c core.Context, expected core.Type) {
	i.Const = false
	i.Expression.Validate(c, expected)
	if i.Expression.IsConstant() {
		c.Error(i.Position, "expect variable")
	}
	if !core.IsInteger(i.Expression.Type()) {
		c.Error(i.Position, "expect integer expression")
	}
}
