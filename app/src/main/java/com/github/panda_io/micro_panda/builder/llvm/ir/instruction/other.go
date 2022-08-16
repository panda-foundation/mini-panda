package instruction

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

type InstICmp struct {
	ir.LocalIdent
	Pred ir.IPred
	X, Y ir.Value // integer scalar, pointer, integer vector or pointer vector.
	Typ  ir.Type  // boolean or boolean vector
}

func NewICmp(pred ir.IPred, x, y ir.Value) *InstICmp {
	inst := &InstICmp{Pred: pred, X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstICmp) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstICmp) Type() ir.Type {
	if inst.Typ == nil {
		switch xType := inst.X.Type().(type) {
		case *ir_types.IntType, *ir_types.PointerType:
			inst.Typ = ir_types.I1
		default:
			panic(fmt.Errorf("invalid icmp operand type; expected *IntType or *PointerType, got %T", xType))
		}
	}
	return inst.Typ
}

func (inst *InstICmp) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = icmp %s %s, %s", inst.Ident(), inst.Pred, inst.X.String(), inst.Y.Ident())
	return err
}

type InstFCmp struct {
	ir.LocalIdent
	Pred ir.FPred
	X, Y ir.Value // floating-point scalar or floating-point vector
	Typ  ir.Type  // boolean or boolean vector
}

func NewFCmp(pred ir.FPred, x, y ir.Value) *InstFCmp {
	inst := &InstFCmp{Pred: pred, X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstFCmp) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstFCmp) Type() ir.Type {
	if inst.Typ == nil {
		switch xType := inst.X.Type().(type) {
		case *ir_types.FloatType:
			inst.Typ = ir_types.I1
		default:
			panic(fmt.Errorf("invalid fcmp operand type; expected *FloatType, got %T", xType))
		}
	}
	return inst.Typ
}

func (inst *InstFCmp) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fcmp %s %s, %s", inst.Ident(), inst.Pred, inst.X.String(), inst.Y.Ident())
	return err
}

type InstCall struct {
	ir.LocalIdent
	Callee ir.Value
	Args   []ir.Value
	Typ    ir.Type
}

func NewCall(callee ir.Value, args ...ir.Value) *InstCall {
	inst := &InstCall{Callee: callee, Args: args}
	inst.Type()
	return inst
}

func (inst *InstCall) String() string {
	return fmt.Sprintf("%s %s", inst.Type().String(), inst.Ident())
}

func (inst *InstCall) Type() ir.Type {
	if inst.Typ == nil {
		sig := inst.Sig()
		inst.Typ = sig.RetType
	}
	return inst.Typ
}

func (inst *InstCall) WriteIR(w io.Writer) error {
	var err error
	if !inst.Type().Equal(ir_types.Void) {
		_, err = fmt.Fprintf(w, "%s = ", inst.Ident())
		if err != nil {
			return err
		}
	}
	calleeType := inst.Type()
	_, err = fmt.Fprintf(w, "call %s %s(", calleeType.String(), inst.Callee.Ident())
	if err != nil {
		return err
	}
	for i, arg := range inst.Args {
		if i != 0 {
			_, err = w.Write([]byte(", "))
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte(arg.String()))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte(")"))
	return err
}

func (inst *InstCall) Sig() *ir_types.FuncType {
	t, ok := inst.Callee.Type().(*ir_types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid callee type; expected *PointerType, got %T", inst.Callee.Type()))
	}
	sig, ok := t.ElemType.(*ir_types.FuncType)
	if !ok {
		panic(fmt.Errorf("invalid callee type; expected *FuncType, got %T", t.ElemType))
	}
	return sig
}
