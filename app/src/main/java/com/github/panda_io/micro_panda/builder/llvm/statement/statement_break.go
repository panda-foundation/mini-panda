package llvm

import (
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
)

func BreakIR(c *Context) {
	c.Block.AddInstruction(instruction.NewBr(c.LeaveBlock))
}
