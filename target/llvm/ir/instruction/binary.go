package instruction

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type InstAdd struct {
	ir.LocalIdent
	X, Y ir.Value
	Typ  ir.Type
}

func NewAdd(x, y ir.Value) *InstAdd {
	inst := &InstAdd{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstAdd) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstAdd) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstAdd) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = add %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstFAdd struct {
	ir.LocalIdent
	X, Y ir.Value // floating-point scalar or floating-point vector
	Typ  ir.Type
}

func NewFAdd(x, y ir.Value) *InstFAdd {
	inst := &InstFAdd{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstFAdd) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstFAdd) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstFAdd) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fadd %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstSub struct {
	ir.LocalIdent
	X, Y ir.Value // integer scalar or integer vector
	Typ  ir.Type
}

func NewSub(x, y ir.Value) *InstSub {
	inst := &InstSub{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstSub) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstSub) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstSub) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = sub %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstFSub struct {
	ir.LocalIdent
	X, Y ir.Value // floating-point scalar or floating-point vector
	Typ  ir.Type
}

func NewFSub(x, y ir.Value) *InstFSub {
	inst := &InstFSub{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstFSub) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstFSub) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstFSub) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fsub %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstMul struct {
	ir.LocalIdent
	X, Y ir.Value // integer scalar or integer vector
	Typ  ir.Type
}

func NewMul(x, y ir.Value) *InstMul {
	inst := &InstMul{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstMul) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstMul) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstMul) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = mul %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstFMul struct {
	ir.LocalIdent
	X, Y ir.Value // floating-point scalar or floating-point vector
	Typ  ir.Type
}

func NewFMul(x, y ir.Value) *InstFMul {
	inst := &InstFMul{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstFMul) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstFMul) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstFMul) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fmul %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstUDiv struct {
	ir.LocalIdent
	X, Y ir.Value // integer scalar or integer vector
	Typ  ir.Type
}

func NewUDiv(x, y ir.Value) *InstUDiv {
	inst := &InstUDiv{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstUDiv) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstUDiv) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstUDiv) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = udiv %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstSDiv struct {
	ir.LocalIdent
	X, Y ir.Value // integer scalar or integer vector
	Typ  ir.Type
}

func NewSDiv(x, y ir.Value) *InstSDiv {
	inst := &InstSDiv{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstSDiv) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstSDiv) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstSDiv) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = sdiv %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstFDiv struct {
	ir.LocalIdent
	X, Y ir.Value // floating-point scalar or floating-point vector
	Typ  ir.Type
}

func NewFDiv(x, y ir.Value) *InstFDiv {
	inst := &InstFDiv{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstFDiv) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstFDiv) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstFDiv) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fdiv %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstURem struct {
	ir.LocalIdent
	X, Y ir.Value // integer scalar or integer vector
	Typ  ir.Type
}

func NewURem(x, y ir.Value) *InstURem {
	inst := &InstURem{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstURem) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstURem) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstURem) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = urem %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstSRem struct {
	ir.LocalIdent
	X, Y ir.Value // integer scalar or integer vector
	Typ  ir.Type
}

func NewSRem(x, y ir.Value) *InstSRem {
	inst := &InstSRem{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstSRem) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstSRem) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstSRem) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = srem %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}

type InstFRem struct {
	ir.LocalIdent
	X, Y ir.Value // floating-point scalar or floating-point vector
	Typ  ir.Type
}

func NewFRem(x, y ir.Value) *InstFRem {
	inst := &InstFRem{X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstFRem) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstFRem) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

func (inst *InstFRem) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = frem %s, %s", inst.Ident(), inst.X.String(), inst.Y.Ident())
	return err
}
