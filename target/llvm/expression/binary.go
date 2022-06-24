package expression

import (
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	ir_expression "github.com/panda-io/micro-panda/target/llvm/ir/constant/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
	"github.com/panda-io/micro-panda/target/llvm/llvm"
	"github.com/panda-io/micro-panda/target/llvm/types"
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

func BinaryIR(c llvm.Context, b *expression.Binary) ir.Value {
	t := types.TypeIR(b.Left.Type())
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

	var inst instruction.Instruction
	var ret ir.Value
	switch b.Operator {
	case token.Assign:
		if ast_types.IsPointer(b.Left.Type()) && ast_types.IsArray(b.Right.Type()) {
			var gep instruction.Instruction = instruction.NewGetElementPtr(types.TypeIR(b.Right.Type()), v2, constant.NewInt(ir_types.I32, 0), constant.NewInt(ir_types.I32, 0))
			c.Block().AddInstruction(gep)
			inst = instruction.NewStore(gep.(ir.Value), v1)
		} else {
			inst = instruction.NewStore(c.AutoLoad(v2), v1)
		}

	case token.PlusAssign:
		if ir_types.IsInt(t) {
			inst = instruction.NewAdd(c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir_types.IsFloat(t) {
			inst = instruction.NewFAdd(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		c.Block().AddInstruction(inst)
		inst = instruction.NewStore(inst.(ir.Value), v1)

	case token.MinusAssign:
		if ir_types.IsInt(t) {
			inst = instruction.NewSub(c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir_types.IsFloat(t) {
			inst = instruction.NewFSub(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		c.Block().AddInstruction(inst)
		inst = instruction.NewStore(inst.(ir.Value), v1)

	case token.MulAssign:
		if ir_types.IsInt(t) {
			inst = instruction.NewMul(c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir_types.IsFloat(t) {
			inst = instruction.NewFMul(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		c.Block().AddInstruction(inst)
		inst = instruction.NewStore(inst.(ir.Value), v1)

	case token.DivAssign:
		if ir_types.IsInt(t) {
			if t.(*ir_types.IntType).Unsigned {
				inst = instruction.NewUDiv(c.AutoLoad(v1), c.AutoLoad(v2))
			} else {
				inst = instruction.NewSDiv(c.AutoLoad(v1), c.AutoLoad(v2))
			}
		} else if ir_types.IsFloat(t) {
			inst = instruction.NewFDiv(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		c.Block().AddInstruction(inst)
		inst = instruction.NewStore(inst.(ir.Value), v1)
		ret = v1

	case token.RemAssign:
		if t.(*ir_types.IntType).Unsigned {
			inst = instruction.NewURem(c.AutoLoad(v1), c.AutoLoad(v2))
		} else {
			inst = instruction.NewSRem(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		c.Block().AddInstruction(inst)
		inst = instruction.NewStore(inst.(ir.Value), v1)
		ret = v1

	case token.LeftShiftAssign:
		inst = instruction.NewShl(c.AutoLoad(v1), c.AutoLoad(v2))
		c.Block().AddInstruction(inst)
		inst = instruction.NewStore(inst.(ir.Value), v1)
		ret = v1

	case token.RightShiftAssign:
		if t.(*ir_types.IntType).Unsigned {
			inst = instruction.NewLShr(c.AutoLoad(v1), c.AutoLoad(v2))
		} else {
			inst = instruction.NewAShr(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		c.Block().AddInstruction(inst)
		inst = instruction.NewStore(inst.(ir.Value), v1)
		ret = v1

	case token.OrAssign:
		inst = instruction.NewOr(c.AutoLoad(v1), c.AutoLoad(v2))
		c.Block().AddInstruction(inst)
		inst = instruction.NewStore(inst.(ir.Value), v1)
		ret = v1

	case token.XorAssign:
		inst = instruction.NewXor(c.AutoLoad(v1), c.AutoLoad(v2))
		c.Block().AddInstruction(inst)
		inst = instruction.NewStore(inst.(ir.Value), v1)
		ret = v1

	case token.AndAssign:
		inst = instruction.NewAnd(c.AutoLoad(v1), c.AutoLoad(v2))
		c.Block().AddInstruction(inst)
		inst = instruction.NewStore(inst.(ir.Value), v1)
		ret = v1

	case token.Or, token.And:
		if b.Operator == token.Or {
			inst = instruction.NewOr(c.AutoLoad(v1), c.AutoLoad(v2))
		} else {
			inst = instruction.NewAnd(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		ret = inst.(ir.Value)

	case token.BitOr:
		inst = instruction.NewOr(c.AutoLoad(v1), c.AutoLoad(v2))
		ret = inst.(ir.Value)

	case token.BitXor:
		inst = instruction.NewXor(c.AutoLoad(v1), c.AutoLoad(v2))
		ret = inst.(ir.Value)

	case token.BitAnd:

		inst = instruction.NewAnd(c.AutoLoad(v1), c.AutoLoad(v2))
		ret = inst.(ir.Value)

	case token.Equal, token.NotEqual, token.Less, token.LessEqual, token.Greater, token.GreaterEqual:
		if ir_types.IsInt(t) {
			var icmp ir.IPred
			if t.(*ir_types.IntType).Unsigned {
				icmp = UICMP[b.Operator]
			} else {
				icmp = ICMP[b.Operator]
			}
			inst = instruction.NewICmp(icmp, c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir_types.IsFloat(t) {
			fmp := FCMP[b.Operator]
			inst = instruction.NewFCmp(fmp, c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir_types.IsPointer(t) {
			if b.Operator == token.Equal || b.Operator == token.NotEqual {
				icmp := ICMP[b.Operator]
				inst = instruction.NewICmp(icmp, c.AutoLoad(v1), c.AutoLoad(v2))
			}
		}
		ret = inst.(ir.Value)

	case token.LeftShift:
		inst = instruction.NewShl(c.AutoLoad(v1), c.AutoLoad(v2))
		ret = inst.(ir.Value)

	case token.RightShift:
		if t.(*ir_types.IntType).Unsigned {
			inst = instruction.NewLShr(c.AutoLoad(v1), c.AutoLoad(v2))
		} else {
			inst = instruction.NewAShr(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		ret = inst.(ir.Value)

	case token.Plus:
		if ir_types.IsInt(t) {
			inst = instruction.NewAdd(c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir_types.IsFloat(t) {
			inst = instruction.NewFAdd(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		ret = inst.(ir.Value)

	case token.Minus:
		if ir_types.IsInt(t) {
			inst = instruction.NewSub(c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir_types.IsFloat(t) {
			inst = instruction.NewFSub(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		ret = inst.(ir.Value)

	case token.Mul:
		if ir_types.IsInt(t) {
			inst = instruction.NewMul(c.AutoLoad(v1), c.AutoLoad(v2))
		} else if ir_types.IsFloat(t) {
			inst = instruction.NewFMul(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		ret = inst.(ir.Value)

	case token.Div:
		if ir_types.IsInt(t) {
			if t.(*ir_types.IntType).Unsigned {
				inst = instruction.NewUDiv(c.AutoLoad(v1), c.AutoLoad(v2))
			} else {
				inst = instruction.NewSDiv(c.AutoLoad(v1), c.AutoLoad(v2))
			}
		} else if ir_types.IsFloat(t) {
			inst = instruction.NewFDiv(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		ret = inst.(ir.Value)

	case token.Rem:
		if t.(*ir_types.IntType).Unsigned {
			inst = instruction.NewURem(c.AutoLoad(v1), c.AutoLoad(v2))
		} else {
			inst = instruction.NewSRem(c.AutoLoad(v1), c.AutoLoad(v2))
		}
		ret = inst.(ir.Value)
	}
	c.Block().AddInstruction(inst)
	return ret
}

func BinaryConstIR(p llvm.Program, b *expression.Binary) constant.Constant {
	v1 := ExpressionConstIR(p, b.Left)
	v2 := ExpressionConstIR(p, b.Right)
	t := types.TypeIR(b.Left.Type())

	switch b.Operator {
	case token.Assign, token.MulAssign, token.DivAssign, token.RemAssign, token.PlusAssign, token.MinusAssign,
		token.LeftShiftAssign, token.RightShiftAssign, token.AndAssign, token.OrAssign, token.XorAssign:
		return nil

	case token.Or, token.And:
		if b.Operator == token.Or {
			return ir_expression.NewExprOr(v1, v2)
		} else {
			return ir_expression.NewExprAnd(v1, v2)
		}

	case token.BitOr:
		return ir_expression.NewExprOr(v1, v2)

	case token.BitXor:
		return ir_expression.NewExprXor(v1, v2)

	case token.BitAnd:
		return ir_expression.NewExprAnd(v1, v2)

	case token.Equal, token.NotEqual, token.Less, token.LessEqual, token.Greater, token.GreaterEqual:
		if ir_types.IsInt(t) {
			var icmp ir.IPred
			if t.(*ir_types.IntType).Unsigned {
				icmp = UICMP[b.Operator]
			} else {
				icmp = ICMP[b.Operator]
			}
			return ir_expression.NewExprICmp(icmp, v1, v2)
		} else if ir_types.IsFloat(t) {
			fmp := FCMP[b.Operator]
			return ir_expression.NewExprFCmp(fmp, v1, v2)
		} else if ir_types.IsPointer(t) {
			if b.Operator == token.Equal || b.Operator == token.NotEqual {
				icmp := ICMP[b.Operator]
				return ir_expression.NewExprICmp(icmp, v1, v2)
			}
		}

	case token.LeftShift:
		return ir_expression.NewExprShl(v1, v2)

	case token.RightShift:
		if t.(*ir_types.IntType).Unsigned {
			return ir_expression.NewExprLShr(v1, v2)
		} else {
			return ir_expression.NewExprAShr(v1, v2)
		}

	case token.Plus:
		if ir_types.IsInt(t) {
			return ir_expression.NewExprAdd(v1, v2)
		} else if ir_types.IsFloat(t) {
			return ir_expression.NewExprFAdd(v1, v2)
		}

	case token.Minus:
		if ir_types.IsInt(t) {
			return ir_expression.NewExprSub(v1, v2)
		} else if ir_types.IsFloat(t) {
			return ir_expression.NewExprFSub(v1, v2)
		}

	case token.Mul:
		if ir_types.IsInt(t) {
			return ir_expression.NewExprMul(v1, v2)
		} else if ir_types.IsFloat(t) {
			return ir_expression.NewExprFMul(v1, v2)
		}

	case token.Div:
		if ir_types.IsInt(t) {
			if t.(*ir_types.IntType).Unsigned {
				return ir_expression.NewExprUDiv(v1, v2)
			} else {
				return ir_expression.NewExprSDiv(v1, v2)
			}
		} else if ir_types.IsFloat(t) {
			return ir_expression.NewExprFDiv(v1, v2)
		}

	case token.Rem:
		if t.(*ir_types.IntType).Unsigned {
			return ir_expression.NewExprURem(v1, v2)
		} else {
			return ir_expression.NewExprSRem(v1, v2)
		}
	}

	return nil
}
