package constant

import (
	"fmt"
	"strings"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

type Struct struct {
	Typ    *ir_types.StructType
	Fields []Constant
}

func NewStruct(t *ir_types.StructType, fields ...Constant) *Struct {
	c := &Struct{
		Fields: fields,
		Typ:    t,
	}
	c.Type()
	return c
}

func (c *Struct) String() string {
	return fmt.Sprintf("%s %s", c.Type().String(), c.Ident())
}

func (c *Struct) Type() ir.Type {
	if c.Typ == nil {
		var fieldTypes []ir.Type
		for _, field := range c.Fields {
			fieldTypes = append(fieldTypes, field.Type())
		}
		c.Typ = ir_types.NewStructType("", fieldTypes...)
	}
	return c.Typ
}

func (c *Struct) Ident() string {
	if len(c.Fields) == 0 {
		return "{}"
	}
	buf := &strings.Builder{}
	buf.WriteString("{ ")
	for i, field := range c.Fields {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(field.String())
	}
	buf.WriteString(" }")
	return buf.String()
}
