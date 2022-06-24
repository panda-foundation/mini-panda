package llvm

import (
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	ir_expression "github.com/panda-io/micro-panda/target/llvm/ir/constant/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

func SubscriptingIR(c *Context, s *expression.Subscripting) ir.Value {
	src := ExpressionIR(c, s.Parent)
	var srcType ir.Type
	indexes := []ir.Value{}
	if ast_types.IsArray(s.Parent.Type()) {
		indexes = append(indexes, constant.NewInt(ir_types.I32, 0))
		srcType = TypeIR(s.Parent.Type())
	} else {
		load := instruction.NewLoad(src.Type().(*ir_types.PointerType).ElemType, src)
		c.Block.AddInstruction(load)
		src = load
		srcType = TypeIR(ast_types.GetElementType(s.Parent.Type()))
	}
	for _, i := range s.Indexes {
		indexes = append(indexes, c.AutoLoad(ExpressionIR(c, i)))
	}
	v := instruction.NewGetElementPtr(srcType, src, indexes...)
	c.Block.AddInstruction(v)
	return v
}

func SubscriptingConstIR(p *Program, s *expression.Subscripting) constant.Constant {
	if t, ok := s.Parent.Type().(*ast_types.TypeArray); ok {
		indexes := []constant.Constant{constant.NewInt(ir_types.I32, 0)}
		for _, i := range s.Indexes {
			indexes = append(indexes, ExpressionConstIR(p, i))
		}
		return ir_expression.NewExprGetElementPtr(TypeIR(t), ExpressionConstIR(p, s.Parent), indexes...)
	}
	return nil
}
