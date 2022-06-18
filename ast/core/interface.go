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

type Declaration interface {
	Node
	Type() Type
	Validate(c Context)
	IsConstant() bool
	QualifiedName() string
	ResolveType(c Context)
}
