package types

import (
	ast_types "github.com/panda-io/micro-panda/ast/types"
	ir_core "github.com/panda-io/micro-panda/ir/core"
	ir_types "github.com/panda-io/micro-panda/ir/types"
)

func TypeFunctionIR(f *ast_types.TypeFunction) ir_core.Type {
	var types []ir_core.Type
	for _, param := range f.Parameters {
		types = append(types, TypeIR(param))
	}
	var t ir_core.Type = ir_types.Void
	if f.ReturnType != nil {
		t = TypeIR(f.ReturnType)
	}
	//TO-DO
	return ir_types.NewPointerType(ir_types.NewFuncType("", t, types...))
}
