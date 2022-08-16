package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import com.github.panda_io.micro_panda.builder.llvm.ir.Identifier;
import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Bitwise extends Instruction {
    public enum Operator {
        Shl("shl"),
        LShr("lshl"),
        AShr("ashr"),
        And("and"),
        Or("or"),
        Xor("xor");

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
    Value x;
    Value y;
    Type type;

    public Bitwise(Operator operator, Value x, Value y) {
        this.operator = operator;
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
            this.type = this.x.getType();
        }
        return this.type;
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("%s = %s %s, %s", this.identifier.identifier(), this.operator.toString(),
                this.x.string(), this.y.identifier()));
    }
}
