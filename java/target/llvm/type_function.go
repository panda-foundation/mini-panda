package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

func TypeFunctionIR(f *ast.TypeFunction) ir.Type {
	var types []ir.Type
	for _, param := range f.Parameters {
		types = append(types, TypeIR(param))
	}
	var t ir.Type = ir.Void
	if f.ReturnType != nil {
		t = TypeIR(f.ReturnType)
	}
	return ir.NewPointerType(ir.NewFuncType(t, types...))
}
