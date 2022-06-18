package ast

import (
	"fmt"

	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/ast/declaration"
)

func NewContext(p *Program) core.Context {
	return &Context{
		Program: p,
		objects: make(map[string]core.Type),
	}
}

type Context struct {
	Program  *Program
	Function *declaration.Function
	Returned bool

	parent  *Context
	objects map[string]core.Type
}

func (c *Context) Error(offset int, message string) {
	c.Program.Error(offset, message)
}

func (c *Context) NewContext() core.Context {
	return &Context{
		Program:  c.Program,
		Function: c.Function,

		parent:  c,
		objects: make(map[string]core.Type),
	}
}

func (c *Context) AddObject(name string, t core.Type) error {
	if _, ok := c.objects[name]; ok {
		return fmt.Errorf("redeclared variable: %s", name)
	}
	c.objects[name] = t
	return nil
}

func (c *Context) FindObject(name string) core.Type {
	if v, ok := c.objects[name]; ok {
		return v
	}
	if c.parent != nil {
		v := c.parent.FindObject(name)
		if v != nil {
			return v
		}
	}
	if c.Function != nil && c.Function.Parent != nil {
		return c.Function.Parent.MemberType(name)
	}
	return nil
}

func (c *Context) ResolveType(v core.Type) core.Type {
	switch t := v.(type) {
	case *core.TypeName:
		d := c.FindDeclarationByType(t)
		if d == nil {
			c.Program.Error(v.GetPosition(), "type not defined")
		} else {
			if f, ok := d.(*declaration.Function); ok {
				return f.Typ
			} else if _, ok := d.(*declaration.Struct); ok {
				// struct
				t.Qualified = d.QualifiedName()
			} else {
				c.Program.Error(v.GetPosition(), "type not defined")
			}
		}
		return t
	case *core.TypeArray:
		t.ElementType = c.ResolveType(t.ElementType)
		if t.Dimension[0] < 0 {
			c.Program.Error(v.GetPosition(), "invalid array index")
		}
		for i := 1; i < len(t.Dimension); i++ {
			if t.Dimension[i] < 1 {
				c.Program.Error(v.GetPosition(), "invalid array index")
			}
		}
		return t
	case *core.TypePointer:
		t.ElementType = c.ResolveType(t.ElementType)
		return t
	case *core.TypeFunction:
		t.ReturnType = c.ResolveType(t.ReturnType)
		for i := 0; i < len(t.Parameters); i++ {
			t.Parameters[i] = c.ResolveType(t.Parameters[i])
			if core.IsStruct(t.Parameters[i]) {
				c.Program.Error(t.Parameters[i].GetPosition(), "struct is not allowed as parameter, use pointer instead")
			}
			if core.IsArray(t.Parameters[i]) {
				c.Program.Error(t.Parameters[i].GetPosition(), "array is not allowed as parameter, use pointer instead")
			}
		}
		return t
	default:
		return t
	}
}

func (c *Context) FindLocalDeclaration(member string) core.Declaration {
	qualified := c.Program.Module.Namespace + "." + member
	if d, ok := c.Program.Declarations[qualified]; ok {
		return d
	}
	qualified = core.Global + "." + member
	if d, ok := c.Program.Declarations[qualified]; ok {
		return d
	}
	return nil
}

func (c *Context) FindDeclarationByName(selector, member string) core.Declaration {
	if selector == "" {
		return c.FindLocalDeclaration(member)
	}
	for _, i := range c.Program.Module.Imports {
		if i.Alias == selector {
			qualified := i.Namespace + "." + member
			return c.Program.Declarations[qualified]
		}
	}
	return nil
}

func (c *Context) FindDeclarationByType(t *core.TypeName) core.Declaration {
	d := c.FindDeclarationByName(t.Selector, t.Name)
	if _, ok := d.(*declaration.Enum); ok {
		t.IsEnum = true
	}
	t.Qualified = d.QualifiedName()
	return d
}

func (c *Context) FindQualifiedDeclaration(qualified string) core.Declaration {
	return c.Program.Declarations[qualified]
}

func (c *Context) IsNamespace(name string) bool {
	for _, i := range c.Program.Module.Imports {
		if i.Alias == name {
			return true
		}
	}
	return false
}

func (c *Context) SetFunction(f core.Declaration) {
	c.Function = f.(*declaration.Function)
}

func (c *Context) ReturnType() core.Type {
	return c.Function.ReturnType
}
