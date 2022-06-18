package declaration

import (
	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/ast/statement"
)

type Function struct {
	DeclarationBase
	Parameters []*expression.Parameter
	ReturnType core.Type
	Body       *statement.Block

	Parent *Struct
	Typ    *core.TypeFunction
}

func (f *Function) IsConstant() bool {
	return true
}

func (f *Function) Kind() core.DeclarationKind {
	return core.DeclarationFunction
}

func (f *Function) Type() core.Type {
	return f.Typ
}

func (f *Function) ResolveType(c core.Context) {
	f.Typ.ReturnType = f.ReturnType
	if f.HasAttribute(core.Extern) {
		f.Typ.Extern = true
	} else if f.Body == nil {
		f.Typ.TypeDefine = true
	}
	if f.Parent != nil {
		f.Typ.MemberFunction = true
		f.Typ.Parameters = append(f.Typ.Parameters, f.Parent.PointerType())
	}
	if f.Parameters != nil {
		for _, param := range f.Parameters {
			f.Typ.Parameters = append(f.Typ.Parameters, param.Type)
		}
	}
	c.ResolveType(f.Typ)
}

func (f *Function) Validate(ctx core.Context) {
	if f.Body == nil {
		if f.Parent != nil {
			ctx.Error(f.Position, "function body is required for member function")
		}
		if f.Typ.Extern {
			if l := f.GetAttribute(core.Extern, "name"); l != nil {
				if n, ok := l.String(); ok {
					f.Typ.ExternName = n
				}
			}
			if f.Typ.ExternName == "" {
				ctx.Error(f.Position, "'name' of meta data is required for extern function")
			}
		}
	} else {
		c := ctx.NewContext()
		c.SetFunction(f)
		if f.Parent != nil {
			p := &core.TypePointer{
				ElementType: f.Parent.Type(),
			}
			p.Position = f.Parent.Position
			_ = c.AddObject(core.StructThis, p)
		}
		if f.Typ.Extern {
			c.Error(f.Position, "extern function has no body")
		}
		if f.Parameters != nil {
			for _, param := range f.Parameters {
				err := c.AddObject(param.Name, param.Type)
				if err != nil {
					c.Error(param.Position, err.Error())
				}
			}
		}
		f.Body.Validate(c)
	}
	//TO-DO check terminated
	//c.Program.Error(f.Position, "missing return")
}
