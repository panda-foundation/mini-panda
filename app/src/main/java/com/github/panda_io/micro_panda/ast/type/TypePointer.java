package com.github.panda_io.micro_panda.ast.type;

public class TypePointer extends Type {
    public Type elementType;

    public TypePointer(Type elementType) {
        this.elementType = elementType;
    }

    public boolean equal(Type type) {
        if (type instanceof TypePointer) {
            return this.elementType.equal(((TypePointer) type).elementType);
        } else if (type instanceof TypeArray) {
            if (((TypeArray) type).dimensions.size() == 1) {
                return this.elementType.equal(((TypeArray) type).elementType);
            }
        }
        return false;
    }

    public String string() {
        return String.format("pointer<%s>", this.elementType.string());
    }
}
