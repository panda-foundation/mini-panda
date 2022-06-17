package ast

type Declaration interface {
	Node
	HasAttribute(attribute string) bool
	QualifiedName() string
	ValidateType(c *Context)
	Validate(c *Context)
}

type Attribute struct {
	Position int
	Name     string
	Text     string
	Values   map[string]*Literal
}

type DeclarationBase struct {
	NodeBase
	Attributes []*Attribute
	Public     bool
	Name       *Identifier
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

func (b *DeclarationBase) GetAttributeValue(attribute string, value string) *Literal {
	for _, a := range b.Attributes {
		if a.Name == attribute {
			return a.Values[value]
		}
	}
	return nil
}

func (b *DeclarationBase) QualifiedName() string {
	return b.Qualified
}
