package ast

type Break struct {
	StatementBase
}

func (*Break) Validate(c *Context) {
	/* if c.LeaveBlock == nil {
		//TO-DO add check
		//c.Program.Error(b.Position, "invalid break")
	}*/
}
