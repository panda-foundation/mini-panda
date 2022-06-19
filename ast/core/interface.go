package core

type Node interface {
	GetPosition() int
	SetPosition(position int)
}

type Type interface {
	Node
	Equal(t Type) bool
	String() string
}

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
	DeclarationVariable DeclarationKind = iota
	DeclarationFunction
	DeclarationEnum
	DeclarationStruct
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

type Function interface {
	GetReturnType() Type
}

type Context interface {
	NewContext() Context
	Error(offset int, message string)

	AddObject(name string, t Type) error
	FindObject(name string) Type
	ResolveType(v Type) Type

	FindDeclaration(t *TypeName) Declaration
	FindLocalDeclaration(name string) Declaration
	FindQualifiedDeclaration(qualified string) Declaration
	IsNamespace(namespace string) bool

	GetFunction() Function
	SetFunction(f Function)
}
