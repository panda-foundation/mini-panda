package expression

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/token"
)

type Binary struct {
	ExpressionBase
	Left     ast.Expression
	Operator token.Token
	Right    ast.Expression
}

func (b *Binary) Validate(c ast.Context, expected ast.Type) {
	b.Left.Validate(c, expected)
	b.Right.Validate(c, b.Left.Type())

	switch b.Operator {
	case token.LeftShift, token.RightShift, token.BitXor, token.BitOr, token.BitAnd:
		b.Const = b.Left.IsConstant() && b.Right.IsConstant()
		if ast_types.IsInteger(b.Left.Type()) && ast_types.IsInteger(b.Right.Type()) {
			b.Typ = b.Left.Type()
		} else {
			c.Error(b.Left.GetPosition(), "expect integer for bit operation")
		}

	case token.Assign:
		b.Const = false
		if !b.Left.Type().Equal(b.Right.Type()) {
			c.Error(b.Left.GetPosition(), "mismatch type for binary expression")
		}
		if b.Left.IsConstant() {
			c.Error(b.Left.GetPosition(), "expect variable")
		}
		if ast_types.IsArray(b.Left.Type()) && !ast_types.IsPointer(b.Left.Type()) {
			c.Error(b.Left.GetPosition(), "array type is not assignable")
		}
		if ast_types.IsStruct(b.Left.Type()) {
			//TO-DO copy struct?
			c.Error(b.Left.GetPosition(), "struct type is not assignable")
		}

	case token.MulAssign, token.DivAssign, token.RemAssign, token.PlusAssign, token.MinusAssign:
		b.Const = false
		if !b.Left.Type().Equal(b.Right.Type()) {
			c.Error(b.Left.GetPosition(), "mismatch type for binary expression")
		}
		if b.Left.IsConstant() {
			c.Error(b.Left.GetPosition(), "expect variable")
		}
		if !(ast_types.IsNumber(b.Left.Type()) && ast_types.IsNumber(b.Right.Type())) {
			c.Error(b.Left.GetPosition(), "expect number for binary expression")
		}

	case token.LeftShiftAssign, token.RightShiftAssign, token.AndAssign, token.OrAssign, token.XorAssign:
		b.Const = false
		if !(ast_types.IsInteger(b.Left.Type()) && ast_types.IsInteger(b.Right.Type())) {
			c.Error(b.Left.GetPosition(), "expect integer for bit operation assign")
		}
		if b.Left.IsConstant() {
			c.Error(b.Left.GetPosition(), "expect variable")
		}

	case token.Or, token.And:
		b.Const = b.Left.IsConstant() && b.Right.IsConstant()
		if ast_types.IsBool(b.Left.Type()) && ast_types.IsBool(b.Right.Type()) {
			b.Typ = b.Left.Type()
		} else {
			c.Error(b.Left.GetPosition(), "mismatch type for binary expression")
		}

	case token.Less, token.LessEqual, token.Greater, token.GreaterEqual, token.Equal, token.NotEqual:
		b.Const = b.Left.IsConstant() && b.Right.IsConstant()
		if ast_types.IsNumber(b.Left.Type()) && ast_types.IsNumber(b.Right.Type()) {
			b.Typ = ast_types.TypeBool
		} else if ast_types.IsPointer(b.Left.Type()) && ast_types.IsPointer(b.Right.Type()) {
			b.Typ = ast_types.TypeBool
		} else {
			c.Error(b.Left.GetPosition(), "expect number for compare")
		}

	case token.Plus, token.Minus, token.Mul, token.Div, token.Rem:
		b.Const = b.Left.IsConstant() && b.Right.IsConstant()
		if ast_types.IsNumber(b.Left.Type()) && ast_types.IsNumber(b.Right.Type()) && b.Left.Type().Equal(b.Right.Type()) {
			b.Typ = b.Left.Type()
		} else {
			c.Error(b.Left.GetPosition(), "mismatch type for binary expression")
		}
	}
}
