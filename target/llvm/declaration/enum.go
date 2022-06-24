package declaration

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

type Enum struct {
	DeclarationBase

	Members map[string]ir.Constant
}

func (e *Enum) GetMember(member string) ir.Constant {
	return e.Members[member]
}

func (e *Enum) GenerateIR(p *Program, enum *ast.Enum) {
	e.Members = make(map[string]ir.Constant)
	e.Qualified = enum.Qualified
	for i, v := range enum.Members {
		name := e.Qualified + "." + v.Name.Name
		value := ir.NewInt(ir.UI8, int64(enum.Values[i]))
		d := p.Module.NewGlobalDef(name, value)
		d.Immutable = true
		e.Members[v.Name.Name] = d
	}
}
