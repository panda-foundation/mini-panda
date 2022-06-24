package declaration

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	ir_core "github.com/panda-io/micro-panda/ir/core"
	"github.com/panda-io/micro-panda/target/llvm"
	"github.com/panda-io/micro-panda/target/llvm/context"
)

type Struct struct {
	DeclarationBase

	Struct          *ir.StructType
	VariableIndexes map[string]int
	FunctionIndexes map[string]int
	Functions       []*Function
}

func (s *Struct) GenerateDefineIR(p llvm.Program, ss *ast.Struct) {
	s.Qualified = ss.Qualified
	var types []ir.Type
	for _, v := range ss.Variables {
		types = append(types, types.TypeIR(v.Type))
	}
	s.Struct = ir.NewStructType(types...)
	p.Module.NewTypeDef(ss.Qualified, s.Struct)
	s.VariableIndexes = map[string]int{}
	for index, v := range ss.Variables {
		s.VariableIndexes[v.Name.Name] = index
	}
	s.FunctionIndexes = map[string]int{}
	for index, f := range ss.Functions {
		s.FunctionIndexes[f.Name.Name] = index
		ff := &Function{}
		ff.Parent = s
		ff.GenerateDefineIR(p, f)
		s.Functions = append(s.Functions, ff)
	}
}

func (s *Struct) GenerateIR(p llvm.Program, ss *ast.Struct) {
	for i, f := range ss.Functions {
		s.Functions[i].GenerateIR(p, f)
	}
}

func (s *Struct) HasVariable(name string) bool {
	_, ok := s.VariableIndexes[name]
	return ok
}

func (s *Struct) GetMember(ctx context.Context, this ir_core.Value, member string) ir_core.Value {
	if index, ok := s.VariableIndexes[member]; ok {
		v := ir.NewGetElementPtr(s.Struct, this, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, int64(index)))
		ctx.Block().AddInstruction(v)
		return v
	} else if index, ok := s.FunctionIndexes[member]; ok {
		return s.Functions[index].Function
	}
	return nil
}
