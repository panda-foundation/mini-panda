package ast

type Function struct {
	DeclarationBase
	Parameters []*Parameter
	ReturnType Type
	Body       *Block

	Parent *Struct
	Type   *TypeFunction
}

type Parameter struct {
	NodeBase
	Name string
	Type Type
}

type Arguments struct {
	NodeBase
	Arguments []Expression
}

func (f *Function) ValidateType(c *Context) {
	f.Type.ReturnType = f.ReturnType
	if f.HasAttribute(Extern) {
		f.Type.Extern = true
	} else if f.Body == nil {
		f.Type.TypeDefine = true
	}
	if l := f.GetAttributeValue(Extern, Variadic); l != nil {
		if variadic, ok := l.Bool(); ok {
			f.Type.Variadic = variadic
		}
	}
	if f.Parent != nil {
		f.Type.MemberFunction = true
		f.Type.Parameters = append(f.Type.Parameters, f.Parent.PointerType())
	}
	if f.Parameters != nil {
		for _, param := range f.Parameters {
			f.Type.Parameters = append(f.Type.Parameters, param.Type)
		}
	}
	ValidateType(f.Type, c.Program)
}

func (f *Function) Validate(ctx *Context) {
	if f.Body == nil {
		if f.Parent != nil {
			ctx.Program.Error(f.Position, "function body is required for member function")
		}
		if f.Type.Extern {
			if l := f.GetAttributeValue(Extern, Name); l != nil {
				if n, ok := l.String(); ok {
					f.Type.ExternName = n
				}
			}
			if f.Type.ExternName == "" {
				ctx.Program.Error(f.Position, "'name' of meta data is required for extern function")
			}
		}
	} else {
		c := ctx.NewContext()
		c.Function = f
		if f.Parent != nil {
			p := &TypePointer{
				ElementType: f.Parent.Type(),
			}
			p.Position = f.Parent.Position
			_ = c.AddObject(StructThis, p)
			if f.Type.Variadic {
				c.Program.Error(f.Position, "member function cannot be 'extern' or 'variadic'")
			}
		}
		if f.Type.Extern {
			c.Program.Error(f.Position, "extern function has no body")
		}
		if f.Parameters != nil {
			for _, param := range f.Parameters {
				err := c.AddObject(param.Name, param.Type)
				if err != nil {
					c.Program.Error(param.Position, err.Error())
				}
			}
		}
		f.Body.Validate(c)
	}
	//TO-DO check terminated
	//c.Program.Error(f.Position, "missing return")
}
