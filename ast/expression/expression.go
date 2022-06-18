package expression

import (
	"github.com/panda-io/micro-panda/ast/core"
)

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

type Parameter struct {
	core.NodeBase
	Name string
	Type core.Type
}

type Arguments struct {
	core.NodeBase
	Arguments []core.Expression
}
