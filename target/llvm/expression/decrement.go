package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
	"github.com/panda-io/micro-panda/target/llvm/llvm"
	"github.com/panda-io/micro-panda/target/llvm/types"
)

func DecrementIR(c llvm.Context, d *ast.Decrement) ir.Value {
	t := types.TypeIR(d.Expression.Type())
	if ir_types.IsInt(t) {
		e := ExpressionIR(c, d.Expression)
		operand := c.AutoLoad(e)
		sub := ir.NewSub(operand, ir.NewInt(t.(*ir.IntType), 1))
		c.Block().AddInstruction(sub)
		c.Block().AddInstruction(ir.NewStore(sub, e))
	}
	return nil
}

func DecrementConstIR(p llvm.Program, d *ast.Decrement) ir.Constant {
	return nil
}
