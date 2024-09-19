package ast

type Subscripting struct {
	ExpressionBase
	Parent  Expression
	Indexes []Expression
}

func (s *Subscripting) Validate(c *Context, expected Type) {
	s.Const = false
	s.Parent.Validate(c, nil)
	if t, ok := s.Parent.Type().(*TypeArray); ok {
		if len(s.Indexes) == len(t.Dimension) {
			s.Typ = t.ElementType
			for _, e := range s.Indexes {
				e.Validate(c, nil)
				if !IsInteger(e.Type()) {
					c.Program.Error(e.GetPosition(), "expect integer index for array")
				}
			}
		} else if len(s.Indexes) < len(t.Dimension) {
			array := &TypeArray{
				ElementType: t.ElementType,
				Dimension:   []int{0},
			}
			for _, e := range s.Indexes {
				e.Validate(c, nil)
				if !IsInteger(e.Type()) {
					c.Program.Error(e.GetPosition(), "expect integer index for array")
				}
			}
			for i := len(t.Dimension) - len(s.Indexes) - 1; i > 0; i-- {
				array.Dimension = append(array.Dimension, t.Dimension[len(t.Dimension)-1])
			}
			s.Typ = array
		} else {
			c.Program.Error(s.Position, "mismatch array dimension")
		}
	} else {
		c.Program.Error(s.Position, "expect array type")
	}
}
