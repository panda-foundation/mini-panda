package instruction

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type InstShl struct {
	ir.LocalIdent
	X, Y ir.Value // integer scalar or integer vector
	Typ  ir.Type
}

func NewShl(x, y ir.Value) *InstShl {
	inst := &InstShl{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstShl) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstShl) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstShl) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = shl %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstLShr struct {
	ir.LocalIdent
	X, Y ir.Value // integer scalars or vectors
	Typ  ir.Type
}

func NewLShr(x, y ir.Value) *InstLShr {
	inst := &InstLShr{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstLShr) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstLShr) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstLShr) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = lshr %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstAShr struct {
	ir.LocalIdent
	X, Y ir.Value // integer scalars or vectors
	Typ  ir.Type
}

func NewAShr(x, y ir.Value) *InstAShr {
	inst := &InstAShr{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstAShr) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstAShr) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstAShr) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = ashr %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstAnd struct {
	ir.LocalIdent
	X, Y ir.Value // integer scalars or vectors
	Typ  ir.Type
}

func NewAnd(x, y ir.Value) *InstAnd {
	inst := &InstAnd{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstAnd) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstAnd) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstAnd) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = and %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstOr struct {
	ir.LocalIdent
	X, Y ir.Value // integer scalars or vectors
	Typ  ir.Type
}

func NewOr(x, y ir.Value) *InstOr {
	inst := &InstOr{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstOr) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstOr) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstOr) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = or %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstXor struct {
	ir.LocalIdent
	X, Y ir.Value // integer scalars or vectors
	Typ  ir.Type
}

func NewXor(x, y ir.Value) *InstXor {
	inst := &InstXor{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstXor) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstXor) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstXor) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = xor %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}
