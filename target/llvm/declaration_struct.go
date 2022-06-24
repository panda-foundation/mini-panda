package llvm

import (
	"github.com/panda-io/micro-panda/ast/declaration"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
	ir_core "github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

type Struct struct {
	DeclarationBase

	Struct          *ir_types.StructType
	VariableIndexes map[string]int
	FunctionIndexes map[string]int
	Functions       []*Function
}

func (s *Struct) GenerateDefineIR(p *Program, ss *declaration.Struct) {
	s.Qualified = ss.Qualified
	var fields []ir_core.Type
	for _, v := range ss.Variables {
		fields = append(fields, TypeIR(v.Type()))
	}
	s.Struct = ir_types.NewStructType(ss.Qualified, fields...)
	p.Program.NewTypeDef(ss.Qualified, s.Struct)
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

func (s *Struct) GenerateIR(p *Program, ss *declaration.Struct) {
	for i, f := range ss.Functions {
		s.Functions[i].GenerateIR(p, f)
	}
}

func (s *Struct) HasVariable(name string) bool {
	_, ok := s.VariableIndexes[name]
	return ok
}

func (s *Struct) GetMember(ctx *Context, this ir_core.Value, member string) ir_core.Value {
	if index, ok := s.VariableIndexes[member]; ok {
		v := instruction.NewGetElementPtr(s.Struct, this, constant.NewInt(ir_types.I32, 0), constant.NewInt(ir_types.I32, int64(index)))
		ctx.Block.AddInstruction(v)
		return v
	} else if index, ok := s.FunctionIndexes[member]; ok {
		return s.Functions[index].Function
	}
	return nil
}
