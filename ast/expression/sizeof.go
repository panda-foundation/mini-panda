package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/ast/types"
)

type Sizeof struct {
	ExpressionBase
	Target core.Type
}

func (s *Sizeof) Validate(c ast.Context, expected core.Type) {
	s.Target = c.ResolveType(s.Target)
	s.Typ = types.TypeU32
	s.Const = true
}
