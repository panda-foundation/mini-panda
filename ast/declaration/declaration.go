package declaration

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/expression"
)

type Attribute struct {
	Position int
	Name     string
	Text     string
	Values   map[string]*expression.Literal
}

type DeclarationBase struct {
	ast.NodeBase
	Attributes []*Attribute
	Public     bool
	Name       *expression.Identifier
	Qualified  string
}

func (b *DeclarationBase) HasAttribute(attribute string) bool {
	for _, a := range b.Attributes {
		if a.Name == attribute {
			return true
		}
	}
	return false
}

func (b *DeclarationBase) GetAttribute(attribute string, name string) *expression.Literal {
	for _, a := range b.Attributes {
		if a.Name == attribute {
			return a.Values[name]
		}
	}
	return nil
}

func (b *DeclarationBase) QualifiedName() string {
	return b.Qualified
}
