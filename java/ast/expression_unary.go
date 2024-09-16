package ast

import (
	"github.com/panda-io/micro-panda/token"
)

type Unary struct {
	ExpressionBase
	Operator   token.Token
	Expression Expression
}

func (u *Unary) Validate(c *Context, expected Type) {
	u.Expression.Validate(c, expected)
	u.Const = u.Expression.IsConstant()
	u.Typ = u.Expression.Type()
	switch u.Operator {
	case token.Plus, token.Minus:
		if !IsNumber(u.Typ) {
			c.Program.Error(u.Position, "expect number expression")
		}
	case token.Not:
		if !IsBool(u.Typ) {
			c.Program.Error(u.Position, "expect boolean expression")
		}
	case token.Complement:
		if !IsInteger(u.Typ) {
			c.Program.Error(u.Position, "expect integer expression")
		}
	case token.BitAnd:
		if IsPointer(u.Typ) || IsFunction(u.Typ) || IsArray(u.Typ) {
			c.Program.Error(u.Position, "pointer, function and array are not allowed to use '&' operator")
		}
		u.Typ = &TypePointer{
			ElementType: u.Typ,
		}
		switch u.Expression.(type) {
		case *Identifier:
		case *MemberAccess:
		case *Subscripting:
		default:
			c.Program.Error(u.Position, "only identifier, member_access and subscripting are allowed with '&' operator")
		}
	}
}
