package types

import (
	"io"

	"github.com/panda-io/micro-panda/ir/core"
)

type FuncType struct {
	TypeName string
	RetType  core.Type
	Params   []core.Type
}

func NewFuncType(typeName string, retType core.Type, params ...core.Type) *FuncType {
	return &FuncType{
		TypeName: typeName,
		RetType:  retType,
		Params:   params,
	}
}

func (t *FuncType) Equal(u core.Type) bool {
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
	return t.TypeName
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
