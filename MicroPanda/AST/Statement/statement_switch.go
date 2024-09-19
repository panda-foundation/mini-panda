package ast

import (
	"github.com/panda-io/micro-panda/token"
)

type Switch struct {
	StatementBase
	Initialization Statement
	Operand        Expression
	Cases          []*Case
	Default        *Case
}

type Case struct {
	StatementBase
	Token token.Token
	Case  Expression
	Body  Statement
}

func (s *Switch) Validate(c *Context) {
	ctx := c.NewContext()
	if s.Initialization != nil {
		s.Initialization.Validate(ctx)
	}
	var operandType Type
	if s.Operand == nil {
		c.Program.Error(s.Position, "expect switch operand")
		return
	} else {
		s.Operand.Validate(ctx, nil)
		operandType = s.Operand.Type()
		if !IsInteger(operandType) {
			c.Program.Error(s.Operand.GetPosition(), "expect integer operand")
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

func (c *Case) Validate(context *Context, operandType Type) {
	if c.Case == nil {
		context.Program.Error(c.Position, "expect case expression")
	} else {
		c.Case.Validate(context, operandType)
		if !c.Case.Type().Equal(operandType) {
			context.Program.Error(c.Position, "case operand type mismatch with switch operand")
		}
	}
	if c.Body != nil {
		c.Body.Validate(context)
	}
}
