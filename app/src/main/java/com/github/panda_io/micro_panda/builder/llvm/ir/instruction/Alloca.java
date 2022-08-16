package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import com.github.panda_io.micro_panda.builder.llvm.ir.Identifier;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Pointer;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Alloca extends Instruction {
    Identifier identifier;
    Type elementType;
    Type type;

    public Alloca(Type elementType) {
        this.elementType = elementType;
        this.identifier = new Identifier(false);
        this.getType();
    }

    public String string() {
        return String.format("%s %s", this.type.string(), this.identifier.identifier());
    }

    public Type getType() {
        if (this.type == null) {
            this.type = new Pointer(this.elementType);
        }
        return this.type;
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("%s = alloca %s", this.identifier.identifier(), this.elementType.string()));
    }
}