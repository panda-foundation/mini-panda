package declaration

import (
	"fmt"
	"strconv"

	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/token"
)

type Enum struct {
	DeclarationBase
	Members []*Variable
	Values  []uint8
}

func (e *Enum) IsConstant() bool {
	return false
}

func (e *Enum) Kind() core.DeclarationKind {
	return core.DeclarationEnum
}

func (e *Enum) Type() core.Type {
	return nil
}

func (e *Enum) AddMember(m *Variable) error {
	if e.HasMember(m.Name.Name) {
		return fmt.Errorf("%s redeclared", m.Name.Name)
	}
	e.Members = append(e.Members, m)
	return nil
}

func (e *Enum) HasMember(name string) bool {
	for _, v := range e.Members {
		if v.Name.Name == name {
			return true
		}
	}
	return false
}

func (e *Enum) ResolveType(c core.Context) {
}

func (e *Enum) Validate(c core.Context) {
	var index int
	for _, v := range e.Members {
		if index >= 256 {
			c.Error(v.GetPosition(), "enum value shoud be less than 256")
		}
		if v.Value == nil {
			e.Values = append(e.Values, uint8(index))
			index++
		} else {
			if literal, ok := v.Value.(*expression.Literal); ok && literal.Token == token.INT {
				if i, _ := strconv.Atoi(literal.Value); i >= index {
					index = i
					e.Values = append(e.Values, uint8(index))
					index++
				} else {
					c.Error(v.GetPosition(), fmt.Sprintf("enum value here should be greater than %d.", i-1))
				}
			} else {
				c.Error(v.GetPosition(), "enum value must be const integer.")
			}
		}
	}
}
