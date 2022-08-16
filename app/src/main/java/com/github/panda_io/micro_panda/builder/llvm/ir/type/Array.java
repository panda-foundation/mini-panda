package com.github.panda_io.micro_panda.builder.llvm.ir.type;

public class Array extends Type {
    long length;
    Type elementType;

    public Array(long length, Type elementType) {
        this.length = length;
        this.elementType = elementType;
    }

    public boolean equal(Type type) {
        if (type instanceof Array) {
            Array array = (Array) type;
            return array.length == this.length && array.elementType.equal(this.elementType);
        }
        return false;
    }

    public String string() {
        return String.format("[%d x %s]", this.length, this.elementType.string());
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("[%d x %s]", this.length, this.elementType.string()));
    }
}
