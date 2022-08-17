package llvm

import "github.com/panda-io/micro-panda/target/llvm/ir/instruction"

func ContinueIR(c *Context) {
	c.Block.AddInstruction(instruction.NewBr(c.LoopBlock))
}
