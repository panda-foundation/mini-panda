package com.github.panda_io.micro_panda.ast.type;

import com.github.panda_io.micro_panda.ast.Node;
import com.github.panda_io.micro_panda.scanner.Token;

public abstract class Type extends Node {
    public abstract boolean equal(Type type);

    public abstract String string();

    public static final TypeBuiltin bool = new TypeBuiltin(Token.Bool);
    public static final TypeBuiltin u8 = new TypeBuiltin(Token.Uint8);
    public static final TypeBuiltin u16 = new TypeBuiltin(Token.Uint16);
    public static final TypeBuiltin u32 = new TypeBuiltin(Token.Uint32);
    public static final TypeBuiltin u64 = new TypeBuiltin(Token.Uint64);
    public static final TypeBuiltin i8 = new TypeBuiltin(Token.Int8);
    public static final TypeBuiltin i16 = new TypeBuiltin(Token.Int16);
    public static final TypeBuiltin i32 = new TypeBuiltin(Token.Int32);
    public static final TypeBuiltin i64 = new TypeBuiltin(Token.Int64);
    public static final TypeBuiltin f16 = new TypeBuiltin(Token.Float16);
    public static final TypeBuiltin f32 = new TypeBuiltin(Token.Float32);
    public static final TypeBuiltin f64 = new TypeBuiltin(Token.Float64);
    public static final TypePointer rawPointer = new TypePointer(u8);

    public boolean isInteger() {
        if (this instanceof TypeBuiltin) {
            return ((TypeBuiltin) this).token.isInteger();
        }
        return false;
    }

    public boolean isFloat() {
        if (this instanceof TypeBuiltin) {
            return ((TypeBuiltin) this).token.isFloat();
        }
        return false;
    }

    public boolean isNumber() {
        if (this instanceof TypeBuiltin) {
            return ((TypeBuiltin) this).token.isNumber();
        }
        return false;
    }

    public boolean isBool() {
        if (this instanceof TypeBuiltin) {
            return ((TypeBuiltin) this).token == Token.Bool;
        }
        return false;
    }

    public boolean isStruct() {
        if (this instanceof TypeName) {
            return !((TypeName) this).isEnum;
        }
        return false;
    }

    public boolean isArray() {
        if (this instanceof TypeArray) {
            return ((TypeArray) this).dimensions.get(0) != 0;
        }
        return false;
    }

    public boolean isFunction() {
        return (this instanceof TypeFunction);
    }

    public boolean isPointer() {
        if (this instanceof TypePointer) {
            return true;
        } else if (this instanceof TypeArray) {
            return ((TypeArray) this).dimensions.get(0) == 0;
        }
        return false;
    }

    public Type elementType() {
        if (this instanceof TypePointer) {
            return ((TypePointer)this).elementType;
        }
        if (this instanceof TypeArray) {
            TypeArray array = (TypeArray) this;
            if (array.dimensions.get(0) == 0) {
                if (array.dimensions.size() == 1) {
                    return array.elementType;
                } else {
                    TypeArray type = new TypeArray(array.elementType);
                    for (int i = 1; i < array.dimensions.size(); i++) {
                        type.dimensions.add(array.dimensions.get(i));
                    }
                    return type;
                }
            }
        }
        return null;
    }
}