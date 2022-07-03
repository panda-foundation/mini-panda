package com.github.panda_io.micro_panda.ast.declaration;

import java.util.*;

import com.github.panda_io.micro_panda.ast.*;
import com.github.panda_io.micro_panda.ast.expression.Literal;
import com.github.panda_io.micro_panda.ast.type.*;
import com.github.panda_io.micro_panda.scanner.Token;
import com.github.panda_io.micro_panda.ast.statement.Block;

public class Function extends Declaration {
	public static class Parameter extends Node {
		public String name;
		public Type type;
	}

	public List<Parameter> parameters;
	public Type returnType;
	public Block body;
	public Struct parent;
	public com.github.panda_io.micro_panda.ast.type.Function type;

	public boolean isConstant() {
		return true;
	}

	public Type getType() {
		return this.type;
	}

    public void resolveType(Context context) {
		this.type = new com.github.panda_io.micro_panda.ast.type.Function();
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
			for (Parameter parameter:this.parameters) {
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
				if (literal != null) {
					if (literal.token == Token.STRING) {
						f.Typ.ExternName = n
					}
				}
				if (this.type.externName == null) {
					context.addError(this.getOffset(), "'name' of meta data is required for extern function")
				}
			}
		} else {
			Context ctx = context.newContext();
			c.SetFunction(f)
			if f.Parent != nil {
				p := &ast_types.TypePointer{
					ElementType: f.Parent.Type(),
				}
				_ = c.AddObject(ast.StructThis, p)
			}
			if f.Typ.Extern {
				c.Error(f.GetPosition(), "extern function has no body")
			}
			if f.Parameters != nil {
				for _, param := range f.Parameters {
					err := c.AddObject(param.Name, param.Typ)
					if err != nil {
						c.Error(param.GetPosition(), err.Error())
					}
				}
			}
			f.Body.Validate(c)
		}
		//TO-DO check terminated
		//c.Program.Error(f.Position, "missing return")
	}
}
