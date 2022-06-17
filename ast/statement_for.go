package ast

type For struct {
	StatementBase
	Initialization Statement
	Condition      Expression
	Post           Statement
	Body           Statement
}

func (f *For) Validate(c *Context) {
	ctx := c.NewContext()
	if f.Initialization != nil {
		f.Initialization.Validate(ctx)
	}
	if f.Condition != nil {
		conditionCtx := ctx.NewContext()
		f.Condition.Validate(conditionCtx, TypeBool)
		if f.Condition.Type() != nil && !f.Condition.Type().Equal(TypeBool) {
			c.Program.Error(f.Condition.GetPosition(), "expect bool type condition")
		}
	}
	if f.Post != nil {
		postCtx := ctx.NewContext()
		f.Post.Validate(postCtx)
	}
	if f.Body != nil {
		bodyCtx := ctx.NewContext()
		f.Body.Validate(bodyCtx)
	}
}
