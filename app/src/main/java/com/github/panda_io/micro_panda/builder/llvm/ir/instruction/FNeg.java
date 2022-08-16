package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import com.github.panda_io.micro_panda.builder.llvm.ir.Identifier;
import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class FNeg extends Instruction {
    Identifier identifier;
    Value x;
    Type type;

    public FNeg(Value x) {
        this.x = x;
        this.identifier = new Identifier(false);
    }

    public String string() {
        return String.format("%s %s", this.type.string(), this.identifier.identifier());
    }

    public Type getType() {
        if (this.type == null) {
            this.type = this.x.getType();
        }
        return this.type;
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("%s = fneg %s", this.identifier.identifier(), this.x.string()));
    }
}