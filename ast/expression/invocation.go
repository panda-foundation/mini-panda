package expression

import (
	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/token"
)

type Invocation struct {
	ExpressionBase
	Function  core.Expression
	Arguments *Arguments

	FunctionDefine *core.TypeFunction
}

func (i *Invocation) Validate(c core.Context, expected core.Type) {
	i.Function.Validate(c, expected)
	i.Const = false
	t := i.Function.Type()
	if f, ok := t.(*core.TypeFunction); ok {
		if f == nil {
			c.Error(i.Position, "invalid function")
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
				i.Arguments.Arguments = append([]core.Expression{u}, i.Arguments.Arguments...)
			} else if m, ok := i.Function.(*MemberAccess); ok {
				if core.IsStruct(m.Parent.Type()) {
					u := &Unary{
						Operator:   token.BitAnd,
						Expression: m.Parent,
					}
					u.Position = m.Parent.GetPosition()
					i.Arguments.Arguments = append([]core.Expression{u}, i.Arguments.Arguments...)
				} else if core.IsPointer(m.Parent.Type()) {
					i.Arguments.Arguments = append([]core.Expression{m.Parent}, i.Arguments.Arguments...)
				}
			}
		}
		if i.Arguments == nil {
			if len(f.Parameters) > 0 {
				c.Error(i.Position, "expect arguments")
			}
		} else if len(f.Parameters) == len(i.Arguments.Arguments) {
			for n := 0; n < len(f.Parameters); n++ {
				i.Arguments.Arguments[n].Validate(c, f.Parameters[n])
			}
		} else {
			c.Error(i.Position, "mismatch arguments and parameters")
		}
	} else {
		c.Error(i.Position, "expect function type")
	}
}
