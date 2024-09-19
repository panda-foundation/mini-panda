package ast

import (
	"github.com/panda-io/micro-panda/token"
)

type Variable struct {
	DeclarationBase
	Token token.Token
	Type  Type
	Const bool
	Value Expression

	Parent *Struct
}

func (v *Variable) ValidateType(c *Context) {
	v.Type = ValidateType(v.Type, c.Program)
}

func (v *Variable) Validate(c *Context) {
	if v.Value != nil {
		v.Value.Validate(c, v.Type)
	}
	if v.Const {
		if v.Value == nil {
			c.Program.Error(v.Position, "const must be initialized when declare")
		} else if !v.Value.IsConstant() {
			c.Program.Error(v.Value.GetPosition(), "expect const expression")
		}
	}
	if v.Value != nil {
		if v.Value.Type() == nil {
			c.Program.Error(v.Value.GetPosition(), "unknown type")
		} else if !v.Value.Type().Equal(v.Type) {
			c.Program.Error(v.Value.GetPosition(), "init value type mismatch with define")
		}
	}
}
