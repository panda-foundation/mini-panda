package core

import "io"

type IRWriter interface {
	WriteIR(io.Writer) error
}
