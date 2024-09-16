package ast

type Expression interface {
	Node
	IsConstant() bool
	Type() Type
	Validate(c *Context, expected Type)
}

type ExpressionBase struct {
	NodeBase
	Const bool
	Typ   Type
}

func (b *ExpressionBase) IsConstant() bool {
	return b.Const
}

func (b *ExpressionBase) Type() Type {
	return b.Typ
}
