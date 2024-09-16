package ast

type Block struct {
	StatementBase
	Statements []Statement
}

func (b *Block) Validate(c *Context) {
	//TO-DO warning: unreachable code //Start, End of block
	for _, statement := range b.Statements {
		ctx := c
		if _, ok := statement.(*Block); ok {
			ctx = c.NewContext()
		}
		statement.Validate(ctx)
	}
}
