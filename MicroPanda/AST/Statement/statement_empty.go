package ast

type Empty struct {
	StatementBase
}

func (*Empty) Validate(c *Context) {
}
