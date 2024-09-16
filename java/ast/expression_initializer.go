package ast

type Initializer struct {
	ExpressionBase
	Expressions []Expression
}

func (i *Initializer) Validate(c *Context, expected Type) {
	if array, ok := expected.(*TypeArray); ok {
		i.Typ = array
		i.Const = true
		i.ValidateTypeArray(c, array, i.Expressions)
	} else if t, ok := expected.(*TypeName); ok {
		d := c.Program.FindType(t)
		i.Typ = t
		i.Const = true
		if s, ok := d.(*Struct); ok {
			if len(s.Variables) == len(i.Expressions) {
				for idx, e := range i.Expressions {
					e.Validate(c, s.Variables[idx].Type)
					if !e.IsConstant() {
						c.Program.Error(e.GetPosition(), "expect constant expression initializer")
					}
					if e.Type() != nil && !e.Type().Equal(s.Variables[idx].Type) {
						c.Program.Error(e.GetPosition(), "type mismatch")
					}
				}
			} else {
				c.Program.Error(i.Position, "element number mismatch")
			}
		} else {
			c.Program.Error(i.Position, "undefined type")
		}
	} else {
		c.Program.Error(i.Position, "unexpected initializer")
	}
}

func (i *Initializer) ValidateTypeArray(c *Context, t *TypeArray, exprs []Expression) {
	if t.Dimension[0] == 0 {
		c.Program.Error(i.GetPosition(), "initializer is not allowed to pointer type variable")
	}
	if len(t.Dimension) == 1 {
		if len(exprs) == t.Dimension[0] {
			for _, expr := range exprs {
				expr.Validate(c, t.ElementType)
				if !expr.IsConstant() {
					c.Program.Error(expr.GetPosition(), "expect constant expression initializer")
				}
				if !expr.Type().Equal(t.ElementType) {
					c.Program.Error(expr.GetPosition(), "type mismatch")
				}
			}
		} else {
			c.Program.Error(i.Position, "array length mismatch")
		}
	} else {
		if len(exprs) == t.Dimension[0] {
			sub := &TypeArray{
				ElementType: t.ElementType,
				Dimension:   t.Dimension[1:],
			}
			for _, expr := range exprs {
				if subExprs, ok := expr.(*Initializer); ok {
					i.ValidateTypeArray(c, sub, subExprs.Expressions)
				} else {
					c.Program.Error(expr.GetPosition(), "expect array initializer")
				}
			}
		} else {
			c.Program.Error(i.Position, "array length mismatch")
		}
	}
}
