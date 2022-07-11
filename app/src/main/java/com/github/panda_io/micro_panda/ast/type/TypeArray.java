package com.github.panda_io.micro_panda.ast.type;

import java.util.*;

public class TypeArray extends Type {
    public Type elementType;
    public List<Integer> dimensions;

    public TypeArray() {
        this.dimensions = new ArrayList<>();
    }

    public boolean equal(Type type) {
        if (type == null)
            return false;
        if (type instanceof TypeArray) {
            TypeArray array = (TypeArray) type;
            if (this.dimensions.size() == array.dimensions.size()) {
                for (int i = 1; i < this.dimensions.size(); i++) {
                    if (this.dimensions.get(i) != array.dimensions.get(i)) {
                        return false;
                    }
                }
                if (this.dimensions.get(0) == array.dimensions.get(0)) {
                    return true;
                }
                if (this.dimensions.get(0) == 0 || array.dimensions.get(0) == 0) {
                    return true;
                }
            }
        } else if (type instanceof TypePointer) {
            if (this.dimensions.size() == 1) {
                return this.elementType.equal(((TypePointer) type).elementType);
            }
        }
        return false;
    }

    public String string() {
        return String.format("array<%s>", this.elementType.string());
    }
}
