package statement

import "github.com/panda-io/micro-panda/ast/ast"

type Continue struct {
	StatementBase
}

func (*Continue) Validate(c ast.Context) {
	/* if c.LoopBlock == nil {
		//TO-DO add check
		//c.Program.Error(con.Position, "invalid continue")
	}*/
}
