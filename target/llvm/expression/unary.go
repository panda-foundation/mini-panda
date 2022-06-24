package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/ir/instruction"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
	"github.com/panda-io/micro-panda/target/llvm/llvm"
	"github.com/panda-io/micro-panda/target/llvm/types"
	"github.com/panda-io/micro-panda/token"
)

func UnaryIR(c llvm.Context, u *ast.Unary) ir.Value {
	t := types.TypeIR(u.Expression.Type())
	v := ExpressionIR(c, u.Expression)

	var inst ir.Instruction
	switch u.Operator {
	case token.Plus:
		return c.AutoLoad(v)

	case token.Minus:
		if ir_types.IsFloat(t) {
			inst = ir.NewFSub(ir.NewFloat(t.(*ir.FloatType), 0), c.AutoLoad(v))
		} else if ir_types.IsInt(t) {
			inst = ir.NewSub(ir.NewInt(t.(*ir.IntType), 0), c.AutoLoad(v))
		}

	case token.Not:
		inst = instruction.NewXor(c.AutoLoad(v), ir.True)

	case token.Complement:
		inst = instruction.NewXor(c.AutoLoad(v), ir.NewInt(t.(*ir.IntType), -1))

	case token.BitAnd:
		alloca := instruction.NewAlloca(ir.NewPointerType(t))
		c.Block().AddInstruction(alloca)
		c.Block().AddInstruction(ir.NewStore(v, alloca))
		return alloca
	}
	if inst == nil {
		return nil
	}
	c.Block().AddInstruction(inst)
	return inst.(ir.Value)
}

func UnaryConstIR(p llvm.Program, u *ast.Unary) ir.Constant {
	v := ExpressionConstIR(p, u.Expression)
	t := v.Type()

	switch u.Operator {
	case token.Plus:
		if ir_types.IsNumber(t) {
			return v
		}

	case token.Minus:
		if ir_types.IsFloat(t) {
			return ir.NewExprFSub(ir.NewFloat(t.(*ir.FloatType), 0), v)
		} else if ir_types.IsInt(t) {
			return ir.NewExprSub(ir.NewInt(t.(*ir.IntType), 0), v)
		}

	case token.Not:
		if ir_types.IsBool(t) {
			return ir.NewExprXor(v, ir.True)
		}

	case token.Complement:
		if ir_types.IsInt(t) {
			return ir.NewExprXor(v, ir.NewInt(t.(*ir.IntType), -1))
		}

	case token.BitAnd:
		// must be a global
		if g, ok := v.(*ir.Global); ok {
			if g.Immutable {
				return g
			}
		}
	}
	return nil
}
