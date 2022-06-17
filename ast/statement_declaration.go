package ast

type DeclarationStatement struct {
	StatementBase
	Name  *Identifier
	Type  Type
	Value Expression
}

func (d *DeclarationStatement) Validate(c *Context) {
	d.Type = ValidateType(d.Type, c.Program)
	if d.Value != nil {
		d.Value.Validate(c, d.Type)
		if d.Value.Type() != nil && d.Type != nil && !d.Value.Type().Equal(d.Type) {
			c.Program.Error(d.Value.GetPosition(), "init value type mismatch with define")
		}
	}
	if d.Type != nil {
		err := c.AddObject(d.Name.Name, d.Type)
		if err != nil {
			c.Program.Error(d.Position, err.Error())
		}
	}
}
