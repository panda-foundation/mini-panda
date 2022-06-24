package expression

import (
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/llvm"
	"github.com/panda-io/micro-panda/target/llvm/types"
)

func ConversionIR(ctx llvm.Context, c *expression.Conversion) ir.Value {
	if c.Typ.Equal(c.Value.Type()) {
		return ExpressionIR(ctx, c.Value)
	}
	v := ctx.AutoLoad(ExpressionIR(ctx, c.Value))
	if ast_types.IsNumber(c.Typ) {
		t0 := c.Typ.(*ast_types.TypeBuiltin)
		t1 := c.Value.Type().(*ast_types.TypeBuiltin)
		size0 := ast.TypeBuilinBits(t0)
		size1 := ast.TypeBuilinBits(t1)
		if ast_types.IsInteger(c.Typ) {
			if ast_types.IsInteger(c.Value.Type()) {
				// int to int
				if size0 < size1 {
					trunc := instruction.NewTrunc(v, types.TypeIR(c.Typ))
					ctx.Block().AddInstruction(trunc)
					return trunc
				} else if size0 > size1 {
					sext := ir.NewSExt(v, types.TypeIR(c.Typ))
					ctx.Block().AddInstruction(sext)
					return sext
				} else {
					return v
				}
			} else {
				// float to int
				t := types.TypeIR(t0).(*ir.IntType)
				if t.Unsigned {
					f2ui := ir.NewFPToUI(v, types.TypeIR(c.Typ))
					ctx.Block().AddInstruction(f2ui)
					return f2ui
				} else {
					f2si := ir.NewFPToSI(v, types.TypeIR(c.Typ))
					ctx.Block().AddInstruction(f2si)
					return f2si
				}
			}
		} else {
			if ast_types.IsInteger(c.Value.Type()) {
				// int to float
				t := types.TypeIR(c.Value.Type()).(*ir.IntType)
				if t.Unsigned {
					ui2f := ir.NewUIToFP(v, types.TypeIR(c.Typ))
					ctx.Block().AddInstruction(ui2f)
					return ui2f
				} else {
					si2f := ir.NewSIToFP(v, types.TypeIR(c.Typ))
					ctx.Block().AddInstruction(si2f)
					return si2f
				}
			} else {
				// float to float
				if size0 < size1 {
					trunc := ir.NewFPTrunc(v, types.TypeIR(c.Typ))
					ctx.Block().AddInstruction(trunc)
					return trunc
				} else if size0 > size1 {
					sext := ir.NewFPExt(v, types.TypeIR(c.Typ))
					ctx.Block().AddInstruction(sext)
					return sext
				} else {
					return v
				}
			}
		}
	} else if ast_types.IsPointer(c.Typ) {
		// bit cast
		cast := ir.NewBitCast(v, types.TypeIR(c.Typ))
		ctx.Block().AddInstruction(cast)
		return cast
	}
	return nil
}

func ConversionConstIR(p llvm.Program, c *ast.Conversion) ir.Constant {
	if c.Typ.Equal(c.Value.Type()) {
		return ExpressionConstIR(p, c.Value)
	}
	v := ExpressionConstIR(p, c.Value)
	if ast_types.IsNumber(c.Typ) {
		t0 := c.Typ.(*ast.TypeBuiltin)
		t1 := c.Value.Type().(*ast.TypeBuiltin)
		size0 := ast.TypeBuilinBits(t0)
		size1 := ast.TypeBuilinBits(t1)
		if ast_types.IsInteger(c.Typ) {
			if ast_types.IsInteger(c.Value.Type()) {
				// int to int
				if size0 < size1 {
					return ir.NewExprTrunc(v, types.TypeIR(c.Typ))
				} else if size0 > size1 {
					return ir.NewExprSExt(v, types.TypeIR(c.Typ))
				} else {
					return v
				}
			} else {
				// float to int
				t := types.TypeIR(t0).(*ir.IntType)
				if t.Unsigned {
					return ir.NewExprFPToUI(v, types.TypeIR(c.Typ))
				} else {
					return ir.NewExprFPToSI(v, types.TypeIR(c.Typ))
				}
			}
		} else {
			if ast_types.IsInteger(c.Value.Type()) {
				// int to float
				t := types.TypeIR(c.Value.Type()).(*ir.IntType)
				if t.Unsigned {
					return ir.NewExprUIToFP(v, types.TypeIR(c.Typ))
				} else {
					return ir.NewExprSIToFP(v, types.TypeIR(c.Typ))
				}
			} else {
				// float to float
				if size0 < size1 {
					return ir.NewExprFPTrunc(v, types.TypeIR(c.Typ))
				} else if size0 > size1 {
					return ir.NewExprFPExt(v, types.TypeIR(c.Typ))
				} else {
					return v
				}
			}
		}
	} else if ast_types.IsPointer(c.Typ) {
		// bit cast
		return ir.NewExprBitCast(v, types.TypeIR(c.Typ))
	}
	return nil
}
