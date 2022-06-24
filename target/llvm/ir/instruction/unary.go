package instruction

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type InstFNeg struct {
	ir.LocalIdent
	X   ir.Value // floating-point scalar or floating-point vector
	Typ ir.Type
}

func NewFNeg(x ir.Value) *InstFNeg {
	inst := &InstFNeg{X: x}
	inst.Type()
	return inst
}

func (inst *InstFNeg) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstFNeg) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fneg %s", inst.Ident(), inst.X.String())
	return err
}
