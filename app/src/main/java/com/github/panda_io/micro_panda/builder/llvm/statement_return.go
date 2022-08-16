package llvm

import (
	"github.com/panda-io/micro-panda/ast/statement"
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
	ir_core "github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

func ReturnIR(c *Context, r *statement.Return) {
	if r.Expression != nil {
		var value ir_core.Value
		if r.Expression.IsConstant() {
			value = ExpressionConstIR(c.Program, r.Expression)
		} else {
			value = ExpressionIR(c, r.Expression)
		}
		var t ir_core.Type = ir_types.Void
		if c.Function.Function.Sig.RetType != nil {
			t = c.Function.Function.Sig.RetType
		}
		if value.Type().Equal(t) {
			c.Block.AddInstruction(instruction.NewStore(value, c.Function.Return))
		} else if p, ok := value.Type().(*ir_types.PointerType); ok && p.ElemType.Equal(t) {
			load := instruction.NewLoad(p.ElemType, value)
			c.Block.AddInstruction(load)
			c.Block.AddInstruction(instruction.NewStore(load, c.Function.Return))
		}
	}
	c.Returned = true
	c.Block.AddInstruction(instruction.NewBr(c.Function.Exit))
}
