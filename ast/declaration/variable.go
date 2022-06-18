package declaration

import (
	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/token"
)

type Variable struct {
	DeclarationBase
	Token token.Token
	Type  core.Type
	Const bool
	Value core.Expression

	Parent *Struct
}

func (v *Variable) ResolveType(c core.Context) {
	v.Type = c.ResolveType(v.Type)
}

func (v *Variable) Validate(c core.Context) {
	if v.Value != nil {
		v.Value.Validate(c, v.Type)
	}
	if v.Const {
		if v.Value == nil {
			c.Error(v.Position, "const must be initialized when declare")
		} else if !v.Value.IsConstant() {
			c.Error(v.Value.GetPosition(), "expect const expression")
		}
	}
	if v.Value != nil {
		if v.Value.Type() == nil {
			c.Error(v.Value.GetPosition(), "unknown type")
		} else if !v.Value.Type().Equal(v.Type) {
			c.Error(v.Value.GetPosition(), "init value type mismatch with define")
		}
	}
}
