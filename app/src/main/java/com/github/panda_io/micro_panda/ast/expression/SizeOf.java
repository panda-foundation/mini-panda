package com.github.panda_io.micro_panda.ast.expression;

import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;

public class SizeOf extends Expression {
    public Type target;

    public void validate(Context context, Type expected) {
        context.resolveType(this.target);
        this.type = Type.u32;
        this.constant = true;
    }
}
