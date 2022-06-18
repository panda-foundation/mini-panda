// Package ir provides a definition of LLVM IR values.
package value

import (
	"github.com/panda-io/micro-panda/constant"
	"github.com/panda-io/micro-panda/types"
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
	Type() types.Type
	// Ident returns the identifier associated with the value.
	Ident() string
}

func IsConstant(v Value) bool {
	_, ok := v.(constant.Constant)
	return ok
}
