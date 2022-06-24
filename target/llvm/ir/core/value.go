package core

import "fmt"

type Value interface {
	fmt.Stringer
	Type() Type
	Ident() string
}
