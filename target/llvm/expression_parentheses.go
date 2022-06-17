package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

func ParenthesesIR(c *Context, p *ast.Parentheses) ir.Value {
	return ExpressionIR(c, p.Expression)
}

func ParenthesesConstIR(program *Program, p *ast.Parentheses) ir.Constant {
	return ExpressionConstIR(program, p.Expression)
}
