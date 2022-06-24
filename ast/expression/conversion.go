package expression

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/ast/types"
)

type Conversion struct {
	ExpressionBase
	Value ast.Expression
}

func (c *Conversion) Validate(ctx core.Context, expected core.Type) {
	c.Typ = ctx.ResolveType(c.Typ)
	c.Value.Validate(ctx, c.Typ)
	c.Const = c.Value.IsConstant()
	//TO-DO enum convert with int?
	if !((types.IsNumber(c.Typ) && types.IsNumber(c.Value.Type())) || (types.IsPointer(c.Typ) && types.IsPointer(c.Value.Type()))) {
		ctx.Error(c.GetPosition(), "invalid type conversion")
	}
}
