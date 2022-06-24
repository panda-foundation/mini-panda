package types

import (
	ast_core "github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/declaration"
	ast_types "github.com/panda-io/micro-panda/ast/types"
	"github.com/panda-io/micro-panda/ir"
	ir_core "github.com/panda-io/micro-panda/ir/core"
	ir_types "github.com/panda-io/micro-panda/ir/types"
)

func TypeIR(typ ast_core.Type) ir_core.Type {
	switch t := typ.(type) {
	case *ast_types.TypeBuiltin:
		return TypeBuiltinIR(t)

	case *ast_types.TypeName:
		return TypeNameIR(t)

	case *ast_types.TypePointer:
		return TypePointerIR(t)

	case *ast_types.TypeArray:
		return TypeArrayIR(t)

	case *ast_types.TypeFunction:
		return TypeFunctionIR(t)
	}
	return nil
}

func ParamIR(parameter *declaration.Parameter) *ir.Param {
	var param *ir.Param
	var paramType ir_core.Type
	switch t := parameter.Typ.(type) {
	case *ast_types.TypeBuiltin:
		paramType = TypeBuiltinIR(t)

	case *ast_types.TypeName:
		paramType = TypeNameIR(t)

	case *ast_types.TypePointer:
		paramType = TypePointerIR(t)

	case *ast_types.TypeArray:
		paramType = TypeArrayIR(t)

	case *ast_types.TypeFunction:
		paramType = TypeFunctionIR(t)
	}
	param = ir.NewParam(paramType)
	param.LocalName = parameter.Name
	return param
}

func StructIR(qualified string) *ir_types.StructType {
	return ir_types.NewStructType(qualified)
}

func StructPointerIR(qualified string) *ir_types.PointerType {
	t := ir_types.NewStructType(qualified)
	return ir_types.NewPointerType(t)
}
