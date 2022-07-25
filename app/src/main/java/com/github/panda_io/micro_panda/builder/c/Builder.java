package com.github.panda_io.micro_panda.builder.c;

import java.util.Collection;

import com.github.panda_io.micro_panda.ast.Module;
import com.github.panda_io.micro_panda.ast.Program;
import com.github.panda_io.micro_panda.ast.declaration.Enumeration;
import com.github.panda_io.micro_panda.ast.declaration.Function;
import com.github.panda_io.micro_panda.ast.declaration.Struct;
import com.github.panda_io.micro_panda.ast.declaration.Variable;

public class Builder {

    public static StringBuilder build(Program program) {
        StringBuilder builder = new StringBuilder();
        builder.append("#include <stdint.h>\n\n");
        Collection<Module> modules = program.getModules();

        for (Module module : modules) {
            for (Struct struct : module.structs) {
                builder.append("struct ");
                builder.append(struct.qualified.replaceAll("\\.", "_"));
                builder.append(";\n\n");
            }
        }

        for (Module module : modules) {
            for (Function function : module.functions) {
                if (!function.type.isExtern) {
                    DeclarationBuiler.writeFunctionDefine(builder, function);
                    builder.append(";\n\n");
                }
            }
        }

        for (Module module : modules) {
            for (Enumeration enumeration : module.enumerations) {
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
        }

        for (Module module : modules) {
            for (Struct struct : module.structs) {
                builder.append("struct ");
                builder.append(struct.qualified.replaceAll("\\.", "_"));
                builder.append("\n{\n");
                for (Variable variable : struct.variables) {
                    StatementBuiler.writeIndent(builder, 1);
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
        }

        for (Module module : modules) {
            for (Variable variable : module.variables) {
                TypeBuiler.writeType(builder, variable.type);
                builder.append(" ");
                builder.append(variable.qualified.replaceAll("\\.", "_"));
                if (variable.value != null) {
                    builder.append(" = ");
                    ExpressionBuiler.writeExpression(builder, variable.value);
                }
                builder.append(";\n\n");
            }
        }

        for (Module module : modules) {
            for (Function function : module.functions) {
                if (!function.type.isExtern) {
                    DeclarationBuiler.writeFunction(builder, function);
                    builder.append("\n");
                }
            }
        }

        for (Module module : modules) {
            for (Struct struct : module.structs) {
                for (Function function : struct.functions) {
                    DeclarationBuiler.writeFunction(builder, function);
                    builder.append("\n");
                }
            }
        }

        return builder;
    }
}
