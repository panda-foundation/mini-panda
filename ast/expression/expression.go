package expression

type ExpressionBase struct {
	core.NodeBase
	Const bool
	Typ   core.Type
}

func (b *ExpressionBase) IsConstant() bool {
	return b.Const
}

func (b *ExpressionBase) Type() core.Type {
	return b.Typ
}
