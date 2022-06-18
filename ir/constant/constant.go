package constant

import (
	"io"

	"github.com/panda-io/micro-panda/types"
	"github.com/panda-io/micro-panda/value"
)

type irWriter interface {
	writeIR(io.Writer) error
}

type Constant interface {
	value.Value
	isConstant()
}

// === [ Constants ] ===

// Convenience constants.
var (
	// Boolean constants.
	True  = NewInt(types.I1, 1) // true
	False = NewInt(types.I1, 0) // false
)

// Constant is an LLVM IR constant; a value that is immutable at runtime, such
// as an integer or floating-point literal, or the address of a function or
// global variable.
//
// A Constant has one of the following underlying types.
//
// Simple constants
//
// https://llvm.org/docs/LangRef.html#simple-constants
//
//    *constant.Int         // https://godoc.org/github.com/llir/llvm/ir/constant#Int
//    *constant.Float       // https://godoc.org/github.com/llir/llvm/ir/constant#Float
//    *constant.Null        // https://godoc.org/github.com/llir/llvm/ir/constant#Null
//
// Complex constants
//
// https://llvm.org/docs/LangRef.html#complex-constants
//
//    *constant.Struct            // https://godoc.org/github.com/llir/llvm/ir/constant#Struct
//    *constant.Array             // https://godoc.org/github.com/llir/llvm/ir/constant#Array
//    *constant.CharArray         // https://godoc.org/github.com/llir/llvm/ir/constant#CharArray
//
// Global variable and function addresses
//
// https://llvm.org/docs/LangRef.html#global-variable-and-function-addresses
//
//    *ir.Global   // https://godoc.org/github.com/llir/llvm/ir#Global
//    *ir.Func     // https://godoc.org/github.com/llir/llvm/ir#Func
//
//
// Addresses of basic blocks
//
// https://llvm.org/docs/LangRef.html#addresses-of-basic-blocks
//
//    *constant.BlockAddress   // https://godoc.org/github.com/llir/llvm/ir/constant#BlockAddress
//
// Constant expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    constant.Expression   // https://godoc.org/github.com/llir/llvm/ir/constant#Expression

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Int) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Float) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Null) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Struct) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Array) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*CharArray) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ZeroInitializer) isConstant() {}
