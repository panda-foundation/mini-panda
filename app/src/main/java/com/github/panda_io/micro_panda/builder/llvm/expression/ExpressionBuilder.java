package com.github.panda_io.micro_panda.builder.llvm.expression;

import com.github.panda_io.micro_panda.ast.expression.Expression;
import com.github.panda_io.micro_panda.builder.llvm.Context;
import com.github.panda_io.micro_panda.builder.llvm.Program;
import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.constant.Constant;

public class ExpressionBuilder {
    public static Value buildExpression(Context context, Expression expression) {
        /*
        switch t := expr.(type) {
            case *expression.Binary:
                return BinaryIR(c, t)
            case *expression.Decrement:
                return DecrementIR(c, t)
            case *expression.Identifier:
                return IdentifierIR(c, t)
            case *expression.Increment:
                return IncrementIR(c, t)
            case *expression.Initializer:
                return InitializerIR(c, t)
            case *expression.Invocation:
                return InvocationIR(c, t)
            case *expression.Literal:
                return LiteralIR(c, t)
            case *expression.MemberAccess:
                return MemberAccessIR(c, t)
            case *expression.Parentheses:
                return ParenthesesIR(c, t)
            case *expression.Subscripting:
                return SubscriptingIR(c, t)
            case *expression.This:
                return ThisIR(c, t)
            case *expression.Unary:
                return UnaryIR(c, t)
            case *expression.Conversion:
                return ConversionIR(c, t)
            case *expression.Sizeof:
                return SizeofIR(c, t)
            }*/
        return null;
    }

    public static Constant buildConstantExpression(Program program, Expression expression) {
        /*
        switch t := expr.(type) {
	case *expression.Binary:
		return BinaryConstIR(p, t)
	case *expression.Decrement:
		return DecrementConstIR(p, t)
	case *expression.Identifier:
		return IdentifierConstIR(p, t)
	case *expression.Increment:
		return IncrementConstIR(p, t)
	case *expression.Initializer:
		return InitializerConstIR(p, t)
	case *expression.Invocation:
		return InvocationConstIR(p, t)
	case *expression.Literal:
		return LiteralConstIR(p, t)
	case *expression.MemberAccess:
		return MemberAccessConstIR(p, t)
	case *expression.Parentheses:
		return ParenthesesConstIR(p, t)
	case *expression.Subscripting:
		return SubscriptingConstIR(p, t)
	case *expression.This:
		return ThisConstIR(p, t)
	case *expression.Unary:
		return UnaryConstIR(p, t)
	case *expression.Conversion:
		return ConversionConstIR(p, t)
	case *expression.Sizeof:
		return SizeofConstIR(p, t)
	}*/
        return null;
    }
}
