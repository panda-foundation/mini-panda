package statement

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
)

type For struct {
	StatementBase
	Initialization ast.Statement
	Condition      ast.Expression
	Post           ast.Statement
	Body           ast.Statement
}

func (f *For) Validate(c ast.Context) {
	ctx := c.NewContext()
	if f.Initialization != nil {
		f.Initialization.Validate(ctx)
	}
	if f.Condition != nil {
		conditionCtx := ctx.NewContext()
		f.Condition.Validate(conditionCtx, ast_types.TypeBool)
		if f.Condition.Type() != nil && !f.Condition.Type().Equal(ast_types.TypeBool) {
			c.Error(f.Condition.GetPosition(), "expect bool type condition")
		}
	}
	if f.Post != nil {
		postCtx := ctx.NewContext()
		f.Post.Validate(postCtx)
	}
	if f.Body != nil {
		bodyCtx := ctx.NewContext()
		f.Body.Validate(bodyCtx)
	}
}
