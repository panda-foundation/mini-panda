package com.github.panda_io.micro_panda.ast.type;

import org.junit.Test;
import static org.junit.Assert.*;
import com.github.panda_io.micro_panda.scanner.Token;

public class TypeTest {
    @Test
    public void testBuiltinEqual() {
        Builtin type0 = new Builtin(Token.Bool);
        Builtin type1 = new Builtin(Token.Bool);
        Builtin type2 = new Builtin(Token.Uint8);
        Builtin type3 = new Builtin(Token.Uint8);
        Builtin type4 = new Builtin(Token.Float64);
        Builtin type5 = new Builtin(Token.Float64);

        assertTrue(type0.equal(type1));
        assertTrue(type2.equal(type3));
        assertTrue(type4.equal(type5));
        assertFalse(type1.equal(type2));
        assertFalse(type3.equal(type4));
        assertFalse(type5.equal(type0));
    }

    @Test
    public void testBuiltinString() {
        Builtin type0 = new Builtin(Token.Bool);
        assertTrue(type0.string().equals("bool"));

        Builtin type1 = new Builtin(Token.Float16);
        assertTrue(type1.string().equals("f16"));

        Builtin type2 = new Builtin(Token.Void);
        assertTrue(type2.string().equals("void"));
    }
}
