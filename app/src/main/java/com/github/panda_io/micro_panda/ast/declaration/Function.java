package com.github.panda_io.micro_panda.ast.declaration;

import java.util.*;

import com.github.panda_io.micro_panda.ast.*;
import com.github.panda_io.micro_panda.ast.expression.Literal;
import com.github.panda_io.micro_panda.ast.type.*;
import com.github.panda_io.micro_panda.scanner.Token;
import com.github.panda_io.micro_panda.ast.statement.BlockStatement;

public class Function extends Declaration {
	public static class Parameter extends Node {
		public String name;
		public Type type;
	}

	public List<Parameter> parameters;
	public Type returnType;
	public BlockStatement body;
	public Struct parent;
	public TypeFunction type;

	public boolean isConstant() {
		return true;
	}

	public Type getType() {
		return this.type;
	}

	public void resolveType(Context context) {
		this.type = new TypeFunction();
		this.type.returnType = this.returnType;
		if (this.hasAttribute(Constant.attriExtern)) {
			this.type.isExtern = true;
		} else if (this.body == null) {
			this.type.isDefine = true;
		}

		if (this.parent != null) {
			this.type.isMemberFunction = true;
			this.type.parameters.add(this.parent.pointerType());
		}
		if (this.parameters.size() > 0) {
			for (Parameter parameter : this.parameters) {
				this.type.parameters.add(parameter.type);
			}
		}
		context.resolveType(this.type);
	}

	public void validate(Context context) {
		if (this.body == null) {
			if (this.parent == null) {
				context.addError(this.getOffset(), "function body is required for member function");
			}
			if (this.type.isExtern) {
				Literal literal = this.getAttribute(Constant.attriExtern, "name");
				if (literal != null && literal.token == Token.STRING) {
					if (literal.token == Token.STRING) {
						// this.type.externName = literal.value/string? TO-DO check regex name
						// f.Typ.ExternName = n
					}
				} else {
					context.addError(this.getOffset(), "'name' of meta data is required for extern function");
				}
			}
		} else {
			Context ctx = context.newContext();
			ctx.setFunction(this);
			if (this.parent != null) {
				TypePointer pointer = new TypePointer();
				pointer.elementType = this.parent.getType();
				ctx.insertObject(Constant.structThis, pointer);
			}
			if (this.type.isExtern) {
				ctx.addError(this.getOffset(), "extern function has no body");
			}
			if (this.parameters != null) {
				for (Parameter parameter : this.parameters) {
					boolean success = ctx.insertObject(parameter.name, parameter.type);
					if (!success) {
						ctx.addError(parameter.getOffset(),
								String.format("redeclared parameter with name %s", parameter.name));
					}
				}
			}
			this.body.validate(context);
		}
		// TO-DO check terminated
		// c.Program.Error(f.Position, "missing return")
	}
}
