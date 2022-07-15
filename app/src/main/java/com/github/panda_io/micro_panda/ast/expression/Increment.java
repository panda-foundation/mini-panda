package com.github.panda_io.micro_panda.ast.expression;

import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;

public class Increment extends Expression {
    public Expression expression;
    
	public boolean isLvalue() {
		return false;
	}
    
    public void validate(Context context, Type expected) {
        this.constant = false;
        this.expression.validate(context, expected);
        if (this.expression.type == null) {
            return;
        }
        if (this.expression.isConstant()) {
            context.addError(this.expression.getOffset(), "expect variable");
        }
        if (!this.expression.type.isInteger()) {
            context.addError(this.expression.getOffset(), "expect integer expression");
        }
    }
}
