package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/target/llvm/llvm"
)

func InitializerIR(c llvm.Context, i *ast.Initializer) ir.Value {
	return InitializerConstIR(c.Program, i)
}

func InitializerConstIR(p llvm.Program, i *ast.Initializer) ir.Constant {
	values := []ir.Constant{}
	for _, e := range i.Expressions {
		values = append(values, ExpressionConstIR(p, e))
	}
	t := i.Type()
	if array, ok := t.(*ast.TypeArray); ok {
		return ir.NewArray(TypeArrayIR(array).(*ir.ArrayType), values...)
	} else if n, ok := t.(*ast.TypeName); ok {
		return ir.NewStruct(TypeNameIR(n).(*ir.StructType), values...)
	}
	return nil
}
