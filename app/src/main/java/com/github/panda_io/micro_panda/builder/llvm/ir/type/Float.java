package com.github.panda_io.micro_panda.builder.llvm.ir.type;

public class Float extends Type {
    public enum Kind {
        Half("half"),
        Float("float"),
        Double("double");

        final String kind;

        private Kind(String kind) {
            this.kind = kind;
        }

        @Override
        public String toString() {
            return this.kind;
        }
    };

    Kind kind;

    public Float(Kind kind) {
        this.kind = kind;
    }

    public boolean equal(Type type) {
        if (type instanceof Float) {
            return ((Float) type).kind == this.kind;
        }
        return false;
    }

    public String string() {
        return this.kind.toString();
    }

    public void writeIR(StringBuilder builder) {
        builder.append(this.kind.toString());
    }
}
