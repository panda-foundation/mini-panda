package types

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/ir/core"
)

type PointerType struct {
	ElemType core.Type
}

func NewPointerType(elemType core.Type) *PointerType {
	p := &PointerType{
		ElemType: elemType,
	}
	return p
}

func (t *PointerType) Equal(u core.Type) bool {
	if u, ok := u.(*PointerType); ok {
		return t.ElemType.Equal(u.ElemType)
	}
	return false
}

func (t *PointerType) String() string {
	return fmt.Sprintf("%s*", t.ElemType.String())
}

func (t *PointerType) WriteIR(w io.Writer) error {
	err := t.ElemType.WriteIR(w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("*"))
	return err
}
