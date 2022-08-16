package com.github.panda_io.micro_panda.builder.llvm.ir.type;

public class Void extends Type {
    public boolean equal(Type type) {
        return type instanceof Void;
    }

    public String string() {
        return "void";
    }

    public void writeIR(StringBuilder builder) {
        builder.append("void");
    }
}
