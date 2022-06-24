package llvm

import (
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	ir_expression "github.com/panda-io/micro-panda/target/llvm/ir/constant/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

func SizeofIR(c *Context, s *expression.Sizeof) ir.Value {
	return SizeofConstIR(c.Program, s)
}

func SizeofConstIR(p *Program, s *expression.Sizeof) constant.Constant {
	if t, ok := s.Target.(*ast_types.TypeBuiltin); ok {
		return constant.NewInt(ir_types.I32, int64(ast_types.TypeBuiltinSize(t)))
	} else if ast_types.IsPointer(s.Target) || ast_types.IsFunction(s.Target) {
		p := ir_types.NewPointerType(ir_types.I8)
		size := ir_expression.NewExprGetElementPtr(p, constant.NewNull(ir_types.NewPointerType(p)), constant.NewInt(ir_types.I32, 1))
		return ir_expression.NewExprPtrToInt(size, ir_types.I32)
	} else {
		// struct, array
		t := TypeIR(s.Target)
		size := ir_expression.NewExprGetElementPtr(t, constant.NewNull(ir_types.NewPointerType(t)), constant.NewInt(ir_types.I32, 1))
		return ir_expression.NewExprPtrToInt(size, ir_types.I32)
	}
}
