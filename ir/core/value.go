package core

import (
	"github.com/panda-io/micro-panda/ir/constant"
)

// Value is an LLVM IR value, which may be used as an operand of instructions
// and terminators.
//
// A Value has one of the following underlying types.
//
//    constant.Constant   // https://godoc.org/github.com/llir/llvm/ir/constant#Constant
//    value.Named         // https://godoc.org/github.com/llir/llvm/ir/value#Named
type Value interface {
	// Type returns the type of the value.
	Type() Type
	// Ident returns the identifier associated with the value.
	Ident() string
}

func IsConstant(v Value) bool {
	_, ok := v.(constant.Constant)
	return ok
}
