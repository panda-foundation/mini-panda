package statement

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
)

type If struct {
	StatementBase
	Initialization ast.Statement
	Condition      ast.Expression
	Body           ast.Statement
	Else           ast.Statement
}

func (i *If) Validate(c ast.Context) {
	ctx := c.NewContext()
	if i.Initialization != nil {
		i.Initialization.Validate(ctx)
	}
	if i.Condition == nil {
		c.Error(i.GetPosition(), "expect condition expression")
	} else {
		i.Condition.Validate(ctx, ast_types.TypeBool)
		if i.Condition.Type() != nil && !i.Condition.Type().Equal(ast_types.TypeBool) {
			c.Error(i.Condition.GetPosition(), "expect bool type condition")
		}
	}
	if i.Body != nil {
		bodyCtx := ctx.NewContext()
		i.Body.Validate(bodyCtx)
	}
	if i.Else != nil {
		elseCtx := ctx.NewContext()
		i.Else.Validate(elseCtx)
	}
}
