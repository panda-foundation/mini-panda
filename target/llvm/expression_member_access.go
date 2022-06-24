package llvm

import (
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

func MemberAccessIR(c *Context, m *expression.MemberAccess) ir.Value {
	if m.Qualified == "" {
		if m.Const {
			if i, ok := m.Parent.(*expression.Identifier); ok {
				if e, ok := c.Program.FindDeclaration(i.Qualified).(*Enum); ok {
					return e.GetMember(m.Member.Name)
				}
			} else if mm, ok := m.Parent.(*expression.MemberAccess); ok {
				if e, ok := c.Program.FindDeclaration(mm.Qualified).(*Enum); ok {
					return e.GetMember(m.Member.Name)
				}
			}
		} else {
			t := m.Parent.Type()
			v := ExpressionIR(c, m.Parent)
			if p, ok := t.(*ast_types.TypePointer); ok {
				t = p.ElementType
				v = c.AutoLoad(v)
			}
			if n, ok := t.(*ast_types.TypeName); ok {
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

func MemberAccessConstIR(p *Program, m *expression.MemberAccess) constant.Constant {
	if m.Qualified == "" {
		if m.Const {
			if i, ok := m.Parent.(*expression.Identifier); ok {
				if e, ok := p.FindDeclaration(i.Qualified).(*Enum); ok {
					return e.GetMember(m.Member.Name)
				}
			} else if mm, ok := m.Parent.(*expression.MemberAccess); ok {
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
