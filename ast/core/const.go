package core

const (
	Global        = "global"
	FunctionEntry = "entry"
	FunctionBody  = "body"
	FunctionExit  = "exit"
	StructThis    = "this"
	ProgramEntry  = "global.main"

	// meta define
	Extern = "extern"
)

type DeclarationKind int

const (
	DeclarationEnum DeclarationKind = iota
	DeclarationFunction
	DeclarationStruct
	DeclarationVariable
)
