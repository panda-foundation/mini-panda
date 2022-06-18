package expression

import (
	"github.com/panda-io/micro-panda/ast/core"
)

type Parentheses struct {
	ExpressionBase
	Expression core.Expression
}

func (p *Parentheses) Validate(c core.Context, expected core.Type) {
	p.Expression.Validate(c, expected)
	p.Const = p.Expression.IsConstant()
	p.Typ = p.Expression.Type()
}
