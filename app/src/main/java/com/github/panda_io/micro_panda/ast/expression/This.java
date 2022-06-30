package com.github.panda_io.micro_panda.ast.expression;

import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.ast.Constant;
import com.github.panda_io.micro_panda.ast.type.Type;

public class This extends Expression {
    public void validate(Context context, Type expected) {
        this.constant = false;
        this.type = context.findObject(Constant.structThis);
        if (this.type == null) {
            context.addError(this.getOffset(), "undefined \"this\"");
        }
    }
}
