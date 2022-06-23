package statement

import (
	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/ast/types"
)

type For struct {
	StatementBase
	Initialization core.Statement
	Condition      core.Expression
	Post           core.Statement
	Body           core.Statement
}

func (f *For) Validate(c core.Context) {
	ctx := c.NewContext()
	if f.Initialization != nil {
		f.Initialization.Validate(ctx)
	}
	if f.Condition != nil {
		conditionCtx := ctx.NewContext()
		f.Condition.Validate(conditionCtx, types.TypeBool)
		if f.Condition.Type() != nil && !f.Condition.Type().Equal(types.TypeBool) {
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
