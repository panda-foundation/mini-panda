package expression

import (
	"strconv"

	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/target/llvm/llvm"
	"github.com/panda-io/micro-panda/target/llvm/types"
	"github.com/panda-io/micro-panda/token"
)

func LiteralIR(c llvm.Context, l *ast.Literal) ir.Value {
	return LiteralConstIR(c.Program, l)
}

func LiteralConstIR(p llvm.Program, l *ast.Literal) ir.Constant {
	switch l.Token {
	case token.CHAR:
		r := []rune(l.Value[1 : len(l.Value)-1])
		return ir.NewInt(ir.UI8, int64(r[0]))
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
		return ir.NewFloatFromString(types.TypeIR(l.Type()).(*ir.FloatType), l.Value)
	case token.INT:
		return ir.NewIntFromString(types.TypeIR(l.Type()).(*ir.IntType), l.Value)
	case token.BOOL:
		if l.Value == "true" {
			return ir.True
		}
		return ir.False
	case token.NULL:
		return ir.NewNull(types.TypeIR(l.Type()).(*ir.PointerType))
	default:
		return nil
	}
}
