package com.github.panda_io.micro_panda.builder.llvm.ir.constant;

import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Int extends Constant {
    com.github.panda_io.micro_panda.builder.llvm.ir.type.Int type;
    long x;

    public Int(com.github.panda_io.micro_panda.builder.llvm.ir.type.Int type, long x) {
        this.type = type;
        this.x = x;
    }

    public static Int Bool(boolean x) {
        if (x) {
            return Constant.True;
        }
        return Constant.False;
    }

    public static Int FromString(com.github.panda_io.micro_panda.builder.llvm.ir.type.Int type, String value) {
        if (value.equals("true")) {
            return Constant.True;
        } else if (value.equals("false")) {
            return Constant.False;
        }
        try {
            long x = Long.parseLong(value);
            return new Int(type, x);
        } catch (Exception e) {
            return null;
        }
    }

    public String string() {
        return String.format("%s %s", this.type.string(), this.identifier());
    }

    public Type getType() {
        return this.type;
    }

    public String identifier() {
        if (this.type.bitSize == 1) {
            if (this.x == 0) {
                return "false";
            }
            return "true";
        }
        return Long.toString(this.x);
    }
}
