package com.github.panda_io.micro_panda.ast.statement;

import com.github.panda_io.micro_panda.ast.Context;

public class Break extends Statement {
    public void validate(Context context) {
        /*
         * if c.LeaveBlock == nil {
         * //TO-DO add check
         * //c.Program.Error(b.Position, "invalid break")
         * }
         */
    }
}
