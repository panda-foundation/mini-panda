package ir

import "io"

type IRWriter interface {
	WriteIR(io.Writer) error
}
