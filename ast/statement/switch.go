package statement

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
	"github.com/panda-io/micro-panda/token"
)

type Switch struct {
	StatementBase
	Initialization ast.Statement
	Operand        ast.Expression
	Cases          []*Case
	Default        *Case
}

type Case struct {
	StatementBase
	Token token.Token
	Case  ast.Expression
	Body  ast.Statement
}

func (s *Switch) Validate(c ast.Context) {
	ctx := c.NewContext()
	if s.Initialization != nil {
		s.Initialization.Validate(ctx)
	}
	var operandType ast.Type
	if s.Operand == nil {
		c.Error(s.GetPosition(), "expect switch operand")
		return
	} else {
		s.Operand.Validate(ctx, nil)
		operandType = s.Operand.Type()
		if !ast_types.IsInteger(operandType) {
			c.Error(s.Operand.GetPosition(), "expect integer operand")
			return
		}
	}
	for _, ca := range s.Cases {
		caseCtx := ctx.NewContext()
		ca.Validate(caseCtx, operandType)
	}
	if s.Default != nil {
		defaultCtx := ctx.NewContext()
		s.Default.Validate(defaultCtx, operandType)
	}
}

func (c *Case) Validate(ctx ast.Context, operandType ast.Type) {
	if c.Case == nil {
		ctx.Error(c.GetPosition(), "expect case expression")
	} else {
		c.Case.Validate(ctx, operandType)
		if !c.Case.Type().Equal(operandType) {
			ctx.Error(c.GetPosition(), "case operand type mismatch with switch operand")
		}
	}
	if c.Body != nil {
		c.Body.Validate(ctx)
	}
}
