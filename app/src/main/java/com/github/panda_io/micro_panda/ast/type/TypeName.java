package com.github.panda_io.micro_panda.ast.type;

public class TypeName extends Type {
    public String name;
    public String qualified;
    public boolean isEnum;

    public TypeName() {
        this.isEnum = false;
    }

    public boolean equal(Type type) {
        if (type instanceof TypeName) {
            return this.qualified != null && this.qualified.equals(((TypeName) type).qualified);
        }
        return false;
    }

    public String string() {
        if (this.qualified != null) {
            return this.qualified;
        }
        return this.name;
    }
}
