package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/token"
)

func TypeBuiltinIR(b *ast.TypeBuiltin) ir.Type {
	switch b.Token {
	case token.Bool:
		return ir.I1

	case token.Int8:
		return ir.I8

	case token.Uint8:
		return ir.UI8

	case token.Int16:
		return ir.I16

	case token.Uint16:
		return ir.UI16

	case token.Int32:
		return ir.I32

	case token.Uint32:
		return ir.UI32

	case token.Int64:
		return ir.I64

	case token.Uint64:
		return ir.UI64

	case token.Float16:
		return ir.Float16

	case token.Float32:
		return ir.Float32

	case token.Float64:
		return ir.Float64

	case token.Void:
		return ir.Void

	default:
		return nil
	}
}
