package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

func TypeIR(typ ast.Type) ir.Type {
	switch t := typ.(type) {
	case *ast.TypeBuiltin:
		return TypeBuiltinIR(t)

	case *ast.TypeName:
		return TypeNameIR(t)

	case *ast.TypePointer:
		return TypePointerIR(t)

	case *ast.TypeArray:
		return TypeArrayIR(t)

	case *ast.TypeFunction:
		return TypeFunctionIR(t)
	}
	return nil
}

func ParamIR(parameter *ast.Parameter) *ir.Param {
	var param *ir.Param
	var paramType ir.Type
	switch t := parameter.Type.(type) {
	case *ast.TypeBuiltin:
		paramType = TypeBuiltinIR(t)

	case *ast.TypeName:
		paramType = TypeNameIR(t)

	case *ast.TypePointer:
		paramType = TypePointerIR(t)

	case *ast.TypeArray:
		paramType = TypeArrayIR(t)

	case *ast.TypeFunction:
		paramType = TypeFunctionIR(t)
	}
	param = ir.NewParam(paramType)
	param.LocalName = parameter.Name
	return param
}

func StructIR(qualified string) *ir.StructType {
	t := ir.NewStructType()
	t.TypeName = qualified
	return t
}

func StructPointerIR(qualified string) *ir.PointerType {
	t := ir.NewStructType()
	t.TypeName = qualified
	return ir.NewPointerType(t)
}
