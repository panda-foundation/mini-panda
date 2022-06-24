package expression

import (
	"fmt"

	"github.com/panda-io/micro-panda/ast/ast"
)

type Identifier struct {
	ExpressionBase
	Name string

	Qualified   string
	IsNamespace bool
}

func (i *Identifier) Validate(c ast.Context, expected ast.Type) {
	t := c.FindObject(i.Name)
	if t == nil {
		d := c.FindLocalDeclaration(i.Name)
		if d == nil {
			i.IsNamespace = c.IsNamespace(i.Name)
		} else if d.Kind() != ast.DeclarationStruct {
			i.Const = d.IsConstant()
			i.Typ = d.Type()
			i.Qualified = d.QualifiedName()
		}
	} else {
		i.Const = false
		i.Typ = t
	}
	// * type would be nil for enum (its member has type u8)
	// * type is nil when identifier is namespacee
	if i.Typ == nil && i.Qualified == "" && !i.IsNamespace {
		c.Error(i.GetPosition(), fmt.Sprintf("undefined %s", i.Name))
	}
}
