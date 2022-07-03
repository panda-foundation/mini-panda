package com.github.panda_io.micro_panda.ast.statement;

import com.github.panda_io.micro_panda.ast.Context;

public class Expression extends Statement {
    public com.github.panda_io.micro_panda.ast.expression.Expression expression;

    public void validate(Context context) {
        this.expression.validate(context, null);
    }
}