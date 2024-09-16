namespace MicroPanda.Token;

using System.Collections.Generic;

internal enum Token
{
	ILLEGAL,
	EOF,
	METADATA,

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
	Import,
	Namespace,
	Pointer,
	Public,
	Return,
	Sizeof,
	Struct,
	Switch,
	This,
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
	operatorBegin,
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
	Rem,
	BitAnd,
	BitOr,
	BitXor,
	Complement,
	Not,
	Less,
	Greater,

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
	LeftShift,
	RightShift,
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

internal class TokenUtil
{
	internal static bool IsLiteral(Token token) => Token._literalBegin < token && token < Token._literalEnd;

	internal static bool IsOperator(Token token) => Token.operatorBegin < token && token < Token._operatorEnd;

	internal static bool IsKeyword(Token token) => Token._keywordBegin < token && token < Token._keywordEnd;

	internal static bool IsScalar(Token token) => Token._scalarBegin < token && token < Token._scalarEnd;

	internal static bool IsInteger(Token token) => Token._integerBegin < token && token < Token._integerEnd;

	internal static bool IsFloat(Token token) => Token._floatBegin < token && token < Token._floatEnd;

	internal static bool IsNumber(Token token) => Token._numberBegin < token && token < Token._numberEnd;

	internal static bool IsAssign(Token token) => Token._assignBegin < token && token < Token._assignEnd;

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

	internal static Dictionary<Token, string> _token2String = new()
	{
		{ Token.ILLEGAL, "ILLEGAL" },
		{ Token.EOF, "EOF" },
		{ Token.METADATA, "METADATA" },

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
		{ Token.Import, "import" },
		{ Token.Namespace, "namespace" },
		{ Token.Pointer, "pointer" },
		{ Token.Public, "public" },
		{ Token.Return, "return" },
		{ Token.Sizeof, "sizeof" },
		{ Token.Struct, "struct" },
		{ Token.Switch, "switch" },
		{ Token.This, "this" },
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
		{ Token.Rem, "%" },
		{ Token.BitAnd, "&" },
		{ Token.BitOr, "|" },
		{ Token.BitXor, "^" },
		{ Token.Complement, "~" },
		{ Token.Not, "!" },
		{ Token.Less, "<" },
		{ Token.Greater, ">" },

		{ Token.Assign, "=" },
		{ Token.PlusAssign, "+=" },
		{ Token.MinusAssign, "-=" },
		{ Token.MulAssign, "*=" },
		{ Token.DivAssign, "/=" },
		{ Token.RemAssign, "%=" },
		{ Token.XorAssign, "^=" },
		{ Token.AndAssign, "&=" },
		{ Token.OrAssign, "|=" },
		{ Token.LeftShift, "<<" },
		{ Token.RightShift, ">>" },
		{ Token.LeftShiftAssign, "<<|" },
		{ Token.RightShiftAssign, ">>|" },

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

	internal static Dictionary<string, Token> _string2Token = new()
	{
		{ "ILLEGAL", Token.ILLEGAL },
		{ "EOF", Token.EOF },
		{ "METADATA", Token.METADATA },

		{ "identifier", Token.IDENT },
		{ "bool_literal", Token.BOOL },
		{ "char_literal", Token.CHAR },
		{ "int_literal", Token.INT },
		{ "float_literal", Token.FLOAT },
		{ "string_literal", Token.STRING },
		{ "null", Token.NULL },

		{ "break", Token.Break },
		{ "case", Token.Case },
		{ "const", Token.Const },
		{ "continue", Token.Continue },
		{ "default", Token.Default },
		{ "else", Token.Else },
		{ "enum", Token.Enum },
		{ "for", Token.For },
		{ "function", Token.Function },
		{ "if", Token.If },
		{ "import", Token.Import },
		{ "namespace", Token.Namespace },
		{ "pointer", Token.Pointer },
		{ "public", Token.Public },
		{ "return", Token.Return },
		{ "sizeof", Token.Sizeof },
		{ "struct", Token.Struct },
		{ "switch", Token.Switch },
		{ "this", Token.This },
		{ "var", Token.Var },

		{ "bool", Token.Bool },
		{ "i8", Token.Int8 },
		{ "i16", Token.Int16 },
		{ "i32", Token.Int32 },
		{ "i64", Token.Int64 },
		{ "u8", Token.Uint8 },
		{ "u16", Token.Uint16 },
		{ "u32", Token.Uint32 },
		{ "u64", Token.Uint64 },
		{ "f16", Token.Float16 },
		{ "f32", Token.Float32 },
		{ "f64", Token.Float64 },
		{ "void", Token.Void },

		{ "(", Token.LeftParen },
		{ ")", Token.RightParen },
		{ "[", Token.LeftBracket },
		{ "]", Token.RightBracket },
		{ "{", Token.LeftBrace },
		{ "}", Token.RightBrace },

		{ "+", Token.Plus },
		{ "-", Token.Minus },
		{ "*", Token.Mul },
		{ "/", Token.Div },
		{ "%", Token.Rem },
		{ "&", Token.BitAnd },
		{ "|", Token.BitOr },
		{ "^", Token.BitXor },
		{ "~", Token.Complement },
		{ "!", Token.Not },
		{ "<", Token.Less },
		{ ">", Token.Greater },

		{ "=", Token.Assign },
		{ "+=", Token.PlusAssign },
		{ "-=", Token.MinusAssign },
		{ "*=", Token.MulAssign },
		{ "/=", Token.DivAssign },
		{ "%=", Token.RemAssign },
		{ "^=", Token.XorAssign },
		{ "&=", Token.AndAssign },
		{ "|=", Token.OrAssign },
		{ "<<", Token.LeftShift },
		{ ">>", Token.RightShift },
		{ "<<|", Token.LeftShiftAssign },
		{ ">>|", Token.RightShiftAssign },

		{ "==", Token.Equal },
		{ "!=", Token.NotEqual },
		{ "<=", Token.LessEqual },
		{ ">=", Token.GreaterEqual },
		{ "&&", Token.And },
		{ "||", Token.Or },
		{ "++", Token.PlusPlus },
		{ "--", Token.MinusMinus },

		{ ",", Token.Comma },
		{ ":", Token.Colon },
		{ ";", Token.Semi },
		{ ".", Token.Dot }
	};
}