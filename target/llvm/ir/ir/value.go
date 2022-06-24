package ir

import "fmt"

type Value interface {
	fmt.Stringer
	Type() Type
	Ident() string
}
