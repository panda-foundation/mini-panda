package llvm

import (
	"strconv"

	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
	"github.com/panda-io/micro-panda/token"
)

func LiteralIR(c *Context, l *expression.Literal) ir.Value {
	return LiteralConstIR(c.Program, l)
}

func LiteralConstIR(p *Program, l *expression.Literal) constant.Constant {
	switch l.Token {
	case token.CHAR:
		r := []rune(l.Value[1 : len(l.Value)-1])
		return constant.NewInt(ir_types.UI8, int64(r[0]))
	case token.STRING:
		if l.Value[0] == '"' {
			// string
			str, _ := strconv.Unquote(l.Value)
			return p.AddString(str)
		} else {
			// `` raw string
			return p.AddString(l.Value[1 : len(l.Value)-1])
		}
	case token.FLOAT:
		return constant.NewFloatFromString(TypeIR(l.Type()).(*ir_types.FloatType), l.Value)
	case token.INT:
		return constant.NewIntFromString(TypeIR(l.Type()).(*ir_types.IntType), l.Value)
	case token.BOOL:
		if l.Value == "true" {
			return constant.True
		}
		return constant.False
	case token.NULL:
		return constant.NewNull(TypeIR(l.Type()).(*ir_types.PointerType))
	default:
		return nil
	}
}
