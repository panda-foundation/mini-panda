package types

import (
	"io"

	"github.com/panda-io/micro-panda/ir/core"
)

type VoidType struct {
}

func (t *VoidType) Equal(u core.Type) bool {
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
