package com.github.panda_io.micro_panda.builder.c;

import java.util.Collection;

import com.github.panda_io.micro_panda.ast.Module;
import com.github.panda_io.micro_panda.ast.Program;
import com.github.panda_io.micro_panda.ast.declaration.*;

public class Builder {

    public static StringBuilder build(Program program) {
        StringBuilder builder = new StringBuilder();
        builder.append("#include <stdint.h>\n#include <stdio.h>\n\n");
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
                if (!function.type.isExtern && function.type.isDefine) {
                    DeclarationBuiler.writeFunctionDefine(builder, function);
                    builder.append(";\n\n");
                }
            }
        }

        for (Module module : modules) {
            for (Function function : module.functions) {
                if (!function.type.isExtern && !function.type.isDefine) {
                    DeclarationBuiler.writeFunctionDefine(builder, function);
                    builder.append(";\n\n");
                }
            }
        }

        for (Module module : modules) {
            for (Enumeration enumeration : module.enumerations) {
                DeclarationBuiler.writeEnumDefine(builder, enumeration);
            }
        }

        for (Module module : modules) {
            for (Struct struct : module.structs) {
                DeclarationBuiler.writeStructDefine(builder, struct);
            }
        }

        for (Module module : modules) {
            for (Variable variable : module.variables) {
                DeclarationBuiler.writeVariable(builder, variable);
            }
        }

        for (Module module : modules) {
            for (Function function : module.functions) {
                if (!(function.type.isExtern || function.type.isDefine)) {
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
