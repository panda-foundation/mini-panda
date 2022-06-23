package expression

import (
	"fmt"

	"github.com/panda-io/micro-panda/ir/constant"
	"github.com/panda-io/micro-panda/ir/core"
)

// --- [ Unary expressions ] ---------------------------------------------------

type ExprFNeg struct {
	X   constant.Constant // floating-point scalar or vector constant
	Typ core.Type
}

func NewExprFNeg(x constant.Constant) *ExprFNeg {
	e := &ExprFNeg{X: x}
	e.Type()
	return e
}

func (e *ExprFNeg) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

func (e *ExprFNeg) Type() core.Type {
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

func (e *ExprFNeg) Ident() string {
	return fmt.Sprintf("fneg (%s)", e.X)
}

func (e *ExprFNeg) Simplify() constant.Constant {
	return e
}
