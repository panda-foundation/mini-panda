package llvm

import (
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

func ParenthesesIR(c *Context, p *expression.Parentheses) ir.Value {
	return ExpressionIR(c, p.Expression)
}

func ParenthesesConstIR(program *Program, p *expression.Parentheses) constant.Constant {
	return ExpressionConstIR(program, p.Expression)
}
