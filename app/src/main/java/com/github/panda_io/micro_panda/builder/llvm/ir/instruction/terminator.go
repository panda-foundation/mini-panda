package instruction

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type Terminator interface {
	isTerminator()
}

type TermRet struct {
	X ir.Value
}

func NewRet(x ir.Value) *TermRet {
	return &TermRet{X: x}
}

func (term *TermRet) WriteIR(w io.Writer) error {
	var err error
	if term.X == nil {
		_, err = w.Write([]byte("ret void"))
	} else {
		_, err = fmt.Fprintf(w, "ret %s", term.X.String())
	}
	return err
}

type TermBr struct {
	Target ir.Value // *ir.Block
}

func NewBr(target ir.Value) *TermBr {
	return &TermBr{Target: target}
}

func (term *TermBr) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "br %s", term.Target.String())
	return err
}

type TermCondBr struct {
	Cond        ir.Value
	TargetTrue  ir.Value // *ir.Block
	TargetFalse ir.Value // *ir.Block
}

func NewCondBr(cond, targetTrue, targetFalse ir.Value) *TermCondBr {
	return &TermCondBr{Cond: cond, TargetTrue: targetTrue, TargetFalse: targetFalse}
}

func (term *TermCondBr) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "br %s, %s, %s", term.Cond.String(), term.TargetTrue.String(), term.TargetFalse.String())
	return err
}

type TermSwitch struct {
	X             ir.Value
	TargetDefault ir.Value // *ir.Block
	Cases         []*Case
}

func NewSwitch(x, targetDefault ir.Value, cases ...*Case) *TermSwitch {
	return &TermSwitch{X: x, TargetDefault: targetDefault, Cases: cases}
}

func (term *TermSwitch) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "switch %s, %s [\n", term.X.String(), term.TargetDefault.String())
	if err != nil {
		return err
	}
	for _, c := range term.Cases {
		_, err = fmt.Fprintf(w, "\t\t%s\n", c.String())
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("\t]"))
	return err
}

type Case struct {
	X      ir.Value // Constant (integer constant or integer constant expression)
	Target ir.Value // *ir.Block
}

func NewCase(x, target ir.Value) *Case {
	return &Case{X: x, Target: target}
}

func (c *Case) String() string {
	return fmt.Sprintf("%s, %s", c.X.String(), c.Target.String())
}

func (*TermRet) isTerminator()    {}
func (*TermBr) isTerminator()     {}
func (*TermCondBr) isTerminator() {}
func (*TermSwitch) isTerminator() {}
