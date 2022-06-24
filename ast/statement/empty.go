package statement

import "github.com/panda-io/micro-panda/ast"

type Empty struct {
	StatementBase
}

func (*Empty) Validate(c ast.Context) {
}
