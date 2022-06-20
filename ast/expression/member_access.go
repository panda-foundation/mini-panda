package expression

import (
	"fmt"

	"github.com/panda-io/micro-panda/ast/core"
)

type MemberAccess struct {
	ExpressionBase
	Parent core.Expression
	Member *Identifier

	Qualified   string
	IsNamespace bool
}

/*
parent expression could be: identifier$, member_access$, parentheses, subscripting, 'this', invocation
possible incomplete parent expression. it need to combine with member access
*/

func (m *MemberAccess) Validate(c core.Context, expected core.Type) {
	m.Parent.Validate(c, nil)
	if m.Parent.Type() == nil {
		if i, ok := m.Parent.(*Identifier); ok {
			if i.IsNamespace {
				qualified := fmt.Sprintf("%s.%s", i.Name, m.Member.Name)
				d := c.FindQualifiedDeclaration(qualified)
				// struct has no static members
				if d != nil && d.Kind() != core.DeclarationStruct {
					m.Typ = d.Type()
					m.Const = d.IsConstant()
					m.Qualified = d.QualifiedName()
				} else if c.IsNamespace(qualified) {
					m.IsNamespace = true
					m.Qualified = qualified
				}
			} else if d := c.FindQualifiedDeclaration(i.Qualified); d != nil {
				if d.Kind() == core.DeclarationEnum && d.(core.Enum).HasMember(m.Member.Name) {
					m.Typ = core.TypeU8
					m.Const = true
				}
			}
		} else if mm, ok := m.Parent.(*MemberAccess); ok {
			if mm.IsNamespace {
				qualified := fmt.Sprintf("%s.%s", mm.Qualified, m.Member.Name)
				d := c.FindQualifiedDeclaration(qualified)
				// struct has no static members
				if d != nil && d.Kind() != core.DeclarationStruct {
					m.Typ = d.Type()
					m.Const = d.IsConstant()
					m.Qualified = d.QualifiedName()
				} else if c.IsNamespace(qualified) {
					m.IsNamespace = true
					m.Qualified = qualified
				}
			} else if d := c.FindQualifiedDeclaration(mm.Qualified); d != nil {
				if d.Kind() == core.DeclarationEnum && d.(core.Enum).HasMember(m.Member.Name) {
					m.Typ = core.TypeU8
					m.Const = true
				}
			}
		}
	} else {
		t := m.Parent.Type()
		if p, ok := t.(*core.TypePointer); ok {
			t = p.ElementType
		}
		if n, ok := t.(*core.TypeName); ok {
			d := c.FindDeclaration(n)
			if d != nil {
				if d.Kind() == core.DeclarationStruct {
					m.Typ = d.(core.Struct).MemberType(m.Member.Name)
					m.Const = false
				}
			}
		}
	}
	// * type would be nil for enum (its member has type u8)
	if m.Typ == nil && m.Qualified == "" {
		c.Error(m.GetPosition(), fmt.Sprintf("undefined: %s", m.Member.Name))
	}
}
