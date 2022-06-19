package expression

import (
	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/token"
)

type Invocation struct {
	ExpressionBase
	Function  core.Expression
	Arguments []core.Expression

	FunctionDefine *core.TypeFunction
}

func (i *Invocation) Validate(c core.Context, expected core.Type) {
	i.Function.Validate(c, expected)
	i.Const = false
	t := i.Function.Type()
	if f, ok := t.(*core.TypeFunction); ok {
		i.FunctionDefine = f
		i.Typ = f.ReturnType
		if f.MemberFunction {
			// implicit conversion
			if _, ok := i.Function.(*This); ok {
				u := &Unary{
					Operator:   token.BitAnd,
					Expression: i.Function,
				}
				u.SetPosition(i.Function.GetPosition())
				i.Arguments = append([]core.Expression{u}, i.Arguments...)
			} else if m, ok := i.Function.(*MemberAccess); ok {
				if core.IsStruct(m.Parent.Type()) {
					u := &Unary{
						Operator:   token.BitAnd,
						Expression: m.Parent,
					}
					u.SetPosition(m.Parent.GetPosition())
					i.Arguments.Arguments = append([]core.Expression{u}, i.Arguments.Arguments...)
				} else if core.IsPointer(m.Parent.Type()) {
					i.Arguments.Arguments = append([]core.Expression{m.Parent}, i.Arguments.Arguments...)
				}
			}
		}
		if i.Arguments == nil {
			if len(f.Parameters) > 0 {
				c.Error(i.GetPosition(), "expect arguments")
			}
		} else if len(f.Parameters) == len(i.Arguments.Arguments) {
			for n := 0; n < len(f.Parameters); n++ {
				i.Arguments.Arguments[n].Validate(c, f.Parameters[n])
			}
		} else {
			c.Error(i.GetPosition(), "mismatch arguments and parameters")
		}
	} else {
		c.Error(i.GetPosition(), "expect function type")
	}
}
