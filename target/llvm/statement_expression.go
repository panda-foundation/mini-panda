package llvm

import "github.com/panda-io/micro-panda/ast"

func ExpressionStatementIR(c *Context, e *ast.ExpressionStatement) {
	if e.Expression.IsConstant() {
		ExpressionConstIR(c.Program, e.Expression)
	} else {
		ExpressionIR(c, e.Expression)
	}
}
