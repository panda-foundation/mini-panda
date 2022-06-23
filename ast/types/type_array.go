package types

import (
	"fmt"

	"github.com/panda-io/micro-panda/ast/core"
)

type TypeArray struct {
	TypeBase
	ElementType core.Type
	Dimension   []int
}

func (a *TypeArray) Equal(t core.Type) bool {
	if array, ok := t.(*TypeArray); ok {
		if len(a.Dimension) == len(array.Dimension) {
			for i := 1; i < len(a.Dimension); i++ {
				if a.Dimension[i] != array.Dimension[i] {
					return false
				}
			}
			return true
		}
	} else if pointer, ok := t.(*TypePointer); ok {
		if len(a.Dimension) == 1 {
			return a.ElementType.Equal(pointer.ElementType)
		}
	}
	return false
}

func (a *TypeArray) String() string {
	return fmt.Sprintf("array[%s]", a.ElementType.String())
}
