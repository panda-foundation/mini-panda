package com.github.panda_io.micro_panda.ast.type;

import java.util.ArrayList;

public class Function extends Type {
    public Type returnType;
    public ArrayList<Type> parameters;

    public boolean isMemberFunction;
    public boolean isExtern;
    public String externName;
    public boolean isDefine;

    public Function() {
        this.parameters = new ArrayList<>();
    }

    public boolean equal(Type type) {
        if (type instanceof Function) {
            Function function = (Function) type;
            if (this.returnType != null && function.returnType != null) {
                if (!this.returnType.equal(function.returnType)) {
                    return false;
                }
            } else if (this.returnType != null || function.returnType != null) {
                return false;
            }
            if (this.parameters.size() != function.parameters.size()) {
                return false;
            }
            for (int i = 0; i < this.parameters.size(); i++) {
                if (!this.parameters.get(i).equal(function.parameters.get(i))) {
                    return false;
                }
            }
            return true;
        }
        return false;
    }

    public String string() {
        StringBuffer buffer = new StringBuffer();
        buffer.append("function(");
        for (int i = 0; i < this.parameters.size(); i++) {
            if (i != 0) {
                buffer.append(", ");
            }
            buffer.append(this.parameters.get(i).string());
        }
        return buffer.toString();
    }
}