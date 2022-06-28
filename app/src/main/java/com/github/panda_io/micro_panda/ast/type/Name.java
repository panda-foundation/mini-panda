package com.github.panda_io.micro_panda.ast.type;

public class Name extends Type {
    String name;

    public String qualified;
    public boolean isEnum;

    public Name(String name) {
        this.name = name;
    }

    public boolean equal(Type type) {
        if (type instanceof Name) {
            return this.qualified != null && this.qualified.equals(((Name) type).qualified);
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
