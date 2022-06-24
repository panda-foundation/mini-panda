package declaration

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/ast/statement"
)

type Parameter struct {
	ast.NodeBase
	Name string
	Typ  ast.Type
}

type Function struct {
	DeclarationBase
	Parameters []*Parameter
	ReturnType ast.Type
	Body       *statement.Block

	Parent *Struct
	Typ    *ast_types.TypeFunction
}

func (f *Function) IsConstant() bool {
	return true
}

func (f *Function) Kind() ast.DeclarationKind {
	return ast.DeclarationFunction
}

func (f *Function) Type() ast.Type {
	return f.Typ
}

func (f *Function) GetReturnType() ast.Type {
	return f.ReturnType
}

func (f *Function) ResolveType(c ast.Context) {
	f.Typ.ReturnType = f.ReturnType
	if f.HasAttribute(ast.AttriExtern) {
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
			f.Typ.Parameters = append(f.Typ.Parameters, param.Typ)
		}
	}
	c.ResolveType(f.Typ)
}

func (f *Function) Validate(ctx ast.Context) {
	if f.Body == nil {
		if f.Parent != nil {
			ctx.Error(f.GetPosition(), "function body is required for member function")
		}
		if f.Typ.Extern {
			if l := f.GetAttribute(ast.AttriExtern, "name"); l != nil {
				if n, ok := l.String(); ok {
					f.Typ.ExternName = n
				}
			}
			if f.Typ.ExternName == "" {
				ctx.Error(f.GetPosition(), "'name' of meta data is required for extern function")
			}
		}
	} else {
		c := ctx.NewContext()
		c.SetFunction(f)
		if f.Parent != nil {
			p := &ast_types.TypePointer{
				ElementType: f.Parent.Type(),
			}
			_ = c.AddObject(ast.StructThis, p)
		}
		if f.Typ.Extern {
			c.Error(f.GetPosition(), "extern function has no body")
		}
		if f.Parameters != nil {
			for _, param := range f.Parameters {
				err := c.AddObject(param.Name, param.Typ)
				if err != nil {
					c.Error(param.GetPosition(), err.Error())
				}
			}
		}
		f.Body.Validate(c)
	}
	//TO-DO check terminated
	//c.Program.Error(f.Position, "missing return")
}
