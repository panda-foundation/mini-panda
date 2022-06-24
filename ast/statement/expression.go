package statement

import (
	"github.com/panda-io/micro-panda/ast"
)

type ExpressionStatement struct {
	StatementBase
	Expression ast.Expression
}

func (e *ExpressionStatement) Validate(c ast.Context) {
	e.Expression.Validate(c, nil)
}
