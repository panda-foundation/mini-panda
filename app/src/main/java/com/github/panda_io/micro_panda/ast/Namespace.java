package com.github.panda_io.micro_panda.ast;

import java.util.*;

import com.github.panda_io.micro_panda.ast.declaration.*;

public class Namespace {
    String name;
    String qualified;
    Map<String, Declaration> declarations;
    Map<String, Namespace> children;

    public Namespace(String name, String qualified) {
        this.name = name;
        this.qualified = qualified;
        this.declarations = new HashMap<>();
        this.children = new HashMap<>();
    }

    public void addDeclaration(Program program, Declaration declaration) {
        String[] names = declaration.qualified.split("\\.");
        Namespace namespace = this;
        for (int i = 0; i < names.length; i++) {
            String name = names[i];
            if (i == names.length - 1) {
                if (namespace.children.get(name) != null) {
                    program.addError(declaration.getOffset(),
                            String.format("declaration qualified name conflict with existing namespace '%s'",
                                    namespace.children.get(name).qualified));
                    return;
                }
                namespace.declarations.put(name, declaration);
            } else {
                if (namespace.declarations.get(name) != null) {
                    program.addError(declaration.getOffset(),
                            String.format("namespace conflict with existing declaration's qualified name '%s'",
                                    namespace.declarations.get(name).qualified));
                    return;
                }
                if (!namespace.children.containsKey(name)) {
                    String qualified = name;
                    if (namespace.qualified != null) {
                        qualified = String.format("%s.%s", namespace.qualified, name);
                    }
                    namespace.children.put(name,
                            new Namespace(name, qualified));
                }
                namespace = namespace.children.get(name);
            }
        }
    }

    public boolean isNamespace(String ns) {
        String[] names = ns.split("\\.");
        Namespace namespace = this;
        for (int i = 0; i < names.length; i++) {
            String name = names[i];
            if (i == names.length - 1) {
                return namespace.children.get(name) != null;
            } else {
                if (!namespace.children.containsKey(name)) {
                    return false;
                }
                namespace = namespace.children.get(name);
            }
        }
        return false;
    }
}
