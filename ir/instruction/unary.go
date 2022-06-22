package instruction

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/ir/core"
)

type InstFNeg struct {
	core.LocalIdent
	X   core.Value // floating-point scalar or floating-point vector
	Typ core.Type
}

func NewFNeg(x core.Value) *InstFNeg {
	inst := &InstFNeg{X: x}
	inst.Type()
	return inst
}

func (inst *InstFNeg) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstFNeg) writeIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fneg %s", inst.Ident(), inst.X)
	return err
}
