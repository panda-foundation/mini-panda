package llvm

import "github.com/panda-io/micro-panda/ir"

func ContinueIR(c *Context) {
	c.Block.AddInstruction(ir.NewBr(c.LoopBlock))
}
