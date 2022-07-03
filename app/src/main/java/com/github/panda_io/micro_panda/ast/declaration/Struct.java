package com.github.panda_io.micro_panda.ast.declaration;

import java.util.*;

import com.github.panda_io.micro_panda.ast.type.Name;
import com.github.panda_io.micro_panda.ast.type.Pointer;
import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;

public class Struct extends Declaration {
	public List<Variable> variables;
	public List<Function> functions;

	public boolean isConstant() {
		return false;
	}

    public Type getType()  {
		Name name = new Name(this.name.name);
		name.qualified = this.qualified;
		name.isEnum = false;
        return name;
    }

	public void resolveType(Context context) {
		if (this.variables == null || this.variables.size() == 0) {
			context.addError(this.getOffset(), "struct should contain at least 1 variable member");
		}
		for (Variable variable:this.variables) {
			variable.resolveType(context);
		}
		for (Function function:this.functions) {
			function.resolveType(context);
			function.qualified = String.format("%s.%s", this.qualified, function.name.name);
		}
    }

	public void validate(Context context) {
		for (Variable variable:this.variables) {
			variable.validate(context);
			if (variable.value != null) {
				context.addError(variable.getOffset(), "struct member has no initialize value");
			}
		}
		for (Function function:this.functions) {
			function.validate(context);
		}
	}

	public boolean hasMember(String member) {
		for (Variable variable:this.variables) {
			if (variable.name.name.equals(member)) {
				return true;
			}
		}
		for (Function function:this.functions) {
			if (function.name.name.equals(member)) {
				return true;
			}
		}
		return false;
	}
	
	public Type memberType(String member) {
		for (Variable variable:this.variables) {
			if (variable.name.name.equals(member)) {
				return variable.getType();
			}
		}
		for (Function function:this.functions) {
			if (function.name.name.equals(member)) {
				return function.getType();
			}
		}
		return null;
	}

	public boolean addVariable(Variable variable) {
		if (this.hasMember(variable.name.name)) {
			return false;
		}
		variable.parent = this;
		this.variables.add(variable);
		return true;
	}

	public boolean  addFunction(Function function) {
		if (this.hasMember(function.name.name)) {
			return false;
		}
		function.parent = this;
		this.functions.add(function);
		return true;
	}

	public Type pointerType() {
		return new Pointer(this.getType());
	}
}
