package com.github.panda_io.micro_panda.parser;

import java.util.ArrayList;

import com.github.panda_io.micro_panda.ast.expression.*;
import com.github.panda_io.micro_panda.scanner.Token;

public class ExpressionParser {

	static Expression parseExpression(Context context) throws Exception {
		return parseBinaryExpression(context, 0);
	}

	static Expression parseOperand(Context context) throws Exception {
		switch (context.token) {
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
				conversion.setOffset(context.position);
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
				literal.setOffset(context.position);
				literal.token = context.token;
				literal.value = context.literal;
				context.next();
				return literal;

			case LeftParen:
				Parentheses parentheses = new Parentheses();
				parentheses.setOffset(context.position);
				context.next();
				parentheses.expression = parseExpression(context);
				context.expect(Token.RightParen);
				return parentheses;

			case LeftBrace:
				Initializer initializer = new Initializer();
				initializer.expressions = new ArrayList<>();
				initializer.setOffset(context.position);
				context.next();
				while (true) {
					initializer.expressions.add(parseExpression(context));
					if (context.token == Token.Comma) {
						context.next();
					} else if (context.token == Token.RightBrace) {
						context.next();
						break;
					} else {
						context.addError(context.position, "unexpected " + context.token.toString());
						return null;
					}
				}
				return initializer;

			case This:
				This thisExpr = new This();
				thisExpr.setOffset(context.position);
				context.next();
				return thisExpr;

			default:
				context.addError(context.position, "unexpected " + context.token.toString());
				return null;
		}
	}

	static Expression parsePrimaryExpression(Context context) throws Exception {
		Expression x = parseOperand(context);
		while (true) {
			switch (context.token) {
				case Dot:
					MemberAccess memberAccess = new MemberAccess();
					memberAccess.setOffset(context.position);
					context.next();
					memberAccess.parent = x;
					memberAccess.member = context.parseIdentifier();
					x = memberAccess;
					break;

				case LeftBracket:
					Subscripting subscripting = new Subscripting();
					subscripting.indexes = new ArrayList<>();
					subscripting.setOffset(context.position);
					subscripting.parent = x;
					while (context.token == Token.LeftBracket) {
						context.next();
						subscripting.indexes.add(parseExpression(context));
						context.expect(Token.RightBracket);
					}
					x = subscripting;
					break;

				case LeftParen:
					Invocation invocation = new Invocation();
					invocation.setOffset(context.position);
					invocation.function = x;
					invocation.arguments = TypeParser.parseArguments(context);
					x = invocation;
					break;

				case PlusPlus:
					Increment increment = new Increment();
					increment.setOffset(context.position);
					increment.expression = x;
					context.next();
					return increment;

				case MinusMinus:
					Decrement decrement = new Decrement();
					decrement.setOffset(context.position);
					decrement.expression = x;
					context.next();
					return decrement;

				default:
					return x;
			}
		}
	}

	static Expression parseUnaryExpression(Context context) throws Exception {
		switch (context.token) {
			case Plus:
			case Minus:
			case Not:
			case Complement:
			case BitAnd:
			case Mul:
				Unary unary = new Unary();
				unary.setOffset(context.position);
				unary.operator = context.token;
				context.next();
				unary.expression = parseUnaryExpression(context);
				return unary;

			default:
				return parsePrimaryExpression(context);
		}
	}

	static Expression parseBinaryExpression(Context context, int precedence) throws Exception {
		Expression x = parseUnaryExpression(context);
		while (true) {
			if (context.token == Token.Semi) {
				return x;
			}
			Token operator = context.token;
			int offset = context.position;
			int operatorPrecedenc = operator.precedence();
			if (operatorPrecedenc <= precedence) {
				return x;
			}
			context.next();
			Expression y = parseBinaryExpression(context, operatorPrecedenc);
			Binary binary = new Binary();
			binary.setOffset(offset);
			binary.left = x;
			binary.right = y;
			binary.operator = operator;
			x = binary;
		}
	}
}
