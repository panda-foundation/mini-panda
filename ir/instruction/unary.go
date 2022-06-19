package instruction

import (
	"fmt"
	"strings"

	"github.com/panda-io/micro-panda/ir/core"
)

// --- [ Unary instructions ] --------------------------------------------------

// ~~~ [ fneg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFNeg is an LLVM IR fneg instruction.
type InstFNeg struct {
	core.LocalIdent
	X   core.Value // floating-point scalar or floating-point vector
	Typ core.Type
}

// NewFNeg returns a new fneg instruction based on the given operand.
func NewFNeg(x core.Value) *InstFNeg {
	inst := &InstFNeg{X: x}
	inst.Type()
	return inst
}

// Type returns the type of the instruction.
func (inst *InstFNeg) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
// 'fneg' FastMathFlags=FastMathFlag* X=TypeValue Metadata=(',' MetadataAttachment)+?
func (inst *InstFNeg) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("fneg")
	fmt.Fprintf(buf, " %s", inst.X)
	return buf.String()
}
