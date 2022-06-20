package ir

import (
	"fmt"
	"math/big"
)

// --- [ Conversion expressions ] ----------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprTrunc is an LLVM IR trunc expression.
type ExprTrunc struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To Type
}

// NewTrunc returns a new trunc expression based on the given source value and
// target type.
func NewExprTrunc(from Constant, to Type) *ExprTrunc {
	e := &ExprTrunc{From: from, To: to}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprTrunc) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprTrunc) Type() Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprTrunc) Ident() string {
	// 'trunc' '(' From=TypeConst 'to' To=Type ')'
	return fmt.Sprintf("trunc (%s to %s)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprTrunc) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSExt is an LLVM IR sext expression.
type ExprSExt struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To Type
}

// NewSExt returns a new sext expression based on the given source value and
// target type.
func NewExprSExt(from Constant, to Type) *ExprSExt {
	e := &ExprSExt{From: from, To: to}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSExt) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSExt) Type() Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSExt) Ident() string {
	// 'sext' '(' From=TypeConst 'to' To=Type ')'
	return fmt.Sprintf("sext (%s to %s)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprSExt) Simplify() Constant {
	// TODO: validate if we use the right approach here for sign extension. Since
	// big.Int already contains an explicit sign for the arbitrary precision
	// integer, perhaps we should not change the underlying value but instead
	// just return a new LLVM IR integer constant of the To LLVM IR type.
	from := e.From
	if fromExpr, ok := from.(Expression); ok {
		from = fromExpr.Simplify()
	}
	switch expr := from.(type) {
	case *Int:
		fromType := expr.Typ
		toType := e.To.(*IntType)
		fromX := expr.X
		// Copy e.X big.Int.
		toX := new(big.Int).Set(fromX)
		// Create simplified return constant.
		c := &Int{
			Typ: toType,
			X:   toX,
		}
		// Check if from is signed.
		fromBits := int(fromType.BitSize)
		toBits := int(toType.BitSize)
		if fromX.Bit(fromBits-1) == 1 {
			// Sign extend.
			for i := fromBits; i < toBits; i++ {
				toX.SetBit(c.X, i, 1)
			}
		}
		return c
	default:
		panic(fmt.Errorf("support for sext constant expression from type %T not yet implemented", e.From))
	}
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFPTrunc is an LLVM IR fptrunc expression.
type ExprFPTrunc struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To Type
}

// NewFPTrunc returns a new fptrunc expression based on the given source value
// and target type.
func NewExprFPTrunc(from Constant, to Type) *ExprFPTrunc {
	e := &ExprFPTrunc{From: from, To: to}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFPTrunc) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFPTrunc) Type() Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFPTrunc) Ident() string {
	// 'fptrunc' '(' From=TypeConst 'to' To=Type ')'
	return fmt.Sprintf("fptrunc (%s to %s)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFPTrunc) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFPExt is an LLVM IR fpext expression.
type ExprFPExt struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To Type
}

// NewFPExt returns a new fpext expression based on the given source value and
// target type.
func NewExprFPExt(from Constant, to Type) *ExprFPExt {
	e := &ExprFPExt{From: from, To: to}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFPExt) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFPExt) Type() Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFPExt) Ident() string {
	// 'fpext' '(' From=TypeConst 'to' To=Type ')'
	return fmt.Sprintf("fpext (%s to %s)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFPExt) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFPToUI is an LLVM IR fptoui expression.
type ExprFPToUI struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To Type
}

// NewFPToUI returns a new fptoui expression based on the given source value and
// target type.
func NewExprFPToUI(from Constant, to Type) *ExprFPToUI {
	e := &ExprFPToUI{From: from, To: to}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFPToUI) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFPToUI) Type() Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFPToUI) Ident() string {
	// 'fptoui' '(' From=TypeConst 'to' To=Type ')'
	return fmt.Sprintf("fptoui (%s to %s)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFPToUI) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFPToSI is an LLVM IR fptosi expression.
type ExprFPToSI struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To Type
}

// NewFPToSI returns a new fptosi expression based on the given source value and
// target type.
func NewExprFPToSI(from Constant, to Type) *ExprFPToSI {
	e := &ExprFPToSI{From: from, To: to}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFPToSI) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFPToSI) Type() Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFPToSI) Ident() string {
	// 'fptosi' '(' From=TypeConst 'to' To=Type ')'
	return fmt.Sprintf("fptosi (%s to %s)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFPToSI) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprUIToFP is an LLVM IR uitofp expression.
type ExprUIToFP struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To Type
}

// NewUIToFP returns a new uitofp expression based on the given source value and
// target type.
func NewExprUIToFP(from Constant, to Type) *ExprUIToFP {
	e := &ExprUIToFP{From: from, To: to}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprUIToFP) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprUIToFP) Type() Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprUIToFP) Ident() string {
	// 'uitofp' '(' From=TypeConst 'to' To=Type ')'
	return fmt.Sprintf("uitofp (%s to %s)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprUIToFP) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSIToFP is an LLVM IR sitofp expression.
type ExprSIToFP struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To Type
}

// NewSIToFP returns a new sitofp expression based on the given source value and
// target type.
func NewExprSIToFP(from Constant, to Type) *ExprSIToFP {
	e := &ExprSIToFP{From: from, To: to}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSIToFP) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSIToFP) Type() Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSIToFP) Ident() string {
	// 'sitofp' '(' From=TypeConst 'to' To=Type ')'
	return fmt.Sprintf("sitofp (%s to %s)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprSIToFP) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprBitCast is an LLVM IR bitcast expression.
type ExprBitCast struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To Type
}

// NewBitCast returns a new bitcast expression based on the given source value
// and target type.
func NewExprBitCast(from Constant, to Type) *ExprBitCast {
	e := &ExprBitCast{From: from, To: to}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprBitCast) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprBitCast) Type() Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprBitCast) Ident() string {
	// 'bitcast' '(' From=TypeConst 'to' To=Type ')'
	return fmt.Sprintf("bitcast (%s to %s)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprBitCast) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}
