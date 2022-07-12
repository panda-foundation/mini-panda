package com.github.panda_io.micro_panda.ast.expression;

import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;

public class Conversion extends Expression {
    public Expression value;

    public void validate(Context context, Type expected) {
        this.type = context.resolveType(this.type);
        this.value.validate(context, null);
        if (this.value.type == null) {
            return;
        }
        this.constant = this.value.isConstant();
        boolean isNumber = this.type.isNumber() && this.value.type.isNumber();
        boolean isPointer = this.type.isPointer() && this.value.type.isPointer();
        if (!(isNumber || isPointer)) {
            context.addError(this.type.getOffset(), "invalid type conversion");
        }
    }
}
