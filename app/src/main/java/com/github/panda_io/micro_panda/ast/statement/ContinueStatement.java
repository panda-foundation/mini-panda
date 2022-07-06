package com.github.panda_io.micro_panda.ast.statement;

import com.github.panda_io.micro_panda.ast.Context;

public class ContinueStatement extends Statement {
    public void validate(Context context) {
        /*
         * if c.LoopBlock == nil {
         * //TO-DO add check
         * //c.Program.Error(con.Position, "invalid continue")
         * }
         */
    }
}
