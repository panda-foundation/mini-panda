package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/token"
)

var (
	ICMP = map[token.Token]ir.IPred{
		token.Equal:        ir.IPredEQ,
		token.NotEqual:     ir.IPredNE,
		token.Less:         ir.IPredSLT,
		token.LessEqual:    ir.IPredSLE,
		token.Greater:      ir.IPredSGT,
		token.GreaterEqual: ir.IPredSGE,
	}

	UICMP = map[token.Token]ir.IPred{
		token.Equal:        ir.IPredEQ,
		token.NotEqual:     ir.IPredNE,
		token.Less:         ir.IPredULT,
		token.LessEqual:    ir.IPredULE,
		token.Greater:      ir.IPredUGT,
		token.GreaterEqual: ir.IPredUGE,
	}

	FCMP = map[token.Token]ir.FPred{
		token.Equal:        ir.FPredOEQ,
		token.NotEqual:     ir.FPredONE,
		token.Less:         ir.FPredOLT,
		token.LessEqual:    ir.FPredOLE,
		token.Greater:      ir.FPredOGT,
		token.GreaterEqual: ir.FPredOGE,
	}
)

func BinaryIR(c *Context, b *ast.Binary) ir.Value {
	t := TypeIR(b.Left.Type())
	var v1 ir.Value
	var v2 ir.Value
	if b.Left.IsConstant() {
		v1 = ExpressionConstIR(c.Program, b.Left)
	} else {
		v1 = ExpressionIR(c, b.Left)
	}
	if b.Right.IsConstant() {
		v2 = ExpressionConstIR(c.Program, b.Right)
	} else {
		v2 = ExpressionIR(c, b.Right)
	}

	var inst ir.Instruction
	var ret ir.Value
	switch b.Operator {
	case token.Assign:
		if ast.IsPointer(b.Left.Type()) && ast.IsArray(b.Right.Type()) {
			var gep ir.Instruction = ir.NewGetElementPtr(TypeIR(b.Right.Type()), v2, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, 0))
			c.Block.AddInstruction(gep)
			inst = ir.NewStore(gep.(ir.Value), v1)
		} else {
			inst = ir.NewStore(c.AutoLoad(v2), v1)
		}

	case token.PlusAssign:
		if ir.IsInt(t) {
			inst = ir.NewAdd(c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir.IsFloat(t) {
			inst = ir.NewFAdd(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		c.Block.AddInstruction(inst)
		inst = ir.NewStore(inst.(ir.Value), v1)

	case token.MinusAssign:
		if ir.IsInt(t) {
			inst = ir.NewSub(c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir.IsFloat(t) {
			inst = ir.NewFSub(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		c.Block.AddInstruction(inst)
		inst = ir.NewStore(inst.(ir.Value), v1)

	case token.MulAssign:
		if ir.IsInt(t) {
			inst = ir.NewMul(c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir.IsFloat(t) {
			inst = ir.NewFMul(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		c.Block.AddInstruction(inst)
		inst = ir.NewStore(inst.(ir.Value), v1)

	case token.DivAssign:
		if ir.IsInt(t) {
			if t.(*ir.IntType).Unsigned {
				inst = ir.NewUDiv(c.AutoLoad(v1), c.AutoLoad(v2))
			} else {
				inst = ir.NewSDiv(c.AutoLoad(v1), c.AutoLoad(v2))
			}
		} else if ir.IsFloat(t) {
			inst = ir.NewFDiv(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		c.Block.AddInstruction(inst)
		inst = ir.NewStore(inst.(ir.Value), v1)
		ret = v1

	case token.RemAssign:
		if t.(*ir.IntType).Unsigned {
			inst = ir.NewURem(c.AutoLoad(v1), c.AutoLoad(v2))
		} else {
			inst = ir.NewSRem(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		c.Block.AddInstruction(inst)
		inst = ir.NewStore(inst.(ir.Value), v1)
		ret = v1

	case token.LeftShiftAssign:
		inst = ir.NewShl(c.AutoLoad(v1), c.AutoLoad(v2))
		c.Block.AddInstruction(inst)
		inst = ir.NewStore(inst.(ir.Value), v1)
		ret = v1

	case token.RightShiftAssign:
		if t.(*ir.IntType).Unsigned {
			inst = ir.NewLShr(c.AutoLoad(v1), c.AutoLoad(v2))
		} else {
			inst = ir.NewAShr(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		c.Block.AddInstruction(inst)
		inst = ir.NewStore(inst.(ir.Value), v1)
		ret = v1

	case token.OrAssign:
		inst = ir.NewOr(c.AutoLoad(v1), c.AutoLoad(v2))
		c.Block.AddInstruction(inst)
		inst = ir.NewStore(inst.(ir.Value), v1)
		ret = v1

	case token.XorAssign:
		inst = ir.NewXor(c.AutoLoad(v1), c.AutoLoad(v2))
		c.Block.AddInstruction(inst)
		inst = ir.NewStore(inst.(ir.Value), v1)
		ret = v1

	case token.AndAssign:
		inst = ir.NewAnd(c.AutoLoad(v1), c.AutoLoad(v2))
		c.Block.AddInstruction(inst)
		inst = ir.NewStore(inst.(ir.Value), v1)
		ret = v1

	case token.Or, token.And:
		if b.Operator == token.Or {
			inst = ir.NewOr(c.AutoLoad(v1), c.AutoLoad(v2))
		} else {
			inst = ir.NewAnd(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		ret = inst.(ir.Value)

	case token.BitOr:
		inst = ir.NewOr(c.AutoLoad(v1), c.AutoLoad(v2))
		ret = inst.(ir.Value)

	case token.BitXor:
		inst = ir.NewXor(c.AutoLoad(v1), c.AutoLoad(v2))
		ret = inst.(ir.Value)

	case token.BitAnd:

		inst = ir.NewAnd(c.AutoLoad(v1), c.AutoLoad(v2))
		ret = inst.(ir.Value)

	case token.Equal, token.NotEqual, token.Less, token.LessEqual, token.Greater, token.GreaterEqual:
		if ir.IsInt(t) {
			var icmp ir.IPred
			if t.(*ir.IntType).Unsigned {
				icmp = UICMP[b.Operator]
			} else {
				icmp = ICMP[b.Operator]
			}
			inst = ir.NewICmp(icmp, c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir.IsFloat(t) {
			fmp := FCMP[b.Operator]
			inst = ir.NewFCmp(fmp, c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir.IsPointer(t) {
			if b.Operator == token.Equal || b.Operator == token.NotEqual {
				icmp := ICMP[b.Operator]
				inst = ir.NewICmp(icmp, c.AutoLoad(v1), c.AutoLoad(v2))
			}
		}
		ret = inst.(ir.Value)

	case token.LeftShift:
		inst = ir.NewShl(c.AutoLoad(v1), c.AutoLoad(v2))
		ret = inst.(ir.Value)

	case token.RightShift:
		if t.(*ir.IntType).Unsigned {
			inst = ir.NewLShr(c.AutoLoad(v1), c.AutoLoad(v2))
		} else {
			inst = ir.NewAShr(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		ret = inst.(ir.Value)

	case token.Plus:
		if ir.IsInt(t) {
			inst = ir.NewAdd(c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir.IsFloat(t) {
			inst = ir.NewFAdd(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		ret = inst.(ir.Value)

	case token.Minus:
		if ir.IsInt(t) {
			inst = ir.NewSub(c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir.IsFloat(t) {
			inst = ir.NewFSub(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		ret = inst.(ir.Value)

	case token.Mul:
		if ir.IsInt(t) {
			inst = ir.NewMul(c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir.IsFloat(t) {
			inst = ir.NewFMul(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		ret = inst.(ir.Value)

	case token.Div:
		if ir.IsInt(t) {
			if t.(*ir.IntType).Unsigned {
				inst = ir.NewUDiv(c.AutoLoad(v1), c.AutoLoad(v2))
			} else {
				inst = ir.NewSDiv(c.AutoLoad(v1), c.AutoLoad(v2))
			}
		} else if ir.IsFloat(t) {
			inst = ir.NewFDiv(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		ret = inst.(ir.Value)

	case token.Rem:
		if t.(*ir.IntType).Unsigned {
			inst = ir.NewURem(c.AutoLoad(v1), c.AutoLoad(v2))
		} else {
			inst = ir.NewSRem(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		ret = inst.(ir.Value)
	}
	c.Block.AddInstruction(inst)
	return ret
}

func BinaryConstIR(p *Program, b *ast.Binary) ir.Constant {
	v1 := ExpressionConstIR(p, b.Left)
	v2 := ExpressionConstIR(p, b.Right)
	t := TypeIR(b.Left.Type())

	switch b.Operator {
	case token.Assign, token.MulAssign, token.DivAssign, token.RemAssign, token.PlusAssign, token.MinusAssign,
		token.LeftShiftAssign, token.RightShiftAssign, token.AndAssign, token.OrAssign, token.XorAssign:
		return nil

	case token.Or, token.And:
		if b.Operator == token.Or {
			return ir.NewExprOr(v1, v2)
		} else {
			return ir.NewExprAnd(v1, v2)
		}

	case token.BitOr:
		return ir.NewExprOr(v1, v2)

	case token.BitXor:
		return ir.NewExprXor(v1, v2)

	case token.BitAnd:
		return ir.NewExprAnd(v1, v2)

	case token.Equal, token.NotEqual, token.Less, token.LessEqual, token.Greater, token.GreaterEqual:
		if ir.IsInt(t) {
			var icmp ir.IPred
			if t.(*ir.IntType).Unsigned {
				icmp = UICMP[b.Operator]
			} else {
				icmp = ICMP[b.Operator]
			}
			return ir.NewExprICmp(icmp, v1, v2)
		} else if ir.IsFloat(t) {
			fmp := FCMP[b.Operator]
			return ir.NewExprFCmp(fmp, v1, v2)
		} else if ir.IsPointer(t) {
			if b.Operator == token.Equal || b.Operator == token.NotEqual {
				icmp := ICMP[b.Operator]
				return ir.NewExprICmp(icmp, v1, v2)
			}
		}

	case token.LeftShift:
		return ir.NewExprShl(v1, v2)

	case token.RightShift:
		if t.(*ir.IntType).Unsigned {
			return ir.NewExprLShr(v1, v2)
		} else {
			return ir.NewExprAShr(v1, v2)
		}

	case token.Plus:
		if ir.IsInt(t) {
			return ir.NewExprAdd(v1, v2)
		} else if ir.IsFloat(t) {
			return ir.NewExprFAdd(v1, v2)
		}

	case token.Minus:
		if ir.IsInt(t) {
			return ir.NewExprSub(v1, v2)
		} else if ir.IsFloat(t) {
			return ir.NewExprFSub(v1, v2)
		}

	case token.Mul:
		if ir.IsInt(t) {
			return ir.NewExprMul(v1, v2)
		} else if ir.IsFloat(t) {
			return ir.NewExprFMul(v1, v2)
		}

	case token.Div:
		if ir.IsInt(t) {
			if t.(*ir.IntType).Unsigned {
				return ir.NewExprUDiv(v1, v2)
			} else {
				return ir.NewExprSDiv(v1, v2)
			}
		} else if ir.IsFloat(t) {
			return ir.NewExprFDiv(v1, v2)
		}

	case token.Rem:
		if t.(*ir.IntType).Unsigned {
			return ir.NewExprURem(v1, v2)
		} else {
			return ir.NewExprSRem(v1, v2)
		}
	}

	return nil
}
