package com.github.panda_io.micro_panda.ast.expression;

import com.github.panda_io.micro_panda.ast.Node;
import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;

public abstract class Expression extends Node {
    boolean constant;
    Type type;

    public Type getType() {
        return this.type;
    }

	public boolean isConstant() {
        return this.constant;
    }

	public abstract void validate(Context context, Type expected);
}
