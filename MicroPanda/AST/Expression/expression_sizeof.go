package ast

type Sizeof struct {
	ExpressionBase
	Target Type
}

func (s *Sizeof) Validate(c *Context, expected Type) {
	s.Target = ValidateType(s.Target, c.Program)
	s.Typ = TypeU32
	s.Const = true
}
