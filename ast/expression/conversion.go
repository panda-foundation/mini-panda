package expression

import (
	"github.com/panda-io/micro-panda/ast/core"
)

type Conversion struct {
	ExpressionBase
	Value core.Expression
}

func (c *Conversion) Validate(ctx core.Context, expected core.Type) {
	c.Typ = ctx.ResolveType(c.Typ)
	c.Value.Validate(ctx, c.Typ)
	c.Const = c.Value.IsConstant()
	if !((core.IsNumber(c.Typ) && core.IsNumber(c.Value.Type())) || (core.IsPointer(c.Typ) && core.IsPointer(c.Value.Type()))) {
		ctx.Error(c.Position, "invalid type conversion")
	}
}
