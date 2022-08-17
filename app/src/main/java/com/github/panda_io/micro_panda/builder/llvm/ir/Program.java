package com.github.panda_io.micro_panda.builder.llvm.ir;

import java.util.ArrayList;
import java.util.List;

import com.github.panda_io.micro_panda.builder.llvm.ir.constant.Constant;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Program {
    List<Type> structs;
    List<Global> globals;
    List<Function> functions;

    public Program() {
        this.structs = new ArrayList<>();
        this.globals = new ArrayList<>();
        this.functions = new ArrayList<>();
    }

    public Function addFunction(String name, Type returnType, List<Function.Parameter> parameters) {
        Function function = new Function(name, returnType, parameters);
        this.functions.add(function);
        return function;
    }

    public Global addGlobal(String name, Constant initializer) {
        Global global = new Global(name, initializer);
        this.globals.add(global);
        return global;
    }

    public Type addStruct(String name, Type type) {
        this.structs.add(type);
        return type;
    }

    public void writeIR(StringBuilder builder) {
        for (Type type : this.structs) {
            builder.append(String.format("%s = type ", type.string()));
            type.writeIR(builder);
            builder.append("\n");
        }
        builder.append("\n");
        for (Global global : this.globals) {
            global.writeIR(builder);
        }
        builder.append("\n");
        for (int i = 0; i < this.functions.size(); i++) {
            if (i != 0) {
                builder.append("\n");
            }
            this.functions.get(i).writeIR(builder);
        }
    }
}
