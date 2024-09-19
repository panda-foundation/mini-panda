package ast

type Decrement struct {
	ExpressionBase
	Expression Expression
}

func (d *Decrement) Validate(c *Context, expected Type) {
	d.Expression.Validate(c, expected)
	d.Const = false
	if d.Expression.IsConstant() {
		c.Program.Error(d.Position, "expect variable")
	}
	if !IsInteger(d.Expression.Type()) {
		c.Program.Error(d.Position, "expect integer expression")
	}
}
