package com.github.panda_io.micro_panda.parser;

import java.util.ArrayList;

import com.github.panda_io.micro_panda.ast.statement.*;
import com.github.panda_io.micro_panda.scanner.Token;

public class StatementParser {

	static Statement parseStatement(Context context) throws Exception {
		switch (context.token) {
			case Break:
				BreakStatement breakStmt = new BreakStatement();
				breakStmt.setOffset(context.position);
				context.next();
				context.expect(Token.Semi);
				return breakStmt;

			case Continue:
				ContinueStatement continueStmt = new ContinueStatement();
				continueStmt.setOffset(context.position);
				context.next();
				context.expect(Token.Semi);
				return continueStmt;

			case Return:
				ReturnStatement returnStmt = new ReturnStatement();
				returnStmt.setOffset(context.position);
				context.next();
				if (context.token != Token.Semi) {
					returnStmt.expression = ExpressionParser.parseExpression(context);
				}
				context.expect(Token.Semi);
				return returnStmt;

			case LeftBrace:
				return parseBlockStatement(context);

			case If:
				return parseIfStatement(context);

			case Switch:
				return parseSwitchStatement(context);

			case For:
				return parseForStatement(context);

			default:
				return parseSimpleStatement(context, true);
		}
	}

	static Statement parseSimpleStatement(Context context, boolean consumeSemi) throws Exception {
		switch (context.token) {
			case Semi:
				EmptyStatement empty = new EmptyStatement();
				empty.setOffset(context.position);
				if (consumeSemi) {
					context.expect(Token.Semi);
				}
				return empty;

			case Var:
				return parseDeclarationStatement(context, consumeSemi);

			case IDENT:
			case This:
			case CHAR:
			case INT:
			case FLOAT:
			case STRING:
			case BOOL:
			case NULL:
			case Void:
			case LeftParen:
			case LeftBracket:
			case Plus:
			case Minus:
			case Not:
			case BitXor:
				ExpressionStatement expression = new ExpressionStatement();
				expression.setOffset(context.position);
				expression.expression = ExpressionParser.parseExpression(context);
				if (consumeSemi) {
					context.expect(Token.Semi);
				}
				return expression;

			default:
				context.expectedError(context.position, "statement");
				return null;
		}
	}

	static DeclarationStatement parseDeclarationStatement(Context context, boolean consumeSemi) throws Exception {
		DeclarationStatement declaration = new DeclarationStatement();
		declaration.setOffset(context.position);
		context.next();
		declaration.name = context.parseIdentifier();
		if (context.token != Token.Assign && context.token != Token.Semi
				&& context.token != Token.Colon) {
			declaration.type = TypeParser.parseType(context);
		}
		if (context.token == Token.Assign) {
			context.next();
			declaration.value = ExpressionParser.parseExpression(context);
		}
		if (consumeSemi) {
			context.expect(Token.Semi);
		}
		return declaration;
	}

	static BlockStatement parseBlockStatement(Context context) throws Exception {
		BlockStatement block = new BlockStatement();
		block.setOffset(context.position);
		context.next();
		block.statements = new ArrayList<>();
		while (context.token != Token.RightBrace) {
			block.statements.add(parseStatement(context));
		}
		context.next();
		return block;
	}

	static IfStatement parseIfStatement(Context context) throws Exception {
		IfStatement ifStmt = new IfStatement();
		ifStmt.setOffset(context.position);
		context.next();
		context.expect(Token.LeftParen);

		Statement first = parseSimpleStatement(context, false);
		if (context.token == Token.Semi) {
			context.next();
			ifStmt.initialization = first;
			Statement condiftion = parseSimpleStatement(context, false);
			if (condiftion instanceof ExpressionStatement) {
				ifStmt.condition = ((ExpressionStatement) condiftion).expression;
			} else {
				context.addError(condiftion.getOffset(), "if condition must be an expression");
			}
		} else if (first instanceof ExpressionStatement) {
			ifStmt.condition = ((ExpressionStatement) first).expression;
		} else {
			context.addError(first.getOffset(), "if condition must be an expression");
		}
		context.expect(Token.RightParen);
		ifStmt.body = parseStatement(context);
		if (context.token == Token.Else) {
			context.next();
			if (context.token == Token.If) {
				ifStmt.elseStatement = parseIfStatement(context);
			} else {
				ifStmt.elseStatement = parseStatement(context);
			}
		}
		return ifStmt;
	}

	static SwitchStatement parseSwitchStatement(Context context) throws Exception {
		SwitchStatement switchStmt = new SwitchStatement();
		switchStmt.cases = new ArrayList<>();
		switchStmt.setOffset(context.position);
		context.next();
		context.expect(Token.LeftParen);
		Statement first = parseSimpleStatement(context, false);
		Statement operand;
		if (context.token == Token.Semi) {
			context.next();
			switchStmt.initialization = first;
			operand = parseSimpleStatement(context, false);
		} else {
			operand = first;
		}
		if (operand instanceof ExpressionStatement) {
			switchStmt.operand = ((ExpressionStatement) operand).expression;
		} else {
			context.addError(operand.getOffset(), "expect expression");
		}
		context.expect(Token.RightParen);
		context.expect(Token.LeftBrace);
		while (context.token == Token.Case) {
			switchStmt.cases.add(parseCaseStatement(context));
		}
		if (context.token == Token.Default) {
			switchStmt.defaultCase = parseCaseStatement(context);
		}
		if (switchStmt.cases.size() == 0) {
			context.addError(switchStmt.getOffset(), "expect 'case'");
		}
		context.expect(Token.RightBrace);
		return switchStmt;
	}

	static SwitchStatement.Case parseCaseStatement(Context context) throws Exception {
		SwitchStatement.Case caseStmt = new SwitchStatement.Case();
		caseStmt.setOffset(context.position);
		caseStmt.token = context.token;
		if (caseStmt.token == Token.Case) {
			context.next();
			caseStmt.caseExpr = ExpressionParser.parseExpression(context);
		} else {
			context.expect(Token.Default);
		}
		context.expect(Token.Colon);
		caseStmt.body = parseStatement(context);
		return caseStmt;
	}

	// for {}
	// for (condition) {}
	// for (init; condition; post) {}
	static ForStatement parseForStatement(Context context) throws Exception {
		ForStatement forStmt = new ForStatement();
		forStmt.setOffset(context.position);
		context.next();
		if (context.token == Token.LeftParen) {
			context.next();
			Statement first = parseSimpleStatement(context, false);
			if (context.token == Token.RightParen) {
				context.next();
				if (first instanceof ExpressionStatement) {
					forStmt.condition = ((ExpressionStatement) first).expression;
				} else {
					context.addError(first.getOffset(), "expect expression");
				}
				forStmt.body = parseStatement(context);
				return forStmt;
			} else {
				context.expect(Token.Semi);
				forStmt.initialization = first;
				Statement second = parseSimpleStatement(context, false);
				if (second instanceof ExpressionStatement) {
					forStmt.condition = ((ExpressionStatement) second).expression;
				} else {
					context.addError(second.getOffset(), "expect expression");
				}
				context.expect(Token.Semi);
				forStmt.post = parseSimpleStatement(context, false);
				context.expect(Token.RightParen);
				forStmt.body = parseStatement(context);
				return forStmt;
			}
		} else {
			forStmt.body = parseStatement(context);
			return forStmt;
		}
	}
}
