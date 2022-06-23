package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

func TypePointerIR(t *ast.TypePointer) ir.Type {
	return ir.NewPointerType(TypeIR(t.ElementType))
}
