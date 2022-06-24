package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/ast/types"
	"github.com/panda-io/micro-panda/token"
)

type Invocation struct {
	ExpressionBase
	Function  ast.Expression
	Arguments []ast.Expression

	FunctionDefine *types.TypeFunction
}

func (i *Invocation) Validate(c ast.Context, expected core.Type) {
	i.Function.Validate(c, expected)
	i.Const = false
	t := i.Function.Type()
	if f, ok := t.(*types.TypeFunction); ok {
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
				i.Arguments = append([]ast.Expression{u}, i.Arguments...)
			} else if m, ok := i.Function.(*MemberAccess); ok {
				if types.IsStruct(m.Parent.Type()) {
					u := &Unary{
						Operator:   token.BitAnd,
						Expression: m.Parent,
					}
					u.SetPosition(m.Parent.GetPosition())
					i.Arguments = append([]ast.Expression{u}, i.Arguments...)
				} else if types.IsPointer(m.Parent.Type()) {
					i.Arguments = append([]ast.Expression{m.Parent}, i.Arguments...)
				}
			}
		}
		if len(i.Arguments) == 0 {
			if len(f.Parameters) > 0 {
				c.Error(i.GetPosition(), "expect arguments")
			}
		} else if len(f.Parameters) == len(i.Arguments) {
			for n := 0; n < len(f.Parameters); n++ {
				i.Arguments[n].Validate(c, f.Parameters[n])
			}
		} else {
			c.Error(i.GetPosition(), "mismatch arguments and parameters")
		}
	} else {
		c.Error(i.GetPosition(), "expect function type")
	}
}
