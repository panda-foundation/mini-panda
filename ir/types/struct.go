package types

import (
	"io"

	"github.com/panda-io/micro-panda/ir/core"
)

type StructType struct {
	TypeName string
	Fields   []core.Type
}

func NewStructType(typeName string, fields ...core.Type) *StructType {
	return &StructType{
		TypeName: typeName,
		Fields:   fields,
	}
}

func (t *StructType) Equal(u core.Type) bool {
	if u, ok := u.(*StructType); ok {
		if len(t.Fields) != len(u.Fields) {
			return false
		}
		for i := range t.Fields {
			if !t.Fields[i].Equal(u.Fields[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (t *StructType) String() string {
	return t.TypeName
}

func (t *StructType) WriteIR(w io.Writer) error {
	if len(t.Fields) == 0 {
		_, err := w.Write([]byte("{}"))
		return err
	}
	_, err := w.Write([]byte("{ "))
	if err != nil {
		return err
	}
	for i, field := range t.Fields {
		if i != 0 {
			_, err := w.Write([]byte(", "))
			if err != nil {
				return err
			}
		}
		err = field.WriteIR(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte(" }"))
	return err
}
