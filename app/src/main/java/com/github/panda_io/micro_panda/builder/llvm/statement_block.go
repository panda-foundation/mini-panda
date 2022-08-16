package llvm

import (
	"github.com/panda-io/micro-panda/ast/statement"
)

func BlockIR(c *Context, b *statement.Block) {
	for _, stmt := range b.Statements {
		ctx := c
		if _, ok := stmt.(*statement.Block); ok {
			ctx = c.NewContext()
		}
		StatementIR(ctx, stmt)
		if ctx.Block.Terminated {
			return
		}
	}
}
