package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import com.github.panda_io.micro_panda.builder.llvm.ir.Identifier;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public abstract class Instruction {
    Identifier identifier;

    public Identifier getIdentifier() {
        return this.identifier;
    }

    public abstract Type getType();

    public abstract String string();

    public abstract void writeIR(StringBuilder builder);
}
