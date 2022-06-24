package declaration

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

type Function struct {
	DeclarationBase

	Parent *Struct

	Params       []*ir.Param
	ParamTypes   []ir.Type
	Function     *ir.Func
	FunctionType *ir.FuncType
	Entry        *ir.Block
	Body         *ir.Block
	Exit         *ir.Block
	Return       ir.Value
}

func (ff *Function) GenerateDefineIR(p *Program, f *ast.Function) {
	if f.Parent != nil {
		this := ir.NewPointerType(ff.Parent.Struct)
		param := ir.NewParam(this)
		param.LocalName = ast.StructThis
		ff.Params = append(ff.Params, param)
	}
	if f.Parameters != nil {
		for _, parameter := range f.Parameters {
			t := TypeIR(parameter.Type)
			if _, ok := parameter.Type.(*ast.TypeName); ok {
				t = ir.NewPointerType(t)
			}
			param := ir.NewParam(t)
			param.LocalName = parameter.Name
			ff.Params = append(ff.Params, param)
			ff.ParamTypes = append(ff.ParamTypes, t)
		}
	}
	var t ir.Type = ir.Void
	if f.ReturnType != nil {
		t = TypeIR(f.ReturnType)
	}
	n := f.Qualified
	if f.Type.Extern {
		n = f.Type.ExternName
	}
	if n == ast.ProgramEntry {
		n = "main"
	}
	if f.Type.TypeDefine {
		ff.FunctionType = ir.NewFuncType(t, ff.ParamTypes...)
	} else {
		ff.Function = p.Module.NewFunc(n, t, ff.Params...)
		ff.Function.Sig.Variadic = f.Type.Variadic
	}
}

func (ff *Function) GenerateIR(p *Program, f *ast.Function) {
	if f.Body != nil {
		c := NewContext(p)
		c.Function = ff
		ff.Entry = ff.Function.NewBlock(ast.FunctionEntry)
		ff.Body = ff.Function.NewBlock(ast.FunctionBody)
		ff.Exit = ff.Function.NewBlock(ast.FunctionExit)
		c.Block = ff.Entry

		// prepare params
		for _, param := range ff.Params {
			alloc := ir.NewAlloca(param.Typ)
			ff.Entry.AddInstruction(alloc)
			store := ir.NewStore(param, alloc)
			ff.Entry.AddInstruction(store)
			c.AddObject(param.LocalName, alloc)
		}

		// prepare return value
		if f.ReturnType != nil {
			alloca := ir.NewAlloca(TypeIR(f.ReturnType))
			ff.Entry.AddInstruction(alloca)
			ff.Return = alloca
		}

		ff.Entry.AddInstruction(ir.NewBr(ff.Body))
		c.Block = ff.Body
		StatementIR(c, f.Body)

		if !c.Block.Terminated {
			if c.Returned || f.ReturnType == nil {
				c.Block.AddInstruction(ir.NewBr(ff.Exit))
			}
		}

		// return
		if f.ReturnType == nil {
			ff.Exit.AddInstruction(ir.NewRet(nil))
		} else {
			load := ir.NewLoad(TypeIR(f.ReturnType), ff.Return)
			ff.Exit.AddInstruction(load)
			ff.Exit.AddInstruction(ir.NewRet(load))
		}
	}
}

func GenerateArgumentsIR(c *Context, args *ast.Arguments, call *ir.InstCall) {
	function := call.Callee.Type().(*ir.PointerType).ElemType.(*ir.FuncType)
	if args == nil {
		return
	}
	for _, arg := range args.Arguments {
		v := ExpressionIR(c, arg)
		if !function.Variadic {
			call.Args = append(call.Args, v)
		} else {
			call.Args = append(call.Args, c.AutoLoad(v))
		}
	}
}
