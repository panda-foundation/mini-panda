package expression

import (
	"github.com/panda-io/micro-panda/ast/core"
)

type Subscripting struct {
	ExpressionBase
	Parent  core.Expression
	Indexes []core.Expression
}

func (s *Subscripting) Validate(c core.Context, expected core.Type) {
	s.Const = false
	s.Parent.Validate(c, nil)
	if t, ok := s.Parent.Type().(*core.TypeArray); ok {
		if len(s.Indexes) == len(t.Dimension) {
			s.Typ = t.ElementType
			for _, e := range s.Indexes {
				e.Validate(c, nil)
				if !core.IsInteger(e.Type()) {
					c.Error(e.GetPosition(), "expect integer index for array")
				}
			}
		} else if len(s.Indexes) < len(t.Dimension) {
			array := &core.TypeArray{
				ElementType: t.ElementType,
				Dimension:   []int{0},
			}
			for _, e := range s.Indexes {
				e.Validate(c, nil)
				if !core.IsInteger(e.Type()) {
					c.Error(e.GetPosition(), "expect integer index for array")
				}
			}
			for i := len(t.Dimension) - len(s.Indexes) - 1; i > 0; i-- {
				array.Dimension = append(array.Dimension, t.Dimension[len(t.Dimension)-1])
			}
			s.Typ = array
		} else {
			c.Error(s.Position, "mismatch array dimension")
		}
	} else {
		c.Error(s.Position, "expect array type")
	}
}
