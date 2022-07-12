package com.github.panda_io.micro_panda.ast.expression;

import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.ast.declaration.Declaration;
import com.github.panda_io.micro_panda.ast.declaration.Struct;

public class Identifier extends Expression {
    public String name;
    public String qualified;
    public boolean isNamespace; // part of namespace

    public void validate(Context context, Type expected) {
        Type type = context.findObject(this.name);
        if (type == null) {
            Declaration declaration = context.findLocalDeclaration(this.name);
            if (declaration == null) {
                this.isNamespace = context.isNamespace(this.name);
            } else if (!(declaration instanceof Struct)) {
                this.constant = declaration.isConstant();
                this.type = declaration.getType();
                this.qualified = declaration.qualified;
            }
        } else {
            this.constant = false;
            this.type = type;
            if (expected != null && !this.type.equal(expected)) {
                context.addError(this.getOffset(),
                        String.format("mismatch type, expect %s, got %s", expected.string(), this.type.string()));
            }
        }
        // type is null for enum (enum member has type u8)
        // type is null when identifier is namespacee (or part of namespace)
        if (this.type == null && this.qualified == null && !this.isNamespace) {
            context.addError(this.getOffset(), String.format("undefined: %s", this.name));
        }
    }
}
