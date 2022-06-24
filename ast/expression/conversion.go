package expression

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
)

type Conversion struct {
	ExpressionBase
	Value ast.Expression
}

func (c *Conversion) Validate(ctx ast.Context, expected ast.Type) {
	c.Typ = ctx.ResolveType(c.Typ)
	c.Value.Validate(ctx, c.Typ)
	c.Const = c.Value.IsConstant()
	//TO-DO enum convert with int?
	if !((ast_types.IsNumber(c.Typ) && ast_types.IsNumber(c.Value.Type())) || (ast_types.IsPointer(c.Typ) && ast_types.IsPointer(c.Value.Type()))) {
		ctx.Error(c.GetPosition(), "invalid type conversion")
	}
}
