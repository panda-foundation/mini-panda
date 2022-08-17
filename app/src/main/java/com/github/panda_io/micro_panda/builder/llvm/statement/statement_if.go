package llvm

import (
	"github.com/panda-io/micro-panda/ast/statement"
	"github.com/panda-io/micro-panda/target/llvm/ir/instruction"
	ir_core "github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

func IfIR(c *Context, i *statement.If) {
	ctx := c.NewContext()
	ctx.Block = c.Block
	if i.Initialization != nil {
		StatementIR(ctx, i.Initialization)
	}

	nextBlock := c.Function.Function.NewBlock("")
	bodyBlock := c.Function.Function.NewBlock("")
	elseBlock := nextBlock

	bodyContext := ctx.NewContext()
	bodyContext.Block = bodyBlock
	StatementIR(bodyContext, i.Body)
	if bodyContext.Returned {
		ctx.Returned = true
	} else if !bodyContext.Block.Terminated {
		bodyContext.Block.AddInstruction(instruction.NewBr(nextBlock))
	}

	elseContext := ctx.NewContext()
	if i.Else == nil {
		ctx.Returned = false
	} else {
		elseBlock = c.Function.Function.NewBlock("")
		elseContext.Block = elseBlock
		StatementIR(elseContext, i.Else)
		ctx.Returned = elseContext.Returned
		if !elseContext.Block.Terminated {
			elseContext.Block.AddInstruction(instruction.NewBr(nextBlock))
		}
	}

	var condition ir_core.Value
	if i.Condition.IsConstant() {
		condition = ExpressionConstIR(ctx.Program, i.Condition)
	} else {
		condition = ExpressionIR(ctx, i.Condition)
	}
	ctx.Block.AddInstruction(instruction.NewCondBr(condition, bodyBlock, elseBlock))
	c.Block = nextBlock
	c.Returned = ctx.Returned
}
