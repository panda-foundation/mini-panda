package com.github.panda_io.micro_panda.builder.llvm.ir.constant.expression;

import java.util.ArrayList;
import java.util.List;

import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;
import com.github.panda_io.micro_panda.builder.llvm.ir.constant.Constant;

public class GetElementPtr extends Expression {
    public class Index extends Expression {
        Constant index;

        public Index(Constant index) {
            this.index = index;
        }

        public String string() {
            return this.index.string();
        }

        public Type getType() {
            return this.index.getType();
        }

        public String identifier() {
            return this.index.identifier();
        }
    }

    Type elementType;
    Constant source;
    List<Constant> indexes;
    Type type;

    public GetElementPtr(Type elementType, Constant source, List<Constant> indexes) {
        this.elementType = elementType;
        this.source = source;
        this.indexes = indexes;
        this.getType();
    }

    public String string() {
        return String.format("%s %s", this.type.string(), this.identifier());
    }

    public Type getType() {
        if (this.type == null) {
            List<com.github.panda_io.micro_panda.builder.llvm.ir.instruction.GetElementPtr.Index> idxs = new ArrayList<>();
            for (Constant index : this.indexes) {
                com.github.panda_io.micro_panda.builder.llvm.ir.instruction.GetElementPtr.Index idx = com.github.panda_io.micro_panda.builder.llvm.ir.instruction.GetElementPtr
                        .getGepIndex(index);
                idxs.add(idx);
            }
            this.type = com.github.panda_io.micro_panda.builder.llvm.ir.instruction.GetElementPtr
                    .getGepType(this.elementType, idxs);
        }
        return this.type;
    }

    public String identifier() {
        StringBuilder builder = new StringBuilder();
        builder.append(String.format("getelementptr (%s, %s", this.elementType.string(), this.source.string()));
        for (Constant index : this.indexes) {
            builder.append(String.format(", %s", index.string()));
        }
        builder.append(")");
        return builder.toString();
    }
}