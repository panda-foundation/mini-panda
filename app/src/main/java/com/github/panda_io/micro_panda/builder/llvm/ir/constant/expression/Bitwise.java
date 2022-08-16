package com.github.panda_io.micro_panda.builder.llvm.ir.constant.expression;

import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;
import com.github.panda_io.micro_panda.builder.llvm.ir.constant.Constant;

public class Bitwise extends Expression {
    com.github.panda_io.micro_panda.builder.llvm.ir.instruction.Bitwise.Operator operator;
    Constant x;
    Constant y;
    Type type;

    public Bitwise(com.github.panda_io.micro_panda.builder.llvm.ir.instruction.Bitwise.Operator operator, Constant x,
            Constant y) {
        this.operator = operator;
        this.x = x;
        this.y = y;
        this.getType();
    }

    public String string() {
        return String.format("%s %s", this.type.string(), this.identifier());
    }

    public Type getType() {
        if (this.type == null) {
            this.type = this.x.getType();
        }
        return this.type;
    }

    public String identifier() {
        return String.format("%s (%s, %s)", this.operator.toString(), this.x.string(), this.y.string());
    }
}