package llvm

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/statement"
)

func StatementIR(c *Context, stmt ast.Statement) {
	switch s := stmt.(type) {
	case *statement.Block:
		BlockIR(c, s)
	case *statement.Break:
		BreakIR(c)
	case *statement.Continue:
		ContinueIR(c)
	case *statement.DeclarationStatement:
		DeclarationIR(c, s)
	case *statement.ExpressionStatement:
		ExpressionStatementIR(c, s)
	case *statement.For:
		ForIR(c, s)
	case *statement.If:
		IfIR(c, s)
	case *statement.Return:
		ReturnIR(c, s)
	case *statement.Switch:
		SwitchIR(c, s)
	}
}
