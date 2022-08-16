package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import com.github.panda_io.micro_panda.builder.llvm.ir.Identifier;
import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Conversion extends Instruction {
    public enum Operator {
        Trunc("trunc"),
        SExt("sext"),
        FPTrunc("fptrunc"),
        FPExt("fpext"),
        FPToUI("fptoui"),
        FPToSI("fptosi"),
        UIToFP("uitofp"),
        SIToFP("sitofp"),
        BitCast("bitcast");

        final String operator;

        private Operator(String operator) {
            this.operator = operator;
        }

        @Override
        public String toString() {
            return this.operator;
        }
    }

    Operator operator;
    Identifier identifier;
    Value from;
    Type to;

    public Conversion(Operator operator, Value from, Type to) {
        this.operator = operator;
        this.from = from;
        this.to = to;
        this.identifier = new Identifier(false);
    }

    public String string() {
        return String.format("%s %s", this.to.string(), this.identifier.identifier());
    }

    public Type getType() {
        return this.to;
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("%s = %s %s to %s", this.identifier.identifier(), this.operator.toString(),
                this.from.string(), this.to.string()));
    }
}
