package llvm

import "github.com/panda-io/micro-panda/ast"

func StatementIR(c *Context, stmt ast.Statement) {
	switch s := stmt.(type) {
	case *ast.Block:
		BlockIR(c, s)
	case *ast.Break:
		BreakIR(c)
	case *ast.Continue:
		ContinueIR(c)
	case *ast.DeclarationStatement:
		DeclarationIR(c, s)
	case *ast.ExpressionStatement:
		ExpressionStatementIR(c, s)
	case *ast.For:
		ForIR(c, s)
	case *ast.If:
		IfIR(c, s)
	case *ast.Return:
		ReturnIR(c, s)
	case *ast.Switch:
		SwitchIR(c, s)
	}
}
