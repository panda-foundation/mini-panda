package expression

import (
	"fmt"

	"github.com/panda-io/micro-panda/ir/constant"
	"github.com/panda-io/micro-panda/ir/core"
	"github.com/panda-io/micro-panda/ir/types"
)

type ExprICmp struct {
	Pred core.IPred
	X, Y constant.Constant
	Typ  core.Type
}

func NewExprICmp(pred core.IPred, x, y constant.Constant) *ExprICmp {
	e := &ExprICmp{Pred: pred, X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprICmp) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

func (e *ExprICmp) Type() core.Type {
	if e.Typ == nil {
		switch xType := e.X.Type().(type) {
		case *types.IntType, *types.PointerType:
			e.Typ = types.I1

		default:
			panic(fmt.Errorf("invalid icmp operand type; expected *IntType, *PointerType or *VectorType, got %T", xType))
		}
	}
	return e.Typ
}

func (e *ExprICmp) Ident() string {
	return fmt.Sprintf("icmp %s (%s, %s)", e.Pred, e.X, e.Y)
}

type ExprFCmp struct {
	Pred core.FPred
	X, Y constant.Constant
	Typ  core.Type
}

func NewExprFCmp(pred core.FPred, x, y constant.Constant) *ExprFCmp {
	e := &ExprFCmp{Pred: pred, X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprFCmp) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

func (e *ExprFCmp) Type() core.Type {
	if e.Typ == nil {
		switch xType := e.X.Type().(type) {
		case *types.FloatType:
			e.Typ = types.I1

		default:
			panic(fmt.Errorf("invalid fcmp operand type; expected *FloatType or *VectorType, got %T", xType))
		}
	}
	return e.Typ
}

func (e *ExprFCmp) Ident() string {
	return fmt.Sprintf("fcmp %s (%s, %s)", e.Pred, e.X, e.Y)
}
