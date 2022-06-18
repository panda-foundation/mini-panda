package statement

import "github.com/panda-io/micro-panda/ast/core"

type Empty struct {
	StatementBase
}

func (*Empty) Validate(c core.Context) {
}
