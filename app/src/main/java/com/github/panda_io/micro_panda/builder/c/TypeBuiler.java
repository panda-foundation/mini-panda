package com.github.panda_io.micro_panda.builder.c;

import java.util.HashMap;
import java.util.Map;

import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.type.TypeArray;
import com.github.panda_io.micro_panda.ast.type.TypeBuiltin;
import com.github.panda_io.micro_panda.ast.type.TypeFunction;
import com.github.panda_io.micro_panda.ast.type.TypeName;
import com.github.panda_io.micro_panda.ast.type.TypePointer;
import com.github.panda_io.micro_panda.scanner.Token;

public class TypeBuiler {

    static Map<Token, String> builtinTypes;
    static {
        builtinTypes = new HashMap<>();
        builtinTypes.put(Token.Bool, "uint8_t");
        builtinTypes.put(Token.Int8, "int8_t");
        builtinTypes.put(Token.Uint8, "uint8_t");
        builtinTypes.put(Token.Int16, "int16_t");
        builtinTypes.put(Token.Uint16, "uint16_t");
        builtinTypes.put(Token.Int32, "int32_t");
        builtinTypes.put(Token.Uint32, "uint32_t");
        builtinTypes.put(Token.Int64, "int64_t");
        builtinTypes.put(Token.Uint64, "uint64_t");
        builtinTypes.put(Token.Float16, "float");
        builtinTypes.put(Token.Float32, "float");
        builtinTypes.put(Token.Float64, "double");
        builtinTypes.put(Token.Void, "void");
    }

    static void writeType(StringBuilder builder, Type type) {
        if (type == null) {
            builder.append("void");
        } else if (type instanceof TypeBuiltin) {
            TypeBuiltin builtin = (TypeBuiltin) type;
            builder.append(builtinTypes.get(builtin.token));
        } else if (type instanceof TypeName) {
            builder.append(((TypeName) type).qualified.replaceAll("\\.", "_"));
        } else if (type instanceof TypeFunction) {
            // TO-DO
        } else if (type instanceof TypePointer) {
            writeType(builder, ((TypePointer) type).elementType);
            builder.append("*");
        } else if (type instanceof TypeArray) {
            TypeArray array = (TypeArray)type;
            if (array.dimensions.get(0) == 0 && array.dimensions.size() > 1) {
                builder.append("(");
            }
            writeType(builder, array.elementType);
            if (array.dimensions.get(0) == 0) {
                builder.append("*");
                if (array.dimensions.size() == 0) {
                    builder.append(")");
                }
            }
            //TO-DO check return pointer of multi dimension array
        }
    }

    static void writeArrayIndex(StringBuilder builder, TypeArray array) {
        if (array.dimensions.get(0) > 0) {
            builder.append(String.format("[%d]", array.dimensions.get(0)));
        }
        for (int i = 1; i < array.dimensions.size(); i++) {
            builder.append(String.format("[%d]", array.dimensions.get(i)));
        }
    }
}
