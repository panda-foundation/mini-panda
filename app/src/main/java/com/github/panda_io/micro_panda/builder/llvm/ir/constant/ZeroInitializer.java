package com.github.panda_io.micro_panda.builder.llvm.ir.constant;

import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class ZeroInitializer extends Constant {
    Type type;

    public ZeroInitializer(Type type) {
        this.type = type;
    }

    public String string() {
        return String.format("%s zeroinitializer", this.type.string());
    }

    public Type getType() {
        return this.type;
    }

    public String identifier() {
        return "zeroinitializer";
    }
}
