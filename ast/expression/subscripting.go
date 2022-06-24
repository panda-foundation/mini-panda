package expression

import (
	"fmt"

	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
)

type Subscripting struct {
	ExpressionBase
	Parent  ast.Expression
	Indexes []ast.Expression
}

func (s *Subscripting) Validate(c ast.Context, expected ast.Type) {
	s.Const = false
	s.Parent.Validate(c, nil)
	if t, ok := s.Parent.Type().(*ast_types.TypeArray); ok {
		if len(s.Indexes) == len(t.Dimension) {
			s.Typ = t.ElementType
			for _, e := range s.Indexes {
				e.Validate(c, nil)
				if !ast_types.IsInteger(e.Type()) {
					c.Error(e.GetPosition(), fmt.Sprintf("expect integer index for array, got '%s'", e.Type().String()))
				}
			}
		} else if len(s.Indexes) < len(t.Dimension) {
			array := &ast_types.TypeArray{
				ElementType: t.ElementType,
				Dimension:   []int{0},
			}
			for _, e := range s.Indexes {
				e.Validate(c, nil)
				if !ast_types.IsInteger(e.Type()) {
					c.Error(e.GetPosition(), fmt.Sprintf("expect integer index for array, got '%s'", e.Type().String()))
				}
			}
			for i := len(t.Dimension) - len(s.Indexes) - 1; i > 0; i-- {
				array.Dimension = append(array.Dimension, t.Dimension[len(t.Dimension)-1])
			}
			s.Typ = array
		} else {
			c.Error(s.GetPosition(), "mismatch array dimension")
		}
	} else {
		c.Error(s.GetPosition(), "expect array type")
	}
}
