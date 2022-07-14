package com.github.panda_io.micro_panda.builder.c;

import com.github.panda_io.micro_panda.ast.declaration.Function;
import com.github.panda_io.micro_panda.ast.declaration.Struct;
import com.github.panda_io.micro_panda.ast.declaration.Function.Parameter;

public class DeclarationBuiler {
    
    static void writeFunctionDefine(StringBuilder builder, Function function) {
        TypeBuiler.writeType(builder, function.returnType);
        builder.append(String.format(" %s(", function.qualified.replaceAll("\\.", "_")));
        for (int i = 0; i < function.parameters.size(); i++) {
            Parameter parameter = function.parameters.get(i);
            if (i > 0) {
                builder.append(", ");
            }
            TypeBuiler.writeType(builder, parameter.type);
            builder.append(String.format(" %s", parameter.name));            
        }
        builder.append(")");
    }
        
    static void writeStructDefine(StringBuilder builder, Struct struct) {

    }

    static void writeFunction(StringBuilder builder, Function function) {
        writeFunctionDefine(builder, function);
        StatementBuiler.writeBlockStatement(builder, function.body, 0);
    }
}
