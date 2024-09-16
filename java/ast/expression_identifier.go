package ast

import "fmt"

type Identifier struct {
	ExpressionBase
	Name string

	Qualified   string
	IsNamespace bool
}

func (i *Identifier) Validate(c *Context, expected Type) {
	t := c.FindObject(i.Name)
	if t == nil {
		_, d := c.Program.FindMember(i.Name)
		if v, ok := d.(*Variable); ok {
			i.Const = v.Const
			i.Typ = v.Type
			i.Qualified = d.QualifiedName()
		}
		if f, ok := d.(*Function); ok {
			i.Const = true
			i.Typ = f.Type
			i.Qualified = d.QualifiedName()
		}
		if _, ok := d.(*Enum); ok {
			i.Qualified = d.QualifiedName()
		}
		if d == nil {
			i.IsNamespace = c.Program.IsNamespace(i.Name)
		}
	} else {
		i.Const = false
		i.Typ = t
	}
	// * type would be nil for enum (its member has type u8)
	// * type is nil when identifier is namespacee
	if i.Typ == nil && i.Qualified == "" && !i.IsNamespace {
		c.Program.Error(i.Position, fmt.Sprintf("undefined %s", i.Name))
	}
}
