package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

func DeclarationIR(c *Context, d *ast.DeclarationStatement) {
	t := TypeIR(d.Type)
	alloca := ir.NewAlloca(t)
	c.Function.Entry.InsertAlloca(alloca)
	var store *ir.InstStore
	if d.Value == nil {
		store = ir.NewStore(ir.NewZeroInitializer(t), alloca)
	} else {
		value := ExpressionIR(c, d.Value)
		if ast.IsPointer(d.Type) && ast.IsArray(d.Value.Type()) {
			var gep ir.Instruction = ir.NewGetElementPtr(TypeIR(d.Value.Type()), value, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, 0))
			c.Block.AddInstruction(gep)
			value = gep.(ir.Value)
		} else {
			value = c.AutoLoad(value)
		}
		store = ir.NewStore(value, alloca)
	}
	c.Block.AddInstruction(store)
	c.AddObject(d.Name.Name, alloca)
}
