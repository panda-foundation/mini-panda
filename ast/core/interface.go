package core

type Expression interface {
	Node
	Type() Type
	Validate(c Context, expected Type)
	IsConstant() bool
}

type Statement interface {
	Node
	Validate(c Context)
}

type DeclarationKind int

const (
	DeclarationEnum DeclarationKind = iota
	DeclarationFunction
	DeclarationStruct
	DeclarationVariable
)

type Declaration interface {
	Node
	Type() Type
	Validate(c Context)
	IsConstant() bool
	QualifiedName() string
	ResolveType(c Context)

	Kind() DeclarationKind
}

type Struct interface {
	HasMember(name string) bool
	MemberType(name string) Type
	ValidateInitializer(c Context, expressions []Expression)
}

type Enum interface {
	HasMember(name string) bool
}
