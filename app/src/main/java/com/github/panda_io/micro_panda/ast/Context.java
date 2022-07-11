package com.github.panda_io.micro_panda.ast;

import java.util.*;

import com.github.panda_io.micro_panda.ast.type.*;
import com.github.panda_io.micro_panda.ast.declaration.*;
import com.github.panda_io.micro_panda.ast.declaration.Enumeration;

public class Context {
    Program  program;
    Function function;
	Context parent;
	Map<String, Type> objects;

    public int loopLevel;

    public Context(Program program) {
        this.program = program;
        this.objects = new HashMap<>();
        this.loopLevel = 0;
    }

    public Context newContext() {
        Context context = new Context(this.program);
        context.function = this.function;
        context.loopLevel = this.loopLevel;
        context.parent = this;
        return context;
    }

	public void addError(int offset, String message) {
        this.program.addError(offset, message);
    }

	public boolean insertObject(String name, Type type) {
        if (this.objects.containsKey(name)) {
            return false;
        }
        this.objects.put(name, type);
        return true;
    }

    public Type findObject(String name) {
        if (this.objects.containsKey(name)) {
            return this.objects.get(name);
        }
        if (this.parent != null) {
            Type object = this.parent.findObject(name);
            if (object != null) {
                return object;
            }
        }
        if (this.function != null && this.function.parent != null) {
            return this.function.parent.memberType(name);
        }
        return null;
    }

    public Type resolveType(Type type) {
        if (type instanceof TypeName) {
            Declaration declaration = this.findDeclaration(type);
            if (declaration == null) {
                this.addError(type.getOffset(), "type not defined");
            } else {
                if (declaration instanceof Function) {
                    return ((Function)declaration).type;
                } else if (declaration instanceof Struct) {
                } else if (declaration instanceof Enumeration) {
                    ((TypeName)type).isEnum = true;
                } else {
                    this.addError(type.getOffset(), "type not defined");
                }
            }
        } else if (type instanceof TypeArray) {
            TypeArray array = (TypeArray)type;
            array.elementType = this.resolveType(array.elementType);
            if (array.dimensions.size() == 0 || array.dimensions.get(0) < 0) {
                this.addError(type.getOffset(), "invalid array index");
            }
            for (int dimension : array.dimensions) {
                if (dimension < 1) {
                    this.addError(type.getOffset(), "invalid array index");
                }
            }
        } else if (type instanceof TypePointer) {
            TypePointer pointer = (TypePointer)type;
            pointer.elementType = this.resolveType(pointer.elementType);
        } else if (type instanceof TypeFunction) {
            TypeFunction function = (TypeFunction)type;
            function.returnType = this.resolveType(function.returnType);
            for (int i = 0; i < function.parameters.size(); i++) {
                Type parameter = this.resolveType(function.parameters.get(i));
                function.parameters.set(i, parameter);
                if (parameter.isStruct()) {
                    this.addError(parameter.getOffset(), "struct is not allowed as parameter, use pointer instead");
                } else if (parameter.isArrayWithSize()) {
                    this.addError(parameter.getOffset(), "array is not allowed as parameter, use pointer instead");
                }
            }
        }
        return type;
    }

	public Declaration findDeclaration(Type type) {
        TypeName name = (TypeName)type;
        if (name.qualified == null) {
            return this.findLocalDeclaration(name.name);
        }
        Declaration declaration = this.findQualifiedDeclaration(name.qualified);
        if (declaration instanceof Enumeration) {
            name.isEnum = true;
        }
        name.qualified = declaration.qualified;
        return declaration;
    }

	public Declaration findLocalDeclaration(String name) {
        String qualified = String.format("%s.%s", this.program.module.namespace, name);
        if (this.program.declarations.containsKey(qualified)) {
            return this.program.declarations.get(qualified);
        }
        qualified = String.format("%s.%s", Constant.global, name);
        if (this.program.declarations.containsKey(qualified)) {
            return this.program.declarations.get(qualified);
        }
        for (Module.Using imported : this.program.module.usings) {
            qualified = String.format("%s.%s", imported.namespace, name);
            if (this.program.declarations.containsKey(qualified)) {
                return this.program.declarations.get(qualified);
            }
        }
        return null;
    }

	public Declaration findQualifiedDeclaration(String qualified) {
    	return this.program.declarations.get(qualified);
    }

	public boolean isNamespace(String ns) {
        return this.program.isNamespace(ns);
    }

    public void setFunction(Function function) {
        this.function = function;
    }

    public Function getFunction() {
        return this.function;
    }
}
