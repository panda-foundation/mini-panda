package instruction

import (
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

// === [ Instructions ] ========================================================

// Instruction is an LLVM IR instruction. All instructions (except store and
// fence) implement the value.Named interface and may thus be used directly as
// values.
//
// An Instruction has one of the following underlying ir_types.
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
//    *InstAdd    // https://godoc.org/github.com/llir/llvm/ir#InstAdd
//    *InstFAdd   // https://godoc.org/github.com/llir/llvm/ir#InstFAdd
//    *InstSub    // https://godoc.org/github.com/llir/llvm/ir#InstSub
//    *InstFSub   // https://godoc.org/github.com/llir/llvm/ir#InstFSub
//    *InstMul    // https://godoc.org/github.com/llir/llvm/ir#InstMul
//    *InstFMul   // https://godoc.org/github.com/llir/llvm/ir#InstFMul
//    *InstUDiv   // https://godoc.org/github.com/llir/llvm/ir#InstUDiv
//    *InstSDiv   // https://godoc.org/github.com/llir/llvm/ir#InstSDiv
//    *InstFDiv   // https://godoc.org/github.com/llir/llvm/ir#InstFDiv
//    *InstURem   // https://godoc.org/github.com/llir/llvm/ir#InstURem
//    *InstSRem   // https://godoc.org/github.com/llir/llvm/ir#InstSRem
//    *InstFRem   // https://godoc.org/github.com/llir/llvm/ir#InstFRem
//
// Bitwise instructions
//
// https://llvm.org/docs/LangRef.html#bitwise-binary-operations
//
//    *InstShl    // https://godoc.org/github.com/llir/llvm/ir#InstShl
//    *InstLShr   // https://godoc.org/github.com/llir/llvm/ir#InstLShr
//    *InstAShr   // https://godoc.org/github.com/llir/llvm/ir#InstAShr
//    *InstAnd    // https://godoc.org/github.com/llir/llvm/ir#InstAnd
//    *InstOr     // https://godoc.org/github.com/llir/llvm/ir#InstOr
//    *InstXor    // https://godoc.org/github.com/llir/llvm/ir#InstXor
//
// Memory instructions
//
// https://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations
//
//    *InstAlloca          // https://godoc.org/github.com/llir/llvm/ir#InstAlloca
//    *InstLoad            // https://godoc.org/github.com/llir/llvm/ir#InstLoad
//    *InstStore           // https://godoc.org/github.com/llir/llvm/ir#InstStore
//    *InstGetElementPtr   // https://godoc.org/github.com/llir/llvm/ir#InstGetElementPtr
//
// Conversion instructions
//
// https://llvm.org/docs/LangRef.html#conversion-operations
//
//    *InstTrunc           // https://godoc.org/github.com/llir/llvm/ir#InstTrunc
//    *InstSExt            // https://godoc.org/github.com/llir/llvm/ir#InstSExt
//    *InstFPTrunc         // https://godoc.org/github.com/llir/llvm/ir#InstFPTrunc
//    *InstFPExt           // https://godoc.org/github.com/llir/llvm/ir#InstFPExt
//    *InstFPToUI          // https://godoc.org/github.com/llir/llvm/ir#InstFPToUI
//    *InstFPToSI          // https://godoc.org/github.com/llir/llvm/ir#InstFPToSI
//    *InstUIToFP          // https://godoc.org/github.com/llir/llvm/ir#InstUIToFP
//    *InstSIToFP          // https://godoc.org/github.com/llir/llvm/ir#InstSIToFP
//    *InstBitCast         // https://godoc.org/github.com/llir/llvm/ir#InstBitCast
//
// Other instructions
//
// https://llvm.org/docs/LangRef.html#other-operations
//
//    *InstICmp         // https://godoc.org/github.com/llir/llvm/ir#InstICmp
//    *InstFCmp         // https://godoc.org/github.com/llir/llvm/ir#InstFCmp
//    *InstCall         // https://godoc.org/github.com/llir/llvm/ir#InstCall
//
// Terminator instructions
//
// https://llvm.org/docs/LangRef.html#terminator-instructions
//
//    *TermRet           // https://godoc.org/github.com/llir/llvm/ir#TermRet
//    *TermBr            // https://godoc.org/github.com/llir/llvm/ir#TermBr
//    *TermCondBr        // https://godoc.org/github.com/llir/llvm/ir#TermCondBr
//    *TermSwitch        // https://godoc.org/github.com/llir/llvm/ir#TermSwitch

type Instruction interface {
	ir.IRWriter
	IsInstruction()
}

//TO-DO remove String()

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
