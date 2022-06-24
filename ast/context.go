package ast

import (
	"fmt"

	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/ast/declaration"
)

func NewContext(p *Program) ast.Context {
	return &Context{
		Program: p,
		objects: make(map[string]ast.Type),
	}
}

type Context struct {
	Program  *Program
	Function *declaration.Function

	parent  *Context
	objects map[string]ast.Type
}

func (c *Context) Error(offset int, message string) {
	c.Program.Error(offset, message)
}

func (c *Context) NewContext() ast.Context {
	return &Context{
		Program:  c.Program,
		Function: c.Function,

		parent:  c,
		objects: make(map[string]ast.Type),
	}
}

func (c *Context) AddObject(name string, t ast.Type) error {
	if _, ok := c.objects[name]; ok {
		return fmt.Errorf("redeclared variable: %s", name)
	}
	c.objects[name] = t
	return nil
}

func (c *Context) FindObject(name string) ast.Type {
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

func (c *Context) ResolveType(v ast.Type) ast.Type {
	switch t := v.(type) {
	case *ast_types.TypeName:
		d := c.FindDeclaration(t)
		if d == nil {
			c.Program.Error(v.GetPosition(), "type not defined")
		} else {
			t.Qualified = d.QualifiedName()
			switch i := d.(type) {
			case *declaration.Function:
				return i.Typ

			case *declaration.Struct:

			case *declaration.Enum:
				t.IsEnum = true

			default:
				c.Program.Error(v.GetPosition(), "type not defined")
			}
		}
		return t

	case *ast_types.TypeArray:
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

	case *ast_types.TypePointer:
		t.ElementType = c.ResolveType(t.ElementType)
		return t

	case *ast_types.TypeFunction:
		t.ReturnType = c.ResolveType(t.ReturnType)
		for i := 0; i < len(t.Parameters); i++ {
			t.Parameters[i] = c.ResolveType(t.Parameters[i])
			if ast_types.IsStruct(t.Parameters[i]) {
				c.Program.Error(t.Parameters[i].GetPosition(), "struct is not allowed as parameter, use pointer instead")
			}
			if ast_types.IsArray(t.Parameters[i]) {
				c.Program.Error(t.Parameters[i].GetPosition(), "array is not allowed as parameter, use pointer instead")
			}
		}
		return t

	default:
		return t
	}
}

func (c *Context) FindDeclaration(t ast.Type) ast.Declaration {
	typeName := t.(*ast_types.TypeName)
	if typeName.Qualified == "" {
		return c.FindLocalDeclaration(typeName.Name)
	}
	d := c.FindQualifiedDeclaration(typeName.Qualified)
	if _, ok := d.(*declaration.Enum); ok {
		typeName.IsEnum = true
	}
	typeName.Qualified = d.QualifiedName()
	return d
}

func (c *Context) FindLocalDeclaration(name string) ast.Declaration {
	qualified := fmt.Sprintf("%s.%s", c.Program.Module.Namespace, name)
	if d, ok := c.Program.Declarations[qualified]; ok {
		return d
	}
	qualified = fmt.Sprintf("%s.%s", ast.Global, name)
	if d, ok := c.Program.Declarations[qualified]; ok {
		return d
	}
	for _, i := range c.Program.Module.Imports {
		qualified = fmt.Sprintf("%s.%s", i.Namespace, name)
		if d, ok := c.Program.Declarations[qualified]; ok {
			return d
		}
	}
	return nil
}

func (c *Context) FindQualifiedDeclaration(qualified string) ast.Declaration {
	return c.Program.Declarations[qualified]
}

func (c *Context) IsNamespace(name string) bool {
	return c.Program.IsNamespace(name)
}

func (c *Context) SetFunction(f ast.Function) {
	c.Function = f.(*declaration.Function)
}

func (c *Context) GetFunction() ast.Function {
	return c.Function
}
