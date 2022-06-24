package llvm

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/ast/declaration"
	"github.com/panda-io/micro-panda/target/llvm/ir"
	ir_core "github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

func TypeIR(typ ast.Type) ir_core.Type {
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
