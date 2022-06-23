package core

import "fmt"

type Type interface {
	IRWriter
	fmt.Stringer
	Equal(u Type) bool
}
