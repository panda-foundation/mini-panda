package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

func MemberAccessIR(c *Context, m *ast.MemberAccess) ir.Value {
	if m.Qualified == "" {
		if m.Const {
			if i, ok := m.Parent.(*ast.Identifier); ok {
				if e, ok := c.Program.FindDeclaration(i.Qualified).(*Enum); ok {
					return e.GetMember(m.Member.Name)
				}
			} else if mm, ok := m.Parent.(*ast.MemberAccess); ok {
				if e, ok := c.Program.FindDeclaration(mm.Qualified).(*Enum); ok {
					return e.GetMember(m.Member.Name)
				}
			}
		} else {
			t := m.Parent.Type()
			v := ExpressionIR(c, m.Parent)
			if p, ok := t.(*ast.TypePointer); ok {
				t = p.ElementType
				v = c.AutoLoad(v)
			}
			if n, ok := t.(*ast.TypeName); ok {
				if s, ok := c.Program.FindDeclaration(n.Qualified).(*Struct); ok {
					return s.GetMember(c, v, m.Member.Name)
				}
			}
		}
	} else {
		d := c.Program.FindDeclaration(m.Qualified)
		switch t := d.(type) {
		case *Variable:
			return t.Variable
		case *Function:
			return t.Function
		}
	}
	return nil
}

func MemberAccessConstIR(p *Program, m *ast.MemberAccess) ir.Constant {
	if m.Qualified == "" {
		if m.Const {
			if i, ok := m.Parent.(*ast.Identifier); ok {
				if e, ok := p.FindDeclaration(i.Qualified).(*Enum); ok {
					return e.GetMember(m.Member.Name)
				}
			} else if mm, ok := m.Parent.(*ast.MemberAccess); ok {
				if e, ok := p.FindDeclaration(mm.Qualified).(*Enum); ok {
					return e.GetMember(m.Member.Name)
				}
			}
		}
	} else {
		d := p.FindDeclaration(m.Qualified)
		switch t := d.(type) {
		case *Variable:
			if t.Const {
				return t.Variable
			}
		case *Function:
			return t.Function
		}
	}
	return nil
}
