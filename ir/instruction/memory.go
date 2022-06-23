package instruction

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/ir/constant"
	"github.com/panda-io/micro-panda/ir/constant/expression"
	"github.com/panda-io/micro-panda/ir/core"
	"github.com/panda-io/micro-panda/ir/types"
)

type InstAlloca struct {
	core.LocalIdent
	ElemType core.Type
	Typ      *types.PointerType
}

func NewAlloca(elemType core.Type) *InstAlloca {
	inst := &InstAlloca{ElemType: elemType}
	inst.Type()
	return inst
}

func (inst *InstAlloca) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstAlloca) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = types.NewPointerType(inst.ElemType)
	}
	return inst.Typ
}

func (inst *InstAlloca) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = alloca %s", inst.Ident(), inst.ElemType)
	return err
}

type InstLoad struct {
	core.LocalIdent
	ElemType core.Type
	Src      core.Value
}

func NewLoad(elemType core.Type, src core.Value) *InstLoad {
	inst := &InstLoad{ElemType: elemType, Src: src}
	return inst
}

func (inst *InstLoad) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstLoad) Type() core.Type {
	return inst.ElemType
}

func (inst *InstLoad) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = load %s, %s", inst.Ident(), inst.ElemType, inst.Src)
	return err
}

type InstStore struct {
	Src core.Value
	Dst core.Value
}

func NewStore(src, dst core.Value) *InstStore {
	dstPtrType, ok := dst.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid store dst operand type; expected *Pointer, got %T", dst.Type()))
	}
	if !src.Type().Equal(dstPtrType.ElemType) {
		panic(fmt.Errorf("store operands are not compatible: src=%v; dst=%v", src.Type(), dst.Type()))
	}
	return &InstStore{Src: src, Dst: dst}
}

func (inst *InstStore) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "store %s, %s", inst.Src, inst.Dst)
	return err
}

type InstGetElementPtr struct {
	core.LocalIdent
	ElemType core.Type
	Src      core.Value
	Indices  []core.Value
	Typ      core.Type // *PointerType or *VectorType (with elements of pointer type)
}

func NewGetElementPtr(elemType core.Type, src core.Value, indices ...core.Value) *InstGetElementPtr {
	inst := &InstGetElementPtr{ElemType: elemType, Src: src, Indices: indices}
	inst.Type()
	return inst
}

func (inst *InstGetElementPtr) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstGetElementPtr) Type() core.Type {
	if inst.Typ == nil {
		inst.Typ = inst.gepInstType(inst.ElemType, inst.Indices)
	}
	return inst.Typ
}

func (inst *InstGetElementPtr) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = getelementptr %s, %s", inst.Ident(), inst.ElemType, inst.Src)
	if err != nil {
		return err
	}
	for _, index := range inst.Indices {
		_, err := fmt.Fprintf(w, ", %s", index)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *InstGetElementPtr) gepInstType(elemType core.Type, indices []core.Value) core.Type {
	var idxs []*expression.GepIndex
	for _, index := range indices {
		var idx *expression.GepIndex
		switch index := index.(type) {
		case constant.Constant:
			idx = expression.GetGepIndex(index)
		}
		idxs = append(idxs, idx)
	}
	return expression.GepResultType(elemType, idxs)
}
