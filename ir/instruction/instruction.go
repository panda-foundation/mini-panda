package instruction

import "github.com/panda-io/micro-panda/ir/core"

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
//
// Terminator instructions
//
// https://llvm.org/docs/LangRef.html#terminator-instructions
//
//    *ir.TermRet           // https://godoc.org/github.com/llir/llvm/ir#TermRet
//    *ir.TermBr            // https://godoc.org/github.com/llir/llvm/ir#TermBr
//    *ir.TermCondBr        // https://godoc.org/github.com/llir/llvm/ir#TermCondBr
//    *ir.TermSwitch        // https://godoc.org/github.com/llir/llvm/ir#TermSwitch

type Instruction interface {
	core.Value
	core.IRWriter
	IsInstruction()
}

// Binary instructions.
func (*InstFNeg) IsInstruction() {}

// Binary instructions.
func (*InstAdd) IsInstruction()  {}
func (*InstFAdd) IsInstruction() {}
func (*InstSub) IsInstruction()  {}
func (*InstFSub) IsInstruction() {}
func (*InstMul) IsInstruction()  {}
func (*InstFMul) IsInstruction() {}
func (*InstUDiv) IsInstruction() {}
func (*InstSDiv) IsInstruction() {}
func (*InstFDiv) IsInstruction() {}
func (*InstURem) IsInstruction() {}
func (*InstSRem) IsInstruction() {}
func (*InstFRem) IsInstruction() {}

// Bitwise instructions.
func (*InstShl) IsInstruction()  {}
func (*InstLShr) IsInstruction() {}
func (*InstAShr) IsInstruction() {}
func (*InstAnd) IsInstruction()  {}
func (*InstOr) IsInstruction()   {}
func (*InstXor) IsInstruction()  {}

// Memory instructions.
func (*InstAlloca) IsInstruction()        {}
func (*InstLoad) IsInstruction()          {}
func (*InstStore) IsInstruction()         {}
func (*InstGetElementPtr) IsInstruction() {}

// Conversion instructions.
func (*InstTrunc) IsInstruction()   {}
func (*InstSExt) IsInstruction()    {}
func (*InstFPTrunc) IsInstruction() {}
func (*InstFPExt) IsInstruction()   {}
func (*InstFPToUI) IsInstruction()  {}
func (*InstFPToSI) IsInstruction()  {}
func (*InstUIToFP) IsInstruction()  {}
func (*InstSIToFP) IsInstruction()  {}
func (*InstBitCast) IsInstruction() {}

// Other instructions.
func (*InstICmp) IsInstruction() {}
func (*InstFCmp) IsInstruction() {}
func (*InstCall) IsInstruction() {}

// Terminator instructions
func (*TermRet) IsInstruction()    {}
func (*TermBr) IsInstruction()     {}
func (*TermCondBr) IsInstruction() {}
func (*TermSwitch) IsInstruction() {}
