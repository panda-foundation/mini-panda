package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

func ConversionIR(ctx *Context, c *ast.Conversion) ir.Value {
	if c.Typ.Equal(c.Value.Type()) {
		return ExpressionIR(ctx, c.Value)
	}
	v := ctx.AutoLoad(ExpressionIR(ctx, c.Value))
	if ast.IsNumber(c.Typ) {
		t0 := c.Typ.(*ast.TypeBuiltin)
		t1 := c.Value.Type().(*ast.TypeBuiltin)
		size0 := ast.TypeBuilinBits(t0)
		size1 := ast.TypeBuilinBits(t1)
		if ast.IsInteger(c.Typ) {
			if ast.IsInteger(c.Value.Type()) {
				// int to int
				if size0 < size1 {
					trunc := ir.NewTrunc(v, TypeIR(c.Typ))
					ctx.Block.AddInstruction(trunc)
					return trunc
				} else if size0 > size1 {
					sext := ir.NewSExt(v, TypeIR(c.Typ))
					ctx.Block.AddInstruction(sext)
					return sext
				} else {
					return v
				}
			} else {
				// float to int
				t := TypeIR(t0).(*ir.IntType)
				if t.Unsigned {
					f2ui := ir.NewFPToUI(v, TypeIR(c.Typ))
					ctx.Block.AddInstruction(f2ui)
					return f2ui
				} else {
					f2si := ir.NewFPToSI(v, TypeIR(c.Typ))
					ctx.Block.AddInstruction(f2si)
					return f2si
				}
			}
		} else {
			if ast.IsInteger(c.Value.Type()) {
				// int to float
				t := TypeIR(c.Value.Type()).(*ir.IntType)
				if t.Unsigned {
					ui2f := ir.NewUIToFP(v, TypeIR(c.Typ))
					ctx.Block.AddInstruction(ui2f)
					return ui2f
				} else {
					si2f := ir.NewSIToFP(v, TypeIR(c.Typ))
					ctx.Block.AddInstruction(si2f)
					return si2f
				}
			} else {
				// float to float
				if size0 < size1 {
					trunc := ir.NewFPTrunc(v, TypeIR(c.Typ))
					ctx.Block.AddInstruction(trunc)
					return trunc
				} else if size0 > size1 {
					sext := ir.NewFPExt(v, TypeIR(c.Typ))
					ctx.Block.AddInstruction(sext)
					return sext
				} else {
					return v
				}
			}
		}
	} else if ast.IsPointer(c.Typ) {
		// bit cast
		cast := ir.NewBitCast(v, TypeIR(c.Typ))
		ctx.Block.AddInstruction(cast)
		return cast
	}
	return nil
}

func ConversionConstIR(p *Program, c *ast.Conversion) ir.Constant {
	if c.Typ.Equal(c.Value.Type()) {
		return ExpressionConstIR(p, c.Value)
	}
	v := ExpressionConstIR(p, c.Value)
	if ast.IsNumber(c.Typ) {
		t0 := c.Typ.(*ast.TypeBuiltin)
		t1 := c.Value.Type().(*ast.TypeBuiltin)
		size0 := ast.TypeBuilinBits(t0)
		size1 := ast.TypeBuilinBits(t1)
		if ast.IsInteger(c.Typ) {
			if ast.IsInteger(c.Value.Type()) {
				// int to int
				if size0 < size1 {
					return ir.NewExprTrunc(v, TypeIR(c.Typ))
				} else if size0 > size1 {
					return ir.NewExprSExt(v, TypeIR(c.Typ))
				} else {
					return v
				}
			} else {
				// float to int
				t := TypeIR(t0).(*ir.IntType)
				if t.Unsigned {
					return ir.NewExprFPToUI(v, TypeIR(c.Typ))
				} else {
					return ir.NewExprFPToSI(v, TypeIR(c.Typ))
				}
			}
		} else {
			if ast.IsInteger(c.Value.Type()) {
				// int to float
				t := TypeIR(c.Value.Type()).(*ir.IntType)
				if t.Unsigned {
					return ir.NewExprUIToFP(v, TypeIR(c.Typ))
				} else {
					return ir.NewExprSIToFP(v, TypeIR(c.Typ))
				}
			} else {
				// float to float
				if size0 < size1 {
					return ir.NewExprFPTrunc(v, TypeIR(c.Typ))
				} else if size0 > size1 {
					return ir.NewExprFPExt(v, TypeIR(c.Typ))
				} else {
					return v
				}
			}
		}
	} else if ast.IsPointer(c.Typ) {
		// bit cast
		return ir.NewExprBitCast(v, TypeIR(c.Typ))
	}
	return nil
}
