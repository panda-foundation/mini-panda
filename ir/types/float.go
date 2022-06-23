package types

import (
	"io"

	"github.com/panda-io/micro-panda/ir/core"
)

type FloatType struct {
	Kind FloatKind
}

func (t *FloatType) Equal(u core.Type) bool {
	if u, ok := u.(*FloatType); ok {
		return t.Kind == u.Kind
	}
	return false
}

func (t *FloatType) String() string {
	return string(t.Kind)
}

func (t *FloatType) WriteIR(w io.Writer) error {
	_, err := w.Write([]byte(t.Kind))
	return err
}

type FloatKind string

const (
	FloatKindHalf   FloatKind = "half"
	FloatKindFloat  FloatKind = "float"
	FloatKindDouble FloatKind = "double"
)
