package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import java.util.List;

import com.github.panda_io.micro_panda.builder.llvm.ir.Identifier;
import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Function;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Pointer;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Call extends Instruction {
    Value callee;
    List<Value> arguments;
    Type type;

    public Call(Value callee, List<Value> arguments) {
        this.callee = callee;
        this.arguments = arguments;
        this.identifier = new Identifier(false);
        this.getType();
    }

    public String string() {
        return String.format("%s %s", this.type.string(), this.identifier.identifier());
    }

    public Type getType() {
        if (this.type == null) {
            Type calleeType = this.callee.getType();
            Function function = (Function) (((Pointer) calleeType).elementType);
            this.type = function.returnType;
        }
        return this.type;
    }

    public void writeIR(StringBuilder builder) {
        if (!this.type.equal(Type.Void)) {
            builder.append(String.format("%s = ", this.identifier.identifier()));
        }
        builder.append(String.format("call %s %s(", this.type.string(), this.callee.identifier()));
        for (int i = 0; i < this.arguments.size(); i++) {
            if (i != 0) {
                builder.append(", ");
            }
            builder.append(this.arguments.get(i).string());
        }
        builder.append(")");
    }
}