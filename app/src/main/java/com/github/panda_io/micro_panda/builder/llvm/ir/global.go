package ir

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir_types"
)

// constant
type Global struct {
	ir.GlobalIdent
	Immutable   bool
	ContentType ir.Type
	Init        constant.Constant
	Typ         *ir_types.PointerType
}

func NewGlobal(name string, contentType ir.Type) *Global {
	global := &Global{ContentType: contentType}
	global.SetName(name)
	global.Type()
	return global
}

func NewGlobalDef(name string, init constant.Constant) *Global {
	global := &Global{ContentType: init.Type(), Init: init}
	global.SetName(name)
	global.Type()
	return global
}

func (g *Global) String() string {
	return fmt.Sprintf("%s %s", g.Type().String(), g.Ident())
}

func (g *Global) Type() ir.Type {
	if g.Typ == nil {
		g.Typ = ir_types.NewPointerType(g.ContentType)
	}
	return g.Typ
}

func (g *Global) WriteIR(w io.Writer) error {
	declaration := "global"
	if g.Immutable {
		declaration = "constant"
	}
	_, err := fmt.Fprintf(w, "%s = %s %s", g.Ident(), declaration, g.ContentType.String())
	if err != nil {
		return err
	}
	if g.Init != nil {
		_, err = fmt.Fprintf(w, " %s", g.Init.Ident())
	}
	return err
}

func (g *Global) IsConstant() {}
