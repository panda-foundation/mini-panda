package expression

import (
	"github.com/panda-io/micro-panda/ast/core"
)

type Sizeof struct {
	ExpressionBase
	Target core.Type
}

func (s *Sizeof) Validate(c core.Context, expected core.Type) {
	s.Target = c.ResolveType(s.Target)
	s.Typ = core.TypeU32
	s.Const = true
}
