package types

import (
	ast_types "github.com/panda-io/micro-panda/ast/types"
	ir_core "github.com/panda-io/micro-panda/ir/core"
	ir_types "github.com/panda-io/micro-panda/ir/types"
	"github.com/panda-io/micro-panda/token"
)

func TypeBuiltinIR(b *ast_types.TypeBuiltin) ir_core.Type {
	switch b.Token {
	case token.Bool:
		return ir_types.I1

	case token.Int8:
		return ir_types.I8

	case token.Uint8:
		return ir_types.UI8

	case token.Int16:
		return ir_types.I16

	case token.Uint16:
		return ir_types.UI16

	case token.Int32:
		return ir_types.I32

	case token.Uint32:
		return ir_types.UI32

	case token.Int64:
		return ir_types.I64

	case token.Uint64:
		return ir_types.UI64

	case token.Float16:
		return ir_types.Float16

	case token.Float32:
		return ir_types.Float32

	case token.Float64:
		return ir_types.Float64

	case token.Void:
		return ir_types.Void

	default:
		return nil
	}
}
