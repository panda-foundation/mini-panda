package com.github.panda_io.micro_panda.builder.llvm.ir.type;

public class Pointer extends Type {
    Type elementType;

    public Pointer(Type elementType) {
        this.elementType = elementType;
    }

    public boolean equal(Type type) {
        if (type instanceof Pointer) {
            return ((Pointer) type).elementType.equal(this.elementType);
        }
        return false;
    }

    public String string() {
        return String.format("%s*", this.elementType.string());
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("%s*", this.elementType.string()));
    }
}
