package com.github.panda_io.micro_panda.parser;

import java.util.List;
import java.util.ArrayList;

import com.github.panda_io.micro_panda.ast.declaration.Function.Parameter;
import com.github.panda_io.micro_panda.ast.expression.Expression;
import com.github.panda_io.micro_panda.ast.type.*;
import com.github.panda_io.micro_panda.scanner.Scanner;
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
				context.program.addError(context.scanner.position, "array count must > 0");
			}
			context.scanner.scan();
		}
		context.expect(Token.RightBracket);
		TypeArray array = new TypeArray(parseType(context));
		array.setOffset(offset);
		if (count > 0) {
			array.dimensions.add(count);
		}
		return array;
	}

	static TypeName parseTypeName(Context context) {
		/*
		t := &ast_types.TypeName{}
		t.SetPosition(p.position)
		qualified := p.parseQualified()
		if strings.Contains(qualified, ".") {
			t.Qualified = qualified
			names := strings.Split(qualified, ".")
			t.Name = names[len(names)-1]
		} else {
			t.Name = qualified
		}*/
		return null;
	}

	static TypePointer parseTypePointer(Context context) {
		/*
		t := &ast_types.TypePointer{}
		if p.token == token.Less {
			p.next()
			t.ElementType = p.parseType()
			p.expect(token.Greater)
		} else {
			t.ElementType = ast_types.TypeU8
		}*/
		return null;
	}

	static List<Parameter> parseParameters(Context context) {
		/*
		t := []*declaration.Parameter{}
		p.expect(token.LeftParen)
		if p.token == token.RightParen {
			p.next()
			return nil
		}
		t = append(t, p.parseParameter())
		for p.token == token.Comma {
			p.next()
			t = append(t, p.parseParameter())
		}
		p.expect(token.RightParen)*/
		return null;
	}

	static Parameter parseParameter(Context context) {
		/*
		t := &declaration.Parameter{}
		t.SetPosition(p.position)
		t.Name = p.parseIdentifier().Name
		t.Typ = p.parseType()*/
		return null;
	}

	static List<Expression> parseArguments(Context context) {
		/*
		expressions := []ast.Expression{}
		p.expect(token.LeftParen)
		if p.token == token.RightParen {
			p.next()
			return expressions
		}
		expressions = append(expressions, p.parseExpression())
		for p.token == token.Comma {
			p.next()
			expressions = append(expressions, p.parseExpression())
		}
		p.expect(token.RightParen)
		return expressions*/
		return null;
	}

	static TypeFunction parseFunctionType(Context context) throws Exception {
		/*
		t := &ast_types.TypeFunction{}
		t.SetPosition(p.position)
		p.expect(token.LeftParen)
		if p.token == token.RightParen {
			p.next()
			return t
		}
		t.Parameters = append(t.Parameters, p.parseType())
		for p.token == token.Comma {
			p.next()
			t.Parameters = append(t.Parameters, p.parseType())
		}
		p.expect(token.RightParen)
		if p.token != token.Semi && p.token != token.Assign {
			t.ReturnType = p.parseType()
		}*/
		return null;
	}
}
