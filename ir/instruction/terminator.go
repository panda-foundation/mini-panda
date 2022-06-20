package instruction

import (
	"fmt"
	"strings"
)

// === [ Terminators ] =========================================================

// Terminator is an LLVM IR terminator instruction (a control flow instruction).
//
// A Terminator has one of the following underlying
//
// Terminators
//
// https://llvm.org/docs/LangRef.html#terminator-instructions
//
//    *ir.TermRet           // https://godoc.org/github.com/llir/llvm/ir#TermRet
//    *ir.TermBr            // https://godoc.org/github.com/llir/llvm/ir#TermBr
//    *ir.TermCondBr        // https://godoc.org/github.com/llir/llvm/ir#TermCondBr
//    *ir.TermSwitch        // https://godoc.org/github.com/llir/llvm/ir#TermSwitch
//    *ir.TermIndirectBr    // https://godoc.org/github.com/llir/llvm/ir#TermIndirectBr
//    *ir.TermInvoke        // https://godoc.org/github.com/llir/llvm/ir#TermInvoke
//    *ir.TermCallBr        // https://godoc.org/github.com/llir/llvm/ir#TermCallBr
//    *ir.TermResume        // https://godoc.org/github.com/llir/llvm/ir#TermResume
//    *ir.TermCatchSwitch   // https://godoc.org/github.com/llir/llvm/ir#TermCatchSwitch
//    *ir.TermCatchRet      // https://godoc.org/github.com/llir/llvm/ir#TermCatchRet
//    *ir.TermCleanupRet    // https://godoc.org/github.com/llir/llvm/ir#TermCleanupRet
//    *ir.TermUnreachable   // https://godoc.org/github.com/llir/llvm/ir#TermUnreachable

type Terminator interface {
	isTerminator()
}

// --- [ ret ] -----------------------------------------------------------------

// TermRet is an LLVM IR ret terminator.
type TermRet struct {
	// Return value; or nil if void return.
	X Value
}

// NewRet returns a new ret terminator based on the given return  A nil
// return value indicates a void return.
func NewRet(x Value) *TermRet {
	return &TermRet{X: x}
}

// LLString returns the LLVM syntax representation of the terminator.
//
// Void return instruction.
//
//    'ret' XTyp=VoidType Metadata=(',' MetadataAttachment)+?
//
// Value return instruction.
//
//    'ret' XTyp=ConcreteType X=Value Metadata=(',' MetadataAttachment)+?
func (term *TermRet) LLString() string {
	buf := &strings.Builder{}
	if term.X == nil {
		buf.WriteString("ret void")
	} else {
		fmt.Fprintf(buf, "ret %s", term.X)
	}
	return buf.String()
}

// --- [ br ] ------------------------------------------------------------------

// TermBr is an unconditional LLVM IR br terminator.
type TermBr struct {
	// Target branch.
	Target Value // *ir.Block
}

// NewBr returns a new unconditional br terminator based on the given target
// basic block.
func NewBr(target Value) *TermBr {
	return &TermBr{Target: target}
}

// LLString returns the LLVM syntax representation of the terminator.
//
// 'br' Target=Label Metadata=(',' MetadataAttachment)+?
func (term *TermBr) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "br %s", term.Target)
	return buf.String()
}

// --- [ conditional br ] ------------------------------------------------------

// TermCondBr is a conditional LLVM IR br terminator.
type TermCondBr struct {
	// Branching condition.
	Cond Value
	// True condition target branch.
	TargetTrue Value // *ir.Block
	// False condition target branch.
	TargetFalse Value // *ir.Block
}

// NewCondBr returns a new conditional br terminator based on the given
// branching condition and conditional target basic blocks.
func NewCondBr(cond Value, targetTrue, targetFalse *Block) *TermCondBr {
	return &TermCondBr{Cond: cond, TargetTrue: targetTrue, TargetFalse: targetFalse}
}

// LLString returns the LLVM syntax representation of the terminator.
//
// 'br' CondTyp=IntType Cond=Value ',' TargetTrue=Label ',' TargetFalse=Label Metadata=(',' MetadataAttachment)+?
func (term *TermCondBr) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "br %s, %s, %s", term.Cond, term.TargetTrue, term.TargetFalse)
	return buf.String()
}

// --- [ switch ] --------------------------------------------------------------

// TermSwitch is an LLVM IR switch terminator.
type TermSwitch struct {
	// Control variable.
	X Value
	// Default target branch.
	TargetDefault Value // *ir.Block
	// Switch cases.
	Cases []*Case
}

// NewSwitch returns a new switch terminator based on the given control
// variable, default target basic block and switch cases.
func NewSwitch(x Value, targetDefault *Block, cases ...*Case) *TermSwitch {
	return &TermSwitch{X: x, TargetDefault: targetDefault, Cases: cases}
}

// LLString returns the LLVM syntax representation of the terminator.
//
// 'switch' X=TypeValue ',' Default=Label '[' Cases=Case* ']' Metadata=(',' MetadataAttachment)+?
func (term *TermSwitch) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "switch %s, %s [\n", term.X, term.TargetDefault)
	for _, c := range term.Cases {
		fmt.Fprintf(buf, "\t\t%s\n", c)
	}
	buf.WriteString("\t]")
	return buf.String()
}

// ~~~ [ Switch case ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Case is a switch case.
type Case struct {
	// Case comparand.
	X Value // Constant (integer constant or integer constant expression)
	// Case target branch.
	Target Value // *ir.Block
}

// NewCase returns a new switch case based on the given case comparand and
// target basic block.
func NewCase(x Constant, target *Block) *Case {
	return &Case{X: x, Target: target}
}

// String returns the string representation of the switch case.
func (c *Case) String() string {
	// X=TypeConst ',' Target=Label
	return fmt.Sprintf("%s, %s", c.X, c.Target)
}

// Terminator instructions
func (*TermRet) isTerminator()    {}
func (*TermBr) isTerminator()     {}
func (*TermCondBr) isTerminator() {}
func (*TermSwitch) isTerminator() {}
