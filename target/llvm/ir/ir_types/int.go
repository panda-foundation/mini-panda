package ir_types

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type IntType struct {
	BitSize  int
	Unsigned bool
}

func NewIntType(bitSize int) *IntType {
	return &IntType{
		BitSize: bitSize,
	}
}

func (t *IntType) Equal(u ir.Type) bool {
	if u, ok := u.(*IntType); ok {
		return t.BitSize == u.BitSize
	}
	return false
}

func (t *IntType) String() string {
	return fmt.Sprintf("i%d", t.BitSize)
}

func (t *IntType) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "i%d", t.BitSize)
	return err
}
