package instruction

// Memory instructions
//
// https://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations
//
//    *InstAlloca          // https://godoc.org/github.com/llir/llvm/ir#InstAlloca
//    *InstLoad            // https://godoc.org/github.com/llir/llvm/ir#InstLoad
//    *InstStore           // https://godoc.org/github.com/llir/llvm/ir#InstStore
//    *InstGetElementPtr   // https://godoc.org/github.com/llir/llvm/ir#InstGetElementPtr
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
