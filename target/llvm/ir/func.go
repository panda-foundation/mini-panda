package ir

import (
	"fmt"
	"io"
	"strings"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

type Func struct {
	ir.GlobalIdent
	Sig    *ir_types.FuncType
	Params []*Param
	Blocks []*Block // nil if declaration.
	Typ    *ir_types.PointerType
}

func NewFunc(name string, retType ir.Type, params ...*Param) *Func {
	paramTypes := make([]ir.Type, len(params))
	for i, param := range params {
		paramTypes[i] = param.Type()
	}
	sig := ir_types.NewFuncType(name, retType, paramTypes...)
	f := &Func{Sig: sig, Params: params}
	f.SetName(name)
	f.Type()
	return f
}

func (f *Func) NewBlock(name string) *Block {
	block := NewBlock(name)
	f.Blocks = append(f.Blocks, block)
	return block
}

func (f *Func) String() string {
	return fmt.Sprintf("%s %s", f.Type().String(), f.Ident())
}

func (f *Func) Type() ir.Type {
	if f.Typ == nil {
		f.Typ = ir_types.NewPointerType(f.Sig)
	}
	return f.Typ
}

func (f *Func) WriteIR(w io.Writer) error {
	if err := f.assignIDs(); err != nil {
		panic(fmt.Errorf("unable to assign IDs of function %q; %v", f.Ident(), err))
	}
	if len(f.Blocks) == 0 {
		_, err := fmt.Fprintf(w, "declare %s", f.headerString())
		return err
	} else {
		_, err := fmt.Fprintf(w, "define %s %s", f.headerString(), f.bodyString())
		return err
	}
}

func (f *Func) IsConstant() {}

func (f *Func) assignIDs() error {
	id := int64(0)
	for _, param := range f.Params {
		if err := f.setName(param, &id); err != nil {
			return err
		}
	}
	for _, block := range f.Blocks {
		if err := f.setName(block, &id); err != nil {
			return err
		}
		for _, inst := range block.Insts {
			n, ok := inst.(ir.Ident)
			if !ok {
				continue
			}
			if inst.Type().Equal(ir_types.Void) {
				continue
			}
			if err := f.setName(n, &id); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *Func) headerString() string {
	buf := &strings.Builder{}
	_, _ = fmt.Fprintf(buf, " %s", f.Sig.RetType.String())
	_, _ = fmt.Fprintf(buf, " %s(", f.Ident())
	for i, param := range f.Params {
		if i != 0 {
			_, _ = buf.WriteString(", ")
		}
		_ = param.WriteIR(buf)
	}
	_, _ = buf.WriteString(")")
	return buf.String()
}

func (f *Func) bodyString() string {
	buf := &strings.Builder{}
	_, _ = buf.WriteString("{\n")
	for i, block := range f.Blocks {
		if i != 0 {
			_, _ = buf.WriteString("\n")
		}
		_ = block.WriteIR(buf)
		_, _ = buf.WriteString("\n")
	}
	_, _ = buf.WriteString("}")
	return buf.String()
}

func (f *Func) setName(name ir.Ident, id *int64) error {
	if name.Name() == "" {
		if name.ID() != 0 && *id != name.ID() {
			want := *id
			got := name.ID()
			return fmt.Errorf("invalid local ID, expected %s, got %s", ir.LocalID(want), ir.LocalID(got))
		}
		name.SetID(*id)
		(*id)++
	}
	return nil
}

type Param struct {
	ir.LocalIdent
	Typ ir.Type
}

func NewParam(typ ir.Type) *Param {
	return &Param{
		Typ: typ,
	}
}

func (p *Param) String() string {
	return fmt.Sprintf("%s %s", p.Type().String(), p.Ident())
}

func (p *Param) Type() ir.Type {
	return p.Typ
}

func (p *Param) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s %s", p.Typ.String(), p.Ident())
	return err
}
