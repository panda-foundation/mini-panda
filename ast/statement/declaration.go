package statement

import (
	"fmt"

	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/ast/expression"
)

type DeclarationStatement struct {
	StatementBase
	Name  *expression.Identifier
	Type  core.Type
	Value core.Expression
}

func (d *DeclarationStatement) Validate(c core.Context) {
	d.Type = c.ResolveType(d.Type)
	if d.Value != nil {
		d.Value.Validate(c, d.Type)
		if d.Value.Type() != nil && d.Type != nil && !d.Value.Type().Equal(d.Type) {
			c.Error(d.Value.GetPosition(), fmt.Sprintf("init value type mismatch with define, expect '%s', got '%s'", d.Type.String(), d.Value.Type().String()))
		}
	}
	if d.Type != nil {
		err := c.AddObject(d.Name.Name, d.Type)
		if err != nil {
			c.Error(d.GetPosition(), err.Error())
		}
	}
}
