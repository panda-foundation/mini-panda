package instruction

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/ir/core"
)

type InstAdd struct {
	core.LocalIdent
	X, Y core.Value
	Typ  core.Type
}

func NewAdd(x, y core.Value) *InstAdd {
	inst := &InstAdd{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstAdd) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstAdd) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstAdd) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = add %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstFAdd struct {
	core.LocalIdent
	X, Y core.Value // floating-point scalar or floating-point vector
	Typ  core.Type
}

func NewFAdd(x, y core.Value) *InstFAdd {
	inst := &InstFAdd{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstFAdd) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstFAdd) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstFAdd) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fadd %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstSub struct {
	core.LocalIdent
	X, Y core.Value // integer scalar or integer vector
	Typ  core.Type
}

func NewSub(x, y core.Value) *InstSub {
	inst := &InstSub{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstSub) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstSub) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstSub) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = sub %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstFSub struct {
	core.LocalIdent
	X, Y core.Value // floating-point scalar or floating-point vector
	Typ  core.Type
}

func NewFSub(x, y core.Value) *InstFSub {
	inst := &InstFSub{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstFSub) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstFSub) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstFSub) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fsub %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstMul struct {
	core.LocalIdent
	X, Y core.Value // integer scalar or integer vector
	Typ  core.Type
}

func NewMul(x, y core.Value) *InstMul {
	inst := &InstMul{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstMul) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstMul) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstMul) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = mul %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstFMul struct {
	core.LocalIdent
	X, Y core.Value // floating-point scalar or floating-point vector
	Typ  core.Type
}

func NewFMul(x, y core.Value) *InstFMul {
	inst := &InstFMul{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstFMul) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstFMul) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstFMul) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fmul %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstUDiv struct {
	core.LocalIdent
	X, Y core.Value // integer scalar or integer vector
	Typ  core.Type
}

func NewUDiv(x, y core.Value) *InstUDiv {
	inst := &InstUDiv{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstUDiv) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstUDiv) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstUDiv) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = udiv %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstSDiv struct {
	core.LocalIdent
	X, Y core.Value // integer scalar or integer vector
	Typ  core.Type
}

func NewSDiv(x, y core.Value) *InstSDiv {
	inst := &InstSDiv{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstSDiv) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstSDiv) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstSDiv) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = sdiv %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstFDiv struct {
	core.LocalIdent
	X, Y core.Value // floating-point scalar or floating-point vector
	Typ  core.Type
}

func NewFDiv(x, y core.Value) *InstFDiv {
	inst := &InstFDiv{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstFDiv) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstFDiv) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstFDiv) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fdiv %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstURem struct {
	core.LocalIdent
	X, Y core.Value // integer scalar or integer vector
	Typ  core.Type
}

func NewURem(x, y core.Value) *InstURem {
	inst := &InstURem{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstURem) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstURem) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstURem) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = urem %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstSRem struct {
	core.LocalIdent
	X, Y core.Value // integer scalar or integer vector
	Typ  core.Type
}

func NewSRem(x, y core.Value) *InstSRem {
	inst := &InstSRem{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstSRem) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstSRem) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstSRem) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = srem %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}

type InstFRem struct {
	core.LocalIdent
	X, Y core.Value // floating-point scalar or floating-point vector
	Typ  core.Type
}

func NewFRem(x, y core.Value) *InstFRem {
	inst := &InstFRem{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstFRem) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstFRem) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstFRem) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = frem %s, %s", inst.Ident(), inst.X, inst.Y.Ident())
	return err
}
