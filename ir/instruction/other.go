package instruction

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/ir/core"
	"github.com/panda-io/micro-panda/ir/types"
)

type InstICmp struct {
	core.LocalIdent
	Pred core.IPred
	X, Y core.Value // integer scalar, pointer, integer vector or pointer vector.
	Typ  core.Type  // boolean or boolean vector
}

func NewICmp(pred core.IPred, x, y core.Value) *InstICmp {
	inst := &InstICmp{Pred: pred, X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstICmp) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstICmp) Type() core.Type {
	if inst.Typ == nil {
		switch xType := inst.X.Type().(type) {
		case *types.IntType, *types.PointerType:
			inst.Typ = types.I1
		default:
			panic(fmt.Errorf("invalid icmp operand type; expected *IntType or *PointerType, got %T", xType))
		}
	}
	return inst.Typ
}

func (inst *InstICmp) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = icmp %s %s, %s", inst.Ident(), inst.Pred, inst.X, inst.Y.Ident())
	return err
}

type InstFCmp struct {
	core.LocalIdent
	Pred core.FPred
	X, Y core.Value // floating-point scalar or floating-point vector
	Typ  core.Type  // boolean or boolean vector
}

func NewFCmp(pred core.FPred, x, y core.Value) *InstFCmp {
	inst := &InstFCmp{Pred: pred, X: x, Y: y}
	inst.Type()
	return inst
}

func (inst *InstFCmp) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstFCmp) Type() core.Type {
	if inst.Typ == nil {
		switch xType := inst.X.Type().(type) {
		case *types.FloatType:
			inst.Typ = types.I1
		default:
			panic(fmt.Errorf("invalid fcmp operand type; expected *FloatType, got %T", xType))
		}
	}
	return inst.Typ
}

func (inst *InstFCmp) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s = fcmp %s %s, %s", inst.Ident(), inst.Pred, inst.X, inst.Y.Ident())
	return err
}

type InstCall struct {
	core.LocalIdent
	Callee core.Value
	Args   []core.Value
	Typ    core.Type
}

func NewCall(callee core.Value, args ...core.Value) *InstCall {
	inst := &InstCall{Callee: callee, Args: args}
	inst.Type()
	return inst
}

func (inst *InstCall) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstCall) Type() core.Type {
	if inst.Typ == nil {
		sig := inst.Sig()
		inst.Typ = sig.RetType
	}
	return inst.Typ
}

func (inst *InstCall) WriteIR(w io.Writer) error {
	var err error
	if !inst.Type().Equal(types.Void) {
		_, err = fmt.Fprintf(w, "%s = ", inst.Ident())
		if err != nil {
			return err
		}
	}
	calleeType := inst.Type()
	_, err = fmt.Fprintf(w, "call %s %s(", calleeType, inst.Callee.Ident())
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

func (inst *InstCall) Sig() *types.FuncType {
	t, ok := inst.Callee.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid callee type; expected *PointerType, got %T", inst.Callee.Type()))
	}
	sig, ok := t.ElemType.(*types.FuncType)
	if !ok {
		panic(fmt.Errorf("invalid callee type; expected *FuncType, got %T", t.ElemType))
	}
	return sig
}
