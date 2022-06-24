package types

import (
	ast_types "github.com/panda-io/micro-panda/ast/types"
	ir_core "github.com/panda-io/micro-panda/ir/core"
	ir_types "github.com/panda-io/micro-panda/ir/types"
)

func TypeNameIR(t *ast_types.TypeName) ir_core.Type {
	if t.IsEnum {
		return ir_types.UI8
	} else {
		return ir_types.NewStructType(t.Qualified)
	}
}
