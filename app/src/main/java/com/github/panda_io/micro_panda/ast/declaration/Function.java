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
				parameter.type = context.resolveType(parameter.type);
				if (this.type.isDefine && parameter.type instanceof TypeFunction) {
					context.addError(parameter.getOffset(), "function pointer paramenter is not allowed in function define");
				}
				this.type.parameters.add(parameter.type);
			}
		}
		context.resolveType(this.type);
	}

	public void validate(Context context) {
		if (this.body == null) {
			if (this.parent != null) {
				context.addError(this.getOffset(), "function body is required for member function");
			}
			if (this.type.isExtern) {
				Literal literal = this.getAttribute(Constant.attriExtern, "name");
				if (literal != null && literal.token == Token.STRING) {
					if (literal.token == Token.STRING) {
						this.type.externName = literal.value.substring(1, literal.value.length() - 1);
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
			this.body.validate(ctx);
		}
		// TO-DO check terminated
		// c.Program.Error(f.Position, "missing return")
		if (this.qualified.equals(Constant.programEntry)) {
			if (this.parameters.size() > 0) {
				context.addError(this.parameters.get(0).getOffset(),
								"program entry has no parameters");
			}
			//TO-DO check return type, should be null or void
		}
		//TO-DO cannot return struct or array
	}
}
