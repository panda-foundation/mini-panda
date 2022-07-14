package com.github.panda_io.micro_panda.ast.statement;

import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.expression.Expression;

public class IfStatement extends Statement {
    public Expression condition;
    public BlockStatement body;
    public Statement elseStatement;

    public void validate(Context context) {
        Context ctx = context.newContext();
        if (this.condition == null) {
            context.addError(this.getOffset(), "expect condition expression");
        } else {
            this.condition.validate(ctx, Type.bool);
            if (this.condition.getType() != null && !this.condition.getType().equal(Type.bool)) {
                context.addError(this.condition.getOffset(), "expect bool type condition");
            }
        }
        if (this.body != null) {
            Context bodyCtx = ctx.newContext();
            this.body.validate(bodyCtx);
        }
        if (this.elseStatement != null) {
            Context elseCtx = ctx.newContext();
            this.elseStatement.validate(elseCtx);
        }
    }
}
