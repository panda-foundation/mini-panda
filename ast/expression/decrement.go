package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/ast/types"
)

type Decrement struct {
	ExpressionBase
	Expression ast.Expression
}

func (d *Decrement) Validate(c ast.Context, expected core.Type) {
	d.Expression.Validate(c, expected)
	d.Const = false
	if d.Expression.IsConstant() {
		c.Error(d.GetPosition(), "expect variable")
	}
	if !types.IsInteger(d.Expression.Type()) {
		c.Error(d.GetPosition(), "expect integer expression")
	}
}
