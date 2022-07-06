package com.github.panda_io.micro_panda.ast.statement;

import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.ast.expression.Expression;

public class ExpressionStatement extends Statement {
    public Expression expression;

    public void validate(Context context) {
        this.expression.validate(context, null);
    }
}