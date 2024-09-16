package ast

type Return struct {
	StatementBase
	Expression Expression
}

func (r *Return) Validate(c *Context) {
	r.Expression.Validate(c, c.Function.ReturnType)
	if r.Expression == nil {
		if c.Function.ReturnType != nil {
			c.Program.Error(r.Position, "mismatch return type, expect 'null'")
		}
	} else if r.Expression.Type() != nil && !r.Expression.Type().Equal(c.Function.ReturnType) {
		c.Program.Error(r.Position, "mismatch return type")
	}
}
