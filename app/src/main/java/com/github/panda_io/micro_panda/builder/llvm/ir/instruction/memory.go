package instruction

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

type InstAlloca struct {
	ir.LocalIdent
	ElemType ir.Type
	Typ      *ir_types.PointerType
}

func NewAlloca(elemType ir.Type) *InstAlloca {
	inst := &InstAlloca{ElemType: elemType}
	inst.Type()
	return inst
}

func (inst *InstAlloca) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstAlloca) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = ir_types.NewPointerType(inst.ElemType)
	}
	return inst.Typ
}

func (inst *InstAlloca) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = alloca %s", inst.Ident(), inst.ElemType.String())
	return err
}

type InstLoad struct {
	ir.LocalIdent
	ElemType ir.Type
	Src      ir.Value
}

func NewLoad(elemType ir.Type, src ir.Value) *InstLoad {
	inst := &InstLoad{ElemType: elemType, Src: src}
	return inst
}

func (inst *InstLoad) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstLoad) Type() ir.Type {
	return inst.ElemType
}

func (inst *InstLoad) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = load %s, %s", inst.Ident(), inst.ElemType.String(), inst.Src.String())
	return err
}

type InstStore struct {
	Src ir.Value
	Dst ir.Value
}

func NewStore(src, dst ir.Value) *InstStore {
	dstPtrType, ok := dst.Type().(*ir_types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid store dst operand type; expected *Pointer, got %T", dst.Type()))
	}
	if !src.Type().Equal(dstPtrType.ElemType) {
		panic(fmt.Errorf("store operands are not compatible: src=%v; dst=%v", src.Type(), dst.Type()))
	}
	return &InstStore{Src: src, Dst: dst}
}

func (inst *InstStore) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "store %s, %s", inst.Src.String(), inst.Dst.String())
	return err
}

type InstGetElementPtr struct {
	ir.LocalIdent
	ElemType ir.Type
	Src      ir.Value
	Indices  []ir.Value
	Typ      ir.Type // *PointerType or *VectorType (with elements of pointer type)
}

func NewGetElementPtr(elemType ir.Type, src ir.Value, indices ...ir.Value) *InstGetElementPtr {
	inst := &InstGetElementPtr{ElemType: elemType, Src: src, Indices: indices}
	inst.Type()
	return inst
}

func (inst *InstGetElementPtr) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstGetElementPtr) Type() ir.Type {
	if inst.Typ == nil {
		inst.Typ = inst.gepInstType(inst.ElemType, inst.Indices)
	}
	return inst.Typ
}

func (inst *InstGetElementPtr) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = getelementptr %s, %s", inst.Ident(), inst.ElemType.String(), inst.Src.String())
	if err != nil {
		return err
	}
	for _, index := range inst.Indices {
		_, err := fmt.Fprintf(w, ", %s", index.String())
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *InstGetElementPtr) gepInstType(elemType ir.Type, indices []ir.Value) ir.Type {
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
