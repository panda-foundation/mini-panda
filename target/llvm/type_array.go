package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

func TypeArrayIR(t *ast.TypeArray) ir.Type {
	e := TypeIR(t.ElementType)
	if t.Dimension[0] == 0 {
		if len(t.Dimension) == 1 {
			return ir.NewPointerType(e)
		} else {
			array := ir.NewArrayType(uint64(t.Dimension[len(t.Dimension)-1]), e)
			for i := len(t.Dimension) - 3; i >= 0; i-- {
				array = ir.NewArrayType(uint64(t.Dimension[i]), array)
			}
			return ir.NewPointerType(array)
		}
	} else {
		array := ir.NewArrayType(uint64(t.Dimension[len(t.Dimension)-1]), e)
		for i := len(t.Dimension) - 2; i >= 0; i-- {
			array = ir.NewArrayType(uint64(t.Dimension[i]), array)
		}
		return array
	}
}
