package com.github.panda_io.micro_panda.ast.statement;

import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.expression.Expression;

public class ForStatement extends Statement {
    public Statement initialization;
    public Expression condition;
    public Statement post;
    public BlockStatement body;

    public void validate(Context context) {
        Context ctx = context.newContext();
        ctx.loopLevel++;
        if (this.initialization != null) {
            this.initialization.validate(context);
            if (!this.initialization.isSimpleStatement()) {
                context.addError(this.initialization.getOffset(), "invalid init statement, expect simple statment");
            }
        }
        if (this.condition != null) {
            Context conditionCtx = ctx.newContext();
            this.condition.validate(conditionCtx, Type.bool);
            if (this.condition.getType() != null && !this.condition.getType().equal(Type.bool)) {
                context.addError(this.condition.getOffset(), "expect bool type condition");
            }
        }
        if (this.post != null) {
            Context postCtx = ctx.newContext();
            this.post.validate(postCtx);
        }
        if (this.body != null) {
            Context bodyCtx = ctx.newContext();
            this.body.validate(bodyCtx);
        }
        ctx.loopLevel--;
    }
}
