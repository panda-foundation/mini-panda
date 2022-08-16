package com.github.panda_io.micro_panda.builder.llvm.ir.type;

import java.util.List;

public class Function extends Type {
    public String qualified;
    public Type returnType;
    public List<Type> parameters;

    public Function(String qualified, Type returnType, List<Type> parameters) {
        this.qualified = qualified;
        this.returnType = returnType;
        this.parameters = parameters;
    }

    public boolean equal(Type type) {
        if (type instanceof Function) {
            return ((Function) type).qualified.equals(this.qualified);
        }
        return false;
    }

    public String string() {
        return this.qualified;
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("%s (", this.returnType.string()));
        for (int i = 0; i < this.parameters.size(); i++) {
            if (i != 0) {
                builder.append(", ");
            }
            builder.append(this.parameters.get(i).string());
        }
        builder.append(")");
    }
}
