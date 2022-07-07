package com.github.panda_io.micro_panda.parser;

import java.util.List;
import java.util.ArrayList;

import com.github.panda_io.micro_panda.ast.Module;
import com.github.panda_io.micro_panda.ast.Module.Import;
import com.github.panda_io.micro_panda.ast.declaration.*;
import com.github.panda_io.micro_panda.ast.Constant;
import com.github.panda_io.micro_panda.scanner.*;

public class ModuleParser {

	static Module parseModule(Context context, File file) throws Exception {
		Module module = new Module();
		module.file = file;
		module.attributes = DeclarationParser.parseAttributes(context);
		module.namespace = parseNamespace(context);
		module.imports = parseImports(context);

		while (context.token != Token.EOF) {
			List<Declaration.Attribute> attributes = DeclarationParser.parseAttributes(context);
			boolean isPublic = DeclarationParser.parseModifier(context);
			switch (context.token) {
				case Const:
				case Var:
					Variable variable = DeclarationParser.parseVariable(context, isPublic, attributes);
					variable.qualified = String.format("%s.%s", module.namespace, variable.name.name);
					module.variables.add(variable);
					context.program.addDeclaration(variable);

				case Function:
					Function function = DeclarationParser.parseFunction(context, isPublic, attributes);
					function.qualified = String.format("%s.%s", module.namespace, function.name.name);
					module.functions.add(function);
					context.program.addDeclaration(function);

				case Enum:
					Enumeration enumeration = DeclarationParser.parseEnum(context, isPublic, attributes);
					enumeration.qualified = String.format("%s.%s", module.namespace, enumeration.name.name);
					module.enumerations.add(enumeration);
					context.program.addDeclaration(enumeration);

				case Struct:
					Struct struct = DeclarationParser.parseStruct(context, isPublic, attributes);
					struct.qualified = String.format("%s.%s", module.namespace, struct.name.name);
					module.structs.add(struct);
					context.program.addDeclaration(struct);

				default:
					context.expectedError(context.position, "declaration");
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
		String namespace = context.parseQualified();
		context.expect(Token.Semi);
		return namespace;
	}

	static List<Import> parseImports(Context context) throws Exception {
		List<Import> imports = new ArrayList<>();
		while (context.token == Token.Import) {
			context.expect(Token.Import);
			Import imp = new Import();
			imp.namespace = context.parseQualified();
			context.expect(Token.Semi);
			imports.add(imp);
		}
		return imports;
	}
}
