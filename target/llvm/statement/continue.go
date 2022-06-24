package llvm

import (
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/target/llvm"
)

func ContinueIR(c llvm.Context) {
	c.Block().AddInstruction(ir.NewBr(c.LoopBlock))
}
