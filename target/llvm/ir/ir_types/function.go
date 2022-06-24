package ir_types

import (
	"io"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type FuncType struct {
	Qualified string //TO-DO
	RetType   ir.Type
	Params    []ir.Type
}

func NewFuncType(qualified string, retType ir.Type, params ...ir.Type) *FuncType {
	return &FuncType{
		Qualified: qualified,
		RetType:   retType,
		Params:    params,
	}
}

func (t *FuncType) Equal(u ir.Type) bool {
	if u, ok := u.(*FuncType); ok {
		if !t.RetType.Equal(u.RetType) {
			return false
		}
		if len(t.Params) != len(u.Params) {
			return false
		}
		for i := range t.Params {
			if !t.Params[i].Equal(u.Params[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (t *FuncType) String() string {
	return t.Qualified
}

func (t *FuncType) WriteIR(w io.Writer) error {
	err := t.RetType.WriteIR(w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(" ("))
	if err != nil {
		return err
	}
	for i, param := range t.Params {
		if i != 0 {
			_, err = w.Write([]byte(", "))
			if err != nil {
				return err
			}
		}
		err = param.WriteIR(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte(")"))
	return err
}
