package statement

import (
	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/token"
)

type Switch struct {
	StatementBase
	Initialization core.Statement
	Operand        core.Expression
	Cases          []*Case
	Default        *Case
}

type Case struct {
	StatementBase
	Token token.Token
	Case  core.Expression
	Body  core.Statement
}

func (s *Switch) Validate(c core.Context) {
	ctx := c.NewContext()
	if s.Initialization != nil {
		s.Initialization.Validate(ctx)
	}
	var operandType core.Type
	if s.Operand == nil {
		c.Error(s.GetPosition(), "expect switch operand")
		return
	} else {
		s.Operand.Validate(ctx, nil)
		operandType = s.Operand.Type()
		if !core.IsInteger(operandType) {
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

func (c *Case) Validate(ctx core.Context, operandType core.Type) {
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
