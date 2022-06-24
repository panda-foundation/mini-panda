package ir_types

import (
	"io"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type VoidType struct {
}

func (t *VoidType) Equal(u ir.Type) bool {
	if _, ok := u.(*VoidType); ok {
		return true
	}
	return false
}

func (t *VoidType) String() string {
	return "void"
}

func (t *VoidType) WriteIR(w io.Writer) error {
	_, err := w.Write([]byte("void"))
	return err
}
