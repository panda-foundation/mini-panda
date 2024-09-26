namespace MicroPanda.Token;

using System.Collections.Generic;
using System.Text;

internal enum Token
{
	ILLEGAL,
	EOF,
	COMMENT,
	ANNOTATION,
	NEWLINE,

	// literals
	_literalBegin,
	IDENT,
	BOOL,
	CHAR,
	INT,
	FLOAT,
	STRING,
	NULL,
	_literalEnd,

	// keywords
	_keywordBegin,
	Break,
	Case,
	Const,
	Continue,
	Default,
	Else,
	Enum,
	For,
	Function,
	If,
	Namespace,
	Pointer,
	Public,
	Return,
	Sizeof,
	Struct,
	Switch,
	This,
	Using,
	Var,
	_keywordEnd,

	// scalar
	_scalarBegin,
	Bool,
	_numberBegin,
	_integerBegin,
	Int8,
	Int16,
	Int32,
	Int64,
	Uint8,
	Uint16,
	Uint32,
	Uint64,
	_integerEnd,
	_floatBegin,
	integerEnd,
	floatBegin,
	Float16,
	Float32,
	Float64,
	_floatEnd,
	_numberEnd,
	Void,
	_scalarEnd,

	// operators
	_operatorBegin,
	LeftParen,
	RightParen,
	LeftBracket,
	RightBracket,
	LeftBrace,
	RightBrace,

	Plus,
	Minus,
	Mul,
	Div,
	Less,
	Greater,
	Rem,
	BitAnd,
	BitOr,
	BitXor,
	Complement,
	Not,
	LeftShift,
	RightShift,

	_assignBegin,
	Assign,
	PlusAssign,
	MinusAssign,
	MulAssign,
	DivAssign,
	RemAssign,
	XorAssign,
	AndAssign,
	OrAssign,
	LeftShiftAssign,
	RightShiftAssign,
	_assignEnd,

	Equal,
	NotEqual,
	LessEqual,
	GreaterEqual,
	And,
	Or,
	PlusPlus,
	MinusMinus,

	Comma,
	Colon,
	Semi,
	Dot,
	_operatorEnd
}

internal static class TokenHelper
{
	internal static bool IsLiteral(Token token) => Token._literalBegin < token && token < Token._literalEnd;

	internal static bool IsOperator(Token token) => Token._operatorBegin < token && token < Token._operatorEnd;

	internal static bool IsKeyword(Token token) => Token._keywordBegin < token && token < Token._keywordEnd;

	internal static bool IsScalar(Token token) => Token._scalarBegin < token && token < Token._scalarEnd;

	internal static bool IsInteger(Token token) => Token._integerBegin < token && token < Token._integerEnd;

	internal static bool IsFloat(Token token) => Token._floatBegin < token && token < Token._floatEnd;

	internal static bool IsNumber(Token token) => Token._numberBegin < token && token < Token._numberEnd;

	internal static bool IsAssign(Token token) => Token._assignBegin < token && token < Token._assignEnd;

	internal static bool IsOperator(string literal) => _operations.Contains(literal);

	internal static Token FromString(string literal)
	{
		if (_string2Token.TryGetValue(literal, out var token))
		{
			return token;
		}
		else if (literal == "true" || literal == "false")
		{
			return Token.BOOL;
		}
		else if (literal == "null")
		{
			return Token.NULL;
		}
		return Token.IDENT;
	}

	internal static string ToString(Token token)
	{
		return _token2String[token];
	}

	internal static int Precedence(Token token)
	{
        return token switch
        {
            Token.Assign or Token.MulAssign or Token.DivAssign or Token.RemAssign or Token.PlusAssign or Token.MinusAssign or Token.LeftShiftAssign or Token.RightShiftAssign or Token.AndAssign or Token.OrAssign or Token.XorAssign => 1,
            Token.Or => 2,
            Token.And => 3,
            Token.BitOr => 4,
            Token.BitXor => 5,
            Token.BitAnd => 6,
            Token.Equal or Token.NotEqual => 7,
            Token.Less or Token.LessEqual or Token.Greater or Token.GreaterEqual => 8,
            Token.LeftShift or Token.RightShift => 9,
            Token.Plus or Token.Minus => 10,
            Token.Mul or Token.Div or Token.Rem => 11,
            _ => 0,
        };
    }

	static TokenHelper()
	{
		foreach (var pair in _token2String)
		{
			_string2Token[pair.Value] = pair.Key;
		}

		for (var i = Token._operatorBegin + 1; i < Token._operatorEnd; i++)
		{
			if (i != Token._assignBegin && i != Token._assignEnd)
			{
				_operations.Add(_token2String[i]);
			}
		}
	}

	private static readonly Dictionary<Token, string> _token2String = new()
	{
		{ Token.ILLEGAL, "ILLEGAL" },
		{ Token.EOF, "EOF" },
		{ Token.COMMENT, "COMMENT" },
		{ Token.ANNOTATION, "ANNOTATION" },
		{ Token.NEWLINE, "NEWLINE" },

		{ Token.IDENT, "identifier" },
		{ Token.BOOL, "bool_literal" },
		{ Token.CHAR, "char_literal" },
		{ Token.INT, "int_literal" },
		{ Token.FLOAT, "float_literal" },
		{ Token.STRING, "string_literal" },
		{ Token.NULL, "null" },

		{ Token.Break, "break" },
		{ Token.Case, "case" },
		{ Token.Const, "const" },
		{ Token.Continue, "continue" },
		{ Token.Default, "default" },
		{ Token.Else, "else" },
		{ Token.Enum, "enum" },
		{ Token.For, "for" },
		{ Token.Function, "function" },
		{ Token.If, "if" },
		{ Token.Namespace, "namespace" },
		{ Token.Pointer, "pointer" },
		{ Token.Public, "public" },
		{ Token.Return, "return" },
		{ Token.Sizeof, "sizeof" },
		{ Token.Struct, "struct" },
		{ Token.Switch, "switch" },
		{ Token.This, "this" },
		{ Token.Using, "using" },
		{ Token.Var, "var" },

		{ Token.Bool, "bool" },
		{ Token.Int8, "i8" },
		{ Token.Int16, "i16" },
		{ Token.Int32, "i32" },
		{ Token.Int64, "i64" },
		{ Token.Uint8, "u8" },
		{ Token.Uint16, "u16" },
		{ Token.Uint32, "u32" },
		{ Token.Uint64, "u64" },
		{ Token.Float16, "f16" },
		{ Token.Float32, "f32" },
		{ Token.Float64, "f64" },
		{ Token.Void, "void" },

		{ Token.LeftParen, "(" },
		{ Token.RightParen, ")" },
		{ Token.LeftBracket, "[" },
		{ Token.RightBracket, "]" },
		{ Token.LeftBrace, "{" },
		{ Token.RightBrace, "}" },

		{ Token.Plus, "+" },
		{ Token.Minus, "-" },
		{ Token.Mul, "*" },
		{ Token.Div, "/" },
		{ Token.Less, "<" },
		{ Token.Greater, ">" },
		{ Token.Rem, "%" },
		{ Token.BitAnd, "&" },
		{ Token.BitOr, "|" },
		{ Token.BitXor, "^" },
		{ Token.Complement, "~" },
		{ Token.Not, "!" },
		{ Token.LeftShift, "<<" },
		{ Token.RightShift, ">>" },

		{ Token.Assign, "=" },
		{ Token.PlusAssign, "+=" },
		{ Token.MinusAssign, "-=" },
		{ Token.MulAssign, "*=" },
		{ Token.DivAssign, "/=" },
		{ Token.RemAssign, "%=" },
		{ Token.XorAssign, "^=" },
		{ Token.AndAssign, "&=" },
		{ Token.OrAssign, "|=" },
		{ Token.LeftShiftAssign, "<<=" },
		{ Token.RightShiftAssign, ">>=" },

		{ Token.Equal, "==" },
		{ Token.NotEqual, "!=" },
		{ Token.LessEqual, "<=" },
		{ Token.GreaterEqual, ">=" },
		{ Token.And, "&&" },
		{ Token.Or, "||" },
		{ Token.PlusPlus, "++" },
		{ Token.MinusMinus, "--" },

		{ Token.Comma, "," },
		{ Token.Colon, ":" },
		{ Token.Semi, ";" },
		{ Token.Dot, "." }
	};

	private static readonly Dictionary<string, Token> _string2Token = [];

	private static readonly HashSet<string> _operations = [];
}