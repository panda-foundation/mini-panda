package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/target/llvm/llvm"
	"github.com/panda-io/micro-panda/target/llvm/types"
)

func SizeofIR(c llvm.Context, s *ast.Sizeof) ir.Value {
	return SizeofConstIR(c.Program, s)
}

func SizeofConstIR(p llvm.Program, s *ast.Sizeof) ir.Constant {
	if t, ok := s.Target.(*ast.TypeBuiltin); ok {
		return ir.NewInt(ir.I32, int64(ast.TypeBuilinSize(t)))
	} else if ast_types.IsPointer(s.Target) || ast_types.IsFunction(s.Target) {
		p := ir.NewPointerType(ir.I8)
		size := ir.NewExprGetElementPtr(p, ir.NewNull(ir.NewPointerType(p)), ir.NewInt(ir.I32, 1))
		return ir.NewExprPtrToInt(size, ir.I32)
	} else {
		// struct, array
		t := types.TypeIR(s.Target)
		size := ir.NewExprGetElementPtr(t, ir.NewNull(ir.NewPointerType(t)), ir.NewInt(ir.I32, 1))
		return ir.NewExprPtrToInt(size, ir.I32)
	}
}
