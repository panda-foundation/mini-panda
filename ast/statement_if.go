package ast

type If struct {
	StatementBase
	Initialization Statement
	Condition      Expression
	Body           Statement
	Else           Statement
}

func (i *If) Validate(c *Context) {
	ctx := c.NewContext()
	if i.Initialization != nil {
		i.Initialization.Validate(ctx)
	}
	if i.Condition == nil {
		c.Program.Error(i.Position, "expect condition expression")
	} else {
		i.Condition.Validate(ctx, TypeBool)
		if i.Condition.Type() != nil && !i.Condition.Type().Equal(TypeBool) {
			c.Program.Error(i.Condition.GetPosition(), "expect bool type condition")
		}
	}
	if i.Body != nil {
		bodyCtx := ctx.NewContext()
		i.Body.Validate(bodyCtx)
	}
	if i.Else != nil {
		elseCtx := ctx.NewContext()
		i.Else.Validate(elseCtx)
	}
}
