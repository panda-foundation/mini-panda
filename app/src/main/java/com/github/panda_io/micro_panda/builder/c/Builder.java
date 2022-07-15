package com.github.panda_io.micro_panda.builder.c;

import java.util.Collection;

import com.github.panda_io.micro_panda.ast.Module;
import com.github.panda_io.micro_panda.ast.Program;
import com.github.panda_io.micro_panda.ast.declaration.Function;
import com.github.panda_io.micro_panda.ast.declaration.Struct;

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

        //TO-DO enum define

        //TO-DO struct define

        //TO-DO variable & init

        for (Module module : modules) {
            for (Function function : module.functions) {
                if (!function.type.isExtern) {
                    DeclarationBuiler.writeFunction(builder, function);
                    builder.append("\n");
                }
            }
        }

        return builder;
    }
}
