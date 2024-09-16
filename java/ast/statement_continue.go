package ast

type Continue struct {
	StatementBase
}

func (*Continue) Validate(c *Context) {
	/* if c.LoopBlock == nil {
		//TO-DO add check
		//c.Program.Error(con.Position, "invalid continue")
	}*/
}
