package expression

import (
	"github.com/panda-io/micro-panda/ast/ast"
)

type Parentheses struct {
	ExpressionBase
	Expression ast.Expression
}

func (p *Parentheses) Validate(c ast.Context, expected ast.Type) {
	p.Expression.Validate(c, expected)
	p.Const = p.Expression.IsConstant()
	p.Typ = p.Expression.Type()
}
