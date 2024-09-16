package ast

type Parentheses struct {
	ExpressionBase
	Expression Expression
}

func (p *Parentheses) Validate(c *Context, expected Type) {
	p.Expression.Validate(c, expected)
	p.Const = p.Expression.IsConstant()
	p.Typ = p.Expression.Type()
}
