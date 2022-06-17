package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

func DecrementIR(c *Context, d *ast.Decrement) ir.Value {
	t := TypeIR(d.Expression.Type())
	if ir.IsInt(t) {
		e := ExpressionIR(c, d.Expression)
		operand := c.AutoLoad(e)
		sub := ir.NewSub(operand, ir.NewInt(t.(*ir.IntType), 1))
		c.Block.AddInstruction(sub)
		c.Block.AddInstruction(ir.NewStore(sub, e))
	}
	return nil
}

func DecrementConstIR(p *Program, d *ast.Decrement) ir.Constant {
	return nil
}
