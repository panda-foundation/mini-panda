package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/target/llvm"
)

func BlockIR(c llvm.Context, b *ast.Block) {
	for _, stmt := range b.Statements {
		ctx := c
		if _, ok := stmt.(*ast.Block); ok {
			ctx = c.NewContext()
		}
		StatementIR(ctx, stmt)
		if ctx.Block.Terminated {
			return
		}
	}
}
