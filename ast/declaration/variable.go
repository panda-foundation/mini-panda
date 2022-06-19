package declaration

import (
	"github.com/panda-io/micro-panda/ast/core"
)

type Variable struct {
	DeclarationBase
	Typ   core.Type
	Const bool
	Value core.Expression

	Parent *Struct
}

func (v *Variable) IsConstant() bool {
	return v.Const
}

func (v *Variable) Kind() core.DeclarationKind {
	return core.DeclarationVariable
}

func (v *Variable) Type() core.Type {
	return v.Typ
}

func (v *Variable) ResolveType(c core.Context) {
	v.Typ = c.ResolveType(v.Typ)
}

func (v *Variable) Validate(c core.Context) {
	if v.Value != nil {
		v.Value.Validate(c, v.Typ)
	}
	if v.Const {
		if v.Value == nil {
			c.Error(v.GetPosition(), "const must be initialized when declare")
		} else if !v.Value.IsConstant() {
			c.Error(v.Value.GetPosition(), "expect const expression")
		}
	}
	if v.Value != nil {
		if v.Value.Type() == nil {
			c.Error(v.Value.GetPosition(), "unknown type")
		} else if !v.Value.Type().Equal(v.Typ) {
			c.Error(v.Value.GetPosition(), "init value type mismatch with define")
		}
	}
}
