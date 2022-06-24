package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/target/llvm"
)

func ExpressionStatementIR(c llvm.Context, e *ast.ExpressionStatement) {
	if e.Expression.IsConstant() {
		ExpressionConstIR(c.Program, e.Expression)
	} else {
		ExpressionIR(c, e.Expression)
	}
}
