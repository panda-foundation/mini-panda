package com.github.panda_io.micro_panda.ast.statement;

import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.ast.type.*;
import com.github.panda_io.micro_panda.ast.expression.Expression;

public class ReturnStatement extends Statement {
    public Expression expression;

    public void validate(Context context) {
        Type returnType = context.getFunction().returnType;
        if (returnType == null) {
            if (this.expression != null) {
                context.addError(this.getOffset(), "mismatch return type, expect 'void'");
            }
        } else {
            if (this.expression == null) {
                context.addError(this.getOffset(),
                        String.format("mismatch return type, expect '%s' got 'void'", returnType.string()));
            }
            this.expression.validate(context, returnType);
            if (this.expression.getType() != null && !this.expression.getType().equal(returnType)) {
                context.addError(this.getOffset(), String.format("mismatch return type, expect '%s' got '%s'",
                        returnType.string(), this.expression.getType().string()));
            }
        }
        //TO-DO cannot return struct or array
    }
}
