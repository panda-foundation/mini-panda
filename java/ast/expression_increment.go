package ast

type Increment struct {
	ExpressionBase
	Expression Expression
}

func (i *Increment) Validate(c *Context, expected Type) {
	i.Const = false
	i.Expression.Validate(c, expected)
	if i.Expression.IsConstant() {
		c.Program.Error(i.Position, "expect variable")
	}
	if !IsInteger(i.Expression.Type()) {
		c.Program.Error(i.Position, "expect integer expression")
	}
}
