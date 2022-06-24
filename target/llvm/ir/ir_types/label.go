package ir_types

import (
	"io"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type LabelType struct {
}

func (t *LabelType) Equal(u ir.Type) bool {
	if _, ok := u.(*LabelType); ok {
		return true
	}
	return false
}

func (t *LabelType) String() string {
	return "label"
}

func (t *LabelType) WriteIR(w io.Writer) error {
	_, err := w.Write([]byte("label"))
	return err
}
