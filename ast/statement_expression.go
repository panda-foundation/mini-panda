package ast

type ExpressionStatement struct {
	StatementBase
	Expression Expression
}

func (e *ExpressionStatement) Validate(c *Context) {
	e.Expression.Validate(c, nil)
}
