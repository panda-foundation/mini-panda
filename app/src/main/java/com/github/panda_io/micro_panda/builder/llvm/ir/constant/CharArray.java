package com.github.panda_io.micro_panda.builder.llvm.ir.constant;

import com.github.panda_io.micro_panda.builder.llvm.ir.Encode;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class CharArray {
    com.github.panda_io.micro_panda.builder.llvm.ir.type.Array type;
    byte[] x;

    public CharArray(byte[] x) {
        this.x = x;
        this.type = new com.github.panda_io.micro_panda.builder.llvm.ir.type.Array(this.x.length, Type.I8);
    }

    public static CharArray FromString(String value) {
        return new CharArray(value.getBytes());
    }

    public String string() {
        return String.format("%s %s", this.type.string(), this.identifier());
    }

    public Type getType() {
        return this.type;
    }

    public String identifier() {
        return "c" + Encode.quote(this.x);
    }
}
