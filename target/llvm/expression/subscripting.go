package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/target/llvm/llvm"
	"github.com/panda-io/micro-panda/target/llvm/types"
)

func SubscriptingIR(c llvm.Context, s *ast.Subscripting) ir.Value {
	src := ExpressionIR(c, s.Parent)
	var srcType ir.Type
	indexes := []ir.Value{}
	if ast_types.IsArray(s.Parent.Type()) {
		indexes = append(indexes, ir.NewInt(ir.I32, 0))
		srcType = types.TypeIR(s.Parent.Type())
	} else {
		load := ir.NewLoad(src.Type().(*ir.PointerType).ElemType, src)
		c.Block().AddInstruction(load)
		src = load
		srcType = types.TypeIR(ast.ElementType(s.Parent.Type()))
	}
	for _, i := range s.Indexes {
		indexes = append(indexes, c.AutoLoad(ExpressionIR(c, i)))
	}
	v := ir.NewGetElementPtr(srcType, src, indexes...)
	c.Block().AddInstruction(v)
	return v
}

func SubscriptingConstIR(p llvm.Program, s *ast.Subscripting) ir.Constant {
	if t, ok := s.Parent.Type().(*ast.TypeArray); ok {
		indexes := []ir.Constant{ir.NewInt(ir.I32, 0)}
		for _, i := range s.Indexes {
			indexes = append(indexes, ExpressionConstIR(p, i))
		}
		return ir.NewExprGetElementPtr(types.TypeIR(t), ExpressionConstIR(p, s.Parent), indexes...)
	}
	return nil
}
