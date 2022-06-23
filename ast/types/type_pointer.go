package types

import (
	"fmt"

	"github.com/panda-io/micro-panda/ast/core"
)

type TypePointer struct {
	TypeBase
	ElementType core.Type
}

func (p *TypePointer) Equal(t core.Type) bool {
	if pointer, ok := t.(*TypePointer); ok {
		return p.ElementType.Equal(pointer.ElementType)
	} else if array, ok := t.(*TypeArray); ok {
		if len(array.Dimension) == 1 {
			return p.ElementType.Equal(array.ElementType)
		}
	}
	return false
}

func (p *TypePointer) String() string {
	return fmt.Sprintf("pointer<%s>", p.ElementType.String())
}
