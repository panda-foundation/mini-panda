package llvm

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

func ThisIR(c *Context, this *expression.This) ir.Value {
	return c.FindObject(ast.StructThis)
}

func ThisConstIR(p *Program, this *expression.This) constant.Constant {
	return nil
}
