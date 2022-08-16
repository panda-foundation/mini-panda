package expression

import (
	"fmt"

	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

// --- [ Unary expressions ] ---------------------------------------------------

type ExprFNeg struct {
	X   constant.Constant // floating-point scalar or vector constant
	Typ ir.Type
}

func NewExprFNeg(x constant.Constant) *ExprFNeg {
	e := &ExprFNeg{X: x}
	e.Type()
	return e
}

func (e *ExprFNeg) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

func (e *ExprFNeg) Type() ir.Type {
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

func (e *ExprFNeg) Ident() string {
	return fmt.Sprintf("fneg (%s)", e.X.String())
}

func (e *ExprFNeg) Simplify() constant.Constant {
	return e
}
