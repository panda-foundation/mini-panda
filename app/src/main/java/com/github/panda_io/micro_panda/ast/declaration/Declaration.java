package com.github.panda_io.micro_panda.ast.declaration;

import java.util.*;

import com.github.panda_io.micro_panda.ast.expression.Identifier;
import com.github.panda_io.micro_panda.ast.expression.Literal;
import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.ast.Node;

public abstract class Declaration extends Node {
    public static class Attribute {
        public int offset;
        public String name;
        public String text;
        public Map<String, Literal> values;
    }

    public List<Attribute> attributes;
    public boolean isPublic;
    public Identifier name;
    public String qualified;

    public abstract Type getType();

    public abstract void validate(Context context);

    public abstract boolean isConstant();

    public abstract void resolveType(Context context);

    public boolean hasAttribute(String name) {
        for (Attribute attribute : this.attributes) {
            if (attribute.name.equals(name)) {
                return true;
            }
        }
        return false;
    }

    public Literal getAttribute(String name, String valueName) {
        for (Attribute attribute : this.attributes) {
            if (attribute.name.equals(name)) {
                return attribute.values.get(valueName);
            }
        }
        return null;
    }
}
