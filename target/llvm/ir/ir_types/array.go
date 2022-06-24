package ir_types

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type ArrayType struct {
	Len      uint64
	ElemType ir.Type
}

func NewArrayType(len uint64, elemType ir.Type) *ArrayType {
	return &ArrayType{
		Len:      len,
		ElemType: elemType,
	}
}

func (t *ArrayType) Equal(u ir.Type) bool {
	if u, ok := u.(*ArrayType); ok {
		if t.Len != u.Len {
			return false
		}
		return t.ElemType.Equal(u.ElemType)
	}
	return false
}

func (t *ArrayType) String() string {
	return fmt.Sprintf("[%d x %s]", t.Len, t.ElemType.String())
}

func (t *ArrayType) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "[%d x ", t.Len)
	if err != nil {
		return nil
	}
	err = t.ElemType.WriteIR(w)
	if err != nil {
		return nil
	}
	_, err = w.Write([]byte("]"))
	return err
}
