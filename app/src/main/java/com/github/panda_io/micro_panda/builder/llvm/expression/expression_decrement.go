package llvm

import (
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

func DecrementIR(c *Context, d *expression.Decrement) ir.Value {
	t := TypeIR(d.Expression.Type())
	if ir_types.IsInt(t) {
		e := ExpressionIR(c, d.Expression)
		operand := c.AutoLoad(e)
		sub := instruction.NewSub(operand, constant.NewInt(t.(*ir_types.IntType), 1))
		c.Block.AddInstruction(sub)
		c.Block.AddInstruction(instruction.NewStore(sub, e))
	}
	return nil
}

func DecrementConstIR(p *Program, d *expression.Decrement) constant.Constant {
	return nil
}
