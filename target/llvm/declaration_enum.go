package llvm

import (
	"github.com/panda-io/micro-panda/ast/declaration"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

type Enum struct {
	DeclarationBase

	Members map[string]constant.Constant
}

func (e *Enum) GetMember(member string) constant.Constant {
	return e.Members[member]
}

func (e *Enum) GenerateIR(p *Program, enum *declaration.Enum) {
	e.Members = make(map[string]constant.Constant)
	e.Qualified = enum.Qualified
	for i, v := range enum.Members {
		name := e.Qualified + "." + v.Name.Name
		value := constant.NewInt(ir_types.UI8, int64(enum.Values[i]))
		d := p.Program.NewGlobalDef(name, value)
		d.Immutable = true
		e.Members[v.Name.Name] = d
	}
}
