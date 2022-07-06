package com.github.panda_io.micro_panda.parser;

import java.util.List;
import java.util.ArrayList;

import com.github.panda_io.micro_panda.ast.declaration.Function.Parameter;
import com.github.panda_io.micro_panda.ast.expression.Expression;
import com.github.panda_io.micro_panda.ast.type.*;
import com.github.panda_io.micro_panda.scanner.Token;

public class TypeParser {

	static Type parseType(Context context) throws Exception {
		if (context.scanner.token.isScalar()) {
			Type type = new TypeBuiltin(context.scanner.token);
			type.setOffset(context.scanner.position);
			context.scanner.scan();
			return type;
		}
		if (context.scanner.token == Token.Function) {
			context.scanner.scan();
			return parseFunctionType(context);
		}
		if (context.scanner.token == Token.LeftBracket) {
			return parseTypeArray(context);
		}
		if (context.scanner.token == Token.Pointer) {
			context.scanner.scan();
			return parseTypePointer(context);
		}
		return parseTypeName(context);
	}

	static TypeArray parseTypeArray(Context context) throws Exception {
		int offset = context.scanner.position;
		context.scanner.scan();
		int count = 0;
		if (context.scanner.token == Token.INT) {
			count = Integer.parseInt(context.scanner.literal);
			if (count < 1) {
				context.addError(context.scanner.position, "array count must > 0");
			}
			context.scanner.scan();
		}
		context.expect(Token.RightBracket);
		TypeArray array = new TypeArray();
		array.elementType = parseType(context);
		array.setOffset(offset);
		if (count > 0) {
			array.dimensions.add(count);
		}
		return array;
	}

	static TypeName parseTypeName(Context context) throws Exception {
		TypeName name = new TypeName();
		name.setOffset(context.scanner.position);
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
		if (context.scanner.token == Token.Less) {
			context.scanner.scan();
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
		if (context.scanner.token == Token.RightParen) {
			context.scanner.scan();
			return parameters;
		}
		parameters.add(parseParameter(context));
		while (context.scanner.token == Token.Comma) {
			context.scanner.scan();
			parameters.add(parseParameter(context));
		}
		context.expect(Token.RightParen);
		return parameters;
	}

	static Parameter parseParameter(Context context) throws Exception {
		Parameter parameter = new Parameter();
		parameter.setOffset(context.scanner.position);
		parameter.name = context.parseIdentifier().name;
		parameter.type = parseType(context);
		return parameter;
	}

	static List<Expression> parseArguments(Context context) throws Exception {
		List<Expression> expressions = new ArrayList<>();
		context.expect(Token.LeftParen);
		if (context.scanner.token == Token.RightParen) {
			context.scanner.scan();
			return expressions;
		}
		expressions.add(ExpressionParser.parseExpression(context));
		while (context.scanner.token == Token.Comma) {
			context.scanner.scan();
			expressions.add(ExpressionParser.parseExpression(context));
		}
		context.expect(Token.RightParen);
		return expressions;
	}

	static TypeFunction parseFunctionType(Context context) throws Exception {
		TypeFunction function = new TypeFunction();
		function.setOffset(context.scanner.position);
		context.expect(Token.LeftParen);
		if (context.scanner.token == Token.RightParen) {
			context.scanner.scan();
			return function;
		}
		function.parameters.add(parseType(context));
		while (context.scanner.token == Token.Comma) {
			context.scanner.scan();
			function.parameters.add(parseType(context));
		}
		context.expect(Token.RightParen);
		if (context.scanner.token != Token.Semi && context.scanner.token != Token.Assign) {
			function.returnType = parseType(context);
		}
		return function;
	}
}
