package com.github.panda_io.micro_panda.builder.llvm.ir.type;

import java.util.List;

public class Struct extends Type {
    String qualified;
    List<Type> fields;

    public Struct(String qualified, List<Type> fields) {
        this.qualified = qualified;
        this.fields = fields;
    }

    public boolean equal(Type type) {
        if (type instanceof Struct) {
            return ((Struct) type).qualified.equals(this.qualified);
        }
        return false;
    }

    public String string() {
        return this.qualified;
    }

    public void writeIR(StringBuilder builder) {
        if (this.fields == null || this.fields.size() == 0) {
            builder.append("{}");
            return;
        }
        builder.append("{ ");
        for (int i = 0; i < this.fields.size(); i++) {
            if (i != 0) {
                builder.append(", ");
            }
            builder.append(this.fields.get(i).string());
        }
        builder.append(" }");
    }
}
