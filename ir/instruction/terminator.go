package instruction

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/ir/core"
)

type Terminator interface {
	isTerminator()
}

type TermRet struct {
	X core.Value
}

func NewRet(x core.Value) *TermRet {
	return &TermRet{X: x}
}

func (term *TermRet) WriteIR(w io.Writer) error {
	var err error
	if term.X == nil {
		_, err = w.Write([]byte("ret void"))
	} else {
		_, err = fmt.Fprintf(w, "ret %s", term.X)
	}
	return err
}

type TermBr struct {
	Target core.Value // *ir.Block
}

func NewBr(target core.Value) *TermBr {
	return &TermBr{Target: target}
}

func (term *TermBr) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "br %s", term.Target)
	return err
}

type TermCondBr struct {
	Cond        core.Value
	TargetTrue  core.Value // *ir.Block
	TargetFalse core.Value // *ir.Block
}

func NewCondBr(cond, targetTrue, targetFalse core.Value) *TermCondBr {
	return &TermCondBr{Cond: cond, TargetTrue: targetTrue, TargetFalse: targetFalse}
}

func (term *TermCondBr) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "br %s, %s, %s", term.Cond, term.TargetTrue, term.TargetFalse)
	return err
}

type TermSwitch struct {
	X             core.Value
	TargetDefault core.Value // *ir.Block
	Cases         []*Case
}

func NewSwitch(x, targetDefault core.Value, cases ...*Case) *TermSwitch {
	return &TermSwitch{X: x, TargetDefault: targetDefault, Cases: cases}
}

func (term *TermSwitch) WriteIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "switch %s, %s [\n", term.X, term.TargetDefault)
	if err != nil {
		return err
	}
	for _, c := range term.Cases {
		_, err = fmt.Fprintf(w, "\t\t%s\n", c)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("\t]"))
	return err
}

type Case struct {
	X      core.Value // Constant (integer constant or integer constant expression)
	Target core.Value // *ir.Block
}

func NewCase(x, target core.Value) *Case {
	return &Case{X: x, Target: target}
}

func (c *Case) String() string {
	return fmt.Sprintf("%s, %s", c.X, c.Target)
}

func (*TermRet) isTerminator()    {}
func (*TermBr) isTerminator()     {}
func (*TermCondBr) isTerminator() {}
func (*TermSwitch) isTerminator() {}
