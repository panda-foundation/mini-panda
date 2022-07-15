package com.github.panda_io.micro_panda.ast;

import java.util.*;

import com.github.panda_io.micro_panda.ast.declaration.Declaration;
import com.github.panda_io.micro_panda.scanner.Position;

//TO-DO Ternary Operator
//TO-DO int myNumbers[] = {25, 50, 75, 100};
//TO-DO copy(assign) struct
//TO-DO compile global offset when debug == true

public class Program {
    static class Error {
        Position position;
        String message;

        Error(Position position, String message) {
            this.position = position;
            this.message = message;
        }
    }

    Map<String, Module> modules;
    Module module;
    Map<String, Declaration> declarations; // by qualified name
    Namespace namespace;
    List<Error> errors;

    public Program() {
        this.modules = new HashMap<>();
        this.declarations = new HashMap<>();
        this.namespace = new Namespace(null, null);
        this.errors = new ArrayList<>();
    }

    public void addModule(String filename, Module module) {
        this.modules.put(filename, module);
    }

    public void setModule(Module module) {
        this.module = module;
    }

    public Collection<Module> getModules() {
        return this.modules.values();
    }

    public void addDeclaration(Declaration declaration) {
        if (this.declarations.containsKey(declaration.qualified)) {
            this.addError(declaration.getOffset(),
                    String.format("duplicated declaration with qualified name: %s", declaration.qualified));
        }
        this.declarations.put(declaration.qualified, declaration);
        this.namespace.addDeclaration(this, declaration);
    }

    public boolean hasDeclaration(String qualified) {
        return this.declarations.containsKey(qualified);
    }

    public boolean isNamespace(String ns) {
        return this.namespace.isNamespace(ns);
    }

    public void validate() {
        //TO-DO check public modifer (member access)
        //TO-DO check const (assign)
        for (Module module : this.modules.values()) {
            // TO-DO check if import is valid // must be valid, cannot import self, cannot
            // duplicated
            module.resolveType(this);
        }
        for (Module module : this.modules.values()) {
            module.validate(this);
        }
    }

    public void addError(int offset, String message) {
        this.errors.add(new Error(this.module.file.getPosition(offset), message));
    }

    public boolean hasError() {
        return this.errors.size() > 0;
    }

    public void printErrors() {
        for (Error error : this.errors) {
            System.out.printf("%s : %s \n", error.position.string(), error.message);
        }
    }

    public void printLocation(int offset) {
        System.out.println(this.module.file.getPosition(offset).string());
    }
}
