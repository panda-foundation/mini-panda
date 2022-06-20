package expression

import (
	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/token"
)

type Binary struct {
	ExpressionBase
	Left     core.Expression
	Operator token.Token
	Right    core.Expression
}

func (b *Binary) Validate(c core.Context, expected core.Type) {
	b.Left.Validate(c, expected)
	b.Right.Validate(c, b.Left.Type())

	switch b.Operator {
	case token.LeftShift, token.RightShift, token.BitXor, token.BitOr, token.BitAnd:
		b.Const = b.Left.IsConstant() && b.Right.IsConstant()
		if core.IsInteger(b.Left.Type()) && core.IsInteger(b.Right.Type()) {
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
		if core.IsArray(b.Left.Type()) && !core.IsPointer(b.Left.Type()) {
			c.Error(b.Left.GetPosition(), "array type is not assignable")
		}
		if core.IsStruct(b.Left.Type()) {
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
		if !(core.IsNumber(b.Left.Type()) && core.IsNumber(b.Right.Type())) {
			c.Error(b.Left.GetPosition(), "expect number for binary expression")
		}

	case token.LeftShiftAssign, token.RightShiftAssign, token.AndAssign, token.OrAssign, token.XorAssign:
		b.Const = false
		if !(core.IsInteger(b.Left.Type()) && core.IsInteger(b.Right.Type())) {
			c.Error(b.Left.GetPosition(), "expect integer for bit operation assign")
		}
		if b.Left.IsConstant() {
			c.Error(b.Left.GetPosition(), "expect variable")
		}

	case token.Or, token.And:
		b.Const = b.Left.IsConstant() && b.Right.IsConstant()
		if core.IsBool(b.Left.Type()) && core.IsBool(b.Right.Type()) {
			b.Typ = b.Left.Type()
		} else {
			c.Error(b.Left.GetPosition(), "mismatch type for binary expression")
		}

	case token.Less, token.LessEqual, token.Greater, token.GreaterEqual, token.Equal, token.NotEqual:
		b.Const = b.Left.IsConstant() && b.Right.IsConstant()
		if core.IsNumber(b.Left.Type()) && core.IsNumber(b.Right.Type()) {
			b.Typ = core.TypeBool
		} else if core.IsPointer(b.Left.Type()) && core.IsPointer(b.Right.Type()) {
			b.Typ = core.TypeBool
		} else {
			c.Error(b.Left.GetPosition(), "expect number for compare")
		}

	case token.Plus, token.Minus, token.Mul, token.Div, token.Rem:
		b.Const = b.Left.IsConstant() && b.Right.IsConstant()
		if core.IsNumber(b.Left.Type()) && core.IsNumber(b.Right.Type()) && b.Left.Type().Equal(b.Right.Type()) {
			b.Typ = b.Left.Type()
		} else {
			c.Error(b.Left.GetPosition(), "mismatch type for binary expression")
		}
	}
}
