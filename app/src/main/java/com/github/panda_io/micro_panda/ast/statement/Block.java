package com.github.panda_io.micro_panda.ast.statement;

import java.util.*;

import com.github.panda_io.micro_panda.ast.Context;

public class Block extends Statement {
    public List<Statement> statements;
    
    //TO-DO warning: unreachable code //Start, End of block
    public void validate(Context context) {
        context = context.newContext();
        for (Statement statement:this.statements) {
            statement.validate(context);
        }
    }
}
