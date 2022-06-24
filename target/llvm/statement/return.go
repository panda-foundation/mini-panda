package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/target/llvm"
)

func ReturnIR(c llvm.Context, r *ast.Return) {
	if r.Expression != nil {
		var value ir.Value
		if r.Expression.IsConstant() {
			value = ExpressionConstIR(c.Program, r.Expression)
		} else {
			value = ExpressionIR(c, r.Expression)
		}
		var t ir.Type = ir.Void
		if c.Function.Function.Sig.RetType != nil {
			t = c.Function.Function.Sig.RetType
		}
		if value.Type().Equal(t) {
			c.Block().AddInstruction(ir.NewStore(value, c.Function.Return))
		} else if p, ok := value.Type().(*ir.PointerType); ok && p.ElemType.Equal(t) {
			load := ir.NewLoad(p.ElemType, value)
			c.Block().AddInstruction(load)
			c.Block().AddInstruction(ir.NewStore(load, c.Function.Return))
		}
	}
	c.Returned = true
	c.Block().AddInstruction(ir.NewBr(c.Function.Exit))
}
