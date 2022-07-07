package com.github.panda_io.micro_panda.parser;

import com.github.panda_io.micro_panda.ast.expression.*;
import com.github.panda_io.micro_panda.scanner.Token;

public class ExpressionParser {

	static Expression parseExpression(Context context) throws Exception {
		return parseBinaryExpression(context, 0);
	}

	static Expression parseOperand(Context context) throws Exception {
		switch (context.scanner.token) {
			case IDENT:
				return context.parseIdentifier();

			case Bool:
			case Int8:
			case Int16:
			case Int32:
			case Int64:
			case Uint8:
			case Uint16:
			case Uint32:
			case Uint64:
			case Float16:
			case Float32:
			case Float64:
			case Pointer:
			case LeftBracket:
				Conversion conversion = new Conversion();
				conversion.setOffset(context.scanner.position);
				conversion.setType(TypeParser.parseType(context));
				context.expect(Token.LeftParen);
				conversion.value = parseExpression(context);
				context.expect(Token.RightParen);
				return conversion;

			case CHAR:
			case INT:
			case FLOAT:
			case STRING:
			case BOOL:
			case NULL:
			case Void:
				Literal literal = new Literal();
				literal.setOffset(context.scanner.position);
				literal.token = context.scanner.token;
				literal.value = context.scanner.literal;
				context.scanner.scan();
				return literal;

			case LeftParen:
				Parentheses parentheses = new Parentheses();
				parentheses.setOffset(context.scanner.position);
				context.scanner.scan();
				parentheses.expression = parseExpression(context);
				context.expect(Token.RightParen);
				return parentheses;

			case LeftBrace:
				Initializer initializer = new Initializer();
				initializer.setOffset(context.scanner.position);
				context.scanner.scan();
				while (true) {
					initializer.expressions.add(parseExpression(context));
					if (context.scanner.token == Token.Comma) {
						context.scanner.scan();
					} else if (context.scanner.token == Token.RightBrace) {
						context.scanner.scan();
						break;
					} else {
						context.addError(context.scanner.position, "unexpected " + context.scanner.token.toString());
						return null;
					}
				}
				return initializer;

			case This:
				This thisExpr = new This();
				thisExpr.setOffset(context.scanner.position);
				context.scanner.scan();
				return thisExpr;

			default:
				context.addError(context.scanner.position, "unexpected " + context.scanner.token.toString());
				return null;
		}
	}

	static Expression parsePrimaryExpression(Context context) throws Exception {
		Expression x = parseOperand(context);
		while (true) {
			switch (context.scanner.token) {
				case Dot:
					MemberAccess memberAccess = new MemberAccess();
					memberAccess.setOffset(context.scanner.position);
					context.scanner.scan();
					memberAccess.parent = x;
					memberAccess.member = context.parseIdentifier();
					x = memberAccess;
					break;

				case LeftBracket:
					Subscripting subscripting = new Subscripting();
					subscripting.setOffset(context.scanner.position);
					subscripting.parent = x;
					while (context.scanner.token == Token.LeftBracket) {
						context.scanner.scan();
						subscripting.indexes.add(parseExpression(context));
						context.expect(Token.RightBracket);
					}
					x = subscripting;
					break;

				case LeftParen:
					Invocation invocation = new Invocation();
					invocation.setOffset(context.scanner.position);
					invocation.function = x;
					invocation.arguments = TypeParser.parseArguments(context);
					x = invocation;
					break;

				case PlusPlus:
					Increment increment = new Increment();
					increment.setOffset(context.scanner.position);
					increment.expression = x;
					context.scanner.scan();
					return increment;

				case MinusMinus:
					Decrement decrement = new Decrement();
					decrement.setOffset(context.scanner.position);
					decrement.expression = x;
					context.scanner.scan();
					return decrement;

				default:
					return x;
			}
		}
	}

	static Expression parseUnaryExpression(Context context) throws Exception {
		switch (context.scanner.token) {
			case Plus:
			case Minus:
			case Not:
			case Complement:
			case BitAnd:
			case Mul:
				Unary unary = new Unary();
				unary.setOffset(context.scanner.position);
				unary.operator = context.scanner.token;
				context.scanner.scan();
				unary.expression = parseUnaryExpression(context);
				return unary;

			default:
				return parsePrimaryExpression(context);
		}
	}

	static Expression parseBinaryExpression(Context context, int precedence) throws Exception {
		Expression x = parseUnaryExpression(context);
		while (true) {
			if (context.scanner.token == Token.Semi) {
				return x;
			}
			Token operator = context.scanner.token;
			int operatorPrecedenc = operator.precedence();
			if (operatorPrecedenc <= precedence) {
				return x;
			}
			context.scanner.scan();
			Expression y = parseBinaryExpression(context, operatorPrecedenc);
			Binary binary = new Binary();
			binary.left = x;
			binary.right = y;
			binary.operator = operator;
			x = binary;
		}
	}
}
