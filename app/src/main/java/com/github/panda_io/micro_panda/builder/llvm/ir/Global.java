package com.github.panda_io.micro_panda.builder.llvm.ir;

import com.github.panda_io.micro_panda.builder.llvm.ir.constant.Constant;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Pointer;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Global extends Constant {
    Identifier _identifier;
    boolean immutable;
    Type contentType;
    Constant initializer;
    Type type;

    public Global(String name, Constant initializer) {
        this._identifier = new Identifier(true);
        this._identifier.setName(name);
        this.initializer = initializer;
        this.contentType = this.initializer.getType();
        this.getType();
    }

    public String string() {
        return String.format("%s %s", this.type.string(), this.identifier());
    }

    public Type getType() {
        if (this.type == null) {
            this.type = new Pointer(this.contentType);
        }
        return this.type;
    }

    public String identifier() {
        return this._identifier.identifier();
    }

    public void writeIR(StringBuilder builder) {
        String declaration = "global";
        if (this.immutable) {
            declaration = "constant";
        }
        builder.append(String.format("%s = %s %s %s", this.identifier(), declaration, this.contentType.string(),
                this.initializer.identifier()));
    }
}