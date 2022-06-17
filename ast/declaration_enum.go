package ast

import (
	"fmt"
	"strconv"

	"github.com/panda-io/micro-panda/token"
)

type Enum struct {
	DeclarationBase
	Members []*Variable
	Values  []uint8
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

func (e *Enum) ValidateType(c *Context) {
}

func (e *Enum) Validate(c *Context) {
	var index int
	for _, v := range e.Members {
		if index >= 256 {
			c.Program.Error(v.Position, "enum value shoud be less than 256")
		}
		if v.Value == nil {
			e.Values = append(e.Values, uint8(index))
			index++
		} else {
			if literal, ok := v.Value.(*Literal); ok {
				if literal.Token == token.INT {
					if i, _ := strconv.Atoi(literal.Value); i >= index {
						index = i
						e.Values = append(e.Values, uint8(index))
						index++
					} else {
						c.Program.Error(v.Position, fmt.Sprintf("enum value here should be greater than %d.", i-1))
					}
				} else {
					c.Program.Error(v.Position, "enum value must be integer.")
				}
			} else {
				c.Program.Error(v.Position, "enum value must be const integer.")
			}
		}
	}
}
