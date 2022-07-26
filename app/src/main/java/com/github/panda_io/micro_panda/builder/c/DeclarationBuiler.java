package com.github.panda_io.micro_panda.builder.c;

import com.github.panda_io.micro_panda.ast.Constant;
import com.github.panda_io.micro_panda.ast.type.*;
import com.github.panda_io.micro_panda.ast.declaration.*;
import com.github.panda_io.micro_panda.ast.declaration.Function.Parameter;

public class DeclarationBuiler {
    
    static void writeFunctionDefine(StringBuilder builder, Function function) {
        TypeBuiler.writeType(builder, function.returnType);
        if (function.qualified.equals(Constant.programEntry)) {
            builder.append(" main(");
        } else {
            builder.append(String.format(" %s(", function.qualified.replaceAll("\\.", "_")));
        }
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
        
    static void writeEnumDefine(StringBuilder builder, Enumeration enumeration) {
        for (Variable member : enumeration.members) {
            TypeBuiler.writeType(builder, member.type);
            builder.append(" ");
            builder.append(member.qualified.replaceAll("\\.", "_"));
            builder.append(" = ");
            ExpressionBuiler.writeExpression(builder, member.value);
            builder.append(";\n");
        }
        builder.append("\n");
    }

    static void writeStructDefine(StringBuilder builder, Struct struct) {
        builder.append("struct ");
        builder.append(struct.qualified.replaceAll("\\.", "_"));
        builder.append("\n{\n");
        for (Variable variable : struct.variables) {
            StatementBuiler.writeIndent(builder, 1);
            if (variable.type instanceof TypeName && !((TypeName)variable.type).isEnum) {
                builder.append("struct ");
            }
            if (variable.type instanceof TypeArray) {
                TypeArray array = (TypeArray)variable.type;
                if (array.elementType instanceof TypeName && !((TypeName)array.elementType).isEnum) {
                    builder.append("struct ");
                }
                if (array.elementType instanceof TypePointer) {
                    TypePointer pointer = (TypePointer)array.elementType;
                    if (pointer.elementType instanceof TypeName && !((TypeName)pointer.elementType).isEnum) {
                        builder.append("struct ");
                    }
                }
            }
            if (variable.type instanceof TypePointer) {
                TypePointer pointer = (TypePointer)variable.type;
                if (pointer.elementType instanceof TypeName && !((TypeName)pointer.elementType).isEnum) {
                    builder.append("struct ");
                }
            }
            TypeBuiler.writeType(builder, variable.type);
            builder.append(" ");
            builder.append(variable.name.name);
            if (variable.value != null) {
                builder.append(" = ");
                ExpressionBuiler.writeExpression(builder, variable.value);
            }
            builder.append(";\n");
        }
        builder.append("};\n\n");
    }

    static void writeVariable(StringBuilder builder, Variable variable) {
        TypeBuiler.writeType(builder, variable.type);
        builder.append(" ");
        builder.append(variable.qualified.replaceAll("\\.", "_"));
        if (variable.value != null) {
            builder.append(" = ");
            ExpressionBuiler.writeExpression(builder, variable.value);
        }
        builder.append(";\n\n");
    }

    static void writeFunction(StringBuilder builder, Function function) {
        writeFunctionDefine(builder, function);
        StatementBuiler.writeBlockStatement(builder, function.body, 0);
    }
}
