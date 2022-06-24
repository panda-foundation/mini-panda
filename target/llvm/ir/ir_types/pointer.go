package ir_types

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type PointerType struct {
	ElemType ir.Type
}

func NewPointerType(elemType ir.Type) *PointerType {
	p := &PointerType{
		ElemType: elemType,
	}
	return p
}

func (t *PointerType) Equal(u ir.Type) bool {
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
