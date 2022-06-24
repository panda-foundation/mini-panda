package ast_types

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/token"
)

type TypeBuiltin struct {
	TypeBase
	Token token.Token
}

func (b *TypeBuiltin) Equal(t ast.Type) bool {
	if t, ok := t.(*TypeBuiltin); ok {
		return t.Token == b.Token
	}
	return false
}

func (b *TypeBuiltin) String() string {
	return b.Token.String()
}
