package instruction

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/ir/core"
)

type InstShl struct {
	core.LocalIdent
	X, Y core.Value // integer scalar or integer vector
	Typ  core.Type
}

func NewShl(x, y core.Value) *InstShl {
	inst := &InstShl{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstShl) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstShl) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstShl) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = shl %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstLShr struct {
	core.LocalIdent
	X, Y core.Value // integer scalars or vectors
	Typ  core.Type
}

func NewLShr(x, y core.Value) *InstLShr {
	inst := &InstLShr{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstLShr) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstLShr) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstLShr) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = lshr %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstAShr struct {
	core.LocalIdent
	X, Y core.Value // integer scalars or vectors
	Typ  core.Type
}

func NewAShr(x, y core.Value) *InstAShr {
	inst := &InstAShr{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstAShr) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstAShr) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstAShr) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = ashr %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstAnd struct {
	core.LocalIdent
	X, Y core.Value // integer scalars or vectors
	Typ  core.Type
}

func NewAnd(x, y core.Value) *InstAnd {
	inst := &InstAnd{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstAnd) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstAnd) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstAnd) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = and %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstOr struct {
	core.LocalIdent
	X, Y core.Value // integer scalars or vectors
	Typ  core.Type
}

func NewOr(x, y core.Value) *InstOr {
	inst := &InstOr{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstOr) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstOr) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstOr) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = or %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstXor struct {
	core.LocalIdent
	X, Y core.Value // integer scalars or vectors
	Typ  core.Type
}

func NewXor(x, y core.Value) *InstXor {
	inst := &InstXor{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstXor) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstXor) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstXor) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = xor %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}
