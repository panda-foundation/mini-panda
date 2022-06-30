package statement

import "github.com/panda-io/micro-panda/ast/ast"

type Empty struct {
	StatementBase
}

func (*Empty) Validate(c ast.Context) {
}
