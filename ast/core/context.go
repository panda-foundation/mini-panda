package core

type Context interface {
	NewContext() Context
	Error(offset int, message string)

	AddObject(name string, t Type) error
	FindObject(name string) Type
	ResolveType(v Type) Type

	//FindType(selector string, member string) Type
	FindDeclarationByName(selector, member string) Declaration //find FindSelector in program
	FindDeclarationByType(t *TypeName) Declaration             //find type in program
	FindLocalDeclaration(name string) Declaration              //find member in program
	IsNamespace(namespace string) bool                         // in program

	SetFunction(f Declaration)
	ReturnType() Type
}

/*
	_, d := c.Program.FindSelector(i.Name, m.Member.Name)
	if d != nil {
		if v, ok := d.(*Variable); ok {
			m.Typ = v.Type
			m.Const = v.Const
			m.Qualified = d.QualifiedName()
		}
		if f, ok := d.(*Function); ok {
			m.Typ = f.Type
			m.Const = true
			m.Qualified = d.QualifiedName()
		}
		if _, ok := d.(*Enum); ok {
			m.Qualified = d.QualifiedName()
		}
	}
*/

/*
import (
	"fmt"
)

func NewContext(p *Program) *Context {
	return &Context{
		Program: p,
		objects: make(map[string]Type),
	}
}

type Context struct {
	Program  *Program
	Function *Function
	Returned bool

	parent  *Context
	objects map[string]Type
}

func (c *Context) NewContext() *Context {
	return &Context{
		Program:  c.Program,
		Function: c.Function,

		parent:  c,
		objects: make(map[string]Type),
	}
}

func (c *Context) AddObject(name string, t Type) error {
	if _, ok := c.objects[name]; ok {
		return fmt.Errorf("redeclared variable: %s", name)
	}
	c.objects[name] = t
	return nil
}

func (c *Context) FindObject(name string) Type {
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

func (c *Context) FindSelector(selector string, member string) Type {
	parent := c.FindObject(selector)
	if parent == nil {
		_, d := c.Program.FindSelector(selector, member)
		if d == nil {
			// could be an enum
			_, e := c.Program.FindMember(selector)
			if ee, ok := e.(*Enum); ok {
				if ee.HasMember(member) {
					return TypeU8
				}
			}
			return nil
		}
		switch t := d.(type) {
		case *Variable:
			return t.Type

		case *Function:
			return t.Type

		case *Enum:
			// enum itself is not a type, its member has u8 type
			return nil
		}
	} else if t, ok := parent.(*TypeName); ok {
		d := c.Program.FindType(t)
		if s, ok := d.(*Struct); ok {
			return s.MemberType(member)
		}
	}
	return nil
}
*/
