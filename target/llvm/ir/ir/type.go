package ir

type Type interface {
	IRWriter
	String() string
	Equal(u Type) bool
}
