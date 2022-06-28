package com.github.panda_io.micro_panda.ast.type;

public class Pointer extends Type {
    Type elementType;

    public Pointer(Type elementType) {
        this.elementType = elementType;
    }

    public boolean equal(Type type) {
        if (type instanceof Pointer) {
            return this.elementType.equal(((Pointer) type).elementType);
        } else if (type instanceof Array) {
            if (((Array) type).dimensions.size() == 1) {
                return this.elementType.equal(((Array) type).elementType);
            }
        }
        return false;
    }

    public String string() {
        return String.format("pointer<%s>", this.elementType.string());
    }
}
