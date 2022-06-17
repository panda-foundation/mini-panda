package ast

import "github.com/panda-io/micro-panda/token"

type Invocation struct {
	ExpressionBase
	Function  Expression
	Arguments *Arguments

	FunctionDefine *TypeFunction
}

func (i *Invocation) Validate(c *Context, expected Type) {
	i.Function.Validate(c, expected)
	i.Const = false
	t := i.Function.Type()
	if f, ok := t.(*TypeFunction); ok {
		if f == nil {
			c.Program.Error(i.Position, "invalid function")
			return
		}
		i.FunctionDefine = f
		i.Typ = f.ReturnType
		if f.MemberFunction {
			if i.Arguments == nil {
				i.Arguments = &Arguments{}
				i.Arguments.Position = i.Function.GetPosition()
			}
			// implicit conversion
			if _, ok := i.Function.(*This); ok {
				u := &Unary{
					Operator:   token.BitAnd,
					Expression: i.Function,
				}
				u.Position = i.Function.GetPosition()
				i.Arguments.Arguments = append([]Expression{u}, i.Arguments.Arguments...)
			} else if m, ok := i.Function.(*MemberAccess); ok {
				if IsStruct(m.Parent.Type()) {
					u := &Unary{
						Operator:   token.BitAnd,
						Expression: m.Parent,
					}
					u.Position = m.Parent.GetPosition()
					i.Arguments.Arguments = append([]Expression{u}, i.Arguments.Arguments...)
				} else if IsPointer(m.Parent.Type()) {
					i.Arguments.Arguments = append([]Expression{m.Parent}, i.Arguments.Arguments...)
				}
			}
		}
		if i.Arguments == nil {
			if len(f.Parameters) > 0 {
				c.Program.Error(i.Position, "expect arguments")
			}
		} else if len(f.Parameters) == len(i.Arguments.Arguments) {
			for n := 0; n < len(f.Parameters); n++ {
				i.Arguments.Arguments[n].Validate(c, f.Parameters[n])
			}
		} else if len(i.Arguments.Arguments) < len(f.Parameters) {
			c.Program.Error(i.Position, "mismatch arguments and parameters")
		} else if len(i.Arguments.Arguments) > len(f.Parameters) {
			if f.Variadic {
				for n := 0; n < len(f.Parameters); n++ {
					i.Arguments.Arguments[n].Validate(c, f.Parameters[n])
					if i.Arguments.Arguments[n].Type() != nil && !i.Arguments.Arguments[n].Type().Equal(f.Parameters[n]) {
						c.Program.Error(i.Arguments.Arguments[n].GetPosition(), "mismatch argument and parameter type")
					}
				}
				for n := len(f.Parameters); n < len(i.Arguments.Arguments); n++ {
					i.Arguments.Arguments[n].Validate(c, nil)
					if IsStruct(i.Arguments.Arguments[n].Type()) {
						c.Program.Error(i.Arguments.Arguments[n].GetPosition(), "struct is not allowed as argement, use pointer instead")
					}
				}
			} else {
				c.Program.Error(i.Position, "mismatch arguments and parameters")
			}
		}
	} else {
		c.Program.Error(i.Position, "expect function type")
	}
}
