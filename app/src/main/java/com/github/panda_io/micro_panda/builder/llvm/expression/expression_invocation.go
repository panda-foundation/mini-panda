package llvm

import (
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

func InvocationIR(c *Context, i *expression.Invocation) ir.Value {
	t := i.Function.Type()
	if _, ok := t.(*ast_types.TypeFunction); ok {
		var call *instruction.InstCall
		callee := ExpressionIR(c, i.Function)
		call = instruction.NewCall(c.AutoLoad(callee))
		if i.Arguments != nil {
			for _, arg := range i.Arguments {
				v := ExpressionIR(c, arg)
				if t, ok := arg.Type().(*ast_types.TypeArray); ok && t.Dimension[0] != 0 {
					var gep instruction.Instruction = instruction.NewGetElementPtr(TypeIR(t), v, constant.NewInt(ir_types.I32, 0), constant.NewInt(ir_types.I32, 0))
					c.Block.AddInstruction(gep)
					v = gep.(ir.Value)
				} else {
					v = c.AutoLoad(v)
				}
				call.Args = append(call.Args, v)
			}
		}
		c.Block.AddInstruction(call)
		return call
	}
	return nil
}

func InvocationConstIR(p *Program, i *expression.Invocation) constant.Constant {
	return nil
}
