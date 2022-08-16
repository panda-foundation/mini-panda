package com.github.panda_io.micro_panda.builder.llvm.ir.type;

public class Label extends Type {
    public boolean equal(Type type) {
        return type instanceof Label;
    }

    public String string() {
        return "label";
    }

    public void writeIR(StringBuilder builder) {
        builder.append("label");
    }
}
