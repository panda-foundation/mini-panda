package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import java.util.List;

import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Switch extends Terminator {
    public class Case {
        Value x;
        Value target;

        public Case(Value x, Value target) {
            this.x = x;
            this.target = target;
        }

        public String string() {
            return String.format("%s, %s", this.x.string(), this.target.string());
        }
    }

    Value x;
    Value targetDefault;
    List<Case> cases;

    public Switch(Value x, Value targetDefault, List<Case> cases) {
        this.x = x;
        this.targetDefault = targetDefault;
        this.cases = cases;
    }

    public String string() {
        return "";
    }

    public Type getType() {
        return null;
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("switch %s, %s [\n", this.x.string(), this.targetDefault.string()));
        for (Case target : this.cases) {
            builder.append(String.format("\t\t%s\n", target.string()));
        }
        builder.append("\t]");
    }
}