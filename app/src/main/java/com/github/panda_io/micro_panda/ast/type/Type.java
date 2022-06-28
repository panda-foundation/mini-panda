package com.github.panda_io.micro_panda.ast.type;

import com.github.panda_io.micro_panda.ast.Node;
import com.github.panda_io.micro_panda.scanner.Token;

public abstract class Type extends Node {
    public abstract boolean equal(Type type);

    public abstract String string();

    public static final Builtin bool = new Builtin(Token.Bool);
    public static final Builtin u8 = new Builtin(Token.Uint8);
    public static final Builtin u16 = new Builtin(Token.Uint16);
    public static final Builtin u32 = new Builtin(Token.Uint32);
    public static final Builtin u64 = new Builtin(Token.Uint64);
    public static final Builtin i8 = new Builtin(Token.Int8);
    public static final Builtin i16 = new Builtin(Token.Int16);
    public static final Builtin i32 = new Builtin(Token.Int32);
    public static final Builtin i64 = new Builtin(Token.Int64);
    public static final Builtin f16 = new Builtin(Token.Float16);
    public static final Builtin f32 = new Builtin(Token.Float32);
    public static final Builtin f64 = new Builtin(Token.Float64);
    public static final Pointer rawPointer = new Pointer(u8);

    public boolean isInteger() {
        if (this instanceof Builtin) {
            return ((Builtin) this).token.isInteger();
        }
        return false;
    }

    public boolean isFloat() {
        if (this instanceof Builtin) {
            return ((Builtin) this).token.isFloat();
        }
        return false;
    }

    public boolean isNumber() {
        if (this instanceof Builtin) {
            return ((Builtin) this).token.isNumber();
        }
        return false;
    }

    public boolean isBool() {
        if (this instanceof Builtin) {
            return ((Builtin) this).token == Token.Bool;
        }
        return false;
    }

    public boolean isStruct() {
        if (this instanceof Name) {
            return !((Name) this).isEnum;
        }
        return false;
    }

    public boolean isArray() {
        if (this instanceof Array) {
            return ((Array) this).dimensions.get(0) != 0;
        }
        return false;
    }

    public boolean isFunction() {
        return (this instanceof Function);
    }

    public boolean isPointer() {
        if (this instanceof Pointer) {
            return true;
        } else if (this instanceof Array) {
            return ((Array) this).dimensions.get(0) == 0;
        }
        return false;
    }

    public Type elementType() {
        if (this instanceof Pointer) {
            return ((Pointer)this).elementType;
        }
        if (this instanceof Array) {
            Array array = (Array) this;
            if (array.dimensions.get(0) == 0) {
                if (array.dimensions.size() == 1) {
                    return array.elementType;
                } else {
                    Array type = new Array(array.elementType);
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