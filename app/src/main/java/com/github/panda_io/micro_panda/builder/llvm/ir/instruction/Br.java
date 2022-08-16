package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Br extends Terminator {
    Value target;

    public Br(Value target) {
        this.target = target;
    }

    public String string() {
        return "";
    }

    public Type getType() {
        return null;
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("br %s", this.target.string()));
    }
}