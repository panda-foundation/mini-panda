package statement

import "github.com/panda-io/micro-panda/ast/core"

type Break struct {
	StatementBase
}

func (*Break) Validate(c core.Context) {
	/* if c.LeaveBlock == nil {
		//TO-DO add check
		//c.Program.Error(b.Position, "invalid break")
	}*/
}
