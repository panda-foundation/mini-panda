package statement

import "github.com/panda-io/micro-panda/ast/core"

type ExpressionStatement struct {
	StatementBase
	Expression core.Expression
}

func (e *ExpressionStatement) Validate(c core.Context) {
	e.Expression.Validate(c, nil)
}
