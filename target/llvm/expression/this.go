package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/ir/constant"
)

func ThisIR(c *Context, this *ast.This) ir.Value {
	return c.FindObject(ast.StructThis)
}

func ThisConstIR(p *Program, this *ast.This) constant.Constant {
	return nil
}
