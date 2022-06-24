package llvm

import (
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/target/llvm"
)

func BreakIR(c llvm.Context) {
	c.Block().AddInstruction(ir.NewBr(c.LeaveBlock))
}
