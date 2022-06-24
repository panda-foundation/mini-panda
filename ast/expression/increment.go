package expression

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
)

type Increment struct {
	ExpressionBase
	Expression ast.Expression
}

func (i *Increment) Validate(c ast.Context, expected ast.Type) {
	i.Const = false
	i.Expression.Validate(c, expected)
	if i.Expression.IsConstant() {
		c.Error(i.GetPosition(), "expect variable")
	}
	if !ast_types.IsInteger(i.Expression.Type()) {
		c.Error(i.GetPosition(), "expect integer expression")
	}
}
