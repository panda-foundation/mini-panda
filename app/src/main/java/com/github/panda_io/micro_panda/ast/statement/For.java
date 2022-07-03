package com.github.panda_io.micro_panda.ast.statement;

import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.ast.type.Type;

public class For extends Statement {
    public Statement initialization;
    public com.github.panda_io.micro_panda.ast.expression.Expression condition;
    public Statement post;
    public Statement body;

    public void validate(Context context) {
        Context ctx = context.newContext();
        if (this.initialization != null) {
            this.initialization.validate(context);
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
    }
}
