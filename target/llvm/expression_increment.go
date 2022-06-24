package llvm

import (
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

func IncrementIR(c *Context, i *expression.Increment) ir.Value {
	t := TypeIR(i.Expression.Type())
	if ir_types.IsInt(t) {
		e := ExpressionIR(c, i.Expression)
		operand := c.AutoLoad(e)
		add := instruction.NewAdd(operand, constant.NewInt(t.(*ir_types.IntType), 1))
		c.Block.AddInstruction(add)
		c.Block.AddInstruction(instruction.NewStore(add, e))
	}
	return nil
}

func IncrementConstIR(p *Program, i *expression.Increment) constant.Constant {
	return nil
}
