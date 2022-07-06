package com.github.panda_io.micro_panda.ast.type;

import org.junit.Test;
import static org.junit.Assert.*;
import com.github.panda_io.micro_panda.scanner.Token;

public class TypeTest {
    @Test
    public void testBuiltinEqual() {
        TypeBuiltin type0 = new TypeBuiltin(Token.Bool);
        TypeBuiltin type1 = new TypeBuiltin(Token.Bool);
        TypeBuiltin type2 = new TypeBuiltin(Token.Uint8);
        TypeBuiltin type3 = new TypeBuiltin(Token.Uint8);
        TypeBuiltin type4 = new TypeBuiltin(Token.Float64);
        TypeBuiltin type5 = new TypeBuiltin(Token.Float64);

        assertTrue(type0.equal(type1));
        assertTrue(type2.equal(type3));
        assertTrue(type4.equal(type5));
        assertFalse(type1.equal(type2));
        assertFalse(type3.equal(type4));
        assertFalse(type5.equal(type0));
    }

    @Test
    public void testBuiltinString() {
        TypeBuiltin type0 = new TypeBuiltin(Token.Bool);
        assertTrue(type0.string().equals("bool"));

        TypeBuiltin type1 = new TypeBuiltin(Token.Float16);
        assertTrue(type1.string().equals("f16"));

        TypeBuiltin type2 = new TypeBuiltin(Token.Void);
        assertTrue(type2.string().equals("void"));
    }
}
