package ast

type This struct {
	ExpressionBase
}

func (t *This) Validate(c *Context, expected Type) {
	t.Const = false
	t.Typ = c.FindObject(StructThis)
	if t.Typ == nil {
		c.Program.Error(t.Position, "undefined 'this'")
	}
}
