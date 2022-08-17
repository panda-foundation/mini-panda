package com.github.panda_io.micro_panda.builder.llvm.ir;

import java.util.ArrayList;
import java.util.List;

import com.github.panda_io.micro_panda.builder.llvm.ir.instruction.Alloca;
import com.github.panda_io.micro_panda.builder.llvm.ir.instruction.Instruction;
import com.github.panda_io.micro_panda.builder.llvm.ir.instruction.Terminator;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Block {
    Identifier identifier;
    List<Instruction> instructions;
    boolean terminated;

    public Block(String name) {
        this.identifier = new Identifier(false);
        this.identifier.setName(name);
        this.instructions = new ArrayList<>();
        this.terminated = false;
    }

    public String string() {
        return String.format("%s %s", this.getType().string(), this.identifier.identifier());
    }

    public Type getType() {
        return Type.Label;
    }

    public void addInstruction(Instruction instruction) {
        this.instructions.add(instruction);
        if (instruction instanceof Terminator) {
            this.terminated = true;
        }
    }

    public void insertAlloca(Alloca alloca) {
        if (this.terminated) {
            int position = this.instructions.size() - 1;
            this.instructions.add(position, alloca);
        } else {
            this.instructions.add(alloca);
        }
    }

    public void writeIR(StringBuilder builder) {
        String name = "";
        if (this.identifier.name.isEmpty()) {
            name = Encode.localID(this.identifier.id);
        } else {
            name = Encode.localName(this.identifier.name);
        }
        builder.append(String.format("%s\n", name));
        for (Instruction instruction : this.instructions) {
            builder.append("\t");
            instruction.writeIR(builder);
            builder.append("\n");
        }
    }
}