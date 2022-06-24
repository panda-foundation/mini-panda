package llvm

import (
	"github.com/panda-io/micro-panda/target/llvm/ir"
	ir_core "github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

type Context interface {
	NewContext() Context
	AddObject(name string, value ir_core.Value)
	FindObject(name string) ir_core.Value
	AutoLoad(value ir_core.Value) ir_core.Value

	Program() Program
	Block() *ir.Block
	LeaveBlock() *ir.Block
	LoopBlock() *ir.Block
}
