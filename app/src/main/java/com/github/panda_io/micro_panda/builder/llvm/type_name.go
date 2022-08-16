package llvm

import (
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

func TypeNameIR(t *ast_types.TypeName) ir.Type {
	if t.IsEnum {
		return ir_types.UI8
	} else {
		return ir_types.NewStructType(t.Qualified)
	}
}
