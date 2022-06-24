package expression

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/token"
)

type Unary struct {
	ExpressionBase
	Operator   token.Token
	Expression ast.Expression
}

func (u *Unary) Validate(c ast.Context, expected ast.Type) {
	u.Expression.Validate(c, expected)
	u.Const = u.Expression.IsConstant()
	u.Typ = u.Expression.Type()
	switch u.Operator {
	case token.Plus, token.Minus:
		if !ast_types.IsNumber(u.Typ) {
			c.Error(u.GetPosition(), "expect number expression")
		}

	case token.Not:
		if !ast_types.IsBool(u.Typ) {
			c.Error(u.GetPosition(), "expect boolean expression")
		}

	case token.Complement:
		if !ast_types.IsInteger(u.Typ) {
			c.Error(u.GetPosition(), "expect integer expression")
		}

	case token.BitAnd:
		if ast_types.IsPointer(u.Typ) || ast_types.IsFunction(u.Typ) || ast_types.IsArray(u.Typ) {
			c.Error(u.GetPosition(), "pointer, function and array are not allowed to use '&' operator")
			return
		}
		u.Typ = &ast_types.TypePointer{
			ElementType: u.Typ,
		}
		switch u.Expression.(type) {
		case *Identifier:
		case *MemberAccess:
		case *Subscripting:
		default:
			c.Error(u.GetPosition(), "only identifier, member_access and subscripting are allowed with '&' operator")
		}

	case token.Mul:
		if ast_types.IsPointer(u.Typ) {
			u.Typ = u.Typ.(*ast_types.TypePointer).ElementType
		} else {
			c.Error(u.GetPosition(), "only pointer type is allowed with '*' operator")
		}
	}
}
