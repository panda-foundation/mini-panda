package expression

import (
	"fmt"

	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

type ExprICmp struct {
	Pred ir.IPred
	X, Y constant.Constant
	Typ  ir.Type
}

func NewExprICmp(pred ir.IPred, x, y constant.Constant) *ExprICmp {
	e := &ExprICmp{Pred: pred, X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprICmp) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprICmp) Type() ir.Type {
	if e.Typ == nil {
		switch xType := e.X.Type().(type) {
		case *ir_types.IntType, *ir_types.PointerType:
			e.Typ = ir_types.I1

		default:
			panic(fmt.Errorf("invalid icmp operand type; expected *IntType, *PointerType or *VectorType, got %T", xType))
		}
	}
	return e.Typ
}

func (e *ExprICmp) Ident() string {
	return fmt.Sprintf("icmp %s (%s, %s)", e.Pred, e.X.String(), e.Y.String())
}

type ExprFCmp struct {
	Pred ir.FPred
	X, Y constant.Constant
	Typ  ir.Type
}

func NewExprFCmp(pred ir.FPred, x, y constant.Constant) *ExprFCmp {
	e := &ExprFCmp{Pred: pred, X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprFCmp) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprFCmp) Type() ir.Type {
	if e.Typ == nil {
		switch xType := e.X.Type().(type) {
		case *ir_types.FloatType:
			e.Typ = ir_types.I1

		default:
			panic(fmt.Errorf("invalid fcmp operand type; expected *FloatType or *VectorType, got %T", xType))
		}
	}
	return e.Typ
}

func (e *ExprFCmp) Ident() string {
	return fmt.Sprintf("fcmp %s (%s, %s)", e.Pred, e.X.String(), e.Y.String())
}
