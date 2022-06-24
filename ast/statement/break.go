package statement

import "github.com/panda-io/micro-panda/ast/ast"

type Break struct {
	StatementBase
}

func (*Break) Validate(c ast.Context) {
	/* if c.LeaveBlock == nil {
		//TO-DO add check
		//c.Program.Error(b.Position, "invalid break")
	}*/
}
