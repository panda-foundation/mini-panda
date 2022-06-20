package instruction

import "io"

// === [ Instructions ] ========================================================

// Instruction is an LLVM IR instruction. All instructions (except store and
// fence) implement the value.Named interface and may thus be used directly as
// values.
//
// An Instruction has one of the following underlying types.
//
// Unary instructions
//
// https://llvm.org/docs/LangRef.html#unary-operations
//
//    *ir.InstFNeg   // https://godoc.org/github.com/llir/llvm/ir#InstFNeg
//
// Binary instructions
//
// https://llvm.org/docs/LangRef.html#binary-operations
//
//    *ir.InstAdd    // https://godoc.org/github.com/llir/llvm/ir#InstAdd
//    *ir.InstFAdd   // https://godoc.org/github.com/llir/llvm/ir#InstFAdd
//    *ir.InstSub    // https://godoc.org/github.com/llir/llvm/ir#InstSub
//    *ir.InstFSub   // https://godoc.org/github.com/llir/llvm/ir#InstFSub
//    *ir.InstMul    // https://godoc.org/github.com/llir/llvm/ir#InstMul
//    *ir.InstFMul   // https://godoc.org/github.com/llir/llvm/ir#InstFMul
//    *ir.InstUDiv   // https://godoc.org/github.com/llir/llvm/ir#InstUDiv
//    *ir.InstSDiv   // https://godoc.org/github.com/llir/llvm/ir#InstSDiv
//    *ir.InstFDiv   // https://godoc.org/github.com/llir/llvm/ir#InstFDiv
//    *ir.InstURem   // https://godoc.org/github.com/llir/llvm/ir#InstURem
//    *ir.InstSRem   // https://godoc.org/github.com/llir/llvm/ir#InstSRem
//    *ir.InstFRem   // https://godoc.org/github.com/llir/llvm/ir#InstFRem
//
// Bitwise instructions
//
// https://llvm.org/docs/LangRef.html#bitwise-binary-operations
//
//    *ir.InstShl    // https://godoc.org/github.com/llir/llvm/ir#InstShl
//    *ir.InstLShr   // https://godoc.org/github.com/llir/llvm/ir#InstLShr
//    *ir.InstAShr   // https://godoc.org/github.com/llir/llvm/ir#InstAShr
//    *ir.InstAnd    // https://godoc.org/github.com/llir/llvm/ir#InstAnd
//    *ir.InstOr     // https://godoc.org/github.com/llir/llvm/ir#InstOr
//    *ir.InstXor    // https://godoc.org/github.com/llir/llvm/ir#InstXor
//
// Memory instructions
//
// https://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations
//
//    *ir.InstAlloca          // https://godoc.org/github.com/llir/llvm/ir#InstAlloca
//    *ir.InstLoad            // https://godoc.org/github.com/llir/llvm/ir#InstLoad
//    *ir.InstStore           // https://godoc.org/github.com/llir/llvm/ir#InstStore
//    *ir.InstGetElementPtr   // https://godoc.org/github.com/llir/llvm/ir#InstGetElementPtr
//
// Conversion instructions
//
// https://llvm.org/docs/LangRef.html#conversion-operations
//
//    *ir.InstTrunc           // https://godoc.org/github.com/llir/llvm/ir#InstTrunc
//    *ir.InstSExt            // https://godoc.org/github.com/llir/llvm/ir#InstSExt
//    *ir.InstFPTrunc         // https://godoc.org/github.com/llir/llvm/ir#InstFPTrunc
//    *ir.InstFPExt           // https://godoc.org/github.com/llir/llvm/ir#InstFPExt
//    *ir.InstFPToUI          // https://godoc.org/github.com/llir/llvm/ir#InstFPToUI
//    *ir.InstFPToSI          // https://godoc.org/github.com/llir/llvm/ir#InstFPToSI
//    *ir.InstUIToFP          // https://godoc.org/github.com/llir/llvm/ir#InstUIToFP
//    *ir.InstSIToFP          // https://godoc.org/github.com/llir/llvm/ir#InstSIToFP
//    *ir.InstBitCast         // https://godoc.org/github.com/llir/llvm/ir#InstBitCast
//
// Other instructions
//
// https://llvm.org/docs/LangRef.html#other-operations
//
//    *ir.InstICmp         // https://godoc.org/github.com/llir/llvm/ir#InstICmp
//    *ir.InstFCmp         // https://godoc.org/github.com/llir/llvm/ir#InstFCmp
//    *ir.InstCall         // https://godoc.org/github.com/llir/llvm/ir#InstCall

type irWriter interface {
	writeIR(io.Writer) error
}

type Instruction interface {
	irWriter
	// isInstruction ensures that only instructions can be assigned to the
	// instruction.Instruction interface.
	isInstruction()
}

// === [ ir.Instruction ] ======================================================

// Binary instructions.
func (*InstFNeg) isInstruction() {}

// Binary instructions.
func (*InstAdd) isInstruction()  {}
func (*InstFAdd) isInstruction() {}
func (*InstSub) isInstruction()  {}
func (*InstFSub) isInstruction() {}
func (*InstMul) isInstruction()  {}
func (*InstFMul) isInstruction() {}
func (*InstUDiv) isInstruction() {}
func (*InstSDiv) isInstruction() {}
func (*InstFDiv) isInstruction() {}
func (*InstURem) isInstruction() {}
func (*InstSRem) isInstruction() {}
func (*InstFRem) isInstruction() {}

// Bitwise instructions.
func (*InstShl) isInstruction()  {}
func (*InstLShr) isInstruction() {}
func (*InstAShr) isInstruction() {}
func (*InstAnd) isInstruction()  {}
func (*InstOr) isInstruction()   {}
func (*InstXor) isInstruction()  {}

// Memory instructions.
func (*InstAlloca) isInstruction()        {}
func (*InstLoad) isInstruction()          {}
func (*InstStore) isInstruction()         {}
func (*InstFence) isInstruction()         {}
func (*InstCmpXchg) isInstruction()       {}
func (*InstAtomicRMW) isInstruction()     {}
func (*InstGetElementPtr) isInstruction() {}

// Conversion instructions.
func (*InstTrunc) isInstruction()   {}
func (*InstSExt) isInstruction()    {}
func (*InstFPTrunc) isInstruction() {}
func (*InstFPExt) isInstruction()   {}
func (*InstFPToUI) isInstruction()  {}
func (*InstFPToSI) isInstruction()  {}
func (*InstUIToFP) isInstruction()  {}
func (*InstSIToFP) isInstruction()  {}
func (*InstBitCast) isInstruction() {}

// Other instructions.
func (*InstICmp) isInstruction() {}
func (*InstFCmp) isInstruction() {}
func (*InstCall) isInstruction() {}

// Terminator instructions
func (*TermRet) isInstruction()         {}
func (*TermBr) isInstruction()          {}
func (*TermCondBr) isInstruction()      {}
func (*TermSwitch) isInstruction()      {}
func (*TermResume) isInstruction()      {}
func (*TermCatchSwitch) isInstruction() {}
func (*TermCatchRet) isInstruction()    {}
func (*TermCleanupRet) isInstruction()  {}
func (*TermUnreachable) isInstruction() {}
