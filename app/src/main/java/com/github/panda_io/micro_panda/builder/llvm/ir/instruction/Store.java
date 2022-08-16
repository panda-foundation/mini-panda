package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Store extends Instruction {
    Value source;
    Value destination;

    public Store(Value source, Value destination) {
        this.source = source;
        this.destination = destination;
    }

    public String string() {
        return "";
    }

    public Type getType() {
        return null;
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("store %s, %s", this.source.string(), this.destination.string()));
    }
}