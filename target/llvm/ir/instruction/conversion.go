package instruction

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

type InstTrunc struct {
	ir.LocalIdent
	From ir.Value
	To   ir.Type
}

func NewTrunc(from ir.Value, to ir.Type) *InstTrunc {
	fromType := from.Type()
	toType := to
	if fromIntT, ok := fromType.(*ir_types.IntType); ok {
		toIntT, ok := toType.(*ir_types.IntType)
		if !ok {
			panic(fmt.Errorf("trunc operands are not compatible: from=%v; to=%T", fromIntT, to))
		}
		fromSize := fromIntT.BitSize
		toSize := toIntT.BitSize
		if fromSize < toSize {
			panic(fmt.Errorf("invalid trunc operands: from.BitSize < to.BitSize (%v is smaller than %v)", from.Type(), to))
		}
	}
	return &InstTrunc{From: from, To: to}
}

func (inst *InstTrunc) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstTrunc) Type() ir.Type {
	return inst.To
}

func (inst *InstTrunc) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = trunc %s to %s", inst.Ident(), inst.From.String(), inst.To.String())
	return err
}

type InstSExt struct {
	ir.LocalIdent
	From ir.Value
	To   ir.Type
}

func NewSExt(from ir.Value, to ir.Type) *InstSExt {
	return &InstSExt{From: from, To: to}
}

func (inst *InstSExt) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstSExt) Type() ir.Type {
	return inst.To
}

func (inst *InstSExt) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = sext %s to %s", inst.Ident(), inst.From.String(), inst.To.String())
	return err
}

type InstFPTrunc struct {
	ir.LocalIdent
	From ir.Value
	To   ir.Type
}

func NewFPTrunc(from ir.Value, to ir.Type) *InstFPTrunc {
	return &InstFPTrunc{From: from, To: to}
}

func (inst *InstFPTrunc) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstFPTrunc) Type() ir.Type {
	return inst.To
}

func (inst *InstFPTrunc) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fptrunc %s to %s", inst.Ident(), inst.From.String(), inst.To.String())
	return err
}

type InstFPExt struct {
	ir.LocalIdent
	From ir.Value
	To   ir.Type
}

func NewFPExt(from ir.Value, to ir.Type) *InstFPExt {
	return &InstFPExt{From: from, To: to}
}

func (inst *InstFPExt) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstFPExt) Type() ir.Type {
	return inst.To
}

func (inst *InstFPExt) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fpext %s to %s", inst.Ident(), inst.From.String(), inst.To.String())
	return err
}

type InstFPToUI struct {
	ir.LocalIdent
	From ir.Value
	To   ir.Type
}

func NewFPToUI(from ir.Value, to ir.Type) *InstFPToUI {
	return &InstFPToUI{From: from, To: to}
}

func (inst *InstFPToUI) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstFPToUI) Type() ir.Type {
	return inst.To
}

func (inst *InstFPToUI) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fptoui %s to %s", inst.Ident(), inst.From.String(), inst.To.String())
	return err
}

type InstFPToSI struct {
	ir.LocalIdent
	From ir.Value
	To   ir.Type
}

func NewFPToSI(from ir.Value, to ir.Type) *InstFPToSI {
	return &InstFPToSI{From: from, To: to}
}

func (inst *InstFPToSI) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstFPToSI) Type() ir.Type {
	return inst.To
}

func (inst *InstFPToSI) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fptosi %s to %s", inst.Ident(), inst.From.String(), inst.To.String())
	return err
}

type InstUIToFP struct {
	ir.LocalIdent
	From ir.Value
	To   ir.Type
}

func NewUIToFP(from ir.Value, to ir.Type) *InstUIToFP {
	return &InstUIToFP{From: from, To: to}
}

func (inst *InstUIToFP) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstUIToFP) Type() ir.Type {
	return inst.To
}

func (inst *InstUIToFP) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = uitofp %s to %s", inst.Ident(), inst.From.String(), inst.To.String())
	return err
}

type InstSIToFP struct {
	ir.LocalIdent
	From ir.Value
	To   ir.Type
}

func NewSIToFP(from ir.Value, to ir.Type) *InstSIToFP {
	return &InstSIToFP{From: from, To: to}
}

func (inst *InstSIToFP) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstSIToFP) Type() ir.Type {
	return inst.To
}

func (inst *InstSIToFP) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = sitofp %s to %s", inst.Ident(), inst.From.String(), inst.To.String())
	return err
}

type InstBitCast struct {
	ir.LocalIdent
	From ir.Value
	To   ir.Type
}

func NewBitCast(from ir.Value, to ir.Type) *InstBitCast {
	return &InstBitCast{From: from, To: to}
}

func (inst *InstBitCast) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstBitCast) Type() ir.Type {
	return inst.To
}

func (inst *InstBitCast) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = bitcast %s to %s", inst.Ident(), inst.From.String(), inst.To.String())
	return err
}
