package instruction

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/ir/core"
)

type InstTrunc struct {
	core.LocalIdent
	From core.Value
	To   core.Type
}

func NewTrunc(from core.Value, to core.Type) *InstTrunc {
	fromType := from.Type()
	toType := to
	if fromIntT, ok := fromType.(*core.IntType); ok {
		toIntT, ok := toType.(*core.IntType)
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
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstTrunc) Type() core.Type {
	return inst.To
}

func (inst *InstTrunc) writeIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = trunc %s to %s", inst.Ident(), inst.From, inst.To)
	return err
}

type InstSExt struct {
	core.LocalIdent
	From core.Value
	To   core.Type
}

func NewSExt(from core.Value, to core.Type) *InstSExt {
	return &InstSExt{From: from, To: to}
}

func (inst *InstSExt) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstSExt) Type() core.Type {
	return inst.To
}

func (inst *InstSExt) writeIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = sext %s to %s", inst.Ident(), inst.From, inst.To)
	return err
}

type InstFPTrunc struct {
	core.LocalIdent
	From core.Value
	To   core.Type
}

func NewFPTrunc(from core.Value, to core.Type) *InstFPTrunc {
	return &InstFPTrunc{From: from, To: to}
}

func (inst *InstFPTrunc) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstFPTrunc) Type() core.Type {
	return inst.To
}

func (inst *InstFPTrunc) writeIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fptrunc %s to %s", inst.Ident(), inst.From, inst.To)
	return err
}

type InstFPExt struct {
	core.LocalIdent
	From core.Value
	To   core.Type
}

func NewFPExt(from core.Value, to core.Type) *InstFPExt {
	return &InstFPExt{From: from, To: to}
}

func (inst *InstFPExt) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstFPExt) Type() core.Type {
	return inst.To
}

func (inst *InstFPExt) writeIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fpext %s to %s", inst.Ident(), inst.From, inst.To)
	return err
}

type InstFPToUI struct {
	core.LocalIdent
	From core.Value
	To   core.Type
}

func NewFPToUI(from core.Value, to core.Type) *InstFPToUI {
	return &InstFPToUI{From: from, To: to}
}

func (inst *InstFPToUI) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstFPToUI) Type() core.Type {
	return inst.To
}

func (inst *InstFPToUI) writeIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fptoui %s to %s", inst.Ident(), inst.From, inst.To)
	return err
}

type InstFPToSI struct {
	core.LocalIdent
	From core.Value
	To   core.Type
}

func NewFPToSI(from core.Value, to core.Type) *InstFPToSI {
	return &InstFPToSI{From: from, To: to}
}

func (inst *InstFPToSI) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstFPToSI) Type() core.Type {
	return inst.To
}

func (inst *InstFPToSI) writeIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fptosi %s to %s", inst.Ident(), inst.From, inst.To)
	return err
}

type InstUIToFP struct {
	core.LocalIdent
	From core.Value
	To   core.Type
}

func NewUIToFP(from core.Value, to core.Type) *InstUIToFP {
	return &InstUIToFP{From: from, To: to}
}

func (inst *InstUIToFP) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstUIToFP) Type() core.Type {
	return inst.To
}

func (inst *InstUIToFP) writeIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = uitofp %s to %s", inst.Ident(), inst.From, inst.To)
	return err
}

type InstSIToFP struct {
	core.LocalIdent
	From core.Value
	To   core.Type
}

func NewSIToFP(from core.Value, to core.Type) *InstSIToFP {
	return &InstSIToFP{From: from, To: to}
}

func (inst *InstSIToFP) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstSIToFP) Type() core.Type {
	return inst.To
}

func (inst *InstSIToFP) writeIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = sitofp %s to %s", inst.Ident(), inst.From, inst.To)
	return err
}

type InstBitCast struct {
	core.LocalIdent
	From core.Value
	To   core.Type
}

func NewBitCast(from core.Value, to core.Type) *InstBitCast {
	return &InstBitCast{From: from, To: to}
}

func (inst *InstBitCast) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstBitCast) Type() core.Type {
	return inst.To
}

func (inst *InstBitCast) writeIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = bitcast %s to %s", inst.Ident(), inst.From, inst.To)
	return err
}
