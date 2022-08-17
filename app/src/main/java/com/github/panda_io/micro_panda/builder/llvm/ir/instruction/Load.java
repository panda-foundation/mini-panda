package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import com.github.panda_io.micro_panda.builder.llvm.ir.Identifier;
import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Load extends Instruction {
    Type elementType;
    Value source;

    public Load(Type elementType, Value source) {
        this.elementType = elementType;
        this.source = source;
        this.identifier = new Identifier(false);
    }

    public String string() {
        return String.format("%s %s", this.elementType.string(), this.identifier.identifier());
    }

    public Type getType() {
        return this.elementType;
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("%s = load %s, %s", this.identifier.identifier(), this.elementType.string(),
                this.source.string()));
    }
}