package com.github.panda_io.micro_panda.ast.statement;

import com.github.panda_io.micro_panda.ast.Context;

public class BreakStatement extends Statement {
    public void validate(Context context) {
        if (context.loopLevel == 0) {
            context.addError(this.getOffset(), "invalid 'break'");
        }
    }
}
