package ir

import (
	"fmt"
	"strings"
)

// --- [ Bitwise expressions ] -------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprShl is an LLVM IR shl expression.
type ExprShl struct {
	// Operands.
	X, Y Constant // integer scalars or vectors

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewShl returns a new shl expression based on the given operands.
func NewExprShl(x, y Constant) *ExprShl {
	e := &ExprShl{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprShl) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprShl) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprShl) Ident() string {
	// 'shl' OverflowFlags=OverflowFlag* '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("shl")
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprShl) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprLShr is an LLVM IR lshr expression.
type ExprLShr struct {
	// Operands.
	X, Y Constant // integer scalars or vectors

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewLShr returns a new lshr expression based on the given operands.
func NewExprLShr(x, y Constant) *ExprLShr {
	e := &ExprLShr{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprLShr) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprLShr) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprLShr) Ident() string {
	// 'lshr' Exactopt '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("lshr")
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprLShr) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprAShr is an LLVM IR ashr expression.
type ExprAShr struct {
	// Operands.
	X, Y Constant // integer scalars or vectors

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewAShr returns a new ashr expression based on the given operands.
func NewExprAShr(x, y Constant) *ExprAShr {
	e := &ExprAShr{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprAShr) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprAShr) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprAShr) Ident() string {
	// 'ashr' Exactopt '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("ashr")
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprAShr) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprAnd is an LLVM IR and expression.
type ExprAnd struct {
	// Operands.
	X, Y Constant // integer scalars or vectors

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewAnd returns a new and expression based on the given operands.
func NewExprAnd(x, y Constant) *ExprAnd {
	e := &ExprAnd{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprAnd) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprAnd) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprAnd) Ident() string {
	// 'and' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("and (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprAnd) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprOr is an LLVM IR or expression.
type ExprOr struct {
	// Operands.
	X, Y Constant // integer scalars or vectors

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewOr returns a new or expression based on the given operands.
func NewExprOr(x, y Constant) *ExprOr {
	e := &ExprOr{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprOr) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprOr) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprOr) Ident() string {
	// 'or' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("or (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprOr) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprXor is an LLVM IR xor expression.
type ExprXor struct {
	// Operands.
	X, Y Constant // integer scalars or vectors

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewXor returns a new xor expression based on the given operands.
func NewExprXor(x, y Constant) *ExprXor {
	e := &ExprXor{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprXor) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprXor) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprXor) Ident() string {
	// 'xor' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("xor (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprXor) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}
