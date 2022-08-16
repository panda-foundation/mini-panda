package com.github.panda_io.micro_panda.builder.llvm.ir.type;

public class Int extends Type {
    int bitSize;
    boolean unsigned;

    public Int(int bitSize, boolean unsigned) {
        this.bitSize = bitSize;
        this.unsigned = unsigned;
    }

    public boolean equal(Type type) {
        if (type instanceof Int) {
            return ((Int) type).bitSize == this.bitSize;
        }
        return false;
    }

    public String string() {
        return String.format("i%d", this.bitSize);
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("i%d", this.bitSize));
    }
}
