package com.github.panda_io.micro_panda.ast.type;

import java.util.ArrayList;

public class Array extends Type {
	Type elementType;
	ArrayList<Integer> dimensions;

    public Array(Type elementType) {
        this.elementType = elementType;
        this.dimensions = new ArrayList<>();
    }

    public boolean equal(Type type) {
        if (type instanceof Array) {
            Array array = (Array)type;
            if (this.dimensions.size() == array.dimensions.size()) {
                for (int i = 1; i < this.dimensions.size(); i++) {
                    if (this.dimensions.get(i) != array.dimensions.get(i)) {
                        return false;
                    }
                }
                return true;
            }
        } else if (type instanceof Pointer) {
            if (this.dimensions.size() == 1){
                return this.elementType.equal(((Pointer)type).elementType);
            }
        }
        return false;
    }
    
    public String string() {
        return String.format("array<%s>", this.elementType.string());
    }    
}
