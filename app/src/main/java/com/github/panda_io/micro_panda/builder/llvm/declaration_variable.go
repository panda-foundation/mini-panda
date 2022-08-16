package llvm

import (
	"github.com/panda-io/micro-panda/ast/declaration"
	"github.com/panda-io/micro-panda/target/llvm/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	ir_core "github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type Variable struct {
	DeclarationBase

	Const    bool
	Variable *ir.Global
	Type     ir_core.Type
}

func (vv *Variable) GenerateIR(p *Program, v *declaration.Variable) {
	vv.Const = v.Const
	vv.Type = TypeIR(v.Type())
	if v.Value != nil {
		value := ExpressionConstIR(p, v.Value)
		vv.Variable = p.Program.NewGlobalDef(v.Qualified, value)
		if v.Const {
			vv.Variable.Immutable = true
		}
	} else {
		vv.Variable = p.Program.NewGlobalDef(v.Qualified, constant.NewZeroInitializer(vv.Type))
	}
}
