package expression

import "github.com/panda-io/micro-panda/ast/ast"

type This struct {
	ExpressionBase
}

func (t *This) Validate(c ast.Context, expected ast.Type) {
	t.Const = false
	t.Typ = c.FindObject(ast.StructThis)
	if t.Typ == nil {
		c.Error(t.GetPosition(), "undefined 'this'")
	}
}
