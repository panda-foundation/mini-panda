package com.github.panda_io.micro_panda.builder.llvm.ir.constant.expression;

import com.github.panda_io.micro_panda.builder.llvm.ir.constant.Constant;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class FNeg extends Expression {
    Constant x;
    Type type;

    public FNeg(Constant x) {
        this.x = x;
        this.getType();
    }

    public String string() {
        return String.format("%s %s", this.type.string(), this.identifier());
    }

    public Type getType() {
        if (this.type == null) {
            this.type = this.x.getType();
        }
        return this.type;
    }

    public String identifier() {
        return String.format("fneg (%s)", this.x.string());
    }
}