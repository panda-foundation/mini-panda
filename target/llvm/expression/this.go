package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/llvm"
)

func ThisIR(c llvm.Context, this *ast.This) ir.Value {
	return c.FindObject(ast.StructThis)
}

func ThisConstIR(p llvm.Program, this *ast.This) constant.Constant {
	return nil
}
