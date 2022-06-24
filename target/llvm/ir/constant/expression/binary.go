package expression

import (
	"fmt"

	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

// --- [ Binary expressions ] --------------------------------------------------

type ExprAdd struct {
	X, Y constant.Constant // integer scalar or vector constants
	Typ  ir.Type
}

func NewExprAdd(x, y constant.Constant) *ExprAdd {
	e := &ExprAdd{X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprAdd) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprAdd) Type() ir.Type {
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

func (e *ExprAdd) Ident() string {
	return fmt.Sprintf("add (%s, %s)", e.X.String(), e.Y.String())
}

type ExprFAdd struct {
	X, Y constant.Constant // floating-point scalar or vector constants
	Typ  ir.Type
}

func NewExprFAdd(x, y constant.Constant) *ExprFAdd {
	e := &ExprFAdd{X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprFAdd) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprFAdd) Type() ir.Type {
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

func (e *ExprFAdd) Ident() string {
	return fmt.Sprintf("fadd (%s, %s)", e.X.String(), e.Y.String())
}

type ExprSub struct {
	X, Y constant.Constant // integer scalar or vector constants
	Typ  ir.Type
}

func NewExprSub(x, y constant.Constant) *ExprSub {
	e := &ExprSub{X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprSub) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprSub) Type() ir.Type {
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

func (e *ExprSub) Ident() string {
	return fmt.Sprintf("sub (%s, %s)", e.X.String(), e.Y.String())
}

type ExprFSub struct {
	X, Y constant.Constant // floating-point scalar or vector constants
	Typ  ir.Type
}

func NewExprFSub(x, y constant.Constant) *ExprFSub {
	e := &ExprFSub{X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprFSub) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprFSub) Type() ir.Type {
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

func (e *ExprFSub) Ident() string {
	return fmt.Sprintf("fsub (%s, %s)", e.X.String(), e.Y.String())
}

type ExprMul struct {
	X, Y constant.Constant // integer scalar or vector constants
	Typ  ir.Type
}

func NewExprMul(x, y constant.Constant) *ExprMul {
	e := &ExprMul{X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprMul) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprMul) Type() ir.Type {
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

func (e *ExprMul) Ident() string {
	return fmt.Sprintf("mul (%s, %s)", e.X.String(), e.Y.String())
}

type ExprFMul struct {
	X, Y constant.Constant // floating-point scalar or vector constants
	Typ  ir.Type
}

func NewExprFMul(x, y constant.Constant) *ExprFMul {
	e := &ExprFMul{X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprFMul) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprFMul) Type() ir.Type {
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

func (e *ExprFMul) Ident() string {
	return fmt.Sprintf("fmul (%s, %s)", e.X.String(), e.Y.String())
}

type ExprUDiv struct {
	X, Y constant.Constant // integer scalar or vector constants
	Typ  ir.Type
}

func NewExprUDiv(x, y constant.Constant) *ExprUDiv {
	e := &ExprUDiv{X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprUDiv) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprUDiv) Type() ir.Type {
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

func (e *ExprUDiv) Ident() string {
	return fmt.Sprintf("udiv (%s, %s)", e.X.String(), e.Y.String())
}

type ExprSDiv struct {
	X, Y constant.Constant // integer scalar or vector constants
	Typ  ir.Type
}

func NewExprSDiv(x, y constant.Constant) *ExprSDiv {
	e := &ExprSDiv{X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprSDiv) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprSDiv) Type() ir.Type {
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

func (e *ExprSDiv) Ident() string {
	return fmt.Sprintf("sdiv (%s, %s)", e.X.String(), e.Y.String())
}

type ExprFDiv struct {
	X, Y constant.Constant // floating-point scalar or vector constants
	Typ  ir.Type
}

func NewExprFDiv(x, y constant.Constant) *ExprFDiv {
	e := &ExprFDiv{X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprFDiv) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprFDiv) Type() ir.Type {
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

func (e *ExprFDiv) Ident() string {
	return fmt.Sprintf("fdiv (%s, %s)", e.X.String(), e.Y.String())
}

type ExprURem struct {
	X, Y constant.Constant // integer scalar or vector constants
	Typ  ir.Type
}

func NewExprURem(x, y constant.Constant) *ExprURem {
	e := &ExprURem{X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprURem) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprURem) Type() ir.Type {
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

func (e *ExprURem) Ident() string {
	return fmt.Sprintf("urem (%s, %s)", e.X.String(), e.Y.String())
}

type ExprSRem struct {
	X, Y constant.Constant // integer scalar or vector constants
	Typ  ir.Type
}

func NewExprSRem(x, y constant.Constant) *ExprSRem {
	e := &ExprSRem{X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprSRem) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprSRem) Type() ir.Type {
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

func (e *ExprSRem) Ident() string {
	return fmt.Sprintf("srem (%s, %s)", e.X.String(), e.Y.String())
}

type ExprFRem struct {
	X, Y constant.Constant // floating-point scalar or vector constants
	Typ  ir.Type
}

func NewExprFRem(x, y constant.Constant) *ExprFRem {
	e := &ExprFRem{X: x, Y: y}
	e.Type()
	return e
}

func (e *ExprFRem) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprFRem) Type() ir.Type {
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

func (e *ExprFRem) Ident() string {
	return fmt.Sprintf("frem (%s, %s)", e.X.String(), e.Y.String())
}
