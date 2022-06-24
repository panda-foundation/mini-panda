package expression

import (
	"fmt"
	"strings"

	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type ExprGetElementPtr struct {
	ElemType ir.Type
	Src      constant.Constant
	Indices  []constant.Constant // *Int, *Vector or *Index
	Typ      ir.Type             // *PointerType or *VectorType (with elements of pointer type)
}

func NewExprGetElementPtr(elemType ir.Type, src constant.Constant, indices ...constant.Constant) *ExprGetElementPtr {
	e := &ExprGetElementPtr{ElemType: elemType, Src: src, Indices: indices}
	e.Type()
	return e
}

func (e *ExprGetElementPtr) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprGetElementPtr) Type() ir.Type {
	if e.Typ == nil {
		e.Typ = e.gepExprType(e.ElemType, e.Src.Type(), e.Indices)
	}
	return e.Typ
}

func (e *ExprGetElementPtr) Ident() string {
	buf := &strings.Builder{}
	buf.WriteString("getelementptr")
	fmt.Fprintf(buf, " (%s, %s", e.ElemType.String(), e.Src.String())
	for _, index := range e.Indices {
		fmt.Fprintf(buf, ", %s", index.String())
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

func (index *Index) Type() ir.Type {
	return index.Index.Type()
}

func (e *ExprGetElementPtr) gepExprType(elemType, src ir.Type, indices []constant.Constant) ir.Type {
	var idxs []*GepIndex
	for _, index := range indices {
		idx := GetGepIndex(index)
		idxs = append(idxs, idx)
	}
	return GepResultType(elemType, idxs)
}
