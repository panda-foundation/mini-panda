package constant

import (
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

func IsConstant(v ir.Value) bool {
	_, ok := v.(Constant)
	return ok
}

type Constant interface {
	ir.Value
	IsConstant()
}

var (
	True  = NewInt(ir_types.I1, 1) // true
	False = NewInt(ir_types.I1, 0) // false
)

// Constant is an LLVM IR constant; a value that is immutable at runtime, such
// as an integer or floating-point literal, or the address of a function or
// global variable.
//
// A Constant has one of the following underlying ir_types.
//
// Simple constants
//
// https://llvm.org/docs/LangRef.html#simple-constants
//
//    *Int         // https://godoc.org/github.com/llir/llvm/ir/constant#Int
//    *Float       // https://godoc.org/github.com/llir/llvm/ir/constant#Float
//    *Null        // https://godoc.org/github.com/llir/llvm/ir/constant#Null
//
// Complex constants
//
// https://llvm.org/docs/LangRef.html#complex-constants
//
//    *Struct            // https://godoc.org/github.com/llir/llvm/ir/constant#Struct
//    *Array             // https://godoc.org/github.com/llir/llvm/ir/constant#Array
//    *CharArray         // https://godoc.org/github.com/llir/llvm/ir/constant#CharArray
//
// Global variable and function addresses
//
// https://llvm.org/docs/LangRef.html#global-variable-and-function-addresses
//
//    *Global   // https://godoc.org/github.com/llir/llvm/ir#Global
//    *Func     // https://godoc.org/github.com/llir/llvm/ir#Func
//
// Constant expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    constant.Expression   // https://godoc.org/github.com/llir/llvm/ir/constant#Expression

func (*Int) IsConstant()   {}
func (*Float) IsConstant() {}
func (*Null) IsConstant()  {}

func (*Struct) IsConstant()          {}
func (*Array) IsConstant()           {}
func (*CharArray) IsConstant()       {}
func (*ZeroInitializer) IsConstant() {}
