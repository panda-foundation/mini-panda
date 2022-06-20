package core

// Value is an LLVM IR value, which may be used as an operand of instructions
// and terminators.
type Value interface {
	// Type returns the type of the value.
	Type() Type
	// Ident returns the identifier associated with the value.
	Ident() string
}
