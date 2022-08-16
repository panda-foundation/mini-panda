package expression

import (
	"fmt"

	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type ExprTrunc struct {
	From constant.Constant
	To   ir.Type
}

func NewExprTrunc(from constant.Constant, to ir.Type) *ExprTrunc {
	e := &ExprTrunc{From: from, To: to}
	e.Type()
	return e
}

func (e *ExprTrunc) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprTrunc) Type() ir.Type {
	return e.To
}

func (e *ExprTrunc) Ident() string {
	return fmt.Sprintf("trunc (%s to %s)", e.From.String(), e.To.String())
}

type ExprSExt struct {
	From constant.Constant
	To   ir.Type
}

func NewExprSExt(from constant.Constant, to ir.Type) *ExprSExt {
	e := &ExprSExt{From: from, To: to}
	e.Type()
	return e
}

func (e *ExprSExt) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprSExt) Type() ir.Type {
	return e.To
}

func (e *ExprSExt) Ident() string {
	return fmt.Sprintf("sext (%s to %s)", e.From.String(), e.To.String())
}

type ExprFPTrunc struct {
	From constant.Constant
	To   ir.Type
}

func NewExprFPTrunc(from constant.Constant, to ir.Type) *ExprFPTrunc {
	e := &ExprFPTrunc{From: from, To: to}
	e.Type()
	return e
}

func (e *ExprFPTrunc) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprFPTrunc) Type() ir.Type {
	return e.To
}

func (e *ExprFPTrunc) Ident() string {
	return fmt.Sprintf("fptrunc (%s to %s)", e.From.String(), e.To.String())
}

type ExprFPExt struct {
	From constant.Constant
	To   ir.Type
}

func NewExprFPExt(from constant.Constant, to ir.Type) *ExprFPExt {
	e := &ExprFPExt{From: from, To: to}
	e.Type()
	return e
}

func (e *ExprFPExt) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprFPExt) Type() ir.Type {
	return e.To
}

func (e *ExprFPExt) Ident() string {
	return fmt.Sprintf("fpext (%s to %s)", e.From.String(), e.To.String())
}

type ExprFPToUI struct {
	From constant.Constant
	To   ir.Type
}

func NewExprFPToUI(from constant.Constant, to ir.Type) *ExprFPToUI {
	e := &ExprFPToUI{From: from, To: to}
	e.Type()
	return e
}

func (e *ExprFPToUI) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprFPToUI) Type() ir.Type {
	return e.To
}

func (e *ExprFPToUI) Ident() string {
	return fmt.Sprintf("fptoui (%s to %s)", e.From.String(), e.To.String())
}

type ExprFPToSI struct {
	From constant.Constant
	To   ir.Type
}

func NewExprFPToSI(from constant.Constant, to ir.Type) *ExprFPToSI {
	e := &ExprFPToSI{From: from, To: to}
	e.Type()
	return e
}

func (e *ExprFPToSI) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprFPToSI) Type() ir.Type {
	return e.To
}

func (e *ExprFPToSI) Ident() string {
	return fmt.Sprintf("fptosi (%s to %s)", e.From.String(), e.To.String())
}

type ExprUIToFP struct {
	From constant.Constant
	To   ir.Type
}

func NewExprUIToFP(from constant.Constant, to ir.Type) *ExprUIToFP {
	e := &ExprUIToFP{From: from, To: to}
	e.Type()
	return e
}

func (e *ExprUIToFP) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprUIToFP) Type() ir.Type {
	return e.To
}

func (e *ExprUIToFP) Ident() string {
	return fmt.Sprintf("uitofp (%s to %s)", e.From.String(), e.To.String())
}

type ExprSIToFP struct {
	From constant.Constant
	To   ir.Type
}

func NewExprSIToFP(from constant.Constant, to ir.Type) *ExprSIToFP {
	e := &ExprSIToFP{From: from, To: to}
	e.Type()
	return e
}

func (e *ExprSIToFP) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprSIToFP) Type() ir.Type {
	return e.To
}

func (e *ExprSIToFP) Ident() string {
	return fmt.Sprintf("sitofp (%s to %s)", e.From.String(), e.To.String())
}

type ExprBitCast struct {
	From constant.Constant
	To   ir.Type
}

func NewExprBitCast(from constant.Constant, to ir.Type) *ExprBitCast {
	e := &ExprBitCast{From: from, To: to}
	e.Type()
	return e
}

func (e *ExprBitCast) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprBitCast) Type() ir.Type {
	return e.To
}

func (e *ExprBitCast) Ident() string {
	return fmt.Sprintf("bitcast (%s to %s)", e.From.String(), e.To.String())
}

type ExprPtrToInt struct {
	From constant.Constant
	To   ir.Type
}

func NewExprPtrToInt(from constant.Constant, to ir.Type) *ExprPtrToInt {
	e := &ExprPtrToInt{From: from, To: to}
	e.Type()
	return e
}

func (e *ExprPtrToInt) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprPtrToInt) Type() ir.Type {
	return e.To
}

func (e *ExprPtrToInt) Ident() string {
	return fmt.Sprintf("ptrtoint (%s to %s)", e.From.String(), e.To.String())
}

type ExprIntToPtr struct {
	From constant.Constant
	To   ir.Type
}

func NewExprIntToPtr(from constant.Constant, to ir.Type) *ExprIntToPtr {
	e := &ExprIntToPtr{From: from, To: to}
	e.Type()
	return e
}

func (e *ExprIntToPtr) String() string {
	return fmt.Sprintf("%s %s", e.Type().String(), e.Ident())
}

func (e *ExprIntToPtr) Type() ir.Type {
	return e.To
}

func (e *ExprIntToPtr) Ident() string {
	return fmt.Sprintf("inttoptr (%s to %s)", e.From.String(), e.To.String())
}
