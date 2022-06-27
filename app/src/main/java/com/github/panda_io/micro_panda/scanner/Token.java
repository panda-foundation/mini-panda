package com.github.panda_io.micro_panda.scanner;

import java.util.HashMap;

public enum Token {
    ILLEGAL("ILLEGAL"),
    EOF("EOF"),
    META("META"),

    IDENT("identifier"),
    BOOL("bool_literal"),
    CHAR("char_literal"),
    INT("int_literal"),
    FLOAT("float_literal"),
    STRING("string_literal"),
    NULL("null"),

    Break("break"),
    Case("case"),
    Const("const"),
    Continue("continue"),
    Default("default"),
    Else("else"),
    Enum("enum"),
    For("for"),
    Function("function"),
    If("if"),
    Import("import"),
    Namespace("namespace"),
    Pointer("pointer"),
    Public("public"),
    Sizeof("sizeof"),
    Return("return"),
    Struct("struct"),
    Switch("switch"),
    This("this"),
    Var("var"),

    Bool("bool"),
    Int8("i8"),
    Int16("i16"),
    Int32("i32"),
    Int64("i64"),
    Uint8("u8"),
    Uint16("u16"),
    Uint32("u32"),
    Uint64("u64"),
    Float16("f16"),
    Float32("f32"),
    Float64("f64"),
    Void("void"),

    LeftParen("("),
    RightParen(")"),
    LeftBracket("["),
    RightBracket("]"),
    LeftBrace("{"),
    RightBrace("}"),
    Plus("+"),
    Minus("-"),
    Mul("*"),
    Div("/"),
    Rem("%"),
    BitAnd("&"),
    BitOr("|"),
    BitXor("^"),
    Complement("~"),
    Not("!"),
    Assign("="),
    Less("<"),
    Greater(">"),
    PlusAssign("+="),
    MinusAssign("-="),
    MulAssign("*="),
    DivAssign("/="),
    RemAssign("%="),
    XorAssign("^="),
    AndAssign("&="),
    OrAssign("|="),
    LeftShift("<<"),
    RightShift(">>"),
    LeftShiftAssign("<<="),
    RightShiftAssign(">>="),
    Equal("=="),
    NotEqual("!="),
    LessEqual("<="),
    GreaterEqual(">="),
    And("&&"),
    Or("||"),
    PlusPlus("++"),
    MinusMinus("--"),
    Comma(","),
    Colon(":"),
    Semi(";"),
    Dot(".");

    private static final HashMap<String, Token> tokens = new HashMap<>();
    static {
        for (Token t : values()) {
            tokens.put(t.token, t);
        }
    }

    public static Token valueOfToken(String token) {
        return tokens.get(token);
    }

    public static Token readToken(String literal) {
        if (tokens.containsKey(literal)) {
            return tokens.get(literal);
        }
        if (literal.equals("true") || literal.equals("false")) {
            return BOOL;
        }
        if (literal.equals("null")) {
            return NULL;
        }
        return IDENT;
    }

    final String token;

    Token(String token) {
        this.token = token;
    }

    @Override
    public String toString() {
        return this.token;
    }

    public int Precedence() {
        switch (this) {
            case Assign:
            case MulAssign:
            case DivAssign:
            case RemAssign:
            case PlusAssign:
            case MinusAssign:
            case LeftShiftAssign:
            case RightShiftAssign:
            case AndAssign:
            case OrAssign:
            case XorAssign:
                return 1;

            case Or:
                return 2;

            case And:
                return 3;

            case BitOr:
                return 4;

            case BitXor:
                return 5;

            case BitAnd:
                return 6;

            case Equal:
            case NotEqual:
                return 7;

            case Less:
            case LessEqual:
            case Greater:
            case GreaterEqual:
                return 8;

            case LeftShift:
            case RightShift:
                return 9;

            case Plus:
            case Minus:
                return 10;

            case Mul:
            case Div:
            case Rem:
                return 11;

            default:
                return 0;
        }
    }

    public static boolean IsLiteral(Token token) {
        return token.compareTo(IDENT) >= 0 && token.compareTo(NULL) <= 0;
    }

    public static boolean IsOperator(Token token) {
        return token.compareTo(LeftParen) >= 0 && token.compareTo(Dot) <= 0;
    }

    public static boolean IsKeyword(Token token) {
        return token.compareTo(Break) >= 0 && token.compareTo(Var) <= 0;
    }

    public static boolean IsScalar(Token token) {
        return token.compareTo(Bool) >= 0 && token.compareTo(Void) <= 0;
    }

    public static boolean IsInteger(Token token) {
        return token.compareTo(Int8) >= 0 && token.compareTo(Uint64) <= 0;
    }

    public static boolean IsFloat(Token token) {
        return token.compareTo(Float16) >= 0 && token.compareTo(Float64) <= 0;
    }

    public static boolean IsNumber(Token token) {
        return token.compareTo(Int8) >= 0 && token.compareTo(Float64) <= 0;
    }

    public static boolean IsAssign(Token token) {
        return token.compareTo(Assign) >= 0 && token.compareTo(RightShiftAssign) <= 0;
    }
}
