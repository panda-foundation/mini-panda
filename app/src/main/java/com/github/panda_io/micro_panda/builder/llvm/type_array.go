package llvm

import (
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

func TypeArrayIR(t *ast_types.TypeArray) ir.Type {
	e := TypeIR(t.ElementType)
	if t.Dimension[0] == 0 {
		if len(t.Dimension) == 1 {
			return ir_types.NewPointerType(e)
		} else {
			array := ir_types.NewArrayType(uint64(t.Dimension[len(t.Dimension)-1]), e)
			for i := len(t.Dimension) - 3; i >= 0; i-- {
				array = ir_types.NewArrayType(uint64(t.Dimension[i]), array)
			}
			return ir_types.NewPointerType(array)
		}
	} else {
		array := ir_types.NewArrayType(uint64(t.Dimension[len(t.Dimension)-1]), e)
		for i := len(t.Dimension) - 2; i >= 0; i-- {
			array = ir_types.NewArrayType(uint64(t.Dimension[i]), array)
		}
		return array
	}
}
