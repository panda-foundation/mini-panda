package statement

import (
	"fmt"

	"github.com/panda-io/micro-panda/ast/core"
)

type Return struct {
	StatementBase
	Expression core.Expression
}

func (r *Return) Validate(c core.Context) {
	returnType := c.GetFunction().GetReturnType()
	if returnType == nil {
		if r.Expression != nil {
			c.Error(r.GetPosition(), "mismatch return type, expect 'null'")
		}
	} else {
		if r.Expression == nil {
			c.Error(r.GetPosition(), fmt.Sprintf("mismatch return type, expect '%s' got 'null'", returnType.String()))
		}
		r.Expression.Validate(c, returnType)
		if r.Expression.Type() != nil && !r.Expression.Type().Equal(returnType) {
			c.Error(r.GetPosition(), fmt.Sprintf("mismatch return type, expect '%s' got '%s'", returnType.String(), r.Expression.Type().String()))
		}
	}
}
