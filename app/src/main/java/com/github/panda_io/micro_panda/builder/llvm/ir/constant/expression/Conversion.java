package com.github.panda_io.micro_panda.builder.llvm.ir.constant.expression;

import com.github.panda_io.micro_panda.builder.llvm.ir.constant.Constant;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Conversion extends Expression {
    com.github.panda_io.micro_panda.builder.llvm.ir.instruction.Conversion.Operator operator;
    Constant from;
    Type to;

    public Conversion(com.github.panda_io.micro_panda.builder.llvm.ir.instruction.Conversion.Operator operator,
            Constant from, Type to) {
        this.operator = operator;
        this.from = from;
        this.to = to;
    }

    public String string() {
        return String.format("%s %s", this.to.string(), this.identifier());
    }

    public Type getType() {
        return this.to;
    }

    public String identifier() {
        return String.format("%s (%s to %s)", this.operator.toString(), this.from.string(), this.to.string());
    }
}
