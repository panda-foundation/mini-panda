package ir

import "io"

type irWriter interface {
	writeIR(io.Writer) error
}

// === [ constant.Constant ] ===================================================

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Global) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Func) isConstant() {}

// --- [ Index ] --------------------------------

func (*Index) isConstant() {}
