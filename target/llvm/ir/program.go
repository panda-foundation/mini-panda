package ir

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type Program struct {
	TypeDefs []ir.Type
	Globals  []*Global
	Funcs    []*Func
}

func NewProgram() *Program {
	return &Program{}
}

func (m *Program) WriteIR(w io.Writer) error {
	for _, t := range m.TypeDefs {
		_, err := fmt.Fprintf(w, "%s = type ", t.String())
		if err != nil {
			return err
		}
		err = t.WriteIR(w)
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}

	_, err := w.Write([]byte("\n"))
	if err != nil {
		return err
	}

	for _, g := range m.Globals {
		err := g.WriteIR(w)
		if err != nil {
			return err
		}
	}

	_, err = w.Write([]byte("\n"))
	if err != nil {
		return err
	}

	for i, f := range m.Funcs {
		if i != 0 {
			_, err := w.Write([]byte("\n"))
			if err != nil {
				return err
			}
		}
		err := f.WriteIR(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Program) NewFunc(name string, retType ir.Type, params ...*Param) *Func {
	f := NewFunc(name, retType, params...)
	m.Funcs = append(m.Funcs, f)
	return f
}

func (m *Program) NewGlobal(name string, contentType ir.Type) *Global {
	g := NewGlobal(name, contentType)
	m.Globals = append(m.Globals, g)
	return g
}

func (m *Program) NewGlobalDef(name string, init constant.Constant) *Global {
	g := NewGlobalDef(name, init)
	m.Globals = append(m.Globals, g)
	return g
}

func (m *Program) NewTypeDef(name string, typ ir.Type) ir.Type {
	//TO-DO if struct, set name
	//typ.SetName(name)
	m.TypeDefs = append(m.TypeDefs, typ)
	return typ
}
