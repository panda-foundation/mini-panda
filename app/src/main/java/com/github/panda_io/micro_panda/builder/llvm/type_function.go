package llvm

import (
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

func TypeFunctionIR(f *ast_types.TypeFunction) ir.Type {
	var types []ir.Type
	for _, param := range f.Parameters {
		types = append(types, TypeIR(param))
	}
	var t ir.Type = ir_types.Void
	if f.ReturnType != nil {
		t = TypeIR(f.ReturnType)
	}
	//TO-DO
	return ir_types.NewPointerType(ir_types.NewFuncType("", t, types...))
}
