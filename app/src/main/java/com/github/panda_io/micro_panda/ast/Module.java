package com.github.panda_io.micro_panda.ast;

import java.util.List;
import java.util.ArrayList;

import com.github.panda_io.micro_panda.ast.declaration.*;
import com.github.panda_io.micro_panda.ast.declaration.Enumeration;
import com.github.panda_io.micro_panda.scanner.*;

public class Module {
    public static class Import extends Node {
        public String namespace;
    }

    public File file;
    public String namespace;
    public List<Import> imports;
    public List<Declaration.Attribute> attributes;
    public List<Variable> variables;
    public List<Function> functions;
    public List<Enumeration> enumerations;
    public List<Struct> structs;

    public Module() {
        this.imports = new ArrayList<>();
        this.attributes = new ArrayList<>();
        this.variables = new ArrayList<>();
        this.functions = new ArrayList<>();
        this.enumerations = new ArrayList<>();
        this.structs = new ArrayList<>();
    }

    public void resolveType(Program program) {
        program.module = this;
        Context context = new Context(program);

        for (Variable variable : this.variables) {
            variable.resolveType(context);
        }
        for (Function function : this.functions) {
            function.resolveType(context);
        }
        for (Enumeration enumeration : this.enumerations) {
            enumeration.resolveType(context);
        }
        for (Struct struct : this.structs) {
            struct.resolveType(context);
        }
    }

    public void validate(Program program) {
        program.module = this;
        Context context = new Context(program);

        for (Variable variable : this.variables) {
            variable.validate(context);
        }
        for (Function function : this.functions) {
            function.validate(context);
        }
        for (Enumeration enumeration : this.enumerations) {
            enumeration.validate(context);
        }
        for (Struct struct : this.structs) {
            struct.validate(context);
        }
    }
}
