package llvm

import "github.com/panda-io/micro-panda/ir"

func BreakIR(c *Context) {
	c.Block.AddInstruction(ir.NewBr(c.LeaveBlock))
}
