package instruction

import (
	"fmt"
	"strings"
)

// --- [ Other instructions ] --------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstICmp is an LLVM IR icmp instruction.
type InstICmp struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Integer comparison predicate.
	Pred IPred
	// Integer scalar or vector operands.
	X, Y Value // integer scalar, pointer, integer vector or pointer vector.

	// extra.

	// Type of result produced by the instruction.
	Typ Type // boolean or boolean vector
}

// NewICmp returns a new icmp instruction based on the given integer comparison
// predicate and integer scalar or vector operands.
func NewICmp(pred IPred, x, y Value) *InstICmp {
	inst := &InstICmp{Pred: pred, X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstICmp) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction. The result type is either boolean
// type or vector of booleans type.
func (inst *InstICmp) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		switch xType := inst.X.Type().(type) {
		case *IntType, *PointerType:
			inst.Typ = I1
		case *VectorType:
			inst.Typ = NewVectorType(xType.Len, I1)
		default:
			panic(fmt.Errorf("invalid icmp operand type; expected *IntType, *PointerType or *VectorType, got %T", xType))
		}
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'icmp' Pred=IPred X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstICmp) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "icmp %s %s, %s", inst.Pred, inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFCmp is an LLVM IR fcmp instruction.
type InstFCmp struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Floating-point comparison predicate.
	Pred FPred
	// Floating-point scalar or vector operands.
	X, Y Value // floating-point scalar or floating-point vector

	// extra.

	// Type of result produced by the instruction.
	Typ Type // boolean or boolean vector
}

// NewFCmp returns a new fcmp instruction based on the given floating-point
// comparison predicate and floating-point scalar or vector operands.
func NewFCmp(pred FPred, x, y Value) *InstFCmp {
	inst := &InstFCmp{Pred: pred, X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFCmp) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction. The result type is either boolean
// type or vector of booleans type.
func (inst *InstFCmp) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		switch xType := inst.X.Type().(type) {
		case *FloatType:
			inst.Typ = I1
		case *VectorType:
			inst.Typ = NewVectorType(xType.Len, I1)
		default:
			panic(fmt.Errorf("invalid fcmp operand type; expected *FloatType or *VectorType, got %T", xType))
		}
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'fcmp' FastMathFlags=FastMathFlag* Pred=FPred X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstFCmp) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("fcmp")
	fmt.Fprintf(buf, " %s %s, %s", inst.Pred, inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ call ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstCall is an LLVM IR call instruction.
type InstCall struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Callee.
	Callee Value
	// Function arguments.
	Args []Value
	// Type of result produced by the instruction.
	Typ Type
}

// NewCall returns a new call instruction based on the given callee and function
// arguments.
//
// TODO: specify the set of underlying types of callee.
func NewCall(callee Value, args ...Value) *InstCall {
	inst := &InstCall{Callee: callee, Args: args}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstCall) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstCall) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		sig := inst.Sig()
		inst.Typ = sig.RetType
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// Tailopt 'call' FastMathFlags=FastMathFlag* CallingConvopt ReturnAttrs=ReturnAttribute* AddrSpaceopt Typ=Type Callee=Value '(' Args ')' FuncAttrs=FuncAttribute* OperandBundles=('[' (OperandBundle separator ',')+ ']')? Metadata=(',' MetadataAttachment)+?
func (inst *InstCall) LLString() string {
	buf := &strings.Builder{}
	if !inst.Type().Equal(Void) {
		fmt.Fprintf(buf, "%s = ", inst.Ident())
	}
	buf.WriteString("call")
	// Use function signature instead of return type for variadic functions.
	calleeType := inst.Type()
	if sig := inst.Sig(); sig.Variadic {
		calleeType = sig
	}
	fmt.Fprintf(buf, " %s %s(", calleeType, inst.Callee.Ident())
	for i, arg := range inst.Args {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	buf.WriteString(")")
	return buf.String()
}

// Sig returns the function signature of the callee.
func (inst *InstCall) Sig() *FuncType {
	t, ok := inst.Callee.Type().(*PointerType)
	if !ok {
		panic(fmt.Errorf("invalid callee type; expected *PointerType, got %T", inst.Callee.Type()))
	}
	sig, ok := t.ElemType.(*FuncType)
	if !ok {
		panic(fmt.Errorf("invalid callee type; expected *FuncType, got %T", t.ElemType))
	}
	return sig
}
