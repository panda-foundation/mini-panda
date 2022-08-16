package llvm

import (
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

func InitializerIR(c *Context, i *expression.Initializer) ir.Value {
	return InitializerConstIR(c.Program, i)
}

func InitializerConstIR(p *Program, i *expression.Initializer) constant.Constant {
	values := []constant.Constant{}
	for _, e := range i.Expressions {
		values = append(values, ExpressionConstIR(p, e))
	}
	t := i.Type()
	if array, ok := t.(*ast_types.TypeArray); ok {
		return constant.NewArray(TypeArrayIR(array).(*ir_types.ArrayType), values...)
	} else if n, ok := t.(*ast_types.TypeName); ok {
		return constant.NewStruct(TypeNameIR(n).(*ir_types.StructType), values...)
	}
	return nil
}
