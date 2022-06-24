package llvm

import (
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	ir_expression "github.com/panda-io/micro-panda/target/llvm/ir/constant/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
	ir_core "github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
	"github.com/panda-io/micro-panda/token"
)

func UnaryIR(c *Context, u *expression.Unary) ir_core.Value {
	t := TypeIR(u.Expression.Type())
	v := ExpressionIR(c, u.Expression)

	var inst instruction.Instruction
	switch u.Operator {
	case token.Plus:
		return c.AutoLoad(v)

	case token.Minus:
		if ir_types.IsFloat(t) {
			inst = instruction.NewFSub(constant.NewFloat(t.(*ir_types.FloatType), 0), c.AutoLoad(v))
		} else if ir_types.IsInt(t) {
			inst = instruction.NewSub(constant.NewInt(t.(*ir_types.IntType), 0), c.AutoLoad(v))
		}

	case token.Not:
		inst = instruction.NewXor(c.AutoLoad(v), constant.True)

	case token.Complement:
		inst = instruction.NewXor(c.AutoLoad(v), constant.NewInt(t.(*ir_types.IntType), -1))

	case token.BitAnd:
		alloca := instruction.NewAlloca(ir_types.NewPointerType(t))
		c.Block.AddInstruction(alloca)
		c.Block.AddInstruction(instruction.NewStore(v, alloca))
		return alloca
	}
	if inst == nil {
		return nil
	}
	c.Block.AddInstruction(inst)
	return inst.(ir_core.Value)
}

func UnaryConstIR(p *Program, u *expression.Unary) constant.Constant {
	v := ExpressionConstIR(p, u.Expression)
	t := v.Type()

	switch u.Operator {
	case token.Plus:
		if ir_types.IsNumber(t) {
			return v
		}

	case token.Minus:
		if ir_types.IsFloat(t) {
			return ir_expression.NewExprFSub(constant.NewFloat(t.(*ir_types.FloatType), 0), v)
		} else if ir_types.IsInt(t) {
			return ir_expression.NewExprSub(constant.NewInt(t.(*ir_types.IntType), 0), v)
		}

	case token.Not:
		if ir_types.IsBool(t) {
			return ir_expression.NewExprXor(v, constant.True)
		}

	case token.Complement:
		if ir_types.IsInt(t) {
			return ir_expression.NewExprXor(v, constant.NewInt(t.(*ir_types.IntType), -1))
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
