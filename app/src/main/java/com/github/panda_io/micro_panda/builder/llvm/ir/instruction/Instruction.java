package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public abstract class Instruction {
    public abstract Type getType();
    public abstract String string();
    public abstract void writeIR(StringBuilder builder);
}
