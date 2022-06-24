package ir

type Value interface {
	String() string
	Type() Type
	Ident() string
}
