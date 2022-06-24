package declaration

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

type Variable struct {
	DeclarationBase

	Const    bool
	Variable *ir.Global
	Type     ir.Type
}

func (vv *Variable) GenerateIR(p *Program, v *ast.Variable) {
	vv.Const = v.Const
	vv.Type = TypeIR(v.Type)
	if v.Value != nil {
		value := ExpressionConstIR(p, v.Value)
		vv.Variable = p.Module.NewGlobalDef(v.Qualified, value)
		if v.Const {
			vv.Variable.Immutable = true
		}
	} else {
		vv.Variable = p.Module.NewGlobalDef(v.Qualified, ir.NewZeroInitializer(vv.Type))
	}
}
