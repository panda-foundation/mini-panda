package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/target/llvm"
)

func SwitchIR(c llvm.Context, s *ast.Switch) {
	ctx := c.NewContext()
	ctx.Block = c.Block
	ctx.Returned = true
	if s.Initialization != nil {
		StatementIR(ctx, s.Initialization)
	}
	var operand ir.Value
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
		defaultContext.Block().AddInstruction(ir.NewBr(nextBlock))
	}
	if !defaultContext.Returned {
		ctx.Returned = false
	}

	var caseBlocks []*ir.Case
	for _, cc := range s.Cases {
		caseContext := ctx.NewContext()
		caseBlock := c.Function.Function.NewBlock("")
		caseContext.Block = caseBlock
		StatementIR(caseContext, cc.Body)
		if !caseContext.Returned {
			ctx.Returned = false
		}
		caseBlocks = append(caseBlocks, ir.NewCase(ExpressionConstIR(c.Program, cc.Case), caseBlock))
	}

	for i, cc := range caseBlocks {
		b := cc.Target.(*ir.Block)
		if !b.Terminated {
			if i == len(caseBlocks)-1 {
				// last one
				b.AddInstruction(ir.NewBr(defaultBlock))
			} else {
				b.AddInstruction(ir.NewBr(caseBlocks[i+1].Target))
			}
		}
	}

	ctx.Block().AddInstruction(ir.NewSwitch(operand, defaultBlock, caseBlocks...))
	c.Block = nextBlock
	c.Returned = ctx.Returned
}
