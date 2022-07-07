package com.github.panda_io.micro_panda.parser;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

import com.github.panda_io.micro_panda.ast.declaration.*;
import com.github.panda_io.micro_panda.ast.expression.*;
import com.github.panda_io.micro_panda.scanner.Token;

public class DeclarationParser {

	static Variable parseVariable(Context context, boolean isPublic, List<Declaration.Attribute> attributes)
			throws Exception {
		Variable variable = new Variable();
		variable.isPublic = isPublic;
		variable.attributes = attributes;
		if (context.scanner.token == Token.Const) {
			variable.constant = true;
		}
		context.scanner.scan();
		variable.name = context.parseIdentifier();
		variable.type = TypeParser.parseType(context);
		if (context.scanner.token == Token.Assign) {
			context.scanner.scan();
			variable.value = ExpressionParser.parseExpression(context);
		}
		if (variable.constant && variable.value == null) {
			context.addError(variable.name.getOffset(), "constant declaration must be initalized");
		}
		context.expect(Token.Semi);
		return variable;
	}

	static Function parseFunction(Context context, boolean isPublic, List<Declaration.Attribute> attributes)
			throws Exception {
		Function function = new Function();
		function.isPublic = isPublic;
		function.attributes = attributes;
		context.scanner.scan();
		function.name = context.parseIdentifier();
		function.parameters = TypeParser.parseParameters(context);
		if (context.scanner.token != Token.Semi && context.scanner.token != Token.LeftBrace) {
			function.returnType = TypeParser.parseType(context);
		}
		if (context.scanner.token == Token.LeftBrace) {
			function.body = StatementParser.parseBlockStatement(context);
		} else if (context.scanner.token == Token.Semi) {
			context.scanner.scan();
		}
		return function;
	}

	static Enumeration parseEnum(Context context, boolean isPublic, List<Declaration.Attribute> attributes)
			throws Exception {
		Enumeration enumeration = new Enumeration();
		enumeration.isPublic = isPublic;
		enumeration.attributes = attributes;
		context.scanner.scan();
		enumeration.name = context.parseIdentifier();
		context.expect(Token.LeftBrace);
		while (context.scanner.token != Token.RightBrace) {
			Variable variable = new Variable();
			variable.constant = true;
			variable.name = context.parseIdentifier();
			if (context.scanner.token == Token.Assign) {
				context.scanner.scan();
				variable.value = ExpressionParser.parseExpression(context);
			}
			boolean success = enumeration.addMember(variable);
			if (!success) {
				context.addError(variable.getOffset(),
						String.format("duplicated enum member '%s'", variable.name.name));
			}
			if (context.scanner.token != Token.Comma) {
				break;
			}
			context.scanner.scan();
		}
		context.expect(Token.RightBrace);
		return enumeration;
	}

	static Struct parseStruct(Context context, boolean isPublic, List<Declaration.Attribute> attributes)
			throws Exception {
		Struct struct = new Struct();
		struct.isPublic = isPublic;
		struct.attributes = attributes;
		context.scanner.scan();
		struct.name = context.parseIdentifier();
		context.expect(Token.LeftBrace);
		while (context.scanner.token != Token.RightBrace) {
			List<Declaration.Attribute> memberAttri = parseAttributes(context);
			boolean isPublicMember = parseModifier(context);
			switch (context.scanner.token) {
				case Const:
				case Var:
					Variable variable = parseVariable(context, isPublicMember, memberAttri);
					boolean success = struct.addVariable(variable);
					if (!success) {
						context.addError(variable.getOffset(),
								String.format("duplicated member '%s'", variable.name.name));
					}
					break;

				case Function:
					Function function = parseFunction(context, isPublicMember, memberAttri);
					success = struct.addFunction(function);
					if (!success) {
						context.addError(function.getOffset(),
								String.format("duplicated member '%s'", function.name.name));
					}
					break;

				default:
					context.expectedError(struct.getOffset(), "member declaration");
			}
		}
		context.expect(Token.RightBrace);
		return struct;
	}

	static boolean parseModifier(Context context) throws Exception {
		if (context.scanner.token == Token.Public) {
			context.scanner.scan();
			return true;
		}
		return false;
	}

	static List<Declaration.Attribute> parseAttributes(Context context) throws Exception {
		List<Declaration.Attribute> attributes = new ArrayList<>();
		if (context.scanner.token != Token.META) {
			return attributes;
		}
		while (context.scanner.token == Token.META) {
			context.scanner.scan();
			if (context.scanner.token != Token.IDENT) {
				context.expect(Token.IDENT);
			}
			Declaration.Attribute attribute = new Declaration.Attribute();
			attribute.values = new HashMap<>();
			attribute.setOffset(context.scanner.position);
			attribute.name = context.scanner.literal;
			context.scanner.scan();

			if (context.scanner.token == Token.STRING) {
				attribute.text = context.scanner.literal;
				context.scanner.scan();
			} else if (context.scanner.token == Token.LeftParen) {
				context.scanner.scan();
				if (context.scanner.token == Token.STRING) {
					attribute.text = context.scanner.literal;
					context.scanner.scan();
				} else {
					while (true) {
						if (context.scanner.token != Token.IDENT) {
							String name = context.scanner.literal;
							context.scanner.scan();
							context.expect(Token.Assign);
							switch (context.scanner.token) {
								case CHAR:
								case INT:
								case FLOAT:
								case STRING:
								case BOOL:
									if (attribute.values.containsKey(name)) {
										context.addError(context.scanner.position, "duplicated attribute " + name);
									}
									Literal literal = new Literal();
									literal.setOffset(context.scanner.position);
									literal.token = context.scanner.token;
									literal.value = context.scanner.literal;
									attribute.values.put(name, literal);
									break;

								default:
									context.expectedError(context.scanner.position,
											"basic literal (bool, char, int, float, string)");
							}
							context.scanner.scan();
							if (context.scanner.token == Token.RightParen) {
								break;
							}
							context.expect(Token.Comma);
						} else {
							context.expect(Token.IDENT);
						}
					}
				}
				context.expect(Token.RightParen);
			}
			attributes.add(attribute);
		}
		return attributes;
	}
}
