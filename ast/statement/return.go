package statement

import "github.com/panda-io/micro-panda/ast/core"

type Return struct {
	StatementBase
	Expression core.Expression
}

func (r *Return) Validate(c core.Context) {
	r.Expression.Validate(c, c.ReturnType())
	if r.Expression == nil {
		if c.ReturnType() != nil {
			c.Error(r.Position, "mismatch return type, expect 'null'")
		}
	} else if r.Expression.Type() != nil && !r.Expression.Type().Equal(c.ReturnType()) {
		c.Error(r.Position, "mismatch return type")
	}
}
