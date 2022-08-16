package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class If extends Terminator {
    Value condition;
    Value targetTrue;
    Value targetFalse;

    public If(Value condition, Value targetTrue, Value targetFalse) {
        this.condition = condition;
        this.targetTrue = targetTrue;
        this.targetFalse = targetFalse;
    }

    public String string() {
        return "";
    }

    public Type getType() {
        return null;
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("br %s, %s, %s", this.condition.string(), this.targetTrue.string(),
                this.targetFalse.string()));
    }
}