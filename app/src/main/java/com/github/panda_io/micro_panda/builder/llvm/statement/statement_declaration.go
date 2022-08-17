package llvm

import (
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/ast/statement"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
	ir_core "github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

func DeclarationIR(c *Context, d *statement.DeclarationStatement) {
	t := TypeIR(d.Type)
	alloca := instruction.NewAlloca(t)
	c.Function.Entry.InsertAlloca(alloca)
	var store *instruction.InstStore
	if d.Value == nil {
		store = instruction.NewStore(constant.NewZeroInitializer(t), alloca)
	} else {
		value := ExpressionIR(c, d.Value)
		if ast_types.IsPointer(d.Type) && ast_types.IsArray(d.Value.Type()) {
			var gep instruction.Instruction = instruction.NewGetElementPtr(TypeIR(d.Value.Type()), value, constant.NewInt(ir_types.I32, 0), constant.NewInt(ir_types.I32, 0))
			c.Block.AddInstruction(gep)
			value = gep.(ir_core.Value)
		} else {
			value = c.AutoLoad(value)
		}
		store = instruction.NewStore(value, alloca)
	}
	c.Block.AddInstruction(store)
	c.AddObject(d.Name.Name, alloca)
}
