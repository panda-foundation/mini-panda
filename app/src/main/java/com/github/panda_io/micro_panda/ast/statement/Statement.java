package com.github.panda_io.micro_panda.ast.statement;

import com.github.panda_io.micro_panda.ast.*;

public abstract class Statement extends Node {
    public abstract void validate(Context context);
}
