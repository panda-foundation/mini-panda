package com.github.panda_io.micro_panda.builder.llvm.ir.constant.expression;

import com.github.panda_io.micro_panda.builder.llvm.ir.constant.Constant;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class ICmp extends Expression {
    com.github.panda_io.micro_panda.builder.llvm.ir.instruction.ICmp.Operand operand;
    Constant x;
    Constant y;
    Type type;

    public ICmp(com.github.panda_io.micro_panda.builder.llvm.ir.instruction.ICmp.Operand operand, Constant x,
            Constant y) {
        this.operand = operand;
        this.x = x;
        this.y = y;
        this.getType();
    }

    public String string() {
        return String.format("%s %s", this.type.string(), this.identifier());
    }

    public Type getType() {
        if (this.type == null) {
            this.type = Type.I1;
        }
        return this.type;
    }

    public String identifier() {
        return String.format("icmp %s (%s, %s)", this.operand.toString(), this.x.string(), this.y.string());
    }
}