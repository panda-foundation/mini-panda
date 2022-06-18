package expression

import (
	"fmt"

	"github.com/panda-io/micro-panda/ast/core"
)

type MemberAccess struct {
	ExpressionBase
	Parent core.Expression
	Member *Identifier

	Qualified string
}

/*
1 level member access
	* package.variable $
	* package.function $
	* package.enum $
	* enum.member $
	* struct_instance.member

2 level member access
	* package.enum.member $
	* package.struct_instance.member
	* struct_instance.member.member

> 2 level member access
	* package.struct_instance.member ...
	* struct_instance.member.member ...

parent expression could be: identifier$, member_access$, parentheses, subscripting, 'this', invocation
$ possible incomplete parent expression. it need to combine with member access
*/

func (m *MemberAccess) Validate(c core.Context, expected core.Type) {
	m.Parent.Validate(c, nil)
	if m.Parent.Type() == nil {
		if i, ok := m.Parent.(*Identifier); ok {
			if i.IsNamespace {
				d := c.FindDeclarationByName(i.Name, m.Member.Name)
				if d != nil && !d.IsStruct() {
					m.Typ = d.Type()
					m.Const = d.IsConstant()
					m.Qualified = d.QualifiedName()
				}
			} else if e, ok := c.Program.FindQualified(i.Qualified).(*Enum); ok {
				if e.HasMember(m.Member.Name) {
					m.Typ = TypeU8
					m.Const = true
				}
			}
		} else if mm, ok := m.Parent.(*MemberAccess); ok {
			if e, ok := c.Program.FindQualified(mm.Qualified).(*Enum); ok {
				if e.HasMember(m.Member.Name) {
					m.Typ = TypeU8
					m.Const = true
				}
			}
		}
	} else {
		t := m.Parent.Type()
		if p, ok := t.(*TypePointer); ok {
			t = p.ElementType
		}
		if n, ok := t.(*TypeName); ok {
			d := c.Program.FindType(n)
			if d != nil {
				if s, ok := d.(*Struct); ok {
					m.Typ = s.MemberType(m.Member.Name)
					m.Const = false
				}
			}
		}
	}
	// * type would be nil for enum (its member has type u8)
	if m.Typ == nil && m.Qualified == "" {
		c.Error(m.Position, fmt.Sprintf("undefined: %s", m.Member.Name))
	}
}
