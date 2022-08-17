package com.github.panda_io.micro_panda.builder.llvm.ir;

import java.util.ArrayList;
import java.util.List;

import com.github.panda_io.micro_panda.builder.llvm.ir.constant.Constant;
import com.github.panda_io.micro_panda.builder.llvm.ir.instruction.Instruction;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Pointer;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Function extends Constant {
    public class Parameter {
        Identifier identifier;
        Type type;

        public Parameter(Type type) {
            this.type = type;
            this.identifier = new Identifier(false);
        }

        public String string() {
            return String.format("%s %s", this.type.string(), this.identifier.identifier());
        }

        public Type getType() {
            return this.type;
        }

        public void writeIR(StringBuilder builder) {
            builder.append(String.format("%s %s", this.type.string(), this.identifier.identifier()));
        }
    }

    Identifier _identifier;
    com.github.panda_io.micro_panda.builder.llvm.ir.type.Function signature;
    List<Parameter> parameters;
    List<Block> blocks;
    Type type;

    public Function(String name, Type returnType, List<Parameter> parameters) {
        this._identifier = new Identifier(true);
        this._identifier.setName(name);
        this.parameters = parameters;
        List<Type> parameterTypes = new ArrayList<>();
        for (Parameter parameter : parameters) {
            parameterTypes.add(parameter.getType());
        }
        this.signature = new com.github.panda_io.micro_panda.builder.llvm.ir.type.Function(name, returnType,
                parameterTypes);
        this.blocks = new ArrayList<>();
        this.getType();
    }

    public String string() {
        return String.format("%s %s", this.type.string(), this.identifier());
    }

    public Type getType() {
        if (this.type == null) {
            this.type = new Pointer(this.signature);
        }
        return this.type;
    }

    public String identifier() {
        return this._identifier.identifier();
    }

    public void writeIR(StringBuilder builder) {
        this.setIds();
        if (this.blocks.size() == 0) {
            builder.append("declare ");
            this.writeHeaderIR(builder);
        } else {
            builder.append("define ");
            this.writeHeaderIR(builder);
            this.writeBodyIR(builder);
        }
    }

    public Block createBlock(String name) {
        Block block = new Block(name);
        this.blocks.add(block);
        return block;
    }

    public void writeHeaderIR(StringBuilder builder) {
        builder.append(String.format(" %s %s", this.signature.returnType.string(), this.identifier()));
        for (int i = 0; i < this.parameters.size(); i++) {
            if (i != 0) {
                builder.append(", ");
            }
            this.parameters.get(i).writeIR(builder);
        }
        builder.append(")");
    }

    public void writeBodyIR(StringBuilder builder) {
        builder.append(" {\n");
        for (int i = 0; i < this.blocks.size(); i++) {
            if (i != 0) {
                builder.append("\n");
            }
            this.blocks.get(i).writeIR(builder);
            builder.append("\n");
        }
        builder.append("}");
    }

    void setIds() {
        int id = 0;
        for (Parameter parameter : this.parameters) {
            id = this.setId(parameter.identifier, id);
        }
        for (Block block : this.blocks) {
            id = this.setId(block.identifier, id);
            for (Instruction instruction : block.instructions) {
                Identifier identifier = instruction.getIdentifier();
                if (identifier != null) {
                    if (!instruction.getType().equal(Type.Void)) {
                        id = this.setId(identifier, id);
                    }
                }
            }
        }
    }

    int setId(Identifier name, int id) {
        if (name.name.isEmpty()) {
            name.setId(id);
            id++;
        }
        return id;
    }
}
