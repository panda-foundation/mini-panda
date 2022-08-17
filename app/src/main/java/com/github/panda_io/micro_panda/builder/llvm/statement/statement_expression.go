package llvm

import (
	"github.com/panda-io/micro-panda/ast/statement"
)

func ExpressionStatementIR(c *Context, e *statement.ExpressionStatement) {
	if e.Expression.IsConstant() {
		ExpressionConstIR(c.Program, e.Expression)
	} else {
		ExpressionIR(c, e.Expression)
	}
}
