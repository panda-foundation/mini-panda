package types

import (
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

func TypePointerIR(t *ast_types.TypePointer) ir.Type {
	return ir_types.NewPointerType(TypeIR(t.ElementType))
}
