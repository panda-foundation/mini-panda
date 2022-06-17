package ast

type Statement interface {
	Node
	Validate(c *Context)
}

type StatementBase struct {
	NodeBase
}
