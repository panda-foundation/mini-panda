package statement

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ast/core"
)

type Block struct {
	StatementBase
	Statements []core.Statement
}

func (b *Block) Validate(c ast.Context) {
	//TO-DO warning: unreachable code //Start, End of block
	for _, statement := range b.Statements {
		ctx := c
		if _, ok := statement.(*Block); ok {
			ctx = c.NewContext()
		}
		statement.Validate(ctx)
	}
}
