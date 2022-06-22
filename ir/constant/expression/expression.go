package expression

// === [ Expressions ] =========================================================

// Expression is an LLVM IR constant expression.
//
// An Expression has one of the following underlying types.
//
// Unary expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprFNeg   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFNeg
//
// Binary expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprAdd    // https://godoc.org/github.com/llir/llvm/ir/constant#ExprAdd
//    *constant.ExprFAdd   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFAdd
//    *constant.ExprSub    // https://godoc.org/github.com/llir/llvm/ir/constant#ExprSub
//    *constant.ExprFSub   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFSub
//    *constant.ExprMul    // https://godoc.org/github.com/llir/llvm/ir/constant#ExprMul
//    *constant.ExprFMul   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFMul
//    *constant.ExprUDiv   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprUDiv
//    *constant.ExprSDiv   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprSDiv
//    *constant.ExprFDiv   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFDiv
//    *constant.ExprURem   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprURem
//    *constant.ExprSRem   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprSRem
//    *constant.ExprFRem   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFRem
//
// Bitwise expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprShl    // https://godoc.org/github.com/llir/llvm/ir/constant#ExprShl
//    *constant.ExprLShr   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprLShr
//    *constant.ExprAShr   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprAShr
//    *constant.ExprAnd    // https://godoc.org/github.com/llir/llvm/ir/constant#ExprAnd
//    *constant.ExprOr     // https://godoc.org/github.com/llir/llvm/ir/constant#ExprOr
//    *constant.ExprXor    // https://godoc.org/github.com/llir/llvm/ir/constant#ExprXor
//
// Memory expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprGetElementPtr   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprGetElementPtr
//
// Conversion expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprTrunc           // https://godoc.org/github.com/llir/llvm/ir/constant#ExprTrunc
//    *constant.ExprSExt            // https://godoc.org/github.com/llir/llvm/ir/constant#ExprSExt
//    *constant.ExprFPTrunc         // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPTrunc
//    *constant.ExprFPExt           // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPExt
//    *constant.ExprFPToUI          // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPToUI
//    *constant.ExprFPToSI          // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPToSI
//    *constant.ExprUIToFP          // https://godoc.org/github.com/llir/llvm/ir/constant#ExprUIToFP
//    *constant.ExprSIToFP          // https://godoc.org/github.com/llir/llvm/ir/constant#ExprSIToFP
//    *constant.ExprBitCast         // https://godoc.org/github.com/llir/llvm/ir/constant#ExprBitCast
//
// Other expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprICmp     // https://godoc.org/github.com/llir/llvm/ir/constant#ExprICmp
//    *constant.ExprFCmp     // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFCmp

type Expression interface {
	Constant
	// Simplify returns an equivalent (and potentially simplified) constant to
	// the constant expression.
	Simplify() Constant
}

// --- [ Unary expressions ] ---------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFNeg) isConstant() {}

// --- [ Binary expressions ] --------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAdd) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFAdd) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSub) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFSub) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprMul) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFMul) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprUDiv) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSDiv) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFDiv) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprURem) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSRem) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFRem) isConstant() {}

// --- [ Bitwise expressions ] -------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprShl) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprLShr) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAShr) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAnd) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprOr) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprXor) isConstant() {}

// --- [ Memory expressions ] --------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprGetElementPtr) isConstant() {}

// --- [ Conversion expressions ] ----------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprTrunc) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSExt) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPTrunc) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPExt) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPToUI) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPToSI) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprUIToFP) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSIToFP) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprBitCast) isConstant() {}

// --- [ Other expressions ] ---------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprICmp) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFCmp) isConstant() {}
