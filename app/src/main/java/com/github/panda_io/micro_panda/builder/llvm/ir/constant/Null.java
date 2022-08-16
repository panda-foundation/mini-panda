package com.github.panda_io.micro_panda.builder.llvm.ir.constant;

import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Null extends Constant {
    Type type;

    public Null(Type type) {
        this.type = type;
    }

    public String string() {
        return String.format("%s null", this.type.string());
    }

    public Type getType() {
        return this.type;
    }

    public String identifier() {
        return "null";
    }
}
