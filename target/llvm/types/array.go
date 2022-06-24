package types

import (
	ast_types "github.com/panda-io/micro-panda/ast/types"
	ir_core "github.com/panda-io/micro-panda/ir/core"
	ir_types "github.com/panda-io/micro-panda/ir/types"
)

func TypeArrayIR(t *ast_types.TypeArray) ir_core.Type {
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
