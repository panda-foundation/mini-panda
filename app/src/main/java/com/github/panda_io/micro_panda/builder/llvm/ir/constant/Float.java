package com.github.panda_io.micro_panda.builder.llvm.ir.constant;

import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Float extends Constant {
    com.github.panda_io.micro_panda.builder.llvm.ir.type.Float type;
    double x;

    public Float(com.github.panda_io.micro_panda.builder.llvm.ir.type.Float type, double x) {
        this.type = type;
        this.x = x;
    }

    public static Float FromString(com.github.panda_io.micro_panda.builder.llvm.ir.type.Float type, String value) {
        return new Float(type, Double.parseDouble(value));
    }

    public String string() {
        return String.format("%s %s", this.type.string(), this.identifier());
    }

    public Type getType() {
        return this.type;
    }

    public String identifier() {
        long value = Double.doubleToRawLongBits(this.x);
        return String.format("0x%s", Long.toHexString(value));
    }
}
