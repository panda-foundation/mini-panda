package expression

import "github.com/panda-io/micro-panda/ast/ast"

type ExpressionBase struct {
	ast.NodeBase
	Const bool
	Typ   ast.Type
}

func (b *ExpressionBase) IsConstant() bool {
	return b.Const
}

func (b *ExpressionBase) Type() ast.Type {
	return b.Typ
}
