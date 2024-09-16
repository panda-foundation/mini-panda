package ast

type Conversion struct {
	ExpressionBase
	Value Expression
}

func (c *Conversion) Validate(ctx *Context, expected Type) {
	c.Typ = ValidateType(c.Typ, ctx.Program)
	c.Value.Validate(ctx, c.Typ)
	c.Const = c.Value.IsConstant()
	if !((IsNumber(c.Typ) && IsNumber(c.Value.Type())) || (IsPointer(c.Typ) && IsPointer(c.Value.Type()))) {
		ctx.Program.Error(c.Position, "invalid type conversion")
	}
}
