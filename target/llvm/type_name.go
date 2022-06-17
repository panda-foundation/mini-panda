package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

func TypeNameIR(t *ast.TypeName) ir.Type {
	if t.IsEnum {
		return ir.UI8
	} else {
		s := ir.NewStructType()
		s.TypeName = t.Qualified
		return s
	}
}
