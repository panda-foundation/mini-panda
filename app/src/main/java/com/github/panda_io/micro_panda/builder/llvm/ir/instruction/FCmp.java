package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import com.github.panda_io.micro_panda.builder.llvm.ir.Identifier;
import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class FCmp extends Instruction {
    public enum Operand {
        False("false"),
        OEQ("oeq"),
        OGE("oge"),
        OGT("ogt"),
        OLE("ole"),
        OLT("olt"),
        ONE("one"),
        ORD("ord"),
        True("true"),
        UEQ("ueq"),
        UGE("uge"),
        UGT("ugt"),
        ULE("ule"),
        ULT("ult"),
        UNE("une"),
        UNO("uno");

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
    Value x;
    Value y;
    Type type;

    public FCmp(Operand operand, Value x, Value y) {
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
        builder.append(String.format("%s = fcmp %s %s, %s", this.identifier.identifier(), this.operand.toString(),
                this.x.string(), this.y.identifier()));
    }
}