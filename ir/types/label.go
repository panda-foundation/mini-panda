package types

import (
	"io"

	"github.com/panda-io/micro-panda/ir/core"
)

type LabelType struct {
}

func (t *LabelType) Equal(u core.Type) bool {
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
