package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

func ThisIR(c *Context, this *ast.This) ir.Value {
	return c.FindObject(ast.StructThis)
}

func ThisConstIR(p *Program, this *ast.This) ir.Constant {
	return nil
}
