package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
	"github.com/panda-io/micro-panda/target/llvm/llvm"
	"github.com/panda-io/micro-panda/target/llvm/types"
)

func IncrementIR(c llvm.Context, i *ast.Increment) ir.Value {
	t := types.TypeIR(i.Expression.Type())
	if ir_types.IsInt(t) {
		e := ExpressionIR(c, i.Expression)
		operand := c.AutoLoad(e)
		add := ir.NewAdd(operand, ir.NewInt(t.(*ir.IntType), 1))
		c.Block().AddInstruction(add)
		c.Block().AddInstruction(ir.NewStore(add, e))
	}
	return nil
}

func IncrementConstIR(p llvm.Program, i *ast.Increment) ir.Constant {
	return nil
}
