package llvm

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/ast/declaration"
	"github.com/panda-io/micro-panda/target/llvm/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
	ir_core "github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

type Function struct {
	DeclarationBase

	Parent *Struct

	Params       []*ir.Param
	ParamTypes   []ir_core.Type
	Function     *ir.Func
	FunctionType *ir_types.FuncType
	Entry        *ir.Block
	Body         *ir.Block
	Exit         *ir.Block
	Return       ir_core.Value
}

func (ff *Function) GenerateDefineIR(p *Program, f *declaration.Function) {
	if f.Parent != nil {
		this := ir_types.NewPointerType(ff.Parent.Struct)
		param := ir.NewParam(this)
		param.LocalName = ast.StructThis
		ff.Params = append(ff.Params, param)
	}
	if f.Parameters != nil {
		for _, parameter := range f.Parameters {
			t := TypeIR(parameter.Typ)
			if _, ok := parameter.Typ.(*ast_types.TypeName); ok {
				t = ir_types.NewPointerType(t)
			}
			param := ir.NewParam(t)
			param.LocalName = parameter.Name
			ff.Params = append(ff.Params, param)
			ff.ParamTypes = append(ff.ParamTypes, t)
		}
	}
	var t ir_core.Type = ir_types.Void
	if f.ReturnType != nil {
		t = TypeIR(f.ReturnType)
	}
	n := f.Qualified
	if f.Typ.Extern {
		n = f.Typ.ExternName
	}
	if n == ast.ProgramEntry {
		n = "main"
	}
	if f.Typ.TypeDefine {
		ff.FunctionType = ir_types.NewFuncType(f.Qualified, t, ff.ParamTypes...)
	} else {
		ff.Function = p.Program.NewFunc(n, t, ff.Params...)
	}
}

func (ff *Function) GenerateIR(p *Program, f *declaration.Function) {
	if f.Body != nil {
		c := NewContext(p)
		c.Function = ff
		ff.Entry = ff.Function.NewBlock(ast.FunctionEntry)
		ff.Body = ff.Function.NewBlock(ast.FunctionBody)
		ff.Exit = ff.Function.NewBlock(ast.FunctionExit)
		c.Block = ff.Entry

		// prepare params
		for _, param := range ff.Params {
			alloc := instruction.NewAlloca(param.Typ)
			ff.Entry.AddInstruction(alloc)
			store := instruction.NewStore(param, alloc)
			ff.Entry.AddInstruction(store)
			c.AddObject(param.LocalName, alloc)
		}

		// prepare return value
		if f.ReturnType != nil {
			alloca := instruction.NewAlloca(TypeIR(f.ReturnType))
			ff.Entry.AddInstruction(alloca)
			ff.Return = alloca
		}

		ff.Entry.AddInstruction(instruction.NewBr(ff.Body))
		c.Block = ff.Body
		StatementIR(c, f.Body)

		if !c.Block.Terminated {
			if c.Returned || f.ReturnType == nil {
				c.Block.AddInstruction(instruction.NewBr(ff.Exit))
			}
		}

		// return
		if f.ReturnType == nil {
			ff.Exit.AddInstruction(instruction.NewRet(nil))
		} else {
			load := instruction.NewLoad(TypeIR(f.ReturnType), ff.Return)
			ff.Exit.AddInstruction(load)
			ff.Exit.AddInstruction(instruction.NewRet(load))
		}
	}
}

func GenerateArgumentsIR(c *Context, args []ast.Expression, call *instruction.InstCall) {
	if args == nil {
		return
	}
	for _, arg := range args {
		v := ExpressionIR(c, arg)
		call.Args = append(call.Args, v)
	}
}
