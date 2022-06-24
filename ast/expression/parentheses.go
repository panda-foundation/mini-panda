package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ast/core"
)

type Parentheses struct {
	ExpressionBase
	Expression ast.Expression
}

func (p *Parentheses) Validate(c ast.Context, expected core.Type) {
	p.Expression.Validate(c, expected)
	p.Const = p.Expression.IsConstant()
	p.Typ = p.Expression.Type()
}
