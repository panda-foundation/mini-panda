package types

import (
	ast_types "github.com/panda-io/micro-panda/ast/types"
	ir_core "github.com/panda-io/micro-panda/ir/core"
	ir_types "github.com/panda-io/micro-panda/ir/types"
)

func TypePointerIR(t *ast_types.TypePointer) ir_core.Type {
	return ir_types.NewPointerType(TypeIR(t.ElementType))
}
