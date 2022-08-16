package llvm

import (
	"github.com/panda-io/micro-panda/ast/statement"
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
	ir_core "github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

func ForIR(c *Context, f *statement.For) {
	ctx := c.NewContext()
	ctx.Block = c.Block
	if f.Initialization != nil {
		StatementIR(ctx, f.Initialization)
	}

	nextBlock := c.Function.Function.NewBlock("")
	ctx.LeaveBlock = nextBlock

	conditionBlock := c.Function.Function.NewBlock("")
	conditionContext := ctx.NewContext()
	conditionContext.Block = conditionBlock

	postBlock := c.Function.Function.NewBlock("")
	ctx.LoopBlock = postBlock
	postContext := ctx.NewContext()
	postContext.Block = postBlock
	if f.Post != nil {
		StatementIR(postContext, f.Post)
	}
	postContext.Block.AddInstruction(instruction.NewBr(conditionBlock))

	bodyBlock := c.Function.Function.NewBlock("")
	bodyContext := ctx.NewContext()
	bodyContext.Block = bodyBlock
	StatementIR(bodyContext, f.Body)
	if bodyContext.Returned {
		ctx.Returned = true
	} else if !bodyContext.Block.Terminated {
		bodyContext.Block.AddInstruction(instruction.NewBr(postBlock))
	}

	var condition ir_core.Value
	if f.Condition.IsConstant() {
		condition = ExpressionConstIR(conditionContext.Program, f.Condition)
	} else {
		condition = ExpressionIR(conditionContext, f.Condition)
	}
	conditionContext.Block.AddInstruction(instruction.NewCondBr(condition, bodyBlock, nextBlock))
	ctx.Block.AddInstruction(instruction.NewBr(conditionBlock))
	c.Block = nextBlock
	c.Returned = ctx.Returned
}
