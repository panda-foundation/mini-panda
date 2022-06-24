package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

func IncrementIR(c *Context, i *ast.Increment) ir.Value {
	t := TypeIR(i.Expression.Type())
	if ir.IsInt(t) {
		e := ExpressionIR(c, i.Expression)
		operand := c.AutoLoad(e)
		add := ir.NewAdd(operand, ir.NewInt(t.(*ir.IntType), 1))
		c.Block.AddInstruction(add)
		c.Block.AddInstruction(ir.NewStore(add, e))
	}
	return nil
}

func IncrementConstIR(p *Program, i *ast.Increment) ir.Constant {
	return nil
}
