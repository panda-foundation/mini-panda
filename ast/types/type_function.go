package types

import (
	"strings"

	"github.com/panda-io/micro-panda/ast/core"
)

type TypeFunction struct {
	TypeBase
	ReturnType core.Type
	Parameters []core.Type

	MemberFunction bool
	Extern         bool
	ExternName     string
	TypeDefine     bool
}

func (f *TypeFunction) Equal(t core.Type) bool {
	if t, ok := t.(*TypeFunction); ok {
		if f.ReturnType != nil && t.ReturnType != nil {
			if !f.ReturnType.Equal(t.ReturnType) {
				return false
			}
		} else if f.ReturnType != t.ReturnType {
			return false
		}
		if len(f.Parameters) != len(t.Parameters) {
			return false
		}
		for i := 0; i < len(f.Parameters); i++ {
			if !f.Parameters[i].Equal(t.Parameters[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (f *TypeFunction) String() string {
	var b strings.Builder
	b.WriteString("function(")
	for i, t := range f.Parameters {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(t.String())
	}
	b.WriteString(")->")
	if f.ReturnType == nil {
		b.WriteString("null")
	} else {
		b.WriteString(f.ReturnType.String())
	}
	return b.String()
}
