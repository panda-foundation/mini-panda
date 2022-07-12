package com.github.panda_io.micro_panda.parser;

import java.util.List;
import java.util.ArrayList;

import com.github.panda_io.micro_panda.ast.declaration.Function.Parameter;
import com.github.panda_io.micro_panda.ast.expression.Expression;
import com.github.panda_io.micro_panda.ast.type.*;
import com.github.panda_io.micro_panda.scanner.Token;

public class TypeParser {

	static Type parseType(Context context) throws Exception {
		if (context.token.isScalar()) {
			Type type = new TypeBuiltin(context.token);
			type.setOffset(context.position);
			context.next();
			return type;
		}
		if (context.token == Token.Function) {
			context.next();
			return parseFunctionType(context);
		}
		if (context.token == Token.LeftBracket) {
			return parseTypeArray(context);
		}
		if (context.token == Token.Pointer) {
			context.next();
			return parseTypePointer(context);
		}
		return parseTypeName(context);
	}

	static TypeArray parseTypeArray(Context context) throws Exception {
		int offset = context.position;
		context.next();
		int count = 0;
		if (context.token == Token.INT) {
			count = Integer.parseInt(context.literal);
			if (count < 1) {
				context.addError(context.position, "array count must > 0");
			}
			context.next();
		}
		context.expect(Token.RightBracket);
		TypeArray array = new TypeArray();
		array.elementType = parseType(context);
		array.setOffset(offset);
		array.dimensions.add(count);
		return array;
	}

	static TypeName parseTypeName(Context context) throws Exception {
		TypeName name = new TypeName();
		name.setOffset(context.position);
		String qualified = context.parseQualified();
		if (qualified.contains(".")) {
			name.qualified = qualified;
			String[] names = qualified.split(".");
			name.name = names[names.length - 1];
		} else {
			name.name = qualified;
		}
		return name;
	}

	static TypePointer parseTypePointer(Context context) throws Exception {
		TypePointer pointer = new TypePointer();
		if (context.token == Token.Less) {
			context.next();
			pointer.elementType = parseType(context);
			context.expect(Token.Greater);
		} else {
			pointer.elementType = Type.u8;
		}
		return pointer;
	}

	static List<Parameter> parseParameters(Context context) throws Exception {
		List<Parameter> parameters = new ArrayList<>();
		context.expect(Token.LeftParen);
		if (context.token == Token.RightParen) {
			context.next();
			return parameters;
		}
		parameters.add(parseParameter(context));
		while (context.token == Token.Comma) {
			context.next();
			parameters.add(parseParameter(context));
		}
		context.expect(Token.RightParen);
		return parameters;
	}

	static Parameter parseParameter(Context context) throws Exception {
		Parameter parameter = new Parameter();
		parameter.setOffset(context.position);
		parameter.name = context.parseIdentifier().name;
		parameter.type = parseType(context);
		return parameter;
	}

	static List<Expression> parseArguments(Context context) throws Exception {
		List<Expression> expressions = new ArrayList<>();
		context.expect(Token.LeftParen);
		if (context.token == Token.RightParen) {
			context.next();
			return expressions;
		}
		expressions.add(ExpressionParser.parseExpression(context));
		while (context.token == Token.Comma) {
			context.next();
			expressions.add(ExpressionParser.parseExpression(context));
		}
		context.expect(Token.RightParen);
		return expressions;
	}

	static TypeFunction parseFunctionType(Context context) throws Exception {
		TypeFunction function = new TypeFunction();
		function.setOffset(context.position);
		context.expect(Token.LeftParen);
		if (context.token == Token.RightParen) {
			context.next();
			return function;
		}
		function.parameters.add(parseType(context));
		while (context.token == Token.Comma) {
			context.next();
			function.parameters.add(parseType(context));
		}
		context.expect(Token.RightParen);
		if (context.token != Token.Semi && context.token != Token.Assign) {
			function.returnType = parseType(context);
		}
		return function;
	}
}
