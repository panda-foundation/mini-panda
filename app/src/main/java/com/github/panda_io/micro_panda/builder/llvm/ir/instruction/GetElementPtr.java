package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import java.util.ArrayList;
import java.util.List;

import com.github.panda_io.micro_panda.builder.llvm.ir.Identifier;
import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.constant.Constant;
import com.github.panda_io.micro_panda.builder.llvm.ir.constant.expression.GepIndex;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class GetElementPtr extends Instruction {
    Identifier identifier;
    Type elementType;
    Value source;
    List<Value> indexes;
    Type type;

    public GetElementPtr(Type elementType, Value source, List<Value> indexes) {
        this.elementType = elementType;
        this.source = source;
        this.indexes = indexes;
        this.identifier = new Identifier(false);
        this.getType();
    }

    public String string() {
        return String.format("%s %s", this.type.string(), this.identifier.identifier());
    }

    public Type getType() {
        if (this.type == null) {
            List<GepIndex> gepIndexes = new ArrayList<>();
            for (Value index : this.indexes) {
                if (index instanceof Constant) {
                    GepIndex gepIndex = GepIndex.getGepIndex((Constant) index);
                    gepIndexes.add(gepIndex);
                }
            }
            this.type = GepIndex.getGepType(this.elementType, gepIndexes);
        }
        return this.type;
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("%s = getelementptr %s, %s", this.identifier.identifier(),
                this.elementType.string(), this.source.string()));
        for (Value index : this.indexes) {
            builder.append(String.format(", %s", index.string()));
        }
    }
}