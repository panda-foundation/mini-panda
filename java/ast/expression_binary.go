package ast

import (
	"github.com/panda-io/micro-panda/token"
)

type Binary struct {
	ExpressionBase
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (b *Binary) Validate(c *Context, expected Type) {
	b.Left.Validate(c, expected)
	b.Right.Validate(c, b.Left.Type())
	if b.Left.Type() == nil {
		c.Program.Error(b.Left.GetPosition(), "invalid type for binary expression")
		return
	}
	if b.Right.Type() == nil {
		c.Program.Error(b.Right.GetPosition(), "invalid type for binary expression")
		return
	}
	switch b.Operator {
	case token.LeftShift, token.RightShift, token.BitXor, token.BitOr, token.BitAnd:
		b.Const = b.Left.IsConstant() && b.Right.IsConstant()
		if IsInteger(b.Left.Type()) && IsInteger(b.Right.Type()) {
			b.Typ = b.Left.Type()
		} else {
			c.Program.Error(b.Left.GetPosition(), "expect integer for bit operation")
		}
	case token.Assign:
		b.Const = false
		if !b.Left.Type().Equal(b.Right.Type()) {
			c.Program.Error(b.Left.GetPosition(), "mismatch type for binary expression")
		}
		if b.Left.IsConstant() {
			c.Program.Error(b.Left.GetPosition(), "expect variable")
		}
		if IsArray(b.Left.Type()) && !IsPointer(b.Left.Type()) {
			c.Program.Error(b.Left.GetPosition(), "array type is not assignable")
		}
		if IsStruct(b.Left.Type()) {
			c.Program.Error(b.Left.GetPosition(), "struct type is not assignable")
		}
	case token.MulAssign, token.DivAssign, token.RemAssign, token.PlusAssign, token.MinusAssign:
		b.Const = false
		if !b.Left.Type().Equal(b.Right.Type()) {
			c.Program.Error(b.Left.GetPosition(), "mismatch type for binary expression")
		}
		if b.Left.IsConstant() {
			c.Program.Error(b.Left.GetPosition(), "expect variable")
		}
		if !(IsNumber(b.Left.Type()) && IsNumber(b.Right.Type())) {
			c.Program.Error(b.Left.GetPosition(), "expect number for binary expression")
		}
	case token.LeftShiftAssign, token.RightShiftAssign, token.AndAssign, token.OrAssign, token.XorAssign:
		b.Const = false
		if !(IsInteger(b.Left.Type()) && IsInteger(b.Right.Type())) {
			c.Program.Error(b.Left.GetPosition(), "expect integer for bit operation assign")
		}
		if b.Left.IsConstant() {
			c.Program.Error(b.Left.GetPosition(), "expect variable")
		}
	case token.Or, token.And:
		b.Const = b.Left.IsConstant() && b.Right.IsConstant()
		if IsBool(b.Left.Type()) && IsBool(b.Right.Type()) {
			b.Typ = b.Left.Type()
		} else {
			c.Program.Error(b.Left.GetPosition(), "mismatch type for binary expression")
		}
	case token.Less, token.LessEqual, token.Greater, token.GreaterEqual, token.Equal, token.NotEqual:
		b.Const = b.Left.IsConstant() && b.Right.IsConstant()
		if IsNumber(b.Left.Type()) && IsNumber(b.Right.Type()) {
			b.Typ = TypeBool
		} else if IsPointer(b.Left.Type()) && IsPointer(b.Right.Type()) {
			b.Typ = TypeBool
		} else {
			c.Program.Error(b.Left.GetPosition(), "expect number for compare")
		}
	case token.Plus, token.Minus, token.Mul, token.Div, token.Rem:
		b.Const = b.Left.IsConstant() && b.Right.IsConstant()
		if IsNumber(b.Left.Type()) && IsNumber(b.Right.Type()) && b.Left.Type().Equal(b.Right.Type()) {
			b.Typ = b.Left.Type()
		} else {
			c.Program.Error(b.Left.GetPosition(), "mismatch type for binary expression")
		}
	}
}
