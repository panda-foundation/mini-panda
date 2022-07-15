package com.github.panda_io.micro_panda.ast.declaration;

import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.ast.expression.Expression;

public class Variable extends Declaration {
    public Type type;
    public boolean constant;
    public Expression value;
    public Struct parent;

    public boolean isConstant() {
        return this.constant;
    }

    public Type getType() {
        return this.type;
    }

    public void resolveType(Context context) {
        this.type = context.resolveType(this.type);
    }

    public void validate(Context context) {
        if (this.value != null) {
            this.value.validate(context, this.type);
        }
        if (this.constant) {
            if (this.value == null) {
                context.addError(this.getOffset(), "const must be initialized when declare");
            } else if (!this.value.isConstant()) {
                context.addError(this.getOffset(), "expect constant expression");
            }
        }
        if (this.value != null) {
            if (this.value.getType() == null) {
                context.addError(this.value.getOffset(), "unknown type");
            } else if (!this.value.getType().equal(this.type)) {
                context.addError(this.value.getOffset(), "init value type mismatch with define");
            }
            if (!this.value.isConstant()) {
                context.addError(this.value.getOffset(), "expect const expression");
            }
        }
    }
}
