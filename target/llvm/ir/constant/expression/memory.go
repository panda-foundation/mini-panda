package expression

import (
	"fmt"
	"strings"

	"github.com/panda-io/micro-panda/ir/constant"
	"github.com/panda-io/micro-panda/ir/core"
)

type ExprGetElementPtr struct {
	ElemType core.Type
	Src      constant.Constant
	Indices  []constant.Constant // *Int, *Vector or *Index
	Typ      core.Type           // *PointerType or *VectorType (with elements of pointer type)
}

func NewExprGetElementPtr(elemType core.Type, src constant.Constant, indices ...constant.Constant) *ExprGetElementPtr {
	e := &ExprGetElementPtr{ElemType: elemType, Src: src, Indices: indices}
	e.Type()
	return e
}

func (e *ExprGetElementPtr) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

func (e *ExprGetElementPtr) Type() core.Type {
	if e.Typ == nil {
		e.Typ = e.gepExprType(e.ElemType, e.Src.Type(), e.Indices)
	}
	return e.Typ
}

func (e *ExprGetElementPtr) Ident() string {
	buf := &strings.Builder{}
	buf.WriteString("getelementptr")
	fmt.Fprintf(buf, " (%s, %s", e.ElemType, e.Src)
	for _, index := range e.Indices {
		fmt.Fprintf(buf, ", %s", index)
	}
	buf.WriteString(")")
	return buf.String()
}

type Index struct {
	Index constant.Constant
}

func NewIndex(index constant.Constant) *Index {
	return &Index{Index: index}
}

func (index *Index) String() string {
	return index.Index.String()
}

func (index *Index) Ident() string {
	return index.Index.Ident()
}

func (index *Index) Type() core.Type {
	return index.Index.Type()
}

func (e *ExprGetElementPtr) gepExprType(elemType, src core.Type, indices []constant.Constant) core.Type {
	var idxs []*GepIndex
	for _, index := range indices {
		idx := GetGepIndex(index)
		idxs = append(idxs, idx)
	}
	return GepResultType(elemType, idxs)
}
