package com.github.panda_io.micro_panda.parser;

import java.util.List;
import java.util.ArrayList;

import com.github.panda_io.micro_panda.ast.Module;
import com.github.panda_io.micro_panda.ast.Module.Using;
import com.github.panda_io.micro_panda.ast.declaration.*;
import com.github.panda_io.micro_panda.ast.type.TypeFunction;
import com.github.panda_io.micro_panda.ast.Constant;
import com.github.panda_io.micro_panda.scanner.*;

public class ModuleParser {

	static Module parseModule(Context context, File file) throws Exception {
		Module module = new Module();
		module.file = file;
		context.program.addModule(file.filename(), module);
		context.program.setModule(module);
		module.attributes = DeclarationParser.parseAttributes(context);
		module.namespace = parseNamespace(context);
		module.usings = parseUsings(context);

		while (context.token != Token.EOF) {
			List<Declaration.Attribute> attributes = DeclarationParser.parseAttributes(context);
			boolean isPublic = DeclarationParser.parseModifier(context);
			switch (context.token) {
				case Var:
					Variable variable = DeclarationParser.parseVariable(context, isPublic, attributes);
					variable.qualified = String.format("%s.%s", module.namespace, variable.name.name);
					module.variables.add(variable);
					context.program.addDeclaration(variable);
					break;

				case Function:
					Function function = DeclarationParser.parseFunction(context, isPublic, attributes);
					function.qualified = String.format("%s.%s", module.namespace, function.name.name);
					function.type = new TypeFunction();
                    function.type.qualified = function.qualified;
					module.functions.add(function);
					context.program.addDeclaration(function);
					break;

				case Enum:
					Enumeration enumeration = DeclarationParser.parseEnum(context, isPublic, attributes);
					enumeration.qualified = String.format("%s.%s", module.namespace, enumeration.name.name);
					module.enumerations.add(enumeration);
					context.program.addDeclaration(enumeration);
					break;

				case Struct:
					Struct struct = DeclarationParser.parseStruct(context, isPublic, attributes, module.namespace);
					struct.qualified = String.format("%s.%s", module.namespace, struct.name.name);
					module.structs.add(struct);
					context.program.addDeclaration(struct);
					break;

				default:
					context.unexpected(context.position, "declaration");
			}
		}
		return module;
	}

	static String parseNamespace(Context context) throws Exception {
		context.expect(Token.Namespace);
		if (context.token == Token.Semi) {
			context.next();
			return Constant.global;
		}
		int position = context.position;
		String namespace = context.parseQualified();
		context.expect(Token.Semi);
		if (namespace.equals(Constant.global)) {
			context.addError(position, "'global' is reserved namespace");
		}
		return namespace;
	}

	static List<Using> parseUsings(Context context) throws Exception {
		List<Using> imports = new ArrayList<>();
		while (context.token == Token.Using) {
			context.expect(Token.Using);
			Using imp = new Using();
			imp.namespace = context.parseQualified();
			context.expect(Token.Semi);
			imports.add(imp);
		}
		return imports;
	}
}
