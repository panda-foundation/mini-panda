package token

type Token int

const (
	ILLEGAL Token = iota
	EOF
	META

	// literals
	literalBegin
	IDENT
	BOOL
	CHAR
	INT
	FLOAT
	STRING
	NULL
	literalEnd

	// keywords
	keywordBegin
	Break
	Case
	Const
	Continue
	Default
	Else
	Enum
	For
	Function
	If
	Import
	Namespace
	Pointer
	Public
	Return
	Sizeof
	Struct
	Switch
	This
	Var
	keywordEnd

	// scalars
	scalarBegin
	Bool
	numberBegin
	integerBegin
	Int8
	Int16
	Int32
	Int64
	Uint8
	Uint16
	Uint32
	Uint64
	integerEnd
	floatBegin
	Float16
	Float32
	Float64
	floatEnd
	numberEnd
	Void
	scalarEnd

	// operators
	operatorBegin
	LeftParen
	RightParen
	LeftBracket
	RightBracket
	LeftBrace
	RightBrace

	Plus
	Minus
	Mul
	Div
	Rem
	BitAnd
	BitOr
	BitXor
	Complement
	Not
	Less
	Greater

	assignBegin
	Assign
	PlusAssign
	MinusAssign
	MulAssign
	DivAssign
	RemAssign
	XorAssign
	AndAssign
	OrAssign
	LeftShift
	RightShift
	LeftShiftAssign
	RightShiftAssign
	assignEnd

	Equal
	NotEqual
	LessEqual
	GreaterEqual
	And
	Or
	PlusPlus
	MinusMinus

	Comma
	Colon
	Semi
	Dot
	operatorEnd
)

var (
	tokenStrings = [...]string{
		IDENT:  "identifier",
		BOOL:   "bool_literal",
		CHAR:   "char_literal",
		INT:    "int_literal",
		FLOAT:  "float_literal",
		STRING: "string_literal",
		NULL:   "null",

		Break:     "break",
		Case:      "case",
		Const:     "const",
		Continue:  "continue",
		Default:   "default",
		Else:      "else",
		Enum:      "enum",
		For:       "for",
		Function:  "function",
		If:        "if",
		Import:    "import",
		Namespace: "namespace",
		Pointer:   "pointer",
		Public:    "public",
		Return:    "return",
		Sizeof:    "sizeof",
		Struct:    "struct",
		Switch:    "switch",
		This:      "this",
		Var:       "var",

		Bool:    "bool",
		Int8:    "i8",
		Int16:   "i16",
		Int32:   "i32",
		Int64:   "i64",
		Uint8:   "u8",
		Uint16:  "u16",
		Uint32:  "u32",
		Uint64:  "u64",
		Float16: "f16",
		Float32: "f32",
		Float64: "f64",
		Void:    "void",

		LeftParen:    "(",
		RightParen:   ")",
		LeftBracket:  "[",
		RightBracket: "]",
		LeftBrace:    "{",
		RightBrace:   "}",

		Plus:             "+",
		Minus:            "-",
		Mul:              "*",
		Div:              "/",
		Rem:              "%",
		BitAnd:           "&",
		BitOr:            "|",
		BitXor:           "^",
		Complement:       "~",
		Not:              "!",
		Assign:           "=",
		Less:             "<",
		Greater:          ">",
		PlusAssign:       "+=",
		MinusAssign:      "-=",
		MulAssign:        "*=",
		DivAssign:        "/=",
		RemAssign:        "%=",
		XorAssign:        "^=",
		AndAssign:        "&=",
		OrAssign:         "|=",
		LeftShift:        "<<",
		RightShift:       ">>",
		LeftShiftAssign:  "<<=",
		RightShiftAssign: ">>=",
		Equal:            "==",
		NotEqual:         "!=",
		LessEqual:        "<=",
		GreaterEqual:     ">=",
		And:              "&&",
		Or:               "||",
		PlusPlus:         "++",
		MinusMinus:       "--",
		Comma:            ",",
		Colon:            ":",
		Semi:             ";",
		Dot:              ".",
	}

	tokens map[string]Token
)

func init() {
	tokens = make(map[string]Token)

	for i := keywordBegin + 1; i < keywordEnd; i++ {
		tokens[tokenStrings[i]] = i
	}

	for i := scalarBegin + 1; i < scalarEnd; i++ {
		tokens[tokenStrings[i]] = i
	}

	for i := operatorBegin + 1; i < operatorEnd; i++ {
		tokens[tokenStrings[i]] = i
	}

	operatorRoot = &operatorNode{
		children: make(map[byte]*operatorNode),
		token:    ILLEGAL,
	}

	for i := operatorBegin + 1; i < operatorEnd; i++ {
		operatorRoot.insert(tokenStrings[i])
	}
}

func ReadToken(literal string) Token {
	if token, ok := tokens[literal]; ok {
		return token
	}
	if literal == "true" || literal == "false" {
		return BOOL
	}
	if literal == "null" {
		return NULL
	}
	return IDENT
}

func (t Token) String() string {
	if 0 <= t && t < Token(len(tokenStrings)) {
		return tokenStrings[t]
	}
	return ""
}

func (t Token) IsLiteral() bool {
	return literalBegin < t && t < literalEnd
}

func (t Token) IsOperator() bool {
	return operatorBegin < t && t < operatorEnd
}

func (t Token) IsKeyword() bool {
	return keywordBegin < t && t < keywordEnd
}

func (t Token) IsScalar() bool {
	return scalarBegin < t && t < scalarEnd
}

func (t Token) IsInteger() bool {
	return integerBegin < t && t < integerEnd
}

func (t Token) IsFloat() bool {
	return floatBegin < t && t < floatEnd
}

func (t Token) IsNumber() bool {
	return numberBegin < t && t < numberEnd
}

func (t Token) IsAssign() bool {
	return assignBegin < t && t < assignEnd
}

func (t Token) Precedence() int {
	switch t {
	case Assign, MulAssign, DivAssign, RemAssign, PlusAssign, MinusAssign,
		LeftShiftAssign, RightShiftAssign, AndAssign, OrAssign, XorAssign:
		return 1

	case Or:
		return 2

	case And:
		return 3

	case BitOr:
		return 4

	case BitXor:
		return 5

	case BitAnd:
		return 6

	case Equal, NotEqual:
		return 7

	case Less, LessEqual, Greater, GreaterEqual:
		return 8

	case LeftShift, RightShift:
		return 9

	case Plus, Minus:
		return 10

	case Mul, Div, Rem:
		return 11
	}
	return 0
}
