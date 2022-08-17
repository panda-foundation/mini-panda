package llvm

import (
	"github.com/panda-io/micro-panda/ast/statement"
	"github.com/panda-io/micro-panda/target/llvm/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
	ir_core "github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

func SwitchIR(c *Context, s *statement.Switch) {
	ctx := c.NewContext()
	ctx.Block = c.Block
	ctx.Returned = true
	if s.Initialization != nil {
		StatementIR(ctx, s.Initialization)
	}
	var operand ir_core.Value
	if s.Operand.IsConstant() {
		operand = ExpressionConstIR(c.Program, s.Operand)
	} else {
		operand = c.AutoLoad(ExpressionIR(ctx, s.Operand))
	}

	nextBlock := c.Function.Function.NewBlock("")
	ctx.LeaveBlock = nextBlock

	defaultContext := ctx.NewContext()
	defaultBlock := c.Function.Function.NewBlock("")
	defaultContext.Block = defaultBlock
	if s.Default != nil {
		StatementIR(defaultContext, s.Default.Body)
	}
	if !defaultContext.Block.Terminated {
		defaultContext.Block.AddInstruction(instruction.NewBr(nextBlock))
	}
	if !defaultContext.Returned {
		ctx.Returned = false
	}

	var caseBlocks []*instruction.Case
	for _, cc := range s.Cases {
		caseContext := ctx.NewContext()
		caseBlock := c.Function.Function.NewBlock("")
		caseContext.Block = caseBlock
		StatementIR(caseContext, cc.Body)
		if !caseContext.Returned {
			ctx.Returned = false
		}
		caseBlocks = append(caseBlocks, instruction.NewCase(ExpressionConstIR(c.Program, cc.Case), caseBlock))
	}

	for i, cc := range caseBlocks {
		b := cc.Target.(*ir.Block)
		if !b.Terminated {
			if i == len(caseBlocks)-1 {
				// last one
				b.AddInstruction(instruction.NewBr(defaultBlock))
			} else {
				b.AddInstruction(instruction.NewBr(caseBlocks[i+1].Target))
			}
		}
	}

	ctx.Block.AddInstruction(instruction.NewSwitch(operand, defaultBlock, caseBlocks...))
	c.Block = nextBlock
	c.Returned = ctx.Returned
}
