package llvm

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/target/llvm/declaration"
	"github.com/panda-io/micro-panda/target/llvm/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
	ir_core "github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
	"github.com/panda-io/micro-panda/target/llvm/llvm"
)

func NewContext(p *Program) llvm.Context {
	return &Context{
		Program: p,
		objects: make(map[string]ir_core.Value),
	}
}

type Context struct {
	Program  *Program
	Function *declaration.Function

	Block      *ir.Block
	LeaveBlock *ir.Block
	LoopBlock  *ir.Block
	Returned   bool

	parent  llvm.Context
	objects map[string]ir_core.Value
}

func (c *Context) NewContext() llvm.Context {
	return &Context{
		Program:  c.Program,
		Function: c.Function,

		LeaveBlock: c.LeaveBlock,
		LoopBlock:  c.LoopBlock,

		parent:  c,
		objects: make(map[string]ir_core.Value),
	}
}

func (c *Context) AddObject(name string, value ir_core.Value) {
	c.objects[name] = value
}

func (c *Context) FindObject(name string) ir_core.Value {
	if v, ok := c.objects[name]; ok {
		return v
	}
	if c.parent != nil {
		v := c.parent.FindObject(name)
		if v != nil {
			return v
		}
	}
	if c.Function.Parent != nil && c.Function.Parent.HasVariable(name) {
		this := c.FindObject(ast.StructThis)
		return c.Function.Parent.GetMember(c, this, name)
	}
	return nil
}

func (c *Context) AutoLoad(value ir_core.Value) ir_core.Value {
	switch t := value.(type) {
	// global define
	case *ir.Global:
		load := instruction.NewLoad(t.ContentType, t)
		c.Block().AddInstruction(load)
		return load

	// parameter
	case *ir.Param:
		if typ, ok := t.Type().(*ir_types.PointerType); ok {
			load := instruction.NewLoad(typ.ElemType, t)
			c.Block().AddInstruction(load)
			return load
		}

	// alloca in function
	case *instruction.InstAlloca:
		load := instruction.NewLoad(t.ElemType, t)
		c.Block().AddInstruction(load)
		return load

	// struct member
	case *instruction.InstGetElementPtr:
		typ := t.Type().(*ir_types.PointerType)
		load := instruction.NewLoad(typ.ElemType, t)
		c.Block().AddInstruction(load)
		return load
	}
	return value
}
