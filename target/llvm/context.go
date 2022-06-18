package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/ir/value"
)

func NewContext(p *Program) *Context {
	return &Context{
		Program: p,
		objects: make(map[string]value.Value),
	}
}

type Context struct {
	Program  *Program
	Function *Function

	Block      *ir.Block
	LeaveBlock *ir.Block
	LoopBlock  *ir.Block
	Returned   bool

	parent  *Context
	objects map[string]value.Value
}

func (c *Context) NewContext() *Context {
	return &Context{
		Program:  c.Program,
		Function: c.Function,

		LeaveBlock: c.LeaveBlock,
		LoopBlock:  c.LoopBlock,

		parent:  c,
		objects: make(map[string]value.Value),
	}
}

func (c *Context) AddObject(name string, value value.Value) {
	c.objects[name] = value
}

func (c *Context) FindObject(name string) value.Value {
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

func (c *Context) AutoLoad(value value.Value) value.Value {
	switch t := value.(type) {
	// global define
	case *ir.Global:
		load := ir.NewLoad(t.ContentType, t)
		c.Block.AddInstruction(load)
		return load

	// parameter
	case *ir.Param:
		if typ, ok := t.Type().(*ir.PointerType); ok {
			load := ir.NewLoad(typ.ElemType, t)
			c.Block.AddInstruction(load)
			return load
		}

	// alloca in function
	case *ir.InstAlloca:
		load := ir.NewLoad(t.ElemType, t)
		c.Block.AddInstruction(load)
		return load

	// struct member
	case *ir.InstGetElementPtr:
		typ := t.Type().(*ir.PointerType)
		load := ir.NewLoad(typ.ElemType, t)
		c.Block.AddInstruction(load)
		return load
	}
	return value
}
