package com.github.panda_io.micro_panda.ast;

import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.declaration.Declaration;

public class Context {
    public Context newContext() {
        return null;
    }

	public void addError(int offset, String message) {

    }

	public void insertObject(String name, Type type) {

    }

    public Type findObject(String name) {
        return null;
    }

    public Type resolveType(Type type) {
        return null;
    }

	public Declaration findDeclaration(Type type) {
        return null;
    }

	public Declaration findLocalDeclaration(String name) {
        return null;
    }

	public Declaration findQualifiedDeclaration(String qualified) {
        return null;
    }

	public boolean isNamespace(String namespace) {
        return false;
    }
}
