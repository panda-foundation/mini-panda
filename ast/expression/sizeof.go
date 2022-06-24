package expression

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
)

type Sizeof struct {
	ExpressionBase
	Target ast.Type
}

func (s *Sizeof) Validate(c ast.Context, expected ast.Type) {
	s.Target = c.ResolveType(s.Target)
	s.Typ = ast_types.TypeU32
	s.Const = true
}
