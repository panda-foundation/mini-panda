package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import com.github.panda_io.micro_panda.builder.llvm.ir.Identifier;
import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class ICmp extends Instruction {
    public enum Operand {
        EQ("eq"),
        NE("ne"),
        SGE("sge"),
        SGT("sgt"),
        SLE("sle"),
        SLT("slt"),
        UGE("uge"),
        UGT("ugt"),
        ULE("ule"),
        ULT("ult");

        final String operand;

        private Operand(String operand) {
            this.operand = operand;
        }

        @Override
        public String toString() {
            return this.operand;
        }
    }

    Operand operand;
    Identifier identifier;
    Value x;
    Value y;
    Type type;

    public ICmp(Operand operand, Value x, Value y) {
        this.operand = operand;
        this.x = x;
        this.y = y;
        this.identifier = new Identifier(false);
        this.getType();
    }

    public String string() {
        return String.format("%s %s", this.type.string(), this.identifier.identifier());
    }

    public Type getType() {
        if (this.type == null) {
            this.type = Type.I1;
        }
        return this.type;
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("%s = icmp %s %s, %s", this.identifier.identifier(), this.operand.toString(),
                this.x.string(), this.y.identifier()));
    }
}