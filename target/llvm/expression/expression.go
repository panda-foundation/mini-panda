package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
	"github.com/panda-io/micro-panda/llvm/context"
	"github.com/panda-io/micro-panda/target/llvm/llvm"
)

func ExpressionIR(c context.Context, expression ast_ast.Expression) ir.Value {
	switch t := expression.(type) {
	case *expression.Binary:
		return BinaryIR(c, t)
	case *ast.Decrement:
		return DecrementIR(c, t)
	case *ast.Identifier:
		return IdentifierIR(c, t)
	case *ast.Increment:
		return IncrementIR(c, t)
	case *ast.Initializer:
		return InitializerIR(c, t)
	case *ast.Invocation:
		return InvocationIR(c, t)
	case *ast.Literal:
		return LiteralIR(c, t)
	case *ast.MemberAccess:
		return MemberAccessIR(c, t)
	case *ast.Parentheses:
		return ParenthesesIR(c, t)
	case *ast.Subscripting:
		return SubscriptingIR(c, t)
	case *ast.This:
		return ThisIR(c, t)
	case *ast.Unary:
		return UnaryIR(c, t)
	case *ast.Conversion:
		return ConversionIR(c, t)
	case *ast.Sizeof:
		return SizeofIR(c, t)
	}
	return nil
}

func ExpressionConstIR(p llvm.Program, expression ast_ast.Expression) ir.Constant {
	switch t := expression.(type) {
	case *ast.Binary:
		return BinaryConstIR(p, t)
	case *ast.Decrement:
		return DecrementConstIR(p, t)
	case *ast.Identifier:
		return IdentifierConstIR(p, t)
	case *ast.Increment:
		return IncrementConstIR(p, t)
	case *ast.Initializer:
		return InitializerConstIR(p, t)
	case *ast.Invocation:
		return InvocationConstIR(p, t)
	case *ast.Literal:
		return LiteralConstIR(p, t)
	case *ast.MemberAccess:
		return MemberAccessConstIR(p, t)
	case *ast.Parentheses:
		return ParenthesesConstIR(p, t)
	case *ast.Subscripting:
		return SubscriptingConstIR(p, t)
	case *ast.This:
		return ThisConstIR(p, t)
	case *ast.Unary:
		return UnaryConstIR(p, t)
	case *ast.Conversion:
		return ConversionConstIR(p, t)
	case *ast.Sizeof:
		return SizeofConstIR(p, t)
	}
	return nil
}
