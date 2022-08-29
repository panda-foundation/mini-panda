package com.github.panda_io.micro_panda.scanner;

import java.util.*;

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
    Continue("continue"),
    Default("default"),
    Else("else"),
    Enum("enum"),
    For("for"),
    Function("function"),
    If("if"),
    Namespace("namespace"),
    Pointer("pointer"),
    Public("public"),
    Return("return"),
    Struct("struct"),
    Switch("switch"),
    This("this"),
    Using("using"),
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

    private static final Map<String, Token> tokens = new HashMap<>();
    static {
        for (Token token : values()) {
            tokens.put(token.token, token);
        }

        OperatorNode.root = new OperatorNode();
        for (int i = LeftParen.ordinal(); i <= Dot.ordinal(); i++) {
            OperatorNode.root.insert(values()[i].toString());
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

    public int precedence() {
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

    public boolean isLiteral() {
        return this.compareTo(IDENT) >= 0 && this.compareTo(NULL) <= 0;
    }

    public boolean isOperator() {
        return this.compareTo(LeftParen) >= 0 && this.compareTo(Dot) <= 0;
    }

    public boolean isKeyword() {
        return this.compareTo(Break) >= 0 && this.compareTo(Var) <= 0;
    }

    public boolean isScalar() {
        return this.compareTo(Bool) >= 0 && this.compareTo(Void) <= 0;
    }

    public boolean isInteger() {
        return this.compareTo(Int8) >= 0 && this.compareTo(Uint64) <= 0;
    }

    public boolean isFloat() {
        return this.compareTo(Float16) >= 0 && this.compareTo(Float64) <= 0;
    }

    public boolean isNumber() {
        return this.compareTo(Int8) >= 0 && this.compareTo(Float64) <= 0;
    }

    public boolean isAssign() {
        return this.compareTo(Assign) >= 0 && this.compareTo(RightShiftAssign) <= 0;
    }
}
