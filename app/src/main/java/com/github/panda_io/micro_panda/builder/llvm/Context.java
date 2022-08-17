package com.github.panda_io.micro_panda.builder.llvm;

import com.github.panda_io.micro_panda.builder.llvm.ir.Value;

public class Context {
    public Context() {

    }

    public Context createContext() {
        return null;
    }

    public void addObject(String name, Value value) {
    }

    public Value findObject(String name) {
        return null;
    }
    
    public Value autoLoad(Value value) {
        return null;
    }
    /*
//TO-DO
func NewContext(p *Program) *Context {
	return &Context{
		Program: p,
		objects: make(map[string]ir_core.Value),
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
	objects map[string]ir_core.Value
}

func (c *Context) NewContext() *Context {
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
		c.Block.AddInstruction(load)
		return load

	// parameter
	case *ir.Param:
		if typ, ok := t.Type().(*ir_types.PointerType); ok {
			load := instruction.NewLoad(typ.ElemType, t)
			c.Block.AddInstruction(load)
			return load
		}

	// alloca in function
	case *instruction.InstAlloca:
		load := instruction.NewLoad(t.ElemType, t)
		c.Block.AddInstruction(load)
		return load

	// struct member
	case *instruction.InstGetElementPtr:
		typ := t.Type().(*ir_types.PointerType)
		load := instruction.NewLoad(typ.ElemType, t)
		c.Block.AddInstruction(load)
		return load
	}
	return value
}
*/
}
