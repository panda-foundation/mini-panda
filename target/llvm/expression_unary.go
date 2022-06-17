package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/token"
)

func UnaryIR(c *Context, u *ast.Unary) ir.Value {
	t := TypeIR(u.Expression.Type())
	v := ExpressionIR(c, u.Expression)

	var inst ir.Instruction
	switch u.Operator {
	case token.Plus:
		return c.AutoLoad(v)

	case token.Minus:
		if ir.IsFloat(t) {
			inst = ir.NewFSub(ir.NewFloat(t.(*ir.FloatType), 0), c.AutoLoad(v))
		} else if ir.IsInt(t) {
			inst = ir.NewSub(ir.NewInt(t.(*ir.IntType), 0), c.AutoLoad(v))
		}

	case token.Not:
		inst = ir.NewXor(c.AutoLoad(v), ir.True)

	case token.Complement:
		inst = ir.NewXor(c.AutoLoad(v), ir.NewInt(t.(*ir.IntType), -1))

	case token.BitAnd:
		alloca := ir.NewAlloca(ir.NewPointerType(t))
		c.Block.AddInstruction(alloca)
		c.Block.AddInstruction(ir.NewStore(v, alloca))
		return alloca
	}
	if inst == nil {
		return nil
	}
	c.Block.AddInstruction(inst)
	return inst.(ir.Value)
}

func UnaryConstIR(p *Program, u *ast.Unary) ir.Constant {
	v := ExpressionConstIR(p, u.Expression)
	t := v.Type()

	switch u.Operator {
	case token.Plus:
		if ir.IsNumber(t) {
			return v
		}

	case token.Minus:
		if ir.IsFloat(t) {
			return ir.NewExprFSub(ir.NewFloat(t.(*ir.FloatType), 0), v)
		} else if ir.IsInt(t) {
			return ir.NewExprSub(ir.NewInt(t.(*ir.IntType), 0), v)
		}

	case token.Not:
		if ir.IsBool(t) {
			return ir.NewExprXor(v, ir.True)
		}

	case token.Complement:
		if ir.IsInt(t) {
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
