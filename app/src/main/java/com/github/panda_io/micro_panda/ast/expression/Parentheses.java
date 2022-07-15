package com.github.panda_io.micro_panda.ast.expression;

import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;

public class Parentheses extends Expression {
    public Expression expression;
    
	public boolean isLvalue() {
		return this.expression.isLvalue();
	}

    public void validate(Context context, Type expected) {
        this.expression.validate(context, expected);
        this.constant = this.expression.isConstant();
        this.type = this.expression.type;
    }
}
