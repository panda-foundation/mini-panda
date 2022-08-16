package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Ret extends Terminator {
    Value x;

    public Ret(Value x) {
        this.x = x;
    }

    public String string() {
        return "";
    }

    public Type getType() {
        return null;
    }

    public void writeIR(StringBuilder builder) {
        if (this.x == null) {
            builder.append("ret void");
        } else {
            builder.append(String.format("ret %s", this.x.string()));
        }
    }
}