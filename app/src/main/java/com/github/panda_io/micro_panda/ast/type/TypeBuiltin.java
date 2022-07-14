package com.github.panda_io.micro_panda.ast.type;

import com.github.panda_io.micro_panda.scanner.Token;

public class TypeBuiltin extends Type {
    public Token token;

    public TypeBuiltin(Token token) {
        this.token = token;
    }

    public boolean equal(Type type) {
        if (type == null)
            return false;
        if (type instanceof TypeBuiltin) {
            return this.token == ((TypeBuiltin) type).token;
        }
        return false;
    }

    public String string() {
        return this.token.toString();
    }

    /*
    public int bits() {
        switch (this.token) {
            case Bool:
                return 1;

            case Int8:
            case Uint8:
                return 8;

            case Int16:
            case Uint16:
            case Float16:
                return 16;

            case Int32:
            case Uint32:
            case Float32:
                return 32;

            case Int64:
            case Uint64:
            case Float64:
                return 64;

            default:
                return 0;
        }
    }

    public int size() {
        switch (this.token) {
            case Bool:
                return 1;

            case Int8:
            case Uint8:
                return 1;

            case Int16:
            case Uint16:
            case Float16:
                return 2;

            case Int32:
            case Uint32:
            case Float32:
                return 4;

            case Int64:
            case Uint64:
            case Float64:
                return 8;

            default:
                return 0;
        }
    }*/
}