package com.github.panda_io.micro_panda.ast.expression;

import java.util.*;

import com.github.panda_io.micro_panda.ast.type.*;
import com.github.panda_io.micro_panda.scanner.Token;
import com.github.panda_io.micro_panda.ast.Context;

public class Invocation extends Expression {
	public Expression function;
	public List<Expression> arguments;
	public TypeFunction define;
    
	public boolean isLvalue() {
		return false;
	}
	
	public void validate(Context context, Type expected) {
		this.function.validate(context, null);
		this.constant = false;
		if (this.function.type == null) {
            return;
        }
		Type functionType = this.function.type;
		if (functionType instanceof TypeFunction) {
			this.define = (TypeFunction) functionType;
			this.type = this.define.returnType;
			if (this.type != null) {
				if (expected != null && !this.type.equal(expected)) {
					context.addError(this.getOffset(), String.format("mismatch return type, expect %s got %s",
							expected.string(), this.type.string()));
				}
			} else if (expected != null) {
				context.addError(this.getOffset(),
						String.format("mismatch return type, expect %s got null", expected.string()));
			}
			if (this.define.isMemberFunction) {
				// implicit conversion
				if (this.function instanceof MemberAccess) {
					MemberAccess memberAccess = (MemberAccess) this.function;
					if (memberAccess.parent.type.isStruct()) {
						Unary parent = new Unary();
						parent.operator = Token.BitAnd;
						parent.expression = memberAccess.parent;
						parent.setOffset(memberAccess.parent.getOffset());
						this.arguments.add(0, parent);
					} else if (memberAccess.parent.type.isPointer()) {
						this.arguments.add(0, memberAccess.parent);
					}
				}
			}
			if (this.define.parameters.size() == this.arguments.size()) {
				for (int i = 0; i < this.arguments.size(); i++) {
					this.arguments.get(i).validate(context, this.define.parameters.get(i));
				}
			} else {
				context.addError(this.getOffset(), "mismatch arguments and parameters");
			}
		} else {
			context.addError(this.getOffset(), "expect function");
		}
	}
}
