package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/target/llvm/llvm"
	"github.com/panda-io/micro-panda/target/llvm/types"
)

func InvocationIR(c llvm.Context, i *ast.Invocation) ir.Value {
	t := i.Function.Type()
	if f, ok := t.(*ast.TypeFunction); ok {
		var call *ir.InstCall
		callee := ExpressionIR(c, i.Function)
		call = ir.NewCall(c.AutoLoad(callee))
		if i.Arguments != nil {
			for i, arg := range i.Arguments.Arguments {
				v := ExpressionIR(c, arg)
				// for variadic function call
				// since the prototype doesn’t specify types for optional arguments,
				// in a call to a variadic function the default argument promotions are performed on the optional argument values.
				// this means the objects of type char or short int (whether signed or not) are promoted to either int or unsigned int, as appropriate;
				// and that objects of type float are promoted to type double.
				// so, if the caller passes a char as an optional argument, it is promoted to an int, and the function can access it with va_arg (ap, int).
				if f.Variadic && i >= len(f.Parameters) {
					if arg.Type().Equal(ast.TypeF16) || arg.Type().Equal(ast.TypeF32) {
						double := ir.NewFPExt(c.AutoLoad(v), ir.Float64)
						c.Block().AddInstruction(double)
						v = double
					} else {
						argTyp := types.TypeIR(arg.Type())
						if intTyp, ok := argTyp.(*ir.IntType); ok {
							if intTyp.BitSize < 32 {
								i32 := ir.NewSExt(c.AutoLoad(v), ir.I32)
								c.Block().AddInstruction(i32)
								v = i32
							}
						}
					}
				}
				if t, ok := arg.Type().(*ast.TypeArray); ok && t.Dimension[0] != 0 {
					var gep ir.Instruction = ir.NewGetElementPtr(types.TypeIR(t), v, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, 0))
					c.Block().AddInstruction(gep)
					v = gep.(ir.Value)
				} else {
					v = c.AutoLoad(v)
				}
				call.Args = append(call.Args, v)
			}
		}
		c.Block().AddInstruction(call)
		return call
	}
	return nil
}

func InvocationConstIR(p llvm.Program, i *ast.Invocation) ir.Constant {
	return nil
}
